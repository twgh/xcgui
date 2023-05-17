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
func XProgBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xProgBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 进度条_置范围, 设置范围.
//
// hEle: 元素句柄.
//
// range_: 范围.
func XProgBar_SetRange(hEle int, range_ int) int {
	r, _, _ := xProgBar_SetRange.Call(uintptr(hEle), uintptr(range_))
	return int(r)
}

// 进度条_取范围.
//
// hEle: 元素句柄.
func XProgBar_GetRange(hEle int) int {
	r, _, _ := xProgBar_GetRange.Call(uintptr(hEle))
	return int(r)
}

// 进度条_置进度图片.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XProgBar_SetImageLoad(hEle int, hImage int) int {
	r, _, _ := xProgBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 进度条_置进度, 设置位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func XProgBar_SetPos(hEle int, pos int) int {
	r, _, _ := xProgBar_SetPos.Call(uintptr(hEle), uintptr(pos))
	return int(r)
}

// 进度条_取进度, 获取当前位置点.
//
// hEle: 元素句柄.
func XProgBar_GetPos(hEle int) int {
	r, _, _ := xProgBar_GetPos.Call(uintptr(hEle))
	return int(r)
}

// 进度条_置水平, 设置水平或垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.
func XProgBar_EnableHorizon(hEle int, bHorizon bool) int {
	r, _, _ := xProgBar_EnableHorizon.Call(uintptr(hEle), common.BoolPtr(bHorizon))
	return int(r)
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// hEle: 元素句柄.
//
// bStretch: 缩放.
func XProgBar_EnableStretch(hEle int, bStretch bool) bool {
	r, _, _ := xProgBar_EnableStretch.Call(uintptr(hEle), common.BoolPtr(bStretch))
	return r != 0
}

// 进度条_启用进度文本 显示进度值文本.
//
// hEle: 元素句柄.
//
// bShow: 是否启用.
func XProgBar_EnableShowText(hEle int, bShow bool) bool {
	r, _, _ := xProgBar_EnableShowText.Call(uintptr(hEle), common.BoolPtr(bShow))
	return r != 0
}

// 进度条_置进度颜色. 设置进度颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
func XProgBar_SetColorLoad(hEle int, color int) bool {
	r, _, _ := xProgBar_SetColorLoad.Call(uintptr(hEle), uintptr(color))
	return r != 0
}
