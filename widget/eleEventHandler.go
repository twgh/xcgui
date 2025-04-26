package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"sync"
	"syscall"
)

var EventHandler = newEleEventHandler()

type eleEventHandler struct {
	sync.RWMutex
	EventInfoMap map[int]map[xcc.XE_]eventInfo // 元素事件信息map
}

type eventInfo struct {
	Cbs          []interface{}
	EvnetFuncPtr uintptr
}

func newEleEventHandler() *eleEventHandler {
	return &eleEventHandler{
		EventInfoMap: make(map[int]map[xcc.XE_]eventInfo),
	}
}

// AddCallBack 添加回调函数, 返回回调函数索引, 失败返回 -1.
//
// hEle: 元素句柄.
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
func (h *eleEventHandler) AddCallBack(hEle int, eventType xcc.XE_, eventFunc interface{}, fun interface{}, allowAddingMultiple ...bool) int {
	h.Lock()
	defer h.Unlock()

	// 获取元素的事件回调函数map
	eventMap := h.EventInfoMap[hEle]
	if eventMap == nil {
		eventMap = make(map[xcc.XE_]eventInfo)
	}

	info, ok := eventMap[eventType]
	// 判断没有注册过这个类型的事件, 就注册
	if !ok {
		// 注册元素销毁完成事件. 在这里移除元素的所有事件.
		if !h.regEleDestroyEnd(hEle) {
			return -1
		}
		// 不是元素销毁完成事件, 才注册
		if eventType != xcc.XE_DESTROY_END {
			cbPtr := syscall.NewCallback(eventFunc)
			if !xc.XEle_RegEventC1Ex(hEle, eventType, cbPtr) {
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
	h.EventInfoMap[hEle] = eventMap
	return idnex
}

// GetCallBacks 获取指定元素指定事件的回调函数数组.
func (h *eleEventHandler) GetCallBacks(hEle int, eventType xcc.XE_) []interface{} {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[hEle][eventType].Cbs
}

// RemoveAllCallBack 从 map 里移除指定元素的所有事件.
func (h *eleEventHandler) RemoveAllCallBack(hEle int) {
	h.Lock()
	delete(h.EventInfoMap, hEle)
	h.Unlock()
}

// RemoveCallBack 从 map 里移除指定元素指定事件的指定索引的 CallBack.
func (h *eleEventHandler) RemoveCallBack(hEle int, eventType xcc.XE_, index int) bool {
	h.Lock()
	defer h.Unlock()
	info := h.EventInfoMap[hEle][eventType]
	l := len(info.Cbs)
	if l == 0 {
		return true
	}
	if index >= l {
		return false
	}
	info.Cbs = append(info.Cbs[:index], info.Cbs[index+1:]...)
	h.EventInfoMap[hEle][eventType] = info
	return true
}

// RemoveEvent 从 map 里移除指定元素的指定事件.
func (h *eleEventHandler) RemoveEvent(hWindow int, eventType xcc.XE_) {
	h.Lock()
	defer h.Unlock()
	delete(h.EventInfoMap[hWindow], eventType)
}

// regEleDestroyEnd 注册元素销毁完成事件. 每个元素只注册一次.
func (h *eleEventHandler) regEleDestroyEnd(hEle int) bool {
	if xc.XC_GetProperty(hEle, "IsRegEleDestroyEnd") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(onXE_DESTROY_END)
	if xc.XEle_RegEventC1Ex(hEle, xcc.XE_DESTROY_END, cbPtr) {
		xc.XC_SetProperty(hEle, "IsRegEleDestroyEnd", "1")
		m := h.EventInfoMap[hEle]
		if m == nil {
			m = make(map[xcc.XE_]eventInfo)
		}
		m[xcc.XE_DESTROY_END] = eventInfo{EvnetFuncPtr: cbPtr}
		h.EventInfoMap[hEle] = m
		return true
	}
	return false
}
