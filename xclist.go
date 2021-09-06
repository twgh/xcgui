package xcgui

import "unsafe"

// 列表_创建
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func XList_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xList_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表_增加列
// hEle: 元素句柄.
// width: 列宽度.
func XList_AddColumn(hEle int, width int) int {
	r, _, _ := xList_AddColumn.Call(uintptr(hEle), uintptr(width))
	return int(r)
}

// 列表_插入列
// hEle: 元素句柄.
// width: 列宽度.
// iItem: 插入位置索引.
func XList_InsertColumn(hEle int, width int, iItem int) int {
	r, _, _ := xList_InsertColumn.Call(uintptr(hEle), uintptr(width), uintptr(iItem))
	return int(r)
}

// 列表_启用多选
// hEle: 元素句柄.
// bEnable: 是否启用.
func XList_EnableMultiSel(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableMultiSel.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表_启用拖动更改列宽
// hEle: 元素句柄.
// bEnable: 是否启用.
func XList_EnableDragChangeColumnWidth(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableDragChangeColumnWidth.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表_启用垂直滚动条顶部对齐
// hEle: 元素句柄.
// bTop: 是否启用.
func XList_EnableVScrollBarTop(hEle int, bTop bool) int {
	r, _, _ := xList_EnableVScrollBarTop.Call(uintptr(hEle), boolPtr(bTop))
	return int(r)
}

// 列表_启用项背景全行模式
// hEle: 元素句柄.
// bFull: 是否启用.
func XList_EnableItemBkFullRow(hEle int, bFull bool) int {
	r, _, _ := xList_EnableItemBkFullRow.Call(uintptr(hEle), boolPtr(bFull))
	return int(r)
}

// 列表_启用固定行高
// hEle: 元素句柄
// bEnable: 是否启用
func XList_EnableFixedRowHeight(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableFixedRowHeight.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表_启用模板复用
// hEle: 元素句柄
// bEnable: 是否启用
func XList_EnablemTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xList_EnablemTemplateReuse.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表_启用虚表
// hEle: 元素句柄
// bEnable: 是否启用
func XList_EnableVirtualTable(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableVirtualTable.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表_置虚表行数
// hEle: 元素句柄
// nRowCount: 行数
func XList_SetVirtualRowCount(hEle int, nRowCount int) int {
	r, _, _ := xList_SetVirtualRowCount.Call(uintptr(hEle), uintptr(nRowCount))
	return int(r)
}

// 列表_置排序
// hEle: 元素句柄.
// iColumn: 列索引.
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
// bEnable: 是否启用排序功能.
func XList_SetSort(hEle int, iColumn int, iColumnAdapter int, bEnable bool) int {
	r, _, _ := xList_SetSort.Call(uintptr(hEle), uintptr(iColumn), uintptr(iColumnAdapter), boolPtr(bEnable))
	return int(r)
}

// 列表_置绘制项背景标志
// hEle: 元素句柄.
// nFlags: 标志位@reflist_drawItemBk_flag_.
func XList_SetDrawItemBkFlags(hEle int, nFlags int) int {
	r, _, _ := xList_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表_置列宽
// hEle: 元素句柄.
// iItem: 列索引.
// width: 宽度.
func XList_SetColumnWidth(hEle int, iItem int, width int) int {
	r, _, _ := xList_SetColumnWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
	return int(r)
}

// 列表_置列最小宽度
// hEle: 元素句柄.
// iItem: 列索引.
// width: 宽度.
func XList_SetColumnMinWidth(hEle int, iItem int, width int) int {
	r, _, _ := xList_SetColumnMinWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
	return int(r)
}

// 列表_置列宽度固定
// hEle: 元素句柄.
// iColumn: 列索引.
// bFixed: 是否固定宽度.
func XList_SetColumnWidthFixed(hEle int, iColumn int, bFixed bool) int {
	r, _, _ := xList_SetColumnWidthFixed.Call(uintptr(hEle), uintptr(iColumn), boolPtr(bFixed))
	return int(r)
}

// 列表_取列宽度
// hEle: 元素句柄.
// iColumn: 列索引.
func XList_GetColumnWidth(hEle int, iColumn int) int {
	r, _, _ := xList_GetColumnWidth.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_取列数量
// hEle: 元素句柄.
func XList_GetColumnCount(hEle int) int {
	r, _, _ := xList_GetColumnCount.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项数据
// hEle: 元素句柄.
// iItem: 项索引.
// iSubItem: 子项索引.
// data: 用户数据.
func XList_SetItemData(hEle int, iItem int, iSubItem int, data int) bool {
	r, _, _ := xList_SetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem), uintptr(data))
	return int(r) != 0
}

// 列表_取项数据
// hEle: 元素句柄.
// iItem: 项索引.
// iSubItem: 子项索引.
func XList_GetItemData(hEle int, iItem int, iSubItem int) int {
	r, _, _ := xList_GetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem))
	return int(r)
}

// 列表_置选择项
// hEle: 元素句柄.
// iItem: 项索引.
func XList_SetSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_SetSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r) != 0
}

// 列表_取选择项
// hEle: 元素句柄.
func XList_GetSelectItem(hEle int) int {
	r, _, _ := xList_GetSelectItem.Call(uintptr(hEle))
	return int(r)
}

// 列表_取选择项数量
// hEle: 元素句柄.
func XList_GetSelectItemCount(hEle int) int {
	r, _, _ := xList_GetSelectItemCount.Call(uintptr(hEle))
	return int(r)
}

// 列表_添加选择项
// hEle: 元素句柄
// iItem: 项索引
func XList_AddSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_AddSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r) != 0
}

