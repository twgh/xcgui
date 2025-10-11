package widget

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Menu 弹出菜单.
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
func NewMenuByUID(nUID int32) *Menu {
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
func (m *Menu) AddItem(nID int32, pText string, nParentID int32, nFlags xcc.Menu_Item_Flag_) *Menu {
	xc.XMenu_AddItem(m.Handle, nID, pText, nParentID, nFlags)
	return m
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
func (m *Menu) AddItemIcon(nID int32, pText string, nParentID int32, hIcon int, nFlags xcc.Menu_Item_Flag_) *Menu {
	xc.XMenu_AddItemIcon(m.Handle, nID, pText, nParentID, hIcon, nFlags)
	return m
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
func (m *Menu) InsertItem(nID int32, pText string, nFlags xcc.Menu_Item_Flag_, insertID int32) *Menu {
	xc.XMenu_InsertItem(m.Handle, nID, pText, nFlags, insertID)
	return m
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
func (m *Menu) InsertItemIcon(nID int32, pText string, hIcon int, nFlags xcc.Menu_Item_Flag_, insertID int32) *Menu {
	xc.XMenu_InsertItemIcon(m.Handle, nID, pText, hIcon, nFlags, insertID)
	return m
}

// 菜单_取第一个子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetFirstChildItem(nID int32) int32 {
	return xc.XMenu_GetFirstChildItem(m.Handle, nID)
}

// 菜单_取末尾子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetEndChildItem(nID int32) int32 {
	return xc.XMenu_GetEndChildItem(m.Handle, nID)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetPrevSiblingItem(nID int32) int32 {
	return xc.XMenu_GetPrevSiblingItem(m.Handle, nID)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetNextSiblingItem(nID int32) int32 {
	return xc.XMenu_GetNextSiblingItem(m.Handle, nID)
}

// 菜单_取父项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) GetParentItem(nID int32) int32 {
	return xc.XMenu_GetParentItem(m.Handle, nID)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// bAuto: 是否自动销毁.
func (m *Menu) SetAutoDestroy(bAuto bool) *Menu {
	xc.XMenu_SetAutoDestroy(m.Handle, bAuto)
	return m
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// bEnable: 是否启用.
func (m *Menu) EnableDrawBackground(bEnable bool) *Menu {
	xc.XMenu_EnableDrawBackground(m.Handle, bEnable)
	return m
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// bEnable: 是否启用.
func (m *Menu) EnableDrawItem(bEnable bool) *Menu {
	xc.XMenu_EnableDrawItem(m.Handle, bEnable)
	return m
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
func (m *Menu) Popup(hParentWnd uintptr, x, y int32, hParentEle int, nPosition xcc.Menu_Popup_Position_) bool {
	return xc.XMenu_Popup(m.Handle, hParentWnd, x, y, hParentEle, nPosition)
}

// 菜单_销毁.
func (m *Menu) DestroyMenu() *Menu {
	xc.XMenu_DestroyMenu(m.Handle)
	return m
}

// 菜单_关闭.
func (m *Menu) CloseMenu() *Menu {
	xc.XMenu_CloseMenu(m.Handle)
	return m
}

// 菜单_置背景图片.
//
// hImage: 图片句柄.
func (m *Menu) SetBkImage(hImage int) *Menu {
	xc.XMenu_SetBkImage(m.Handle, hImage)
	return m
}

// 菜单_置项文本.
//
// nID: 项ID.
//
// pText: 文本内容.
func (m *Menu) SetItemText(nID int32, pText string) bool {
	return xc.XMenu_SetItemText(m.Handle, nID, pText)
}

// 菜单_取项文本.
//
// nID: 项ID.
func (m *Menu) GetItemText(nID int32) string {
	return xc.XMenu_GetItemText(m.Handle, nID)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// nID: 项ID.
func (m *Menu) GetItemTextLength(nID int32) int32 {
	return xc.XMenu_GetItemTextLength(m.Handle, nID)
}

// 菜单_置项图标.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func (m *Menu) SetItemIcon(nID int32, hIcon int) bool {
	return xc.XMenu_SetItemIcon(m.Handle, nID, hIcon)
}

// 菜单_置项标志.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func (m *Menu) SetItemFlags(nID int32, uFlags xcc.Menu_Item_Flag_) bool {
	return xc.XMenu_SetItemFlags(m.Handle, nID, uFlags)
}

// 菜单_置项高度.
//
// height: 高度.
func (m *Menu) SetItemHeight(height int32) *Menu {
	xc.XMenu_SetItemHeight(m.Handle, height)
	return m
}

// 菜单_取项高度.
func (m *Menu) GetItemHeight() int32 {
	return xc.XMenu_GetItemHeight(m.Handle)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// crColor: xc.RGBA 颜色.
func (m *Menu) SetBorderColor(crColor uint32) *Menu {
	xc.XMenu_SetBorderColor(m.Handle, crColor)
	return m
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
func (m *Menu) SetBorderSize(nLeft, nTop, nRight, nBottom int32) *Menu {
	xc.XMenu_SetBorderSize(m.Handle, nLeft, nTop, nRight, nBottom)
	return m
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
func (m *Menu) GetLeftWidth() int32 {
	return xc.XMenu_GetLeftWidth(m.Handle)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
func (m *Menu) GetLeftSpaceText() int32 {
	return xc.XMenu_GetLeftSpaceText(m.Handle)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
func (m *Menu) GetItemCount() int32 {
	return xc.XMenu_GetItemCount(m.Handle)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func (m *Menu) SetItemCheck(nID int32, bCheck bool) bool {
	return xc.XMenu_SetItemCheck(m.Handle, nID, bCheck)
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// nID: 菜单项ID.
func (m *Menu) IsItemCheck(nID int32) bool {
	return xc.XMenu_IsItemCheck(m.Handle, nID)
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func (m *Menu) SetItemWidth(nID, nWidth int32) bool {
	return xc.XMenu_SetItemWidth(m.Handle, nID, nWidth)
}

// 菜单_取菜单条, 返回菜单条句柄.
func (m *Menu) GetMenuBar() int {
	return xc.XMenu_GetMenuBar(m.Handle)
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_Menu_Select 添加弹出菜单项选择事件.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_Select(hWindowOrhEle int, pFun func(hWindowOrhEle int, nID int32, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_SELECT, xc.OnXWM_MENU_SELECT, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_SELECT, onXE_MENU_SELECT, pFun, allowAddingMultiple...)
	}
	return ret
}

// AddEvent_Menu_Popup 添加菜单弹出事件.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_Popup(hWindowOrhEle int, pFun func(hWindowOrhEle int, HMENUX int, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_POPUP, xc.OnXWM_MENU_POPUP, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_POPUP, onXE_MENU_POPUP, pFun, allowAddingMultiple...)
	}
	return ret
}

// AddEvent_Menu_Exit 添加菜单退出事件.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_Exit(hWindowOrhEle int, pFun func(hWindowOrhEle int, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_EXIT, xc.OnXWM_MENU_EXIT, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_EXIT, onXE_MENU_EXIT, pFun, allowAddingMultiple...)
	}
	return ret
}

// AddEvent_Menu_Popup_Wnd 添加菜单弹出窗口事件.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_Popup_Wnd(hWindowOrhEle int, pFun func(hWindowOrhEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_POPUP_WND, xc.OnXWM_MENU_POPUP_WND, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_POPUP_WND, onXE_MENU_POPUP_WND, pFun, allowAddingMultiple...)
	}
	return ret
}

// AddEvent_Menu_Draw_Background 添加菜单绘制背景事件. 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_Draw_Background(hWindowOrhEle int, pFun func(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_DRAW_BACKGROUND, xc.OnXWM_MENU_DRAW_BACKGROUND, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_DRAW_BACKGROUND, onXE_MENU_DRAW_BACKGROUND, pFun, allowAddingMultiple...)
	}
	return ret
}

// AddEvent_Menu_DrawItem 添加菜单项绘制事件. 启用该功能需要调用 xc.XMenu_EnableDrawItem.
//
// hWindowOrhEle: 炫彩元素或窗口句柄.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (m *Menu) AddEvent_Menu_DrawItem(hWindowOrhEle int, pFun func(hWindowOrhEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	ret := -1
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.WndEventHandler.AddCallBack(hWindowOrhEle, xcc.XWM_MENU_DRAWITEM, xc.OnXWM_MENU_DRAWITEM, pFun, allowAddingMultiple...)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.EleEventHandler.AddCallBack(hWindowOrhEle, xcc.XE_MENU_DRAWITEM, onXE_MENU_DRAWITEM, pFun, allowAddingMultiple...)
	}
	return ret
}

// ------------------------------------- 事件 ------------------------------------- //

// 事件_弹出菜单项被选择.
func (m *Menu) Event_MENU_SELECT(hWindowOrhEle int, pFun xc.XE_MENU_SELECT) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_SELECT, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_SELECT, pFun)
	}
	return ret
}

// 事件_菜单弹出.
func (m *Menu) Event_MENU_POPUP(hWindowOrhEle int, pFun xc.XE_MENU_POPUP) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_POPUP, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_POPUP, pFun)
	}
	return ret
}

