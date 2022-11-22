package xc

import "github.com/twgh/xcgui/xcc"

// Menu_PopupWnd_ 菜单-弹出窗口信息
type Menu_PopupWnd_ struct {
	HWindow   int   // 窗口句柄
	NParentID int32 // 父项ID
}

// Menu_DrawBackground_ 菜单背景自绘结构
type Menu_DrawBackground_ struct {
	HMenu     int   // 菜单句柄
	HWindow   int   // 当前弹出菜单项的窗口句柄
	NParentID int32 // 父项ID
}

// Menu_DrawItem_ 菜单项自绘结构
type Menu_DrawItem_ struct {
	HMenu             int                 // 菜单句柄
	HWindow           int                 // 当前弹出菜单项的窗口句柄
	NID               int32               // ID
	NState            xcc.Menu_Item_Flag_ // 状态: Menu_Item_Flag_
	NShortcutKeyWidth int32               // 右侧快捷键占位宽度
	RcItem            RECT                // 坐标
	HIcon             int                 // 菜单项图标句柄
	PText             uintptr             // 文本, 使用xc.UintPtrToString()函数转换到string
}

// ListBox_Item_ 列表框项信息
type ListBox_Item_ struct {
	Index      int32                // 项索引
	NUserData  int                  // 用户绑定数据
	NHeight    int32                // 项默认高度
	NSelHeight int32                // 项选中时高度
	NState     xcc.List_Item_State_ // 状态: List_Item_State_
	RcItem     RECT                 // 项坐标
	HLayout    int                  // 布局对象句柄
	HTemp      int                  // 列表项模板句柄
}

// List 列表头项信息
type List_Header_Item_ struct {
	Index          int32              // 项索引
	NUserData      int                // 用户绑定数据
	BSort          bool               // 是否支持排序
	NSortType      int32              // 排序方式,0无效,1升序,2降序
	IColumnAdapter int32              // 对应数据适配器中的列索引
	NState         xcc.Common_State3_ // 状态: Common_State3_
	RcItem         RECT               // 项坐标
	HLayout        int                // 布局对象句柄
	HTemp          int                // 列表项模板句柄
}

// List 列表项信息
type List_Item_ struct {
	Index     int32                // 项索引
	ISubItem  int32                // 子项索引
	NUserData int                  // 用户数据
	NState    xcc.List_Item_State_ // 状态: List_Item_State_
	RcItem    RECT                 // 未使用
	HLayout   int                  // 布局对象句柄
	HTemp     int                  // 列表项模板句柄
}

// Tree 树项信息
type Tree_Item_ struct {
	NID        int32                // 项ID
	NDepth     int32                // 深度
	NHeight    int32                // 项高度
	NSelHeight int32                // 项选中状态高度
	NUserData  int                  // 用户数据
	BExpand    bool                 // 展开
	NState     xcc.Tree_Item_State_ // 状态:　Tree_Item_State_
	RcItem     RECT                 // 坐标
	HLayout    int                  // 布局对象句柄
	HTemp      int                  // 列表项模板句柄
}

// Tree_Drag_Item_ 树UI元素拖动项
type Tree_Drag_Item_ struct {
	NDragItem int32 // 拖动项ID
	NDestItem int32 // 目标项ID
	NType     int32 // 停放相对目标位置, 0:(上)停放到目标的上面, 1:(下)停放到目标的下面, 3:(中)停放到目标的的子项,
}

