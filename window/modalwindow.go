package window

import (
	"github.com/twgh/xcgui/xc"
)

// 模态窗口.
type ModalWindow struct {
	windowBase
}

// 模态窗口_创建, 创建模态窗口, 然后你需要调用DoModal()来显示窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// pTitle: 窗口标题内容.
//
// hWndParent: 父窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func NewModalWindow(nWidth int, nHeight int, pTitle string, hWndParent int, XCStyle int) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XModalWnd_Create(nWidth, nHeight, pTitle, hWndParent, XCStyle))
	return p
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能.
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
func NewModalWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XModalWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle))
	return p
}

// 从句柄创建对象.
func NewModalWindowByHandle(handle int) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewModalWindowByName(name string) *ModalWindow {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewModalWindowByUID(nUID int) *ModalWindow {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewModalWindowByUIDName(name string) *ModalWindow {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//
// bEnable: 开启开关.
func (m *ModalWindow) EnableAutoClose(bEnable bool) int {
	return xc.XModalWnd_EnableAutoClose(m.Handle, bEnable)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//
// bEnable: 是否启用.
func (m *ModalWindow) EnableEscClose(bEnable bool) int {
	return xc.XModalWnd_EnableEscClose(m.Handle, bEnable)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
func (m *ModalWindow) DoModal() int {
	return xc.XModalWnd_DoModal(m.Handle)
}

// 模态窗口_结束, 结束模态窗口, 返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
//
// nResult: XModalWnd_DoModal()返回值.
func (m *ModalWindow) EndModal(nResult int) int {
	return xc.XModalWnd_EndModal(m.Handle, nResult)
}

// 模态窗口_附加窗口, 返回窗口资源句柄.
func (m *ModalWindow) Attach() int {
	return xc.XModalWnd_Attach(m.Handle)
}
