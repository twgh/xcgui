package xc

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 窗口_创建, 返回: GUI库窗口资源句柄.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式, Window_Style_.
func XWnd_Create(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xWnd_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 窗口_创建扩展, 返回: GUI库窗口资源句柄.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式, Window_Style_.
func XWnd_CreateEx(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xWnd_CreateEx.Call(uintptr(dwExStyle), uintptr(dwStyle), common.StrPtr(lpClassName), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 窗口_显示.
//
// hWindow: 窗口句柄.
//
// nCmdShow: 显示方式: xcc.SW_.
func XWnd_ShowWindow(hWindow int, nCmdShow xcc.SW_) int {
	r, _, _ := xWnd_ShowWindow.Call(uintptr(hWindow), uintptr(nCmdShow))
	return int(r)
}

// 窗口_置顶.
//
// hWindow: 窗口句柄.
func XWnd_SetTop(hWindow int) int {
	r, _, _ := xWnd_SetTop.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_注册事件C.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func XWnd_RegEventC(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RegEventC.Call(uintptr(hWindow), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_注册事件C1.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func XWnd_RegEventC1(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RegEventC1.Call(uintptr(hWindow), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_移除事件C.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func XWnd_RemoveEventC(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RemoveEventC.Call(uintptr(hWindow), uintptr(nEvent), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_注册事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_ .
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XWnd_RegEventCEx(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RegEventC.Call(uintptr(hWindow), uintptr(nEvent), pFun)
	return r != 0
}

// 窗口_注册事件C1Ex, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XWnd_RegEventC1Ex(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RegEventC1.Call(uintptr(hWindow), uintptr(nEvent), pFun)
	return r != 0
}

// 窗口_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func XWnd_RemoveEventCEx(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RemoveEventC.Call(uintptr(hWindow), uintptr(nEvent), pFun)
	return r != 0
}

// 窗口_添加子对象.
//
// hWindow: 窗口句柄.
//
// hChild: 要添加的对象句柄.
func XWnd_AddChild(hWindow int, hChild int) bool {
	r, _, _ := xWnd_AddChild.Call(uintptr(hWindow), uintptr(hChild))
	return r != 0
}

// 窗口_插入子对象.
//
// hWindow: 窗口句柄.
//
// hChild: 要插入的对象句柄.
//
// index: 插入位置索引.
func XWnd_InsertChild(hWindow int, hChild int, index int) bool {
	r, _, _ := xWnd_InsertChild.Call(uintptr(hWindow), uintptr(hChild), uintptr(index))
	return r != 0
}

// 窗口_取HWND.
//
// hWindow: 窗口句柄.
func XWnd_GetHWND(hWindow int) uintptr {
	r, _, _ := xWnd_GetHWND.Call(uintptr(hWindow))
	return r
}

// 窗口_重绘.
//
// hWindow: 窗口资源句柄.
//
// bImmediate: 是否立即重绘, 默认为否.
func XWnd_Redraw(hWindow int, bImmediate bool) int {
	r, _, _ := xWnd_Redraw.Call(uintptr(hWindow), common.BoolPtr(bImmediate))
	return int(r)
}

// 窗口_重绘指定区域.
//
// hWindow: 窗口资源句柄.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func XWnd_RedrawRect(hWindow int, pRect *RECT, bImmediate bool) int {
	r, _, _ := xWnd_RedrawRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bImmediate))
	return int(r)
}

// 窗口_置坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.
func XWnd_SetRect(hWindow int, pRect *RECT) bool {
	r, _, _ := xWnd_SetRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}

// 窗口_置焦点.
//
// hWindow: 窗口资源句柄.
//
// hFocusEle: 将获得焦点的元素.
func XWnd_SetFocusEle(hWindow int, hFocusEle int) bool {
	r, _, _ := xWnd_SetFocusEle.Call(uintptr(hWindow), uintptr(hFocusEle))
	return r != 0
}

// 窗口_取焦点.
//
// hWindow: 窗口资源句柄.
func XWnd_GetFocusEle(hWindow int) int {
	r, _, _ := xWnd_GetFocusEle.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取鼠标停留元素.
//
// hWindow: 窗口资源句柄.
func XWnd_GetStayEle(hWindow int) int {
	r, _, _ := xWnd_GetStayEle.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_绘制, 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序.
//
// hWindow: 窗口资源句柄.
//
// hDraw: 图形绘制句柄.
func XWnd_DrawWindow(hWindow int, hDraw int) int {
	r, _, _ := xWnd_DrawWindow.Call(uintptr(hWindow), uintptr(hDraw))
	return int(r)
}

// 窗口_居中.
//
// hWindow: 窗口资源句柄.
func XWnd_Center(hWindow int) int {
	r, _, _ := xWnd_Center.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_居中扩展.
//
// hWindow: 窗口资源句柄.
//
// width: 窗口宽度.
//
// height: 窗口高度.
func XWnd_CenterEx(hWindow, width, height int) int {
	r, _, _ := xWnd_CenterEx.Call(uintptr(hWindow), uintptr(width), uintptr(height))
	return int(r)
}

// 窗口_置光标.
//
// hWindow: 窗口句柄.
//
// hCursor: 鼠标光标句柄.
func XWnd_SetCursor(hWindow int, hCursor int) int {
	r, _, _ := xWnd_SetCursor.Call(uintptr(hWindow), uintptr(hCursor))
	return int(r)
}

// 窗口_取光标.
//
// hWindow: 窗口句柄.
func XWnd_GetCursor(hWindow int) int {
	r, _, _ := xWnd_GetCursor.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_启用拖动边框.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableDragBorder(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableDragBorder.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用拖动窗口.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableDragWindow(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableDragWindow.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用拖动标题栏.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableDragCaption(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableDragCaption.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用绘制背景.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableDrawBk(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableDrawBk.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用自动焦点.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableAutoFocus(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableAutoFocus.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用允许最大化.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableMaxWindow(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableMaxWindow.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用限制窗口大小.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableLimitWindowSize(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableLimitWindowSize.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用布局.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableLayout(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableLayout.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_启用布局覆盖边框.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableLayoutOverlayBorder(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_EnableLayoutOverlayBorder.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_显示布局边界.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_ShowLayoutFrame(hWindow int, bEnable bool) int {
	r, _, _ := xWnd_ShowLayoutFrame.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 窗口_判断启用布局.
//
// hWindow: 窗口句柄.
func XWnd_IsEnableLayout(hWindow int) bool {
	r, _, _ := xWnd_IsEnableLayout.Call(uintptr(hWindow))
	return r != 0
}

// 窗口_是否最大化.
//
// hWindow: 窗口句柄.
func XWnd_IsMaxWindow(hWindow int) bool {
	r, _, _ := xWnd_IsMaxWindow.Call(uintptr(hWindow))
	return r != 0
}

// 窗口_置鼠标捕获元素.
//
// hWindow: 窗口句柄.
//
// hEle: 元素句柄.
func XWnd_SetCaptureEle(hWindow, hEle int) int {
	r, _, _ := xWnd_SetCaptureEle.Call(uintptr(hWindow), uintptr(hEle))
	return int(r)
}

// 窗口_取鼠标捕获元素.
//
// hWindow: 窗口句柄.
func XWnd_GetCaptureEle(hWindow int) int {
	r, _, _ := xWnd_GetCaptureEle.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取绘制矩形.
//
// hWindow: 窗口句柄.
//
// pRcPaint: 重绘区域坐标.
func XWnd_GetDrawRect(hWindow int, pRcPaint *RECT) int {
	r, _, _ := xWnd_GetDrawRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRcPaint)))
	return int(r)
}

// 窗口_置系统光标.
//
// hWindow: 窗口句柄.
//
// hCursor: 光标句柄.
func XWnd_SetCursorSys(hWindow, hCursor int) int {
	r, _, _ := xWnd_SetCursorSys.Call(uintptr(hWindow), uintptr(hCursor))
	return int(r)
}

// 窗口_置字体.
//
// hWindow: 窗口句柄.
//
// hFontx: 炫彩字体句柄.
func XWnd_SetFont(hWindow, hFontx int) int {
	r, _, _ := xWnd_SetFont.Call(uintptr(hWindow), uintptr(hFontx))
	return int(r)
}

// 窗口_置文本颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色值.
func XWnd_SetTextColor(hWindow, color int) int {
	r, _, _ := xWnd_SetTextColor.Call(uintptr(hWindow), uintptr(color))
	return int(r)
}

// 窗口_取文本颜色.
//
// hWindow: 窗口句柄.
func XWnd_GetTextColor(hWindow int) int {
	r, _, _ := xWnd_GetTextColor.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取文本颜色扩展.
//
// hWindow: 窗口句柄.
func XWnd_GetTextColorEx(hWindow int) int {
	r, _, _ := xWnd_GetTextColorEx.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_置ID.
//
// hWindow: 窗口句柄.
//
// nID: ID值.
func XWnd_SetID(hWindow, nID int) int {
	r, _, _ := xWnd_SetID.Call(uintptr(hWindow), uintptr(nID))
	return int(r)
}

// 窗口_取ID.
//
// hWindow: 窗口句柄.
func XWnd_GetID(hWindow int) int {
	r, _, _ := xWnd_GetID.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_置名称.
//
// hWindow: 窗口句柄.
//
// pName: name值, 字符串.
func XWnd_SetName(hWindow int, pName string) int {
	r, _, _ := xWnd_SetName.Call(uintptr(hWindow), common.StrPtr(pName))
	return int(r)
}

// 窗口_取名称.
//
// hWindow: 窗口句柄.
func XWnd_GetName(hWindow int) string {
	r, _, _ := xWnd_GetName.Call(uintptr(hWindow))
	return common.UintPtrToString(r)
}

// 窗口_置边大小.
//
// hWindow: 窗口句柄.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底部大小.
func XWnd_SetBorderSize(hWindow, left, top, right, bottom int) int {
	r, _, _ := xWnd_SetBorderSize.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口_取边大小.
//
// hWindow: 窗口句柄.
func XWnd_GetBorderSize(hWindow int, pBorder *RECT) int {
	r, _, _ := xWnd_GetBorderSize.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pBorder)))
	return int(r)
}

// 窗口_置布局内填充大小.
//
// hWindow: 窗口句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func XWnd_SetPadding(hWindow, left, top, right, bottom int) int {
	r, _, _ := xWnd_SetPadding.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口_置拖动边框大小.
//
// hWindow: 窗口句柄.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底边大小.
func XWnd_SetDragBorderSize(hWindow, left, top, right, bottom int) int {
	r, _, _ := xWnd_SetDragBorderSize.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口_取拖动边框大小.
//
// hWindow: 窗口句柄.
//
// pSize: 拖动边框大小.
func XWnd_GetDragBorderSize(hWindow int, pBorder *RECT) int {
	r, _, _ := xWnd_GetDragBorderSize.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pBorder)))
	return int(r)
}

// 窗口_置最小大小.
//
// hWindow: 窗口句柄.
//
// width: 最小宽度.
//
// height: 最小高度.
func XWnd_SetMinimumSize(hWindow, width, height int) int {
	r, _, _ := xWnd_SetMinimumSize.Call(uintptr(hWindow), uintptr(width), uintptr(height))
	return int(r)
}

// 窗口_测试点击子元素.
//
// hWindow: 窗口句柄.
//
// pPt: 左边点.
func XWnd_HitChildEle(hWindow int, pPt *POINT) int {
	r, _, _ := xWnd_HitChildEle.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 窗口_取子对象数量.
//
// hWindow: 窗口句柄.
func XWnd_GetChildCount(hWindow int) int {
	r, _, _ := xWnd_GetChildCount.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取子对象从索引.
//
// hWindow: 窗口句柄.
//
// index: 元素索引.
func XWnd_GetChildByIndex(hWindow, index int) int {
	r, _, _ := xWnd_GetChildByIndex.Call(uintptr(hWindow), uintptr(index))
	return int(r)
}

// 窗口_取子对象从ID.
//
// hWindow: 窗口句柄.
//
// nID: 元素ID, ID必须大于0.
func XWnd_GetChildByID(hWindow, nID int) int {
	r, _, _ := xWnd_GetChildByID.Call(uintptr(hWindow), uintptr(nID))
	return int(r)
}

// 窗口_取子对象.
//
// hWindow: 窗口句柄.
//
// nID: 对象ID,ID必须大于0.
func XWnd_GetChild(hWindow, nID int) int {
	r, _, _ := xWnd_GetChild.Call(uintptr(hWindow), uintptr(nID))
	return int(r)
}

// 窗口_关闭.
//
// hWindow: 窗口句柄.
func XWnd_CloseWindow(hWindow int) int {
	r, _, _ := xWnd_CloseWindow.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_调整布局.
//
// hWindow: 窗口句柄.
func XWnd_AdjustLayout(hWindow int) int {
	r, _, _ := xWnd_AdjustLayout.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_调整布局扩展.
//
// hWindow: 窗口句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
func XWnd_AdjustLayoutEx(hWindow int, nFlags xcc.AdjustLayout_) int {
	r, _, _ := xWnd_AdjustLayoutEx.Call(uintptr(hWindow), uintptr(nFlags))
	return int(r)
}

// 窗口_创建插入符.
//
// hWindow: 窗口句柄.
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
func XWnd_CreateCaret(hWindow, hEle, x, y, width, height int) int {
	r, _, _ := xWnd_CreateCaret.Call(uintptr(hWindow), uintptr(hEle), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return int(r)
}

// 窗口_置插入符位置, 设置插入符位置.
//
// hWindow: 窗口句柄.
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
func XWnd_SetCaretPos(hWindow, x, y, width, height int, bUpdate bool) int {
	r, _, _ := xWnd_SetCaretPos.Call(uintptr(hWindow), uintptr(x), uintptr(y), uintptr(width), uintptr(height), common.BoolPtr(bUpdate))
	return int(r)
}

// 窗口_置插入符颜色.
//
// hWindow: 窗口句柄.
//
// color: 颜色值.
func XWnd_SetCaretColor(hWindow, color int) int {
	r, _, _ := xWnd_SetCaretColor.Call(uintptr(hWindow), uintptr(color))
	return int(r)
}

// 窗口_显示插入符.
//
// hWindow: 窗口句柄.
//
// bShow: 是否显示.
func XWnd_ShowCaret(hWindow int, bShow bool) int {
	r, _, _ := xWnd_ShowCaret.Call(uintptr(hWindow), common.BoolPtr(bShow))
	return int(r)
}

// 窗口_销毁插入符.
//
// hWindow: 窗口句柄.
func XWnd_DestroyCaret(hWindow int) int {
	r, _, _ := xWnd_DestroyCaret.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取插入符元素.
//
// hWindow: 窗口句柄.
func XWnd_GetCaretHELE(hWindow int) int {
	r, _, _ := xWnd_GetCaretHELE.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取客户区坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.
func XWnd_GetClientRect(hWindow int, pRect *RECT) int {
	r, _, _ := xWnd_GetClientRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 窗口_取Body坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.
func XWnd_GetBodyRect(hWindow int, pRect *RECT) int {
	r, _, _ := xWnd_GetBodyRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 窗口_取布局坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 接收返回坐标.
func XWnd_GetLayoutRect(hWindow int, pRect *RECT) int {
	r, _, _ := xWnd_GetLayoutRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 窗口_移动窗口.
//
// hWindow: 窗口句柄.
//
// x: X坐标.
//
// y: Y坐标.
func XWnd_SetPosition(hWindow int, x, y int32) {
	xWnd_SetPosition.Call(uintptr(hWindow), uintptr(x), uintptr(y))
}

// 窗口_取坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.
func XWnd_GetRect(hWindow int, pRect *RECT) int {
	r, _, _ := xWnd_GetRect.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 窗口_最大化.
//
// hWindow: 窗口句柄.
//
// bMaximize: 是否最大化.
func XWnd_MaxWindow(hWindow int, bMaximize bool) int {
	r, _, _ := xWnd_MaxWindow.Call(uintptr(hWindow), common.BoolPtr(bMaximize))
	return int(r)
}

// 窗口_置定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func XWnd_SetTimer(hWindow, nIDEvent, uElapse int) int {
	r, _, _ := xWnd_SetTimer.Call(uintptr(hWindow), uintptr(nIDEvent), uintptr(uElapse))
	return int(r)
}

// 窗口_关闭定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
func XWnd_KillTimer(hWindow, nIDEvent int) int {
	r, _, _ := xWnd_KillTimer.Call(uintptr(hWindow), uintptr(nIDEvent))
	return int(r)
}

// 窗口_置炫彩定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func XWnd_SetXCTimer(hWindow, nIDEvent, uElapse int) int {
	r, _, _ := xWnd_SetXCTimer.Call(uintptr(hWindow), uintptr(nIDEvent), uintptr(uElapse))
	return int(r)
}

// 窗口_关闭炫彩定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
func XWnd_KillXCTimer(hWindow, nIDEvent int) int {
	r, _, _ := xWnd_KillXCTimer.Call(uintptr(hWindow), uintptr(nIDEvent))
	return int(r)
}

// 窗口_取背景管理器.
//
// hWindow: 窗口句柄.
func XWnd_GetBkManager(hWindow int) int {
	r, _, _ := xWnd_GetBkManager.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取背景管理器扩展.
//
// hWindow: 窗口句柄.
func XWnd_GetBkManagerEx(hWindow int) int {
	r, _, _ := xWnd_GetBkManagerEx.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_置背景管理器.
//
// hWindow: 窗口句柄.
//
// hBkInfoM: 背景管理器.
func XWnd_SetBkMagager(hWindow, hBkInfoM int) int {
	r, _, _ := xWnd_SetBkMagager.Call(uintptr(hWindow), uintptr(hBkInfoM))
	return int(r)
}

// 窗口_置透明类型.
//
// hWindow: 窗口句柄.
//
// nType: 窗口透明类型: xcc.Window_Transparent_.
func XWnd_SetTransparentType(hWindow int, nType xcc.Window_Transparent_) int {
	r, _, _ := xWnd_SetTransparentType.Call(uintptr(hWindow), uintptr(nType))
	return int(r)
}

// 窗口_置透明度.
//
// hWindow: 窗口句柄.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func XWnd_SetTransparentAlpha(hWindow int, alpha byte) int {
	r, _, _ := xWnd_SetTransparentAlpha.Call(uintptr(hWindow), uintptr(alpha))
	return int(r)
}

// 窗口_置透明色.
//
// hWindow: 窗口句柄.
//
// color: 窗口透明色.
func XWnd_SetTransparentColor(hWindow, color int) int {
	r, _, _ := xWnd_SetTransparentColor.Call(uintptr(hWindow), uintptr(color))
	return int(r)
}

// 窗口_置阴影信息.
//
// hWindow: 窗口句柄.
//
// nSize: 阴影大小.
//
// nDepth: 阴影深度, 0-255.
//
// nAngeleSize: 圆角阴影内收大小.
//
// bRightAngle: 是否强制直角.
//
// color: 阴影颜色.
func XWnd_SetShadowInfo(hWindow, nSize int, nDepth byte, nAngeleSize int, bRightAngle bool, color int) int {
	r, _, _ := xWnd_SetShadowInfo.Call(uintptr(hWindow), uintptr(nSize), uintptr(nDepth), uintptr(nAngeleSize), common.BoolPtr(bRightAngle), uintptr(color))
	return int(r)
}

// 窗口_取阴影信息.
//
// hWindow: 窗口句柄.
//
// pnSize: 阴影大小.
//
// pnDepth: 阴影深度.
//
// pnAngeleSize: 圆角阴影内收大小 .
//
// pbRightAngle: 是否强制直角.
//
// pColor: 阴影颜色.
func XWnd_GetShadowInfo(hWindow int, pnSize, pnDepth, pnAngeleSize *int32, pbRightAngle *bool, pColor *int) int {
	r, _, _ := xWnd_GetShadowInfo.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pnSize)), uintptr(unsafe.Pointer(pnDepth)), uintptr(unsafe.Pointer(pnAngeleSize)), uintptr(unsafe.Pointer(pbRightAngle)), uintptr(unsafe.Pointer(pColor)))
	return int(r)
}

// 窗口_取透明类型, 返回: xcc.Window_Transparent_.
//
// hWindow: 窗口句柄.
func XWnd_GetTransparentType(hWindow int) xcc.Window_Transparent_ {
	r, _, _ := xWnd_GetTransparentType.Call(uintptr(hWindow))
	return xcc.Window_Transparent_(r)
}

// 窗口_附加窗口, 返回窗口资源句柄.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func XWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xWnd_Attach.Call(hWnd, uintptr(XCStyle))
	return int(r)
}

// 窗口_启用拖放文件.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.
func XWnd_EnableDragFiles(hWindow int, bEnable bool) bool {
	r, _, _ := xWnd_EnableDragFiles.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return r != 0
}

// 窗口_显示 显示隐藏窗口.
//
// hWindow: 窗口句柄.
//
// bShow: 是否显示.
func XWnd_Show(hWindow int, bShow bool) int {
	r, _, _ := xWnd_Show.Call(uintptr(hWindow), common.BoolPtr(bShow))
	return int(r)
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
func XWnd_GetCaretInfo(hWindow int, pX, pY, pWidth, pHeight *int32) int {
	r, _, _ := xWnd_GetCaretInfo.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)), uintptr(unsafe.Pointer(pWidth)), uintptr(unsafe.Pointer(pHeight)))
	return int(r)
}

// 窗口_置图标.
//
// hWindow: 窗口句柄.
//
// hImage: 图标句柄.
func XWnd_SetIcon(hWindow, hImage int) int {
	r, _, _ := xWnd_SetIcon.Call(uintptr(hWindow), uintptr(hImage))
	return int(r)
}

// 窗口_置标题.
//
// hWindow: 窗口句柄.
//
// pTitle: 标题文本.
func XWnd_SetTitle(hWindow int, pTitle string) int {
	r, _, _ := xWnd_SetTitle.Call(uintptr(hWindow), common.StrPtr(pTitle))
	return int(r)
}

// 窗口_置标题颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色.
func XWnd_SetTitleColor(hWindow, color int) int {
	r, _, _ := xWnd_SetTitleColor.Call(uintptr(hWindow), uintptr(color))
	return int(r)
}

// 窗口_取控制按钮, 返回按钮句柄.
//
// hWindow: 窗口句柄.
//
// nFlag: xcc.Window_Style_ . 可用值: xcc.Window_Style_Btn_Min , xcc.Window_Style_Btn_Max , xcc.Window_Style_Btn_Close .
func XWnd_GetButton(hWindow int, nFlag xcc.Window_Style_) int {
	r, _, _ := xWnd_GetButton.Call(uintptr(hWindow), uintptr(nFlag))
	return int(r)
}

// 窗口_取图标, 返回图标句柄.
//
// hWindow: 窗口句柄.
func XWnd_GetIcon(hWindow int) int {
	r, _, _ := xWnd_GetIcon.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_取标题, 返回标题文本.
//
// hWindow: 窗口句柄.
func XWnd_GetTitle(hWindow int) string {
	r, _, _ := xWnd_GetTitle.Call(uintptr(hWindow))
	return common.UintPtrToString(r)
}

// 窗口_取标题颜色, 返回ABGR 颜色.
//
// hWindow: 窗口句柄.
func XWnd_GetTitleColor(hWindow int) int {
	r, _, _ := xWnd_GetTitleColor.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_添加背景边框.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func XWnd_AddBkBorder(hWindow int, nState xcc.Window_State_Flag_, color int, width int) int {
	r, _, _ := xWnd_AddBkBorder.Call(uintptr(hWindow), uintptr(nState), uintptr(color), uintptr(width))
	return int(r)
}

// 窗口_添加背景填充.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func XWnd_AddBkFill(hWindow int, nState xcc.Window_State_Flag_, color int) int {
	r, _, _ := xWnd_AddBkFill.Call(uintptr(hWindow), uintptr(nState), uintptr(color))
	return int(r)
}

// 窗口_添加背景图片.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func XWnd_AddBkImage(hWindow int, nState xcc.Window_State_Flag_, hImage int) int {
	r, _, _ := xWnd_AddBkImage.Call(uintptr(hWindow), uintptr(nState), uintptr(hImage))
	return int(r)
}

// 窗口_取背景对象数量.
//
// hWindow: 窗口句柄.
func XWnd_GetBkInfoCount(hWindow int) int {
	r, _, _ := xWnd_GetBkInfoCount.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_清空背景对象.
//
// hWindow: 窗口句柄.
func XWnd_ClearBkInfo(hWindow int) int {
	r, _, _ := xWnd_ClearBkInfo.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// hWindow: 窗口句柄.
//
// pText: 背景内容字符串.
func XWnd_SetBkInfo(hWindow int, pText string) int {
	r, _, _ := xWnd_SetBkInfo.Call(uintptr(hWindow), common.StrPtr(pText))
	return int(r)
}

// 窗口_是否可拖动标题栏.
//
// hWindow: 窗口句柄.
func XWnd_IsDragCaption(hWindow int) bool {
	r, _, _ := xWnd_IsDragCaption.Call(uintptr(hWindow))
	return r != 0
}

// 窗口_是否可拖动窗口.
//
// hWindow: 窗口句柄.
func XWnd_IsDragWindow(hWindow int) bool {
	r, _, _ := xWnd_IsDragWindow.Call(uintptr(hWindow))
	return r != 0
}

// 窗口_是否可拖动边框.
//
// hWindow: 窗口句柄.
func XWnd_IsDragBorder(hWindow int) bool {
	r, _, _ := xWnd_IsDragBorder.Call(uintptr(hWindow))
	return r != 0
}

// 窗口_置标题外间距, 设置标题内容(图标, 标题, 控制按钮)外间距.
//
// hWindow: 窗口句柄.
//
// left: 左边间距.
//
// top: 上边间距.
//
// right: 右边间距.
//
// bottom: 下边间距.
func XWnd_SetCaptionMargin(hWindow int, left int, top int, right int, bottom int) int {
	r, _, _ := xWnd_SetCaptionMargin.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口_置窗口位置. 封装系统API SetWindowPos(), 内部做了DPI适配.
//
// hWindow: 窗口句柄.
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
func XWnd_SetWindowPos(hWindow int, hWndInsertAfter xcc.HWND_, x, y, cx, cy int32, uFlags xcc.SWP_) int {
	r, _, _ := xWnd_SetWindowPos.Call(uintptr(hWindow), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return int(r)
}

// 窗口_取DPI. 获取当前窗口所在显示器DPI, 返回窗口DPI.
//
// hWindow: 窗口句柄.
func XWnd_GetDPI(hWindow int) int {
	r, _, _ := xWnd_GetDPI.Call(uintptr(hWindow))
	return int(r)
}

// 窗口_坐标转换DPI. 窗口客户区坐标转换到缩放后DPI坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 接收返回坐标.
func XWnd_RectToDPI(hWindow int, pRect *RECT) int {
	r, _, _ := xWnd_RectToDPI.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 窗口_坐标点转换DPI. 窗口客户区坐标点转换到缩放后.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.
func XWnd_PointToDPI(hWindow int, pPt *POINT) int {
	r, _, _ := xWnd_PointToDPI.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pPt)))
	return int(r)
}

// 窗口_取光标位置. 封装的系统API: GetCursorPos, 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.
func XWnd_GetCursorPos(hWindow int, pPt *POINT) bool {
	r, _, _ := xWnd_GetCursorPos.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pPt)))
	return r != 0
}

// 窗口_客户区坐标点到屏幕坐标点. 封装的系统API: ClientToScreen, 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.
func XWnd_ClientToScreen(hWindow int, pPt *POINT) bool {
	r, _, _ := xWnd_ClientToScreen.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pPt)))
	return r != 0
}

// 窗口_屏幕坐标点到客户区坐标点. 封装的系统API: ScreenToClient(), 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.
func XWnd_ScreenToClient(hWindow int, pPt *POINT) bool {
	r, _, _ := xWnd_ScreenToClient.Call(uintptr(hWindow), uintptr(unsafe.Pointer(pPt)))
	return r != 0
}

// 窗口_置DPI. 设置当前窗口DPI, 默认DPI为96.
//
// hWindow: 窗口句柄.
//
// nDPI: DPI值.
func XWnd_SetDPI(hWindow int, nDPI int) int {
	r, _, _ := xWnd_SetDPI.Call(uintptr(hWindow), uintptr(nDPI))
	return int(r)
}
