package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 列表视图.
type ListView struct {
	ScrollView
}

// 列表视_创建.
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
func NewListView(x int, y int, cx int, cy int, hParent int) *ListView {
	p := &ListView{}
	p.SetHandle(xc.XListView_Create(x, y, cx, cy, hParent))
	return p
}

// 列表视_创建Ex. 创建列表视图元素, 使用内置项模板, 自动创建数据适配器.
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
func NewListViewEx(x, y, cx, cy int32, hParent, col_extend_count int32) *ListView {
	p := &ListView{}
	p.SetHandle(xc.XListView_CreateEx(x, y, cx, cy, hParent, col_extend_count))
	return p
}

// 从句柄创建对象.
func NewListViewByHandle(handle int) *ListView {
	p := &ListView{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewListViewByName(name string) *ListView {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewListViewByUID(nUID int) *ListView {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewListViewByUIDName(name string) *ListView {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListView) CreateAdapter() int {
	return xc.XListView_CreateAdapter(l.Handle)
}

// 列表视_绑定数据适配器.
//
// hAdapter: 数据适配器XAdListView.
func (l *ListView) BindAdapter(hAdapter int) int {
	return xc.XListView_BindAdapter(l.Handle, hAdapter)
}

// 列表视_取数据适配器, 返回数据适配器句柄.
func (l *ListView) GetAdapter() int {
	return xc.XListView_GetAdapter(l.Handle)
}

// 列表视_置项模板文件.
//
// pXmlFile: 文件名.
func (l *ListView) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XListView_SetItemTemplateXML(l.Handle, pXmlFile)
}

// 列表视_置项模板从字符串.
//
// pStringXML: 字符串.
func (l *ListView) SetItemTemplateXMLFromString(pStringXML string) bool {
	return xc.XListView_SetItemTemplateXMLFromString(l.Handle, pStringXML)
}

// 列表视_置项模板, 置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListView) SetItemTemplate(hTemp int) bool {
	return xc.XListView_SetItemTemplate(l.Handle, hTemp)
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) GetTemplateObject(iGroup int, iItem int, nTempItemID int) int {
	return xc.XListView_GetTemplateObject(l.Handle, iGroup, iItem, nTempItemID)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) GetTemplateObjectGroup(iGroup int, nTempItemID int) int {
	return xc.XListView_GetTemplateObjectGroup(l.Handle, iGroup, nTempItemID)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) GetItemIDFromHXCGUI(hXCGUI int, piGroup *int32, piItem *int32) bool {
	return xc.XListView_GetItemIDFromHXCGUI(l.Handle, hXCGUI, piGroup, piItem)
}

// 列表视_测试点击项, 检查坐标点所在项.
//
// pPt: 坐标点.
//
// pOutGroup: 接收组索引.
//
// pOutItem: 接收项索引.
func (l *ListView) HitTest(pPt *xc.POINT, pOutGroup *int32, pOutItem *int32) bool {
	return xc.XListView_HitTest(l.Handle, pPt, pOutGroup, pOutItem)
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// pOutGroup: 接收做索引.
//
// pOutItem: 接收项索引.
func (l *ListView) HitTestOffset(pPt *xc.POINT, pOutGroup *int32, pOutItem *int32) bool {
	return xc.XListView_HitTestOffset(l.Handle, pPt, pOutGroup, pOutItem)
}

// 列表视_启用多选.
//
// bEnable: 是否启用.
func (l *ListView) EnableMultiSel(bEnable bool) int {
	return xc.XListView_EnableMultiSel(l.Handle, bEnable)
}

// 列表视_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListView) EnableTemplateReuse(bEnable bool) int {
	return xc.XListView_EnableTemplateReuse(l.Handle, bEnable)
}

// 列表视_启用虚表.
//
// bEnable: 是否启用.
func (l *ListView) EnableVirtualTable(bEnable bool) int {
	return xc.XListView_EnableVirtualTable(l.Handle, bEnable)
}

// 列表视_置虚表项数量.
//
// iGroup: 组索引.
//
// nCount: 项数量.
func (l *ListView) SetVirtualItemCount(iGroup int, nCount int) bool {
	return xc.XListView_SetVirtualItemCount(l.Handle, iGroup, nCount)
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func (l *ListView) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XListView_SetDrawItemBkFlags(l.Handle, nFlags)
}

// 列表视_置选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) SetSelectItem(iGroup int, iItem int) bool {
	return xc.XListView_SetSelectItem(l.Handle, iGroup, iItem)
}

// 列表视_取选择项.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) GetSelectItem(piGroup *int32, piItem *int32) bool {
	return xc.XListView_GetSelectItem(l.Handle, piGroup, piItem)
}

