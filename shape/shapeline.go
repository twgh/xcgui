package shape

import (
	"github.com/twgh/xcgui/xc"
)

type ShapeLine struct {
	Shape
}

// 形状线_创建
// x1: 坐标.
// y1: 坐标.
// x2: 坐标.
// y2: 坐标.
// hParent: 父对象句柄.
func NewShapeLine(x1 int, y1 int, x2 int, y2 int, hParent int) *ShapeLine {
	p := &ShapeLine{}
	p.SetHXCGUI(xc.XShapeLine_Create(x1, y1, x2, y2, hParent))
	return p
}

// 形状线_置位置
// x1: 坐标.
// y1: 坐标.
// x2: 坐标.
// y2: 坐标.
func (s *ShapeLine) SetPosition(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XShapeLine_SetPosition(s.HXCGUI, x1, y1, x2, y2)
}

// 形状线_置颜色, 设置直线颜色
// color: RGB颜色值.
// alpha: 透明度.
func (s *ShapeLine) SetColor(color int, alpha uint8) int {
	return xc.XShapeLine_SetColor(s.HXCGUI, color, alpha)
}
