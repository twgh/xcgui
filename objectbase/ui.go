package objectbase

import "github.com/twgh/xcgui/xc"

// 可视对象.
type UI struct {
	ObjectBase
}

// 可视对象_置样式, 设置UI对象样式.
//
// nStyle: 样式值: Button_Style_ , FrameWnd_Style_ , ToolBar_Style_ , ListBox_Style_.
func (u *UI) SetStyle(nStyle int) int {
	return xc.XUI_SetStyle(u.Handle, nStyle)
}

// 可视对象_取样式, 获取UI对象样式, 返回: Button_Style_ , FrameWnd_Style_ , ToolBar_Style_ , ListBox_Style_.
func (u *UI) GetStyle() int {
	return xc.XUI_GetStyle(u.Handle)
}

// 可视对象_启用CSS, 启用或禁用样式.
//
// bEnable: 是否启用.
func (u *UI) EnableCSS(bEnable bool) int {
	return xc.XUI_EnableCSS(u.Handle, bEnable)
}

// 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
// pName: 套用样式名称.
func (u *UI) SetCssName(pName string) int {
	return xc.XUI_SetCssName(u.Handle, pName)
}

// 可视对象_取CSS名称, 获取CSS样式名称.
func (u *UI) GetCssName() string {
	return xc.XUI_GetCssName(u.Handle)
}
