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

// MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_. 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *windowBase) MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {
	return xc.XC_MessageBox(pTitle, pText, nFlags, w.GetHWND(), XCStyle)
}

// Msg_Create 消息框_创建, 然后请调用 DoModal() 方法显示模态窗口.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return *ModalWindow 模态窗口对象.
func (w *windowBase) Msg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XMsg_Create(pTitle, pText, nFlags, w.GetHWND(), XCStyle))
	return p
}

// Msg_CreateEx 消息框_创建扩展, 然后请调用 DoModal() 方法显示模态窗口.
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return *ModalWindow 模态窗口对象.
func (w *windowBase) Msg_CreateEx(dwExStyle, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XMsg_CreateEx(dwExStyle, dwStyle, lpClassName, pTitle, pText, nFlags, w.GetHWND(), XCStyle))
	return p
}

// 炫彩_发送窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) SendMessage(msg uint32, wParam, lParam uint) uint {
	return xc.XC_SendMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_投递窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) PostMessage(msg uint32, wParam int32, lParam int32) bool {
	return xc.XC_PostMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
func (w *windowBase) IsHWINDOW() bool {
	return xc.XC_IsHWINDOW(w.Handle)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// nID: ID值.
func (w *windowBase) GetObjectByID(nID int) int {
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
func (w *windowBase) ShowWindow(nCmdShow xcc.SW_) int {
	return xc.XWnd_ShowWindow(w.Handle, nCmdShow)
}

// 窗口_置顶.
func (w *windowBase) SetTop() int {
	return xc.XWnd_SetTop(w.Handle)
}

// 窗口_注册事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) RegEventC(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RegEventC(w.Handle, nEvent, pFun)
}

// 窗口_注册事件C1.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) RegEventC1(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RegEventC1(w.Handle, nEvent, pFun)
}

// 窗口_移除事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) RemoveEventC(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RemoveEventC(w.Handle, nEvent, pFun)
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
func (w *windowBase) InsertChild(hChild int, index int) bool {
	return xc.XWnd_InsertChild(w.Handle, hChild, index)
}

// 窗口_取HWND.
func (w *windowBase) GetHWND() uintptr {
	return xc.XWnd_GetHWND(w.Handle)
}

// 窗口_重绘.
//
// bImmediate: 是否立即重绘, 默认为否.
func (w *windowBase) Redraw(bImmediate bool) int {
	return xc.XWnd_Redraw(w.Handle, bImmediate)
}

// 窗口_重绘指定区域.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *windowBase) RedrawRect(pRect *xc.RECT, bImmediate bool) int {
	return xc.XWnd_RedrawRect(w.Handle, pRect, bImmediate)
}

// 窗口_置坐标.
//
// pRect: 坐标.
func (w *windowBase) SetRect(pRect *xc.RECT) bool {
	return xc.XWnd_SetRect(w.Handle, pRect)
}

// 窗口_置焦点.
//
// hFocusEle: 将获得焦点的元素.
func (w *windowBase) SetFocusEle(hFocusEle int) bool {
	return xc.XWnd_SetFocusEle(w.Handle, hFocusEle)
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
func (w *windowBase) DrawWindow(hDraw int) int {
	return xc.XWnd_DrawWindow(w.Handle, hDraw)
}

// 窗口_居中.
func (w *windowBase) Center() int {
	return xc.XWnd_Center(w.Handle)
}

// 窗口_居中扩展.
//
// width: 窗口宽度.
//
// height: 窗口高度.
func (w *windowBase) CenterEx(width, height int) int {
	return xc.XWnd_CenterEx(w.Handle, width, height)
}

// 窗口_置光标.
//
// hCursor: 鼠标光标句柄.
func (w *windowBase) SetCursor(hCursor int) int {
	return xc.XWnd_SetCursor(w.Handle, hCursor)
}

// 窗口_取光标.
func (w *windowBase) GetCursor() int {
	return xc.XWnd_GetCursor(w.Handle)
}

// 窗口_启用拖动边框.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragBorder(bEnable bool) int {
	return xc.XWnd_EnableDragBorder(w.Handle, bEnable)
}

