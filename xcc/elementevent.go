package xcc

// 元素事件
type XE_ uint32

const (
	// XE_ELEPROCE 元素处理过程事件.
	//	func onEleProce(nEvent int, wParam, lParam uint, pbHandled *bool) int {
	//		return 0
	//	}
	XE_ELEPROCE XE_ = 1

	/*元素绘制事件.
	func onElePaint(hDraw int, pbHandled *bool) int{
		return 0
	}*/
	XE_PAINT XE_ = 2

	/*该元素及子元素绘制完成事件. 启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
	func onPaintEnd(hDraw int, pbHandled *bool) int {
		return 0
	}*/
	XE_PAINT_END XE_ = 3

	/*滚动视图绘制事件.
	func onPaintScrollview(hDraw int, pbHandled *bool) int {
		return 0
	}*/
	XE_PAINT_SCROLLVIEW XE_ = 4

	/*元素鼠标移动事件.
	func onMouseMove(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	   	return 0
	}*/
	XE_MOUSEMOVE XE_ = 5

	XE_MOUSESTAY        XE_ = 6  // 元素鼠标进入事件.
	XE_MOUSEHOVER       XE_ = 7  // 元素鼠标悬停事件.
	XE_MOUSELEAVE       XE_ = 8  // 元素鼠标离开事件.
	XE_MOUSEWHEEL       XE_ = 9  // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL
	XE_LBUTTONDOWN      XE_ = 10 // 鼠标左键按下事件.
	XE_LBUTTONUP        XE_ = 11 // 鼠标左键弹起事件.
	XE_RBUTTONDOWN      XE_ = 12 // 鼠标右键按下事件.
	XE_RBUTTONUP        XE_ = 13 // 鼠标右键弹起事件.
	XE_LBUTTONDBCLICK   XE_ = 14 // 鼠标左键双击事件.
	XE_XC_TIMER         XE_ = 16 // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
	XE_ADJUSTLAYOUT     XE_ = 17 // 调整布局事件. 暂停使用
	XE_ADJUSTLAYOUT_END XE_ = 18 // 调整布局完成事件.
	XE_TOOLTIP_POPUP    XE_ = 19 // 元素工具提示弹出事件.
	XE_SETFOCUS         XE_ = 31 // 元素获得焦点事件.
	XE_KILLFOCUS        XE_ = 32 // 元素失去焦点事件.
	XE_DESTROY          XE_ = 33 // 元素即将销毁事件. 在销毁子对象之前触发
	XE_DESTROY_END      XE_ = 42 // 元素销毁完成事件. 在销毁子对象之后触发

	/*按钮点击事件.
	func onBnClick(pbHandled *bool) int {
		return 0
	}*/
	XE_BNCLICK XE_ = 34

	XE_BUTTON_CHECK                       XE_ = 35  // 按钮选中事件.
	XE_SIZE                               XE_ = 36  // 元素大小改变事件.
	XE_SHOW                               XE_ = 37  // 元素显示隐藏事件.
	XE_SETFONT                            XE_ = 38  // 元素设置字体事件.
	XE_KEYDOWN                            XE_ = 39  // 元素按键事件.
	XE_KEYUP                              XE_ = 40  // 元素按键事件.
	XE_CHAR                               XE_ = 41  // 通过TranslateMessage函数翻译的字符事件.
	XE_SETCAPTURE                         XE_ = 51  // 元素设置鼠标捕获.
	XE_KILLCAPTURE                        XE_ = 52  // 元素失去鼠标捕获.
	XE_SETCURSOR                          XE_ = 53  // 设置鼠标光标
	XE_SCROLLVIEW_SCROLL_H                XE_ = 54  // 滚动视图元素水平滚动事件,滚动视图触发.
	XE_SCROLLVIEW_SCROLL_V                XE_ = 55  // 滚动视图元素垂直滚动事件,滚动视图触发.
	XE_SBAR_SCROLL                        XE_ = 56  // 滚动条元素滚动事件,滚动条触发.
	XE_MENU_POPUP                         XE_ = 57  // 菜单弹出
	XE_MENU_POPUP_WND                     XE_ = 58  // 菜单弹出窗口
	XE_MENU_SELECT                        XE_ = 59  // 弹出菜单项选择事件.
	XE_MENU_DRAW_BACKGROUND               XE_ = 60  // 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground
	XE_MENU_DRAWITEM                      XE_ = 61  // 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem
	XE_MENU_EXIT                          XE_ = 62  // 弹出菜单退出事件.
	XE_SLIDERBAR_CHANGE                   XE_ = 63  // 滑动条元素,滑块位置改变事件.
	XE_PROGRESSBAR_CHANGE                 XE_ = 64  // 进度条元素,进度改变事件.
	XE_COMBOBOX_SELECT                    XE_ = 71  // 组合框下拉列表项选择事件.
	XE_COMBOBOX_SELECT_END                XE_ = 74  // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
	XE_COMBOBOX_POPUP_LIST                XE_ = 72  // 组合框下拉列表弹出事件.
	XE_COMBOBOX_EXIT_LIST                 XE_ = 73  // 组合框下拉列表退出事件.
	XE_LISTBOX_TEMP_CREATE                XE_ = 81  // 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LISTBOX_TEMP_CREATE_END            XE_ = 82  // 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LISTBOX_TEMP_DESTROY               XE_ = 83  // 列表框元素,项模板销毁.
	XE_LISTBOX_TEMP_ADJUST_COORDINATE     XE_ = 84  // 列表框元素,项模板调整坐标. 已停用
	XE_LISTBOX_DRAWITEM                   XE_ = 85  // 列表框元素,项绘制事件.
	XE_LISTBOX_SELECT                     XE_ = 86  // 列表框元素,项选择事件.
	XE_LIST_TEMP_CREATE                   XE_ = 101 // 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LIST_TEMP_CREATE_END               XE_ = 102 // 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LIST_TEMP_DESTROY                  XE_ = 103 // 列表元素,项模板销毁.
	XE_LIST_TEMP_ADJUST_COORDINATE        XE_ = 104 // 列表元素,项模板调整坐标. 已停用
	XE_LIST_DRAWITEM                      XE_ = 105 // 列表元素,绘制项.
	XE_LIST_SELECT                        XE_ = 106 // 列表元素,项选择事件.
	XE_LIST_HEADER_DRAWITEM               XE_ = 107 // 列表元素绘制列表头项.
	XE_LIST_HEADER_CLICK                  XE_ = 108 // 列表元素,列表头项点击事件.
	XE_LIST_HEADER_WIDTH_CHANGE           XE_ = 109 // 列表元素,列表头项宽度改变事件.
	XE_LIST_HEADER_TEMP_CREATE            XE_ = 110 // 列表元素,列表头项模板创建.
	XE_LIST_HEADER_TEMP_CREATE_END        XE_ = 111 // 列表元素,列表头项模板创建完成事件.
	XE_LIST_HEADER_TEMP_DESTROY           XE_ = 112 // 列表元素,列表头项模板销毁.
	XE_LIST_HEADER_TEMP_ADJUST_COORDINATE XE_ = 113 // 列表元素,列表头项模板调整坐标. 已停用
	XE_TREE_TEMP_CREATE                   XE_ = 121 // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_TREE_TEMP_CREATE_END               XE_ = 122 // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_TREE_TEMP_DESTROY                  XE_ = 123 // 列表树元素-项模板销毁,模板复用机制需先启用;
	XE_TREE_TEMP_ADJUST_COORDINATE        XE_ = 124 // 树元素,项模板,调整项坐标. 已停用
	XE_TREE_DRAWITEM                      XE_ = 125 // 树元素,绘制项.
	XE_TREE_SELECT                        XE_ = 126 // 树元素,项选择事件.
	XE_TREE_EXPAND                        XE_ = 127 // 树元素,项展开收缩事件.
	XE_TREE_DRAG_ITEM_ING                 XE_ = 128 // 树元素,用户正在拖动项, 可对参数值修改.
	XE_TREE_DRAG_ITEM                     XE_ = 129 // 树元素,拖动项事件.
	XE_LISTVIEW_TEMP_CREATE               XE_ = 141 // 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LISTVIEW_TEMP_CREATE_END           XE_ = 142 // 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LISTVIEW_TEMP_DESTROY              XE_ = 143 // 列表视元素-项模板销毁, 模板复用机制需先启用;
	XE_LISTVIEW_TEMP_ADJUST_COORDINATE    XE_ = 144 // 列表视元素,项模板调整坐标.已停用
	XE_LISTVIEW_DRAWITEM                  XE_ = 145 // 列表视元素,自绘项.
	XE_LISTVIEW_SELECT                    XE_ = 146 // 列表视元素,项选择事件.
	XE_LISTVIEW_EXPAND                    XE_ = 147 // 列表视元素,组展开收缩事件.
	XE_PGRID_VALUE_CHANGE                 XE_ = 151 // 属性网格元素 项值改变事件
	XE_PGRID_ITEM_SET                     XE_ = 152 //
	XE_PGRID_ITEM_SELECT                  XE_ = 153 //
	XE_PGRID_ITEM_ADJUST_COORDINATE       XE_ = 154 //
	XE_PGRID_ITEM_DESTROY                 XE_ = 155 //
	XE_PGRID_ITEM_EXPAND                  XE_ = 156 //
	XE_RICHEDIT_CHANGE                    XE_ = 161 // 富文本元素 用户修改内容事件,只有当用户操作时才会触发,需要开启后有效, XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE()； XRichEdit_SetText()、 XRichEdit_InsertString()不会触发此事件
	XE_EDIT_SET                           XE_ = 162 // 元素事件_编辑框设置
	XE_EDIT_DRAWROW                       XE_ = 181 // 暂未使用
	XE_EDIT_CHANGED                       XE_ = 182 // 编辑框_内容被改变
	XE_EDIT_POS_CHANGED                   XE_ = 183 // 编辑框_光标位置_被改变
	XE_EDIT_STYLE_CHANGED                 XE_ = 184 // 编辑框_样式_被改变
	XE_EDIT_ENTER_GET_TABALIGN            XE_ = 185 // 回车TAB对齐,返回需要TAB数量
	XE_EDIT_SWAPROW                       XE_ = 186 // 元素事件_交换行
	XE_EDIT_COLOR_CHANGE                  XE_ = 260 // 编辑框_颜色被改变

	XE_EDITOR_MODIFY_ROWS      XE_ = 190 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化
	XE_EDITOR_SETBREAKPOINT    XE_ = 191 // 代码编辑框_设置断点
	XE_EDITOR_REMOVEBREAKPOINT XE_ = 192 // 代码编辑框_移除断点
	XE_EDIT_ROW_CHANGED        XE_ = 193 // 编辑框_行_被改变
	XE_EDITOR_AUTOMATCH_SELECT XE_ = 194 // 编辑框_自动匹配选择
	XE_TABBAR_SELECT           XE_ = 221 // TabBar标签按钮选择改变事件
	XE_TABBAR_DELETE           XE_ = 222 // TabBar标签按钮删除事件
	XE_MONTHCAL_CHANGE         XE_ = 231 // 月历元素日期改变事件
	XE_DATETIME_CHANGE         XE_ = 241 // 日期时间元素,内容改变事件
	XE_DATETIME_POPUP_MONTHCAL XE_ = 242 // 日期时间元素,弹出月历卡片事件
	XE_DATETIME_EXIT_MONTHCAL  XE_ = 243 // 日期时间元素,弹出的月历卡片退出事件
	XE_DROPFILES               XE_ = 250 // 文件拖放事件.
)
