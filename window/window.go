package window

import (
	"github.com/twgh/xcgui/xc"
)

type Window struct {
	WindowPublic

	HWindow int
	HWND    int
}

// 窗口_创建, 返回: GUI库窗口资源句柄
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// pTitle: 窗口标题.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式, Xc_Window_Style_
func NewWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle int) *Window {
	p := &Window{
		HWindow: xc.XFrameWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle),
	}
	p.HWindow_ = p.HWindow
	p.HWND = p.GetHWND()
	return p
}

// 窗口_创建扩展, 返回: GUI库窗口资源句柄
// dwExStyle: 窗口扩展样式.
// lpClassName: 窗口类名.
// lpWindowName: 窗口名.
// dwStyle: 窗口样式
// x: 窗口左上角x坐标.
// y: 窗口左上角y坐标.
// cx: 窗口宽度.
// cy: 窗口高度.
// hWndParent: 父窗口.
// XCStyle: GUI库窗口样式, Xc_Window_Style_
func NewWindowEx(dwExStyle int, lpClassName string, lpWindowName string, dwStyle int, x int, y int, cx int, cy int, hWndParent int, XCStyle int) *Window {
	p := &Window{
		HWindow: xc.XWnd_CreateEx(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, cx, cy, hWndParent, XCStyle),
	}
	return p
}

// 窗口_显示
// nCmdShow: 显示方式, SW_
func (w *Window) ShowWindow(nCmdShow int) int {
	return xc.XWnd_ShowWindow(w.HWindow, nCmdShow)
}

// 窗口_置顶
func (w *Window) SetTop() int {
	return xc.XWnd_SetTop(w.HWindow)
}

// 窗口_注册事件C
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *Window) RegEventC(nEvent int, pFun int) bool {
	return xc.XWnd_RegEventC(w.HWindow, nEvent, pFun)
}

// 窗口_注册事件C1
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *Window) RegEventC1(nEvent int, pFun int) bool {
	return xc.XWnd_RegEventC1(w.HWindow, nEvent, pFun)
}

// 窗口_移除事件C
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *Window) RemoveEventC(nEvent int, pFun int) bool {
	return xc.XWnd_RemoveEventC(w.HWindow, nEvent, pFun)
}

// 窗口_添加子对象
// hChild: 要添加的对象句柄.
func (w *Window) AddChild(hChild int) bool {
	return xc.XWnd_AddChild(w.HWindow, hChild)
}

// 窗口_插入子对象
// hChild: 要插入的对象句柄.
// index: 插入位置索引.
func (w *Window) InsertChild(hChild int, index int) bool {
	return xc.XWnd_InsertChild(w.HWindow, hChild, index)
}

// 窗口_取HWND
func (w *Window) GetHWND() int {
	return xc.XWnd_GetHWND(w.HWindow)
}

// 窗口_重绘
// bImmediate: 是否立即重绘, 默认为否.
func (w *Window) Redraw(bImmediate bool) int {
	return xc.XWnd_Redraw(w.HWindow, bImmediate)
}

// 窗口_重绘指定区域
// pRect: 需要重绘的区域坐标.
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *Window) RedrawRect(pRect *xc.RECT, bImmediate bool) int {
	return xc.XWnd_RedrawRect(w.HWindow, pRect, bImmediate)
}

// 窗口_置坐标
// pRect: 坐标
func (w *Window) SetRect(pRect *xc.RECT) bool {
	return xc.XWnd_SetRect(w.HWindow, pRect)
}

// 窗口_置焦点
// hFocusEle: 将获得焦点的元素.
func (w *Window) SetFocusEle(hFocusEle int) bool {
	return xc.XWnd_SetFocusEle(w.HWindow, hFocusEle)
}

// 窗口_取焦点
func (w *Window) GetFocusEle() int {
	return xc.XWnd_GetFocusEle(w.HWindow)
}

// 窗口_取鼠标停留元素
func (w *Window) GetStayEle() int {
	return xc.XWnd_GetStayEle(w.HWindow)
}

// 窗口_绘制
// hDraw: 图形绘制句柄.
func (w *Window) DrawWindow(hDraw int) int {
	return xc.XWnd_DrawWindow(w.HWindow, hDraw)
}

// 窗口_居中
func (w *Window) Center() int {
	return xc.XWnd_Center(w.HWindow)
}

// 窗口_居中扩展
// width: 窗口宽度
// height: 窗口高度
func (w *Window) CenterEx(width, height int) int {
	return xc.XWnd_CenterEx(w.HWindow, width, height)
}

