package xcgui

import "unsafe"

// 窗口组件_判断显示, 判断UI对象是否显示
// hXCGUI: 对象句柄.
func XWidget_IsShow(hXCGUI int) bool {
	r, _, _ := xWidget_IsShow.Call(uintptr(hXCGUI))
	return int(r) != 0
}

// 窗口组件_显示
// hXCGUI: 对象句柄
// bShow: 是否显示
func XWidget_Show(hXCGUI int, bShow bool) int {
	r, _, _ := xWidget_Show.Call(uintptr(hXCGUI), boolPtr(bShow))
	return int(r)
}

// 窗口组件_启用布局控制
// hXCGUI:
// bEnable:
func XWidget_EnableLayoutControl(hXCGUI int, bEnable bool) int {
	r, _, _ := xWidget_EnableLayoutControl.Call(uintptr(hXCGUI), boolPtr(bEnable))
	return int(r)
}

// 窗口组件_是否布局控制
// hXCGUI:
func XWidget_IsLayoutControl(hXCGUI int) bool {
	r, _, _ := xWidget_IsLayoutControl.Call(uintptr(hXCGUI))
	return int(r) != 0
}

// 窗口组件_取父元素
// hXCGUI: 对象句柄.
func XWidget_GetParentEle(hXCGUI int) int {
	r, _, _ := xWidget_GetParentEle.Call(uintptr(hXCGUI))
	return int(r)
}

// 窗口组件_取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父
// hXCGUI: 对象句柄.
func XWidget_GetParent(hXCGUI int) int {
	r, _, _ := xWidget_GetParent.Call(uintptr(hXCGUI))
	return int(r)
}

// 窗口组件_取HWND
// hXCGUI: 对象句柄.
func XWidget_GetHWND(hXCGUI int) int {
	r, _, _ := xWidget_GetHWND.Call(uintptr(hXCGUI))
	return int(r)
}

// 窗口组件_取HWINDOW
// hXCGUI: 对象句柄.
func XWidget_GetHWINDOW(hXCGUI int) int {
	r, _, _ := xWidget_GetHWINDOW.Call(uintptr(hXCGUI))
	return int(r)
}

// 窗口组件_布局项_启用换行, 强制换行
// hXCGUI: UI对象句柄
// bWrap: 是否换行
func XWidget_LayoutItem_EnableWrap(hXCGUI int, bWrap bool) int {
	r, _, _ := xWidget_LayoutItem_EnableWrap.Call(uintptr(hXCGUI), boolPtr(bWrap))
	return int(r)
}

// 窗口组件_布局项_启用交换, 根据水平垂直布局变换, 交换属性(宽度,高度,最小宽度,最小高度)
// hXCGUI: UI对象句柄
// bEnable: 是否启用
func XWidget_LayoutItem_EnableSwap(hXCGUI int, bEnable bool) int {
	r, _, _ := xWidget_LayoutItem_EnableSwap.Call(uintptr(hXCGUI), boolPtr(bEnable))
	return int(r)
}

// 窗口组件_布局项_启用浮动, 向反方向对齐
// hXCGUI: UI对象句柄
// bFloat: 是否浮动
func XWidget_LayoutItem_EnableFloat(hXCGUI int, bFloat bool) int {
	r, _, _ := xWidget_LayoutItem_EnableFloat.Call(uintptr(hXCGUI), boolPtr(bFloat))
	return int(r)
}

// 窗口组件_布局项_置宽度
// hXCGUI: UI对象句柄
// nType: 类型, Layout_Size_
// nWidth: 宽度
func XWidget_LayoutItem_SetWidth(hXCGUI, nType, nWidth int) int {
	r, _, _ := xWidget_LayoutItem_SetWidth.Call(uintptr(hXCGUI), uintptr(nType), uintptr(nWidth))
	return int(r)
}

// 窗口组件_布局项_置高度
// hXCGUI: UI对象句柄
// nType: 类型, Layout_Size_
// nHeight: 高度
func XWidget_LayoutItem_SetHeight(hXCGUI, nType, nHeight int) int {
	r, _, _ := xWidget_LayoutItem_SetHeight.Call(uintptr(hXCGUI), uintptr(nType), uintptr(nHeight))
	return int(r)
}

// 窗口组件_布局项_取宽度
// hXCGUI: UI对象句柄
// pType: 返回类型
// pWidth: 返回宽度
func XWidget_LayoutItem_GetWidth(hXCGUI int, pType, pWidth *int) int {
	r, _, _ := xWidget_LayoutItem_GetWidth.Call(uintptr(hXCGUI), uintptr(unsafe.Pointer(pType)), uintptr(unsafe.Pointer(pWidth)))
	return int(r)
}

// 窗口组件_布局项_取高度
// hXCGUI: UI对象句柄
// pType: 返回类型
// pHeight: 返回高度
func XWidget_LayoutItem_GetHeight(hXCGUI int, pType, pHeight *int) int {
	r, _, _ := xWidget_LayoutItem_GetHeight.Call(uintptr(hXCGUI), uintptr(unsafe.Pointer(pType)), uintptr(unsafe.Pointer(pHeight)))
	return int(r)
}

// 窗口组件_布局项_置对齐, 根据水平垂直轴变化对齐
// hXCGUI: UI对象句柄
// nAlign: 对齐方式, Layout_Align_Axis_
func XWidget_LayoutItem_SetAlign(hXCGUI, nAlign int) int {
	r, _, _ := xWidget_LayoutItem_SetAlign.Call(uintptr(hXCGUI), uintptr(nAlign))
	return int(r)
}

// 窗口组件_布局项_置外间距
// hXCGUI: UI对象句柄
func XWidget_LayoutItem_SetMargin(hXCGUI, left, top, right, bottom int) int {
	r, _, _ := xWidget_LayoutItem_SetMargin.Call(uintptr(hXCGUI), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口组件_布局项_置对齐
// hXCGUI: UI对象句柄
// pMargin: 接收返回
func XWidget_LayoutItem_GetMargin(hXCGUI int, pMargin *RECT) int {
	r, _, _ := xWidget_LayoutItem_GetMargin.Call(uintptr(hXCGUI), uintptr(unsafe.Pointer(pMargin)))
	return int(r)
}

// 窗口组件_布局项_置最小大小, 限制大小仅针对缩放有效(自动, 填充父, 比例, 百分比)
// hXCGUI: UI对象句柄
// width: 最小宽度
// height: 最小高度
func XWidget_LayoutItem_SetMinSize(hXCGUI, width, height int) int {
	r, _, _ := xWidget_LayoutItem_SetMinSize.Call(uintptr(hXCGUI), uintptr(width), uintptr(height))
	return int(r)
}

// 窗口组件_布局项_置位置, 相对位置, 值大于等于0有效
// hXCGUI: UI对象句柄
// left: 左边距离
// top: 上边距离
// right: 右边距离
// bottom: 下边距离
func XWidget_LayoutItem_SetPosition(hXCGUI, left, top, right, bottom int) int {
	r, _, _ := xWidget_LayoutItem_SetPosition.Call(uintptr(hXCGUI), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}
