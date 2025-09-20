package xc

import "github.com/twgh/xcgui/common"

// 进度条_创建, 创建进度条元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func XProgBar_Create(x, y, cx, cy int32, hParent int) int {
	r, _, _ := xProgBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 进度条_置范围, 设置范围.
//
// hEle: 元素句柄.
//
// range_: 范围.
func XProgBar_SetRange(hEle int, range_ int32) {
	xProgBar_SetRange.Call(uintptr(hEle), uintptr(range_))
}

// 进度条_取范围.
//
// hEle: 元素句柄.
func XProgBar_GetRange(hEle int) int32 {
	r, _, _ := xProgBar_GetRange.Call(uintptr(hEle))
	return int32(r)
}

// 进度条_置进度图片.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XProgBar_SetImageLoad(hEle int, hImage int) {
	xProgBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
}

// 进度条_置进度, 设置位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func XProgBar_SetPos(hEle int, pos int32) {
	xProgBar_SetPos.Call(uintptr(hEle), uintptr(pos))
}

// 进度条_取进度, 获取当前位置点.
//
// hEle: 元素句柄.
func XProgBar_GetPos(hEle int) int32 {
	r, _, _ := xProgBar_GetPos.Call(uintptr(hEle))
	return int32(r)
}

// 进度条_置水平, 设置水平或垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.
func XProgBar_EnableHorizon(hEle int, bHorizon bool) {
	xProgBar_EnableHorizon.Call(uintptr(hEle), common.BoolPtr(bHorizon))
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// hEle: 元素句柄.
//
// bStretch: 缩放.
func XProgBar_EnableStretch(hEle int, bStretch bool) {
	xProgBar_EnableStretch.Call(uintptr(hEle), common.BoolPtr(bStretch))
}

// 进度条_启用进度文本 显示进度值文本.
//
// hEle: 元素句柄.
//
// bShow: 是否启用.
func XProgBar_EnableShowText(hEle int, bShow bool) {
	xProgBar_EnableShowText.Call(uintptr(hEle), common.BoolPtr(bShow))
}

// 进度条_置进度颜色. 设置进度颜色.
//
// hEle: 元素句柄.
//
// color: xc.RGBA 颜色.
func XProgBar_SetColorLoad(hEle int, color uint32) {
	xProgBar_SetColorLoad.Call(uintptr(hEle), uintptr(color))
}
