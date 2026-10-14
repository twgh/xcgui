package d2d

import "unsafe"

// ID2D1EllipseGeometry 椭圆几何.
//
// 它是设备无关资源, 由 ID2D1Factory.CreateEllipseGeometry 创建, 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1ellipsegeometry
type ID2D1EllipseGeometry struct {
	ID2D1Geometry
}

type ID2D1EllipseGeometryVtbl struct {
	ID2D1GeometryVtbl
	GetEllipse ComProc // 17
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1EllipseGeometry) vtbl() *ID2D1EllipseGeometryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1EllipseGeometryVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetEllipse 取椭圆数据.
func (i *ID2D1EllipseGeometry) GetEllipse() (D2D1_ELLIPSE, error) {
	var e D2D1_ELLIPSE
	if i == nil {
		return e, errGeometryNil
	}
	r, _, _ := i.vtbl().GetEllipse.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&e)))
	if r != 0 {
		return e, errno(r)
	}
	return e, nil
}

// Base 取基类视图(绘制/填充方法要的是 *ID2D1Geometry).
func (i *ID2D1EllipseGeometry) Base() *ID2D1Geometry {
	if i == nil {
		return nil
	}
	return &i.ID2D1Geometry
}
