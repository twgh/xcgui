// xcgui const
package xcc

// 窗口消息
// SW_

const (
	SW_SHOW           = 5 // 显示
	SW_HIDE           = 0 // 隐藏
	SW_SHOWMAXIMIZED  = 3 // 最大化
	SW_SHOWMINIMIZED  = 2 // 最小化
	SW_SHOWNOACTIVATE = 4 // 不激活
	SW_SHOWNA         = 8 // 原来的状态显示
)

// XC_NAME

const (
	XC_NAME1 = "name1"
	XC_NAME2 = "name2"
	XC_NAME3 = "name3"
	XC_NAME4 = "name4"
	XC_NAME5 = "name5"
	XC_NAME6 = "name6"
)

// 特殊ID

const (
	XC_ID_ROOT  = 0  // 根节点
	XC_ID_ERROR = -1 // ID错误
	XC_ID_FIRST = -2 // 插入开始位置
	XC_ID_LAST  = -3 // 插入末尾位置
)

// 调整布局标识位
// AdjustLayout_

const (
	AdjustLayout_No   = iota // 不调整布局
	AdjustLayout_All         // 强制调整自身和子对象布局
	AdjustLayout_Self        // 只调整自身布局, 不调整子对象布局
)

// 炫彩窗口透明标识
// Window_Transparent_

const (
	Window_Transparent_False  = iota // 窗口透明_无
	Window_Transparent_Shaped        // 窗口透明_异型
	Window_Transparent_Shadow        // 窗口透明_阴影,带透明通道，边框阴影，窗口透明或半透明。
	Window_Transparent_Simple        // 窗口透明_简单,指定透明色.
	Window_Transparent_Win7          // 窗口透明_玻璃
)

// 布局轴对齐
// Layout_Align_Axis_

const (
	Layout_Align_Axis_Auto   = iota // 无
	Layout_Align_Axis_Start         // 水平布局(顶部), 垂直布局(左侧)
	Layout_Align_Axis_Center        // 居中
	Layout_Align_Axis_End           // 水平布局(底部), 垂直布局(右侧)
)

// 布局大小类型
// Layout_Size_

const (
	Layout_Size_Fixed   = iota // 固定大小
	Layout_Size_Fill           // 填充父
	Layout_Size_Auto           // 根据内容计算大小
	Layout_Size_Weight         // 按照比例分配剩余空间
	Layout_Size_Percent        // 百分比
	Layout_Size_Disable        // 不使用
)

// 布局_对齐_
// Layout_Align_

const (
	Layout_Align_Left        = iota // 左侧
	Layout_Align_Top                // 顶部
	Layout_Align_Right              // 右侧
	Layout_Align_Bottom             // 底部
	Layout_Align_Center             // 居中
	Layout_Align_Equidistant        // 等距
)

// 对象样式(用于区分外观)
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

// 弹出消息框
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

// 普通三种状态
// Common_State3_

const (
	Common_State3_Leave = iota // 离开
	Common_State3_Stay         // 停留
	Common_State3_Down         // 按下
)

// 按钮状态
// Button_State_

const (
	Button_State_Leave   = iota // 离开状态
	Button_State_Stay           // 停留状态
	Button_State_Down           // 按下状态
	Button_State_Check          // 选中状态
	Button_State_Disable        // 禁用状态
)

// 对象扩展类型(功能扩展)
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

// 文本对齐
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

// 按钮图标对齐方式
// Button_Icon_Align_

const (
	Button_Icon_Align_Left   = iota // 图标在左边
	Button_Icon_Align_Top           // 图标在顶部
	Button_Icon_Align_Right         // 图标在右边
	Button_Icon_Align_Bottom        // 图标在底部
)

// 窗口样式
// Window_Style_

