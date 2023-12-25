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
// hAdapter: 数据适配器句柄.
//
// iColumn: 要排序的列索引。.
//
// bAscending: 是否按照升序方式排序.
func XAdTable_Sort(hAdapter int, iColumn int, bAscending bool) int {
	r, _, _ := xAdTable_Sort.Call(uintptr(hAdapter), uintptr(iColumn), common.BoolPtr(bAscending))
	return int(r)
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XAdTable_GetItemDataType(hAdapter int, iItem int, iColumn int) xcc.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataType.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return xcc.Adapter_Date_Type_(r)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
func XAdTable_GetItemDataTypeEx(hAdapter int, iItem int, pName string) xcc.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataTypeEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return xcc.Adapter_Date_Type_(r)
}

// 数据适配器表_添加列, 添加数据列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdTable_AddColumn(hAdapter int, pName string) int {
	r, _, _ := xAdTable_AddColumn.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int(r)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
// hAdapter: 数据适配器句柄.
//
// pColName: 列名, 多个列名用逗号分开.
func XAdTable_SetColumn(hAdapter int, pColName string) int {
	r, _, _ := xAdTable_SetColumn.Call(uintptr(hAdapter), common.StrPtr(pColName))
	return int(r)
}

// 数据适配器表_添加项文本, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.
func XAdTable_AddItemText(hAdapter int, pValue string) int {
	r, _, _ := xAdTable_AddItemText.Call(uintptr(hAdapter), common.StrPtr(pValue))
	return int(r)
}

// 数据适配器表_添加项文本扩展, 添加数据项, 填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
func XAdTable_AddItemTextEx(hAdapter int, pName string, pValue string) int {
	r, _, _ := xAdTable_AddItemTextEx.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue))
	return int(r)
}

// 数据适配器表_添加项图片, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.
func XAdTable_AddItemImage(hAdapter int, hImage int) int {
	r, _, _ := xAdTable_AddItemImage.Call(uintptr(hAdapter), uintptr(hImage))
	return int(r)
}

// 数据适配器表_添加项图片扩展, 添加数据项, 并填充指定列数据.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdTable_AddItemImageEx(hAdapter int, pName string, hImage int) int {
	r, _, _ := xAdTable_AddItemImageEx.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 数据适配器表_插入项文本, 插入数据项, 填充第一列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pValue: 值.
func XAdTable_InsertItemText(hAdapter int, iItem int, pValue string) int {
	r, _, _ := xAdTable_InsertItemText.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pValue))
	return int(r)
}

// 数据适配器表_插入项文本扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// pValue: 值.
func XAdTable_InsertItemTextEx(hAdapter int, iItem int, pName string, pValue string) int {
	r, _, _ := xAdTable_InsertItemTextEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return int(r)
}

// 数据适配器表_插入项图片, 插入数据项, 填充第一列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// hImage: 图片句柄.
func XAdTable_InsertItemImage(hAdapter int, iItem int, hImage int) int {
	r, _, _ := xAdTable_InsertItemImage.Call(uintptr(hAdapter), uintptr(iItem), uintptr(hImage))
	return int(r)
}

// 数据适配器表_插入项图片扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdTable_InsertItemImageEx(hAdapter int, iItem int, pName string, hImage int) int {
	r, _, _ := xAdTable_InsertItemImageEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 数据适配器表_置项文本, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func XAdTable_SetItemText(hAdapter int, iItem int, iColumn int, pValue string) bool {
	r, _, _ := xAdTable_SetItemText.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func XAdTable_SetItemTextEx(hAdapter int, iItem int, pName string, pValue string) bool {
	r, _, _ := xAdTable_SetItemTextEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器表_置项整数值, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 值.
func XAdTable_SetItemInt(hAdapter int, iItem int, iColumn int, nValue int32) bool {
	r, _, _ := xAdTable_SetItemInt.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// nValue: 值.
func XAdTable_SetItemIntEx(hAdapter int, iItem int, pName string, nValue int32) bool {
	r, _, _ := xAdTable_SetItemIntEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(nValue))
	return r != 0
}

// 数据适配器表_置项浮点值, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fValue: 值.
func XAdTable_SetItemFloat(hAdapter int, iItem int, iColumn int, fValue float32) bool {
	r, _, _ := xAdTable_SetItemFloat.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fValue))
	return r != 0
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// fValue: 值.
func XAdTable_SetItemFloatEx(hAdapter int, iItem int, pName string, fValue float32) bool {
	r, _, _ := xAdTable_SetItemFloatEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), common.Float32Ptr(fValue))
	return r != 0
}

// 数据适配器表_置项图片, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XAdTable_SetItemImage(hAdapter int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xAdTable_SetItemImage.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdTable_SetItemImageEx(hAdapter int, iItem int, pName string, hImage int) bool {
	r, _, _ := xAdTable_SetItemImageEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器表_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
func XAdTable_DeleteItem(hAdapter int, iItem int) bool {
	r, _, _ := xAdTable_DeleteItem.Call(uintptr(hAdapter), uintptr(iItem))
	return r != 0
}

// 数据适配器表_删除项扩展, 删除多个项.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项起始索引.
//
// nCount: 删除项数量.
func XAdTable_DeleteItemEx(hAdapter int, iItem int, nCount int) bool {
	r, _, _ := xAdTable_DeleteItemEx.Call(uintptr(hAdapter), uintptr(iItem), uintptr(nCount))
	return r != 0
}

// 数据适配器表_删除项全部, 删除所有项.
//
// hAdapter: 数据适配器句柄.
func XAdTable_DeleteItemAll(hAdapter int) int {
	r, _, _ := xAdTable_DeleteItemAll.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.
//
// hAdapter: 数据适配器句柄.
func XAdTable_DeleteColumnAll(hAdapter int) int {
	r, _, _ := xAdTable_DeleteColumnAll.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器表_取项数量, 获取项数量.
//
// hAdapter: 数据适配器句柄.
func XAdTable_GetCount(hAdapter int) int {
	r, _, _ := xAdTable_GetCount.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器表_取列数量, 获取列数量.
//
// hAdapter: 数据适配器句柄.
func XAdTable_GetCountColumn(hAdapter int) int {
	r, _, _ := xAdTable_GetCountColumn.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XAdTable_GetItemText(hAdapter int, iItem int, iColumn int) string {
	r, _, _ := xAdTable_GetItemText.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
func XAdTable_GetItemTextEx(hAdapter int, iItem int, pName string) string {
	r, _, _ := xAdTable_GetItemTextEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XAdTable_GetItemImage(hAdapter int, iItem int, iColumn int) int {
	r, _, _ := xAdTable_GetItemImage.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
func XAdTable_GetItemImageEx(hAdapter int, iItem int, pName string) int {
	r, _, _ := xAdTable_GetItemImageEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 数据适配器表_取项整数值, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func XAdTable_GetItemInt(hAdapter int, iItem int, iColumn int, pOutValue *int32) bool {
	r, _, _ := xAdTable_GetItemInt.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func XAdTable_GetItemIntEx(hAdapter int, iItem int, pName string, pOutValue *int32) bool {
	r, _, _ := xAdTable_GetItemIntEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项浮点值, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func XAdTable_GetItemFloat(hAdapter int, iItem int, iColumn int, pOutValue *float32) bool {
	r, _, _ := xAdTable_GetItemFloat.Call(uintptr(hAdapter), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func XAdTable_GetItemFloatEx(hAdapter int, iItem int, pName string, pOutValue *float32) bool {
	r, _, _ := xAdTable_GetItemFloatEx.Call(uintptr(hAdapter), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}
