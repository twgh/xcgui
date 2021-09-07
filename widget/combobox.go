package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

type ComboBox struct {
	Element

	HEle int
}

// 组合框_创建, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func NewComboBox(x int, y int, cx int, cy int, hParent int) *ComboBox {
	p := &ComboBox{
		HEle: xc.XComboBox_Create(x, y, cx, cy, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 组合框_置选择项
// iIndex: 项索引.
func (c *ComboBox) SetSelItem(iIndex int) bool {
	return xc.XComboBox_SetSelItem(c.HEle, iIndex)
}

// 组合框_创建数据适配器, 返回数据适配器句柄
func (c *ComboBox) CreateAdapter() int {
	return xc.XComboBox_CreateAdapter(c.HEle)
}

// 组合框_绑定数据适配器
// hAdapter: 适配器句柄.
func (c *ComboBox) BindAdapter(hAdapter int) int {
	return xc.XComboBox_BindAdapter(c.HEle, hAdapter)
}

// 组合框_取数据适配器, 获取绑定的数据适配器
func (c *ComboBox) GetAdapter() int {
	return xc.XComboBox_GetAdapter(c.HEle)
}

// 组合框_置绑定名称
// pName: 字段名
func (c *ComboBox) SetBindName(pName string) int {
	return xc.XComboBox_SetBindName(c.HEle, pName)
}

// 组合框_取下拉按钮坐标
// pRect: 坐标.
func (c *ComboBox) GetButtonRect(pRect *xc.RECT) int {
	return xc.XComboBox_GetButtonRect(c.HEle, pRect)
}

// 组合框_置下拉按钮大小
// size: 大小.
func (c *ComboBox) SetButtonSize(size int) int {
	return xc.XComboBox_SetButtonSize(c.HEle, size)
}

// 组合框_置下拉列表高度
// height: 高度, -1自动计算高度
func (c *ComboBox) SetDropHeight(height int) int {
	return xc.XComboBox_SetDropHeight(c.HEle, height)
}

// 组合框_取下拉列表高度
func (c *ComboBox) GetDropHeight() int {
	return xc.XComboBox_GetDropHeight(c.HEle)
}

// 组合框_置项模板, 设置下拉列表项模板文件
// pXmlFile: 项模板文件.
func (c *ComboBox) SetItemTemplateXML(pXmlFile string) int {
	return xc.XComboBox_SetItemTemplateXML(c.HEle, pXmlFile)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板
// pStringXML: 字符串指针.
func (c *ComboBox) SetItemTemplateXMLFromString(pStringXML string) int {
	return xc.XComboBox_SetItemTemplateXMLFromString(c.HEle, pStringXML)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮
// bEnable: 是否绘制.
func (c *ComboBox) EnableDrawButton(bEnable bool) int {
	return xc.XComboBox_EnableDrawButton(c.HEle, bEnable)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容
// bEdit: TRUE可编辑
func (c *ComboBox) EnableEdit(bEdit bool) int {
	return xc.XComboBox_EnableEdit(c.HEle, bEdit)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小
// bEnable: 是否启用.
func (c *ComboBox) EnableDropHeightFixed(bEnable bool) int {
	return xc.XComboBox_EnableDropHeightFixed(c.HEle, bEnable)
}

// 组合框_添加背景边框, 添加背景内容边框
// nState: 按钮状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (c *ComboBox) AddBkBorder(nState int, color int, alpha uint8, width int) int {
	return xc.XComboBox_AddBkBorder(c.HEle, nState, color, alpha, width)
}

// 组合框_添加背景填充, 添加背景内容填充
// nState: 按钮状态.
// color: RGB颜色.
// alpha: 透明度.
func (c *ComboBox) AddBkFill(nState int, color int, alpha uint8) int {
	return xc.XComboBox_AddBkFill(c.HEle, nState, color, alpha)
}

// 组合框_添加背景图片, 添加背景内容图片
// nState: 按钮状态.
// hImage: 图片句柄.
func (c *ComboBox) AddBkImage(nState int, hImage int) int {
	return xc.XComboBox_AddBkImage(c.HEle, nState, hImage)
}

// 组合框_取背景对象数量, 成功返回背景内容数量, 否则返回: XC_ID_ERROR
func (c *ComboBox) GetBkInfoCount() int {
	return xc.XComboboX_GetBkInfoCount(c.HEle)
}

// 组合框_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确
func (c *ComboBox) ClearBkInfo() int {
	return xc.XComboBox_ClearBkInfo(c.HEle)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引
func (c *ComboBox) GetSelItem() int {
	return xc.XComboBox_GetSelItem(c.HEle)
}

// 组合框_取状态, 返回: ComboBox_State_
func (c *ComboBox) GetState() int {
	return xc.XComboBox_GetState(c.HEle)
}

// 组合框_添加项文本, 返回项索引
// pText:
func (c *ComboBox) AddItemText(pText string) int {
	return xc.XComboBox_AddItemText(c.HEle, pText)
}

// 组合框_添加项文本扩展, 返回项索引
// pName: 字段名
// pText: 文本
func (c *ComboBox) AddItemTextEx(pName string, pText string) int {
	return xc.XComboBox_AddItemTextEx(c.HEle, pName, pText)
}

// 组合框_添加项图片, 返回项索引
// hImage: 图片句柄
func (c *ComboBox) AddItemImage(hImage int) int {
	return xc.XComboBox_AddItemImage(c.HEle, hImage)
}

// 组合框_添加项图片扩展, 返回项索引
// pName: 字段名
// hImage: 图片句柄
func (c *ComboBox) AddItemImageEx(pName string, hImage int) int {
	return xc.XComboBox_AddItemImageEx(c.HEle, pName, hImage)
}

// 组合框_插入项文本, 返回项索引
// iItem: 项索引
// pText: 文本
func (c *ComboBox) InsertItemText(iItem int, pText string) int {
	return xc.XComboBox_InsertItemText(c.HEle, iItem, pText)
}

// 组合框_插入项文本扩展, 返回项索引
// iItem: 项索引
// pName: 字段名
// pText: 文本
func (c *ComboBox) InsertItemTextEx(iItem int, pName string, pText string) int {
	return xc.XComboBox_InsertItemTextEx(c.HEle, iItem, pName, pText)
}

// 组合框_插入项图片, 返回项索引
// iItem: 项索引
// hImage: 图片句柄
func (c *ComboBox) InsertItemImage(iItem int, hImage int) int {
	return xc.XComboBox_InsertItemImage(c.HEle, iItem, hImage)
}

// 组合框_插入项图片扩展, 返回项索引
// iItem: 项索引
// pName: 字段名
// hImage: 图片句柄
func (c *ComboBox) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XComboBox_InsertItemImageEx(c.HEle, iItem, pName, hImage)
}

// 组合框_置项文本
// iItem: 项索引
// iColumn: 列索引
// pText: 文本
func (c *ComboBox) SetItemText(iItem int, iColumn int, pText string) bool {
	return xc.XComboBox_SetItemText(c.HEle, iItem, iColumn, pText)
}

// 组合框_置项文本扩展
// iItem: 项索引
// pName: 字段名
// pText: 文本
func (c *ComboBox) SetItemTextEx(iItem int, pName string, pText string) bool {
	return xc.XComboBox_SetItemTextEx(c.HEle, iItem, pName, pText)
}

// 组合框_置项图片
// iItem: 项索引
// iColumn: 列索引
// hImage: 图片句柄
func (c *ComboBox) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XComboBox_SetItemImage(c.HEle, iItem, iColumn, hImage)
}

// 组合框_置项图片扩展
// iItem: 项索引
// pName: 字段名
// hImage: 图片句柄
func (c *ComboBox) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XComboBox_SetItemImageEx(c.HEle, iItem, pName, hImage)
}

// 组合框_置项整数值
// iItem: 项索引
// iColumn: 列索引
// nValue: 整数值
func (c *ComboBox) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XComboBox_SetItemInt(c.HEle, iItem, iColumn, nValue)
}

// 组合框_置项指数值扩展
// iItem: 项索引
// pName: 字段名
// nValue: 整数值
func (c *ComboBox) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XComboBox_SetItemIntEx(c.HEle, iItem, pName, nValue)
}

// 组合框_置项浮点值
// iItem: 项索引
// iColumn: 列索引
// fFloat: 浮点数
func (c *ComboBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {
	return xc.XComboBox_SetItemFloat(c.HEle, iItem, iColumn, fFloat)
}

// 组合框_置项浮点值扩展
// iItem: 项索引
// pName: 字段名
// fFloat: 浮点数
func (c *ComboBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {
	return xc.XComboBox_SetItemFloatEx(c.HEle, iItem, pName, fFloat)
}

// 组合框_取项文本
// iItem: 项索引
// iColumn: 列索引
func (c *ComboBox) GetItemText(iItem int, iColumn int) string {
	return xc.XComboBox_GetItemText(c.HEle, iItem, iColumn)
}

// 组合框_取项文本扩展
// iItem: 项索引
// pName: 字段名
func (c *ComboBox) GetItemTextEx(iItem int, pName string) string {
	return xc.XComboBox_GetItemTextEx(c.HEle, iItem, pName)
}

// 组合框_取项图片
// iItem: 项索引
// iColumn: 列索引
func (c *ComboBox) GetItemImage(iItem int, iColumn int) int {
	return xc.XComboBox_GetItemImage(c.HEle, iItem, iColumn)
}

// 组合框_取项图片扩展
// iItem: 项索引
// pName: 字段名
func (c *ComboBox) GetItemImageEx(iItem int, pName string) int {
	return xc.XComboBox_GetItemImageEx(c.HEle, iItem, pName)
}

// 组合框_取项整数值
// iItem: 项索引
// iColumn: 列索引
// pOutValue: 接收返回整数值
func (c *ComboBox) GetItemInt(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XComboBox_GetItemInt(c.HEle, iItem, iColumn, pOutValue)
}

// 组合框_取项整数值扩展
// iItem: 项索引
// pName: 字段名
// pOutValue: 接收返回整数值
func (c *ComboBox) GetItemIntEx(iItem int, pName string, pOutValue *int) bool {
	return xc.XComboBox_GetItemIntEx(c.HEle, iItem, pName, pOutValue)
}

// 组合框_取项浮点值
// iItem: 项索引
// iColumn: 列索引
// pOutValue: 接收返回浮点值
func (c *ComboBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloat(c.HEle, iItem, iColumn, pOutValue)
}

// 组合框_取项浮点值扩展
// iItem: 项索引
// pName: 字段名
// pOutValue: 接收返回浮点值
func (c *ComboBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloatEx(c.HEle, iItem, pName, pOutValue)
}

// 组合框_删除项
// iItem: 项索引
func (c *ComboBox) DeleteItem(iItem int) bool {
	return xc.XComboBox_DeleteItem(c.HEle, iItem)
}

// 组合框_删除项扩展
// iItem: 项索引
// nCount: 删除数量
func (c *ComboBox) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XComboBox_DeleteItemEx(c.HEle, iItem, nCount)
}

// 组合框_删除项全部
func (c *ComboBox) DeleteItemAll() int {
	return xc.XComboBox_DeleteItemAll(c.HEle)
}

// 组合框_删除列全部
func (c *ComboBox) DeleteColumnAll() int {
	return xc.XComboBox_DeleteColumnAll(c.HEle)
}

// 组合框_取项数量
func (c *ComboBox) GetCount() int {
	return xc.XComboBox_GetCount(c.HEle)
}

// 组合框_取列数量
func (c *ComboBox) GetCountColumn() int {
	return xc.XComboBox_GetCountColumn(c.HEle)
}

/*
下面都是事件
*/

// 组合框_下拉列表项选择完成事件, 编辑框内容已经改变
// pFun: 事件函数指针.
func (c *ComboBox) SetComboBoxSelectEnd(pFun func(iItem int, pbHandled *bool) int) bool {
	return xc.XEle_RegEventC(c.HEle, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 组合框_下拉列表项选择完成事件, 编辑框内容已经改变
// pFun: 事件函数指针.
func (c *ComboBox) SetComboBoxSelectEnd1(pFun func(hEle int, iItem int, pbHandled *bool) int) bool {
	return xc.XEle_RegEventC1(c.HEle, xcc.XE_COMBOBOX_SELECT_END, pFun)
}
