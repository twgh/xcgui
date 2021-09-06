package xcc // xcgui const

// SW_

const (
	SW_SHOW           = 5 // 显示
	SW_HIDE           = 0 // 隐藏
	SW_SHOWMAXIMIZED  = 3 // 最大化
	SW_SHOWMINIMIZED  = 2 // 最小化
	SW_SHOWNOACTIVATE = 4 // 不激活
	SW_SHOWNA         = 8 // 原来的状态显示
)

// 特殊ID

const (
	XC_ID_ROOT  = 0  // 根节点
	XC_ID_ERROR = -1 // ID错误
	XC_ID_FIRST = -2 // 插入开始位置
	XC_ID_LAST  = -3 // 插入末尾位置
)

// XC_AdjustLayout_

const (
	XC_AdjustLayout_No   = iota // 不调整布局
	XC_AdjustLayout_All         // 强制调整自身和子对象布局
	XC_AdjustLayout_Self        // 只调整自身布局, 不调整子对象布局
)

// Window_Transparent_

const (
	Window_Transparent_False  = iota // 窗口透明_无
	Window_Transparent_Shaped        // 窗口透明_异型
	Window_Transparent_Shadow        // 窗口透明_阴影,带透明通道，边框阴影，窗口透明或半透明。
	Window_Transparent_Simple        // 窗口透明_简单,指定透明色.
	Window_Transparent_Win7          // 窗口透明_玻璃
)

// Layout_Align_Axis_

const (
	Layout_Align_Axis_Auto   = iota // 无
	Layout_Align_Axis_Start         // 水平布局(顶部), 垂直布局(左侧)
	Layout_Align_Axis_Center        // 居中
	Layout_Align_Axis_End           // 水平布局(底部), 垂直布局(右侧)
)

// Layout_Size_

const (
	Layout_Size_Fixed   = iota // 固定大小
	Layout_Size_Fill           // 填充父
	Layout_Size_Auto           // 根据内容计算大小
	Layout_Size_Weight         // 按照比例分配剩余空间
	Layout_Size_Percent        // 百分比
	Layout_Size_Disable        // 不使用
)

// Layout_Align_

const (
	Layout_Align_Left        = iota // 左侧
	Layout_Align_Top                // 顶部
	Layout_Align_Right              // 右侧
	Layout_Align_Bottom             // 底部
	Layout_Align_Center             // 居中
	Layout_Align_Equidistant        // 等距
)

// XC_OBJECT_STYLE

const (
	Button_Style_Default            = iota // 默认风格
	Button_Style_Check                     // 复选按钮
	Button_Style_Radio                     // 单选按钮
	Button_Style_Expand                    // 展开收缩按钮
	Button_Style_Scrollbar_Left            // 水平滚动条,左按钮
	Button_Style_Scrollbar_Right           // 水平滚动条,右按钮
	Button_Style_Scrollbar_Up              // 垂直滚动条,上按钮
	Button_Style_Scrollbar_Down            // 垂直滚动条,下按钮
	Button_Style_Scrollbar_Slider_H        // 水平滚动条,滑块
	Button_Style_Scrollbar_Slider_V        // 垂直滚动条,滑块
	Button_Style_TabBar_Button             // TabBar上的标签按钮
	Button_Style_ToolBar_Left              // ToolBar左滚动按钮
	Button_Style_ToolBar_Right             // ToolBar右滚动按钮
	Button_Style_Pane_Close                // 窗格关闭按钮
	Button_Style_Pane_Lock                 // 窗格锁定按钮
	Button_Style_Pane_Menu                 // 窗格下拉菜单按钮
	Button_Style_Pane_DockH                // 框架窗口左边或右边码头上按钮
	Button_Style_Pane_DockV                // 框架窗口上边或下边码头上按钮
	FrameWnd_Style_Dock_Left               // 框架窗口停靠码头
	FrameWnd_Style_Dock_Top                // 框架窗口停靠码头
	FrameWnd_Style_Dock_Right              // 框架窗口停靠码头
	FrameWnd_Style_Dock_Bottom             // 框架窗口停靠码头
	ToolBar_Style_Separator                // 工具条上的分割线
	ListBox_Style_ComboBox                 // 下拉组合框弹出的listBox
)

// MessageBox_Flag_

