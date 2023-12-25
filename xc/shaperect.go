package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

// 形状矩形_创建, 创建矩形形状对象, 返回句柄.
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
func XShapeRect_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xShapeRect_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 形状矩形_置边框色, 设置边框颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeRect_SetBorderColor(hShape int, color int) int {
	r, _, _ := xShapeRect_SetBorderColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状矩形_置填充色, 设置填充颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func XShapeRect_SetFillColor(hShape int, color int) int {
	r, _, _ := xShapeRect_SetFillColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 形状矩形_置圆角大小, 设置圆角大小.
//
// hShape: 形状对象句柄.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XShapeRect_SetRoundAngle(hShape int, nWidth int, nHeight int32) int {
	r, _, _ := xShapeRect_SetRoundAngle.Call(uintptr(hShape), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 形状矩形_取圆角大小, 获取圆角大小.
//
// hShape: 形状对象句柄.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func XShapeRect_GetRoundAngle(hShape int, pWidth *int, pHeight *int32) int {
	r, _, _ := xShapeRect_GetRoundAngle.Call(uintptr(hShape), uintptr(unsafe.Pointer(pWidth)), uintptr(unsafe.Pointer(pHeight)))
	return int(r)
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeRect_EnableBorder(hShape int, bEnable bool) int {
	r, _, _ := xShapeRect_EnableBorder.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}

// 形状矩形_启用填充, 启用填充矩形.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeRect_EnableFill(hShape int, bEnable bool) int {
	r, _, _ := xShapeRect_EnableFill.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}

// 形状矩形_启用圆角.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func XShapeRect_EnableRoundAngle(hShape int, bEnable bool) int {
	r, _, _ := xShapeRect_EnableRoundAngle.Call(uintptr(hShape), common.BoolPtr(bEnable))
	return int(r)
}
