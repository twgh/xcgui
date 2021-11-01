package widget

import (
	"github.com/twgh/xcgui/xc"
)

// 布局元素.
type LayoutEle struct {
	Element
}

// 布局_创建, 创建布局元素.
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
func NewLayout(x int, y int, cx int, cy int, hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_Create(x, y, cx, cy, hParent))
	return p
}

// 布局_创建扩展, 创建布局元素.
//
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutEx(hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_CreateEx(hParent))
	return p
}

// 从句柄创建对象.
func NewLayoutEleByHandle(handle int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewLayoutEleByName(name string) *LayoutEle {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewLayoutEleByUID(nUID int) *LayoutEle {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewLayoutEleByUIDName(name string) *LayoutEle {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 布局_判断启用, 是否已经启用布局功能.
func (l *LayoutEle) IsEnableLayout() bool {
	return xc.XLayout_IsEnableLayout(l.Handle)
}

// 布局_启用, 启用布局功能.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableLayout(bEnable bool) int {
	return xc.XLayout_EnableLayout(l.Handle, bEnable)
}

// 布局_显示布局边界, 显示布局边界.
//
// bEnable: 是否显示.
func (l *LayoutEle) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayout_ShowLayoutFrame(l.Handle, bEnable)
}

// 布局_取内宽度, 获取宽度,不包含内边距大小.
func (l *LayoutEle) GetWidthIn() int {
	return xc.XLayout_GetWidthIn(l.Handle)
}

// 布局_取内高度, 获取高度,不包含内边距大小.
func (l *LayoutEle) GetHeightIn() int {
	return xc.XLayout_GetHeightIn(l.Handle)
}

/*
LayoutBox-布局盒子
*/

// 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
}

// 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
}

// 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
}

// 布局盒子_置水平对齐.
//
// nAlign: 对齐方式.
func (l *LayoutEle) SetAlignH(nAlign int) int {
	return xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
}

// 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式.
func (l *LayoutEle) SetAlignV(nAlign int) int {
	return xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
}

// 布局盒子_置对齐基线.
//
// nAlign: 对齐方式.
func (l *LayoutEle) SetAlignBaseline(nAlign int) int {
	return xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
}

// 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (l *LayoutEle) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.Handle, nSpace)
}

// 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (l *LayoutEle) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
}
