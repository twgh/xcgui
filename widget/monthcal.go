package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// MonthCal 月历卡片.
type MonthCal struct {
	Element
}

// 月历_创建, 创建日期时间元素.
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
func NewMonthCal(x, y, cx, cy int32, hParent int) *MonthCal {
	p := &MonthCal{}
	p.SetHandle(xc.XMonthCal_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewMonthCalByHandle(handle int) *MonthCal {
	p := &MonthCal{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewMonthCalByName(name string) *MonthCal {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewMonthCalByUID(nUID int32) *MonthCal {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewMonthCalByUIDName(name string) *MonthCal {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 月历_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型, MonthCal_Button_Type_.
func (m *MonthCal) GetButton(nType xcc.MonthCal_Button_Type_) int {
	return xc.XMonthCal_GetButton(m.Handle, nType)
}

// 月历_置当前日期, 设置月历选中的年月日.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func (m *MonthCal) SetToday(nYear int32, nMonth int32, nDay int32) *MonthCal {
	xc.XMonthCal_SetToday(m.Handle, nYear, nMonth, nDay)
	return m
}

// 月历_取当前日期, 获取月历当前年月日.
//
// pnYear: 年.
//
// pnMonth: 月.
//
// pnDay: 日.
func (m *MonthCal) GetToday(pnYear *int32, pnMonth *int32, pnDay *int32) *MonthCal {
	xc.XMonthCal_GetToday(m.Handle, pnYear, pnMonth, pnDay)
	return m
}

// 月历_取选择日期, 获取月历选中的年月日.
//
// pnYear: 年.
//
// pnMonth: 月.
//
// pnDay: 日.
func (m *MonthCal) GetSelDate(pnYear *int32, pnMonth *int32, pnDay *int32) *MonthCal {
	xc.XMonthCal_GetSelDate(m.Handle, pnYear, pnMonth, pnDay)
	return m
}

// 月历_置文本颜色.
//
// nFlag: 1:周六, 周日文字颜色, 2:日期文字的颜色; 其它周文字颜色, 使用元素自身颜色.
//
// color: xc.RGBA 颜色值.
func (m *MonthCal) SetTextColor(nFlag int32, color int) *MonthCal {
	xc.XMonthCal_SetTextColor(m.Handle, nFlag, color)
	return m
}

/*
以下都是事件
*/

type XE_MONTHCAL_CHANGE func(pbHandled *bool) int            // 月历元素日期改变事件.
type XE_MONTHCAL_CHANGE1 func(hEle int, pbHandled *bool) int // 月历元素日期改变事件.

// 月历元素日期改变事件.
func (m *MonthCal) Event_MONTHCAL_CHANGE(pFun XE_MONTHCAL_CHANGE) bool {
	return xc.XEle_RegEventC(m.Handle, xcc.XE_MONTHCAL_CHANGE, pFun)
}

// 月历元素日期改变事件.
func (m *MonthCal) Event_MONTHCAL_CHANGE1(pFun XE_MONTHCAL_CHANGE1) bool {
	return xc.XEle_RegEventC1(m.Handle, xcc.XE_MONTHCAL_CHANGE, pFun)
}