// ListView 列表视项信息
type ListView_Item_ struct {
	IGroup    int32                // 项所述组索引 -1没有组
	IItem     int32                // 项在数组中位置索引,如果此致为-1,那么为组
	NUserData int                  // 用户绑定数据
	NState    xcc.List_Item_State_ // 状态: List_Item_State_
	RcItem    RECT                 // 整个区域,包含边框
	HLayout   int                  // 布局对象
	HTemp     int                  // 列表项模板
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

// LOGFONTW 字体属性
type LOGFONTW struct {
	LfHeight         int32      // 高度
	LfWidth          int32      // 宽度
	LfEscapement     int32      // 指定角度
	LfOrientation    int32      // 字符基线
	LfWeight         int32      // 字体粗细
	LfItalic         byte       // 是否斜体
	LfUnderline      byte       // 是否下划线
	LfStrikeOut      byte       // 是否删除线
	LfCharSet        byte       // 字符集
	LfOutPrecision   byte       // 输出精度
	LfClipPrecision  byte       // 剪辑精度
	LfQuality        byte       // 输出质量
	LfPitchAndFamily byte       // 字符间距
	LfFaceName       [32]uint16 // 字体名称, 使用xc.Font_Info_Name()函数转换为string.
}

// Editor 颜色信息
type Editor_Color_ struct {
	ClrMargin1         int32 // 背景色1, ABGR 颜色
	ClrMargin2         int32 // 背景色2, ABGR 颜色
	ClrMarginText      int32 // 文本色, ABGR 颜色
	ClrBreakpoint      int32 // 断点色, ABGR 颜色
	ClrBreakpointArrow int32 // 断点箭头色, ABGR 颜色
	ClrRun             int32 // 当前执行位置指示色, ABGR 颜色
	ClrCurRow          int32 // 突出显示当前行颜色, ABGR 颜色
	ClrMatch           int32 // 设置匹配文本背景色, ABGR 颜色
}

// Edit 数据复制
type Edit_Data_Copy_ struct {
	NCount      int32                  // 内容数量
	NStyleCount int32                  // 样式数量
	PStyle      *Edit_Data_Copy_Style_ // 样式数组
	PData       *uint32                // 内容数组
}

// Edit 样式信息
type Edit_Style_Info_ struct {
	Type            xcc.Edit_Style_Type_ // 样式类型: Edit_Style_Type_
	NRef            uint16               // 引用计数
	HFont_image_obj int                  // 字体,图片,UI对象句柄
	Color           int32                // 颜色
	BColor          bool                 // 是否使用颜色
}

// Edit 数据复制-样式
type Edit_Data_Copy_Style_ struct {
	HFont_image_obj int   // 字体,图片,UI对象句柄
	Color           int32 // 颜色
	BColor          bool  // 是否使用颜色
}

// Position_ 位置点
type Position_ struct {
	IRow    int32 // 行索引
	IColumn int32 // 列索引
}

// Font_Info_ 字体信息
type Font_Info_ struct {
	NSize  int32          // 字体大小, 单位(pt,磅).
	NStyle xcc.FontStyle_ // 字体样式: FontStyle_
	Name   [32]uint16     // 字体名称, 使用xc.Font_Info_Name()函数转换为string.
}

// ListBox 列表框项信息2
type ListBox_Item_Info_ struct {
	NUserData  int   // 用户绑定数据
	NHeight    int32 // 项高度, -1使用默认高度
	NSelHeight int32 // 项选中时高度, -1使用默认高度
}

// MonthCal_item_ 月历元素项数据
type MonthCal_item_ struct {
	NDay   int32             // 日期
	NType  int32             // 1上月, 2当月, 3下月
	NState xcc.CombinedState // 组合状态, MonthCal_State_Flag_
	RcItem RECT              // 项坐标
}

// PGrid 属性网格项信息
type PropertyGrid_item_ struct {
	NType         int32 // 类型
	NID           int32 // 项ID
	NDepth        int32 // 深度
	NUserData     int   // 用户数据
	NNameColWidth int32 // 名称列宽度
	RcItem        RECT  // 坐标
	RcExpand      RECT  // 展开
	BExpand       bool  // 是否展开
	BShow         bool  // 是否可见
}

// 列表视项ID
type ListView_Item_Id_ struct {
	IGroup int32 // 组索引
	IItem  int32 // 项索引
}
