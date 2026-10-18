package d2d

import (
	"math"
	"testing"
	"unsafe"

	"github.com/twgh/xcgui/wapi/gdi"
	"github.com/twgh/xcgui/wapi/types"
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

func TestMatrixConstructors(t *testing.T) {
	id := IdentityMatrix()
	if id != (D2D1_MATRIX_3X2_F{M11: 1, M22: 1}) {
		t.Errorf("IdentityMatrix = %+v", id)
	}
	tr := TranslationMatrix(10, 20)
	if tr.M11 != 1 || tr.M22 != 1 || tr.DX != 10 || tr.DY != 20 {
		t.Errorf("TranslationMatrix = %+v", tr)
	}
}

func TestPointArg(t *testing.T) {
	got := pointArg(D2D1_POINT_2F{X: 1.5, Y: -2.25})
	want := uint64(math.Float32bits(1.5)) | uint64(math.Float32bits(-2.25))<<32
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
	check("ID2D1Factory.CreateRectangleGeometry", unsafe.Offsetof(factory.CreateRectangleGeometry), 5*p)
	check("ID2D1Factory.CreateRoundedRectangleGeometry", unsafe.Offsetof(factory.CreateRoundedRectangleGeometry), 6*p)
	check("ID2D1Factory.CreateEllipseGeometry", unsafe.Offsetof(factory.CreateEllipseGeometry), 7*p)
	check("ID2D1Factory.CreatePathGeometry", unsafe.Offsetof(factory.CreatePathGeometry), 10*p)
	check("ID2D1Factory.CreateStrokeStyle", unsafe.Offsetof(factory.CreateStrokeStyle), 11*p)
	check("ID2D1Factory.CreateDCRenderTarget", unsafe.Offsetof(factory.CreateDCRenderTarget), 16*p)
	check("ID2D1Factory.Release", unsafe.Offsetof(factory.Release), 2*p)

	var rt ID2D1RenderTargetVtbl
	check("ID2D1RenderTarget.GetFactory", unsafe.Offsetof(rt.GetFactory), 3*p)
	check("ID2D1RenderTarget.CreateSolidColorBrush", unsafe.Offsetof(rt.CreateSolidColorBrush), 8*p)
	check("ID2D1RenderTarget.DrawLine", unsafe.Offsetof(rt.DrawLine), 15*p)
	check("ID2D1RenderTarget.DrawRectangle", unsafe.Offsetof(rt.DrawRectangle), 16*p)
	check("ID2D1RenderTarget.FillRectangle", unsafe.Offsetof(rt.FillRectangle), 17*p)
	check("ID2D1RenderTarget.DrawEllipse", unsafe.Offsetof(rt.DrawEllipse), 20*p)
	check("ID2D1RenderTarget.FillEllipse", unsafe.Offsetof(rt.FillEllipse), 21*p)
	check("ID2D1RenderTarget.DrawGeometry", unsafe.Offsetof(rt.DrawGeometry), 22*p)
	check("ID2D1RenderTarget.FillGeometry", unsafe.Offsetof(rt.FillGeometry), 23*p)
	check("ID2D1RenderTarget.SetTransform", unsafe.Offsetof(rt.SetTransform), 30*p)
	// Flush 之后到 IsSupported 这段是"离屏渲染才发现少列了"的槽(d2d1.h 里它们一直都在):
	// SaveDrawingState 起 6 个 + BeginDraw 补齐之后, 整张虚表才与 SDK 逐槽对齐.
	// 照着 d2d1.h 一个方法一个槽数过来(0=QueryInterface), 顺序错一格就会调到邻居方法.
	check("ID2D1RenderTarget.Flush", unsafe.Offsetof(rt.Flush), 42*p)
	check("ID2D1RenderTarget.SaveDrawingState", unsafe.Offsetof(rt.SaveDrawingState), 43*p)
	check("ID2D1RenderTarget.RestoreDrawingState", unsafe.Offsetof(rt.RestoreDrawingState), 44*p)
	check("ID2D1RenderTarget.PushAxisAlignedClip", unsafe.Offsetof(rt.PushAxisAlignedClip), 45*p)
	check("ID2D1RenderTarget.PopAxisAlignedClip", unsafe.Offsetof(rt.PopAxisAlignedClip), 46*p)
	check("ID2D1RenderTarget.Clear", unsafe.Offsetof(rt.Clear), 47*p)
	check("ID2D1RenderTarget.BeginDraw", unsafe.Offsetof(rt.BeginDraw), 48*p)
	check("ID2D1RenderTarget.EndDraw", unsafe.Offsetof(rt.EndDraw), 49*p)
	check("ID2D1RenderTarget.GetPixelFormat", unsafe.Offsetof(rt.GetPixelFormat), 50*p)
	check("ID2D1RenderTarget.SetDpi", unsafe.Offsetof(rt.SetDpi), 51*p)
	check("ID2D1RenderTarget.GetDpi", unsafe.Offsetof(rt.GetDpi), 52*p)
	check("ID2D1RenderTarget.GetSize", unsafe.Offsetof(rt.GetSize), 53*p)
	check("ID2D1RenderTarget.GetPixelSize", unsafe.Offsetof(rt.GetPixelSize), 54*p)
	check("ID2D1RenderTarget.GetMaximumBitmapSize", unsafe.Offsetof(rt.GetMaximumBitmapSize), 55*p)
	check("ID2D1RenderTarget.IsSupported", unsafe.Offsetof(rt.IsSupported), 56*p)

	// DC 渲染目标: ID2D1RenderTarget 的 57 个槽(0~56)之后就是 BindDC.
	var dcrt ID2D1DCRenderTargetVtbl
	check("ID2D1DCRenderTarget.BindDC", unsafe.Offsetof(dcrt.BindDC), 57*p)

	var stroke ID2D1StrokeStyleVtbl
	check("ID2D1StrokeStyle.GetStartCap", unsafe.Offsetof(stroke.GetStartCap), 4*p)
	check("ID2D1StrokeStyle.GetLineJoin", unsafe.Offsetof(stroke.GetLineJoin), 8*p)
	check("ID2D1StrokeStyle.GetDashStyle", unsafe.Offsetof(stroke.GetDashStyle), 10*p)
	check("ID2D1StrokeStyle.GetDashes", unsafe.Offsetof(stroke.GetDashes), 12*p)

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

	// 三种形状几何: 都以 GetXxx 收尾(槽 17)
	var rect ID2D1RectangleGeometryVtbl
	check("ID2D1RectangleGeometry.GetRect", unsafe.Offsetof(rect.GetRect), 17*p)
	var rounded ID2D1RoundedRectangleGeometryVtbl
	check("ID2D1RoundedRectangleGeometry.GetRoundedRect", unsafe.Offsetof(rounded.GetRoundedRect), 17*p)
	var ellipse ID2D1EllipseGeometryVtbl
	check("ID2D1EllipseGeometry.GetEllipse", unsafe.Offsetof(ellipse.GetEllipse), 17*p)

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
		{"ID2D1DCRenderTarget", unsafe.Sizeof(ID2D1DCRenderTarget{})},
		{"ID2D1Geometry", unsafe.Sizeof(ID2D1Geometry{})},
		{"ID2D1PathGeometry", unsafe.Sizeof(ID2D1PathGeometry{})},
		{"ID2D1RectangleGeometry", unsafe.Sizeof(ID2D1RectangleGeometry{})},
		{"ID2D1RoundedRectangleGeometry", unsafe.Sizeof(ID2D1RoundedRectangleGeometry{})},
		{"ID2D1EllipseGeometry", unsafe.Sizeof(ID2D1EllipseGeometry{})},
		{"ID2D1Brush", unsafe.Sizeof(ID2D1Brush{})},
		{"ID2D1SolidColorBrush", unsafe.Sizeof(ID2D1SolidColorBrush{})},
		{"ID2D1StrokeStyle", unsafe.Sizeof(ID2D1StrokeStyle{})},
		{"ID2D1SimplifiedGeometrySink", unsafe.Sizeof(ID2D1SimplifiedGeometrySink{})},
		{"ID2D1GeometrySink", unsafe.Sizeof(ID2D1GeometrySink{})},
	}
	for _, c := range cases {
		if c.size != ptrSize {
			t.Errorf("%s 大小 = %d 字节, 期望 %d(一个虚表指针) —— 嵌入链上不允许加字段", c.name, c.size, ptrSize)
		}
	}
}

