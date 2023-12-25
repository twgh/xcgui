package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 日期时间.
type DateTime struct {
	Element
}

// 日期_创建, 创建日期时间元素.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewDateTime(x int, y int, cx int, cy int, hParent int) *DateTime {
	p := &DateTime{}
	p.SetHandle(xc.XDateTime_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewDateTimeByHandle(handle int) *DateTime {
	p := &DateTime{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewDateTimeByName(name string) *DateTime {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewDateTimeByUID(nUID int) *DateTime {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewDateTimeByUIDName(name string) *DateTime {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 日期_置样式, 设置样式.
//
// nStyle: 样式: 0为日期元素, 1为时间元素.
func (d *DateTime) SetStyle(nStyle int) int {
	return xc.XDateTime_SetStyle(d.Handle, nStyle)
}

// 日期_取样式, 返回元素样式.
func (d *DateTime) GetStyle() int {
	return xc.XDateTime_GetStyle(d.Handle)
}

// 日期_启用分割栏为斜线, 切换分割栏为: 斜线或横线.
//
// bSlash: TRUE: 斜线, FALSE: 横线.
func (d *DateTime) EnableSplitSlash(bSlash bool) int {
	return xc.XDateTime_EnableSplitSlash(d.Handle, bSlash)
}

// 日期_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型.
func (d *DateTime) GetButton(nType int) int {
	return xc.XDateTime_GetButton(d.Handle, nType)
}

// 日期_取选择日期背景颜色, 获取被选择文字的背景颜色.
func (d *DateTime) GetSelBkColor() int {
	return xc.XDateTime_GetSelBkColor(d.Handle)
}

// 日期_置选择日期背景颜色, 设置被选择文字的背景颜色.
//
// crSelectBk: 文字被选中背景色, ABGR 颜色.
func (d *DateTime) SetSelBkColor(crSelectBk int) int {
	return xc.XDateTime_SetSelBkColor(d.Handle, crSelectBk)
}

// 日期_取当前日期.
//
// pnYear: 年.[OUT].
//
// pnMonth: 月.[OUT].
//
// pnDay: 日.[OUT].
func (d *DateTime) GetDate(pnYear *int32, pnMonth *int32, pnDay *int32) int {
	return xc.XDateTime_GetDate(d.Handle, pnYear, pnMonth, pnDay)
}

// 日期_置当前日期.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func (d *DateTime) SetDate(nYear int32, nMonth int32, nDay int32) int {
	return xc.XDateTime_SetDate(d.Handle, nYear, nMonth, nDay)
}

// 日期_取当前时间.
//
// pnHour: 时.[OUT].
//
// pnMinute: 分.[OUT].
//
// pnSecond: 秒.[OUT].
func (d *DateTime) GetTime(pnHour *int32, pnMinute *int32, pnSecond *int32) int {
	return xc.XDateTime_GetTime(d.Handle, pnHour, pnMinute, pnSecond)
}

// 日期_社区当前时间, 设置当前时分秒.
//
// nHour: 时.
//
// nMinute: 分.
//
// nSecond: 秒.
func (d *DateTime) SetTime(nHour int32, nMinute int32, nSecond int32) int {
	return xc.XDateTime_SetTime(d.Handle, nHour, nMinute, nSecond)
}

// 日期_弹出, 弹出月历卡片.
func (d *DateTime) Popup() int {
	return xc.XDateTime_Popup(d.Handle)
}

/*
以下都是事件
*/

type XE_DATETIME_CHANGE func(pbHandled *bool) int                                                     // 日期时间元素,内容改变事件.
type XE_DATETIME_CHANGE1 func(hEle int, pbHandled *bool) int                                          // 日期时间元素,内容改变事件.
type XE_DATETIME_POPUP_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int            // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_POPUP_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_EXIT_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int             // 日期时间元素,弹出的月历卡片退出事件.
type XE_DATETIME_EXIT_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int  // 日期时间元素,弹出的月历卡片退出事件.

// 日期时间元素,内容改变事件.
func (d *DateTime) Event_DATETIME_CHANGE(pFun XE_DATETIME_CHANGE) bool {
	return xc.XEle_RegEventC(d.Handle, xcc.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,内容改变事件.
func (d *DateTime) Event_DATETIME_CHANGE1(pFun XE_DATETIME_CHANGE1) bool {
	return xc.XEle_RegEventC1(d.Handle, xcc.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,弹出月历卡片事件.
func (d *DateTime) Event_DATETIME_POPUP_MONTHCAL(pFun XE_DATETIME_POPUP_MONTHCAL) bool {
	return xc.XEle_RegEventC(d.Handle, xcc.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,弹出月历卡片事件.
func (d *DateTime) Event_DATETIME_POPUP_MONTHCAL1(pFun XE_DATETIME_POPUP_MONTHCAL1) bool {
	return xc.XEle_RegEventC1(d.Handle, xcc.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,弹出的月历卡片退出事件.
func (d *DateTime) Event_DATETIME_EXIT_MONTHCAL(pFun XE_DATETIME_EXIT_MONTHCAL) bool {
	return xc.XEle_RegEventC(d.Handle, xcc.XE_DATETIME_EXIT_MONTHCAL, pFun)
}

// 日期时间元素,弹出的月历卡片退出事件.
func (d *DateTime) Event_DATETIME_EXIT_MONTHCAL1(pFun XE_DATETIME_EXIT_MONTHCAL1) bool {
	return xc.XEle_RegEventC1(d.Handle, xcc.XE_DATETIME_EXIT_MONTHCAL, pFun)
}
