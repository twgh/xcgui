package window

import (
	xc "github.com/twgh/xcgui"
)

type FrameWindow struct {
	WindowPublic

	HWindow int
}

// 框架窗口_创建, 返回GUI库窗口资源句柄
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// pTitle: 窗口标题.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式, Xc_Window_Style_
func NewFreamWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle int) *FrameWindow {
	p := &FrameWindow{
		HWindow: xc.XFrameWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle),
	}
	p.HWindow_ = p.HWindow
	return p
}

// 框架窗口_创建扩展, 返回GUI库窗口资源句柄
// dwExStyle: 窗口扩展样式.
// lpClassName: 窗口类名.
// lpWindowName: 窗口名.
// dwStyle: 窗口样式
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式, Xc_Window_Style_
func NewFreamWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *FrameWindow {
	p := &FrameWindow{
		HWindow: xc.XFrameWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle),
	}
	p.HWindow_ = p.HWindow
	return p
}

// 框架窗口_取布局区域坐标, 用来布局窗格的区域坐标, 不包含码头
// pRect: 返回坐标.
func (fw *FrameWindow) GetLayoutAreaRect(pRect *xc.RECT) int {
	return xc.XFrameWnd_GetLayoutAreaRect(fw.HWindow, pRect)
}

// 框架窗口_置视图, 设置窗格组TabBar高度
// hEle: 元素句柄.
func (fw *FrameWindow) SetView(hEle int) int {
	return xc.XFrameWnd_SetView(fw.HWindow, hEle)
}

// 框架窗口_置窗格分隔条颜色
// color: RGB颜色值
// alpha: 透明度
func (fw *FrameWindow) SetPaneSplitBarColor(color int, alpha uint8) int {
	return xc.XFrameWnd_SetPaneSplitBarColor(fw.HWindow, color, alpha)
}

// 框架窗口_置TabBar条高度, 设置窗格组TabBar高度
// nHeight: 高度
func (fw *FrameWindow) SetTabBarHeight(nHeight int) int {
	return xc.XFrameWnd_SetTabBarHeight(fw.HWindow, nHeight)
}

// 框架窗口_保存布局到文件, 保存布局信息到文件
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) SaveLayoutToFile(pFileName string) bool {
	return xc.XFrameWnd_SaveLayoutToFile(fw.HWindow, pFileName)
}

// 框架窗口_加载布局信息文件, 加载布局信息文件
// aPaneList: 窗格句柄数组.
// nPaneCount: 窗格数量.
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) LoadLayoutFile(aPaneList int, nPaneCount int, pFileName string) bool {
	return xc.XFrameWnd_LoadLayoutFile(fw.HWindow, aPaneList, nPaneCount, pFileName)
}

// 框架窗口_添加窗格, 添加窗格到框架窗口
// hPaneDest: 目标窗格.
// hPaneNew: 当前窗格.
// align: 对齐方式, Pane_Align_
func (fw *FrameWindow) AddPane(hPaneDest int, hPaneNew int, align int) bool {
	return xc.XFrameWnd_AddPane(fw.HWindow, hPaneDest, hPaneNew, align)
}

// 框架窗口_合并窗格
// hPaneDest: 目标窗格.
// hPaneNew: 当前窗格.
func (fw *FrameWindow) MergePane(hPaneDest int, hPaneNew int) bool {
	return xc.XFrameWnd_MergePane(fw.HWindow, hPaneDest, hPaneNew)
}
