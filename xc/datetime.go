package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

// 日期_创建, 创建日期时间元素, 返回元素句柄.
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
func XDateTime_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xDateTime_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 日期_置样式, 设置样式.
//
// hEle: 元素句柄.
//
// nStyle: 样式: 0为日期元素, 1为时间元素.
func XDateTime_SetStyle(hEle int, nStyle int) int {
	r, _, _ := xDateTime_SetStyle.Call(uintptr(hEle), uintptr(nStyle))
	return int(r)
}

// 日期_取样式, 返回元素样式.
//
// hEle: 元素句柄.
func XDateTime_GetStyle(hEle int) int {
	r, _, _ := xDateTime_GetStyle.Call(uintptr(hEle))
	return int(r)
}

// 日期_启用分割栏为斜线, 切换分割栏为: 斜线或横线.
//
// hEle: 元素句柄.
//
// bSlash: TRUE: 斜线, FALSE: 横线.
func XDateTime_EnableSplitSlash(hEle int, bSlash bool) int {
	r, _, _ := xDateTime_EnableSplitSlash.Call(uintptr(hEle), common.BoolPtr(bSlash))
	return int(r)
}

// 日期_取内部按钮, 获取内部按钮元素.
//
// hEle: 元素句柄.
//
// nType: 按钮类型.
func XDateTime_GetButton(hEle int, nType int) int {
	r, _, _ := xDateTime_GetButton.Call(uintptr(hEle), uintptr(nType))
	return int(r)
}

// 日期_取选择日期背景颜色, 获取被选择文字的背景颜色.
//
// hEle: 元素句柄.
func XDateTime_GetSelBkColor(hEle int) int {
	r, _, _ := xDateTime_GetSelBkColor.Call(uintptr(hEle))
	return int(r)
}

// 日期_置选择日期背景颜色, 设置被选择文字的背景颜色.
//
// hEle: 元素句柄.
//
// crSelectBk: 文字被选中背景色, ABGR 颜色.
func XDateTime_SetSelBkColor(hEle int, crSelectBk int) int {
	r, _, _ := xDateTime_SetSelBkColor.Call(uintptr(hEle), uintptr(crSelectBk))
	return int(r)
}

// 日期_取当前日期.
//
// hEle: 元素句柄.
//
// pnYear: 年.[OUT].
//
// pnMonth: 月.[OUT].
//
// pnDay: 日.[OUT].
func XDateTime_GetDate(hEle int, pnYear *int32, pnMonth *int32, pnDay *int32) int {
	r, _, _ := xDateTime_GetDate.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnYear)), uintptr(unsafe.Pointer(pnMonth)), uintptr(unsafe.Pointer(pnDay)))
	return int(r)
}

// 日期_置当前日期.
//
// hEle: 元素句柄.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func XDateTime_SetDate(hEle int, nYear int32, nMonth int32, nDay int32) int {
	r, _, _ := xDateTime_SetDate.Call(uintptr(hEle), uintptr(nYear), uintptr(nMonth), uintptr(nDay))
	return int(r)
}

// 日期_取当前时间.
//
// hEle: 元素句柄.
//
// pnHour: 时.[OUT].
//
// pnMinute: 分.[OUT].
//
// pnSecond: 秒.[OUT].
func XDateTime_GetTime(hEle int, pnHour *int32, pnMinute *int32, pnSecond *int32) int {
	r, _, _ := xDateTime_GetTime.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnHour)), uintptr(unsafe.Pointer(pnMinute)), uintptr(unsafe.Pointer(pnSecond)))
	return int(r)
}

// 日期_社区当前时间, 设置当前时分秒.
//
// hEle: 元素句柄.
//
// nHour: 时.
//
// nMinute: 分.
//
// nSecond: 秒.
func XDateTime_SetTime(hEle int, nHour int32, nMinute int32, nSecond int32) int {
	r, _, _ := xDateTime_SetTime.Call(uintptr(hEle), uintptr(nHour), uintptr(nMinute), uintptr(nSecond))
	return int(r)
}

// 日期_弹出, 弹出月历卡片.
//
// hEle: 元素句柄.
func XDateTime_Popup(hEle int) int {
	r, _, _ := xDateTime_Popup.Call(uintptr(hEle))
	return int(r)
}