const (
	Window_Style_Nothing         = 0x0000 // 什么也没有
	Window_Style_Caption         = 0x0001 // 标题栏
	Window_Style_Border          = 0x0002 // 边框,如果没有指定,那么边框大小为0
	Window_Style_Center          = 0x0004 // 窗口居中
	Window_Style_Drag_Border     = 0x0008 // 拖动窗口边框
	Window_Style_Drag_Window     = 0x0010 // 拖动窗口
	Window_Style_Allow_MaxWindow = 0x0020 // 允许窗口最大化
	Window_Style_Icon            = 0x0040 // 图标
	Window_Style_Title           = 0x0080 // 标题
	Window_Style_Btn_Min         = 0x0100 // 控制按钮-最小化
	Window_Style_Btn_Max         = 0x0200 // 控制按钮-最大化
	Window_Style_Btn_Close       = 0x0400 // 控制按钮-关闭

	Window_Style_Default = (Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow | Window_Style_Icon | Window_Style_Title | Window_Style_Btn_Min | Window_Style_Btn_Max | Window_Style_Btn_Close) // 窗口样式-控制按钮: 居中 图标, 标题, 关闭按钮, 最大化按钮, 最小化按钮

	Window_Style_Simple = (Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow) // 窗口样式-简单: 居中

	Window_Style_Pop = (Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow | Window_Style_Icon |
		Window_Style_Title) // 窗口样式-弹出窗口: 图标, 标题, 关闭按钮

	Window_Style_Modal = (Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Icon | Window_Style_Title | Window_Style_Btn_Close) // 模态窗口样式-控制按钮: 居中, 图标, 标题, 关闭按钮

	Window_Style_Modal_Simple = (Window_Style_Caption | Window_Style_Border | Window_Style_Center) // 模态窗口样式-简单: 居中
)

// 对象句柄类型
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

// UI元素位置
// Element_Position_

const (
	Element_Position_No     = 0x00 // 无效
	Element_Position_Left   = 0x01 // 左边
	Element_Position_Top    = 0x02 // 上边
	Element_Position_Right  = 0x04 // 右边
	Element_Position_Bottom = 0x08 // 下边
)

// 位置标识
// Position_Flag_

const (
	Position_Flag_Left        = iota // 左
	Position_Flag_Top                // 上
	Position_Flag_Right              // 右
	Position_Flag_Bottom             // 下
	Position_Flag_LeftTop            // 左上角
	Position_Flag_LeftBottom         // 左下角
	Position_Flag_RightTop           // 右上角
	Position_Flag_RightBottom        // 右下角
	Position_Flag_Center             // 中心
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
	XE_EDIT_SET                           = 162 // 编辑框_置文本
	XE_EDIT_DRAWROW                       = 181 // 和XE_EDIT_CHANGED的对换?
	XE_EDIT_CHANGED                       = 182 // 编辑框_内容被改变
	XE_EDIT_POS_CHANGED                   = 183 // 编辑框_光标位置_被改变
	XE_EDIT_STYLE_CHANGED                 = 184 // 编辑框_样式_被改变
	XE_EDIT_ENTER_GET_TABALIGN            = 185 // 编辑框_回车_获取标签?
	XE_EDITOR_MODIFY_ROWS                 = 186 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化
	XE_EDITOR_SETBREAKPOINT               = 191 // 代码编辑框_设置断点
	XE_EDITOR_REMOVEBREAKPOINT            = 192 // 代码编辑框_移除断点
	XE_EDIT_ROW_CHANGED                   = 193 // 编辑框_行_被改变
	XE_EDITOR_AUTOMATCH_SELECT            = 194 // 编辑框_自动匹配选择
	XE_TABBAR_SELECT                      = 221 // TabBar标签按钮选择改变事件
	XE_TABBAR_DELETE                      = 222 // TabBar标签按钮删除事件
	XE_MONTHCAL_CHANGE                    = 231 // 月历元素日期改变事件
	XE_DATETIME_CHANGE                    = 241 // 日期时间元素,内容改变事件
	XE_DATETIME_POPUP_MONTHCAL            = 242 // 日期时间元素,弹出月历卡片事件
	XE_DATETIME_EXIT_MONTHCAL             = 243 // 日期时间元素,弹出的月历卡片退出事件
	XE_DROPFILES                          = 250 // 文件拖放事件.
)

// 窗格_对齐_
// Pane_Align_

const (
	Pane_Align_Left   = iota // 左侧
	Pane_Align_Top           // 顶部
	Pane_Align_Right         // 右侧
	Pane_Align_Bottom        // 底部
	Pane_Align_Center        // 居中
	Pane_Align_Error  = -1   // 错误
)

// 弹出菜单项标识
// Menu_Item_Flag_

const (
	Menu_Item_Flag_Normal    = 0x00 // 正常
	Menu_Item_Flag_Select    = 0x01 // 选择
	Menu_Item_Flag_Check     = 0x02 // 勾选
	Menu_Item_Flag_Popup     = 0x04 // 弹出
	Menu_Item_Flag_Separator = 0x08 // 分隔栏 ID号任意,ID号被忽略
	Menu_Item_Flag_Disable   = 0x10 // 禁用
)

// 弹出菜单方向
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

// ComboBox状态
// ComboBox_State_

