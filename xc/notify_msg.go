package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 通知消息_弹出, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
func XNotifyMsg_Popup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_Popup.Call(uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_弹出扩展, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func XNotifyMsg_PopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_PopupEx.Call(uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin), common.BoolPtr(bBtnClose), common.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// hWindow: 窗口句柄.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
func XNotifyMsg_WindowPopup(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_WindowPopup.Call(uintptr(hWindow), uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// hWindow: 窗口句柄.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func XNotifyMsg_WindowPopupEx(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_WindowPopupEx.Call(uintptr(hWindow), uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin), common.BoolPtr(bBtnClose), common.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_置持续时间.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// duration: 持续时间.
func XNotifyMsg_SetDuration(hWindow, duration int) int {
	r, _, _ := xNotifyMsg_SetDuration.Call(uintptr(hWindow), uintptr(duration))
	return int(r)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func XNotifyMsg_SetParentMargin(hWindow, left, top, right, bottom int) int {
	r, _, _ := xNotifyMsg_SetParentMargin.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 通知消息_置标题高度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nHeight: 高度.
func XNotifyMsg_SetCaptionHeight(hWindow, nHeight int) int {
	r, _, _ := xNotifyMsg_SetCaptionHeight.Call(uintptr(hWindow), uintptr(nHeight))
	return int(r)
}

// 通知消息_置宽度, 设置默认宽度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nWidth: 宽度.
func XNotifyMsg_SetWidth(hWindow, nWidth int) int {
	r, _, _ := xNotifyMsg_SetWidth.Call(uintptr(hWindow), uintptr(nWidth))
	return int(r)
}

// 通知消息_置间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nSpace: 间隔大小.
func XNotifyMsg_SetSpace(hWindow, nSpace int) int {
	r, _, _ := xNotifyMsg_SetSpace.Call(uintptr(hWindow), uintptr(nSpace))
	return int(r)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func XNotifyMsg_SetBorderSize(hWindow, left, top, right, bottom int) int {
	r, _, _ := xNotifyMsg_SetBorderSize.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}
