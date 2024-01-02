package xcc

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

// SW_ 是显示窗口的形式.
type SW_ int32

const (
	SW_HIDE            SW_ = iota // 隐藏窗口，活动状态给另一个窗口
	SW_SHOWNORMAL                 // 用原来的大小和位置显示一个窗口，并将其激活
	SW_SHOWMINIMIZED              // 最小化窗口，并将其激活
	SW_SHOWMAXIMIZED              // 最大化窗口，并将其激活
	SW_SHOWNOACTIVATE             // 用最近的大小和位置显示一个窗口，同时不激活该窗口
	SW_SHOW                       // 用当前的大小和位置显示一个窗口，并将其激活
	SW_MINIMIZE                   // 最小化窗口，活动状态给另一个窗口
	SW_SHOWMINNOACTIVE            // 最小化窗口，同时不改变活动窗口
	SW_SHOWNA                     // 原来的状态显示
	SW_RESTORE                    // 激活并还原显示窗口
	SW_SHOWDEFAULT                // 按默认状态显示
	SW_FORCEMINIMIZE              // 最小化窗口，即使拥有该窗口的线程被挂起也会最小化。在从其他线程最小化窗口时才使用这个参数
)

// AdjustLayout_ 是调整布局标识位.
type AdjustLayout_ int32

const (
	AdjustLayout_No   AdjustLayout_ = iota // 不调整布局
	AdjustLayout_All                       // 强制调整自身和子对象布局
	AdjustLayout_Self                      // 只调整自身布局, 不调整子对象布局
)

// Window_Transparent_ 是炫彩窗口透明标识.
type Window_Transparent_ int

const (
	Window_Transparent_False  Window_Transparent_ = iota // 默认窗口, 不透明
	Window_Transparent_Shaped                            // 透明窗口, 带透明通道, 异型
	Window_Transparent_Shadow                            // 阴影窗口, 带透明通道, 边框阴影, 窗口透明或半透明
	Window_Transparent_Simple                            // 透明窗口, 不带透明通道, 指定半透明度, 指定透明色
	Window_Transparent_Win7                              // WIN7玻璃窗口, 需要WIN7开启特效, 当前未启用
)

// Layout_Align_Axis_ 是布局轴对齐.
type Layout_Align_Axis_ int

const (
	Layout_Align_Axis_Auto   Layout_Align_Axis_ = iota // 无
	Layout_Align_Axis_Start                            // 水平布局(顶部), 垂直布局(左侧)
	Layout_Align_Axis_Center                           // 居中
	Layout_Align_Axis_End                              // 水平布局(底部), 垂直布局(右侧)
)

// Layout_Size_ 是布局大小类型.
type Layout_Size_ int

const (
	Layout_Size_Fixed   Layout_Size_ = iota // 固定大小
	Layout_Size_Fill                        // 填充父
	Layout_Size_Auto                        // 自动大小, 根据内容计算大小
	Layout_Size_Weight                      // 按照比例分配剩余空间
	Layout_Size_Percent                     // 百分比
	Layout_Size_Disable                     // 不使用
)

// Layout_Align_ 是布局_对齐.
type Layout_Align_ int

const (
	Layout_Align_Left        Layout_Align_ = iota // 左侧
	Layout_Align_Top                              // 顶部
	Layout_Align_Right                            // 右侧
	Layout_Align_Bottom                           // 底部
	Layout_Align_Center                           // 居中
	Layout_Align_Equidistant                      // 等距
)

// XC_OBJECT_STYLE 对象样式(用于区分外观).
type XC_OBJECT_STYLE int

const (
	Button_Style_Default XC_OBJECT_STYLE = iota // 默认风格
	Button_Style_Radio                          // 单选按钮
	Button_Style_Check                          // 复选按钮
	Button_Style_Icon                           // 图标按钮
	Button_Style_Expand                         // 展开收缩按钮

	Button_Style_Close // 关闭按钮
	Button_Style_Max   // 最大化按钮
	Button_Style_Min   // 最小化按钮

	Button_Style_Scrollbar_Left     // 水平滚动条-左按钮
	Button_Style_Scrollbar_Right    // 水平滚动条-右按钮
	Button_Style_Scrollbar_Up       // 垂直滚动条-上按钮
	Button_Style_Scrollbar_Down     // 垂直滚动条-下按钮
	Button_Style_Scrollbar_Slider_H // 水平滚动条-滑块
	Button_Style_Scrollbar_Slider_V // 垂直滚动条-滑块

	Button_Style_TabBar // TabBar-按钮
	Button_Style_Slider // 滑动条-滑块

	Button_Style_ToolBar       // 工具条-按钮
	Button_Style_ToolBar_Left  // 工具条-左滚动按钮
	Button_Style_ToolBar_Right // 工具条-右滚动按钮

	Button_Style_Pane_Close // 窗格-关闭按钮
	Button_Style_Pane_Lock  // 窗格-锁定按钮
	Button_Style_Pane_Menu  // 窗格-下拉菜单按钮

	Button_Style_Pane_Dock_Left   // 窗格-码头按钮左
	Button_Style_Pane_Dock_Top    // 窗格-码头按钮上
	Button_Style_Pane_Dock_Right  // 窗格-码头按钮右
	Button_Style_Pane_Dock_Bottom // 窗格-码头按钮下

	Element_Style_FrameWnd_Dock_Left   // 框架窗口-停靠码头左
	Element_Style_FrameWnd_Dock_Top    // 框架窗口-停靠码头上
	Element_Style_FrameWnd_Dock_Right  // 框架窗口-停靠码头右
	Element_Style_FrameWnd_Dock_Bottom // 框架窗口-停靠码头下

	Element_Style_ToolBar_Separator // 工具条-分割线
	ListBox_Style_ComboBox          // 组合框-下拉列表框, 下拉组合框弹出的ListBox
)

// MessageBox_Flag_ 弹出消息框.
type MessageBox_Flag_ int

