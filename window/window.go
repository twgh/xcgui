package window

import (
	"github.com/twgh/xcgui/xc"
)

// 普通窗口
type Window struct {
	windowBase
}

// 窗口_创建
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// pTitle: 窗口标题.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式: Xc_Window_Style_
func NewWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XFrameWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// 窗口_创建扩展
// dwExStyle: 窗口扩展样式.
// lpClassName: 窗口类名.
// lpWindowName: 窗口名.
// dwStyle: 窗口样式
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式: Xc_Window_Style_
func NewWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle))
	return p
}
