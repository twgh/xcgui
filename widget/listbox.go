package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ListBox 列表框.
type ListBox struct {
	ScrollView
}

// 列表框_创建, 创建列表框元素.
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
func NewListBox(x, y, cx, cy int32, hParent int) *ListBox {
	p := &ListBox{}
	p.SetHandle(xc.XListBox_Create(x, y, cx, cy, hParent))
	return p
}

// 列表框_创建Ex, 创建列表框元素, 使用内置项模板, 自动创建数据适配器.
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
func NewListBoxEx(x, y, cx, cy int32, hParent, col_extend_count int32) *ListBox {
	p := &ListBox{}
	p.SetHandle(xc.XListBox_CreateEx(x, y, cx, cy, hParent, col_extend_count))
	return p
}

// 从句柄创建对象.
func NewListBoxByHandle(handle int) *ListBox {
	p := &ListBox{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewListBoxByName(name string) *ListBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewListBoxByUID(nUID int32) *ListBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewListBoxByUIDName(name string) *ListBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表框_启用固定行高.
//
// bEnable: 是否启用.
func (l *ListBox) EnableFixedRowHeight(bEnable bool) *ListBox {
	xc.XListBox_EnableFixedRowHeight(l.Handle, bEnable)
	return l
}

// 列表框_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListBox) EnableTemplateReuse(bEnable bool) *ListBox {
	xc.XListBox_EnableTemplateReuse(l.Handle, bEnable)
	return l
}

// 列表框_启用虚表.
//
// bEnable: 是否启用.
func (l *ListBox) EnableVirtualTable(bEnable bool) *ListBox {
	xc.XListBox_EnableVirtualTable(l.Handle, bEnable)
	return l
}

// 列表框_置虚表行数.
//
// nRowCount: 行数.
func (l *ListBox) SetVirtualRowCount(nRowCount int32) *ListBox {
	xc.XListBox_SetVirtualRowCount(l.Handle, nRowCount)
	return l
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func (l *ListBox) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) *ListBox {
	xc.XListBox_SetDrawItemBkFlags(l.Handle, nFlags)
	return l
}

// 列表框_置项数据, 设置项用户数据.
//
// iItem: 想索引.
//
// nUserData: 用户数据.
func (l *ListBox) SetItemData(iItem int32, nUserData int) bool {
	return xc.XListBox_SetItemData(l.Handle, iItem, nUserData)
}

// 列表框_取项数据, 获取项用户数据.
//
// iItem: 项索引.
func (l *ListBox) GetItemData(iItem int32) int {
	return xc.XListBox_GetItemData(l.Handle, iItem)
}

// 列表框_置项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) SetItemInfo(iItem int32, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_SetItemInfo(l.Handle, iItem, pItem)
}

// 列表框_取项背景信息, 获取项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) GetItemInfo(iItem int32, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_GetItemInfo(l.Handle, iItem, pItem)
}

// 列表框_置选择项, 设置选择选.
//
// iItem: 项索引.
func (l *ListBox) SetSelectItem(iItem int32) bool {
	return xc.XListBox_SetSelectItem(l.Handle, iItem)
}

// 列表框_取选择项, 返回项索引.
func (l *ListBox) GetSelectItem() int32 {
	return xc.XListBox_GetSelectItem(l.Handle)
}

// 列表框_添加选择项.
//
// iItem: 项索引.
func (l *ListBox) AddSelectItem(iItem int32) bool {
	return xc.XListBox_AddSelectItem(l.Handle, iItem)
}

