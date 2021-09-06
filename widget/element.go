package widget

import (
	xc "github.com/twgh/xcgui"
)

type Element struct {
	HEle_ int
}

// 元素_创建, 创建基础元素
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父为窗口句柄或元素句柄.
func NewElement(x int, y int, cx int, cy int, hParent int) *Element {
	p := &Element{
		HEle_: xc.XEle_Create(x, y, cx, cy, hParent),
	}
	return p
}

// 元素_注册事件C, 注册事件C方式, 省略2参数
// nEvent: 事件类型, XE_
// pFun: 事件函数指针.
func (e *Element) RegEventC(nEvent int, pFun interface{}) bool {
	return xc.XEle_RegEventC(e.HEle_, nEvent, pFun)
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数
// nEvent: 事件类型, XE_
// pFun: 事件函数指针.
func (e *Element) RegEventC1(nEvent int, pFun interface{}) bool {
	return xc.XEle_RegEventC1(e.HEle_, nEvent, pFun)
}

// 元素_移除事件C
// nEvent: 事件类型, XE_
// pFun: 事件函数指针.
func (e *Element) RemoveEventC(nEvent int, pFun interface{}) bool {
	return xc.XEle_RemoveEventC(e.HEle_, nEvent, pFun)
}

// 元素_发送事件
// nEvent: 事件类型, XE_
// wParam: 参数.
// lParam: 参数.
func (e *Element) SendEvent(nEvent int, wParam int, lParam int) int {
	return xc.XEle_SendEvent(e.HEle_, nEvent, wParam, lParam)
}

// 元素_投递事件
// nEvent: 事件类型, XE_
// wParam: 参数.
// lParam: 参数.
func (e *Element) PostEvent(nEvent int, wParam int, lParam int) int {
	return xc.XEle_PostEvent(e.HEle_, nEvent, wParam, lParam)
}

// 元素_取坐标
// pRect: 坐标.
func (e *Element) GetRect(pRect *xc.RECT) int {
	return xc.XEle_GetRect(e.HEle_, pRect)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移
// pRect: 坐标.
func (e *Element) GetRectLogic(pRect *xc.RECT) int {
	return xc.XEle_GetRectLogic(e.HEle_, pRect)
}

// 元素_取客户区坐标
// pRect: 坐标.
func (e *Element) GetClientRect(pRect *xc.RECT) int {
	return xc.XEle_GetClientRect(e.HEle_, pRect)
}

// 元素_置宽度
// nWidth: 宽度
func (e *Element) SetWidth(nWidth int) int {
	return xc.XEle_SetWidth(e.HEle_, nWidth)
}

// 元素_置高度
// nHeight: 高度
func (e *Element) SetHeight(nHeight int) int {
	return xc.XEle_SetHeight(e.HEle_, nHeight)
}

// 元素_取宽度
func (e *Element) GetWidth() int {
	return xc.XEle_GetWidth(e.HEle_)
}

// 元素_取高度
func (e *Element) GetHeight() int {
	return xc.XEle_GetHeight(e.HEle_)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标
// pRect: 坐标.
func (e *Element) RectWndClientToEleClient(pRect *xc.RECT) int {
	return xc.XEle_RectWndClientToEleClient(e.HEle_, pRect)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标
// pPt: 坐标.
func (e *Element) PointWndClientToEleClient(pPt *xc.POINT) int {
	return xc.XEle_PointWndClientToEleClient(e.HEle_, pPt)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标
// pRect: 坐标.
func (e *Element) RectClientToWndClient(pRect *xc.RECT) int {
	return xc.XEle_RectClientToWndClient(e.HEle_, pRect)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标
// pPt: 坐标.
func (e *Element) PointClientToWndClient(pPt *xc.POINT) int {
	return xc.XEle_PointClientToWndClient(e.HEle_, pPt)
}

// 元素_基于窗口客户区坐标
// pRect: 坐标.
func (e *Element) GetWndClientRect(pRect *xc.RECT) int {
	return xc.XEle_GetWndClientRect(e.HEle_, pRect)
}

// 元素_取类型
func (e *Element) GetType() int {
	return xc.XEle_GetType(e.HEle_)
}

// 元素_取HWND
func (e *Element) GetHWND() int {
	return xc.XEle_GetHWND(e.HEle_)
}

// 元素_取HWINDOW
func (e *Element) GetHWINDOW() int {
	return xc.XEle_GetHWINDOW(e.HEle_)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄
func (e *Element) GetCursor() int {
	return xc.XEle_GetCursor(e.HEle_)
}

// 元素_置光标, 设置元素鼠标光标
// hCursor: 光标句柄.
func (e *Element) SetCursor(hCursor int) int {
	return xc.XEle_SetCursor(e.HEle_, hCursor)
}

// 元素_添加子对象
// hChild: 要添加的子元素句柄或形状对象句柄.
func (e *Element) AddChild(hChild int) bool {
	return xc.XEle_AddChild(e.HEle_, hChild)
}

// 元素_插入子对象, 插入子对象到指定位置
// hChild: 要插入的元素句柄或形状对象句柄.
// index: 插入位置索引.
func (e *Element) InsertChild(hChild int, index int) bool {
	return xc.XEle_InsertChild(e.HEle_, hChild, index)
}

// 元素_显示
// bShow: TRUE显示元素
func (e *Element) Show(bShow bool) int {
	return xc.XEle_Show(e.HEle_, bShow)
}

// 元素_置坐标, 如果返回0坐标没有改变, 如果大小改变返回2(触发XE_SIZE), 否则返回1(仅改变left,top,没有改变大小)
// pRect: 坐标.
// bRedraw: 是否重绘.
// nFlags: 调整布局标识位, XC_AdjustLayout_
// nAdjustNo:
func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags int, nAdjustNo int) int {
	return xc.XEle_SetRect(e.HEle_, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_置坐标扩展, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// bRedraw: 是否重绘.
// nFlags: 调整布局标识位, XC_AdjustLayout_
// nAdjustNo:
func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw bool, nFlags int, nAdjustNo int) int {
	return xc.XEle_SetRectEx(e.HEle_, x, y, cx, cy, bRedraw, nFlags, nAdjustNo)
}

// 元素_置逻辑坐标, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1
// pRect: 坐标.
// bRedraw: 是否重绘.
// nFlags: 参数将被带入XE_SIZE ,XE_ADJUSTLAYOUT 事件回调.xc_adjustLayout_
// nAdjustNo:
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags int, nAdjustNo int) int {
	return xc.XEle_SetRectLogic(e.HEle_, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1
// x: X坐标.
// y: Y坐标.
// bRedraw: 是否重绘.
// nFlags: 调整布局标识位, XC_AdjustLayout_
// nAdjustNo:
func (e *Element) Move(x int, y int, bRedraw bool, nFlags int, nAdjustNo int) int {
	return xc.XEle_Move(e.HEle_, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动逻辑坐标, 移动元素坐标, 逻辑坐标, 包含滚动视图偏移. 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1
// x: X坐标.
// y: Y坐标.
// bRedraw: 是否重绘.
// nFlags: 调整布局标识位, XC_AdjustLayout_
// nAdjustNo:
func (e *Element) MoveLogic(x int, y int, bRedraw bool, nFlags int, nAdjustNo int) int {
	return xc.XEle_MoveLogic(e.HEle_, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_置ID
// nID: ID值.
func (e *Element) SetID(nID int) int {
	return xc.XEle_SetID(e.HEle_, nID)
}

// 元素_取ID
func (e *Element) GetID() int {
	return xc.XEle_GetID(e.HEle_)
}

// 元素_置UID, 置元素UID, 全局唯一标识符
// nUID: UID值.
func (e *Element) SetUID(nUID int) int {
	return xc.XEle_SetUID(e.HEle_, nUID)
}

// 元素_取UID
func (e *Element) GetUID() int {
	return xc.XEle_GetUID(e.HEle_)
}

// 元素_置名称
// pName: name值
func (e *Element) SetName(pName string) int {
	return xc.XEle_SetName(e.HEle_, pName)
}

// 元素_取名称
func (e *Element) GetName() int {
	return xc.XEle_GetName(e.HEle_)
}

// 元素_判断显示
func (e *Element) IsShow() bool {
	return xc.XEle_IsShow(e.HEle_)
}

// 元素_判断绘制焦点
func (e *Element) IsDrawFocus() bool {
	return xc.XEle_IsDrawFocus(e.HEle_)
}

// 元素_判断启用, 元素是否为启用状态
func (e *Element) IsEnable() bool {
	return xc.XEle_IsEnable(e.HEle_)
}

// 元素_判断启用焦点, 元素是否启用焦点
func (e *Element) IsEnableFocus() bool {
	return xc.XEle_IsEnableFocus(e.HEle_)
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透
func (e *Element) IsMouseThrough() bool {
	return xc.XEle_IsMouseThrough(e.HEle_)
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素
// pPt: 坐标点.
func (e *Element) HitChildEle(pPt *xc.POINT) int {
	return xc.XEle_HitChildEle(e.HEle_, pPt)
}

// 元素_判断背景透明
func (e *Element) IsBkTransparent() bool {
	return xc.XEle_IsBkTransparent(e.HEle_)
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件
func (e *Element) IsEnableEvent_XE_PAINT_END() bool {
	return xc.XEle_IsEnableEvent_XE_PAINT_END(e.HEle_)
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit
func (e *Element) IsKeyTab() bool {
	return xc.XEle_IsKeyTab(e.HEle_)
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键)
func (e *Element) IsSwitchFocus() bool {
	return xc.XEle_IsSwitchFocus(e.HEle_)
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素
func (e *Element) IsEnable_XE_MOUSEWHEEL() bool {
	return xc.XEle_IsEnable_XE_MOUSEWHEEL(e.HEle_)
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素
// hChildEle: 子元素句柄
func (e *Element) IsChildEle(hChildEle int) bool {
	return xc.XEle_IsChildEle(e.HEle_, hChildEle)
}

// 元素_判断启用画布, 判断是否启用画布
func (e *Element) IsEnableCanvas() bool {
	return xc.XEle_IsEnableCanvas(e.HEle_)
}

// 元素_判断焦点, 判断是否拥有焦点
func (e *Element) IsFocus() bool {
	return xc.XEle_IsFocus(e.HEle_)
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点
func (e *Element) IsFocusEx() bool {
	return xc.XEle_IsFocusEx(e.HEle_)
}

// 元素_启用, 启用或禁用元素
// bEnable: 启用或禁用.
func (e *Element) Enable(bEnable bool) int {
	return xc.XEle_Enable(e.HEle_, bEnable)
}

// 元素_启用焦点, 启用焦点
// bEnable: 是否启用.
func (e *Element) EnableFocus(bEnable bool) int {
	return xc.XEle_EnableFocus(e.HEle_, bEnable)
}

// 元素_启用绘制焦点
// bEnable: 是否启用.
func (e *Element) EnableDrawFocus(bEnable bool) int {
	return xc.XEle_EnableDrawFocus(e.HEle_, bEnable)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框
// bEnable: 是否启用.
func (e *Element) EnableDrawBorder(bEnable bool) int {
	return xc.XEle_EnableDrawBorder(e.HEle_, bEnable)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布
// bEnable: 是否启用.
func (e *Element) EnableCanvas(bEnable bool) int {
	return xc.XEle_EnableCanvas(e.HEle_, bEnable)
}

// 元素_启用事件_XE_PAINT_END
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_PAINT_END(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_PAINT_END(e.HEle_, bEnable)
}

// 元素_启用背景透明
// bEnable: 是否启用.
func (e *Element) EnableBkTransparent(bEnable bool) int {
	return xc.XEle_EnableBkTransparent(e.HEle_, bEnable)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件
// bEnable: 是否启用.
func (e *Element) EnableMouseThrough(bEnable bool) int {
	return xc.XEle_EnableMouseThrough(e.HEle_, bEnable)
}

// 元素_启用接收TAB, 启用接收Tab输入
// bEnable: 是否启用.
func (e *Element) EnableKeyTab(bEnable bool) int {
	return xc.XEle_EnableKeyTab(e.HEle_, bEnable)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点
// bEnable: 是否启用.
func (e *Element) EnableSwitchFocus(bEnable bool) int {
	return xc.XEle_EnableSwitchFocus(e.HEle_, bEnable)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素
// bEnable: 是否启用.
func (e *Element) EnableEvent_XE_MOUSEWHEEL(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_MOUSEWHEEL(e.HEle_, bEnable)
}

// 元素_取父元素
func (e *Element) GetParentEle() int {
	return xc.XEle_GetParentEle(e.HEle_)
}

// 元素_取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父
func (e *Element) GetParent() int {
	return xc.XEle_GetParent(e.HEle_)
}

// 元素_移除, 移除元素, 但不销毁
func (e *Element) Remove() int {
	return xc.XEle_Remove(e.HEle_)
}

// 元素_置Z序, 设置元素Z序
// index: 位置索引.
func (e *Element) SetZOrder(index int) bool {
	return xc.XEle_SetZOrder(e.HEle_, index)
}

// 元素_置Z序扩展, 设置元素Z序
// hDestEle: 目标元素.
// nType: 类型.
func (e *Element) SetZOrderEx(hDestEle int, nType int) bool {
	return xc.XEle_SetZOrderEx(e.HEle_, hDestEle, nType)
}

// 元素_取Z序, 获取元素Z序索引, 位置索引
func (e *Element) GetZOrder() int {
	return xc.XEle_GetZOrder(e.HEle_)
}

// 元素_启用置顶, 设置元素置顶.
// bTopmost: 是否置顶显示
func (e *Element) EnableTopmost(bTopmost bool) bool {
	return xc.XEle_EnableTopmost(e.HEle_, bTopmost)
}

// 元素_重绘
// bImmediate: 是否立即重绘
func (e *Element) Redraw(bImmediate bool) int {
	return xc.XEle_Redraw(e.HEle_, bImmediate)
}

// 元素_重绘指定区域
// pRect: 相对于元素客户区坐标.
// bImmediate: 是否立即重绘
func (e *Element) RedrawRect(pRect *xc.RECT, bImmediate bool) int {
	return xc.XEle_RedrawRect(e.HEle_, pRect, bImmediate)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象
func (e *Element) GetChildCount() int {
	return xc.XEle_GetChildCount(e.HEle_)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象
// index: 索引.
func (e *Element) GetChildByIndex(index int) int {
	return xc.XEle_GetChildByIndex(e.HEle_, index)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象
// nID: 元素ID.
func (e *Element) GetChildByID(nID int) int {
	return xc.XEle_GetChildByID(e.HEle_, nID)
}

// 元素_置边框大小
// left: 左边大小.
// top: 上边大小.
// right: 右边大小.
// bottom: 下边大小.
func (e *Element) SetBorderSize(left int, top int, right int, bottom int) int {
	return xc.XEle_SetBorderSize(e.HEle_, left, top, right, bottom)
}

// 元素_取边框大小
// pBorder: 大小.
func (e *Element) GetBorderSize(pBorder *xc.RECT) int {
	return xc.XEle_GetBorderSize(e.HEle_, pBorder)
}

// 元素_置内填充大小
// left: 左边大小.
// top: 上边大小.
// right: 右边大小.
// bottom: 下边大小.
func (e *Element) SetPadding(left int, top int, right int, bottom int) int {
	return xc.XEle_SetPadding(e.HEle_, left, top, right, bottom)
}

// 元素_取内填充大小
// pPadding: 大小.
func (e *Element) GetPadding(pPadding *xc.RECT) int {
	return xc.XEle_GetPadding(e.HEle_, pPadding)
}

// 元素_置拖动边框
// nFlags: 边框位置组合, Element_Position_
func (e *Element) SetDragBorder(nFlags int) int {
	return xc.XEle_SetDragBorder(e.HEle_, nFlags)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小
// nFlags: 边框位置标识, Element_Position_
// hBindEle: 绑定元素.
// nSpace: 元素间隔大小
func (e *Element) SetDragBorderBindEle(nFlags int, hBindEle int, nSpace int) int {
	return xc.XEle_SetDragBorderBindEle(e.HEle_, nFlags, hBindEle, nSpace)
}

// 元素_置最小大小
// nWidth: 最小宽度
// nHeight: 最小高度.
func (e *Element) SetMinSize(nWidth int, nHeight int) int {
	return xc.XEle_SetMinSize(e.HEle_, nWidth, nHeight)
}

// 元素_置最大大小
// nWidth: 最大宽度.
// nHeight: 最大高度.
func (e *Element) SetMaxSize(nWidth int, nHeight int) int {
	return xc.XEle_SetMaxSize(e.HEle_, nWidth, nHeight)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动
// bHorizon: 是否锁定水平滚动.
// bVertical: 是否锁定垂直滚动.
func (e *Element) SetLockScroll(bHorizon bool, bVertical bool) int {
	return xc.XEle_SetLockScroll(e.HEle_, bHorizon, bVertical)
}

// 元素_置文本颜色
// color: RGB颜色值.
// alpha: 透明度.
func (e *Element) SetTextColor(color int, alpha uint8) int {
	return xc.XEle_SetTextColor(e.HEle_, color, alpha)
}

// 元素_取文本颜色
func (e *Element) GetTextColor() int {
	return xc.XEle_GetTextColor(e.HEle_)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取
func (e *Element) GetTextColorEx() int {
	return xc.XEle_GetTextColorEx(e.HEle_)
}

// 元素_置焦点边框颜色
// color: RGB颜色值.
// alpha: 透明度.
func (e *Element) SetFocusBorderColor(color int, alpha uint8) int {
	return xc.XEle_SetFocusBorderColor(e.HEle_, color, alpha)
}

// 元素_取焦点边框颜色
func (e *Element) GetFocusBorderColor() int {
	return xc.XEle_GetFocusBorderColor(e.HEle_)
}

// 元素_置字体
// hFontx: 炫彩字体.
func (e *Element) SetFont(hFontx int) int {
	return xc.XEle_SetFont(e.HEle_, hFontx)
}

// 元素_取字体
func (e *Element) GetFont() int {
	return xc.XEle_GetFont(e.HEle_)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取
func (e *Element) GetFontEx() int {
	return xc.XEle_GetFontEx(e.HEle_)
}

// 元素_置透明度
// alpha: 透明度.
func (e *Element) SetAlpha(alpha uint8) int {
	return xc.XEle_SetAlpha(e.HEle_, alpha)
}

// 元素_销毁
func (e *Element) Destroy() int {
	return xc.XEle_Destroy(e.HEle_)
}

// 元素_添加背景边框, 添加背景内容边框
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (e *Element) AddBkBorder(color int, alpha uint8, width int) int {
	return xc.XEle_AddBkBorder(e.HEle_, color, alpha, width)
}

// 元素_添加背景填充, 添加背景内容填充
// color: RGB颜色.
// alpha: 透明度.
func (e *Element) AddBkFill(color int, alpha uint8) int {
	return xc.XEle_AddBkFill(e.HEle_, color, alpha)
}

// 元素_添加背景图片, 添加背景内容图片
// hImage: 图片句柄.
func (e *Element) AddBkImage(hImage int) int {
	return xc.XEle_AddBkImage(e.HEle_, hImage)
}

// 元素_取背景对象数量, 获取背景内容数量
func (e *Element) GetBkInfoCount() int {
	return xc.XEle_GetBkInfoCount(e.HEle_)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确
func (e *Element) ClearBkInfo() int {
	return xc.XEle_ClearBkInfo(e.HEle_)
}

// 元素_取背景管理器, 获取元素背景管理器
func (e *Element) GetBkManager() int {
	return xc.XEle_GetBkManager(e.HEle_)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取
func (e *Element) GetBkManagerEx() int {
	return xc.XEle_GetBkManagerEx(e.HEle_)
}

// 元素_置背景管理器
// hBkInfoM: 背景管理器
func (e *Element) SetBkManager(hBkInfoM int) int {
	return xc.XEle_SetBkManager(e.HEle_, hBkInfoM)
}

// 元素_取状态, 获取组合状态
func (e *Element) GetStateFlags() int {
	return xc.XEle_GetStateFlags(e.HEle_)
}

// 元素_绘制焦点, 绘制元素焦点
// hDraw: 图形绘制句柄.
// pRect: 区域坐标.
func (e *Element) DrawFocus(hDraw int, pRect *xc.RECT) bool {
	return xc.XEle_DrawFocus(e.HEle_, hDraw, pRect)
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序
// hDraw: 图形绘制句柄.
func (e *Element) DrawEle(hDraw int) int {
	return xc.XEle_DrawEle(e.HEle_, hDraw)
}

// 元素_置用户数据
// nData: 用户数据.
func (e *Element) SetUserData(nData int) int {
	return xc.XEle_SetUserData(e.HEle_, nData)
}

// 元素_取用户数据
func (e *Element) GetUserData() int {
	return xc.XEle_GetUserData(e.HEle_)
}

// 元素_取内容大小
// bHorizon:
// pSize: 返回大小.
// cx: 宽度
// cy: 高度
func (e *Element) GetContentSize(bHorizon bool, pSize *xc.SIZE, cx int, cy int) int {
	return xc.XEle_GetContentSize(e.HEle_, bHorizon, pSize, cx, cy)
}

// 元素_置鼠标捕获
// b: TRUE设置
func (e *Element) SetCapture(b bool) int {
	return xc.XEle_SetCapture(e.HEle_, b)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题
// bEnable: 启用或关闭.
func (e *Element) EnableTransparentChannel(bEnable bool) int {
	return xc.XEle_EnableTransparentChannel(e.HEle_, bEnable)
}

// 元素_置炫彩定时器, 设置元素定时器
// nIDEvent: 事件ID.
// uElapse: 延时毫秒.
func (e *Element) SetXCTimer(nIDEvent int, uElapse int) bool {
	return xc.XEle_SetXCTimer(e.HEle_, nIDEvent, uElapse)
}

// 元素_关闭炫彩定时器, 关闭元素定时器
// nIDEvent: 事件ID.
func (e *Element) KillXCTimer(nIDEvent int) bool {
	return xc.XEle_KillXCTimer(e.HEle_, nIDEvent)
}

// 元素_置工具提示, 设置工具提示内容
// pText: 工具提示内容.
func (e *Element) SetToolTip(pText string) int {
	return xc.XEle_SetToolTip(e.HEle_, pText)
}

// 元素_置工具提示扩展, 设置工具提示内容
// pText: 工具提示内容.
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_
func (e *Element) SetToolTipEx(pText string, nTextAlign int) int {
	return xc.XEle_SetToolTipEx(e.HEle_, pText, nTextAlign)
}

// 元素_取工具提示, 获取工具提示内容
func (e *Element) GetToolTip() int {
	return xc.XEle_GetToolTip(e.HEle_)
}

// 元素_弹出工具提示, 弹出工具提示
// x: X坐标.
// y: Y坐标.
func (e *Element) PopupToolTip(x int, y int) int {
	return xc.XEle_PopupToolTip(e.HEle_, x, y)
}

// 元素_调整布局
// nAdjustNo:
func (e *Element) AdjustLayout(nAdjustNo int) int {
	return xc.XEle_AdjustLayout(e.HEle_, nAdjustNo)
}

// 元素_调整布局扩展
// nFlags: 调整标识
// nAdjustNo:
func (e *Element) AdjustLayoutEx(nFlags int, nAdjustNo int) int {
	return xc.XEle_AdjustLayoutEx(e.HEle_, nFlags, nAdjustNo)
}
