package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Pane Pane元素.
type Pane struct {
	Element
}

// 窗格_创建, 创建窗格元素, 返回元素句柄.
//
//	pName: 窗格标题.
//
//	nWidth: 宽度.
//
//	nHeight: 高度.
//
//	hFrameWnd: 框架窗口.
func NewPane(pName string, nWidth, nHeight int32, hFrameWnd int) *Pane {
	p := &Pane{}
	p.SetHandle(xc.XPane_Create(pName, nWidth, nHeight, hFrameWnd))
	return p
}

// 从句柄创建对象.
func NewPaneByHandle(handle int) *Pane {
	p := &Pane{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewPaneByName(name string) *Pane {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewPaneByUID(nUID int32) *Pane {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewPaneByUIDName(name string) *Pane {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 窗格_置视图, 设置窗格视图元素.
//
//	hView: 绑定视图元素.
func (p *Pane) SetView(hView int) *Pane {
	xc.XPane_SetView(p.Handle, hView)
	return p
}

// 窗格_置标题, 设置标题文本.
//
//	pTitle: 文本内容.
func (p *Pane) SetTitle(pTitle string) *Pane {
	xc.XPane_SetTitle(p.Handle, pTitle)
	return p
}

// 窗格_取标题, 获取标题文本.
func (p *Pane) GetTitle() string {
	return xc.XPane_GetTitle(p.Handle)
}

// 窗格_置标题栏高度, 设置标题栏高度.
//
//	nHeight: 高度.
func (p *Pane) SetCaptionHeight(nHeight int32) *Pane {
	xc.XPane_SetCaptionHeight(p.Handle, nHeight)
	return p
}

// 窗格_取标题栏高度, 获取标题栏高度.
func (p *Pane) GetCaptionHeight() int32 {
	return xc.XPane_GetCaptionHeight(p.Handle)
}

// 窗格_判断显示, 判断窗格是否显示.
func (p *Pane) IsShowPane() bool {
	return xc.XPane_IsShowPane(p.Handle)
}

// 窗格_置大小, 设置窗格大小.
//
//	nWidth: 宽度.
//
//	nHeight: 高度.
func (p *Pane) SetSize(nWidth, nHeight int32) *Pane {
	xc.XPane_SetSize(p.Handle, nWidth, nHeight)
	return p
}

// 窗格_取状态, 获取窗格停靠状态, 返回: Pane_State_.
func (p *Pane) GetState() xcc.Pane_State_ {
	return xc.XPane_GetState(p.Handle)
}

// 窗格_取视图坐标, 获取窗格视图坐标.
//
//	pRect: 接收返回的坐标值.
func (p *Pane) GetViewRect(pRect *xc.RECT) *Pane {
	xc.XPane_GetViewRect(p.Handle, pRect)
	return p
}

// HidePane 窗格_隐藏.
//
//	bGroupActivate: 当为窗格组成员时, 延迟处理窗格组成员激活的切换.
func (p *Pane) HidePane(bGroupActivate bool) *Pane {
	xc.XPane_HidePane(p.Handle, bGroupActivate)
	return p
}

// ShowPane 窗格_显示.
//
//	bGroupActivate: 如果是窗格组成员, 那么窗格组切换当前窗格为显示状态.
func (p *Pane) ShowPane(bGroupActivate bool) *Pane {
	xc.XPane_ShowPane(p.Handle, bGroupActivate)
	return p
}

// 窗格_停靠, 窗格停靠到码头.
func (p *Pane) DockPane() *Pane {
	xc.XPane_DockPane(p.Handle)
	return p
}

// 窗格_锁定, 锁定窗格.
func (p *Pane) LockPane() *Pane {
	xc.XPane_LockPane(p.Handle)
	return p
}

// 窗格_浮动.
func (p *Pane) FloatPane() *Pane {
	xc.XPane_FloatPane(p.Handle)
	return p
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序.
//
//	hDraw: 图形绘制句柄.
func (p *Pane) DrawPane(hDraw int) *Pane {
	xc.XPane_DrawPane(p.Handle, hDraw)
	return p
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见.
func (p *Pane) SetSelect() bool {
	return xc.XPane_SetSelect(p.Handle)
}

// 窗格_是否激活. 判断窗格是否激活, 当为组成员时有效.
func (p *Pane) IsGroupActivate() bool {
	return xc.XPane_IsGroupActivate(p.Handle)
}
