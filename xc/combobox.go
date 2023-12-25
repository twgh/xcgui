package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 组合框_创建, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func XComboBox_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xComboBox_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 组合框_置选择项.
//
// hEle: 元素句柄.
//
// iIndex: 项索引.
func XComboBox_SetSelItem(hEle int, iIndex int) bool {
	r, _, _ := xComboBox_SetSelItem.Call(uintptr(hEle), uintptr(iIndex))
	return r != 0
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func XComboBox_CreateAdapter(hEle int) int {
	r, _, _ := xComboBox_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 组合框_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 适配器句柄.
func XComboBox_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xComboBox_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
//
// hEle: 元素句柄.
func XComboBox_GetAdapter(hEle int) int {
	r, _, _ := xComboBox_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 组合框_置绑定名称.
//
// hEle: 元素句柄.
//
// pName: 字段名.
func XComboBox_SetBindName(hEle int, pName string) int {
	r, _, _ := xComboBox_SetBindName.Call(uintptr(hEle), common.StrPtr(pName))
	return int(r)
}

// 组合框_取下拉按钮坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XComboBox_GetButtonRect(hEle int, pRect *RECT) int {
	r, _, _ := xComboBox_GetButtonRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 组合框_置下拉按钮大小.
//
// hEle: 元素句柄.
//
// size: 大小.
func XComboBox_SetButtonSize(hEle int, size int) int {
	r, _, _ := xComboBox_SetButtonSize.Call(uintptr(hEle), uintptr(size))
	return int(r)
}

// 组合框_置下拉列表高度.
//
// hEle: 元素句柄.
//
// height: 高度, -1自动计算高度.
func XComboBox_SetDropHeight(hEle int, height int) int {
	r, _, _ := xComboBox_SetDropHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 组合框_取下拉列表高度.
//
// hEle: 元素句柄.
func XComboBox_GetDropHeight(hEle int) int {
	r, _, _ := xComboBox_GetDropHeight.Call(uintptr(hEle))
	return int(r)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 项模板文件.
func XComboBox_SetItemTemplateXML(hEle int, pXmlFile string) int {
	r, _, _ := xComboBox_SetItemTemplateXML.Call(uintptr(hEle), common.StrPtr(pXmlFile))
	return int(r)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func XComboBox_SetItemTemplateXMLFromString(hEle int, pStringXML string) int {
	r, _, _ := xComboBox_SetItemTemplateXMLFromString.Call(uintptr(hEle), XC_wtoa(pStringXML))
	return int(r)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// hEle: 元素句柄.
//
// bEnable: 是否绘制.
func XComboBox_EnableDrawButton(hEle int, bEnable bool) int {
	r, _, _ := xComboBox_EnableDrawButton.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// hEle: 元素句柄.
//
// bEdit: TRUE可编辑.
func XComboBox_EnableEdit(hEle int, bEdit bool) int {
	r, _, _ := xComboBox_EnableEdit.Call(uintptr(hEle), common.BoolPtr(bEdit))
	return int(r)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XComboBox_EnableDropHeightFixed(hEle int, bEnable bool) int {
	r, _, _ := xComboBox_EnableDropHeightFixed.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
//
// hEle: 元素句柄.
func XComboBox_GetSelItem(hEle int) int {
	r, _, _ := xComboBox_GetSelItem.Call(uintptr(hEle))
	return int(r)
}

// 组合框_取状态, 返回: ComboBox_State_.
//
// hEle: 元素句柄.
func XComboBox_GetState(hEle int) xcc.ComboBox_State_ {
	r, _, _ := xComboBox_GetState.Call(uintptr(hEle))
	return xcc.ComboBox_State_(r)
}

// 组合框_添加项文本, 返回项索引.
//
// hEle: 元素句柄.
//
// pText:.
func XComboBox_AddItemText(hEle int, pText string) int {
	r, _, _ := xComboBox_AddItemText.Call(uintptr(hEle), common.StrPtr(pText))
	return int(r)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
//
// pText: 文本.
func XComboBox_AddItemTextEx(hEle int, pName string, pText string) int {
	r, _, _ := xComboBox_AddItemTextEx.Call(uintptr(hEle), common.StrPtr(pName), common.StrPtr(pText))
	return int(r)
}

// 组合框_添加项图片, 返回项索引.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XComboBox_AddItemImage(hEle int, hImage int) int {
	r, _, _ := xComboBox_AddItemImage.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func XComboBox_AddItemImageEx(hEle int, pName string, hImage int) int {
	r, _, _ := xComboBox_AddItemImageEx.Call(uintptr(hEle), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 组合框_插入项文本, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pText: 文本.
func XComboBox_InsertItemText(hEle int, iItem int, pText string) int {
	r, _, _ := xComboBox_InsertItemText.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pText))
	return int(r)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func XComboBox_InsertItemTextEx(hEle int, iItem int, pName string, pText string) int {
	r, _, _ := xComboBox_InsertItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pText))
	return int(r)
}

// 组合框_插入项图片, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func XComboBox_InsertItemImage(hEle int, iItem int, hImage int) int {
	r, _, _ := xComboBox_InsertItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(hImage))
	return int(r)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func XComboBox_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {
	r, _, _ := xComboBox_InsertItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 组合框_置项文本.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pText: 文本.
func XComboBox_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {
	r, _, _ := xComboBox_SetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.StrPtr(pText))
	return r != 0
}

// 组合框_置项文本扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func XComboBox_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {
	r, _, _ := xComboBox_SetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pText))
	return r != 0
}

// 组合框_置项图片.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XComboBox_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xComboBox_SetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 组合框_置项图片扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func XComboBox_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {
	r, _, _ := xComboBox_SetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 组合框_置项整数值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func XComboBox_SetItemInt(hEle int, iItem int, iColumn int, nValue int32) bool {
	r, _, _ := xComboBox_SetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 组合框_置项指数值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// nValue: 整数值.
func XComboBox_SetItemIntEx(hEle int, iItem int, pName string, nValue int32) bool {
	r, _, _ := xComboBox_SetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(nValue))
	return r != 0
}

// 组合框_置项浮点值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func XComboBox_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {
	r, _, _ := xComboBox_SetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fFloat))
	return r != 0
}

// 组合框_置项浮点值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// fFloat: 浮点数.
func XComboBox_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {
	r, _, _ := xComboBox_SetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.Float32Ptr(fFloat))
	return r != 0
}

// 组合框_取项文本.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XComboBox_GetItemText(hEle int, iItem int32, iColumn int32) string {
	r, _, _ := xComboBox_GetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 组合框_取项文本扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
func XComboBox_GetItemTextEx(hEle int, iItem int, pName string) string {
	r, _, _ := xComboBox_GetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 组合框_取项图片.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XComboBox_GetItemImage(hEle int, iItem int, iColumn int) int {
	r, _, _ := xComboBox_GetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 组合框_取项图片扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
func XComboBox_GetItemImageEx(hEle int, iItem int, pName string) int {
	r, _, _ := xComboBox_GetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 组合框_取项整数值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func XComboBox_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {
	r, _, _ := xComboBox_GetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 组合框_取项整数值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回整数值.
func XComboBox_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {
	r, _, _ := xComboBox_GetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 组合框_取项浮点值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func XComboBox_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {
	r, _, _ := xComboBox_GetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 组合框_取项浮点值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回浮点值.
func XComboBox_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {
	r, _, _ := xComboBox_GetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 组合框_删除项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XComboBox_DeleteItem(hEle int, iItem int) bool {
	r, _, _ := xComboBox_DeleteItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 组合框_删除项扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func XComboBox_DeleteItemEx(hEle int, iItem int, nCount int) bool {
	r, _, _ := xComboBox_DeleteItemEx.Call(uintptr(hEle), uintptr(iItem), uintptr(nCount))
	return r != 0
}

// 组合框_删除项全部.
//
// hEle: 元素句柄.
func XComboBox_DeleteItemAll(hEle int) int {
	r, _, _ := xComboBox_DeleteItemAll.Call(uintptr(hEle))
	return int(r)
}

// 组合框_删除列全部.
//
// hEle: 元素句柄.
func XComboBox_DeleteColumnAll(hEle int) int {
	r, _, _ := xComboBox_DeleteColumnAll.Call(uintptr(hEle))
	return int(r)
}

// 组合框_取项数量.
//
// hEle:.
func XComboBox_GetCount(hEle int) int {
	r, _, _ := xComboBox_GetCount.Call(uintptr(hEle))
	return int(r)
}

// 组合框_取列数量.
//
// hEle: 元素句柄.
func XComboBox_GetCountColumn(hEle int) int {
	r, _, _ := xComboBox_GetCountColumn.Call(uintptr(hEle))
	return int(r)
}

// 组合框_弹出下拉列表.
//
// hEle: 元素句柄.
func XComboBox_PopupDropList(hEle int) int {
	r, _, _ := xComboBox_PopupDropList.Call(uintptr(hEle))
	return int(r)
}

// 组合框_设置项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func XComboBox_SetItemTemplate(hEle, hTemp int) int {
	r, _, _ := xComboBox_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return int(r)
}

// 组合框_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func XComboBox_SetItemTemplateXMLFromMem(hEle int, data []byte) bool {
	r, _, _ := xComboBox_SetItemTemplateXMLFromMem.Call(uintptr(hEle), common.ByteSliceDataPtr(&data), uintptr(len(data)))
	return r != 0
}

// 组合框_置项模板从资源ZIP.
//
// hEle: 元素句柄.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func XComboBox_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xComboBox_SetItemTemplateXMLFromZipRes.Call(uintptr(hEle), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}

// 组合框_取项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func XComboBox_GetItemTemplate(hEle int) int {
	r, _, _ := xComboBox_GetItemTemplate.Call(uintptr(hEle))
	return int(r)
}