// 窗口_置光标
// hCursor: 鼠标光标句柄.
func (w *Window) SetCursor(hCursor int) int {
	return xc.XWnd_SetCursor(w.HWindow, hCursor)
}

// 窗口_取光标
func (w *Window) GetCursor() int {
	return xc.XWnd_GetCursor(w.HWindow)
}

// 窗口_启用拖动边框
// bEnable: 是否启用.
func (w *Window) EnableDragBorder(bEnable bool) int {
	return xc.XWnd_EnableDragBorder(w.HWindow, bEnable)
}

// 窗口_启用拖动窗口
// bEnable: 是否启用.
func (w *Window) EnableDragWindow(bEnable bool) int {
	return xc.XWnd_EnableDragWindow(w.HWindow, bEnable)
}

// 窗口_启用拖动标题栏
// bEnable: 是否启用.
func (w *Window) EnableDragCaption(bEnable bool) int {
	return xc.XWnd_EnableDragCaption(w.HWindow, bEnable)
}

// 窗口_启用绘制背景
// bEnable: 是否启用.
func (w *Window) EnableDrawBk(bEnable bool) int {
	return xc.XWnd_EnableDrawBk(w.HWindow, bEnable)
}

// 窗口_启用自动焦点
// bEnable: 是否启用.
func (w *Window) EnableAutoFocus(bEnable bool) int {
	return xc.XWnd_EnableAutoFocus(w.HWindow, bEnable)
}

// 窗口_启用允许最大化
// bEnable: 是否启用.
func (w *Window) EnableMaxWindow(bEnable bool) int {
	return xc.XWnd_EnableMaxWindow(w.HWindow, bEnable)
}

// 窗口_启用限制窗口大小
// bEnable: 是否启用
func (w *Window) EnablemLimitWindowSize(bEnable bool) int {
	return xc.XWnd_EnablemLimitWindowSize(w.HWindow, bEnable)
}

// 窗口_启用布局
// bEnable: 是否启用.
func (w *Window) EnableLayout(bEnable bool) int {
	return xc.XWnd_EnableLayout(w.HWindow, bEnable)
}

// 窗口_启用布局覆盖边框
// bEnable: 是否启用
func (w *Window) EnableLayoutOverlayBorder(bEnable bool) int {
	return xc.XWnd_EnableLayoutOverlayBorder(w.HWindow, bEnable)
}

// 窗口_显示布局边界
// bEnable: 是否启用.
func (w *Window) ShowLayoutFrame(bEnable bool) int {
	return xc.XWnd_ShowLayoutFrame(w.HWindow, bEnable)
}

// 窗口_判断启用布局
func (w *Window) IsEnableLayout() bool {
	return xc.XWnd_IsEnableLayout(w.HWindow)
}

// 窗口_是否最大化
func (w *Window) IsMaxWindow() bool {
	return xc.XWnd_IsMaxWindow(w.HWindow)
}

// 窗口_置鼠标捕获元素
// hEle: 元素句柄.
func (w *Window) SetCaptureEle(hEle int) int {
	return xc.XWnd_SetCaptureEle(w.HWindow, hEle)
}

// 窗口_取鼠标捕获元素
func (w *Window) GetCaptureEle() int {
	return xc.XWnd_GetCaptureEle(w.HWindow)
}

// 窗口_取绘制矩形
// pRcPaint: 重绘区域坐标.
func (w *Window) GetDrawRect(pRcPaint *xc.RECT) int {
	return xc.XWnd_GetDrawRect(w.HWindow, pRcPaint)
}

// 窗口_置系统光标
// hCursor: 光标句柄.
func (w *Window) SetCursorSys(hCursor int) int {
	return xc.XWnd_SetCursorSys(w.HWindow, hCursor)
}

// 窗口_置字体
// hFontx: 炫彩字体句柄.
func (w *Window) SetFont(hFontx int) int {
	return xc.XWnd_SetFont(w.HWindow, hFontx)
}

// 窗口_置文本颜色
// color: RGB颜色值.
// alpha: 透明度.
func (w *Window) SetTextColor(color int, alpha uint8) int {
	return xc.XWnd_SetTextColor(w.HWindow, color, alpha)
}

// 窗口_取文本颜色
func (w *Window) GetTextColor() int {
	return xc.XWnd_GetTextColor(w.HWindow)
}

// 窗口_取文本颜色扩展
func (w *Window) GetTextColorEx() int {
	return xc.XWnd_GetTextColorEx(w.HWindow)
}

// 窗口_置ID
// nID: ID值.
func (w *Window) SetID(nID int) int {
	return xc.XWnd_SetID(w.HWindow, nID)
}

