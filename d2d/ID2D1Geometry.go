package d2d

import (
	"math"
	"unsafe"
)

// ID2D1Geometry 是所有几何对象的基类.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1geometry
type ID2D1Geometry struct {
	ID2D1Resource
}

type ID2D1GeometryVtbl struct {
	ID2D1ResourceVtbl
	GetBounds            ComProc // 4
	GetWidenedBounds     ComProc // 5
	StrokeContainsPoint  ComProc // 6
	FillContainsPoint    ComProc // 7
	CompareWithGeometry  ComProc // 8
	Simplify             ComProc // 9
	Tessellate           ComProc // 10
	CombineWithGeometry  ComProc // 11
	Outline              ComProc // 12
	ComputeArea          ComProc // 13
	ComputeLength        ComProc // 14
	ComputePointAtLength ComProc // 15
	Widen                ComProc // 16
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1Geometry) vtbl() *ID2D1GeometryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1GeometryVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetBounds 取几何的包围盒(不做变换, transform 传 NULL).
//
// 注意: GetBounds 的 transform 是 _In_opt_(可传 NULL); 而 ComputeArea/ComputeLength 的首参
// worldTransform 是 _In_(不能传 NULL, 传了就崩) —— 本包因此只封装了允许 NULL 的 GetBounds.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-getbounds
func (i *ID2D1Geometry) GetBounds() (D2D1_RECT_F, error) {
	var rc D2D1_RECT_F
	if i == nil {
		return rc, errGeometryNil
	}
	r, _, _ := i.vtbl().GetBounds.Call(uintptr(unsafe.Pointer(i)), 0, uintptr(unsafe.Pointer(&rc)))
	if r != 0 {
		return rc, errno(r)
	}
	return rc, nil
}

// finiteRect 判断包围盒是否有限(几何损坏时可能出现 NaN/Inf).
func finiteRect(rc D2D1_RECT_F) bool {
	for _, v := range [4]float32{rc.Left, rc.Top, rc.Right, rc.Bottom} {
		f := float64(v)
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return false
		}
	}
	return true
}
