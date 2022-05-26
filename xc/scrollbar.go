package xc

import "github.com/twgh/xcgui/common"

// 滚动条_创建, 创建滚动条元素, 返回元素句柄.
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
func XSBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xSBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 滚动条_置范围, 设置滚动范围.
//
// hEle: 元素句柄.
//
// range_: 范围.
func XSBar_SetRange(hEle int, range_ int) int {
	r, _, _ := xSBar_SetRange.Call(uintptr(hEle), uintptr(range_))
	return int(r)
}

// 滚动条_取范围, 获取滚动范围.
//
// hEle: 元素句柄.
func XSBar_GetRange(hEle int) int {
	r, _, _ := xSBar_GetRange.Call(uintptr(hEle))
	return int(r)
}

// 滚动条_显示上下按钮, 显示隐藏滚动条上下按钮.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.
func XSBar_ShowButton(hEle int, bShow bool) int {
	r, _, _ := xSBar_ShowButton.Call(uintptr(hEle), common.BoolPtr(bShow))
	return int(r)
}

// 滚动条_置滑块长度.
//
// hEle: 元素句柄.
//
// length: 长度.
func XSBar_SetSliderLength(hEle int, length int) int {
	r, _, _ := xSBar_SetSliderLength.Call(uintptr(hEle), uintptr(length))
	return int(r)
}

// 滚动条_置滑块最小长度.
//
// hEle: 元素句柄.
//
// minLength: 长度.
func XSBar_SetSliderMinLength(hEle int, minLength int) int {
	r, _, _ := xSBar_SetSliderMinLength.Call(uintptr(hEle), uintptr(minLength))
	return int(r)
}

// 滚动条_置滑块两边间隔, 设置滑块两边的间隔大小.
//
// hEle: 元素句柄.
//
// nPadding: 间隔大小.
func XSBar_SetSliderPadding(hEle int, nPadding int) int {
	r, _, _ := xSBar_SetSliderPadding.Call(uintptr(hEle), uintptr(nPadding))
	return int(r)
}

// 滚动条_置水平, 设置水平或者垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.
func XSBar_EnableHorizon(hEle int, bHorizon bool) bool {
	r, _, _ := xSBar_EnableHorizon.Call(uintptr(hEle), common.BoolPtr(bHorizon))
	return r != 0
}

// 滚动条_取滑块最大长度.
//
// hEle: 元素句柄.
func XSBar_GetSliderMaxLength(hEle int) int {
	r, _, _ := xSBar_GetSliderMaxLength.Call(uintptr(hEle))
	return int(r)
}

// 滚动条_向上滚动.
//
// hEle: 元素句柄.
func XSBar_ScrollUp(hEle int) bool {
	r, _, _ := xSBar_ScrollUp.Call(uintptr(hEle))
	return r != 0
}

// 滚动条_向下滚动.
//
// hEle: 元素句柄.
func XSBar_ScrollDown(hEle int) bool {
	r, _, _ := xSBar_ScrollDown.Call(uintptr(hEle))
	return r != 0
}

// 滚动条_滚动到顶部.
//
// hEle: 元素句柄.
func XSBar_ScrollTop(hEle int) bool {
	r, _, _ := xSBar_ScrollTop.Call(uintptr(hEle))
	return r != 0
}

// 滚动条_滚动到底部.
//
// hEle: 元素句柄.
func XSBar_ScrollBottom(hEle int) bool {
	r, _, _ := xSBar_ScrollBottom.Call(uintptr(hEle))
	return r != 0
}

// 滚动条_滚动到指定位置, 滚动到指定位置点, 触发事件: XE_SBAR_SCROLL.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func XSBar_ScrollPos(hEle int, pos int) bool {
	r, _, _ := xSBar_ScrollPos.Call(uintptr(hEle), uintptr(pos))
	return r != 0
}

// 滚动条_取上按钮, 获取上按钮, 返回按钮句柄.
//
// hEle: 元素句柄.
func XSBar_GetButtonUp(hEle int) int {
	r, _, _ := xSBar_GetButtonUp.Call(uintptr(hEle))
	return int(r)
}

// 滚动条_取下按钮, 获取下按钮, 返回按钮句柄.
//
// hEle: 元素句柄.
func XSBar_GetButtonDown(hEle int) int {
	r, _, _ := xSBar_GetButtonDown.Call(uintptr(hEle))
	return int(r)
}

// 滚动条_取滑块, 获取滑动按钮, 返回按钮句柄.
//
// hEle: 元素句柄.
func XSBar_GetButtonSlider(hEle int) int {
	r, _, _ := xSBar_GetButtonSlider.Call(uintptr(hEle))
	return int(r)
}
