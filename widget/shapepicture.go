package widget

import (
	"github.com/twgh/xcgui/xc"
)

type ShapePicture struct {
	Shape
}

// 形状图片_创建, 创建形状对象-图片
// x: x坐标.
// y: y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapePicture(x int, y int, cx int, cy int, hParent int) *ShapePicture {
	p := &ShapePicture{}
	p.SetHandle(xc.XShapePic_Create(x, y, cx, cy, hParent))
	return p
}

// 形状图片_置图片, 设置图片
// hImage: 图片句柄.
func (s *ShapePicture) SetImage(hImage int) int {
	return xc.XShapePic_SetImage(s.Handle, hImage)
}

// 形状图片_取图片, 获取图片句柄
func (s *ShapePicture) GetImage() int {
	return xc.XShapePic_GetImage(s.Handle)
}
