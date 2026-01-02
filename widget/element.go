package widget

import (
	"github.com/twgh/xcgui/bkmanager"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Element 基础元素.
type Element struct {
	objectbase.Widget
}

// 元素_创建, 创建基础元素, 失败返回 nil.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewElement(x, y, cx, cy int32, hParent int) *Element {
	return NewElementByHandle(xc.XEle_Create(x, y, cx, cy, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewElementByHandle(handle int) *Element {
	if handle <= 0 {
		return nil
	}
	p := &Element{}
	p.SetHandle(handle)
	return p
}

// 从 name 创建对象, 失败返回 nil.
func NewElementByName(name string) *Element {
	return NewElementByHandle(xc.XC_GetObjectByName(name))
}

// 从 UID 创建对象, 失败返回 nil.
func NewElementByUID(nUID int32) *Element {
	return NewElementByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从 UID 名称创建对象, 失败返回 nil.
func NewElementByUIDName(name string) *Element {
	return NewElementByHandle(xc.XC_GetObjectByUIDName(name))
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// fun: 事件函数.
func (e *Element) RegEventC(nEvent xcc.XE_, fun interface{}) bool {
	return xc.XEle_RegEventC(e.Handle, nEvent, fun)
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// fun: 事件函数.
func (e *Element) RegEventC1(nEvent xcc.XE_, fun interface{}) bool {
	return xc.XEle_RegEventC1(e.Handle, nEvent, fun)
}

// 元素_移除事件C.
//
// nEvent: 事件类型: xcc.XE_.
//
// fun: 事件函数.
func (e *Element) RemoveEventC(nEvent xcc.XE_, fun interface{}) bool {
	return xc.XEle_RemoveEventC(e.Handle, nEvent, fun)
}

// 元素_移除事件. 只适用于 AddEvent_ 方式添加的事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// index: 使用 AddEvent_ 函数返回的回调函数索引.
//   - 为空时, 直接移除事件.
//   - 不为空时, 移除指定索引的回调函数.
func (e *Element) RemoveEvent(nEvent xcc.XE_, index ...int) *Element {
	if len(index) > 0 {
		xc.EleEventHandler.RemoveCallBack(e.Handle, nEvent, index[0])
	} else {
		cbPtr := xc.EleEventHandler.EventInfoMap[e.Handle][nEvent].EvnetFuncPtr
		if cbPtr > 0 {
			xc.XEle_RegEventCEx(e.Handle, nEvent, cbPtr)
		}
		xc.EleEventHandler.RemoveEvent(e.Handle, nEvent)
	}
	return e
}

// 元素_注册事件CEx, 注册事件C方式, 省略2参数, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func (e *Element) RegEventCEx(nEvent xcc.XE_, pFun uintptr) bool {
	return xc.XEle_RegEventCEx(e.Handle, nEvent, pFun)
}

// 元素_注册事件C1Ex, 注册事件C1方式, 省略1参数, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func (e *Element) RegEventC1Ex(nEvent xcc.XE_, pFun uintptr) bool {
	return xc.XEle_RegEventC1Ex(e.Handle, nEvent, pFun)
}

// 元素_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func (e *Element) RemoveEventCEx(nEvent xcc.XE_, pFun uintptr) bool {
	return xc.XEle_RemoveEventCEx(e.Handle, nEvent, pFun)
}

// 元素_发送事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func (e *Element) SendEvent(nEvent xcc.XE_, wParam, lParam uintptr) int32 {
	return xc.XEle_SendEvent(e.Handle, nEvent, wParam, lParam)
}

// 元素_投递事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func (e *Element) PostEvent(nEvent xcc.XE_, wParam, lParam uintptr) bool {
	return xc.XEle_PostEvent(e.Handle, nEvent, wParam, lParam)
}

// 元素_取坐标.
//
// pRect: 坐标.
func (e *Element) GetRect(pRect *xc.RECT) *Element {
	xc.XEle_GetRect(e.Handle, pRect)
	return e
}

// 元素_取坐标ex.
func (e *Element) GetRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetRect(e.Handle, &rc)
	return rc
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
func (e *Element) GetRectLogic(pRect *xc.RECT) *Element {
	xc.XEle_GetRectLogic(e.Handle, pRect)
	return e
}

// 元素_取逻辑坐标ex, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
func (e *Element) GetRectLogicEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetRectLogic(e.Handle, &rc)
	return rc
}

// 元素_取客户区坐标.
//
// pRect: 坐标.
func (e *Element) GetClientRect(pRect *xc.RECT) *Element {
	xc.XEle_GetClientRect(e.Handle, pRect)
	return e
}

// 元素_取客户区坐标ex.
func (e *Element) GetClientRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetClientRect(e.Handle, &rc)
	return rc
}

// 元素_置宽度.
//
// nWidth: 宽度.
func (e *Element) SetWidth(nWidth int32) *Element {
	xc.XEle_SetWidth(e.Handle, nWidth)
	return e
}

// 元素_置高度.
//
// nHeight: 高度.
func (e *Element) SetHeight(nHeight int32) *Element {
	xc.XEle_SetHeight(e.Handle, nHeight)
	return e
}

// 元素_取宽度.
func (e *Element) GetWidth() int32 {
	return xc.XEle_GetWidth(e.Handle)
}

