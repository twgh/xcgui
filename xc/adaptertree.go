package xc

import "github.com/twgh/xcgui/common"

// 数据适配器树_创建, 创建树元素数据适配器, 返回数据适配器句柄.
func XAdTree_Create() int {
	r, _, _ := xAdTree_Create.Call()
	return int(r)
}

// 数据适配器树_添加列, 添加列, 返回索引值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdTree_AddColumn(hAdapter int, pName string) int {
	r, _, _ := xAdTree_AddColumn.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int(r)
}

// 数据适配器树_置列, 设置列, 返回列数量.
//
// hAdapter: 数据适配器句柄.
//
// pColName: 列名, 列名, 多个列名用逗号分开.
func XAdTree_SetColumn(hAdapter int, pColName string) int {
	r, _, _ := xAdTree_SetColumn.Call(uintptr(hAdapter), common.StrPtr(pColName))
	return int(r)
}

// 数据适配器树_插入项文本, 插入项, 数据填充到第一列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func XAdTree_InsertItemText(hAdapter int, pValue string, nParentID int, insertID int) int {
	r, _, _ := xAdTree_InsertItemText.Call(uintptr(hAdapter), common.StrPtr(pValue), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 数据适配器树_插入项文本扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func XAdTree_InsertItemTextEx(hAdapter int, pName string, pValue string, nParentID int, insertID int) int {
	r, _, _ := xAdTree_InsertItemTextEx.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 数据适配器树_插入项图片, 插入项, 数据填充到第一列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func XAdTree_InsertItemImage(hAdapter int, hImage int, nParentID int, insertID int) int {
	r, _, _ := xAdTree_InsertItemImage.Call(uintptr(hAdapter), uintptr(hImage), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 数据适配器树_插入项图片扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func XAdTree_InsertItemImageEx(hAdapter int, pName string, hImage int, nParentID int, insertID int) int {
	r, _, _ := xAdTree_InsertItemImageEx.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 数据适配器树_取项数量, 获取项数量.
//
// hAdapter: 数据适配器句柄.
func XAdTree_GetCount(hAdapter int) int {
	r, _, _ := xAdTree_GetCount.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器树_取列数量, 获取列数量.
//
// hAdapter: 数据适配器句柄.
func XAdTree_GetCountColumn(hAdapter int) int {
	r, _, _ := xAdTree_GetCountColumn.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器树_置项文本, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// pValue: 值.
func XAdTree_SetItemText(hAdapter int, nID int, iColumn int, pValue string) bool {
	r, _, _ := xAdTree_SetItemText.Call(uintptr(hAdapter), uintptr(nID), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器树_置项文本扩展, 设置项文件内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
//
// pValue: 值.
func XAdTree_SetItemTextEx(hAdapter int, nID int, pName string, pValue string) bool {
	r, _, _ := xAdTree_SetItemTextEx.Call(uintptr(hAdapter), uintptr(nID), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器树_置项图片, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XAdTree_SetItemImage(hAdapter int, nID int, iColumn int, hImage int) bool {
	r, _, _ := xAdTree_SetItemImage.Call(uintptr(hAdapter), uintptr(nID), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 数据适配器树_置项图片扩展, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdTree_SetItemImageEx(hAdapter int, nID int, pName string, hImage int) bool {
	r, _, _ := xAdTree_SetItemImageEx.Call(uintptr(hAdapter), uintptr(nID), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器树_取项文本, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
func XAdTree_GetItemText(hAdapter int, nID int, iColumn int) string {
	r, _, _ := xAdTree_GetItemText.Call(uintptr(hAdapter), uintptr(nID), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 数据适配器树_取项文本扩展, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
func XAdTree_GetItemTextEx(hAdapter int, nID int, pName string) string {
	r, _, _ := xAdTree_GetItemTextEx.Call(uintptr(hAdapter), uintptr(nID), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器树_取项图片, 获取项内容, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
func XAdTree_GetItemImage(hAdapter int, nID int, iColumn int) int {
	r, _, _ := xAdTree_GetItemImage.Call(uintptr(hAdapter), uintptr(nID), uintptr(iColumn))
	return int(r)
}

// 数据适配器树_取项图片扩展, 获取项内容, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
func XAdTree_GetItemImageEx(hAdapter int, nID int, pName string) int {
	r, _, _ := xAdTree_GetItemImageEx.Call(uintptr(hAdapter), uintptr(nID), common.StrPtr(pName))
	return int(r)
}

// 数据适配器树_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
func XAdTree_DeleteItem(hAdapter int, nID int) bool {
	r, _, _ := xAdTree_DeleteItem.Call(uintptr(hAdapter), uintptr(nID))
	return r != 0
}

// 数据适配器树_删除项全部, 删除所有项.
//
// hAdapter: 数据适配器句柄.
func XAdTree_DeleteItemAll(hAdapter int) int {
	r, _, _ := xAdTree_DeleteItemAll.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器树_删除列全部, 删除所有列, 并且清空数据.
//
// hAdapter: 数据适配器句柄.
func XAdTree_DeleteColumnAll(hAdapter int) int {
	r, _, _ := xAdTree_DeleteColumnAll.Call(uintptr(hAdapter))
	return int(r)
}
