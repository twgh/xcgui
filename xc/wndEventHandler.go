package xc

import (
	"sync"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

// WndEventBus 窗口事件总线
var WndEventBus = newWindowEventBus()

// 窗口事件总线
type windowEventBus struct {
	EventInfoMap map[int]map[xcc.WM_]EventInfo // 事件信息map
	mu           sync.RWMutex
}

// 创建新的窗口事件总线
func newWindowEventBus() *windowEventBus {
	return &windowEventBus{
		EventInfoMap: make(map[int]map[xcc.WM_]EventInfo),
	}
}

// AddCallback 添加回调函数, 返回回调函数 ID, 失败返回 -1.
//
// hWindow: 炫彩窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
//
// eventFunc: 事件函数.
//
// cb: 回调函数. 不能为 nil.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 不填默认为 true.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
func (w *windowEventBus) AddCallback(hWindow int, eventType xcc.WM_, eventFunc interface{}, cb interface{}, allowAddingMultiple ...bool) int {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 注册窗口非客户区销毁事件. 在这里移除窗口的所有事件.
	if !w.regWndNCDestroy(hWindow) {
		return -1
	}

	// 获取窗口的事件回调函数map
	eventMap := w.EventInfoMap[hWindow]
	if eventMap == nil {
		eventMap = make(map[xcc.WM_]EventInfo)
	}

	info, ok := eventMap[eventType]
	// 判断没有注册过这个类型的事件, 就注册
	if !ok {
		// 不是窗口非客户区销毁事件, 才注册
		if eventType != xcc.WM_NCDESTROY {
			cbPtr := syscall.NewCallback(eventFunc)
			if !XWnd_RegEventC1Ex(hWindow, eventType, cbPtr) {
				return -1
			}
			info.EvnetFuncPtr = cbPtr
		}
	}

	isAddingMultiple := true
	if len(allowAddingMultiple) > 0 {
		isAddingMultiple = allowAddingMultiple[0]
	}

	newCbInfo := CbInfo{ID: info.nextID, CB: cb}

	if isAddingMultiple {
		info.nextID++
		info.Cbs = append(info.Cbs, newCbInfo)
	} else {
		info.nextID = 0 // 因为覆盖了旧的回调函数, 所以重置为0
		newCbInfo.ID = 0
		info.Cbs = []CbInfo{newCbInfo}
	}
	eventMap[eventType] = info
	w.EventInfoMap[hWindow] = eventMap
	return newCbInfo.ID
}

// GetEventFuncPtr 获取指定窗口指定事件的事件函数指针.
//
// hWindow: 窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
func (w *windowEventBus) GetEventFuncPtr(hWindow int, eventType xcc.WM_) uintptr {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.EventInfoMap[hWindow][eventType].EvnetFuncPtr
}

// GetCallbacks 获取指定窗口指定事件的回调函数数组.
//
// hWindow: 窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
func (w *windowEventBus) GetCallbacks(hWindow int, eventType xcc.WM_) []CbInfo {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.EventInfoMap[hWindow][eventType].Cbs
}

// RemoveEvent 移除指定元素的指定事件以及该事件的 Callbacks.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型, xcc.WM_, xcc.XWM_.
func (w *windowEventBus) RemoveEvent(hWindow int, nEvent xcc.WM_) *windowEventBus {
	cbPtr := WndEventBus.GetEventFuncPtr(hWindow, nEvent)
	if cbPtr > 0 {
		XWnd_RemoveEventCEx(hWindow, nEvent, cbPtr)
	}
	WndEventBus.RemoveCallbacks(hWindow, nEvent)
	return w
}

// RemoveAllCallback 移除指定窗口的所有事件的 Callback.
//
// hWindow: 窗口句柄.
func (w *windowEventBus) RemoveAllCallback(hWindow int) *windowEventBus {
	w.mu.Lock()
	delete(w.EventInfoMap, hWindow)
	w.mu.Unlock()
	return w
}

// RemoveCallbacks 移除指定窗口指定事件的所有 Callback.
//
// hWindow: 窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
func (w *windowEventBus) RemoveCallbacks(hWindow int, eventType xcc.WM_) *windowEventBus {
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.EventInfoMap[hWindow], eventType)
	return w
}

// RemoveCallback 移除指定窗口指定事件的指定 ID 的 Callback.
//
// hWindow: 窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
//
// id: 回调函数 ID.
func (w *windowEventBus) RemoveCallback(hWindow int, eventType xcc.WM_, id int) *windowEventBus {
	w.mu.Lock()
	defer w.mu.Unlock()

	eInfo := w.EventInfoMap[hWindow][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs = append(eInfo.Cbs[:i], eInfo.Cbs[i+1:]...)
			w.EventInfoMap[hWindow][eventType] = eInfo
			break
		}
	}
	return w
}

// SetCallback 设置指定窗口指定事件的指定 ID 的 Callback.
//
// hWindow: 窗口句柄.
//
// eventType: 事件类型, xcc.WM_, xcc.XWM_.
//
// id: 回调函数 ID.
//
// cb: 回调函数.
func (w *windowEventBus) SetCallback(hWindow int, eventType xcc.WM_, id int, cb interface{}) *windowEventBus {
	w.mu.Lock()
	defer w.mu.Unlock()

	eInfo := w.EventInfoMap[hWindow][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs[i].CB = cb
			w.EventInfoMap[hWindow][eventType] = eInfo
			break
		}
	}
	return w
}

// regWndNCDestroy 注册窗口非客户区销毁事件. 每个窗口只注册一次.
func (w *windowEventBus) regWndNCDestroy(hWindow int) bool {
	if XC_GetProperty(hWindow, "IsRegWndNCDestroy") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(OnWM_NCDESTROY)
	if XWnd_RegEventC1Ex(hWindow, xcc.WM_NCDESTROY, cbPtr) {
		XC_SetProperty(hWindow, "IsRegWndNCDestroy", "1")
		m := w.EventInfoMap[hWindow]
		if m == nil {
			m = make(map[xcc.WM_]EventInfo)
		}
		m[xcc.WM_NCDESTROY] = EventInfo{EvnetFuncPtr: cbPtr}
		w.EventInfoMap[hWindow] = m
		return true
	}
	return false
}

// ![仅供内部使用]
//
// OnWM_NCDESTROY 窗口非客户区销毁事件.
func OnWM_NCDESTROY(hWindow int, pbHandled *bool) int {
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.WM_NCDESTROY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(WM_NCDESTROY1); ok {
			ret = cb(hWindow, pbHandled)
		}
		if *pbHandled {
			break
		}
	}

	WndEventBus.RemoveAllCallback(hWindow)
	return ret
}

// ![仅供内部使用]
//
// OnXWM_MENU_POPUP 菜单弹出事件.
func OnXWM_MENU_POPUP(hWindow int, hMenu int, pbHandled *bool) int {
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_POPUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, hMenu int, pbHandled *bool) int); ok {
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
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_POPUP_WND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int); ok {
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
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_SELECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, nID int32, pbHandled *bool) int); ok {
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
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_EXIT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, pbHandled *bool) int); ok {
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
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_DRAW_BACKGROUND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int); ok {
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
	cbs := WndEventBus.GetCallbacks(hWindow, xcc.XWM_MENU_DRAWITEM)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(func(hWindow int, hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int); ok {
			ret = cb(hWindow, hDraw, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}
