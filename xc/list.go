package xc

import (
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 列表_创建, 创建列表元素, 返回元素句柄.
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
func XList_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xList_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表_创建Ex, 创建列表元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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
func XList_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {
	r, _, _ := xList_CreateEx.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent), uintptr(col_extend_count))
	return int(r)
}

// 列表_增加列, 返回位置索引.
//
// hEle: 元素句柄.
//
// width: 列宽度.
func XList_AddColumn(hEle int, width int) int {
	r, _, _ := xList_AddColumn.Call(uintptr(hEle), uintptr(width))
	return int(r)
}

// 列表_插入列, 返回插入位置索引.
//
// hEle: 元素句柄.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func XList_InsertColumn(hEle int, width int, iItem int) int {
	r, _, _ := xList_InsertColumn.Call(uintptr(hEle), uintptr(width), uintptr(iItem))
	return int(r)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableMultiSel(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableMultiSel.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableDragChangeColumnWidth(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableDragChangeColumnWidth.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_启用垂直滚动条顶部对齐.
//
// hEle: 元素句柄.
//
// bTop: 是否启用.
func XList_EnableVScrollBarTop(hEle int, bTop bool) int {
	r, _, _ := xList_EnableVScrollBarTop.Call(uintptr(hEle), common.BoolPtr(bTop))
	return int(r)
}

// 列表_启用项背景全行模式, 启用项背景全行填充模式.
//
// hEle: 元素句柄.
//
// bFull: 是否启用.
func XList_EnableItemBkFullRow(hEle int, bFull bool) int {
	r, _, _ := xList_EnableItemBkFullRow.Call(uintptr(hEle), common.BoolPtr(bFull))
	return int(r)
}

// 列表_启用固定行高.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableFixedRowHeight(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableFixedRowHeight.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableTemplateReuse.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableVirtualTable(hEle int, bEnable bool) int {
	r, _, _ := xList_EnableVirtualTable.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_置虚表行数.
//
// hEle: 元素句柄.
//
// nRowCount: 行数.
func XList_SetVirtualRowCount(hEle int, nRowCount int) int {
	r, _, _ := xList_SetVirtualRowCount.Call(uintptr(hEle), uintptr(nRowCount))
	return int(r)
}

// 列表_置排序, 设置排序属性.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
//
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
//
// bEnable: 是否启用排序功能.
func XList_SetSort(hEle int, iColumn int, iColumnAdapter int, bEnable bool) int {
	r, _, _ := xList_SetSort.Call(uintptr(hEle), uintptr(iColumn), uintptr(iColumnAdapter), common.BoolPtr(bEnable))
	return int(r)
}

// 列表_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func XList_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {
	r, _, _ := xList_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表_置列宽.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func XList_SetColumnWidth(hEle int, iItem int, width int) int {
	r, _, _ := xList_SetColumnWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
	return int(r)
}

// 列表_置列最小宽度.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func XList_SetColumnMinWidth(hEle int, iItem int, width int) int {
	r, _, _ := xList_SetColumnMinWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
	return int(r)
}

// 列表_置列宽度固定.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func XList_SetColumnWidthFixed(hEle int, iColumn int, bFixed bool) int {
	r, _, _ := xList_SetColumnWidthFixed.Call(uintptr(hEle), uintptr(iColumn), common.BoolPtr(bFixed))
	return int(r)
}

// 列表_取列宽度.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func XList_GetColumnWidth(hEle int, iColumn int) int {
	r, _, _ := xList_GetColumnWidth.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_取列数量.
//
// hEle: 元素句柄.
func XList_GetColumnCount(hEle int) int {
	r, _, _ := xList_GetColumnCount.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项数据, 设置项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// data: 用户数据.
func XList_SetItemData(hEle int, iItem int, iSubItem int, data int) bool {
	r, _, _ := xList_SetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem), uintptr(data))
	return r != 0
}

// 列表_取项数据, 获取项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
func XList_GetItemData(hEle int, iItem int, iSubItem int) int {
	r, _, _ := xList_GetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem))
	return int(r)
}

// 列表_置选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_SetSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_SetSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_取选择项, 返回项索引.
//
// hEle: 元素句柄.
func XList_GetSelectItem(hEle int) int {
	r, _, _ := xList_GetSelectItem.Call(uintptr(hEle))
	return int(r)
}

// 列表_取选择项数量, 获取选择项数量.
//
// hEle: 元素句柄.
func XList_GetSelectItemCount(hEle int) int {
	r, _, _ := xList_GetSelectItemCount.Call(uintptr(hEle))
	return int(r)
}

