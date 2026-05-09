package xc

import (
	"sync"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

// EleEventBus 元素事件总线
var EleEventBus = newEleEventBus()

// 元素事件总线
type eleEventBus struct {
	EventInfoMap map[int]map[xcc.XE_]EventInfo // 事件信息map
	mu           sync.RWMutex
}

// CbInfo 回调函数信息
type CbInfo struct {
	ID int         // 身份标识
	CB interface{} // 实际的回调函数
}

// EventInfo 事件信息
type EventInfo struct {
	Cbs          []CbInfo
	EvnetFuncPtr uintptr
	nextID       int // 用于生成身份标识
}

// 创建新的元素事件总线
func newEleEventBus() *eleEventBus {
	return &eleEventBus{
		EventInfoMap: make(map[int]map[xcc.XE_]EventInfo),
	}
}

// AddCallBack 添加回调函数, 返回回调函数 ID, 失败返回 -1.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
//
// eventFunc: 事件函数.
//
// cb: 回调函数. 不能为 nil.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 默认为 false.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
func (h *eleEventBus) AddCallBack(hEle int, eventType xcc.XE_, eventFunc interface{}, cb interface{}, allowAddingMultiple ...bool) int {
	h.mu.Lock()
	defer h.mu.Unlock()

	// 注册元素销毁完成事件. 在这里移除元素的所有事件.
	if !h.regEleDestroyEnd(hEle) {
		return -1
	}

	// 获取元素的事件回调函数map
	eventMap := h.EventInfoMap[hEle]
	if eventMap == nil {
		eventMap = make(map[xcc.XE_]EventInfo)
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

	// 生成一个身份标识
	id := info.nextID
	info.nextID++
	newCbInfo := CbInfo{ID: id, CB: cb}

	if isAddingMultiple {
		info.Cbs = append(info.Cbs, newCbInfo)
	} else {
		info.Cbs = []CbInfo{newCbInfo}
	}
	eventMap[eventType] = info
	h.EventInfoMap[hEle] = eventMap
	return id
}

// GetCallBacks 获取指定元素指定事件的回调函数数组.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
func (h *eleEventBus) GetCallBacks(hEle int, eventType xcc.XE_) []CbInfo {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.EventInfoMap[hEle][eventType].Cbs
}

// RemoveAllCallBack 移除指定元素的所有事件以及 CallBack.
//
// hEle: 元素句柄.
func (h *eleEventBus) RemoveAllCallBack(hEle int) {
	h.mu.Lock()
	delete(h.EventInfoMap, hEle)
	h.mu.Unlock()
}

// RemoveCallBack 移除指定元素指定事件的指定 ID 的 CallBack.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
//
// id: 回调函数 ID.
func (h *eleEventBus) RemoveCallBack(hEle int, eventType xcc.XE_, id int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	eInfo := h.EventInfoMap[hEle][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs = append(eInfo.Cbs[:i], eInfo.Cbs[i+1:]...)
			h.EventInfoMap[hEle][eventType] = eInfo
		}
	}
}

// SetCallBack 设置指定元素指定事件的指定 ID 的回调函数.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
//
// id: 回调函数 ID.
//
// cb: 回调函数.
func (h *eleEventBus) SetCallBack(hEle int, eventType xcc.XE_, id int, cb interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	eInfo := h.EventInfoMap[hEle][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs[i].CB = cb
			h.EventInfoMap[hEle][eventType] = eInfo
		}
	}
}

// RemoveEvent 移除指定元素指定事件的所有 CallBack.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
func (h *eleEventBus) RemoveEvent(hEle int, eventType xcc.XE_) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.EventInfoMap[hEle], eventType)
}

// regEleDestroyEnd 注册元素销毁完成事件. 每个元素只注册一次.
func (h *eleEventBus) regEleDestroyEnd(hEle int) bool {
	if XC_GetProperty(hEle, "IsRegEleDestroyEnd") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(OnXE_DESTROY_END)
	if XEle_RegEventC1Ex(hEle, xcc.XE_DESTROY_END, cbPtr) {
		XC_SetProperty(hEle, "IsRegEleDestroyEnd", "1")
		m := h.EventInfoMap[hEle]
		if m == nil {
			m = make(map[xcc.XE_]EventInfo)
		}
		m[xcc.XE_DESTROY_END] = EventInfo{EvnetFuncPtr: cbPtr}
		h.EventInfoMap[hEle] = m
		return true
	}
	return false
}

// ![仅供内部使用]
//
// OnXE_DESTROY_END 元素销毁完成事件. 在销毁子对象之后触发.
func OnXE_DESTROY_END(hEle int, pbHandled *bool) int {
	cbs := EleEventBus.GetCallBacks(hEle, xcc.XE_DESTROY_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(XE_DESTROY_END1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}

	EleEventBus.RemoveAllCallBack(hEle)
	return ret
}
