package d2d

import (
	"errors"
	"math"
	"unsafe"

	"github.com/twgh/xcgui/wapi/com"
)

// ID2D1RenderTarget 是 Direct2D 的渲染目标, 所有绘制调用都落在它上面.
//
// 它由 xcgui 的绘制模块持有(见 xc.XDraw_GetD2dRenderTarget), 这里只是借用, 不要 Release;
// 只在绘制回调内使用.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type ID2D1RenderTarget struct {
	ID2D1Resource
}

type ID2D1RenderTargetVtbl struct {
	ID2D1ResourceVtbl
	CreateBitmap                 ComProc // 4
	CreateBitmapFromWicBitmap    ComProc // 5
	CreateSharedBitmap           ComProc // 6
	CreateBitmapBrush            ComProc // 7
	CreateSolidColorBrush        ComProc // 8
	CreateGradientStopCollection ComProc // 9
	CreateLinearGradientBrush    ComProc // 10
	CreateRadialGradientBrush    ComProc // 11
	CreateCompatibleRenderTarget ComProc // 12
	CreateLayer                  ComProc // 13
	CreateMesh                   ComProc // 14
	DrawLine                     ComProc // 15
	DrawRectangle                ComProc // 16
	FillRectangle                ComProc // 17
	DrawRoundedRectangle         ComProc // 18
	FillRoundedRectangle         ComProc // 19
	DrawEllipse                  ComProc // 20
	FillEllipse                  ComProc // 21
	DrawGeometry                 ComProc // 22
	FillGeometry                 ComProc // 23
	FillMesh                     ComProc // 24
	FillOpacityMask              ComProc // 25
	DrawBitmap                   ComProc // 26
	DrawText                     ComProc // 27
	DrawTextLayout               ComProc // 28
	DrawGlyphRun                 ComProc // 29
	SetTransform                 ComProc // 30
	GetTransform                 ComProc // 31
	SetAntialiasMode             ComProc // 32
	GetAntialiasMode             ComProc // 33
	SetTextAntialiasMode         ComProc // 34
	GetTextAntialiasMode         ComProc // 35
	SetTextRenderingParams       ComProc // 36
	GetTextRenderingParams       ComProc // 37
	SetTags                      ComProc // 38
	GetTags                      ComProc // 39
	PushLayer                    ComProc // 40
	PopLayer                     ComProc // 41
	Flush                        ComProc // 42
	EndDraw                      ComProc // 43
	GetPixelFormat               ComProc // 44
	SetDpi                       ComProc // 45
	GetDpi                       ComProc // 46
	GetSize                      ComProc // 47
	GetPixelSize                 ComProc // 48
	GetMaximumBitmapSize         ComProc // 49
	IsSupported                  ComProc // 50
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1RenderTarget) vtbl() *ID2D1RenderTargetVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1RenderTargetVtbl)(unsafe.Pointer(i.Vtbl))
}

// NewRenderTargetByHandle 从 xc.XDraw_GetD2dRenderTarget(hDraw) 的返回值构造渲染目标对象.
//
// GDI 渲染模式(GDI 模式没有 D2D 目标)下句柄为 0, 此时返回 nil —— 调用方据此走回落路径.
func NewRenderTargetByHandle(h uintptr) *ID2D1RenderTarget {
	if h == 0 {
		return nil
	}
	return (*ID2D1RenderTarget)(com.InterfacePtr(h))
}

// SetTransform 设置渲染目标的变换矩阵(影响之后的所有绘制).
//
// 矩阵按**指针**传递, 所以走普通 syscall 调用即可 —— 不受包注释 ABI 规则 1(按值 float 形参走 XMM)
// 的限制. 用完后记得把矩阵设回单位矩阵(IdentityMatrix()), 否则同一渲染目标上的后续绘制会被变换.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-settransform
func (i *ID2D1RenderTarget) SetTransform(transform *D2D1_MATRIX_3X2_F) {
	if i == nil {
		return
	}
	// SetTransform 返回 void, 没有 HRESULT 可查.
	i.vtbl().SetTransform.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(transform)))
}

// GetTransform 读回当前变换矩阵(出参, 同样是 void 方法).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-gettransform
func (i *ID2D1RenderTarget) GetTransform() (D2D1_MATRIX_3X2_F, error) {
	var m D2D1_MATRIX_3X2_F
	if i == nil {
		return m, errors.New("d2d: 渲染目标为空")
	}
	i.vtbl().GetTransform.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&m)))
	return m, nil
}