const (
	MessageBox_Flag_Other        = 0      // 其他
	MessageBox_Flag_Ok           = 1      // 确定按钮
	MessageBox_Flag_Cancel       = 2      // 取消按钮
	MessageBox_Flag_Icon_Appicon = 4096   // 图标 应用程序 IDI_APPLICATION
	MessageBox_Flag_Icon_Info    = 8192   // 图标 信息 IDI_ASTERISK
	MessageBox_Flag_Icon_Qustion = 16384  // 图标 问询/帮助/提问 IDI_QUESTION
	MessageBox_Flag_Icon_Error   = 32768  // 图标 错误/拒绝/禁止 IDI_ERROR
	MessageBox_Flag_Icon_Warning = 65536  // 图标 警告 IDI_WARNING
	MessageBox_Flag_Icon_Shield  = 131072 // 图标 盾牌/安全 IDI_SHIELD
)

// Common_State3_

const (
	Common_State3_Leave = iota // 离开
	Common_State3_Stay         // 停留
	Common_State3_Down         // 按下
)

// Button_State_

const (
	Button_State_Leave   = iota // 离开状态
	Button_State_Stay           // 停留状态
	Button_State_Down           // 按下状态
	Button_State_Check          // 选中状态
	Button_State_Disable        // 禁用状态
)

// XC_OBJECT_TYPE_EX

const (
	Button_Type_Default = iota // 默认类型
	Button_Type_Radio          // 单选按钮
	Button_Type_Check          // 复选按钮
	Button_Type_Close          // 窗口关闭按钮
	Button_Type_Min            // 窗口最小化按钮
	Button_Type_Max            // 窗口最大化还原按钮
	Element_Type_Layout        // 布局元素,启用布局功能的元素
	Xc_Ex_Error         = -1   // 错误类型
)

// TextFormatFlag_

const (
	TextAlignFlag_Left                   = 0       // 左对齐
	TextAlignFlag_Top                    = 0       // 垂直顶对齐
	TextAlignFlag_Left_Top               = 0x4000  // 内部保留
	TextAlignFlag_Center                 = 0x1     // 水平居中
	TextAlignFlag_Right                  = 0x2     // 右对齐
	TextAlignFlag_Vcenter                = 0x4     // 垂直居中
	TextAlignFlag_Bottom                 = 0x8     // 垂直底对齐
	TextFormatFlag_DirectionRightToLeft  = 0x10    // 从右向左顺序显示文本
	TextFormatFlag_NoWrap                = 0x20    // 禁止换行
	TextFormatFlag_DirectionVertical     = 0x40    // 垂直显示文本
	TextFormatFlag_NoFitBlackBox         = 0x80    // 允许部分字符延伸该字符串的布局矩形。默认情况下，将重新定位字符以避免任何延伸
	TextFormatFlag_DisplayFormatControl  = 0x100   // 控制字符（如从左到右标记）随具有代表性的标志符号一起显示在输出中。
	TextFormatFlag_NoFontFallback        = 0x200   // 对于请求的字体中不支持的字符，禁用回退到可选字体。缺失的任何字符都用缺失标志符号的字体显示，通常是一个空的方块
	TextFormatFlag_MeasureTrailingSpaces = 0x400   // 包括每一行结尾处的尾随空格。在默认情况下，MeasureString 方法返回的边框都将排除每一行结尾处的空格。设置此标记以便在测定时将空格包括进去
	TextFormatFlag_LineLimit             = 0x800   // 如果内容显示高度不够一行,那么不显示
	TextFormatFlag_NoClip                = 0x1000  // 允许显示标志符号的伸出部分和延伸到边框外的未换行文本。在默认情况下，延伸到边框外侧的所有文本和标志符号部分都被剪裁
	TextTrimming_None                    = 0       // 不使用去尾
	TextTrimming_Character               = 0x40000 // 以字符为单位去尾
	TextTrimming_Word                    = 0x80000 // 以单词为单位去尾
	TextTrimming_EllipsisCharacter       = 0x8000  // 以字符为单位去尾,省略部分使用且略号表示
	TextTrimming_EllipsisWord            = 0x10000 // 以单词为单位去尾,
	TextTrimming_EllipsisPath            = 0x20000 // 略去字符串中间部分，保证字符的首尾都能够显示
)

// Button_Icon_Align_

const (
	Button_Icon_Align_Left   = iota // 图标在左边
	Button_Icon_Align_Top           // 图标在顶部
	Button_Icon_Align_Right         // 图标在右边
	Button_Icon_Align_Bottom        // 图标在底部
)

// Xc_Window_Style_

