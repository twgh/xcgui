package objectbase

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ObjectBase 炫彩对象基类.
type ObjectBase struct {
	Handle int // 句柄.
}

// 给本类的Handle赋值.
func (o *ObjectBase) SetHandle(handle int) *ObjectBase {
	o.Handle = handle
	return o
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
func (o *ObjectBase) GetType() xcc.XC_OBJECT_TYPE {
	return xc.XObj_GetType(o.Handle)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
func (o *ObjectBase) GetTypeBase() xcc.XC_OBJECT_TYPE {
	return xc.XObj_GetTypeBase(o.Handle)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) GetTypeEx() xcc.XC_OBJECT_TYPE_EX {
	return xc.XObj_GetTypeEx(o.Handle)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) SetTypeEx(nType xcc.XC_OBJECT_TYPE_EX) *ObjectBase {
	xc.XObj_SetTypeEx(o.Handle, nType)
	return o
}

// 炫彩_置属性, 设置对象属性.
//
// name: 属性名.
//
// pValue: 属性值.
func (o *ObjectBase) SetProperty(name string, pValue string) bool {
	return xc.XC_SetProperty(o.Handle, name, pValue)
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// name: 属性名.
func (o *ObjectBase) GetProperty(name string) string {
	return xc.XC_GetProperty(o.Handle, name)
}