// CreateSolidColorBrush 创建纯色画刷(画刷属性传 NULL, 用默认值; 透明度由颜色自身的 alpha 承担).
//
// 画刷是**设备相关**资源: 只能给创建它的渲染目标使用, 用完必须 Release. 一般不必缓存 ——
// 直接调 FillGeometryColor 即可.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createsolidcolorbrush(constd2d1_color_f_constd2d1_brush_properties_id2d1solidcolorbrush)
func (i *ID2D1RenderTarget) CreateSolidColorBrush(color *D2D1_COLOR_F) (*ID2D1SolidColorBrush, error) {
	if i == nil {
		return nil, errors.New("d2d: 渲染目标为空")
	}
	var brush *ID2D1SolidColorBrush
	r, _, _ := i.vtbl().CreateSolidColorBrush.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(color)), 0, uintptr(unsafe.Pointer(&brush)))
	if r != 0 {
		return nil, errno(r)
	}
	return brush, nil
}

// FillGeometry 用画刷填充几何(图层传 NULL, 即不做图层混合).
//
// 出图前先做一次几何合法性自检: 坏几何(索引错位 / 没 Close 成功)会让 xcgui 整帧不呈现 ——
// 宁可这一笔不画, 也不能把坏几何喂进去.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillgeometry
func (i *ID2D1RenderTarget) FillGeometry(geometry *ID2D1Geometry, brush *ID2D1Brush) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if !geometry.Usable() {
		return errors.New("d2d: 几何不可用(未 Close 成功或已释放)")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	// FillGeometry 返回 void, 没有 HRESULT 可查(照直觉去查返回值的做法在这里是错的).
	i.vtbl().FillGeometry.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(geometry)),
		uintptr(unsafe.Pointer(brush)), 0)
	return nil
}

// DrawGeometry 用画刷给几何**描边**(线宽 + 可选描边样式) —— 折线的展开由 Direct2D 负责, 线帽与连接
// 样式由 style 决定(传 nil = 默认: 平头 + 斜接; 圆角语义见 RoundStrokeStyleProperties).
//
// 几何可以是任意几何(路径/矩形/圆角矩形/椭圆), 传基类视图 —— 路径几何用 Base(),
// 其它几何类型自带 Base().
//
// 为什么值得用它而不是"CPU 把描边展开成填充多边形再 FillGeometry": 实测同样画 2 万次, DrawGeometry
// 约 20ms、填充**已展开**的环约 80ms(4 倍), 而且省掉展开代码、圆角/圆帽由 D2D 保证. 代价是线宽参数
// 在 XMM3、style 压栈 —— 见 wapi/com.CallWithFloat32Stack 的说明(布局真机实测过).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawgeometry
func (i *ID2D1RenderTarget) DrawGeometry(geometry *ID2D1Geometry, brush *ID2D1Brush, strokeWidth float32, style *ID2D1StrokeStyle) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if !geometry.Usable() {
		return errors.New("d2d: 几何不可用(未 Close 成功或已释放)")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	// DrawGeometry 返回 void, 没有 HRESULT 可查.
	com.CallWithFloat32Stack(uintptr(i.vtbl().DrawGeometry),
		[4]uintptr{
			uintptr(unsafe.Pointer(i)),
			uintptr(unsafe.Pointer(geometry)),
			uintptr(unsafe.Pointer(brush)),
			0, // R9 无对应参数(按位置计数: 线宽在 XMM3, style 压栈)
		},
		[4]float32{0, 0, 0, strokeWidth},
		[4]uintptr{uintptr(unsafe.Pointer(style))})
	return nil
}

// DrawGeometryColor 建画刷 → 描边 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) DrawGeometryColor(geometry *ID2D1Geometry, color D2D1_COLOR_F, strokeWidth float32, style *ID2D1StrokeStyle) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.DrawGeometry(geometry, &brush.ID2D1Brush, strokeWidth, style)
}

// FillGeometryColor 建画刷 → 填充 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
//
// 画刷是设备相关资源, "每帧即建即放"最省心; 实测建画刷的代价远小于几何本身.
func (i *ID2D1RenderTarget) FillGeometryColor(geometry *ID2D1Geometry, color D2D1_COLOR_F) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.FillGeometry(geometry, &brush.ID2D1Brush) // ID2D1SolidColorBrush 嵌入 ID2D1Brush, 直接取地址即可
}

// ---------------- 基本图元(直线 / 矩形 / 椭圆) ----------------
//
// 这几个是 D2D 的"第一梯队"绘制口, 不必先建几何(矩形/椭圆若要反复用, 也可以建几何再绘制/填充,
// 见 CreateRectangleGeometry / CreateRoundedRectangleGeometry / CreateEllipseGeometry).
//
// 形参里的线宽(FLOAT)同样踩"按位置计数"布局(见 wapi/com.CallWithFloat32Stack):
//   - 4 个参数的(矩形/椭圆): 线宽在 XMM3, 样式在栈[0];
//   - 5 个参数的(直线): 线宽在**栈[0]**(位置 4 —— 前 4 个 XMM 槽已被位置 0~3 占掉), 样式在栈[1].