// 列表视_添加选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) AddSelectItem(iGroup int, iItem int) bool {
	return xc.XListView_AddSelectItem(l.Handle, iGroup, iItem)
}

// 列表视_显示指定项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) VisibleItem(iGroup int, iItem int) int {
	return xc.XListView_VisibleItem(l.Handle, iGroup, iItem)
}

// 列表视_取可视项范围, 获取当前可见项范围.
//
// piGroup1: 可视开始组.
//
// piGroup2: 可视结束组.
//
// piStartGroup: 可视开始组.
//
// piStartItem: 可视开始项.
//
// piEndGroup: 可视结束组.
//
// piEndItem: 可视结束项.
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup *int32, piEndItem *int32) int {
	return xc.XListView_GetVisibleItemRange(l.Handle, piGroup1, piGroup2, piStartGroup, piStartItem, piEndGroup, piEndItem)
}

// 列表视_取选择项数量.
func (l *ListView) GetSelectItemCount() int {
	return xc.XListView_GetSelectItemCount(l.Handle)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量.
//
// pArray: 数组.
//
// nArraySize: 数组大小.
func (l *ListView) GetSelectAll(pArray *[]xc.ListView_Item_Id_, nArraySize int) int {
	return xc.XListView_GetSelectAll(l.Handle, pArray, nArraySize)
}

// 列表视_置选择项全部, 选择所有的项.
func (l *ListView) SetSelectAll() int {
	return xc.XListView_SetSelectAll(l.Handle)
}

// 列表视_取消选择项全部, 取消选择所有项.
func (l *ListView) CancelSelectAll() int {
	return xc.XListView_CancelSelectAll(l.Handle)
}

// 列表视_置列间隔, 置列间隔大小.
//
// space: 间隔大小.
func (l *ListView) SetColumnSpace(space int) int {
	return xc.XListView_SetColumnSpace(l.Handle, space)
}

// 列表视_置行间隔, 置行间隔大小.
//
// space: 间隔大小.
func (l *ListView) SetRowSpace(space int) int {
	return xc.XListView_SetRowSpace(l.Handle, space)
}

// 列表视_置项大小.
//
// width: 宽度.
//
// height: 高度.
func (l *ListView) SetItemSize(width int, height int) int {
	return xc.XListView_SetItemSize(l.Handle, width, height)
}

// 列表视_取项大小.
//
// pSize: 接收返回大小.
func (l *ListView) GetItemSize(pSize *xc.SIZE) int {
	return xc.XListView_GetItemSize(l.Handle, pSize)
}

// 列表视_置组高度.
//
// height: 高度.
func (l *ListView) SetGroupHeight(height int) int {
	return xc.XListView_SetGroupHeight(l.Handle, height)
}

// 列表视_取组高度.
func (l *ListView) GetGroupHeight() int {
	return xc.XListView_GetGroupHeight(l.Handle)
}

// 列表视_置组用户数据.
//
// iGroup: 组索引.
//
// nData: 数据.
func (l *ListView) SetGroupUserData(iGroup int, nData int) int {
	return xc.XListView_SetGroupUserData(l.Handle, iGroup, nData)
}

// 列表视_置项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nData: 数据.
func (l *ListView) SetItemUserData(iGroup int, iItem int, nData int) int {
	return xc.XListView_SetItemUserData(l.Handle, iGroup, iItem, nData)
}

// 列表视_取组用户数据.
//
// iGroup: 组索引.
func (l *ListView) GetGroupUserData(iGroup int) int {
	return xc.XListView_GetGroupUserData(l.Handle, iGroup)
}

// 列表视_取项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) GetItemUserData(iGroup int, iItem int) int {
	return xc.XListView_GetItemUserData(l.Handle, iGroup, iItem)
}

