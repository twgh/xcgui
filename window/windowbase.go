package window

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// windowBase 窗口基类.
type windowBase struct {
	objectbase.UI
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//
// bEnable: 开启开关.
func (w *windowBase) EnableAutoClose(bEnable bool) *windowBase {
	xc.XModalWnd_EnableAutoClose(w.Handle, bEnable)
	return w
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//
// bEnable: 是否启用.
func (w *windowBase) EnableEscClose(bEnable bool) *windowBase {
	xc.XModalWnd_EnableEscClose(w.Handle, bEnable)
	return w
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回:
//   - xcc.MessageBox_Flag_Ok: 点击确定按钮退出.
//   - xcc.MessageBox_Flag_Cancel: 点击取消按钮退出.
//   - xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *windowBase) DoModal() xcc.MessageBox_Flag_ {
	return xc.XModalWnd_DoModal(w.Handle)
}

// 模态窗口_结束, 结束模态窗口.
//
// nResult: 用作 xc.XModalWnd_DoModal 的返回值:
//   - xcc.MessageBox_Flag_Ok: 点击确定按钮退出.
//   - xcc.MessageBox_Flag_Cancel: 点击取消按钮退出.
//   - xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *windowBase) EndModal(nResult xcc.MessageBox_Flag_) *windowBase {
	xc.XModalWnd_EndModal(w.Handle, nResult)
	return w
}

// MessageBox 炫彩_消息框. 返回值如下:
//   - xcc.MessageBox_Flag_Ok: 点击确定按钮退出.
//   - xcc.MessageBox_Flag_Cancel: 点击取消按钮退出.
//   - xcc.MessageBox_Flag_Other: 其他方式退出.
//
// pTitle: 标题.
//
// pText: 内容文本.
//
// nFlags: 标识: xcc.MessageBox_Flag_.
//
// XCStyle: xcc.Window_Style_.
func (w *windowBase) MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {
	return xc.XC_MessageBox(pTitle, pText, nFlags, w.GetHWND(), XCStyle)
}

// Msg_Create 消息框_创建, 返回消息框窗口对象.
//
// pTitle: 标题.
//
// pText: 内容文本.
//
// nFlags: 标识: xcc.MessageBox_Flag_.
//
// XCStyle: xcc.Window_Style_.
func (w *windowBase) Msg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *Window {
	wd := &Window{}
	hWindow := xc.XMsg_Create(pTitle, pText, nFlags, w.GetHWND(), XCStyle)
	if xc.XC_IsHWINDOW(hWindow) {
		wd.SetHandle(hWindow)
	}
	return wd
}

// Msg_CreateEx 消息框_创建扩展, 返回消息框窗口对象.
//
// dwExStyle: 窗口扩展样式, xcc.WS_EX_ 常量组合.
//
// dwStyle: 窗口样式, xcc.WS_ 常量组合.
//
// lpClassName: 窗口类名.
//
// pTitle: 标题.
//
// pText: 内容文本.
//
// nFlags: 标识: xcc.MessageBox_Flag_.
//
// XCStyle: xcc.Window_Style_.
func (w *windowBase) Msg_CreateEx(dwExStyle xcc.WS_EX_, dwStyle xcc.WS_, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *Window {
	wd := &Window{}
	hWindow := xc.XMsg_CreateEx(dwExStyle, dwStyle, lpClassName, pTitle, pText, nFlags, w.GetHWND(), XCStyle)
	if xc.XC_IsHWINDOW(hWindow) {
		wd.SetHandle(hWindow)
	}
	return wd
}

// 炫彩_发送窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) SendMessage(msg uint32, wParam, lParam uintptr) uintptr {
	return xc.XC_SendMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_投递窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) PostMessage(msg uint32, wParam, lParam uintptr) bool {
	return xc.XC_PostMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
func (w *windowBase) IsHWINDOW() bool {
	return xc.XC_IsHWINDOW(w.Handle)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// nID: ID值.
func (w *windowBase) GetObjectByID(nID int32) int {
	return xc.XC_GetObjectByID(w.Handle, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// pName: ID名称.
func (w *windowBase) GetObjectByIDName(pName string) int {
	return xc.XC_GetObjectByIDName(w.Handle, pName)
}

// 窗口_显示.
//
// nCmdShow: 显示方式: xcc.SW_.
func (w *windowBase) ShowWindow(nCmdShow xcc.SW_) bool {
	return xc.XWnd_ShowWindow(w.Handle, nCmdShow)
}

// 窗口_置顶.
//
// bTop: 是否置顶.
func (w *windowBase) SetTop(bTop ...bool) *windowBase {
	xc.XWnd_SetTop(w.Handle, bTop...)
	return w
}

// 窗口_移除事件. 只适用于 AddEvent_ 方式添加的事件.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// index: 使用 AddEvent_ 函数返回的回调函数索引.
//   - 为空时, 直接移除事件.
//   - 不为空时, 移除指定索引的回调函数.
func (w *windowBase) RemoveEvent(nEvent xcc.WM_, index ...int) *windowBase {
	if len(index) > 0 {
		xc.WndEventHandler.RemoveCallBack(w.Handle, nEvent, index[0])
	} else { // 移除事件
		cbPtr := xc.WndEventHandler.EventInfoMap[w.Handle][nEvent].EvnetFuncPtr
		if cbPtr > 0 {
			xc.XWnd_RemoveEventCEx(w.Handle, nEvent, cbPtr)
		}
		xc.WndEventHandler.RemoveEvent(w.Handle, nEvent)
	}
	return w
}

// 窗口_注册事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// fun: 事件函数.
func (w *windowBase) RegEventC(nEvent xcc.WM_, fun interface{}) bool {
	return xc.XWnd_RegEventC(w.Handle, nEvent, fun)
}

// 窗口_注册事件C1.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// fun: 事件函数.
func (w *windowBase) RegEventC1(nEvent xcc.WM_, fun interface{}) bool {
	return xc.XWnd_RegEventC1(w.Handle, nEvent, fun)
}

// 窗口_移除事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// fun: 事件函数.
func (w *windowBase) RemoveEventC(nEvent xcc.WM_, fun interface{}) bool {
	return xc.XWnd_RemoveEventC(w.Handle, nEvent, fun)
}

// 窗口_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成..
func (w *windowBase) RemoveEventCEx(nEvent xcc.WM_, pFun uintptr) bool {
	return xc.XWnd_RemoveEventCEx(w.Handle, nEvent, pFun)
}

// 窗口_添加子对象.
//
// hChild: 要添加的对象句柄.
func (w *windowBase) AddChild(hChild int) bool {
	return xc.XWnd_AddChild(w.Handle, hChild)
}

// 窗口_插入子对象.
//
// hChild: 要插入的对象句柄.
//
// index: 插入位置索引.
func (w *windowBase) InsertChild(hChild int, index int32) bool {
	return xc.XWnd_InsertChild(w.Handle, hChild, index)
}

// 窗口_取HWND.
func (w *windowBase) GetHWND() uintptr {
	return xc.XWnd_GetHWND(w.Handle)
}

// 窗口_重绘.
//
// bImmediate: 是否立即重绘, 通常为 false 即可.
func (w *windowBase) Redraw(bImmediate bool) *windowBase {
	xc.XWnd_Redraw(w.Handle, bImmediate)
	return w
}

// 窗口_重绘指定区域.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *windowBase) RedrawRect(pRect *xc.RECT, bImmediate bool) *windowBase {
	xc.XWnd_RedrawRect(w.Handle, pRect, bImmediate)
	return w
}

// 窗口_置坐标.
//
// pRect: 坐标.
func (w *windowBase) SetRect(pRect *xc.RECT) *windowBase {
	xc.XWnd_SetRect(w.Handle, pRect)
	return w
}

// 窗口_置焦点.
//
// hFocusEle: 将获得焦点的元素.
func (w *windowBase) SetFocusEle(hFocusEle int) *windowBase {
	xc.XWnd_SetFocusEle(w.Handle, hFocusEle)
	return w
}

// 窗口_取焦点.
func (w *windowBase) GetFocusEle() int {
	return xc.XWnd_GetFocusEle(w.Handle)
}

// 窗口_取鼠标停留元素.
func (w *windowBase) GetStayEle() int {
	return xc.XWnd_GetStayEle(w.Handle)
}

// 窗口_绘制, 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (w *windowBase) DrawWindow(hDraw int) *windowBase {
	xc.XWnd_DrawWindow(w.Handle, hDraw)
	return w
}

// 窗口_居中.
func (w *windowBase) Center() *windowBase {
	xc.XWnd_Center(w.Handle)
	return w
}

// 窗口_居中扩展.
//
// width: 窗口宽度.
//
// height: 窗口高度.
func (w *windowBase) CenterEx(width, height int32) *windowBase {
	xc.XWnd_CenterEx(w.Handle, width, height)
	return w
}

// 窗口_置光标.
//
// hCursor: 鼠标光标句柄.
func (w *windowBase) SetCursor(hCursor uintptr) *windowBase {
	xc.XWnd_SetCursor(w.Handle, hCursor)
	return w
}

// 窗口_取光标.
func (w *windowBase) GetCursor() uintptr {
	return xc.XWnd_GetCursor(w.Handle)
}

// 窗口_启用拖动边框.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragBorder(bEnable bool) *windowBase {
	xc.XWnd_EnableDragBorder(w.Handle, bEnable)
	return w
}

// 窗口_启用拖动窗口.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragWindow(bEnable bool) *windowBase {
	xc.XWnd_EnableDragWindow(w.Handle, bEnable)
	return w
}

// 窗口_启用拖动标题栏.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragCaption(bEnable bool) *windowBase {
	xc.XWnd_EnableDragCaption(w.Handle, bEnable)
	return w
}

