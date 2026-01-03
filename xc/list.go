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
func XList_Create(x, y, cx, cy int32, hParent int) int {
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
func XList_AddColumn(hEle int, width int32) int32 {
	r, _, _ := xList_AddColumn.Call(uintptr(hEle), uintptr(width))
	return int32(r)
}

// 列表_插入列, 返回插入位置索引.
//
// hEle: 元素句柄.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func XList_InsertColumn(hEle int, width int32, iItem int32) int32 {
	r, _, _ := xList_InsertColumn.Call(uintptr(hEle), uintptr(width), uintptr(iItem))
	return int32(r)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableMultiSel(hEle int, bEnable bool) {
	xList_EnableMultiSel.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableDragChangeColumnWidth(hEle int, bEnable bool) {
	xList_EnableDragChangeColumnWidth.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 列表_启用垂直滚动条顶部对齐.
//
// hEle: 元素句柄.
//
// bTop: 是否启用.
func XList_EnableVScrollBarTop(hEle int, bTop bool) {
	xList_EnableVScrollBarTop.Call(uintptr(hEle), common.BoolPtr(bTop))
}

// 列表_启用行背景铺满.
//
// hEle: 元素句柄.
//
// bFull: 是否启用.
func XList_EnableRowBkFull(hEle int, bFull bool) {
	xList_EnableRowBkFull.Call(uintptr(hEle), common.BoolPtr(bFull))
}

// 列表_启用固定行高.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableFixedRowHeight(hEle int, bEnable bool) {
	xList_EnableFixedRowHeight.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 列表_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableTemplateReuse(hEle int, bEnable bool) {
	xList_EnableTemplateReuse.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 列表_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XList_EnableVirtualTable(hEle int, bEnable bool) {
	xList_EnableVirtualTable.Call(uintptr(hEle), common.BoolPtr(bEnable))
}

// 列表_置虚表行数.
//
// hEle: 元素句柄.
//
// nRowCount: 行数.
func XList_SetVirtualRowCount(hEle int, nRowCount int32) {
	xList_SetVirtualRowCount.Call(uintptr(hEle), uintptr(nRowCount))
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
func XList_SetSort(hEle int, iColumn int32, iColumnAdapter int32, bEnable bool) {
	xList_SetSort.Call(uintptr(hEle), uintptr(iColumn), uintptr(iColumnAdapter), common.BoolPtr(bEnable))
}

// 列表_置行背景绘制标志. 设置是否绘制指定状态下行的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位, xcc.List_DrawItemBk_Flag_.
func XList_SetDrawRowBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) {
	xList_SetDrawRowBkFlags.Call(uintptr(hEle), uintptr(nFlags))
}

// 列表_置列宽.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func XList_SetColumnWidth(hEle int, iItem, width int32) {
	xList_SetColumnWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
}

// 列表_置列最小宽度.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func XList_SetColumnMinWidth(hEle int, iItem, width int32) {
	xList_SetColumnMinWidth.Call(uintptr(hEle), uintptr(iItem), uintptr(width))
}

// 列表_置列宽度固定.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func XList_SetColumnWidthFixed(hEle int, iColumn int32, bFixed bool) {
	xList_SetColumnWidthFixed.Call(uintptr(hEle), uintptr(iColumn), common.BoolPtr(bFixed))
}

// 列表_取列宽度.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func XList_GetColumnWidth(hEle int, iColumn int32) int32 {
	r, _, _ := xList_GetColumnWidth.Call(uintptr(hEle), uintptr(iColumn))
	return int32(r)
}

// 列表_取列数量.
//
// hEle: 元素句柄.
func XList_GetColumnCount(hEle int) int32 {
	r, _, _ := xList_GetColumnCount.Call(uintptr(hEle))
	return int32(r)
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
func XList_SetItemData(hEle int, iItem int32, iSubItem int32, data int) bool {
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
func XList_GetItemData(hEle int, iItem, iSubItem int32) int {
	r, _, _ := xList_GetItemData.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem))
	return int(r)
}

// 列表_置选择行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XList_SetSelectRow(hEle int, iRow int32) bool {
	r, _, _ := xList_SetSelectRow.Call(uintptr(hEle), uintptr(iRow))
	return r != 0
}

