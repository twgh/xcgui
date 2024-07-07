package xc

import "github.com/twgh/xcgui/common"

// 菜单条_创建, 创建菜单条元素; 如果指定了父为窗口, 默认调用XWnd_AddMenuBar()函数, 将菜单条添加到窗口非客户区. 返回元素句柄.
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
func XMenuBar_Create(x, y, cx, cy int32, hParent int) int {
	r, _, _ := xMenuBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 菜单条_添加按钮, 添加弹出菜单按钮, 返回菜单按钮索引.
//
// hEle: 元素句柄.
//
// pText: 文本内容.
func XMenuBar_AddButton(hEle int, pText string) int32 {
	r, _, _ := xMenuBar_AddButton.Call(uintptr(hEle), common.StrPtr(pText))
	return int32(r)
}

// 菜单条_置按钮高度, 根据内容自动调整宽度. (已废弃)请使用内填充限制高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func XMenuBar_SetButtonHeight(hEle int, height int32) {
	xMenuBar_SetButtonHeight.Call(uintptr(hEle), uintptr(height))
}

// 菜单条_取菜单, 返回菜单句柄.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条上菜单按钮的索引.
func XMenuBar_GetMenu(hEle int, nIndex int32) int {
	r, _, _ := xMenuBar_GetMenu.Call(uintptr(hEle), uintptr(nIndex))
	return int(r)
}

// 菜单条_删除按钮, 删除菜单条上的菜单按钮, 同时该按钮下的弹出菜单也被销毁.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条按钮索引.
func XMenuBar_DeleteButton(hEle int, nIndex int32) bool {
	r, _, _ := xMenuBar_DeleteButton.Call(uintptr(hEle), uintptr(nIndex))
	return r != 0
}

// 菜单条_启用自动宽度, 根据内容自动调整宽度.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XMenuBar_EnableAutoWidth(hEle int, bEnable bool) {
	xMenuBar_EnableAutoWidth.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 菜单条_取菜单按钮. 返回按钮句柄.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条按钮索引.
func XMenuBar_GetButton(hEle int, nIndex int32) bool {
	r, _, _ := xMenuBar_GetButton.Call(uintptr(hEle), uintptr(nIndex))
	return r != 0
}

// 菜单条_取选择项. 菜单条当前选择项, 当前弹出菜单的按钮索引.
//
// hEle: 元素句柄.
func XMenuBar_GetSelect(hEle int) int32 {
	r, _, _ := xMenuBar_GetSelect.Call(uintptr(hEle))
	return int32(r)
}
