package xc

import "github.com/twgh/xcgui/common"

// 滑动条_创建, 创建滑动条元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func XSliderBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xSliderBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 滑动条_置范围, 设置滑动范围.
//
// hEle: 元素句柄.
//
// range_: 范围.
func XSliderBar_SetRange(hEle int, range_ int) int {
	r, _, _ := xSliderBar_SetRange.Call(uintptr(hEle), uintptr(range_))
	return int(r)
}

// 滑动条_取范围, 获取滚动范围.
//
// hEle: 元素句柄.
func XSliderBar_GetRange(hEle int) int {
	r, _, _ := xSliderBar_GetRange.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_置进度图片, 设置进度贴图.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XSliderBar_SetImageLoad(hEle int, hImage int) int {
	r, _, _ := xSliderBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 滑动条_置滑块宽度, 设置滑块按钮宽度.
//
// hEle: 元素句柄.
//
// width: 宽度.
func XSliderBar_SetButtonWidth(hEle int, width int) int {
	r, _, _ := xSliderBar_SetButtonWidth.Call(uintptr(hEle), uintptr(width))
	return int(r)
}

// 滑动条_置滑块高度, 设置滑块按钮高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func XSliderBar_SetButtonHeight(hEle int, height int) int {
	r, _, _ := xSliderBar_SetButtonHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 滑动条_置当前位置, 设置当前进度点.
//
// hEle: 元素句柄.
//
// pos: 进度点.
func XSliderBar_SetPos(hEle int, pos int) int {
	r, _, _ := xSliderBar_SetPos.Call(uintptr(hEle), uintptr(pos))
	return int(r)
}

// 滑动条_取当前位置, 获取当前进度点.
//
// hEle: 元素句柄.
func XSliderBar_GetPos(hEle int) int {
	r, _, _ := xSliderBar_GetPos.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_取滑块, 返回滑块按钮句柄.
//
// hEle: 元素句柄.
func XSliderBar_GetButton(hEle int) int {
	r, _, _ := xSliderBar_GetButton.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_置水平, 设置水平或垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.
func XSliderBar_EnableHorizon(hEle int, bHorizon bool) int {
	r, _, _ := xSliderBar_EnableHorizon.Call(uintptr(hEle), common.BoolPtr(bHorizon))
	return int(r)
}
