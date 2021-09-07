package xc

// 菜单条_创建
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func XMenuBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xMenuBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 菜单条_添加按钮
// hEle: 元素句柄.
// pText: 文本内容.
func XMenuBar_AddButton(hEle int, pText string) int {
	r, _, _ := xMenuBar_AddButton.Call(uintptr(hEle), strPtr(pText))
	return int(r)
}

// 菜单条_置按钮高度
// hEle: 元素句柄.
// height: 高度.
func XMenuBar_SetButtonHeight(hEle int, height int) int {
	r, _, _ := xMenuBar_SetButtonHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 菜单条_取菜单
// hEle: 元素句柄.
// nIndex: 菜单条上菜单按钮的索引.
func XMenuBar_GetMenu(hEle int, nIndex int) int {
	r, _, _ := xMenuBar_GetMenu.Call(uintptr(hEle), uintptr(nIndex))
	return int(r)
}

// 菜单条_删除按钮
// hEle: 元素句柄.
// nIndex: 菜单条按钮索引.
func XMenuBar_DeleteButton(hEle int, nIndex int) bool {
	r, _, _ := xMenuBar_DeleteButton.Call(uintptr(hEle), uintptr(nIndex))
	return int(r) != 0
}

// 菜单条_启用自动宽度
// hEle: 元素句柄.
// bEnable: 是否启用.
func XMenuBar_EnableAutoWidth(hEle int, bEnable bool) int {
	r, _, _ := xMenuBar_EnableAutoWidth.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}
