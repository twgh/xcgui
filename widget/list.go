package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 列表.
type List struct {
	ScrollView
}

// 列表_创建, 创建列表元素.
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
func NewList(x, y, cx, cy int32, hParent int) *List {
	p := &List{}
	p.SetHandle(xc.XList_Create(x, y, cx, cy, hParent))
	return p
}

// 列表_创建Ex, 创建列表元素, 使用内置项模板, 自动创建数据适配器.
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
func NewListEx(x, y, cx, cy int32, hParent, col_extend_count int32) *List {
	p := &List{}
	p.SetHandle(xc.XList_CreateEx(x, y, cx, cy, hParent, col_extend_count))
	return p
}

// 从句柄创建对象.
func NewListByHandle(handle int) *List {
	p := &List{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewListByName(name string) *List {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewListByUID(nUID int32) *List {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewListByUIDName(name string) *List {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表_增加列, 返回位置索引.
//
// width: 列宽度.
func (l *List) AddColumn(width int32) int32 {
	return xc.XList_AddColumn(l.Handle, width)
}

// 列表_插入列, 返回插入位置索引.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func (l *List) InsertColumn(width, iItem int32) int32 {
	return xc.XList_InsertColumn(l.Handle, width, iItem)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// bEnable: 是否启用.
func (l *List) EnableMultiSel(bEnable bool) *List {
	xc.XList_EnableMultiSel(l.Handle, bEnable)
	return l
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// bEnable: 是否启用.
func (l *List) EnableDragChangeColumnWidth(bEnable bool) *List {
	xc.XList_EnableDragChangeColumnWidth(l.Handle, bEnable)
	return l
}

// 列表_启用垂直滚动条顶部对齐.
//
// bTop: 是否启用.
func (l *List) EnableVScrollBarTop(bTop bool) *List {
	xc.XList_EnableVScrollBarTop(l.Handle, bTop)
	return l
}

// 列表_启用行背景铺满.
//
// bFull: 是否启用.
func (l *List) EnableRowBkFull(bFull bool) *List {
	xc.XList_EnableRowBkFull(l.Handle, bFull)
	return l
}

// 列表_启用固定行高.
//
// bEnable: 是否启用.
func (l *List) EnableFixedRowHeight(bEnable bool) *List {
	xc.XList_EnableFixedRowHeight(l.Handle, bEnable)
	return l
}

// 列表_启用模板复用.
//
// bEnable: 是否启用.
func (l *List) EnableTemplateReuse(bEnable bool) *List {
	xc.XList_EnableTemplateReuse(l.Handle, bEnable)
	return l
}

// 列表_启用虚表.
//
// bEnable: 是否启用.
func (l *List) EnableVirtualTable(bEnable bool) *List {
	xc.XList_EnableVirtualTable(l.Handle, bEnable)
	return l
}

// 列表_置虚表行数.
//
// nRowCount: 行数.
func (l *List) SetVirtualRowCount(nRowCount int32) *List {
	xc.XList_SetVirtualRowCount(l.Handle, nRowCount)
	return l
}

// 列表_置排序, 设置排序属性.
//
// iColumn: 列索引.
//
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
//
// bEnable: 是否启用排序功能.
func (l *List) SetSort(iColumn int32, iColumnAdapter int32, bEnable bool) *List {
	xc.XList_SetSort(l.Handle, iColumn, iColumnAdapter, bEnable)
	return l
}

// 列表_置行背景绘制标志. 设置是否绘制指定状态下行的背景.
//
// nFlags: 标志位, xcc.List_DrawItemBk_Flag_.
func (l *List) SetDrawRowBkFlags(nFlags xcc.List_DrawItemBk_Flag_) *List {
	xc.XList_SetDrawRowBkFlags(l.Handle, nFlags)
	return l
}

// 列表_置列宽.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) SetColumnWidth(iItem, width int32) *List {
	xc.XList_SetColumnWidth(l.Handle, iItem, width)
	return l
}

// 列表_置列最小宽度.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) SetColumnMinWidth(iItem, width int32) *List {
	xc.XList_SetColumnMinWidth(l.Handle, iItem, width)
	return l
}

// 列表_置列宽度固定.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func (l *List) SetColumnWidthFixed(iColumn int32, bFixed bool) *List {
	xc.XList_SetColumnWidthFixed(l.Handle, iColumn, bFixed)
	return l
}

// 列表_取列宽度.
//
// iColumn: 列索引.
func (l *List) GetColumnWidth(iColumn int32) int32 {
	return xc.XList_GetColumnWidth(l.Handle, iColumn)
}

// 列表_取列数量.
func (l *List) GetColumnCount() int32 {
	return xc.XList_GetColumnCount(l.Handle)
}

// 列表_置项数据, 设置项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// data: 用户数据.
func (l *List) SetItemData(iItem, iSubItem int32, data int) bool {
	return xc.XList_SetItemData(l.Handle, iItem, iSubItem, data)
}

// 列表_取项数据, 获取项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
func (l *List) GetItemData(iItem, iSubItem int32) int {
	return xc.XList_GetItemData(l.Handle, iItem, iSubItem)
}

// 列表_置选择行.
//
// iRow: 行索引.
func (l *List) SetSelectRow(iRow int32) bool {
	return xc.XList_SetSelectRow(l.Handle, iRow)
}

// 列表_取选择行, 返回行索引.
func (l *List) GetSelectRow() int32 {
	return xc.XList_GetSelectRow(l.Handle)
}

// 列表_取选择行数量, 获取选择行数量.
func (l *List) GetSelectRowCount() int32 {
	return xc.XList_GetSelectRowCount(l.Handle)
}

// 列表_添加选择行.
//
// iRow: 行索引.
func (l *List) AddSelectRow(iRow int32) bool {
	return xc.XList_AddSelectRow(l.Handle, iRow)
}

// 列表_置选择全部, 选择全部行.
func (l *List) SetSelectAll() *List {
	xc.XList_SetSelectAll(l.Handle)
	return l
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// pArray: 接收行索引数组.
//
// nArraySize: 数组大小.
func (l *List) GetSelectAll(pArray *[]int32, nArraySize int32) int32 {
	return xc.XList_GetSelectAll(l.Handle, pArray, nArraySize)
}

// 列表_显示指定行, 滚动视图让指定行可见.
//
// iRow: 行索引.
func (l *List) VisibleRow(iRow int32) *List {
	xc.XList_VisibleRow(l.Handle, iRow)
	return l
}

// 列表_取消选择行, 取消选择指定行.
//
// iRow: 行索引.
func (l *List) CancelSelectRow(iRow int32) bool {
	return xc.XList_CancelSelectRow(l.Handle, iRow)
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
func (l *List) CancelSelectAll() *List {
	xc.XList_CancelSelectAll(l.Handle)
	return l
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄.
func (l *List) GetHeaderHELE() int {
	return xc.XList_GetHeaderHELE(l.Handle)
}

// 列表_删除列.
//
// iItem: 项索引.
func (l *List) DeleteColumn(iItem int32) bool {
	return xc.XList_DeleteColumn(l.Handle, iItem)
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
func (l *List) DeleteColumnAll() *List {
	xc.XList_DeleteColumnAll(l.Handle)
	return l
}

// 列表_绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *List) BindAdapter(hAdapter int) *List {
	xc.XList_BindAdapter(l.Handle, hAdapter)
	return l
}

// 列表_列表头绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdMap.
func (l *List) BindAdapterHeader(hAdapter int) *List {
	xc.XList_BindAdapterHeader(l.Handle, hAdapter)
	return l
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func (l *List) CreateAdapter(colExtend_count int32) int {
	return xc.XList_CreateAdapter(l.Handle, colExtend_count)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *List) CreateAdapterHeader() int {
	return xc.XList_CreateAdapterHeader(l.Handle)
}

// 列表_取数据适配器, 返回数据适配器句柄.
func (l *List) GetAdapter() int {
	return xc.XList_GetAdapter(l.Handle)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄.
func (l *List) GetAdapterHeader() int {
	return xc.XList_GetAdapterHeader(l.Handle)
}

// 列表_置项模板文件, 设置项布局模板文件.
//
// pXmlFile: 文件名.
func (l *List) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XList_SetItemTemplateXML(l.Handle, pXmlFile)
}

// 列表_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *List) SetItemTemplateXMLFromString(pStringXML string) bool {
	return xc.XList_SetItemTemplateXMLFromString(l.Handle, pStringXML)
}

// 列表_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *List) SetItemTemplate(hTemp int) bool {
	return xc.XList_SetItemTemplate(l.Handle, hTemp)
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// nTempItemID: 模板项itemID.
func (l *List) GetTemplateObject(iItem, iSubItem int32, nTempItemID int32) int {
	return xc.XList_GetTemplateObject(l.Handle, iItem, iSubItem, nTempItemID)
}

// 列表_取对象所在行. 获取当前对象所在模板实例, 属于列表中哪一行. 成功返回行索引, 否则返回 xcc.XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *List) GetRowIndexFromHXCGUI(hXCGUI int) int32 {
	return xc.XList_GetRowIndexFromHXCGUI(l.Handle, hXCGUI)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func (l *List) GetHeaderTemplateObject(iItem int32, nTempItemID int32) int {
	return xc.XList_GetHeaderTemplateObject(l.Handle, iItem, nTempItemID)
}

// 列表_取列表头对象所在列. 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个列. 成功返回列索引, 否则返回 xcc.XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (l *List) GetHeaderColumnIndexFromHXCGUI(hXCGUI int) int32 {
	return xc.XList_GetHeaderColumnIndexFromHXCGUI(l.Handle, hXCGUI)
}

// 列表_置列表头高度.
//
// height: 高度.
func (l *List) SetHeaderHeight(height int32) *List {
	xc.XList_SetHeaderHeight(l.Handle, height)
	return l
}

// 列表_取列表头高度.
func (l *List) GetHeaderHeight() int32 {
	return xc.XList_GetHeaderHeight(l.Handle)
}

// 列表_取可视行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *List) GetVisibleRowRange(piStart *int32, piEnd *int32) *List {
	xc.XList_GetVisibleRowRange(l.Handle, piStart, piEnd)
	return l
}

// 列表_置行默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) SetRowHeightDefault(nHeight int32, nSelHeight int32) *List {
	xc.XList_SetRowHeightDefault(l.Handle, nHeight, nSelHeight)
	return l
}

// 列表_取行默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) GetRowHeightDefault(pHeight *int32, pSelHeight *int32) *List {
	xc.XList_GetRowHeightDefault(l.Handle, pHeight, pSelHeight)
	return l
}

// 列表_置行间距.
//
// nSpace: 行间距大小.
func (l *List) SetRowSpace(nSpace int32) *List {
	xc.XList_SetRowSpace(l.Handle, nSpace)
	return l
}

// 列表_取行间距.
func (l *List) GetRowSpace() int32 {
	return xc.XList_GetRowSpace(l.Handle)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// iColumn: 列索引, -1代表不锁定.
func (l *List) SetLockColumnLeft(iColumn int32) *List {
	xc.XList_SetLockColumnLeft(l.Handle, iColumn)
	return l
}

// 列表_置锁定列右侧.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func (l *List) SetLockColumnRight(iColumn int32) *List {
	xc.XList_SetLockColumnRight(l.Handle, iColumn)
	return l
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// bLock: 是否锁定.
func (l *List) SetLockRowBottom(bLock bool) *List {
	xc.XList_SetLockRowBottom(l.Handle, bLock)
	return l
}

// 列表_置锁定行底部重叠.
//
// bOverlap: 是否重叠.
func (l *List) SetLockRowBottomOverlap(bOverlap bool) *List {
	xc.XList_SetLockRowBottomOverlap(l.Handle, bOverlap)
	return l
}

// 列表_测试点击项, 检测坐标点所在项.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) HitTest(pPt *xc.POINT, piItem *int32, piSubItem *int32) bool {
	return xc.XList_HitTest(l.Handle, pPt, piItem, piSubItem)
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) HitTestOffset(pPt *xc.POINT, piItem *int32, piSubItem *int32) bool {
	return xc.XList_HitTestOffset(l.Handle, pPt, piItem, piSubItem)
}

// 列表_刷新项数据.
func (l *List) RefreshData() *List {
	xc.XList_RefreshData(l.Handle)
	return l
}

// 列表_刷新指定行. 修改数据后, 刷新指定行模板, 以便更新数据到模板(如果当前行可见).
//
// iRow: 行索引.
func (l *List) RefreshRow(iRow int32) *List {
	xc.XList_RefreshRow(l.Handle, iRow)
	return l
}

// 列表_添加列文本.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen.
//
// pText: 文本.
func (l *List) AddColumnText(nWidth int32, pName string, pText string) int32 {
	return xc.XList_AddColumnText(l.Handle, nWidth, pName, pText)
}

// 列表_添加列图片.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen.
//
// hImage: 图片句柄.
func (l *List) AddColumnImage(nWidth int32, pName string, hImage int) int32 {
	return xc.XList_AddColumnImage(l.Handle, nWidth, pName, hImage)
}

// 列表_添加行文本.
//
// pText:.
func (l *List) AddRowText(pText string) int32 {
	return xc.XList_AddRowText(l.Handle, pText)
}

// 列表_添加行文本扩展.
//
// pName:.
//
// pText:.
func (l *List) AddRowTextEx(pName string, pText string) int32 {
	return xc.XList_AddRowTextEx(l.Handle, pName, pText)
}

// 列表_添加行图片.
//
// hImage:.
func (l *List) AddRowImage(hImage int) int32 {
	return xc.XList_AddRowImage(l.Handle, hImage)
}

// 列表_添加行图片扩展.
//
// pName:.
//
// hImage:.
func (l *List) AddRowImageEx(pName string, hImage int) int32 {
	return xc.XList_AddRowImageEx(l.Handle, pName, hImage)
}

// 列表_插入行文本.
//
// iRow:.
//
// pValue:.
func (l *List) InsertRowText(iRow int32, pValue string) int32 {
	return xc.XList_InsertRowText(l.Handle, iRow, pValue)
}

// 列表_插入行文本扩展.
//
// iRow:.
//
// pName:.
//
// pValue:.
func (l *List) InsertRowTextEx(iRow int32, pName string, pValue string) int32 {
	return xc.XList_InsertRowTextEx(l.Handle, iRow, pName, pValue)
}

// 列表_插入行图片.
//
// iRow:.
//
// hImage:.
func (l *List) InsertRowImage(iRow int32, hImage int) int32 {
	return xc.XList_InsertRowImage(l.Handle, iRow, hImage)
}

// 列表_插入行图片扩展.
//
// iRow:.
//
// pName:.
//
// hImage:.
func (l *List) InsertRowImageEx(iRow int32, pName string, hImage int) int32 {
	return xc.XList_InsertRowImageEx(l.Handle, iRow, pName, hImage)
}

// 列表_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *List) SetItemText(iItem, iColumn int32, pText string) bool {
	return xc.XList_SetItemText(l.Handle, iItem, iColumn, pText)
}

// 列表_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *List) SetItemTextEx(iItem int32, pName string, pText string) bool {
	return xc.XList_SetItemTextEx(l.Handle, iItem, pName, pText)
}

