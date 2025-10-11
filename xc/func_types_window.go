package xc

type XWM_WINDPROC func(message uint32, wParam, lParam uintptr, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER func(nTimerID uint, pbHandled *bool) int                          // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_SETFOCUS_ELE func(hEle int, pbHandled *bool) int                           // 窗口事件_置焦点元素. 指定元素获得焦点

type XWM_TRAYICON func(wParam, lParam uintptr, pbHandled *bool) int // 托盘图标事件.

type XWM_WINDPROC1 func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER1 func(hWindow int, nTimerID uint, pbHandled *bool) int                          // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_SETFOCUS_ELE1 func(hWindow int, hEle int, pbHandled *bool) int                           // 窗口事件_置焦点元素. 指定元素获得焦点

type XWM_FLOAT_PANE func(hFloatWnd int, hPane int, pbHandled *bool) int               // 浮动窗格.
type XWM_FLOAT_PANE1 func(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int // 浮动窗格.
type XWM_PAINT_END func(hDraw int, pbHandled *bool) int                               // 窗口绘制完成消息.
type XWM_PAINT_END1 func(hWindow int, hDraw int, pbHandled *bool) int                 // 窗口绘制完成消息.
type XWM_PAINT_DISPLAY func(pbHandled *bool) int                                      // 窗口绘制完成并且已经显示到屏幕.
type XWM_PAINT_DISPLAY1 func(hWindow int, pbHandled *bool) int                        // 窗口绘制完成并且已经显示到屏幕.
type XWM_DOCK_POPUP func(hWindowDock, hPane int, pbHandled *bool) int                 // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_DOCK_POPUP1 func(hWindow int, hWindowDock, hPane int, pbHandled *bool) int   // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_BODYVIEW_RECT func(width, height int32, pbHandled *bool) int                 // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
type XWM_BODYVIEW_RECT1 func(hWindow int, width, height int32, pbHandled *bool) int   // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG func(hFloatWnd int, hArray *[6]int, pbHandled *bool) int

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG1 func(hWindow int, hFloatWnd int, hArray *[6]int, pbHandled *bool) int

type WM_PAINT func(hDraw int, pbHandled *bool) int                                     // 窗口绘制消息.
type WM_PAINT1 func(hWindow int, hDraw int, pbHandled *bool) int                       // 窗口绘制消息.
type WM_CLOSE func(pbHandled *bool) int                                                // 窗口关闭消息.
type WM_CLOSE1 func(hWindow int, pbHandled *bool) int                                  // 窗口关闭消息.
type WM_DESTROY func(pbHandled *bool) int                                              // 窗口销毁消息.
type WM_DESTROY1 func(hWindow int, pbHandled *bool) int                                // 窗口销毁消息.
type WM_NCDESTROY func(pbHandled *bool) int                                            // 窗口非客户区销毁消息.
type WM_NCDESTROY1 func(hWindow int, pbHandled *bool) int                              // 窗口非客户区销毁消息.
type WM_MOUSEMOVE func(nFlags uint, pPt *POINT, pbHandled *bool) int                   // 窗口鼠标移动消息.
type WM_MOUSEMOVE1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int     // 窗口鼠标移动消息.
type WM_LBUTTONDOWN func(nFlags uint, pPt *POINT, pbHandled *bool) int                 // 窗口鼠标左键按下消息.
type WM_LBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int   // 窗口鼠标左键按下消息.
type WM_LBUTTONUP func(nFlags uint, pPt *POINT, pbHandled *bool) int                   // 窗口鼠标左键弹起消息.
type WM_LBUTTONUP1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int     // 窗口鼠标左键弹起消息.
type WM_RBUTTONDOWN func(nFlags uint, pPt *POINT, pbHandled *bool) int                 // 窗口鼠标右键按下消息.
type WM_RBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int   // 窗口鼠标右键按下消息.
type WM_RBUTTONUP func(nFlags uint, pPt *POINT, pbHandled *bool) int                   // 窗口鼠标右键弹起消息.
type WM_RBUTTONUP1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int     // 窗口鼠标右键弹起消息.
type WM_LBUTTONDBLCLK func(nFlags uint, pPt *POINT, pbHandled *bool) int               // 窗口鼠标左键双击消息.
type WM_LBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int // 窗口鼠标左键双击消息.
type WM_RBUTTONDBLCLK func(nFlags uint, pPt *POINT, pbHandled *bool) int               // 窗口鼠标右键双击消息.
type WM_RBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int // 窗口鼠标右键双击消息.
type WM_MOUSEWHEEL func(nFlags uint, pPt *POINT, pbHandled *bool) int                  // 窗口鼠标滚轮滚动消息.
type WM_MOUSEWHEEL1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int    // 窗口鼠标滚轮滚动消息.
type WM_EXITSIZEMOVE func(pbHandled *bool) int                                         // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_EXITSIZEMOVE1 func(hWindow int, pbHandled *bool) int                           // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_MOUSEHOVER func(nFlags uint, pPt *POINT, pbHandled *bool) int                  // 窗口鼠标进入消息.
type WM_MOUSEHOVER1 func(hWindow int, nFlags uint, pPt *POINT, pbHandled *bool) int    // 窗口鼠标进入消息.
type WM_MOUSELEAVE func(pbHandled *bool) int                                           // 窗口鼠标离开消息.
type WM_MOUSELEAVE1 func(hWindow int, pbHandled *bool) int                             // 窗口鼠标离开消息.
type WM_SIZE func(nFlags uint, pPt *SIZE, pbHandled *bool) int                         // 窗口大小改变消息.
type WM_SIZE1 func(hWindow int, nFlags uint, pPt *SIZE, pbHandled *bool) int           // 窗口大小改变消息.
type WM_TIMER func(nIDEvent uint, pbHandled *bool) int                                 // 窗口定时器消息.
type WM_TIMER1 func(hWindow int, nIDEvent uint, pbHandled *bool) int                   // 窗口定时器消息.
type WM_SETFOCUS func(pbHandled *bool) int                                             // 窗口获得焦点.
type WM_SETFOCUS1 func(hWindow int, pbHandled *bool) int                               // 窗口获得焦点.
type WM_KILLFOCUS func(pbHandled *bool) int                                            // 窗口失去焦点.
type WM_KILLFOCUS1 func(hWindow int, pbHandled *bool) int                              // 窗口失去焦点.
type WM_KEYDOWN func(wParam, lParam uintptr, pbHandled *bool) int                      // 窗口键盘按键消息.
type WM_KEYDOWN1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int        // 窗口键盘按键消息.
type WM_CAPTURECHANGED func(hWnd uintptr, pbHandled *bool) int                         // 窗口鼠标捕获改变消息.
type WM_CAPTURECHANGED1 func(hWindow int, hWnd uintptr, pbHandled *bool) int           // 窗口鼠标捕获改变消息.
type WM_SETCURSOR func(wParam, lParam uintptr, pbHandled *bool) int                    // 窗口设置鼠标光标.
type WM_SETCURSOR1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int      // 窗口设置鼠标光标.
type WM_CHAR func(wParam, lParam uintptr, pbHandled *bool) int                         // 窗口字符消息.
type WM_CHAR1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int           // 窗口字符消息.
type WM_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int                         // 拖动文件到窗口.
type WM_DROPFILES1 func(hWindow int, hDropInfo uintptr, pbHandled *bool) int           // 拖动文件到窗口.

type XWM_MENU_POPUP func(hMenu int, pbHandled *bool) int                                                      // 菜单弹出.
type XWM_MENU_POPUP1 func(hWindow int, hMenu int, pbHandled *bool) int                                        // 菜单弹出.
type XWM_MENU_POPUP_WND func(hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int                           // 菜单弹出窗口.
type XWM_MENU_POPUP_WND1 func(hWindow int, hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XWM_MENU_SELECT func(nID int32, pbHandled *bool) int                                                     // 菜单选择.
type XWM_MENU_SELECT1 func(hWindow int, nID int32, pbHandled *bool) int                                       // 菜单选择.
type XWM_MENU_EXIT func(pbHandled *bool) int                                                                  // 菜单退出.
type XWM_MENU_EXIT1 func(hWindow int, pbHandled *bool) int                                                    // 菜单退出.
type XWM_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int               // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAW_BACKGROUND1 func(hWindow int, hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAWITEM func(hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int                            // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
type XWM_MENU_DRAWITEM1 func(hWindow int, hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