// 事件_菜单退出.
func (m *Menu) Event_MENU_EXIT(hWindowOrhEle int, pFun xc.XE_MENU_EXIT) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_EXIT, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_EXIT, pFun)
	}
	return ret
}

// 菜单弹出窗口.
func (m *Menu) Event_MENU_POPUP_WND(hWindowOrhEle int, pFun xc.XE_MENU_POPUP_WND) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_POPUP_WND, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_POPUP_WND, pFun)
	}
	return ret
}

// 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
func (m *Menu) Event_MENU_DRAW_BACKGROUND(hWindowOrhEle int, pFun xc.XE_MENU_DRAW_BACKGROUND) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
	}
	return ret
}

// 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.
func (m *Menu) Event_MENU_DRAWITEM(hWindowOrhEle int, pFun xc.XE_MENU_DRAWITEM) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC(hWindowOrhEle, xcc.XWM_MENU_DRAWITEM, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC(hWindowOrhEle, xcc.XE_MENU_DRAWITEM, pFun)
	}
	return ret
}

// 事件_弹出菜单项被选择.
func (m *Menu) Event_MENU_SELECT1(hWindowOrhEle int, pFun func(hWindowOrhEle int, nID int32, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_SELECT, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_SELECT, pFun)
	}
	return ret
}