// 列表_置选择全部
// hEle: 元素句柄.
func XList_SetSelectAll(hEle int) int {
	r, _, _ := xList_SetSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_取全部选择
// hEle: 元素句柄.
// pArray: 接收行索引数组.
// nArraySize: 数组大小.
func XList_GetSelectAll(hEle int, pArray int, nArraySize int) int {
	r, _, _ := xList_GetSelectAll.Call(uintptr(hEle), uintptr(pArray), uintptr(nArraySize))
	return int(r)
}

// 列表_显示指定项
// hEle: 元素句柄.
// iItem: 项索引.
func XList_VisibleItem(hEle int, iItem int) int {
	r, _, _ := xList_VisibleItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表_取消选择项
// hEle: 元素句柄.
// iItem: 项索引.
func XList_CancelSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_CancelSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r) != 0
}

// 列表_取消全部选择项
// hEle: 元素句柄.
func XList_CancelSelectAll(hEle int) int {
	r, _, _ := xList_CancelSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_取列表头
// hEle: 元素句柄.
func XList_GetHeaderHELE(hEle int) int {
	r, _, _ := xList_GetHeaderHELE.Call(uintptr(hEle))
	return int(r)
}

// 列表_删除列
// hEle: 元素句柄.
// iItem: 项索引.
func XList_DeleteColumn(hEle int, iItem int) bool {
	r, _, _ := xList_DeleteColumn.Call(uintptr(hEle), uintptr(iItem))
	return int(r) != 0
}

// 列表_删除全部列
// hEle: 元素句柄.
func XList_DeleteColumnAll(hEle int) int {
	r, _, _ := xList_DeleteColumnAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_绑定数据适配器
// hEle: 元素句柄.
// hAdapter: 数据适配器句柄XAdTable.
func XList_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xList_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表_列表头绑定数据适配器
// hEle: 元素句柄.
// hAdapter: 数据适配器句柄XAdMap.
func XList_BindAdapterHeader(hEle int, hAdapter int) int {
	r, _, _ := xList_BindAdapterHeader.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表_创建数据适配器
// hEle: 元素句柄.
func XList_CreateAdapter(hEle int) int {
	r, _, _ := xList_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表_列表头创建数据适配器
// hEle: 元素句柄.
func XList_CreateAdapterHeader(hEle int) int {
	r, _, _ := xList_CreateAdapterHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_取数据适配器
// hEle: 元素句柄.
func XList_GetAdapter(hEle int) int {
	r, _, _ := xList_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表_列表头获取数据适配器
// hEle: 元素句柄.
func XList_GetAdapterHeader(hEle int) int {
	r, _, _ := xList_GetAdapterHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项模板文件
// hEle: 元素句柄.
// pXmlFile: 文件名.
func XList_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xList_SetItemTemplateXML.Call(uintptr(hEle), strPtr(pXmlFile))
	return int(r) != 0
}

// 列表_置项模板从字符串
// hEle: 元素句柄.
// pStringXML: 字符串指针.
func XList_SetItemTemplateXMLFromString(hEle int, pStringXML int) bool {
	r, _, _ := xList_SetItemTemplateXMLFromString.Call(uintptr(hEle), uintptr(pStringXML))
	return int(r) != 0
}

// 列表_置项模板
// hEle: 元素句柄.
// hTemp: 模板句柄.
func XList_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xList_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return int(r) != 0
}

// 列表_取项模板对象
// hEle: 元素句柄.
// iItem: 项索引.
// iSubItem: 子项索引.
// nTempItemID: 模板项itemID.
func XList_GetTemplateObject(hEle int, iItem int, iSubItem int, nTempItemID int) int {
	r, _, _ := xList_GetTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取对象所在行
// hEle: 元素句柄.
// hXCGUI: 对象句柄
func XList_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xList_GetItemIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表_取列表头模板对象
// hEle: 元素句柄.
// iItem: 列表头项ID.
// nTempItemID: 模板项ID.
func XList_GetHeaderTemplateObject(hEle int, iItem int, nTempItemID int) int {
	r, _, _ := xList_GetHeaderTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取列表头对象所在行
// hEle: 元素句柄.
// hXCGUI: 对象句柄.
func XList_GetHeaderItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xList_GetHeaderItemIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表_置列表头高度
// hEle: 元素句柄.
// height: 高度.
func XList_SetHeaderHeight(hEle int, height int) int {
	r, _, _ := xList_SetHeaderHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 列表_取列表头高度
// hEle: 元素句柄.
func XList_GetHeaderHeight(hEle int) int {
	r, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))
	return int(r)
}

// 列表_取可视行范围
// hEle: 元素句柄.
// piStart: 开始行索引.
// piEnd: 结束行索引.
func XList_GetVisibleRowRange(hEle int, piStart *int, piEnd *int) int {
	r, _, _ := xList_GetVisibleRowRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piStart)), uintptr(unsafe.Pointer(piEnd)))
	return int(r)
}

// 列表_添加项背景边框
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func XList_AddItemBkBorder(hEle int, nState int, color int, alpha uint8, width int) int {
	r, _, _ := xList_AddItemBkBorder.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha), uintptr(width))
	return int(r)
}

// 列表_添加项背景填充
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
func XList_AddItemBkFill(hEle int, nState int, color int, alpha uint8) int {
	r, _, _ := xList_AddItemBkFill.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha))
	return int(r)
}

