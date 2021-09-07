package xc

// 工具条_创建
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func XToolBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xToolBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 工具条_插入元素
// hEle: 元素句柄.
// hNewEle: 将要插入的元素.
// index: 插入位置索引
func XToolBar_InsertEle(hEle int, hNewEle int, index int) int {
	r, _, _ := xToolBar_InsertEle.Call(uintptr(hEle), uintptr(hNewEle), uintptr(index))
	return int(r)
}

// 工具条_插入分割栏
// hEle: 元素句柄.
// index: 插入位置索引
// color: RGB颜色
func XToolBar_InsertSeparator(hEle int, index int, color int) int {
	r, _, _ := xToolBar_InsertSeparator.Call(uintptr(hEle), uintptr(index), uintptr(color))
	return int(r)
}

// 工具条_启用下拉菜单
// hEle: 元素句柄.
// bEnable: 是否启用.
func XToolBar_EnableButtonMenu(hEle int, bEnable bool) int {
	r, _, _ := xToolBar_EnableButtonMenu.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 工具条_取元素
// hEle: 元素句柄.
// index: 索引值.
func XToolBar_GetEle(hEle int, index int) int {
	r, _, _ := xToolBar_GetEle.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// 工具条_取左滚动按钮
// hEle: 元素句柄.
func XToolBar_GetButtonLeft(hEle int) int {
	r, _, _ := xToolBar_GetButtonLeft.Call(uintptr(hEle))
	return int(r)
}

// 工具条_取右滚动按钮
// hEle: 元素句柄.
func XToolBar_GetButtonRight(hEle int) int {
	r, _, _ := xToolBar_GetButtonRight.Call(uintptr(hEle))
	return int(r)
}

// 工具条_取菜单按钮
// hEle: 元素句柄.
func XToolBar_GetButtonMenu(hEle int) int {
	r, _, _ := xToolBar_GetButtonMenu.Call(uintptr(hEle))
	return int(r)
}

// 工具条_置间距
// hEle: 元素句柄.
// nSize: 间距大小.
func XToolBar_SetSpace(hEle int, nSize int) int {
	r, _, _ := xToolBar_SetSpace.Call(uintptr(hEle), uintptr(nSize))
	return int(r)
}

// 工具条_删除元素
// hEle: 元素句柄.
// index: 索引值.
func XToolBar_DeleteEle(hEle int, index int) int {
	r, _, _ := xToolBar_DeleteEle.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// 工具条_删除全部
// hEle: 元素句柄.
func XToolBar_DeleteAllEle(hEle int) int {
	r, _, _ := xToolBar_DeleteAllEle.Call(uintptr(hEle))
	return int(r)
}
