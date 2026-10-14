package d2d

import "unsafe"

// ID2D1SimplifiedGeometrySink 是几何写入器的基类.
//
// 注意: 除 Close 以外的写入方法在 COM 里都是 STDMETHOD_(void, ...), 返回值是寄存器残留垃圾,
// 本包因此既不会返回也不检查它们的"HRESULT"(照直觉检查会静默失败, 表现为"代码改了画面没变").
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1simplifiedgeometrysink
type ID2D1SimplifiedGeometrySink struct {
	IUnknown
}

type ID2D1SimplifiedGeometrySinkVtbl struct {
	IUnknownVtbl
	SetFillMode     ComProc // 3
	SetSegmentFlags ComProc // 4
	BeginFigure     ComProc // 5
	AddLines        ComProc // 6
	AddBeziers      ComProc // 7
	EndFigure       ComProc // 8
	Close           ComProc // 9
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1SimplifiedGeometrySink) vtbl() *ID2D1SimplifiedGeometrySinkVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1SimplifiedGeometrySinkVtbl)(unsafe.Pointer(i.Vtbl))
}