// 元素_取高度.
func (e *Element) GetHeight() int32 {
	return xc.XEle_GetHeight(e.Handle)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectWndClientToEleClient(pRect *xc.RECT) *Element {
	xc.XEle_RectWndClientToEleClient(e.Handle, pRect)
	return e
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointWndClientToEleClient(pPt *xc.POINT) *Element {
	xc.XEle_PointWndClientToEleClient(e.Handle, pPt)
	return e
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectClientToWndClient(pRect *xc.RECT) *Element {
	xc.XEle_RectClientToWndClient(e.Handle, pRect)
	return e
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointClientToWndClient(pPt *xc.POINT) *Element {
	xc.XEle_PointClientToWndClient(e.Handle, pPt)
	return e
}

// 元素_基于窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) GetWndClientRect(pRect *xc.RECT) *Element {
	xc.XEle_GetWndClientRect(e.Handle, pRect)
	return e
}

// 元素_基于窗口客户区坐标ex.
func (e *Element) GetWndClientRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetWndClientRect(e.Handle, &rc)
	return rc
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.
func (e *Element) GetCursor() uintptr {
	return xc.XEle_GetCursor(e.Handle)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hCursor: 光标句柄, 使用系统预定义的, 或者从文件加载.
//   - hCur := wapi.LoadImageW(0, wapi.IDC_HAND, wapi.IMAGE_CURSOR, 0, 0, wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
//   - hCur := wapi.LoadImageW(0, common.StrPtr("arrow.cur"), wapi.IMAGE_CURSOR, 0, 0, wapi.LR_LOADFROMFILE)
func (e *Element) SetCursor(hCursor uintptr) *Element {
	xc.XEle_SetCursor(e.Handle, hCursor)
	return e
}

// 元素_添加子对象.
//
// hChild: 要添加的子元素句柄或形状对象句柄.
func (e *Element) AddChild(hChild int) bool {
	return xc.XEle_AddChild(e.Handle, hChild)
}

// 元素_插入子对象, 插入子对象到指定位置.
//
// hChild: 要插入的元素句柄或形状对象句柄.
//
// index: 插入位置索引.
func (e *Element) InsertChild(hChild int, index int32) bool {
	return xc.XEle_InsertChild(e.Handle, hChild, index)
}

// 元素_置坐标, 如果返回0坐标没有改变, 如果大小改变返回2(触发XE_SIZE), 否则返回1(仅改变left,top,没有改变大小).
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetRect(e.Handle, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_置坐标扩展, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetRectEx(x, y, cx, cy int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetRectEx(e.Handle, x, y, cx, cy, bRedraw, nFlags, nAdjustNo)
}

// 元素_置逻辑坐标, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_. 此参数将会传入XE_SIZE ,XE_ADJUSTLAYOUT 事件回调.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetRectLogic(e.Handle, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// x: X坐标.
//
// y: Y坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetPosition(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetPosition(e.Handle, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动逻辑坐标, 移动元素坐标, 逻辑坐标, 包含滚动视图偏移. 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// x: X坐标.
//
// y: Y坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetPositionLogic(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetPositionLogic(e.Handle, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_判断绘制焦点.
func (e *Element) IsDrawFocus() bool {
	return xc.XEle_IsDrawFocus(e.Handle)
}

// 元素_判断启用, 元素是否为启用状态.
func (e *Element) IsEnable() bool {
	return xc.XEle_IsEnable(e.Handle)
}

// 元素_判断启用焦点, 元素是否启用焦点.
func (e *Element) IsEnableFocus() bool {
	return xc.XEle_IsEnableFocus(e.Handle)
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透.
func (e *Element) IsMouseThrough() bool {
	return xc.XEle_IsMouseThrough(e.Handle)
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素. 成功时返回元素句柄.
//
// pPt: 坐标点.
func (e *Element) HitChildEle(pPt *xc.POINT) int {
	return xc.XEle_HitChildEle(e.Handle, pPt)
}

// 元素_判断背景透明.
func (e *Element) IsBkTransparent() bool {
	return xc.XEle_IsBkTransparent(e.Handle)
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件.
func (e *Element) IsEnableEvent_XE_PAINT_END() bool {
	return xc.XEle_IsEnableEvent_XE_PAINT_END(e.Handle)
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit.
func (e *Element) IsKeyTab() bool {
	return xc.XEle_IsKeyTab(e.Handle)
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键).
func (e *Element) IsSwitchFocus() bool {
	return xc.XEle_IsSwitchFocus(e.Handle)
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素.
func (e *Element) IsEnable_XE_MOUSEWHEEL() bool {
	return xc.XEle_IsEnable_XE_MOUSEWHEEL(e.Handle)
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素.
//
// hChildEle: 子元素句柄.
func (e *Element) IsChildEle(hChildEle int) bool {
	return xc.XEle_IsChildEle(e.Handle, hChildEle)
}

// 元素_判断启用画布, 判断是否启用画布.
func (e *Element) IsEnableCanvas() bool {
	return xc.XEle_IsEnableCanvas(e.Handle)
}

// 元素_判断焦点, 判断是否拥有焦点.
func (e *Element) IsFocus() bool {
	return xc.XEle_IsFocus(e.Handle)
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点.
func (e *Element) IsFocusEx() bool {
	return xc.XEle_IsFocusEx(e.Handle)
}

// 元素_启用, 启用或禁用元素.
//
// bEnable: 启用或禁用.
func (e *Element) Enable(bEnable bool) *Element {
	xc.XEle_Enable(e.Handle, bEnable)
	return e
}

// 元素_启用焦点, 启用焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableFocus(bEnable bool) *Element {
	xc.XEle_EnableFocus(e.Handle, bEnable)
	return e
}

// 元素_启用绘制焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableDrawFocus(bEnable bool) *Element {
	xc.XEle_EnableDrawFocus(e.Handle, bEnable)
	return e
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// bEnable: 是否启用.
func (e *Element) EnableDrawBorder(bEnable bool) *Element {
	xc.XEle_EnableDrawBorder(e.Handle, bEnable)
	return e
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// bEnable: 是否启用.
func (e *Element) EnableCanvas(bEnable bool) *Element {
	xc.XEle_EnableCanvas(e.Handle, bEnable)
	return e
}

// 元素_启用事件_XE_PAINT_END.
//
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_PAINT_END(bEnable bool) *Element {
	xc.XEle_EnableEvent_XE_PAINT_END(e.Handle, bEnable)
	return e
}

// 元素_启用背景透明.
//
// bEnable: 是否启用.
func (e *Element) EnableBkTransparent(bEnable bool) *Element {
	xc.XEle_EnableBkTransparent(e.Handle, bEnable)
	return e
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// bEnable: 是否启用.
func (e *Element) EnableMouseThrough(bEnable bool) *Element {
	xc.XEle_EnableMouseThrough(e.Handle, bEnable)
	return e
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// bEnable: 是否启用.
func (e *Element) EnableKeyTab(bEnable bool) *Element {
	xc.XEle_EnableKeyTab(e.Handle, bEnable)
	return e
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableSwitchFocus(bEnable bool) *Element {
	xc.XEle_EnableSwitchFocus(e.Handle, bEnable)
	return e
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_MOUSEWHEEL(bEnable bool) *Element {
	xc.XEle_EnableEvent_XE_MOUSEWHEEL(e.Handle, bEnable)
	return e
}

// 元素_移除, 移除元素, 但不销毁.
func (e *Element) Remove() *Element {
	xc.XEle_Remove(e.Handle)
	return e
}

// 元素_置Z序, 设置元素Z序.
//
// index: 位置索引.
func (e *Element) SetZOrder(index int32) bool {
	return xc.XEle_SetZOrder(e.Handle, index)
}

// 元素_置Z序扩展, 设置元素Z序.
//
// hDestEle: 目标元素.
//
// nType: 类型, Zorder_.
func (e *Element) SetZOrderEx(hDestEle int, nType xcc.Zorder_) bool {
	return xc.XEle_SetZOrderEx(e.Handle, hDestEle, nType)
}

// 元素_取Z序, 获取元素Z序索引, 位置索引.
func (e *Element) GetZOrder() int32 {
	return xc.XEle_GetZOrder(e.Handle)
}

// 元素_启用置顶, 设置元素置顶.
//
// bTopmost: 是否置顶显示.
func (e *Element) EnableTopmost(bTopmost bool) bool {
	return xc.XEle_EnableTopmost(e.Handle, bTopmost)
}

// 元素_重绘.
//
// bImmediate: 是否立即重绘, 通常为false即可.
func (e *Element) Redraw(bImmediate ...bool) *Element {
	b := false
	if len(bImmediate) > 0 {
		b = bImmediate[0]
	}
	xc.XEle_Redraw(e.Handle, b)
	return e
}

// 元素_重绘指定区域.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.
func (e *Element) RedrawRect(pRect *xc.RECT, bImmediate bool) *Element {
	xc.XEle_RedrawRect(e.Handle, pRect, bImmediate)
	return e
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.
func (e *Element) GetChildCount() int32 {
	return xc.XEle_GetChildCount(e.Handle)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// index: 索引.
func (e *Element) GetChildByIndex(index int32) int {
	return xc.XEle_GetChildByIndex(e.Handle, index)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// nID: 元素ID.
func (e *Element) GetChildByID(nID int32) int {
	return xc.XEle_GetChildByID(e.Handle, nID)
}

// 元素_置边框大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (e *Element) SetBorderSize(left, top, right, bottom int32) *Element {
	xc.XEle_SetBorderSize(e.Handle, left, top, right, bottom)
	return e
}

// 元素_取边框大小.
//
// pBorder: 大小.
func (e *Element) GetBorderSize(pBorder *xc.RECT) *Element {
	xc.XEle_GetBorderSize(e.Handle, pBorder)
	return e
}

// 元素_取边框大小ex.
func (e *Element) GetBorderSizeEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetBorderSize(e.Handle, &rc)
	return rc
}

// 元素_置内填充大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (e *Element) SetPadding(left, top, right, bottom int32) *Element {
	xc.XEle_SetPadding(e.Handle, left, top, right, bottom)
	return e
}

// 元素_取内填充大小.
//
// pPadding: 大小.
func (e *Element) GetPadding(pPadding *xc.RECT) *Element {
	xc.XEle_GetPadding(e.Handle, pPadding)
	return e
}

// 元素_取内填充大小ex.
func (e *Element) GetPaddingEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetPadding(e.Handle, &rc)
	return rc
}

// 元素_置拖动边框.
//
// nFlags: 边框位置组合, xcc.Element_Position_.
func (e *Element) SetDragBorder(nFlags xcc.Element_Position_) *Element {
	xc.XEle_SetDragBorder(e.Handle, nFlags)
	return e
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// nFlags: 边框位置标识, xcc.Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.
func (e *Element) SetDragBorderBindEle(nFlags xcc.Element_Position_, hBindEle int, nSpace int32) *Element {
	xc.XEle_SetDragBorderBindEle(e.Handle, nFlags, hBindEle, nSpace)
	return e
}

// 元素_置最小大小.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.
func (e *Element) SetMinSize(nWidth, nHeight int32) *Element {
	xc.XEle_SetMinSize(e.Handle, nWidth, nHeight)
	return e
}

// 元素_置最大大小.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.
func (e *Element) SetMaxSize(nWidth, nHeight int32) *Element {
	xc.XEle_SetMaxSize(e.Handle, nWidth, nHeight)
	return e
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.
func (e *Element) SetLockScroll(bHorizon bool, bVertical bool) *Element {
	xc.XEle_SetLockScroll(e.Handle, bHorizon, bVertical)
	return e
}

// 元素_置文本颜色.
//
// color: xc.RGBA 颜色值.
func (e *Element) SetTextColor(color uint32) *Element {
	xc.XEle_SetTextColor(e.Handle, color)
	return e
}

// 元素_取文本颜色.
func (e *Element) GetTextColor() uint32 {
	return xc.XEle_GetTextColor(e.Handle)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.
func (e *Element) GetTextColorEx() uint32 {
	return xc.XEle_GetTextColorEx(e.Handle)
}

// 元素_置焦点边框颜色.
//
// color: xc.RGBA 颜色值.
func (e *Element) SetFocusBorderColor(color uint32) *Element {
	xc.XEle_SetFocusBorderColor(e.Handle, color)
	return e
}

// 元素_取焦点边框颜色.
func (e *Element) GetFocusBorderColor() uint32 {
	return xc.XEle_GetFocusBorderColor(e.Handle)
}

// 元素_置字体.
//
// hFontx: 炫彩字体.
func (e *Element) SetFont(hFontx int) *Element {
	xc.XEle_SetFont(e.Handle, hFontx)
	return e
}

// 元素_取字体.
func (e *Element) GetFont() int {
	return xc.XEle_GetFont(e.Handle)
}

// 元素_取字体, 返回炫彩字体对象, 失败返回 nil.
func (e *Element) GetFontObj() *font.Font {
	return font.NewByHandle(xc.XEle_GetFont(e.Handle))
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.
func (e *Element) GetFontEx() int {
	return xc.XEle_GetFontEx(e.Handle)
}

// 元素_取字体扩展, 获取元素字体对象, 优先从资源中获取, 失败返回 nil.
func (e *Element) GetFontObjEx() *font.Font {
	return font.NewByHandle(xc.XEle_GetFontEx(e.Handle))
}

// 元素_置透明度.
func (e *Element) SetAlpha(alpha byte) *Element {
	xc.XEle_SetAlpha(e.Handle, alpha)
	return e
}

// 元素_销毁.
func (e *Element) Destroy() *Element {
	xc.XEle_Destroy(e.Handle)
	return e
}

// 元素_添加背景边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: xc.RGBA 颜色.
//
// width: 线宽.
func (e *Element) AddBkBorder(nState xcc.CombinedState, color uint32, width int32) *Element {
	xc.XEle_AddBkBorder(e.Handle, nState, color, width)
	return e
}

// 元素_添加背景填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: xc.RGBA 颜色.
func (e *Element) AddBkFill(nState xcc.CombinedState, color uint32) *Element {
	xc.XEle_AddBkFill(e.Handle, nState, color)
	return e
}

// 元素_添加背景图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (e *Element) AddBkImage(nState xcc.CombinedState, hImage int) *Element {
	xc.XEle_AddBkImage(e.Handle, nState, hImage)
	return e
}

// 元素_取背景对象数量, 获取背景内容数量.
func (e *Element) GetBkInfoCount() int32 {
	return xc.XEle_GetBkInfoCount(e.Handle)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
func (e *Element) ClearBkInfo() *Element {
	xc.XEle_ClearBkInfo(e.Handle)
	return e
}

// 元素_取背景管理器, 获取元素背景管理器.
func (e *Element) GetBkManager() int {
	return xc.XEle_GetBkManager(e.Handle)
}

// 元素_取背景管理器, 获取元素背景管理器对象, 失败返回 nil.
func (e *Element) GetBkManagerObj() *bkmanager.BkManager {
	return bkmanager.NewByHandle(xc.XEle_GetBkManager(e.Handle))
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.
func (e *Element) GetBkManagerEx() int {
	return xc.XEle_GetBkManagerEx(e.Handle)
}

// 元素_取背景管理器扩展, 获取元素背景管理器对象, 优先从资源中获取, 失败返回 nil.
func (e *Element) GetBkManagerObjEx() *bkmanager.BkManager {
	return bkmanager.NewByHandle(xc.XEle_GetBkManagerEx(e.Handle))
}

// 元素_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (e *Element) SetBkManager(hBkInfoM int) *Element {
	xc.XEle_SetBkManager(e.Handle, hBkInfoM)
	return e
}

// 元素_取状态, 获取组合状态.
func (e *Element) GetStateFlags() xcc.CombinedState {
	return xc.XEle_GetStateFlags(e.Handle)
}

// 元素_绘制焦点, 绘制元素焦点.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func (e *Element) DrawFocus(hDraw int, pRect *xc.RECT) bool {
	return xc.XEle_DrawFocus(e.Handle, hDraw, pRect)
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (e *Element) DrawEle(hDraw int) *Element {
	xc.XEle_DrawEle(e.Handle, hDraw)
	return e
}

// 元素_置用户数据.
//
// nData: 用户数据.
func (e *Element) SetUserData(nData int) *Element {
	xc.XEle_SetUserData(e.Handle, nData)
	return e
}

// 元素_取用户数据.
func (e *Element) GetUserData() int {
	return xc.XEle_GetUserData(e.Handle)
}

// 元素_取内容大小.
//
// bHorizon: 水平或垂直, 布局属性交换依赖.
//
// cx: 宽度.
//
// cy: 高度.
//
// pSize: 返回大小.
func (e *Element) GetContentSize(bHorizon bool, cx, cy int32, pSize *xc.SIZE) *Element {
	xc.XEle_GetContentSize(e.Handle, bHorizon, cx, cy, pSize)
	return e
}

// 元素_置鼠标捕获.
//
// b: TRUE设置.
func (e *Element) SetCapture(b bool) *Element {
	xc.XEle_SetCapture(e.Handle, b)
	return e
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// bEnable: 启用或关闭.
func (e *Element) EnableTransparentChannel(bEnable bool) *Element {
	xc.XEle_EnableTransparentChannel(e.Handle, bEnable)
	return e
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.
func (e *Element) SetXCTimer(nIDEvent, uElapse uint32) bool {
	return xc.XEle_SetXCTimer(e.Handle, nIDEvent, uElapse)
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// nIDEvent: 事件ID.
func (e *Element) KillXCTimer(nIDEvent uint32) bool {
	return xc.XEle_KillXCTimer(e.Handle, nIDEvent)
}

// 元素_置工具提示, 设置工具提示内容.
//
// pText: 工具提示内容.
func (e *Element) SetToolTip(pText string) *Element {
	xc.XEle_SetToolTip(e.Handle, pText)
	return e
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (e *Element) SetToolTipEx(pText string, nTextAlign xcc.TextFormatFlag_) *Element {
	xc.XEle_SetToolTipEx(e.Handle, pText, nTextAlign)
	return e
}

// 元素_取工具提示, 获取工具提示内容.
func (e *Element) GetToolTip() string {
	return xc.XEle_GetToolTip(e.Handle)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// x: X坐标.
//
// y: Y坐标.
func (e *Element) PopupToolTip(x, y int32) *Element {
	xc.XEle_PopupToolTip(e.Handle, x, y)
	return e
}

// 元素_调整布局.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) AdjustLayout(nAdjustNo uint32) *Element {
	xc.XEle_AdjustLayout(e.Handle, nAdjustNo)
	return e
}

// 元素_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) AdjustLayoutEx(nFlags xcc.AdjustLayout_, nAdjustNo uint32) *Element {
	xc.XEle_AdjustLayoutEx(e.Handle, nFlags, nAdjustNo)
	return e
}

// 元素_取透明度, 返回透明度.
func (e *Element) GetAlpha() byte {
	return xc.XEle_GetAlpha(e.Handle)
}

// 元素_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func (e *Element) GetPosition(pOutX, pOutY *int32) *Element {
	xc.XEle_GetPosition(e.Handle, pOutX, pOutY)
	return e
}

// 元素_取位置ex.
func (e *Element) GetPositionEx() xc.POINT {
	var pOutX, pOutY int32
	xc.XEle_GetPosition(e.Handle, &pOutX, &pOutY)
	return xc.POINT{X: pOutX, Y: pOutY}
}

// 元素_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) SetSize(nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int32 {
	return xc.XEle_SetSize(e.Handle, nWidth, nHeight, bRedraw, nFlags, nAdjustNo)
}

// 元素_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func (e *Element) GetSize(pOutWidth, pOutHeight *int32) *Element {
	xc.XEle_GetSize(e.Handle, pOutWidth, pOutHeight)
	return e
}

// 元素_取大小ex.
func (e *Element) GetSizeEx() xc.POINT {
	var pOutWidth, pOutHeight int32
	xc.XEle_GetSize(e.Handle, &pOutWidth, &pOutHeight)
	return xc.POINT{X: pOutWidth, Y: pOutHeight}
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (e *Element) SetBkInfo(pText string) int32 {
	return xc.XEle_SetBkInfo(e.Handle, pText)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.
func (e *Element) GetWndClientRectDPI(pRect *xc.RECT) *Element {
	xc.XEle_GetWndClientRectDPI(e.Handle, pRect)
	return e
}

// 元素_取窗口客户区坐标DPIex. 基于DPI缩放后的坐标.
func (e *Element) GetWndClientRectDPIEx() xc.RECT {
	rc := xc.RECT{}
	xc.XEle_GetWndClientRectDPI(e.Handle, &rc)
	return rc
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pPt: 接收返回坐标点.
func (e *Element) PointClientToWndClientDPI(pPt *xc.POINT) *Element {
	xc.XEle_PointClientToWndClientDPI(e.Handle, pPt)
	return e
}

// 元素_取窗口客户区坐标DPIex. 基于DPI缩放后的坐标.
//
// pPt: 接收返回坐标点.
func (e *Element) PointClientToWndClientDPIEx() xc.POINT {
	pt := xc.POINT{}
	xc.XEle_PointClientToWndClientDPI(e.Handle, &pt)
	return pt
}

// 元素_客户区坐标到窗口客户区DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.
func (e *Element) RectClientToWndClientDPI(pRect *xc.RECT) *Element {
	xc.XEle_RectClientToWndClientDPI(e.Handle, pRect)
	return e
}

// SetFocus 元素_置焦点.
func (e *Element) SetFocus() bool {
	hParent := 0
	hEle := e.Handle
	for {
		hParent = xc.XWidget_GetParent(hEle)
		if xc.XC_IsHWINDOW(hParent) {
			break
		}

		if hParent == 0 {
			return false
		}

		hEle = hParent
	}

	xc.XWnd_SetFocusEle(hParent, e.Handle)
	return true
}

// GetLeft 元素_取左边.
func (e *Element) GetLeft() int32 {
	var rc xc.RECT
	xc.XEle_GetRect(e.Handle, &rc)
	return rc.Left
}

// GetTop 元素_取顶边.
func (e *Element) GetTop() int32 {
	var rc xc.RECT
	xc.XEle_GetRect(e.Handle, &rc)
	return rc.Top
}

// GetRight 元素_取右边.
func (e *Element) GetRight() int32 {
	var rc xc.RECT
	xc.XEle_GetRect(e.Handle, &rc)
	return rc.Right
}

// GetBottom 元素_取底边.
func (e *Element) GetBottom() int32 {
	var rc xc.RECT
	xc.XEle_GetRect(e.Handle, &rc)
	return rc.Bottom
}

// SetLeft 元素_置左边.
//
// x: 左边x坐标.
//
// bRedraw: 是否重绘.
func (e *Element) SetLeft(x int32, bRedraw bool) bool {
	return xc.XEle_SetPosition(e.Handle, x, e.GetTop(), bRedraw, xcc.AdjustLayout_All, 0) != 0
}

// SetLeft 元素_置顶边.
//
// y: 顶边y坐标.
//
// bRedraw: 是否重绘.
func (e *Element) SetTop(y int32, bRedraw bool) bool {
	return xc.XEle_SetPosition(e.Handle, e.GetLeft(), y, bRedraw, xcc.AdjustLayout_All, 0) != 0
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_Destroy_End 添加元素销毁完成事件. 在销毁子对象之后触发.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Destroy_End(pFun xc.XE_DESTROY_END1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_DESTROY_END, xc.OnXE_DESTROY_END, pFun, allowAddingMultiple...)
}

// AddEvent_EleProce 添加元素处理过程事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_EleProce(pFun xc.XE_ELEPROCE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_ELEPROCE, onXE_ELEPROCE, pFun, allowAddingMultiple...)
}

// onXE_ELEPROCE 元素处理过程事件.
func onXE_ELEPROCE(hEle int, nEvent uint32, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_ELEPROCE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_ELEPROCE1); ok {
			ret = cb(hEle, nEvent, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint 添加元素绘制事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Paint(pFun xc.XE_PAINT1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_PAINT, onXE_PAINT, pFun, allowAddingMultiple...)
}

// onXE_PAINT 元素绘制事件.
func onXE_PAINT(hEle int, hDraw int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_PAINT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_PAINT1); ok {
			ret = cb(hEle, hDraw, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint_End 添加该元素及子元素绘制完成事件.启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Paint_End(pFun xc.XE_PAINT_END1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_PAINT_END, onXE_PAINT_END, pFun, allowAddingMultiple...)
}

// onXE_PAINT_END 该元素及子元素绘制完成事件.
func onXE_PAINT_END(hEle int, hDraw int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_PAINT_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_PAINT_END1); ok {
			ret = cb(hEle, hDraw, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint_Scrollview 添加滚动视图绘制事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Paint_Scrollview(pFun xc.XE_PAINT_SCROLLVIEW1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_PAINT_SCROLLVIEW, onXE_PAINT_SCROLLVIEW, pFun, allowAddingMultiple...)
}

// onXE_PAINT_SCROLLVIEW 滚动视图绘制事件.
func onXE_PAINT_SCROLLVIEW(hEle int, hDraw int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_PAINT_SCROLLVIEW)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_PAINT_SCROLLVIEW1); ok {
			ret = cb(hEle, hDraw, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseMove 添加元素鼠标移动事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_MouseMove(pFun xc.XE_MOUSEMOVE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MOUSEMOVE, onXE_MOUSEMOVE, pFun, allowAddingMultiple...)
}

// onXE_MOUSEMOVE 元素鼠标移动事件.
func onXE_MOUSEMOVE(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MOUSEMOVE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_MOUSEMOVE1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseStay 添加元素鼠标进入事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_MouseStay(pFun xc.XE_MOUSESTAY1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MOUSESTAY, onXE_MOUSESTAY, pFun, allowAddingMultiple...)
}

// onXE_MOUSESTAY 元素鼠标进入事件.
func onXE_MOUSESTAY(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MOUSESTAY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_MOUSESTAY1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseHover 添加元素鼠标悬停事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_MouseHover(pFun xc.XE_MOUSEHOVER1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MOUSEHOVER, onXE_MOUSEHOVER, pFun, allowAddingMultiple...)
}

// onXE_MOUSEHOVER 元素鼠标悬停事件.
func onXE_MOUSEHOVER(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MOUSEHOVER)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_MOUSEHOVER1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseLeave 添加元素鼠标离开事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_MouseLeave(pFun xc.XE_MOUSELEAVE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MOUSELEAVE, onXE_MOUSELEAVE, pFun, allowAddingMultiple...)
}

// onXE_MOUSELEAVE 元素鼠标离开事件.
func onXE_MOUSELEAVE(hEle int, hEleStay int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MOUSELEAVE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_MOUSELEAVE1); ok {
			ret = cb(hEle, hEleStay, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseWheel 添加元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL. flags: 见MSDN中 WM_MOUSEWHEEL 消息 wParam 参数说明.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_MouseWheel(pFun xc.XE_MOUSEWHEEL1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MOUSEWHEEL, onXE_MOUSEWHEEL, pFun, allowAddingMultiple...)
}

// onXE_MOUSEWHEEL 元素鼠标滚轮滚动事件.
func onXE_MOUSEWHEEL(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MOUSEWHEEL)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_MOUSEWHEEL1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonDown 添加鼠标左键按下事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_LButtonDown(pFun xc.XE_LBUTTONDOWN1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_LBUTTONDOWN, onXE_LBUTTONDOWN, pFun, allowAddingMultiple...)
}

// onXE_LBUTTONDOWN 鼠标左键按下事件.
func onXE_LBUTTONDOWN(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LBUTTONDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LBUTTONDOWN1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonUp 添加鼠标左键弹起事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_LButtonUp(pFun xc.XE_LBUTTONUP1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_LBUTTONUP, onXE_LBUTTONUP, pFun, allowAddingMultiple...)
}

// onXE_LBUTTONUP 鼠标左键弹起事件.
func onXE_LBUTTONUP(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LBUTTONUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LBUTTONUP1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_RButtonDown 添加鼠标右键按下事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_RButtonDown(pFun xc.XE_RBUTTONDOWN1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_RBUTTONDOWN, onXE_RBUTTONDOWN, pFun, allowAddingMultiple...)
}

// onXE_RBUTTONDOWN 鼠标右键按下事件.
func onXE_RBUTTONDOWN(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_RBUTTONDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_RBUTTONDOWN1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_RButtonUp 添加鼠标右键弹起事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_RButtonUp(pFun xc.XE_RBUTTONUP1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_RBUTTONUP, onXE_RBUTTONUP, pFun, allowAddingMultiple...)
}

// onXE_RBUTTONUP 鼠标右键弹起事件.
func onXE_RBUTTONUP(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_RBUTTONUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_RBUTTONUP1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonDBClick 添加鼠标左键双击事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_LButtonDBClick(pFun xc.XE_LBUTTONDBCLICK1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_LBUTTONDBCLICK, onXE_LBUTTONDBCLICK, pFun, allowAddingMultiple...)
}

// onXE_LBUTTONDBCLICK 鼠标左键双击事件.
func onXE_LBUTTONDBCLICK(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_LBUTTONDBCLICK)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_LBUTTONDBCLICK1); ok {
			ret = cb(hEle, nFlags, pPt, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_XC_Timer 添加炫彩定时器事件. 非系统定时器, 定时器消息 XM_TIMER.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_XC_Timer(pFun xc.XE_XC_TIMER1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_XC_TIMER, onXE_XC_TIMER, pFun, allowAddingMultiple...)
}

// onXE_XC_TIMER 炫彩定时器事件.
func onXE_XC_TIMER(hEle int, nTimerID int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_XC_TIMER)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_XC_TIMER1); ok {
			ret = cb(hEle, nTimerID, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_AdjustLayout 添加调整布局事件. 暂停使用.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_AdjustLayout(pFun xc.XE_ADJUSTLAYOUT1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_ADJUSTLAYOUT, onXE_ADJUSTLAYOUT, pFun, allowAddingMultiple...)
}

// onXE_ADJUSTLAYOUT 调整布局事件.
func onXE_ADJUSTLAYOUT(hEle int, nFlags int32, nAdjustNo uint32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_ADJUSTLAYOUT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_ADJUSTLAYOUT1); ok {
			ret = cb(hEle, nFlags, nAdjustNo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ToolTip_Popup 添加工具提示弹出事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_ToolTip_Popup(pFun xc.XE_TOOLTIP_POPUP1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_TOOLTIP_POPUP, onXE_TOOLTIP_POPUP, pFun, allowAddingMultiple...)
}

// onXE_TOOLTIP_POPUP 工具提示弹出事件.
func onXE_TOOLTIP_POPUP(hEle int, hWindow int, pText uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_TOOLTIP_POPUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_TOOLTIP_POPUP1); ok {
			ret = cb(hEle, hWindow, pText, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_AdjustLayout_End 添加调整布局完成事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_AdjustLayout_End(pFun xc.XE_ADJUSTLAYOUT_END1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_ADJUSTLAYOUT_END, onXE_ADJUSTLAYOUT_END, pFun, allowAddingMultiple...)
}

// onXE_ADJUSTLAYOUT_END 调整布局完成事件.
func onXE_ADJUSTLAYOUT_END(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_ADJUSTLAYOUT_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_ADJUSTLAYOUT_END1); ok {
			ret = cb(hEle, nFlags, nAdjustNo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetFocus 添加元素获得焦点事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SetFocus(pFun xc.XE_SETFOCUS1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SETFOCUS, onXE_SETFOCUS, pFun, allowAddingMultiple...)
}

// onXE_SETFOCUS 元素获得焦点事件.
func onXE_SETFOCUS(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SETFOCUS)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SETFOCUS1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KillFocus 添加元素失去焦点事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_KillFocus(pFun xc.XE_KILLFOCUS1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_KILLFOCUS, onXE_KILLFOCUS, pFun, allowAddingMultiple...)
}

// onXE_KILLFOCUS 元素失去焦点事件.
func onXE_KILLFOCUS(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_KILLFOCUS)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_KILLFOCUS1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Destroy 添加元素即将销毁事件. 在销毁子对象之前触发.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Destroy(pFun xc.XE_DESTROY1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_DESTROY, onXE_DESTROY, pFun, allowAddingMultiple...)
}

// onXE_DESTROY 元素即将销毁事件.
func onXE_DESTROY(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_DESTROY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_DESTROY1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Size 添加元素大小改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Size(pFun xc.XE_SIZE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SIZE, onXE_SIZE, pFun, allowAddingMultiple...)
}

// onXE_SIZE 元素大小改变事件.
func onXE_SIZE(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SIZE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SIZE1); ok {
			ret = cb(hEle, nFlags, nAdjustNo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Show 添加元素显示隐藏事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Show(pFun xc.XE_SHOW1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SHOW, onXE_SHOW, pFun, allowAddingMultiple...)
}

// onXE_SHOW 元素显示隐藏事件.
func onXE_SHOW(hEle int, bShow bool, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SHOW)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SHOW1); ok {
			ret = cb(hEle, bShow, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetFont 添加元素设置字体事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SetFont(pFun xc.XE_SETFONT1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SETFONT, onXE_SETFONT, pFun, allowAddingMultiple...)
}

// onXE_SETFONT 元素设置字体事件.
func onXE_SETFONT(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SETFONT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SETFONT1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KeyDown 添加元素按键按下事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_KeyDown(pFun xc.XE_KEYDOWN1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_KEYDOWN, onXE_KEYDOWN, pFun, allowAddingMultiple...)
}

// onXE_KEYDOWN 元素按键按下事件.
func onXE_KEYDOWN(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_KEYDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_KEYDOWN1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KeyUp 添加元素按键弹起事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_KeyUp(pFun xc.XE_KEYUP1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_KEYUP, onXE_KEYUP, pFun, allowAddingMultiple...)
}

// onXE_KEYUP 元素按键弹起事件.
func onXE_KEYUP(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_KEYUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_KEYUP1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SysKeyDown 添加元素系统按键按下事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SysKeyDown(pFun xc.XE_SYSKEYDOWN1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SYSKEYDOWN, onXE_SYSKEYDOWN, pFun, allowAddingMultiple...)
}

// onXE_SYSKEYDOWN 元素系统按键按下事件.
func onXE_SYSKEYDOWN(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SYSKEYDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SYSKEYDOWN1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SysKeyUp 添加元素系统按键弹起事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SysKeyUp(pFun xc.XE_SYSKEYUP1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SYSKEYUP, onXE_SYSKEYUP, pFun, allowAddingMultiple...)
}

// onXE_SYSKEYUP 元素系统按键弹起事件.
func onXE_SYSKEYUP(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SYSKEYUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SYSKEYUP1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Char 添加字符输入事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Char(pFun xc.XE_CHAR1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_CHAR, onXE_CHAR, pFun, allowAddingMultiple...)
}

// onXE_CHAR 字符输入事件.
func onXE_CHAR(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_CHAR)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_CHAR1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetCapture 添加元素设置鼠标捕获事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SetCapture(pFun xc.XE_SETCAPTURE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SETCAPTURE, onXE_SETCAPTURE, pFun, allowAddingMultiple...)
}

// onXE_SETCAPTURE 元素设置鼠标捕获事件.
func onXE_SETCAPTURE(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SETCAPTURE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SETCAPTURE1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KillCapture 添加元素失去鼠标捕获事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_KillCapture(pFun xc.XE_KILLCAPTURE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_KILLCAPTURE, onXE_KILLCAPTURE, pFun, allowAddingMultiple...)
}

// onXE_KILLCAPTURE 元素失去鼠标捕获事件.
func onXE_KILLCAPTURE(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_KILLCAPTURE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_KILLCAPTURE1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetCursor 添加元素设置鼠标光标事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_SetCursor(pFun xc.XE_SETCURSOR1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_SETCURSOR, onXE_SETCURSOR, pFun, allowAddingMultiple...)
}

// onXE_SETCURSOR 元素设置鼠标光标事件.
func onXE_SETCURSOR(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SETCURSOR)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SETCURSOR1); ok {
			ret = cb(hEle, wParam, lParam, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_DropFiles 添加文件拖放事件. 需先启用: xc.XWnd_EnableDragFiles.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_DropFiles(pFun xc.XE_DROPFILES1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_DROPFILES, onXE_DROPFILES, pFun, allowAddingMultiple...)
}

// onXE_DROPFILES 文件拖放事件.
func onXE_DROPFILES(hEle int, hDropInfo uintptr, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_DROPFILES)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_DROPFILES1); ok {
			ret = cb(hEle, hDropInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_Select 添加弹出菜单项选择事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_Select(pFun func(hEle int, nID int32, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_SELECT, onXE_MENU_SELECT, pFun, allowAddingMultiple...)
}

// onXE_MENU_SELECT 弹出菜单项选择事件.
func onXE_MENU_SELECT(hEle int, nID int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_SELECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, nID int32, pbHandled *bool) int); ok {
			ret = cb(hEle, nID, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_Popup 添加菜单弹出事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_Popup(pFun func(hEle int, HMENUX int, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_POPUP, onXE_MENU_POPUP, pFun, allowAddingMultiple...)
}

// onXE_MENU_POPUP 菜单弹出事件.
func onXE_MENU_POPUP(hEle int, HMENUX int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_POPUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, HMENUX int, pbHandled *bool) int); ok {
			ret = cb(hEle, HMENUX, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_Exit 添加菜单退出事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_Exit(pFun func(hEle int, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_EXIT, onXE_MENU_EXIT, pFun, allowAddingMultiple...)
}

// onXE_MENU_EXIT 菜单退出事件.
func onXE_MENU_EXIT(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_EXIT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, pbHandled *bool) int); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_Popup_Wnd 添加菜单弹出窗口事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_Popup_Wnd(pFun func(hEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_POPUP_WND, onXE_MENU_POPUP_WND, pFun, allowAddingMultiple...)
}

// onXE_MENU_POPUP_WND 菜单弹出窗口事件.
func onXE_MENU_POPUP_WND(hEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_POPUP_WND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int); ok {
			ret = cb(hEle, hMenu, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_Draw_Background 添加菜单绘制背景事件. 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_Draw_Background(pFun func(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_DRAW_BACKGROUND, onXE_MENU_DRAW_BACKGROUND, pFun, allowAddingMultiple...)
}

// onXE_MENU_DRAW_BACKGROUND 菜单绘制背景事件.
func onXE_MENU_DRAW_BACKGROUND(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_DRAW_BACKGROUND)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int); ok {
			ret = cb(hEle, hDraw, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Menu_DrawItem 添加菜单项绘制事件. 启用该功能需要调用 xc.XMenu_EnableDrawItem.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Element) AddEvent_Menu_DrawItem(pFun func(hEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_MENU_DRAWITEM, onXE_MENU_DRAWITEM, pFun, allowAddingMultiple...)
}

// onXE_MENU_DRAWITEM 菜单项绘制事件.
func onXE_MENU_DRAWITEM(hEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_MENU_DRAWITEM)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(func(hEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int); ok {
			ret = cb(hEle, hDraw, pInfo, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 元素处理过程事件.
func (e *Element) Event_ELEPROCE(pFun xc.XE_ELEPROCE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ELEPROCE, pFun)
}

// 元素处理过程事件.
func (e *Element) Event_ELEPROCE1(pFun xc.XE_ELEPROCE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ELEPROCE, pFun)
}

// 元素绘制事件.
func (e *Element) Event_PAINT(pFun xc.XE_PAINT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT, pFun)
}

// 元素绘制事件.
func (e *Element) Event_PAINT1(pFun xc.XE_PAINT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
func (e *Element) Event_PAINT_END(pFun xc.XE_PAINT_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT_END, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用 xc.XEle_EnableEvent_XE_PAINT_END.
func (e *Element) Event_PAINT_END1(pFun xc.XE_PAINT_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT_END, pFun)
}

// 滚动视图绘制事件.
func (e *Element) Event_PAINT_SCROLLVIEW(pFun xc.XE_PAINT_SCROLLVIEW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// 滚动视图绘制事件.
func (e *Element) Event_PAINT_SCROLLVIEW1(pFun xc.XE_PAINT_SCROLLVIEW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// 元素鼠标移动事件.
func (e *Element) Event_MOUSEMOVE(pFun xc.XE_MOUSEMOVE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEMOVE, pFun)
}

// 元素鼠标移动事件.
func (e *Element) Event_MOUSEMOVE1(pFun xc.XE_MOUSEMOVE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEMOVE, pFun)
}

// 元素鼠标进入事件.
func (e *Element) Event_MOUSESTAY(pFun xc.XE_MOUSESTAY) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSESTAY, pFun)
}

// 元素鼠标进入事件.
func (e *Element) Event_MOUSESTAY1(pFun xc.XE_MOUSESTAY1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSESTAY, pFun)
}

// 元素鼠标悬停事件.
func (e *Element) Event_MOUSEHOVER(pFun xc.XE_MOUSEHOVER) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEHOVER, pFun)
}

// 元素鼠标悬停事件.
func (e *Element) Event_MOUSEHOVER1(pFun xc.XE_MOUSEHOVER1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEHOVER, pFun)
}

// 元素鼠标离开事件.
func (e *Element) Event_MOUSELEAVE(pFun xc.XE_MOUSELEAVE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSELEAVE, pFun)
}

// 元素鼠标离开事件.
func (e *Element) Event_MOUSELEAVE1(pFun xc.XE_MOUSELEAVE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSELEAVE, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL.
func (e *Element) Event_MOUSEWHEEL(pFun xc.XE_MOUSEWHEEL) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEWHEEL, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 xc.XEle_EnableEvent_XE_MOUSEWHEEL.
func (e *Element) Event_MOUSEWHEEL1(pFun xc.XE_MOUSEWHEEL1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEWHEEL, pFun)
}

// 鼠标左键按下事件.
func (e *Element) Event_LBUTTONDOWN(pFun xc.XE_LBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键按下事件.
func (e *Element) Event_LBUTTONDOWN1(pFun xc.XE_LBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键弹起事件.
func (e *Element) Event_LBUTTONUP(pFun xc.XE_LBUTTONUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONUP, pFun)
}

// 鼠标左键弹起事件.
func (e *Element) Event_LBUTTONUP1(pFun xc.XE_LBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONUP, pFun)
}

// 鼠标右键按下事件.
func (e *Element) Event_RBUTTONDOWN(pFun xc.XE_RBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键按下事件.
func (e *Element) Event_RBUTTONDOWN1(pFun xc.XE_RBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键弹起事件.
func (e *Element) Event_RBUTTONUP(pFun xc.XE_RBUTTONUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_RBUTTONUP, pFun)
}

// 鼠标右键弹起事件.
func (e *Element) Event_RBUTTONUP1(pFun xc.XE_RBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_RBUTTONUP, pFun)
}

// 鼠标左键双击事件.
func (e *Element) Event_LBUTTONDBCLICK(pFun xc.XE_LBUTTONDBCLICK) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONDBCLICK, pFun)
}

// 鼠标左键双击事件.
func (e *Element) Event_LBUTTONDBCLICK1(pFun xc.XE_LBUTTONDBCLICK1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONDBCLICK, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) Event_XC_TIMER(pFun xc.XE_XC_TIMER) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_XC_TIMER, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) Event_XC_TIMER1(pFun xc.XE_XC_TIMER1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_XC_TIMER, pFun)
}

// 调整布局事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT(pFun xc.XE_ADJUSTLAYOUT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT1(pFun xc.XE_ADJUSTLAYOUT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局完成事件.
func (e *Element) Event_ADJUSTLAYOUT_END(pFun xc.XE_ADJUSTLAYOUT_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// 调整布局完成事件.
func (e *Element) Event_ADJUSTLAYOUT_END1(pFun xc.XE_ADJUSTLAYOUT_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// 元素获得焦点事件.
func (e *Element) Event_SETFOCUS(pFun xc.XE_SETFOCUS) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETFOCUS, pFun)
}

// 元素获得焦点事件.
func (e *Element) Event_SETFOCUS1(pFun xc.XE_SETFOCUS1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETFOCUS, pFun)
}

// 元素失去焦点事件.
func (e *Element) Event_KILLFOCUS(pFun xc.XE_KILLFOCUS) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KILLFOCUS, pFun)
}

// 元素失去焦点事件.
func (e *Element) Event_KILLFOCUS1(pFun xc.XE_KILLFOCUS1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KILLFOCUS, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.
func (e *Element) Event_DESTROY(pFun xc.XE_DESTROY) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DESTROY, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.
func (e *Element) Event_DESTROY1(pFun xc.XE_DESTROY1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DESTROY, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.
func (e *Element) Event_DESTROY_END(pFun xc.XE_DESTROY_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DESTROY_END, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.
func (e *Element) Event_DESTROY_END1(pFun xc.XE_DESTROY_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DESTROY_END, pFun)
}

// 元素大小改变事件.
func (e *Element) Event_SIZE(pFun xc.XE_SIZE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SIZE, pFun)
}

// 元素大小改变事件.
func (e *Element) Event_SIZE1(pFun xc.XE_SIZE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SIZE, pFun)
}

// 元素显示隐藏事件.
func (e *Element) Event_SHOW(pFun xc.XE_SHOW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SHOW, pFun)
}

// 元素显示隐藏事件.
func (e *Element) Event_SHOW1(pFun xc.XE_SHOW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SHOW, pFun)
}

// 元素设置字体事件.
func (e *Element) Event_SETFONT(pFun xc.XE_SETFONT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETFONT, pFun)
}

// 元素设置字体事件.
func (e *Element) Event_SETFONT1(pFun xc.XE_SETFONT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETFONT, pFun)
}

// 元素按键按下事件.
func (e *Element) Event_KEYDOWN(pFun xc.XE_KEYDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KEYDOWN, pFun)
}

// 元素按键按下事件.
func (e *Element) Event_KEYDOWN1(pFun xc.XE_KEYDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KEYDOWN, pFun)
}

// 元素按键弹起事件.
func (e *Element) Event_KEYUP(pFun xc.XE_KEYUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KEYUP, pFun)
}

// 元素按键弹起事件.
func (e *Element) Event_KEYUP1(pFun xc.XE_KEYUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KEYUP, pFun)
}

// 元素系统按键按下事件.
func (e *Element) Event_SYSKEYDOWN1(pFun xc.XE_SYSKEYDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SYSKEYDOWN, pFun)
}

// 元素系统按键弹起事件.
func (e *Element) Event_SYSKEYUP1(pFun xc.XE_SYSKEYUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SYSKEYUP, pFun)
}

// 通过 TranslateMessage 函数翻译的字符事件.
func (e *Element) Event_CHAR(pFun xc.XE_CHAR) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_CHAR, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.
func (e *Element) Event_CHAR1(pFun xc.XE_CHAR1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_CHAR, pFun)
}

// 元素设置鼠标捕获.
func (e *Element) Event_SETCAPTURE(pFun xc.XE_SETCAPTURE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETCAPTURE, pFun)
}

// 元素设置鼠标捕获.
func (e *Element) Event_SETCAPTURE1(pFun xc.XE_SETCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETCAPTURE, pFun)
}

// 元素失去鼠标捕获.
func (e *Element) Event_KILLCAPTURE(pFun xc.XE_KILLCAPTURE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KILLCAPTURE, pFun)
}

// 元素失去鼠标捕获.
func (e *Element) Event_KILLCAPTURE1(pFun xc.XE_KILLCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KILLCAPTURE, pFun)
}