const (
	ComboBox_State_Leave = iota // 鼠标离开状态
	ComboBox_State_Stay         // 鼠标停留状态
	ComboBox_State_Down         // 按下状态
)

// 数据适配器数据类型
// Adapter_Date_Type_

const (
	Adapter_Date_Type_Int    = iota // 整形
	Adapter_Date_Type_Float         // 浮点型
	Adapter_Date_Type_String        // 字符串
	Adapter_Date_Type_Image         // 图片
)

// Edit 聊天气泡对齐方式
// Chat_Flag_

const (
	Chat_Flag_Left            = 0x1 // 左侧
	Chat_Flag_Right           = 0x2 // 右侧
	Chat_Flag_Center          = 0x4 // 中间
	Chat_Flag_Next_Row_Bubble = 0x8 // 下一行显示气泡
)

// 编辑框类型
// Edit_Type_

const (
	Edit_Type_None      = iota // 普通编辑框, 每行的高度相同
	Edit_Type_Editor           // 代码编辑
	Edit_Type_Richedit         // 富文本编辑框, 每行的高度可能不同
	Edit_Type_Chat             // 聊天气泡, 每行的高度可能不同
	Edit_Type_CodeTable        // 代码表格,内部使用, 每行的高度相同
)

// 编辑框风格类型
// Edit_Style_Type_

const (
	Edit_Style_Type_Font_Color = iota + 1 // 字体
	Edit_Style_Type_Image                 // 图片
	Edit_Style_Type_Obj                   // UI对象
)

// 编辑框文本对齐标志
// Edit_TextAlign_Flag_

const (
	Edit_TextAlign_Flag_Left     = 0x0 // 左侧
	Edit_TextAlign_Flag_Right    = 0x1 // 右侧
	Edit_TextAlign_Flag_Center   = 0x2 // 水平居中
	Edit_TextAlign_Flag_Top      = 0x0 // 顶部
	Edit_TextAlign_Flag_Bottom   = 0x4 // 底部
	Edit_TextAlign_Flag_Center_V = 0x8 // 垂直居中
)

// Table 标识
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

// 缓动类型
// Ease_Type_

const (
	Ease_Type_In    = iota // 从慢到快
	Ease_Type_Out          // 从快到慢
	Ease_Type_InOut        // 从慢到快再到慢
)

// 缓动标识
// Ease_Flag_

const (
	Ease_Flag_Linear = iota // 线性, 直线
	Ease_Flag_Quad          // 二次方曲线
	Ease_Flag_Cubic         // 三次方曲线, 圆弧
	Ease_Flag_Quart         // 四次方曲线
	Ease_Flag_Quint         // 五次方曲线

	Ease_Flag_Sine    // 正弦, 在末端变化
	Ease_Flag_Expo    // 突击, 突然一下
	Ease_Flag_Circ    // 圆环, 好比绕过一个圆环
	Ease_Flag_Elastic // 强力回弹
	Ease_Flag_Back    // 回弹, 比较缓慢
	Ease_Flag_Bounce  // 弹跳, 模拟小球落地弹跳

	Ease_Flag_In    = 0x010000 // 从慢到快
	Ease_Flag_Out   = 0x020000 // 从快到慢
	Ease_Flag_InOut = 0x030000 // 从慢到快再到慢
)

// 字体样式
// FontStyle_

const (
	FontStyle_Regular    = 0 // 正常
	FontStyle_Bold       = 1 // 粗体
	FontStyle_Italic     = 2 // 斜体
	FontStyle_BoldItalic = 3 // 粗斜体
	FontStyle_Underline  = 4 // 下划线
	FontStyle_Strikeout  = 8 // 删除线
)

// 图片绘制类型
// Image_Draw_Type_

const (
	Image_Draw_Type_Default         = iota // 默认
	Image_Draw_Type_Stretch                // 拉伸
	Image_Draw_Type_Adaptive               // 自适应,九宫格
	Image_Draw_Type_Tile                   // 平铺
	Image_Draw_Type_Fixed_Ratio            // 固定比例,当图片超出显示范围时,按照原始比例压缩显示图片
	Image_Draw_Type_Adaptive_Border        // 九宫格不绘制中间区域
)

// 列表项模板类型
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

// 窗口位置
// Window_Position_

const (
	Window_Position_Top    = iota // 上
	Window_Position_Bottom        // 下
	Window_Position_Left          // 左
	Window_Position_Right         // 右
	Window_Position_Body          // body
	Window_Position_Window        // window 整个窗口
	Window_Position_Error  = -1   // 错误
)

