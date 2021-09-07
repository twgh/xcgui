package widget

import (
	"github.com/twgh/xcgui/xc"
)

type ListBox struct {
	Element

	HEle int
}

// 列表框_创建, 创建列表框元素, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素
func NewListBox(x int, y int, cx int, cy int, hParent int) *ListBox {
	p := &ListBox{
		HEle: xc.XListBox_Create(x, y, cx, cy, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 列表框_启用固定行高
// bEnable: 是否启用
func (l *ListBox) EnableFixedRowHeight(bEnable bool) int {
	return xc.XListBox_EnableFixedRowHeight(l.HEle, bEnable)
}

// 列表框_启用模板复用
// bEnable: 是否启用
func (l *ListBox) EnablemTemplateReuse(bEnable bool) int {
	return xc.XListBox_EnablemTemplateReuse(l.HEle, bEnable)
}

// 列表框_启用虚表
// bEnable: 是否启用
func (l *ListBox) EnableVirtualTable(bEnable bool) int {
	return xc.XListBox_EnableVirtualTable(l.HEle, bEnable)
}

// 列表框_置虚表行数
// nRowCount: 行数
func (l *ListBox) SetVirtualRowCount(nRowCount int) int {
	return xc.XListBox_SetVirtualRowCount(l.HEle, nRowCount)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景
// nFlags: 标志位, List_DrawItemBk_Flag_
func (l *ListBox) SetDrawItemBkFlags(nFlags int) int {
	return xc.XListBox_SetDrawItemBkFlags(l.HEle, nFlags)
}

// 列表框_置项数据, 设置项用户数据
// iItem: 想索引.
// nUserData: 用户数据.
func (l *ListBox) SetItemData(iItem int, nUserData int) bool {
	return xc.XListBox_SetItemData(l.HEle, iItem, nUserData)
}

// 列表框_取项数据, 获取项用户数据
// iItem: 项索引.
func (l *ListBox) GetItemData(iItem int) int {
	return xc.XListBox_GetItemData(l.HEle, iItem)
}

// 列表框_添加项背景边框, 添加项背景内容边框
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (l *ListBox) AddItemBkBorder(nState int, color int, alpha uint8, width int) int {
	return xc.XListBox_AddItemBkBorder(l.HEle, nState, color, alpha, width)
}

// 列表框_添加项背景填充, 添加项背景内容填充
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
func (l *ListBox) AddItemBkFill(nState int, color int, alpha uint8) int {
	return xc.XListBox_AddItemBkFill(l.HEle, nState, color, alpha)
}

// 列表框_添加项背景图片, 添加项背景内容图片
// nState: 项状态.
// hImage: 图片句柄.
func (l *ListBox) AddItemBkImage(nState int, hImage int) int {
	return xc.XListBox_AddItemBkImage(l.HEle, nState, hImage)
}

// 列表框_取项背景对象数量, 获取背景内容数量, 成功返回背景内容数量, 否则返回XC_ID_ERROR
func (l *ListBox) GetItemBkInfoCount() int {
	return xc.XListBox_GetItemBkInfoCount(l.HEle)
}

// 列表框_清空项背景对象, 清空项背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确
func (l *ListBox) ClearItemBkInfo() int {
	return xc.XListBox_ClearItemBkInfo(l.HEle)
}

// 列表框_置项信息
// iItem: 项索引.
// pItem: 项信息.
func (l *ListBox) SetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_SetItemInfo(l.HEle, iItem, pItem)
}

// 列表框_取项背景信息, 获取项信息
// iItem: 项索引.
// pItem: 项信息.
func (l *ListBox) GetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_GetItemInfo(l.HEle, iItem, pItem)
}

// 列表框_置选择项, 设置选择选
// iItem: 项索引.
func (l *ListBox) SetSelectItem(iItem int) bool {
	return xc.XListBox_SetSelectItem(l.HEle, iItem)
}

// 列表框_取选择项, 返回项索引
func (l *ListBox) GetSelectItem() int {
	return xc.XListBox_GetSelectItem(l.HEle)
}

// 列表框_添加选择项
// iItem: 项索引
func (l *ListBox) AddSelectItem(iItem int) bool {
	return xc.XListBox_AddSelectItem(l.HEle, iItem)
}

// 列表框_取消选择项
// iItem: 项索引.
func (l *ListBox) CancelSelectItem(iItem int) bool {
	return xc.XListBox_CancelSelectItem(l.HEle, iItem)
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE
func (l *ListBox) CancelSelectAll() bool {
	return xc.XListBox_CancelSelectAll(l.HEle)
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量
// pArray: 数组缓冲区.
// nArraySize: 数组大小.
func (l *ListBox) GetSelectAll(pArray int, nArraySize int) int {
	return xc.XListBox_GetSelectAll(l.HEle, pArray, nArraySize)
}

// 列表框_取选择项数量, 获取选择项数量
func (l *ListBox) GetSelectCount() int {
	return xc.XListBox_GetSelectCount(l.HEle)
}

// 列表框_取鼠标停留项, 返回鼠标所在项
func (l *ListBox) GetItemMouseStay() int {
	return xc.XListBox_GetItemMouseStay(l.HEle)
}

// 列表框_选择全部项
func (l *ListBox) SelectAll() bool {
	return xc.XListBox_SelectAll(l.HEle)
}

// 列表框_显示指定项, 滚动视图让指定项可见
// iItem: 项索引.
func (l *ListBox) VisibleItem(iItem int) int {
	return xc.XListBox_VisibleItem(l.HEle, iItem)
}

// 列表框_取可视行范围, 获取当前可见行范围
// piStart: 开始行索引
// piEnd: 结束行索引
func (l *ListBox) GetVisibleRowRange(piStart *int, piEnd *int) int {
	return xc.XListBox_GetVisibleRowRange(l.HEle, piStart, piEnd)
}

// 列表框_置项默认高度
// nHeight: 项高度.
// nSelHeight: 选中项高度.
func (l *ListBox) SetItemHeightDefault(nHeight int, nSelHeight int) int {
	return xc.XListBox_SetItemHeightDefault(l.HEle, nHeight, nSelHeight)
}

// 列表框_取项默认高度
// pHeight: 高度.
// pSelHeight: 选中时高度.
func (l *ListBox) GetItemHeightDefault(pHeight *int, pSelHeight *int) int {
	return xc.XListBox_GetItemHeightDefault(l.HEle, pHeight, pSelHeight)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄
func (l *ListBox) GetItemIndexFromHXCGUI(hXCGUI int) int {
	return xc.XListBox_GetItemIndexFromHXCGUI(l.HEle, hXCGUI)
}

// 列表框_置行间距
// nSpace: 间距大小.
func (l *ListBox) SetRowSpace(nSpace int) int {
	return xc.XListBox_SetRowSpace(l.HEle, nSpace)
}

// 列表框_取行间距
func (l *ListBox) GetRowSpace() int {
	return xc.XListBox_GetRowSpace(l.HEle)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引
// pPt: 坐标点.
func (l *ListBox) HitTest(pPt *xc.POINT) int {
	return xc.XListBox_HitTest(l.HEle, pPt)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引
// pPt: 坐标点.
func (l *ListBox) HitTestOffset(pPt *xc.POINT) int {
	return xc.XListBox_HitTestOffset(l.HEle, pPt)
}

// 列表框_置项模板文件, 设置列表项模板文件
// pXmlFile: 文件名.
func (l *ListBox) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XListBox_SetItemTemplateXML(l.HEle, pXmlFile)
}

// 列表框_置项模板, 设置列表项模板
// hTemp: 模板句柄
func (l *ListBox) SetItemTemplate(hTemp int) bool {
	return xc.XListBox_SetItemTemplate(l.HEle, hTemp)
}

// 列表框_置项模板从字符串, 设置项布局模板文件
// pStringXML: 字符串指针.
func (l *ListBox) SetItemTemplateXMLFromString(pStringXML int) bool {
	return xc.XListBox_SetItemTemplateXMLFromString(l.HEle, pStringXML)
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL
// iItem: 项索引.
// nTempItemID: 模板项ID.
func (l *ListBox) GetTemplateObject(iItem int, nTempItemID int) int {
	return xc.XListBox_GetTemplateObject(l.HEle, iItem, nTempItemID)
}

// 列表框_启用多选, 是否启用多行选择功能
// bEnable: 是否启用.
func (l *ListBox) EnableMultiSel(bEnable bool) int {
	return xc.XListBox_EnableMultiSel(l.HEle, bEnable)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄
func (l *ListBox) CreateAdapter() int {
	return xc.XListBox_CreateAdapter(l.HEle)
}

// 列表框_绑定数据适配器, 绑定数据适配器
// hAdapter: 数据适配器句柄 XAdTable.
func (l *ListBox) BindAdapter(hAdapter int) int {
	return xc.XListBox_BindAdapter(l.HEle, hAdapter)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄
func (l *ListBox) GetAdapter() int {
	return xc.XListBox_GetAdapter(l.HEle)
}

// 列表框_排序
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
// bAscending: 升序(TRUE)或降序(FALSE).
func (l *ListBox) Sort(iColumnAdapter int, bAscending bool) int {
	return xc.XListBox_Sort(l.HEle, iColumnAdapter, bAscending)
}

// 列表框_刷新数据
func (l *ListBox) RefreshData() int {
	return xc.XListBox_RefreshData(l.HEle)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI
// iItem: 项索引.
func (l *ListBox) RefreshItem(iItem int) int {
	return xc.XListBox_RefreshItem(l.HEle, iItem)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引
// pText:
func (l *ListBox) AddItemText(pText string) int {
	return xc.XListBox_AddItemText(l.HEle, pText)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx
// pName:
// pText:
func (l *ListBox) AddItemTextEx(pName string, pText string) int {
	return xc.XListBox_AddItemTextEx(l.HEle, pName, pText)
}

// 列表框_添加项图片, XAdTable_AddItemImage
// hImage:
func (l *ListBox) AddItemImage(hImage int) int {
	return xc.XListBox_AddItemImage(l.HEle, hImage)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx
// pName:
// hImage:
func (l *ListBox) AddItemImageEx(pName string, hImage int) int {
	return xc.XListBox_AddItemImageEx(l.HEle, pName, hImage)
}

// 列表框_插入项文本
// iItem:
// pValue:
func (l *ListBox) InsertItemText(iItem int, pValue string) int {
	return xc.XListBox_InsertItemText(l.HEle, iItem, pValue)
}

// 列表框_插入项文本扩展
// iItem:
// pName:
// pValue:
func (l *ListBox) InsertItemTextEx(iItem int, pName string, pValue string) int {
	return xc.XListBox_InsertItemTextEx(l.HEle, iItem, pName, pValue)
}

// 列表框_插入项图片
// iItem:
// hImage:
func (l *ListBox) InsertItemImage(iItem int, hImage int) int {
	return xc.XListBox_InsertItemImage(l.HEle, iItem, hImage)
}

// 列表框_插入项图片扩展
// iItem:
// pName:
// hImage:
func (l *ListBox) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XListBox_InsertItemImageEx(l.HEle, iItem, pName, hImage)
}

// 列表框_置项文本
// iItem:
// iColumn:
// pText:
func (l *ListBox) SetItemText(iItem int, iColumn int, pText string) bool {
	return xc.XListBox_SetItemText(l.HEle, iItem, iColumn, pText)
}

// 列表框_置项文本扩展
// iItem:
// pName:
// pText:
func (l *ListBox) SetItemTextEx(iItem int, pName string, pText string) bool {
	return xc.XListBox_SetItemTextEx(l.HEle, iItem, pName, pText)
}

// 列表框_置项图片
// iItem:
// iColumn:
// hImage:
func (l *ListBox) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XListBox_SetItemImage(l.HEle, iItem, iColumn, hImage)
}

// 列表框_置项图片扩展
// iItem:
// pName:
// hImage:
func (l *ListBox) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XListBox_SetItemImageEx(l.HEle, iItem, pName, hImage)
}

// 列表框_置项整数值
// iItem:
// iColumn:
// nValue:
func (l *ListBox) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XListBox_SetItemInt(l.HEle, iItem, iColumn, nValue)
}

// 列表框_置项整数值扩展
// iItem:
// pName:
// nValue:
func (l *ListBox) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XListBox_SetItemIntEx(l.HEle, iItem, pName, nValue)
}

// 列表框_置项浮点值
// iItem:
// iColumn:
// fFloat:
func (l *ListBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {
	return xc.XListBox_SetItemFloat(l.HEle, iItem, iColumn, fFloat)
}

// 列表框_置项浮点值扩展
// iItem:
// pName:
// fFloat:
func (l *ListBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {
	return xc.XListBox_SetItemFloatEx(l.HEle, iItem, pName, fFloat)
}

// 列表框_取项文本
// iItem:
// iColumn:
func (l *ListBox) GetItemText(iItem int, iColumn int) string {
	return xc.XListBox_GetItemText(l.HEle, iItem, iColumn)
}

// 列表框_取项文本扩展
// iItem:
// pName:
func (l *ListBox) GetItemTextEx(iItem int, pName string) string {
	return xc.XListBox_GetItemTextEx(l.HEle, iItem, pName)
}

// 列表框_取项图片
// iItem:
// iColumn:
func (l *ListBox) GetItemImage(iItem int, iColumn int) int {
	return xc.XListBox_GetItemImage(l.HEle, iItem, iColumn)
}

// 列表框_取项图片扩展
// iItem:
// pName:
func (l *ListBox) GetItemImageEx(iItem int, pName string) int {
	return xc.XListBox_GetItemImageEx(l.HEle, iItem, pName)
}

// 列表框_取项整数值
// iItem:
// iColumn:
// pOutValue:
func (l *ListBox) GetItemInt(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XListBox_GetItemInt(l.HEle, iItem, iColumn, pOutValue)
}

// 列表框_取项整数值扩展
// iItem:
// pName:
// pOutValue:
func (l *ListBox) GetItemIntEx(iItem int, pName string, pOutValue *int) bool {
	return xc.XListBox_GetItemIntEx(l.HEle, iItem, pName, pOutValue)
}

// 列表框_取项浮点值
// iItem:
// iColumn:
// pOutValue:
func (l *ListBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloat(l.HEle, iItem, iColumn, pOutValue)
}

// 列表框_取项浮点值扩展
// iItem:
// pName:
// pOutValue:
func (l *ListBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloatEx(l.HEle, iItem, pName, pOutValue)
}

// 列表框_删除项
// iItem:
func (l *ListBox) DeleteItem(iItem int) bool {
	return xc.XListBox_DeleteItem(l.HEle, iItem)
}

// 列表框_删除项扩展
// iItem:
// nCount:
func (l *ListBox) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XListBox_DeleteItemEx(l.HEle, iItem, nCount)
}

// 列表框_删除项全部
func (l *ListBox) DeleteItemAll() int {
	return xc.XListBox_DeleteItemAll(l.HEle)
}

// 列表框_删除列全部
func (l *ListBox) DeleteColumnAll() int {
	return xc.XListBox_DeleteColumnAll(l.HEle)
}

// 列表框_取项数量AD
func (l *ListBox) GetCount_AD() int {
	return xc.XListBox_GetCount_AD(l.HEle)
}

// 列表框_取列数量AD
func (l *ListBox) GetCountColumn_AD() int {
	return xc.XListBox_GetCountColumn_AD(l.HEle)
}