// 窗口_取ID
func (w *Window) GetID() int {
	return xc.XWnd_GetID(w.HWindow)
}

// 窗口_置名称
// pName: name值, 字符串.
func (w *Window) SetName(pName string) int {
	return xc.XWnd_SetName(w.HWindow, pName)
}

// 窗口_取名称
func (w *Window) GetName() string {
	return xc.XWnd_GetName(w.HWindow)
}

// 窗口_置边大小
// left: 窗口左边大小.
// top: 窗口上边大小.
// right: 窗口右边大小.
// bottom: 窗口底部大小.
func (w *Window) SetBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetBorderSize(w.HWindow, left, top, right, bottom)
}

// 窗口_取边大小
func (w *Window) GetBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetBorderSize(w.HWindow, pBorder)
}

// 窗口_置布局内填充大小
// left: 左边大小.
// top: 上边大小.
// right: 右边大小.
// bottom: 下边大小.
func (w *Window) SetPadding(left, top, right, bottom int) int {
	return xc.XWnd_SetPadding(w.HWindow, left, top, right, bottom)
}

// 窗口_置拖动边框大小
// left: 窗口左边大小.
// top: 窗口上边大小.
// right: 窗口右边大小.
// bottom: 窗口底边大小.
func (w *Window) SetDragBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetDragBorderSize(w.HWindow, left, top, right, bottom)
}

// 窗口_取拖动边框大小
// pSize: 拖动边框大小.
func (w *Window) GetDragBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetDragBorderSize(w.HWindow, pBorder)
}

// 窗口_置最小大小
// width: 最小宽度.
// height: 最小高度.
func (w *Window) SetMinimumSize(width, height int) int {
	return xc.XWnd_SetMinimumSize(w.HWindow, width, height)
}

// 窗口_测试点击子元素
// pPt: 左边点.
func (w *Window) HitChildEle(pPt *xc.POINT) int {
	return xc.XWnd_HitChildEle(w.HWindow, pPt)
}

// 窗口_取子对象数量
func (w *Window) GetChildCount() int {
	return xc.XWnd_GetChildCount(w.HWindow)
}

// 窗口_取子对象从索引
// index: 元素索引.
func (w *Window) GetChildByIndex(index int) int {
	return xc.XWnd_GetChildByIndex(w.HWindow, index)
}

// 窗口_取子对象从ID
// nID: 元素ID, ID必须大于0.
func (w *Window) GetChildByID(nID int) int {
	return xc.XWnd_GetChildByID(w.HWindow, nID)
}

// 窗口_取子对象
// nID: 对象ID,ID必须大于0.
func (w *Window) GetChild(nID int) int {
	return xc.XWnd_GetChild(w.HWindow, nID)
}

// 窗口_关闭
func (w *Window) CloseWindow() int {
	return xc.XWnd_CloseWindow(w.HWindow)
}

// 窗口_调整布局
func (w *Window) AdjustLayout() int {
	return xc.XWnd_AdjustLayout(w.HWindow)
}

// 窗口_调整布局扩展
// nFlags: 调整标识, AdjustLayout_
func (w *Window) AdjustLayoutEx(nFlags int) int {
	return xc.XWnd_AdjustLayoutEx(w.HWindow, nFlags)
}

// 窗口_创建插入符
// hEle: 元素句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (w *Window) CreateCaret(hEle, x, y, width, height int) int {
	return xc.XWnd_CreateCaret(w.HWindow, hEle, x, y, width, height)
}

// 窗口_置插入符大小
// width: 宽度.
// height: 高度.
func (w *Window) SetCaretSize(width, height int) int {
	return xc.XWnd_SetCaretSize(w.HWindow, width, height)
}

// 窗口_置插入符位置
// x: x坐标.
// y: y坐标.
func (w *Window) SetCaretPos(x, y int) int {
	return xc.XWnd_SetCaretPos(w.HWindow, x, y)
}

// 窗口_置插入符位置扩展
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (w *Window) SetCaretPosEx(x, y, width, height int) int {
	return xc.XWnd_SetCaretPosEx(w.HWindow, x, y, width, height)
}

// 窗口_置插入符颜色
// color: 颜色值.
func (w *Window) SetCaretColor(color int) int {
	return xc.XWnd_SetCaretColor(w.HWindow, color)
}

// 窗口_显示插入符
// bShow: 是否显示.
func (w *Window) ShowCaret(bShow bool) int {
	return xc.XWnd_ShowCaret(w.HWindow, bShow)
}

