package adapter

import (
	"github.com/twgh/xcgui/xc"
)

// AdapterListView 数据适配器-列表视元素.
type AdapterListView struct {
	adapter
}

// 数据适配器列表视_创建, 创建列表视元素数据适配器, 失败返回 nil.
func NewAdapterListView() *AdapterListView {
	return NewAdapterListViewByHandle(xc.XAdListView_Create())
}

// 从句柄创建对象, 失败返回 nil.
func NewAdapterListViewByHandle(handle int) *AdapterListView {
	if handle <= 0 {
		return nil
	}
	p := &AdapterListView{}
	p.SetHandle(handle)
	return p
}

// 数据适配器列表视_组添加列, 组操作, 添加数据列, 返回列索引.
//
// name: 字段称.
func (a *AdapterListView) Group_AddColumn(name string) int32 {
	return xc.XAdListView_Group_AddColumn(a.Handle, name)
}

// 数据适配器列表视_添加组文本, 组操作, 添加组, 数据默认填充到第一列, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemText(pValue string, iPos int32) int32 {
	return xc.XAdListView_Group_AddItemText(a.Handle, pValue, iPos)
}

// 数据适配器列表视_添加组文本扩展, 组操作, 添加组, 数据填充指定列, 返回组索引.
//
// name: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemTextEx(name string, pValue string, iPos int32) int32 {
	return xc.XAdListView_Group_AddItemTextEx(a.Handle, name, pValue, iPos)
}

// 数据适配器列表视_添加组图片, 组操作, 添加组, 数据默认填充第一列.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemImage(hImage int, iPos int32) int32 {
	return xc.XAdListView_Group_AddItemImage(a.Handle, hImage, iPos)
}

// 数据适配器列表视_添加组图片扩展, 组操作, 添加组, 数据填充指定列.
//
// name: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemImageEx(name string, hImage int, iPos int32) int32 {
	return xc.XAdListView_Group_AddItemImageEx(a.Handle, name, hImage, iPos)
}

// 数据适配器列表视_组设置文本, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) Group_SetText(iGroup, iColumn int32, pValue string) bool {
	return xc.XAdListView_Group_SetText(a.Handle, iGroup, iColumn, pValue)
}

// 数据适配器列表视_组设置文本扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// name: 字段名.
//
// pValue: 值.
func (a *AdapterListView) Group_SetTextEx(iGroup int32, name string, pValue string) bool {
	return xc.XAdListView_Group_SetTextEx(a.Handle, iGroup, name, pValue)
}

// 数据适配器列表视_组设置图片, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) Group_SetImage(iGroup, iColumn int32, hImage int) bool {
	return xc.XAdListView_Group_SetImage(a.Handle, iGroup, iColumn, hImage)
}

// 数据适配器列表视_组设置图片扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// name: 字段名.
//
// hImage: 图片句柄.
func (a *AdapterListView) Group_SetImageEx(iGroup int32, name string, hImage int) bool {
	return xc.XAdListView_Group_SetImageEx(a.Handle, iGroup, name, hImage)
}

// 数据适配器列表视_项添加列, 项操作, 添加列.
//
// name: 字段称.
func (a *AdapterListView) Item_AddColumn(name string) int32 {
	return xc.XAdListView_Item_AddColumn(a.Handle, name)
}

// 数据适配器列表视_取组数量, 组操作, 获取组数量.
func (a *AdapterListView) Group_GetCount() int32 {
	return xc.XAdListView_Group_GetCount(a.Handle)
}

// 数据适配器列表视_取组中项数量, 项操作, 获取指定组中项数量.
//
// iGroup: 组索引.
func (a *AdapterListView) Item_GetCount(iGroup int32) int32 {
	return xc.XAdListView_Item_GetCount(a.Handle, iGroup)
}

// 数据适配器列表视_添加项文本, 项操作, 添加项.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemText(iGroup int32, pValue string, iPos int32) int32 {
	return xc.XAdListView_Item_AddItemText(a.Handle, iGroup, pValue, iPos)
}

// 数据适配器列表视_添加项文本扩展, 项操作, 添加项, 数据填充指定列.
//
// iGroup: 组索引.
//
// name: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemTextEx(iGroup int32, name string, pValue string, iPos int32) int32 {
	return xc.XAdListView_Item_AddItemTextEx(a.Handle, iGroup, name, pValue, iPos)
}

// 数据适配器列表视_添加项图片, 项操作, 添加项.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemImage(iGroup int32, hImage int, iPos int32) int32 {
	return xc.XAdListView_Item_AddItemImage(a.Handle, iGroup, hImage, iPos)
}

// 数据适配器列表视_添加项图片扩展, 项操作, 添加项, 填充指定列数据.
//
// iGroup: 组索引.
//
// name: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemImageEx(iGroup int32, name string, hImage int, iPos int32) int32 {
	return xc.XAdListView_Item_AddItemImageEx(a.Handle, iGroup, name, hImage, iPos)
}

