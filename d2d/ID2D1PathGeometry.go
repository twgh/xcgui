package d2d

import "unsafe"

// ID2D1PathGeometry 是路径几何: 由若干子路径(直线/贝塞尔/弧)拼成的形状.
//
// 它是设备无关资源, 由本包创建(见 ID2D1Factory.CreatePathGeometry), 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1pathgeometry
type ID2D1PathGeometry struct {
	ID2D1Geometry
}

type ID2D1PathGeometryVtbl struct {
	ID2D1GeometryVtbl
	Open            ComProc // 17
	Stream          ComProc // 18
	GetSegmentCount ComProc // 19
	GetFigureCount  ComProc // 20
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1PathGeometry) vtbl() *ID2D1PathGeometryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1PathGeometryVtbl)(unsafe.Pointer(i.Vtbl))
}

// Open 获取几何写入器, 用于往里写子路径; 写完后必须 Close.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1pathgeometry-open
func (i *ID2D1PathGeometry) Open() (*ID2D1GeometrySink, error) {
	if i == nil {
		return nil, errGeometryNil
	}
	var sink *ID2D1GeometrySink
	r, _, _ := i.vtbl().Open.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&sink)))
	if r != 0 {
		return nil, errno(r)
	}
	return sink, nil
}

// Usable 用 ID2D1Geometry::GetBounds 判断几何是否真的可用(HRESULT 成功且包围盒有限).
//
// 为什么必须有这一关: 索引错位或**没 Close 成功**的几何喂给 FillGeometry 会让 xcgui 整帧不
// 呈现 —— 窗口彻底隐形(连边框都没有)、进程不崩、日志正常, 只信截图. 未 Close 的几何在这里
// 返回 D2DERR_WRONG_STATE(0x88990001).
func (i *ID2D1PathGeometry) Usable() bool {
	if i == nil {
		return false
	}
	rc, err := i.GetBounds() // 继承自嵌入的 ID2D1Geometry
	return err == nil && finiteRect(rc)
}
