package widget

import (
	"github.com/twgh/xcgui/xc"
)

// 组框(形状对象).
type ShapeGroupBox struct {
	Shape
}

// 形状组框_创建, 创建组框形状对象.
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
func NewShapeGroupBox(x int, y int, cx int, cy int, pName string, hParent int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.SetHandle(xc.XShapeGroupBox_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
func NewShapeGroupBoxByHandle(handle int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeGroupBoxByName(name string) *ShapeGroupBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeGroupBoxByUID(nUID int) *ShapeGroupBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeGroupBoxByUIDName(name string) *ShapeGroupBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状组框_置边框颜色.
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) SetBorderColor(color int) int {
	return xc.XShapeGroupBox_SetBorderColor(s.Handle, color)
}

// 形状组框_置文本颜色.
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) SetTextColor(color int) int {
	return xc.XShapeGroupBox_SetTextColor(s.Handle, color)
}

// 形状组框_置字体.
//
// hFontX: 炫彩字体.
func (s *ShapeGroupBox) SetFontX(hFontX int) int {
	return xc.XShapeGroupBox_SetFontX(s.Handle, hFontX)
}

// 形状组框_置文本偏移, 设置文本偏移量.
//
// offsetX: 水平偏移.
//
// offsetY: 垂直偏移.
func (s *ShapeGroupBox) SetTextOffset(offsetX int32, offsetY int32) int {
	return xc.XShapeGroupBox_SetTextOffset(s.Handle, offsetX, offsetY)
}

// 形状组框_置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeGroupBox) SetRoundAngle(nWidth int32, nHeight int32) int {
	return xc.XShapeGroupBox_SetRoundAngle(s.Handle, nWidth, nHeight)
}

// 形状组框_置文本.
//
// pText: 文本内容.
func (s *ShapeGroupBox) SetText(pText string) int {
	return xc.XShapeGroupBox_SetText(s.Handle, pText)
}

// 形状组框_取文本偏移, 获取文本偏移量.
//
// pOffsetX: X坐标偏移量.
//
// pOffsetY: Y坐标偏移量.
func (s *ShapeGroupBox) GetTextOffset(pOffsetX *int32, pOffsetY *int32) int {
	return xc.XShapeGroupBox_GetTextOffset(s.Handle, pOffsetX, pOffsetY)
}

// 形状组框_取圆角大小.
//
// pWidth: 返回圆角宽度.
//
// pHeight: 返回圆角高度.
func (s *ShapeGroupBox) GetRoundAngle(pWidth *int32, pHeight *int32) int {
	return xc.XShapeGroupBox_GetRoundAngle(s.Handle, pWidth, pHeight)
}

// 形状组框_启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeGroupBox) EnableRoundAngle(bEnable bool) int {
	return xc.XShapeGroupBox_EnableRoundAngle(s.Handle, bEnable)
}
