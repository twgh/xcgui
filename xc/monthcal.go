package xc

import "unsafe"

// 月历_创建
// x: x坐标
// y: y坐标
// cx: 宽度
// cy: 高度
// hParent: 父为窗口句柄或元素句柄.
func XMonthCal_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xMonthCal_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 月历_取内部按钮
// hEle: 元素句柄.
// nType: 按钮类型.
func XMonthCal_GetButton(hEle int, nType int) int {
	r, _, _ := xMonthCal_GetButton.Call(uintptr(hEle), uintptr(nType))
	return int(r)
}

// 月历_置当前日期
// hEle: 元素句柄.
// nYear: 年.
// nMonth: 月.
// nDay: 日.
func XMonthCal_SetToday(hEle int, nYear int, nMonth int, nDay int) int {
	r, _, _ := xMonthCal_SetToday.Call(uintptr(hEle), uintptr(nYear), uintptr(nMonth), uintptr(nDay))
	return int(r)
}

// 月历_取当前日期
// hEle: 元素句柄.
// pnYear: 年.[INT
// pnMonth: 月.[INT
// pnDay: 日.[INT
func XMonthCal_GetToday(hEle int, pnYear *int, pnMonth *int, pnDay *int) int {
	r, _, _ := xMonthCal_GetToday.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnYear)), uintptr(unsafe.Pointer(pnMonth)), uintptr(unsafe.Pointer(pnDay)))
	return int(r)
}

// 月历_取选择日期
// hEle: 元素句柄.
// pnYear: 年.[INT
// pnMonth: 月.[INT
// pnDay: 日.[INT
func XMonthCal_GetSelDate(hEle int, pnYear *int, pnMonth *int, pnDay *int) int {
	r, _, _ := xMonthCal_GetSelDate.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnYear)), uintptr(unsafe.Pointer(pnMonth)), uintptr(unsafe.Pointer(pnDay)))
	return int(r)
}
