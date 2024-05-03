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
func (u *UI) SetStyle(nStyle xcc.XC_OBJECT_STYLE) *UI {
	xc.XUI_SetStyle(u.Handle, nStyle)
	return u
}

// GetStyle 可视对象_取样式, 获取UI对象样式.
//
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func (u *UI) GetStyle() xcc.XC_OBJECT_STYLE {
	return xc.XUI_GetStyle(u.Handle)
}

// EnableCSS 可视对象_启用CSS. 启用或禁用样式, 并且覆盖内嵌子元素属性, 例如: 滚动视图上的滚动条, 滚动条上的按钮.
//
//	@param bEnable 是否启用.
func (u *UI) EnableCSS(bEnable bool) *UI {
	xc.XUI_EnableCSS(u.Handle, bEnable)
	return u
}

// EnableCssEx 可视对象_启用CSS. 启用或禁用样式, 仅设置自身属性, 不包含内嵌子元素属性, 例如:滚动视图上的滚动条, 滚动条上的按钮
//
//	@param bEnable 是否启用.
func (u *UI) EnableCssEx(bEnable bool) *UI {
	xc.XUI_EnableCssEx(u.Handle, bEnable)
	return u
}

// SetCssName 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
//	@param pName 套用样式名称.
func (u *UI) SetCssName(pName string) *UI {
	xc.XUI_SetCssName(u.Handle, pName)
	return u
}

// GetCssName 可视对象_取CSS名称, 获取CSS样式名称.
//
//	@return string
func (u *UI) GetCssName() string {
	return xc.XUI_GetCssName(u.Handle)
}
