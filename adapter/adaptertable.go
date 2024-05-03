package adapter

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// AdapterTable 数据适配器-XList-XListBox.
type AdapterTable struct {
	adapter
}

// 数据适配器表_创建, 创建列表框元素数据适配器.
func NewAdapterTable() *AdapterTable {
	p := &AdapterTable{}
	p.SetHandle(xc.XAdTable_Create())
	return p
}

// 从句柄创建对象.
func NewAdapterTableByHandle(handle int) *AdapterTable {
	p := &AdapterTable{}
	p.SetHandle(handle)
	return p
}

// 数据适配器表_排序, 对内容进行排序.
//
// iColumn: 要排序的列索引。.
//
// bAscending: 是否按照升序方式排序.
func (a *AdapterTable) Sort(iColumn int32, bAscending bool) *AdapterTable {
	xc.XAdTable_Sort(a.Handle, iColumn, bAscending)
	return a
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemDataType(iItem, iColumn int32) xcc.Adapter_Date_Type_ {
	return xc.XAdTable_GetItemDataType(a.Handle, iItem, iColumn)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemDataTypeEx(iItem int32, pName string) xcc.Adapter_Date_Type_ {
	return xc.XAdTable_GetItemDataTypeEx(a.Handle, iItem, pName)
}

// 数据适配器表_添加列, 添加数据列.
//
// pName: 字段称.
func (a *AdapterTable) AddColumn(pName string) int32 {
	return xc.XAdTable_AddColumn(a.Handle, pName)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
// pColName: 列名, 多个列名用逗号分开.
func (a *AdapterTable) SetColumn(pColName string) int32 {
	return xc.XAdTable_SetColumn(a.Handle, pColName)
}

// 数据适配器表_添加行文本, 添加数据行, 默认值放到第一列, 返回行索引值.
//
// pValue: 值.
func (a *AdapterTable) AddRowText(pValue string) int32 {
	return xc.XAdTable_AddRowText(a.Handle, pValue)
}

// 数据适配器表_添加行文本扩展, 添加数据行, 填充指定列数据, 返回行索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) AddRowTextEx(pName string, pValue string) int32 {
	return xc.XAdTable_AddRowTextEx(a.Handle, pName, pValue)
}

// 数据适配器表_添加行图片, 添加数据行, 默认值放到第一列, 返回行索引值.
//
// hImage: 图片句柄.
func (a *AdapterTable) AddRowImage(hImage int) int32 {
	return xc.XAdTable_AddRowImage(a.Handle, hImage)
}

// 数据适配器表_添加行图片扩展, 添加数据行, 并填充指定列数据.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) AddRowImageEx(pName string, hImage int) int32 {
	return xc.XAdTable_AddRowImageEx(a.Handle, pName, hImage)
}

// 数据适配器表_插入行文本, 插入数据行, 填充第一列数据, 返回行索引.
//
// iRow: 插入位置索引.
//
// pValue: 值.
func (a *AdapterTable) InsertRowText(iRow int32, pValue string) int32 {
	return xc.XAdTable_InsertRowText(a.Handle, iRow, pValue)
}

// 数据适配器表_插入行文本扩展, 插入数据行, 并填充指定列数据, 返回行索引.
//
// iRow: 插入位置索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) InsertRowTextEx(iRow int32, pName string, pValue string) int32 {
	return xc.XAdTable_InsertRowTextEx(a.Handle, iRow, pName, pValue)
}

// 数据适配器表_插入行图片, 插入数据行, 填充第一列数据, 返回行索引.
//
// iRow: 插入位置索引.
//
// hImage: 图片句柄.
func (a *AdapterTable) InsertRowImage(iRow int32, hImage int) int32 {
	return xc.XAdTable_InsertRowImage(a.Handle, iRow, hImage)
}

// 数据适配器表_插入行图片扩展, 插入数据行, 并填充指定列数据, 返回行索引.
//
// iRow: 插入位置索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) InsertRowImageEx(iRow int32, pName string, hImage int) int32 {
	return xc.XAdTable_InsertRowImageEx(a.Handle, iRow, pName, hImage)
}

