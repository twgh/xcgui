package shape

import (
	xc "github.com/twgh/xcgui"
)

type ShapeGif struct {
	Shape

	HXCGUI int
}

// 形状GIF_创建, 创建形状对象GIF, 成功返回形状对象GIF句柄, 否则返回NULL
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapeGif(x int, y int, cx int, cy int, hParent int) *ShapeGif {
	p := &ShapeGif{
		HXCGUI: xc.XShapeGif_Create(x, y, cx, cy, hParent),
	}
	p.HXCGUI_ = p.HXCGUI
	return p
}

// 形状GIF_置图片, 设置GIF图片
// hImage: 图片句柄.
func (s *ShapeGif) SetImage(hImage int) int {
	return xc.XShapeGif_SetImage(s.HXCGUI, hImage)
}

// 形状GIF_取图片, 获取图片句柄
func (s *ShapeGif) GetImage() int {
	return xc.XShapeGif_GetImage(s.HXCGUI)
}
