package xc

import (
	"github.com/twgh/xcgui/xcc"
)

/*
// //////////////////////////////////// 旧版函数 //////////////////////////////////////
*/

// Deprecated
//
// !这是旧版函数, 请使用 ARGB
func ABGR(r, g, b, a byte) int {
	return ARGB(r, g, b, a)
}

// Deprecated
//
// !这是旧版函数, 请使用 ARGB2
func ABGR2(rgb int, a byte) int {
	return ARGB2(rgb, a)
}

// Deprecated
//
// !这是旧版函数, 请使用 RGB2ARGB
func RGB2ABGR(rgb int, a byte) int {
	return RGB2ARGB(rgb, a)
}

// Deprecated
//
// !这是旧版函数, 请使用 HexRGB2ARGB
func HexRGB2ABGR(str string, a byte) int {
	return HexRGB2ARGB(str, a)
}

// Deprecated
//
// !这是旧版函数, 请使用 xc.XWnd_ClientToScreen 或 wapi.ClientToScreen
func ClientToScreen(hWindow int, pPoint *POINT) {
	XWnd_ClientToScreen(hWindow, pPoint)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_AddRowText
func XList_AddItemText(hEle int, pText string) int {
	return XList_AddRowText(hEle, pText)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_AddRowTextEx
func XList_AddItemTextEx(hEle int, pName string, pText string) int {
	return XList_AddRowTextEx(hEle, pName, pText)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_AddRowImage
func XList_AddItemImage(hEle int, hImage int) int {
	return XList_AddRowImage(hEle, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_AddRowImageEx
func XList_AddItemImageEx(hEle int, pName string, hImage int) int {
	return XList_AddRowImageEx(hEle, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_InsertRowText
func XList_InsertItemText(hEle int, iItem int, pValue string) int {
	return XList_InsertRowText(hEle, iItem, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_InsertRowTextEx
func XList_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {
	return XList_InsertRowTextEx(hEle, iItem, pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_InsertRowImage
func XList_InsertItemImage(hEle int, iItem int, hImage int) int {
	return XList_InsertRowImage(hEle, iItem, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_InsertRowImageEx
func XList_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {
	return XList_InsertRowImageEx(hEle, iItem, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_DeleteRow
func XList_DeleteItem(hEle int, iRow int) bool {
	return XList_DeleteRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_DeleteRowEx
func XList_DeleteItemEx(hEle int, iRow int, nCount int) bool {
	return XList_DeleteRowEx(hEle, iRow, nCount)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_DeleteRowAll
func XList_DeleteItemAll(hEle int) int {
	return XList_DeleteRowAll(hEle)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_SetSelectRow
func XList_SetSelectItem(hEle int, iRow int) bool {
	return XList_SetSelectRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetSelectRow
func XList_GetSelectItem(hEle int) int {
	return XList_GetSelectRow(hEle)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetSelectRowCount
func XList_GetSelectItemCount(hEle int) int {
	return XList_GetSelectRowCount(hEle)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_AddSelectRow
func XList_AddSelectItem(hEle int, iRow int) bool {
	return XList_AddSelectRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_VisibleRow
func XList_VisibleItem(hEle int, iRow int) int {
	return XList_VisibleRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_CancelSelectRow
func XList_CancelSelectItem(hEle int, iRow int) bool {
	return XList_CancelSelectRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetHeaderColumnIndexFromHXCGUI
func XList_GetHeaderItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	return XList_GetHeaderColumnIndexFromHXCGUI(hEle, hXCGUI)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetRowIndexFromHXCGUI
func XList_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	return XList_GetRowIndexFromHXCGUI(hEle, hXCGUI)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_SetRowHeightDefault
func XList_SetItemHeightDefault(hEle int, nHeight int32, nSelHeight int32) int {
	return XList_SetRowHeightDefault(hEle, nHeight, nSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetRowHeightDefault
func XList_GetItemHeightDefault(hEle int, pHeight *int32, pSelHeight *int32) int {
	return XList_GetRowHeightDefault(hEle, pHeight, pSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_SetRowHeight
func XList_SetItemHeight(hEle int, iRow int, nHeight, nSelHeight int32) int {
	return XList_SetRowHeight(hEle, iRow, nHeight, nSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_GetRowHeight
func XList_GetItemHeight(hEle int, iRow int, pHeight, pSelHeight *int32) int {
	return XList_GetRowHeight(hEle, iRow, pHeight, pSelHeight)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_EnableRowBkFull
func XList_EnableItemBkFullRow(hEle int, bFull bool) {
	XList_EnableRowBkFull(hEle, bFull)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_SetDrawRowBkFlags
func XList_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) {
	XList_SetDrawRowBkFlags(hEle, nFlags)
}

// Deprecated
//
// !这是旧版函数, 请使用 XList_RefreshRow
func XList_RefreshItem(hEle int, iRow int) {
	XList_RefreshRow(hEle, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_AddRowText
func XAdTable_AddItemText(hAdapter int, pValue string) int32 {
	return XAdTable_AddRowText(hAdapter, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_AddRowTextEx
func XAdTable_AddItemTextEx(hAdapter int, pName string, pValue string) int32 {
	return XAdTable_AddRowTextEx(hAdapter, pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_AddRowImage
func XAdTable_AddItemImage(hAdapter int, hImage int) int32 {
	return XAdTable_AddRowImage(hAdapter, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_AddRowImageEx
func XAdTable_AddItemImageEx(hAdapter int, pName string, hImage int) int32 {
	return XAdTable_AddRowImageEx(hAdapter, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_InsertRowText
func XAdTable_InsertItemText(hAdapter int, iRow int32, pValue string) int32 {
	return XAdTable_InsertRowText(hAdapter, iRow, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_InsertRowTextEx
func XAdTable_InsertItemTextEx(hAdapter int, iRow int32, pName string, pValue string) int32 {
	return XAdTable_InsertRowTextEx(hAdapter, iRow, pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_InsertRowImage
func XAdTable_InsertItemImage(hAdapter int, iRow int32, hImage int) int32 {
	return XAdTable_InsertRowImage(hAdapter, iRow, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_InsertRowImageEx
func XAdTable_InsertItemImageEx(hAdapter int, iRow int32, pName string, hImage int) int32 {
	return XAdTable_InsertRowImageEx(hAdapter, iRow, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_GetCountRow
func XAdTable_GetCount(hAdapter int) int32 {
	return XAdTable_GetCountRow(hAdapter)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_DeleteRow
func XAdTable_DeleteItem(hAdapter int, iRow int32) bool {
	return XAdTable_DeleteRow(hAdapter, iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_DeleteRowEx
func XAdTable_DeleteItemEx(hAdapter int, iRow int32, nCount int32) bool {
	return XAdTable_DeleteRowEx(hAdapter, iRow, nCount)
}

// Deprecated
//
// !这是旧版函数, 请使用 XAdTable_DeleteRowAll
func XAdTable_DeleteItemAll(hAdapter int) {
	XAdTable_DeleteRowAll(hAdapter)
}

// Deprecated
//
// !这是旧版函数, 请使用 XBkM_SetInfo
func XBkM_SetBkInfo(hBkInfoM int, pText string) int32 {
	return XBkM_SetInfo(hBkInfoM, pText)
}
