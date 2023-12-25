package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 框架窗口_创建, 返回GUI库窗口资源句柄.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口真实句柄.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func XFrameWnd_Create(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xFrameWnd_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 框架窗口_创建扩展, 返回GUI库窗口资源句柄.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func XFrameWnd_CreateEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xFrameWnd_CreateEx.Call(uintptr(dwExStyle), uintptr(dwStyle), common.StrPtr(lpClassName), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 框架窗口_取布局区域坐标, 用来布局窗格的区域坐标, 不包含码头.
//
// hWindow: 窗口句柄.
//
// pRect: 返回坐标.
func XFrameWnd_GetLayoutAreaRect(hWindow int, pRect *RECT) int {
	r, _, _ := xFrameWnd_GetLayoutAreaRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 框架窗口_置视图, 设置窗格组TabBar高度.
//
// hWindow: 窗口句柄.
//
// hEle: 元素句柄.
func XFrameWnd_SetView(hWindow int, hEle int) int {
	r, _, _ := xFrameWnd_SetView.Call(uintptr(hWindow), uintptr(hEle))
	return int(r)
}

// 框架窗口_置窗格分隔条颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色值.
func XFrameWnd_SetPaneSplitBarColor(hWindow int, color int) int {
	r, _, _ := xFrameWnd_SetPaneSplitBarColor.Call(uintptr(hWindow), uintptr(color))
	return int(r)
}

// 框架窗口_置TabBar条高度, 设置窗格组TabBar高度.
//
// hWindow: 窗口句柄.
//
// nHeight: 高度.
func XFrameWnd_SetTabBarHeight(hWindow int, nHeight int32) int {
	r, _, _ := xFrameWnd_SetTabBarHeight.Call(uintptr(hWindow), uintptr(nHeight))
	return int(r)
}

// 框架窗口_保存布局到文件, 保存布局信息到文件.
//
// hWindow: 窗口句柄.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func XFrameWnd_SaveLayoutToFile(hWindow int, pFileName string) bool {
	r, _, _ := xFrameWnd_SaveLayoutToFile.Call(uintptr(hWindow), common.StrPtr(pFileName))
	return r != 0
}

// 框架窗口_加载布局信息文件, 加载布局信息文件.
//
// hWindow: 窗口句柄.
//
// aPaneList: 窗格句柄数组.
//
// nPaneCount: 窗格数量.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func XFrameWnd_LoadLayoutFile(hWindow int, aPaneList []int, nPaneCount int32, pFileName string) bool {
	r, _, _ := xFrameWnd_LoadLayoutFile.Call(uintptr(hWindow), uintptr(unsafe.Pointer(&aPaneList[0])), uintptr(nPaneCount), common.StrPtr(pFileName))
	return r != 0
}

// 框架窗口_添加窗格, 添加窗格到框架窗口.
//
// hWindow: 窗口句柄.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
//
// align: 对齐方式, Pane_Align_.
func XFrameWnd_AddPane(hWindow int, hPaneDest int, hPaneNew int, align xcc.Pane_Align_) bool {
	r, _, _ := xFrameWnd_AddPane.Call(uintptr(hWindow), uintptr(hPaneDest), uintptr(hPaneNew), uintptr(align))
	return r != 0
}

// 框架窗口_合并窗格.
//
// hWindow: 窗口句柄.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
func XFrameWnd_MergePane(hWindow int, hPaneDest int, hPaneNew int) bool {
	r, _, _ := xFrameWnd_MergePane.Call(uintptr(hWindow), uintptr(hPaneDest), uintptr(hPaneNew))
	return r != 0
}

// 框架窗口_附加窗口, 返回窗口资源句柄.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func XFrameWnd_Attach(hWnd uintptr, XCStyle int) int {
	r, _, _ := xFrameWnd_Attach.Call(hWnd, uintptr(XCStyle))
	return int(r)
}

// 框架窗口_取拖动浮动窗格停留位置标识, 返回 拖动窗格 所处框架窗口单元格标识: xcc.FrameWnd_Cell_Type_.
//
// hWindow: 窗口句柄.
func XFrameWnd_GetDragFloatWndTopFlag(hWindow int) xcc.FrameWnd_Cell_Type_ {
	r, _, _ := xFrameWnd_GetDragFloatWndTopFlag.Call(uintptr(hWindow))
	return xcc.FrameWnd_Cell_Type_(r)
}

// 框架窗口_取主视图坐标. 获取框架窗口主视图区域坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 返回坐标.
func XFrameWnd_GetViewRect(hWindow int, pRect *RECT) {
	xFrameWnd_GetViewRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
}

// 框架窗口_置窗格分隔条宽度.
//
// hWindow: 窗口句柄.
//
// nWidth: 宽度.
func XFrameWnd_SetPaneSplitBarWidth(hWindow int, nWidth int32) {
	xFrameWnd_SetPaneSplitBarWidth.Call(uintptr(hWindow), uintptr(nWidth))
}

// 框架窗口_取窗格分隔条宽度.
//
// hWindow: 窗口句柄.
func XFrameWnd_GetPaneSplitBarWidth(hWindow int) int32 {
	r, _, _ := xFrameWnd_GetPaneSplitBarWidth.Call(uintptr(hWindow))
	return int32(r)
}
