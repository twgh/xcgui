package adapter

import (
	"github.com/twgh/xcgui/xc"
)

// AdapterTree 数据适配器-树元素.
type AdapterTree struct {
	adapter
}

// 数据适配器树_创建, 创建树元素数据适配器.
func NewAdapterTree() *AdapterTree {
	p := &AdapterTree{}
	p.SetHandle(xc.XAdTree_Create())
	return p
}

// 从句柄创建对象.
func NewAdapterTreeByHandle(handle int) *AdapterTree {
	p := &AdapterTree{}
	p.SetHandle(handle)
	return p
}

// 数据适配器树_添加列, 添加列, 返回索引值.
//
// pName: 字段称.
func (a *AdapterTree) AddColumn(pName string) int {
	return xc.XAdTree_AddColumn(a.Handle, pName)
}

// 数据适配器树_置列, 设置列, 返回列数量.
//
// pColName: 列名, 列名, 多个列名用逗号分开.
func (a *AdapterTree) SetColumn(pColName string) int {
	return xc.XAdTree_SetColumn(a.Handle, pColName)
}

// 数据适配器树_插入项文本, 插入项, 数据填充到第一列, 返回项ID值.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) InsertItemText(pValue string, nParentID int, insertID int) int {
	return xc.XAdTree_InsertItemText(a.Handle, pValue, nParentID, insertID)
}

// 数据适配器树_插入项文本扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// pName: 字段称.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) InsertItemTextEx(pName string, pValue string, nParentID int, insertID int) int {
	return xc.XAdTree_InsertItemTextEx(a.Handle, pName, pValue, nParentID, insertID)
}

// 数据适配器树_插入项图片, 插入项, 数据填充到第一列, 返回项ID值.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) InsertItemImage(hImage int, nParentID int, insertID int) int {
	return xc.XAdTree_InsertItemImage(a.Handle, hImage, nParentID, insertID)
}

// 数据适配器树_插入项图片扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) InsertItemImageEx(pName string, hImage int, nParentID int, insertID int) int {
	return xc.XAdTree_InsertItemImageEx(a.Handle, pName, hImage, nParentID, insertID)
}

// 数据适配器树_取项数量, 获取项数量.
func (a *AdapterTree) GetCount() int {
	return xc.XAdTree_GetCount(a.Handle)
}

// 数据适配器树_取列数量, 获取列数量.
func (a *AdapterTree) GetCountColumn() int {
	return xc.XAdTree_GetCountColumn(a.Handle)
}

// 数据适配器树_置项文本, 设置项数据.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterTree) SetItemText(nID int, iColumn int, pValue string) bool {
	return xc.XAdTree_SetItemText(a.Handle, nID, iColumn, pValue)
}

// 数据适配器树_置项文本扩展, 设置项文件内容.
//
// nID: 项ID.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTree) SetItemTextEx(nID int, pName string, pValue string) bool {
	return xc.XAdTree_SetItemTextEx(a.Handle, nID, pName, pValue)
}

// 数据适配器树_置项图片, 设置项数据.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterTree) SetItemImage(nID int, iColumn int, hImage int) bool {
	return xc.XAdTree_SetItemImage(a.Handle, nID, iColumn, hImage)
}

// 数据适配器树_置项图片扩展, 设置项内容.
//
// nID: 项ID.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTree) SetItemImageEx(nID int, pName string, hImage int) bool {
	return xc.XAdTree_SetItemImageEx(a.Handle, nID, pName, hImage)
}

// 数据适配器树_取项文本, 获取项文本内容.
//
// nID: 项ID.
//
// iColumn: 列索引.
func (a *AdapterTree) GetItemText(nID int, iColumn int) string {
	return xc.XAdTree_GetItemText(a.Handle, nID, iColumn)
}

// 数据适配器树_取项文本扩展, 获取项文本内容.
//
// nID: 项ID.
//
// pName: 字段称.
func (a *AdapterTree) GetItemTextEx(nID int, pName string) string {
	return xc.XAdTree_GetItemTextEx(a.Handle, nID, pName)
}

// 数据适配器树_取项图片, 获取项内容, 返回图片句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
func (a *AdapterTree) GetItemImage(nID int, iColumn int) int {
	return xc.XAdTree_GetItemImage(a.Handle, nID, iColumn)
}

// 数据适配器树_取项图片扩展, 获取项内容, 返回图片句柄.
//
// nID: 项ID.
//
// pName: 字段称.
func (a *AdapterTree) GetItemImageEx(nID int, pName string) int {
	return xc.XAdTree_GetItemImageEx(a.Handle, nID, pName)
}

// 数据适配器树_删除项, 删除项.
//
// nID: 项ID.
func (a *AdapterTree) DeleteItem(nID int) bool {
	return xc.XAdTree_DeleteItem(a.Handle, nID)
}

// 数据适配器树_删除项全部, 删除所有项.
func (a *AdapterTree) DeleteItemAll() int {
	return xc.XAdTree_DeleteItemAll(a.Handle)
}

// 数据适配器树_删除列全部, 删除所有列, 并且清空数据.
func (a *AdapterTree) DeleteColumnAll() int {
	return xc.XAdTree_DeleteColumnAll(a.Handle)
}
