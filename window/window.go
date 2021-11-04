package window

import (
	"github.com/twgh/xcgui/xc"
)

// 普通窗口.
type Window struct {
	windowBase
}

// 窗口_创建.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func NewWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// 窗口_创建扩展.
//
// dwExStyle: 窗口扩展样式.
//
// lpClassName: 窗口类名.
//
// lpWindowName: 窗口名.
//
// dwStyle: 窗口样式.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func NewWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle))
	return p
}

// 窗口_附加窗口, 返回窗口对象.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func Attach(hWnd, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Attach(hWnd, XCStyle))
	return p
}

// 从句柄创建对象.
func NewWindowByHandle(handle int) *Window {
	p := &Window{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewWindowByName(name string) *Window {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewWindowByUID(nUID int) *Window {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewWindowByUIDName(name string) *Window {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

/*
LayoutBox-布局盒子
*/

// 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (w *Window) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(w.Handle, bEnable)
}

// 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (w *Window) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(w.Handle, bEnable)
}

// 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (w *Window) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(w.Handle, bEnable)
}

// 布局盒子_置水平对齐.
//
// nAlign: 对齐方式.
func (w *Window) SetAlignH(nAlign int) int {
	return xc.XLayoutBox_SetAlignH(w.Handle, nAlign)
}

// 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式.
func (w *Window) SetAlignV(nAlign int) int {
	return xc.XLayoutBox_SetAlignV(w.Handle, nAlign)
}

// 布局盒子_置对齐基线.
//
// nAlign: 对齐方式.
func (w *Window) SetAlignBaseline(nAlign int) int {
	return xc.XLayoutBox_SetAlignBaseline(w.Handle, nAlign)
}

// 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (w *Window) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(w.Handle, nSpace)
}

// 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (w *Window) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(w.Handle, nSpace)
}
