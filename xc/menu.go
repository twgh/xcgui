package xc

// 菜单_创建, 创建菜单, 默认弹出菜单窗口关闭后自动销毁.
func XMenu_Create() int {
	r, _, _ := xMenu_Create.Call()
	return int(r)
}

// 菜单_添加项, 添加菜单项.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// nFlags: 标识, Menu_Item_Flag_.
func XMenu_AddItem(hMenu int, nID int, pText string, nParentID int, nFlags int) int {
	r, _, _ := xMenu_AddItem.Call(uintptr(hMenu), uintptr(nID), strPtr(pText), uintptr(nParentID), uintptr(nFlags))
	return int(r)
}

// 菜单_添加项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, Menu_Item_Flag_.
func XMenu_AddItemIcon(hMenu int, nID int, pText string, nParentID int, hIcon int, nFlags int) int {
	r, _, _ := xMenu_AddItemIcon.Call(uintptr(hMenu), uintptr(nID), strPtr(pText), uintptr(nParentID), uintptr(hIcon), uintptr(nFlags))
	return int(r)
}

// 菜单_插入项.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func XMenu_InsertItem(hMenu int, nID int, pText string, nFlags int, insertID int) int {
	r, _, _ := xMenu_InsertItem.Call(uintptr(hMenu), uintptr(nID), strPtr(pText), uintptr(nFlags), uintptr(insertID))
	return int(r)
}

// 菜单_插入项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func XMenu_InsertItemIcon(hMenu int, nID int, pText string, hIcon int, nFlags int, insertID int) int {
	r, _, _ := xMenu_InsertItemIcon.Call(uintptr(hMenu), uintptr(nID), strPtr(pText), uintptr(hIcon), uintptr(nFlags), uintptr(insertID))
	return int(r)
}

