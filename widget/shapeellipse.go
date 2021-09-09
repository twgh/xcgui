package widget

import (
	"github.com/twgh/xcgui/xc"
)

type ShapeEllipse struct {
	Shape
}

// 形状圆_创建, 创建圆形形状对象
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapeEllipse(x int, y int, cx int, cy int, hParent int) *ShapeEllipse {
	p := &ShapeEllipse{}
	p.SetHandle(xc.XShapeEllipse_Create(x, y, cx, cy, hParent))
	return p
}

// 形状圆_置边框色
// color: RGB颜色值.
// alpha: 透明度.
func (s *ShapeEllipse) SetBorderColor(color int, alpha uint8) int {
	return xc.XShapeEllipse_SetBorderColor(s.Handle, color, alpha)
}

// 形状圆_置填充色
// color: RGB颜色值.
// alpha: 透明度.
func (s *ShapeEllipse) SetFillColor(color int, alpha uint8) int {
	return xc.XShapeEllipse_SetFillColor(s.Handle, color, alpha)
}

// 形状圆_启用边框, 启用绘制圆边框
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableBorder(bEnable bool) int {
	return xc.XShapeEllipse_EnableBorder(s.Handle, bEnable)
}

// 形状圆_启用填充, 启用填充圆
// bEnable: 是否启用.
func (s *ShapeEllipse) EnableFill(bEnable bool) int {
	return xc.XShapeEllipse_EnableFill(s.Handle, bEnable)
}
