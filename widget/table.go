package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Table 表格.
type Table struct {
	Shape
}

// 表格_创建, 失败返回 nil.
//
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewTable(x, y, cx, cy int32, hParent int) *Table {
	return NewTableByHandle(xc.XTable_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewTableByHandle(handle int) *Table {
	if handle <= 0 {
		return nil
	}
	p := &Table{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewTableByName(name string) *Table {
	return NewTableByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewTableByUID(nUID int32) *Table {
	return NewTableByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewTableByUIDName(name string) *Table {
	return NewTableByHandle(xc.XC_GetObjectByUIDName(name))
}

// 表格_重置.
//
// nRow: 行数量.
//
// nCol: 列数量.
func (t *Table) Reset(nRow int32, nCol int32) *Table {
	xc.XTable_Reset(t.Handle, nRow, nCol)
	return t
}

// 表格_组合行.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) ComboRow(iRow int32, iCol int32, count int32) *Table {
	xc.XTable_ComboRow(t.Handle, iRow, iCol, count)
	return t
}

// 表格_组合列.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) ComboCol(iRow, iCol, count int32) *Table {
	xc.XTable_ComboCol(t.Handle, iRow, iCol, count)
	return t
}

// 表格_置列宽.
//
// iCol: 列索引.
//
// width: 宽度.
func (t *Table) SetColWidth(iCol, width int32) *Table {
	xc.XTable_SetColWidth(t.Handle, iCol, width)
	return t
}

// 表格_置行高.
//
// iRow: 行索引.
//
// height: 高度.
func (t *Table) SetRowHeight(iRow, height int32) *Table {
	xc.XTable_SetRowHeight(t.Handle, iRow, height)
	return t
}

// 表格_置边颜色.
//
// color: xc.RGBA 颜色.
func (t *Table) SetBorderColor(color uint32) *Table {
	xc.XTable_SetBorderColor(t.Handle, color)
	return t
}

// 表格_置文本颜色.
//
// color: xc.RGBA 颜色.
func (t *Table) SetTextColor(color uint32) *Table {
	xc.XTable_SetTextColor(t.Handle, color)
	return t
}

// 表格_置字体.
//
// hFont: 字体句柄.
func (t *Table) SetFont(hFont int) *Table {
	xc.XTable_SetFont(t.Handle, hFont)
	return t
}

// 表格_置项内填充.
//
// leftSize: 内填充大小.
//
// topSize: 内填充大小.
//
// rightSize: 内填充大小.
//
// bottomSize: 内填充大小.
func (t *Table) SetItemPadding(leftSize, topSize, rightSize, bottomSize int32) *Table {
	xc.XTable_SetItemPadding(t.Handle, leftSize, topSize, rightSize, bottomSize)
	return t
}

// 表格_置项文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func (t *Table) SetItemText(iRow, iCol int32, pText string) *Table {
	xc.XTable_SetItemText(t.Handle, iRow, iCol, pText)
	return t
}

// 表格_置项字体.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func (t *Table) SetItemFont(iRow, iCol int32, hFont int) *Table {
	xc.XTable_SetItemFont(t.Handle, iRow, iCol, hFont)
	return t
}

// 表格_置项文本对齐.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (t *Table) SetItemTextAlign(iRow, iCol int32, nAlign xcc.TextFormatFlag_) *Table {
	xc.XTable_SetItemTextAlign(t.Handle, iRow, iCol, nAlign)
	return t
}

// 表格_置项文本色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: xc.RGBA 颜色.
//
// bColor: 是否使用.
func (t *Table) SetItemTextColor(iRow, iCol int32, color uint32, bColor bool) *Table {
	xc.XTable_SetItemTextColor(t.Handle, iRow, iCol, color, bColor)
	return t
}

// 表格_置项背景色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: xc.RGBA 颜色.
//
// bColor: 是否使用.
func (t *Table) SetItemBkColor(iRow, iCol int32, color uint32, bColor bool) *Table {
	xc.XTable_SetItemBkColor(t.Handle, iRow, iCol, color, bColor)
	return t
}

// 表格_置项线.
//
// iRow1: 行索引1.
//
// iCol1: 列索引1.
//
// iRow2: 行索引2.
//
// iCol2: 列索引2.
//
// nFlag: 标识, Table_Line_Flag_, 暂时没有, 填0.
//
// color: xc.RGBA 颜色.
func (t *Table) SetItemLine(iRow1, iCol1, iRow2, iCol2 int32, nFlag int32, color uint32) *Table {
	xc.XTable_SetItemLine(t.Handle, iRow1, iCol1, iRow2, iCol2, nFlag, color)
	return t
}

// 表格_置项标识.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, Table_Flag_.
func (t *Table) SetItemFlag(iRow, iCol int32, flag xcc.Table_Flag_) *Table {
	xc.XTable_SetItemFlag(t.Handle, iRow, iCol, flag)
	return t
}

// 表格_取项坐标.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func (t *Table) GetItemRect(iRow, iCol int32, pRect *xc.RECT) bool {
	return xc.XTable_GetItemRect(t.Handle, iRow, iCol, pRect)
}

// 表格_取行数.
func (t *Table) GetRowCount() int32 {
	return xc.XTable_GetRowCount(t.Handle)
}

// 表格_取列数.
func (t *Table) GetColCount() int32 {
	return xc.XTable_GetColCount(t.Handle)
}

// 表格_置项文本Ex.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
//
// textColor: 文本颜色, xc.RGBA 颜色.
//
// bkColor: 背景颜色, xc.RGBA 颜色.
//
// bTextColor: 是否使用文本颜色.
//
// bBkColor: 是否使用背景颜色.
//
// hFont: 炫彩字体句柄, 可为0.
func (t *Table) SetItemTextEx(iRow, iCol int32, pText string, textColor, bkColor uint32, bTextColor, bBkColor bool, hFont int) *Table {
	xc.XTable_SetItemTextEx(t.Handle, iRow, iCol, pText, textColor, bkColor, bTextColor, bBkColor, hFont)
	return t
}