// 列表_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *List) SetItemImage(iItem, iColumn int32, hImage int) bool {
	return xc.XList_SetItemImage(l.Handle, iItem, iColumn, hImage)
}

// 列表_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *List) SetItemImageEx(iItem int32, pName string, hImage int) bool {
	return xc.XList_SetItemImageEx(l.Handle, iItem, pName, hImage)
}

// 列表_置项指数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *List) SetItemInt(iItem, iColumn int32, nValue int32) bool {
	return xc.XList_SetItemInt(l.Handle, iItem, iColumn, nValue)
}

// 列表_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *List) SetItemIntEx(iItem int32, pName string, nValue int32) bool {
	return xc.XList_SetItemIntEx(l.Handle, iItem, pName, nValue)
}

// 列表_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *List) SetItemFloat(iItem int32, iColumn int32, fFloat float32) bool {
	return xc.XList_SetItemFloat(l.Handle, iItem, iColumn, fFloat)
}

// 列表_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *List) SetItemFloatEx(iItem int32, pName string, fFloat float32) bool {
	return xc.XList_SetItemFloatEx(l.Handle, iItem, pName, fFloat)
}

// 列表_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *List) GetItemText(iItem int32, iColumn int32) string {
	return xc.XList_GetItemText(l.Handle, iItem, iColumn)
}