// 列表_取选择行, 返回行索引.
//
// hEle: 元素句柄.
func XList_GetSelectRow(hEle int) int32 {
	r, _, _ := xList_GetSelectRow.Call(uintptr(hEle))
	return int32(r)
}

// 列表_取选择行数量, 获取选择行数量.
//
// hEle: 元素句柄.
func XList_GetSelectRowCount(hEle int) int32 {
	r, _, _ := xList_GetSelectItemRow.Call(uintptr(hEle))
	return int32(r)
}

// 列表_添加选择行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XList_AddSelectRow(hEle int, iRow int32) bool {
	r, _, _ := xList_AddSelectRow.Call(uintptr(hEle), uintptr(iRow))
	return r != 0
}

// 列表_置选择全部, 选择全部行.
//
// hEle: 元素句柄.
func XList_SetSelectAll(hEle int) {
	xList_SetSelectAll.Call(uintptr(hEle))
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// hEle: 元素句柄.
//
// pArray: 接收行索引数组.
//
// nArraySize: 数组大小.
func XList_GetSelectAll(hEle int, pArray *[]int32, nArraySize int32) int32 {
	if nArraySize < 1 {
		return 0
	}
	*pArray = make([]int32, nArraySize)
	r, _, _ := xList_GetSelectAll.Call(uintptr(hEle), uintptr(unsafe.Pointer(&(*pArray)[0])), uintptr(nArraySize))
	return int32(r)
}

// 列表_显示指定行, 滚动视图让指定行可见.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XList_VisibleRow(hEle int, iRow int32) {
	xList_VisibleRow.Call(uintptr(hEle), uintptr(iRow))
}

// 列表_取消选择行, 取消选择指定行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XList_CancelSelectRow(hEle int, iRow int32) bool {
	r, _, _ := xList_CancelSelectRow.Call(uintptr(hEle), uintptr(iRow))
	return r != 0
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
//
// hEle: 元素句柄.
func XList_CancelSelectAll(hEle int) {
	xList_CancelSelectAll.Call(uintptr(hEle))
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
func XList_DeleteColumn(hEle int, iItem int32) bool {
	r, _, _ := xList_DeleteColumn.Call(uintptr(hEle), uintptr(iItem))
	return r != 0
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
//
// hEle: 元素句柄.
func XList_DeleteColumnAll(hEle int) {
	xList_DeleteColumnAll.Call(uintptr(hEle))
}

// 列表_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdTable.
func XList_BindAdapter(hEle int, hAdapter int) {
	xList_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
}

// 列表_列表头绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdMap.
func XList_BindAdapterHeader(hEle int, hAdapter int) {
	xList_BindAdapterHeader.Call(uintptr(hEle), uintptr(hAdapter))
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func XList_CreateAdapter(hEle int, colExtend_count int32) int {
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
func XList_GetTemplateObject(hEle int, iItem, iSubItem, nTempItemID int32) int {
	r, _, _ := xList_GetTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(iSubItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取对象所在行. 获取当前对象所在模板实例, 属于列表中哪一行. 成功返回行索引, 否则返回 xcc.XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func XList_GetRowIndexFromHXCGUI(hEle int, hXCGUI int) int32 {
	r, _, _ := xList_GetRowIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int32(r)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func XList_GetHeaderTemplateObject(hEle int, iItem, nTempItemID int32) int {
	r, _, _ := xList_GetHeaderTemplateObject.Call(uintptr(hEle), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表_取列表头对象所在列. 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个列. 成功返回列索引, 否则返回 xcc.XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄.
func XList_GetHeaderColumnIndexFromHXCGUI(hEle int, hXCGUI int) int32 {
	r, _, _ := xList_GetHeaderColumnIndexFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int32(r)
}

// 列表_置列表头高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func XList_SetHeaderHeight(hEle int, height int32) {
	xList_SetHeaderHeight.Call(uintptr(hEle), uintptr(height))
}

// 列表_取列表头高度.
//
// hEle: 元素句柄.
func XList_GetHeaderHeight(hEle int) int32 {
	r, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))
	return int32(r)
}

// 列表_取可视行范围.
//
// hEle: 元素句柄.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func XList_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) {
	xList_GetVisibleRowRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piStart)), uintptr(unsafe.Pointer(piEnd)))
}

// 列表_置行默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func XList_SetRowHeightDefault(hEle int, nHeight int32, nSelHeight int32) {
	xList_SetRowHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
}

// 列表_取行默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func XList_GetRowHeightDefault(hEle int, pHeight *int32, pSelHeight *int32) {
	xList_GetRowHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
}

