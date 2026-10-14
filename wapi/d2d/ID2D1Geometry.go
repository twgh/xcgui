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

// Usable 用 ID2D1Geometry::GetBounds 判断几何是否真的可用(HRESULT 成功且包围盒有限).
//
// 为什么必须有这一关: 索引错位或**没 Close 成功**的几何喂给 FillGeometry 会让 xcgui 整帧不
// 呈现 —— 窗口彻底隐形(连边框都没有)、进程不崩、日志正常, 只信截图. 未 Close 的几何在这里
// 返回 D2DERR_WRONG_STATE(0x88990001).
//
// 本包所有"画几何"的方法(绘制/填充)内部都会先过这一关.
func (i *ID2D1Geometry) Usable() bool {
	if i == nil {
		return false
	}
	rc, err := i.GetBounds()
	return err == nil && finiteRect(rc)
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