// 列表_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *List) GetItemTextEx(iItem int32, pName string) string {
	return xc.XList_GetItemTextEx(l.Handle, iItem, pName)
}

// 列表_取项图片. 返回图片句柄.
//
// iItem:.
//
// iColumn:.
func (l *List) GetItemImage(iItem int32, iColumn int32) int {
	return xc.XList_GetItemImage(l.Handle, iItem, iColumn)
}

// 列表_取项图片扩展. 返回图片句柄.
//
// iItem:.
//
// pName:.
func (l *List) GetItemImageEx(iItem int32, pName string) int {
	return xc.XList_GetItemImageEx(l.Handle, iItem, pName)
}

// 列表_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) GetItemInt(iItem, iColumn int32, pOutValue *int32) bool {
	return xc.XList_GetItemInt(l.Handle, iItem, iColumn, pOutValue)
}

// 列表_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) GetItemIntEx(iItem int32, pName string, pOutValue *int32) bool {
	return xc.XList_GetItemIntEx(l.Handle, iItem, pName, pOutValue)
}

// 列表_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) GetItemFloat(iItem, iColumn int32, pOutValue *float32) bool {
	return xc.XList_GetItemFloat(l.Handle, iItem, iColumn, pOutValue)
}

// 列表_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) GetItemFloatEx(iItem int32, pName string, pOutValue *float32) bool {
	return xc.XList_GetItemFloatEx(l.Handle, iItem, pName, pOutValue)
}