const (
	Xc_Window_Style_Nothing     = 0x00000000                                                                                                                              // 什么也没有
	Xc_Window_Style_Caption     = 0x00000001                                                                                                                              // top布局,如果指定该属性,默认为绑定标题栏元素
	Xc_Window_Style_Border      = 0x00000002                                                                                                                              // 边框,指定默认上下左右布局大小,如果没有指定,那么边框布局大小为0
	Xc_Window_Style_Center      = 0x00000004                                                                                                                              // 窗口居中
	Xc_Window_Style_Drag_Border = 0x00000008                                                                                                                              // 拖动窗口边框
	Xc_Window_Style_Drag_Window = 0x00000010                                                                                                                              // 拖动窗口
	Xc_Window_Style_Default     = (Xc_Window_Style_Caption | Xc_Window_Style_Border | Xc_Window_Style_Center | Xc_Window_Style_Drag_Border | Xc_Window_Style_Drag_Window) // 允许窗口最大化  窗口默认样式
	Xc_Window_Style_Modal       = (Xc_Window_Style_Caption | Xc_Window_Style_Center | Xc_Window_Style_Border)                                                             // 模态窗口样式
)

// XC_OBJECT_TYPE

const (
	XC_ERROR            = -1       // 错误类型
	XC_WINDOW           = 1        // 窗口
	XC_MODALWINDOW      = 2        // 模态窗口
	XC_FRAMEWND         = 3        // 框架窗口
	XC_FLOATWND         = 4        // 浮动窗口
	XC_OBJECT_UI        = 19       // ...
	XC_ELE              = 21       // 基础元素
	XC_ELE_LAYOUT       = 53       // 布局元素
	XC_LAYOUT_FRAME     = 54       // 流式布局
	XC_BUTTON           = 22       // 按钮
	XC_EDIT             = 45       // 编辑框
	XC_EDITOR           = 46       // 代码编辑框
	XC_RICHEDIT         = 23       // 富文本编辑框
	XC_COMBOBOX         = 24       // 下拉组合框
	XC_SCROLLBAR        = 25       // 滚动条
	XC_SCROLLVIEW       = 26       // 滚动视图
	XC_LIST             = 27       // 列表
	XC_LISTBOX          = 28       // 列表框
	XC_LISTVIEW         = 29       // 列表视图,大图标
	XC_TREE             = 30       // 列表树
	XC_MENUBAR          = 31       // 菜单条
	XC_SLIDERBAR        = 32       // 滑动条
	XC_PROGRESSBAR      = 33       // 进度条
	XC_TOOLBAR          = 34       // 工具条
	XC_MONTHCAL         = 35       // 月历卡片
	XC_DATETIME         = 36       // 日期时间
	XC_PROPERTYGRID     = 37       // 属性网格
	XC_EDIT_COLOR       = 38       // 颜色选择框
	XC_EDIT_SET         = 39       // 设置编辑框
	XC_TABBAR           = 40       // tab条
	XC_TEXTLINK         = 41       // 文本链接按钮
	XC_PANE             = 42       // 窗格
	XC_PANE_SPLIT       = 43       // 窗格拖动分割条
	XC_MENUBAR_BUTTON   = 44       // 菜单条上的按钮
	XC_EDIT_FILE        = 50       // EditFile 文件选择编辑框
	XC_EDIT_FOLDER      = 51       // EditFolder 文件夹选择编辑框
	XC_LIST_HEADER      = 52       // 列表头元素
	XC_SHAPE            = 61       // 形状对象
	XC_SHAPE_TEXT       = 62       // 形状对象-文本
	XC_SHAPE_PICTURE    = 63       // 形状对象-图片
	XC_SHAPE_RECT       = 64       // 形状对象-矩形
	XC_SHAPE_ELLIPSE    = 65       // 形状对象-圆
	XC_SHAPE_LINE       = 66       // 形状对象-直线
	XC_SHAPE_GROUPBOX   = 67       // 形状对象-组框
	XC_SHAPE_GIF        = 68       // 形状对象-GIF
	XC_SHAPE_TABLE      = 69       // 形状对象-表格
	XC_MENU             = 81       // 弹出菜单
	XC_IMAGE            = 82       // 图片
	XC_IMAGE_TEXTURE    = XC_IMAGE // 图片纹理,图片源,图片素材
	XC_HDRAW            = 83       // 绘图操作
	XC_FONT             = 84       // 炫彩字体
	XC_FLASH            = 85       // flash
	XC_PANE_CELL        = 86       // ...
	XC_WEB              = 87       // ...
	XC_IMAGE_FRAME      = 88       // 图片帧,指定图片的渲染属性
	XC_LAYOUT_OBJECT    = 101      // 布局对象LayoutObject
	XC_ADAPTER          = 102      // ...
	XC_ADAPTER_TABLE    = 103      // 数据适配器AdapterTable
	XC_ADAPTER_TREE     = 104      // 数据适配器AdapterTree
	XC_ADAPTER_LISTVIEW = 105      // 数据适配器AdapterListView
	XC_ADAPTER_MAP      = 106      // 数据适配器AdapterMap
)

// Element_Position_