const (
	MessageBox_Flag_Other  MessageBox_Flag_ = 0    // 其他
	MessageBox_Flag_Ok     MessageBox_Flag_ = 0x01 // 确定按钮
	MessageBox_Flag_Cancel MessageBox_Flag_ = 0x02 // 取消按钮

	MessageBox_Flag_Icon_Appicon MessageBox_Flag_ = 0x01000 // 图标 应用程序 IDI_APPLICATION
	MessageBox_Flag_Icon_Info    MessageBox_Flag_ = 0x02000 // 图标 信息 IDI_ASTERISK
	MessageBox_Flag_Icon_Qustion MessageBox_Flag_ = 0x04000 // 图标 问询/帮助/提问 IDI_QUESTION
	MessageBox_Flag_Icon_Error   MessageBox_Flag_ = 0x08000 // 图标 错误/拒绝/禁止 IDI_ERROR
	MessageBox_Flag_Icon_Warning MessageBox_Flag_ = 0x10000 // 图标 警告 IDI_WARNING
	MessageBox_Flag_Icon_Shield  MessageBox_Flag_ = 0x20000 // 图标 盾牌/安全 IDI_SHIELD
)

// Common_State3_ 普通三种状态.
type Common_State3_ int32

const (
	Common_State3_Leave Common_State3_ = iota // 离开
	Common_State3_Stay                        // 停留
	Common_State3_Down                        // 按下
)

// Button_State_ 按钮状态.
type Button_State_ int

const (
	Button_State_Leave   Button_State_ = iota // 离开状态
	Button_State_Stay                         // 停留状态
	Button_State_Down                         // 按下状态
	Button_State_Check                        // 选中状态
	Button_State_Disable                      // 禁用状态
)

// TextFormatFlag_ 文本对齐.
type TextFormatFlag_ int

const (
	TextAlignFlag_Left     TextFormatFlag_ = 0      // 左对齐
	TextAlignFlag_Top      TextFormatFlag_ = 0      // 垂直顶对齐
	TextAlignFlag_Left_Top TextFormatFlag_ = 0x4000 // 内部保留
	TextAlignFlag_Center   TextFormatFlag_ = 0x1    // 水平居中
	TextAlignFlag_Right    TextFormatFlag_ = 0x2    // 右对齐
	TextAlignFlag_Vcenter  TextFormatFlag_ = 0x4    // 垂直居中
	TextAlignFlag_Bottom   TextFormatFlag_ = 0x8    // 垂直底对齐

	TextFormatFlag_DirectionRightToLeft  TextFormatFlag_ = 0x10   // 从右向左顺序显示文本
	TextFormatFlag_NoWrap                TextFormatFlag_ = 0x20   // 禁止换行
	TextFormatFlag_DirectionVertical     TextFormatFlag_ = 0x40   // 垂直显示文本
	TextFormatFlag_NoFitBlackBox         TextFormatFlag_ = 0x80   // 允许部分字符延伸该字符串的布局矩形。默认情况下，将重新定位字符以避免任何延伸
	TextFormatFlag_DisplayFormatControl  TextFormatFlag_ = 0x100  // 控制字符（如从左到右标记）随具有代表性的标志符号一起显示在输出中。
	TextFormatFlag_NoFontFallback        TextFormatFlag_ = 0x200  // 对于请求的字体中不支持的字符，禁用回退到可选字体。缺失的任何字符都用缺失标志符号的字体显示，通常是一个空的方块
	TextFormatFlag_MeasureTrailingSpaces TextFormatFlag_ = 0x400  // 包括每一行结尾处的尾随空格。在默认情况下，MeasureString 方法返回的边框都将排除每一行结尾处的空格。设置此标记以便在测定时将空格包括进去
	TextFormatFlag_LineLimit             TextFormatFlag_ = 0x800  // 如果内容显示高度不够一行,那么不显示
	TextFormatFlag_NoClip                TextFormatFlag_ = 0x1000 // 允许显示标志符号的伸出部分和延伸到边框外的未换行文本。在默认情况下，延伸到边框外侧的所有文本和标志符号部分都被剪裁

	TextTrimming_None              TextFormatFlag_ = 0       // 不使用去尾
	TextTrimming_Character         TextFormatFlag_ = 0x40000 // 以字符为单位去尾
	TextTrimming_Word              TextFormatFlag_ = 0x80000 // 以单词为单位去尾
	TextTrimming_EllipsisCharacter TextFormatFlag_ = 0x8000  // 以字符为单位去尾,省略部分使用且略号表示
	TextTrimming_EllipsisWord      TextFormatFlag_ = 0x10000 // 以单词为单位去尾,
	TextTrimming_EllipsisPath      TextFormatFlag_ = 0x20000 // 略去字符串中间部分，保证字符的首尾都能够显示
)

// Button_Icon_Align_ 按钮图标对齐方式.
type Button_Icon_Align_ int

const (
	Button_Icon_Align_Left   Button_Icon_Align_ = iota // 图标在左边
	Button_Icon_Align_Top                              // 图标在顶部
	Button_Icon_Align_Right                            // 图标在右边
	Button_Icon_Align_Bottom                           // 图标在底部
)

// Window_Style_ 窗口样式.
type Window_Style_ int

const (
	Window_Style_Nothing         Window_Style_ = 0x0000 // 什么也没有
	Window_Style_Caption         Window_Style_ = 0x0001 // 标题栏
	Window_Style_Border          Window_Style_ = 0x0002 // 边框,如果没有指定,那么边框大小为0
	Window_Style_Center          Window_Style_ = 0x0004 // 窗口居中
	Window_Style_Drag_Border     Window_Style_ = 0x0008 // 拖动窗口边框
	Window_Style_Drag_Window     Window_Style_ = 0x0010 // 拖动窗口
	Window_Style_Allow_MaxWindow Window_Style_ = 0x0020 // 允许窗口最大化

	Window_Style_Icon      Window_Style_ = 0x0040 // 图标
	Window_Style_Title     Window_Style_ = 0x0080 // 标题
	Window_Style_Btn_Min   Window_Style_ = 0x0100 // 控制按钮-最小化
	Window_Style_Btn_Max   Window_Style_ = 0x0200 // 控制按钮-最大化
	Window_Style_Btn_Close Window_Style_ = 0x0400 // 控制按钮-关闭

	Window_Style_Default = Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow | Window_Style_Icon | Window_Style_Title | Window_Style_Btn_Min | Window_Style_Btn_Max | Window_Style_Btn_Close // 窗口样式-控制按钮: 居中 图标, 标题, 关闭按钮, 最大化按钮, 最小化按钮

	Window_Style_Simple = Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow // 窗口样式-简单: 居中

	Window_Style_Pop = Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Drag_Border | Window_Style_Allow_MaxWindow | Window_Style_Icon |
		Window_Style_Title // 窗口样式-弹出窗口: 图标, 标题, 关闭按钮

	Window_Style_Modal = Window_Style_Caption | Window_Style_Border | Window_Style_Center | Window_Style_Icon | Window_Style_Title | Window_Style_Btn_Close // 模态窗口样式-控制按钮: 居中, 图标, 标题, 关闭按钮

	Window_Style_Modal_Simple = Window_Style_Caption | Window_Style_Border | Window_Style_Center // 模态窗口样式-简单: 居中
)

