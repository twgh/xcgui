package widget

import "github.com/twgh/xcgui/xc"

// 工具条.
type ToolBar struct {
	Element
}

// 工具条_创建, 创建工具条元素; 如果指定了父为窗口, 默认调用XWnd_AddToolBar()函数, 将工具条添加到窗口非客户区.
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
func NewToolBar(x, y, cx, cy int32, hParent int) *ToolBar {
	p := &ToolBar{}
	p.SetHandle(xc.XToolBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewToolBarByHandle(handle int) *ToolBar {
	p := &ToolBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewToolBarByName(name string) *ToolBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ToolBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewToolBarByUID(nUID int) *ToolBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ToolBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewToolBarByUIDName(name string) *ToolBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ToolBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 工具条_插入元素, 插入元素到工具条, 返回插入位置索引.
//
// hNewEle: 将要插入的元素.
//
// index: 插入位置索引, (-1)插入末尾.
func (t *ToolBar) InsertEle(hNewEle int, index int) int {
	return xc.XToolBar_InsertEle(t.Handle, hNewEle, index)
}

// 工具条_插入分割栏, 插入分隔符到工具条, 返回插入位置索引.
//
// index: 插入位置索引, (-1)插入末尾.
//
// color: ABGR 颜色.
func (t *ToolBar) InsertSeparator(index int, color int) int {
	return xc.XToolBar_InsertSeparator(t.Handle, index, color)
}

// 工具条_启用下拉菜单, 启用下拉菜单, 显示隐藏的项.
//
// bEnable: 是否启用.
func (t *ToolBar) EnableButtonMenu(bEnable bool) int {
	return xc.XToolBar_EnableButtonMenu(t.Handle, bEnable)
}

// 工具条_取元素, 获取工具条上指定元素, 返回元素句柄.
//
// index: 索引值.
func (t *ToolBar) GetEle(index int) int {
	return xc.XToolBar_GetEle(t.Handle, index)
}

// 工具条_取左滚动按钮, 获取左滚动按钮句柄.
func (t *ToolBar) GetButtonLeft() int {
	return xc.XToolBar_GetButtonLeft(t.Handle)
}

// 工具条_取右滚动按钮, 获取右滚动按钮句柄.
func (t *ToolBar) GetButtonRight() int {
	return xc.XToolBar_GetButtonRight(t.Handle)
}

// 工具条_取菜单按钮, 获取菜单按钮句柄.
func (t *ToolBar) GetButtonMenu() int {
	return xc.XToolBar_GetButtonMenu(t.Handle)
}

// 工具条_置间距, 设置对象之间的间距.
//
// nSize: 间距大小.
func (t *ToolBar) SetSpace(nSize int) int {
	return xc.XToolBar_SetSpace(t.Handle, nSize)
}

// 工具条_删除元素, 删除元素, 并且销毁.
//
// index: 索引值.
func (t *ToolBar) DeleteEle(index int) int {
	return xc.XToolBar_DeleteEle(t.Handle, index)
}

// 工具条_删除全部, 删除所有元素, 并且销毁.
func (t *ToolBar) DeleteAllEle() int {
	return xc.XToolBar_DeleteAllEle(t.Handle)
}