// 窗口_启用绘制背景.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDrawBk(bEnable bool) *windowBase {
	xc.XWnd_EnableDrawBk(w.Handle, bEnable)
	return w
}

// 窗口_启用自动焦点. 当鼠标左键按下是否获得焦点.
//
// bEnable: 是否启用.
func (w *windowBase) EnableAutoFocus(bEnable bool) *windowBase {
	xc.XWnd_EnableAutoFocus(w.Handle, bEnable)
	return w
}

// 窗口_启用允许最大化.
//
// bEnable: 是否启用.
func (w *windowBase) EnableMaxWindow(bEnable bool) *windowBase {
	xc.XWnd_EnableMaxWindow(w.Handle, bEnable)
	return w
}

// 窗口_启用限制窗口大小.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLimitWindowSize(bEnable bool) *windowBase {
	xc.XWnd_EnableLimitWindowSize(w.Handle, bEnable)
	return w
}

// 窗口_启用布局.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLayout(bEnable bool) *windowBase {
	xc.XWnd_EnableLayout(w.Handle, bEnable)
	return w
}

// 窗口_启用布局覆盖边框.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLayoutOverlayBorder(bEnable bool) *windowBase {
	xc.XWnd_EnableLayoutOverlayBorder(w.Handle, bEnable)
	return w
}

// 窗口_显示布局边界.
//
// bEnable: 是否启用.
func (w *windowBase) ShowLayoutFrame(bEnable bool) *windowBase {
	xc.XWnd_ShowLayoutFrame(w.Handle, bEnable)
	return w
}

// 窗口_判断启用布局.
func (w *windowBase) IsEnableLayout() bool {
	return xc.XWnd_IsEnableLayout(w.Handle)
}

// 窗口_是否最大化.
func (w *windowBase) IsMaxWindow() bool {
	return xc.XWnd_IsMaxWindow(w.Handle)
}

// 窗口_置鼠标捕获元素.
//
// hEle: 元素句柄.
func (w *windowBase) SetCaptureEle(hEle int) *windowBase {
	xc.XWnd_SetCaptureEle(w.Handle, hEle)
	return w
}

// 窗口_取鼠标捕获元素.
func (w *windowBase) GetCaptureEle() int {
	return xc.XWnd_GetCaptureEle(w.Handle)
}

// 窗口_取绘制矩形.
//
// pRcPaint: 重绘区域坐标.
func (w *windowBase) GetDrawRect(pRcPaint *xc.RECT) *windowBase {
	xc.XWnd_GetDrawRect(w.Handle, pRcPaint)
	return w
}

// 窗口_取绘制矩形ex.
func (w *windowBase) GetDrawRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetDrawRect(w.Handle, &rc)
	return rc
}

// 窗口_置系统光标.
//
// hCursor: 光标句柄.
func (w *windowBase) SetCursorSys(hCursor uintptr) uintptr {
	return xc.XWnd_SetCursorSys(w.Handle, hCursor)
}

// 窗口_置字体.
//
// hFontx: 炫彩字体句柄.
func (w *windowBase) SetFont(hFontx int) *windowBase {
	xc.XWnd_SetFont(w.Handle, hFontx)
	return w
}

// 窗口_置文本颜色.
//
// color: xc.RGBA 颜色值.
func (w *windowBase) SetTextColor(color int) *windowBase {
	xc.XWnd_SetTextColor(w.Handle, color)
	return w
}

// 窗口_取文本颜色, 返回xc.RGBA 颜色.
func (w *windowBase) GetTextColor() int {
	return xc.XWnd_GetTextColor(w.Handle)
}

// 窗口_取文本颜色扩展, 返回xc.RGBA 颜色.
func (w *windowBase) GetTextColorEx() int {
	return xc.XWnd_GetTextColorEx(w.Handle)
}

// 窗口_置ID.
//
// nID: ID值.
func (w *windowBase) SetID(nID int32) *windowBase {
	xc.XWnd_SetID(w.Handle, nID)
	return w
}

// 窗口_取ID.
func (w *windowBase) GetID() int32 {
	return xc.XWnd_GetID(w.Handle)
}

// 窗口_置名称.
//
// pName: name值, 字符串.
func (w *windowBase) SetName(pName string) *windowBase {
	xc.XWnd_SetName(w.Handle, pName)
	return w
}

// 窗口_取名称.
func (w *windowBase) GetName() string {
	return xc.XWnd_GetName(w.Handle)
}

// 窗口_置边大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底部大小.
func (w *windowBase) SetBorderSize(left, top, right, bottom int32) *windowBase {
	xc.XWnd_SetBorderSize(w.Handle, left, top, right, bottom)
	return w
}

// 窗口_取边大小.
//
// pBorder: 返回边大小.
func (w *windowBase) GetBorderSize(pBorder *xc.RECT) *windowBase {
	xc.XWnd_GetBorderSize(w.Handle, pBorder)
	return w
}

// 窗口_取边大小ex.
func (w *windowBase) GetBorderSizeEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetBorderSize(w.Handle, &rc)
	return rc
}

// 窗口_置布局内填充大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (w *windowBase) SetPadding(left, top, right, bottom int32) *windowBase {
	xc.XWnd_SetPadding(w.Handle, left, top, right, bottom)
	return w
}

// 窗口_取布局内填充大小.
//
// pPadding: 返回布局内填充大小.
func (w *windowBase) GetPadding(pPadding *xc.RECT) *windowBase {
	xc.XWnd_GetPadding(w.Handle, pPadding)
	return w
}

// 窗口_取布局内填充大小ex.
func (w *windowBase) GetPaddingEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetPadding(w.Handle, &rc)
	return rc
}

// 窗口_置拖动边框大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底边大小.
func (w *windowBase) SetDragBorderSize(left, top, right, bottom int32) *windowBase {
	xc.XWnd_SetDragBorderSize(w.Handle, left, top, right, bottom)
	return w
}

// 窗口_取拖动边框大小.
//
// pSize: 拖动边框大小.
func (w *windowBase) GetDragBorderSize(pBorder *xc.RECT) *windowBase {
	xc.XWnd_GetDragBorderSize(w.Handle, pBorder)
	return w
}

// 窗口_取拖动边框大小ex.
func (w *windowBase) GetDragBorderSizeEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetDragBorderSize(w.Handle, &rc)
	return rc
}

// 窗口_置最小大小.
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *windowBase) SetMinimumSize(width, height int32) *windowBase {
	xc.XWnd_SetMinimumSize(w.Handle, width, height)
	return w
}

// 窗口_测试点击子元素. 成功则返回元素句柄.
//
// pPt: 左边点.
func (w *windowBase) HitChildEle(pPt *xc.POINT) int {
	return xc.XWnd_HitChildEle(w.Handle, pPt)
}

// 窗口_取子对象数量.
func (w *windowBase) GetChildCount() int32 {
	return xc.XWnd_GetChildCount(w.Handle)
}

// 窗口_取子对象从索引.
//
// index: 元素索引.
func (w *windowBase) GetChildByIndex(index int32) int {
	return xc.XWnd_GetChildByIndex(w.Handle, index)
}

// 窗口_取子对象从ID.
//
// nID: 元素ID, ID必须大于0.
func (w *windowBase) GetChildByID(nID int32) int {
	return xc.XWnd_GetChildByID(w.Handle, nID)
}

// 窗口_取子对象.
//
// nID: 对象ID,ID必须大于0.
func (w *windowBase) GetChild(nID int32) int {
	return xc.XWnd_GetChild(w.Handle, nID)
}

// 窗口_关闭.
func (w *windowBase) CloseWindow() *windowBase {
	xc.XWnd_CloseWindow(w.Handle)
	return w
}

// 窗口_调整布局.
func (w *windowBase) AdjustLayout() *windowBase {
	xc.XWnd_AdjustLayout(w.Handle)
	return w
}

// 窗口_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
func (w *windowBase) AdjustLayoutEx(nFlags xcc.AdjustLayout_) *windowBase {
	xc.XWnd_AdjustLayoutEx(w.Handle, nFlags)
	return w
}

// 窗口_创建插入符.
//
// hEle: 元素句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func (w *windowBase) CreateCaret(hEle int, x, y, width, height int32) *windowBase {
	xc.XWnd_CreateCaret(w.Handle, hEle, x, y, width, height)
	return w
}

// 窗口_置插入符位置, 设置插入符位置.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
//
// bUpdate: 是否立即更新UI.
func (w *windowBase) SetCaretPos(x, y, width, height int32, bUpdate bool) *windowBase {
	xc.XWnd_SetCaretPos(w.Handle, x, y, width, height, bUpdate)
	return w
}

// 窗口_置插入符颜色.
//
// color: 颜色值, xc.RGBA 颜色.
func (w *windowBase) SetCaretColor(color int) *windowBase {
	xc.XWnd_SetCaretColor(w.Handle, color)
	return w
}

// 窗口_显示插入符.
//
// bShow: 是否显示.
func (w *windowBase) ShowCaret(bShow bool) *windowBase {
	xc.XWnd_ShowCaret(w.Handle, bShow)
	return w
}

