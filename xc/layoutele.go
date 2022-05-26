package xc

import "github.com/twgh/xcgui/common"

// 布局_创建, 创建布局元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func XLayout_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xLayout_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 布局_创建扩展, 创建布局元素, 返回元素句柄.
//
// hParent: 父为窗口句柄或元素句柄.
func XLayout_CreateEx(hParent int) int {
	r, _, _ := xLayout_CreateEx.Call(uintptr(hParent))
	return int(r)
}

// 布局_判断启用, 是否已经启用布局功能.
//
// hEle: 元素句柄.
func XLayout_IsEnableLayout(hEle int) bool {
	r, _, _ := xLayout_IsEnableLayout.Call(uintptr(hEle))
	return r != 0
}

// 布局_启用, 启用布局功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XLayout_EnableLayout(hEle int, bEnable bool) int {
	r, _, _ := xLayout_EnableLayout.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 布局_显示布局边界, 显示布局边界.
//
// hEle: 元素句柄.
//
// bEnable: 是否显示.
func XLayout_ShowLayoutFrame(hEle int, bEnable bool) int {
	r, _, _ := xLayout_ShowLayoutFrame.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 布局_取内宽度, 获取宽度,不包含内边距大小.
//
// hEle: 元素句柄.
func XLayout_GetWidthIn(hEle int) int {
	r, _, _ := xLayout_GetWidthIn.Call(uintptr(hEle))
	return int(r)
}

// 布局_取内高度, 获取高度,不包含内边距大小.
//
// hEle: 元素句柄.
func XLayout_GetHeightIn(hEle int) int {
	r, _, _ := xLayout_GetHeightIn.Call(uintptr(hEle))
	return int(r)
}
