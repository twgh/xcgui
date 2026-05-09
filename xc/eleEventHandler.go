package xc

import (
	"sync"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

// EleEventBus 元素事件总线
var EleEventBus = newElementEventBus()

// 元素事件总线
type elementEventBus struct {
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
	nextID       int // 用于生成唯一的 ID
}

// 创建新的元素事件总线
func newElementEventBus() *elementEventBus {
	return &elementEventBus{
		EventInfoMap: make(map[int]map[xcc.XE_]EventInfo),
	}
}

// AddCallback 添加回调函数, 返回回调函数 ID, 失败返回 -1.
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
func (e *elementEventBus) AddCallback(hEle int, eventType xcc.XE_, eventFunc interface{}, cb interface{}, allowAddingMultiple ...bool) int {
	e.mu.Lock()
	defer e.mu.Unlock()

	// 注册元素销毁完成事件. 在这里移除元素的所有事件.
	if !e.regEleDestroyEnd(hEle) {
		return -1
	}

	// 获取元素的事件回调函数map
	eventMap := e.EventInfoMap[hEle]
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

	// 生成唯一的 ID
	id := info.nextID
	info.nextID++
	newCbInfo := CbInfo{ID: id, CB: cb}

	if isAddingMultiple {
		info.Cbs = append(info.Cbs, newCbInfo)
	} else {
		info.Cbs = []CbInfo{newCbInfo}
	}
	eventMap[eventType] = info
	e.EventInfoMap[hEle] = eventMap
	return id
}

// GetCallbacks 获取指定元素指定事件的回调函数数组.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
func (e *elementEventBus) GetCallbacks(hEle int, eventType xcc.XE_) []CbInfo {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.EventInfoMap[hEle][eventType].Cbs
}

// RemoveAllCallback 移除指定元素的所有事件的 Callback.
//
// hEle: 元素句柄.
func (e *elementEventBus) RemoveAllCallback(hEle int) {
	e.mu.Lock()
	delete(e.EventInfoMap, hEle)
	e.mu.Unlock()
}

// RemoveCallbacks 移除指定元素指定事件的所有 Callback.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
func (e *elementEventBus) RemoveCallbacks(hEle int, eventType xcc.XE_) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.EventInfoMap[hEle], eventType)
}

// RemoveCallback 移除指定元素指定事件的指定 ID 的 Callback.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
//
// id: 回调函数 ID.
func (e *elementEventBus) RemoveCallback(hEle int, eventType xcc.XE_, id int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	eInfo := e.EventInfoMap[hEle][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs = append(eInfo.Cbs[:i], eInfo.Cbs[i+1:]...)
			e.EventInfoMap[hEle][eventType] = eInfo
		}
	}
}

// SetCallback 设置指定元素指定事件的指定 ID 的 Callback.
//
// hEle: 元素句柄.
//
// eventType: 事件类型, xcc.XE_.
//
// id: 回调函数 ID.
//
// cb: 回调函数.
func (e *elementEventBus) SetCallback(hEle int, eventType xcc.XE_, id int, cb interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	eInfo := e.EventInfoMap[hEle][eventType]
	for i, cbinfo := range eInfo.Cbs {
		if id == cbinfo.ID {
			eInfo.Cbs[i].CB = cb
			e.EventInfoMap[hEle][eventType] = eInfo
		}
	}
}

// regEleDestroyEnd 注册元素销毁完成事件. 每个元素只注册一次.
func (e *elementEventBus) regEleDestroyEnd(hEle int) bool {
	if XC_GetProperty(hEle, "IsRegEleDestroyEnd") == "1" {
		return true
	}
	cbPtr := syscall.NewCallback(OnXE_DESTROY_END)
	if XEle_RegEventC1Ex(hEle, xcc.XE_DESTROY_END, cbPtr) {
		XC_SetProperty(hEle, "IsRegEleDestroyEnd", "1")
		m := e.EventInfoMap[hEle]
		if m == nil {
			m = make(map[xcc.XE_]EventInfo)
		}
		m[xcc.XE_DESTROY_END] = EventInfo{EvnetFuncPtr: cbPtr}
		e.EventInfoMap[hEle] = m
		return true
	}
	return false
}

// ![仅供内部使用]
//
// OnXE_DESTROY_END 元素销毁完成事件. 在销毁子对象之后触发.
func OnXE_DESTROY_END(hEle int, pbHandled *bool) int {
	cbs := EleEventBus.GetCallbacks(hEle, xcc.XE_DESTROY_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].CB.(XE_DESTROY_END1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}

	EleEventBus.RemoveAllCallback(hEle)
	return ret
}