// 窗口_销毁插入符.
func (w *windowBase) DestroyCaret() *windowBase {
	xc.XWnd_DestroyCaret(w.Handle)
	return w
}

// 窗口_取插入符元素.
func (w *windowBase) GetCaretHELE() int {
	return xc.XWnd_GetCaretHELE(w.Handle)
}

// 窗口_取客户区坐标.
//
// pRect: 坐标.
func (w *windowBase) GetClientRect(pRect *xc.RECT) bool {
	return xc.XWnd_GetClientRect(w.Handle, pRect)
}

// 窗口_取客户区坐标ex.
func (w *windowBase) GetClientRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetClientRect(w.Handle, &rc)
	return rc
}

// 窗口_取Body坐标.
//
// pRect: 坐标.
func (w *windowBase) GetBodyRect(pRect *xc.RECT) *windowBase {
	xc.XWnd_GetBodyRect(w.Handle, pRect)
	return w
}

// 窗口_取Body坐标ex.
func (w *windowBase) GetBodyRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetBodyRect(w.Handle, &rc)
	return rc
}

// 窗口_取布局坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) GetLayoutRect(pRect *xc.RECT) *windowBase {
	xc.XWnd_GetLayoutRect(w.Handle, pRect)
	return w
}

// 窗口_取布局坐标ex.
func (w *windowBase) GetLayoutRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetLayoutRect(w.Handle, &rc)
	return rc
}

// 窗口_移动窗口.
//
// x: X坐标.
//
// y: Y坐标.
func (w *windowBase) SetPosition(x, y int32) *windowBase {
	xc.XWnd_SetPosition(w.Handle, x, y)
	return w
}

// 窗口_取坐标.
//
// pRect: 坐标.
func (w *windowBase) GetRect(pRect *xc.RECT) *windowBase {
	xc.XWnd_GetRect(w.Handle, pRect)
	return w
}

// 窗口_取坐标ex.
func (w *windowBase) GetRectEx() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc
}

// 窗口_取坐标ex. 返回的坐标已经根据窗口dpi进行换算了, 换算后的结果就是屏幕上显示的真实坐标了.
func (w *windowBase) GetRectDPI() xc.RECT {
	rc := xc.RECT{}
	xc.XWnd_GetRect(w.Handle, &rc)
	dpi := xc.XWnd_GetDPI(w.Handle)
	rc.Left = xc.DpiConvRound(dpi, rc.Left)
	rc.Top = xc.DpiConvRound(dpi, rc.Top)
	rc.Right = xc.DpiConvRound(dpi, rc.Right)
	rc.Bottom = xc.DpiConvRound(dpi, rc.Bottom)
	return rc
}

// 窗口_最大化.
//
// bMaximize: 是否最大化.
func (w *windowBase) MaxWindow(bMaximize bool) *windowBase {
	xc.XWnd_MaxWindow(w.Handle, bMaximize)
	return w
}

// 窗口_置定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetTimer(nIDEvent, uElapse uint32) uint32 {
	return xc.XWnd_SetTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) KillTimer(nIDEvent uint32) bool {
	return xc.XWnd_KillTimer(w.Handle, nIDEvent)
}

// 窗口_置炫彩定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetXCTimer(nIDEvent, uElapse uint32) bool {
	return xc.XWnd_SetXCTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭炫彩定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) KillXCTimer(nIDEvent uint32) bool {
	return xc.XWnd_KillXCTimer(w.Handle, nIDEvent)
}

// 窗口_取背景管理器.
func (w *windowBase) GetBkManager() int {
	return xc.XWnd_GetBkManager(w.Handle)
}

// 窗口_取背景管理器扩展.
func (w *windowBase) GetBkManagerEx() int {
	return xc.XWnd_GetBkManagerEx(w.Handle)
}

// 窗口_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (w *windowBase) SetBkMagager(hBkInfoM int) *windowBase {
	xc.XWnd_SetBkMagager(w.Handle, hBkInfoM)
	return w
}

// 窗口_置透明类型.
//
// nType: 窗口透明类型: xcc.Window_Transparent_.
func (w *windowBase) SetTransparentType(nType xcc.Window_Transparent_) *windowBase {
	xc.XWnd_SetTransparentType(w.Handle, nType)
	return w
}

// 窗口_置透明度.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *windowBase) SetTransparentAlpha(alpha byte) *windowBase {
	xc.XWnd_SetTransparentAlpha(w.Handle, alpha)
	return w
}

// 窗口_置透明色.
//
// color: 窗口透明色, xc.RGBA 颜色.
func (w *windowBase) SetTransparentColor(color int) *windowBase {
	xc.XWnd_SetTransparentColor(w.Handle, color)
	return w
}

// 窗口_置阴影信息.
//
// nSize: 阴影大小.
//
// nDepth: 阴影深度, 0-255.
//
// nAngeleSize: 圆角阴影内收大小.
//
// bRightAngle: 是否强制直角.
//
// color: 阴影颜色, xc.RGBA 颜色.
func (w *windowBase) SetShadowInfo(nSize int32, nDepth int32, nAngeleSize int32, bRightAngle bool, color int) *windowBase {
	xc.XWnd_SetShadowInfo(w.Handle, nSize, nDepth, nAngeleSize, bRightAngle, color)
	return w
}

// 窗口_取阴影信息.
//
// pnSize: 阴影大小.
//
// pnDepth: 阴影深度.
//
// pnAngeleSize: 圆角阴影内收大小 .
//
// pbRightAngle: 是否强制直角.
//
// pColor: 阴影颜色, xc.RGBA 颜色.
func (w *windowBase) GetShadowInfo(pnSize, pnDepth, pnAngeleSize *int32, pbRightAngle *bool, pColor *int) *windowBase {
	xc.XWnd_GetShadowInfo(w.Handle, pnSize, pnDepth, pnAngeleSize, pbRightAngle, pColor)
	return w
}

// 窗口_取透明类型, 返回: xcc.Window_Transparent_.
func (w *windowBase) GetTransparentType() xcc.Window_Transparent_ {
	return xc.XWnd_GetTransparentType(w.Handle)
}

// 窗口_启用拖放文件.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragFiles(bEnable bool) *windowBase {
	xc.XWnd_EnableDragFiles(w.Handle, bEnable)
	return w
}

// 窗口_显示 显示隐藏窗口.
//
// bShow: 是否显示.
func (w *windowBase) Show(bShow bool) *windowBase {
	xc.XWnd_Show(w.Handle, bShow)
	return w
}

// 窗口_取插入符信息, 获取插入符信息, 返回: 插入符元素句柄.
//
// hWindow: 窗口句柄.
//
// pX: 接收返回x坐标.
//
// pY: 接收返回y坐标.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (w *windowBase) GetCaretInfo(pX, pY, pWidth, pHeight *int32) int {
	return xc.XWnd_GetCaretInfo(w.Handle, pX, pY, pWidth, pHeight)
}

// 窗口_置图标.
//
// hImage: 图标句柄.
func (w *windowBase) SetIcon(hImage int) *windowBase {
	xc.XWnd_SetIcon(w.Handle, hImage)
	return w
}

// 窗口_置标题.
//
// pTitle: 标题文本.
func (w *windowBase) SetTitle(pTitle string) *windowBase {
	xc.XWnd_SetTitle(w.Handle, pTitle)
	return w
}

// 窗口_置标题颜色.
//
// color: xc.RGBA 颜色.
func (w *windowBase) SetTitleColor(color int) *windowBase {
	xc.XWnd_SetTitleColor(w.Handle, color)
	return w
}

// 窗口_取控制按钮, 返回按钮句柄.
//
// nFlag: xcc.Window_Style_ . 可用值: xcc.Window_Style_Btn_Min , xcc.Window_Style_Btn_Max , xcc.Window_Style_Btn_Close .
func (w *windowBase) GetButton(nFlag xcc.Window_Style_) int {
	return xc.XWnd_GetButton(w.Handle, nFlag)
}

// 窗口_取图标, 返回图标句柄.
func (w *windowBase) GetIcon() int {
	return xc.XWnd_GetIcon(w.Handle)
}

// 窗口_取标题, 返回标题文本.
func (w *windowBase) GetTitle() string {
	return xc.XWnd_GetTitle(w.Handle)
}

// 窗口_取标题颜色, 返回xc.RGBA 颜色.
func (w *windowBase) GetTitleColor() int {
	return xc.XWnd_GetTitleColor(w.Handle)
}

// 窗口_添加背景边框.
//
// nState: 组合状态.
//
// color: xc.RGBA 颜色.
//
// width: 线宽.
func (w *windowBase) AddBkBorder(nState xcc.Window_State_Flag_, color int, width int32) *windowBase {
	xc.XWnd_AddBkBorder(w.Handle, nState, color, width)
	return w
}

// 窗口_添加背景填充.
//
// nState: 组合状态.
//
// color: xc.RGBA 颜色.
func (w *windowBase) AddBkFill(nState xcc.Window_State_Flag_, color int) *windowBase {
	xc.XWnd_AddBkFill(w.Handle, nState, color)
	return w
}

// 窗口_添加背景图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (w *windowBase) AddBkImage(nState xcc.Window_State_Flag_, hImage int) *windowBase {
	xc.XWnd_AddBkImage(w.Handle, nState, hImage)
	return w
}

// 窗口_取背景对象数量.
func (w *windowBase) GetBkInfoCount() int32 {
	return xc.XWnd_GetBkInfoCount(w.Handle)
}