// 列表框_取消选择项.
//
// iItem: 项索引.
func (l *ListBox) CancelSelectItem(iItem int32) bool {
	return xc.XListBox_CancelSelectItem(l.Handle, iItem)
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
func (l *ListBox) CancelSelectAll() bool {
	return xc.XListBox_CancelSelectAll(l.Handle)
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//
// pArray: 数组缓冲区.
//
// nArraySize: 数组大小.
func (l *ListBox) GetSelectAll(pArray *[]int32, nArraySize int32) int32 {
	return xc.XListBox_GetSelectAll(l.Handle, pArray, nArraySize)
}

// 列表框_取选择项数量, 获取选择项数量.
func (l *ListBox) GetSelectCount() int32 {
	return xc.XListBox_GetSelectCount(l.Handle)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
func (l *ListBox) GetItemMouseStay() int32 {
	return xc.XListBox_GetItemMouseStay(l.Handle)
}

// 列表框_选择全部项.
func (l *ListBox) SelectAll() bool {
	return xc.XListBox_SelectAll(l.Handle)
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//
// iItem: 项索引.
func (l *ListBox) VisibleItem(iItem int32) *ListBox {
	xc.XListBox_VisibleItem(l.Handle, iItem)
	return l
}

// 列表框_取可视行范围, 获取当前可见行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *ListBox) GetVisibleRowRange(piStart *int32, piEnd *int32) *ListBox {
	xc.XListBox_GetVisibleRowRange(l.Handle, piStart, piEnd)
	return l
}

// 列表框_置项默认高度.
//
// nHeight: 项高度.
//
// nSelHeight: 选中项高度.
func (l *ListBox) SetItemHeightDefault(nHeight int32, nSelHeight int32) *ListBox {
	xc.XListBox_SetItemHeightDefault(l.Handle, nHeight, nSelHeight)
	return l
}

// 列表框_取项默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *ListBox) GetItemHeightDefault(pHeight *int32, pSelHeight *int32) *ListBox {
	xc.XListBox_GetItemHeightDefault(l.Handle, pHeight, pSelHeight)
	return l
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *ListBox) GetItemIndexFromHXCGUI(hXCGUI int) int32 {
	return xc.XListBox_GetItemIndexFromHXCGUI(l.Handle, hXCGUI)
}

// 列表框_置行间距.
//
// nSpace: 间距大小.
func (l *ListBox) SetRowSpace(nSpace int32) *ListBox {
	xc.XListBox_SetRowSpace(l.Handle, nSpace)
	return l
}

// 列表框_取行间距.
func (l *ListBox) GetRowSpace() int32 {
	return xc.XListBox_GetRowSpace(l.Handle)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) HitTest(pPt *xc.POINT) int32 {
	return xc.XListBox_HitTest(l.Handle, pPt)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) HitTestOffset(pPt *xc.POINT) int32 {
	return xc.XListBox_HitTestOffset(l.Handle, pPt)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//
// pXmlFile: 文件名.
func (l *ListBox) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XListBox_SetItemTemplateXML(l.Handle, pXmlFile)
}

// 列表框_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListBox) SetItemTemplate(hTemp int) bool {
	return xc.XListBox_SetItemTemplate(l.Handle, hTemp)
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *ListBox) SetItemTemplateXMLFromString(pStringXML string) bool {
	return xc.XListBox_SetItemTemplateXMLFromString(l.Handle, pStringXML)
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListBox) GetTemplateObject(iItem, nTempItemID int32) int {
	return xc.XListBox_GetTemplateObject(l.Handle, iItem, nTempItemID)
}

// 列表框_启用多选, 是否启用多行选择功能.
//
// bEnable: 是否启用.
func (l *ListBox) EnableMultiSel(bEnable bool) *ListBox {
	xc.XListBox_EnableMultiSel(l.Handle, bEnable)
	return l
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListBox) CreateAdapter() int {
	return xc.XListBox_CreateAdapter(l.Handle)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *ListBox) BindAdapter(hAdapter int) *ListBox {
	xc.XListBox_BindAdapter(l.Handle, hAdapter)
	return l
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
func (l *ListBox) GetAdapter() int {
	return xc.XListBox_GetAdapter(l.Handle)
}

// 列表框_排序.
//
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
//
// bAscending: 升序(TRUE)或降序(FALSE).
func (l *ListBox) Sort(iColumnAdapter int32, bAscending bool) *ListBox {
	xc.XListBox_Sort(l.Handle, iColumnAdapter, bAscending)
	return l
}

// 列表框_刷新数据.
func (l *ListBox) RefreshData() *ListBox {
	xc.XListBox_RefreshData(l.Handle)
	return l
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iItem: 项索引.
func (l *ListBox) RefreshItem(iItem int32) *ListBox {
	xc.XListBox_RefreshItem(l.Handle, iItem)
	return l
}

// 列表框_添加项文本, XAdTable_AddRowText, 返回项索引.
//
// pText:.
func (l *ListBox) AddItemText(pText string) int32 {
	return xc.XListBox_AddItemText(l.Handle, pText)
}

// 列表框_添加项文本扩展, XAdTable_AddRowTextEx.
//
// pName:.
//
// pText:.
func (l *ListBox) AddItemTextEx(pName string, pText string) int32 {
	return xc.XListBox_AddItemTextEx(l.Handle, pName, pText)
}

// 列表框_添加项图片, XAdTable_AddRowImage.
//
// hImage:.
func (l *ListBox) AddItemImage(hImage int) int32 {
	return xc.XListBox_AddItemImage(l.Handle, hImage)
}

// 列表框_添加项图片扩展, XAdTable_AddRowImageEx.
//
// pName:.
//
// hImage:.
func (l *ListBox) AddItemImageEx(pName string, hImage int) int32 {
	return xc.XListBox_AddItemImageEx(l.Handle, pName, hImage)
}

// 列表框_插入项文本.
//
// iItem:.
//
// pValue:.
func (l *ListBox) InsertItemText(iItem int32, pValue string) int32 {
	return xc.XListBox_InsertItemText(l.Handle, iItem, pValue)
}

// 列表框_插入项文本扩展.
//
// iItem:.
//
// pName:.
//
// pValue:.
func (l *ListBox) InsertItemTextEx(iItem int32, pName string, pValue string) int32 {
	return xc.XListBox_InsertItemTextEx(l.Handle, iItem, pName, pValue)
}

// 列表框_插入项图片.
//
// iItem:.
//
// hImage:.
func (l *ListBox) InsertItemImage(iItem int32, hImage int) int32 {
	return xc.XListBox_InsertItemImage(l.Handle, iItem, hImage)
}

// 列表框_插入项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) InsertItemImageEx(iItem int32, pName string, hImage int) int32 {
	return xc.XListBox_InsertItemImageEx(l.Handle, iItem, pName, hImage)
}

// 列表框_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *ListBox) SetItemText(iItem, iColumn int32, pText string) bool {
	return xc.XListBox_SetItemText(l.Handle, iItem, iColumn, pText)
}

// 列表框_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *ListBox) SetItemTextEx(iItem int32, pName string, pText string) bool {
	return xc.XListBox_SetItemTextEx(l.Handle, iItem, pName, pText)
}