// 独立工厂 + DIB + DC 渲染目标的真机往返: "往离屏位图上画"这条路不依赖 xcgui.dll
// (app.New(false) / 还没建窗口时也能用), 把 SVG 渲染成图片句柄靠的就是它.
// 环境里没有 D2D(d2d1.dll) 时跳过, 不阻塞其它测试.
func TestCreateDCRenderTargetRoundTrip(t *testing.T) {
	f, err := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED)
	if err != nil {
		t.Skipf("d2d1.dll 或 D2D 不可用: %v", err)
	}
	defer f.Release()

	const w, h = 8, 4
	dib, err := gdi.NewDIB(w, h)
	if err != nil {
		t.Fatal(err)
	}
	defer dib.Close()
	if dib.Width() != w || dib.Height() != h {
		t.Errorf("DIB 尺寸 = %dx%d, 想要 %dx%d", dib.Width(), dib.Height(), w, h)
	}
	if got := len(dib.Pixels()); got != w*h*4 {
		t.Fatalf("像素字节数 = %d, 想要 %d", got, w*h*4)
	}

	props := NewDCRenderTargetProperties()
	rt, err := f.CreateDCRenderTarget(&props)
	if err != nil {
		t.Fatal(err)
	}
	defer rt.Release()

	if err := rt.BindDC(dib.HDC(), &types.RECT{Left: 0, Top: 0, Right: w, Bottom: h}); err != nil {
		t.Fatal(err)
	}
	rt.BeginDraw()
	rt.Clear(nil) // 清成透明黑
	rc := D2D1_RECT_F{Left: 0, Top: 0, Right: w / 2, Bottom: h}
	if err := rt.FillRectangleColor(&rc, NewColorF(255, 0, 0, 255)); err != nil {
		t.Fatal(err)
	}
	if err := rt.EndDraw(); err != nil {
		t.Fatal(err)
	}

	px := dib.Pixels()
	// top-down + 预乘: 左上角(已画不透明红)应是 BGRA(0,0,255,255).
	if b, g, r, a := px[0], px[1], px[2], px[3]; b != 0 || g != 0 || r != 255 || a != 255 {
		t.Errorf("左上角像素 = BGRA(%d,%d,%d,%d), 想要 (0,0,255,255)", b, g, r, a)
	}
	// 右半(没画)应还是 Clear 过的透明黑(证明 top-down 行序与 Clear 都生效).
	off := (1*w + (w - 1)) * 4 // 第 2 行最右一列
	if b, g, r, a := px[off], px[off+1], px[off+2], px[off+3]; b != 0 || g != 0 || r != 0 || a != 0 {
		t.Errorf("未绘制区域像素 = BGRA(%d,%d,%d,%d), 想要全 0", b, g, r, a)
	}
}

