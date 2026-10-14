package d2d

import "unsafe"

// ID2D1Resource 是所有 Direct2D 资源的基类.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type ID2D1Resource struct {
	IUnknown
}

type ID2D1ResourceVtbl struct {
	IUnknownVtbl
	GetFactory ComProc // 3
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1Resource) vtbl() *ID2D1ResourceVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1ResourceVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetFactory 获取创建该资源的工厂.
func (i *ID2D1Resource) GetFactory() *ID2D1Factory {
	if i == nil {
		return nil
	}
	var factory *ID2D1Factory
	i.vtbl().GetFactory.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&factory)))
	return factory
}