// XC_OBJECT_TYPE 对象句柄类型.
type XC_OBJECT_TYPE int

const (
	XC_ERROR                XC_OBJECT_TYPE = -1 // 错误类型
	XC_NOTHING              XC_OBJECT_TYPE = 0
	XC_WINDOW               XC_OBJECT_TYPE = 1  // 窗口
	XC_MODALWINDOW          XC_OBJECT_TYPE = 2  // 模态窗口
	XC_FRAMEWND             XC_OBJECT_TYPE = 3  // 框架窗口
	XC_FLOATWND             XC_OBJECT_TYPE = 4  // 浮动窗口
	XC_COMBOBOXWINDOW       XC_OBJECT_TYPE = 11 // comboBoxWindow_ 组合框弹出下拉列表窗口
	XC_POPUPMENUWINDOW      XC_OBJECT_TYPE = 12 // popupMenuWindow_ 弹出菜单主窗口
	XC_POPUPMENUCHILDWINDOW XC_OBJECT_TYPE = 13 // popupMenuChildWindow_ 弹出菜单子窗口
	XC_OBJECT_UI            XC_OBJECT_TYPE = 19 // 可视对象
	XC_WIDGET_UI            XC_OBJECT_TYPE = 20 // 窗口组件
	XC_ELE                  XC_OBJECT_TYPE = 21 // 基础元素
	XC_ELE_LAYOUT           XC_OBJECT_TYPE = 53 // 布局元素
	XC_LAYOUT_FRAME         XC_OBJECT_TYPE = 54 // 流式布局
	XC_BUTTON               XC_OBJECT_TYPE = 22 // 按钮
	XC_EDIT                 XC_OBJECT_TYPE = 45 // 编辑框
	XC_EDITOR               XC_OBJECT_TYPE = 46 // 代码编辑框

	XC_RICHEDIT     XC_OBJECT_TYPE = 23 // 富文本编辑框
	XC_COMBOBOX     XC_OBJECT_TYPE = 24 // 下拉组合框
	XC_SCROLLBAR    XC_OBJECT_TYPE = 25 // 滚动条
	XC_SCROLLVIEW   XC_OBJECT_TYPE = 26 // 滚动视图
	XC_LIST         XC_OBJECT_TYPE = 27 // 列表
	XC_LISTBOX      XC_OBJECT_TYPE = 28 // 列表框
	XC_LISTVIEW     XC_OBJECT_TYPE = 29 // 列表视图,大图标
	XC_TREE         XC_OBJECT_TYPE = 30 // 列表树
	XC_MENUBAR      XC_OBJECT_TYPE = 31 // 菜单条
	XC_SLIDERBAR    XC_OBJECT_TYPE = 32 // 滑动条
	XC_PROGRESSBAR  XC_OBJECT_TYPE = 33 // 进度条
	XC_TOOLBAR      XC_OBJECT_TYPE = 34 // 工具条
	XC_MONTHCAL     XC_OBJECT_TYPE = 35 // 月历卡片
	XC_DATETIME     XC_OBJECT_TYPE = 36 // 日期时间
	XC_PROPERTYGRID XC_OBJECT_TYPE = 37 // 属性网格
	XC_EDIT_COLOR   XC_OBJECT_TYPE = 38 // 颜色选择框
	XC_EDIT_SET     XC_OBJECT_TYPE = 39 // 设置编辑框
	XC_TABBAR       XC_OBJECT_TYPE = 40 // tab条
	XC_TEXTLINK     XC_OBJECT_TYPE = 41 // 文本链接按钮

	XC_PANE           XC_OBJECT_TYPE = 42 // 窗格
	XC_PANE_SPLIT     XC_OBJECT_TYPE = 43 // 窗格拖动分割条
	XC_MENUBAR_BUTTON XC_OBJECT_TYPE = 44 // 菜单条上的按钮
	XC_EDIT_FILE      XC_OBJECT_TYPE = 50 // EditFile 文件选择编辑框
	XC_EDIT_FOLDER    XC_OBJECT_TYPE = 51 // EditFolder 文件夹选择编辑框
	XC_LIST_HEADER    XC_OBJECT_TYPE = 52 // 列表头元素

	XC_SHAPE          XC_OBJECT_TYPE = 61 // 形状对象
	XC_SHAPE_TEXT     XC_OBJECT_TYPE = 62 // 形状对象-文本
	XC_SHAPE_PICTURE  XC_OBJECT_TYPE = 63 // 形状对象-图片
	XC_SHAPE_RECT     XC_OBJECT_TYPE = 64 // 形状对象-矩形
	XC_SHAPE_ELLIPSE  XC_OBJECT_TYPE = 65 // 形状对象-圆
	XC_SHAPE_LINE     XC_OBJECT_TYPE = 66 // 形状对象-直线
	XC_SHAPE_GROUPBOX XC_OBJECT_TYPE = 67 // 形状对象-组框
	XC_SHAPE_GIF      XC_OBJECT_TYPE = 68 // 形状对象-GIF
	XC_SHAPE_TABLE    XC_OBJECT_TYPE = 69 // 形状对象-表格

	XC_MENU          XC_OBJECT_TYPE = 81       // 弹出菜单
	XC_IMAGE         XC_OBJECT_TYPE = 82       // 图片
	XC_IMAGE_TEXTURE                = XC_IMAGE // 图片纹理, 图片源, 图片素材
	XC_HDRAW         XC_OBJECT_TYPE = 83       // 绘图操作
	XC_FONT          XC_OBJECT_TYPE = 84       // 炫彩字体
	XC_FLASH         XC_OBJECT_TYPE = 85       // flash
	XC_PANE_CELL     XC_OBJECT_TYPE = 86       // ...
	XC_WEB           XC_OBJECT_TYPE = 87       // ...
	XC_IMAGE_FRAME   XC_OBJECT_TYPE = 88       // 图片帧, 指定图片的渲染属性
	XC_SVG           XC_OBJECT_TYPE = 89       // SVG矢量图形

	XC_LAYOUT_OBJECT    XC_OBJECT_TYPE = 101 // 布局对象LayoutObject
	XC_ADAPTER          XC_OBJECT_TYPE = 102 // ...
	XC_ADAPTER_TABLE    XC_OBJECT_TYPE = 103 // 数据适配器AdapterTable
	XC_ADAPTER_TREE     XC_OBJECT_TYPE = 104 // 数据适配器AdapterTree
	XC_ADAPTER_LISTVIEW XC_OBJECT_TYPE = 105 // 数据适配器AdapterListView
	XC_ADAPTER_MAP      XC_OBJECT_TYPE = 106 // 数据适配器AdapterMap

	XC_BKINFOM XC_OBJECT_TYPE = 116 // 背景管理器

	// 无实体对象, 只是用来判断布局

	XC_LAYOUT_LISTVIEW     XC_OBJECT_TYPE = 111
	XC_LAYOUT_LIST         XC_OBJECT_TYPE = 112
	XC_LAYOUT_OBJECT_GROUP XC_OBJECT_TYPE = 113
	XC_LAYOUT_OBJECT_ITEM  XC_OBJECT_TYPE = 114
	XC_LAYOUT_PANEL        XC_OBJECT_TYPE = 115

	// 无实体对象, 只是用来判断类型

	XC_LIST_ITEM      XC_OBJECT_TYPE = 121 // 列表项模板 list_item
	XC_LISTVIEW_GROUP XC_OBJECT_TYPE = 122
	XC_LISTVIEW_ITEM  XC_OBJECT_TYPE = 123
	XC_LAYOUT_BOX     XC_OBJECT_TYPE = 124

	XC_ANIMATION_SEQUENCE XC_OBJECT_TYPE = 131 // 动画序列
	XC_ANIMATION_GROUP    XC_OBJECT_TYPE = 132 // 动画同步组
	XC_ANIMATION_ITEM     XC_OBJECT_TYPE = 133 // 动画项
)

