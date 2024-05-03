package widget

import (
	"github.com/twgh/xcgui/xcc"
)

/*
// //////////////////////////////////// 旧版函数 //////////////////////////////////////
*/

// Deprecated
//
// !这是旧版函数, 请使用 AddRowText
func (l *List) AddItemText(pText string) int {
	return l.AddRowText(pText)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowTextEx
func (l *List) AddItemTextEx(pName string, pText string) int {
	return l.AddRowTextEx(pName, pText)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImage
func (l *List) AddItemImage(hImage int) int {
	return l.AddRowImage(hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImageEx
func (l *List) AddItemImageEx(pName string, hImage int) int {
	return l.AddRowImageEx(pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowText
func (l *List) InsertItemText(iItem int, pValue string) int {
	return l.InsertRowText(iItem, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowTextEx
func (l *List) InsertItemTextEx(iItem int, pName string, pValue string) int {
	return l.InsertRowTextEx(iItem, pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImage
func (l *List) InsertItemImage(iItem int, hImage int) int {
	return l.InsertRowImage(iItem, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImageEx
func (l *List) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return l.InsertRowImageEx(iItem, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRow
func (l *List) DeleteItem(iItem int) bool {
	return l.DeleteRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRowEx
func (l *List) DeleteItemEx(iItem int, nCount int) bool {
	return l.DeleteRowEx(iItem, nCount)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRowAll
func (l *List) DeleteItemAll() *List {
	return l.DeleteRowAll()
}

// Deprecated
//
// !这是旧版函数, 请使用 SetSelectRow
func (l *List) SetSelectItem(iItem int) bool {
	return l.SetSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetSelectRow
func (l *List) GetSelectItem() int {
	return l.GetSelectRow()
}

// Deprecated
//
// !这是旧版函数, 请使用 GetSelectRowCount
func (l *List) GetSelectItemCount() int {
	return l.GetSelectRowCount()
}

// Deprecated
//
// !这是旧版函数, 请使用 AddSelectRow
func (l *List) AddSelectItem(iItem int) bool {
	return l.AddSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 VisibleRow
func (l *List) VisibleItem(iItem int) *List {
	return l.VisibleRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 CancelSelectRow
func (l *List) CancelSelectItem(iItem int) bool {
	return l.CancelSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetHeaderColumnIndexFromHXCGUI
func (l *List) GetHeaderItemIndexFromHXCGUI(hXCGUI int) int {
	return l.GetHeaderColumnIndexFromHXCGUI(hXCGUI)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetRowIndexFromHXCGUI
func (l *List) GetItemIndexFromHXCGUI(hXCGUI int) int {
	return l.GetRowIndexFromHXCGUI(hXCGUI)
}

// Deprecated
//
// !这是旧版函数, 请使用 SetRowHeightDefault
func (l *List) SetItemHeightDefault(nHeight int32, nSelHeight int32) *List {
	return l.SetRowHeightDefault(nHeight, nSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetRowHeightDefault
func (l *List) GetItemHeightDefault(pHeight *int32, pSelHeight *int32) *List {
	return l.GetRowHeightDefault(pHeight, pSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 SetRowHeight
func (l *List) SetItemHeight(iRow int, nHeight, nSelHeight int32) *List {
	return l.SetRowHeight(iRow, nHeight, nSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetRowHeight
func (l *List) GetItemHeight(iRow int, pHeight, pSelHeight *int32) *List {
	return l.GetRowHeight(iRow, pHeight, pSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 EnableRowBkFull
func (l *List) EnableItemBkFullRow(bFull bool) *List {
	l.EnableRowBkFull(bFull)
	return l
}

// Deprecated
//
// !这是旧版函数, 请使用 SetDrawRowBkFlags
func (l *List) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) *List {
	l.SetDrawRowBkFlags(nFlags)
	return l
}

// Deprecated
//
// !这是旧版函数, 请使用 RefreshRow
func (l *List) RefreshItem(iRow int) *List {
	l.RefreshRow(iRow)
	return l
}
