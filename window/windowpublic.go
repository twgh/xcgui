package window

import (
	xc "github.com/twgh/xcgui"
)

type WindowPublic struct {
	HWindow_ int
}

// 炫彩_消息框
// pText: 内容文本
// pCaption: 标题
// nFlags: 标识, MessageBox_Flag_
func (w *WindowPublic) MessageBox(pText string, pCaption string, nFlags int) int {
	return xc.XC_MessageBox(w.HWindow_, pText, pCaption, nFlags)
}

// 炫彩_发送窗口消息
// msg:
// wParam:
// lParam:
func (w *WindowPublic) SendMessage(msg int, wParam int, lParam int) int {
	return xc.XC_SendMessage(w.HWindow_, msg, wParam, lParam)
}

// 炫彩_投递窗口消息
// msg:
// wParam:
// lParam:
func (w *WindowPublic) PostMessage(msg int, wParam int, lParam int) bool {
	return xc.XC_PostMessage(w.HWindow_, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄
func (w *WindowPublic) IsHWINDOW() bool {
	return xc.XC_IsHWINDOW(w.HWindow_)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象
// nID: ID值.
func (w *WindowPublic) GetObjectByID(nID int) int {
	return xc.XC_GetObjectByID(w.HWindow_, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄
// pName: ID名称.
func (w *WindowPublic) GetObjectByIDName(pName string) int {
	return xc.XC_GetObjectByIDName(w.HWindow_, pName)
}