// 数据适配器表_置项文本, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterTable) SetItemText(iItem, iColumn int32, pValue string) bool {
	return xc.XAdTable_SetItemText(a.Handle, iItem, iColumn, pValue)
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) SetItemTextEx(iItem int32, pName string, pValue string) bool {
	return xc.XAdTable_SetItemTextEx(a.Handle, iItem, pName, pValue)
}

// 数据适配器表_置项整数值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 值.
func (a *AdapterTable) SetItemInt(iItem, iColumn int32, nValue int32) bool {
	return xc.XAdTable_SetItemInt(a.Handle, iItem, iColumn, nValue)
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// nValue: 值.
func (a *AdapterTable) SetItemIntEx(iItem int32, pName string, nValue int32) bool {
	return xc.XAdTable_SetItemIntEx(a.Handle, iItem, pName, nValue)
}

// 数据适配器表_置项浮点值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fValue: 值.
func (a *AdapterTable) SetItemFloat(iItem, iColumn int32, fValue float32) bool {
	return xc.XAdTable_SetItemFloat(a.Handle, iItem, iColumn, fValue)
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// fValue: 值.
func (a *AdapterTable) SetItemFloatEx(iItem int32, pName string, fValue float32) bool {
	return xc.XAdTable_SetItemFloatEx(a.Handle, iItem, pName, fValue)
}

// 数据适配器表_置项图片, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterTable) SetItemImage(iItem, iColumn int32, hImage int) bool {
	return xc.XAdTable_SetItemImage(a.Handle, iItem, iColumn, hImage)
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) SetItemImageEx(iItem int32, pName string, hImage int) bool {
	return xc.XAdTable_SetItemImageEx(a.Handle, iItem, pName, hImage)
}

// 数据适配器表_删除行.
//
// iRow: 行索引.
func (a *AdapterTable) DeleteRow(iRow int32) bool {
	return xc.XAdTable_DeleteRow(a.Handle, iRow)
}

// 数据适配器表_删除行扩展, 删除多个行.
//
// iRow: 行起始索引.
//
// nCount: 删除行数量.
func (a *AdapterTable) DeleteRowEx(iRow, nCount int32) bool {
	return xc.XAdTable_DeleteRowEx(a.Handle, iRow, nCount)
}

// 数据适配器表_删除行全部, 删除所有行.
func (a *AdapterTable) DeleteRowAll() *AdapterTable {
	xc.XAdTable_DeleteRowAll(a.Handle)
	return a
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.
func (a *AdapterTable) DeleteColumnAll() *AdapterTable {
	xc.XAdTable_DeleteColumnAll(a.Handle)
	return a
}

// 数据适配器表_取列数量, 获取列数量.
func (a *AdapterTable) GetCountColumn() int32 {
	return xc.XAdTable_GetCountColumn(a.Handle)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemText(iItem, iColumn int32) string {
	return xc.XAdTable_GetItemText(a.Handle, iItem, iColumn)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemTextEx(iItem int32, pName string) string {
	return xc.XAdTable_GetItemTextEx(a.Handle, iItem, pName)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemImage(iItem, iColumn int32) int {
	return xc.XAdTable_GetItemImage(a.Handle, iItem, iColumn)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemImageEx(iItem int32, pName string) int {
	return xc.XAdTable_GetItemImageEx(a.Handle, iItem, pName)
}

// 数据适配器表_取项整数值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemInt(iItem, iColumn int32, pOutValue *int32) bool {
	return xc.XAdTable_GetItemInt(a.Handle, iItem, iColumn, pOutValue)
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemIntEx(iItem int32, pName string, pOutValue *int32) bool {
	return xc.XAdTable_GetItemIntEx(a.Handle, iItem, pName, pOutValue)
}

// 数据适配器表_取项浮点值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemFloat(iItem, iColumn int32, pOutValue *float32) bool {
	return xc.XAdTable_GetItemFloat(a.Handle, iItem, iColumn, pOutValue)
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemFloatEx(iItem int32, pName string, pOutValue *float32) bool {
	return xc.XAdTable_GetItemFloatEx(a.Handle, iItem, pName, pOutValue)
}

// 数据适配器表_取行数量, 获取行数量.
func (a *AdapterTable) GetCountRow() int32 {
	return xc.XAdTable_GetCountRow(a.Handle)
}