// 列表_删除行.
//
// iRow:.
func (l *List) DeleteRow(iRow int32) bool {
	return xc.XList_DeleteRow(l.Handle, iRow)
}

// 列表_删除行扩展.
//
// iRow:.
//
// nCount:.
func (l *List) DeleteRowEx(iRow int32, nCount int32) bool {
	return xc.XList_DeleteRowEx(l.Handle, iRow, nCount)
}

// 列表_删除行全部.
func (l *List) DeleteRowAll() *List {
	xc.XList_DeleteRowAll(l.Handle)
	return l
}

// 列表_删除列全部AD.
func (l *List) DeleteColumnAll_AD() *List {
	xc.XList_DeleteColumnAll_AD(l.Handle)
	return l
}

// 列表_取项数量AD.
func (l *List) GetCount_AD() int32 {
	return xc.XList_GetCount_AD(l.Handle)
}

// 列表_取列数量AD.
func (l *List) GetCountColumn_AD() int32 {
	return xc.XList_GetCountColumn_AD(l.Handle)
}

// 列表_置分割线颜色.
//
// color: xc.RGBA 颜色值.
func (l *List) SetSplitLineColor(color int) *List {
	xc.XList_SetSplitLineColor(l.Handle, color)
	return l
}

// 列表_置行高度.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) SetRowHeight(iRow int32, nHeight, nSelHeight int32) *List {
	xc.XList_SetRowHeight(l.Handle, iRow, nHeight, nSelHeight)
	return l
}