// 事件_菜单弹出.
func (m *Menu) Event_MENU_POPUP1(hWindowOrhEle int, pFun func(hWindowOrhEle int, HMENUX int, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_POPUP, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_POPUP, pFun)
	}
	return ret
}

// 事件_菜单退出.
func (m *Menu) Event_MENU_EXIT1(hWindowOrhEle int, pFun func(hWindowOrhEle int, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_EXIT, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_EXIT, pFun)
	}
	return ret
}

// 菜单弹出窗口.
func (m *Menu) Event_MENU_POPUP_WND1(hWindowOrhEle int, pFun func(hWindowOrhEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_POPUP_WND, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_POPUP_WND, pFun)
	}
	return ret
}

// 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
func (m *Menu) Event_MENU_DRAW_BACKGROUND1(hWindowOrhEle int, pFun func(hWindowOrhEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
	}
	return ret
}

// 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.
func (m *Menu) Event_MENU_DRAWITEM1(hWindowOrhEle int, pFun func(hWindowOrhEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int) bool {
	ret := false
	if xc.XC_IsHWINDOW(hWindowOrhEle) {
		ret = xc.XWnd_RegEventC1(hWindowOrhEle, xcc.XWM_MENU_DRAWITEM, pFun)
	} else if xc.XC_IsHELE(hWindowOrhEle) {
		ret = xc.XEle_RegEventC1(hWindowOrhEle, xcc.XE_MENU_DRAWITEM, pFun)
	}
	return ret
}