// 列表_添加项背景图片
// hEle: 元素句柄.
// nState: 项状态.
// hImage: 图片句柄.
func XList_AddItemBkImage(hEle int, nState int, hImage int) int {
	r, _, _ := xList_AddItemBkImage.Call(uintptr(hEle), uintptr(nState), uintptr(hImage))
	return int(r)
}

// 列表_取项背景对象数量
// hEle: 元素句柄.
func XList_GetItemBkInfoCount(hEle int) int {
	r, _, _ := xList_GetItemBkInfoCount.Call(uintptr(hEle))
	return int(r)
}

// 列表_清空项背景对象
// hEle: 元素句柄.
func XList_ClearItemBkInfo(hEle int) int {
	r, _, _ := xList_ClearItemBkInfo.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项默认高度
// hEle: 元素句柄.
// nHeight: 高度.
// nSelHeight: 选中时高度.
func XList_SetItemHeightDefault(hEle int, nHeight int, nSelHeight int) int {
	r, _, _ := xList_SetItemHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表_取项默认高度
// hEle: 元素句柄.
// pHeight: 高度.
// pSelHeight: 选中时高度.
func XList_GetItemHeightDefault(hEle int, pHeight *int, pSelHeight *int) int {
	r, _, _ := xList_GetItemHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表_置行间距
// hEle: 元素句柄.
// nSpace: 行间距大小.
func XList_SetRowSpace(hEle int, nSpace int) int {
	r, _, _ := xList_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 列表_取行间距
// hEle: 元素句柄.
func XList_GetRowSpace(hEle int) int {
	r, _, _ := xList_GetRowSpace.Call(uintptr(hEle))
	return int(r)
}

// 列表_置锁定列左侧
// hEle: 元素句柄.
// iColumn: 列索引
func XList_SetLockColumnLeft(hEle int, iColumn int) int {
	r, _, _ := xList_SetLockColumnLeft.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_置锁定列右侧
// hEle: 元素句柄.
// iColumn: 列索引
func XList_SetLockColumnRight(hEle int, iColumn int) int {
	r, _, _ := xList_SetLockColumnRight.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_置锁定行底部
// hEle: 元素句柄.
// bLock: 是否锁定.
func XList_SetLockRowBottom(hEle int, bLock bool) int {
	r, _, _ := xList_SetLockRowBottom.Call(uintptr(hEle), boolPtr(bLock))
	return int(r)
}

// 列表_置锁定行底部重叠
// hEle: 元素句柄
// bOverlap: 是否重叠
func XList_SetLockRowBottomOverlap(hEle int, bOverlap bool) int {
	r, _, _ := xList_SetLockRowBottomOverlap.Call(uintptr(hEle), boolPtr(bOverlap))
	return int(r)
}

// 列表_测试点击项
// hEle: 元素句柄.
// pPt: 坐标点.
// piItem: 项索引.
// piSubItem: 子项索引.
func XList_HitTest(hEle int, pPt *POINT, piItem *int, piSubItem *int) bool {
	r, _, _ := xList_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(piItem)), uintptr(unsafe.Pointer(piSubItem)))
	return int(r) != 0
}

// 列表_测试点击项扩展
// hEle: 元素句柄.
// pPt: 坐标点.
// piItem: 项索引.
// piSubItem: 子项索引.
func XList_HitTestOffset(hEle int, pPt *POINT, piItem *int, piSubItem *int) bool {
	r, _, _ := xList_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(piItem)), uintptr(unsafe.Pointer(piSubItem)))
	return int(r) != 0
}

// 列表_刷新项数据
// hEle: 元素句柄.
func XList_RefreshData(hEle int) int {
	r, _, _ := xList_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表_刷新指定项
// hEle: 元素句柄.
// iItem: 项索引.
func XList_RefreshItem(hEle int, iItem int) int {
	r, _, _ := xList_RefreshItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表_添加列文本
// hEle:
// nWidth:
// pName:
// pText:
func XList_AddColumnText(hEle int, nWidth int, pName string, pText string) int {
	r, _, _ := xList_AddColumnText.Call(uintptr(hEle), uintptr(nWidth), strPtr(pName), strPtr(pText))
	return int(r)
}

// 列表_添加列图片
// hEle:
// nWidth:
// pName:
// hImage:
func XList_AddColumnImage(hEle int, nWidth int, pName string, hImage int) int {
	r, _, _ := xList_AddColumnImage.Call(uintptr(hEle), uintptr(nWidth), strPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_添加项文本
// hEle:
// pText:
func XList_AddItemText(hEle int, pText string) int {
	r, _, _ := xList_AddItemText.Call(uintptr(hEle), strPtr(pText))
	return int(r)
}

// 列表_添加项文本扩展
// hEle:
// pName:
// pText:
func XList_AddItemTextEx(hEle int, pName string, pText string) int {
	r, _, _ := xList_AddItemTextEx.Call(uintptr(hEle), strPtr(pName), strPtr(pText))
	return int(r)
}

// 列表_添加项图片
// hEle:
// hImage:
func XList_AddItemImage(hEle int, hImage int) int {
	r, _, _ := xList_AddItemImage.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 列表_添加项图片扩展
// hEle:
// pName:
// hImage:
func XList_AddItemImageEx(hEle int, pName string, hImage int) int {
	r, _, _ := xList_AddItemImageEx.Call(uintptr(hEle), strPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_插入项文本
// hEle:
// iItem:
// pValue:
func XList_InsertItemText(hEle int, iItem int, pValue string) int {
	r, _, _ := xList_InsertItemText.Call(uintptr(hEle), uintptr(iItem), strPtr(pValue))
	return int(r)
}

// 列表_插入项文本扩展
// hEle:
// iItem:
// pName:
// pValue:
func XList_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {
	r, _, _ := xList_InsertItemTextEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), strPtr(pValue))
	return int(r)
}

// 列表_插入项图片
// hEle:
// iItem:
// hImage:
func XList_InsertItemImage(hEle int, iItem int, hImage int) int {
	r, _, _ := xList_InsertItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(hImage))
	return int(r)
}

// 列表_插入项图片扩展
// hEle:
// iItem:
// pName:
// hImage:
func XList_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {
	r, _, _ := xList_InsertItemImageEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_置项文本
// hEle:
// iItem:
// iColumn:
// pText:
func XList_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {
	r, _, _ := xList_SetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), strPtr(pText))
	return int(r) != 0
}

// 列表_置项文本扩展
// hEle:
// iItem:
// pName:
// pText:
func XList_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {
	r, _, _ := xList_SetItemTextEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), strPtr(pText))
	return int(r) != 0
}

// 列表_置项图片
// hEle:
// iItem:
// iColumn:
// hImage:
func XList_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xList_SetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return int(r) != 0
}

// 列表_置项图片扩展
// hEle:
// iItem:
// pName:
// hImage:
func XList_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {
	r, _, _ := xList_SetItemImageEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(hImage))
	return int(r) != 0
}

// 列表_置项指数值
// hEle:
// iItem:
// iColumn:
// nValue:
func XList_SetItemInt(hEle int, iItem int, iColumn int, nValue int) bool {
	r, _, _ := xList_SetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return int(r) != 0
}

// 列表_置项整数值扩展
// hEle:
// iItem:
// pName:
// nValue:
func XList_SetItemIntEx(hEle int, iItem int, pName string, nValue int) bool {
	r, _, _ := xList_SetItemIntEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(nValue))
	return int(r) != 0
}