// 菜单_取第一个子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetFirstChildItem(hMenu int, nID int) int {
	r, _, _ := xMenu_GetFirstChildItem.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_取末尾子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetEndChildItem(hMenu int, nID int) int {
	r, _, _ := xMenu_GetEndChildItem.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetPrevSiblingItem(hMenu int, nID int) int {
	r, _, _ := xMenu_GetPrevSiblingItem.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetNextSiblingItem(hMenu int, nID int) int {
	r, _, _ := xMenu_GetNextSiblingItem.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_取父项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetParentItem(hMenu int, nID int) int {
	r, _, _ := xMenu_GetParentItem.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// hMenu: 菜单句柄.
//
// bAuto: 是否自动销毁.
func XMenu_SetAutoDestroy(hMenu int, bAuto bool) int {
	r, _, _ := xMenu_SetAutoDestroy.Call(uintptr(hMenu), boolPtr(bAuto))
	return int(r)
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func XMenu_EnableDrawBackground(hMenu int, bEnable bool) int {
	r, _, _ := xMenu_EnableDrawBackground.Call(uintptr(hMenu), boolPtr(bEnable))
	return int(r)
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func XMenu_EnableDrawItem(hMenu int, bEnable bool) int {
	r, _, _ := xMenu_EnableDrawItem.Call(uintptr(hMenu), boolPtr(bEnable))
	return int(r)
}

// 菜单_弹出.
//
// hMenu: 菜单句柄.
//
// hParentWnd: 父窗口句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// hParentEle: 父元素句柄, 如果该值不为NULL, hParentEle元素将接收菜单消息事件, 否则将由hParentWnd窗口接收菜单的消息事件.
//
// nPosition: 弹出位置, Menu_Popup_Position_.
func XMenu_Popup(hMenu int, hParentWnd int, x int, y int, hParentEle int, nPosition int) bool {
	r, _, _ := xMenu_Popup.Call(uintptr(hMenu), uintptr(hParentWnd), uintptr(x), uintptr(y), uintptr(hParentEle), uintptr(nPosition))
	return int(r) != 0
}

// 菜单_销毁.
//
// hMenu: 菜单句柄.
func XMenu_DestroyMenu(hMenu int) int {
	r, _, _ := xMenu_DestroyMenu.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_关闭.
//
// hMenu: 菜单句柄.
func XMenu_CloseMenu(hMenu int) int {
	r, _, _ := xMenu_CloseMenu.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_置背景图片.
//
// hMenu: 菜单句柄.
//
// hImage: 图片句柄.
func XMenu_SetBkImage(hMenu int, hImage int) int {
	r, _, _ := xMenu_SetBkImage.Call(uintptr(hMenu), uintptr(hImage))
	return int(r)
}

// 菜单_置项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
func XMenu_SetItemText(hMenu int, nID int, pText string) bool {
	r, _, _ := xMenu_SetItemText.Call(uintptr(hMenu), uintptr(nID), strPtr(pText))
	return int(r) != 0
}

// 菜单_取项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetItemText(hMenu int, nID int) string {
	r, _, _ := xMenu_GetItemText.Call(uintptr(hMenu), uintptr(nID))
	return UintPtrToString(r)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetItemTextLength(hMenu int, nID int) int {
	r, _, _ := xMenu_GetItemTextLength.Call(uintptr(hMenu), uintptr(nID))
	return int(r)
}

// 菜单_置项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func XMenu_SetItemIcon(hMenu int, nID int, hIcon int) bool {
	r, _, _ := xMenu_SetItemIcon.Call(uintptr(hMenu), uintptr(nID), uintptr(hIcon))
	return int(r) != 0
}

// 菜单_置项标志.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func XMenu_SetItemFlags(hMenu int, nID int, uFlags int) bool {
	r, _, _ := xMenu_SetItemFlags.Call(uintptr(hMenu), uintptr(nID), uintptr(uFlags))
	return int(r) != 0
}

// 菜单_置项高度.
//
// hMenu: 菜单句柄.
//
// height: 高度.
func XMenu_SetItemHeight(hMenu int, height int) int {
	r, _, _ := xMenu_SetItemHeight.Call(uintptr(hMenu), uintptr(height))
	return int(r)
}

// 菜单_取项高度.
//
// hMenu: 菜单句柄.
func XMenu_GetItemHeight(hMenu int) int {
	r, _, _ := xMenu_GetItemHeight.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// hMenu: 菜单句柄.
//
// crColor: ABGR颜色.
func XMenu_SetBorderColor(hMenu int, crColor int) int {
	r, _, _ := xMenu_SetBorderColor.Call(uintptr(hMenu), uintptr(crColor))
	return int(r)
}

// 菜单_置边框大小, 设置弹出菜单窗口边框大小.
//
// hMenu: 菜单句柄.
//
// nLeft: 边大小.
//
// nTop: 边大小.
//
// nRight: 边大小.
//
// nBottom: 边大小.
func XMenu_SetBorderSize(hMenu int, nLeft int, nTop int, nRight int, nBottom int) int {
	r, _, _ := xMenu_SetBorderSize.Call(uintptr(hMenu), uintptr(nLeft), uintptr(nTop), uintptr(nRight), uintptr(nBottom))
	return int(r)
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
//
// hMenu: 菜单句柄.
func XMenu_GetLeftWidth(hMenu int) int {
	r, _, _ := xMenu_GetLeftWidth.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
//
// hMenu: 菜单句柄.
func XMenu_GetLeftSpaceText(hMenu int) int {
	r, _, _ := xMenu_GetLeftSpaceText.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
//
// hMenu: 菜单句柄.
func XMenu_GetItemCount(hMenu int) int {
	r, _, _ := xMenu_GetItemCount.Call(uintptr(hMenu))
	return int(r)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func XMenu_SetItemCheck(hMenu int, nID int, bCheck bool) bool {
	r, _, _ := xMenu_SetItemCheck.Call(uintptr(hMenu), uintptr(nID), boolPtr(bCheck))
	return int(r) != 0
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
func XMenu_IsItemCheck(hMenu int, nID int) bool {
	r, _, _ := xMenu_IsItemCheck.Call(uintptr(hMenu), uintptr(nID))
	return int(r) != 0
}
