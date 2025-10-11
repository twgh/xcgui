package xc

import (
	"sync"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

var EleEventHandler = newEleEventHandler()

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
// cb: 回调函数. 不能为 nil.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 默认为 false.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
func (h *eleEventHandler) AddCallBack(hEle int, eventType xcc.XE_, eventFunc interface{}, cb interface{}, allowAddingMultiple ...bool) int {
	h.Lock()
	defer h.Unlock()

	// 注册元素销毁完成事件. 在这里移除元素的所有事件.
	if !h.regEleDestroyEnd(hEle) {
		return -1
	}

	// 获取元素的事件回调函数map
	eventMap := h.EventInfoMap[hEle]
	if eventMap == nil {
		eventMap = make(map[xcc.XE_]eventInfo)
	}

	info, ok := eventMap[eventType]
	// 判断没有注册过这个类型的事件, 就注册
	if !ok {
		// 不是元素销毁完成事件, 才注册
		if eventType != xcc.XE_DESTROY_END {
			cbPtr := syscall.NewCallback(eventFunc)
			if !XEle_RegEventC1Ex(hEle, eventType, cbPtr) {
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
	h.EventInfoMap[hEle] = eventMap
	return idnex
}

// GetCallBacks 获取指定元素指定事件的回调函数数组.
func (h *eleEventHandler) GetCallBacks(hEle int, eventType xcc.XE_) []interface{} {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[hEle][eventType].Cbs
}

// RemoveAllCallBack 从 map 里移除指定元素的所有事件以及CallBack.
func (h *eleEventHandler) RemoveAllCallBack(hEle int) {
	h.Lock()
	delete(h.EventInfoMap, hEle)
	h.Unlock()
}

// RemoveCallBack 从 map 里移除指定元素指定事件的指定索引的 CallBack.
//   - 当只是想让一个 CallBack 失效时, 使用 SetCallBack 把该 CallBack 设置为 nil 更好, 因为你移除一个 CallBack, 会导致后面的 CallBack 的索引往前移动一位, 那你添加 CallBack 时记录的索引就没法用了.
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

// SetCallBack 设置指定元素指定事件的指定索引的回调函数.
//   - 当只是想让一个 CallBack 失效时, 直接把该 CallBack 设置为 nil 比使用 RemoveCallBack 更好, 因为你移除一个 CallBack, 会导致后面的 CallBack 的索引往前移动一位, 那你添加 CallBack 时记录的索引就没法用了.
func (h *eleEventHandler) SetCallBack(hEle int, eventType xcc.XE_, index int, cb interface{}) bool {
	h.Lock()
	defer h.Unlock()
	info := h.EventInfoMap[hEle][eventType]
	l := len(info.Cbs)
	if l == 0 { // 空
		return false
	}
	if index >= l { // 越界
		return false
	}
	info.Cbs[index] = cb
	h.EventInfoMap[hEle][eventType] = info
	return true
}

// RemoveEvent 从 map 里移除指定元素的指定事件的所有CallBack.
func (h *eleEventHandler) RemoveEvent(hEle int, eventType xcc.XE_) {
	h.Lock()
	defer h.Unlock()
	delete(h.EventInfoMap[hEle], eventType)
}

// regEleDestroyEnd 注册元素销毁完成事件. 每个元素只注册一次.
func (h *eleEventHandler) regEleDestroyEnd(hEle int) bool {
	if XC_GetProperty(hEle, "IsRegEleDestroyEnd") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(OnXE_DESTROY_END)
	if XEle_RegEventC1Ex(hEle, xcc.XE_DESTROY_END, cbPtr) {
		XC_SetProperty(hEle, "IsRegEleDestroyEnd", "1")
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

// ![仅供内部使用]
//
// OnXE_DESTROY_END 元素销毁完成事件. 在销毁子对象之后触发.
func OnXE_DESTROY_END(hEle int, pbHandled *bool) int {
	cbs := EleEventHandler.GetCallBacks(hEle, xcc.XE_DESTROY_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(XE_DESTROY_END1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}

	EleEventHandler.RemoveAllCallBack(hEle)
	return ret
}
