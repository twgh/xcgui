package xc

import "github.com/twgh/xcgui/xcc"

type XE_ELEPROCE func(nEvent uint32, wParam, lParam uintptr, pbHandled *bool) int            // 元素处理过程事件.
type XE_ELEPROCE1 func(hEle int, nEvent uint32, wParam, lParam uintptr, pbHandled *bool) int // 元素处理过程事件.
type XE_PAINT func(hDraw int, pbHandled *bool) int                                           // 元素绘制事件.
type XE_PAINT1 func(hEle int, hDraw int, pbHandled *bool) int                                // 元素绘制事件.
type XE_PAINT_END func(hDraw int, pbHandled *bool) int                                       // 该元素及子元素绘制完成事件.启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
type XE_PAINT_END1 func(hEle int, hDraw int, pbHandled *bool) int                            // 该元素及子元素绘制完成事件.启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
type XE_PAINT_SCROLLVIEW func(hDraw int, pbHandled *bool) int                                // 滚动视图绘制事件.
type XE_PAINT_SCROLLVIEW1 func(hEle int, hDraw int, pbHandled *bool) int                     // 滚动视图绘制事件.
type XE_MOUSEMOVE func(nFlags int, pPt *POINT, pbHandled *bool) int                          // 元素鼠标移动事件.
type XE_MOUSEMOVE1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int               // 元素鼠标移动事件.
type XE_MOUSESTAY func(pbHandled *bool) int                                                  // 元素鼠标进入事件.
type XE_MOUSESTAY1 func(hEle int, pbHandled *bool) int                                       // 元素鼠标进入事件.
type XE_MOUSEHOVER func(nFlags int, pPt *POINT, pbHandled *bool) int                         // 元素鼠标悬停事件.
type XE_MOUSEHOVER1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int              // 元素鼠标悬停事件.
type XE_MOUSELEAVE func(hEleStay int, pbHandled *bool) int                                   // 元素鼠标离开事件.
type XE_MOUSELEAVE1 func(hEle int, hEleStay int, pbHandled *bool) int                        // 元素鼠标离开事件.
type XE_MOUSEWHEEL func(nFlags int, pPt *POINT, pbHandled *bool) int                         // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL. flags: 见MSDN中 WM_MOUSEWHEEL 消息 wParam 参数说明.
type XE_MOUSEWHEEL1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int              // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL. flags: 见MSDN中 WM_MOUSEWHEEL 消息 wParam 参数说明.
type XE_LBUTTONDOWN func(nFlags int, pPt *POINT, pbHandled *bool) int                        // 鼠标左键按下事件.
type XE_LBUTTONDOWN1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int             // 鼠标左键按下事件.
type XE_LBUTTONUP func(nFlags int, pPt *POINT, pbHandled *bool) int                          // 鼠标左键弹起事件.
type XE_LBUTTONUP1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int               // 鼠标左键弹起事件.
type XE_RBUTTONDOWN func(nFlags int, pPt *POINT, pbHandled *bool) int                        // 鼠标右键按下事件.
type XE_RBUTTONDOWN1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int             // 鼠标右键按下事件.
type XE_RBUTTONUP func(nFlags int, pPt *POINT, pbHandled *bool) int                          // 鼠标右键弹起事件.
type XE_RBUTTONUP1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int               // 鼠标右键弹起事件.
type XE_LBUTTONDBCLICK func(nFlags int, pPt *POINT, pbHandled *bool) int                     // 鼠标左键双击事件.
type XE_LBUTTONDBCLICK1 func(hEle int, nFlags int, pPt *POINT, pbHandled *bool) int          // 鼠标左键双击事件.
type XE_XC_TIMER func(nTimerID int, pbHandled *bool) int                                     // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_XC_TIMER1 func(hEle int, nTimerID int, pbHandled *bool) int                          // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_ADJUSTLAYOUT func(nFlags int32, nAdjustNo uint32, pbHandled *bool) int               // 调整布局事件. 暂停使用.
type XE_ADJUSTLAYOUT1 func(hEle int, nFlags int32, nAdjustNo uint32, pbHandled *bool) int    // 调整布局事件. 暂停使用.
type XE_TOOLTIP_POPUP func(hWindow int, pText uintptr, pbHandled *bool) int                  // 元素工具提示弹出事件.
type XE_TOOLTIP_POPUP1 func(hEle int, hWindow int, pText uintptr, pbHandled *bool) int       // 元素工具提示弹出事件1.