// 列表_添加选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_AddSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_AddSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_置选择全部, 选择全部行.
//
// hEle: 元素句柄.
func XList_SetSelectAll(hEle int) int {
	r, _, _ := xList_SetSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// hEle: 元素句柄.
//
// pArray: 接收行索引数组.
//
// nArraySize: 数组大小.
func XList_GetSelectAll(hEle int, pArray *[]int32, nArraySize int) int {
	if nArraySize < 1 {
		return 0
	}
	*pArray = make([]int32, nArraySize)
	r, _, _ := xList_GetSelectAll.Call(uintptr(hEle), uintptr(unsafe.Pointer(&(*pArray)[0])), uintptr(nArraySize))
	return int(r)
}

// 列表_显示指定项, 滚动视图让指定项可见.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_VisibleItem(hEle int, iItem int) int {
	r, _, _ := xList_VisibleItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表_取消选择项, 取消选择指定项(这里的项可以理解为行).
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_CancelSelectItem(hEle int, iItem int) bool {
	r, _, _ := xList_CancelSelectItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
//
// hEle: 元素句柄.
func XList_CancelSelectAll(hEle int) int {
	r, _, _ := xList_CancelSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄.
//
// hEle: 元素句柄.
func XList_GetHeaderHELE(hEle int) int {
	r, _, _ := xList_GetHeaderHELE.Call(uintptr(hEle))
	return int(r)
}

// 列表_删除列.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_DeleteColumn(hEle int, iItem int) bool {
	r, _, _ := xList_DeleteColumn.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
//
// hEle: 元素句柄.
func XList_DeleteColumnAll(hEle int) int {
	r, _, _ := xList_DeleteColumnAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdTable.
func XList_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xList_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表_列表头绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdMap.
func XList_BindAdapterHeader(hEle int, hAdapter int) int {
	r, _, _ := xList_BindAdapterHeader.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func XList_CreateAdapter(hEle int, colExtend_count int) int {
	r, _, _ := xList_CreateAdapter.Call(uintptr(hEle), uintptr(colExtend_count))
	return int(r)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func XList_CreateAdapterHeader(hEle int) int {
	r, _, _ := xList_CreateAdapterHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_取数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func XList_GetAdapter(hEle int) int {
	r, _, _ := xList_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄.
//
// hEle: 元素句柄.
func XList_GetAdapterHeader(hEle int) int {
	r, _, _ := xList_GetAdapterHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项模板文件, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func XList_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xList_SetItemTemplateXML.Call(uintptr(hEle), common.StrPtr(pXmlFile))
	return r != 0
}

// 列表_置项模板从字符串, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func XList_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {
	r, _, _ := xList_SetItemTemplateXMLFromString.Call(uintptr(hEle), XC_wtoa(pStringXML))
	return r != 0
}

// 列表_置项模板, 设置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func XList_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xList_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return r != 0
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// nTempItemID: 模板项itemID.
func XList_GetTemplateObject(hEle int, iItem int, iSubItem int, nTempItemID int) int {
	r, _, _ := xList_GetTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取对象所在行, 获取当前对象所在模板实例, 属于列表中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func XList_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xList_GetItemIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func XList_GetHeaderTemplateObject(hEle int, iItem int, nTempItemID int) int {
	r, _, _ := xList_GetHeaderTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取列表头对象所在行, 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄.
func XList_GetHeaderItemIndexFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xList_GetHeaderItemIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表_置列表头高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func XList_SetHeaderHeight(hEle int, height int) int {
	r, _, _ := xList_SetHeaderHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 列表_取列表头高度.
//
// hEle: 元素句柄.
func XList_GetHeaderHeight(hEle int) int {
	r, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))
	return int(r)
}

// 列表_取可视行范围.
//
// hEle: 元素句柄.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func XList_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {
	r, _, _ := xList_GetVisibleRowRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piStart)), uintptr(unsafe.Pointer(piEnd)))
	return int(r)
}

// 列表_置项默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func XList_SetItemHeightDefault(hEle int, nHeight int32, nSelHeight int32) int {
	r, _, _ := xList_SetItemHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表_取项默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func XList_GetItemHeightDefault(hEle int, pHeight *int32, pSelHeight *int32) int {
	r, _, _ := xList_GetItemHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 行间距大小.
func XList_SetRowSpace(hEle int, nSpace int) int {
	r, _, _ := xList_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 列表_取行间距.
//
// hEle: 元素句柄.
func XList_GetRowSpace(hEle int) int {
	r, _, _ := xList_GetRowSpace.Call(uintptr(hEle))
	return int(r)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定.
func XList_SetLockColumnLeft(hEle int, iColumn int) int {
	r, _, _ := xList_SetLockColumnLeft.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_置锁定列右侧.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func XList_SetLockColumnRight(hEle int, iColumn int) int {
	r, _, _ := xList_SetLockColumnRight.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// hEle: 元素句柄.
//
// bLock: 是否锁定.
func XList_SetLockRowBottom(hEle int, bLock bool) int {
	r, _, _ := xList_SetLockRowBottom.Call(uintptr(hEle), common.BoolPtr(bLock))
	return int(r)
}

// 列表_置锁定行底部重叠.
//
// hEle: 元素句柄.
//
// bOverlap: 是否重叠.
func XList_SetLockRowBottomOverlap(hEle int, bOverlap bool) int {
	r, _, _ := xList_SetLockRowBottomOverlap.Call(uintptr(hEle), common.BoolPtr(bOverlap))
	return int(r)
}

// 列表_测试点击项, 检测坐标点所在项.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func XList_HitTest(hEle int, pPt *POINT, piItem *int32, piSubItem *int32) bool {
	r, _, _ := xList_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(piItem)), uintptr(unsafe.Pointer(piSubItem)))
	return r != 0
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func XList_HitTestOffset(hEle int, pPt *POINT, piItem *int32, piSubItem *int32) bool {
	r, _, _ := xList_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(piItem)), uintptr(unsafe.Pointer(piSubItem)))
	return r != 0
}

// 列表_刷新项数据.
//
// hEle: 元素句柄.
func XList_RefreshData(hEle int) int {
	r, _, _ := xList_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func XList_RefreshItem(hEle int, iItem int) int {
	r, _, _ := xList_RefreshItem.Call(uintptr(hEle), uintptr(iItem))
	return int(r)
}

// 列表_添加列文本.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// pText: 文本.
func XList_AddColumnText(hEle int, nWidth int, pName string, pText string) int {
	r, _, _ := xList_AddColumnText.Call(uintptr(hEle), uintptr(nWidth), common.StrPtr(pName), common.StrPtr(pText))
	return int(r)
}

// 列表_添加列图片.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// hImage: 图片句柄.
func XList_AddColumnImage(hEle int, nWidth int, pName string, hImage int) int {
	r, _, _ := xList_AddColumnImage.Call(uintptr(hEle), uintptr(nWidth), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_添加项文本.
//
// hEle:.
//
// pText:.
func XList_AddItemText(hEle int, pText string) int {
	r, _, _ := xList_AddItemText.Call(uintptr(hEle), common.StrPtr(pText))
	return int(r)
}

// 列表_添加项文本扩展.
//
// hEle:.
//
// pName:.
//
// pText:.
func XList_AddItemTextEx(hEle int, pName string, pText string) int {
	r, _, _ := xList_AddItemTextEx.Call(uintptr(hEle), common.StrPtr(pName), common.StrPtr(pText))
	return int(r)
}

// 列表_添加项图片.
//
// hEle:.
//
// hImage:.
func XList_AddItemImage(hEle int, hImage int) int {
	r, _, _ := xList_AddItemImage.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 列表_添加项图片扩展.
//
// hEle:.
//
// pName:.
//
// hImage:.
func XList_AddItemImageEx(hEle int, pName string, hImage int) int {
	r, _, _ := xList_AddItemImageEx.Call(uintptr(hEle), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_插入项文本.
//
// hEle:.
//
// iItem:.
//
// pValue:.
func XList_InsertItemText(hEle int, iItem int, pValue string) int {
	r, _, _ := xList_InsertItemText.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pValue))
	return int(r)
}

// 列表_插入项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pValue:.
func XList_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {
	r, _, _ := xList_InsertItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return int(r)
}

// 列表_插入项图片.
//
// hEle:.
//
// iItem:.
//
// hImage:.
func XList_InsertItemImage(hEle int, iItem int, hImage int) int {
	r, _, _ := xList_InsertItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(hImage))
	return int(r)
}

// 列表_插入项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func XList_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {
	r, _, _ := xList_InsertItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return int(r)
}

// 列表_置项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func XList_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {
	r, _, _ := xList_SetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.StrPtr(pText))
	return r != 0
}

// 列表_置项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pText:.
func XList_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {
	r, _, _ := xList_SetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pText))
	return r != 0
}

// 列表_置项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func XList_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xList_SetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 列表_置项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func XList_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {
	r, _, _ := xList_SetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 列表_置项指数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func XList_SetItemInt(hEle int, iItem int, iColumn int, nValue int) bool {
	r, _, _ := xList_SetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 列表_置项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// nValue:.
func XList_SetItemIntEx(hEle int, iItem int, pName string, nValue int) bool {
	r, _, _ := xList_SetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(nValue))
	return r != 0
}

// 列表_置项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func XList_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {
	r, _, _ := xList_SetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表_置项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func XList_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {
	r, _, _ := xList_SetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表_取项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XList_GetItemText(hEle int, iItem int, iColumn int) string {
	r, _, _ := xList_GetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 列表_取项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func XList_GetItemTextEx(hEle int, iItem int, pName string) string {
	r, _, _ := xList_GetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 列表_取项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XList_GetItemImage(hEle int, iItem int, iColumn int) int {
	r, _, _ := xList_GetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 列表_取项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func XList_GetItemImageEx(hEle int, iItem int, pName string) int {
	r, _, _ := xList_GetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 列表_取项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func XList_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {
	r, _, _ := xList_GetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_取项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func XList_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {
	r, _, _ := xList_GetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_取项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func XList_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_取项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func XList_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(pName), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_删除项.
//
// hEle:.
//
// iItem:.
func XList_DeleteItem(hEle int, iItem int) bool {
	r, _, _ := xList_DeleteItem.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_删除项扩展.
//
// hEle:.
//
// iItem:.
//
// nCount:.
func XList_DeleteItemEx(hEle int, iItem int, nCount int) bool {
	r, _, _ := xList_DeleteItemEx.Call(uintptr(hEle), uintptr(iItem), uintptr(nCount))
	return r != 0
}

// 列表_删除项全部.
//
// hEle:.
func XList_DeleteItemAll(hEle int) int {
	r, _, _ := xList_DeleteItemAll.Call(uintptr(hEle))
	return int(r)
}

// 列表_删除列全部AD.
//
// hEle:.
func XList_DeleteColumnAll_AD(hEle int) int {
	r, _, _ := xList_DeleteColumnAll_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表_取项数量AD.
//
// hEle:.
func XList_GetCount_AD(hEle int) int {
	r, _, _ := xList_GetCount_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表_取列数量AD.
//
// hEle:.
func XList_GetCountColumn_AD(hEle int) int {
	r, _, _ := xList_GetCountColumn_AD.Call(uintptr(hEle))
	return int(r)
}

// 列表_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XList_SetSplitLineColor(hEle int, color int) int {
	r, _, _ := xList_SetSplitLineColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 列表_置项高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func XList_SetItemHeight(hEle int, iRow int, nHeight, nSelHeight int32) int {
	r, _, _ := xList_SetItemHeight.Call(uintptr(hEle), uintptr(iRow), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表_取项高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func XList_GetItemHeight(hEle int, iRow int, pHeight, pSelHeight *int32) int {
	r, _, _ := xList_GetItemHeight.Call(uintptr(hEle), uintptr(iRow), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func XList_SetDragRectColor(hEle int, color, width int) int {
	r, _, _ := xList_SetDragRectColor.Call(uintptr(hEle), uintptr(color), uintptr(width))
	return int(r)
}

// 列表_取项模板. 返回列表项模板句柄.
//
// hEle: 元素句柄.
func XList_GetItemTemplate(hEle int) int {
	r, _, _ := xList_GetItemTemplate.Call(uintptr(hEle))
	return int(r)
}

// 列表_取项模板列表头. 返回列表头项模板句柄.
//
// hEle: 元素句柄.
func XList_GetItemTemplateHeader(hEle int) int {
	r, _, _ := xList_GetItemTemplateHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_刷新项数据列表头.
//
// hEle: 元素句柄.
func XList_RefreshDataHeader(hEle int) int {
	r, _, _ := xList_RefreshDataHeader.Call(uintptr(hEle))
	return int(r)
}

// 列表_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func XList_SetItemTemplateXMLFromMem(hEle int, data []byte) bool {
	r, _, _ := xList_SetItemTemplateXMLFromMem.Call(uintptr(hEle), common.ByteSliceDataPtr(&data), uintptr(len(data)))
	return r != 0
}

// 列表_置项模板从资源ZIP. 从RC资源ZIP加载.
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
func XList_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xList_SetItemTemplateXMLFromZipRes.Call(uintptr(hEle), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}
