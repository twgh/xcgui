package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeEllipse 圆形(形状对象).
type ShapeEllipse struct {
	Shape
}

// 形状圆_创建, 创建圆形形状对象, 失败返回 nil.
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
func NewShapeEllipse(x, y, cx, cy int32, hParent int) *ShapeEllipse {
	return NewShapeEllipseByHandle(xc.XShapeEllipse_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewShapeEllipseByHandle(handle int) *ShapeEllipse {
	if handle <= 0 {
		return nil
	}
	p := &ShapeEllipse{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewShapeEllipseByName(name string) *ShapeEllipse {
	return NewShapeEllipseByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewShapeEllipseByUID(nUID int32) *ShapeEllipse {
	return NewShapeEllipseByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewShapeEllipseByUIDName(name string) *ShapeEllipse {
	return NewShapeEllipseByHandle(xc.XC_GetObjectByUIDName(name))
}

// 形状圆_置边框色.
//
// color: xc.RGBA 颜色值.
func (s *ShapeEllipse) SetBorderColor(color uint32) *ShapeEllipse {
	xc.XShapeEllipse_SetBorderColor(s.Handle, color)
	return s
}

// 形状圆_置填充色.
//
// color: xc.RGBA 颜色值.
func (s *ShapeEllipse) SetFillColor(color uint32) *ShapeEllipse {
	xc.XShapeEllipse_SetFillColor(s.Handle, color)
	return s
}

// 形状圆_启用边框, 启用绘制圆边框.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableBorder(bEnable bool) *ShapeEllipse {
	xc.XShapeEllipse_EnableBorder(s.Handle, bEnable)
	return s
}

// 形状圆_启用填充, 启用填充圆.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableFill(bEnable bool) *ShapeEllipse {
	xc.XShapeEllipse_EnableFill(s.Handle, bEnable)
	return s
}
