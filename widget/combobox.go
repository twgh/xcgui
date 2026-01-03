package widget

import (
	"github.com/twgh/xcgui/adapter"
	"github.com/twgh/xcgui/tmpl"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 下拉组合框.
type ComboBox struct {
	Edit
}

// 组合框_创建, 失败返回 nil.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口.
func NewComboBox(x, y, cx, cy int32, hParent int) *ComboBox {
	return NewComboBoxByHandle(xc.XComboBox_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewComboBoxByHandle(handle int) *ComboBox {
	if handle <= 0 {
		return nil
	}
	p := &ComboBox{}
	p.SetHandle(handle)
	return p
}

// 从 name 创建对象, 失败返回 nil.
func NewComboBoxByName(name string) *ComboBox {
	return NewComboBoxByHandle(xc.XC_GetObjectByName(name))
}

// 从 UID 创建对象, 失败返回 nil.
func NewComboBoxByUID(nUID int32) *ComboBox {
	return NewComboBoxByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从 UID 名称创建对象, 失败返回 nil.
func NewComboBoxByUIDName(name string) *ComboBox {
	return NewComboBoxByHandle(xc.XC_GetObjectByUIDName(name))
}

// 组合框_置选择项.
//
// iIndex: 项索引.
func (c *ComboBox) SetSelItem(iIndex int32) bool {
	return xc.XComboBox_SetSelItem(c.Handle, iIndex)
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
func (c *ComboBox) CreateAdapter() int {
	return xc.XComboBox_CreateAdapter(c.Handle)
}

// 组合框_绑定数据适配器.
//
// hAdapter: 适配器句柄.
func (c *ComboBox) BindAdapter(hAdapter int) *ComboBox {
	xc.XComboBox_BindAdapter(c.Handle, hAdapter)
	return c
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
func (c *ComboBox) GetAdapter() int {
	return xc.XComboBox_GetAdapter(c.Handle)
}

// 组合框_取数据适配器对象, 获取绑定的数据适配器对象, 失败返回 nil.
func (c *ComboBox) GetAdapterObj() *adapter.AdapterTable {
	return adapter.NewAdapterTableByHandle(xc.XComboBox_GetAdapter(c.Handle))
}

// 组合框_置绑定名称.
//
// name: 字段名.
func (c *ComboBox) SetBindName(name string) *ComboBox {
	xc.XComboBox_SetBindName(c.Handle, name)
	return c
}

// 组合框_取下拉按钮坐标.
//
// pRect: 坐标.
func (c *ComboBox) GetButtonRect(pRect *xc.RECT) *ComboBox {
	xc.XComboBox_GetButtonRect(c.Handle, pRect)
	return c
}

// 组合框_置下拉按钮大小.
//
// size: 大小.
func (c *ComboBox) SetButtonSize(size int32) *ComboBox {
	xc.XComboBox_SetButtonSize(c.Handle, size)
	return c
}

// 组合框_置下拉列表高度.
//
// height: 高度, -1自动计算高度.
func (c *ComboBox) SetDropHeight(height int32) *ComboBox {
	xc.XComboBox_SetDropHeight(c.Handle, height)
	return c
}

// 组合框_取下拉列表高度.
func (c *ComboBox) GetDropHeight() int32 {
	return xc.XComboBox_GetDropHeight(c.Handle)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// pXmlFile: 项模板文件.
func (c *ComboBox) SetItemTemplateXML(pXmlFile string) *ComboBox {
	xc.XComboBox_SetItemTemplateXML(c.Handle, pXmlFile)
	return c
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// pStringXML: 字符串.
func (c *ComboBox) SetItemTemplateXMLFromString(pStringXML string) *ComboBox {
	xc.XComboBox_SetItemTemplateXMLFromString(c.Handle, pStringXML)
	return c
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// bEnable: 是否绘制.
func (c *ComboBox) EnableDrawButton(bEnable bool) *ComboBox {
	xc.XComboBox_EnableDrawButton(c.Handle, bEnable)
	return c
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// bEdit: TRUE可编辑.
func (c *ComboBox) EnableEdit(bEdit bool) *ComboBox {
	xc.XComboBox_EnableEdit(c.Handle, bEdit)
	return c
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// bEnable: 是否启用.
func (c *ComboBox) EnableDropHeightFixed(bEnable bool) *ComboBox {
	xc.XComboBox_EnableDropHeightFixed(c.Handle, bEnable)
	return c
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
func (c *ComboBox) GetSelItem() int32 {
	return xc.XComboBox_GetSelItem(c.Handle)
}

// 组合框_取状态, 返回: ComboBox_State_.
func (c *ComboBox) GetState() xcc.ComboBox_State_ {
	return xc.XComboBox_GetState(c.Handle)
}

// 组合框_添加项文本, 返回项索引.
//
// text:.
func (c *ComboBox) AddItemText(text string) int32 {
	return xc.XComboBox_AddItemText(c.Handle, text)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// name: 字段名.
//
// text: 文本.
func (c *ComboBox) AddItemTextEx(name string, text string) int32 {
	return xc.XComboBox_AddItemTextEx(c.Handle, name, text)
}

// 组合框_添加项图片, 返回项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) AddItemImage(hImage int) int32 {
	return xc.XComboBox_AddItemImage(c.Handle, hImage)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// name: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) AddItemImageEx(name string, hImage int) int32 {
	return xc.XComboBox_AddItemImageEx(c.Handle, name, hImage)
}

// 组合框_插入项文本, 返回项索引.
//
// iItem: 项索引.
//
// text: 文本.
func (c *ComboBox) InsertItemText(iItem int32, text string) int32 {
	return xc.XComboBox_InsertItemText(c.Handle, iItem, text)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// iItem: 项索引.
//
// name: 字段名.
//
// text: 文本.
func (c *ComboBox) InsertItemTextEx(iItem int32, name string, text string) int32 {
	return xc.XComboBox_InsertItemTextEx(c.Handle, iItem, name, text)
}

// 组合框_插入项图片, 返回项索引.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) InsertItemImage(iItem int32, hImage int) int32 {
	return xc.XComboBox_InsertItemImage(c.Handle, iItem, hImage)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// iItem: 项索引.
//
// name: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) InsertItemImageEx(iItem int32, name string, hImage int) int32 {
	return xc.XComboBox_InsertItemImageEx(c.Handle, iItem, name, hImage)
}

// 组合框_置项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// text: 文本.
func (c *ComboBox) SetItemText(iItem, iColumn int32, text string) bool {
	return xc.XComboBox_SetItemText(c.Handle, iItem, iColumn, text)
}

// 组合框_置项文本扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// text: 文本.
func (c *ComboBox) SetItemTextEx(iItem int32, name string, text string) bool {
	return xc.XComboBox_SetItemTextEx(c.Handle, iItem, name, text)
}

// 组合框_置项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (c *ComboBox) SetItemImage(iItem, iColumn int32, hImage int) bool {
	return xc.XComboBox_SetItemImage(c.Handle, iItem, iColumn, hImage)
}

// 组合框_置项图片扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) SetItemImageEx(iItem int32, name string, hImage int) bool {
	return xc.XComboBox_SetItemImageEx(c.Handle, iItem, name, hImage)
}

// 组合框_置项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func (c *ComboBox) SetItemInt(iItem, iColumn int32, nValue int32) bool {
	return xc.XComboBox_SetItemInt(c.Handle, iItem, iColumn, nValue)
}

// 组合框_置项指数值扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// nValue: 整数值.
func (c *ComboBox) SetItemIntEx(iItem int32, name string, nValue int32) bool {
	return xc.XComboBox_SetItemIntEx(c.Handle, iItem, name, nValue)
}

// 组合框_置项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func (c *ComboBox) SetItemFloat(iItem, iColumn int32, fFloat float32) bool {
	return xc.XComboBox_SetItemFloat(c.Handle, iItem, iColumn, fFloat)
}

// 组合框_置项浮点值扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// fFloat: 浮点数.
func (c *ComboBox) SetItemFloatEx(iItem int32, name string, fFloat float32) bool {
	return xc.XComboBox_SetItemFloatEx(c.Handle, iItem, name, fFloat)
}

// 组合框_取项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) GetItemText(iItem int32, iColumn int32) string {
	return xc.XComboBox_GetItemText(c.Handle, iItem, iColumn)
}

// 组合框_取项文本扩展.
//
// iItem: 项索引.
//
// name: 字段名.
func (c *ComboBox) GetItemTextEx(iItem int32, name string) string {
	return xc.XComboBox_GetItemTextEx(c.Handle, iItem, name)
}

// 组合框_取项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) GetItemImage(iItem, iColumn int32) int {
	return xc.XComboBox_GetItemImage(c.Handle, iItem, iColumn)
}

// 组合框_取项图片扩展.
//
// iItem: 项索引.
//
// name: 字段名.
func (c *ComboBox) GetItemImageEx(iItem int32, name string) int {
	return xc.XComboBox_GetItemImageEx(c.Handle, iItem, name)
}

// 组合框_取项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) GetItemInt(iItem, iColumn int32, pOutValue *int32) bool {
	return xc.XComboBox_GetItemInt(c.Handle, iItem, iColumn, pOutValue)
}

// 组合框_取项整数值扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) GetItemIntEx(iItem int32, name string, pOutValue *int32) bool {
	return xc.XComboBox_GetItemIntEx(c.Handle, iItem, name, pOutValue)
}

// 组合框_取项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) GetItemFloat(iItem, iColumn int32, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloat(c.Handle, iItem, iColumn, pOutValue)
}

// 组合框_取项浮点值扩展.
//
// iItem: 项索引.
//
// name: 字段名.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) GetItemFloatEx(iItem int32, name string, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloatEx(c.Handle, iItem, name, pOutValue)
}

// 组合框_删除项.
//
// iItem: 项索引.
func (c *ComboBox) DeleteItem(iItem int32) bool {
	return xc.XComboBox_DeleteItem(c.Handle, iItem)
}

// 组合框_删除项扩展.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func (c *ComboBox) DeleteItemEx(iItem, nCount int32) bool {
	return xc.XComboBox_DeleteItemEx(c.Handle, iItem, nCount)
}

// 组合框_删除项全部.
func (c *ComboBox) DeleteItemAll() *ComboBox {
	xc.XComboBox_DeleteItemAll(c.Handle)
	return c
}

// 组合框_删除列全部.
func (c *ComboBox) DeleteColumnAll() *ComboBox {
	xc.XComboBox_DeleteColumnAll(c.Handle)
	return c
}

// 组合框_取项数量.
func (c *ComboBox) GetCount() int32 {
	return xc.XComboBox_GetCount(c.Handle)
}

// 组合框_取列数量.
func (c *ComboBox) GetCountColumn() int32 {
	return xc.XComboBox_GetCountColumn(c.Handle)
}

// 组合框_弹出下拉列表.
func (c *ComboBox) PopupDropList() *ComboBox {
	xc.XComboBox_PopupDropList(c.Handle)
	return c
}

// 组合框_设置项模板.
//
// hTemp: 模板句柄.
func (c *ComboBox) SetItemTemplate(hTemp int) bool {
	return xc.XComboBox_SetItemTemplate(c.Handle, hTemp)
}

// 组合框_置项模板从内存.
//
// data: 模板数据.
func (c *ComboBox) SetItemTemplateXMLFromMem(data []byte) bool {
	return xc.XComboBox_SetItemTemplateXMLFromMem(c.Handle, data)
}

// 组合框_置项模板从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (c *ComboBox) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XComboBox_SetItemTemplateXMLFromZipRes(c.Handle, id, pFileName, pPassword, hModule)
}

