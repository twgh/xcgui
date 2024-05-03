package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeRect 矩形形状对象.
type ShapeRect struct {
	Shape
}

// 形状矩形_创建, 创建矩形形状对象.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func NewShapeRect(x int, y int, cx int, cy int, hParent int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(xc.XShapeRect_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewShapeRectByHandle(handle int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeRectByName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeRectByUID(nUID int) *ShapeRect {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeRectByUIDName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状矩形_置边框色, 设置边框颜色.
//
// color: ARGB 颜色值.
func (s *ShapeRect) SetBorderColor(color int) *ShapeRect {
	xc.XShapeRect_SetBorderColor(s.Handle, color)
	return s
}

// 形状矩形_置填充色, 设置填充颜色.
//
// color: ARGB 颜色值.
func (s *ShapeRect) SetFillColor(color int) *ShapeRect {
	xc.XShapeRect_SetFillColor(s.Handle, color)
	return s
}

// 形状矩形_置圆角大小, 设置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeRect) SetRoundAngle(nWidth int, nHeight int32) *ShapeRect {
	xc.XShapeRect_SetRoundAngle(s.Handle, nWidth, nHeight)
	return s
}

// 形状矩形_取圆角大小, 获取圆角大小.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func (s *ShapeRect) GetRoundAngle(pWidth *int, pHeight *int32) *ShapeRect {
	xc.XShapeRect_GetRoundAngle(s.Handle, pWidth, pHeight)
	return s
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableBorder(bEnable bool) *ShapeRect {
	xc.XShapeRect_EnableBorder(s.Handle, bEnable)
	return s
}

// 形状矩形_启用填充, 启用填充矩形.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableFill(bEnable bool) *ShapeRect {
	xc.XShapeRect_EnableFill(s.Handle, bEnable)
	return s
}

// 形状矩形_启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableRoundAngle(bEnable bool) *ShapeRect {
	xc.XShapeRect_EnableRoundAngle(s.Handle, bEnable)
	return s
}
