package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"sync"
	"syscall"
)

var EventHandler = newWndEventHandler()

type wndEventHandler struct {
	sync.RWMutex
	EventInfoMap map[int]map[xcc.WM_]eventInfo // 窗口事件信息map
}

type eventInfo struct {
	Cbs          []interface{}
	EvnetFuncPtr uintptr
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
// fun: 回调函数.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 默认为 false.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
func (h *wndEventHandler) AddCallBack(hWindow int, eventType xcc.WM_, eventFunc interface{}, fun interface{}, allowAddingMultiple ...bool) int {
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
			if !xc.XWnd_RegEventC1Ex(hWindow, eventType, cbPtr) {
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
		info.Cbs = append(info.Cbs, fun)
		idnex = len(info.Cbs) - 1
	} else {
		info.Cbs = []interface{}{fun}
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

// RemoveAllCallBack 从 map 里移除指定窗口的所有事件.
func (h *wndEventHandler) RemoveAllCallBack(hWindow int) {
	h.Lock()
	delete(h.EventInfoMap, hWindow)
	h.Unlock()
}

// RemoveCallBack 从 map 里移除指定窗口指定事件的指定索引的 CallBack.
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

// RemoveEvent 从 map 里移除指定窗口的指定事件.
func (h *wndEventHandler) RemoveEvent(hWindow int, eventType xcc.WM_) {
	h.Lock()
	defer h.Unlock()
	delete(h.EventInfoMap[hWindow], eventType)
}

// regWndNCDestroy 注册窗口非客户区销毁事件. 每个窗口只注册一次.
func (h *wndEventHandler) regWndNCDestroy(hWindow int) bool {
	if xc.XC_GetProperty(hWindow, "IsRegWndNCDestroy") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(onWM_NCDESTROY)
	if xc.XWnd_RegEventC1Ex(hWindow, xcc.WM_NCDESTROY, cbPtr) {
		xc.XC_SetProperty(hWindow, "IsRegWndNCDestroy", "1")
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