const (
	Element_Position_No     = 0x00 // 无效
	Element_Position_Left   = 0x01 // 左边
	Element_Position_Top    = 0x02 // 上边
	Element_Position_Right  = 0x04 // 右边
	Element_Position_Bottom = 0x08 // 下边
)

// 元素事件

const (
	XE_ELEPROCE                           = 1   // 元素处理过程事件.
	XE_PAINT                              = 2   // 元素绘制事件
	XE_PAINT_END                          = 3   // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END()
	XE_PAINT_SCROLLVIEW                   = 4   // 滚动视图绘制事件.
	XE_MOUSEMOVE                          = 5   // 元素鼠标移动事件.
	XE_MOUSESTAY                          = 6   // 元素鼠标进入事件.
	XE_MOUSEHOVER                         = 7   // 元素鼠标悬停事件.
	XE_MOUSELEAVE                         = 8   // 元素鼠标离开事件.
	XE_MOUSEWHEEL                         = 9   // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL()
	XE_LBUTTONDOWN                        = 10  // 鼠标左键按下事件.
	XE_LBUTTONUP                          = 11  // 鼠标左键弹起事件.
	XE_RBUTTONDOWN                        = 12  // 鼠标右键按下事件.
	XE_RBUTTONUP                          = 13  // 鼠标右键弹起事件.
	XE_LBUTTONDBCLICK                     = 14  // 鼠标左键双击事件.
	XE_XC_TIMER                           = 16  // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
	XE_ADJUSTLAYOUT                       = 17  // 调整布局事件. 暂停使用
	XE_ADJUSTLAYOUT_END                   = 18  // 调整布局完成事件.
	XE_SETFOCUS                           = 31  // 元素获得焦点事件.
	XE_KILLFOCUS                          = 32  // 元素失去焦点事件.
	XE_DESTROY                            = 33  // 元素即将销毁事件. 在销毁子对象之前触发
	XE_DESTROY_END                        = 42  // 元素销毁完成事件. 在销毁子对象之后触发
	XE_BNCLICK                            = 34  // 按钮点击事件.
	XE_BUTTON_CHECK                       = 35  // 按钮选中事件.
	XE_SIZE                               = 36  // 元素大小改变事件.
	XE_SHOW                               = 37  // 元素显示隐藏事件.
	XE_SETFONT                            = 38  // 元素设置字体事件.
	XE_KEYDOWN                            = 39  // 元素按键事件.
	XE_KEYUP                              = 40  // 元素按键事件.
	XE_CHAR                               = 41  // 通过TranslateMessage函数翻译的字符事件.
	XE_SETCAPTURE                         = 51  // 元素设置鼠标捕获.
	XE_KILLCAPTURE                        = 52  // 元素失去鼠标捕获.
	XE_SETCURSOR                          = 53  // 设置鼠标光标
	XE_SCROLLVIEW_SCROLL_H                = 54  // 滚动视图元素水平滚动事件,滚动视图触发.
	XE_SCROLLVIEW_SCROLL_V                = 55  // 滚动视图元素垂直滚动事件,滚动视图触发.
	XE_SBAR_SCROLL                        = 56  // 滚动条元素滚动事件,滚动条触发.
	XE_MENU_POPUP                         = 57  // 菜单弹出
	XE_MENU_POPUP_WND                     = 58  // 菜单弹出窗口
	XE_MENU_SELECT                        = 59  // 弹出菜单项选择事件.
	XE_MENU_DRAW_BACKGROUND               = 60  // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
	XE_MENU_DRAWITEM                      = 61  // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
	XE_MENU_EXIT                          = 62  // 弹出菜单退出事件.
	XE_SLIDERBAR_CHANGE                   = 63  // 滑动条元素,滑块位置改变事件.
	XE_PROGRESSBAR_CHANGE                 = 64  // 进度条元素,进度改变事件.
	XE_COMBOBOX_SELECT                    = 71  // 组合框下拉列表项选择事件.
	XE_COMBOBOX_SELECT_END                = 74  // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
	XE_COMBOBOX_POPUP_LIST                = 72  // 组合框下拉列表弹出事件.
	XE_COMBOBOX_EXIT_LIST                 = 73  // 组合框下拉列表退出事件.
	XE_LISTBOX_TEMP_CREATE                = 81  // 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LISTBOX_TEMP_CREATE_END            = 82  // 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LISTBOX_TEMP_DESTROY               = 83  // 列表框元素,项模板销毁.
	XE_LISTBOX_TEMP_ADJUST_COORDINATE     = 84  // 列表框元素,项模板调整坐标. 已停用
	XE_LISTBOX_DRAWITEM                   = 85  // 列表框元素,项绘制事件.
	XE_LISTBOX_SELECT                     = 86  // 列表框元素,项选择事件.
	XE_LIST_TEMP_CREATE                   = 101 // 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LIST_TEMP_CREATE_END               = 102 // 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LIST_TEMP_DESTROY                  = 103 // 列表元素,项模板销毁.
	XE_LIST_TEMP_ADJUST_COORDINATE        = 104 // 列表元素,项模板调整坐标. 已停用
	XE_LIST_DRAWITEM                      = 105 // 列表元素,绘制项.
	XE_LIST_SELECT                        = 106 // 列表元素,项选择事件.
	XE_LIST_HEADER_DRAWITEM               = 107 // 列表元素绘制列表头项.
	XE_LIST_HEADER_CLICK                  = 108 // 列表元素,列表头项点击事件.
	XE_LIST_HEADER_WIDTH_CHANGE           = 109 // 列表元素,列表头项宽度改变事件.
	XE_LIST_HEADER_TEMP_CREATE            = 110 // 列表元素,列表头项模板创建.
	XE_LIST_HEADER_TEMP_CREATE_END        = 111 // 列表元素,列表头项模板创建完成事件.
	XE_LIST_HEADER_TEMP_DESTROY           = 112 // 列表元素,列表头项模板销毁.
	XE_LIST_HEADER_TEMP_ADJUST_COORDINATE = 113 // 列表元素,列表头项模板调整坐标. 已停用
	XE_TREE_TEMP_CREATE                   = 121 // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_TREE_TEMP_CREATE_END               = 122 // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_TREE_TEMP_DESTROY                  = 123 // 列表树元素-项模板销毁,模板复用机制需先启用;
	XE_TREE_TEMP_ADJUST_COORDINATE        = 124 // 树元素,项模板,调整项坐标. 已停用
	XE_TREE_DRAWITEM                      = 125 // 树元素,绘制项.
	XE_TREE_SELECT                        = 126 // 树元素,项选择事件.
	XE_TREE_EXPAND                        = 127 // 树元素,项展开收缩事件.
	XE_TREE_DRAG_ITEM_ING                 = 128 // 树元素,用户正在拖动项, 可对参数值修改.
	XE_TREE_DRAG_ITEM                     = 129 // 树元素,拖动项事件.
	XE_LISTVIEW_TEMP_CREATE               = 141 // 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
	XE_LISTVIEW_TEMP_CREATE_END           = 142 // 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册
	XE_LISTVIEW_TEMP_DESTROY              = 143 // 列表视元素-项模板销毁, 模板复用机制需先启用;
	XE_LISTVIEW_TEMP_ADJUST_COORDINATE    = 144 // 列表视元素,项模板调整坐标.已停用
	XE_LISTVIEW_DRAWITEM                  = 145 // 列表视元素,自绘项.
	XE_LISTVIEW_SELECT                    = 146 // 列表视元素,项选择事件.
	XE_LISTVIEW_EXPAND                    = 147 // 列表视元素,组展开收缩事件.
	XE_PGRID_VALUE_CHANGE                 = 151 // 属性网格元素 项值改变事件
	XE_PGRID_ITEM_SET                     = 152 //
	XE_PGRID_ITEM_SELECT                  = 153 //
	XE_PGRID_ITEM_ADJUST_COORDINATE       = 154 //
	XE_PGRID_ITEM_DESTROY                 = 155 //
	XE_PGRID_ITEM_EXPAND                  = 156 //
	XE_RICHEDIT_CHANGE                    = 161 // 富文本元素 用户修改内容事件,只有当用户操作时才会触发,需要开启后有效, XRichEdit_EnableEvent_XE_RICHEDIT_CHANGE()； XRichEdit_SetText()、 XRichEdit_InsertString()不会触发此事件
	XE_EDIT_SET                           = 162 //
	XE_EDIT_DRAWROW                       = 181 //
	XE_EDIT_CHANGED                       = 182 //
	XE_EDIT_POS_CHANGED                   = 183 //
	XE_EDIT_STYLE_CHANGED                 = 184 //
	XE_EDIT_ENTER_GET_TABALIGN            = 185 //
	XE_EDITOR_MODIFY_ROWS                 = 186 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化
	XE_EDITOR_SETBREAKPOINT               = 191 //
	XE_EDITOR_REMOVEBREAKPOINT            = 192 //
	XE_EDIT_ROW_CHANGED                   = 193 //
	XE_EDITOR_AUTOMATCH_SELECT            = 194 //
	XE_TABBAR_SELECT                      = 221 // TabBar标签按钮选择改变事件
	XE_TABBAR_DELETE                      = 222 // TabBar标签按钮删除事件
	XE_MONTHCAL_CHANGE                    = 231 // 月历元素日期改变事件
	XE_DATETIME_CHANGE                    = 241 // 日期时间元素,内容改变事件
	XE_DATETIME_POPUP_MONTHCAL            = 242 // 日期时间元素,弹出月历卡片事件
	XE_DATETIME_EXIT_MONTHCAL             = 243 // 日期时间元素,弹出的月历卡片退出事件
	XE_DROPFILES                          = 250 // 文件拖放事件.
)