// 列表框_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *ListBox) SetItemImage(iItem, iColumn int32, hImage int) bool {
	return xc.XListBox_SetItemImage(l.Handle, iItem, iColumn, hImage)
}

// 列表框_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) SetItemImageEx(iItem int32, pName string, hImage int) bool {
	return xc.XListBox_SetItemImageEx(l.Handle, iItem, pName, hImage)
}

// 列表框_置项整数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *ListBox) SetItemInt(iItem, iColumn int32, nValue int32) bool {
	return xc.XListBox_SetItemInt(l.Handle, iItem, iColumn, nValue)
}

// 列表框_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *ListBox) SetItemIntEx(iItem int32, pName string, nValue int32) bool {
	return xc.XListBox_SetItemIntEx(l.Handle, iItem, pName, nValue)
}

// 列表框_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *ListBox) SetItemFloat(iItem, iColumn int32, fFloat float32) bool {
	return xc.XListBox_SetItemFloat(l.Handle, iItem, iColumn, fFloat)
}

// 列表框_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *ListBox) SetItemFloatEx(iItem int32, pName string, fFloat float32) bool {
	return xc.XListBox_SetItemFloatEx(l.Handle, iItem, pName, fFloat)
}

// 列表框_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) GetItemText(iItem, iColumn int32) string {
	return xc.XListBox_GetItemText(l.Handle, iItem, iColumn)
}

// 列表框_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *ListBox) GetItemTextEx(iItem int32, pName string) string {
	return xc.XListBox_GetItemTextEx(l.Handle, iItem, pName)
}

// 列表框_取项图片. 返回图片句柄.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) GetItemImage(iItem, iColumn int32) int {
	return xc.XListBox_GetItemImage(l.Handle, iItem, iColumn)
}

// 列表框_取项图片扩展. 返回图片句柄.
//
// iItem:.
//
// pName:.
func (l *ListBox) GetItemImageEx(iItem int32, pName string) int {
	return xc.XListBox_GetItemImageEx(l.Handle, iItem, pName)
}

// 列表框_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) GetItemInt(iItem, iColumn int32, pOutValue *int32) bool {
	return xc.XListBox_GetItemInt(l.Handle, iItem, iColumn, pOutValue)
}

