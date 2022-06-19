package widget

import (
	"github.com/twgh/xcgui/xc"
)

// 圆形(形状对象).
type ShapeEllipse struct {
	Shape
}

// 形状圆_创建, 创建圆形形状对象.
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
func NewShapeEllipse(x int, y int, cx int, cy int, hParent int) *ShapeEllipse {
	p := &ShapeEllipse{}
	p.SetHandle(xc.XShapeEllipse_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewShapeEllipseByHandle(handle int) *ShapeEllipse {
	p := &ShapeEllipse{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeEllipseByName(name string) *ShapeEllipse {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeEllipseByUID(nUID int) *ShapeEllipse {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeEllipseByUIDName(name string) *ShapeEllipse {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状圆_置边框色.
//
// color: ABGR 颜色值.
func (s *ShapeEllipse) SetBorderColor(color int) int {
	return xc.XShapeEllipse_SetBorderColor(s.Handle, color)
}

// 形状圆_置填充色.
//
// color: ABGR 颜色值.
func (s *ShapeEllipse) SetFillColor(color int) int {
	return xc.XShapeEllipse_SetFillColor(s.Handle, color)
}

// 形状圆_启用边框, 启用绘制圆边框.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableBorder(bEnable bool) int {
	return xc.XShapeEllipse_EnableBorder(s.Handle, bEnable)
}

// 形状圆_启用填充, 启用填充圆.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableFill(bEnable bool) int {
	return xc.XShapeEllipse_EnableFill(s.Handle, bEnable)
}