// Pane_Align_

const (
	Pane_Align_Left   = iota // 左侧
	Pane_Align_Top           // 顶部
	Pane_Align_Right         // 右侧
	Pane_Align_Bottom        // 底部
	Pane_Align_Center        // 居中
	Pane_Align_Error  = -1   // 错误
)

// Menu_Item_Flag_

const (
	Menu_Item_Flag_Normal    = 0x00 // 正常
	Menu_Item_Flag_Select    = 0x01 // 选择
	Menu_Item_Flag_Check     = 0x02 // 勾选
	Menu_Item_Flag_Popup     = 0x04 // 弹出
	Menu_Item_Flag_Separator = 0x08 // 分隔栏 ID号任意,ID号被忽略
	Menu_Item_Flag_Disable   = 0x10 // 禁用
)

// Menu_Popup_Position_

const (
	Menu_Popup_Position_Left_Top      = iota // 左上角
	Menu_Popup_Position_Left_Bottom          // 左下角
	Menu_Popup_Position_Right_Top            // 右上角
	Menu_Popup_Position_Right_Bottom         // 右下角
	Menu_Popup_Position_Center_Left          // 左居中
	Menu_Popup_Position_Center_Top           // 上居中
	Menu_Popup_Position_Center_Right         // 右居中
	Menu_Popup_Position_Center_Bottom        // 下居中
)