// 列表_取行高度.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) GetRowHeight(iRow int32, pHeight, pSelHeight *int32) *List {
	xc.XList_GetRowHeight(l.Handle, iRow, pHeight, pSelHeight)
	return l
}

// 列表_置拖动矩形颜色.
//
// color: xc.RGBA 颜色值.
//
// width: 线宽度.
func (l *List) SetDragRectColor(color int, width int32) *List {
	xc.XList_SetDragRectColor(l.Handle, color, width)
	return l
}

// 列表_取项模板. 返回列表项模板句柄.
func (l *List) GetItemTemplate() int {
	return xc.XList_GetItemTemplate(l.Handle)
}

// 列表_取项模板列表头. 返回列表头项模板句柄.
func (l *List) GetItemTemplateHeader() int {
	return xc.XList_GetItemTemplateHeader(l.Handle)
}

// 列表_刷新项数据列表头.
func (l *List) RefreshDataHeader() *List {
	xc.XList_RefreshDataHeader(l.Handle)
	return l
}

// 列表_置项模板从内存.
//
// data: 模板数据.
func (l *List) SetItemTemplateXMLFromMem(data []byte) bool {
	return xc.XList_SetItemTemplateXMLFromMem(l.Handle, data)
}

// 列表_置项模板从资源ZIP. 从RC资源ZIP加载.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *List) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XList_SetItemTemplateXMLFromZipRes(l.Handle, id, pFileName, pPassword, hModule)
}

