package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ScrollView 滚动视图.
type ScrollView struct {
	Element
}

// 滚动视_创建, 创建滚动视图元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func NewScrollView(x, y, cx, cy int32, hParent int) *ScrollView {
	p := &ScrollView{}
	p.SetHandle(xc.XSView_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewScrollViewByHandle(handle int) *ScrollView {
	p := &ScrollView{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewScrollViewByName(name string) *ScrollView {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewScrollViewByUID(nUID int32) *ScrollView {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewScrollViewByUIDName(name string) *ScrollView {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 滚动视_置视图大小, 设置内容大小, 如果内容改变返回TRUE否则返回FALSE.
//
// cx: 宽度.
//
// cy: 高度.
func (s *ScrollView) SetTotalSize(cx, cy int32) bool {
	return xc.XSView_SetTotalSize(s.Handle, cx, cy)
}

// 滚动视_取视图大小, 获取内容总大小.
//
// pSize: 大小.
func (s *ScrollView) GetTotalSize(pSize *xc.SIZE) *ScrollView {
	xc.XSView_GetTotalSize(s.Handle, pSize)
	return s
}

// 滚动视_置滚动单位大小, 设置滚动单位大小, 如果内容改变返回TRUE否则返回FALSE.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *ScrollView) SetLineSize(nWidth, nHeight int32) bool {
	return xc.XSView_SetLineSize(s.Handle, nWidth, nHeight)
}

// 滚动视_取滚动单位大小, 获取滚动单位大小.
//
// pSize: 返回大小.
func (s *ScrollView) GetLineSize(pSize *xc.SIZE) *ScrollView {
	xc.XSView_GetLineSize(s.Handle, pSize)
	return s
}

// 滚动视_置滚动条大小.
//
// size: 滚动条大小.
func (s *ScrollView) SetScrollBarSize(size int32) *ScrollView {
	xc.XSView_SetScrollBarSize(s.Handle, size)
	return s
}

// 滚动视_取视口原点X, 获取视口原点X坐标.
func (s *ScrollView) GetViewPosH() int32 {
	return xc.XSView_GetViewPosH(s.Handle)
}

// 滚动视_取视口原点Y, 获取视口原点Y坐标.
func (s *ScrollView) GetViewPosV() int32 {
	return xc.XSView_GetViewPosV(s.Handle)
}

// 滚动视_取视口宽度.
func (s *ScrollView) GetViewWidth() int32 {
	return xc.XSView_GetViewWidth(s.Handle)
}

// 滚动视_取视口高度.
func (s *ScrollView) GetViewHeight() int32 {
	return xc.XSView_GetViewHeight(s.Handle)
}

// 滚动视_取视口坐标.
//
// pRect: 坐标.
func (s *ScrollView) GetViewRect(pRect *xc.RECT) *ScrollView {
	xc.XSView_GetViewRect(s.Handle, pRect)
	return s
}

// 滚动视_取水平滚动条, 返回滚动条句柄.
func (s *ScrollView) GetScrollBarH() int {
	return xc.XSView_GetScrollBarH(s.Handle)
}

// 滚动视_取垂直滚动条, 返回滚动条句柄.
func (s *ScrollView) GetScrollBarV() int {
	return xc.XSView_GetScrollBarV(s.Handle)
}

// 滚动视_水平滚动, 水平滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) ScrollPosH(pos int32) bool {
	return xc.XSView_ScrollPosH(s.Handle, pos)
}

// 滚动视_垂直滚动, 垂直滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) ScrollPosV(pos int32) bool {
	return xc.XSView_ScrollPosV(s.Handle, pos)
}

// 滚动视_水平滚动到X, 水平滚动条, 滚动到指定坐标.
//
// posX: X坐标.
func (s *ScrollView) ScrollPosXH(posX int32) bool {
	return xc.XSView_ScrollPosXH(s.Handle, posX)
}

// 滚动视_垂直滚动到Y, 垂直滚动条, 滚动到指定坐标.
//
// posY: Y坐标.
func (s *ScrollView) ScrollPosYV(posY int32) bool {
	return xc.XSView_ScrollPosYV(s.Handle, posY)
}

// 滚动视_显示水平滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) ShowSBarH(bShow bool) *ScrollView {
	xc.XSView_ShowSBarH(s.Handle, bShow)
	return s
}

// 滚动视_显示垂直滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) ShowSBarV(bShow bool) *ScrollView {
	xc.XSView_ShowSBarV(s.Handle, bShow)
	return s
}

// 滚动视_启用自动显示滚动条.
//
// bEnable: 是否启用.
func (s *ScrollView) EnableAutoShowScrollBar(bEnable bool) *ScrollView {
	xc.XSView_EnableAutoShowScrollBar(s.Handle, bEnable)
	return s
}

// 滚动视_向左滚动.
func (s *ScrollView) ScrollLeftLine() bool {
	return xc.XSView_ScrollLeftLine(s.Handle)
}

// 滚动视_向右滚动.
func (s *ScrollView) ScrollRightLine() bool {
	return xc.XSView_ScrollRightLine(s.Handle)
}

// 滚动视_向上滚动.
func (s *ScrollView) ScrollTopLine() bool {
	return xc.XSView_ScrollTopLine(s.Handle)
}

// 滚动视_向下滚动.
func (s *ScrollView) ScrollBottomLine() bool {
	return xc.XSView_ScrollBottomLine(s.Handle)
}

// 滚动视_滚动到左侧, 水平滚动到左侧.
func (s *ScrollView) ScrollLeft() bool {
	return xc.XSView_ScrollLeft(s.Handle)
}

// 滚动视_滚动到右侧, 水平滚动到右侧.
func (s *ScrollView) ScrollRight() bool {
	return xc.XSView_ScrollRight(s.Handle)
}

// 滚动视_滚动到顶部, 垂直滚动到顶部.
func (s *ScrollView) ScrollTop() bool {
	return xc.XSView_ScrollTop(s.Handle)
}

// 滚动视_滚动到底部, 垂直滚动到底部.
func (s *ScrollView) ScrollBottom() bool {
	return xc.XSView_ScrollBottom(s.Handle)
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_ScrollView_Scroll_H 添加滚动视图元素水平滚动事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (s *ScrollView) AddEvent_ScrollView_Scroll_H(pFun XE_SCROLLVIEW_SCROLL_H1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_H, onXE_SCROLLVIEW_SCROLL_H, pFun, allowAddingMultiple...)
}

// onXE_SCROLLVIEW_SCROLL_H 滚动视图元素水平滚动事件.
func onXE_SCROLLVIEW_SCROLL_H(hEle int, pos int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SCROLLVIEW_SCROLL_H)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XE_SCROLLVIEW_SCROLL_H1)(hEle, pos, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_ScrollView_Scroll_V 添加滚动视图元素垂直滚动事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (s *ScrollView) AddEvent_ScrollView_Scroll_V(pFun XE_SCROLLVIEW_SCROLL_V1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_V, onXE_SCROLLVIEW_SCROLL_V, pFun, allowAddingMultiple...)
}