// ComboBox_State_

const (
	ComboBox_State_Leave = iota // 鼠标离开状态
	ComboBox_State_Stay         // 鼠标停留状态
	ComboBox_State_Down         // 按下状态
)

// Adapter_Date_Type_

const (
	Adapter_Date_Type_Int    = iota // 整形
	Adapter_Date_Type_Float         // 浮点型
	Adapter_Date_Type_String        // 字符串
	Adapter_Date_Type_Image         // 图片
)

// Chat_Flag_

const (
	Chat_Flag_Left            = 0x1 // 左侧
	Chat_Flag_Right           = 0x2 // 右侧
	Chat_Flag_Center          = 0x4 // 中间
	Chat_Flag_Next_Row_Bubble = 0x8 // 下一行显示气泡
)

// Edit_Type_

const (
	Edit_Type_None      = iota // 普通编辑框, 每行的高度相同
	Edit_Type_Editor           // 代码编辑
	Edit_Type_Richedit         // 富文本编辑框, 每行的高度可能不同
	Edit_Type_Chat             // 聊天气泡, 每行的高度可能不同
	Edit_Type_CodeTable        // 代码表格,内部使用, 每行的高度相同
)

// Edit_Style_Type_

const (
	Edit_Style_Type_Font_Color = iota + 1 // 字体
	Edit_Style_Type_Image                 // 图片
	Edit_Style_Type_Obj                   // UI对象
)

// Edit_TextAlign_Flag_

const (
	Edit_TextAlign_Flag_Left     = 0x0 // 左侧
	Edit_TextAlign_Flag_Right    = 0x1 // 右侧
	Edit_TextAlign_Flag_Center   = 0x2 // 水平居中
	Edit_TextAlign_Flag_Top      = 0x0 // 顶部
	Edit_TextAlign_Flag_Bottom   = 0x4 // 底部
	Edit_TextAlign_Flag_Center_V = 0x8 // 垂直居中
)

// Table_Flag_

const (
	Table_Flag_Full = iota // 铺满组合单元格
	Table_Flag_None        // 正常最小单元格
)

// GRADIENT_FILL_

const (
	GRADIENT_FILL_RECT_H   = iota //水平填充
	GRADIENT_FILL_RECT_V          //垂直填充
	GRADIENT_FILL_TRIANGLE        //三角形
)

// Ease_

const (
	Ease_In    = iota // 从慢到快
	Ease_Out          // 从快到慢
	Ease_InOut        // 从慢到快再到慢
)

// FontStyle_

const (
	FontStyle_Regular    = 0 // 正常
	FontStyle_Bold       = 1 // 粗体
	FontStyle_Italic     = 2 // 斜体
	FontStyle_BoldItalic = 3 // 粗斜体
	FontStyle_Underline  = 4 // 下划线
	FontStyle_Strikeout  = 8 // 删除线
)