// 窗口_销毁插入符
func (w *Window) DestroyCaret() int {
	return xc.XWnd_DestroyCaret(w.HWindow)
}

// 窗口_取插入符元素
func (w *Window) GetCaretHELE() int {
	return xc.XWnd_GetCaretHELE(w.HWindow)
}

// 窗口_取客户区坐标
// pRect: 坐标.
func (w *Window) GetClientRect(pRect *xc.RECT) int {
	return xc.XWnd_GetClientRect(w.HWindow, pRect)
}

// 窗口_取Body坐标
// pRect: 坐标.
func (w *Window) GetBodyRect(pRect *xc.RECT) int {
	return xc.XWnd_GetBodyRect(w.HWindow, pRect)
}

// 窗口_取布局坐标
// pRect: 接收返回坐标
func (w *Window) GetLayoutRect(pRect *xc.RECT) int {
	return xc.XWnd_GetLayoutRect(w.HWindow, pRect)
}

// 窗口_移动窗口
// x: X坐标
// y: Y坐标
func (w *Window) Move(x, y int) int {
	return xc.XWnd_Move(w.HWindow, x, y)
}

// 窗口_取坐标
// pRect: 坐标
func (w *Window) GetRect(pRect *xc.RECT) int {
	return xc.XWnd_GetRect(w.HWindow, pRect)
}

// 窗口_最大化
// bMaximize: 是否最大化
func (w *Window) MaxWindow(bMaximize bool) int {
	return xc.XWnd_MaxWindow(w.HWindow, bMaximize)
}

// 窗口_置定时器
// nIDEvent: 定时器ID.
// uElapse: 间隔值, 单位毫秒.
func (w *Window) SetTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetTimer(w.HWindow, nIDEvent, uElapse)
}

// 窗口_关闭定时器
// nIDEvent: 定时器ID.
func (w *Window) KillTimer(nIDEvent int) int {
	return xc.XWnd_KillTimer(w.HWindow, nIDEvent)
}

// 窗口_置炫彩定时器
// nIDEvent: 定时器ID.
// uElapse: 间隔值, 单位毫秒.
func (w *Window) SetXCTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetXCTimer(w.HWindow, nIDEvent, uElapse)
}

// 窗口_关闭炫彩定时器
// nIDEvent: 定时器ID.
func (w *Window) KillXCTimer(nIDEvent int) int {
	return xc.XWnd_KillXCTimer(w.HWindow, nIDEvent)
}

// 窗口_取背景管理器
func (w *Window) GetBkManager() int {
	return xc.XWnd_GetBkManager(w.HWindow)
}

// 窗口_取背景管理器扩展
func (w *Window) GetBkManagerEx() int {
	return xc.XWnd_GetBkManagerEx(w.HWindow)
}

// 窗口_置背景管理器
// hBkInfoM: 背景管理器
func (w *Window) SetBkMagager(hBkInfoM int) int {
	return xc.XWnd_SetBkMagager(w.HWindow, hBkInfoM)
}

// 窗口_置透明类型
// nType: 窗口透明类型.
func (w *Window) SetTransparentType(nType int) int {
	return xc.XWnd_SetTransparentType(w.HWindow, nType)
}

// 窗口_置透明度
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *Window) SetTransparentAlpha(alpha uint8) int {
	return xc.XWnd_SetTransparentAlpha(w.HWindow, alpha)
}

// 窗口_置透明色
// color: 窗口透明色.
func (w *Window) SetTransparentColor(color int) int {
	return xc.XWnd_SetTransparentColor(w.HWindow, color)
}

// 窗口_置阴影信息
// nSize: 阴影大小
// nDepth: 阴影深度, 0-255.
// nAngeleSize: 圆角阴影内收大小.
// bRightAngle: 是否强制直角.
// color: 阴影颜色.
func (w *Window) SetShadowInfo(nSize int, nDepth uint8, nAngeleSize int, bRightAngle bool, color int) int {
	return xc.XWnd_SetShadowInfo(w.HWindow, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取阴影信息
// pnSize: 阴影大小.
// pnDepth: 阴影深度.
// pnAngeleSize: 圆角阴影内收大小 .
// pbRightAngle: 是否强制直角.
// pColor: 阴影颜色.
func (w *Window) GetShadowInfo(nSize *int, nDepth *uint8, nAngeleSize *int, bRightAngle *bool, color *int) int {
	return xc.XWnd_GetShadowInfo(w.HWindow, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取透明类型
func (w *Window) GetTransparentType() int {
	return xc.XWnd_GetTransparentType(w.HWindow)
}
