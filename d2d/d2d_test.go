package d2d

import (
	"math"
	"testing"
	"unsafe"
)

func TestNewColorF(t *testing.T) {
	c := NewColorF(255, 128, 0, 64)
	if c.R != 1 {
		t.Errorf("R = %v, 期望 1", c.R)
	}
	if c.G < 0.50 || c.G > 0.51 {
		t.Errorf("G = %v, 期望 ~0.502", c.G)
	}
	if c.B != 0 {
		t.Errorf("B = %v, 期望 0", c.B)
	}
	if c.A < 0.25 || c.A > 0.26 {
		t.Errorf("A = %v, 期望 ~0.251", c.A)
	}
}

func TestPointArg(t *testing.T) {
	got := pointArg(D2D1_POINT_2F{X: 1.5, Y: -2.25})
	want := uintptr(math.Float32bits(1.5)) | uintptr(math.Float32bits(-2.25))<<32
	if got != want {
		t.Errorf("pointArg = %#x, 期望 %#x", got, want)
	}
}

// 虚表布局就是 ABI 本身: 索引错一位就会调到邻居方法(实测踩过 —— 把 SetTransform 当
// GetAntialiasMode 调, 用垃圾指针设了变换, 整窗不再呈现). 这里把"方法在虚表里的第几个槽"
// 固化成断言, 防止将来有人插字段/改顺序.
func TestVtblIndexLayout(t *testing.T) {
	p := unsafe.Sizeof(uintptr(0))
	check := func(name string, got, want uintptr) {
		t.Helper()
		if got != want {
			t.Errorf("%s 在虚表里的偏移 = %d 字节(第 %d 个槽), 期望第 %d 个槽", name, got, got/p, want/p)
		}
	}

	var factory ID2D1FactoryVtbl
	check("ID2D1Factory.CreatePathGeometry", unsafe.Offsetof(factory.CreatePathGeometry), 10*p)
	check("ID2D1Factory.Release", unsafe.Offsetof(factory.Release), 2*p)

	var rt ID2D1RenderTargetVtbl
	check("ID2D1RenderTarget.GetFactory", unsafe.Offsetof(rt.GetFactory), 3*p)
	check("ID2D1RenderTarget.CreateSolidColorBrush", unsafe.Offsetof(rt.CreateSolidColorBrush), 8*p)
	check("ID2D1RenderTarget.FillGeometry", unsafe.Offsetof(rt.FillGeometry), 23*p)
	check("ID2D1RenderTarget.SetTransform", unsafe.Offsetof(rt.SetTransform), 30*p)
	check("ID2D1RenderTarget.IsSupported", unsafe.Offsetof(rt.IsSupported), 50*p)

	var path ID2D1PathGeometryVtbl
	check("ID2D1PathGeometry.GetBounds", unsafe.Offsetof(path.GetBounds), 4*p)
	check("ID2D1PathGeometry.Open", unsafe.Offsetof(path.Open), 17*p)
	check("ID2D1PathGeometry.GetFigureCount", unsafe.Offsetof(path.GetFigureCount), 20*p)

	var sink ID2D1GeometrySinkVtbl
	check("ID2D1GeometrySink.SetFillMode", unsafe.Offsetof(sink.SetFillMode), 3*p)
	check("ID2D1GeometrySink.BeginFigure", unsafe.Offsetof(sink.BeginFigure), 5*p)
	check("ID2D1GeometrySink.AddLines", unsafe.Offsetof(sink.AddLines), 6*p)
	check("ID2D1GeometrySink.EndFigure", unsafe.Offsetof(sink.EndFigure), 8*p)
	check("ID2D1GeometrySink.Close", unsafe.Offsetof(sink.Close), 9*p)
	check("ID2D1GeometrySink.AddArc", unsafe.Offsetof(sink.AddArc), 14*p)

	var brush ID2D1SolidColorBrushVtbl
	check("ID2D1SolidColorBrush.SetOpacity", unsafe.Offsetof(brush.SetOpacity), 4*p)
	check("ID2D1SolidColorBrush.SetColor", unsafe.Offsetof(brush.SetColor), 8*p)
}

