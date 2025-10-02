package xc

import (
	"sync"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

var WndEventHandler = newWndEventHandler()

type wndEventHandler struct {
	sync.RWMutex
	EventInfoMap map[int]map[xcc.WM_]eventInfo // 窗口事件信息map
}

func newWndEventHandler() *wndEventHandler {
	return &wndEventHandler{
		EventInfoMap: make(map[int]map[xcc.WM_]eventInfo),
	}
}

// AddCallBack 添加回调函数, 返回回调函数索引, 失败返回 -1.
//
// hWindow: 炫彩窗口句柄.
//
// eventType: 事件类型.
//
// eventFunc: 事件函数.
//
// cb: 回调函数. 不能为 nil.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 默认为 false.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
func (h *wndEventHandler) AddCallBack(hWindow int, eventType xcc.WM_, eventFunc interface{}, cb interface{}, allowAddingMultiple ...bool) int {
	h.Lock()
	defer h.Unlock()

	// 获取窗口的事件回调函数map
	eventMap := h.EventInfoMap[hWindow]
	if eventMap == nil {
		eventMap = make(map[xcc.WM_]eventInfo)
	}

	info, ok := eventMap[eventType]
	// 判断没有注册过这个类型的事件, 就注册
	if !ok {
		// 注册窗口非客户区销毁事件. 在这里移除窗口的所有事件.
		if !h.regWndNCDestroy(hWindow) {
			return -1
		}
		// 不是窗口非客户区销毁事件, 才注册
		if eventType != xcc.WM_NCDESTROY {
			cbPtr := syscall.NewCallback(eventFunc)
			if !XWnd_RegEventC1Ex(hWindow, eventType, cbPtr) {
				return -1
			}
			info.EvnetFuncPtr = cbPtr
		}
	}

	isAddingMultiple := false
	if len(allowAddingMultiple) > 0 {
		isAddingMultiple = allowAddingMultiple[0]
	}
	idnex := 0
	if isAddingMultiple {
		info.Cbs = append(info.Cbs, cb)
		idnex = len(info.Cbs) - 1
	} else {
		info.Cbs = []interface{}{cb}
	}
	eventMap[eventType] = info
	h.EventInfoMap[hWindow] = eventMap
	return idnex
}

// GetCallBacks 获取指定窗口指定事件的回调函数数组.
func (h *wndEventHandler) GetCallBacks(hWindow int, eventType xcc.WM_) []interface{} {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[hWindow][eventType].Cbs
}

// RemoveAllCallBack 从 map 里移除指定窗口的所有事件以及 CallBack.
func (h *wndEventHandler) RemoveAllCallBack(hWindow int) {
	h.Lock()
	delete(h.EventInfoMap, hWindow)
	h.Unlock()
}

// RemoveCallBack 从 map 里移除指定窗口指定事件的指定索引的 CallBack.
//   - 当只是想让一个 CallBack 失效时, 使用 SetCallBack 把该 CallBack 设置为 nil 更好, 因为你移除一个 CallBack, 会导致后面的 CallBack 的索引往前移动一位, 那你添加 CallBack 时记录的索引就没法用了.
func (h *wndEventHandler) RemoveCallBack(hWindow int, eventType xcc.WM_, index int) bool {
	h.Lock()
	defer h.Unlock()
	info := h.EventInfoMap[hWindow][eventType]
	l := len(info.Cbs)
	if l == 0 {
		return true
	}
	if index >= l {
		return false
	}
	info.Cbs = append(info.Cbs[:index], info.Cbs[index+1:]...)
	h.EventInfoMap[hWindow][eventType] = info
	return true
}

// SetCallBack 设置指定窗口指定事件的指定索引的 CallBack.
//   - 当只是想让一个 CallBack 失效时, 直接把该 CallBack 设置为 nil 比使用 RemoveCallBack 更好, 因为你移除一个 CallBack, 会导致后面的 CallBack 的索引往前移动一位, 那你添加 CallBack 时记录的索引就没法用了.
func (h *wndEventHandler) SetCallBack(hWindow int, eventType xcc.WM_, index int, cb interface{}) bool {
	h.Lock()
	defer h.Unlock()
	info := h.EventInfoMap[hWindow][eventType]
	l := len(info.Cbs)
	if l == 0 { // 空
		return false
	}
	if index >= l { // 越界
		return false
	}
	info.Cbs[index] = cb
	h.EventInfoMap[hWindow][eventType] = info
	return true
}

// RemoveEvent 从 map 里移除指定窗口的指定事件的所有CallBack.
func (h *wndEventHandler) RemoveEvent(hWindow int, eventType xcc.WM_) {
	h.Lock()
	defer h.Unlock()
	delete(h.EventInfoMap[hWindow], eventType)
}

// regWndNCDestroy 注册窗口非客户区销毁事件. 每个窗口只注册一次.
func (h *wndEventHandler) regWndNCDestroy(hWindow int) bool {
	if XC_GetProperty(hWindow, "IsRegWndNCDestroy") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(OnWM_NCDESTROY)
	if XWnd_RegEventC1Ex(hWindow, xcc.WM_NCDESTROY, cbPtr) {
		XC_SetProperty(hWindow, "IsRegWndNCDestroy", "1")
		m := h.EventInfoMap[hWindow]
		if m == nil {
			m = make(map[xcc.WM_]eventInfo)
		}
		m[xcc.WM_NCDESTROY] = eventInfo{EvnetFuncPtr: cbPtr}
		h.EventInfoMap[hWindow] = m
		return true
	}
	return false
}

// ![仅供内部使用]
//
// OnWM_NCDESTROY 窗口非客户区销毁事件.
func OnWM_NCDESTROY(hWindow int, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.WM_NCDESTROY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, pbHandled *bool) int); ok {
			ret = cb(hWindow, pbHandled)
		}
		if *pbHandled {
			break
		}
	}

	WndEventHandler.RemoveAllCallBack(hWindow)
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_POPUP 菜单弹出事件.
func OnXWM_MENU_POPUP(hWindow int, hMenu int, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_POPUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, hMenu int, pbHandled *bool) int); ok {
			ret = cb(hWindow, hMenu, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_POPUP_WND 菜单弹出窗口事件.
func OnXWM_MENU_POPUP_WND(hWindow int, hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_POPUP_WND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int); ok {
			ret = cb(hWindow, hMenu, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_SELECT 菜单选择事件.
func OnXWM_MENU_SELECT(hWindow int, nID int32, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_SELECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, nID int32, pbHandled *bool) int); ok {
			ret = cb(hWindow, nID, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_EXIT 菜单退出事件.
func OnXWM_MENU_EXIT(hWindow int, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_EXIT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, pbHandled *bool) int); ok {
			ret = cb(hWindow, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_DRAW_BACKGROUND 绘制菜单背景事件.
func OnXWM_MENU_DRAW_BACKGROUND(hWindow int, hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_DRAW_BACKGROUND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int); ok {
			ret = cb(hWindow, hDraw, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_DRAWITEM 绘制菜单项事件.
func OnXWM_MENU_DRAWITEM(hWindow int, hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int {
	cbs := WndEventHandler.GetCallBacks(hWindow, xcc.XWM_MENU_DRAWITEM)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hWindow int, hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int); ok {
			ret = cb(hWindow, hDraw, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}
