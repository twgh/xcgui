package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

// 形状组框_创建, 创建组框形状对象, 返回句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 名称.
//
// hParent: 父对象句柄.
func XShapeGroupBox_Create(x int, y int, cx int, cy int, pName string, hParent int) int {
	r, _, _ := xShapeGroupBox_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pName), uintptr(hParent))
	return int(r)
}

// 形状组框_置边框颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeGroupBox_SetBorderColor(hShape int, color int) int {
	r, _, _ := xShapeGroupBox_SetBorderColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状组框_置文本颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeGroupBox_SetTextColor(hShape int, color int) int {
	r, _, _ := xShapeGroupBox_SetTextColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状组框_置字体.
//
// hShape: 形状对象句柄.
//
// hFontX: 炫彩字体.
func XShapeGroupBox_SetFontX(hShape int, hFontX int) int {
	r, _, _ := xShapeGroupBox_SetFontX.Call(uintptr(hShape), uintptr(hFontX))
	return int(r)
}

// 形状组框_置文本偏移, 设置文本偏移量.
//
// hShape: 形状对象句柄.
//
// offsetX: 水平偏移.
//
// offsetY: 垂直偏移.
func XShapeGroupBox_SetTextOffset(hShape int, offsetX int32, offsetY int32) int {
	r, _, _ := xShapeGroupBox_SetTextOffset.Call(uintptr(hShape), uintptr(offsetX), uintptr(offsetY))
	return int(r)
}

// 形状组框_置圆角大小.
//
// hShape: 形状对象句柄.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XShapeGroupBox_SetRoundAngle(hShape int, nWidth int32, nHeight int32) int {
	r, _, _ := xShapeGroupBox_SetRoundAngle.Call(uintptr(hShape), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 形状组框_置文本.
//
// hShape: 形状对象句柄.
//
// pText: 文本内容.
func XShapeGroupBox_SetText(hShape int, pText string) int {
	r, _, _ := xShapeGroupBox_SetText.Call(uintptr(hShape), common.StrPtr(pText))
	return int(r)
}

// 形状组框_取文本偏移, 获取文本偏移量.
//
// hShape: 形状对象句柄.
//
// pOffsetX: X坐标偏移量.
//
// pOffsetY: Y坐标偏移量.
func XShapeGroupBox_GetTextOffset(hShape int, pOffsetX *int32, pOffsetY *int32) int {
	r, _, _ := xShapeGroupBox_GetTextOffset.Call(uintptr(hShape), uintptr(unsafe.Pointer(pOffsetX)), uintptr(unsafe.Pointer(pOffsetY)))
	return int(r)
}

// 形状组框_取圆角大小.
//
// hShape: 形状对象句柄.
//
// pWidth: 返回圆角宽度.
//
// pHeight: 返回圆角高度.
func XShapeGroupBox_GetRoundAngle(hShape int, pWidth *int32, pHeight *int32) int {
	r, _, _ := xShapeGroupBox_GetRoundAngle.Call(uintptr(hShape), uintptr(unsafe.Pointer(pWidth)), uintptr(unsafe.Pointer(pHeight)))
	return int(r)
}

// 形状组框_启用圆角.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeGroupBox_EnableRoundAngle(hShape int, bEnable bool) int {
	r, _, _ := xShapeGroupBox_EnableRoundAngle.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}