// 列表框_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) GetItemIntEx(iItem int32, pName string, pOutValue *int32) bool {
	return xc.XListBox_GetItemIntEx(l.Handle, iItem, pName, pOutValue)
}

// 列表框_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) GetItemFloat(iItem, iColumn int32, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloat(l.Handle, iItem, iColumn, pOutValue)
}

// 列表框_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) GetItemFloatEx(iItem int32, pName string, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloatEx(l.Handle, iItem, pName, pOutValue)
}

// 列表框_删除项.
//
// iItem:.
func (l *ListBox) DeleteItem(iItem int32) bool {
	return xc.XListBox_DeleteItem(l.Handle, iItem)
}

// 列表框_删除项扩展.
//
// iItem:.
//
// nCount:.
func (l *ListBox) DeleteItemEx(iItem int32, nCount int32) bool {
	return xc.XListBox_DeleteItemEx(l.Handle, iItem, nCount)
}

// 列表框_删除项全部.
func (l *ListBox) DeleteItemAll() *ListBox {
	xc.XListBox_DeleteItemAll(l.Handle)
	return l
}

// 列表框_删除列全部.
func (l *ListBox) DeleteColumnAll() *ListBox {
	xc.XListBox_DeleteColumnAll(l.Handle)
	return l
}

// 列表框_取项数量AD.
func (l *ListBox) GetCount_AD() int32 {
	return xc.XListBox_GetCount_AD(l.Handle)
}

// 列表框_取列数量AD.
func (l *ListBox) GetCountColumn_AD() int32 {
	return xc.XListBox_GetCountColumn_AD(l.Handle)
}

// 列表框_置分割线颜色.
//
// color: xc.RGBA 颜色值.
func (l *ListBox) SetSplitLineColor(color uint32) *ListBox {
	xc.XListBox_SetSplitLineColor(l.Handle, color)
	return l
}

// 列表框_置拖动矩形颜色.
//
// color: xc.RGBA 颜色值.
//
// width: 线宽度.
func (l *ListBox) SetDragRectColor(color uint32, width int32) *ListBox {
	xc.XListBox_SetDragRectColor(l.Handle, color, width)
	return l
}

// 列表框_置项模板从内存. 设置项模板文件.
//
// data: 模板数据.
func (l *ListBox) SetItemTemplateXMLFromMem(data []byte) bool {
	return xc.XListBox_SetItemTemplateXMLFromMem(l.Handle, data)
}

// 列表框_置项模板从资源ZIP. 设置项模板文件.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *ListBox) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XListBox_SetItemTemplateXMLFromZipRes(l.Handle, id, pFileName, pPassword, hModule)
}

// 列表框_取项模板. 获取列表项模板, 返回项模板句柄.
func (l *ListBox) GetItemTemplate() int {
	return xc.XListBox_GetItemTemplate(l.Handle)
}

// 列表框_置项高度. 注意:为了提高性能, 默认使用列表框的全局项高度, 如果需要指定项单独高度, 需要关闭固定行高: xc.XListBox_EnableFixedRowHeight
//
// iItem: 项索引.
//
// nHeight: 项高度.
//
// nSelHeight: 项选中时高度.
func (l *ListBox) SetItemHeight(iItem, nHeight, nSelHeight int32) *ListBox {
	xc.XListBox_SetItemHeight(l.Handle, iItem, nHeight, nSelHeight)
	return l
}

