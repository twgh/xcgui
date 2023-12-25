package xc

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 元素_创建, 创建基础元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func XEle_Create(x, y, cx, cy int32, hParent int) int {
	r, _, _ := xEle_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func XEle_RegEventC(hEle int, nEvent xcc.XE_, pFun interface{}) bool {
	r, _, _ := xEle_RegEventC.Call(uintptr(hEle), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func XEle_RegEventC1(hEle int, nEvent xcc.XE_, pFun interface{}) bool {
	r, _, _ := xEle_RegEventC1.Call(uintptr(hEle), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 元素_移除事件C.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func XEle_RemoveEventC(hEle int, nEvent xcc.XE_, pFun interface{}) bool {
	r, _, _ := xEle_RemoveEventC.Call(uintptr(hEle), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 元素_注册事件CEx, 注册事件C方式, 省略2参数, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XEle_RegEventCEx(hEle int, nEvent xcc.XE_, pFun uintptr) bool {
	r, _, _ := xEle_RegEventC.Call(uintptr(hEle), uintptr(nEvent), pFun)
	return r != 0
}

// 元素_注册事件C1Ex, 注册事件C1方式, 省略1参数, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XEle_RegEventC1Ex(hEle int, nEvent xcc.XE_, pFun uintptr) bool {
	r, _, _ := xEle_RegEventC1.Call(uintptr(hEle), uintptr(nEvent), pFun)
	return r != 0
}

// 元素_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XEle_RemoveEventCEx(hEle int, nEvent xcc.XE_, pFun uintptr) bool {
	r, _, _ := xEle_RemoveEventC.Call(uintptr(hEle), uintptr(nEvent), pFun)
	return r != 0
}

// 元素_发送事件.
//
// hEle: 目标元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func XEle_SendEvent(hEle int, nEvent xcc.XE_, wParam, lParam uint) int {
	r, _, _ := xEle_SendEvent.Call(uintptr(hEle), uintptr(nEvent), uintptr(wParam), uintptr(lParam))
	return int(r)
}

// 元素_投递事件.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func XEle_PostEvent(hEle int, nEvent xcc.XE_, wParam, lParam uint) int {
	r, _, _ := xEle_PostEvent.Call(uintptr(hEle), uintptr(nEvent), uintptr(wParam), uintptr(lParam))
	return int(r)
}

// 元素_取坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_GetRect(hEle int, pRect *RECT) int {
	r, _, _ := xEle_GetRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_GetRectLogic(hEle int, pRect *RECT) int {
	r, _, _ := xEle_GetRectLogic.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_取客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_GetClientRect(hEle int, pRect *RECT) int {
	r, _, _ := xEle_GetClientRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_置宽度.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
func XEle_SetWidth(hEle int, nWidth int) int {
	r, _, _ := xEle_SetWidth.Call(uintptr(hEle), uintptr(nWidth))
	return int(r)
}

// 元素_置高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
func XEle_SetHeight(hEle int, nHeight int) int {
	r, _, _ := xEle_SetHeight.Call(uintptr(hEle), uintptr(nHeight))
	return int(r)
}

// 元素_取宽度.
//
// hEle: 元素句柄.
func XEle_GetWidth(hEle int) int {
	r, _, _ := xEle_GetWidth.Call(uintptr(hEle))
	return int(r)
}

// 元素_取高度.
//
// hEle: 元素句柄.
func XEle_GetHeight(hEle int) int {
	r, _, _ := xEle_GetHeight.Call(uintptr(hEle))
	return int(r)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_RectWndClientToEleClient(hEle int, pRect *RECT) int {
	r, _, _ := xEle_RectWndClientToEleClient.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// hEle: 元素句柄.
//
// pPt: 坐标.
func XEle_PointWndClientToEleClient(hEle int, pPt *POINT) int {
	r, _, _ := xEle_PointWndClientToEleClient.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_RectClientToWndClient(hEle int, pRect *RECT) int {
	r, _, _ := xEle_RectClientToWndClient.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pPt: 坐标.
func XEle_PointClientToWndClient(hEle int, pPt *POINT) int {
	r, _, _ := xEle_PointClientToWndClient.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 元素_基于窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func XEle_GetWndClientRect(hEle int, pRect *RECT) int {
	r, _, _ := xEle_GetWndClientRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.
//
// hEle: 元素句柄.
func XEle_GetCursor(hEle int) int {
	r, _, _ := xEle_GetCursor.Call(uintptr(hEle))
	return int(r)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hEle: 元素句柄.
//
// hCursor: 光标句柄.
func XEle_SetCursor(hEle int, hCursor int) int {
	r, _, _ := xEle_SetCursor.Call(uintptr(hEle), uintptr(hCursor))
	return int(r)
}

// 元素_添加子对象.
//
// hEle: 元素句柄.
//
// hChild: 要添加的子元素句柄或形状对象句柄.
func XEle_AddChild(hEle int, hChild int) bool {
	r, _, _ := xEle_AddChild.Call(uintptr(hEle), uintptr(hChild))
	return r != 0
}

// 元素_插入子对象, 插入子对象到指定位置.
//
// hEle: 元素句柄.
//
// hChild: 要插入的元素句柄或形状对象句柄.
//
// index: 插入位置索引.
func XEle_InsertChild(hEle int, hChild int, index int) bool {
	r, _, _ := xEle_InsertChild.Call(uintptr(hEle), uintptr(hChild), uintptr(index))
	return r != 0
}

// 元素_置坐标, 如果返回0坐标没有改变, 如果大小改变返回2(触发XE_SIZE), 否则返回1(仅改变left,top,没有改变大小).
//
// hEle: 元素句柄.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func XEle_SetRect(hEle int, pRect *RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_置坐标扩展, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
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
func XEle_SetRectEx(hEle int, x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetRectEx.Call(uintptr(hEle), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_置逻辑坐标, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_ , 此参数将会传入XE_SIZE ,XE_ADJUSTLAYOUT 事件回调.
//
// nAdjustNo: 调整布局流水号, 可填0.
func XEle_SetRectLogic(hEle int, pRect *RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetRectLogic.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_移动, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
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
func XEle_SetPosition(hEle int, x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetPosition.Call(uintptr(hEle), uintptr(x), uintptr(y), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_移动逻辑坐标, 移动元素坐标, 逻辑坐标, 包含滚动视图偏移. 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
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
func XEle_SetPositionLogic(hEle int, x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetPositionLogic.Call(uintptr(hEle), uintptr(x), uintptr(y), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_判断绘制焦点.
//
// hEle: 元素句柄.
func XEle_IsDrawFocus(hEle int) bool {
	r, _, _ := xEle_IsDrawFocus.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断启用, 元素是否为启用状态.
//
// hEle: 元素句柄.
func XEle_IsEnable(hEle int) bool {
	r, _, _ := xEle_IsEnable.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断启用焦点, 元素是否启用焦点.
//
// hEle: 元素句柄.
func XEle_IsEnableFocus(hEle int) bool {
	r, _, _ := xEle_IsEnableFocus.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透.
//
// hEle: 元素句柄.
func XEle_IsMouseThrough(hEle int) bool {
	r, _, _ := xEle_IsMouseThrough.Call(uintptr(hEle))
	return r != 0
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func XEle_HitChildEle(hEle int, pPt *POINT) int {
	r, _, _ := xEle_HitChildEle.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 元素_判断背景透明.
//
// hEle: 元素句柄.
func XEle_IsBkTransparent(hEle int) bool {
	r, _, _ := xEle_IsBkTransparent.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件.
//
// hEle: 元素句柄.
func XEle_IsEnableEvent_XE_PAINT_END(hEle int) bool {
	r, _, _ := xEle_IsEnableEvent_XE_PAINT_END.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit.
//
// hEle: 元素句柄.
func XEle_IsKeyTab(hEle int) bool {
	r, _, _ := xEle_IsKeyTab.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键).
//
// hEle: 元素句柄.
func XEle_IsSwitchFocus(hEle int) bool {
	r, _, _ := xEle_IsSwitchFocus.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素.
//
// hEle: 元素句柄.
func XEle_IsEnable_XE_MOUSEWHEEL(hEle int) bool {
	r, _, _ := xEle_IsEnable_XE_MOUSEWHEEL.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素.
//
// hEle: 元素句柄.
//
// hChildEle: 子元素句柄.
func XEle_IsChildEle(hEle int, hChildEle int) bool {
	r, _, _ := xEle_IsChildEle.Call(uintptr(hEle), uintptr(hChildEle))
	return r != 0
}

// 元素_判断启用画布, 判断是否启用画布.
//
// hEle: 元素句柄.
func XEle_IsEnableCanvas(hEle int) bool {
	r, _, _ := xEle_IsEnableCanvas.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断焦点, 判断是否拥有焦点.
//
// hEle: 元素句柄.
func XEle_IsFocus(hEle int) bool {
	r, _, _ := xEle_IsFocus.Call(uintptr(hEle))
	return r != 0
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点.
//
// hEle: 元素句柄.
func XEle_IsFocusEx(hEle int) bool {
	r, _, _ := xEle_IsFocusEx.Call(uintptr(hEle))
	return r != 0
}

// 元素_启用, 启用或禁用元素.
//
// hEle: 元素句柄.
//
// bEnable: 启用或禁用.
func XEle_Enable(hEle int, bEnable bool) int {
	r, _, _ := xEle_Enable.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用焦点, 启用焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableFocus(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableFocus.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用绘制焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableDrawFocus(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableDrawFocus.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableDrawBorder(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableDrawBorder.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableCanvas(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableCanvas.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用事件_XE_PAINT_END.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableEvent_XE_PAINT_END(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableEvent_XE_PAINT_END.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用背景透明.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableBkTransparent(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableBkTransparent.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableMouseThrough(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableMouseThrough.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableKeyTab(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableKeyTab.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableSwitchFocus(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableSwitchFocus.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEle_EnableEvent_XE_MOUSEWHEEL(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableEvent_XE_MOUSEWHEEL.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_移除, 移除元素, 但不销毁.
//
// hEle: 元素句柄.
func XEle_Remove(hEle int) int {
	r, _, _ := xEle_Remove.Call(uintptr(hEle))
	return int(r)
}

// 元素_置Z序, 设置元素Z序.
//
// hEle: 元素句柄.
//
// index: 位置索引.
func XEle_SetZOrder(hEle int, index int) bool {
	r, _, _ := xEle_SetZOrder.Call(uintptr(hEle), uintptr(index))
	return r != 0
}

// 元素_置Z序扩展, 设置元素Z序.
//
// hEle: 元素句柄.
//
// hDestEle: 目标元素.
//
// nType: 类型, Zorder_.
func XEle_SetZOrderEx(hEle int, hDestEle int, nType xcc.Zorder_) bool {
	r, _, _ := xEle_SetZOrderEx.Call(uintptr(hEle), uintptr(hDestEle), uintptr(nType))
	return r != 0
}

// 元素_取Z序, 获取元素Z序索引, 位置索引.
//
// hEle: 元素句柄.
func XEle_GetZOrder(hEle int) int {
	r, _, _ := xEle_GetZOrder.Call(uintptr(hEle))
	return int(r)
}

// 元素_启用置顶, 设置元素置顶.
//
// hEle: 元素句柄.
//
// bTopmost: 是否置顶显示.
func XEle_EnableTopmost(hEle int, bTopmost bool) bool {
	r, _, _ := xEle_EnableTopmost.Call(uintptr(hEle), common.BoolPtr(bTopmost))
	return r != 0
}

// 元素_重绘.
//
// hEle: 元素句柄.
//
// bImmediate: 是否立即重绘.
func XEle_Redraw(hEle int, bImmediate bool) int {
	r, _, _ := xEle_Redraw.Call(uintptr(hEle), common.BoolPtr(bImmediate))
	return int(r)
}

// 元素_重绘指定区域.
//
// hEle: 元素句柄.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.
func XEle_RedrawRect(hEle int, pRect *RECT, bImmediate bool) int {
	r, _, _ := xEle_RedrawRect.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bImmediate))
	return int(r)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.
//
// hEle: 元素句柄.
func XEle_GetChildCount(hEle int) int {
	r, _, _ := xEle_GetChildCount.Call(uintptr(hEle))
	return int(r)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// hEle: 元素句柄.
//
// index: 索引.
func XEle_GetChildByIndex(hEle int, index int) int {
	r, _, _ := xEle_GetChildByIndex.Call(uintptr(hEle), uintptr(index))
	return int(r)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// hEle: 元素句柄.
//
// nID: 元素ID.
func XEle_GetChildByID(hEle int, nID int) int {
	r, _, _ := xEle_GetChildByID.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 元素_置边框大小.
//
// hEle: 元素句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func XEle_SetBorderSize(hEle int, left int, top int, right int, bottom int) int {
	r, _, _ := xEle_SetBorderSize.Call(uintptr(hEle), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 元素_取边框大小.
//
// hEle: 元素句柄.
//
// pBorder: 大小.
func XEle_GetBorderSize(hEle int, pBorder *RECT) int {
	r, _, _ := xEle_GetBorderSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pBorder)))
	return int(r)
}

// 元素_置内填充大小.
//
// hEle: 元素句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func XEle_SetPadding(hEle int, left int, top int, right int, bottom int) int {
	r, _, _ := xEle_SetPadding.Call(uintptr(hEle), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 元素_取内填充大小.
//
// hEle: 元素句柄.
//
// pPadding: 大小.
func XEle_GetPadding(hEle int, pPadding *RECT) int {
	r, _, _ := xEle_GetPadding.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPadding)))
	return int(r)
}

// 元素_置拖动边框.
//
// hEle: 元素句柄.
//
// nFlags: 边框位置组合, Element_Position_.
func XEle_SetDragBorder(hEle int, nFlags xcc.Element_Position_) int {
	r, _, _ := xEle_SetDragBorder.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// hEle: 元素句柄.
//
// nFlags: 边框位置标识, Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.
func XEle_SetDragBorderBindEle(hEle int, nFlags xcc.Element_Position_, hBindEle int, nSpace int) int {
	r, _, _ := xEle_SetDragBorderBindEle.Call(uintptr(hEle), uintptr(nFlags), uintptr(hBindEle), uintptr(nSpace))
	return int(r)
}

// 元素_置最小大小.
//
// hEle: 元素句柄.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.
func XEle_SetMinSize(hEle int, nWidth int, nHeight int) int {
	r, _, _ := xEle_SetMinSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 元素_置最大大小.
//
// hEle: 元素句柄.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.
func XEle_SetMaxSize(hEle int, nWidth int, nHeight int) int {
	r, _, _ := xEle_SetMaxSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// hEle: 元素句柄.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.
func XEle_SetLockScroll(hEle int, bHorizon bool, bVertical bool) int {
	r, _, _ := xEle_SetLockScroll.Call(uintptr(hEle), common.BoolPtr(bHorizon), common.BoolPtr(bVertical))
	return int(r)
}

// 元素_置文本颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XEle_SetTextColor(hEle int, color int) int {
	r, _, _ := xEle_SetTextColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 元素_取文本颜色.
//
// hEle: 元素句柄.
func XEle_GetTextColor(hEle int) int {
	r, _, _ := xEle_GetTextColor.Call(uintptr(hEle))
	return int(r)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.
//
// hEle: 元素句柄.
func XEle_GetTextColorEx(hEle int) int {
	r, _, _ := xEle_GetTextColorEx.Call(uintptr(hEle))
	return int(r)
}

// 元素_置焦点边框颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XEle_SetFocusBorderColor(hEle int, color int) int {
	r, _, _ := xEle_SetFocusBorderColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 元素_取焦点边框颜色.
//
// hEle: 元素句柄.
func XEle_GetFocusBorderColor(hEle int) int {
	r, _, _ := xEle_GetFocusBorderColor.Call(uintptr(hEle))
	return int(r)
}

// 元素_置字体.
//
// hEle: 元素句柄.
//
// hFontx: 炫彩字体.
func XEle_SetFont(hEle int, hFontx int) int {
	r, _, _ := xEle_SetFont.Call(uintptr(hEle), uintptr(hFontx))
	return int(r)
}

// 元素_取字体.
//
// hEle: 元素句柄.
func XEle_GetFont(hEle int) int {
	r, _, _ := xEle_GetFont.Call(uintptr(hEle))
	return int(r)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.
//
// hEle: 元素句柄.
func XEle_GetFontEx(hEle int) int {
	r, _, _ := xEle_GetFontEx.Call(uintptr(hEle))
	return int(r)
}

// 元素_置透明度.
//
// hEle: 元素句柄.
func XEle_SetAlpha(hEle int, alpha uint8) int {
	r, _, _ := xEle_SetAlpha.Call(uintptr(hEle), uintptr(alpha))
	return int(r)
}

// 元素_销毁.
//
// hEle: 元素句柄.
func XEle_Destroy(hEle int) int {
	r, _, _ := xEle_Destroy.Call(uintptr(hEle))
	return int(r)
}

// 元素_添加背景边框, 添加背景内容边框.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func XEle_AddBkBorder(hEle int, nState xcc.CombinedState, color int, width int) int {
	r, _, _ := xEle_AddBkBorder.Call(uintptr(hEle), uintptr(nState), uintptr(color), uintptr(width))
	return int(r)
}

// 元素_添加背景填充, 添加背景内容填充.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func XEle_AddBkFill(hEle int, nState xcc.CombinedState, color int) int {
	r, _, _ := xEle_AddBkFill.Call(uintptr(hEle), uintptr(nState), uintptr(color))
	return int(r)
}

// 元素_添加背景图片, 添加背景内容图片.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func XEle_AddBkImage(hEle int, nState xcc.CombinedState, hImage int) int {
	r, _, _ := xEle_AddBkImage.Call(uintptr(hEle), uintptr(nState), uintptr(hImage))
	return int(r)
}

// 元素_取背景对象数量, 获取背景内容数量.
//
// hEle: 元素句柄.
func XEle_GetBkInfoCount(hEle int) int {
	r, _, _ := xEle_GetBkInfoCount.Call(uintptr(hEle))
	return int(r)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
//
// hEle: 元素句柄.
func XEle_ClearBkInfo(hEle int) int {
	r, _, _ := xEle_ClearBkInfo.Call(uintptr(hEle))
	return int(r)
}

// 元素_取背景管理器, 获取元素背景管理器.
//
// hEle: 元素句柄.
func XEle_GetBkManager(hEle int) int {
	r, _, _ := xEle_GetBkManager.Call(uintptr(hEle))
	return int(r)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.
//
// hEle: 元素句柄.
func XEle_GetBkManagerEx(hEle int) int {
	r, _, _ := xEle_GetBkManagerEx.Call(uintptr(hEle))
	return int(r)
}

// 元素_置背景管理器.
//
// hEle: 元素句柄.
//
// hBkInfoM: 背景管理器.
func XEle_SetBkManager(hEle int, hBkInfoM int) int {
	r, _, _ := xEle_SetBkManager.Call(uintptr(hEle), uintptr(hBkInfoM))
	return int(r)
}

// 元素_取状态, 获取组合状态.
//
// hEle: 元素句柄.
func XEle_GetStateFlags(hEle int) xcc.CombinedState {
	r, _, _ := xEle_GetStateFlags.Call(uintptr(hEle))
	return xcc.CombinedState(r)
}

// 元素_绘制焦点, 绘制元素焦点.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func XEle_DrawFocus(hEle int, hDraw int, pRect *RECT) bool {
	r, _, _ := xEle_DrawFocus.Call(uintptr(hEle), uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.
func XEle_DrawEle(hEle int, hDraw int) int {
	r, _, _ := xEle_DrawEle.Call(uintptr(hEle), uintptr(hDraw))
	return int(r)
}

// 元素_置用户数据.
//
// hEle: 元素句柄.
//
// nData: 用户数据.
func XEle_SetUserData(hEle int, nData int) int {
	r, _, _ := xEle_SetUserData.Call(uintptr(hEle), uintptr(nData))
	return int(r)
}

// 元素_取用户数据.
//
// hEle: 元素句柄.
func XEle_GetUserData(hEle int) int {
	r, _, _ := xEle_GetUserData.Call(uintptr(hEle))
	return int(r)
}

// 元素_取内容大小.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直, 布局属性交换依赖.
//
// cx: 宽度.
//
// cy: 高度.
//
// pSize: 返回大小.
func XEle_GetContentSize(hEle int, bHorizon bool, cx int, cy int, pSize *SIZE) int {
	r, _, _ := xEle_GetContentSize.Call(uintptr(hEle), common.BoolPtr(bHorizon), uintptr(cx), uintptr(cy), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 元素_置鼠标捕获.
//
// hEle: 元素句柄.
//
// b: TRUE设置.
func XEle_SetCapture(hEle int, b bool) int {
	r, _, _ := xEle_SetCapture.Call(uintptr(hEle), common.BoolPtr(b))
	return int(r)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// hEle: 元素句柄.
//
// bEnable: 启用或关闭.
func XEle_EnableTransparentChannel(hEle int, bEnable bool) int {
	r, _, _ := xEle_EnableTransparentChannel.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// hEle: 元素句柄.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.
func XEle_SetXCTimer(hEle int, nIDEvent int, uElapse int) bool {
	r, _, _ := xEle_SetXCTimer.Call(uintptr(hEle), uintptr(nIDEvent), uintptr(uElapse))
	return r != 0
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// hEle: 元素句柄.
//
// nIDEvent: 事件ID.
func XEle_KillXCTimer(hEle int, nIDEvent int) bool {
	r, _, _ := xEle_KillXCTimer.Call(uintptr(hEle), uintptr(nIDEvent))
	return r != 0
}

// 元素_置工具提示, 设置工具提示内容.
//
// hEle: 元素句柄.
//
// pText: 工具提示内容.
func XEle_SetToolTip(hEle int, pText string) int {
	r, _, _ := xEle_SetToolTip.Call(uintptr(hEle), common.StrPtr(pText))
	return int(r)
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// hEle: 元素句柄.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func XEle_SetToolTipEx(hEle int, pText string, nTextAlign xcc.TextFormatFlag_) int {
	r, _, _ := xEle_SetToolTipEx.Call(uintptr(hEle), common.StrPtr(pText), uintptr(nTextAlign))
	return int(r)
}

// 元素_取工具提示, 获取工具提示内容.
//
// hEle: 元素句柄.
func XEle_GetToolTip(hEle int) int {
	r, _, _ := xEle_GetToolTip.Call(uintptr(hEle))
	return int(r)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// hEle: 元素句柄.
//
// x: X坐标.
//
// y: Y坐标.
func XEle_PopupToolTip(hEle int, x int, y int) int {
	r, _, _ := xEle_PopupToolTip.Call(uintptr(hEle), uintptr(x), uintptr(y))
	return int(r)
}

// 元素_调整布局.
//
// hEle: 元素句柄.
//
// nAdjustNo: 调整布局流水号, 可填0.
func XEle_AdjustLayout(hEle int, nAdjustNo uint32) int {
	r, _, _ := xEle_AdjustLayout.Call(uintptr(hEle), uintptr(nAdjustNo))
	return int(r)
}

// 元素_调整布局扩展.
//
// hEle: 元素句柄.
//
// nFlags: 调整标识.
//
// nAdjustNo: 调整布局流水号, 可填0.
func XEle_AdjustLayoutEx(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_AdjustLayoutEx.Call(uintptr(hEle), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_取透明度, 返回透明度.
//
// hEle: 元素句柄.
func XEle_GetAlpha(hEle int) byte {
	r, _, _ := xEle_GetAlpha.Call(uintptr(hEle))
	return byte(r)
}

// 元素_取位置.
//
// hEle: 元素句柄.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func XEle_GetPosition(hEle int, pOutX, pOutY *int32) int {
	r, _, _ := xEle_GetPosition.Call(uintptr(hEle), uintptr(unsafe.Pointer(pOutX)), uintptr(unsafe.Pointer(pOutY)))
	return int(r)
}

// 元素_置大小.
//
// hEle: 元素句柄.
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
func XEle_SetSize(hEle int, nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	r, _, _ := xEle_SetSize.Call(uintptr(hEle), uintptr(nWidth), uintptr(nHeight), common.BoolPtr(bRedraw), uintptr(nFlags), uintptr(nAdjustNo))
	return int(r)
}

// 元素_取大小.
//
// hEle: 元素句柄.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func XEle_GetSize(hEle int, pOutWidth, pOutHeight *int32) int {
	r, _, _ := xEle_GetSize.Call(uintptr(hEle), uintptr(unsafe.Pointer(pOutWidth)), uintptr(unsafe.Pointer(pOutHeight)))
	return int(r)
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// hEle: 元素句柄.
//
// pText: 背景内容字符串.
func XEle_SetBkInfo(hEle int, pText string) int {
	r, _, _ := xEle_SetBkInfo.Call(uintptr(hEle), common.StrPtr(pText))
	return int(r)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回坐标.
func XEle_GetWndClientRectDPI(hEle int, pRect *RECT) int {
	r, _, _ := xEle_GetWndClientRectDPI.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pPt: 接收返回坐标点.
func XEle_PointClientToWndClientDPI(hEle int, pPt *POINT) int {
	r, _, _ := xEle_PointClientToWndClientDPI.Call(uintptr(hEle), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 元素_客户区坐标到窗口客户区DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回坐标.
func XEle_RectClientToWndClientDPI(hEle int, pRect *RECT) int {
	r, _, _ := xEle_RectClientToWndClientDPI.Call(uintptr(hEle), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}
