package xc

import (
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 列表框_创建, 创建列表框元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func XListBox_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xListBox_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表框_创建Ex. 创建列表框元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
//
// col_extend_count: 列数量. 例如: 内置模板是1列, 如果数据有5列, 那么此参数填5.
func XListBox_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {
	r, _, _ := xListBox_CreateEx.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent), uintptr(col_extend_count))
	return int(r)
}

// 列表框_启用固定行高.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListBox_EnableFixedRowHeight(hEle int, bEnable bool) int {
	r, _, _ := xListBox_EnableFixedRowHeight.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表框_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListBox_EnableTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xListBox_EnableTemplateReuse.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表框_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListBox_EnableVirtualTable(hEle int, bEnable bool) int {
	r, _, _ := xListBox_EnableVirtualTable.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表框_置虚表行数.
//
// hEle: 元素句柄.
//
// nRowCount: 行数.
func XListBox_SetVirtualRowCount(hEle int, nRowCount int) int {
	r, _, _ := xListBox_SetVirtualRowCount.Call(uintptr(hEle), uintptr(nRowCount))
	return int(r)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func XListBox_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {
	r, _, _ := xListBox_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表框_置项数据, 设置项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 想索引.
//
// nUserData: 用户数据.
func XListBox_SetItemData(hEle int, iItem int, nUserData int) bool {
	r, _, _ := xListBox_SetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(nUserData))
	return r != 0
}

// 列表框_取项数据, 获取项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_GetItemData(hEle int, iItem int) int {
	r, _, _ := xListBox_GetItemData.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表框_置项信息.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pItem: 项信息.
func XListBox_SetItemInfo(hEle int, iItem int, pItem *ListBox_Item_Info_) bool {
	r, _, _ := xListBox_SetItemInfo.Call(uintptr(hEle), uintptr(iItem), uintptr(unsafe.Pointer(pItem)))
	return r != 0
}

// 列表框_取项背景信息, 获取项信息.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pItem: 项信息.
func XListBox_GetItemInfo(hEle int, iItem int, pItem *ListBox_Item_Info_) bool {
	r, _, _ := xListBox_GetItemInfo.Call(uintptr(hEle), uintptr(iItem), uintptr(unsafe.Pointer(pItem)))
	return r != 0
}

// 列表框_置选择项, 设置选择选.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_SetSelectItem(hEle int, iItem int) bool {
	r, _, _ := xListBox_SetSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表框_取选择项, 返回项索引.
//
// hEle: 元素句柄.
func XListBox_GetSelectItem(hEle int) int {
	r, _, _ := xListBox_GetSelectItem.Call(uintptr(hEle))
	return int(r)
}

// 列表框_添加选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_AddSelectItem(hEle int, iItem int) bool {
	r, _, _ := xListBox_AddSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表框_取消选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_CancelSelectItem(hEle int, iItem int) bool {
	r, _, _ := xListBox_CancelSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
//
// hEle: 元素句柄.
func XListBox_CancelSelectAll(hEle int) bool {
	r, _, _ := xListBox_CancelSelectAll.Call(uintptr(hEle))
	return r != 0
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//
// hEle: 元素句柄.
//
// pArray: 数组缓冲区.
//
// nArraySize: 数组大小.
func XListBox_GetSelectAll(hEle int, pArray *[]int32, nArraySize int) int {
	if nArraySize < 1 {
		return 0
	}
	*pArray = make([]int32, nArraySize)
	r, _, _ := xListBox_GetSelectAll.Call(uintptr(hEle), uintptr(unsafe.Pointer(&(*pArray)[0])), uintptr(nArraySize))
	return int(r)
}

// 列表框_取选择项数量, 获取选择项数量.
//
// hEle: 元素句柄.
func XListBox_GetSelectCount(hEle int) int {
	r, _, _ := xListBox_GetSelectCount.Call(uintptr(hEle))
	return int(r)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
//
// hEle: 元素句柄.
func XListBox_GetItemMouseStay(hEle int) int {
	r, _, _ := xListBox_GetItemMouseStay.Call(uintptr(hEle))
	return int(r)
}

// 列表框_选择全部项.
//
// hEle: 元素句柄.
func XListBox_SelectAll(hEle int) bool {
	r, _, _ := xListBox_SelectAll.Call(uintptr(hEle))
	return r != 0
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_VisibleItem(hEle int, iItem int) int {
	r, _, _ := xListBox_VisibleItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表框_取可视行范围, 获取当前可见行范围.
//
// hEle: 元素句柄.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func XListBox_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {
	r, _, _ := xListBox_GetVisibleRowRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piStart)), uintptr(unsafe.Pointer(piEnd)))
	return int(r)
}

// 列表框_置项默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 项高度.
//
// nSelHeight: 选中项高度.
func XListBox_SetItemHeightDefault(hEle int, nHeight, nSelHeight int32) int {
	r, _, _ := xListBox_SetItemHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表框_取项默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func XListBox_GetItemHeightDefault(hEle int, pHeight, pSelHeight *int32) int {
	r, _, _ := xListBox_GetItemHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func XListBox_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xListBox_GetItemIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表框_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 间距大小.
func XListBox_SetRowSpace(hEle int, nSpace int) int {
	r, _, _ := xListBox_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 列表框_取行间距.
//
// hEle: 元素句柄.
func XListBox_GetRowSpace(hEle int) int {
	r, _, _ := xListBox_GetRowSpace.Call(uintptr(hEle))
	return int(r)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func XListBox_HitTest(hEle int, pPt *POINT) int {
	r, _, _ := xListBox_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func XListBox_HitTestOffset(hEle int, pPt *POINT) int {
	r, _, _ := xListBox_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func XListBox_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xListBox_SetItemTemplateXML.Call(uintptr(hEle), common.StrPtr(pXmlFile))
	return r != 0
}

// 列表框_置项模板, 设置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func XListBox_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xListBox_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return r != 0
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func XListBox_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromString.Call(uintptr(hEle), XC_wtoa(pStringXML))
	return r != 0
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func XListBox_GetTemplateObject(hEle int, iItem int, nTempItemID int) int {
	r, _, _ := xListBox_GetTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表框_启用多选, 是否启用多行选择功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListBox_EnableMultiSel(hEle int, bEnable bool) int {
	r, _, _ := xListBox_EnableMultiSel.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func XListBox_CreateAdapter(hEle int) int {
	r, _, _ := xListBox_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdTable.
func XListBox_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xListBox_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func XListBox_GetAdapter(hEle int) int {
	r, _, _ := xListBox_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表框_排序.
//
// hEle: 元素句柄.
//
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
//
// bAscending: 升序(TRUE)或降序(FALSE).
func XListBox_Sort(hEle int, iColumnAdapter int, bAscending bool) int {
	r, _, _ := xListBox_Sort.Call(uintptr(hEle), uintptr(iColumnAdapter), common.BoolPtr(bAscending))
	return int(r)
}

// 列表框_刷新数据.
//
// hEle: 元素句柄.
func XListBox_RefreshData(hEle int) int {
	r, _, _ := xListBox_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XListBox_RefreshItem(hEle int, iItem int) int {
	r, _, _ := xListBox_RefreshItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引.
//
// hEle:.
//
// pText:.
func XListBox_AddItemText(hEle int, pText string) int {
	r, _, _ := xListBox_AddItemText.Call(uintptr(hEle), common.StrPtr(pText))
	return int(r)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx.
//
// hEle:.
//
// pName:.
//
// pText:.
func XListBox_AddItemTextEx(hEle int, pName string, pText string) int {
	r, _, _ := xListBox_AddItemTextEx.Call(uintptr(hEle), common.StrPtr(pName), common.StrPtr(pText))
	return int(r)
}

// 列表框_添加项图片, XAdTable_AddItemImage.
//
// hEle:.
//
// hImage:.
func XListBox_AddItemImage(hEle int, hImage int) int {
	r, _, _ := xListBox_AddItemImage.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx.
//
// hEle:.
//
// pName:.
//
// hImage:.
func XListBox_AddItemImageEx(hEle int, pName string, hImage int) int {
	r, _, _ := xListBox_AddItemImageEx.Call(uintptr(hEle), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表框_插入项文本.
//
// hEle:.
//
// iItem:.
//
// pValue:.
func XListBox_InsertItemText(hEle int, iItem int, pValue string) int {
	r, _, _ := xListBox_InsertItemText.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pValue))
	return int(r)
}

// 列表框_插入项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pValue:.
func XListBox_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {
	r, _, _ := xListBox_InsertItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return int(r)
}

// 列表框_插入项图片.
//
// hEle:.
//
// iItem:.
//
// hImage:.
func XListBox_InsertItemImage(hEle int, iItem int, hImage int) int {
	r, _, _ := xListBox_InsertItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(hImage))
	return int(r)
}

// 列表框_插入项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func XListBox_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {
	r, _, _ := xListBox_InsertItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表框_置项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func XListBox_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {
	r, _, _ := xListBox_SetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.StrPtr(pText))
	return r != 0
}

// 列表框_置项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pText:.
func XListBox_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {
	r, _, _ := xListBox_SetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pText))
	return r != 0
}

// 列表框_置项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func XListBox_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xListBox_SetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 列表框_置项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func XListBox_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {
	r, _, _ := xListBox_SetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 列表框_置项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func XListBox_SetItemInt(hEle int, iItem int, iColumn int, nValue int) bool {
	r, _, _ := xListBox_SetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 列表框_置项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// nValue:.
func XListBox_SetItemIntEx(hEle int, iItem int, pName string, nValue int) bool {
	r, _, _ := xListBox_SetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(nValue))
	return r != 0
}

// 列表框_置项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func XListBox_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {
	r, _, _ := xListBox_SetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表框_置项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func XListBox_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {
	r, _, _ := xListBox_SetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表框_取项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XListBox_GetItemText(hEle int, iItem int, iColumn int) string {
	r, _, _ := xListBox_GetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 列表框_取项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func XListBox_GetItemTextEx(hEle int, iItem int, pName string) string {
	r, _, _ := xListBox_GetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 列表框_取项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XListBox_GetItemImage(hEle int, iItem int, iColumn int) int {
	r, _, _ := xListBox_GetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 列表框_取项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func XListBox_GetItemImageEx(hEle int, iItem int, pName string) int {
	r, _, _ := xListBox_GetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 列表框_取项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func XListBox_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {
	r, _, _ := xListBox_GetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表框_取项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func XListBox_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {
	r, _, _ := xListBox_GetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表框_取项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func XListBox_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {
	r, _, _ := xListBox_GetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表框_取项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func XListBox_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {
	r, _, _ := xListBox_GetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表框_删除项.
//
// hEle:.
//
// iItem:.
func XListBox_DeleteItem(hEle int, iItem int) bool {
	r, _, _ := xListBox_DeleteItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表框_删除项扩展.
//
// hEle:.
//
// iItem:.
//
// nCount:.
func XListBox_DeleteItemEx(hEle int, iItem int, nCount int) bool {
	r, _, _ := xListBox_DeleteItemEx.Call(uintptr(hEle), uintptr(iItem), uintptr(nCount))
	return r != 0
}

// 列表框_删除项全部.
//
// hEle:.
func XListBox_DeleteItemAll(hEle int) int {
	r, _, _ := xListBox_DeleteItemAll.Call(uintptr(hEle))
	return int(r)
}

// 列表框_删除列全部.
//
// hEle:.
func XListBox_DeleteColumnAll(hEle int) int {
	r, _, _ := xListBox_DeleteColumnAll.Call(uintptr(hEle))
	return int(r)
}

// 列表框_取项数量AD.
//
// hEle:.
func XListBox_GetCount_AD(hEle int) int {
	r, _, _ := xListBox_GetCount_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表框_取列数量AD.
//
// hEle:.
func XListBox_GetCountColumn_AD(hEle int) int {
	r, _, _ := xListBox_GetCountColumn_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表框_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XListBox_SetSplitLineColor(hEle int, color int) int {
	r, _, _ := xListBox_SetSplitLineColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 列表框_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func XListBox_SetDragRectColor(hEle int, color, width int) int {
	r, _, _ := xListBox_SetDragRectColor.Call(uintptr(hEle), uintptr(color), uintptr(width))
	return int(r)
}

// 列表框_置项模板从内存. 设置项模板文件.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func XListBox_SetItemTemplateXMLFromMem(hEle int, data []byte) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromMem.Call(uintptr(hEle), common.ByteSliceDataPtr(&data), uintptr(len(data)))
	return r != 0
}

// 列表框_置项模板从资源ZIP. 设置项模板文件.
//
// hEle: 元素句柄.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func XListBox_SetItemTemplateXMLFromZipRes(hEle, id int, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromZipRes.Call(uintptr(hEle), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}

// 列表框_取项模板. 获取列表项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func XListBox_GetItemTemplate(hEle int) int {
	r, _, _ := xListBox_GetItemTemplate.Call(uintptr(hEle))
	return int(r)
}
