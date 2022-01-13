package widget

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// 弹出菜单.
type Menu struct {
	objectbase.ObjectBase
}

// 菜单_创建, 创建菜单, 默认弹出菜单窗口关闭后自动销毁.
func NewMenu() *Menu {
	p := &Menu{}
	p.SetHandle(xc.XMenu_Create())
	return p
}

// 从句柄创建对象.
func NewMenuByHandle(handle int) *Menu {
	p := &Menu{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewMenuByName(name string) *Menu {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewMenuByUID(nUID int) *Menu {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewMenuByUIDName(name string) *Menu {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 菜单_添加项, 添加菜单项.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// nFlags: 标识, Menu_Item_Flag_.
func (m *Menu) AddItem(nID int, pText string, nParentID int, nFlags int) int {
	return xc.XMenu_AddItem(m.Handle, nID, pText, nParentID, nFlags)
}

// 菜单_添加项图标.
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
func (m *Menu) AddItemIcon(nID int, pText string, nParentID int, hIcon int, nFlags int) int {
	return xc.XMenu_AddItemIcon(m.Handle, nID, pText, nParentID, hIcon, nFlags)
}

// 菜单_插入项.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func (m *Menu) InsertItem(nID int, pText string, nFlags int, insertID int) int {
	return xc.XMenu_InsertItem(m.Handle, nID, pText, nFlags, insertID)
}

// 菜单_插入项图标.
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
func (m *Menu) InsertItemIcon(nID int, pText string, hIcon int, nFlags int, insertID int) int {
	return xc.XMenu_InsertItemIcon(m.Handle, nID, pText, hIcon, nFlags, insertID)
}

// 菜单_取第一个子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetFirstChildItem(nID int) int {
	return xc.XMenu_GetFirstChildItem(m.Handle, nID)
}

// 菜单_取末尾子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetEndChildItem(nID int) int {
	return xc.XMenu_GetEndChildItem(m.Handle, nID)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetPrevSiblingItem(nID int) int {
	return xc.XMenu_GetPrevSiblingItem(m.Handle, nID)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetNextSiblingItem(nID int) int {
	return xc.XMenu_GetNextSiblingItem(m.Handle, nID)
}

// 菜单_取父项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetParentItem(nID int) int {
	return xc.XMenu_GetParentItem(m.Handle, nID)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// bAuto: 是否自动销毁.
func (m *Menu) SetAutoDestroy(bAuto bool) int {
	return xc.XMenu_SetAutoDestroy(m.Handle, bAuto)
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// bEnable: 是否启用.
func (m *Menu) EnableDrawBackground(bEnable bool) int {
	return xc.XMenu_EnableDrawBackground(m.Handle, bEnable)
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// bEnable: 是否启用.
func (m *Menu) EnableDrawItem(bEnable bool) int {
	return xc.XMenu_EnableDrawItem(m.Handle, bEnable)
}

// 菜单_弹出.
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
func (m *Menu) Popup(hParentWnd int, x int, y int, hParentEle int, nPosition int) bool {
	return xc.XMenu_Popup(m.Handle, hParentWnd, x, y, hParentEle, nPosition)
}

// 菜单_销毁.
func (m *Menu) DestroyMenu() int {
	return xc.XMenu_DestroyMenu(m.Handle)
}

// 菜单_关闭.
func (m *Menu) CloseMenu() int {
	return xc.XMenu_CloseMenu(m.Handle)
}

// 菜单_置背景图片.
//
// hImage: 图片句柄.
func (m *Menu) SetBkImage(hImage int) int {
	return xc.XMenu_SetBkImage(m.Handle, hImage)
}

// 菜单_置项文本.
//
// nID: 项ID.
//
// pText: 文本内容.
func (m *Menu) SetItemText(nID int, pText string) bool {
	return xc.XMenu_SetItemText(m.Handle, nID, pText)
}

// 菜单_取项文本.
//
// nID: 项ID.
func (m *Menu) GetItemText(nID int) string {
	return xc.XMenu_GetItemText(m.Handle, nID)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// nID: 项ID.
func (m *Menu) GetItemTextLength(nID int) int {
	return xc.XMenu_GetItemTextLength(m.Handle, nID)
}

// 菜单_置项图标.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func (m *Menu) SetItemIcon(nID int, hIcon int) bool {
	return xc.XMenu_SetItemIcon(m.Handle, nID, hIcon)
}

// 菜单_置项标志.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func (m *Menu) SetItemFlags(nID int, uFlags int) bool {
	return xc.XMenu_SetItemFlags(m.Handle, nID, uFlags)
}

// 菜单_置项高度.
//
// height: 高度.
func (m *Menu) SetItemHeight(height int) int {
	return xc.XMenu_SetItemHeight(m.Handle, height)
}

// 菜单_取项高度.
func (m *Menu) GetItemHeight() int {
	return xc.XMenu_GetItemHeight(m.Handle)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// crColor: ABGR颜色.
func (m *Menu) SetBorderColor(crColor int) int {
	return xc.XMenu_SetBorderColor(m.Handle, crColor)
}

// 菜单_置边框大小, 设置弹出菜单窗口边框大小.
//
// nLeft: 边大小.
//
// nTop: 边大小.
//
// nRight: 边大小.
//
// nBottom: 边大小.
func (m *Menu) SetBorderSize(nLeft int, nTop int, nRight int, nBottom int) int {
	return xc.XMenu_SetBorderSize(m.Handle, nLeft, nTop, nRight, nBottom)
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
func (m *Menu) GetLeftWidth() int {
	return xc.XMenu_GetLeftWidth(m.Handle)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
func (m *Menu) GetLeftSpaceText() int {
	return xc.XMenu_GetLeftSpaceText(m.Handle)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
func (m *Menu) GetItemCount() int {
	return xc.XMenu_GetItemCount(m.Handle)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func (m *Menu) SetItemCheck(nID int, bCheck bool) bool {
	return xc.XMenu_SetItemCheck(m.Handle, nID, bCheck)
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// nID: 菜单项ID.
func (m *Menu) IsItemCheck(nID int) bool {
	return xc.XMenu_IsItemCheck(m.Handle, nID)
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func (m *Menu) SetItemWidth(nID int, nWidth int) bool {
	return xc.XMenu_SetItemWidth(m.Handle, nID, nWidth)
}
