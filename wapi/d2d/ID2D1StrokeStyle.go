package d2d

import "unsafe"

// ID2D1StrokeStyle 描述描边样式(线帽 / 连接 / 虚线), 由 ID2D1Factory.CreateStrokeStyle 创建.
//
// 它是**设备无关**资源(归属工厂), 可以长期缓存复用, 但用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1strokestyle
type ID2D1StrokeStyle struct {
	ID2D1Resource
}

type ID2D1StrokeStyleVtbl struct {
	ID2D1ResourceVtbl
	GetStartCap    ComProc // 4
	GetEndCap      ComProc // 5
	GetDashCap     ComProc // 6
	GetMiterLimit  ComProc // 7 —— 按值返回 FLOAT(在 XMM0 里), syscall 读不到, 故不封装
	GetLineJoin    ComProc // 8
	GetDashOffset  ComProc // 9 —— 同样按值返回 FLOAT
	GetDashStyle   ComProc // 10
	GetDashesCount ComProc // 11
	GetDashes      ComProc // 12
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1StrokeStyle) vtbl() *ID2D1StrokeStyleVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1StrokeStyleVtbl)(unsafe.Pointer(i.Vtbl))
}

// GetStartCap 取起点线帽样式.
func (i *ID2D1StrokeStyle) GetStartCap() D2D1_CAP_STYLE {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetStartCap.Call(uintptr(unsafe.Pointer(i)))
	return D2D1_CAP_STYLE(int32(r))
}

// GetEndCap 取终点线帽样式.
func (i *ID2D1StrokeStyle) GetEndCap() D2D1_CAP_STYLE {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetEndCap.Call(uintptr(unsafe.Pointer(i)))
	return D2D1_CAP_STYLE(int32(r))
}

// GetDashCap 取虚线端点样式.
func (i *ID2D1StrokeStyle) GetDashCap() D2D1_DASH_CAP_STYLE {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetDashCap.Call(uintptr(unsafe.Pointer(i)))
	return D2D1_DASH_CAP_STYLE(int32(r))
}

// GetLineJoin 取连接样式.
func (i *ID2D1StrokeStyle) GetLineJoin() D2D1_LINE_JOIN {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetLineJoin.Call(uintptr(unsafe.Pointer(i)))
	return D2D1_LINE_JOIN(int32(r))
}

// GetDashStyle 取虚线样式.
func (i *ID2D1StrokeStyle) GetDashStyle() D2D1_DASH_STYLE {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetDashStyle.Call(uintptr(unsafe.Pointer(i)))
	return D2D1_DASH_STYLE(int32(r))
}

// GetDashesCount 取自定义虚线段数.
func (i *ID2D1StrokeStyle) GetDashesCount() uint32 {
	if i == nil {
		return 0
	}
	r, _, _ := i.vtbl().GetDashesCount.Call(uintptr(unsafe.Pointer(i)))
	return uint32(r)
}