// 列表视_刷新项数据.
func (l *ListView) RefreshData() int {
	return xc.XListView_RefreshData(l.Handle)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iGroup: 组索引.
//
// iItem: 项索引, 如果为-1, 代表为组.
func (l *ListView) RefreshItem(iGroup int, iItem int) int {
	return xc.XListView_RefreshItem(l.Handle, iGroup, iItem)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE.
//
// iGroup: 组索引.
//
// bExpand: 是否展开.
func (l *ListView) ExpandGroup(iGroup int, bExpand bool) bool {
	return xc.XListView_ExpandGroup(l.Handle, iGroup, bExpand)
}

// 列表视_组添加列, 返回列索引.
//
// pName: 字段称.
func (l *ListView) Group_AddColumn(pName string) int {
	return xc.XListView_Group_AddColumn(l.Handle, pName)
}

// 列表视_组添加项文本, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) Group_AddItemText(pValue string, iPos int) int {
	return xc.XListView_Group_AddItemText(l.Handle, pValue, iPos)
}

// 列表视_组添加项文本扩展, 返回组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) Group_AddItemTextEx(pName string, pValue string, iPos int) int {
	return xc.XListView_Group_AddItemTextEx(l.Handle, pName, pValue, iPos)
}

// 列表视_组添加项图片, 返回组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) Group_AddItemImage(hImage int, iPos int) int {
	return xc.XListView_Group_AddItemImage(l.Handle, hImage, iPos)
}

// 列表视_组添加项图片扩展, 返回组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) Group_AddItemImageEx(pName string, hImage int, iPos int) int {
	return xc.XListView_Group_AddItemImageEx(l.Handle, pName, hImage, iPos)
}

// 列表视_组置文本.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) Group_SetText(iGroup int, iColumn int, pValue string) bool {
	return xc.XListView_Group_SetText(l.Handle, iGroup, iColumn, pValue)
}

// 列表视_组置文本扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) Group_SetTextEx(iGroup int, pName string, pValue string) bool {
	return xc.XListView_Group_SetTextEx(l.Handle, iGroup, pName, pValue)
}

// 列表视_组置图片.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) Group_SetImage(iGroup int, iColumn int, hImage int) bool {
	return xc.XListView_Group_SetImage(l.Handle, iGroup, iColumn, hImage)
}

// 列表视_组置图片扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (l *ListView) Group_SetImageEx(iGroup int, pName string, hImage int) bool {
	return xc.XListView_Group_SetImageEx(l.Handle, iGroup, pName, hImage)
}

// 列表视_组获取数量, 返回组数量.
func (l *ListView) Group_GetCount() int {
	return xc.XListView_Group_GetCount(l.Handle)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 XC_ID_ERROR.
//
// iGroup: 组索引.
func (l *ListView) Item_GetCount(iGroup int) int {
	return xc.XListView_Item_GetCount(l.Handle, iGroup)
}

// 列表视_项添加列, 返回列索引.
//
// pName: 字段名.
func (l *ListView) Item_AddColumn(pName string) int {
	return xc.XListView_Item_AddColumn(l.Handle, pName)
}

// 列表视_项添加文本, 返回项索引.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) Item_AddItemText(iGroup int, pValue string, iPos int) int {
	return xc.XListView_Item_AddItemText(l.Handle, iGroup, pValue, iPos)
}

// 列表视_项添加文本扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) Item_AddItemTextEx(iGroup int, pName string, pValue string, iPos int) int {
	return xc.XListView_Item_AddItemTextEx(l.Handle, iGroup, pName, pValue, iPos)
}

// 列表视_项添加图片, 返回项索引.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) Item_AddItemImage(iGroup int, hImage int, iPos int) int {
	return xc.XListView_Item_AddItemImage(l.Handle, iGroup, hImage, iPos)
}

