package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 数据适配器表_创建, 创建列表框元素数据适配器, 返回数据适配器句柄.
func XAdTable_Create() int {
	r, _, _ := xAdTable_Create.Call()
	return int(r)
}

// 数据适配器表_排序, 对内容进行排序.
//
//	hAdapter: 数据适配器句柄.
//
//	iColumn: 要排序的列索引。.
//
//	bAscending: 是否按照升序方式排序.
func XAdTable_Sort(hAdapter int, iColumn int32, bAscending bool) {
	xAdTable_Sort.Call(uintptr(hAdapter), uintptr(iColumn), common.BoolPtr(bAscending))
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
func XAdTable_GetItemDataType(hAdapter int, iItem, iColumn int32) xcc.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataType.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return xcc.Adapter_Date_Type_(r)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
func XAdTable_GetItemDataTypeEx(hAdapter int, iItem int32, pName string) xcc.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataTypeEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return xcc.Adapter_Date_Type_(r)
}

// 数据适配器表_添加列, 添加数据列.
//
//	hAdapter: 数据适配器句柄.
//
//	pName: 字段称.
func XAdTable_AddColumn(hAdapter int, pName string) int32 {
	r, _, _ := xAdTable_AddColumn.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int32(r)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
//	hAdapter: 数据适配器句柄.
//
//	pColName: 列名, 多个列名用逗号分开.
func XAdTable_SetColumn(hAdapter int, pColName string) int32 {
	r, _, _ := xAdTable_SetColumn.Call(uintptr(hAdapter), common.StrPtr(pColName))
	return int32(r)
}

// 数据适配器表_添加行文本, 添加数据行, 默认值放到第一列, 返回行索引值.
//
//	hAdapter: 数据适配器句柄.
//
//	pValue: 值.
func XAdTable_AddRowText(hAdapter int, pValue string) int32 {
	r, _, _ := xAdTable_AddRowText.Call(uintptr(hAdapter), common.StrPtr(pValue))
	return int32(r)
}

// 数据适配器表_添加行文本扩展, 添加数据行, 填充指定列数据, 返回行索引.
//
//	hAdapter: 数据适配器句柄.
//
//	pName: 字段称.
//
//	pValue: 值.
func XAdTable_AddRowTextEx(hAdapter int, pName string, pValue string) int32 {
	r, _, _ := xAdTable_AddRowTextEx.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue))
	return int32(r)
}

// 数据适配器表_添加行图片, 添加数据行, 默认值放到第一列, 返回行索引值.
//
//	hAdapter: 数据适配器句柄.
//
//	hImage: 图片句柄.
func XAdTable_AddRowImage(hAdapter int, hImage int) int32 {
	r, _, _ := xAdTable_AddRowImage.Call(uintptr(hAdapter), uintptr(hImage))
	return int32(r)
}

// 数据适配器表_添加行图片扩展, 添加数据行, 并填充指定列数据.
//
//	hAdapter: 数据适配器句柄.
//
//	pName: 字段称.
//
//	hImage: 图片句柄.
func XAdTable_AddRowImageEx(hAdapter int, pName string, hImage int) int32 {
	r, _, _ := xAdTable_AddRowImageEx.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage))
	return int32(r)
}

// 数据适配器表_插入行文本, 插入数据行, 填充第一列数据, 返回行索引.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 插入位置索引.
//
//	pValue: 值.
func XAdTable_InsertRowText(hAdapter int, iRow int32, pValue string) int32 {
	r, _, _ := xAdTable_InsertRowText.Call(uintptr(hAdapter), uintptr(iRow), common.StrPtr(pValue))
	return int32(r)
}

// 数据适配器表_插入行文本扩展, 插入数据行, 并填充指定列数据, 返回行索引.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 插入位置索引.
//
//	pName: 字段称.
//
//	pValue: 值.
func XAdTable_InsertRowTextEx(hAdapter int, iRow int32, pName string, pValue string) int32 {
	r, _, _ := xAdTable_InsertRowTextEx.Call(uintptr(hAdapter), uintptr(iRow), common.StrPtr(pName), common.StrPtr(pValue))
	return int32(r)
}

// 数据适配器表_插入行图片, 插入数据行, 填充第一列数据, 返回行索引.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 插入位置索引.
//
//	hImage: 图片句柄.
func XAdTable_InsertRowImage(hAdapter int, iRow int32, hImage int) int32 {
	r, _, _ := xAdTable_InsertRowImage.Call(uintptr(hAdapter), uintptr(iRow), uintptr(hImage))
	return int32(r)
}

// 数据适配器表_插入行图片扩展, 插入数据行, 并填充指定列数据, 返回行索引.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 插入位置索引.
//
//	pName: 字段称.
//
//	hImage: 图片句柄.
func XAdTable_InsertRowImageEx(hAdapter int, iRow int32, pName string, hImage int) int32 {
	r, _, _ := xAdTable_InsertRowImageEx.Call(uintptr(hAdapter), uintptr(iRow), common.StrPtr(pName), uintptr(hImage))
	return int32(r)
}

