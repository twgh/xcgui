package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

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
func XMenu_AddItem(hMenu int, nID int32, pText string, nParentID int32, nFlags xcc.Menu_Item_Flag_) {
	xMenu_AddItem.Call(uintptr(hMenu), uintptr(nID), common.StrPtr(pText), uintptr(nParentID), uintptr(nFlags))
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
func XMenu_AddItemIcon(hMenu int, nID int32, pText string, nParentID int32, hIcon int, nFlags xcc.Menu_Item_Flag_) {
	xMenu_AddItemIcon.Call(uintptr(hMenu), uintptr(nID), common.StrPtr(pText), uintptr(nParentID), uintptr(hIcon), uintptr(nFlags))
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
func XMenu_InsertItem(hMenu int, nID int32, pText string, nFlags xcc.Menu_Item_Flag_, insertID int32) {
	xMenu_InsertItem.Call(uintptr(hMenu), uintptr(nID), common.StrPtr(pText), uintptr(nFlags), uintptr(insertID))
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
func XMenu_InsertItemIcon(hMenu int, nID int32, pText string, hIcon int, nFlags xcc.Menu_Item_Flag_, insertID int32) {
	xMenu_InsertItemIcon.Call(uintptr(hMenu), uintptr(nID), common.StrPtr(pText), uintptr(hIcon), uintptr(nFlags), uintptr(insertID))
}

// 菜单_取第一个子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetFirstChildItem(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetFirstChildItem.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_取末尾子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetEndChildItem(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetEndChildItem.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetPrevSiblingItem(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetPrevSiblingItem.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetNextSiblingItem(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetNextSiblingItem.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_取父项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetParentItem(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetParentItem.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// hMenu: 菜单句柄.
//
// bAuto: 是否自动销毁.
func XMenu_SetAutoDestroy(hMenu int, bAuto bool) {
	xMenu_SetAutoDestroy.Call(uintptr(hMenu), common.BoolPtr(bAuto))
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func XMenu_EnableDrawBackground(hMenu int, bEnable bool) {
	xMenu_EnableDrawBackground.Call(uintptr(hMenu), common.BoolPtr(bEnable))
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func XMenu_EnableDrawItem(hMenu int, bEnable bool) {
	xMenu_EnableDrawItem.Call(uintptr(hMenu), common.BoolPtr(bEnable))
}

// 菜单_弹出.
//
// hMenu: 菜单句柄.
//
// hParentWnd: 父窗口真实句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// hParentEle: 父元素句柄, 如果该值不为NULL, hParentEle元素将接收菜单消息事件, 否则将由hParentWnd窗口接收菜单的消息事件.
//
// nPosition: 弹出位置, Menu_Popup_Position_.
func XMenu_Popup(hMenu int, hParentWnd uintptr, x, y int32, hParentEle int, nPosition xcc.Menu_Popup_Position_) bool {
	r, _, _ := xMenu_Popup.Call(uintptr(hMenu), hParentWnd, uintptr(x), uintptr(y), uintptr(hParentEle), uintptr(nPosition))
	return r != 0
}

// 菜单_销毁.
//
// hMenu: 菜单句柄.
func XMenu_DestroyMenu(hMenu int) {
	xMenu_DestroyMenu.Call(uintptr(hMenu))
}

// 菜单_关闭.
//
// hMenu: 菜单句柄.
func XMenu_CloseMenu(hMenu int) {
	xMenu_CloseMenu.Call(uintptr(hMenu))
}

// 菜单_置背景图片.
//
// hMenu: 菜单句柄.
//
// hImage: 图片句柄.
func XMenu_SetBkImage(hMenu int, hImage int) {
	xMenu_SetBkImage.Call(uintptr(hMenu), uintptr(hImage))
}

// 菜单_置项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
func XMenu_SetItemText(hMenu int, nID int32, pText string) bool {
	r, _, _ := xMenu_SetItemText.Call(uintptr(hMenu), uintptr(nID), common.StrPtr(pText))
	return r != 0
}

// 菜单_取项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetItemText(hMenu int, nID int32) string {
	r, _, _ := xMenu_GetItemText.Call(uintptr(hMenu), uintptr(nID))
	return common.UintPtrToString(r)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func XMenu_GetItemTextLength(hMenu int, nID int32) int32 {
	r, _, _ := xMenu_GetItemTextLength.Call(uintptr(hMenu), uintptr(nID))
	return int32(r)
}

// 菜单_置项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func XMenu_SetItemIcon(hMenu int, nID int32, hIcon int) bool {
	r, _, _ := xMenu_SetItemIcon.Call(uintptr(hMenu), uintptr(nID), uintptr(hIcon))
	return r != 0
}

// 菜单_置项标志.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func XMenu_SetItemFlags(hMenu int, nID int32, uFlags xcc.Menu_Item_Flag_) bool {
	r, _, _ := xMenu_SetItemFlags.Call(uintptr(hMenu), uintptr(nID), uintptr(uFlags))
	return r != 0
}

// 菜单_置项高度.
//
// hMenu: 菜单句柄.
//
// height: 高度.
func XMenu_SetItemHeight(hMenu int, height int32) {
	xMenu_SetItemHeight.Call(uintptr(hMenu), uintptr(height))
}

// 菜单_取项高度.
//
// hMenu: 菜单句柄.
func XMenu_GetItemHeight(hMenu int) int32 {
	r, _, _ := xMenu_GetItemHeight.Call(uintptr(hMenu))
	return int32(r)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// hMenu: 菜单句柄.
//
// crColor: ABGR 颜色.
func XMenu_SetBorderColor(hMenu int, crColor int) {
	xMenu_SetBorderColor.Call(uintptr(hMenu), uintptr(crColor))
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
func XMenu_SetBorderSize(hMenu int, nLeft, nTop, nRight, nBottom int32) {
	xMenu_SetBorderSize.Call(uintptr(hMenu), uintptr(nLeft), uintptr(nTop), uintptr(nRight), uintptr(nBottom))
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
//
// hMenu: 菜单句柄.
func XMenu_GetLeftWidth(hMenu int) int32 {
	r, _, _ := xMenu_GetLeftWidth.Call(uintptr(hMenu))
	return int32(r)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
//
// hMenu: 菜单句柄.
func XMenu_GetLeftSpaceText(hMenu int) int32 {
	r, _, _ := xMenu_GetLeftSpaceText.Call(uintptr(hMenu))
	return int32(r)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
//
// hMenu: 菜单句柄.
func XMenu_GetItemCount(hMenu int) int32 {
	r, _, _ := xMenu_GetItemCount.Call(uintptr(hMenu))
	return int32(r)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func XMenu_SetItemCheck(hMenu int, nID int32, bCheck bool) bool {
	r, _, _ := xMenu_SetItemCheck.Call(uintptr(hMenu), uintptr(nID), common.BoolPtr(bCheck))
	return r != 0
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
func XMenu_IsItemCheck(hMenu int, nID int32) bool {
	r, _, _ := xMenu_IsItemCheck.Call(uintptr(hMenu), uintptr(nID))
	return r != 0
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func XMenu_SetItemWidth(hMenu int, nID, nWidth int32) bool {
	r, _, _ := xMenu_SetItemWidth.Call(uintptr(hMenu), uintptr(nID), uintptr(nWidth))
	return r != 0
}

// 菜单_取菜单条, 返回菜单条句柄.
//
// hMenu: 菜单句柄.
func XMenu_GetMenuBar(hMenu int) int {
	r, _, _ := xMenu_GetMenuBar.Call(uintptr(hMenu))
	return int(r)
}
