package widget

import "github.com/twgh/xcgui/xc"

// 布局框架.
type LayoutFrame struct {
	Element
}

// 布局框架_创建.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutFrame(x int, y int, cx int, cy int, hParent int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewLayoutFrameByHandle(handle int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewLayoutFrameByName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewLayoutFrameByUID(nUID int) *LayoutFrame {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewLayoutFrameByUIDName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 布局框架_显示布局边界.
//
// bEnable: 是否启用.
func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayoutFrame_ShowLayoutFrame(l.Handle, bEnable)
}

/*
LayoutBox-布局盒子
*/

// 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
}

// 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
}

// 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
}

// 布局盒子_置水平对齐.
//
// nAlign: 对齐方式.
func (l *LayoutFrame) SetAlignH(nAlign int) int {
	return xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
}

// 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式.
func (l *LayoutFrame) SetAlignV(nAlign int) int {
	return xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
}

// 布局盒子_置对齐基线.
//
// nAlign: 对齐方式.
func (l *LayoutFrame) SetAlignBaseline(nAlign int) int {
	return xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
}

// 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (l *LayoutFrame) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.Handle, nSpace)
}

// 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (l *LayoutFrame) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
}
