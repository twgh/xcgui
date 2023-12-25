package widget

import "github.com/twgh/xcgui/xc"

// MenuBar 菜单条.
type MenuBar struct {
	Element
}

// 菜单条_创建, 创建菜单条元素; 如果指定了父为窗口, 默认调用XWnd_AddMenuBar()函数, 将菜单条添加到窗口非客户区.
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
func NewMenuBar(x int32, y int32, cx int32, cy int32, hParent int) *MenuBar {
	p := &MenuBar{}
	p.SetHandle(xc.XMenuBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewMenuBarByHandle(handle int) *MenuBar {
	p := &MenuBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewMenuBarByName(name string) *MenuBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewMenuBarByUID(nUID int) *MenuBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewMenuBarByUIDName(name string) *MenuBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 菜单条_添加按钮, 添加弹出菜单按钮, 返回菜单按钮索引.
//
// pText: 文本内容.
func (m *MenuBar) AddButton(pText string) int32 {
	return xc.XMenuBar_AddButton(m.Handle, pText)
}

// 菜单条_置按钮高度, 根据内容自动调整宽度.
//
// height: 高度.
func (m *MenuBar) SetButtonHeight(height int32) {
	xc.XMenuBar_SetButtonHeight(m.Handle, height)
}

// 菜单条_取菜单, 返回菜单句柄.
//
// nIndex: 菜单条上菜单按钮的索引.
func (m *MenuBar) GetMenu(nIndex int32) int {
	return xc.XMenuBar_GetMenu(m.Handle, nIndex)
}

// 菜单条_删除按钮, 删除菜单条上的菜单按钮, 同时该按钮下的弹出菜单也被销毁.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) DeleteButton(nIndex int32) bool {
	return xc.XMenuBar_DeleteButton(m.Handle, nIndex)
}

// 菜单条_启用自动宽度, 根据内容自动调整宽度. (已废弃)请使用内填充限制高度
//
// bEnable: 是否启用.
func (m *MenuBar) EnableAutoWidth(bEnable bool) {
	xc.XMenuBar_EnableAutoWidth(m.Handle, bEnable)
}

// 菜单条_取菜单按钮. 返回按钮句柄.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) GetButton(nIndex int32) bool {
	return xc.XMenuBar_GetButton(m.Handle, nIndex)
}

// 菜单条_取选择项. 菜单条当前选择项, 当前弹出菜单的按钮索引.
func (m *MenuBar) GetSelect() int32 {
	return xc.XMenuBar_GetSelect(m.Handle)
}