// 列表视_项添加图片扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) Item_AddItemImageEx(iGroup int, pName string, hImage int, iPos int) int {
	return xc.XListView_Item_AddItemImageEx(l.Handle, iGroup, pName, hImage, iPos)
}

// 列表视_项置文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) Item_SetText(iGroup int, iItem int, iColumn int, pValue string) bool {
	return xc.XListView_Item_SetText(l.Handle, iGroup, iItem, iColumn, pValue)
}

// 列表视_项置文本扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) Item_SetTextEx(iGroup int, iItem int, pName string, pValue string) bool {
	return xc.XListView_Item_SetTextEx(l.Handle, iGroup, iItem, pName, pValue)
}

// 列表视_项置图片.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) Item_SetImage(iGroup int, iItem int, iColumn int, hImage int) bool {
	return xc.XListView_Item_SetImage(l.Handle, iGroup, iItem, iColumn, hImage)
}

// 列表视_项置图片扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 列名称.
//
// hImage: 图片句柄.
func (l *ListView) Item_SetImageEx(iGroup int, iItem int, pName string, hImage int) bool {
	return xc.XListView_Item_SetImageEx(l.Handle, iGroup, iItem, pName, hImage)
}

// 列表视_组删除项.
//
// iGroup: 组索引.
func (l *ListView) Group_DeleteItem(iGroup int) bool {
	return xc.XListView_Group_DeleteItem(l.Handle, iGroup)
}

// 列表视_组删除全部子项.
//
// iGroup: 组索引.
func (l *ListView) Group_DeleteAllChildItem(iGroup int) int {
	return xc.XListView_Group_DeleteAllChildItem(l.Handle, iGroup)
}

// 列表视_项删除.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) Item_DeleteItem(iGroup int, iItem int) bool {
	return xc.XListView_Item_DeleteItem(l.Handle, iGroup, iItem)
}

// 列表视_删除全部.
func (l *ListView) DeleteAll() int {
	return xc.XListView_DeleteAll(l.Handle)
}

// 列表视_删除全部组.
func (l *ListView) DeleteAllGroup() int {
	return xc.XListView_DeleteAllGroup(l.Handle)
}

// 列表视_删除全部项.
func (l *ListView) DeleteAllItem() int {
	return xc.XListView_DeleteAllItem(l.Handle)
}

// 列表视_组删除列.
//
// iColumn: 列索引.
func (l *ListView) DeleteColumnGroup(iColumn int) int {
	return xc.XListView_DeleteColumnGroup(l.Handle, iColumn)
}

// 列表视_项删除列.
//
// iColumn: 列索引.
func (l *ListView) DeleteColumnItem(iColumn int) int {
	return xc.XListView_DeleteColumnItem(l.Handle, iColumn)
}

// 列表视_项获取文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) Item_GetTextEx(iGroup int, iItem int, pName string) string {
	return xc.XListView_Item_GetTextEx(l.Handle, iGroup, iItem, pName)
}

// 列表视_项获取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) Item_GetImageEx(iGroup int, iItem int, pName string) int {
	return xc.XListView_Item_GetImageEx(l.Handle, iGroup, iItem, pName)
}

// 列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) Group_GetText(iGroup int, iColumn int) string {
	return xc.XListView_Group_GetText(l.Handle, iGroup, iColumn)
}

// 列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) Group_GetTextEx(iGroup int, pName string) string {
	return xc.XListView_Group_GetTextEx(l.Handle, iGroup, pName)
}

// 列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) Group_GetImage(iGroup int, iColumn int) int {
	return xc.XListView_Group_GetImage(l.Handle, iGroup, iColumn)
}

// 列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) Group_GetImageEx(iGroup int, pName string) int {
	return xc.XListView_Group_GetImageEx(l.Handle, iGroup, pName)
}

// 列表视_项取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) Item_GetText(iGroup int, iItem int, iColumn int) string {
	return xc.XListView_Item_GetText(l.Handle, iGroup, iItem, iColumn)
}

// 列表视_项取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) Item_GetImage(iGroup int, iItem int, iColumn int) int {
	return xc.XListView_Item_GetImage(l.Handle, iGroup, iItem, iColumn)
}

