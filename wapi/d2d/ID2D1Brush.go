package d2d

import "unsafe"

// ID2D1Brush 是所有画刷的基类.
//
// 画刷是**设备相关**资源: 只能给创建它的渲染目标使用, 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1brush
type ID2D1Brush struct {
	ID2D1Resource
}

type ID2D1BrushVtbl struct {
	ID2D1ResourceVtbl
	SetOpacity   ComProc // 4 (void; 形参是**按值 float** → 要用 wapi/com 的转发器, 见包注释 ABI 规则 1)
	SetTransform ComProc // 5 (void; 形参是 D2D1_MATRIX_3X2_F **指针**, 可直接调 —— 本包暂未封装)
	GetOpacity   ComProc // 6 (注意: **按值返回 FLOAT**, 在 XMM0 里, syscall 读不到 —— 要 com.CallWithFloat32RetFloat)
	GetTransform ComProc // 7 (void; 出参 D2D1_MATRIX_3X2_F*)
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1Brush) vtbl() *ID2D1BrushVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1BrushVtbl)(unsafe.Pointer(i.Vtbl))
}