// 组合框_取项模板, 返回项模板句柄.
func (c *ComboBox) GetItemTemplate() int {
	return xc.XComboBox_GetItemTemplate(c.Handle)
}

// 组合框_取项模板, 返回项模板对象, 失败返回 nil.
func (c *ComboBox) GetItemTemplateObj() *tmpl.ListItemTemplate {
	return tmpl.NewByHandle(xc.XComboBox_GetItemTemplate(c.Handle))
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_ComboBox_Select_End 添加组合框下拉列表项选择完成事件, 编辑框内容已经改变.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (c *ComboBox) AddEvent_ComboBox_Select_End(pFun xc.XE_COMBOBOX_SELECT_END1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(c.Handle, xcc.XE_COMBOBOX_SELECT_END, onXE_COMBOBOX_SELECT_END, pFun, allowAddingMultiple...)
}

// onXE_COMBOBOX_SELECT_END 组合框下拉列表项选择完成事件, 编辑框内容已经改变.
func onXE_COMBOBOX_SELECT_END(hEle int, iItem int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_COMBOBOX_SELECT_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_COMBOBOX_SELECT_END1); ok {
			ret = cb(hEle, iItem, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Combobox_Select 添加组合框下拉列表项选择事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (c *ComboBox) AddEvent_Combobox_Select(pFun xc.XE_COMBOBOX_SELECT1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(c.Handle, xcc.XE_COMBOBOX_SELECT, onXE_COMBOBOX_SELECT, pFun, allowAddingMultiple...)
}

// onXE_COMBOBOX_SELECT 组合框下拉列表项选择事件.
func onXE_COMBOBOX_SELECT(hEle int, iItem int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_COMBOBOX_SELECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_COMBOBOX_SELECT1); ok {
			ret = cb(hEle, iItem, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Combobox_Popup_List 添加组合框下拉列表弹出事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (c *ComboBox) AddEvent_Combobox_Popup_List(pFun xc.XE_COMBOBOX_POPUP_LIST1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(c.Handle, xcc.XE_COMBOBOX_POPUP_LIST, onXE_COMBOBOX_POPUP_LIST, pFun, allowAddingMultiple...)
}

// onXE_COMBOBOX_POPUP_LIST 组合框下拉列表弹出事件.
func onXE_COMBOBOX_POPUP_LIST(hEle int, hWindow int, hListBox int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_COMBOBOX_POPUP_LIST)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_COMBOBOX_POPUP_LIST1); ok {
			ret = cb(hEle, hWindow, hListBox, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Combobox_Exit_List 添加组合框下拉列表退出事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (c *ComboBox) AddEvent_Combobox_Exit_List(pFun xc.XE_COMBOBOX_EXIT_LIST1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(c.Handle, xcc.XE_COMBOBOX_EXIT_LIST, onXE_COMBOBOX_EXIT_LIST, pFun, allowAddingMultiple...)
}

// onXE_COMBOBOX_EXIT_LIST 组合框下拉列表退出事件.
func onXE_COMBOBOX_EXIT_LIST(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_COMBOBOX_EXIT_LIST)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_COMBOBOX_EXIT_LIST1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) Event_ComboBox_Select_End(pFun xc.XE_COMBOBOX_SELECT_END) bool {
	return xc.XEle_RegEventC(c.Handle, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) Event_ComboBox_Select_End1(pFun xc.XE_COMBOBOX_SELECT_END1) bool {
	return xc.XEle_RegEventC1(c.Handle, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) Event_COMBOBOX_SELECT(pFun xc.XE_COMBOBOX_SELECT) bool {
	return xc.XEle_RegEventC(c.Handle, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) Event_COMBOBOX_SELECT1(pFun xc.XE_COMBOBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(c.Handle, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) Event_COMBOBOX_POPUP_LIST(pFun xc.XE_COMBOBOX_POPUP_LIST) bool {
	return xc.XEle_RegEventC(c.Handle, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) Event_COMBOBOX_POPUP_LIST1(pFun xc.XE_COMBOBOX_POPUP_LIST1) bool {
	return xc.XEle_RegEventC1(c.Handle, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) Event_COMBOBOX_EXIT_LIST(pFun xc.XE_COMBOBOX_EXIT_LIST) bool {
	return xc.XEle_RegEventC(c.Handle, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) Event_COMBOBOX_EXIT_LIST1(pFun xc.XE_COMBOBOX_EXIT_LIST1) bool {
	return xc.XEle_RegEventC1(c.Handle, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}
