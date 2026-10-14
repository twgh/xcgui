package d2d

import "unsafe"

// ID2D1RectangleGeometry 矩形几何.
//
// 它是设备无关资源, 由 ID2D1Factory.CreateRectangleGeometry 创建, 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1rectanglegeometry
type ID2D1RectangleGeometry struct {
	ID2D1Geometry
}

type ID2D1RectangleGeometryVtbl struct {
	ID2D1GeometryVtbl
	GetRect ComProc // 17
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1RectangleGeometry) vtbl() *ID2D1RectangleGeometryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1RectangleGeometryVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetRect 取矩形数据.
func (i *ID2D1RectangleGeometry) GetRect() (D2D1_RECT_F, error) {
	var rc D2D1_RECT_F
	if i == nil {
		return rc, errGeometryNil
	}
	r, _, _ := i.vtbl().GetRect.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&rc)))
	if r != 0 {
		return rc, errno(r)
	}
	return rc, nil
}

// Base 取基类视图(绘制/填充方法要的是 *ID2D1Geometry).
func (i *ID2D1RectangleGeometry) Base() *ID2D1Geometry {
	if i == nil {
		return nil
	}
	return &i.ID2D1Geometry
}