// 窗口位置

const (
	WINDOW_TOP         = iota + 1 // 上
	WINDOW_BOTTOM                 // 下
	WINDOW_LEFT                   // 左
	WINDOW_RIGHT                  // 右
	WINDOW_TOPLEFT                // 左上角
	WINDOW_TOPRIGHT               // 右上角
	WINDOW_BOTTOMLEFT             // 左下角
	WINDOW_BOTTOMRIGHT            // 右下角
	WINDOW_CAPTION                // 标题栏移动窗口区域
	WINDOW_BODY
)

// List项状态
// List_Item_State_

const (
	List_Item_State_Leave  = iota // 项鼠标离开状态
	List_Item_State_Stay          // 项鼠标停留状态
	List_Item_State_Select        // 项选择状态
	List_Item_State_Cache         // 缓存的项
)

// Tree项状态
// Tree_Item_State_

const (
	Tree_Item_State_Leave  = iota // 项鼠标离开状态
	Tree_Item_State_Stay          // 项鼠标停留状态
	Tree_Item_State_Select        // 项选择状态
)

// 项背景绘制标志位(List, ListBox, ListView, Tree)
// List_DrawItemBk_Flag_

const (
	List_DrawItemBk_Flag_Leave       = 1 << iota // 绘制鼠标离开状态项背景
	List_DrawItemBk_Flag_Stay                    // 绘制鼠标选择状态项背景
	List_DrawItemBk_Flag_Select                  // 绘制鼠标停留状态项项背景
	List_DrawItemBk_Flag_Group_Leave             // 绘制鼠标离开状态组背景, 当项为组时
	List_DrawItemBk_Flag_Group_Stay              // 绘制鼠标停留状态组背景, 当项为组时
	List_DrawItemBk_Flag_Line                    // 列表绘制水平分割线
	List_DrawItemBk_Flag_LineV                   // 列表绘制垂直分割线
	List_DrawItemBk_Flag_Nothing     = 0         // 不绘制
)

// 属性网格项类型
// PropertyGrid_Item_Type_

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

// Z序位置
// Zorder_

const (
	Zorder_Top    = iota // 最上面
	Zorder_Bottom        // 最下面
	Zorder_Before        // 指定目标下面
	Zorder_After         // 指定目标上面
)

// 窗格状态
// Pane_State_

const (
	Pane_State_Lock  = iota // 锁定
	Pane_State_Dock         // 停靠码头
	Pane_State_Float        // 浮动窗格
)

// 组合状态
// Window_State_Flag_

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

// 组合状态
// 列表树状态
// Tree_State_Flag_

const (
	Tree_State_Flag_Item_Leave     = 0x0080 // 项鼠标离开
	Tree_State_Flag_Item_Stay      = 0x0100 // 项鼠标停留,保留值, 暂未使用
	Tree_State_Flag_Item_Select    = 0x0200 // 项选择
	Tree_State_Flag_Item_Select_No = 0x0400 // 项未选择
	Tree_State_Flag_Group          = 0x0800 // 项为组
	Tree_State_Flag_Group_No       = 0x1000 // 项不为组
)

// 组合状态
// 元素状态标志
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

// 组合状态
// 按钮状态标志
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

// 组合状态
// 组合框状态标志
// ComboBox_State_Flag_

const (
	ComboBox_State_Flag_Leave = Element_State_Flag_Leave // 鼠标离开
	ComboBox_State_Flag_Stay  = Element_State_Flag_Stay  // 鼠标停留
	ComboBox_State_Flag_Down  = Element_State_Flag_Down  // 鼠标按下
)

// 组合状态
// 列表项状态标志
// List_State_Flag_Item_

const (
	List_State_Flag_Item_Leave     = iota // 项鼠标离开
	List_State_Flag_Item_Stay             // 项鼠标停留
	List_State_Flag_Item_Select           // 项选择
	List_State_Flag_Item_Select_No        // 项未选择
)

// 组合状态
// 列表框项状态标志
// ListBox_State_Flag

const (
	ListBox_State_Flag_Item_Leave     = 0x0080 // 项鼠标离开
	ListBox_State_Flag_Item_Stay      = 0x0100 // 项鼠标停留
	ListBox_State_Flag_Item_Select    = 0x0200 // 项选择
	ListBox_State_Flag_Item_Select_No = 0x0400 // 项未选择
)

// 组合状态
// 列表视图状态标志
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

