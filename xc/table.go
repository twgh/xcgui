package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 表格_创建, 返回句柄.
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
func XTable_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xTable_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 表格_重置.
//
// hShape: 形状对象句柄.
//
// nRow: 行数量.
//
// nCol: 列数量.
func XTable_Reset(hShape int, nRow int, nCol int) int {
	r, _, _ := xTable_Reset.Call(uintptr(hShape), uintptr(nRow), uintptr(nCol))
	return int(r)
}

// 表格_组合行.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func XTable_ComboRow(hShape int, iRow int, iCol int, count int) int {
	r, _, _ := xTable_ComboRow.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(count))
	return int(r)
}

// 表格_组合列.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func XTable_ComboCol(hShape int, iRow int, iCol int, count int) int {
	r, _, _ := xTable_ComboCol.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(count))
	return int(r)
}

// 表格_置列宽.
//
// hShape: 形状对象句柄.
//
// iCol: 列索引.
//
// width: 宽度.
func XTable_SetColWidth(hShape int, iCol int, width int) int {
	r, _, _ := xTable_SetColWidth.Call(uintptr(hShape), uintptr(iCol), uintptr(width))
	return int(r)
}

// 表格_置行高.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// height: 高度.
func XTable_SetRowHeight(hShape int, iRow int, height int) int {
	r, _, _ := xTable_SetRowHeight.Call(uintptr(hShape), uintptr(iRow), uintptr(height))
	return int(r)
}

// 表格_置边颜色.
//
// hShape: 形状对象句柄.
//
// color: 颜色.
func XTable_SetBorderColor(hShape int, color int) int {
	r, _, _ := xTable_SetBorderColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 表格_置文本颜色.
//
// hShape: 形状对象句柄.
//
// color: 颜色.
func XTable_SetTextColor(hShape int, color int) int {
	r, _, _ := xTable_SetTextColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}

// 表格_置字体.
//
// hShape: 形状对象句柄.
//
// hFont: 字体句柄.
func XTable_SetFont(hShape int, hFont int) int {
	r, _, _ := xTable_SetFont.Call(uintptr(hShape), uintptr(hFont))
	return int(r)
}

// 表格_置项内填充.
//
// hShape: 形状对象句柄.
//
// leftSize: 内填充大小.
//
// topSize: 内填充大小.
//
// rightSize: 内填充大小.
//
// bottomSize: 内填充大小.
func XTable_SetItemPadding(hShape int, leftSize int, topSize int, rightSize int, bottomSize int) int {
	r, _, _ := xTable_SetItemPadding.Call(uintptr(hShape), uintptr(leftSize), uintptr(topSize), uintptr(rightSize), uintptr(bottomSize))
	return int(r)
}

// 表格_置项文本.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func XTable_SetItemText(hShape int, iRow int, iCol int, pText string) int {
	r, _, _ := xTable_SetItemText.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), common.StrPtr(pText))
	return int(r)
}

// 表格_置项字体.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func XTable_SetItemFont(hShape int, iRow int, iCol int, hFont int) int {
	r, _, _ := xTable_SetItemFont.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(hFont))
	return int(r)
}

// 表格_置项文本对齐.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func XTable_SetItemTextAlign(hShape int, iRow int, iCol int, nAlign xcc.TextFormatFlag_) int {
	r, _, _ := xTable_SetItemTextAlign.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(nAlign))
	return int(r)
}

// 表格_置项文本色.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func XTable_SetItemTextColor(hShape int, iRow int, iCol int, color int, bColor bool) int {
	r, _, _ := xTable_SetItemTextColor.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(color), common.BoolPtr(bColor))
	return int(r)
}

// 表格_置项背景色.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func XTable_SetItemBkColor(hShape int, iRow int, iCol int, color int, bColor bool) int {
	r, _, _ := xTable_SetItemBkColor.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(color), common.BoolPtr(bColor))
	return int(r)
}

// 表格_置项线.
//
// hShape: 形状对象句柄.
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
// color: 颜色.
func XTable_SetItemLine(hShape int, iRow1 int, iCol1 int, iRow2 int, iCol2 int, nFlag int, color int) int {
	r, _, _ := xTable_SetItemLine.Call(uintptr(hShape), uintptr(iRow1), uintptr(iCol1), uintptr(iRow2), uintptr(iCol2), uintptr(nFlag), uintptr(color))
	return int(r)
}

// 表格_置项标识.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, Table_Flag_.
func XTable_SetItemFlag(hShape int, iRow int, iCol int, flag xcc.Table_Flag_) int {
	r, _, _ := xTable_SetItemFlag.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(flag))
	return int(r)
}

// 表格_取项坐标.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func XTable_GetItemRect(hShape int, iRow int, iCol int, pRect *RECT) bool {
	r, _, _ := xTable_GetItemRect.Call(uintptr(hShape), uintptr(iRow), uintptr(iCol), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}
