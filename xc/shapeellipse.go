package xc

import "github.com/twgh/xcgui/common"

// 形状圆_创建, 创建圆形形状对象, 返回句柄.
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
func XShapeEllipse_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xShapeEllipse_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 形状圆_置边框色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeEllipse_SetBorderColor(hShape int, color int) int {
	r, _, _ := xShapeEllipse_SetBorderColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状圆_置填充色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeEllipse_SetFillColor(hShape int, color int) int {
	r, _, _ := xShapeEllipse_SetFillColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状圆_启用边框, 启用绘制圆边框.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeEllipse_EnableBorder(hShape int, bEnable bool) int {
	r, _, _ := xShapeEllipse_EnableBorder.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}

// 形状圆_启用填充, 启用填充圆.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeEllipse_EnableFill(hShape int, bEnable bool) int {
	r, _, _ := xShapeEllipse_EnableFill.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}