// 设置鼠标光标.
func (e *Element) Event_SETCURSOR(pFun xc.XE_SETCURSOR) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETCURSOR, pFun)
}

// 设置鼠标光标.
func (e *Element) Event_SETCURSOR1(pFun xc.XE_SETCURSOR1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETCURSOR, pFun)
}

// 文件拖放事件, 需先启用: xc.XWnd_EnableDragFiles.
func (e *Element) Event_DROPFILES(pFun xc.XE_DROPFILES) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DROPFILES, pFun)
}

// 文件拖放事件, 需先启用: xc.XWnd_EnableDragFiles.
func (e *Element) Event_DROPFILES1(pFun xc.XE_DROPFILES1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DROPFILES, pFun)
}

// 事件_元素工具提示弹出, 可使用 common.UintPtrToString 把uintptr转换到文本.
func (e *Element) Event_TOOLTIP_POPUP(pFun xc.XE_TOOLTIP_POPUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_TOOLTIP_POPUP, pFun)
}

// 事件_元素工具提示弹出1, 可使用 common.UintPtrToString 把uintptr转换到文本.
func (e *Element) Event_TOOLTIP_POPUP1(pFun xc.XE_TOOLTIP_POPUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_TOOLTIP_POPUP, pFun)
}

// 事件_弹出菜单项被选择.
func (e *Element) Event_MENU_SELECT(pFun xc.XE_MENU_SELECT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.
func (e *Element) Event_MENU_POPUP(pFun xc.XE_MENU_POPUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.
func (e *Element) Event_MENU_EXIT(pFun xc.XE_MENU_EXIT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.
func (e *Element) Event_MENU_POPUP_WND(pFun xc.XE_MENU_POPUP_WND) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
func (e *Element) Event_MENU_DRAW_BACKGROUND(pFun xc.XE_MENU_DRAW_BACKGROUND) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.
func (e *Element) Event_MENU_DRAWITEM(pFun xc.XE_MENU_DRAWITEM) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_DRAWITEM, pFun)
}

// 事件_弹出菜单项被选择.
func (e *Element) Event_MENU_SELECT1(pFun xc.XE_MENU_SELECT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.
func (e *Element) Event_MENU_POPUP1(pFun xc.XE_MENU_POPUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.
func (e *Element) Event_MENU_EXIT1(pFun xc.XE_MENU_EXIT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.
func (e *Element) Event_MENU_POPUP_WND1(pFun xc.XE_MENU_POPUP_WND1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用 xc.XMenu_EnableDrawBackground.
func (e *Element) Event_MENU_DRAW_BACKGROUND1(pFun xc.XE_MENU_DRAW_BACKGROUND1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用 xc.XMenu_EnableDrawItem.
func (e *Element) Event_MENU_DRAWITEM1(pFun xc.XE_MENU_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_DRAWITEM, pFun)
}
