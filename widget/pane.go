package widget

import "github.com/twgh/xcgui/xc"

// Pane元素
type Pane struct {
	Element
}

// 窗格_创建, 创建窗格元素, 返回元素句柄
// pName: 窗格标题.
// nWidth: 宽度.
// nHeight: 高度.
// hFrameWnd: 框架窗口.
func NewPane(pName string, nWidth int, nHeight int, hFrameWnd int) *Pane {
	p := &Pane{}
	p.SetHandle(xc.XPane_Create(pName, nWidth, nHeight, hFrameWnd))
	return p
}

// 窗格_置视图, 设置窗格视图元素
// hView: 绑定视图元素.
func (p *Pane) SetView(hView int) int {
	return xc.XPane_SetView(p.Handle, hView)
}

// 窗格_置标题, 设置标题文本
// pTitle: 文本内容.
func (p *Pane) SetTitle(pTitle string) int {
	return xc.XPane_SetTitle(p.Handle, pTitle)
}

// 窗格_取标题, 获取标题文本
func (p *Pane) GetTitle() string {
	return xc.XPane_GetTitle(p.Handle)
}

// 窗格_置标题栏高度, 设置标题栏高度
// nHeight: 高度.
func (p *Pane) SetCaptionHeight(nHeight int) int {
	return xc.XPane_SetCaptionHeight(p.Handle, nHeight)
}

// 窗格_取标题栏高度, 获取标题栏高度
func (p *Pane) GetCaptionHeight() int {
	return xc.XPane_GetCaptionHeight(p.Handle)
}

// 窗格_判断显示, 判断窗格是否显示
func (p *Pane) IsShowPane() bool {
	return xc.XPane_IsShowPane(p.Handle)
}

// 窗格_置大小, 设置窗格大小
// nWidth: 宽度.
// nHeight: 高度.
func (p *Pane) SetSize(nWidth int, nHeight int) int {
	return xc.XPane_SetSize(p.Handle, nWidth, nHeight)
}

// 窗格_取状态, 获取窗格停靠状态, 返回: Pane_State_
func (p *Pane) GetState() int {
	return xc.XPane_GetState(p.Handle)
}

// 窗格_取视图坐标, 获取窗格视图坐标
// pRect: 接收返回的坐标值.
func (p *Pane) GetViewRect(pRect *xc.RECT) int {
	return xc.XPane_GetViewRect(p.Handle, pRect)
}

// 窗格_隐藏
func (p *Pane) HidePane() int {
	return xc.XPane_HidePane(p.Handle)
}

// 窗格_显示
func (p *Pane) ShowPane() int {
	return xc.XPane_ShowPane(p.Handle)
}

// 窗格_停靠, 窗格停靠到码头
func (p *Pane) DockPane() int {
	return xc.XPane_DockPane(p.Handle)
}

// 窗格_锁定, 锁定窗格
func (p *Pane) LockPane() int {
	return xc.XPane_LockPane(p.Handle)
}

// 窗格_浮动
func (p *Pane) FloatPane() int {
	return xc.XPane_FloatPane(p.Handle)
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序
// hDraw: 图形绘制句柄.
func (p *Pane) DrawPane(hDraw int) int {
	return xc.XPane_DrawPane(p.Handle, hDraw)
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见
func (p *Pane) SetSelect() bool {
	return xc.XPane_SetSelect(p.Handle)
}
