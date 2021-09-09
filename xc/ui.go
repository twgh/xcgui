package xc

// 可视对象_置样式, 设置UI对象样式
// hXCGUI: 对象句柄.
// nStyle: 样式值: Button_Style_ , FrameWnd_Style_ , ToolBar_Style_ , ListBox_Style_
func XUI_SetStyle(hXCGUI, nStyle int) int {
	r, _, _ := xUI_SetStyle.Call(uintptr(hXCGUI), uintptr(nStyle))
	return int(r)
}

// 可视对象_取样式, 获取UI对象样式, 返回: Button_Style_ , FrameWnd_Style_ , ToolBar_Style_ , ListBox_Style_
// hXCGUI: 对象句柄.
func XUI_GetStyle(hXCGUI int) int {
	r, _, _ := xUI_GetStyle.Call(uintptr(hXCGUI))
	return int(r)
}

// 可视对象_启用CSS, 启用或禁用样式
// hXCGUI: 对象句柄.
// bEnable: 是否启用.
func XUI_EnableCSS(hXCGUI int, bEnable bool) int {
	r, _, _ := xUI_EnableCSS.Call(uintptr(hXCGUI), boolPtr(bEnable))
	return int(r)
}

// 可视对象_置CSS名称, 设置CSS[套用样式]名称
// hXCGUI: 对象句柄.
// pName: 套用样式名称.
func XUI_SetCssName(hXCGUI int, pName string) int {
	r, _, _ := xUI_SetCssName.Call(uintptr(hXCGUI), strPtr(pName))
	return int(r)
}

// 可视对象_取CSS名称, 获取CSS样式名称
// hXCGUI: 对象句柄.
func XUI_GetCssName(hXCGUI int) string {
	r, _, _ := xUI_GetCssName.Call(uintptr(hXCGUI))
	return uintPtrToString(r)
}
