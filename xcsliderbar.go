package xcgui

// 滑动条_创建
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func XSliderBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xSliderBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 滑动条_置范围
// hEle: 元素句柄.
// range_: 范围.
func XSliderBar_SetRange(hEle int, range_ int) int {
	r, _, _ := xSliderBar_SetRange.Call(uintptr(hEle), uintptr(range_))
	return int(r)
}

// 滑动条_取范围
// hEle: 元素句柄.
func XSliderBar_GetRange(hEle int) int {
	r, _, _ := xSliderBar_GetRange.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_置进度图片
// hEle: 元素句柄.
// hImage: 图片句柄.
func XSliderBar_SetImageLoad(hEle int, hImage int) int {
	r, _, _ := xSliderBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 滑动条_置滑块宽度
// hEle: 元素句柄.
// width: 宽度.
func XSliderBar_SetButtonWidth(hEle int, width int) int {
	r, _, _ := xSliderBar_SetButtonWidth.Call(uintptr(hEle), uintptr(width))
	return int(r)
}

// 滑动条_置滑块高度
// hEle: 元素句柄.
// height: 高度.
func XSliderBar_SetButtonHeight(hEle int, height int) int {
	r, _, _ := xSliderBar_SetButtonHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 滑动条_置两端大小
// hEle: 元素句柄.
// leftSize: 左边间隔大小.
// rightSize: 右边间隔大小.
func XSliderBar_SetSpaceTwo(hEle int, leftSize int, rightSize int) int {
	r, _, _ := xSliderBar_SetSpaceTwo.Call(uintptr(hEle), uintptr(leftSize), uintptr(rightSize))
	return int(r)
}

// 滑动条_置当前位置
// hEle: 元素句柄.
// pos: 进度点.
func XSliderBar_SetPos(hEle int, pos int) int {
	r, _, _ := xSliderBar_SetPos.Call(uintptr(hEle), uintptr(pos))
	return int(r)
}

// 滑动条_取当前位置
// hEle: 元素句柄.
func XSliderBar_GetPos(hEle int) int {
	r, _, _ := xSliderBar_GetPos.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_取滑块
// hEle: 元素句柄.
func XSliderBar_GetButton(hEle int) int {
	r, _, _ := xSliderBar_GetButton.Call(uintptr(hEle))
	return int(r)
}

// 滑动条_置水平
// hEle: 元素句柄.
// bHorizon: 水平或垂直.
func XSliderBar_SetHorizon(hEle int, bHorizon bool) int {
	r, _, _ := xSliderBar_SetHorizon.Call(uintptr(hEle), boolPtr(bHorizon))
	return int(r)
}
