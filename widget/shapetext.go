package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ShapeText 形状对象文本.
type ShapeText struct {
	Shape
}

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
func NewShapeText(x, y, cx, cy int32, pName string, hParent int) *ShapeText {
	p := &ShapeText{}
	p.SetHandle(xc.XShapeText_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
func NewShapeTextByHandle(handle int) *ShapeText {
	p := &ShapeText{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeTextByName(name string) *ShapeText {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeTextByUID(nUID int) *ShapeText {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeTextByUIDName(name string) *ShapeText {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状文本_置文本, 设置文本内容.
//
// pName: 文本内容.
func (s *ShapeText) SetText(pName string) int {
	return xc.XShapeText_SetText(s.Handle, pName)
}

// 形状文本_取文本, 获取文本内容.
func (s *ShapeText) GetText() string {
	return xc.XShapeText_GetText(s.Handle)
}

// 形状文本_取文本长度, 获取文本长度.
func (s *ShapeText) GetTextLength() int {
	return xc.XShapeText_GetTextLength(s.Handle)
}

// 形状文本_置字体.
//
// hFontx: 字体句柄.
func (s *ShapeText) SetFont(hFontx int) int {
	return xc.XShapeText_SetFont(s.Handle, hFontx)
}

// 形状文本_取字体, 返回字体句柄.
func (s *ShapeText) GetFont() int {
	return xc.XShapeText_GetFont(s.Handle)
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// color: ABGR 颜色值.
func (s *ShapeText) SetTextColor(color int) int {
	return xc.XShapeText_SetTextColor(s.Handle, color)
}

// 形状文本_取文本颜色.
func (s *ShapeText) GetTextColor() int {
	return xc.XShapeText_GetTextColor(s.Handle)
}

// 形状文本_置文本对齐.
//
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (s *ShapeText) SetTextAlign(align xcc.TextFormatFlag_) int {
	return xc.XShapeText_SetTextAlign(s.Handle, align)
}

// 形状文本_置偏移, 设置内容偏移.
//
// x: X坐标.
//
// y: Y坐标.
func (s *ShapeText) SetOffset(x int, y int) int {
	return xc.XShapeText_SetOffset(s.Handle, x, y)
}