// 窗口_启用拖动窗口.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragWindow(bEnable bool) int {
	return xc.XWnd_EnableDragWindow(w.Handle, bEnable)
}

// 窗口_启用拖动标题栏.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragCaption(bEnable bool) int {
	return xc.XWnd_EnableDragCaption(w.Handle, bEnable)
}

// 窗口_启用绘制背景.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDrawBk(bEnable bool) int {
	return xc.XWnd_EnableDrawBk(w.Handle, bEnable)
}

// 窗口_启用自动焦点.
//
// bEnable: 是否启用.
func (w *windowBase) EnableAutoFocus(bEnable bool) int {
	return xc.XWnd_EnableAutoFocus(w.Handle, bEnable)
}

// 窗口_启用允许最大化.
//
// bEnable: 是否启用.
func (w *windowBase) EnableMaxWindow(bEnable bool) int {
	return xc.XWnd_EnableMaxWindow(w.Handle, bEnable)
}

// 窗口_启用限制窗口大小.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLimitWindowSize(bEnable bool) int {
	return xc.XWnd_EnableLimitWindowSize(w.Handle, bEnable)
}

// 窗口_启用布局.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLayout(bEnable bool) int {
	return xc.XWnd_EnableLayout(w.Handle, bEnable)
}

// 窗口_启用布局覆盖边框.
//
// bEnable: 是否启用.
func (w *windowBase) EnableLayoutOverlayBorder(bEnable bool) int {
	return xc.XWnd_EnableLayoutOverlayBorder(w.Handle, bEnable)
}

// 窗口_显示布局边界.
//
// bEnable: 是否启用.
func (w *windowBase) ShowLayoutFrame(bEnable bool) int {
	return xc.XWnd_ShowLayoutFrame(w.Handle, bEnable)
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
func (w *windowBase) SetCaptureEle(hEle int) int {
	return xc.XWnd_SetCaptureEle(w.Handle, hEle)
}

// 窗口_取鼠标捕获元素.
func (w *windowBase) GetCaptureEle() int {
	return xc.XWnd_GetCaptureEle(w.Handle)
}

// 窗口_取绘制矩形.
//
// pRcPaint: 重绘区域坐标.
func (w *windowBase) GetDrawRect(pRcPaint *xc.RECT) int {
	return xc.XWnd_GetDrawRect(w.Handle, pRcPaint)
}

// 窗口_置系统光标.
//
// hCursor: 光标句柄.
func (w *windowBase) SetCursorSys(hCursor int) int {
	return xc.XWnd_SetCursorSys(w.Handle, hCursor)
}

// 窗口_置字体.
//
// hFontx: 炫彩字体句柄.
func (w *windowBase) SetFont(hFontx int) int {
	return xc.XWnd_SetFont(w.Handle, hFontx)
}

// 窗口_置文本颜色.
//
// color: ABGR 颜色值.
func (w *windowBase) SetTextColor(color int) int {
	return xc.XWnd_SetTextColor(w.Handle, color)
}

// 窗口_取文本颜色, 返回ABGR 颜色.
func (w *windowBase) GetTextColor() int {
	return xc.XWnd_GetTextColor(w.Handle)
}

// 窗口_取文本颜色扩展, 返回ABGR 颜色.
func (w *windowBase) GetTextColorEx() int {
	return xc.XWnd_GetTextColorEx(w.Handle)
}

// 窗口_置ID.
//
// nID: ID值.
func (w *windowBase) SetID(nID int) int {
	return xc.XWnd_SetID(w.Handle, nID)
}

// 窗口_取ID.
func (w *windowBase) GetID() int {
	return xc.XWnd_GetID(w.Handle)
}

// 窗口_置名称.
//
// pName: name值, 字符串.
func (w *windowBase) SetName(pName string) int {
	return xc.XWnd_SetName(w.Handle, pName)
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
func (w *windowBase) SetBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetBorderSize(w.Handle, left, top, right, bottom)
}

// 窗口_取边大小.
func (w *windowBase) GetBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetBorderSize(w.Handle, pBorder)
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
func (w *windowBase) SetPadding(left, top, right, bottom int) int {
	return xc.XWnd_SetPadding(w.Handle, left, top, right, bottom)
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
func (w *windowBase) SetDragBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetDragBorderSize(w.Handle, left, top, right, bottom)
}

// 窗口_取拖动边框大小.
//
// pSize: 拖动边框大小.
func (w *windowBase) GetDragBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetDragBorderSize(w.Handle, pBorder)
}