// 窗口_清空背景对象.
func (w *windowBase) ClearBkInfo() *windowBase {
	xc.XWnd_ClearBkInfo(w.Handle)
	return w
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
func (w *windowBase) NotifyMsg_WindowPopup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	return xc.XNotifyMsg_WindowPopup(w.Handle, position, pTitle, pText, hIcon, skin)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func (w *windowBase) NotifyMsg_WindowPopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	return xc.XNotifyMsg_WindowPopupEx(w.Handle, position, pTitle, pText, hIcon, skin, bBtnClose, bAutoClose, nWidth, nHeight)
}

// 通知消息_置持续时间.
//
// duration: 持续时间.
func (w *windowBase) NotifyMsg_SetDuration(duration uint32) *windowBase {
	xc.XNotifyMsg_SetDuration(w.Handle, duration)
	return w
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func (w *windowBase) NotifyMsg_SetParentMargin(left, top, right, bottom int32) *windowBase {
	xc.XNotifyMsg_SetParentMargin(w.Handle, left, top, right, bottom)
	return w
}

// 通知消息_置标题高度.
//
// nHeight: 高度.
func (w *windowBase) NotifyMsg_SetCaptionHeight(nHeight int32) *windowBase {
	xc.XNotifyMsg_SetCaptionHeight(w.Handle, nHeight)
	return w
}

// 通知消息_置宽度, 设置默认宽度.
//
// nWidth: 宽度.
func (w *windowBase) NotifyMsg_SetWidth(nWidth int32) *windowBase {
	xc.XNotifyMsg_SetWidth(w.Handle, nWidth)
	return w
}

// 通知消息_置间隔.
//
// nSpace: 间隔大小.
func (w *windowBase) NotifyMsg_SetSpace(nSpace int32) *windowBase {
	xc.XNotifyMsg_SetSpace(w.Handle, nSpace)
	return w
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func (w *windowBase) NotifyMsg_SetBorderSize(left, top, right, bottom int32) *windowBase {
	xc.XNotifyMsg_SetBorderSize(w.Handle, left, top, right, bottom)
	return w
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (w *windowBase) SetBkInfo(pText string) int32 {
	return xc.XWnd_SetBkInfo(w.Handle, pText)
}

// 窗口_是否可拖动标题栏.
func (w *windowBase) IsDragCaption() bool {
	return xc.XWnd_IsDragCaption(w.Handle)
}

// 窗口_是否可拖动窗口.
func (w *windowBase) IsDragWindow() bool {
	return xc.XWnd_IsDragWindow(w.Handle)
}

// 窗口_是否可拖动边框.
func (w *windowBase) IsDragBorder() bool {
	return xc.XWnd_IsDragBorder(w.Handle)
}

// 窗口_置标题外间距, 设置标题内容(图标, 标题, 控制按钮)外间距.
//
// left: 左边间距.
//
// top: 上边间距.
//
// right: 右边间距.
//
// bottom: 下边间距.
func (w *windowBase) SetCaptionMargin(left, top, right, bottom int32) *windowBase {
	xc.XWnd_SetCaptionMargin(w.Handle, left, top, right, bottom)
	return w
}

// SetTopEx 窗口_置顶Ex.
//
// b: 是否置顶.
func (w *windowBase) SetTopEx(b bool) bool {
	return wnd.SetTop(w.GetHWND(), b)
}

// 窗口_置窗口位置. 封装系统API SetWindowPos(), 内部做了DPI适配.
//
// hWndInsertAfter: 在Z序中位于定位窗口之前的窗口句柄. 此参数必须是窗口HWND或以下值之一: xcc.HWND_.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// uFlags: 窗口大小调整和定位标志. 可以是以下值的组合: xcc.SWP_.
func (w *windowBase) SetWindowPos(hWndInsertAfter xcc.HWND_, x, y, cx, cy int32, uFlags xcc.SWP_) bool {
	return xc.XWnd_SetWindowPos(w.Handle, hWndInsertAfter, x, y, cx, cy, uFlags)
}

// 窗口_取DPI. 获取当前窗口所在显示器DPI, 返回窗口DPI.
func (w *windowBase) GetDPI() int32 {
	return xc.XWnd_GetDPI(w.Handle)
}

// 窗口_坐标转换DPI. 窗口客户区坐标转换到缩放后DPI坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) RectToDPI(pRect *xc.RECT) *windowBase {
	xc.XWnd_RectToDPI(w.Handle, pRect)
	return w
}

// 窗口_坐标点转换DPI. 窗口客户区坐标点转换到缩放后.
//
// pPt: 接收返回坐标点.
func (w *windowBase) PointToDPI(pPt *xc.POINT) *windowBase {
	xc.XWnd_PointToDPI(w.Handle, pPt)
	return w
}

// 窗口_取光标位置. 封装的系统API: GetCursorPos(), 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) GetCursorPos(pPt *xc.POINT) bool {
	return xc.XWnd_GetCursorPos(w.Handle, pPt)
}

// 窗口_客户区坐标点到屏幕坐标点. 封装的系统API: ClientToScreen, 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) ClientToScreen(pPt *xc.POINT) bool {
	return xc.XWnd_ClientToScreen(w.Handle, pPt)
}

// 窗口_屏幕坐标点到客户区坐标点. 封装的系统API: ScreenToClient(), 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) ScreenToClient(pPt *xc.POINT) bool {
	return xc.XWnd_ScreenToClient(w.Handle, pPt)
}

// 窗口_置DPI. 设置当前窗口DPI, 默认DPI为96.
//
// nDPI: DPI值.
func (w *windowBase) SetDPI(nDPI int32) *windowBase {
	xc.XWnd_SetDPI(w.Handle, nDPI)
	return w
}

// 窗口_销毁.
func (w *windowBase) DestroyWindow() *windowBase {
	xc.XWnd_DestroyWindow(w.Handle)
	return w
}

// 窗口_置大小. 设置窗口宽高.
//
// 宽: width.
//
// 高: height.
func (w *windowBase) SetSize(width, height int32) *windowBase {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Right = rc.Left + width
	rc.Bottom = rc.Top + height
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// 窗口_置宽度.
//
// 宽: width.
func (w *windowBase) SetWidth(width int32) *windowBase {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Right = rc.Left + width
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// 窗口_置高度.
//
// 高: height.
func (w *windowBase) SetHeight(height int32) *windowBase {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Bottom = rc.Top + height
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// 窗口_取宽度.
func (w *windowBase) GetWidth() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Right - rc.Left
}

// 窗口_取高度.
func (w *windowBase) GetHeight() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Bottom - rc.Top
}

// 窗口_取左边.
func (w *windowBase) GetLeft() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Left
}

// 窗口_取顶边.
func (w *windowBase) GetTop() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Top
}

// 窗口_取右边.
func (w *windowBase) GetRight() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Right
}

// 窗口_取底边.
func (w *windowBase) GetBottom() int32 {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	return rc.Bottom
}

// 窗口_置左边.
//
// x: 左边x坐标.
func (w *windowBase) SetLeft(x int32) *windowBase {
	rc := w.GetRectEx()
	width := rc.Right - rc.Left
	rc.Left = x
	rc.Right = rc.Left + width
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// 窗口_置顶边.
//
// y: 顶边y坐标.
func (w *windowBase) SetTopEdge(y int32) *windowBase {
	rc := w.GetRectEx()
	height := rc.Bottom - rc.Top
	rc.Top = y
	rc.Bottom = rc.Top + height
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// 窗口_置左边和顶边.
//
// x: 左边x坐标.
//
// y: 顶边y坐标.
func (w *windowBase) SetLeftAndTop(x, y int32) *windowBase {
	rc := w.GetRectEx()
	width := rc.Right - rc.Left
	height := rc.Bottom - rc.Top
	rc.Left = x
	rc.Top = y
	rc.Right = rc.Left + width
	rc.Bottom = rc.Top + height
	xc.XWnd_SetRect(w.Handle, &rc)
	return w
}

// CreateTrayIcon 创建托盘图标对象. 之后需调用托盘图标对象的Show函数来显示到托盘.
//
// hIcon: 图标句柄. 可使用 wutil.HIcon 函数生成.
//
// tips: 提示文本, 长度不能超过 127 个字符.
func (w *windowBase) CreateTrayIcon(hIcon uintptr, tips string) *TrayIcon {
	return NewTrayIcon(w.Handle, hIcon, tips)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (w *windowBase) EnableHorizon(bEnable bool) *windowBase {
	xc.XLayoutBox_EnableHorizon(w.Handle, bEnable)
	return w
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (w *windowBase) EnableAutoWrap(bEnable bool) *windowBase {
	xc.XLayoutBox_EnableAutoWrap(w.Handle, bEnable)
	return w
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (w *windowBase) EnableOverflowHide(bEnable bool) *windowBase {
	xc.XLayoutBox_EnableOverflowHide(w.Handle, bEnable)
	return w
}

// SetAlignH 布局盒子_置水平对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (w *windowBase) SetAlignH(nAlign xcc.Layout_Align_) *windowBase {
	xc.XLayoutBox_SetAlignH(w.Handle, nAlign)
	return w
}

// SetAlignV 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (w *windowBase) SetAlignV(nAlign xcc.Layout_Align_) *windowBase {
	xc.XLayoutBox_SetAlignV(w.Handle, nAlign)
	return w
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func (w *windowBase) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) *windowBase {
	xc.XLayoutBox_SetAlignBaseline(w.Handle, nAlign)
	return w
}

// SetSpace 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (w *windowBase) SetSpace(nSpace int32) *windowBase {
	xc.XLayoutBox_SetSpace(w.Handle, nSpace)
	return w
}

// SetSpaceRow 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (w *windowBase) SetSpaceRow(nSpace int32) *windowBase {
	xc.XLayoutBox_SetSpaceRow(w.Handle, nSpace)
	return w
}

// DpiConv 将 int32 类型的整数根据窗口 dpi 进行换算.
//   - 开启自动 DPI 后, 你可能感觉到一些函数获取的坐标不对了, 因为你在用高分辨率屏幕, 然后系统里会有个缩放的设置, 可能是 150%, 这时获取到的坐标应该乘以 1.5 才是屏幕上显示的真实的坐标.
//
// i: int32 类型的整数.
func (w *windowBase) DpiConv(i int32) int32 {
	return xc.DpiConv(w.GetDPI(), i)
}

// DpiConvRound 将 int32 类型的整数根据窗口 dpi 进行换算, 计算结果会四舍五入.
//   - 开启自动 DPI 后, 你可能感觉到一些函数获取的坐标不对了, 因为你在用高分辨率屏幕, 然后系统里会有个缩放的设置, 可能是 150%, 这时获取到的坐标应该乘以 1.5 才是屏幕上显示的真实的坐标.
//
// i: int32 类型的整数.
func (w *windowBase) DpiConvRound(i int32) int32 {
	return xc.DpiConvRound(w.GetDPI(), i)
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_NCDestroy 添加窗口非客户区销毁事件. 在窗口销毁之后触发.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_NCDestroy(fun WM_NCDESTROY1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_NCDESTROY, xc.OnWM_NCDESTROY, fun, allowAddingMultiple...)
}

// AddEvent_DockPopup 添加框架窗口码头弹出窗格事件. 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_DockPopup(fun XWM_DOCK_POPUP1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_DOCK_POPUP, onXWM_DOCK_POPUP, fun, allowAddingMultiple...)
}