// 列表框_取项高度.
//
// iItem: 项索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *ListBox) GetItemHeight(iItem int32, pHeight, pSelHeight *int32) *ListBox {
	xc.XListBox_GetItemHeight(l.Handle, iItem, pHeight, pSelHeight)
	return l
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_ListBox_Temp_Create 添加列表框项模板创建事件, 模板复用机制需先启用. 替换模板无效判断 nFlag, 因为内部会检查模板是否改变, 不用担心重复.
//   - nFlag:  0: 状态改变; 1: 新模板实例; 2: 旧模板复用
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_Temp_Create(pFun xc.XE_LISTBOX_TEMP_CREATE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE, onXE_LISTBOX_TEMP_CREATE, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_TEMP_CREATE 列表框项模板创建事件.
func onXE_LISTBOX_TEMP_CREATE(hEle int, pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_TEMP_CREATE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_TEMP_CREATE1); ok {
			ret = cb(hEle, pItem, nFlag, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ListBox_Temp_Create_End 添加列表框项模板创建完成事件, 模板复用机制需先启用. 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//   - nFlag:  0: 状态改变(复用); 1: 新模板实例; 2: 旧模板复用
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_Temp_Create_End(pFun xc.XE_LISTBOX_TEMP_CREATE_END1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE_END, onXE_LISTBOX_TEMP_CREATE_END, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_TEMP_CREATE_END 列表框项模板创建完成事件.
func onXE_LISTBOX_TEMP_CREATE_END(hEle int, pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_TEMP_CREATE_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_TEMP_CREATE_END1); ok {
			ret = cb(hEle, pItem, nFlag, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ListBox_Temp_Destroy 添加列表框项模板销毁事件.
//   - nFlag: 0:正常销毁; 1:移动到缓存(不会被销毁, 临时缓存备用, 当需要时被复用)
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_Temp_Destroy(pFun xc.XE_LISTBOX_TEMP_DESTROY1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_TEMP_DESTROY, onXE_LISTBOX_TEMP_DESTROY, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_TEMP_DESTROY 列表框项模板销毁事件.
func onXE_LISTBOX_TEMP_DESTROY(hEle int, pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_TEMP_DESTROY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_TEMP_DESTROY1); ok {
			ret = cb(hEle, pItem, nFlag, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ListBox_Temp_Adjust_Coordinate 添加列表框项模板调整坐标事件. 已停用.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_Temp_Adjust_Coordinate(pFun xc.XE_LISTBOX_TEMP_ADJUST_COORDINATE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, onXE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_TEMP_ADJUST_COORDINATE 列表框项模板调整坐标事件.
func onXE_LISTBOX_TEMP_ADJUST_COORDINATE(hEle int, pItem *xc.ListBox_Item_, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_TEMP_ADJUST_COORDINATE1); ok {
			ret = cb(hEle, pItem, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ListBox_DrawItem 添加列表框项绘制事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_DrawItem(pFun xc.XE_LISTBOX_DRAWITEM1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_DRAWITEM, onXE_LISTBOX_DRAWITEM, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_DRAWITEM 列表框项绘制事件.
func onXE_LISTBOX_DRAWITEM(hEle int, hDraw int, pItem *xc.ListBox_Item_, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_DRAWITEM)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_DRAWITEM1); ok {
			ret = cb(hEle, hDraw, pItem, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ListBox_Select 添加列表框项选择事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (l *ListBox) AddEvent_ListBox_Select(pFun xc.XE_LISTBOX_SELECT1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(l.Handle, xcc.XE_LISTBOX_SELECT, onXE_LISTBOX_SELECT, pFun, allowAddingMultiple...)
}

// onXE_LISTBOX_SELECT 列表框项选择事件.
func onXE_LISTBOX_SELECT(hEle int, iItem int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LISTBOX_SELECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LISTBOX_SELECT1); ok {
			ret = cb(hEle, iItem, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *ListBox) Event_LISTBOX_TEMP_CREATE(pFun xc.XE_LISTBOX_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *ListBox) Event_LISTBOX_TEMP_CREATE1(pFun xc.XE_LISTBOX_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END(pFun xc.XE_LISTBOX_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END1(pFun xc.XE_LISTBOX_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *ListBox) Event_LISTBOX_TEMP_DESTROY(pFun xc.XE_LISTBOX_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *ListBox) Event_LISTBOX_TEMP_DESTROY1(pFun xc.XE_LISTBOX_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE(pFun xc.XE_LISTBOX_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE1(pFun xc.XE_LISTBOX_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项绘制事件.
func (l *ListBox) Event_LISTBOX_DRAWITEM(pFun xc.XE_LISTBOX_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项绘制事件.
func (l *ListBox) Event_LISTBOX_DRAWITEM1(pFun xc.XE_LISTBOX_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项选择事件.
func (l *ListBox) Event_LISTBOX_SELECT(pFun xc.XE_LISTBOX_SELECT) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_SELECT, pFun)
}

// 列表框元素,项选择事件.
func (l *ListBox) Event_LISTBOX_SELECT1(pFun xc.XE_LISTBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_SELECT, pFun)
}