// XC_OBJECT_TYPE_EX 对象扩展类型(功能扩展).
type XC_OBJECT_TYPE_EX int

const (
	Button_Type_Default XC_OBJECT_TYPE_EX = iota // 默认类型
	Button_Type_Radio                            // 单选按钮
	Button_Type_Check                            // 复选按钮
	Button_Type_Close                            // 窗口关闭按钮
	Button_Type_Min                              // 窗口最小化按钮
	Button_Type_Max                              // 窗口最大化还原按钮

	Element_Type_Layout // 布局元素, 启用布局功能的元素

	Xc_Ex_Error XC_OBJECT_TYPE_EX = -1 // 错误类型
)

// Element_Position_ UI元素位置.
type Element_Position_ int

const (
	Element_Position_No     Element_Position_ = 0x00 // 无效
	Element_Position_Left   Element_Position_ = 0x01 // 左边
	Element_Position_Top    Element_Position_ = 0x02 // 上边
	Element_Position_Right  Element_Position_ = 0x04 // 右边
	Element_Position_Bottom Element_Position_ = 0x08 // 下边
)

// Position_Flag_ 位置标识.
type Position_Flag_ int

const (
	Position_Flag_Left        Position_Flag_ = iota // 左
	Position_Flag_Top                               // 上
	Position_Flag_Right                             // 右
	Position_Flag_Bottom                            // 下
	Position_Flag_LeftTop                           // 左上角
	Position_Flag_LeftBottom                        // 左下角
	Position_Flag_RightTop                          // 右上角
	Position_Flag_RightBottom                       // 右下角
	Position_Flag_Center                            // 中心
)

// Pane_Align_ 窗格_对齐.
type Pane_Align_ int

const (
	Pane_Align_Left   Pane_Align_ = iota // 左侧
	Pane_Align_Top                       // 顶部
	Pane_Align_Right                     // 右侧
	Pane_Align_Bottom                    // 底部
	Pane_Align_Center                    // 居中
	Pane_Align_Error  = -1               // 错误
)

// Menu_Item_Flag_ 弹出菜单项标识.
type Menu_Item_Flag_ int32

const (
	Menu_Item_Flag_Normal    Menu_Item_Flag_ = 0x00 // 正常
	Menu_Item_Flag_Select    Menu_Item_Flag_ = 0x01 // 选择或鼠标停留
	Menu_Item_Flag_Stay      Menu_Item_Flag_ = 0x01 // 选择或鼠标停留
	Menu_Item_Flag_Check     Menu_Item_Flag_ = 0x02 // 勾选
	Menu_Item_Flag_Popup     Menu_Item_Flag_ = 0x04 // 弹出
	Menu_Item_Flag_Separator Menu_Item_Flag_ = 0x08 // 分隔栏 ID号任意, ID号被忽略
	Menu_Item_Flag_Disable   Menu_Item_Flag_ = 0x10 // 禁用
)

// Menu_Popup_Position_ 弹出菜单方向.
type Menu_Popup_Position_ int

const (
	Menu_Popup_Position_Left_Top      Menu_Popup_Position_ = iota // 左上角
	Menu_Popup_Position_Left_Bottom                               // 左下角
	Menu_Popup_Position_Right_Top                                 // 右上角
	Menu_Popup_Position_Right_Bottom                              // 右下角
	Menu_Popup_Position_Center_Left                               // 左居中
	Menu_Popup_Position_Center_Top                                // 上居中
	Menu_Popup_Position_Center_Right                              // 右居中
	Menu_Popup_Position_Center_Bottom                             // 下居中
)

// ComboBox_State_ ComboBox状态.
type ComboBox_State_ int

const (
	ComboBox_State_Leave ComboBox_State_ = iota // 鼠标离开状态
	ComboBox_State_Stay                         // 鼠标停留状态
	ComboBox_State_Down                         // 按下状态
)

// Adapter_Date_Type_ 数据适配器数据类型.
type Adapter_Date_Type_ int

const (
	Adapter_Date_Type_Int    Adapter_Date_Type_ = iota // 整形
	Adapter_Date_Type_Float                            // 浮点型
	Adapter_Date_Type_String                           // 字符串
	Adapter_Date_Type_Image                            // 图片
	Adapter_Date_Type_Error  Adapter_Date_Type_ = -1
)

// Chat_Flag_ Edit 聊天气泡对齐方式.
type Chat_Flag_ int

