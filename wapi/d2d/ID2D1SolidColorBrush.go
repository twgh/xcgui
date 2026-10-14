package d2d

import "unsafe"

// ID2D1SolidColorBrush 是纯色画刷.
//
// 由本包创建(见 ID2D1RenderTarget.CreateSolidColorBrush), 用完必须 Release.
// 它嵌入 ID2D1Brush, 因此需要基类指针时直接取 `&brush.ID2D1Brush` 即可(内存布局与 COM 对象一致).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1solidcolorbrush
type ID2D1SolidColorBrush struct {
	ID2D1Brush
}

type ID2D1SolidColorBrushVtbl struct {
	ID2D1BrushVtbl
	SetColor ComProc // 8 (void)
	GetColor ComProc // 9 (void, 出参)
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1SolidColorBrush) vtbl() *ID2D1SolidColorBrushVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1SolidColorBrushVtbl)(unsafe.Pointer(i.Vtbl))
}

// SetColor 改画刷颜色(void 方法; 画刷已创建的情况下改色比重新建一个更省).
func (i *ID2D1SolidColorBrush) SetColor(color D2D1_COLOR_F) {
	if i == nil {
		return
	}
	i.vtbl().SetColor.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&color)))
}