// 列表_置项浮点值
// hEle:
// iItem:
// iColumn:
// fFloat:
func XList_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {
	r, _, _ := xList_SetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(fFloat))
	return int(r) != 0
}

// 列表_置项浮点值扩展
// hEle:
// iItem:
// pName:
// fFloat:
func XList_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {
	r, _, _ := xList_SetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(fFloat))
	return int(r) != 0
}

// 列表_取项文本
// hEle:
// iItem:
// iColumn:
func XList_GetItemText(hEle int, iItem int, iColumn int) string {
	r, _, _ := xList_GetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return uintPtrToString(r)
}

// 列表_取项文本扩展
// hEle:
// iItem:
// pName:
func XList_GetItemTextEx(hEle int, iItem int, pName string) string {
	r, _, _ := xList_GetItemTextEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName))
	return uintPtrToString(r)
}

// 列表_取项图片
// hEle:
// iItem:
// iColumn:
func XList_GetItemImage(hEle int, iItem int, iColumn int) int {
	r, _, _ := xList_GetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 列表_取项图片扩展
// hEle:
// iItem:
// pName:
func XList_GetItemImageEx(hEle int, iItem int, pName string) int {
	r, _, _ := xList_GetItemImageEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName))
	return int(r)
}

// 列表_取项整数值
// hEle:
// iItem:
// iColumn:
// pOutValue:
func XList_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int) bool {
	r, _, _ := xList_GetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return int(r) != 0
}