// 数据适配器表_置项文本, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	pValue: 值.
func XAdTable_SetItemText(hAdapter int, iItem, iColumn int32, pValue string) bool {
	r, _, _ := xAdTable_SetItemText.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	pValue: 值.
func XAdTable_SetItemTextEx(hAdapter int, iItem int32, pName string, pValue string) bool {
	r, _, _ := xAdTable_SetItemTextEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器表_置项整数值, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	nValue: 值.
func XAdTable_SetItemInt(hAdapter int, iItem, iColumn int32, nValue int32) bool {
	r, _, _ := xAdTable_SetItemInt.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	nValue: 值.
func XAdTable_SetItemIntEx(hAdapter int, iItem int32, pName string, nValue int32) bool {
	r, _, _ := xAdTable_SetItemIntEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(nValue))
	return r != 0
}

// 数据适配器表_置项浮点值, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	fValue: 值.
func XAdTable_SetItemFloat(hAdapter int, iItem, iColumn int32, fValue float32) bool {
	r, _, _ := xAdTable_SetItemFloat.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fValue))
	return r != 0
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	fValue: 值.
func XAdTable_SetItemFloatEx(hAdapter int, iItem int32, pName string, fValue float32) bool {
	r, _, _ := xAdTable_SetItemFloatEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), common.Float32Ptr(fValue))
	return r != 0
}

// 数据适配器表_置项图片, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	hImage: 图片句柄.
func XAdTable_SetItemImage(hAdapter int, iItem, iColumn int32, hImage int) bool {
	r, _, _ := xAdTable_SetItemImage.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	hImage: 图片句柄.
func XAdTable_SetItemImageEx(hAdapter int, iItem int32, pName string, hImage int) bool {
	r, _, _ := xAdTable_SetItemImageEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器表_删除行.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 行索引.
func XAdTable_DeleteRow(hAdapter int, iRow int32) bool {
	r, _, _ := xAdTable_DeleteRow.Call(uintptr(hAdapter), uintptr(iRow))
	return r != 0
}

// 数据适配器表_删除行扩展, 删除多个行.
//
//	hAdapter: 数据适配器句柄.
//
//	iRow: 行起始索引.
//
//	nCount: 删除行数量.
func XAdTable_DeleteRowEx(hAdapter int, iRow int32, nCount int32) bool {
	r, _, _ := xAdTable_DeleteRowEx.Call(uintptr(hAdapter), uintptr(iRow), uintptr(nCount))
	return r != 0
}

// 数据适配器表_删除行全部, 删除所有行.
//
//	hAdapter: 数据适配器句柄.
func XAdTable_DeleteRowAll(hAdapter int) {
	xAdTable_DeleteRowAll.Call(uintptr(hAdapter))
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.
//
//	hAdapter: 数据适配器句柄.
func XAdTable_DeleteColumnAll(hAdapter int) {
	xAdTable_DeleteColumnAll.Call(uintptr(hAdapter))
}

// 数据适配器表_取列数量, 获取列数量.
//
//	hAdapter: 数据适配器句柄.
func XAdTable_GetCountColumn(hAdapter int) int32 {
	r, _, _ := xAdTable_GetCountColumn.Call(uintptr(hAdapter))
	return int32(r)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
func XAdTable_GetItemText(hAdapter int, iItem, iColumn int32) string {
	r, _, _ := xAdTable_GetItemText.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
func XAdTable_GetItemTextEx(hAdapter int, iItem int32, pName string) string {
	r, _, _ := xAdTable_GetItemTextEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
func XAdTable_GetItemImage(hAdapter int, iItem, iColumn int32) int {
	r, _, _ := xAdTable_GetItemImage.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
func XAdTable_GetItemImageEx(hAdapter int, iItem int32, pName string) int {
	r, _, _ := xAdTable_GetItemImageEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 数据适配器表_取项整数值, 获取项值.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	pOutValue: 接收返还值.
func XAdTable_GetItemInt(hAdapter int, iItem, iColumn int32, pOutValue *int32) bool {
	r, _, _ := xAdTable_GetItemInt.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	pOutValue: 接收返还值.
func XAdTable_GetItemIntEx(hAdapter int, iItem int32, pName string, pOutValue *int32) bool {
	r, _, _ := xAdTable_GetItemIntEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项浮点值, 获取项值.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	iColumn: 列索引.
//
//	pOutValue: 接收返还值.
func XAdTable_GetItemFloat(hAdapter int, iItem, iColumn int32, pOutValue *float32) bool {
	r, _, _ := xAdTable_GetItemFloat.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
//	hAdapter: 数据适配器句柄.
//
//	iItem: 项索引.
//
//	pName: 字段称.
//
//	pOutValue: 接收返还值.
func XAdTable_GetItemFloatEx(hAdapter int, iItem int32, pName string, pOutValue *float32) bool {
	r, _, _ := xAdTable_GetItemFloatEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取行数量, 获取行数量.
//
//	hAdapter: 数据适配器句柄.
func XAdTable_GetCountRow(hAdapter int) int32 {
	r, _, _ := xAdTable_GetCountRow.Call(uintptr(hAdapter))
	return int32(r)
}