// 每个接口结构体必须恰好一个指针大小: 它的地址就是 COM 对象地址, 它的首字段就是真实虚表指针 ——
// 各接口的 vtbl() 重解释、以及"嵌入基类复用方法"都建立在这条上(d2d.go 还有一份编译期断言).
func TestStructLayout(t *testing.T) {
	ptrSize := unsafe.Sizeof(uintptr(0))
	cases := []struct {
		name string
		size uintptr
	}{
		{"IUnknown", unsafe.Sizeof(IUnknown{})},
		{"ID2D1Resource", unsafe.Sizeof(ID2D1Resource{})},
		{"ID2D1Factory", unsafe.Sizeof(ID2D1Factory{})},
		{"ID2D1RenderTarget", unsafe.Sizeof(ID2D1RenderTarget{})},
		{"ID2D1Geometry", unsafe.Sizeof(ID2D1Geometry{})},
		{"ID2D1PathGeometry", unsafe.Sizeof(ID2D1PathGeometry{})},
		{"ID2D1Brush", unsafe.Sizeof(ID2D1Brush{})},
		{"ID2D1SolidColorBrush", unsafe.Sizeof(ID2D1SolidColorBrush{})},
		{"ID2D1SimplifiedGeometrySink", unsafe.Sizeof(ID2D1SimplifiedGeometrySink{})},
		{"ID2D1GeometrySink", unsafe.Sizeof(ID2D1GeometrySink{})},
	}
	for _, c := range cases {
		if c.size != ptrSize {
			t.Errorf("%s 大小 = %d 字节, 期望 %d(一个虚表指针) —— 嵌入链上不允许加字段", c.name, c.size, ptrSize)
		}
	}
}

// 嵌入让子接口直接继承基类方法(编译期保证): AddRef/Release/QueryInterface 只在 IUnknown 上实现一次,
// GetFactory 只在 ID2D1Resource 上实现一次.
//
// 运行期只承诺 **Release 在 nil 上安全** —— 值嵌入下"被提升的基类方法"在 nil 外层指针上会先解引用到
// 嵌入字段而 panic(这是 Go 的语义, 基类方法里判空也救不了), 所以构造拿到 nil 时调用方必须判空.
// 自有方法(Usable/Open/FillGeometry/... 定义在本类型上的)都各自判了空.
func TestMethodReuseThroughEmbedding(t *testing.T) {
	var brush *ID2D1SolidColorBrush
	if brush.Release() != 0 {
		t.Error("nil 画刷 Release 应返回 0")
	}
	var rt *ID2D1RenderTarget
	if rt.Release() != 0 {
		t.Error("nil 渲染目标 Release 应返回 0")
	}
	var path *ID2D1PathGeometry
	if path.Release() != 0 {
		t.Error("nil 几何 Release 应返回 0")
	}
	if path.Usable() { // Usable 是自己判空的自有方法
		t.Error("nil 几何不该 Usable")
	}
	var sink *ID2D1GeometrySink
	if sink.Release() != 0 {
		t.Error("nil 写入器 Release 应返回 0")
	}
}