// 窗口_置最小大小.
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *windowBase) SetMinimumSize(width, height int) int {
	return xc.XWnd_SetMinimumSize(w.Handle, width, height)
}

// 窗口_测试点击子元素.
//
// pPt: 左边点.
func (w *windowBase) HitChildEle(pPt *xc.POINT) int {
	return xc.XWnd_HitChildEle(w.Handle, pPt)
}

// 窗口_取子对象数量.
func (w *windowBase) GetChildCount() int {
	return xc.XWnd_GetChildCount(w.Handle)
}

// 窗口_取子对象从索引.
//
// index: 元素索引.
func (w *windowBase) GetChildByIndex(index int) int {
	return xc.XWnd_GetChildByIndex(w.Handle, index)
}

// 窗口_取子对象从ID.
//
// nID: 元素ID, ID必须大于0.
func (w *windowBase) GetChildByID(nID int) int {
	return xc.XWnd_GetChildByID(w.Handle, nID)
}

// 窗口_取子对象.
//
// nID: 对象ID,ID必须大于0.
func (w *windowBase) GetChild(nID int) int {
	return xc.XWnd_GetChild(w.Handle, nID)
}

// 窗口_关闭.
func (w *windowBase) CloseWindow() int {
	return xc.XWnd_CloseWindow(w.Handle)
}

// 窗口_调整布局.
func (w *windowBase) AdjustLayout() int {
	return xc.XWnd_AdjustLayout(w.Handle)
}

// 窗口_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
func (w *windowBase) AdjustLayoutEx(nFlags xcc.AdjustLayout_) int {
	return xc.XWnd_AdjustLayoutEx(w.Handle, nFlags)
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
func (w *windowBase) CreateCaret(hEle, x, y, width, height int) int {
	return xc.XWnd_CreateCaret(w.Handle, hEle, x, y, width, height)
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
func (w *windowBase) SetCaretPos(x, y, width, height int, bUpdate bool) int {
	return xc.XWnd_SetCaretPos(w.Handle, x, y, width, height, bUpdate)
}

// 窗口_置插入符颜色.
//
// color: 颜色值, ABGR 颜色.
func (w *windowBase) SetCaretColor(color int) int {
	return xc.XWnd_SetCaretColor(w.Handle, color)
}

// 窗口_显示插入符.
//
// bShow: 是否显示.
func (w *windowBase) ShowCaret(bShow bool) int {
	return xc.XWnd_ShowCaret(w.Handle, bShow)
}

// 窗口_销毁插入符.
func (w *windowBase) DestroyCaret() int {
	return xc.XWnd_DestroyCaret(w.Handle)
}

// 窗口_取插入符元素.
func (w *windowBase) GetCaretHELE() int {
	return xc.XWnd_GetCaretHELE(w.Handle)
}

// 窗口_取客户区坐标.
//
// pRect: 坐标.
func (w *windowBase) GetClientRect(pRect *xc.RECT) int {
	return xc.XWnd_GetClientRect(w.Handle, pRect)
}

// 窗口_取Body坐标.
//
// pRect: 坐标.
func (w *windowBase) GetBodyRect(pRect *xc.RECT) int {
	return xc.XWnd_GetBodyRect(w.Handle, pRect)
}

// 窗口_取布局坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) GetLayoutRect(pRect *xc.RECT) int {
	return xc.XWnd_GetLayoutRect(w.Handle, pRect)
}

// 窗口_移动窗口.
//
// x: X坐标.
//
// y: Y坐标.
func (w *windowBase) SetPosition(x, y int32) {
	xc.XWnd_SetPosition(w.Handle, x, y)
}

// 窗口_取坐标.
//
// pRect: 坐标.
func (w *windowBase) GetRect(pRect *xc.RECT) int {
	return xc.XWnd_GetRect(w.Handle, pRect)
}

// 窗口_最大化.
//
// bMaximize: 是否最大化.
func (w *windowBase) MaxWindow(bMaximize bool) int {
	return xc.XWnd_MaxWindow(w.Handle, bMaximize)
}

// 窗口_置定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) KillTimer(nIDEvent int) int {
	return xc.XWnd_KillTimer(w.Handle, nIDEvent)
}

