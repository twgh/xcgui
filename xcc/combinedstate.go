package xcc

// CombinedState 组合状态.
type CombinedState uint32

// Window_State_Flag_ 组合状态
// 窗口状态
type Window_State_Flag_ uint32

const (
	Window_State_Flag_Nothing      Window_State_Flag_ = 0x0000 // 无
	Window_State_Flag_Leave        Window_State_Flag_ = 0x0001 // 整个窗口
	Window_State_Flag_Body_Leave   Window_State_Flag_ = 0x0002 // 窗口-body
	Window_State_Flag_Top_Leave    Window_State_Flag_ = 0x0004 // 窗口-top
	Window_State_Flag_Bottom_Leave Window_State_Flag_ = 0x0008 // 窗口-bottom
	Window_State_Flag_Left_Leave   Window_State_Flag_ = 0x0010 // 窗口-left
	Window_State_Flag_Right_Leave  Window_State_Flag_ = 0x0020 // 窗口-right

	Window_State_Flag_Layout_Body Window_State_Flag_ = 0x20000000 // 布局内容区
)

// 组合状态
// 列表树状态
// type Tree_State_Flag_ CombinedState

const (
	Tree_State_Flag_Item_Leave CombinedState = 0x0080 // 项鼠标离开
	Tree_State_Flag_Item_Stay  CombinedState = 0x0100 // 项鼠标停留,保留值, 暂未使用

	Tree_State_Flag_Item_Select    CombinedState = 0x0200 // 项选择
	Tree_State_Flag_Item_Select_No CombinedState = 0x0400 // 项未选择

	Tree_State_Flag_Group    CombinedState = 0x0800 // 项为组
	Tree_State_Flag_Group_No CombinedState = 0x1000 // 项不为组
)

// 组合状态
// 元素状态标志
// type Element_State_Flag_ CombinedState

const (
	Element_State_Flag_Nothing                  = CombinedState(Window_State_Flag_Nothing) // 无
	Element_State_Flag_Enable     CombinedState = 0x0001                                   // 启用
	Element_State_Flag_Disable    CombinedState = 0x0002                                   // 禁用
	Element_State_Flag_Focus      CombinedState = 0x0004                                   // 焦点
	Element_State_Flag_Focus_No   CombinedState = 0x0008                                   // 无焦点
	Element_State_Flag_FocusEx    CombinedState = 0x40000000                               // 该元素或该元素的子元素拥有焦点
	Element_State_Flag_FocusEx_No CombinedState = 0x80000000                               // 无焦点Ex

	Layout_State_Flag_Layout_Body = CombinedState(Window_State_Flag_Layout_Body) // 布局内容区

	Element_State_Flag_Leave CombinedState = 0x0010 // 鼠标离开
	Element_State_Flag_Stay  CombinedState = 0x0020 // 为扩展模块保留
	Element_State_Flag_Down  CombinedState = 0x0040 // 为扩展模块保留
)

// 组合状态
// 按钮状态标志
// type Button_State_Flag_ CombinedState

const (
	Button_State_Flag_Leave = Element_State_Flag_Leave // 鼠标离开
	Button_State_Flag_Stay  = Element_State_Flag_Stay  // 鼠标停留
	Button_State_Flag_Down  = Element_State_Flag_Down  // 鼠标按下

	Button_State_Flag_Check    CombinedState = 0x0080 // 选中
	Button_State_Flag_Check_No CombinedState = 0x0100 // 未选中

	Button_State_Flag_WindowRestore  CombinedState = 0x0200 // 窗口还原
	Button_State_Flag_WindowMaximize CombinedState = 0x0400 // 窗口最大化
)

// 组合状态
// 组合框状态标志
// type ComboBox_State_Flag_ CombinedState

const (
	ComboBox_State_Flag_Leave = Element_State_Flag_Leave // 鼠标离开
	ComboBox_State_Flag_Stay  = Element_State_Flag_Stay  // 鼠标停留
	ComboBox_State_Flag_Down  = Element_State_Flag_Down  // 鼠标按下
)

// 组合状态
// 列表项状态标志
// type List_State_Flag_ CombinedState

const (
	List_State_Flag_Item_Leave CombinedState = 0x0080 // 项鼠标离开
	List_State_Flag_Item_Stay  CombinedState = 0x0100 // 项鼠标停留

	List_State_Flag_Item_Select    CombinedState = 0x0200 // 项选择
	List_State_Flag_Item_Select_No CombinedState = 0x0400 // 项未选择
)

// 组合状态
// 列表框项状态标志
// type ListBox_State_Flag_ CombinedState

