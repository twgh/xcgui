package xc

import "unsafe"

// 列表视_创建, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func XListView_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xListView_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
// hEle: 元素句柄.
func XListView_CreateAdapter(hEle int) int {
	r, _, _ := xListView_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表视_绑定数据适配器
// hEle: 元素句柄.
// hAdapter: 数据适配器XAdListView.
func XListView_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xListView_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表视_取数据适配器, 返回数据适配器句柄
// hEle: 元素句柄.
func XListView_GetAdapter(hEle int) int {
	r, _, _ := xListView_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置项模板文件
// hEle: 元素句柄.
// pXmlFile: 文件名.
func XListView_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xListView_SetItemTemplateXML.Call(uintptr(hEle), strPtr(pXmlFile))
	return int(r) != 0
}

// 列表视_置项模板从字符串
// hEle: 元素句柄.
// pStringXML: 字符串指针.
func XListView_SetItemTemplateXMLFromString(hEle int, pStringXML int) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromString.Call(uintptr(hEle), uintptr(pStringXML))
	return int(r) != 0
}

// 列表视_置项模板, 置列表项模板
// hEle: 元素句柄.
// hTemp: 模板句柄.
func XListView_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xListView_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return int(r) != 0
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引.
// nTempItemID: 模板项ID.
func XListView_GetTemplateObject(hEle int, iGroup int, iItem int, nTempItemID int) int {
	r, _, _ := xListView_GetTemplateObject.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(nTempItemID))
	return int(r)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄
// hEle: 元素句柄.
// iGroup: 组索引.
// nTempItemID: 模板项ID.
func XListView_GetTemplateObjectGroup(hEle int, iGroup int, nTempItemID int) int {
	r, _, _ := xListView_GetTemplateObjectGroup.Call(uintptr(hEle), uintptr(iGroup), uintptr(nTempItemID))
	return int(r)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项
// hEle: 元素句柄.
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄
// piGroup: 接收组索引.
// piItem: 接收项索引.
func XListView_GetItemIDFromHXCGUI(hEle int, hXCGUI int, piGroup *int, piItem *int) bool {
	r, _, _ := xListView_GetItemIDFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI), uintptr(unsafe.Pointer(piGroup)), uintptr(unsafe.Pointer(piItem)))
	return int(r) != 0
}

// 列表视_测试点击项, 检查坐标点所在项
// hEle: 元素句柄.
// pPt: 坐标点.
// pOutGroup: 接收组索引.
// pOutItem: 接收项索引.
func XListView_HitTest(hEle int, pPt *POINT, pOutGroup *int, pOutItem *int) bool {
	r, _, _ := xListView_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(pOutGroup)), uintptr(unsafe.Pointer(pOutItem)))
	return int(r) != 0
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量
// hEle: 元素句柄.
// pPt: 坐标点.
// pOutGroup: 接收做索引.
// pOutItem: 接收项索引.
func XListView_HitTestOffset(hEle int, pPt *POINT, pOutGroup *int, pOutItem *int) bool {
	r, _, _ := xListView_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)), uintptr(unsafe.Pointer(pOutGroup)), uintptr(unsafe.Pointer(pOutItem)))
	return int(r) != 0
}

// 列表视_启用多选
// hEle: 元素句柄.
// bEnable: 是否启用.
func XListView_EnableMultiSel(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnableMultiSel.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表视_启用模板复用
// hEle: 元素句柄
// bEnable: 是否启用
func XListView_EnablemTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnablemTemplateReuse.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表视_启用虚表
// hEle: 元素句柄
// bEnable: 是否启用
func XListView_EnableVirtualTable(hEle int, bEnable bool) int {
	r, _, _ := xListView_EnableVirtualTable.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表视_置虚表项数量
// hEle: 元素句柄
// iGroup: 组索引
// nCount: 项数量
func XListView_SetVirtualItemCount(hEle int, iGroup int, nCount int) bool {
	r, _, _ := xListView_SetVirtualItemCount.Call(uintptr(hEle), uintptr(iGroup), uintptr(nCount))
	return int(r) != 0
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景
// hEle: 元素句柄.
// nFlags: 标志位: List_DrawItemBk_Flag_
func XListView_SetDrawItemBkFlags(hEle int, nFlags int) int {
	r, _, _ := xListView_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表视_置选择项
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引.
func XListView_SetSelectItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_SetSelectItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r) != 0
}

// 列表视_取选择项
// hEle: 元素句柄.
// piGroup: 接收组索引.
// piItem: 接收项索引.
func XListView_GetSelectItem(hEle int, piGroup *int, piItem *int) bool {
	r, _, _ := xListView_GetSelectItem.Call(uintptr(hEle), uintptr(unsafe.Pointer(piGroup)), uintptr(unsafe.Pointer(piItem)))
	return int(r) != 0
}

// 列表视_添加选择项
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
func XListView_AddSelectItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_AddSelectItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r) != 0
}