// Image_Draw_Type_

const (
	Image_Draw_Type_Default         = iota // 默认
	Image_Draw_Type_Stretch                // 拉伸
	Image_Draw_Type_Adaptive               // 自适应,九宫格
	Image_Draw_Type_Tile                   // 平铺
	Image_Draw_Type_Fixed_Ratio            // 固定比例,当图片超出显示范围时,按照原始比例压缩显示图片
	Image_Draw_Type_Adaptive_Border        // 九宫格不绘制中间区域
)

// ListItemTemp_Type_

const (
	ListItemTemp_Type_Tree           = 0x01                                                               // tree
	ListItemTemp_Type_ListBox        = 0x02                                                               // listBox
	ListItemTemp_Type_List_Head      = 0x04                                                               // list 列表头
	ListItemTemp_Type_List_Item      = 0x08                                                               // list 列表项
	ListItemTemp_Type_ListView_Group = 0x10                                                               // listView 列表视组
	ListItemTemp_Type_ListView_Item  = 0x20                                                               // listView 列表视项
	ListItemTemp_Type_List           = ListItemTemp_Type_List_Head | ListItemTemp_Type_List_Item          // list (列表头)与(列表项)组合
	ListItemTemp_Type_ListView       = ListItemTemp_Type_ListView_Group | ListItemTemp_Type_ListView_Item // listView (列表视组)与(列表视项)组合
)

// Window_Position_

const (
	Window_Position_Top    = iota // top
	Window_Position_Bottom        // bottom
	Window_Position_Left          // left
	Window_Position_Right         // right
	Window_Position_Body          // body
	Window_Position_Window        // window 整个窗口
	Window_Position_Error  = -1   // 错误
)

//  List_Item_State_

const (
	List_Item_State_Leave  = iota // 项鼠标离开状态
	List_Item_State_Stay          // 项鼠标停留状态
	List_Item_State_Select        // 项选择状态
	List_Item_State_Cache         // 缓存的项
)

// Tree_Item_State_

const (
	Tree_Item_State_Leave  = iota // 项鼠标离开状态
	Tree_Item_State_Stay          // 项鼠标停留状态
	Tree_Item_State_Select        // 项选择状态
)

// List_DrawItemBk_Flag_

const (
	List_DrawItemBk_Flag_Nothing     = 0x000 // 不绘制
	List_DrawItemBk_Flag_Leave       = 0x001 // 绘制鼠标离开状态项背景
	List_DrawItemBk_Flag_Stay        = 0x002 // 绘制鼠标选择状态项背景
	List_DrawItemBk_Flag_Select      = 0x004 // 绘制鼠标停留状态项项背景
	List_DrawItemBk_Flag_Group_Leave = 0x008 // 绘制鼠标离开状态组背景,当项为组时
	List_DrawItemBk_Flag_Group_Stay  = 0x010 // 绘制鼠标停留状态组背景,当项为组时
)

//  PropertyGrid_Item_Type_

const (
	PropertyGrid_Item_Type_Text       = iota // 默认,字符串类型
	PropertyGrid_Item_Type_Edit              // 编辑框
	PropertyGrid_Item_Type_Edit_Color        // 颜色选择元素
	PropertyGrid_Item_Type_Edit_File         // 文件选择编辑框
	PropertyGrid_Item_Type_Edit_Set          // 设置
	PropertyGrid_Item_Type_ComboBox          // 组合框
	PropertyGrid_Item_Type_Group             // 组
	PropertyGrid_Item_Type_Panel             // 面板
)

// Zorder_

const (
	Zorder_Top    = iota // 最上面
	Zorder_Bottom        // 最下面
	Zorder_Before        // 指定目标下面
	Zorder_After         // 指定目标上面
)

// Pane_State_

const (
	Pane_State_Lock  = iota // 锁定
	Pane_State_Dock         // 停靠码头
	Pane_State_Float        // 浮动窗格
)

//  Window_State_Flag_

const (
	Window_State_Flag_Nothing      = 0x0000     // 无
	Window_State_Flag_Leave        = 0x0001     // 整个窗口
	Window_State_Flag_Body_Leave   = 0x0002     // 窗口-body
	Window_State_Flag_Top_Leave    = 0x0004     // 窗口-top
	Window_State_Flag_Bottom_Leave = 0x0008     // 窗口-bottom
	Window_State_Flag_Left_Leave   = 0x0010     // 窗口-left
	Window_State_Flag_Right_Leave  = 0x0020     // 窗口-right
	Window_State_Flag_Layout_Body  = 0x20000000 // 布局内容区
)

// Tree_State_Flag_

