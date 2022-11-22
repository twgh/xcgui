package objectbase

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// UI 可视对象.
type UI struct {
	ObjectBase
}

// SetStyle 可视对象_置样式, 设置UI对象样式.
//
//	@param nStyle xcc.XC_OBJECT_STYLE, 样式值: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
//	@return int
func (u *UI) SetStyle(nStyle xcc.XC_OBJECT_STYLE) int {
	return xc.XUI_SetStyle(u.Handle, nStyle)
}

// GetStyle 可视对象_取样式, 获取UI对象样式.
//
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func (u *UI) GetStyle() xcc.XC_OBJECT_STYLE {
	return xc.XUI_GetStyle(u.Handle)
}

// EnableCSS 可视对象_启用CSS, 启用或禁用样式.
//
//	@param bEnable 是否启用.
//	@return int
func (u *UI) EnableCSS(bEnable bool) int {
	return xc.XUI_EnableCSS(u.Handle, bEnable)
}

// SetCssName 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
//	@param pName 套用样式名称.
//	@return int
func (u *UI) SetCssName(pName string) int {
	return xc.XUI_SetCssName(u.Handle, pName)
}

// GetCssName 可视对象_取CSS名称, 获取CSS样式名称.
//
//	@return string
func (u *UI) GetCssName() string {
	return xc.XUI_GetCssName(u.Handle)
}
