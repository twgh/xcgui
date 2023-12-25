package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 形状文本_创建, 创建形状对象文本.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父对象句柄.
func XShapeText_Create(x, y, cx, cy int32, pName string, hParent int) int {
	r, _, _ := xShapeText_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pName), uintptr(hParent))
	return int(r)
}

// 形状文本_置文本, 设置文本内容.
//
// hTextBlock: 形状对象文本句柄.
//
// pName: 文本内容.
func XShapeText_SetText(hTextBlock int, pName string) int {
	r, _, _ := xShapeText_SetText.Call(uintptr(hTextBlock), common.StrPtr(pName))
	return int(r)
}

// 形状文本_取文本, 获取文本内容.
//
// hTextBlock: 形状对象文本句柄.
func XShapeText_GetText(hTextBlock int) string {
	r, _, _ := xShapeText_GetText.Call(uintptr(hTextBlock))
	return common.UintPtrToString(r)
}

// 形状文本_取文本长度, 获取文本长度.
//
// hTextBlock: 形状对象文本句柄.
func XShapeText_GetTextLength(hTextBlock int) int {
	r, _, _ := xShapeText_GetTextLength.Call(uintptr(hTextBlock))
	return int(r)
}

// 形状文本_置字体.
//
// hTextBlock: 形状对象文本句柄.
//
// hFontx: 字体句柄.
func XShapeText_SetFont(hTextBlock int, hFontx int) int {
	r, _, _ := xShapeText_SetFont.Call(uintptr(hTextBlock), uintptr(hFontx))
	return int(r)
}

// 形状文本_取字体, 返回字体句柄.
//
// hTextBlock: 形状对象文本句柄.
func XShapeText_GetFont(hTextBlock int) int {
	r, _, _ := xShapeText_GetFont.Call(uintptr(hTextBlock))
	return int(r)
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// hTextBlock: 形状对象文本句柄.
//
// color: ABGR 颜色值.
func XShapeText_SetTextColor(hTextBlock int, color int) int {
	r, _, _ := xShapeText_SetTextColor.Call(uintptr(hTextBlock), uintptr(color))
	return int(r)
}

// 形状文本_取文本颜色.
//
// hTextBlock: 形状对象文本句柄.
func XShapeText_GetTextColor(hTextBlock int) int {
	r, _, _ := xShapeText_GetTextColor.Call(uintptr(hTextBlock))
	return int(r)
}

// 形状文本_置文本对齐.
//
// hTextBlock: 形状对象文本句柄.
//
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func XShapeText_SetTextAlign(hTextBlock int, align xcc.TextFormatFlag_) int {
	r, _, _ := xShapeText_SetTextAlign.Call(uintptr(hTextBlock), uintptr(align))
	return int(r)
}

// 形状文本_置偏移, 设置内容偏移.
//
// hTextBlock: 形状对象文本句柄.
//
// x: X坐标.
//
// y: Y坐标.
func XShapeText_SetOffset(hTextBlock int, x int, y int) int {
	r, _, _ := xShapeText_SetOffset.Call(uintptr(hTextBlock), uintptr(x), uintptr(y))
	return int(r)
}