// 窗口_置炫彩定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetXCTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetXCTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭炫彩定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) KillXCTimer(nIDEvent int) int {
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
func (w *windowBase) SetBkMagager(hBkInfoM int) int {
	return xc.XWnd_SetBkMagager(w.Handle, hBkInfoM)
}

// 窗口_置透明类型.
//
// nType: 窗口透明类型: xcc.Window_Transparent_.
func (w *windowBase) SetTransparentType(nType xcc.Window_Transparent_) int {
	return xc.XWnd_SetTransparentType(w.Handle, nType)
}

// 窗口_置透明度.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *windowBase) SetTransparentAlpha(alpha byte) int {
	return xc.XWnd_SetTransparentAlpha(w.Handle, alpha)
}

// 窗口_置透明色.
//
// color: 窗口透明色, ABGR 颜色.
func (w *windowBase) SetTransparentColor(color int) int {
	return xc.XWnd_SetTransparentColor(w.Handle, color)
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
// color: 阴影颜色, ABGR 颜色.
func (w *windowBase) SetShadowInfo(nSize int, nDepth byte, nAngeleSize int, bRightAngle bool, color int) int {
	return xc.XWnd_SetShadowInfo(w.Handle, nSize, nDepth, nAngeleSize, bRightAngle, color)
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
// pColor: 阴影颜色, ABGR 颜色.
func (w *windowBase) GetShadowInfo(pnSize, pnDepth, pnAngeleSize *int32, pbRightAngle *bool, pColor *int) int {
	return xc.XWnd_GetShadowInfo(w.Handle, pnSize, pnDepth, pnAngeleSize, pbRightAngle, pColor)
}

// 窗口_取透明类型, 返回: xcc.Window_Transparent_.
func (w *windowBase) GetTransparentType() xcc.Window_Transparent_ {
	return xc.XWnd_GetTransparentType(w.Handle)
}

// 窗口_启用拖放文件.
//
// bEnable: 是否启用.
func (w *windowBase) EnableDragFiles(bEnable bool) bool {
	return xc.XWnd_EnableDragFiles(w.Handle, bEnable)
}

// 窗口_显示 显示隐藏窗口.
//
// bShow: 是否显示.
func (w *windowBase) Show(bShow bool) int {
	return xc.XWnd_Show(w.Handle, bShow)
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
func (w *windowBase) SetIcon(hImage int) int {
	return xc.XWnd_SetIcon(w.Handle, hImage)
}

// 窗口_置标题.
//
// pTitle: 标题文本.
func (w *windowBase) SetTitle(pTitle string) int {
	return xc.XWnd_SetTitle(w.Handle, pTitle)
}

// 窗口_置标题颜色.
//
// color: ABGR 颜色.
func (w *windowBase) SetTitleColor(color int) int {
	return xc.XWnd_SetTitleColor(w.Handle, color)
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

// 窗口_取标题颜色, 返回ABGR 颜色.
func (w *windowBase) GetTitleColor() int {
	return xc.XWnd_GetTitleColor(w.Handle)
}

// 窗口_添加背景边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func (w *windowBase) AddBkBorder(nState xcc.Window_State_Flag_, color int, width int) int {
	return xc.XWnd_AddBkBorder(w.Handle, nState, color, width)
}

// 窗口_添加背景填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func (w *windowBase) AddBkFill(nState xcc.Window_State_Flag_, color int) int {
	return xc.XWnd_AddBkFill(w.Handle, nState, color)
}

// 窗口_添加背景图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (w *windowBase) AddBkImage(nState xcc.Window_State_Flag_, hImage int) int {
	return xc.XWnd_AddBkImage(w.Handle, nState, hImage)
}

// 窗口_取背景对象数量.
func (w *windowBase) GetBkInfoCount() int {
	return xc.XWnd_GetBkInfoCount(w.Handle)
}

// 窗口_清空背景对象.
func (w *windowBase) ClearBkInfo() int {
	return xc.XWnd_ClearBkInfo(w.Handle)
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
func (w *windowBase) NotifyMsg_SetDuration(duration int) int {
	return xc.XNotifyMsg_SetDuration(w.Handle, duration)
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
func (w *windowBase) NotifyMsg_SetParentMargin(left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetParentMargin(w.Handle, left, top, right, bottom)
}

// 通知消息_置标题高度.
//
// nHeight: 高度.
func (w *windowBase) NotifyMsg_SetCaptionHeight(nHeight int) int {
	return xc.XNotifyMsg_SetCaptionHeight(w.Handle, nHeight)
}

// 通知消息_置宽度, 设置默认宽度.
//
// nWidth: 宽度.
func (w *windowBase) NotifyMsg_SetWidth(nWidth int) int {
	return xc.XNotifyMsg_SetWidth(w.Handle, nWidth)
}

// 通知消息_置间隔.
//
// nSpace: 间隔大小.
func (w *windowBase) NotifyMsg_SetSpace(nSpace int) int {
	return xc.XNotifyMsg_SetSpace(w.Handle, nSpace)
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
func (w *windowBase) NotifyMsg_SetBorderSize(left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetBorderSize(w.Handle, left, top, right, bottom)
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (w *windowBase) SetBkInfo(pText string) int {
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
func (w *windowBase) SetCaptionMargin(left int, top int, right int, bottom int) int {
	return xc.XWnd_SetCaptionMargin(w.Handle, left, top, right, bottom)
}

// SetTopEx 窗口_置顶Ex.
//
//	@param b 是否置顶.
//	@return bool
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
func (w *windowBase) SetWindowPos(hWndInsertAfter xcc.HWND_, x, y, cx, cy int32, uFlags xcc.SWP_) int {
	return xc.XWnd_SetWindowPos(w.Handle, hWndInsertAfter, x, y, cx, cy, uFlags)
}

// 窗口_取DPI. 获取当前窗口所在显示器DPI, 返回窗口DPI.
func (w *windowBase) GetDPI() int {
	return xc.XWnd_GetDPI(w.Handle)
}

// 窗口_坐标转换DPI. 窗口客户区坐标转换到缩放后DPI坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) RectToDPI(pRect *xc.RECT) int {
	return xc.XWnd_RectToDPI(w.Handle, pRect)
}

// 窗口_坐标点转换DPI. 窗口客户区坐标点转换到缩放后.
//
// pPt: 接收返回坐标点.
func (w *windowBase) PointToDPI(pPt *xc.POINT) int {
	return xc.XWnd_PointToDPI(w.Handle, pPt)
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
func (w *windowBase) SetDPI(nDPI int) int {
	return xc.XWnd_SetDPI(w.Handle, nDPI)
}

// 窗口_置大小. 设置窗口宽高.
//
// 宽: width.
//
// 高: height.
func (w *windowBase) SetSize(width, height int32) bool {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Right = rc.Left + width
	rc.Bottom = rc.Top + height
	return xc.XWnd_SetRect(w.Handle, &rc)
}

// 窗口_置宽度.
//
// 宽: width.
func (w *windowBase) SetWidth(width int32) bool {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Right = rc.Left + width
	return xc.XWnd_SetRect(w.Handle, &rc)
}

// 窗口_置高度.
//
// 高: height.
func (w *windowBase) SetHeight(height int32) bool {
	var rc xc.RECT
	xc.XWnd_GetRect(w.Handle, &rc)
	rc.Bottom = rc.Top + height
	return xc.XWnd_SetRect(w.Handle, &rc)
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
func (w *windowBase) SetLeft(x int32) {
	xc.XWnd_SetPosition(w.Handle, x, w.GetTop())
}

// 窗口_置顶边.
//
// y: 顶边y坐标.
func (w *windowBase) SetTopEdge(y int32) {
	xc.XWnd_SetPosition(w.Handle, w.GetLeft(), y)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(w.Handle, bEnable)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(w.Handle, bEnable)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(w.Handle, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (w *windowBase) SetAlignH(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignH(w.Handle, nAlign)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (w *windowBase) SetAlignV(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignV(w.Handle, nAlign)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
func (w *windowBase) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {
	return xc.XLayoutBox_SetAlignBaseline(w.Handle, nAlign)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
func (w *windowBase) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(w.Handle, nSpace)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
func (w *windowBase) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(w.Handle, nSpace)
}

/*
下面都是事件
*/

type XWM_TRAYICON func(wParam, lParam uint, pbHandled *bool) int // 托盘图标事件.

type XWM_WINDPROC func(message uint32, wParam, lParam uint, pbHandled *bool) int               // 窗口消息过程.
type XWM_WINDPROC1 func(hWindow int, message uint32, wParam, lParam uint, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER func(nTimerID uint, pbHandled *bool) int                                     // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_XC_TIMER1 func(hWindow int, nTimerID uint, pbHandled *bool) int                       // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_SETFOCUS_ELE func(hEle int, pbHandled *bool) int                                      // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_SETFOCUS_ELE1 func(hWindow int, hEle int, pbHandled *bool) int                        // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_FLOAT_PANE func(hFloatWnd int, hPane int, pbHandled *bool) int                        // 浮动窗格.
type XWM_FLOAT_PANE1 func(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int          // 浮动窗格.
type XWM_PAINT_END func(hDraw int, pbHandled *bool) int                                        // 窗口绘制完成消息.
type XWM_PAINT_END1 func(hWindow int, hDraw int, pbHandled *bool) int                          // 窗口绘制完成消息.
type XWM_PAINT_DISPLAY func(pbHandled *bool) int                                               // 窗口绘制完成并且已经显示到屏幕.
type XWM_PAINT_DISPLAY1 func(hWindow int, pbHandled *bool) int                                 // 窗口绘制完成并且已经显示到屏幕.
type XWM_DOCK_POPUP func(hWindowDock, hPane int, pbHandled *bool) int                          // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_DOCK_POPUP1 func(hWindow int, hWindowDock, hPane int, pbHandled *bool) int            // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_BODYVIEW_RECT func(width, height int32, pbHandled *bool) int                          // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
type XWM_BODYVIEW_RECT1 func(hWindow int, width, height int32, pbHandled *bool) int            // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
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
type WM_KEYDOWN func(wParam, lParam uint, pbHandled *bool) int                            // 窗口键盘按键消息.
type WM_KEYDOWN1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int              // 窗口键盘按键消息.
type WM_CAPTURECHANGED func(hWnd uintptr, pbHandled *bool) int                            // 窗口鼠标捕获改变消息.
type WM_CAPTURECHANGED1 func(hWindow int, hWnd uintptr, pbHandled *bool) int              // 窗口鼠标捕获改变消息.
type WM_SETCURSOR func(wParam, lParam uint, pbHandled *bool) int                          // 窗口设置鼠标光标.
type WM_SETCURSOR1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int            // 窗口设置鼠标光标.
type WM_CHAR func(wParam, lParam uint, pbHandled *bool) int                               // 窗口字符消息.
type WM_CHAR1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int                 // 窗口字符消息.
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
