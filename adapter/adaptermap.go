package adapter

import (
	xc "github.com/twgh/xcgui"
)

type AdapterMap struct {
	Adapter

	HAdapter int
}

// 数据适配器MAP_创建, 创建数据适配器, 单列数据, 返回数据适配器句柄
func NewAdapterMap() *AdapterMap {
	p := &AdapterMap{
		HAdapter: xc.XAdMap_Create(),
	}
	p.HAdapter_ = p.HAdapter
	return p
}

// 数据适配器MAP_添加项文本, 增加数据项
// pName: 字段称.
// pValue: 值.
func (a *AdapterMap) AddItemText(pName string, pValue string) bool {
	return xc.XAdMap_AddItemText(a.HAdapter, pName, pValue)
}

// 数据适配器MAP_添加项图片, 增加数据项
// pName: 字段称.
// hImage: 图片句柄.
func (a *AdapterMap) AddItemImage(pName string, hImage int) bool {
	return xc.XAdMap_AddItemImage(a.HAdapter, pName, hImage)
}

// 数据适配器MAP_删除项, 删除数据项
// pName: 字段称.
func (a *AdapterMap) DeleteItem(pName string) bool {
	return xc.XAdMap_DeleteItem(a.HAdapter, pName)
}

// 数据适配器MAP_取项数量, 返回项数量
func (a *AdapterMap) GetCount() int {
	return xc.XAdMap_GetCount(a.HAdapter)
}

// 数据适配器MAP_取项文本, 获取项内容, 如果内容为文本
// pName: 字段称.
func (a *AdapterMap) GetItemText(pName string) string {
	return xc.XAdMap_GetItemText(a.HAdapter, pName)
}

// 数据适配器MAP_取项图片, 获取项内容, 如果内容为图片句柄, 返回图片句柄
// pName: 字段称.
func (a *AdapterMap) GetItemImage(pName string) int {
	return xc.XAdMap_GetItemImage(a.HAdapter, pName)
}

// 数据适配器MAP_置项文本, 设置项内容
// pName: 字段称.
// pValue: 值.
func (a *AdapterMap) SetItemText(pName string, pValue string) bool {
	return xc.XAdMap_SetItemText(a.HAdapter, pName, pValue)
}

// 数据适配器MAP_置项图片, 设置项内容
// pName: 字段称.
// hImage: 值.
func (a *AdapterMap) SetItemImage(pName string, hImage int) bool {
	return xc.XAdMap_SetItemImage(a.HAdapter, pName, hImage)
}
