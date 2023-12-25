package xc

import (
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 列表视_创建, 返回元素句柄.
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
func XListView_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xListView_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表视_创建Ex. 创建列表视图元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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
func XListView_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {
	r, _, _ := xListView_CreateEx.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent), uintptr(col_extend_count))
	return int(r)
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func XListView_CreateAdapter(hEle int) int {
	r, _, _ := xListView_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表视_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器XAdListView.
func XListView_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xListView_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表视_取数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func XListView_GetAdapter(hEle int) int {
	r, _, _ := xListView_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func XListView_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xListView_SetItemTemplateXML.Call(uintptr(hEle), common.StrPtr(pXmlFile))
	return r != 0
}

// 列表视_置项模板从字符串.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func XListView_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromString.Call(uintptr(hEle), XC_wtoa(pStringXML))
	return r != 0
}

// 列表视_置项模板, 置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func XListView_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xListView_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return r != 0
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func XListView_GetTemplateObject(hEle int, iGroup int, iItem int, nTempItemID int) int {
	r, _, _ := xListView_GetTemplateObject.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nTempItemID: 模板项ID.
func XListView_GetTemplateObjectGroup(hEle int, iGroup int, nTempItemID int) int {
	r, _, _ := xListView_GetTemplateObjectGroup.Call(uintptr(hEle), uintptr(iGroup), uintptr(nTempItemID))
	return int(r)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func XListView_GetItemIDFromHXCGUI(hEle int, hXCGUI int, piGroup *int32, piItem *int32) bool {
	r, _, _ := xListView_GetItemIDFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI), uintptr(unsafe.Pointer(piGroup)), uintptr(unsafe.Pointer(piItem)))
	return r != 0
}

// 列表视_测试点击项, 检查坐标点所在项.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// pOutGroup: 接收组索引.
//
// pOutItem: 接收项索引.
func XListView_HitTest(hEle int, pPt *POINT, pOutGroup *int32, pOutItem *int32) bool {
	r, _, _ := xListView_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(pOutGroup)), uintptr(unsafe.Pointer(pOutItem)))
	return r != 0
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// pOutGroup: 接收做索引.
//
// pOutItem: 接收项索引.
func XListView_HitTestOffset(hEle int, pPt *POINT, pOutGroup *int32, pOutItem *int32) bool {
	r, _, _ := xListView_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(pOutGroup)), uintptr(unsafe.Pointer(pOutItem)))
	return r != 0
}

// 列表视_启用多选.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListView_EnableMultiSel(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnableMultiSel.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表视_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListView_EnableTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnableTemplateReuse.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表视_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XListView_EnableVirtualTable(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnableVirtualTable.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 列表视_置虚表项数量.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nCount: 项数量.
func XListView_SetVirtualItemCount(hEle int, iGroup int, nCount int) bool {
	r, _, _ := xListView_SetVirtualItemCount.Call(uintptr(hEle), uintptr(iGroup), uintptr(nCount))
	return r != 0
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func XListView_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {
	r, _, _ := xListView_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表视_置选择项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XListView_SetSelectItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_SetSelectItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return r != 0
}

// 列表视_取选择项.
//
// hEle: 元素句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func XListView_GetSelectItem(hEle int, piGroup *int32, piItem *int32) bool {
	r, _, _ := xListView_GetSelectItem.Call(uintptr(hEle), uintptr(unsafe.Pointer(piGroup)), uintptr(unsafe.Pointer(piItem)))
	return r != 0
}

// 列表视_添加选择项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XListView_AddSelectItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_AddSelectItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return r != 0
}

// 列表视_显示指定项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XListView_VisibleItem(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_VisibleItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_取可视项范围, 获取当前可见项范围.
//
// hEle: 元素句柄.
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
func XListView_GetVisibleItemRange(hEle int, piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup *int32, piEndItem *int32) int {
	r, _, _ := xListView_GetVisibleItemRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piGroup1)), uintptr(unsafe.Pointer(piGroup2)), uintptr(unsafe.Pointer(piStartGroup)), uintptr(unsafe.Pointer(piStartItem)), uintptr(unsafe.Pointer(piEndGroup)), uintptr(unsafe.Pointer(piEndItem)))
	return int(r)
}

