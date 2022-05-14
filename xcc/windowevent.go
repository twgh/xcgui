package xcc

// 窗口事件
// XWM_

const (
	XWM_WINDPROC             WM_ = 0x7000 + 2  // 窗口消息过程
	XWM_XC_TIMER             WM_ = 0x7000 + 5  // 炫彩定时器, 非系统定时器, 注册消息 XWM_TIMER 接收
	XWM_MENU_POPUP           WM_ = 0x7000 + 11 // 菜单弹出
	XWM_MENU_POPUP_WND       WM_ = 0x7000 + 12 // 菜单弹出窗口
	XWM_MENU_SELECT          WM_ = 0x7000 + 13 // 菜单选择
	XWM_MENU_EXIT            WM_ = 0x7000 + 14 // 菜单退出
	XWM_MENU_DRAW_BACKGROUND WM_ = 0x7000 + 15 // 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground
	XWM_MENU_DRAWITEM        WM_ = 0x7000 + 16 // 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem
	XWM_FLOAT_PANE           WM_ = 0x7000 + 18 // 浮动窗格
	XWM_PAINT_END            WM_ = 0x7000 + 19 // 窗口绘制完成消息
	XWM_PAINT_DISPLAY        WM_ = 0x7000 + 20 // 窗口绘制完成并且已经显示到屏幕
	XWM_DOCK_POPUP           WM_ = 0x7000 + 21 // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格

	// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
	//
	// hFloatWnd: 拖动的浮动窗口句柄.
	//
	// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
	XWM_FLOATWND_DRAG WM_ = 0x7000 + 22
)

// 窗口事件
type WM_ uint32

const (
	WM_PAINT          WM_ = 15  // 窗口绘制消息
	WM_CLOSE          WM_ = 16  // 窗口关闭消息.
	WM_DESTROY        WM_ = 2   // 窗口销毁消息.
	WM_NCDESTROY      WM_ = 130 // 窗口非客户区销毁消息.
	WM_MOUSEMOVE      WM_ = 512 // 窗口鼠标移动消息.
	WM_LBUTTONDOWN    WM_ = 513 // 窗口鼠标左键按下消息
	WM_LBUTTONUP      WM_ = 514 // 窗口鼠标左键弹起消息.
	WM_RBUTTONDOWN    WM_ = 516 // 窗口鼠标右键按下消息.
	WM_RBUTTONUP      WM_ = 517 // 窗口鼠标右键弹起消息.
	WM_LBUTTONDBLCLK  WM_ = 515 // 窗口鼠标左键双击消息.
	WM_RBUTTONDBLCLK  WM_ = 518 // 窗口鼠标右键双击消息.
	WM_MOUSEWHEEL     WM_ = 522 // 窗口鼠标滚轮滚动消息.
	WM_EXITSIZEMOVE   WM_ = 562 // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
	WM_MOUSEHOVER     WM_ = 673 // 窗口鼠标进入消息
	WM_MOUSELEAVE     WM_ = 675 // 窗口鼠标离开消息.
	WM_SIZE           WM_ = 5   // 窗口大小改变消息.
	WM_TIMER          WM_ = 275 // 窗口定时器消息.
	WM_SETFOCUS       WM_ = 7   // 窗口获得焦点.
	WM_KILLFOCUS      WM_ = 8   // 窗口失去焦点.
	WM_KEYDOWN        WM_ = 256 // 窗口键盘按键消息.
	WM_CAPTURECHANGED WM_ = 533 // 窗口鼠标捕获改变消息.
	WM_SETCURSOR      WM_ = 32  // 窗口设置鼠标光标.
	WM_CHAR           WM_ = 258 // 窗口字符消息.
	WM_DROPFILES      WM_ = 563 // 拖动文件到窗口.
)