// 调整布局完成事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 调整布局完成事件.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SETFOCUS func(pbHandled *bool) int               // 元素获得焦点事件.
type XE_SETFOCUS1 func(hEle int, pbHandled *bool) int    // 元素获得焦点事件.
type XE_KILLFOCUS func(pbHandled *bool) int              // 元素失去焦点事件.
type XE_KILLFOCUS1 func(hEle int, pbHandled *bool) int   // 元素失去焦点事件.
type XE_DESTROY func(pbHandled *bool) int                // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY1 func(hEle int, pbHandled *bool) int     // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY_END func(pbHandled *bool) int            // 元素销毁完成事件. 在销毁子对象之后触发.
type XE_DESTROY_END1 func(hEle int, pbHandled *bool) int // 元素销毁完成事件. 在销毁子对象之后触发.

// 元素大小改变事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 元素大小改变事件1.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SHOW func(bShow bool, pbHandled *bool) int                              // 元素显示隐藏事件.
type XE_SHOW1 func(hEle int, bShow bool, pbHandled *bool) int                   // 元素显示隐藏事件.
type XE_SETFONT func(pbHandled *bool) int                                       // 元素设置字体事件.
type XE_SETFONT1 func(hEle int, pbHandled *bool) int                            // 元素设置字体事件.
type XE_KEYDOWN func(wParam, lParam uintptr, pbHandled *bool) int               // 元素按键按下事件.
type XE_KEYDOWN1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int    // 元素按键按下事件.
type XE_KEYUP func(wParam, lParam uintptr, pbHandled *bool) int                 // 元素按键弹起事件.
type XE_KEYUP1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int      // 元素按键弹起事件.
type XE_SYSKEYDOWN1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int // 元素系统按键按下事件.
type XE_SYSKEYUP1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int   // 元素系统按键弹起事件

type XE_CHAR func(wParam, lParam uintptr, pbHandled *bool) int                 // 通过TranslateMessage函数翻译的字符事件.
type XE_CHAR1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int      // 通过TranslateMessage函数翻译的字符事件.
type XE_SETCAPTURE func(pbHandled *bool) int                                   // 元素设置鼠标捕获.
type XE_SETCAPTURE1 func(hEle int, pbHandled *bool) int                        // 元素设置鼠标捕获.
type XE_KILLCAPTURE func(pbHandled *bool) int                                  // 元素失去鼠标捕获.
type XE_KILLCAPTURE1 func(hEle int, pbHandled *bool) int                       // 元素失去鼠标捕获.
type XE_SETCURSOR func(wParam, lParam uintptr, pbHandled *bool) int            // 设置鼠标光标.
type XE_SETCURSOR1 func(hEle int, wParam, lParam uintptr, pbHandled *bool) int // 设置鼠标光标.
type XE_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int                 // 文件拖放事件, 需先启用: xc.XWnd_EnableDragFiles.
type XE_DROPFILES1 func(hEle int, hDropInfo uintptr, pbHandled *bool) int      // 文件拖放事件, 需先启用: xc.XWnd_EnableDragFiles.

