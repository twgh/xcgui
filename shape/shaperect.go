package shape

import (
	"github.com/twgh/xcgui/xc"
)

type ShapeRect struct {
	Shape

	HXCGUI int
}

// 形状矩形_创建, 创建矩形形状对象, 返回句柄
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapeRect(x int, y int, cx int, cy int, hParent int) *ShapeRect {
	p := &ShapeRect{
		HXCGUI: xc.XShapeRect_Create(x, y, cx, cy, hParent),
	}
	p.HXCGUI_ = p.HXCGUI
	return p
}

// 形状矩形_置边框色, 设置边框颜色
// color: RGB颜色值.
// alpha: 透明度.
func (s *ShapeRect) SetBorderColor(color int, alpha uint8) int {
	return xc.XShapeRect_SetBorderColor(s.HXCGUI, color, alpha)
}

// 形状矩形_置填充色, 设置填充颜色
// color: RGB颜色值.
// alpha: 透明度.
func (s *ShapeRect) SetFillColor(color int, alpha uint8) int {
	return xc.XShapeRect_SetFillColor(s.HXCGUI, color, alpha)
}

// 形状矩形_置圆角大小, 设置圆角大小
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func (s *ShapeRect) SetRoundAngle(nWidth int, nHeight int) int {
	return xc.XShapeRect_SetRoundAngle(s.HXCGUI, nWidth, nHeight)
}

// 形状矩形_取圆角大小, 获取圆角大小
// pWidth: 圆角宽度.
// pHeight: 圆角高度.
func (s *ShapeRect) GetRoundAngle(pWidth *int, pHeight *int) int {
	return xc.XShapeRect_GetRoundAngle(s.HXCGUI, pWidth, pHeight)
}

// 形状矩形_启用边框, 启用绘制矩形边框
// bEnable: 是否启用.
func (s *ShapeRect) EnableBorder(bEnable bool) int {
	return xc.XShapeRect_EnableBorder(s.HXCGUI, bEnable)
}

// 形状矩形_启用填充, 启用填充矩形
// bEnable: 是否启用.
func (s *ShapeRect) EnableFill(bEnable bool) int {
	return xc.XShapeRect_EnableFill(s.HXCGUI, bEnable)
}

// 形状矩形_启用圆角
// bEnable: 是否启用.
func (s *ShapeRect) EnableRoundAngle(bEnable bool) int {
	return xc.XShapeRect_EnableRoundAngle(s.HXCGUI, bEnable)
}
