package widget

import "github.com/twgh/xcgui/xc"

// 布局框架
type LayoutFrame struct {
	Element
}

// 布局框架_创建
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutFrame(x int, y int, cx int, cy int, hParent int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
	return p
}

// 布局框架_显示布局边界
// bEnable: 是否启用
func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayoutFrame_ShowLayoutFrame(l.Handle, bEnable)
}