type XE_MENU_SELECT func(nID int32, pbHandled *bool) int                                       // 弹出菜单项选择事件.
type XE_MENU_POPUP func(HMENUX int, pbHandled *bool) int                                       // 菜单弹出.
type XE_MENU_EXIT func(pbHandled *bool) int                                                    // 弹出菜单退出事件.
type XE_MENU_POPUP_WND func(hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
type XE_MENU_DRAWITEM func(hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.

type XE_MENU_SELECT1 func(hEle int, nID int32, pbHandled *bool) int                                       // 弹出菜单项选择事件.
type XE_MENU_POPUP1 func(hEle int, HMENUX int, pbHandled *bool) int                                       // 菜单弹出.
type XE_MENU_EXIT1 func(hEle int, pbHandled *bool) int                                                    // 弹出菜单退出事件.
type XE_MENU_POPUP_WND1 func(hEle int, hMenu int, pInfo *Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND1 func(hEle int, hDraw int, pInfo *Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
type XE_MENU_DRAWITEM1 func(hEle int, hDraw int, pInfo *Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.

type XE_BNCLICK func(pbHandled *bool) int                              // 按钮点击事件.
type XE_BNCLICK1 func(hEle int, pbHandled *bool) int                   // 按钮点击事件.
type XE_BUTTON_CHECK func(bCheck bool, pbHandled *bool) int            // 按钮选中事件.
type XE_BUTTON_CHECK1 func(hEle int, bCheck bool, pbHandled *bool) int // 按钮选中事件.

type XE_COMBOBOX_SELECT_END func(iItem int32, pbHandled *bool) int                          // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT_END1 func(hEle int, iItem int32, pbHandled *bool) int               // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT func(iItem int32, pbHandled *bool) int                              // 组合框下拉列表项选择事件.
type XE_COMBOBOX_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                   // 组合框下拉列表项选择事件.
type XE_COMBOBOX_POPUP_LIST func(hWindow int, hListBox int, pbHandled *bool) int            // 组合框下拉列表弹出事件.
type XE_COMBOBOX_POPUP_LIST1 func(hEle int, hWindow int, hListBox int, pbHandled *bool) int // 组合框下拉列表弹出事件.
type XE_COMBOBOX_EXIT_LIST func(pbHandled *bool) int                                        // 组合框下拉列表退出事件.
type XE_COMBOBOX_EXIT_LIST1 func(hEle int, pbHandled *bool) int                             // 组合框下拉列表退出事件.

type XE_DATETIME_CHANGE func(pbHandled *bool) int                                                     // 日期时间元素,内容改变事件.
type XE_DATETIME_CHANGE1 func(hEle int, pbHandled *bool) int                                          // 日期时间元素,内容改变事件.
type XE_DATETIME_POPUP_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int            // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_POPUP_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_EXIT_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int             // 日期时间元素,弹出的月历卡片退出事件.
type XE_DATETIME_EXIT_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int  // 日期时间元素,弹出的月历卡片退出事件.

type XE_EDIT_SET func(pbHandled *bool) int                                       // 元素事件_编辑框设置.
type XE_EDIT_SET1 func(hEle int, pbHandled *bool) int                            // 元素事件_编辑框设置.
type XE_EDIT_DRAWROW func(hDraw int, iRow int32, pbHandled *bool) int            // 暂未使用.
type XE_EDIT_DRAWROW1 func(hEle int, hDraw int, iRow int32, pbHandled *bool) int // 暂未使用.
type XE_EDIT_CHANGED func(pbHandled *bool) int                                   // 编辑框_内容被改变.
type XE_EDIT_CHANGED1 func(hEle int, pbHandled *bool) int                        // 编辑框_内容被改变.
type XE_EDIT_POS_CHANGED func(iPos int32, pbHandled *bool) int                   // 编辑框_光标位置_被改变.
type XE_EDIT_POS_CHANGED1 func(hEle int, iPos int32, pbHandled *bool) int        // 编辑框_光标位置_被改变.
type XE_EDIT_STYLE_CHANGED func(iStyle int32, pbHandled *bool) int               // 编辑框_样式_被改变.
type XE_EDIT_STYLE_CHANGED1 func(hEle int, iStyle int32, pbHandled *bool) int    // 编辑框_样式_被改变.
type XE_EDIT_ENTER_GET_TABALIGN func(pbHandled *bool) int                        // 回车TAB对齐,返回需要TAB数量.
type XE_EDIT_ENTER_GET_TABALIGN1 func(hEle int, pbHandled *bool) int             // 回车TAB对齐,返回需要TAB数量.
// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
type XE_EDIT_ROW_CHANGED func(iRow int32, nChangeRows int32, pbHandled *bool) int

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
type XE_EDIT_ROW_CHANGED1 func(hEle int, iRow int32, nChangeRows int32, pbHandled *bool) int
type XE_EDIT_SWAPROW func(iRow, bArrowUp int32, pbHandled *bool) int            // 元素事件_交换行
type XE_EDIT_SWAPROW1 func(hEle int, iRow, bArrowUp int32, pbHandled *bool) int // 元素事件_交换行
type XE_EDIT_COLOR_CHANGE func(color uint32, pbHandled *bool) int               // 编辑框_颜色被改变
type XE_EDIT_COLOR_CHANGE1 func(hEle int, color uint32, pbHandled *bool) int    // 编辑框_颜色被改变

type XE_EDITOR_MODIFY_ROWS func(iRow int32, nRows int32, pbHandled *bool) int                 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_MODIFY_ROWS1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int      // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_SETBREAKPOINT func(iRow int32, bCheck bool, pbHandled *bool) int               // 代码编辑框_设置断点.
type XE_EDITOR_SETBREAKPOINT1 func(hEle int, iRow int32, bCheck bool, pbHandled *bool) int    // 代码编辑框_设置断点.
type XE_EDITOR_REMOVEBREAKPOINT func(iRow int32, pbHandled *bool) int                         // 代码编辑框_移除断点.
type XE_EDITOR_REMOVEBREAKPOINT1 func(hEle int, iRow int32, pbHandled *bool) int              // 代码编辑框_移除断点.
type XE_EDITOR_AUTOMATCH_SELECT func(iRow int32, nRows int32, pbHandled *bool) int            // 代码编辑框_自动匹配选择.
type XE_EDITOR_AUTOMATCH_SELECT1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int // 代码编辑框_自动匹配选择.

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE func(pItem *List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE1 func(hEle int, pItem *List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END func(pItem *List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END1 func(hEle int, pItem *List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY func(pItem *List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY1 func(hEle int, pItem *List_Item_, nFlag int32, pbHandled *bool) int
type XE_LIST_TEMP_ADJUST_COORDINATE func(pItem *List_Item_, pbHandled *bool) int                          // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *List_Item_, pbHandled *bool) int               // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_DRAWITEM func(hDraw int, pItem *List_Item_, pbHandled *bool) int                             // 列表元素,绘制项.
type XE_LIST_DRAWITEM1 func(hEle int, hDraw int, pItem *List_Item_, pbHandled *bool) int                  // 列表元素,绘制项.
type XE_LIST_SELECT func(iItem int32, pbHandled *bool) int                                                // 列表元素,项选择事件.
type XE_LIST_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                                     // 列表元素,项选择事件.
type XE_LIST_HEADER_DRAWITEM func(hDraw int, pItem *List_Header_Item_, pbHandled *bool) int               // 列表元素绘制列表头项.
type XE_LIST_HEADER_DRAWITEM1 func(hEle int, hDraw int, pItem *List_Header_Item_, pbHandled *bool) int    // 列表元素绘制列表头项.
type XE_LIST_HEADER_CLICK func(iItem int32, pbHandled *bool) int                                          // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_CLICK1 func(hEle int, iItem int32, pbHandled *bool) int                               // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_WIDTH_CHANGE func(iItem int32, nWidth int32, pbHandled *bool) int                     // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_WIDTH_CHANGE1 func(hEle int, iItem int32, nWidth int32, pbHandled *bool) int          // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_TEMP_CREATE func(pItem *List_Header_Item_, pbHandled *bool) int                       // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE1 func(hEle int, pItem *List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE_END func(pItem *List_Header_Item_, pbHandled *bool) int                   // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_CREATE_END1 func(hEle int, pItem *List_Header_Item_, pbHandled *bool) int        // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_DESTROY func(pItem *List_Header_Item_, pbHandled *bool) int                      // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_DESTROY1 func(hEle int, pItem *List_Header_Item_, pbHandled *bool) int           // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE func(pItem *List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板调整坐标. 已停用.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *List_Header_Item_, pbHandled *bool) int // 列表元素,列表头项模板调整坐标. 已停用.

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE func(pItem *ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE1 func(hEle int, pItem *ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END func(pItem *ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END1 func(hEle int, pItem *ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY func(pItem *ListBox_Item_, nFlag int, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY1 func(hEle int, pItem *ListBox_Item_, nFlag int, pbHandled *bool) int
type XE_LISTBOX_TEMP_ADJUST_COORDINATE func(pItem *ListBox_Item_, pbHandled *bool) int            // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *ListBox_Item_, pbHandled *bool) int // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_DRAWITEM func(hDraw int, pItem *ListBox_Item_, pbHandled *bool) int               // 列表框元素,项绘制事件.
type XE_LISTBOX_DRAWITEM1 func(hEle int, hDraw int, pItem *ListBox_Item_, pbHandled *bool) int    // 列表框元素,项绘制事件.
type XE_LISTBOX_SELECT func(iItem int32, pbHandled *bool) int                                     // 列表框元素,项选择事件.
type XE_LISTBOX_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                          // 列表框元素,项选择事件.

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE func(pItem *ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE1 func(hEle int, pItem *ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END func(pItem *ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END1 func(hEle int, pItem *ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY func(pItem *ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY1 func(hEle int, pItem *ListView_Item_, nFlag int32, pbHandled *bool) int
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE func(pItem *ListView_Item_, pbHandled *bool) int            // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *ListView_Item_, pbHandled *bool) int // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_DRAWITEM func(hDraw int, pItem *ListView_Item_, pbHandled *bool) int               // 列表视元素,自绘项.
type XE_LISTVIEW_DRAWITEM1 func(hEle int, hDraw int, pItem *ListView_Item_, pbHandled *bool) int    // 列表视元素,自绘项.
type XE_LISTVIEW_SELECT func(iGroup int32, iItem int32, pbHandled *bool) int                        // 列表视元素,项选择事件.
type XE_LISTVIEW_SELECT1 func(hEle int, iGroup int32, iItem int32, pbHandled *bool) int             // 列表视元素,项选择事件.
type XE_LISTVIEW_EXPAND func(iGroup int32, bExpand bool, pbHandled *bool) int                       // 列表视元素,组展开收缩事件.
type XE_LISTVIEW_EXPAND1 func(hEle int, iGroup int32, bExpand bool, pbHandled *bool) int            // 列表视元素,组展开收缩事件.

type XE_MONTHCAL_CHANGE func(pbHandled *bool) int            // 月历元素日期改变事件.
type XE_MONTHCAL_CHANGE1 func(hEle int, pbHandled *bool) int // 月历元素日期改变事件.

type XE_PROGRESSBAR_CHANGE func(pos int32, pbHandled *bool) int            // 进度条元素,进度改变事件.
type XE_PROGRESSBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int // 进度条元素,进度改变事件.

type XE_SBAR_SCROLL func(pos int32, pbHandled *bool) int            // 滚动条元素滚动事件,滚动条触发.
type XE_SBAR_SCROLL1 func(hEle int, pos int32, pbHandled *bool) int // 滚动条元素滚动事件,滚动条触发.

type XE_SCROLLVIEW_SCROLL_H func(pos int32, pbHandled *bool) int            // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_H1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V func(pos int32, pbHandled *bool) int            // 滚动视图元素垂直滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素垂直滚动事件,滚动视图触发.

type XE_SLIDERBAR_CHANGE func(pos int32, pbHandled *bool) int            // 滑动条元素,滑块位置改变事件.
type XE_SLIDERBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int // 滑动条元素,滑块位置改变事件.

type XE_TABBAR_SELECT func(iItem int32, pbHandled *bool) int            // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE func(iItem int32, pbHandled *bool) int            // TabBar标签按钮删除事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮删除事件. iItem: 标签位置索引.

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE func(pItem *Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE1 func(hEle int, pItem *Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE_END func(pItem *Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE_END1 func(hEle int, pItem *Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_TREE_TEMP_DESTROY func(pItem *Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_TREE_TEMP_DESTROY1 func(hEle int, pItem *Tree_Item_, nFlag int32, pbHandled *bool) int
type XE_TREE_TEMP_ADJUST_COORDINATE func(pItem *Tree_Item_, pbHandled *bool) int            // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *Tree_Item_, pbHandled *bool) int // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_DRAWITEM func(hDraw int, pItem *Tree_Item_, pbHandled *bool) int               // 树元素,绘制项.
type XE_TREE_DRAWITEM1 func(hEle int, hDraw int, pItem *Tree_Item_, pbHandled *bool) int    // 树元素,绘制项.
type XE_TREE_SELECT func(nItemID int32, pbHandled *bool) int                                // 树元素,项选择事件.
type XE_TREE_SELECT1 func(hEle int, nItemID int32, pbHandled *bool) int                     // 树元素,项选择事件.
type XE_TREE_EXPAND func(id int32, bExpand bool, pbHandled *bool) int                       // 树元素,项展开收缩事件.
type XE_TREE_EXPAND1 func(hEle int, id int32, bExpand bool, pbHandled *bool) int            // 树元素,项展开收缩事件.
type XE_TREE_DRAG_ITEM_ING func(pInfo *Tree_Drag_Item_, pbHandled *bool) int                // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM_ING1 func(hEle int, pInfo *Tree_Drag_Item_, pbHandled *bool) int     // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM func(pInfo *Tree_Drag_Item_, pbHandled *bool) int                    // 树元素,拖动项事件.
type XE_TREE_DRAG_ITEM1 func(hEle int, pInfo *Tree_Drag_Item_, pbHandled *bool) int         // 树元素,拖动项事件.