// 列表视_显示指定项
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引.
func XListView_VisibleItem(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_VisibleItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_取可视项范围, 获取当前可见项范围
// hEle: 元素句柄
// piGroup1: 可视开始组
// piGroup2: 可视结束组
// piStartGroup: 可视开始组
// piStartItem: 可视开始项
// piEndGroup: 可视结束组
// piEndItem: 可视结束项
func XListView_GetVisibleItemRange(hEle int, piGroup1 *int, piGroup2 *int, piStartGroup *int, piStartItem *int, piEndGroup *int, piEndItem *int) int {
	r, _, _ := xListView_GetVisibleItemRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piGroup1)), uintptr(unsafe.Pointer(piGroup2)), uintptr(unsafe.Pointer(piStartGroup)), uintptr(unsafe.Pointer(piStartItem)), uintptr(unsafe.Pointer(piEndGroup)), uintptr(unsafe.Pointer(piEndItem)))
	return int(r)
}

// 列表视_取选择项数量
// hEle: 元素句柄.
func XListView_GetSelectItemCount(hEle int) int {
	r, _, _ := xListView_GetSelectItemCount.Call(uintptr(hEle))
	return int(r)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量
// hEle: 元素句柄.
// pArray: 数组
// nArraySize: 数组大小.
func XListView_GetSelectAll(hEle int, pArray int, nArraySize int) int {
	r, _, _ := xListView_GetSelectAll.Call(uintptr(hEle), uintptr(pArray), uintptr(nArraySize))
	return int(r)
}

// 列表视_置选择项全部, 选择所有的项
// hEle: 元素句柄.
func XListView_SetSelectAll(hEle int) int {
	r, _, _ := xListView_SetSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_取消选择项全部, 取消选择所有项
// hEle: 元素句柄.
func XListView_CancelSelectAll(hEle int) int {
	r, _, _ := xListView_CancelSelectAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置列间隔, 置列间隔大小
// hEle: 元素句柄.
// space: 间隔大小.
func XListView_SetColumnSpace(hEle int, space int) int {
	r, _, _ := xListView_SetColumnSpace.Call(uintptr(hEle), uintptr(space))
	return int(r)
}

// 列表视_置行间隔, 置行间隔大小
// hEle: 元素句柄.
// space: 间隔大小.
func XListView_SetRowSpace(hEle int, space int) int {
	r, _, _ := xListView_SetRowSpace.Call(uintptr(hEle), uintptr(space))
	return int(r)
}

// 列表视_置项大小
// hEle: 元素句柄.
// width: 宽度.
// height: 高度.
func XListView_SetItemSize(hEle int, width int, height int) int {
	r, _, _ := xListView_SetItemSize.Call(uintptr(hEle), uintptr(width), uintptr(height))
	return int(r)
}

// 列表视_取项大小
// hEle: 元素句柄.
// pSize: 接收返回大小.
func XListView_GetItemSize(hEle int, pSize *SIZE) int {
	r, _, _ := xListView_GetItemSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 列表视_置组高度
// hEle: 元素句柄.
// height: 高度.
func XListView_SetGroupHeight(hEle int, height int) int {
	r, _, _ := xListView_SetGroupHeight.Call(uintptr(hEle), uintptr(height))
	return int(r)
}

// 列表视_取组高度
// hEle: 元素句柄.
func XListView_GetGroupHeight(hEle int) int {
	r, _, _ := xListView_GetGroupHeight.Call(uintptr(hEle))
	return int(r)
}

// 列表视_置组用户数据
// hEle: 元素句柄.
// iGroup: 组索引.
// nData: 数据.
func XListView_SetGroupUserData(hEle int, iGroup int, nData int) int {
	r, _, _ := xListView_SetGroupUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(nData))
	return int(r)
}

// 列表视_置项用户数据
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引.
// nData: 数据.
func XListView_SetItemUserData(hEle int, iGroup int, iItem int, nData int) int {
	r, _, _ := xListView_SetItemUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(nData))
	return int(r)
}

// 列表视_取组用户数据
// hEle: 元素句柄.
// iGroup: 组索引.
func XListView_GetGroupUserData(hEle int, iGroup int) int {
	r, _, _ := xListView_GetGroupUserData.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_取项用户数据
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引.
func XListView_GetItemUserData(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_GetItemUserData.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_添加项背景边框
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func XListView_AddItemBkBorder(hEle int, nState int, color int, alpha uint8, width int) int {
	r, _, _ := xListView_AddItemBkBorder.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha), uintptr(width))
	return int(r)
}

// 列表视_添加项背景填充
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
func XListView_AddItemBkFill(hEle int, nState int, color int, alpha uint8) int {
	r, _, _ := xListView_AddItemBkFill.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha))
	return int(r)
}

