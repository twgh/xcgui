package xc

import "github.com/twgh/xcgui/common"

// 数据适配器MAP_创建, 创建数据适配器, 单列数据, 返回数据适配器句柄.
func XAdMap_Create() int {
	r, _, _ := xAdMap_Create.Call()
	return int(r)
}

// 数据适配器MAP_添加项文本, 增加数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
func XAdMap_AddItemText(hAdapter int, pName string, pValue string) bool {
	r, _, _ := xAdMap_AddItemText.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器MAP_添加项图片, 增加数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdMap_AddItemImage(hAdapter int, pName string, hImage int) bool {
	r, _, _ := xAdMap_AddItemImage.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器MAP_删除项, 删除数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdMap_DeleteItem(hAdapter int, pName string) bool {
	r, _, _ := xAdMap_DeleteItem.Call(uintptr(hAdapter), common.StrPtr(pName))
	return r != 0
}

// 数据适配器MAP_取项数量, 返回项数量.
//
// hAdapter: 数据适配器句柄.
func XAdMap_GetCount(hAdapter int) int {
	r, _, _ := xAdMap_GetCount.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器MAP_取项文本, 获取项内容, 如果内容为文本.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdMap_GetItemText(hAdapter int, pName string) string {
	r, _, _ := xAdMap_GetItemText.Call(uintptr(hAdapter), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器MAP_取项图片, 获取项内容, 如果内容为图片句柄, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdMap_GetItemImage(hAdapter int, pName string) int {
	r, _, _ := xAdMap_GetItemImage.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int(r)
}

// 数据适配器MAP_置项文本, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
func XAdMap_SetItemText(hAdapter int, pName string, pValue string) bool {
	r, _, _ := xAdMap_SetItemText.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器MAP_置项图片, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 值.
func XAdMap_SetItemImage(hAdapter int, pName string, hImage int) bool {
	r, _, _ := xAdMap_SetItemImage.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}
