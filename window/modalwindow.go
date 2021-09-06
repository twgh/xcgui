package window

import (
	xc "github.com/twgh/xcgui"
)

type ModalWindow struct {
	WindowPublic

	HWindow int
	HWND    int
}

// 模态窗口_创建, 创建模态窗口, 然后调用DoModal()显示; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄
// nWidth: 宽度.
// nHeight: 高度.
// pTitle: 窗口标题内容.
// hWndParent: 父窗口句柄.
// XCStyle: 炫彩窗口样式, Xc_Window_Style_
func NewModalWindow(nWidth int, nHeight int, pTitle string, hWndParent int, XCStyle int) *ModalWindow {
	p := &ModalWindow{
		HWindow: xc.XModalWnd_Create(nWidth, nHeight, pTitle, hWndParent, XCStyle),
	}
	p.HWindow_ = p.HWindow
	p.HWND = xc.XWnd_GetHWND(p.HWindow)
	return p
}

// 模态窗口_创建扩展, 创建模态窗口,增强功能
// dwExStyle: 窗口扩展样式.
// lpClassName: 窗口类名.
// lpWindowName: 窗口名.
// dwStyle: 窗口样式.
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式, Xc_Window_Style_
func NewModalWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *ModalWindow {
	p := &ModalWindow{
		HWindow: xc.XModalWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle),
	}
	p.HWindow_ = p.HWindow
	return p
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时
// bEnable: 开启开关.
func (m *ModalWindow) EnableAutoClose(bEnable bool) int {
	return xc.XModalWnd_EnableAutoClose(m.HWindow, bEnable)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口
// bEnable: 是否启用
func (m *ModalWindow) EnableEscClose(bEnable bool) int {
	return xc.XModalWnd_EnableEscClose(m.HWindow, bEnable)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出
func (m *ModalWindow) DoModal() int {
	return xc.XModalWnd_DoModal(m.HWindow)
}

// 模态窗口_结束, 结束模态窗口, 返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出
// nResult: XModalWnd_DoModal()返回值.
func (m *ModalWindow) EndModal(nResult int) int {
	return xc.XModalWnd_EndModal(m.HWindow, nResult)
}
