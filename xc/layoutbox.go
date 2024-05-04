package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 布局盒子_启用水平.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableHorizon(hLayoutBox int, bEnable bool) {
	xLayoutBox_EnableHorizon.Call(uintptr(hLayoutBox), common.BoolPtr(bEnable))
}

// 布局盒子_启用自动换行.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableAutoWrap(hLayoutBox int, bEnable bool) {
	xLayoutBox_EnableAutoWrap.Call(uintptr(hLayoutBox), common.BoolPtr(bEnable))
}

// 布局盒子_启用溢出隐藏.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableOverflowHide(hLayoutBox int, bEnable bool) {
	xLayoutBox_EnableOverflowHide.Call(uintptr(hLayoutBox), common.BoolPtr(bEnable))
}

// XLayoutBox_SetAlignH 布局盒子_置水平对齐.
//
//	@param hLayoutBox 窗口或布局元素或布局框架句柄.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
func XLayoutBox_SetAlignH(hLayoutBox int, nAlign xcc.Layout_Align_) {
	xLayoutBox_SetAlignH.Call(uintptr(hLayoutBox), uintptr(nAlign))
}

// XLayoutBox_SetAlignV 布局盒子_置垂直对齐.
//
//	@param hLayoutBox 窗口或布局元素或布局框架句柄.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
func XLayoutBox_SetAlignV(hLayoutBox int, nAlign xcc.Layout_Align_) {
	xLayoutBox_SetAlignV.Call(uintptr(hLayoutBox), uintptr(nAlign))
}

// 布局盒子_置对齐基线.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func XLayoutBox_SetAlignBaseline(hLayoutBox int, nAlign xcc.Layout_Align_Axis_) {
	xLayoutBox_SetAlignBaseline.Call(uintptr(hLayoutBox), uintptr(nAlign))
}

// 布局盒子_置间距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 项间距大小.
func XLayoutBox_SetSpace(hLayoutBox int, nSpace int) {
	xLayoutBox_SetSpace.Call(uintptr(hLayoutBox), uintptr(nSpace))
}

// 布局盒子_置行距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 行间距大小.
func XLayoutBox_SetSpaceRow(hLayoutBox int, nSpace int) {
	xLayoutBox_SetSpaceRow.Call(uintptr(hLayoutBox), uintptr(nSpace))
}
