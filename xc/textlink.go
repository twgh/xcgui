package xc

import "github.com/twgh/xcgui/common"

// 文本链接_创建, 创建静态文本链接元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func XTextLink_Create(x int, y int, cx int, cy int, pName string, hParent int) int {
	r, _, _ := xTextLink_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pName), uintptr(hParent))
	return int(r)
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XTextLink_EnableUnderlineLeave(hEle int, bEnable bool) int {
	r, _, _ := xTextLink_EnableUnderlineLeave.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XTextLink_EnableUnderlineStay(hEle int, bEnable bool) int {
	r, _, _ := xTextLink_EnableUnderlineStay.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XTextLink_SetTextColorStay(hEle int, color int) int {
	r, _, _ := xTextLink_SetTextColorStay.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XTextLink_SetUnderlineColorLeave(hEle int, color int) int {
	r, _, _ := xTextLink_SetUnderlineColorLeave.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XTextLink_SetUnderlineColorStay(hEle int, color int) int {
	r, _, _ := xTextLink_SetUnderlineColorStay.Call(uintptr(hEle), uintptr(color))
	return int(r)
}