const (
	ListBox_State_Flag_Item_Leave CombinedState = 0x0080 // 项鼠标离开
	ListBox_State_Flag_Item_Stay  CombinedState = 0x0100 // 项鼠标停留

	ListBox_State_Flag_Item_Select    CombinedState = 0x0200 // 项选择
	ListBox_State_Flag_Item_Select_No CombinedState = 0x0400 // 项未选择
)

// 组合状态
// 列表视图状态标志
// type ListView_State_Flag_ CombinedState

const (
	ListView_State_Flag_Item_Leave CombinedState = 0x0080 // 项鼠标离开
	ListView_State_Flag_Item_Stay  CombinedState = 0x0100 // 项鼠标停留

	ListView_State_Flag_Item_Select    CombinedState = 0x0200 // 项选择
	ListView_State_Flag_Item_Select_No CombinedState = 0x0400 // 项未选择

	ListView_State_Flag_Group_Leave CombinedState = 0x0800 // 组鼠标离开
	ListView_State_Flag_Group_Stay  CombinedState = 0x1000 // 组鼠标停留

	ListView_State_Flag_Group_Select    CombinedState = 0x2000 // 组选择
	ListView_State_Flag_Group_Select_No CombinedState = 0x4000 // 组未选择
)

// 组合状态
// 列表头状态
// type ListHeader_State_Flag_Item_ CombinedState

const (
	ListHeader_State_Flag_Item_Leave CombinedState = 0x0080 // 项鼠标离开
	ListHeader_State_Flag_Item_Stay  CombinedState = 0x0100 // 项鼠标停留
	ListHeader_State_Flag_Item_Down  CombinedState = 0x0200 // 项鼠标按下
)

// 组合状态
// 月历卡片状态标志
// type MonthCal_State_Flag_ CombinedState

const (
	MonthCal_State_Flag_Leave = Element_State_Flag_Leave // 离开状态

	MonthCal_State_Flag_Item_Leave CombinedState = 0x0080 // 项-离开
	MonthCal_State_Flag_Item_Stay  CombinedState = 0x0100 // 项-停留
	MonthCal_State_Flag_Item_Down  CombinedState = 0x0200 // 项-按下

	MonthCal_State_Flag_Item_Select    CombinedState = 0x0400 // 项-选择
	MonthCal_State_Flag_Item_Select_No CombinedState = 0x0800 // 项-未选择

	MonthCal_State_Flag_Item_Today      CombinedState = 0x1000  // 项-今天
	MonthCal_State_Flag_Item_Other      CombinedState = 0x2000  // 项-上月及下月
	MonthCal_State_Flag_Item_Last_Month CombinedState = 0x4000  // 项-上月
	MonthCal_State_Flag_Item_Cur_Month  CombinedState = 0x8000  // 项-当月
	MonthCal_State_Flag_Item_Next_Month CombinedState = 0x10000 // 项-下月
)

// 组合状态
// 布局状态标志
// type Layout_State_Flag_ CombinedState

const (
	Layout_State_Flag_Nothing               = CombinedState(Window_State_Flag_Nothing) // 无
	Layout_State_Flag_Full    CombinedState = 0x0001                                   // 完整背景
	Layout_State_Flag_Body    CombinedState = 0x0002                                   // 内容区域, 不包含边大小
)

// 组合状态
// 属性网格状态
// type PropertyGrid_State_Flag_ CombinedState

const (
	PropertyGrid_State_Flag_Item_Leave CombinedState = 0x0080 // 离开
	PropertyGrid_State_Flag_Item_Stay  CombinedState = 0x0100 // 停留

	PropertyGrid_State_Flag_Item_Select    CombinedState = 0x0200 // 选择
	PropertyGrid_State_Flag_Item_Select_No CombinedState = 0x0400 // 未选择

	PropertyGrid_State_Flag_Group_Leave     CombinedState = 0x0800 // 组离开
	PropertyGrid_State_Flag_Group_Expand    CombinedState = 0x1000 // 组展开
	PropertyGrid_State_Flag_Group_Expand_No CombinedState = 0x2000 // 组未展开
)

// 组合状态
// 窗格状态
// type Pane_State_Flag_ CombinedState

const (
	Pane_State_Flag_Leave = Element_State_Flag_Leave // 离开
	Pane_State_Flag_Stay  = Element_State_Flag_Stay  // 停留

	Pane_State_Flag_Caption CombinedState = 0x0080 // 标题
	Pane_State_Flag_Body    CombinedState = 0x0100 // 内容区
)
