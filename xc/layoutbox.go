package xc

// 布局盒子_启用水平.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableHorizon(hLayoutBox int, bEnable bool) int {
	r, _, _ := xLayoutBox_EnableHorizon.Call(uintptr(hLayoutBox), boolPtr(bEnable))
	return int(r)
}

// 布局盒子_启用自动换行.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableAutoWrap(hLayoutBox int, bEnable bool) int {
	r, _, _ := xLayoutBox_EnableAutoWrap.Call(uintptr(hLayoutBox), boolPtr(bEnable))
	return int(r)
}

// 布局盒子_启用溢出隐藏.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func XLayoutBox_EnableOverflowHide(hLayoutBox int, bEnable bool) int {
	r, _, _ := xLayoutBox_EnableOverflowHide.Call(uintptr(hLayoutBox), boolPtr(bEnable))
	return int(r)
}

// 布局盒子_置水平对齐.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nAlign: 对齐方式.
func XLayoutBox_SetAlignH(hLayoutBox int, nAlign int) int {
	r, _, _ := xLayoutBox_SetAlignH.Call(uintptr(hLayoutBox), uintptr(nAlign))
	return int(r)
}

// 布局盒子_置垂直对齐.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nAlign: 对齐方式.
func XLayoutBox_SetAlignV(hLayoutBox int, nAlign int) int {
	r, _, _ := xLayoutBox_SetAlignV.Call(uintptr(hLayoutBox), uintptr(nAlign))
	return int(r)
}

// 布局盒子_置对齐基线.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nAlign: 对齐方式.
func XLayoutBox_SetAlignBaseline(hLayoutBox int, nAlign int) int {
	r, _, _ := xLayoutBox_SetAlignBaseline.Call(uintptr(hLayoutBox), uintptr(nAlign))
	return int(r)
}

// 布局盒子_置间距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 项间距大小.
func XLayoutBox_SetSpace(hLayoutBox int, nSpace int) int {
	r, _, _ := xLayoutBox_SetSpace.Call(uintptr(hLayoutBox), uintptr(nSpace))
	return int(r)
}

// 布局盒子_置行距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 行间距大小.
func XLayoutBox_SetSpaceRow(hLayoutBox int, nSpace int) int {
	r, _, _ := xLayoutBox_SetSpaceRow.Call(uintptr(hLayoutBox), uintptr(nSpace))
	return int(r)
}