// 列表_添加列文本2, 返回列索引. 简化版本.
//
// nWidth: 列宽度.
//
// pText: 标题文本.
func (l *List) AddColumnText2(nWidth int32, pText string) int32 {
	return xc.XList_AddColumnText2(l.Handle, nWidth, pText)
}

// 列表_添加列图片2, 返回列索引. 简化版本.
//
// nWidth: 列宽度.
//
// hImage: 图片句柄.
func (l *List) AddColumnImage2(nWidth int32, hImage int) int32 {
	return xc.XList_AddColumnImage2(l.Handle, nWidth, hImage)
}

// 列表_创建数据适配器2. 创建数据适配器，根据绑定的项模板初始化数据适配器的列(字段名); 数据适配器存储数据, UI对象根据绑定的字段名显示数据适配器中对应的数据; 此接口是简化接口, 合并了 xc.XList_CreateAdapter() 和 xc.XList_CreateAdapterHeader() 接口;
//
// col_extend_count: 列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段. 例如默认模板是1列, 但是数据有5列,那么列延伸填5.
func (l *List) CreateAdapters(col_extend_count int32) bool {
	return xc.XList_CreateAdapters(l.Handle, col_extend_count)
}

/*
// //////////////////////////////////// 以下是事件 //////////////////////////////////////
*/

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE func(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE1 func(hEle int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END func(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END1 func(hEle int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY func(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY1 func(hEle int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int
type XE_LIST_TEMP_ADJUST_COORDINATE func(pItem *xc.List_Item_, pbHandled *bool) int                          // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.List_Item_, pbHandled *bool) int               // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_DRAWITEM func(hDraw int, pItem *xc.List_Item_, pbHandled *bool) int                             // 列表元素,绘制项.
type XE_LIST_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.List_Item_, pbHandled *bool) int                  // 列表元素,绘制项.
type XE_LIST_SELECT func(iItem int32, pbHandled *bool) int                                                   // 列表元素,项选择事件.
type XE_LIST_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                                        // 列表元素,项选择事件.
type XE_LIST_HEADER_DRAWITEM func(hDraw int, pItem *xc.List_Header_Item_, pbHandled *bool) int               // 列表元素绘制列表头项.
type XE_LIST_HEADER_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.List_Header_Item_, pbHandled *bool) int    // 列表元素绘制列表头项.
type XE_LIST_HEADER_CLICK func(iItem int32, pbHandled *bool) int                                             // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_CLICK1 func(hEle int, iItem int32, pbHandled *bool) int                                  // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_WIDTH_CHANGE func(iItem int32, nWidth int32, pbHandled *bool) int                        // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_WIDTH_CHANGE1 func(hEle int, iItem int32, nWidth int32, pbHandled *bool) int             // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_TEMP_CREATE func(pItem *xc.List_Header_Item_, pbHandled *bool) int                       // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE_END func(pItem *xc.List_Header_Item_, pbHandled *bool) int                   // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_CREATE_END1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int        // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_DESTROY func(pItem *xc.List_Header_Item_, pbHandled *bool) int                      // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_DESTROY1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int           // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE func(pItem *xc.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板调整坐标. 已停用.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int // 列表元素,列表头项模板调整坐标. 已停用.

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *List) Event_LIST_TEMP_CREATE(pFun XE_LIST_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_TEMP_CREATE, pFun)
}

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *List) Event_LIST_TEMP_CREATE1(pFun XE_LIST_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_TEMP_CREATE, pFun)
}

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *List) Event_LIST_TEMP_CREATE_END(pFun XE_LIST_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_TEMP_CREATE_END, pFun)
}

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *List) Event_LIST_TEMP_CREATE_END1(pFun XE_LIST_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_TEMP_CREATE_END, pFun)
}

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *List) Event_LIST_TEMP_DESTROY(pFun XE_LIST_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_TEMP_DESTROY, pFun)
}

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *List) Event_LIST_TEMP_DESTROY1(pFun XE_LIST_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_TEMP_DESTROY, pFun)
}

