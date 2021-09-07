package xc

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type POINT struct {
	X int32
	Y int32
}

type SIZE struct {
	CX int32
	CY int32
}

type LOGFONTW struct {
	LfHeight         int32     // 高度
	LfWidth          int32     // 宽度
	LfEscapement     int32     // 指定角度
	LfOrientation    int32     // 字符基线
	LfWeight         int32     // 字体粗细
	LfItalic         uint8     // 是否斜体
	LfUnderline      uint8     // 是否下划线
	LfStrikeOut      uint8     // 是否删除线
	LfCharSet        uint8     // 字符集
	LfOutPrecision   uint8     // 输出精度
	LfClipPrecision  uint8     // 剪辑精度
	LfQuality        uint8     // 输出质量
	LfPitchAndFamily uint8     // 字符间距
	LfFaceName       [64]uint8 // 字体名称
}

type Editor_Color_ struct {
	ClrMargin1         int32 // 背景色1
	ClrMargin2         int32 // 背景色2
	ClrMarginText      int32 // 文本色
	ClrBreakpoint      int32 // 断点色
	ClrBreakpointArrow int32 // 断点箭头色
	ClrRun             int32 // 当前执行位置指示色
	ClrMatch           int32 // 设置匹配文本背景色
}

type Edit_Data_Copy_ struct {
	NCount      int32                  // 内容数量
	NStyleCount int32                  // 样式数量
	PStyle      *Edit_Data_Copy_Style_ // 样式数组
	PData       *int32                 // 内容数组
}

type Edit_Style_Info_ struct {
	Type            int32  // 样式类型
	NRef            uint16 // 引用计数
	HFont_image_obj int32  // 字体,图片,UI对象
	Color           int32  // 颜色
	BColor          bool   // 是否使用颜色
}

type Edit_Data_Copy_Style_ struct {
	HFont_image_obj int32 // 字体,图片,UI对象
	Color           int32 // 颜色
	BColor          bool  // 是否使用颜色
}

type Position_ struct {
	IRow    int32 // 行号
	IColumn int32 // 列号
}

type Font_Info_ struct {
	NSize  int32      // 字体大小,单位(pt,磅).
	NStyle int32      // 字体样式 FontStyle_
	Name   [32]uint16 // 字体名称
}

type ListBox_Item_Info_ struct {
	NUserData  int32 // 用户绑定数据
	NHeight    int32 // 默认高度
	NSelHeight int32 // 选中高度
}
