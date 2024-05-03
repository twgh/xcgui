package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// XUI_SetStyle 可视对象_置样式, 设置UI对象样式.
//
//	@param hXCGUI 对象句柄.
//	@param nStyle xcc.XC_OBJECT_STYLE, 样式值: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
//	@return int
func XUI_SetStyle(hXCGUI int, nStyle xcc.XC_OBJECT_STYLE) int {
	r, _, _ := xUI_SetStyle.Call(uintptr(hXCGUI), uintptr(nStyle))
	return int(r)
}

// XUI_GetStyle 可视对象_取样式, 获取UI对象样式
//
//	@param hXCGUI 对象句柄.
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func XUI_GetStyle(hXCGUI int) xcc.XC_OBJECT_STYLE {
	r, _, _ := xUI_GetStyle.Call(uintptr(hXCGUI))
	return xcc.XC_OBJECT_STYLE(r)
}

// XUI_EnableCSS 可视对象_启用CSS. 启用或禁用样式, 并且覆盖内嵌子元素属性, 例如: 滚动视图上的滚动条, 滚动条上的按钮.
//
//	@param hXCGUI 对象句柄.
//	@param bEnable 是否启用.
func XUI_EnableCSS(hXCGUI int, bEnable bool) {
	xUI_EnableCSS.Call(uintptr(hXCGUI), common.BoolPtr(bEnable))
}

// XUI_EnableCssEx 可视对象_启用CSS, 启用或禁用样式, 仅设置自身属性, 不包含内嵌子元素属性, 例如:滚动视图上的滚动条, 滚动条上的按钮
//
//	@param hXCGUI 对象句柄.
//	@param bEnable 是否启用.
func XUI_EnableCssEx(hXCGUI int, bEnable bool) {
	xUI_EnableCssEx.Call(uintptr(hXCGUI), common.BoolPtr(bEnable))
}

// XUI_SetCssName 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
//	@param hXCGUI 对象句柄.
//	@param pName 套用样式名称.
//	@return int
func XUI_SetCssName(hXCGUI int, pName string) int {
	r, _, _ := xUI_SetCssName.Call(uintptr(hXCGUI), common.StrPtr(pName))
	return int(r)
}

// XUI_GetCssName 可视对象_取CSS名称, 获取CSS样式名称.
//
//	@param hXCGUI 对象句柄.
//	@return string
func XUI_GetCssName(hXCGUI int) string {
	r, _, _ := xUI_GetCssName.Call(uintptr(hXCGUI))
	return common.UintPtrToString(r)
}