// onXE_SCROLLVIEW_SCROLL_V 滚动视图元素垂直滚动事件.
func onXE_SCROLLVIEW_SCROLL_V(hEle int, pos int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SCROLLVIEW_SCROLL_V)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XE_SCROLLVIEW_SCROLL_V1)(hEle, pos, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

type XE_SCROLLVIEW_SCROLL_H func(pos int32, pbHandled *bool) int            // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_H1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V func(pos int32, pbHandled *bool) int            // 滚动视图元素垂直滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素垂直滚动事件,滚动视图触发.

// 滚动视图元素水平滚动事件,滚动视图触发.
func (s *ScrollView) Event_SCROLLVIEW_SCROLL_H(pFun XE_SCROLLVIEW_SCROLL_H) bool {
	return xc.XEle_RegEventC(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素水平滚动事件,滚动视图触发.
func (s *ScrollView) Event_SCROLLVIEW_SCROLL_H1(pFun XE_SCROLLVIEW_SCROLL_H1) bool {
	return xc.XEle_RegEventC1(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素垂直滚动事件,滚动视图触发.
func (s *ScrollView) Event_SCROLLVIEW_SCROLL_V(pFun XE_SCROLLVIEW_SCROLL_V) bool {
	return xc.XEle_RegEventC(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_V, pFun)
}

// 滚动视图元素垂直滚动事件,滚动视图触发.
func (s *ScrollView) Event_SCROLLVIEW_SCROLL_V1(pFun XE_SCROLLVIEW_SCROLL_V1) bool {
	return xc.XEle_RegEventC1(s.Handle, xcc.XE_SCROLLVIEW_SCROLL_V, pFun)
}