// 列表视_置拖动矩形颜色.
//
// color: ABGR 颜色.
//
// width: 线宽度.
func (l *ListView) SetDragRectColor(color int, width int) int {
	return xc.XListView_SetDragRectColor(l.Handle, color, width)
}

// 列表视_置项模板从内存.
//
// data: 模板数据.
func (l *ListView) SetItemTemplateXMLFromMem(data []byte) bool {
	return xc.XListView_SetItemTemplateXMLFromMem(l.Handle, data)
}

// 列表视_置项模板从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *ListView) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XListView_SetItemTemplateXMLFromZipRes(l.Handle, id, pFileName, pPassword, hModule)
}

// 列表视_取项模板, 返回项模板句柄.
func (l *ListView) GetItemTemplate() int {
	return xc.XListView_GetItemTemplate(l.Handle)
}

// 列表视_取项模板组, 返回项模板组句柄.
func (l *ListView) GetItemTemplateGroup() int {
	return xc.XListView_GetItemTemplateGroup(l.Handle)
}

/*
以下都是事件
*/

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE func(pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE1 func(hEle int, pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END func(pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END1 func(hEle int, pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY func(pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY1 func(hEle int, pItem *xc.ListView_Item_, nFlag int32, pbHandled *bool) int
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE func(pItem *xc.ListView_Item_, pbHandled *bool) int            // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.ListView_Item_, pbHandled *bool) int // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_DRAWITEM func(hDraw int, pItem *xc.ListView_Item_, pbHandled *bool) int               // 列表视元素,自绘项.
type XE_LISTVIEW_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.ListView_Item_, pbHandled *bool) int    // 列表视元素,自绘项.
type XE_LISTVIEW_SELECT func(iGroup int32, iItem int32, pbHandled *bool) int                           // 列表视元素,项选择事件.
type XE_LISTVIEW_SELECT1 func(hEle int, iGroup int32, iItem int32, pbHandled *bool) int                // 列表视元素,项选择事件.
type XE_LISTVIEW_EXPAND func(iGroup int32, bExpand bool, pbHandled *bool) int                          // 列表视元素,组展开收缩事件.
type XE_LISTVIEW_EXPAND1 func(hEle int, iGroup int32, bExpand bool, pbHandled *bool) int               // 列表视元素,组展开收缩事件.

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
func (l *ListView) Event_LISTVIEW_TEMP_CREATE(pFun XE_LISTVIEW_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
func (l *ListView) Event_LISTVIEW_TEMP_CREATE1(pFun XE_LISTVIEW_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END(pFun XE_LISTVIEW_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END1(pFun XE_LISTVIEW_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
func (l *ListView) Event_LISTVIEW_TEMP_DESTROY(pFun XE_LISTVIEW_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
func (l *ListView) Event_LISTVIEW_TEMP_DESTROY1(pFun XE_LISTVIEW_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// 列表视元素,项模板调整坐标.已停用.
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表视元素,项模板调整坐标.已停用.
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE1(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表视元素,自绘项.
func (l *ListView) Event_LISTVIEW_DRAWITEM(pFun XE_LISTVIEW_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_DRAWITEM, pFun)
}

// 列表视元素,自绘项.
func (l *ListView) Event_LISTVIEW_DRAWITEM1(pFun XE_LISTVIEW_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_DRAWITEM, pFun)
}

// 列表视元素,项选择事件.
func (l *ListView) Event_LISTVIEW_SELECT(pFun XE_LISTVIEW_SELECT) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_SELECT, pFun)
}

// 列表视元素,项选择事件.
func (l *ListView) Event_LISTVIEW_SELECT1(pFun XE_LISTVIEW_SELECT1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_SELECT, pFun)
}

// 列表视元素,组展开收缩事件.
func (l *ListView) Event_LISTVIEW_EXPAND(pFun XE_LISTVIEW_EXPAND) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTVIEW_EXPAND, pFun)
}

// 列表视元素,组展开收缩事件.
func (l *ListView) Event_LISTVIEW_EXPAND1(pFun XE_LISTVIEW_EXPAND1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTVIEW_EXPAND, pFun)
}