// 列表视_取选择项数量.
//
// hEle: 元素句柄.
func XListView_GetSelectItemCount(hEle int) int {
	r, _, _ := xListView_GetSelectItemCount.Call(uintptr(hEle))
	return int(r)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量.
//
// hEle: 元素句柄.
//
// pArray: 数组.
//
// nArraySize: 数组大小.
func XListView_GetSelectAll(hEle int, pArray *[]ListView_Item_Id_, nArraySize int) int {
	if nArraySize < 1 {
		return 0
	}
	*pArray = make([]ListView_Item_Id_, nArraySize)
	r, _, _ := xListView_GetSelectAll.Call(uintptr(hEle), uintptr(unsafe.Pointer(&(*pArray)[0])), uintptr(nArraySize))
	return int(r)
}

// 列表视_置选择项全部, 选择所有的项.
//
// hEle: 元素句柄.
func XListView_SetSelectAll(hEle int) int {
	r, _, _ := xListView_SetSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_取消选择项全部, 取消选择所有项.
//
// hEle: 元素句柄.
func XListView_CancelSelectAll(hEle int) int {
	r, _, _ := xListView_CancelSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置列间隔, 置列间隔大小.
//
// hEle: 元素句柄.
//
// space: 间隔大小.
func XListView_SetColumnSpace(hEle int, space int) int {
	r, _, _ := xListView_SetColumnSpace.Call(uintptr(hEle), uintptr(space))
	return int(r)
}

// 列表视_置行间隔, 置行间隔大小.
//
// hEle: 元素句柄.
//
// space: 间隔大小.
func XListView_SetRowSpace(hEle int, space int) int {
	r, _, _ := xListView_SetRowSpace.Call(uintptr(hEle), uintptr(space))
	return int(r)
}

// 列表视_置项大小.
//
// hEle: 元素句柄.
//
// width: 宽度.
//
// height: 高度.
func XListView_SetItemSize(hEle int, width int, height int) int {
	r, _, _ := xListView_SetItemSize.Call(uintptr(hEle), uintptr(width), uintptr(height))
	return int(r)
}

// 列表视_取项大小.
//
// hEle: 元素句柄.
//
// pSize: 接收返回大小.
func XListView_GetItemSize(hEle int, pSize *SIZE) int {
	r, _, _ := xListView_GetItemSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 列表视_置组高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func XListView_SetGroupHeight(hEle int, height int) int {
	r, _, _ := xListView_SetGroupHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 列表视_取组高度.
//
// hEle: 元素句柄.
func XListView_GetGroupHeight(hEle int) int {
	r, _, _ := xListView_GetGroupHeight.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置组用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nData: 数据.
func XListView_SetGroupUserData(hEle int, iGroup int, nData int) int {
	r, _, _ := xListView_SetGroupUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(nData))
	return int(r)
}

// 列表视_置项用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nData: 数据.
func XListView_SetItemUserData(hEle int, iGroup int, iItem int, nData int) int {
	r, _, _ := xListView_SetItemUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(nData))
	return int(r)
}

// 列表视_取组用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func XListView_GetGroupUserData(hEle int, iGroup int) int {
	r, _, _ := xListView_GetGroupUserData.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_取项用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XListView_GetItemUserData(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_GetItemUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_刷新项数据.
//
// hEle: 元素句柄.
func XListView_RefreshData(hEle int) int {
	r, _, _ := xListView_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引, 如果为-1, 代表为组.
func XListView_RefreshItem(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_RefreshItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// bExpand: 是否展开.
func XListView_ExpandGroup(hEle int, iGroup int, bExpand bool) bool {
	r, _, _ := xListView_ExpandGroup.Call(uintptr(hEle), uintptr(iGroup), common.BoolPtr(bExpand))
	return r != 0
}

// 列表视_组添加列, 返回列索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
func XListView_Group_AddColumn(hEle int, pName string) int {
	r, _, _ := xListView_Group_AddColumn.Call(uintptr(hEle), common.StrPtr(pName))
	return int(r)
}

// 列表视_组添加项文本, 返回组索引.
//
// hEle: 元素句柄.
//
// pValue: 值.
//
// iPos: 插入位置.
func XListView_Group_AddItemText(hEle int, pValue string, iPos int) int {
	r, _, _ := xListView_Group_AddItemText.Call(uintptr(hEle), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项文本扩展, 返回组索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置.
func XListView_Group_AddItemTextEx(hEle int, pName string, pValue string, iPos int) int {
	r, _, _ := xListView_Group_AddItemTextEx.Call(uintptr(hEle), common.StrPtr(pName), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项图片, 返回组索引.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func XListView_Group_AddItemImage(hEle int, hImage int, iPos int) int {
	r, _, _ := xListView_Group_AddItemImage.Call(uintptr(hEle), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项图片扩展, 返回组索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func XListView_Group_AddItemImageEx(hEle int, pName string, hImage int, iPos int) int {
	r, _, _ := xListView_Group_AddItemImageEx.Call(uintptr(hEle), common.StrPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_组置文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func XListView_Group_SetText(hEle int, iGroup int, iColumn int, pValue string) bool {
	r, _, _ := xListView_Group_SetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 列表视_组置文本扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func XListView_Group_SetTextEx(hEle int, iGroup int, pName string, pValue string) bool {
	r, _, _ := xListView_Group_SetTextEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 列表视_组置图片.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XListView_Group_SetImage(hEle int, iGroup int, iColumn int, hImage int) bool {
	r, _, _ := xListView_Group_SetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 列表视_组置图片扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func XListView_Group_SetImageEx(hEle int, iGroup int, pName string, hImage int) bool {
	r, _, _ := xListView_Group_SetImageEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 列表视_组获取数量, 返回组数量.
//
// hEle: 元素句柄.
func XListView_Group_GetCount(hEle int) int {
	r, _, _ := xListView_Group_GetCount.Call(uintptr(hEle))
	return int(r)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func XListView_Item_GetCount(hEle int, iGroup int) int {
	r, _, _ := xListView_Item_GetCount.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_项添加列, 返回列索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
func XListView_Item_AddColumn(hEle int, pName string) int {
	r, _, _ := xListView_Item_AddColumn.Call(uintptr(hEle), common.StrPtr(pName))
	return int(r)
}

// 列表视_项添加文本, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func XListView_Item_AddItemText(hEle int, iGroup int, pValue string, iPos int) int {
	r, _, _ := xListView_Item_AddItemText.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_项添加文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func XListView_Item_AddItemTextEx(hEle int, iGroup int, pName string, pValue string, iPos int) int {
	r, _, _ := xListView_Item_AddItemTextEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName), common.StrPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_项添加图片, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func XListView_Item_AddItemImage(hEle int, iGroup int, hImage int, iPos int) int {
	r, _, _ := xListView_Item_AddItemImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_项添加图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func XListView_Item_AddItemImageEx(hEle int, iGroup int, pName string, hImage int, iPos int) int {
	r, _, _ := xListView_Item_AddItemImageEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_项置文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func XListView_Item_SetText(hEle int, iGroup int, iItem int, iColumn int, pValue string) bool {
	r, _, _ := xListView_Item_SetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), common.StrPtr(pValue))
	return r != 0
}

// 列表视_项置文本扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pValue: 值.
func XListView_Item_SetTextEx(hEle int, iGroup int, iItem int, pName string, pValue string) bool {
	r, _, _ := xListView_Item_SetTextEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 列表视_项置图片.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func XListView_Item_SetImage(hEle int, iGroup int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xListView_Item_SetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return r != 0
}

// 列表视_项置图片扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 列名称.
//
// hImage: 图片句柄.
func XListView_Item_SetImageEx(hEle int, iGroup int, iItem int, pName string, hImage int) bool {
	r, _, _ := xListView_Item_SetImageEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName), uintptr(hImage))
	return r != 0
}

// 列表视_组删除项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func XListView_Group_DeleteItem(hEle int, iGroup int) bool {
	r, _, _ := xListView_Group_DeleteItem.Call(uintptr(hEle), uintptr(iGroup))
	return r != 0
}

// 列表视_组删除全部子项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func XListView_Group_DeleteAllChildItem(hEle int, iGroup int) int {
	r, _, _ := xListView_Group_DeleteAllChildItem.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_项删除.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func XListView_Item_DeleteItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_Item_DeleteItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return r != 0
}

// 列表视_删除全部.
//
// hEle: 元素句柄.
func XListView_DeleteAll(hEle int) int {
	r, _, _ := xListView_DeleteAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_删除全部组.
//
// hEle: 元素句柄.
func XListView_DeleteAllGroup(hEle int) int {
	r, _, _ := xListView_DeleteAllGroup.Call(uintptr(hEle))
	return int(r)
}

// 列表视_删除全部项.
//
// hEle: 元素句柄.
func XListView_DeleteAllItem(hEle int) int {
	r, _, _ := xListView_DeleteAllItem.Call(uintptr(hEle))
	return int(r)
}

// 列表视_组删除列.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func XListView_DeleteColumnGroup(hEle int, iColumn int) int {
	r, _, _ := xListView_DeleteColumnGroup.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表视_项删除列.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func XListView_DeleteColumnItem(hEle int, iColumn int) int {
	r, _, _ := xListView_DeleteColumnItem.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表视_项获取文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func XListView_Item_GetTextEx(hEle int, iGroup int, iItem int, pName string) string {
	r, _, _ := xListView_Item_GetTextEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 列表视_项获取图片扩展, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func XListView_Item_GetImageEx(hEle int, iGroup int, iItem int, pName string) int {
	r, _, _ := xListView_Item_GetImageEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), common.StrPtr(pName))
	return int(r)
}

// 列表视_组取文本, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func XListView_Group_GetText(hEle int, iGroup int, iColumn int) string {
	r, _, _ := xListView_Group_GetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 列表视_组取文本扩展, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func XListView_Group_GetTextEx(hEle int, iGroup int, pName string) string {
	r, _, _ := xListView_Group_GetTextEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 列表视_组取图片, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func XListView_Group_GetImage(hEle int, iGroup int, iColumn int) int {
	r, _, _ := xListView_Group_GetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn))
	return int(r)
}

// 列表视_组取图片扩展, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func XListView_Group_GetImageEx(hEle int, iGroup int, pName string) int {
	r, _, _ := xListView_Group_GetImageEx.Call(uintptr(hEle), uintptr(iGroup), common.StrPtr(pName))
	return int(r)
}

// 列表视_项取文本, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XListView_Item_GetText(hEle int, iGroup int, iItem int, iColumn int) string {
	r, _, _ := xListView_Item_GetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn))
	return common.UintPtrToString(r)
}

// 列表视_项取图片, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func XListView_Item_GetImage(hEle int, iGroup int, iItem int, iColumn int) int {
	r, _, _ := xListView_Item_GetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn))
	return int(r)
}

// 列表视_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
//
// width: 线宽度.
func XListView_SetDragRectColor(hEle int, color int, width int) int {
	r, _, _ := xListView_SetDragRectColor.Call(uintptr(hEle), uintptr(color), uintptr(width))
	return int(r)
}

// 列表视_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func XListView_SetItemTemplateXMLFromMem(hEle int, data []byte) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromMem.Call(uintptr(hEle), common.ByteSliceDataPtr(&data), uintptr(len(data)))
	return r != 0
}

// 列表视_置项模板从资源ZIP.
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
func XListView_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromZipRes.Call(uintptr(hEle), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}

// 列表视_取项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func XListView_GetItemTemplate(hEle int) int {
	r, _, _ := xListView_GetItemTemplate.Call(uintptr(hEle))
	return int(r)
}

// 列表视_取项模板组, 返回项模板组句柄.
//
// hEle: 元素句柄.
func XListView_GetItemTemplateGroup(hEle int) int {
	r, _, _ := xListView_GetItemTemplateGroup.Call(uintptr(hEle))
	return int(r)
}
