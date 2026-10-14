package d2d

import "unsafe"

// ID2D1RoundedRectangleGeometry 圆角矩形几何.
//
// 它是设备无关资源, 由 ID2D1Factory.CreateRoundedRectangleGeometry 创建, 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1roundedrectanglegeometry
type ID2D1RoundedRectangleGeometry struct {
	ID2D1Geometry
}

type ID2D1RoundedRectangleGeometryVtbl struct {
	ID2D1GeometryVtbl
	GetRoundedRect ComProc // 17
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1RoundedRectangleGeometry) vtbl() *ID2D1RoundedRectangleGeometryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1RoundedRectangleGeometryVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetRoundedRect 取圆角矩形数据.
func (i *ID2D1RoundedRectangleGeometry) GetRoundedRect() (D2D1_ROUNDED_RECT, error) {
	var rr D2D1_ROUNDED_RECT
	if i == nil {
		return rr, errGeometryNil
	}
	r, _, _ := i.vtbl().GetRoundedRect.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&rr)))
	if r != 0 {
		return rr, errno(r)
	}
	return rr, nil
}

// Base 取基类视图(绘制/填充方法要的是 *ID2D1Geometry).
func (i *ID2D1RoundedRectangleGeometry) Base() *ID2D1Geometry {
	if i == nil {
		return nil
	}
	return &i.ID2D1Geometry
}
