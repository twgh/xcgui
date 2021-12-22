package xc

// 菜单-弹出窗口信息
type Menu_PopupWnd_ struct {
	HWindow   int // 窗口句柄
	NParentID int // 父项ID
}

// 菜单背景自绘结构
type Menu_DrawBackground_ struct {
	HMenu     int // 菜单句柄
	HWindow   int // 当前弹出菜单项的窗口句柄
	NParentID int // 父项ID
}

// 菜单项自绘结构
type Menu_DrawItem_ struct {
	HMenu   int     // 菜单句柄
	HWindow int     // 当前弹出菜单项的窗口句柄
	NID     int     // ID
	NState  int     // 状态: Menu_Item_Flag_
	RcItem  RECT    // 坐标
	HIcon   int     // 菜单项图标句柄
	PText   uintptr // 文本, 使用xc.UintPtrToString()函数转换到string
}

// 列表框项信息
type ListBox_Item_ struct {
	Index      int  // 项索引
	NUserData  int  // 用户绑定数据
	NHeight    int  // 项默认高度
	NSelHeight int  // 项选中时高度
	NState     int  // 状态: List_Item_State_
	RcItem     RECT // 项坐标
	HLayout    int  // 布局对象句柄
	HTemp      int  // 列表项模板句柄
}

// List 列表头项信息
type List_Header_Item_ struct {
	Index          int  // 项索引
	NUserData      int  // 用户绑定数据
	BSort          bool // 是否支持排序
	NSortType      int  // 排序方式,0无效,1升序,2降序
	IColumnAdapter int  // 对应数据适配器中的列索引
	NState         int  // 状态: Common_State3_
	RcItem         RECT // 项坐标
	HLayout        int  // 布局对象句柄
	HTemp          int  // 列表项模板句柄
}

// List 列表项信息
type List_Item_ struct {
	Index     int  // 项索引
	ISubItem  int  // 子项索引
	NUserData int  // 用户数据
	NState    int  // 状态: List_Item_State_
	RcItem    RECT // 未使用
	HLayout   int  // 布局对象句柄
	HTemp     int  // 列表项模板句柄
}

// Tree 树项信息
type Tree_Item_ struct {
	NID        int  // 项ID
	NDepth     int  // 深度
	NHeight    int  // 项高度
	NSelHeight int  // 项选中状态高度
	NUserData  int  // 用户数据
	BExpand    bool // 展开
	NState     int  // 状态:　Tree_Item_State_
	RcItem     RECT // 坐标
	HLayout    int  // 布局对象句柄
	HTemp      int  // 列表项模板句柄
}

// 树UI元素拖动项
type Tree_Drag_Item_ struct {
	NDragItem int // 拖动项ID
	NDestItem int // 目标项ID
	NType     int // 停放相对目标位置, 0:(上)停放到目标的上面, 1:(下)停放到目标的下面, 3:(中)停放到目标的的子项,
}

// ListView 列表视项信息
type ListView_Item_ struct {
	IGroup    int  // 项所述组索引 -1没有组
	IItem     int  // 项在数组中位置索引,如果此致为-1,那么为组
	NUserData int  // 用户绑定数据
	NState    int  // 状态: List_Item_State_
	RcItem    RECT // 整个区域,包含边框
	HLayout   int  // 布局对象
	HTemp     int  // 列表项模板
}

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type RECTF struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type POINT struct {
	X int32
	Y int32
}

type POINTF struct {
	X float32
	Y float32
}

type SIZE struct {
	CX int32
	CY int32
}

// 字体属性
type LOGFONTW struct {
	LfHeight         int      // 高度
	LfWidth          int      // 宽度
	LfEscapement     int      // 指定角度
	LfOrientation    int      // 字符基线
	LfWeight         int      // 字体粗细
	LfItalic         uint8    // 是否斜体
	LfUnderline      uint8    // 是否下划线
	LfStrikeOut      uint8    // 是否删除线
	LfCharSet        uint8    // 字符集
	LfOutPrecision   uint8    // 输出精度
	LfClipPrecision  uint8    // 剪辑精度
	LfQuality        uint8    // 输出质量
	LfPitchAndFamily uint8    // 字符间距
	LfFaceName       [64]byte // 字体名称
}

// Editor 颜色信息
type Editor_Color_ struct {
	ClrMargin1         int32 // 背景色1, ABGR颜色
	ClrMargin2         int32 // 背景色2, ABGR颜色
	ClrMarginText      int32 // 文本色, ABGR颜色
	ClrBreakpoint      int32 // 断点色, ABGR颜色
	ClrBreakpointArrow int32 // 断点箭头色, ABGR颜色
	ClrRun             int32 // 当前执行位置指示色, ABGR颜色
	ClrCurRow          int32 // 突出显示当前行颜色, ABGR颜色
	ClrMatch           int32 // 设置匹配文本背景色, ABGR颜色
}

// Edit 数据复制
type Edit_Data_Copy_ struct {
	NCount      int32                  // 内容数量
	NStyleCount int32                  // 样式数量
	PStyle      *Edit_Data_Copy_Style_ // 样式数组
	PData       *int32                 // 内容数组
}

// Edit 样式信息
type Edit_Style_Info_ struct {
	Type            int32  // 样式类型: Edit_Style_Type_
	NRef            uint16 // 引用计数
	HFont_image_obj int32  // 字体,图片,UI对象
	Color           int32  // 颜色
	BColor          bool   // 是否使用颜色
}

// Edit 数据复制-样式
type Edit_Data_Copy_Style_ struct {
	HFont_image_obj int32 // 字体,图片,UI对象
	Color           int32 // 颜色
	BColor          bool  // 是否使用颜色
}

// 位置点
type Position_ struct {
	IRow    int32 // 行索引
	IColumn int32 // 列索引
}

// 字体信息
type Font_Info_ struct {
	NSize  int32      // 字体大小, 单位(pt,磅).
	NStyle int32      // 字体样式: FontStyle_
	Name   [32]uint16 // 字体名称, 使用xc.Font_Info_Name()函数转换为string.
}

// ListBox 列表框项信息2
type ListBox_Item_Info_ struct {
	NUserData  int32 // 用户绑定数据
	NHeight    int32 // 项高度, -1使用默认高度
	NSelHeight int32 // 项选中时高度, -1使用默认高度
}

// 月历元素项数据
type MonthCal_item_ struct {
	NDay   int  // 日期
	NType  int  // 1上月, 2当月, 3下月
	NState int  // 组合状态, MonthCal_State_Flag_
	RcItem RECT // 项坐标
}

// PGrid 属性网格项信息
type PropertyGrid_item_ struct {
	NType         int  // 类型
	NID           int  // 项ID
	NDepth        int  // 深度
	NUserData     int  // 用户数据
	NNameColWidth int  // 名称列宽度
	RcItem        RECT // 坐标
	RcExpand      RECT // 展开
	BExpand       bool // 是否展开
	BShow         bool // 是否可见
}

// 列表视项ID
type ListView_Item_Id_ struct {
	IGroup int // 组索引
	IItem  int // 项索引
}
