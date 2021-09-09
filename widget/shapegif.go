package widget

import (
	"github.com/twgh/xcgui/xc"
)

type ShapeGif struct {
	Shape
}

// 形状GIF_创建, 创建形状对象GIF
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapeGif(x int, y int, cx int, cy int, hParent int) *ShapeGif {
	p := &ShapeGif{}
	p.SetHandle(xc.XShapeGif_Create(x, y, cx, cy, hParent))
	return p
}

// 形状GIF_置图片, 设置GIF图片
// hImage: 图片句柄.
func (s *ShapeGif) SetImage(hImage int) int {
	return xc.XShapeGif_SetImage(s.Handle, hImage)
}

// 形状GIF_取图片, 获取图片句柄
func (s *ShapeGif) GetImage() int {
	return xc.XShapeGif_GetImage(s.Handle)
}