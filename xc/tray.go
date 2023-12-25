package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 托盘图标_重置. 重置数据, 当未添加到系统托盘状态才可调用.
func XTrayIcon_Reset() {
	xTrayIcon_Reset.Call()
}

// 托盘图标_删除. 从系统托盘删除.
func XTrayIcon_Del() bool {
	r, _, _ := xTrayIcon_Del.Call()
	return r != 0
}

// 托盘图标_修改. 修改托盘图标.
func XTrayIcon_Modify() bool {
	r, _, _ := xTrayIcon_Modify.Call()
	return r != 0
}

// 托盘图标_置焦点. 将焦点设置回系统托盘.
func XTrayIcon_SetFocus() bool {
	r, _, _ := xTrayIcon_SetFocus.Call()
	return r != 0
}

// 托盘图标_添加. 将图标添加到系统托盘.
//
// hWindow: 关联窗口句柄.
//
// id: 托盘图标唯一标识符.
func XTrayIcon_Add(hWindow int, id int32) bool {
	r, _, _ := xTrayIcon_Add.Call(uintptr(hWindow), uintptr(id))
	return r != 0
}

// 托盘图标_置图标. 设置图标.
//
// hIcon: 图标句柄.
func XTrayIcon_SetIcon(hIcon uintptr) {
	xTrayIcon_SetIcon.Call(hIcon)
}

// 托盘图标_置提示文本. 设置工具提示内容.
//
// pTips: 提示文本内容, 长度不能超过127个字符.
func XTrayIcon_SetTips(pTips string) {
	xTrayIcon_SetTips.Call(common.StrPtr(pTips))
}

// 托盘图标_置回调消息. 设置用户自定义的回调消息类型, 触发事件后, 系统会发送到此消息.
//
// user_message: 用户自定义消息, 界面库默认定义消息为: xcc.XWM_TRAYICON.
func XTrayIcon_SetCallbackMessage(user_message uint32) {
	xTrayIcon_SetCallbackMessage.Call(uintptr(user_message))
}

// 托盘图标_置弹出气泡. 设置弹出气泡信息.
//
// pTitle: 弹出气泡标题.
//
// pText: 弹出气泡内容.
//
// hBalloonIcon: 气泡图标. 可填0.
//
// flags: 标识, 可设置默认图标类型, 禁用声音等: xcc.TrayIcon_Flag_
func XTrayIcon_SetPopupBalloon(pTitle, pText string, hBalloonIcon uintptr, flags xcc.TrayIcon_Flag_) {
	xTrayIcon_SetPopupBalloon.Call(common.StrPtr(pTitle), common.StrPtr(pText), hBalloonIcon, uintptr(flags))
}
