package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ScrollBar 滚动条.
type ScrollBar struct {
	Element
}

// 滚动条_创建, 创建滚动条元素, 返回元素句柄.
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
func NewScrollBar(x, y, cx, cy int32, hParent int) *ScrollBar {
	p := &ScrollBar{}
	p.SetHandle(xc.XSBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewScrollBarByHandle(handle int) *ScrollBar {
	p := &ScrollBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewScrollBarByName(name string) *ScrollBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewScrollBarByUID(nUID int32) *ScrollBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewScrollBarByUIDName(name string) *ScrollBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 滚动条_置范围, 设置滚动范围.
//
// range_: 范围.
func (s *ScrollBar) SetRange(range_ int32) *ScrollBar {
	xc.XSBar_SetRange(s.Handle, range_)
	return s
}

// 滚动条_取范围, 获取滚动范围.
func (s *ScrollBar) GetRange() int32 {
	return xc.XSBar_GetRange(s.Handle)
}

// 滚动条_显示上下按钮, 显示隐藏滚动条上下按钮.
//
// bShow: 是否显示.
func (s *ScrollBar) ShowButton(bShow bool) *ScrollBar {
	xc.XSBar_ShowButton(s.Handle, bShow)
	return s
}

// 滚动条_置滑块长度.
//
// length: 长度.
func (s *ScrollBar) SetSliderLength(length int32) *ScrollBar {
	xc.XSBar_SetSliderLength(s.Handle, length)
	return s
}

// 滚动条_置滑块最小长度.
//
// minLength: 长度.
func (s *ScrollBar) SetSliderMinLength(minLength int32) *ScrollBar {
	xc.XSBar_SetSliderMinLength(s.Handle, minLength)
	return s
}

// 滚动条_置滑块两边间隔, 设置滑块两边的间隔大小.
//
// nPadding: 间隔大小.
func (s *ScrollBar) SetSliderPadding(nPadding int32) *ScrollBar {
	xc.XSBar_SetSliderPadding(s.Handle, nPadding)
	return s
}

// 滚动条_置水平, 设置水平或者垂直.
//
// bHorizon: 水平或垂直.
func (s *ScrollBar) EnableHorizon(bHorizon bool) bool {
	return xc.XSBar_EnableHorizon(s.Handle, bHorizon)
}

// 滚动条_取滑块最大长度.
func (s *ScrollBar) GetSliderMaxLength() int32 {
	return xc.XSBar_GetSliderMaxLength(s.Handle)
}

// 滚动条_向上滚动.
func (s *ScrollBar) ScrollUp() bool {
	return xc.XSBar_ScrollUp(s.Handle)
}

// 滚动条_向下滚动.
func (s *ScrollBar) ScrollDown() bool {
	return xc.XSBar_ScrollDown(s.Handle)
}

// 滚动条_滚动到顶部.
func (s *ScrollBar) ScrollTop() bool {
	return xc.XSBar_ScrollTop(s.Handle)
}

// 滚动条_滚动到底部.
func (s *ScrollBar) ScrollBottom() bool {
	return xc.XSBar_ScrollBottom(s.Handle)
}

// 滚动条_滚动到指定位置, 滚动到指定位置点, 触发事件: XE_SBAR_SCROLL.
//
// pos: 位置点.
func (s *ScrollBar) ScrollPos(pos int32) bool {
	return xc.XSBar_ScrollPos(s.Handle, pos)
}

// 滚动条_取上按钮, 获取上按钮, 返回按钮句柄.
func (s *ScrollBar) GetButtonUp() int {
	return xc.XSBar_GetButtonUp(s.Handle)
}

// 滚动条_取下按钮, 获取下按钮, 返回按钮句柄.
func (s *ScrollBar) GetButtonDown() int {
	return xc.XSBar_GetButtonDown(s.Handle)
}

// 滚动条_取滑块, 获取滑动按钮, 返回按钮句柄.
func (s *ScrollBar) GetButtonSlider() int {
	return xc.XSBar_GetButtonSlider(s.Handle)
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_SBar_Scroll 添加滚动条元素滚动事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (s *ScrollBar) AddEvent_SBar_Scroll(pFun xc.XE_SBAR_SCROLL1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(s.Handle, xcc.XE_SBAR_SCROLL, onXE_SBAR_SCROLL, pFun, allowAddingMultiple...)
}

// onXE_SBAR_SCROLL 滚动条元素滚动事件处理.
func onXE_SBAR_SCROLL(hEle int, pos int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_SBAR_SCROLL)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_SBAR_SCROLL1); ok {
			ret = cb(hEle, pos, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 滚动条元素滚动事件,滚动条触发.
func (s *ScrollBar) Event_SBAR_SCROLL(pFun xc.XE_SBAR_SCROLL) bool {
	return xc.XEle_RegEventC(s.Handle, xcc.XE_SBAR_SCROLL, pFun)
}

// 滚动条元素滚动事件,滚动条触发.
func (s *ScrollBar) Event_SBAR_SCROLL1(pFun xc.XE_SBAR_SCROLL1) bool {
	return xc.XEle_RegEventC1(s.Handle, xcc.XE_SBAR_SCROLL, pFun)
}
