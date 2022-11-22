package xc

import (
	"sync"
	"syscall"
)

var (
	uiThreadCallBackFunc func(data int) int
	uiThreadCallBackPtr  = syscall.NewCallback(uiThreadCallBack)
	rwm                  sync.RWMutex
)

// uiThreadCallBack UI线程回调函数
func uiThreadCallBack(data int) int {
	return uiThreadCallBackFunc(data)
}

// XC_CallUiThreadEx 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 xc.XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param f 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func XC_CallUiThreadEx(f func(data int) int, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = f
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}

// XC_CallUT 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 xc.XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数. 回调函数没有参数也没有返回值.
//	@param f 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func XC_CallUT(f func()) {
	rwm.Lock()
	uiThreadCallBackFunc = func(data int) int {
		f()
		return 0
	}
	xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(0))
	rwm.Unlock()
}

// UiThreader 用于在UI线程操作UI.
type UiThreader interface {
	UiThreadCallBack(data int) int // 回调函数, 把操作UI的代码写在这里面.
}

// XC_CallUiThreader 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 xc.XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param u xc.UiThreader.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func XC_CallUiThreader(u UiThreader, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = u.UiThreadCallBack
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}
