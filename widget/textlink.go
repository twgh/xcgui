package widget

import (
	"github.com/twgh/xcgui/xc"
)

// 静态文本链接按钮.
type TextLink struct {
	Button
}

// 文本链接_创建, 创建静态文本链接元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func NewTextLink(x int, y int, cx int, cy int, pName string, hParent int) *TextLink {
	p := &TextLink{}
	p.SetHandle(xc.XTextLink_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
func NewTextLinkByHandle(handle int) *TextLink {
	p := &TextLink{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewTextLinkByName(name string) *TextLink {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewTextLinkByUID(nUID int) *TextLink {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewTextLinkByUIDName(name string) *TextLink {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态.
//
// bEnable: 是否启用.
func (t *TextLink) EnableUnderlineLeave(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineLeave(t.Handle, bEnable)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态.
//
// bEnable: 是否启用.
func (t *TextLink) EnableUnderlineStay(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineStay(t.Handle, bEnable)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) SetTextColorStay(color int) int {
	return xc.XTextLink_SetTextColorStay(t.Handle, color)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态.
//
// color: ABGR 颜色值.
func (t *TextLink) SetUnderlineColorLeave(color int) int {
	return xc.XTextLink_SetUnderlineColorLeave(t.Handle, color)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) SetUnderlineColorStay(color int) int {
	return xc.XTextLink_SetUnderlineColorStay(t.Handle, color)
}
