package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

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
func XTabBar_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xTabBar_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// TAB条_添加标签, 添加一个标签, 返回标签索引.
//
// hEle: 元素句柄.
//
// pName: 标签文本内容.
func XTabBar_AddLabel(hEle int, pName string) int {
	r, _, _ := xTabBar_AddLabel.Call(uintptr(hEle), common.StrPtr(pName))
	return int(r)
}

// TAB条插入_标签, 插入一个标签, 返回标签索引.
//
// hEle: 元素句柄.
//
// index: 插入位置.
//
// pName: 标签文本内容.
func XTabBar_InsertLabel(hEle int, index int, pName string) int {
	r, _, _ := xTabBar_InsertLabel.Call(uintptr(hEle), uintptr(index), common.StrPtr(pName))
	return int(r)
}

// TAB条_移动标签.
//
// hEle: 元素句柄.
//
// iSrc: 源位置索引.
//
// iDest: 目标位置索引.
func XTabBar_MoveLabel(hEle int, iSrc int, iDest int) bool {
	r, _, _ := xTabBar_MoveLabel.Call(uintptr(hEle), uintptr(iSrc), uintptr(iDest))
	return r != 0
}

// TAB条_删除标签, 删除一个标签.
//
// hEle: 元素句柄.
//
// index: 位置索引.
func XTabBar_DeleteLabel(hEle int, index int) bool {
	r, _, _ := xTabBar_DeleteLabel.Call(uintptr(hEle), uintptr(index))
	return r != 0
}

// TAB条_删除全部, 删除所有标签.
//
// hEle: 元素句柄.
func XTabBar_DeleteLabelAll(hEle int) int {
	r, _, _ := xTabBar_DeleteLabelAll.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取标签, 获取标签按钮句柄.
//
// hEle: 元素句柄.
//
// index: 位置索引.
func XTabBar_GetLabel(hEle int, index int) int {
	r, _, _ := xTabBar_GetLabel.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// TAB条_取标签上的关闭按钮, 获取标签上关闭按钮句柄.
//
// hEle: 元素句柄.
//
// index: 位置索引.
func XTabBar_GetLabelClose(hEle int, index int) int {
	r, _, _ := xTabBar_GetLabelClose.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// TAB条_取左滚动按钮, 获取左滚动按钮句柄.
//
// hEle: 元素句柄.
func XTabBar_GetButtonLeft(hEle int) int {
	r, _, _ := xTabBar_GetButtonLeft.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取右滚动按钮, 获取右滚动按钮句柄.
//
// hEle: 元素句柄.
func XTabBar_GetButtonRight(hEle int) int {
	r, _, _ := xTabBar_GetButtonRight.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取下拉菜单按钮句柄.
//
// hEle: 元素句柄.
func XTabBar_GetButtonDropMenu(hEle int) int {
	r, _, _ := xTabBar_GetButtonDropMenu.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取当前选择, 获取选择的标签索引.
//
// hEle: 元素句柄.
func XTabBar_GetSelect(hEle int) int {
	r, _, _ := xTabBar_GetSelect.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取间隔, 获取标签间距, 0没有间距.
//
// hEle: 元素句柄.
func XTabBar_GetLabelSpacing(hEle int) int {
	r, _, _ := xTabBar_GetLabelSpacing.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取标签数量, 获取标签项数量.
//
// hEle: 元素句柄.
func XTabBar_GetLabelCount(hEle int) int {
	r, _, _ := xTabBar_GetLabelCount.Call(uintptr(hEle))
	return int(r)
}

// TAB条_取标签位置索引, 获取标签按钮位置索引, 成功返回索引值, 否则返回 XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hLabel: 标签按钮句柄.
func XTabBar_GetindexByEle(hEle int, hLabel int) int {
	r, _, _ := xTabBar_GetindexByEle.Call(uintptr(hEle), uintptr(hLabel))
	return int(r)
}

// TAB条_置间隔, 设置标签间距, 0没有间距.
//
// hEle: 元素句柄.
//
// spacing: 标签间隔大小.
func XTabBar_SetLabelSpacing(hEle int, spacing int) int {
	r, _, _ := xTabBar_SetLabelSpacing.Call(uintptr(hEle), uintptr(spacing))
	return int(r)
}

// TAB条_置边距, 设置内容与边框的间隔大小.
//
// hEle: 元素句柄.
//
// left: 左边间隔大小.
//
// top: 上边间隔大小.
//
// right: 右边间隔大小.
//
// bottom: 下边间隔大小.
func XTabBar_SetPadding(hEle int, left int, top int, right int, bottom int) int {
	r, _, _ := xTabBar_SetPadding.Call(uintptr(hEle), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// TAB条_置选择, 设置选择标签.
//
// hEle: 元素句柄.
//
// index: 标签位置索引.
func XTabBar_SetSelect(hEle int, index int) int {
	r, _, _ := xTabBar_SetSelect.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// TAB条_左滚动, 左按钮滚动.
//
// hEle: 元素句柄.
func XTabBar_SetUp(hEle int) int {
	r, _, _ := xTabBar_SetUp.Call(uintptr(hEle))
	return int(r)
}

// TAB条_右滚动, 右按钮滚动.
//
// hEle: 元素句柄.
func XTabBar_SetDown(hEle int) int {
	r, _, _ := xTabBar_SetDown.Call(uintptr(hEle))
	return int(r)
}

// TAB条_启用平铺, 平铺标签, 每个标签显示相同大小.
//
// hEle: 元素句柄.
//
// bTile: 是否启用.
func XTabBar_EnableTile(hEle int, bTile bool) int {
	r, _, _ := xTabBar_EnableTile.Call(uintptr(hEle), common.BoolPtr(bTile))
	return int(r)
}

// TAB条_启用下拉菜单按钮.
//
// hEle: 元素句柄.
//
// bEnable:.
func XTabBar_EnableDropMenu(hEle int, bEnable bool) int {
	r, _, _ := xTabBar_EnableDropMenu.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XTabBar_EnableClose(hEle int, bEnable bool) int {
	r, _, _ := xTabBar_EnableClose.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// TAB条_置关闭按钮大小, 设置关闭按钮大小.
//
// hEle: 元素句柄.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func XTabBar_SetCloseSize(hEle int, pSize *SIZE) int {
	r, _, _ := xTabBar_SetCloseSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// TAB条_置滚动按钮大小.
//
// hEle: 元素句柄.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func XTabBar_SetTurnButtonSize(hEle int, pSize *SIZE) int {
	r, _, _ := xTabBar_SetTurnButtonSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// TAB条_置指定标签固定宽度.
//
// hEle: 元素句柄.
//
// index: 索引.
//
// nWidth: 宽度, , 如果值为-1, 那么自动计算宽度.
func XTabBar_SetLabelWidth(hEle int, index int, nWidth int) int {
	r, _, _ := xTabBar_SetLabelWidth.Call(uintptr(hEle), uintptr(index), uintptr(nWidth))
	return int(r)
}

// TAB条_显示标签, 显示或隐藏指定标签.
//
// hEle: 元素句柄.
//
// index: 标签索引.
//
// bShow: 是否显示.
func XTabBar_ShowLabel(hEle int, index int, bShow bool) bool {
	r, _, _ := xTabBar_ShowLabel.Call(uintptr(hEle), uintptr(index), common.BoolPtr(bShow))
	return r != 0
}
