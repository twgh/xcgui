package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// TabBar Tab条.
type TabBar struct {
	Element
}

// TAB条_创建, 创建tabBar元素, 返回元素句柄.
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
func NewTabBar(x int, y int, cx int, cy int, hParent int) *TabBar {
	p := &TabBar{}
	p.SetHandle(xc.XTabBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewTabBarByHandle(handle int) *TabBar {
	p := &TabBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewTabBarByName(name string) *TabBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewTabBarByUID(nUID int) *TabBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewTabBarByUIDName(name string) *TabBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// TAB条_添加标签, 添加一个标签, 返回标签索引.
//
// pName: 标签文本内容.
func (t *TabBar) AddLabel(pName string) int {
	return xc.XTabBar_AddLabel(t.Handle, pName)
}

// TAB条插入_标签, 插入一个标签, 返回标签索引.
//
// index: 插入位置.
//
// pName: 标签文本内容.
func (t *TabBar) InsertLabel(index int, pName string) int {
	return xc.XTabBar_InsertLabel(t.Handle, index, pName)
}

// TAB条_移动标签.
//
// iSrc: 源位置索引.
//
// iDest: 目标位置索引.
func (t *TabBar) MoveLabel(iSrc int, iDest int) bool {
	return xc.XTabBar_MoveLabel(t.Handle, iSrc, iDest)
}

// TAB条_删除标签, 删除一个标签.
//
// index: 位置索引.
func (t *TabBar) DeleteLabel(index int) bool {
	return xc.XTabBar_DeleteLabel(t.Handle, index)
}

// TAB条_删除全部, 删除所有标签.
func (t *TabBar) DeleteLabelAll() int {
	return xc.XTabBar_DeleteLabelAll(t.Handle)
}

// TAB条_取标签, 获取标签按钮句柄.
//
// index: 位置索引.
func (t *TabBar) GetLabel(index int) int {
	return xc.XTabBar_GetLabel(t.Handle, index)
}

// TAB条_取标签上的关闭按钮, 获取标签上关闭按钮句柄.
//
// index: 位置索引.
func (t *TabBar) GetLabelClose(index int) int {
	return xc.XTabBar_GetLabelClose(t.Handle, index)
}

// TAB条_取左滚动按钮, 获取左滚动按钮句柄.
func (t *TabBar) GetButtonLeft() int {
	return xc.XTabBar_GetButtonLeft(t.Handle)
}

// TAB条_取右滚动按钮, 获取右滚动按钮句柄.
func (t *TabBar) GetButtonRight() int {
	return xc.XTabBar_GetButtonRight(t.Handle)
}

// TAB条_取下拉菜单按钮句柄.
func (t *TabBar) GetButtonDropMenu() int {
	return xc.XTabBar_GetButtonDropMenu(t.Handle)
}

// TAB条_取当前选择, 获取选择的标签索引.
func (t *TabBar) GetSelect() int {
	return xc.XTabBar_GetSelect(t.Handle)
}

// TAB条_取间隔, 获取标签间距, 0没有间距.
func (t *TabBar) GetLabelSpacing() int {
	return xc.XTabBar_GetLabelSpacing(t.Handle)
}

// TAB条_取标签数量, 获取标签项数量.
func (t *TabBar) GetLabelCount() int {
	return xc.XTabBar_GetLabelCount(t.Handle)
}

// TAB条_取标签位置索引, 获取标签按钮位置索引, 成功返回索引值, 否则返回 XC_ID_ERROR.
//
// hLabel: 标签按钮句柄.
func (t *TabBar) GetindexByEle(hLabel int) int {
	return xc.XTabBar_GetindexByEle(t.Handle, hLabel)
}

// TAB条_置间隔, 设置标签间距, 0没有间距.
//
// spacing: 标签间隔大小.
func (t *TabBar) SetLabelSpacing(spacing int) int {
	return xc.XTabBar_SetLabelSpacing(t.Handle, spacing)
}

// TAB条_置边距, 设置内容与边框的间隔大小.
//
// left: 左边间隔大小.
//
// top: 上边间隔大小.
//
// right: 右边间隔大小.
//
// bottom: 下边间隔大小.
func (t *TabBar) SetPadding(left int, top int, right int, bottom int) int {
	return xc.XTabBar_SetPadding(t.Handle, left, top, right, bottom)
}

// TAB条_置选择, 设置选择标签.
//
// index: 标签位置索引.
func (t *TabBar) SetSelect(index int) int {
	return xc.XTabBar_SetSelect(t.Handle, index)
}

// TAB条_左滚动, 左按钮滚动.
func (t *TabBar) SetUp() int {
	return xc.XTabBar_SetUp(t.Handle)
}

// TAB条_右滚动, 右按钮滚动.
func (t *TabBar) SetDown() int {
	return xc.XTabBar_SetDown(t.Handle)
}

// TAB条_启用平铺, 平铺标签, 每个标签显示相同大小.
//
// bTile: 是否启用.
func (t *TabBar) EnableTile(bTile bool) int {
	return xc.XTabBar_EnableTile(t.Handle, bTile)
}

// TAB条_启用下拉菜单按钮.
//
// bEnable:.
func (t *TabBar) EnableDropMenu(bEnable bool) int {
	return xc.XTabBar_EnableDropMenu(t.Handle, bEnable)
}

// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
//
// bEnable: 是否启用.
func (t *TabBar) EnableClose(bEnable bool) int {
	return xc.XTabBar_EnableClose(t.Handle, bEnable)
}

// TAB条_置关闭按钮大小, 设置关闭按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func (t *TabBar) SetCloseSize(pSize *xc.SIZE) int {
	return xc.XTabBar_SetCloseSize(t.Handle, pSize)
}

// TAB条_置滚动按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func (t *TabBar) SetTurnButtonSize(pSize *xc.SIZE) int {
	return xc.XTabBar_SetTurnButtonSize(t.Handle, pSize)
}

// TAB条_置指定标签固定宽度.
//
// index: 索引.
//
// nWidth: 宽度, , 如果值为-1, 那么自动计算宽度.
func (t *TabBar) SetLabelWidth(index int, nWidth int) int {
	return xc.XTabBar_SetLabelWidth(t.Handle, index, nWidth)
}

// TAB条_显示标签, 显示或隐藏指定标签.
//
// index: 标签索引.
//
// bShow: 是否显示.
func (t *TabBar) ShowLabel(index int, bShow bool) bool {
	return xc.XTabBar_ShowLabel(t.Handle, index, bShow)
}

/*
以下都是事件
*/

type XE_TABBAR_SELECT func(iItem int32, pbHandled *bool) int            // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE func(iItem int32, pbHandled *bool) int            // TabBar标签按钮删除事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮删除事件. iItem: 标签位置索引.

// TabBar标签按钮选择改变事件.
func (t *TabBar) Event_TABBAR_SELECT(pFun XE_TABBAR_SELECT) bool {
	return xc.XEle_RegEventC(t.Handle, xcc.XE_TABBAR_SELECT, pFun)
}

// TabBar标签按钮选择改变事件.
func (t *TabBar) Event_TABBAR_SELECT1(pFun XE_TABBAR_SELECT1) bool {
	return xc.XEle_RegEventC1(t.Handle, xcc.XE_TABBAR_SELECT, pFun)
}

// TabBar标签按钮删除事件.
func (t *TabBar) Event_TABBAR_DELETE(pFun XE_TABBAR_DELETE) bool {
	return xc.XEle_RegEventC(t.Handle, xcc.XE_TABBAR_DELETE, pFun)
}

// TabBar标签按钮删除事件.
func (t *TabBar) Event_TABBAR_DELETE1(pFun XE_TABBAR_DELETE1) bool {
	return xc.XEle_RegEventC1(t.Handle, xcc.XE_TABBAR_DELETE, pFun)
}
