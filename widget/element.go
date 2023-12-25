package widget

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Element 基础元素.
type Element struct {
	objectbase.Widget
}

// 元素_创建, 创建基础元素.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewElement(x, y, cx, cy int32, hParent int) *Element {
	p := &Element{}
	p.SetHandle(xc.XEle_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewElementByHandle(handle int) *Element {
	p := &Element{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewElementByName(name string) *Element {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewElementByUID(nUID int) *Element {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewElementByUIDName(name string) *Element {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func (e *Element) RegEventC(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RegEventC(e.Handle, nEvent, pFun)
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func (e *Element) RegEventC1(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RegEventC1(e.Handle, nEvent, pFun)
}

// 元素_移除事件C.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func (e *Element) RemoveEventC(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RemoveEventC(e.Handle, nEvent, pFun)
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
func (e *Element) SendEvent(nEvent xcc.XE_, wParam, lParam uint) int {
	return xc.XEle_SendEvent(e.Handle, nEvent, wParam, lParam)
}

// 元素_投递事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func (e *Element) PostEvent(nEvent xcc.XE_, wParam, lParam uint) int {
	return xc.XEle_PostEvent(e.Handle, nEvent, wParam, lParam)
}

// 元素_取坐标.
//
// pRect: 坐标.
func (e *Element) GetRect(pRect *xc.RECT) int {
	return xc.XEle_GetRect(e.Handle, pRect)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
func (e *Element) GetRectLogic(pRect *xc.RECT) int {
	return xc.XEle_GetRectLogic(e.Handle, pRect)
}

// 元素_取客户区坐标.
//
// pRect: 坐标.
func (e *Element) GetClientRect(pRect *xc.RECT) int {
	return xc.XEle_GetClientRect(e.Handle, pRect)
}

// 元素_置宽度.
//
// nWidth: 宽度.
func (e *Element) SetWidth(nWidth int) int {
	return xc.XEle_SetWidth(e.Handle, nWidth)
}

// 元素_置高度.
//
// nHeight: 高度.
func (e *Element) SetHeight(nHeight int) int {
	return xc.XEle_SetHeight(e.Handle, nHeight)
}

// 元素_取宽度.
func (e *Element) GetWidth() int {
	return xc.XEle_GetWidth(e.Handle)
}

// 元素_取高度.
func (e *Element) GetHeight() int {
	return xc.XEle_GetHeight(e.Handle)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectWndClientToEleClient(pRect *xc.RECT) int {
	return xc.XEle_RectWndClientToEleClient(e.Handle, pRect)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointWndClientToEleClient(pPt *xc.POINT) int {
	return xc.XEle_PointWndClientToEleClient(e.Handle, pPt)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectClientToWndClient(pRect *xc.RECT) int {
	return xc.XEle_RectClientToWndClient(e.Handle, pRect)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointClientToWndClient(pPt *xc.POINT) int {
	return xc.XEle_PointClientToWndClient(e.Handle, pPt)
}

// 元素_基于窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) GetWndClientRect(pRect *xc.RECT) int {
	return xc.XEle_GetWndClientRect(e.Handle, pRect)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.
func (e *Element) GetCursor() int {
	return xc.XEle_GetCursor(e.Handle)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hCursor: 光标句柄.
func (e *Element) SetCursor(hCursor int) int {
	return xc.XEle_SetCursor(e.Handle, hCursor)
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
func (e *Element) InsertChild(hChild int, index int) bool {
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
func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
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
func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
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
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
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
func (e *Element) SetPosition(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
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
func (e *Element) SetPositionLogic(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
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

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素.
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
func (e *Element) Enable(bEnable bool) int {
	return xc.XEle_Enable(e.Handle, bEnable)
}

// 元素_启用焦点, 启用焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableFocus(bEnable bool) int {
	return xc.XEle_EnableFocus(e.Handle, bEnable)
}

// 元素_启用绘制焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableDrawFocus(bEnable bool) int {
	return xc.XEle_EnableDrawFocus(e.Handle, bEnable)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// bEnable: 是否启用.
func (e *Element) EnableDrawBorder(bEnable bool) int {
	return xc.XEle_EnableDrawBorder(e.Handle, bEnable)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// bEnable: 是否启用.
func (e *Element) EnableCanvas(bEnable bool) int {
	return xc.XEle_EnableCanvas(e.Handle, bEnable)
}

// 元素_启用事件_XE_PAINT_END.
//
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_PAINT_END(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_PAINT_END(e.Handle, bEnable)
}

// 元素_启用背景透明.
//
// bEnable: 是否启用.
func (e *Element) EnableBkTransparent(bEnable bool) int {
	return xc.XEle_EnableBkTransparent(e.Handle, bEnable)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// bEnable: 是否启用.
func (e *Element) EnableMouseThrough(bEnable bool) int {
	return xc.XEle_EnableMouseThrough(e.Handle, bEnable)
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// bEnable: 是否启用.
func (e *Element) EnableKeyTab(bEnable bool) int {
	return xc.XEle_EnableKeyTab(e.Handle, bEnable)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// bEnable: 是否启用.
func (e *Element) EnableSwitchFocus(bEnable bool) int {
	return xc.XEle_EnableSwitchFocus(e.Handle, bEnable)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_MOUSEWHEEL(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_MOUSEWHEEL(e.Handle, bEnable)
}

// 元素_移除, 移除元素, 但不销毁.
func (e *Element) Remove() int {
	return xc.XEle_Remove(e.Handle)
}

// 元素_置Z序, 设置元素Z序.
//
// index: 位置索引.
func (e *Element) SetZOrder(index int) bool {
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
func (e *Element) GetZOrder() int {
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
// bImmediate: 是否立即重绘.
func (e *Element) Redraw(bImmediate bool) int {
	return xc.XEle_Redraw(e.Handle, bImmediate)
}

// 元素_重绘指定区域.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.
func (e *Element) RedrawRect(pRect *xc.RECT, bImmediate bool) int {
	return xc.XEle_RedrawRect(e.Handle, pRect, bImmediate)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.
func (e *Element) GetChildCount() int {
	return xc.XEle_GetChildCount(e.Handle)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// index: 索引.
func (e *Element) GetChildByIndex(index int) int {
	return xc.XEle_GetChildByIndex(e.Handle, index)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// nID: 元素ID.
func (e *Element) GetChildByID(nID int) int {
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
func (e *Element) SetBorderSize(left int, top int, right int, bottom int) int {
	return xc.XEle_SetBorderSize(e.Handle, left, top, right, bottom)
}

// 元素_取边框大小.
//
// pBorder: 大小.
func (e *Element) GetBorderSize(pBorder *xc.RECT) int {
	return xc.XEle_GetBorderSize(e.Handle, pBorder)
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
func (e *Element) SetPadding(left int, top int, right int, bottom int) int {
	return xc.XEle_SetPadding(e.Handle, left, top, right, bottom)
}

// 元素_取内填充大小.
//
// pPadding: 大小.
func (e *Element) GetPadding(pPadding *xc.RECT) int {
	return xc.XEle_GetPadding(e.Handle, pPadding)
}

// 元素_置拖动边框.
//
// nFlags: 边框位置组合, Element_Position_.
func (e *Element) SetDragBorder(nFlags xcc.Element_Position_) int {
	return xc.XEle_SetDragBorder(e.Handle, nFlags)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// nFlags: 边框位置标识, Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.
func (e *Element) SetDragBorderBindEle(nFlags xcc.Element_Position_, hBindEle int, nSpace int) int {
	return xc.XEle_SetDragBorderBindEle(e.Handle, nFlags, hBindEle, nSpace)
}

// 元素_置最小大小.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.
func (e *Element) SetMinSize(nWidth int, nHeight int) int {
	return xc.XEle_SetMinSize(e.Handle, nWidth, nHeight)
}

// 元素_置最大大小.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.
func (e *Element) SetMaxSize(nWidth int, nHeight int) int {
	return xc.XEle_SetMaxSize(e.Handle, nWidth, nHeight)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.
func (e *Element) SetLockScroll(bHorizon bool, bVertical bool) int {
	return xc.XEle_SetLockScroll(e.Handle, bHorizon, bVertical)
}

// 元素_置文本颜色.
//
// color: ABGR 颜色值.
func (e *Element) SetTextColor(color int) int {
	return xc.XEle_SetTextColor(e.Handle, color)
}

// 元素_取文本颜色.
func (e *Element) GetTextColor() int {
	return xc.XEle_GetTextColor(e.Handle)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.
func (e *Element) GetTextColorEx() int {
	return xc.XEle_GetTextColorEx(e.Handle)
}

// 元素_置焦点边框颜色.
//
// color: ABGR 颜色值.
func (e *Element) SetFocusBorderColor(color int) int {
	return xc.XEle_SetFocusBorderColor(e.Handle, color)
}

// 元素_取焦点边框颜色.
func (e *Element) GetFocusBorderColor() int {
	return xc.XEle_GetFocusBorderColor(e.Handle)
}

// 元素_置字体.
//
// hFontx: 炫彩字体.
func (e *Element) SetFont(hFontx int) int {
	return xc.XEle_SetFont(e.Handle, hFontx)
}

// 元素_取字体.
func (e *Element) GetFont() int {
	return xc.XEle_GetFont(e.Handle)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.
func (e *Element) GetFontEx() int {
	return xc.XEle_GetFontEx(e.Handle)
}

// 元素_置透明度.
func (e *Element) SetAlpha(alpha uint8) int {
	return xc.XEle_SetAlpha(e.Handle, alpha)
}

// 元素_销毁.
func (e *Element) Destroy() int {
	return xc.XEle_Destroy(e.Handle)
}

// 元素_添加背景边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func (e *Element) AddBkBorder(nState xcc.CombinedState, color int, width int) int {
	return xc.XEle_AddBkBorder(e.Handle, nState, color, width)
}

// 元素_添加背景填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func (e *Element) AddBkFill(nState xcc.CombinedState, color int) int {
	return xc.XEle_AddBkFill(e.Handle, nState, color)
}

// 元素_添加背景图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (e *Element) AddBkImage(nState xcc.CombinedState, hImage int) int {
	return xc.XEle_AddBkImage(e.Handle, nState, hImage)
}

// 元素_取背景对象数量, 获取背景内容数量.
func (e *Element) GetBkInfoCount() int {
	return xc.XEle_GetBkInfoCount(e.Handle)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
func (e *Element) ClearBkInfo() int {
	return xc.XEle_ClearBkInfo(e.Handle)
}

// 元素_取背景管理器, 获取元素背景管理器.
func (e *Element) GetBkManager() int {
	return xc.XEle_GetBkManager(e.Handle)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.
func (e *Element) GetBkManagerEx() int {
	return xc.XEle_GetBkManagerEx(e.Handle)
}

// 元素_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (e *Element) SetBkManager(hBkInfoM int) int {
	return xc.XEle_SetBkManager(e.Handle, hBkInfoM)
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
func (e *Element) DrawEle(hDraw int) int {
	return xc.XEle_DrawEle(e.Handle, hDraw)
}

// 元素_置用户数据.
//
// nData: 用户数据.
func (e *Element) SetUserData(nData int) int {
	return xc.XEle_SetUserData(e.Handle, nData)
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
func (e *Element) GetContentSize(bHorizon bool, cx int, cy int, pSize *xc.SIZE) int {
	return xc.XEle_GetContentSize(e.Handle, bHorizon, cx, cy, pSize)
}

// 元素_置鼠标捕获.
//
// b: TRUE设置.
func (e *Element) SetCapture(b bool) int {
	return xc.XEle_SetCapture(e.Handle, b)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// bEnable: 启用或关闭.
func (e *Element) EnableTransparentChannel(bEnable bool) int {
	return xc.XEle_EnableTransparentChannel(e.Handle, bEnable)
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.
func (e *Element) SetXCTimer(nIDEvent int, uElapse int) bool {
	return xc.XEle_SetXCTimer(e.Handle, nIDEvent, uElapse)
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// nIDEvent: 事件ID.
func (e *Element) KillXCTimer(nIDEvent int) bool {
	return xc.XEle_KillXCTimer(e.Handle, nIDEvent)
}

// 元素_置工具提示, 设置工具提示内容.
//
// pText: 工具提示内容.
func (e *Element) SetToolTip(pText string) int {
	return xc.XEle_SetToolTip(e.Handle, pText)
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (e *Element) SetToolTipEx(pText string, nTextAlign xcc.TextFormatFlag_) int {
	return xc.XEle_SetToolTipEx(e.Handle, pText, nTextAlign)
}

// 元素_取工具提示, 获取工具提示内容.
func (e *Element) GetToolTip() int {
	return xc.XEle_GetToolTip(e.Handle)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// x: X坐标.
//
// y: Y坐标.
func (e *Element) PopupToolTip(x int, y int) int {
	return xc.XEle_PopupToolTip(e.Handle, x, y)
}

// 元素_调整布局.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) AdjustLayout(nAdjustNo uint32) int {
	return xc.XEle_AdjustLayout(e.Handle, nAdjustNo)
}

// 元素_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) AdjustLayoutEx(nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_AdjustLayoutEx(e.Handle, nFlags, nAdjustNo)
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
func (e *Element) GetPosition(pOutX, pOutY *int32) int {
	return xc.XEle_GetPosition(e.Handle, pOutX, pOutY)
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
func (e *Element) SetSize(nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetSize(e.Handle, nWidth, nHeight, bRedraw, nFlags, nAdjustNo)
}

// 元素_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func (e *Element) GetSize(pOutWidth, pOutHeight *int32) int {
	return xc.XEle_GetSize(e.Handle, pOutWidth, pOutHeight)
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (e *Element) SetBkInfo(pText string) int {
	return xc.XEle_SetBkInfo(e.Handle, pText)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.
func (e *Element) GetWndClientRectDPI(pRect *xc.RECT) int {
	return xc.XEle_GetWndClientRectDPI(e.Handle, pRect)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pPt: 接收返回坐标点.
func (e *Element) PointClientToWndClientDPI(pPt *xc.POINT) int {
	return xc.XEle_PointClientToWndClientDPI(e.Handle, pPt)
}

// 元素_客户区坐标到窗口客户区DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.
func (e *Element) RectClientToWndClientDPI(pRect *xc.RECT) int {
	return xc.XEle_RectClientToWndClientDPI(e.Handle, pRect)
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

/*
下面都是事件
*/

type XE_ELEPROCE func(nEvent int, wParam, lParam uint, pbHandled *bool) int               // 元素处理过程事件.
type XE_ELEPROCE1 func(hEle int, nEvent int, wParam, lParam uint, pbHandled *bool) int    // 元素处理过程事件.
type XE_PAINT func(hDraw int, pbHandled *bool) int                                        // 元素绘制事件.
type XE_PAINT1 func(hEle int, hDraw int, pbHandled *bool) int                             // 元素绘制事件.
type XE_PAINT_END func(hDraw int, pbHandled *bool) int                                    // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_END1 func(hEle int, hDraw int, pbHandled *bool) int                         // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_SCROLLVIEW func(hDraw int, pbHandled *bool) int                             // 滚动视图绘制事件.
type XE_PAINT_SCROLLVIEW1 func(hEle int, hDraw int, pbHandled *bool) int                  // 滚动视图绘制事件.
type XE_MOUSEMOVE func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 元素鼠标移动事件.
type XE_MOUSEMOVE1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 元素鼠标移动事件.
type XE_MOUSESTAY func(pbHandled *bool) int                                               // 元素鼠标进入事件.
type XE_MOUSESTAY1 func(hEle int, pbHandled *bool) int                                    // 元素鼠标进入事件.
type XE_MOUSEHOVER func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                   // 元素鼠标悬停事件.
type XE_MOUSEHOVER1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int        // 元素鼠标悬停事件.
type XE_MOUSELEAVE func(hEleStay int, pbHandled *bool) int                                // 元素鼠标离开事件.
type XE_MOUSELEAVE1 func(hEle int, hEleStay int, pbHandled *bool) int                     // 元素鼠标离开事件.
type XE_MOUSEWHEEL func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                   // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_MOUSEWHEEL1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int        // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_LBUTTONDOWN func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                  // 鼠标左键按下事件.
type XE_LBUTTONDOWN1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int       // 鼠标左键按下事件.
type XE_LBUTTONUP func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 鼠标左键弹起事件.
type XE_LBUTTONUP1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 鼠标左键弹起事件.
type XE_RBUTTONDOWN func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                  // 鼠标右键按下事件.
type XE_RBUTTONDOWN1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int       // 鼠标右键按下事件.
type XE_RBUTTONUP func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 鼠标右键弹起事件.
type XE_RBUTTONUP1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 鼠标右键弹起事件.
type XE_LBUTTONDBCLICK func(nFlags int, pPt *xc.POINT, pbHandled *bool) int               // 鼠标左键双击事件.
type XE_LBUTTONDBCLICK1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int    // 鼠标左键双击事件.
type XE_XC_TIMER func(nTimerID int, pbHandled *bool) int                                  // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_XC_TIMER1 func(hEle int, nTimerID int, pbHandled *bool) int                       // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_ADJUSTLAYOUT func(nFlags int32, nAdjustNo uint32, pbHandled *bool) int            // 调整布局事件. 暂停使用.
type XE_ADJUSTLAYOUT1 func(hEle int, nFlags int32, nAdjustNo uint32, pbHandled *bool) int // 调整布局事件. 暂停使用.
type XE_TOOLTIP_POPUP func(hWindow int, pText uintptr, pbHandled *bool) int               // 元素工具提示弹出事件.
type XE_TOOLTIP_POPUP1 func(hEle int, hWindow int, pText uintptr, pbHandled *bool) int    // 元素工具提示弹出事件1.

// 调整布局完成事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 调整布局完成事件.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SETFOCUS func(pbHandled *bool) int               // 元素获得焦点事件.
type XE_SETFOCUS1 func(hEle int, pbHandled *bool) int    // 元素获得焦点事件.
type XE_KILLFOCUS func(pbHandled *bool) int              // 元素失去焦点事件.
type XE_KILLFOCUS1 func(hEle int, pbHandled *bool) int   // 元素失去焦点事件.
type XE_DESTROY func(pbHandled *bool) int                // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY1 func(hEle int, pbHandled *bool) int     // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY_END func(pbHandled *bool) int            // 元素销毁完成事件. 在销毁子对象之后触发.
type XE_DESTROY_END1 func(hEle int, pbHandled *bool) int // 元素销毁完成事件. 在销毁子对象之后触发.

// 元素大小改变事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 元素大小改变事件1.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SHOW func(bShow bool, pbHandled *bool) int                          // 元素显示隐藏事件.
type XE_SHOW1 func(hEle int, bShow bool, pbHandled *bool) int               // 元素显示隐藏事件.
type XE_SETFONT func(pbHandled *bool) int                                   // 元素设置字体事件.
type XE_SETFONT1 func(hEle int, pbHandled *bool) int                        // 元素设置字体事件.
type XE_KEYDOWN func(wParam, lParam uint, pbHandled *bool) int              // 元素按键事件.
type XE_KEYDOWN1 func(hEle int, wParam, lParam uint, pbHandled *bool) int   // 元素按键事件.
type XE_KEYUP func(wParam, lParam uint, pbHandled *bool) int                // 元素按键事件.
type XE_KEYUP1 func(hEle int, wParam, lParam uint, pbHandled *bool) int     // 元素按键事件.
type XE_CHAR func(wParam, lParam uint, pbHandled *bool) int                 // 通过TranslateMessage函数翻译的字符事件.
type XE_CHAR1 func(hEle int, wParam, lParam uint, pbHandled *bool) int      // 通过TranslateMessage函数翻译的字符事件.
type XE_SETCAPTURE func(pbHandled *bool) int                                // 元素设置鼠标捕获.
type XE_SETCAPTURE1 func(hEle int, pbHandled *bool) int                     // 元素设置鼠标捕获.
type XE_KILLCAPTURE func(pbHandled *bool) int                               // 元素失去鼠标捕获.
type XE_KILLCAPTURE1 func(hEle int, pbHandled *bool) int                    // 元素失去鼠标捕获.
type XE_SETCURSOR func(wParam, lParam uint, pbHandled *bool) int            // 设置鼠标光标.
type XE_SETCURSOR1 func(hEle int, wParam, lParam uint, pbHandled *bool) int // 设置鼠标光标.
type XE_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int              // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().
type XE_DROPFILES1 func(hEle int, hDropInfo uintptr, pbHandled *bool) int   // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().

// 元素处理过程事件.
func (e *Element) Event_ELEPROCE(pFun XE_ELEPROCE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ELEPROCE, pFun)
}

// 元素处理过程事件.
func (e *Element) Event_ELEPROCE1(pFun XE_ELEPROCE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ELEPROCE, pFun)
}

// 元素绘制事件.
func (e *Element) Event_PAINT(pFun XE_PAINT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT, pFun)
}

// 元素绘制事件.
func (e *Element) Event_PAINT1(pFun XE_PAINT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
func (e *Element) Event_PAINT_END(pFun XE_PAINT_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT_END, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
func (e *Element) Event_PAINT_END1(pFun XE_PAINT_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT_END, pFun)
}

// 滚动视图绘制事件.
func (e *Element) Event_PAINT_SCROLLVIEW(pFun XE_PAINT_SCROLLVIEW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// 滚动视图绘制事件.
func (e *Element) Event_PAINT_SCROLLVIEW1(pFun XE_PAINT_SCROLLVIEW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// 元素鼠标移动事件.
func (e *Element) Event_MOUSEMOVE(pFun XE_MOUSEMOVE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEMOVE, pFun)
}

// 元素鼠标移动事件.
func (e *Element) Event_MOUSEMOVE1(pFun XE_MOUSEMOVE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEMOVE, pFun)
}

// 元素鼠标进入事件.
func (e *Element) Event_MOUSESTAY(pFun XE_MOUSESTAY) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSESTAY, pFun)
}

// 元素鼠标进入事件.
func (e *Element) Event_MOUSESTAY1(pFun XE_MOUSESTAY1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSESTAY, pFun)
}

// 元素鼠标悬停事件.
func (e *Element) Event_MOUSEHOVER(pFun XE_MOUSEHOVER) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEHOVER, pFun)
}

// 元素鼠标悬停事件.
func (e *Element) Event_MOUSEHOVER1(pFun XE_MOUSEHOVER1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEHOVER, pFun)
}

// 元素鼠标离开事件.
func (e *Element) Event_MOUSELEAVE(pFun XE_MOUSELEAVE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSELEAVE, pFun)
}

// 元素鼠标离开事件.
func (e *Element) Event_MOUSELEAVE1(pFun XE_MOUSELEAVE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSELEAVE, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().
func (e *Element) Event_MOUSEWHEEL(pFun XE_MOUSEWHEEL) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MOUSEWHEEL, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().
func (e *Element) Event_MOUSEWHEEL1(pFun XE_MOUSEWHEEL1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MOUSEWHEEL, pFun)
}

// 鼠标左键按下事件.
func (e *Element) Event_LBUTTONDOWN(pFun XE_LBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键按下事件.
func (e *Element) Event_LBUTTONDOWN1(pFun XE_LBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键弹起事件.
func (e *Element) Event_LBUTTONUP(pFun XE_LBUTTONUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONUP, pFun)
}

// 鼠标左键弹起事件.
func (e *Element) Event_LBUTTONUP1(pFun XE_LBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONUP, pFun)
}

// 鼠标右键按下事件.
func (e *Element) Event_RBUTTONDOWN(pFun XE_RBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键按下事件.
func (e *Element) Event_RBUTTONDOWN1(pFun XE_RBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键弹起事件.
func (e *Element) Event_RBUTTONUP(pFun XE_RBUTTONUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_RBUTTONUP, pFun)
}

// 鼠标右键弹起事件.
func (e *Element) Event_RBUTTONUP1(pFun XE_RBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_RBUTTONUP, pFun)
}

// 鼠标左键双击事件.
func (e *Element) Event_LBUTTONDBCLICK(pFun XE_LBUTTONDBCLICK) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_LBUTTONDBCLICK, pFun)
}

// 鼠标左键双击事件.
func (e *Element) Event_LBUTTONDBCLICK1(pFun XE_LBUTTONDBCLICK1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_LBUTTONDBCLICK, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) Event_XC_TIMER(pFun XE_XC_TIMER) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_XC_TIMER, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) Event_XC_TIMER1(pFun XE_XC_TIMER1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_XC_TIMER, pFun)
}

// 调整布局事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT(pFun XE_ADJUSTLAYOUT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT1(pFun XE_ADJUSTLAYOUT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局完成事件.
func (e *Element) Event_ADJUSTLAYOUT_END(pFun XE_ADJUSTLAYOUT_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// 调整布局完成事件.
func (e *Element) Event_ADJUSTLAYOUT_END1(pFun XE_ADJUSTLAYOUT_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// 元素获得焦点事件.
func (e *Element) Event_SETFOCUS(pFun XE_SETFOCUS) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETFOCUS, pFun)
}

// 元素获得焦点事件.
func (e *Element) Event_SETFOCUS1(pFun XE_SETFOCUS1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETFOCUS, pFun)
}

// 元素失去焦点事件.
func (e *Element) Event_KILLFOCUS(pFun XE_KILLFOCUS) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KILLFOCUS, pFun)
}

// 元素失去焦点事件.
func (e *Element) Event_KILLFOCUS1(pFun XE_KILLFOCUS1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KILLFOCUS, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.
func (e *Element) Event_DESTROY(pFun XE_DESTROY) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DESTROY, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.
func (e *Element) Event_DESTROY1(pFun XE_DESTROY1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DESTROY, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.
func (e *Element) Event_DESTROY_END(pFun XE_DESTROY_END) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DESTROY_END, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.
func (e *Element) Event_DESTROY_END1(pFun XE_DESTROY_END1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DESTROY_END, pFun)
}

// 元素大小改变事件.
func (e *Element) Event_SIZE(pFun XE_SIZE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SIZE, pFun)
}

// 元素大小改变事件.
func (e *Element) Event_SIZE1(pFun XE_SIZE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SIZE, pFun)
}

// 元素显示隐藏事件.
func (e *Element) Event_SHOW(pFun XE_SHOW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SHOW, pFun)
}

// 元素显示隐藏事件.
func (e *Element) Event_SHOW1(pFun XE_SHOW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SHOW, pFun)
}

// 元素设置字体事件.
func (e *Element) Event_SETFONT(pFun XE_SETFONT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETFONT, pFun)
}

// 元素设置字体事件.
func (e *Element) Event_SETFONT1(pFun XE_SETFONT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETFONT, pFun)
}

// 元素按键事件.
func (e *Element) Event_KEYDOWN(pFun XE_KEYDOWN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KEYDOWN, pFun)
}

// 元素按键事件.
func (e *Element) Event_KEYDOWN1(pFun XE_KEYDOWN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KEYDOWN, pFun)
}

// 元素按键事件.
func (e *Element) Event_KEYUP(pFun XE_KEYUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KEYUP, pFun)
}

// 元素按键事件.
func (e *Element) Event_KEYUP1(pFun XE_KEYUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KEYUP, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.
func (e *Element) Event_CHAR(pFun XE_CHAR) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_CHAR, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.
func (e *Element) Event_CHAR1(pFun XE_CHAR1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_CHAR, pFun)
}

// 元素设置鼠标捕获.
func (e *Element) Event_SETCAPTURE(pFun XE_SETCAPTURE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETCAPTURE, pFun)
}

// 元素设置鼠标捕获.
func (e *Element) Event_SETCAPTURE1(pFun XE_SETCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETCAPTURE, pFun)
}

// 元素失去鼠标捕获.
func (e *Element) Event_KILLCAPTURE(pFun XE_KILLCAPTURE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_KILLCAPTURE, pFun)
}

// 元素失去鼠标捕获.
func (e *Element) Event_KILLCAPTURE1(pFun XE_KILLCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_KILLCAPTURE, pFun)
}

// 设置鼠标光标.
func (e *Element) Event_SETCURSOR(pFun XE_SETCURSOR) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_SETCURSOR, pFun)
}

// 设置鼠标光标.
func (e *Element) Event_SETCURSOR1(pFun XE_SETCURSOR1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_SETCURSOR, pFun)
}

// 文件拖放事件, 需先启用:XWnd_EnableDragFiles().
func (e *Element) Event_DROPFILES(pFun XE_DROPFILES) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_DROPFILES, pFun)
}

// 文件拖放事件, 需先启用:XWnd_EnableDragFiles().
func (e *Element) Event_DROPFILES1(pFun XE_DROPFILES1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_DROPFILES, pFun)
}

// 事件_元素工具提示弹出, 可使用 common.UintPtrToString 把uintptr转换到文本.
func (e *Element) Event_TOOLTIP_POPUP(pFun XE_TOOLTIP_POPUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_TOOLTIP_POPUP, pFun)
}

// 事件_元素工具提示弹出1, 可使用 common.UintPtrToString 把uintptr转换到文本.
func (e *Element) Event_TOOLTIP_POPUP1(pFun XE_TOOLTIP_POPUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_TOOLTIP_POPUP, pFun)
}

type XE_MENU_SELECT func(nID int32, pbHandled *bool) int                                          // 弹出菜单项选择事件.
type XE_MENU_POPUP func(HMENUX int, pbHandled *bool) int                                          // 菜单弹出.
type XE_MENU_EXIT func(pbHandled *bool) int                                                       // 弹出菜单退出事件.
type XE_MENU_POPUP_WND func(hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM func(hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 事件_弹出菜单项被选择.
func (e *Element) Event_MENU_SELECT(pFun XE_MENU_SELECT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.
func (e *Element) Event_MENU_POPUP(pFun XE_MENU_POPUP) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.
func (e *Element) Event_MENU_EXIT(pFun XE_MENU_EXIT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.
func (e *Element) Event_MENU_POPUP_WND(pFun XE_MENU_POPUP_WND) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (e *Element) Event_MENU_DRAW_BACKGROUND(pFun XE_MENU_DRAW_BACKGROUND) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (e *Element) Event_MENU_DRAWITEM(pFun XE_MENU_DRAWITEM) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_MENU_DRAWITEM, pFun)
}

type XE_MENU_SELECT1 func(hEle int, nID int, pbHandled *bool) int                                            // 弹出菜单项选择事件.
type XE_MENU_POPUP1 func(hEle int, HMENUX int, pbHandled *bool) int                                          // 菜单弹出.
type XE_MENU_EXIT1 func(hEle int, pbHandled *bool) int                                                       // 弹出菜单退出事件.
type XE_MENU_POPUP_WND1 func(hEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND1 func(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM1 func(hEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 事件_弹出菜单项被选择.
func (e *Element) Event_MENU_SELECT1(pFun XE_MENU_SELECT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.
func (e *Element) Event_MENU_POPUP1(pFun XE_MENU_POPUP1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.
func (e *Element) Event_MENU_EXIT1(pFun XE_MENU_EXIT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.
func (e *Element) Event_MENU_POPUP_WND1(pFun XE_MENU_POPUP_WND1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (e *Element) Event_MENU_DRAW_BACKGROUND1(pFun XE_MENU_DRAW_BACKGROUND1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (e *Element) Event_MENU_DRAWITEM1(pFun XE_MENU_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_MENU_DRAWITEM, pFun)
}