// 列表_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 行间距大小.
func XList_SetRowSpace(hEle int, nSpace int32) {
	xList_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
}

// 列表_取行间距.
//
// hEle: 元素句柄.
func XList_GetRowSpace(hEle int) int32 {
	r, _, _ := xList_GetRowSpace.Call(uintptr(hEle))
	return int32(r)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定.
func XList_SetLockColumnLeft(hEle int, iColumn int32) {
	xList_SetLockColumnLeft.Call(uintptr(hEle), uintptr(iColumn))
}

// 列表_置锁定列右侧.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func XList_SetLockColumnRight(hEle int, iColumn int32) {
	xList_SetLockColumnRight.Call(uintptr(hEle), uintptr(iColumn))
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// hEle: 元素句柄.
//
// bLock: 是否锁定.
func XList_SetLockRowBottom(hEle int, bLock bool) {
	xList_SetLockRowBottom.Call(uintptr(hEle), common.BoolPtr(bLock))
}

// 列表_置锁定行底部重叠.
//
// hEle: 元素句柄.
//
// bOverlap: 是否重叠.
func XList_SetLockRowBottomOverlap(hEle int, bOverlap bool) {
	xList_SetLockRowBottomOverlap.Call(uintptr(hEle), common.BoolPtr(bOverlap))
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
func XList_RefreshData(hEle int) {
	xList_RefreshData.Call(uintptr(hEle))
}

// 列表_刷新指定行. 修改数据后, 刷新指定行模板, 以便更新数据到模板(如果当前行可见).
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XList_RefreshRow(hEle int, iRow int32) {
	xList_RefreshRow.Call(uintptr(hEle), uintptr(iRow))
}

// 列表_添加列文本.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// name: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen.
//
// text: 文本.
func XList_AddColumnText(hEle int, nWidth int32, name string, text string) int32 {
	r, _, _ := xList_AddColumnText.Call(uintptr(hEle), uintptr(nWidth), common.StrPtr(name), common.StrPtr(text))
	return int32(r)
}

// 列表_添加列图片.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// name: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen.
//
// hImage: 图片句柄.
func XList_AddColumnImage(hEle int, nWidth int32, name string, hImage int) int32 {
	r, _, _ := xList_AddColumnImage.Call(uintptr(hEle), uintptr(nWidth), common.StrPtr(name), uintptr(hImage))
	return int32(r)
}

// 列表_添加行文本.
//
// hEle:.
//
// text:.
func XList_AddRowText(hEle int, text string) int32 {
	r, _, _ := xList_AddRowText.Call(uintptr(hEle), common.StrPtr(text))
	return int32(r)
}

// 列表_添加行文本扩展.
//
// hEle:.
//
// name:.
//
// text:.
func XList_AddRowTextEx(hEle int, name string, text string) int32 {
	r, _, _ := xList_AddRowTextEx.Call(uintptr(hEle), common.StrPtr(name), common.StrPtr(text))
	return int32(r)
}

// 列表_添加行图片.
//
// hEle:.
//
// hImage:.
func XList_AddRowImage(hEle int, hImage int) int32 {
	r, _, _ := xList_AddRowImage.Call(uintptr(hEle), uintptr(hImage))
	return int32(r)
}

// 列表_添加行图片扩展.
//
// hEle:.
//
// name:.
//
// hImage:.
func XList_AddRowImageEx(hEle int, name string, hImage int) int32 {
	r, _, _ := xList_AddRowImageEx.Call(uintptr(hEle), common.StrPtr(name), uintptr(hImage))
	return int32(r)
}

// 列表_插入行文本.
//
// hEle:.
//
// iRow:.
//
// pValue:.
func XList_InsertRowText(hEle int, iRow int32, pValue string) int32 {
	r, _, _ := xList_InsertRowText.Call(uintptr(hEle), uintptr(iRow), common.StrPtr(pValue))
	return int32(r)
}

// 列表_插入行文本扩展.
//
// hEle:.
//
// iRow:.
//
// name:.
//
// pValue:.
func XList_InsertRowTextEx(hEle int, iRow int32, name string, pValue string) int32 {
	r, _, _ := xList_InsertRowTextEx.Call(uintptr(hEle), uintptr(iRow), common.StrPtr(name), common.StrPtr(pValue))
	return int32(r)
}

// 列表_插入行图片.
//
// hEle:.
//
// iRow:.
//
// hImage:.
func XList_InsertRowImage(hEle int, iRow int32, hImage int) int32 {
	r, _, _ := xList_InsertRowImage.Call(uintptr(hEle), uintptr(iRow), uintptr(hImage))
	return int32(r)
}

// 列表_插入行图片扩展.
//
// hEle:.
//
// iRow:.
//
// name:.
//
// hImage:.
func XList_InsertRowImageEx(hEle int, iRow int32, name string, hImage int) int32 {
	r, _, _ := xList_InsertRowImageEx.Call(uintptr(hEle), uintptr(iRow), common.StrPtr(name), uintptr(hImage))
	return int32(r)
}

// 列表_置项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// text:.
func XList_SetItemText(hEle int, iItem, iColumn int32, text string) bool {
	r, _, _ := xList_SetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.StrPtr(text))
	return r != 0
}

// 列表_置项文本扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// text:.
func XList_SetItemTextEx(hEle int, iItem int32, name string, text string) bool {
	r, _, _ := xList_SetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), common.StrPtr(text))
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
func XList_SetItemImage(hEle int, iItem, iColumn int32, hImage int) bool {
	r, _, _ := xList_SetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 列表_置项图片扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// hImage:.
func XList_SetItemImageEx(hEle int, iItem int32, name string, hImage int) bool {
	r, _, _ := xList_SetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), uintptr(hImage))
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
func XList_SetItemInt(hEle int, iItem, iColumn int32, nValue int32) bool {
	r, _, _ := xList_SetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(nValue))
	return r != 0
}

