// 基类
package objectbase

import "github.com/twgh/xcgui/xc"

// 炫彩对象基类
type ObjectBase struct {
	Handle int // 句柄
}

// 给本类的Handle赋值
func (o *ObjectBase) SetHandle(handle int) {
	o.Handle = handle
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_
func (o *ObjectBase) GetType() int {
	return xc.XObj_GetType(o.Handle)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
func (o *ObjectBase) GetTypeBase() int {
	return xc.XObj_GetTypeBase(o.Handle)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error
func (o *ObjectBase) GetTypeEx() int {
	return xc.XObj_GetTypeEx(o.Handle)
}