// 组合状态
// 列表头状态
// ListHeader_State_Flag_Item_

const (
	ListHeader_State_Flag_Item_Leave = iota // 项鼠标离开
	ListHeader_State_Flag_Item_Stay         // 项鼠标停留
	ListHeader_State_Flag_Item_Down         // 项鼠标按下
)

// 组合状态
// 月历卡片状态标志
// MonthCal_State_Flag_

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

// 组合状态
// 布局状态标志
// Layout_State_Flag_

const (
	Layout_State_Flag_Nothing = Window_State_Flag_Nothing // 无
	Layout_State_Flag_Full    = 0x0001                    // 完整背景
	Layout_State_Flag_Body    = 0x0002                    // 内容区域, 不包含边大小
)

// 组合状态
// 属性网格状态
// propertyGrid_state_flag_

const (
	PropertyGrid_State_Flag_Item_Leave      = 0x0080 //离开
	PropertyGrid_State_Flag_Item_Stay       = 0x0100 //停留
	PropertyGrid_State_Flag_Item_Select     = 0x0200 //选择
	PropertyGrid_State_Flag_Item_Select_No  = 0x0400 //未选择
	PropertyGrid_State_Flag_Group_Leave     = 0x0800 //组离开
	PropertyGrid_State_Flag_Group_Expand    = 0x1000 //组展开
	PropertyGrid_State_Flag_Group_Expand_No = 0x2000 //组未展开
)

// 组合状态
// 窗格状态
// Pane_State_Flag_

const (
	Pane_State_Flag_Leave   = Element_State_Flag_Leave // 离开
	Pane_State_Flag_Stay    = Element_State_Flag_Stay  // 停留
	Pane_State_Flag_Caption = 0x0080                   // 标题
	Pane_State_Flag_Body    = 0x0100                   // 内容区
)

// 窗口事件
// XWM_

const (
	XWM_WINDPROC             = 0x7000 + 2  // 窗口消息过程
	XWM_XC_TIMER             = 0x7000 + 5  // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收
	XWM_MENU_POPUP           = 0x7000 + 11 // 菜单弹出
	XWM_MENU_POPUP_WND       = 0x7000 + 12 // 菜单弹出窗口
	XWM_MENU_SELECT          = 0x7000 + 13 // 菜单选择
	XWM_MENU_EXIT            = 0x7000 + 14 // 菜单退出
	XWM_MENU_DRAW_BACKGROUND = 0x7000 + 15 // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
	XWM_MENU_DRAWITEM        = 0x7000 + 16 // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
	XWM_FLOAT_PANE           = 0x7000 + 18 // 浮动窗格
	XWM_PAINT_END            = 0x7000 + 19 // 窗口绘制完成消息
	XWM_PAINT_DISPLAY        = 0x7000 + 20 // 窗口绘制完成并且已经显示到屏幕
	XWM_DOCK_POPUP           = 0x7000 + 21 // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格

	// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
	//
	// hFloatWnd: 拖动的浮动窗口句柄.
	//
	// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
	XWM_FLOATWND_DRAG = 0x7000 + 22
)

// 窗口事件
// WM_

const (
	WM_PAINT          = 15  // 窗口绘制消息
	WM_CLOSE          = 16  // 窗口关闭消息.
	WM_DESTROY        = 2   // 窗口销毁消息.
	WM_NCDESTROY      = 130 // 窗口非客户区销毁消息.
	WM_MOUSEMOVE      = 512 // 窗口鼠标移动消息.
	WM_LBUTTONDOWN    = 513 // 窗口鼠标左键按下消息
	WM_LBUTTONUP      = 514 // 窗口鼠标左键弹起消息.
	WM_RBUTTONDOWN    = 516 // 窗口鼠标右键按下消息.
	WM_RBUTTONUP      = 517 // 窗口鼠标右键弹起消息.
	WM_LBUTTONDBLCLK  = 515 // 窗口鼠标左键双击消息.
	WM_RBUTTONDBLCLK  = 518 // 窗口鼠标右键双击消息.
	WM_MOUSEWHEEL     = 522 // 窗口鼠标滚轮滚动消息.
	WM_EXITSIZEMOVE   = 562 // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
	WM_MOUSEHOVER     = 673 // 窗口鼠标进入消息
	WM_MOUSELEAVE     = 675 // 窗口鼠标离开消息.
	WM_SIZE           = 5   // 窗口大小改变消息.
	WM_TIMER          = 275 // 窗口定时器消息.
	WM_SETFOCUS       = 7   // 窗口获得焦点.
	WM_KILLFOCUS      = 8   // 窗口失去焦点.
	WM_KEYDOWN        = 256 // 窗口键盘按键消息.
	WM_CAPTURECHANGED = 533 // 窗口鼠标捕获改变消息.
	WM_SETCURSOR      = 32  // 窗口设置鼠标光标.
	WM_CHAR           = 258 // 窗口字符消息.
	WM_DROPFILES      = 563 // 拖动文件到窗口.
)

