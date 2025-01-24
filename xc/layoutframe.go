package xc

import "github.com/twgh/xcgui/common"

// 布局框架_创建.
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
func XLayoutFrame_Create(x, y, cx, cy int32, hParent int) int {
	r, _, _ := xLayoutFrame_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 布局框架_显示布局边界.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XLayoutFrame_ShowLayoutFrame(hEle int, bEnable bool) {
	xLayoutFrame_ShowLayoutFrame.Call(uintptr(hEle), common.BoolPtr(bEnable))
}