const (
	Tree_State_Flag_Item_Leave     = 0x0080 // 项鼠标离开
	Tree_State_Flag_Item_Stay      = 0x0100 // 项鼠标停留,保留值, 暂未使用
	Tree_State_Flag_Item_Select    = 0x0200 // 项选择
	Tree_State_Flag_Item_Select_No = 0x0400 // 项未选择
	Tree_State_Flag_Group          = 0x0800 // 项为组
	Tree_State_Flag_Group_No       = 0x1000 // 项不为组
)

// Element_State_Flag_

const (
	Element_State_Flag_Nothing    = Window_State_Flag_Nothing     // 无
	Element_State_Flag_Enable     = 0x0001                        // 启用
	Element_State_Flag_Disable    = 0x0002                        // 禁用
	Element_State_Flag_Focus      = 0x0004                        // 焦点
	Element_State_Flag_Focus_No   = 0x0008                        // 无焦点
	Element_State_Flag_FocusEx    = 0x40000000                    // 该元素或该元素的子元素拥有焦点
	Element_State_Flag_FocusEx_No = 0x80000000                    // 无焦点Ex
	Layout_State_Flag_Layout_Body = Window_State_Flag_Layout_Body // 布局内容区
	Element_State_Flag_Leave      = 0x0010                        // 鼠标离开
	Element_State_Flag_Stay       = 0x0020                        // 为扩展模块保留
	Element_State_Flag_Down       = 0x0040                        // 为扩展模块保留
)

// Button_State_Flag_

const (
	Button_State_Flag_Leave          = Element_State_Flag_Leave // 鼠标离开
	Button_State_Flag_Stay           = Element_State_Flag_Stay  // 鼠标停留
	Button_State_Flag_Down           = Element_State_Flag_Down  // 鼠标按下
	Button_State_Flag_Check          = 0x0080                   // 选中
	Button_State_Flag_Check_No       = 0x0100                   // 未选中
	Button_State_Flag_WindowRestore  = 0x0200                   // 窗口还原
	Button_State_Flag_WindowMaximize = 0x0400                   // 窗口最大化
)

// ComboBox_State_Flag_

const (
	ComboBox_State_Flag_Leave = Element_State_Flag_Leave // 鼠标离开
	ComboBox_State_Flag_Stay  = Element_State_Flag_Stay  // 鼠标停留
	ComboBox_State_Flag_Down  = Element_State_Flag_Down  // 鼠标按下
)

// ListBox_State_Flag

const (
	ListBox_State_Flag_Item_Leave     = 0x0080 // 项鼠标离开
	ListBox_State_Flag_Item_Stay      = 0x0100 // 项鼠标停留
	ListBox_State_Flag_Item_Select    = 0x0200 // 项选择
	ListBox_State_Flag_Item_Select_No = 0x0400 // 项未选择
)

// ListView_State_Flag_

const (
	ListView_State_Flag_Item_Leave      = 0x0080 // 项鼠标离开
	ListView_State_Flag_Item_Stay       = 0x0100 // 项鼠标停留
	ListView_State_Flag_Item_Select     = 0x0200 // 项选择
	ListView_State_Flag_Item_Select_No  = 0x0400 // 项未选择
	ListView_State_Flag_Group_Leave     = 0x0800 // 组鼠标离开
	ListView_State_Flag_Group_Stay      = 0x1000 // 组鼠标停留
	ListView_State_Flag_Group_Select    = 0x2000 // 组选择
	ListView_State_Flag_Group_Select_No = 0x4000 // 组未选择
)

//  MonthCal_State_Flag_

const (
	MonthCal_State_Flag_Leave           = Element_State_Flag_Leave // 离开状态
	MonthCal_State_Flag_Item_Leave      = 0x0080                   // 项-离开
	MonthCal_State_Flag_Item_Stay       = 0x0100                   // 项-停留
	MonthCal_State_Flag_Item_Down       = 0x0200                   // 项-按下
	MonthCal_State_Flag_Item_Select     = 0x0400                   // 项-选择
	MonthCal_State_Flag_Item_Select_No  = 0x0800                   // 项-未选择
	MonthCal_State_Flag_Item_Today      = 0x1000                   // 项-今天
	MonthCal_State_Flag_Item_Other      = 0x2000                   // 项-上月及下月
	MonthCal_State_Flag_Item_Last_Month = 0x4000                   // 项-上月
	MonthCal_State_Flag_Item_Cur_Month  = 0x8000                   // 项-当月
	MonthCal_State_Flag_Item_Next_Month = 0x10000                  // 项-下月
)

//  Layout_State_Flag_

const (
	Layout_State_Flag_Nothing = iota // 无
)