// 列表_置项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// nValue:.
func XList_SetItemIntEx(hEle int, iItem int32, name string, nValue int32) bool {
	r, _, _ := xList_SetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), uintptr(nValue))
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
func XList_SetItemFloat(hEle int, iItem, iColumn int32, fFloat float32) bool {
	r, _, _ := xList_SetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表_置项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// fFloat:.
func XList_SetItemFloatEx(hEle int, iItem int32, name string, fFloat float32) bool {
	r, _, _ := xList_SetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), common.Float32Ptr(fFloat))
	return r != 0
}

// 列表_取项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XList_GetItemText(hEle int, iItem, iColumn int32) string {
	r, _, _ := xList_GetItemText.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 列表_取项文本扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
func XList_GetItemTextEx(hEle int, iItem int32, name string) string {
	r, _, _ := xList_GetItemTextEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name))
	return common.UintPtrToString(r)
}

// 列表_取项图片. 返回图片句柄.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func XList_GetItemImage(hEle int, iItem, iColumn int32) int {
	r, _, _ := xList_GetItemImage.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 列表_取项图片扩展. 返回图片句柄.
//
// hEle:.
//
// iItem:.
//
// name:.
func XList_GetItemImageEx(hEle int, iItem int32, name string) int {
	r, _, _ := xList_GetItemImageEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name))
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
func XList_GetItemInt(hEle int, iItem int32, iColumn int32, pOutValue *int32) bool {
	r, _, _ := xList_GetItemInt.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_取项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// pOutValue:.
func XList_GetItemIntEx(hEle int, iItem int32, name string, pOutValue *int32) bool {
	r, _, _ := xList_GetItemIntEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), uintptr(unsafe.Pointer(pOutValue)))
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
func XList_GetItemFloat(hEle int, iItem, iColumn int32, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloat.Call(uintptr(hEle), uintptr(iItem), uintptr(iColumn), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_取项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// name:.
//
// pOutValue:.
func XList_GetItemFloatEx(hEle int, iItem int32, name string, pOutValue *float32) bool {
	r, _, _ := xList_GetItemFloatEx.Call(uintptr(hEle), uintptr(iItem), common.StrPtr(name), uintptr(unsafe.Pointer(pOutValue)))
	return r != 0
}

// 列表_删除行.
//
// hEle:.
//
// iRow:.
func XList_DeleteRow(hEle int, iRow int32) bool {
	r, _, _ := xList_DeleteRow.Call(uintptr(hEle), uintptr(iRow))
	return r != 0
}

// 列表_删除行扩展.
//
// hEle:.
//
// iRow:.
//
// nCount:.
func XList_DeleteRowEx(hEle int, iRow, nCount int32) bool {
	r, _, _ := xList_DeleteRowEx.Call(uintptr(hEle), uintptr(iRow), uintptr(nCount))
	return r != 0
}

// 列表_删除行全部.
//
// hEle:.
func XList_DeleteRowAll(hEle int) {
	xList_DeleteRowAll.Call(uintptr(hEle))
}

// 列表_删除列全部AD.
//
// hEle:.
func XList_DeleteColumnAll_AD(hEle int) {
	xList_DeleteColumnAll_AD.Call(uintptr(hEle))
}

// 列表_取项数量AD.
//
// hEle:.
func XList_GetCount_AD(hEle int) int32 {
	r, _, _ := xList_GetCount_AD.Call(uintptr(hEle))
	return int32(r)
}

// 列表_取列数量AD.
//
// hEle:.
func XList_GetCountColumn_AD(hEle int) int32 {
	r, _, _ := xList_GetCountColumn_AD.Call(uintptr(hEle))
	return int32(r)
}

// 列表_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: xc.RGBA 颜色值.
func XList_SetSplitLineColor(hEle int, color uint32) {
	xList_SetSplitLineColor.Call(uintptr(hEle), uintptr(color))
}

// 列表_置行高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func XList_SetRowHeight(hEle int, iRow, nHeight, nSelHeight int32) {
	xList_SetRowHeight.Call(uintptr(hEle), uintptr(iRow), uintptr(nHeight), uintptr(nSelHeight))
}

// 列表_取行高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func XList_GetRowHeight(hEle int, iRow int32, pHeight, pSelHeight *int32) {
	xList_GetRowHeight.Call(uintptr(hEle), uintptr(iRow), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
}

// 列表_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: xc.RGBA 颜色值.
//
// width: 线宽度.
func XList_SetDragRectColor(hEle int, color uint32, width int32) {
	xList_SetDragRectColor.Call(uintptr(hEle), uintptr(color), uintptr(width))
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
func XList_RefreshDataHeader(hEle int) {
	xList_RefreshDataHeader.Call(uintptr(hEle))
}

// 列表_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func XList_SetItemTemplateXMLFromMem(hEle int, data []byte) bool {
	r, _, _ := xList_SetItemTemplateXMLFromMem.Call(uintptr(hEle), common.ByteSliceDataPtr(&data), uintptr(int32(len(data))))
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

// 列表_添加列文本2, 返回列索引. 简化版本.
//
// hEle: 元素句柄.
//
// nWidth: 列宽度.
//
// text: 标题文本.
func XList_AddColumnText2(hEle int, nWidth int32, text string) int32 {
	r, _, _ := xList_AddColumnText2.Call(uintptr(hEle), uintptr(nWidth), common.StrPtr(text))
	return int32(r)
}

// 列表_添加列图片2, 返回列索引. 简化版本.
//
// hEle: 元素句柄.
//
// nWidth: 列宽度.
//
// hImage: 图片句柄.
func XList_AddColumnImage2(hEle int, nWidth int32, hImage int) int32 {
	r, _, _ := xList_AddColumnImage2.Call(uintptr(hEle), uintptr(nWidth), uintptr(hImage))
	return int32(r)
}

// 列表_创建数据适配器2. 创建数据适配器，根据绑定的项模板初始化数据适配器的列(字段名); 数据适配器存储数据, UI对象根据绑定的字段名显示数据适配器中对应的数据; 此接口是简化接口, 合并了 xc.XList_CreateAdapter() 和 xc.XList_CreateAdapterHeader() 接口;
//
// hEle: 元素句柄.
//
// col_extend_count: 列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段. 例如默认模板是1列, 但是数据有5列,那么列延伸填5.
func XList_CreateAdapters(hEle int, col_extend_count int32) bool {
	r, _, _ := xList_CreateAdapters.Call(uintptr(hEle), uintptr(col_extend_count))
	return r != 0
}