// onXWM_DOCK_POPUP 框架窗口码头弹出窗格事件.
func onXWM_DOCK_POPUP(hWindow int, hWindowDock, hPane int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_DOCK_POPUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_DOCK_POPUP1)(hWindow, hWindowDock, hPane, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_WindProc 添加窗口消息过程事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_WindProc(fun XWM_WINDPROC1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_WINDPROC, onXWM_WINDPROC, fun, allowAddingMultiple...)
}

// onXWM_WINDPROC 窗口消息过程事件.
func onXWM_WINDPROC(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_WINDPROC)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_WINDPROC1)(hWindow, message, wParam, lParam, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_XC_Timer 添加炫彩定时器事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_XC_Timer(fun XWM_XC_TIMER1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_XC_TIMER, onXWM_XC_TIMER, fun, allowAddingMultiple...)
}

// onXWM_XC_TIMER 炫彩定时器事件.
func onXWM_XC_TIMER(hWindow int, nIDEvent uint, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_XC_TIMER)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_XC_TIMER1)(hWindow, nIDEvent, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetFocus_Ele 添加窗口置焦点元素事件. 指定元素获得焦点.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_SetFocus_Ele(fun XWM_SETFOCUS_ELE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_SETFOCUS_ELE, onXWM_SETFOCUS_ELE, fun, allowAddingMultiple...)
}

// onXWM_SETFOCUS_ELE 窗口置焦点元素事件. 指定元素获得焦点.
func onXWM_SETFOCUS_ELE(hWindow int, hEle int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_SETFOCUS_ELE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_SETFOCUS_ELE1)(hWindow, hEle, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Float_Pane 添加浮动窗格事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Float_Pane(fun XWM_FLOAT_PANE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_FLOAT_PANE, onXWM_FLOAT_PANE, fun, allowAddingMultiple...)
}

// onXWM_FLOAT_PANE 浮动窗格事件.
func onXWM_FLOAT_PANE(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_FLOAT_PANE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_FLOAT_PANE1)(hWindow, hFloatWnd, hPane, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint_End 添加窗口绘制完成事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Paint_End(fun XWM_PAINT_END1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_PAINT_END, onXWM_PAINT_END, fun, allowAddingMultiple...)
}

// onXWM_PAINT_END 窗口绘制完成事件.
func onXWM_PAINT_END(hWindow int, hDraw int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_PAINT_END)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_PAINT_END1)(hWindow, hDraw, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint_Display 添加窗口绘制完成并且已经显示到屏幕事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Paint_Display(fun XWM_PAINT_DISPLAY1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_PAINT_DISPLAY, onXWM_PAINT_DISPLAY, fun, allowAddingMultiple...)
}

// onXWM_PAINT_DISPLAY 窗口绘制完成并且已经显示到屏幕事件.
func onXWM_PAINT_DISPLAY(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_PAINT_DISPLAY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_PAINT_DISPLAY1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_BodyView_Rect 添加框架窗口主视图坐标改变事件. 如果主视图没有绑定元素, 那么当坐标改变时触发此事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_BodyView_Rect(fun XWM_BODYVIEW_RECT1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_BODYVIEW_RECT, onXWM_BODYVIEW_RECT, fun, allowAddingMultiple...)
}

// onXWM_BODYVIEW_RECT 框架窗口主视图坐标改变事件.
func onXWM_BODYVIEW_RECT(hWindow int, width, height int32, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_BODYVIEW_RECT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_BODYVIEW_RECT1)(hWindow, width, height, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_FloatWnd_Drag 添加浮动窗口拖动事件. 用户拖动浮动窗口移动, 显示停靠提示.
//   - hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_FloatWnd_Drag(fun XWM_FLOATWND_DRAG1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_FLOATWND_DRAG, onXWM_FLOATWND_DRAG, fun, allowAddingMultiple...)
}

// onXWM_FLOATWND_DRAG 浮动窗口拖动事件.
func onXWM_FLOATWND_DRAG(hWindow int, hFloatWnd int, hArray *[6]int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_FLOATWND_DRAG)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_FLOATWND_DRAG1)(hWindow, hFloatWnd, hArray, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Paint 添加窗口绘制事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Paint(fun WM_PAINT1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_PAINT, onWM_PAINT, fun, allowAddingMultiple...)
}

// onWM_PAINT 窗口绘制事件.
func onWM_PAINT(hWindow int, hDraw int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_PAINT)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_PAINT1)(hWindow, hDraw, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Close 添加窗口关闭事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Close(fun WM_CLOSE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_CLOSE, onWM_CLOSE, fun, allowAddingMultiple...)
}

// onWM_CLOSE 窗口关闭事件.
func onWM_CLOSE(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_CLOSE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_CLOSE1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Destroy 添加窗口销毁事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Destroy(fun WM_DESTROY1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_DESTROY, onWM_DESTROY, fun, allowAddingMultiple...)
}

// onWM_DESTROY 窗口销毁事件.
func onWM_DESTROY(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_DESTROY)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_DESTROY1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseMove 添加窗口鼠标移动事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_MouseMove(fun WM_MOUSEMOVE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_MOUSEMOVE, onWM_MOUSEMOVE, fun, allowAddingMultiple...)
}

// onWM_MOUSEMOVE 窗口鼠标移动事件.
func onWM_MOUSEMOVE(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_MOUSEMOVE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_MOUSEMOVE1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonDown 添加窗口鼠标左键按下事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_LButtonDown(fun WM_LBUTTONDOWN1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_LBUTTONDOWN, onWM_LBUTTONDOWN, fun, allowAddingMultiple...)
}

// onWM_LBUTTONDOWN 窗口鼠标左键按下事件.
func onWM_LBUTTONDOWN(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_LBUTTONDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_LBUTTONDOWN1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonUp 添加窗口鼠标左键弹起事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_LButtonUp(fun WM_LBUTTONUP1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_LBUTTONUP, onWM_LBUTTONUP, fun, allowAddingMultiple...)
}

// onWM_LBUTTONUP 窗口鼠标左键弹起事件.
func onWM_LBUTTONUP(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_LBUTTONUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_LBUTTONUP1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_RButtonDown 添加窗口鼠标右键按下事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_RButtonDown(fun WM_RBUTTONDOWN1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_RBUTTONDOWN, onWM_RBUTTONDOWN, fun, allowAddingMultiple...)
}

// onWM_RBUTTONDOWN 窗口鼠标右键按下事件.
func onWM_RBUTTONDOWN(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_RBUTTONDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_RBUTTONDOWN1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_RButtonUp 添加窗口鼠标右键弹起事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_RButtonUp(fun WM_RBUTTONUP1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_RBUTTONUP, onWM_RBUTTONUP, fun, allowAddingMultiple...)
}

// onWM_RBUTTONUP 窗口鼠标右键弹起事件.
func onWM_RBUTTONUP(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_RBUTTONUP)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_RBUTTONUP1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_LButtonDBClick 添加窗口鼠标左键双击事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_LButtonDBClick(fun WM_LBUTTONDBLCLK1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_LBUTTONDBLCLK, onWM_LBUTTONDBLCLK, fun, allowAddingMultiple...)
}

// onWM_LBUTTONDBLCLK 窗口鼠标左键双击事件.
func onWM_LBUTTONDBLCLK(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_LBUTTONDBLCLK)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_LBUTTONDBLCLK1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_RButtonDBClick 添加窗口鼠标右键双击事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_RButtonDBClick(fun WM_RBUTTONDBLCLK1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_RBUTTONDBLCLK, onWM_RBUTTONDBLCLK, fun, allowAddingMultiple...)
}

