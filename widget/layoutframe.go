package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// LayoutFrame 布局框架.
type LayoutFrame struct {
	ScrollView
}

// NewLayoutFrame 布局框架_创建, 失败返回 nil.
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
func NewLayoutFrame(x, y, cx, cy int32, hParent int) *LayoutFrame {
	return NewLayoutFrameByHandle(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
}

// NewLayoutFrameByHandle 从句柄创建对象, 失败返回 nil.
//
// handle: 布局框架句柄.
func NewLayoutFrameByHandle(handle int) *LayoutFrame {
	if handle <= 0 {
		return nil
	}
	p := &LayoutFrame{}
	p.SetHandle(handle)
	return p
}

// NewLayoutFrameByName 从name创建对象, 失败返回 nil.
//
// name: name名称.
func NewLayoutFrameByName(name string) *LayoutFrame {
	return NewLayoutFrameByHandle(xc.XC_GetObjectByName(name))
}

// NewLayoutFrameByUID 从UID创建对象, 失败返回 nil.
//
// nUID: ID值.
func NewLayoutFrameByUID(nUID int32) *LayoutFrame {
	return NewLayoutFrameByHandle(xc.XC_GetObjectByUID(nUID))
}

// NewLayoutFrameByUIDName 从UID名称创建对象, 失败返回 nil.
//
// name: name名称.
func NewLayoutFrameByUIDName(name string) *LayoutFrame {
	return NewLayoutFrameByHandle(xc.XC_GetObjectByUIDName(name))
}

// ShowLayoutFrame 布局框架_显示布局边界.
//
// bEnable: 是否启用.
func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) *LayoutFrame {
	xc.XLayoutFrame_ShowLayoutFrame(l.Handle, bEnable)
	return l
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableHorizon(bEnable bool) *LayoutFrame {
	xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
	return l
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableAutoWrap(bEnable bool) *LayoutFrame {
	xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
	return l
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (l *LayoutFrame) EnableOverflowHide(bEnable bool) *LayoutFrame {
	xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
	return l
}

// SetAlignH 布局盒子_置水平对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (l *LayoutFrame) SetAlignH(nAlign xcc.Layout_Align_) *LayoutFrame {
	xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
	return l
}

// SetAlignV 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (l *LayoutFrame) SetAlignV(nAlign xcc.Layout_Align_) *LayoutFrame {
	xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
	return l
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func (l *LayoutFrame) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) *LayoutFrame {
	xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
	return l
}

// SetSpace 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (l *LayoutFrame) SetSpace(nSpace int32) *LayoutFrame {
	xc.XLayoutBox_SetSpace(l.Handle, nSpace)
	return l
}

// SetSpaceRow 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (l *LayoutFrame) SetSpaceRow(nSpace int32) *LayoutFrame {
	xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
	return l
}