// DrawLine 画一条直线.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawline
func (i *ID2D1RenderTarget) DrawLine(p0, p1 D2D1_POINT_2F, brush *ID2D1Brush, strokeWidth float32, style *ID2D1StrokeStyle) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	// DrawLine 返回 void. 栈上的 float 占 8 字节槽, 值放低 4 字节.
	com.CallWithFloat32Stack(uintptr(i.vtbl().DrawLine),
		[4]uintptr{
			uintptr(unsafe.Pointer(i)),
			pointArg(p0),
			pointArg(p1),
			uintptr(unsafe.Pointer(brush)),
		},
		[4]float32{},
		[4]uintptr{uintptr(math.Float32bits(strokeWidth)), uintptr(unsafe.Pointer(style))})
	return nil
}

// DrawLineColor 建画刷 → 画线 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) DrawLineColor(p0, p1 D2D1_POINT_2F, color D2D1_COLOR_F, strokeWidth float32, style *ID2D1StrokeStyle) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.DrawLine(p0, p1, &brush.ID2D1Brush, strokeWidth, style)
}

// DrawRectangle 描边矩形.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawrectangle
func (i *ID2D1RenderTarget) DrawRectangle(rect *D2D1_RECT_F, brush *ID2D1Brush, strokeWidth float32, style *ID2D1StrokeStyle) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if rect == nil {
		return errors.New("d2d: 矩形为空")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	com.CallWithFloat32Stack(uintptr(i.vtbl().DrawRectangle),
		[4]uintptr{uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(rect)),
			uintptr(unsafe.Pointer(brush)), 0},
		[4]float32{0, 0, 0, strokeWidth},
		[4]uintptr{uintptr(unsafe.Pointer(style))})
	return nil
}

// DrawRectangleColor 建画刷 → 描边矩形 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) DrawRectangleColor(rect *D2D1_RECT_F, color D2D1_COLOR_F, strokeWidth float32, style *ID2D1StrokeStyle) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.DrawRectangle(rect, &brush.ID2D1Brush, strokeWidth, style)
}

// FillRectangle 填充矩形.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillrectangle
func (i *ID2D1RenderTarget) FillRectangle(rect *D2D1_RECT_F, brush *ID2D1Brush) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if rect == nil {
		return errors.New("d2d: 矩形为空")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	i.vtbl().FillRectangle.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)), 0)
	return nil
}

// FillRectangleColor 建画刷 → 填充矩形 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) FillRectangleColor(rect *D2D1_RECT_F, color D2D1_COLOR_F) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.FillRectangle(rect, &brush.ID2D1Brush)
}

// DrawEllipse 描边椭圆.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawellipse
func (i *ID2D1RenderTarget) DrawEllipse(ellipse *D2D1_ELLIPSE, brush *ID2D1Brush, strokeWidth float32, style *ID2D1StrokeStyle) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if ellipse == nil {
		return errors.New("d2d: 椭圆为空")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	com.CallWithFloat32Stack(uintptr(i.vtbl().DrawEllipse),
		[4]uintptr{uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(ellipse)),
			uintptr(unsafe.Pointer(brush)), 0},
		[4]float32{0, 0, 0, strokeWidth},
		[4]uintptr{uintptr(unsafe.Pointer(style))})
	return nil
}

// DrawEllipseColor 建画刷 → 描边椭圆 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) DrawEllipseColor(ellipse *D2D1_ELLIPSE, color D2D1_COLOR_F, strokeWidth float32, style *ID2D1StrokeStyle) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.DrawEllipse(ellipse, &brush.ID2D1Brush, strokeWidth, style)
}

// FillEllipse 填充椭圆.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillellipse
func (i *ID2D1RenderTarget) FillEllipse(ellipse *D2D1_ELLIPSE, brush *ID2D1Brush) error {
	if i == nil {
		return errors.New("d2d: 渲染目标为空")
	}
	if ellipse == nil {
		return errors.New("d2d: 椭圆为空")
	}
	if brush == nil {
		return errors.New("d2d: 画刷为空")
	}
	i.vtbl().FillEllipse.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(brush)), 0)
	return nil
}

// FillEllipseColor 建画刷 → 填充椭圆 → 立刻释放画刷(便利方法, 不是 SDK 原生 API).
func (i *ID2D1RenderTarget) FillEllipseColor(ellipse *D2D1_ELLIPSE, color D2D1_COLOR_F) error {
	brush, err := i.CreateSolidColorBrush(&color)
	if err != nil {
		return err
	}
	defer brush.Release()
	return i.FillEllipse(ellipse, &brush.ID2D1Brush)
}
