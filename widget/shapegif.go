package widget

import (
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/xc"
)

// ShapeGif 形状对象GIF.
type ShapeGif struct {
	Shape
}

// 形状GIF_创建, 创建形状对象GIF, 失败返回 nil.
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
func NewShapeGif(x, y, cx, cy int32, hParent int) *ShapeGif {
	return NewShapeGifByHandle(xc.XShapeGif_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewShapeGifByHandle(handle int) *ShapeGif {
	if handle <= 0 {
		return nil
	}
	p := &ShapeGif{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewShapeGifByName(name string) *ShapeGif {
	return NewShapeGifByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewShapeGifByUID(nUID int32) *ShapeGif {
	return NewShapeGifByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewShapeGifByUIDName(name string) *ShapeGif {
	return NewShapeGifByHandle(xc.XC_GetObjectByUIDName(name))
}

// 形状GIF_置图片, 设置GIF图片.
//
// hImage: 图片句柄.
func (s *ShapeGif) SetImage(hImage int) *ShapeGif {
	xc.XShapeGif_SetImage(s.Handle, hImage)
	return s
}

// 形状GIF_取图片, 获取图片句柄.
func (s *ShapeGif) GetImage() int {
	return xc.XShapeGif_GetImage(s.Handle)
}

// 形状GIF_取图片, 获取图片对象, 失败返回 nil.
func (s *ShapeGif) GetImageObj() *imagex.Image {
	return imagex.NewByHandle(xc.XShapeGif_GetImage(s.Handle))
}
