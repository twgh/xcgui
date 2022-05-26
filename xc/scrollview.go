package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

// 滚动视_创建, 创建滚动视图元素, 返回元素句柄.
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
func XSView_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xSView_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 滚动视_置视图大小, 设置内容大小, 如果内容改变返回TRUE否则返回FALSE.
//
// hEle: 元素句柄.
//
// cx: 宽度.
//
// cy: 高度.
func XSView_SetTotalSize(hEle int, cx int, cy int) bool {
	r, _, _ := xSView_SetTotalSize.Call(uintptr(hEle), uintptr(cx), uintptr(cy))
	return r != 0
}

// 滚动视_取视图大小, 获取内容总大小.
//
// hEle: 元素句柄.
//
// pSize: 大小.
func XSView_GetTotalSize(hEle int, pSize *SIZE) int {
	r, _, _ := xSView_GetTotalSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 滚动视_置滚动单位大小, 设置滚动单位大小, 如果内容改变返回TRUE否则返回FALSE.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XSView_SetLineSize(hEle int, nWidth int, nHeight int) bool {
	r, _, _ := xSView_SetLineSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight))
	return r != 0
}

// 滚动视_取滚动单位大小, 获取滚动单位大小.
//
// hEle: 元素句柄.
//
// pSize: 返回大小.
func XSView_GetLineSize(hEle int, pSize *SIZE) int {
	r, _, _ := xSView_GetLineSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 滚动视_置滚动条大小.
//
// hEle: 元素句柄.
//
// size: 滚动条大小.
func XSView_SetScrollBarSize(hEle int, size int) int {
	r, _, _ := xSView_SetScrollBarSize.Call(uintptr(hEle), uintptr(size))
	return int(r)
}

// 滚动视_取视口原点X, 获取视口原点X坐标.
//
// hEle: 元素句柄.
func XSView_GetViewPosH(hEle int) int {
	r, _, _ := xSView_GetViewPosH.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_取视口原点Y, 获取视口原点Y坐标.
//
// hEle: 元素句柄.
func XSView_GetViewPosV(hEle int) int {
	r, _, _ := xSView_GetViewPosV.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_取视口宽度.
//
// hEle: 元素句柄.
func XSView_GetViewWidth(hEle int) int {
	r, _, _ := xSView_GetViewWidth.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_取视口高度.
//
// hEle: 元素句柄.
func XSView_GetViewHeight(hEle int) int {
	r, _, _ := xSView_GetViewHeight.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_取视口坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XSView_GetViewRect(hEle int, pRect *RECT) int {
	r, _, _ := xSView_GetViewRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 滚动视_取水平滚动条, 返回滚动条句柄.
//
// hEle: 元素句柄.
func XSView_GetScrollBarH(hEle int) int {
	r, _, _ := xSView_GetScrollBarH.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_取垂直滚动条, 返回滚动条句柄.
//
// hEle: 元素句柄.
func XSView_GetScrollBarV(hEle int) int {
	r, _, _ := xSView_GetScrollBarV.Call(uintptr(hEle))
	return int(r)
}

// 滚动视_水平滚动, 水平滚动条, 滚动到指定位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func XSView_ScrollPosH(hEle int, pos int) bool {
	r, _, _ := xSView_ScrollPosH.Call(uintptr(hEle), uintptr(pos))
	return r != 0
}

// 滚动视_垂直滚动, 垂直滚动条, 滚动到指定位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func XSView_ScrollPosV(hEle int, pos int) bool {
	r, _, _ := xSView_ScrollPosV.Call(uintptr(hEle), uintptr(pos))
	return r != 0
}

// 滚动视_水平滚动到X, 水平滚动条, 滚动到指定坐标.
//
// hEle: 元素句柄.
//
// posX: X坐标.
func XSView_ScrollPosXH(hEle int, posX int) bool {
	r, _, _ := xSView_ScrollPosXH.Call(uintptr(hEle), uintptr(posX))
	return r != 0
}

// 滚动视_垂直滚动到Y, 垂直滚动条, 滚动到指定坐标.
//
// hEle: 元素句柄.
//
// posY: Y坐标.
func XSView_ScrollPosYV(hEle int, posY int) bool {
	r, _, _ := xSView_ScrollPosYV.Call(uintptr(hEle), uintptr(posY))
	return r != 0
}

// 滚动视_显示水平滚动条.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.
func XSView_ShowSBarH(hEle int, bShow bool) int {
	r, _, _ := xSView_ShowSBarH.Call(uintptr(hEle), common.BoolPtr(bShow))
	return int(r)
}

// 滚动视_显示垂直滚动条.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.
func XSView_ShowSBarV(hEle int, bShow bool) int {
	r, _, _ := xSView_ShowSBarV.Call(uintptr(hEle), common.BoolPtr(bShow))
	return int(r)
}

// 滚动视_启用自动显示滚动条.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XSView_EnableAutoShowScrollBar(hEle int, bEnable bool) int {
	r, _, _ := xSView_EnableAutoShowScrollBar.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 滚动视_向左滚动.
//
// hEle: 元素句柄.
func XSView_ScrollLeftLine(hEle int) bool {
	r, _, _ := xSView_ScrollLeftLine.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_向右滚动.
//
// hEle: 元素句柄.
func XSView_ScrollRightLine(hEle int) bool {
	r, _, _ := xSView_ScrollRightLine.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_向上滚动.
//
// hEle: 元素句柄.
func XSView_ScrollTopLine(hEle int) bool {
	r, _, _ := xSView_ScrollTopLine.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_向下滚动.
//
// hEle: 元素句柄.
func XSView_ScrollBottomLine(hEle int) bool {
	r, _, _ := xSView_ScrollBottomLine.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_滚动到左侧, 水平滚动到左侧.
//
// hEle: 元素句柄.
func XSView_ScrollLeft(hEle int) bool {
	r, _, _ := xSView_ScrollLeft.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_滚动到右侧, 水平滚动到右侧.
//
// hEle: 元素句柄.
func XSView_ScrollRight(hEle int) bool {
	r, _, _ := xSView_ScrollRight.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_滚动到顶部, 垂直滚动到顶部.
//
// hEle: 元素句柄.
func XSView_ScrollTop(hEle int) bool {
	r, _, _ := xSView_ScrollTop.Call(uintptr(hEle))
	return r != 0
}

// 滚动视_滚动到底部, 垂直滚动到底部.
//
// hEle: 元素句柄.
func XSView_ScrollBottom(hEle int) bool {
	r, _, _ := xSView_ScrollBottom.Call(uintptr(hEle))
	return r != 0
}