const (
	Chat_Flag_Left            Chat_Flag_ = 0x1 // 左侧
	Chat_Flag_Right           Chat_Flag_ = 0x2 // 右侧
	Chat_Flag_Center          Chat_Flag_ = 0x4 // 中间
	Chat_Flag_Next_Row_Bubble Chat_Flag_ = 0x8 // 下一行显示气泡
)

// Edit_Type_ 编辑框类型.
type Edit_Type_ int

const (
	Edit_Type_None      Edit_Type_ = iota // 普通编辑框, 每行的高度相同
	Edit_Type_Editor                      // 代码编辑
	Edit_Type_Richedit                    // 富文本编辑框, 每行的高度可能不同
	Edit_Type_Chat                        // 聊天气泡, 每行的高度可能不同
	Edit_Type_CodeTable                   // 代码表格, 内部使用, 每行的高度相同
)

// Edit_Style_Type_ 编辑框风格类型.
type Edit_Style_Type_ int32

const (
	Edit_Style_Type_Font_Color Edit_Style_Type_ = iota + 1 // 字体
	Edit_Style_Type_Image                                  // 图片
	Edit_Style_Type_Obj                                    // UI对象
)

// Edit_TextAlign_Flag_ 编辑框文本对齐标志.
type Edit_TextAlign_Flag_ int

const (
	Edit_TextAlign_Flag_Left   Edit_TextAlign_Flag_ = 0x0 // 左侧
	Edit_TextAlign_Flag_Right  Edit_TextAlign_Flag_ = 0x1 // 右侧
	Edit_TextAlign_Flag_Center Edit_TextAlign_Flag_ = 0x2 // 水平居中

	Edit_TextAlign_Flag_Top      Edit_TextAlign_Flag_ = 0x0 // 顶部
	Edit_TextAlign_Flag_Bottom   Edit_TextAlign_Flag_ = 0x4 // 底部
	Edit_TextAlign_Flag_Center_V Edit_TextAlign_Flag_ = 0x8 // 垂直居中
)

// Table_Flag_ Table 标识.
type Table_Flag_ int

const (
	Table_Flag_Full Table_Flag_ = iota // 铺满组合单元格
	Table_Flag_None                    // 正常最小单元格
)

// GRADIENT_FILL_ 渐变填充模式.
type GRADIENT_FILL_ int

const (
	GRADIENT_FILL_RECT_H   GRADIENT_FILL_ = iota // 水平填充
	GRADIENT_FILL_RECT_V                         // 垂直填充
	GRADIENT_FILL_TRIANGLE                       // 三角形
)

// Ease_Type_ 缓动类型.
type Ease_Type_ int

const (
	Ease_Type_In    Ease_Type_ = iota // 从慢到快
	Ease_Type_Out                     // 从快到慢
	Ease_Type_InOut                   // 从慢到快再到慢
)

// Ease_Flag_ 缓动标识.
type Ease_Flag_ int

const (
	Ease_Flag_Linear Ease_Flag_ = iota // 线性, 直线
	Ease_Flag_Quad                     // 二次方曲线
	Ease_Flag_Cubic                    // 三次方曲线, 圆弧
	Ease_Flag_Quart                    // 四次方曲线
	Ease_Flag_Quint                    // 五次方曲线

	Ease_Flag_Sine    // 正弦, 在末端变化
	Ease_Flag_Expo    // 突击, 突然一下
	Ease_Flag_Circ    // 圆环, 好比绕过一个圆环
	Ease_Flag_Elastic // 强力回弹
	Ease_Flag_Back    // 回弹, 比较缓慢
	Ease_Flag_Bounce  // 弹跳, 模拟小球落地弹跳

	Ease_Flag_In    Ease_Flag_ = 0x010000 // 从慢到快
	Ease_Flag_Out   Ease_Flag_ = 0x020000 // 从快到慢
	Ease_Flag_InOut Ease_Flag_ = 0x030000 // 从慢到快再到慢
)

// FontStyle_ 字体样式.
type FontStyle_ int32

const (
	FontStyle_Regular    FontStyle_ = 0 // 正常
	FontStyle_Bold       FontStyle_ = 1 // 粗体
	FontStyle_Italic     FontStyle_ = 2 // 斜体
	FontStyle_BoldItalic FontStyle_ = 3 // 粗斜体
	FontStyle_Underline  FontStyle_ = 4 // 下划线
	FontStyle_Strikeout  FontStyle_ = 8 // 删除线
)

// Image_Draw_Type_ 图片绘制类型.
type Image_Draw_Type_ int

const (
	Image_Draw_Type_Default         Image_Draw_Type_ = iota // 默认
	Image_Draw_Type_Stretch                                 // 拉伸
	Image_Draw_Type_Adaptive                                // 自适应,九宫格
	Image_Draw_Type_Tile                                    // 平铺
	Image_Draw_Type_Fixed_Ratio                             // 固定比例,当图片超出显示范围时,按照原始比例压缩显示图片
	Image_Draw_Type_Adaptive_Border                         // 九宫格不绘制中间区域
)

// ListItemTemp_Type_ 列表项模板类型.
type ListItemTemp_Type_ int

const (
	ListItemTemp_Type_Tree           ListItemTemp_Type_ = 0x01                                                               // tree
	ListItemTemp_Type_ListBox        ListItemTemp_Type_ = 0x02                                                               // listBox
	ListItemTemp_Type_List_Head      ListItemTemp_Type_ = 0x04                                                               // list 列表头
	ListItemTemp_Type_List_Item      ListItemTemp_Type_ = 0x08                                                               // list 列表项
	ListItemTemp_Type_ListView_Group ListItemTemp_Type_ = 0x10                                                               // listView 列表视组
	ListItemTemp_Type_ListView_Item  ListItemTemp_Type_ = 0x20                                                               // listView 列表视项
	ListItemTemp_Type_List                              = ListItemTemp_Type_List_Head | ListItemTemp_Type_List_Item          // list (列表头)与(列表项)组合
	ListItemTemp_Type_ListView                          = ListItemTemp_Type_ListView_Group | ListItemTemp_Type_ListView_Item // listView (列表视组)与(列表视项)组合
)

// Window_Position_ 窗口位置.
type Window_Position_ int

