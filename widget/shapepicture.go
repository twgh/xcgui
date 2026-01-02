package widget

import (
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/xc"
)

// ShapePicture 形状对象图片.
type ShapePicture struct {
	Shape
}

// NewShapePicture 形状图片_创建, 创建形状对象-图片, 失败返回 nil.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func NewShapePicture(x, y, cx, cy int32, hParent int) *ShapePicture {
	return NewShapePictureByHandle(xc.XShapePic_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewShapePictureByHandle(handle int) *ShapePicture {
	if handle <= 0 {
		return nil
	}
	p := &ShapePicture{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewShapePictureByName(name string) *ShapePicture {
	return NewShapePictureByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewShapePictureByUID(nUID int32) *ShapePicture {
	return NewShapePictureByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewShapePictureByUIDName(name string) *ShapePicture {
	return NewShapePictureByHandle(xc.XC_GetObjectByUIDName(name))
}

// 形状图片_置图片, 设置图片.
//
// hImage: 图片句柄.
func (s *ShapePicture) SetImage(hImage int) *ShapePicture {
	xc.XShapePic_SetImage(s.Handle, hImage)
	return s
}

// 形状图片_取图片, 获取图片句柄.
func (s *ShapePicture) GetImage() int {
	return xc.XShapePic_GetImage(s.Handle)
}

// 形状图片_取图片, 获取图片对象, 失败返回 nil.
func (s *ShapePicture) GetImageObj() *imagex.Image {
	return imagex.NewByHandle(xc.XShapePic_GetImage(s.Handle))
}
