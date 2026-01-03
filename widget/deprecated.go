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
func (l *List) AddItemText(text string) int32 {
	return l.AddRowText(text)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowTextEx
func (l *List) AddItemTextEx(name string, text string) int32 {
	return l.AddRowTextEx(name, text)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImage
func (l *List) AddItemImage(hImage int) int32 {
	return l.AddRowImage(hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImageEx
func (l *List) AddItemImageEx(name string, hImage int) int32 {
	return l.AddRowImageEx(name, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowText
func (l *List) InsertItemText(iItem int32, pValue string) int32 {
	return l.InsertRowText(iItem, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowTextEx
func (l *List) InsertItemTextEx(iItem int32, name string, pValue string) int32 {
	return l.InsertRowTextEx(iItem, name, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImage
func (l *List) InsertItemImage(iItem int32, hImage int) int32 {
	return l.InsertRowImage(iItem, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImageEx
func (l *List) InsertItemImageEx(iItem int32, name string, hImage int) int32 {
	return l.InsertRowImageEx(iItem, name, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRow
func (l *List) DeleteItem(iItem int32) bool {
	return l.DeleteRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRowEx
func (l *List) DeleteItemEx(iItem, nCount int32) bool {
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
func (l *List) SetSelectItem(iItem int32) bool {
	return l.SetSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetSelectRow
func (l *List) GetSelectItem() int32 {
	return l.GetSelectRow()
}

// Deprecated
//
// !这是旧版函数, 请使用 GetSelectRowCount
func (l *List) GetSelectItemCount() int32 {
	return l.GetSelectRowCount()
}

// Deprecated
//
// !这是旧版函数, 请使用 AddSelectRow
func (l *List) AddSelectItem(iItem int32) bool {
	return l.AddSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 VisibleRow
func (l *List) VisibleItem(iItem int32) *List {
	return l.VisibleRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 CancelSelectRow
func (l *List) CancelSelectItem(iItem int32) bool {
	return l.CancelSelectRow(iItem)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetHeaderColumnIndexFromHXCGUI
func (l *List) GetHeaderItemIndexFromHXCGUI(hXCGUI int) int32 {
	return l.GetHeaderColumnIndexFromHXCGUI(hXCGUI)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetRowIndexFromHXCGUI
func (l *List) GetItemIndexFromHXCGUI(hXCGUI int) int32 {
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
func (l *List) SetItemHeight(iRow int32, nHeight, nSelHeight int32) *List {
	return l.SetRowHeight(iRow, nHeight, nSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetRowHeight
func (l *List) GetItemHeight(iRow int32, pHeight, pSelHeight *int32) *List {
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
func (l *List) RefreshItem(iRow int32) *List {
	l.RefreshRow(iRow)
	return l
}
