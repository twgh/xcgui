package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 窗格_创建, 创建窗格元素, 返回元素句柄.
//
// pName: 窗格标题.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// hFrameWnd: 框架窗口.
func XPane_Create(pName string, nWidth int, nHeight int, hFrameWnd int) int {
	r, _, _ := xPane_Create.Call(common.StrPtr(pName), uintptr(nWidth), uintptr(nHeight), uintptr(hFrameWnd))
	return int(r)
}

// 窗格_置视图, 设置窗格视图元素.
//
// hEle: 元素句柄.
//
// hView: 绑定视图元素.
func XPane_SetView(hEle int, hView int) int {
	r, _, _ := xPane_SetView.Call(uintptr(hEle), uintptr(hView))
	return int(r)
}

// 窗格_置标题, 设置标题文本.
//
// hEle: 元素句柄.
//
// pTitle: 文本内容.
func XPane_SetTitle(hEle int, pTitle string) int {
	r, _, _ := xPane_SetTitle.Call(uintptr(hEle), common.StrPtr(pTitle))
	return int(r)
}

// 窗格_取标题, 获取标题文本.
//
// hEle: 元素句柄.
func XPane_GetTitle(hEle int) string {
	r, _, _ := xPane_GetTitle.Call(uintptr(hEle))
	return common.UintPtrToString(r)
}

// 窗格_置标题栏高度, 设置标题栏高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
func XPane_SetCaptionHeight(hEle int, nHeight int) int {
	r, _, _ := xPane_SetCaptionHeight.Call(uintptr(hEle), uintptr(nHeight))
	return int(r)
}

// 窗格_取标题栏高度, 获取标题栏高度.
//
// hEle: 元素句柄.
func XPane_GetCaptionHeight(hEle int) int {
	r, _, _ := xPane_GetCaptionHeight.Call(uintptr(hEle))
	return int(r)
}

// 窗格_判断显示, 判断窗格是否显示.
//
// hEle: 元素句柄.
func XPane_IsShowPane(hEle int) bool {
	r, _, _ := xPane_IsShowPane.Call(uintptr(hEle))
	return r != 0
}

// 窗格_置大小, 设置窗格大小.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XPane_SetSize(hEle int, nWidth int, nHeight int) int {
	r, _, _ := xPane_SetSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 窗格_取状态, 获取窗格停靠状态, 返回: Pane_State_.
//
// hEle: 元素句柄.
func XPane_GetState(hEle int) xcc.Pane_State_ {
	r, _, _ := xPane_GetState.Call(uintptr(hEle))
	return xcc.Pane_State_(r)
}

// 窗格_取视图坐标, 获取窗格视图坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回的坐标值.
func XPane_GetViewRect(hEle int, pRect *RECT) int {
	r, _, _ := xPane_GetViewRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// XPane_HidePane 窗格_隐藏.
//
//	@param hEle 元素句柄.
//	@param bGroupActivate 当为窗格组成员时, 延迟处理窗格组成员激活的切换.
//	@return int
func XPane_HidePane(hEle int, bGroupActivate bool) int {
	r, _, _ := xPane_HidePane.Call(uintptr(hEle), common.BoolPtr(bGroupActivate))
	return int(r)
}

// XPane_ShowPane 窗格_显示.
//
//	@param hEle 元素句柄.
//	@param bGroupActivate 如果是窗格组成员, 那么窗格组切换当前窗格为显示状态.
//	@return int
func XPane_ShowPane(hEle int, bGroupActivate bool) int {
	r, _, _ := xPane_ShowPane.Call(uintptr(hEle), common.BoolPtr(bGroupActivate))
	return int(r)
}

// 窗格_停靠, 窗格停靠到码头.
//
// hEle: 元素句柄.
func XPane_DockPane(hEle int) int {
	r, _, _ := xPane_DockPane.Call(uintptr(hEle))
	return int(r)
}

// 窗格_锁定, 锁定窗格.
//
// hEle: 元素句柄.
func XPane_LockPane(hEle int) int {
	r, _, _ := xPane_LockPane.Call(uintptr(hEle))
	return int(r)
}

// 窗格_浮动.
//
// hEle: 元素句柄.
func XPane_FloatPane(hEle int) int {
	r, _, _ := xPane_FloatPane.Call(uintptr(hEle))
	return int(r)
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.
func XPane_DrawPane(hEle int, hDraw int) int {
	r, _, _ := xPane_DrawPane.Call(uintptr(hEle), uintptr(hDraw))
	return int(r)
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见.
//
// hEle: 元素句柄.
func XPane_SetSelect(hEle int) bool {
	r, _, _ := xPane_SetSelect.Call(uintptr(hEle))
	return r != 0
}

// 窗格_是否激活. 判断窗格是否激活, 当为组成员时有效.
//
// hEle: 元素句柄.
func XPane_IsGroupActivate(hEle int) bool {
	r, _, _ := xPane_IsGroupActivate.Call(uintptr(hEle))
	return r != 0
}
