package xc

import (
	"github.com/twgh/xcgui/common"
)

// 浮动窗口_启用标题栏内容, 启用后,将创建标题(形状文本)和关闭按钮.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用, 不填默认为 true.
func XFloatWnd_EnableCaptionContent(hWindow int, bEnable ...bool) bool {
	enable := true
	if len(bEnable) > 0 {
		enable = bEnable[0]
	}
	r, _, _ := xFloatWnd_EnableCaptionContent.Call(uintptr(hWindow), common.BoolPtr(enable))
	return r != 0
}

// 浮动窗口_取标题形状文本, 启用标题栏内容后有效, 返回形状文本句柄.
//
// hWindow: 窗口句柄.
func XFloatWnd_GetCaptionShapeText(hWindow int) int {
	r, _, _ := xFloatWnd_GetCaptionShapeText.Call(uintptr(hWindow))
	return int(r)
}

// 浮动窗口_取标题关闭按钮, 启用标题栏内容后有效, 返回关闭按钮句柄.
//
// hWindow: 窗口句柄.
func XFloatWnd_GetCaptionButtonClose(hWindow int) int {
	r, _, _ := xFloatWnd_GetCaptionButtonClose.Call(uintptr(hWindow))
	return int(r)
}

// 浮动窗口_置标题, 启用标题栏内容后有效.
//
// hWindow: 窗口句柄.
//
// pTitle: 标题文本.
func XFloatWnd_SetTitle(hWindow int, pTitle string) {
	xFloatWnd_SetTitle.Call(uintptr(hWindow), common.StrPtr(pTitle))
}

// 浮动窗口_取标题, 启用标题栏内容后有效, 返回标题文本.
//
// hWindow: 窗口句柄.
func XFloatWnd_GetTitle(hWindow int) string {
	r, _, _ := xFloatWnd_GetTitle.Call(uintptr(hWindow))
	return common.UintPtrToString(r)
}
