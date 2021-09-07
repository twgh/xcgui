package xc

// 进度条_创建, 创建进度条元素, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func XProgBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xProgBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 进度条_置范围, 设置范围
// hEle: 元句柄.
// range_: 范围.
func XProgBar_SetRange(hEle int, range_ int) int {
	r, _, _ := xProgBar_SetRange.Call(uintptr(hEle), uintptr(range_))
	return int(r)
}

// 进度条_取范围
// hEle: 元句柄.
func XProgBar_GetRange(hEle int) int {
	r, _, _ := xProgBar_GetRange.Call(uintptr(hEle))
	return int(r)
}

// 进度条_置进度图片
// hEle: 元句柄.
// hImage: 图片句柄.
func XProgBar_SetImageLoad(hEle int, hImage int) int {
	r, _, _ := xProgBar_SetImageLoad.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 进度条_置两端大小, 设置两端间隔大小
// hEle: 元句柄.
// leftSize: 左边间隔大小.
// rightSize: 右边间隔大小.
func XProgBar_SetSpaceTwo(hEle int, leftSize int, rightSize int) int {
	r, _, _ := xProgBar_SetSpaceTwo.Call(uintptr(hEle), uintptr(leftSize), uintptr(rightSize))
	return int(r)
}

// 进度条_置进度, 设置位置点
// hEle: 元句柄.
// pos: 位置点.
func XProgBar_SetPos(hEle int, pos int) int {
	r, _, _ := xProgBar_SetPos.Call(uintptr(hEle), uintptr(pos))
	return int(r)
}

// 进度条_取进度, 获取当前位置点
// hEle: 元句柄.
func XProgBar_GetPos(hEle int) int {
	r, _, _ := xProgBar_GetPos.Call(uintptr(hEle))
	return int(r)
}

// 进度条_置水平, 设置水平或垂直
// hEle: 元句柄.
// bHorizon: 水平或垂直.
func XProgBar_SetHorizon(hEle int, bHorizon bool) int {
	r, _, _ := xProgBar_SetHorizon.Call(uintptr(hEle), boolPtr(bHorizon))
	return int(r)
}
