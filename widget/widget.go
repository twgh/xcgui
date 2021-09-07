package widget

import (
	"github.com/twgh/xcgui/xc"
)

type Widget struct {
	HXCGUI int
}

// 窗口组件_判断显示, 判断UI对象是否显示
func (w *Widget) IsShow() bool {
	return xc.XWidget_IsShow(w.HXCGUI)
}

// 窗口组件_显示
// bShow: 是否显示
func (w *Widget) Show(bShow bool) int {
	return xc.XWidget_Show(w.HXCGUI, bShow)
}

// 窗口组件_启用布局控制
// bEnable:
func (w *Widget) EnableLayoutControl(bEnable bool) int {
	return xc.XWidget_EnableLayoutControl(w.HXCGUI, bEnable)
}

// 窗口组件_是否布局控制
func (w *Widget) IsLayoutControl() bool {
	return xc.XWidget_IsLayoutControl(w.HXCGUI)
}

// 窗口组件_取父元素
func (w *Widget) GetParentEle() int {
	return xc.XWidget_GetParentEle(w.HXCGUI)
}

// 窗口组件_取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父
func (w *Widget) GetParent() int {
	return xc.XWidget_GetParent(w.HXCGUI)
}

// 窗口组件_取HWND
func (w *Widget) GetHWND() int {
	return xc.XWidget_GetHWND(w.HXCGUI)
}

// 窗口组件_取HWINDOW
func (w *Widget) GetHWINDOW() int {
	return xc.XWidget_GetHWINDOW(w.HXCGUI)
}

// 窗口组件_布局项_启用换行, 强制换行
// bWrap: 是否换行
func (w *Widget) LayoutItem_EnableWrap(bWrap bool) int {
	return xc.XWidget_LayoutItem_EnableWrap(w.HXCGUI, bWrap)
}

// 窗口组件_布局项_启用交换, 根据水平垂直布局变换, 交换属性(宽度,高度,最小宽度,最小高度)
// bEnable: 是否启用
func (w *Widget) LayoutItem_EnableSwap(bEnable bool) int {
	return xc.XWidget_LayoutItem_EnableSwap(w.HXCGUI, bEnable)
}

// 窗口组件_布局项_启用浮动, 向反方向对齐
// bFloat: 是否浮动
func (w *Widget) LayoutItem_EnableFloat(bFloat bool) int {
	return xc.XWidget_LayoutItem_EnableFloat(w.HXCGUI, bFloat)
}

// 窗口组件_布局项_置宽度
// nType: 类型, Layout_Size_
// nWidth: 宽度
func (w *Widget) LayoutItem_SetWidth(nType, nWidth int) int {
	return xc.XWidget_LayoutItem_SetWidth(w.HXCGUI, nType, nWidth)
}

// 窗口组件_布局项_置高度
// nType: 类型, Layout_Size_
// nHeight: 高度
func (w *Widget) LayoutItem_SetHeight(nType, nHeight int) int {
	return xc.XWidget_LayoutItem_SetHeight(w.HXCGUI, nType, nHeight)
}

// 窗口组件_布局项_取宽度
// pType: 返回类型
// pWidth: 返回宽度
func (w *Widget) LayoutItem_GetWidth(pType, pWidth *int) int {
	return xc.XWidget_LayoutItem_GetWidth(w.HXCGUI, pType, pWidth)
}

// 窗口组件_布局项_取高度
// pType: 返回类型
// pHeight: 返回高度
func (w *Widget) LayoutItem_GetHeight(pType, pHeight *int) int {
	return xc.XWidget_LayoutItem_GetHeight(w.HXCGUI, pType, pHeight)
}

// 窗口组件_布局项_置对齐, 根据水平垂直轴变化对齐
// nAlign: 对齐方式, Layout_Align_Axis_
func (w *Widget) LayoutItem_SetAlign(nAlign int) int {
	return xc.XWidget_LayoutItem_SetAlign(w.HXCGUI, nAlign)
}

// 窗口组件_布局项_置外间距
func (w *Widget) LayoutItem_SetMargin(left, top, right, bottom int) int {
	return xc.XWidget_LayoutItem_SetMargin(w.HXCGUI, left, top, right, bottom)
}

// 窗口组件_布局项_置对齐
// pMargin: 接收返回
func (w *Widget) LayoutItem_GetMargin(pMargin *xc.RECT) int {
	return xc.XWidget_LayoutItem_GetMargin(w.HXCGUI, pMargin)
}

// 窗口组件_布局项_置最小大小, 限制大小仅针对缩放有效(自动, 填充父, 比例, 百分比)
// width: 最小宽度
// height: 最小高度
func (w *Widget) LayoutItem_SetMinSize(width, height int) int {
	return xc.XWidget_LayoutItem_SetMinSize(w.HXCGUI, width, height)
}

// 窗口组件_布局项_置位置, 相对位置, 值大于等于0有效
// left: 左边距离
// top: 上边距离
// right: 右边距离
// bottom: 下边距离
func (w *Widget) LayoutItem_SetPosition(left, top, right, bottom int) int {
	return xc.XWidget_LayoutItem_SetPosition(w.HXCGUI, left, top, right, bottom)
}
