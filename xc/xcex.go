package xc

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

var (
	uiThreadCallBackFunc func(data int) int                      // 真实执行的
	uiThreadCallBackPtr  = syscall.NewCallback(uiThreadCallBack) // 壳
	utLock               sync.Mutex
)

// uiThreadCallBack UI线程回调函数
func uiThreadCallBack(data int) int {
	return uiThreadCallBackFunc(data)
}

// XC_CallUiThreadEx 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//   - 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//
// f: 回调函数.
//
// data: 传进回调函数的用户自定义数据.
func XC_CallUiThreadEx(f func(data int) int, data int) int {
	utLock.Lock()
	defer utLock.Unlock()
	uiThreadCallBackFunc = f
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	return int(r)
}

// XC_CallUT 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//   - 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//   - 回调函数没有参数也没有返回值.
//
// f: 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func XC_CallUT(f func()) {
	utLock.Lock()
	defer utLock.Unlock()
	uiThreadCallBackFunc = func(data int) int {
		f()
		return 0
	}
	xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(0))
}

// U 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//   - 是 XC_CallUT 的简写.
//
// f: 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func U(f func()) {
	XC_CallUT(f)
}

// A 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//   - 是 CallUTAny 的简写.
//   - 可以传入任意参数.
//
// f: 回调函数.
//
// data: 传进回调函数的用户自定义数据.
func A(f func(data ...interface{}) int, data ...interface{}) int {
	return CallUTAny(f, data...)
}

// CallUTAny 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//   - 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//   - 可以传入任意参数.
//
// f: 回调函数.
//
// data: 传进回调函数的用户自定义数据.
func CallUTAny(f func(data ...interface{}) int, data ...interface{}) int {
	utLock.Lock()
	defer utLock.Unlock()

	uiThreadCallBackFunc = func(args int) int {
		s := common.UintPtrToSliceWithCap(uintptr(args))
		return f(s...)
	}

	var dataPtr uintptr
	dataLen := len(data)
	if dataLen > 0 {
		// 创建新切片, 第0个元素是切片长度, 后面是参数数据
		dataCopy := make([]interface{}, dataLen+1)
		dataCopy[0] = dataLen + 1

		for i := 1; i < dataLen+1; i++ {
			dataCopy[i] = data[i-1]
		}

		dataPtr = uintptr(unsafe.Pointer(&dataCopy[0]))
	}
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, dataPtr)
	return int(r)
}

// Auto 用于在UI线程操作UI.
//   - 会自动判断当前是否在UI线程, 如果在UI线程则直接执行, 如果不在UI线程则会调用 XC_CallUT.
func Auto(f func()) {
	if IsUiThread() {
		f()
	} else {
		XC_CallUT(f)
	}
}

// AutoInt 用于在UI线程操作UI.
//   - 会自动判断当前是否在UI线程, 如果在UI线程则直接执行, 如果不在UI线程则会调用 XC_CallUiThreadEx.
func AutoInt(f func(data int) int, data int) int {
	if IsUiThread() {
		return f(data)
	} else {
		return XC_CallUiThreadEx(f, data)
	}
}

// AutoAny 用于在UI线程操作UI.
//   - 会自动判断当前是否在UI线程, 如果在UI线程则直接执行, 如果不在UI线程则会调用 CallUTAny.
func AutoAny(f func(data ...interface{}) int, data ...interface{}) int {
	if IsUiThread() {
		return f(data...)
	} else {
		return CallUTAny(f, data...)
	}
}

// UiThreader 用于在UI线程操作UI.
type UiThreader interface {
	UiThreadCallBack(data int) int // 回调函数, 把操作UI的代码写在这里面.
}

// XC_CallUiThreader 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI. 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//
// u: xc.UiThreader.
//
// data: 传进回调函数的用户自定义数据.
func XC_CallUiThreader(u UiThreader, data int) int {
	utLock.Lock()
	defer utLock.Unlock()
	uiThreadCallBackFunc = u.UiThreadCallBack
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	return int(r)
}

// SetBnClicks 给一个窗口或布局元素中所有的按钮注册按钮单击事件.
//   - 本函数是通过遍历子元素实现的, 只会给窗口和布局元素中的按钮注册事件, 像 List, Tree, 滑块条等元素中的按钮不会注册, 想要更多可以自己实现一个.
//
// HXCGUI: 炫彩窗口句柄或布局元素句柄.
//
// onBnClick: 按钮单击事件回调函数.
func SetBnClicks(HXCGUI int, onBnClick func(hEle int, pbHandled *bool) int) {
	if XC_IsHWINDOW(HXCGUI) {
		for i := int32(0); i < XWnd_GetChildCount(HXCGUI); i++ {
			hEle := XWnd_GetChildByIndex(HXCGUI, i)
			switch XC_GetObjectType(hEle) {
			case xcc.XC_ELE_LAYOUT:
				SetBnClicks(hEle, onBnClick)
			case xcc.XC_BUTTON:
				XEle_RegEventC1(hEle, xcc.XE_BNCLICK, onBnClick)
			}
		}
	} else if XC_GetObjectType(HXCGUI) == xcc.XC_ELE_LAYOUT {
		for i := int32(0); i < XEle_GetChildCount(HXCGUI); i++ {
			hEle := XEle_GetChildByIndex(HXCGUI, i)
			switch XC_GetObjectType(hEle) {
			case xcc.XC_ELE_LAYOUT:
				SetBnClicks(hEle, onBnClick)
			case xcc.XC_BUTTON:
				XEle_RegEventC1(hEle, xcc.XE_BNCLICK, onBnClick)
			}
		}
	}
}