// onWM_RBUTTONDBLCLK 窗口鼠标右键双击事件.
func onWM_RBUTTONDBLCLK(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_RBUTTONDBLCLK)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_RBUTTONDBLCLK1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseWheel 添加窗口鼠标滚轮滚动事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_MouseWheel(fun WM_MOUSEWHEEL1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_MOUSEWHEEL, onWM_MOUSEWHEEL, fun, allowAddingMultiple...)
}

// onWM_MOUSEWHEEL 窗口鼠标滚轮滚动事件.
func onWM_MOUSEWHEEL(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_MOUSEWHEEL)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_MOUSEWHEEL1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ExitSizeMove 添加窗口退出移动或调整大小模式循环事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_ExitSizeMove(fun WM_EXITSIZEMOVE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_EXITSIZEMOVE, onWM_EXITSIZEMOVE, fun, allowAddingMultiple...)
}

// onWM_EXITSIZEMOVE 窗口退出移动或调整大小模式循环事件.
func onWM_EXITSIZEMOVE(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_EXITSIZEMOVE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_EXITSIZEMOVE1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseHover 添加窗口鼠标进入事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_MouseHover(fun WM_MOUSEHOVER1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_MOUSEHOVER, onWM_MOUSEHOVER, fun, allowAddingMultiple...)
}

// onWM_MOUSEHOVER 窗口鼠标进入事件.
func onWM_MOUSEHOVER(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_MOUSEHOVER)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_MOUSEHOVER1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_MouseLeave 添加窗口鼠标离开事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_MouseLeave(fun WM_MOUSELEAVE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_MOUSELEAVE, onWM_MOUSELEAVE, fun, allowAddingMultiple...)
}

// onWM_MOUSELEAVE 窗口鼠标离开事件.
func onWM_MOUSELEAVE(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_MOUSELEAVE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_MOUSELEAVE1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Size 添加窗口大小改变事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Size(fun WM_SIZE1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_SIZE, onWM_SIZE, fun, allowAddingMultiple...)
}

// onWM_SIZE 窗口大小改变事件.
func onWM_SIZE(hWindow int, nFlags uint, pPt *xc.SIZE, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_SIZE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_SIZE1)(hWindow, nFlags, pPt, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Timer 添加窗口定时器事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Timer(fun WM_TIMER1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_TIMER, onWM_TIMER, fun, allowAddingMultiple...)
}

// onWM_TIMER 窗口定时器事件.
func onWM_TIMER(hWindow int, nIDEvent uint, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_TIMER)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_TIMER1)(hWindow, nIDEvent, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetFocus 添加窗口获得焦点事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_SetFocus(fun WM_SETFOCUS1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_SETFOCUS, onWM_SETFOCUS, fun, allowAddingMultiple...)
}

// onWM_SETFOCUS 窗口获得焦点事件.
func onWM_SETFOCUS(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_SETFOCUS)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_SETFOCUS1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KillFocus 添加窗口失去焦点事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_KillFocus(fun WM_KILLFOCUS1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_KILLFOCUS, onWM_KILLFOCUS, fun, allowAddingMultiple...)
}

// onWM_KILLFOCUS 窗口失去焦点事件.
func onWM_KILLFOCUS(hWindow int, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_KILLFOCUS)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_KILLFOCUS1)(hWindow, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_KeyDown 添加窗口键盘按键事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_KeyDown(fun WM_KEYDOWN1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_KEYDOWN, onWM_KEYDOWN, fun, allowAddingMultiple...)
}

// onWM_KEYDOWN 窗口键盘按键事件.
func onWM_KEYDOWN(hWindow int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_KEYDOWN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_KEYDOWN1)(hWindow, wParam, lParam, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_CaptureChanged 添加窗口鼠标捕获改变事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_CaptureChanged(fun WM_CAPTURECHANGED1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_CAPTURECHANGED, onWM_CAPTURECHANGED, fun, allowAddingMultiple...)
}

// onWM_CAPTURECHANGED 窗口鼠标捕获改变事件.
func onWM_CAPTURECHANGED(hWindow int, hWnd uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_CAPTURECHANGED)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_CAPTURECHANGED1)(hWindow, hWnd, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_SetCursor 添加窗口设置鼠标光标事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_SetCursor(fun WM_SETCURSOR1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_SETCURSOR, onWM_SETCURSOR, fun, allowAddingMultiple...)
}

// onWM_SETCURSOR 窗口设置鼠标光标事件.
func onWM_SETCURSOR(hWindow int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_SETCURSOR)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_SETCURSOR1)(hWindow, wParam, lParam, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Char 添加窗口字符事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Char(fun WM_CHAR1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_CHAR, onWM_CHAR, fun, allowAddingMultiple...)
}

// onWM_CHAR 窗口字符事件.
func onWM_CHAR(hWindow int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_CHAR)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_CHAR1)(hWindow, wParam, lParam, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_DropFiles 添加拖动文件到窗口事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_DropFiles(fun WM_DROPFILES1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.WM_DROPFILES, onWM_DROPFILES, fun, allowAddingMultiple...)
}

// onWM_DROPFILES 拖动文件到窗口事件.
func onWM_DROPFILES(hWindow int, hDropInfo uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.WM_DROPFILES)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(WM_DROPFILES1)(hWindow, hDropInfo, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_TrayIcon 添加托盘图标事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_TrayIcon(fun XWM_TRAYICON, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_TRAYICON, onXWM_TRAYICON, fun, allowAddingMultiple...)
}

// onXWM_TRAYICON 托盘图标事件.
func onXWM_TRAYICON(hWindow int, wParam, lParam uintptr, pbHandled *bool) int {
	cbs := xc.WndEventHandler.GetCallBacks(hWindow, xcc.XWM_TRAYICON)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XWM_TRAYICON)(wParam, lParam, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// ---------------------------------- 菜单事件 ----------------------------------

// AddEvent_Menu_Popup 添加菜单弹出事件.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (w *windowBase) AddEvent_Menu_Popup(fun XWM_MENU_POPUP1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_POPUP, xc.OnXWM_MENU_POPUP, fun, allowAddingMultiple...)
}

// 菜单弹出窗口事件.
func (w *windowBase) AddEvent_Menu_PopupWnd(fun XWM_MENU_POPUP_WND1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_POPUP_WND, xc.OnXWM_MENU_POPUP_WND, fun, allowAddingMultiple...)
}

// 菜单选择事件.
func (w *windowBase) AddEvent_Menu_Select(fun XWM_MENU_SELECT1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_SELECT, xc.OnXWM_MENU_SELECT, fun, allowAddingMultiple...)
}

// 菜单退出事件.
func (w *windowBase) AddEvent_Menu_Exit(fun XWM_MENU_EXIT1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_EXIT, xc.OnXWM_MENU_EXIT, fun, allowAddingMultiple...)
}

// 绘制菜单背景事件.
func (w *windowBase) AddEvent_Menu_DrawBackground(fun XWM_MENU_DRAW_BACKGROUND1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_DRAW_BACKGROUND, xc.OnXWM_MENU_DRAW_BACKGROUND, fun, allowAddingMultiple...)
}

// 绘制菜单项事件.
func (w *windowBase) AddEvent_Menu_DrawItem(fun XWM_MENU_DRAWITEM1, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(w.Handle, xcc.XWM_MENU_DRAWITEM, xc.OnXWM_MENU_DRAWITEM, fun, allowAddingMultiple...)
}

// ------------------------- 事件 ------------------------- //

type XWM_TRAYICON func(wParam, lParam uintptr, pbHandled *bool) int // 托盘图标事件.