// 数据适配器列表视_组删除项, 删除组, 自动删除子项.
//
// iGroup: 组索引.
func (a *AdapterListView) Group_DeleteItem(iGroup int32) bool {
	return xc.XAdListView_Group_DeleteItem(a.Handle, iGroup)
}

// 数据适配器列表视_删除全部子项, 删除指定组的所有子项.
//
// iGroup: 组索引.
func (a *AdapterListView) Group_DeleteAllChildItem(iGroup int32) *AdapterListView {
	xc.XAdListView_Group_DeleteAllChildItem(a.Handle, iGroup)
	return a
}

// 数据适配器列表视_删除项, 删除项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (a *AdapterListView) Item_DeleteItem(iGroup, iItem int32) bool {
	return xc.XAdListView_Item_DeleteItem(a.Handle, iGroup, iItem)
}

// 数据适配器列表视_删除全部, 删除所有(组,项,列).
func (a *AdapterListView) DeleteAll() *AdapterListView {
	xc.XAdListView_DeleteAll(a.Handle)
	return a
}

// 数据适配器列表视_删除全部组, 删除所有的组.
func (a *AdapterListView) DeleteAllGroup() *AdapterListView {
	xc.XAdListView_DeleteAllGroup(a.Handle)
	return a
}

// 数据适配器列表视_删除全部项, 删除所有的项.
func (a *AdapterListView) DeleteAllItem() *AdapterListView {
	xc.XAdListView_DeleteAllItem(a.Handle)
	return a
}

// 数据适配器列表视_组删除列, 删除组的列.
//
// iColumn: 列索引.
func (a *AdapterListView) DeleteColumnGroup(iColumn int32) *AdapterListView {
	xc.XAdListView_DeleteColumnGroup(a.Handle, iColumn)
	return a
}

// 数据适配器列表视_项删除列, 删除项的列.
//
// iColumn: 列索引.
func (a *AdapterListView) DeleteColumnItem(iColumn int32) *AdapterListView {
	xc.XAdListView_DeleteColumnItem(a.Handle, iColumn)
	return a
}

// 数据适配器列表视_项获取文本扩展, 项操作, 获取项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// name: 字段称.
func (a *AdapterListView) Item_GetTextEx(iGroup, iItem int32, name string) string {
	return xc.XAdListView_Item_GetTextEx(a.Handle, iGroup, iItem, name)
}

// 数据适配器列表视_项获取图片扩展, 项操作, 获取项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// name: 字段称.
func (a *AdapterListView) Item_GetImageEx(iGroup, iItem int32, name string) int {
	return xc.XAdListView_Item_GetImageEx(a.Handle, iGroup, iItem, name)
}

// 数据适配器列表视_项置文本, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) Item_SetText(iGroup, iItem, iColumn int32, pValue string) bool {
	return xc.XAdListView_Item_SetText(a.Handle, iGroup, iItem, iColumn, pValue)
}

// 数据适配器列表视_项置文本扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// name: 字段称.
//
// pValue: 值.
func (a *AdapterListView) Item_SetTextEx(iGroup, iItem int32, name string, pValue string) bool {
	return xc.XAdListView_Item_SetTextEx(a.Handle, iGroup, iItem, name, pValue)
}

// 数据适配器列表视_项置图片, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) Item_SetImage(iGroup, iItem, iColumn int32, hImage int) bool {
	return xc.XAdListView_Item_SetImage(a.Handle, iGroup, iItem, iColumn, hImage)
}

// 数据适配器列表视_项置图片扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// name: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterListView) Item_SetImageEx(iGroup, iItem int32, name string, hImage int) bool {
	return xc.XAdListView_Item_SetImageEx(a.Handle, iGroup, iItem, name, hImage)
}

// 数据适配器列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Group_GetText(iGroup, iColumn int32) string {
	return xc.XAdListView_Group_GetText(a.Handle, iGroup, iColumn)
}

// 数据适配器列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// name: 字段名称.
func (a *AdapterListView) Group_GetTextEx(iGroup int32, name string) string {
	return xc.XAdListView_Group_GetTextEx(a.Handle, iGroup, name)
}

// 数据适配器列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Group_GetImage(iGroup, iColumn int32) int {
	return xc.XAdListView_Group_GetImage(a.Handle, iGroup, iColumn)
}

// 数据适配器列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// name: 字段名称.
func (a *AdapterListView) Group_GetImageEx(iGroup int32, name string) int {
	return xc.XAdListView_Group_GetImageEx(a.Handle, iGroup, name)
}

// 数据适配器列表视_项取文本. 项操作, 返回项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Item_GetText(iGroup, iItem, iColumn int32) string {
	return xc.XAdListView_Item_GetText(a.Handle, iGroup, iItem, iColumn)
}

// 数据适配器列表视_项取图片. 项操作, 返回项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// name: 字段名称.
func (a *AdapterListView) Item_GetImage(iGroup, iItem int32, name string) int {
	return xc.XAdListView_Item_GetImage(a.Handle, iGroup, iItem, name)
}
