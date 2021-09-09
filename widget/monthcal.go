package widget

import "github.com/twgh/xcgui/xc"

// 月历卡片
type MonthCal struct {
	Element
}

// 月历_创建, 创建日期时间元素
// x: x坐标
// y: y坐标
// cx: 宽度
// cy: 高度
// hParent: 父为窗口句柄或元素句柄.
func NewMonthCal(x int, y int, cx int, cy int, hParent int) *MonthCal {
	p := &MonthCal{}
	p.SetHandle(xc.XMonthCal_Create(x, y, cx, cy, hParent))
	return p
}

// 月历_取内部按钮, 获取内部按钮元素
// nType: 按钮类型.
func (m *MonthCal) GetButton(nType int) int {
	return xc.XMonthCal_GetButton(m.Handle, nType)
}

// 月历_置当前日期, 设置月历选中的年月日
// nYear: 年.
// nMonth: 月.
// nDay: 日.
func (m *MonthCal) SetToday(nYear int, nMonth int, nDay int) int {
	return xc.XMonthCal_SetToday(m.Handle, nYear, nMonth, nDay)
}

// 月历_取当前日期, 获取月历当前年月日
// pnYear: 年.[INT
// pnMonth: 月.[INT
// pnDay: 日.[INT
func (m *MonthCal) GetToday(pnYear *int, pnMonth *int, pnDay *int) int {
	return xc.XMonthCal_GetToday(m.Handle, pnYear, pnMonth, pnDay)
}

// 月历_取选择日期, 获取月历选中的年月日
// pnYear: 年.[INT
// pnMonth: 月.[INT
// pnDay: 日.[INT
func (m *MonthCal) GetSelDate(pnYear *int, pnMonth *int, pnDay *int) int {
	return xc.XMonthCal_GetSelDate(m.Handle, pnYear, pnMonth, pnDay)
}
