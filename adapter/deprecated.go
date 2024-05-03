package adapter

/*
// //////////////////////////////////// 旧版函数 //////////////////////////////////////
*/

// Deprecated
//
// !这是旧版函数, 请使用 AddRowText
func (a *AdapterTable) AddItemText(pValue string) int32 {
	return a.AddRowText(pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowTextEx
func (a *AdapterTable) AddItemTextEx(pName string, pValue string) int32 {
	return a.AddRowTextEx(pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImage
func (a *AdapterTable) AddItemImage(hImage int) int32 {
	return a.AddRowImage(hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 AddRowImageEx
func (a *AdapterTable) AddItemImageEx(pName string, hImage int) int32 {
	return a.AddRowImageEx(pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowText
func (a *AdapterTable) InsertItemText(iItem int32, pValue string) int32 {
	return a.InsertRowText(iItem, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowTextEx
func (a *AdapterTable) InsertItemTextEx(iItem int32, pName string, pValue string) int32 {
	return a.InsertRowTextEx(iItem, pName, pValue)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImage
func (a *AdapterTable) InsertItemImage(iItem int32, hImage int) int32 {
	return a.InsertRowImage(iItem, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 InsertRowImageEx
func (a *AdapterTable) InsertItemImageEx(iItem int32, pName string, hImage int) int32 {
	return a.InsertRowImageEx(iItem, pName, hImage)
}

// Deprecated
//
// !这是旧版函数, 请使用 GetCountRow
func (a *AdapterTable) GetCount() int32 {
	return a.GetCountRow()
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRow
func (a *AdapterTable) DeleteItem(iRow int32) bool {
	return a.DeleteRow(iRow)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRowEx
func (a *AdapterTable) DeleteItemEx(iRow, nCount int32) bool {
	return a.DeleteRowEx(iRow, nCount)
}

// Deprecated
//
// !这是旧版函数, 请使用 DeleteRowAll
func (a *AdapterTable) DeleteItemAll() *AdapterTable {
	return a.DeleteRowAll()
}
