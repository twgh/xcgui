package xc

import (
	"github.com/twgh/xcgui/xcc"
	"sync"
	"syscall"
)

var (
	uiThreadCallBackFunc func(data int) int                      // 真实执行的
	uiThreadCallBackPtr  = syscall.NewCallback(uiThreadCallBack) // 壳
	rwm                  sync.RWMutex
)

// uiThreadCallBack UI线程回调函数
func uiThreadCallBack(data int) int {
	return uiThreadCallBackFunc(data)
}

// XC_CallUiThreadEx 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI. 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//
// f: 回调函数.
//
// data: 传进回调函数的用户自定义数据.
func XC_CallUiThreadEx(f func(data int) int, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = f
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}

// XC_CallUT 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI. 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数. 回调函数没有参数也没有返回值.
//
// f: 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
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

// XC_CallUiThreader 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI. 与 XC_CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//
// u: xc.UiThreader.
//
// data: 传进回调函数的用户自定义数据.
func XC_CallUiThreader(u UiThreader, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = u.UiThreadCallBack
	r, _, _ := xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}

// SetBnClicks 给一个窗口或布局元素中所有的按钮注册按钮单击事件.
//
// 说明: 本函数是通过遍历子元素实现的, 只会给窗口和布局元素中的按钮注册事件, 像List, Tree, 滑块条等元素中的按钮不会注册, 想要更多可以自己实现一个.
//
//	HXCGUI 炫彩窗口句柄或布局元素句柄.
//
//	onBnClick 按钮单击事件回调函数.
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

// Font_Info_Name 将[32]uint16转换到string.
//
// arr: [32]uint16.
func Font_Info_Name(arr [32]uint16) string {
	return syscall.UTF16ToString(arr[0:])
}

func Rect2RectF(rc RECT) RECTF {
	return RECTF{
		Left:   float32(rc.Left),
		Right:  float32(rc.Right),
		Top:    float32(rc.Top),
		Bottom: float32(rc.Bottom),
	}
}

// PointInRect 判断一个点是否在矩形范围内.
func PointInRect(pt POINT, rc RECT) bool {
	return pt.X <= rc.Right && pt.X >= rc.Left && pt.Y <= rc.Bottom && pt.Y >= rc.Top
}
