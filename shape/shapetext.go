package shape

import (
	"github.com/twgh/xcgui/xc"
)

type ShapeText struct {
	Shape
}

// 形状文本_创建, 创建形状对象文本
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// pName: 文本内容.
// hParent: 父对象句柄.
func NewShapeText(x int, y int, cx int, cy int, pName string, hParent int) *ShapeText {
	p := &ShapeText{}
	p.SetHXCGUI(xc.XShapeText_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 形状文本_置文本, 设置文本内容
// pName: 文本内容.
func (s *ShapeText) SetText(pName string) int {
	return xc.XShapeText_SetText(s.HXCGUI, pName)
}

// 形状文本_取文本, 获取文本内容
func (s *ShapeText) GetText() string {
	return xc.XShapeText_GetText(s.HXCGUI)
}

// 形状文本_取文本长度, 获取文本长度
func (s *ShapeText) GetTextLength() int {
	return xc.XShapeText_GetTextLength(s.HXCGUI)
}

// 形状文本_置字体
// hFontx: 字体句柄.
func (s *ShapeText) SetFont(hFontx int) int {
	return xc.XShapeText_SetFont(s.HXCGUI, hFontx)
}

// 形状文本_取字体, 返回字体句柄
func (s *ShapeText) GetFont() int {
	return xc.XShapeText_GetFont(s.HXCGUI)
}

// 形状文本_置文本颜色, 设置文本颜色
// color: RGB颜色值.
// alpha: 透明度
func (s *ShapeText) SetTextColor(color int, alpha uint8) int {
	return xc.XShapeText_SetTextColor(s.HXCGUI, color, alpha)
}

// 形状文本_取文本颜色
func (s *ShapeText) GetTextColor() int {
	return xc.XShapeText_GetTextColor(s.HXCGUI)
}

// 形状文本_置文本对齐
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_
func (s *ShapeText) SetTextAlign(align int) int {
	return xc.XShapeText_SetTextAlign(s.HXCGUI, align)
}

// 形状文本_置偏移, 设置内容偏移
// x: X坐标.
// y: Y坐标.
func (s *ShapeText) SetOffset(x int, y int) int {
	return xc.XShapeText_SetOffset(s.HXCGUI, x, y)
}