// 列表视_添加项背景图片
// hEle: 元素句柄.
// nState: 项状态.
// hImage: 图片句柄.
func XListView_AddItemBkImage(hEle int, nState int, hImage int) int {
	r, _, _ := xListView_AddItemBkImage.Call(uintptr(hEle), uintptr(nState), uintptr(hImage))
	return int(r)
}

// 列表视_取项背景对象数量, 成功返回背景内容数量, 否则返回XC_ID_ERROR
// hEle: 元素句柄.
func XListView_GetItemBkInfoCount(hEle int) int {
	r, _, _ := xListView_GetItemBkInfoCount.Call(uintptr(hEle))
	return int(r)
}

// 列表视_清空项背景对象, 清空项背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
// hEle: 元素句柄.
func XListView_ClearItemBkInfo(hEle int) int {
	r, _, _ := xListView_ClearItemBkInfo.Call(uintptr(hEle))
	return int(r)
}

// 列表视_刷新项数据
// hEle: 元素句柄.
func XListView_RefreshData(hEle int) int {
	r, _, _ := xListView_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI
// hEle: 元素句柄.
// iGroup: 组索引.
// iItem: 项索引, 如果为-1, 代表为组
func XListView_RefreshItem(hEle int, iGroup int, iItem int) int {
	r, _, _ := xListView_RefreshItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE
// hEle: 元素句柄.
// iGroup: 组索引.
// bExpand: 是否展开.
func XListView_ExpandGroup(hEle int, iGroup int, bExpand bool) bool {
	r, _, _ := xListView_ExpandGroup.Call(uintptr(hEle), uintptr(iGroup), boolPtr(bExpand))
	return int(r) != 0
}

// 列表视_组添加列, 返回列索引
// hEle: 元素句柄
// pName: 字段称
func XListView_Group_AddColumn(hEle int, pName string) int {
	r, _, _ := xListView_Group_AddColumn.Call(uintptr(hEle), strPtr(pName))
	return int(r)
}

// 列表视_组添加项文本, 返回组索引
// hEle: 元素句柄
// pValue: 值
// iPos: 插入位置
func XListView_Group_AddItemText(hEle int, pValue string, iPos int) int {
	r, _, _ := xListView_Group_AddItemText.Call(uintptr(hEle), strPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项文本扩展, 返回组索引
// hEle: 元素句柄
// pName: 字段称
// pValue: 值
// iPos: 插入位置
func XListView_Group_AddItemTextEx(hEle int, pName string, pValue string, iPos int) int {
	r, _, _ := xListView_Group_AddItemTextEx.Call(uintptr(hEle), strPtr(pName), strPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项图片, 返回组索引
// hEle: 元素句柄
// hImage: 图片句柄
// iPos: 插入位置
func XListView_Group_AddItemImage(hEle int, hImage int, iPos int) int {
	r, _, _ := xListView_Group_AddItemImage.Call(uintptr(hEle), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_组添加项图片扩展, 返回组索引
// hEle: 元素句柄
// pName: 字段称
// hImage: 图片句柄
// iPos: 插入位置
func XListView_Group_AddItemImageEx(hEle int, pName string, hImage int, iPos int) int {
	r, _, _ := xListView_Group_AddItemImageEx.Call(uintptr(hEle), strPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_组置文本
// hEle: 元素句柄
// iGroup: 组索引
// iColumn: 列索引
// pValue: 值
func XListView_Group_SetText(hEle int, iGroup int, iColumn int, pValue string) bool {
	r, _, _ := xListView_Group_SetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn), strPtr(pValue))
	return int(r) != 0
}

// 列表视_组置文本扩展
// hEle: 元素句柄
// iGroup: 组索引
// pName: 字段名
// pValue: 值
func XListView_Group_SetTextEx(hEle int, iGroup int, pName string, pValue string) bool {
	r, _, _ := xListView_Group_SetTextEx.Call(uintptr(hEle), uintptr(iGroup), strPtr(pName), strPtr(pValue))
	return int(r) != 0
}

// 列表视_组置图片
// hEle: 元素句柄
// iGroup: 组索引
// iColumn: 列索引
// hImage: 图片句柄
func XListView_Group_SetImage(hEle int, iGroup int, iColumn int, hImage int) bool {
	r, _, _ := xListView_Group_SetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iColumn), uintptr(hImage))
	return int(r) != 0
}

// 列表视_组置图片扩展
// hEle: 元素句柄
// iGroup: 组索引
// pName: 字段名
// hImage: 图片句柄
func XListView_Group_SetImageEx(hEle int, iGroup int, pName string, hImage int) bool {
	r, _, _ := xListView_Group_SetImageEx.Call(uintptr(hEle), uintptr(iGroup), strPtr(pName), uintptr(hImage))
	return int(r) != 0
}

// 列表视_组获取数量, 返回组数量
// hEle: 元素句柄
func XListView_Group_GetCount(hEle int) int {
	r, _, _ := xListView_Group_GetCount.Call(uintptr(hEle))
	return int(r)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 XC_ID_ERROR
// hEle: 元素句柄
// iGroup: 组索引
func XListView_Item_GetCount(hEle int, iGroup int) int {
	r, _, _ := xListView_Item_GetCount.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_项添加列, 返回列索引
// hEle: 元素句柄
// pName: 字段名
func XListView_Item_AddColumn(hEle int, pName string) int {
	r, _, _ := xListView_Item_AddColumn.Call(uintptr(hEle), strPtr(pName))
	return int(r)
}

// 列表视_项添加文本, 返回项索引
// hEle: 元素句柄
// iGroup: 组索引
// pValue: 值
// iPos: 插入位置, -1添加到末尾
func XListView_Item_AddItemText(hEle int, iGroup int, pValue string, iPos int) int {
	r, _, _ := xListView_Item_AddItemText.Call(uintptr(hEle), uintptr(iGroup), strPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_项添加文本扩展, 返回项索引
// hEle: 元素句柄
// iGroup: 组索引
// pName: 字段名
// pValue: 值
// iPos: 插入位置, -1添加到末尾
func XListView_Item_AddItemTextEx(hEle int, iGroup int, pName string, pValue string, iPos int) int {
	r, _, _ := xListView_Item_AddItemTextEx.Call(uintptr(hEle), uintptr(iGroup), strPtr(pName), strPtr(pValue), uintptr(iPos))
	return int(r)
}

// 列表视_项添加图片, 返回项索引
// hEle: 元素句柄
// iGroup: 组索引
// hImage: 图片句柄
// iPos: 插入位置, -1添加到末尾
func XListView_Item_AddItemImage(hEle int, iGroup int, hImage int, iPos int) int {
	r, _, _ := xListView_Item_AddItemImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_项添加图片扩展, 返回项索引
// hEle: 元素句柄
// iGroup: 组索引
// pName: 字段名
// hImage: 图片句柄
// iPos: 插入位置, -1添加到末尾
func XListView_Item_AddItemImageEx(hEle int, iGroup int, pName string, hImage int, iPos int) int {
	r, _, _ := xListView_Item_AddItemImageEx.Call(uintptr(hEle), uintptr(iGroup), strPtr(pName), uintptr(hImage), uintptr(iPos))
	return int(r)
}

// 列表视_项置文本
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// iColumn: 列索引
// pValue: 值
func XListView_Item_SetText(hEle int, iGroup int, iItem int, iColumn int, pValue string) bool {
	r, _, _ := xListView_Item_SetText.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), strPtr(pValue))
	return int(r) != 0
}

// 列表视_项置文本扩展
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// pName: 字段名
// pValue: 值
func XListView_Item_SetTextEx(hEle int, iGroup int, iItem int, pName string, pValue string) bool {
	r, _, _ := xListView_Item_SetTextEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), strPtr(pName), strPtr(pValue))
	return int(r) != 0
}

// 列表视_项置图片
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// iColumn: 列索引
// hImage: 图片句柄
func XListView_Item_SetImage(hEle int, iGroup int, iItem int, iColumn int, hImage int) bool {
	r, _, _ := xListView_Item_SetImage.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), uintptr(iColumn), uintptr(hImage))
	return int(r) != 0
}

// 列表视_项置图片扩展
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// pName: 列名称
// hImage: 图片句柄
func XListView_Item_SetImageEx(hEle int, iGroup int, iItem int, pName string, hImage int) bool {
	r, _, _ := xListView_Item_SetImageEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), strPtr(pName), uintptr(hImage))
	return int(r) != 0
}

// 列表视_组删除项
// hEle: 元素句柄
// iGroup: 组索引
func XListView_Group_DeleteItem(hEle int, iGroup int) bool {
	r, _, _ := xListView_Group_DeleteItem.Call(uintptr(hEle), uintptr(iGroup))
	return int(r) != 0
}

// 列表视_组删除全部子项
// hEle: 元素句柄
// iGroup: 组索引
func XListView_Group_DeleteAllChildItem(hEle int, iGroup int) int {
	r, _, _ := xListView_Group_DeleteAllChildItem.Call(uintptr(hEle), uintptr(iGroup))
	return int(r)
}

// 列表视_项删除
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
func XListView_Item_DeleteItem(hEle int, iGroup int, iItem int) bool {
	r, _, _ := xListView_Item_DeleteItem.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem))
	return int(r) != 0
}

// 列表视_删除全部
// hEle: 元素句柄
func XListView_DeleteAll(hEle int) int {
	r, _, _ := xListView_DeleteAll.Call(uintptr(hEle))
	return int(r)
}

// 列表视_删除全部组
// hEle: 元素句柄
func XListView_DeleteAllGroup(hEle int) int {
	r, _, _ := xListView_DeleteAllGroup.Call(uintptr(hEle))
	return int(r)
}

// 列表视_删除全部项
// hEle: 元素句柄
func XListView_DeleteAllItem(hEle int) int {
	r, _, _ := xListView_DeleteAllItem.Call(uintptr(hEle))
	return int(r)
}

// 列表视_组删除列
// hEle: 元素句柄
// iColumn: 列索引
func XListView_DeleteColumnGroup(hEle int, iColumn int) int {
	r, _, _ := xListView_DeleteColumnGroup.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表视_项删除列
// hEle: 元素句柄
// iColumn: 列索引
func XListView_DeleteColumnItem(hEle int, iColumn int) int {
	r, _, _ := xListView_DeleteColumnItem.Call(uintptr(hEle), uintptr(iColumn))
	return int(r)
}

// 列表视_项获取文本
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// pName: 字段称
func XListView_Item_GetTextEx(hEle int, iGroup int, iItem int, pName string) string {
	r, _, _ := xListView_Item_GetTextEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), strPtr(pName))
	return uintPtrToString(r)
}

// 列表视_项获取图片扩展, 返回图片句柄
// hEle: 元素句柄
// iGroup: 组索引
// iItem: 项索引
// pName: 字段称
func XListView_Item_GetImageEx(hEle int, iGroup int, iItem int, pName string) int {
	r, _, _ := xListView_Item_GetImageEx.Call(uintptr(hEle), uintptr(iGroup), uintptr(iItem), strPtr(pName))
	return int(r)
}