// vtbl() 必须取"对象里存的虚表指针"(i.Vtbl), 而不是把对象地址当成虚表 —— 后者会读到对象内存
// 偏移 8/16/... 字节处的垃圾当函数指针, 真实渲染时表现为跳到非法地址.
// (实测血案: demo 启动即 panic, 堆栈 PC=0x989680, 全屏没有一个图标画出来.)
func TestVtblResolution(t *testing.T) {
	factoryVtbl := ID2D1FactoryVtbl{}
	factory := ID2D1Factory{}
	factory.Vtbl = (*IUnknownVtbl)(unsafe.Pointer(&factoryVtbl))
	if got := factory.vtbl(); got != &factoryVtbl {
		t.Errorf("ID2D1Factory.vtbl() = %p, 期望 %p", got, &factoryVtbl)
	}

	pathVtbl := ID2D1PathGeometryVtbl{}
	path := ID2D1PathGeometry{}
	path.Vtbl = (*IUnknownVtbl)(unsafe.Pointer(&pathVtbl))
	if got := path.vtbl(); got != &pathVtbl {
		t.Errorf("ID2D1PathGeometry.vtbl() = %p, 期望 %p", got, &pathVtbl)
	}

	brushVtbl := ID2D1SolidColorBrushVtbl{}
	brush := ID2D1SolidColorBrush{}
	brush.Vtbl = (*IUnknownVtbl)(unsafe.Pointer(&brushVtbl))
	if got := brush.vtbl(); got != &brushVtbl {
		t.Errorf("ID2D1SolidColorBrush.vtbl() = %p, 期望 %p", got, &brushVtbl)
	}

	sinkVtbl := ID2D1GeometrySinkVtbl{}
	sink := ID2D1GeometrySink{}
	sink.Vtbl = (*IUnknownVtbl)(unsafe.Pointer(&sinkVtbl))
	if got := sink.vtbl(); got != &sinkVtbl {
		t.Errorf("ID2D1GeometrySink.vtbl() = %p, 期望 %p", got, &sinkVtbl)
	}
}

// 未开启 D2D(GDI 渲染)时句柄为 0, 构造函数返回 nil; 调用方判空后走回落路径.
func TestZeroHandleGivesNil(t *testing.T) {
	if NewFactoryByHandle(0) != nil {
		t.Error("句柄为 0 应返回 nil 工厂")
	}
	if NewRenderTargetByHandle(0) != nil {
		t.Error("句柄为 0 应返回 nil 渲染目标")
	}
}

// nil 对象上的方法必须安全失败(而不是把空指针喂给 syscall 直接崩进程).
func TestNilSafety(t *testing.T) {
	var factory *ID2D1Factory
	if _, err := factory.NewGeometry([][]float32{{0, 0, 10, 0, 10, 10}}); err == nil {
		t.Error("空工厂建几何应报错")
	}
	if _, err := factory.CreatePathGeometry(); err == nil {
		t.Error("空工厂建几何应报错")
	}

	var rt *ID2D1RenderTarget
	if _, err := rt.CreateSolidColorBrush(&D2D1_COLOR_F{}); err == nil {
		t.Error("空渲染目标建画刷应报错")
	}
	if err := rt.FillGeometry(nil, nil); err == nil {
		t.Error("空渲染目标填充应报错")
	}
	if err := rt.FillGeometryColor(nil, NewColorF(0, 0, 0, 255)); err == nil {
		t.Error("空渲染目标填充应报错")
	}

	var geometry *ID2D1PathGeometry
	if geometry.Usable() {
		t.Error("nil 几何不该 Usable")
	}
	geometry.Release()
	if _, err := geometry.Open(); err == nil {
		t.Error("nil 几何 Open 应报错")
	}

	var sink *ID2D1GeometrySink
	sink.SetFillMode(D2D1_FILL_MODE_WINDING)
	sink.BeginFigure(D2D1_POINT_2F{}, D2D1_FIGURE_BEGIN_FILLED)
	sink.AddLines([]float32{1, 2, 3, 4})
	sink.AddLines(nil) // 点数不足直接返回
	sink.EndFigure(D2D1_FIGURE_END_CLOSED)
	if err := sink.Close(); err == nil {
		t.Error("nil 几何写入器 Close 应报错")
	}
	sink.Release()

	var brush *ID2D1SolidColorBrush
	brush.SetColor(NewColorF(1, 2, 3, 4))
	brush.Release()

	var base *ID2D1Brush
	if base.Release() != 0 {
		t.Error("nil 画刷 Release 应返回 0")
	}
}