// 列表_取项整数值扩展
// hEle:
// iItem:
// pName:
// pOutValue:
func XList_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int) bool {
	r, _, _ := xList_GetItemIntEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return int(r) != 0
}

// 列表_取项浮点值
// hEle:
// iItem:
// iColumn:
// pOutValue:
func XList_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return int(r) != 0
}

// 列表_取项浮点值扩展
// hEle:
// iItem:
// pName:
// pOutValue:
func XList_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), strPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return int(r) != 0
}

// 列表_删除项
// hEle:
// iItem:
func XList_DeleteItem(hEle int, iItem int) bool {
	r, _, _ := xList_DeleteItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r) != 0
}

// 列表_删除项扩展
// hEle:
// iItem:
// nCount:
func XList_DeleteItemEx(hEle int, iItem int, nCount int) bool {
	r, _, _ := xList_DeleteItemEx.Call(uintptr(hEle), uintptr(iItem), uintptr(nCount))
	return int(r) != 0
}

// 列表_删除项全部
// hEle:
func XList_DeleteItemAll(hEle int) int {
	r, _, _ := xList_DeleteItemAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_删除列全部AD
// hEle:
func XList_DeleteColumnAll_AD(hEle int) int {
	r, _, _ := xList_DeleteColumnAll_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表_取项数量AD
// hEle:
func XList_GetCount_AD(hEle int) int {
	r, _, _ := xList_GetCount_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表_取列数量AD
// hEle:
func XList_GetCountColumn_AD(hEle int) int {
	r, _, _ := xList_GetCountColumn_AD.Call(uintptr(hEle))
	return int(r)
}
