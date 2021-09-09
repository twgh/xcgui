package widget

import (
	"github.com/twgh/xcgui/xc"
)

// 布局元素
type LayoutEle struct {
	Element
}

// 布局_创建, 创建布局元素
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父为窗口句柄或元素句柄.
func NewLayout(x int, y int, cx int, cy int, hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_Create(x, y, cx, cy, hParent))
	return p
}

// 布局_创建扩展, 创建布局元素
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutEx(hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_CreateEx(hParent))
	return p
}

// 布局_判断启用, 是否已经启用布局功能
func (l *LayoutEle) IsEnableLayout() bool {
	return xc.XLayout_IsEnableLayout(l.Handle)
}

// 布局_启用, 启用布局功能
// bEnable: 是否启用.
func (l *LayoutEle) EnableLayout(bEnable bool) int {
	return xc.XLayout_EnableLayout(l.Handle, bEnable)
}

// 布局_显示布局边界, 显示布局边界
// bEnable: 是否显示.
func (l *LayoutEle) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayout_ShowLayoutFrame(l.Handle, bEnable)
}

// 布局_取内宽度, 获取宽度,不包含内边距大小
func (l *LayoutEle) GetWidthIn() int {
	return xc.XLayout_GetWidthIn(l.Handle)
}

// 布局_取内高度, 获取高度,不包含内边距大小
func (l *LayoutEle) GetHeightIn() int {
	return xc.XLayout_GetHeightIn(l.Handle)
}
