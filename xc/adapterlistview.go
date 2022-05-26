package xc

import "github.com/twgh/xcgui/common"

// 数据适配器列表视_创建, 创建列表视元素数据适配器, 返回数据适配器句柄.
func XAdListView_Create() int {
	r, _, _ := xAdListView_Create.Call()
	return int(r)
}

// 数据适配器列表视_组添加列, 组操作, 添加数据列, 返回列索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdListView_Group_AddColumn(hAdapter int, pName string) int {
	r, _, _ := xAdListView_Group_AddColumn.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int(r)
}

// 数据适配器列表视_添加组文本, 组操作, 添加组, 数据默认填充到第一列, 返回组索引.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Group_AddItemText(hAdapter int, pValue string, iPos int) int {
	r, _, _ := xAdListView_Group_AddItemText.Call(uintptr(hAdapter), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加组文本扩展, 组操作, 添加组, 数据填充指定列, 返回组索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Group_AddItemTextEx(hAdapter int, pName string, pValue string, iPos int) int {
	r, _, _ := xAdListView_Group_AddItemTextEx.Call(uintptr(hAdapter), common.StrPtr(pName), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加组图片, 组操作, 添加组, 数据默认填充第一列.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Group_AddItemImage(hAdapter int, hImage int, iPos int) int {
	r, _, _ := xAdListView_Group_AddItemImage.Call(uintptr(hAdapter), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加组图片扩展, 组操作, 添加组, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Group_AddItemImageEx(hAdapter int, pName string, hImage int, iPos int) int {
	r, _, _ := xAdListView_Group_AddItemImageEx.Call(uintptr(hAdapter), common.StrPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_组设置文本, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func XAdListView_Group_SetText(hAdapter int, iGroup int, iColumn int, pValue string) bool {
	r, _, _ := xAdListView_Group_SetText.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器列表视_组设置文本扩展, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func XAdListView_Group_SetTextEx(hAdapter int, iGroup int, pName string, pValue string) bool {
	r, _, _ := xAdListView_Group_SetTextEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器列表视_组设置图片, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XAdListView_Group_SetImage(hAdapter int, iGroup int, iColumn int, hImage int) bool {
	r, _, _ := xAdListView_Group_SetImage.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 数据适配器列表视_组设置图片扩展, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func XAdListView_Group_SetImageEx(hAdapter int, iGroup int, pName string, hImage int) bool {
	r, _, _ := xAdListView_Group_SetImageEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器列表视_项添加列, 项操作, 添加列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func XAdListView_Item_AddColumn(hAdapter int, pName string) int {
	r, _, _ := xAdListView_Item_AddColumn.Call(uintptr(hAdapter), common.StrPtr(pName))
	return int(r)
}

// 数据适配器列表视_取组数量, 组操作, 获取组数量.
//
// hAdapter: 数据适配器句柄.
func XAdListView_Group_GetCount(hAdapter int) int {
	r, _, _ := xAdListView_Group_GetCount.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器列表视_取组中项数量, 项操作, 获取指定组中项数量.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func XAdListView_Item_GetCount(hAdapter int, iGroup int) int {
	r, _, _ := xAdListView_Item_GetCount.Call(uintptr(hAdapter), uintptr(iGroup))
	return int(r)
}

// 数据适配器列表视_添加项文本, 项操作, 添加项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Item_AddItemText(hAdapter int, iGroup int, pValue string, iPos int) int {
	r, _, _ := xAdListView_Item_AddItemText.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加项文本扩展, 项操作, 添加项, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Item_AddItemTextEx(hAdapter int, iGroup int, pName string, pValue string, iPos int) int {
	r, _, _ := xAdListView_Item_AddItemTextEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加项图片, 项操作, 添加项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Item_AddItemImage(hAdapter int, iGroup int, hImage int, iPos int) int {
	r, _, _ := xAdListView_Item_AddItemImage.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_添加项图片扩展, 项操作, 添加项, 填充指定列数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func XAdListView_Item_AddItemImageEx(hAdapter int, iGroup int, pName string, hImage int, iPos int) int {
	r, _, _ := xAdListView_Item_AddItemImageEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 数据适配器列表视_组删除项, 删除组, 自动删除子项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func XAdListView_Group_DeleteItem(hAdapter int, iGroup int) bool {
	r, _, _ := xAdListView_Group_DeleteItem.Call(uintptr(hAdapter), uintptr(iGroup))
	return r != 0
}

// 数据适配器列表视_删除全部子项, 删除指定组的所有子项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func XAdListView_Group_DeleteAllChildItem(hAdapter int, iGroup int) int {
	r, _, _ := xAdListView_Group_DeleteAllChildItem.Call(uintptr(hAdapter), uintptr(iGroup))
	return int(r)
}

// 数据适配器列表视_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XAdListView_Item_DeleteItem(hAdapter int, iGroup int, iItem int) bool {
	r, _, _ := xAdListView_Item_DeleteItem.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem))
	return r != 0
}

// 数据适配器列表视_删除全部, 删除所有(组,项,列).
//
// hAdapter: 数据适配器句柄.
func XAdListView_DeleteAll(hAdapter int) int {
	r, _, _ := xAdListView_DeleteAll.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器列表视_删除全部组, 删除所有的组.
//
// hAdapter: 数据适配器句柄.
func XAdListView_DeleteAllGroup(hAdapter int) int {
	r, _, _ := xAdListView_DeleteAllGroup.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器列表视_删除全部项, 删除所有的项.
//
// hAdapter: 数据适配器句柄.
func XAdListView_DeleteAllItem(hAdapter int) int {
	r, _, _ := xAdListView_DeleteAllItem.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器列表视_组删除列, 删除组的列.
//
// hAdapter: 数据适配器句柄.
//
// iColumn: 列索引.
func XAdListView_DeleteColumnGroup(hAdapter int, iColumn int) int {
	r, _, _ := xAdListView_DeleteColumnGroup.Call(uintptr(hAdapter), uintptr(iColumn))
	return int(r)
}

// 数据适配器列表视_项删除列, 删除项的列.
//
// hAdapter: 数据适配器句柄.
//
// iColumn: 列索引.
func XAdListView_DeleteColumnItem(hAdapter int, iColumn int) int {
	r, _, _ := xAdListView_DeleteColumnItem.Call(uintptr(hAdapter), uintptr(iColumn))
	return int(r)
}

// 数据适配器列表视_项获取文本扩展, 项操作, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func XAdListView_Item_GetTextEx(hAdapter int, iGroup int, iItem int, pName string) string {
	r, _, _ := xAdListView_Item_GetTextEx.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器列表视_项获取图片扩展, 项操作, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func XAdListView_Item_GetImageEx(hAdapter int, iGroup int, iItem int, pName string) int {
	r, _, _ := xAdListView_Item_GetImageEx.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 数据适配器列表视_项置文本, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func XAdListView_Item_SetText(hAdapter int, iGroup int, iItem int, iColumn int, pValue string) bool {
	r, _, _ := xAdListView_Item_SetText.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器列表视_项置文本扩展, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func XAdListView_Item_SetTextEx(hAdapter int, iGroup int, iItem int, pName string, pValue string) bool {
	r, _, _ := xAdListView_Item_SetTextEx.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 数据适配器列表视_项置图片, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XAdListView_Item_SetImage(hAdapter int, iGroup int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xAdListView_Item_SetImage.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 数据适配器列表视_项置图片扩展, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func XAdListView_Item_SetImageEx(hAdapter int, iGroup int, iItem int, pName string, hImage int) bool {
	r, _, _ := xAdListView_Item_SetImageEx.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 数据适配器列表视_组取文本, 返回文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func XAdListView_Group_GetText(hAdapter int, iGroup int, iColumn int) string {
	r, _, _ := xAdListView_Group_GetText.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 数据适配器列表视_组取文本扩展, 返回文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func XAdListView_Group_GetTextEx(hAdapter int, iGroup int, pName string) string {
	r, _, _ := xAdListView_Group_GetTextEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 数据适配器列表视_组取图片, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func XAdListView_Group_GetImage(hAdapter int, iGroup int, iColumn int) int {
	r, _, _ := xAdListView_Group_GetImage.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iColumn))
	return int(r)
}

// 数据适配器列表视_组取图片扩展, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func XAdListView_Group_GetImageEx(hAdapter int, iGroup int, pName string) int {
	r, _, _ := xAdListView_Group_GetImageEx.Call(uintptr(hAdapter), uintptr(iGroup), common.StrPtr(pName))
	return int(r)
}

// 数据适配器列表视_项取文本. 项操作, 返回项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XAdListView_Item_GetText(hAdapter int, iGroup int, iItem int, iColumn int) string {
	r, _, _ := xAdListView_Item_GetText.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 数据适配器列表视_项取图片. 项操作, 返回项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XAdListView_Item_GetImage(hAdapter int, iGroup int, iItem int, iColumn int) int {
	r, _, _ := xAdListView_Item_GetImage.Call(uintptr(hAdapter), uintptr(iGroup), uintptr(iItem), uintptr(iColumn))
	return int(r)
}
