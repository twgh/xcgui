package xc

import "github.com/twgh/xcgui/xcc"

// XObj_GetType 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
//
// hXCGUI: 对象句柄.
func XObj_GetType(hXCGUI int) xcc.XC_OBJECT_TYPE {
	r, _, _ := xObj_GetType.Call(uintptr(hXCGUI))
	return xcc.XC_OBJECT_TYPE(r)
}

// XObj_GetTypeBase 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
//
// hXCGUI: 对象句柄.
func XObj_GetTypeBase(hXCGUI int) int {
	r, _, _ := xObj_GetTypeBase.Call(uintptr(hXCGUI))
	return int(r)
}

// XObj_GetTypeEx 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
//
// hXCGUI: 对象句柄.
func XObj_GetTypeEx(hXCGUI int) xcc.XC_OBJECT_TYPE_EX {
	r, _, _ := xObj_GetTypeEx.Call(uintptr(hXCGUI))
	return xcc.XC_OBJECT_TYPE_EX(r)
}

// XObj_SetTypeEx 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// hXCGUI: 对象句柄.
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func XObj_SetTypeEx(hXCGUI int, nType xcc.XC_OBJECT_TYPE_EX) int {
	r, _, _ := xObj_SetTypeEx.Call(uintptr(hXCGUI), uintptr(nType))
	return int(r)
}