type XWM_WINDPROC func(message uint32, wParam, lParam uintptr, pbHandled *bool) int               // 窗口消息过程.
type XWM_WINDPROC1 func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER func(nTimerID uint, pbHandled *bool) int                                        // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_XC_TIMER1 func(hWindow int, nTimerID uint, pbHandled *bool) int                          // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_SETFOCUS_ELE func(hEle int, pbHandled *bool) int                                         // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_SETFOCUS_ELE1 func(hWindow int, hEle int, pbHandled *bool) int                           // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_FLOAT_PANE func(hFloatWnd int, hPane int, pbHandled *bool) int                           // 浮动窗格.
type XWM_FLOAT_PANE1 func(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int             // 浮动窗格.
type XWM_PAINT_END func(hDraw int, pbHandled *bool) int                                           // 窗口绘制完成消息.
type XWM_PAINT_END1 func(hWindow int, hDraw int, pbHandled *bool) int                             // 窗口绘制完成消息.
type XWM_PAINT_DISPLAY func(pbHandled *bool) int                                                  // 窗口绘制完成并且已经显示到屏幕.
type XWM_PAINT_DISPLAY1 func(hWindow int, pbHandled *bool) int                                    // 窗口绘制完成并且已经显示到屏幕.
type XWM_DOCK_POPUP func(hWindowDock, hPane int, pbHandled *bool) int                             // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_DOCK_POPUP1 func(hWindow int, hWindowDock, hPane int, pbHandled *bool) int               // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_BODYVIEW_RECT func(width, height int32, pbHandled *bool) int                             // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
type XWM_BODYVIEW_RECT1 func(hWindow int, width, height int32, pbHandled *bool) int               // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG func(hFloatWnd int, hArray *[6]int, pbHandled *bool) int

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG1 func(hWindow int, hFloatWnd int, hArray *[6]int, pbHandled *bool) int

type WM_PAINT func(hDraw int, pbHandled *bool) int                                        // 窗口绘制消息.
type WM_PAINT1 func(hWindow int, hDraw int, pbHandled *bool) int                          // 窗口绘制消息.
type WM_CLOSE func(pbHandled *bool) int                                                   // 窗口关闭消息.
type WM_CLOSE1 func(hWindow int, pbHandled *bool) int                                     // 窗口关闭消息.
type WM_DESTROY func(pbHandled *bool) int                                                 // 窗口销毁消息.
type WM_DESTROY1 func(hWindow int, pbHandled *bool) int                                   // 窗口销毁消息.
type WM_NCDESTROY func(pbHandled *bool) int                                               // 窗口非客户区销毁消息.
type WM_NCDESTROY1 func(hWindow int, pbHandled *bool) int                                 // 窗口非客户区销毁消息.
type WM_MOUSEMOVE func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标移动消息.
type WM_MOUSEMOVE1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标移动消息.
type WM_LBUTTONDOWN func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                 // 窗口鼠标左键按下消息.
type WM_LBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int   // 窗口鼠标左键按下消息.
type WM_LBUTTONUP func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标左键弹起消息.
type WM_LBUTTONUP1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标左键弹起消息.
type WM_RBUTTONDOWN func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                 // 窗口鼠标右键按下消息.
type WM_RBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int   // 窗口鼠标右键按下消息.
type WM_RBUTTONUP func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标右键弹起消息.
type WM_RBUTTONUP1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标右键弹起消息.
type WM_LBUTTONDBLCLK func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int               // 窗口鼠标左键双击消息.
type WM_LBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int // 窗口鼠标左键双击消息.
type WM_RBUTTONDBLCLK func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int               // 窗口鼠标右键双击消息.
type WM_RBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int // 窗口鼠标右键双击消息.
type WM_MOUSEWHEEL func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                  // 窗口鼠标滚轮滚动消息.
type WM_MOUSEWHEEL1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int    // 窗口鼠标滚轮滚动消息.
type WM_EXITSIZEMOVE func(pbHandled *bool) int                                            // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_EXITSIZEMOVE1 func(hWindow int, pbHandled *bool) int                              // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_MOUSEHOVER func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                  // 窗口鼠标进入消息.
type WM_MOUSEHOVER1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int    // 窗口鼠标进入消息.
type WM_MOUSELEAVE func(pbHandled *bool) int                                              // 窗口鼠标离开消息.
type WM_MOUSELEAVE1 func(hWindow int, pbHandled *bool) int                                // 窗口鼠标离开消息.
type WM_SIZE func(nFlags uint, pPt *xc.SIZE, pbHandled *bool) int                         // 窗口大小改变消息.
type WM_SIZE1 func(hWindow int, nFlags uint, pPt *xc.SIZE, pbHandled *bool) int           // 窗口大小改变消息.
type WM_TIMER func(nIDEvent uint, pbHandled *bool) int                                    // 窗口定时器消息.
type WM_TIMER1 func(hWindow int, nIDEvent uint, pbHandled *bool) int                      // 窗口定时器消息.
type WM_SETFOCUS func(pbHandled *bool) int                                                // 窗口获得焦点.
type WM_SETFOCUS1 func(hWindow int, pbHandled *bool) int                                  // 窗口获得焦点.
type WM_KILLFOCUS func(pbHandled *bool) int                                               // 窗口失去焦点.
type WM_KILLFOCUS1 func(hWindow int, pbHandled *bool) int                                 // 窗口失去焦点.
type WM_KEYDOWN func(wParam, lParam uintptr, pbHandled *bool) int                         // 窗口键盘按键消息.
type WM_KEYDOWN1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int           // 窗口键盘按键消息.
type WM_CAPTURECHANGED func(hWnd uintptr, pbHandled *bool) int                            // 窗口鼠标捕获改变消息.
type WM_CAPTURECHANGED1 func(hWindow int, hWnd uintptr, pbHandled *bool) int              // 窗口鼠标捕获改变消息.
type WM_SETCURSOR func(wParam, lParam uintptr, pbHandled *bool) int                       // 窗口设置鼠标光标.
type WM_SETCURSOR1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int         // 窗口设置鼠标光标.
type WM_CHAR func(wParam, lParam uintptr, pbHandled *bool) int                            // 窗口字符消息.
type WM_CHAR1 func(hWindow int, wParam, lParam uintptr, pbHandled *bool) int              // 窗口字符消息.
type WM_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int                            // 拖动文件到窗口.
type WM_DROPFILES1 func(hWindow int, hDropInfo uintptr, pbHandled *bool) int              // 拖动文件到窗口.

// 托盘图标事件.
func (w *windowBase) Event_TRAYICON(pFun XWM_TRAYICON) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_TRAYICON, pFun)
}

// 窗口消息过程.
func (w *windowBase) Event_WINDPROC(pFun XWM_WINDPROC) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_WINDPROC, pFun)
}

// 窗口消息过程.
func (w *windowBase) Event_WINDPROC1(pFun XWM_WINDPROC1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_WINDPROC, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) Event_XC_TIMER(pFun XWM_XC_TIMER) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_XC_TIMER, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) Event_XC_TIMER1(pFun XWM_XC_TIMER1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_XC_TIMER, pFun)
}

// 窗口事件_置焦点元素. 指定元素获得焦点.
func (w *windowBase) Event_SETFOCUS_ELE(pFun XWM_SETFOCUS_ELE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_SETFOCUS_ELE, pFun)
}

// 窗口事件_置焦点元素. 指定元素获得焦点.
func (w *windowBase) Event_SETFOCUS_ELE1(pFun XWM_SETFOCUS_ELE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_SETFOCUS_ELE, pFun)
}

// 浮动窗格.
func (w *windowBase) Event_FLOAT_PANE(pFun XWM_FLOAT_PANE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_FLOAT_PANE, pFun)
}

// 浮动窗格.
func (w *windowBase) Event_FLOAT_PANE1(pFun XWM_FLOAT_PANE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_FLOAT_PANE, pFun)
}

// 窗口绘制完成消息.
func (w *windowBase) Event_PAINT_END(pFun XWM_PAINT_END) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_PAINT_END, pFun)
}

// 窗口绘制完成消息.
func (w *windowBase) Event_PAINT_END1(pFun XWM_PAINT_END1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_PAINT_END, pFun)
}

// 窗口绘制完成并且已经显示到屏幕.
func (w *windowBase) Event_PAINT_DISPLAY(pFun XWM_PAINT_DISPLAY) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_PAINT_DISPLAY, pFun)
}

// 窗口绘制完成并且已经显示到屏幕.
func (w *windowBase) Event_PAINT_DISPLAY1(pFun XWM_PAINT_DISPLAY1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_PAINT_DISPLAY, pFun)
}

// 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) Event_DOCK_POPUP(pFun XWM_DOCK_POPUP) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_DOCK_POPUP, pFun)
}

// 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) Event_DOCK_POPUP1(pFun XWM_DOCK_POPUP1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_DOCK_POPUP, pFun)
}

// 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件.
func (w *windowBase) Event_BODYVIEW_RECT(pFun XWM_BODYVIEW_RECT) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_BODYVIEW_RECT, pFun)
}

// 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件.
func (w *windowBase) Event_BODYVIEW_RECT1(pFun XWM_BODYVIEW_RECT) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_BODYVIEW_RECT, pFun)
}

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) Event_FLOATWND_DRAG(pFun XWM_FLOATWND_DRAG) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_FLOATWND_DRAG, pFun)
}

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) Event_FLOATWND_DRAG1(pFun XWM_FLOATWND_DRAG1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_FLOATWND_DRAG, pFun)
}

// 窗口绘制消息.
func (w *windowBase) Event_PAINT(pFun WM_PAINT) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_PAINT, pFun)
}

// 窗口绘制消息.
func (w *windowBase) Event_PAINT1(pFun WM_PAINT1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_PAINT, pFun)
}

// 窗口关闭消息.
func (w *windowBase) Event_CLOSE(pFun WM_CLOSE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_CLOSE, pFun)
}

// 窗口关闭消息.
func (w *windowBase) Event_CLOSE1(pFun WM_CLOSE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_CLOSE, pFun)
}

// 窗口销毁消息.
func (w *windowBase) Event_DESTROY(pFun WM_DESTROY) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_DESTROY, pFun)
}

// 窗口销毁消息.
func (w *windowBase) Event_DESTROY1(pFun WM_DESTROY1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_DESTROY, pFun)
}

// 窗口非客户区销毁消息.
func (w *windowBase) Event_NCDESTROY(pFun WM_NCDESTROY) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_NCDESTROY, pFun)
}

// 窗口非客户区销毁消息.
func (w *windowBase) Event_NCDESTROY1(pFun WM_NCDESTROY1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_NCDESTROY, pFun)
}

// 窗口鼠标移动消息.
func (w *windowBase) Event_MOUSEMOVE(pFun WM_MOUSEMOVE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_MOUSEMOVE, pFun)
}

// 窗口鼠标移动消息.
func (w *windowBase) Event_MOUSEMOVE1(pFun WM_MOUSEMOVE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_MOUSEMOVE, pFun)
}

