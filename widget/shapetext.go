package widget

import (
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ShapeText 形状对象文本.
type ShapeText struct {
	Shape
}

// 形状文本_创建, 创建形状对象文本, 失败返回 nil.
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
func NewShapeText(x, y, cx, cy int32, pName string, hParent int) *ShapeText {
	return NewShapeTextByHandle(xc.XShapeText_Create(x, y, cx, cy, pName, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewShapeTextByHandle(handle int) *ShapeText {
	if handle <= 0 {
		return nil
	}
	p := &ShapeText{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewShapeTextByName(name string) *ShapeText {
	return NewShapeTextByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewShapeTextByUID(nUID int32) *ShapeText {
	return NewShapeTextByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewShapeTextByUIDName(name string) *ShapeText {
	return NewShapeTextByHandle(xc.XC_GetObjectByUIDName(name))
}

// 形状文本_置文本, 设置文本内容.
//
// pName: 文本内容.
func (s *ShapeText) SetText(pName string) *ShapeText {
	xc.XShapeText_SetText(s.Handle, pName)
	return s
}

// 形状文本_取文本, 获取文本内容.
func (s *ShapeText) GetText() string {
	return xc.XShapeText_GetText(s.Handle)
}

// 形状文本_取文本长度, 获取文本长度.
func (s *ShapeText) GetTextLength() int32 {
	return xc.XShapeText_GetTextLength(s.Handle)
}

// 形状文本_置字体.
//
// hFontx: 字体句柄.
func (s *ShapeText) SetFont(hFontx int) *ShapeText {
	xc.XShapeText_SetFont(s.Handle, hFontx)
	return s
}

// 形状文本_取字体, 返回字体句柄.
func (s *ShapeText) GetFont() int {
	return xc.XShapeText_GetFont(s.Handle)
}

// 形状文本_取字体, 返回字体对象, 失败返回 nil.
func (s *ShapeText) GetFontObj() *font.Font {
	return font.NewByHandle(xc.XShapeText_GetFont(s.Handle))
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// color: xc.RGBA 颜色值.
func (s *ShapeText) SetTextColor(color uint32) *ShapeText {
	xc.XShapeText_SetTextColor(s.Handle, color)
	return s
}

// 形状文本_取文本颜色.
func (s *ShapeText) GetTextColor() uint32 {
	return xc.XShapeText_GetTextColor(s.Handle)
}

// 形状文本_置文本对齐.
//
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (s *ShapeText) SetTextAlign(align xcc.TextFormatFlag_) *ShapeText {
	xc.XShapeText_SetTextAlign(s.Handle, align)
	return s
}

// 形状文本_置偏移, 设置内容偏移.
//
// x: X坐标.
//
// y: Y坐标.
func (s *ShapeText) SetOffset(x, y int32) *ShapeText {
	xc.XShapeText_SetOffset(s.Handle, x, y)
	return s
}
