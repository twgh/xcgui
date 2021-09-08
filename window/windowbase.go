package window

import (
	"github.com/twgh/xcgui/xc"
)

type windowBase struct {
	HWindow int
	HWND    int
}

// 给本类的HWindow和HWND赋值
func (w *windowBase) SetHWindow(hWindow int) {
	w.HWindow = hWindow
	w.HWND = xc.XWnd_GetHWND(hWindow)
}

// 炫彩_消息框
// pText: 内容文本
// pCaption: 标题
// nFlags: 标识, MessageBox_Flag_
func (w *windowBase) MessageBox(pText string, pCaption string, nFlags int) int {
	return xc.XC_MessageBox(w.HWindow, pText, pCaption, nFlags)
}

// 炫彩_发送窗口消息
// msg:
// wParam:
// lParam:
func (w *windowBase) SendMessage(msg int, wParam int, lParam int) int {
	return xc.XC_SendMessage(w.HWindow, msg, wParam, lParam)
}

// 炫彩_投递窗口消息
// msg:
// wParam:
// lParam:
func (w *windowBase) PostMessage(msg int, wParam int, lParam int) bool {
	return xc.XC_PostMessage(w.HWindow, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄
func (w *windowBase) IsHWINDOW() bool {
	return xc.XC_IsHWINDOW(w.HWindow)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象
// nID: ID值.
func (w *windowBase) GetObjectByID(nID int) int {
	return xc.XC_GetObjectByID(w.HWindow, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄
// pName: ID名称.
func (w *windowBase) GetObjectByIDName(pName string) int {
	return xc.XC_GetObjectByIDName(w.HWindow, pName)
}