// 窗口鼠标左键按下消息.
func (w *windowBase) Event_LBUTTONDOWN(pFun WM_LBUTTONDOWN) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_LBUTTONDOWN, pFun)
}

// 窗口鼠标左键按下消息.
func (w *windowBase) Event_LBUTTONDOWN1(pFun WM_LBUTTONDOWN1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_LBUTTONDOWN, pFun)
}

// 窗口鼠标左键弹起消息.
func (w *windowBase) Event_LBUTTONUP(pFun WM_LBUTTONUP) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_LBUTTONUP, pFun)
}

// 窗口鼠标左键弹起消息.
func (w *windowBase) Event_LBUTTONUP1(pFun WM_LBUTTONUP1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_LBUTTONUP, pFun)
}

// 窗口鼠标右键按下消息.
func (w *windowBase) Event_RBUTTONDOWN(pFun WM_RBUTTONDOWN) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_RBUTTONDOWN, pFun)
}

// 窗口鼠标右键按下消息.
func (w *windowBase) Event_RBUTTONDOWN1(pFun WM_RBUTTONDOWN1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_RBUTTONDOWN, pFun)
}

// 窗口鼠标右键弹起消息.
func (w *windowBase) Event_RBUTTONUP(pFun WM_RBUTTONUP) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_RBUTTONUP, pFun)
}

// 窗口鼠标右键弹起消息.
func (w *windowBase) Event_RBUTTONUP1(pFun WM_RBUTTONUP1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_RBUTTONUP, pFun)
}

// 窗口鼠标左键双击消息.
func (w *windowBase) Event_LBUTTONDBLCLK(pFun WM_LBUTTONDBLCLK) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_LBUTTONDBLCLK, pFun)
}

// 窗口鼠标左键双击消息.
func (w *windowBase) Event_LBUTTONDBLCLK1(pFun WM_LBUTTONDBLCLK1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_LBUTTONDBLCLK, pFun)
}

// 窗口鼠标右键双击消息.
func (w *windowBase) Event_RBUTTONDBLCLK(pFun WM_RBUTTONDBLCLK) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_RBUTTONDBLCLK, pFun)
}

// 窗口鼠标右键双击消息.
func (w *windowBase) Event_RBUTTONDBLCLK1(pFun WM_RBUTTONDBLCLK1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_RBUTTONDBLCLK, pFun)
}

// 窗口鼠标滚轮滚动消息.
func (w *windowBase) Event_MOUSEWHEEL(pFun WM_MOUSEWHEEL) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_MOUSEWHEEL, pFun)
}

// 窗口鼠标滚轮滚动消息.
func (w *windowBase) Event_MOUSEWHEEL1(pFun WM_MOUSEWHEEL1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_MOUSEWHEEL, pFun)
}

// 窗口退出移动或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) Event_EXITSIZEMOVE(pFun WM_EXITSIZEMOVE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_EXITSIZEMOVE, pFun)
}

// 窗口退出移动或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) Event_EXITSIZEMOVE1(pFun WM_EXITSIZEMOVE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_EXITSIZEMOVE, pFun)
}

// 窗口鼠标进入消息.
func (w *windowBase) Event_MOUSEHOVER(pFun WM_MOUSEHOVER) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_MOUSEHOVER, pFun)
}

// 窗口鼠标进入消息.
func (w *windowBase) Event_MOUSEHOVER1(pFun WM_MOUSEHOVER1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_MOUSEHOVER, pFun)
}

// 窗口鼠标离开消息.
func (w *windowBase) Event_MOUSELEAVE(pFun WM_MOUSELEAVE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_MOUSELEAVE, pFun)
}

// 窗口鼠标离开消息.
func (w *windowBase) Event_MOUSELEAVE1(pFun WM_MOUSELEAVE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_MOUSELEAVE, pFun)
}

// 窗口大小改变消息.
func (w *windowBase) Event_SIZE(pFun WM_SIZE) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_SIZE, pFun)
}

// 窗口大小改变消息.
func (w *windowBase) Event_SIZE1(pFun WM_SIZE1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_SIZE, pFun)
}

// 窗口定时器消息.
func (w *windowBase) Event_TIMER(pFun WM_TIMER) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_TIMER, pFun)
}

// 窗口定时器消息.
func (w *windowBase) Event_TIMER1(pFun WM_TIMER1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_TIMER, pFun)
}

// 窗口获得焦点.
func (w *windowBase) Event_SETFOCUS(pFun WM_SETFOCUS) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_SETFOCUS, pFun)
}

// 窗口获得焦点.
func (w *windowBase) Event_SETFOCUS1(pFun WM_SETFOCUS1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_SETFOCUS, pFun)
}

// 窗口失去焦点.
func (w *windowBase) Event_KILLFOCUS(pFun WM_KILLFOCUS) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_KILLFOCUS, pFun)
}

// 窗口失去焦点.
func (w *windowBase) Event_KILLFOCUS1(pFun WM_KILLFOCUS1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_KILLFOCUS, pFun)
}

// 窗口键盘按键消息.
func (w *windowBase) Event_KEYDOWN(pFun WM_KEYDOWN) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_KEYDOWN, pFun)
}

// 窗口键盘按键消息.
func (w *windowBase) Event_KEYDOWN1(pFun WM_KEYDOWN1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_KEYDOWN, pFun)
}

// 窗口鼠标捕获改变消息.
func (w *windowBase) Event_CAPTURECHANGED(pFun WM_CAPTURECHANGED) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_CAPTURECHANGED, pFun)
}

// 窗口鼠标捕获改变消息.
func (w *windowBase) Event_CAPTURECHANGED1(pFun WM_CAPTURECHANGED1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_CAPTURECHANGED, pFun)
}

// 窗口设置鼠标光标.
func (w *windowBase) Event_SETCURSOR(pFun WM_SETCURSOR) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_SETCURSOR, pFun)
}

// 窗口设置鼠标光标.
func (w *windowBase) Event_SETCURSOR1(pFun WM_SETCURSOR1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_SETCURSOR, pFun)
}

// 窗口字符消息.
func (w *windowBase) Event_CHAR(pFun WM_CHAR) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_CHAR, pFun)
}

// 窗口字符消息.
func (w *windowBase) Event_CHAR1(pFun WM_CHAR1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_CHAR, pFun)
}

// 拖动文件到窗口.
func (w *windowBase) Event_DROPFILES(pFun WM_DROPFILES) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.WM_DROPFILES, pFun)
}

// 拖动文件到窗口.
func (w *windowBase) Event_DROPFILES1(pFun WM_DROPFILES1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.WM_DROPFILES, pFun)
}

type XWM_MENU_POPUP func(hMenu int, pbHandled *bool) int                                                         // 菜单弹出.
type XWM_MENU_POPUP1 func(hWindow int, hMenu int, pbHandled *bool) int                                           // 菜单弹出.
type XWM_MENU_POPUP_WND func(hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int                           // 菜单弹出窗口.
type XWM_MENU_POPUP_WND1 func(hWindow int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XWM_MENU_SELECT func(nID int32, pbHandled *bool) int                                                        // 菜单选择.
type XWM_MENU_SELECT1 func(hWindow int, nID int32, pbHandled *bool) int                                          // 菜单选择.
type XWM_MENU_EXIT func(pbHandled *bool) int                                                                     // 菜单退出.
type XWM_MENU_EXIT1 func(hWindow int, pbHandled *bool) int                                                       // 菜单退出.
type XWM_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int               // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAW_BACKGROUND1 func(hWindow int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAWITEM func(hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int                            // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
type XWM_MENU_DRAWITEM1 func(hWindow int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 菜单弹出.
func (w *windowBase) Event_MENU_POPUP(pFun XWM_MENU_POPUP) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_POPUP, pFun)
}

// 菜单弹出.
func (w *windowBase) Event_MENU_POPUP1(pFun XWM_MENU_POPUP1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_POPUP, pFun)
}

// 菜单弹出窗口.
func (w *windowBase) Event_MENU_POPUP_WND(pFun XWM_MENU_POPUP_WND) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_POPUP_WND, pFun)
}

// 菜单弹出窗口.
func (w *windowBase) Event_MENU_POPUP_WND1(pFun XWM_MENU_POPUP_WND1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_POPUP_WND, pFun)
}

// 菜单选择.
func (w *windowBase) Event_MENU_SELECT(pFun XWM_MENU_SELECT) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_SELECT, pFun)
}

// 菜单选择.
func (w *windowBase) Event_MENU_SELECT1(pFun XWM_MENU_SELECT1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_SELECT, pFun)
}

// 菜单退出.
func (w *windowBase) Event_MENU_EXIT(pFun XWM_MENU_EXIT) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_EXIT, pFun)
}

// 菜单退出.
func (w *windowBase) Event_MENU_EXIT1(pFun XWM_MENU_EXIT1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_EXIT, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) Event_MENU_DRAW_BACKGROUND(pFun XWM_MENU_DRAW_BACKGROUND) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) Event_MENU_DRAW_BACKGROUND1(pFun XWM_MENU_DRAW_BACKGROUND1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) Event_MENU_DRAWITEM(pFun XWM_MENU_DRAWITEM) bool {
	return xc.XWnd_RegEventC(w.Handle, xcc.XWM_MENU_DRAWITEM, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) Event_MENU_DRAWITEM1(pFun XWM_MENU_DRAWITEM1) bool {
	return xc.XWnd_RegEventC1(w.Handle, xcc.XWM_MENU_DRAWITEM, pFun)
}
