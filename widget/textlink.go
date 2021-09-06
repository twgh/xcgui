package widget

import (
	xc "github.com/twgh/xcgui"
)

type TextLink struct {
	Element

	HEle int
}

// 文本链接_创建, 创建静态文本连接元素, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// pName: 文本内容.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素
func NewTextLink(x int, y int, cx int, cy int, pName string, hParent int) *TextLink {
	p := &TextLink{
		HEle: xc.XTextLink_Create(x, y, cx, cy, pName, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态
// bEnable: 是否启用.
func (t *TextLink) EnableUnderlineLeave(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineLeave(t.HEle, bEnable)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态
// bEnable: 是否启用.
func (t *TextLink) EnableUnderlineStay(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineStay(t.HEle, bEnable)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态
// color: RGB颜色值.
// alpha: 透明度.
func (t *TextLink) SetTextColorStay(color int, alpha uint8) int {
	return xc.XTextLink_SetTextColorStay(t.HEle, color, alpha)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态
// color: RGB颜色值.
// alpha: 透明度.
func (t *TextLink) SetUnderlineColorLeave(color int, alpha uint8) int {
	return xc.XTextLink_SetUnderlineColorLeave(t.HEle, color, alpha)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态
// color: RGB颜色值.
// alpha: 透明度.
func (t *TextLink) SetUnderlineColorStay(color int, alpha uint8) int {
	return xc.XTextLink_SetUnderlineColorStay(t.HEle, color, alpha)
}