// 列表元素,项模板调整坐标. 已停用.
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE(pFun XE_LIST_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,项模板调整坐标. 已停用.
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,绘制项.
func (l *List) Event_LIST_DRAWITEM(pFun XE_LIST_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_DRAWITEM, pFun)
}

// 列表元素,绘制项.
func (l *List) Event_LIST_DRAWITEM1(pFun XE_LIST_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_DRAWITEM, pFun)
}

// 列表元素,项选择事件.
func (l *List) Event_LIST_SELECT(pFun XE_LIST_SELECT) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_SELECT, pFun)
}

// 列表元素,项选择事件.
func (l *List) Event_LIST_SELECT1(pFun XE_LIST_SELECT1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_SELECT, pFun)
}

// 列表元素绘制列表头项.
func (l *List) Event_LIST_HEADER_DRAWITEM(pFun XE_LIST_HEADER_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_DRAWITEM, pFun)
}

// 列表元素绘制列表头项.
func (l *List) Event_LIST_HEADER_DRAWITEM1(pFun XE_LIST_HEADER_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_DRAWITEM, pFun)
}

// 列表元素,列表头项点击事件.
func (l *List) Event_LIST_HEADER_CLICK(pFun XE_LIST_HEADER_CLICK) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_CLICK, pFun)
}

// 列表元素,列表头项点击事件.
func (l *List) Event_LIST_HEADER_CLICK1(pFun XE_LIST_HEADER_CLICK1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_CLICK, pFun)
}

// 列表元素,列表头项宽度改变事件.
func (l *List) Event_LIST_HEADER_WIDTH_CHANGE(pFun XE_LIST_HEADER_WIDTH_CHANGE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// 列表元素,列表头项宽度改变事件.
func (l *List) Event_LIST_HEADER_WIDTH_CHANGE1(pFun XE_LIST_HEADER_WIDTH_CHANGE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// 列表元素,列表头项模板创建.
func (l *List) Event_LIST_HEADER_TEMP_CREATE(pFun XE_LIST_HEADER_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// 列表元素,列表头项模板创建.
func (l *List) Event_LIST_HEADER_TEMP_CREATE1(pFun XE_LIST_HEADER_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// 列表元素,列表头项模板创建完成事件.
func (l *List) Event_LIST_HEADER_TEMP_CREATE_END(pFun XE_LIST_HEADER_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// 列表元素,列表头项模板创建完成事件.
func (l *List) Event_LIST_HEADER_TEMP_CREATE_END1(pFun XE_LIST_HEADER_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// 列表元素,列表头项模板销毁.
func (l *List) Event_LIST_HEADER_TEMP_DESTROY(pFun XE_LIST_HEADER_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// 列表元素,列表头项模板销毁.
func (l *List) Event_LIST_HEADER_TEMP_DESTROY1(pFun XE_LIST_HEADER_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// 列表元素,列表头项模板调整坐标. 已停用.
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,列表头项模板调整坐标. 已停用.
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}