// D2D文本渲染模式
// XC_DWRITE_RENDERING_MODE_

const (
	XC_DWRITE_RENDERING_MODE_DEFAULT                     = iota // 指定根据字体和大小自动确定呈现模式。
	XC_DWRITE_RENDERING_MODE_ALIASED                            // 指定不执行抗锯齿。 每个像素要么设置为文本的前景色，要么保留背景的颜色。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_CLASSIC              // 使用与别名文本相同的度量指定 ClearType 呈现。 字形只能定位在整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_NATURAL              // 使用使用 CLEARTYPE_NATURAL_QUALITY 创建的字体，使用与使用 GDI 的文本呈现相同的指标指定 ClearType 呈现。 与使用别名文本相比，字形度量更接近其理想值，但字形仍然位于整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL                  // 仅在水平维度中指定具有抗锯齿功能的 ClearType 渲染。 这通常用于中小字体大小（最多 16 ppem）。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL_SYMMETRIC        // 指定在水平和垂直维度上具有抗锯齿的 ClearType 渲染。这通常用于较大的尺寸，以使曲线和对角线看起来更平滑，但会牺牲一些柔和度。
	XC_DWRITE_RENDERING_MODE_OUTLINE                            // 指定渲染应绕过光栅化器并直接使用轮廓。 这通常用于非常大的尺寸。
)

// 月历元素上的按钮类型
// MonthCal_Button_Type_

const (
	MonthCal_Button_Type_Today      = iota // 今天按钮
	MonthCal_Button_Type_Last_Year         // 上一年
	MonthCal_Button_Type_Next_Year         // 下一年
	MonthCal_Button_Type_Last_Month        // 上一月
	MonthCal_Button_Type_Next_Month        //下一月
)

// 窗格菜单 当前未使用

const (
	IDM_LOCK  = 1000000006 // 锁定
	IDM_DOCK  = 1000000007 // 停靠
	IDM_FLOAT = 1000000008 // 浮动
	IDM_HIDE  = 1000000009 // 隐藏

	Edit_Style_No = 0xFFFF // 无效样式
)

// 菜单ID, 当前未使用

const (
	IDM_CLIP      = 1000000000 // 剪切
	IDM_COPY      = 1000000001 // 复制
	IDM_PASTE     = 1000000002 // 粘贴
	IDM_DELETE    = 1000000003 // 删除
	IDM_SELECTALL = 1000000004 // 全选
	IDM_DELETEALL = 1000000005 // 清空
)

// 通知消息外观
// NotifyMsg_Skin_

const (
	NotifyMsg_Skin_No      = iota // 默认
	NotifyMsg_Skin_Success        // 成功
	NotifyMsg_Skin_Warning        // 警告
	NotifyMsg_Skin_Message        // 消息
	NotifyMsg_Skin_Error          // 错误
)

// 动画移动标识
// Animation_Move_

const (
	Animation_Move_X = 0x01 // X轴移动.
	Animation_Move_Y = 0x02 // Y轴移动.
)

// 背景对象对齐方式
// BkObject_Align_Flag_

const (
	BkObject_Align_Flag_Left     = 1 << iota // 左对齐, 当设置此标识时, 外间距(margin.left)代表左侧间距; 当right未设置时,那么外间距(margin.right)代表宽度;
	BkObject_Align_Flag_Top                  // 顶对齐, 当设置此标识时, 外间距(margin.top)代表顶部间距; 当bottom未设置时,那么外间距(margin.bottom)代表高度;
	BkObject_Align_Flag_Right                // 右对齐, 当设置此标识时, 外间距(margin.right)代表右侧间距; 当left未设置时,那么外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Bottom               // 底对齐, 当设置此标识时, 外间距(margin.bottom)代表底部间距; 当top未设置时,那么外间距(margin.top)代表高度;
	BkObject_Align_Flag_Center               // 水平居中, 当设置此标识时, 外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Center_v             // 垂直居中, 当设置此标识时, 外间距(margin.top)代表高度;
	BkObject_Align_Flag_No       = 0         // 无
)