const (
	Window_Position_Top    Window_Position_ = iota // 上
	Window_Position_Bottom                         // 下
	Window_Position_Left                           // 左
	Window_Position_Right                          // 右
	Window_Position_Body                           // body
	Window_Position_Window                         // window 整个窗口
	Window_Position_Error  Window_Position_ = -1   // 错误
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

// List_Item_State_ List项状态.
type List_Item_State_ int32

const (
	List_Item_State_Leave  List_Item_State_ = iota // 项鼠标离开状态
	List_Item_State_Stay                           // 项鼠标停留状态
	List_Item_State_Select                         // 项选择状态
	List_Item_State_Cache                          // 缓存的项
)

// Tree_Item_State_ Tree项状态.
type Tree_Item_State_ int32

const (
	Tree_Item_State_Leave  Tree_Item_State_ = iota // 项鼠标离开状态
	Tree_Item_State_Stay                           // 项鼠标停留状态
	Tree_Item_State_Select                         // 项选择状态
)

// List_DrawItemBk_Flag_ 项背景绘制标志位(List, ListBox, ListView, Tree).
type List_DrawItemBk_Flag_ int

const (
	List_DrawItemBk_Flag_Leave       List_DrawItemBk_Flag_ = 1 << iota // 绘制鼠标离开状态项背景
	List_DrawItemBk_Flag_Stay                                          // 绘制鼠标停留状态项项背景
	List_DrawItemBk_Flag_Select                                        // 绘制鼠标选择状态项背景
	List_DrawItemBk_Flag_Group_Leave                                   // 绘制鼠标离开状态组背景, 当项为组时
	List_DrawItemBk_Flag_Group_Stay                                    // 绘制鼠标停留状态组背景, 当项为组时

	List_DrawItemBk_Flag_Line  // 列表绘制水平分割线
	List_DrawItemBk_Flag_LineV // 列表绘制垂直分割线

	List_DrawItemBk_Flag_Nothing List_DrawItemBk_Flag_ = 0 // 不绘制
)

// PropertyGrid_Item_Type_ 属性网格项类型.
type PropertyGrid_Item_Type_ int

const (
	PropertyGrid_Item_Type_Text       PropertyGrid_Item_Type_ = iota // 默认,字符串类型
	PropertyGrid_Item_Type_Edit                                      // 编辑框
	PropertyGrid_Item_Type_Edit_Color                                // 颜色选择元素
	PropertyGrid_Item_Type_Edit_File                                 // 文件选择编辑框
	PropertyGrid_Item_Type_Edit_Set                                  // 设置
	PropertyGrid_Item_Type_ComboBox                                  // 组合框
	PropertyGrid_Item_Type_Group                                     // 组
	PropertyGrid_Item_Type_Panel                                     // 面板
)

// Zorder_ Z序位置.
type Zorder_ int

const (
	Zorder_Top    Zorder_ = iota // 最上面
	Zorder_Bottom                // 最下面
	Zorder_Before                // 指定目标下面
	Zorder_After                 // 指定目标上面
)

// Pane_State_ 窗格状态.
type Pane_State_ int

const (
	Pane_State_Any   Pane_State_ = iota
	Pane_State_Lock              // 锁定
	Pane_State_Dock              // 停靠码头
	Pane_State_Float             // 浮动窗格
	Pane_State_Error Pane_State_ = -1
)

// XC_DWRITE_RENDERING_MODE_ D2D文本渲染模式.
type XC_DWRITE_RENDERING_MODE_ int

const (
	XC_DWRITE_RENDERING_MODE_DEFAULT                     XC_DWRITE_RENDERING_MODE_ = iota // 指定根据字体和大小自动确定呈现模式。
	XC_DWRITE_RENDERING_MODE_ALIASED                                                      // 指定不执行抗锯齿。 每个像素要么设置为文本的前景色，要么保留背景的颜色。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_CLASSIC                                        // 使用与别名文本相同的度量指定 ClearType 呈现。 字形只能定位在整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_NATURAL                                        // 使用使用 CLEARTYPE_NATURAL_QUALITY 创建的字体，使用与使用 GDI 的文本呈现相同的指标指定 ClearType 呈现。 与使用别名文本相比，字形度量更接近其理想值，但字形仍然位于整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL                                            // 仅在水平维度中指定具有抗锯齿功能的 ClearType 渲染。 这通常用于中小字体大小（最多 16 ppem）。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL_SYMMETRIC                                  // 指定在水平和垂直维度上具有抗锯齿的 ClearType 渲染。这通常用于较大的尺寸，以使曲线和对角线看起来更平滑，但会牺牲一些柔和度。
	XC_DWRITE_RENDERING_MODE_OUTLINE                                                      // 指定渲染应绕过光栅化器并直接使用轮廓。 这通常用于非常大的尺寸。
)

// MonthCal_Button_Type_ 月历元素上的按钮类型.
type MonthCal_Button_Type_ int

const (
	MonthCal_Button_Type_Today      MonthCal_Button_Type_ = iota // 今天按钮
	MonthCal_Button_Type_Last_Year                               // 上一年
	MonthCal_Button_Type_Next_Year                               // 下一年
	MonthCal_Button_Type_Last_Month                              // 上一月
	MonthCal_Button_Type_Next_Month                              // 下一月
)

// 窗格菜单 当前未使用

const (
	IDM_LOCK  = 1000000006 // 锁定
	IDM_DOCK  = 1000000007 // 停靠
	IDM_FLOAT = 1000000008 // 浮动
	IDM_HIDE  = 1000000009 // 隐藏

	Edit_Style_No      = 0xFFFF // 无效样式
	Edit_Style_Default = 1      // edit 默认样式

	Text_Buffer_Size = 10240 // 共享文本缓冲区大小
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

// NotifyMsg_Skin_ 通知消息外观.
type NotifyMsg_Skin_ int

const (
	NotifyMsg_Skin_No      NotifyMsg_Skin_ = iota // 默认
	NotifyMsg_Skin_Success                        // 成功
	NotifyMsg_Skin_Warning                        // 警告
	NotifyMsg_Skin_Message                        // 消息
	NotifyMsg_Skin_Error                          // 错误
)

// Animation_Move_ 动画移动标识.
type Animation_Move_ int

const (
	Animation_Move_X Animation_Move_ = 0x01 // X轴移动.
	Animation_Move_Y Animation_Move_ = 0x02 // Y轴移动.
)

// BkObject_Align_Flag_ 背景对象对齐方式.
type BkObject_Align_Flag_ int

const (
	BkObject_Align_Flag_Left     BkObject_Align_Flag_ = 1 << iota // 左对齐, 当设置此标识时, 外间距(margin.left)代表左侧间距; 当right未设置时,那么外间距(margin.right)代表宽度;
	BkObject_Align_Flag_Top                                       // 顶对齐, 当设置此标识时, 外间距(margin.top)代表顶部间距; 当bottom未设置时,那么外间距(margin.bottom)代表高度;
	BkObject_Align_Flag_Right                                     // 右对齐, 当设置此标识时, 外间距(margin.right)代表右侧间距; 当left未设置时,那么外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Bottom                                    // 底对齐, 当设置此标识时, 外间距(margin.bottom)代表底部间距; 当top未设置时,那么外间距(margin.top)代表高度;
	BkObject_Align_Flag_Center                                    // 水平居中, 当设置此标识时, 外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Center_v                                  // 垂直居中, 当设置此标识时, 外间距(margin.top)代表高度;
	BkObject_Align_Flag_No       BkObject_Align_Flag_ = 0         // 无
)

// FrameWnd_Cell_Type_ 框架窗口单元格类型
type FrameWnd_Cell_Type_ int

const (
	FrameWnd_Cell_Type_No         FrameWnd_Cell_Type_ = 0 // 无
	FrameWnd_Cell_Type_Pane       FrameWnd_Cell_Type_ = 1 // 窗格
	FrameWnd_Cell_Type_Group      FrameWnd_Cell_Type_ = 2 // 窗格组
	FrameWnd_Cell_Type_BodyView   FrameWnd_Cell_Type_ = 3 // 主视图区
	FrameWnd_Cell_Type_Top_Bottom FrameWnd_Cell_Type_ = 4 // 上下布局
	FrameWnd_Cell_Type_Left_Right FrameWnd_Cell_Type_ = 5 // 左右布局
)

// VK_ 键盘按键常量
// type VK_ uint32

const (
	VK_Lbutton = 1 // 鼠标左键
	VK_RButton = 2 // 鼠标右键
	VK_MButton = 4 // 鼠标中键

	VK_A = 65
	VK_B = 66
	VK_C = 67
	VK_D = 68
	VK_E = 69
	VK_F = 70
	VK_G = 71
	VK_H = 72
	VK_I = 73
	VK_J = 74
	VK_K = 75
	VK_L = 76
	VK_M = 77
	VK_N = 78
	VK_O = 79
	VK_P = 80
	VK_Q = 81
	VK_R = 82
	VK_S = 83
	VK_T = 84
	VK_U = 85
	VK_V = 86
	VK_W = 87
	VK_X = 88
	VK_Y = 89
	VK_Z = 90

	VK_F1  = 112
	VK_F2  = 113
	VK_F3  = 114
	VK_F4  = 115
	VK_F5  = 116
	VK_F6  = 117
	VK_F7  = 118
	VK_F8  = 119
	VK_F9  = 120
	VK_F10 = 121
	VK_F11 = 122
	VK_F12 = 123
	VK_F13 = 124
	VK_F14 = 125
	VK_F15 = 126
	VK_F16 = 127

	VK_0 = 48
	VK_1 = 49
	VK_2 = 50
	VK_3 = 51
	VK_4 = 52
	VK_5 = 53
	VK_6 = 54
	VK_7 = 55
	VK_8 = 56
	VK_9 = 57

	VK_Numpad0   = 96
	VK_Numpad1   = 97
	VK_Numpad2   = 98
	VK_Numpad3   = 99
	VK_Numpad4   = 100
	VK_Numpad5   = 101
	VK_Numpad6   = 102
	VK_Numpad7   = 103
	VK_Numpad8   = 104
	VK_Numpad9   = 105
	VK_Multiply  = 106 // 数字键盘上的*键
	VK_Add       = 107 // 数字键盘上的+键
	VK_Separator = 108
	VK_Subtract  = 109 // 数字键盘上的-键
	VK_Decimal   = 110 // 数字键盘上的.键
	VK_Divide    = 111 // 数字键盘上的/键

	VK_Break       = 3
	VK_Backspace   = 8
	VK_Tab         = 9
	VK_Enter       = 13
	VK_Shift       = 16
	VK_Ctrl        = 17
	VK_Alt         = 18
	VK_Pause       = 19
	VK_CapsLock    = 20
	VK_NumLock     = 144
	VK_ScrollLock  = 145
	VK_Esc         = 27
	VK_Space       = 32
	VK_PageUp      = 33
	VK_PageDown    = 34
	VK_Home        = 36
	VK_End         = 35
	VK_Left        = 37
	VK_Up          = 38
	VK_Right       = 39
	VK_Down        = 40
	VK_Insert      = 45
	VK_Delete      = 46
	VK_LWin        = 91 // 左win键
	VK_Rwin        = 92 // 右win键
	VK_ContextMenu = 93 // 右Ctrl左边键，点击相当于点击鼠标右键，会弹出快捷菜单
	VK_PrintScreen = 44

	VK_Semicolon     = 186 // ;(分号)
	VK_Equals        = 187 //  = 键
	VK_Comma         = 188 // ,键(逗号)
	VK_Minus         = 189 // -键(减号)
	VK_FullStop      = 190 // .键(句号)
	VK_ForwardSlash  = 191 // /键
	VK_Backquote     = 192 // `键(Esc下面)反撇号,反引号
	VK_BracketLeft   = 219 // [键
	VK_Backslash     = 220 // \键
	VK_BracketRight  = 221 // ]键
	VK_QuotationMark = 222 // ‘键(引号)
)

// SWP_ 是窗口大小和定位的标志.
type SWP_ uint32

const (
	SWP_ASYNCWINDOWPOS SWP_ = 0x4000 // 如果调用线程和拥有窗口的线程连接到不同的输入队列，系统会将请求发布到拥有窗口的线程。这可以防止调用线程在其他线程处理请求时阻塞其执行。
	SWP_DEFERERASE     SWP_ = 0x2000 // 防止生成WM_SYNCPAINT消息。
	SWP_DRAWFRAME      SWP_ = 0x0020 // 在窗口周围绘制一个框架（在窗口的类描述中定义）。
	SWP_FRAMECHANGED   SWP_ = 0x0020 // 应用使用 SetWindowLong 函数 设置的新框架样式。向窗口发送WM_NCCALCSIZE消息，即使窗口大小没有改变。如果未指定此标志，则仅在更改窗口大小时发送 WM_NCCALCSIZE 。
	SWP_HIDEWINDOW     SWP_ = 0x0080 // 隐藏窗口。
	SWP_NOACTIVATE     SWP_ = 0x0010 // 不激活窗口。如果未设置此标志，则窗口被激活并移动到最顶层或非最顶层组的顶部（取决于hWndInsertAfter参数的设置）。
	SWP_NOCOPYBITS     SWP_ = 0x0100 // 丢弃客户区的全部内容。如果未指定此标志，则在调整窗口大小或重新定位后，将保存客户区的有效内容并将其复制回客户区。
	SWP_NOMOVE         SWP_ = 0x0002 // 保留当前位置（忽略X和Y参数）。
	SWP_NOOWNERZORDER  SWP_ = 0x0200 // 不改变所有者窗口在 Z 顺序中的位置。
	SWP_NOREDRAW       SWP_ = 0x0008 // 不重绘更改。如果设置了此标志，则不会发生任何类型的重新绘制。这适用于客户区、非客户区（包括标题栏和滚动条）以及由于窗口移动而未覆盖的父窗口的任何部分。设置此标志时，应用程序必须显式地使需要重绘的窗口和父窗口的任何部分无效或重绘。
	SWP_NOREPOSITION   SWP_ = 0x0200 // 与SWP_NOOWNERZORDER标志相同。
	SWP_NOSENDCHANGING SWP_ = 0x0400 // 阻止窗口接收WM_WINDOWPOSCHANGING消息。
	SWP_NOSIZE         SWP_ = 0x0001 // 保留当前大小（忽略cx和cy参数）。
	SWP_NOZORDER       SWP_ = 0x0004 // 保留当前 Z 顺序（忽略hWndInsertAfter参数）。
	SWP_SHOWWINDOW     SWP_ = 0x0040 // 显示窗口。
)

type HWND_ int

const (
	HWND_NOTOPMOST HWND_ = -2 // 将窗口置于所有非顶层窗口之上（即在所有顶层窗口之后）。如果窗口已经是非顶层窗口则该标志不起作用。
	HWND_TOPMOST   HWND_ = -1 // 将窗口置于所有非顶层窗口之上。即使窗口未被激活, 窗口也将保持顶级位置。
	HWND_TOP       HWND_ = 0  // 将窗口置于Z序的顶部。
	HWND_BOTTOM    HWND_ = 1  // 将窗口置于Z序的底部。如果参数hWnd标识了一个顶层窗口，则窗口失去顶级位置，并且被置在所有其他窗口的底部。
)

// TrayIcon_Flag_ 托盘图标标识.
type TrayIcon_Flag_ int32

const (
	TrayIcon_Flag_Icon_None    TrayIcon_Flag_ = 0x0  // 无图标 NIIF_NONE
	TrayIcon_Flag_Icon_Info    TrayIcon_Flag_ = 0x1  // 信息图标 NIIF_INFO
	TrayIcon_Flag_Icon_Warning TrayIcon_Flag_ = 0x2  // 警告图标 NIIF_WARNING
	TrayIcon_Flag_Icon_Error   TrayIcon_Flag_ = 0x3  // 错误图标 NIIF_ERROR
	TrayIcon_Flag_Icon_User    TrayIcon_Flag_ = 0x4  // 用户指定的图标 NIIF_USER
	TrayIcon_Flag_Nosound      TrayIcon_Flag_ = 0x10 // 禁止播放气泡声音 NIIF_NOSOUND
)

// ABGR颜色

const (
	COLOR_BLUE    = 0xFFFF0000 // 蓝色
	COLOR_GREEN   = 0xFF00FF00 // 绿色
	COLOR_RED     = 0xFF0000FF // 红色
	COLOR_CYAN    = 0xFFFFFF00 // 青色
	COLOR_MAGENTA = 0xFFFF00FF // 洋红色
	COLOR_YELLOW  = 0xFF00FFFF // 黄色

	COLOR_BLUE_LIGHT    = 0xFFFF8080 // 浅蓝色
	COLOR_GREEN_LIGHT   = 0xFF80FF80 // 浅绿色
	COLOR_RED_LIGHT     = 0xFF8080FF // 浅红色
	COLOR_CYAN_LIGHT    = 0xFFFFFF80 // 浅青色
	COLOR_MAGENTA_LIGHT = 0xFFFF80FF // 浅洋红色
	COLOR_YELLOW_LIGHT  = 0xFF80FFFF // 浅黄色

	COLOR_BLUE_DARK    = 0xFF800000 // 深蓝色
	COLOR_GREEN_DARK   = 0xFF008000 // 深绿色
	COLOR_RED_DARK     = 0xFF000080 // 深红色
	COLOR_CYAN_DARK    = 0xFF808000 // 深青色
	COLOR_MAGENTA_DARK = 0xFF800080 // 深洋红色
	COLOR_YELLOW_DARK  = 0xFF008080 // 深黄色

	COLOR_WHITE       = 0xFFFFFFFF // 白色
	COLOR_LIGHTGRAY   = 0xFFD3D3D3 // 浅灰色
	COLOR_GRAY        = 0xFF808080 // 灰色
	COLOR_DARKGRAY    = 0xFF404040 // 深灰色
	COLOR_BLACK       = 0xFF000000 // 黑色
	COLOR_BROWN       = 0xFF2A2AA5 // 棕色
	COLOR_ORANGE      = 0xFF00A5FF // 橙色
	COLOR_TRANSPARENT = 0xFF000000 // 透明

	COLOR_GRAY_3F = 0xFF3F3F3F // 灰色值
	COLOR_GRAY_50 = 0xFF505050 // 灰色值
	COLOR_GRAY_55 = 0xFF555555 // 灰色值
	COLOR_GRAY_60 = 0xFF606060 // 灰色值
	COLOR_GRAY_7C = 0xFF7C7C7C // 灰色值
	COLOR_GRAY_9A = 0xFF9A9A9A // 灰色值
	COLOR_GRAY_AA = 0xFFAAAAAA // 灰色值
	COLOR_GRAY_C0 = 0xFFC0C0C0 // 灰色值
	COLOR_GRAY_C8 = 0xFFC8C8C8 // 灰色值
	COLOR_GRAY_D0 = 0xFFD0D0D0 // 灰色值
	COLOR_GRAY_E7 = 0xFFE7E7E7 // 灰色值

	COLOR_BLUE_98 = 0xFF980000 // 蓝色值
)