// 描边路径的端到端探针: 线宽是**按值 float**, 在 amd64 上 DrawRectangle 会把它放进 XMM3(位置 3)、
// 把 style 压到栈[0]; DrawLine 的线宽更靠后, 落在栈[0]. 这两条正是 com.Call 的两种落位分支,
// 抓图能同时看出"寄存器里的线宽生效了"和"压栈参数没串位".
//
// 对 386 同样有意义: 那边全部按序压栈, 8 字节点结构体由 com.U64 拆成两个槽.
func TestStrokePrimitives(t *testing.T) {
	f, err := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED)
	if err != nil {
		t.Skipf("d2d1.dll 或 D2D 不可用: %v", err)
	}
	defer f.Release()

	const w, h = 16, 16
	red := NewColorF(255, 0, 0, 255)

	// draw 在 w×h 的离屏位图上执行一次绘制, 返回 BGRA(预乘) 像素.
	draw := func(t *testing.T, code func(rt *ID2D1DCRenderTarget)) []byte {
		t.Helper()
		dib, err := gdi.NewDIB(w, h)
		if err != nil {
			t.Fatal(err)
		}
		defer dib.Close()
		props := NewDCRenderTargetProperties()
		rt, err := f.CreateDCRenderTarget(&props)
		if err != nil {
			t.Fatal(err)
		}
		defer rt.Release()
		if err := rt.BindDC(dib.HDC(), &types.RECT{Left: 0, Top: 0, Right: w, Bottom: h}); err != nil {
			t.Fatal(err)
		}
		rt.BeginDraw()
		rt.Clear(nil)
		code(rt)
		if err := rt.EndDraw(); err != nil {
			t.Fatal(err)
		}
		// dib 在本函数返回时关闭, 像素要先拷出来, 不能把 DIB 内存里的切片带出去.
		px := make([]byte, len(dib.Pixels()))
		copy(px, dib.Pixels())
		return px
	}
	at := func(px []byte, x, y int) (b, g, r, a byte) {
		o := (y*w + x) * 4
		return px[o], px[o+1], px[o+2], px[o+3]
	}
	isRed := func(t *testing.T, px []byte, x, y int, what string) {
		t.Helper()
		if b, g, r, a := at(px, x, y); r < 200 || g > 60 || b > 60 || a < 200 {
			t.Errorf("%s 的像素 (%d,%d) = BGRA(%d,%d,%d,%d), 想要不透明红", what, x, y, b, g, r, a)
		}
	}
	isEmpty := func(t *testing.T, px []byte, x, y int, what string) {
		t.Helper()
		if b, g, r, a := at(px, x, y); b|g|r|a != 0 {
			t.Errorf("%s 的像素 (%d,%d) = BGRA(%d,%d,%d,%d), 想要全 0", what, x, y, b, g, r, a)
		}
	}

	// 线宽 4 的矩形描边: 上边框覆盖 y∈[0,4), 内部(8,8)不该被填.
	px := draw(t, func(rt *ID2D1DCRenderTarget) {
		rc := D2D1_RECT_F{Left: 2, Top: 2, Right: 14, Bottom: 14}
		if err := rt.DrawRectangleColor(&rc, red, 4, nil); err != nil {
			t.Error(err)
		}
	})
	isRed(t, px, 8, 2, "DrawRectangle 上边框")
	isEmpty(t, px, 8, 8, "DrawRectangle 内部")

	// 线宽 4 的直线: 位置 4 → amd64 的栈[0], 覆盖 y∈[6,10).
	px = draw(t, func(rt *ID2D1DCRenderTarget) {
		if err := rt.DrawLineColor(D2D1_POINT_2F{X: 0, Y: 8}, D2D1_POINT_2F{X: 16, Y: 8}, red, 4, nil); err != nil {
			t.Error(err)
		}
	})
	isRed(t, px, 8, 8, "DrawLine 线身")
	isEmpty(t, px, 8, 0, "DrawLine 线外")
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
	if path.Base() != nil {
		t.Error("nil 几何的 Base 应返回 nil")
	}
	var sink *ID2D1GeometrySink
	if sink.Release() != 0 {
		t.Error("nil 写入器 Release 应返回 0")
	}

	// Usable 定义在 ID2D1Geometry 上(自有方法), 只对基类指针判空安全:
	// 对 nil 的 *ID2D1PathGeometry 调它, 编译器要先解引用到嵌入字段, 会 panic —— 与其它被提升的
	// 基类方法一个语义(调用方拿到 nil 必须先判空).
	var g *ID2D1Geometry
	if g.Usable() {
		t.Error("nil 几何不该 Usable")
	}

	// 三种形状几何同样 nil 安全(Release / Base)
	var rg *ID2D1RectangleGeometry
	if rg.Release() != 0 || rg.Base() != nil {
		t.Error("nil 矩形几何的 Release/Base 行为不对")
	}
	var rr *ID2D1RoundedRectangleGeometry
	if rr.Release() != 0 || rr.Base() != nil {
		t.Error("nil 圆角矩形几何的 Release/Base 行为不对")
	}
	var eg *ID2D1EllipseGeometry
	if eg.Release() != 0 || eg.Base() != nil {
		t.Error("nil 椭圆几何的 Release/Base 行为不对")
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
	if (*ID2D1Geometry)(nil).Usable() { // Usable 是基类的自有方法, 只对基类指针判空安全
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

	var style *ID2D1StrokeStyle
	if style.Release() != 0 {
		t.Error("nil 描边样式 Release 应返回 0")
	}
	if style.GetLineJoin() != 0 {
		t.Error("nil 描边样式 GetLineJoin 应返回 0")
	}

	// 离屏渲染那条路(独立工厂 → DIB → DC 渲染目标)在 nil 上同样只报错不崩.
	var dcf *ID2D1Factory
	if _, err := dcf.CreateDCRenderTarget(nil); err == nil {
		t.Error("空工厂建 DC 渲染目标应报错")
	}
	var dcrt *ID2D1DCRenderTarget
	if dcrt.Release() != 0 {
		t.Error("nil DC 渲染目标 Release 应返回 0")
	}
	if err := dcrt.BindDC(0, nil); err == nil {
		t.Error("nil DC 渲染目标 BindDC 应报错")
	}
	// EndDraw/Clear/BeginDraw 定义在 ID2D1RenderTarget 上, 对 nil 的 **DC 渲染目标**调用会先
	// 解引用嵌入字段而 panic(值嵌入的语义) —— 这里按基类指针测.
	var baseRT *ID2D1RenderTarget
	if err := baseRT.EndDraw(); err == nil {
		t.Error("nil 渲染目标 EndDraw 应报错")
	}
	baseRT.BeginDraw() // void, 只判空不崩
	baseRT.Clear(nil)
	var d *gdi.DIB
	if d.HDC() != 0 || d.Width() != 0 || d.Height() != 0 || d.Pixels() != nil {
		t.Error("nil DIB 的访问器应返回零值")
	}
	d.Close() // nil 安全
}
