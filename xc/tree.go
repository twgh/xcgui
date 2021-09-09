package xc

import "unsafe"

// 列表树_创建, 创建树元素, 返回元素句柄.
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func XTree_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xTree_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 列表树_启用拖动项, 启用拖动项功能
// hEle: 元素句柄.
// bEnable: 是否启用.
func XTree_EnableDragItem(hEle int, bEnable bool) int {
	r, _, _ := xTree_EnableDragItem.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表树_启用连接线, 启用或禁用显示项的连接线
// hEle: 元素句柄.
// bEnable: 是否启用.
// bSolid: 实线或虚线; TRUE: 实线, FALSE: 虚线
func XTree_EnableConnectLine(hEle int, bEnable bool, bSolid bool) int {
	r, _, _ := xTree_EnableConnectLine.Call(uintptr(hEle), boolPtr(bEnable), boolPtr(bSolid))
	return int(r)
}

// 列表树_启用展开, 启动或关闭默认展开功能, 如果开启新插入的项将自动展开
// hEle: 元素句柄.
// bEnable: 是否启用.
func XTree_EnableExpand(hEle int, bEnable bool) int {
	r, _, _ := xTree_EnableExpand.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表树_启用模板复用
// hEle: 元素句柄
// bEnable: 是否启用
func XTree_EnablemTemplateReuse(hEle int, bEnable bool) int {
	r, _, _ := xTree_EnablemTemplateReuse.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 列表树_置连接线颜色
// hEle: 元素句柄.
// color: RGB颜色.
// alpha: 透明度.
func XTree_SetConnectLineColor(hEle int, color int, alpha uint8) int {
	r, _, _ := xTree_SetConnectLineColor.Call(uintptr(hEle), uintptr(color), uintptr(alpha))
	return int(r)
}

// 列表树_置展开按钮大小, 设置展开按钮占用空间大小
// hEle: 元素句柄.
// nWidth: 宽度.
// nHeight: 高度.
func XTree_SetExpandButtonSize(hEle int, nWidth int, nHeight int) int {
	r, _, _ := xTree_SetExpandButtonSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 列表树_置连接线长度, 设置连线绘制长度, 展开按钮与项内容之间的连线.
// hEle: 元素句柄.
// nLength: 连线绘制长度.
func XTree_SetConnectLineLength(hEle int, nLength int) int {
	r, _, _ := xTree_SetConnectLineLength.Call(uintptr(hEle), uintptr(nLength))
	return int(r)
}

// 列表树_置拖动项插入位置颜色, 设置拖动项插入位置颜色提示
// hEle: 元素句柄.
// color: RGB颜色.
// alpha: 透明度.
func XTree_SetDragInsertPositionColor(hEle int, color int, alpha uint8) int {
	r, _, _ := xTree_SetDragInsertPositionColor.Call(uintptr(hEle), uintptr(color), uintptr(alpha))
	return int(r)
}

// 列表树_置项模板文件
// hEle: 元素句柄.
// pXmlFile: 文件名.
func XTree_SetItemTemplateXML(hEle int, pXmlFile string) bool {
	r, _, _ := xTree_SetItemTemplateXML.Call(uintptr(hEle), strPtr(pXmlFile))
	return int(r) != 0
}

// 列表树_置选择项模板文件, 设置项模板文件, 项选中状态
// hEle: 元素句柄.
// pXmlFile: 文件名.
func XTree_SetItemTemplateXMLSel(hEle int, pXmlFile string) bool {
	r, _, _ := xTree_SetItemTemplateXMLSel.Call(uintptr(hEle), strPtr(pXmlFile))
	return int(r) != 0
}

// 列表树_置项模板
// hEle: 元素句柄.
// hTemp: 模板句柄.
func XTree_SetItemTemplate(hEle int, hTemp int) bool {
	r, _, _ := xTree_SetItemTemplate.Call(uintptr(hEle), uintptr(hTemp))
	return int(r) != 0
}

// 列表树_置选择项模板, 设置列表项模板, 项选中状态
// hEle: 元素句柄.
// hTemp: 模板句柄.
func XTree_SetItemTemplateSel(hEle int, hTemp int) bool {
	r, _, _ := xTree_SetItemTemplateSel.Call(uintptr(hEle), uintptr(hTemp))
	return int(r) != 0
}

// 列表树_置项模板从字符串, 设置项模板文件
// hEle: 元素句柄.
// pStringXML: 字符串指针.
func XTree_SetItemTemplateXMLFromString(hEle int, pStringXML int) bool {
	r, _, _ := xTree_SetItemTemplateXMLFromString.Call(uintptr(hEle), uintptr(pStringXML))
	return int(r) != 0
}

// 列表树_置选择项模板从字符串, 设置项模板文件, 项选中状态
// hEle: 元素句柄.
// pStringXML: 字符串指针.
func XTree_SetItemTemplateXMLSelFromString(hEle int, pStringXML int) bool {
	r, _, _ := xTree_SetItemTemplateXMLSelFromString.Call(uintptr(hEle), uintptr(pStringXML))
	return int(r) != 0
}

// 列表树_置项背景绘制标志, 设置是否绘制指定状态下项的背景
// hEle: 元素句柄.
// nFlags: 标志位: List_DrawItemBk_Flag_.
func XTree_SetDrawItemBkFlags(hEle int, nFlags int) int {
	r, _, _ := xTree_SetDrawItemBkFlags.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 列表树_置项数据, 设置项用户数据
// hEle: 元素句柄.
// nID: 项ID.
// nUserData: 用户数据.
func XTree_SetItemData(hEle int, nID int, nUserData int) bool {
	r, _, _ := xTree_SetItemData.Call(uintptr(hEle), uintptr(nID), uintptr(nUserData))
	return int(r) != 0
}

// 列表树_取项数据, 获取项用户数据
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetItemData(hEle int, nID int) int {
	r, _, _ := xTree_GetItemData.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_置选择项
// hEle: 元素句柄.
// nID: 项ID.
func XTree_SetSelectItem(hEle int, nID int) bool {
	r, _, _ := xTree_SetSelectItem.Call(uintptr(hEle), uintptr(nID))
	return int(r) != 0
}

// 列表树_取选择项, 返回项ID
// hEle: 元素句柄.
func XTree_GetSelectItem(hEle int) int {
	r, _, _ := xTree_GetSelectItem.Call(uintptr(hEle))
	return int(r)
}

// 列表树_可视指定项, 滚动视图让指定项可见
// hEle: 元素句柄.
// nID: 项索引.
func XTree_VisibleItem(hEle int, nID int) int {
	r, _, _ := xTree_VisibleItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_判断展开
// hEle: 元素句柄.
// nID: 项ID.
func XTree_IsExpand(hEle int, nID int) bool {
	r, _, _ := xTree_IsExpand.Call(uintptr(hEle), uintptr(nID))
	return int(r) != 0
}

// 列表树_展开项, 判断项是否展开
// hEle: 元素句柄.
// nID: 项ID.
// bExpand: 是否展开.
func XTree_ExpandItem(hEle int, nID int, bExpand bool) bool {
	r, _, _ := xTree_ExpandItem.Call(uintptr(hEle), uintptr(nID), boolPtr(bExpand))
	return int(r) != 0
}

// 列表树_展开全部子项, 展开所有的子项
// hEle: 元素句柄.
// nID: 项ID.
// bExpand: 是否展开.
func XTree_ExpandAllChildItem(hEle int, nID int, bExpand bool) bool {
	r, _, _ := xTree_ExpandAllChildItem.Call(uintptr(hEle), uintptr(nID), boolPtr(bExpand))
	return int(r) != 0
}

// 列表树_测试点击项, 检测坐标点所在项, 返回项ID
// hEle: 元素句柄.
// pPt: 坐标点.
func XTree_HitTest(hEle int, pPt *POINT) int {
	r, _, _ := xTree_HitTest.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 列表树_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移坐标, 返回项ID
// hEle: 元素句柄.
// pPt: 坐标点.
func XTree_HitTestOffset(hEle int, pPt *POINT) int {
	r, _, _ := xTree_HitTestOffset.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 列表树_取第一个子项, 获取第一个子项. 成功返回项ID, 失败返回XC_ID_ERROR.
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetFirstChildItem(hEle int, nID int) int {
	r, _, _ := xTree_GetFirstChildItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_取末尾子项, 获取末尾子项. 成功返回项ID, 失败返回XC_ID_ERROR.
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetEndChildItem(hEle int, nID int) int {
	r, _, _ := xTree_GetEndChildItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_取上一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetPrevSiblingItem(hEle int, nID int) int {
	r, _, _ := xTree_GetPrevSiblingItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_取下一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetNextSiblingItem(hEle int, nID int) int {
	r, _, _ := xTree_GetNextSiblingItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_取父项, 成功返回项ID, 失败返回XC_ID_ERROR.
// hEle: 元素句柄.
// nID: 项ID.
func XTree_GetParentItem(hEle int, nID int) int {
	r, _, _ := xTree_GetParentItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
// hEle: 元素句柄.
func XTree_CreateAdapter(hEle int) int {
	r, _, _ := xTree_CreateAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表树_绑定数据适配器
// hEle: 元素句柄.
// hAdapter: 数据适配器句柄, XAdTree
func XTree_BindAdapter(hEle int, hAdapter int) int {
	r, _, _ := xTree_BindAdapter.Call(uintptr(hEle), uintptr(hAdapter))
	return int(r)
}

// 列表树_取数据视频器, 返回数据适配器句柄
// hEle: 元素句柄.
func XTree_GetAdapter(hEle int) int {
	r, _, _ := xTree_GetAdapter.Call(uintptr(hEle))
	return int(r)
}

// 列表树_刷新数据, 刷新所有项模板, 以便更新UI
// hEle: 元素句柄.
func XTree_RefreshData(hEle int) int {
	r, _, _ := xTree_RefreshData.Call(uintptr(hEle))
	return int(r)
}

// 列表树_刷新指定项, 刷新指定项模板, 以便更新UI
// hEle: 元素句柄.
// nID: 项ID.
func XTree_RefreshItem(hEle int, nID int) int {
	r, _, _ := xTree_RefreshItem.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 列表树_置缩进, 设置缩进大小
// hEle: 元素句柄.
// nWidth: 缩进宽度.
func XTree_SetIndentation(hEle int, nWidth int) int {
	r, _, _ := xTree_SetIndentation.Call(uintptr(hEle), uintptr(nWidth))
	return int(r)
}

// 列表树_取缩进, 返回缩进值大小
// hEle: 元素句柄.
func XTree_GetIndentation(hEle int) int {
	r, _, _ := xTree_GetIndentation.Call(uintptr(hEle))
	return int(r)
}

// 列表树_置项默认高度
// hEle: 元素句柄.
// nHeight: 高度.
// nSelHeight: 选中时高度.
func XTree_SetItemHeightDefault(hEle int, nHeight int, nSelHeight int) int {
	r, _, _ := xTree_SetItemHeightDefault.Call(uintptr(hEle), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表树_取项默认高度
// hEle: 元素句柄.
// pHeight: 接收返回高度.
// pSelHeight: 接收返回值, 当项选中时的高度
func XTree_GetItemHeightDefault(hEle int, pHeight *int, pSelHeight *int) int {
	r, _, _ := xTree_GetItemHeightDefault.Call(uintptr(hEle), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表树_置项高度
// hEle: 元素句柄.
// nID: 项ID.
// nHeight: 高度.
// nSelHeight: 选中时高度.
func XTree_SetItemHeight(hEle int, nID int, nHeight int, nSelHeight int) int {
	r, _, _ := xTree_SetItemHeight.Call(uintptr(hEle), uintptr(nID), uintptr(nHeight), uintptr(nSelHeight))
	return int(r)
}

// 列表树_取项高度
// hEle: 元素句柄.
// nID: 项ID.
// pHeight: 接收返回高度.
// pSelHeight: 接收返回值, 当项选中时的高度
func XTree_GetItemHeight(hEle int, nID int, pHeight *int, pSelHeight *int) int {
	r, _, _ := xTree_GetItemHeight.Call(uintptr(hEle), uintptr(nID), uintptr(unsafe.Pointer(pHeight)), uintptr(unsafe.Pointer(pSelHeight)))
	return int(r)
}

// 列表树_置行间距
// hEle: 元素句柄.
// nSpace: 行间隔大小.
func XTree_SetRowSpace(hEle int, nSpace int) int {
	r, _, _ := xTree_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 列表树_取行间距
// hEle: 元素句柄.
func XTree_GetRowSpace(hEle int) int {
	r, _, _ := xTree_GetRowSpace.Call(uintptr(hEle))
	return int(r)
}

// 列表树_移动项, 移动项的位置
// hEle: 元素句柄.
// nMoveItem: 要移动的项ID.
// nDestItem: 目标项ID, 参照位置
// nFlag: 0:目标前面, 1:目标后面, 2:目标子项首, 3:目标子项尾
func XTree_MoveItem(hEle int, nMoveItem int, nDestItem int, nFlag int) bool {
	r, _, _ := xTree_MoveItem.Call(uintptr(hEle), uintptr(nMoveItem), uintptr(nDestItem), uintptr(nFlag))
	return int(r) != 0
}

// 列表树_添加项背景边框
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func XTree_AddItemBkBorder(hEle int, nState int, color int, alpha uint8, width int) int {
	r, _, _ := xTree_AddItemBkBorder.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha), uintptr(width))
	return int(r)
}

// 列表树_添加项背景填充
// hEle: 元素句柄.
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
func XTree_AddItemBkFill(hEle int, nState int, color int, alpha uint8) int {
	r, _, _ := xTree_AddItemBkFill.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(alpha))
	return int(r)
}

// 列表树_添加项背景图片
// hEle: 元素句柄.
// nState: 项状态.
// hImage: 图片句柄.
func XTree_AddItemBkImage(hEle int, nState int, hImage int) int {
	r, _, _ := xTree_AddItemBkImage.Call(uintptr(hEle), uintptr(nState), uintptr(hImage))
	return int(r)
}

// 列表树_取项背景对象数量, 返回项背景内容数量
// hEle: 元素句柄.
func XTree_GetItemBkInfoCount(hEle int) int {
	r, _, _ := xTree_GetItemBkInfoCount.Call(uintptr(hEle))
	return int(r)
}

// 列表树_清空项背景对象, 清空项背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
// hEle: 元素句柄.
func XTree_ClearItemBkInfo(hEle int) int {
	r, _, _ := xTree_ClearItemBkInfo.Call(uintptr(hEle))
	return int(r)
}

// 列表树_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄
// hEle: 元素句柄.
// nID: 树项ID.
// nTempItemID: 模板项ID.
func XTree_GetTemplateObject(hEle int, nID int, nTempItemID int) int {
	r, _, _ := xTree_GetTemplateObject.Call(uintptr(hEle), uintptr(nID), uintptr(nTempItemID))
	return int(r)
}

// 列表树_取对象所在项, 获取当前对象所在模板实例, 属于列表树中哪一个项. 成功返回项ID, 否则返回XC_ID_ERROR.
// hEle: 元素句柄.
// hXCGUI: 对象句柄
func XTree_GetItemIDFromHXCGUI(hEle int, hXCGUI int) int {
	r, _, _ := xTree_GetItemIDFromHXCGUI.Call(uintptr(hEle), uintptr(hXCGUI))
	return int(r)
}

// 列表树_插入项文本
// hEle:
// pValue:
// nParentID:
// insertID:
func XTree_InsertItemText(hEle int, pValue string, nParentID int, insertID int) int {
	r, _, _ := xTree_InsertItemText.Call(uintptr(hEle), strPtr(pValue), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 列表树_插入项文本扩展
// hEle:
// pName:
// pValue:
// nParentID:
// insertID:
func XTree_InsertItemTextEx(hEle int, pName string, pValue string, nParentID int, insertID int) int {
	r, _, _ := xTree_InsertItemTextEx.Call(uintptr(hEle), strPtr(pName), strPtr(pValue), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 列表树_插入项图片
// hEle:
// hImage:
// nParentID:
// insertID:
func XTree_InsertItemImage(hEle int, hImage int, nParentID int, insertID int) int {
	r, _, _ := xTree_InsertItemImage.Call(uintptr(hEle), uintptr(hImage), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 列表树_插入项图片扩展
// hEle:
// pName:
// hImage:
// nParentID:
// insertID:
func XTree_InsertItemImageEx(hEle int, pName string, hImage int, nParentID int, insertID int) int {
	r, _, _ := xTree_InsertItemImageEx.Call(uintptr(hEle), strPtr(pName), uintptr(hImage), uintptr(nParentID), uintptr(insertID))
	return int(r)
}

// 列表树_取项数量
// hEle:
func XTree_GetCount(hEle int) int {
	r, _, _ := xTree_GetCount.Call(uintptr(hEle))
	return int(r)
}

// 列表树_取列数量
// hEle:
func XTree_GetCountColumn(hEle int) int {
	r, _, _ := xTree_GetCountColumn.Call(uintptr(hEle))
	return int(r)
}

// 列表树_置项文本
// hEle:
// nID:
// iColumn:
// pValue:
func XTree_SetItemText(hEle int, nID int, iColumn int, pValue string) bool {
	r, _, _ := xTree_SetItemText.Call(uintptr(hEle), uintptr(nID), uintptr(iColumn), strPtr(pValue))
	return int(r) != 0
}

// 列表树_置项文本扩展
// hEle:
// nID:
// pName:
// pValue:
func XTree_SetItemTextEx(hEle int, nID int, pName string, pValue string) bool {
	r, _, _ := xTree_SetItemTextEx.Call(uintptr(hEle), uintptr(nID), strPtr(pName), strPtr(pValue))
	return int(r) != 0
}

// 列表树_置项图片
// hEle:
// nID:
// iColumn:
// hImage:
func XTree_SetItemImage(hEle int, nID int, iColumn int, hImage int) bool {
	r, _, _ := xTree_SetItemImage.Call(uintptr(hEle), uintptr(nID), uintptr(iColumn), uintptr(hImage))
	return int(r) != 0
}

// 列表树_置项图片扩展
// hEle:
// nID:
// pName:
// hImage:
func XTree_SetItemImageEx(hEle int, nID int, pName string, hImage int) bool {
	r, _, _ := xTree_SetItemImageEx.Call(uintptr(hEle), uintptr(nID), strPtr(pName), uintptr(hImage))
	return int(r) != 0
}

// 列表树_取项文本
// hEle:
// nID:
// iColumn:
func XTree_GetItemText(hEle int, nID int, iColumn int) string {
	r, _, _ := xTree_GetItemText.Call(uintptr(hEle), uintptr(nID), uintptr(iColumn))
	return uintPtrToString(r)
}

// 列表树_取项文本扩展
// hEle:
// nID:
// pName:
func XTree_GetItemTextEx(hEle int, nID int, pName string) string {
	r, _, _ := xTree_GetItemTextEx.Call(uintptr(hEle), uintptr(nID), strPtr(pName))
	return uintPtrToString(r)
}

// 列表树_取项图片
// hEle:
// nID:
// iColumn:
func XTree_GetItemImage(hEle int, nID int, iColumn int) int {
	r, _, _ := xTree_GetItemImage.Call(uintptr(hEle), uintptr(nID), uintptr(iColumn))
	return int(r)
}

// 列表树_取项图片扩展
// hEle:
// nID:
// pName:
func XTree_GetItemImageEx(hEle int, nID int, pName string) int {
	r, _, _ := xTree_GetItemImageEx.Call(uintptr(hEle), uintptr(nID), strPtr(pName))
	return int(r)
}

// 列表树_删除项
// hEle:
// nID:
func XTree_DeleteItem(hEle int, nID int) bool {
	r, _, _ := xTree_DeleteItem.Call(uintptr(hEle), uintptr(nID))
	return int(r) != 0
}

// 列表树_删除全部项
// hEle:
func XTree_DeleteItemAll(hEle int) int {
	r, _, _ := xTree_DeleteItemAll.Call(uintptr(hEle))
	return int(r)
}

// 列表树_删除列全部
// hEle:
func XTree_DeleteColumnAll(hEle int) int {
	r, _, _ := xTree_DeleteColumnAll.Call(uintptr(hEle))
	return int(r)
}
