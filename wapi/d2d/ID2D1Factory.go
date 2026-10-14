package d2d

import (
	"errors"
	"unsafe"

	"github.com/twgh/xcgui/wapi/com"
)

// ID2D1Factory 创建 Direct2D 资源(设备无关资源都归它).
//
// 它由 xcgui 持有(见 xc.XC_GetD2dFactory), 这里只是借用, 不要 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1factory
type ID2D1Factory struct {
	IUnknown
}

type ID2D1FactoryVtbl struct {
	IUnknownVtbl
	ReloadSystemMetrics            ComProc // 3
	GetDesktopDpi                  ComProc // 4
	CreateRectangleGeometry        ComProc // 5
	CreateRoundedRectangleGeometry ComProc // 6
	CreateEllipseGeometry          ComProc // 7
	CreateGeometryGroup            ComProc // 8
	CreateTransformedGeometry      ComProc // 9
	CreatePathGeometry             ComProc // 10
	CreateStrokeStyle              ComProc // 11
	CreateDrawingStateBlock        ComProc // 12
	CreateWicBitmapRenderTarget    ComProc // 13
	CreateHwndRenderTarget         ComProc // 14
	CreateDxgiSurfaceRenderTarget  ComProc // 15
	CreateDCRenderTarget           ComProc // 16
}

// vtbl 取本接口的虚表.
//
// 注意是 `i.Vtbl`(对象首字段存的**虚表指针**)而不是 `unsafe.Pointer(i)`: COM 对象本身不是虚表,
// 把对象地址当虚表用会读到对象内存偏移 8/16/... 处的垃圾, 表现为跳到非法地址(实测 PC=0x989680).
func (i *ID2D1Factory) vtbl() *ID2D1FactoryVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1FactoryVtbl)(unsafe.Pointer(i.Vtbl))
}

// NewFactoryByHandle 从 xc.XC_GetD2dFactory()(或 app.GetD2dFactory())的返回值构造工厂对象.
//
// 未开启 D2D(GDI 渲染)时句柄为 0, 此时返回 nil —— 调用方据此走回落路径.
func NewFactoryByHandle(h uintptr) *ID2D1Factory {
	if h == 0 {
		return nil
	}
	return (*ID2D1Factory)(com.InterfacePtr(h))
}

// CreatePathGeometry 创建一条空路径几何.
//
// 几何是设备无关资源(归属工厂), 可以长期缓存复用; 但它由本包创建, 用完必须 Release.
func (i *ID2D1Factory) CreatePathGeometry() (*ID2D1PathGeometry, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	var geometry *ID2D1PathGeometry
	r, _, _ := i.vtbl().CreatePathGeometry.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&geometry)))
	if r != 0 {
		return nil, errno(r)
	}
	return geometry, nil
}

// CreateRectangleGeometry 创建矩形几何.
//
// 矩形/圆角矩形/椭圆要反复绘制或填充时, 建几何比每次调 DrawRectangle/FillRectangle 更省
// (几何可缓存复用, 见 GeometryCache).
func (i *ID2D1Factory) CreateRectangleGeometry(rect *D2D1_RECT_F) (*ID2D1RectangleGeometry, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	if rect == nil {
		return nil, errors.New("d2d: 矩形为空")
	}
	var geometry *ID2D1RectangleGeometry
	r, _, _ := i.vtbl().CreateRectangleGeometry.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(&geometry)))
	if r != 0 {
		return nil, errno(r)
	}
	return geometry, nil
}

// CreateRoundedRectangleGeometry 创建圆角矩形几何.
func (i *ID2D1Factory) CreateRoundedRectangleGeometry(rr *D2D1_ROUNDED_RECT) (*ID2D1RoundedRectangleGeometry, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	if rr == nil {
		return nil, errors.New("d2d: 圆角矩形为空")
	}
	var geometry *ID2D1RoundedRectangleGeometry
	r, _, _ := i.vtbl().CreateRoundedRectangleGeometry.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(rr)), uintptr(unsafe.Pointer(&geometry)))
	if r != 0 {
		return nil, errno(r)
	}
	return geometry, nil
}

// CreateEllipseGeometry 创建椭圆几何.
func (i *ID2D1Factory) CreateEllipseGeometry(ellipse *D2D1_ELLIPSE) (*ID2D1EllipseGeometry, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	if ellipse == nil {
		return nil, errors.New("d2d: 椭圆为空")
	}
	var geometry *ID2D1EllipseGeometry
	r, _, _ := i.vtbl().CreateEllipseGeometry.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(ellipse)), uintptr(unsafe.Pointer(&geometry)))
	if r != 0 {
		return nil, errno(r)
	}
	return geometry, nil
}

// CreateStrokeStyle 创建描边样式(线帽 / 连接 / 虚线). 自定义虚线段(dashes)传 NULL, 即用内置虚线样式.
//
// 描边样式是**设备无关**资源(归属工厂), 建议建一次缓存起来(见 RoundStrokeStyleProperties: 圆角连接 +
// 圆头的 SVG 语义).
func (i *ID2D1Factory) CreateStrokeStyle(props *D2D1_STROKE_STYLE_PROPERTIES) (*ID2D1StrokeStyle, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	var style *ID2D1StrokeStyle
	r, _, _ := i.vtbl().CreateStrokeStyle.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(props)), 0, 0, uintptr(unsafe.Pointer(&style)))
	if r != 0 {
		return nil, errno(r)
	}
	return style, nil
}

// CreateDCRenderTarget 创建"画到 GDI 设备上下文"的渲染目标 —— 离屏渲染成一张可读像素的图片
// 就走它(配合 gdi.NewDIB 的 32bpp DIB, 见 RenderToRGBA 那种用法).
//
// props 一般用 NewDCRenderTargetProperties(); 建出来的渲染目标是**设备相关**资源, 用完必须 Release.
func (i *ID2D1Factory) CreateDCRenderTarget(props *D2D1_RENDER_TARGET_PROPERTIES) (*ID2D1DCRenderTarget, error) {
	if i == nil {
		return nil, errors.New("d2d: 工厂为空(未开启 D2D?)")
	}
	if props == nil {
		return nil, errors.New("d2d: 渲染目标属性为空")
	}
	var rt *ID2D1DCRenderTarget
	r, _, _ := i.vtbl().CreateDCRenderTarget.Call(uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(props)), uintptr(unsafe.Pointer(&rt)))
	if r != 0 {
		return nil, errno(r)
	}
	return rt, nil
}

// NewGeometry 一步构造一条可用几何(便利方法, 不是 SDK 原生 API).
//
// figures 是若干**闭合**折线(扁平 x,y 点对, 每条至少 3 个点; 点数不足的会被跳过).
// 填充模式固定 D2D1_FILL_MODE_WINDING —— 大量互相重叠的多边形(例如把描边展开成的一条条
// 轮廓)必须取并集, 用默认的 ALTERNATE 会让重叠处互相镂空.
//
// 内部按 CreatePathGeometry → Open → BeginFigure/AddLines/EndFigure → Close 走完整流程,
// 并做合法性自检; 任一步失败返回错误与 nil(调用方应走回落路径, 而不是硬画).
func (i *ID2D1Factory) NewGeometry(figures [][]float32) (*ID2D1PathGeometry, error) {
	closed := make([]GeometryFigure, 0, len(figures))
	for _, figure := range figures {
		if len(figure) < 6 { // 闭合填充至少 3 个点, 否则围不出面积
			continue
		}
		closed = append(closed, GeometryFigure{Points: figure, Closed: true})
	}
	if len(closed) == 0 {
		return nil, errors.New("d2d: 没有一个 figure 至少有 3 个点")
	}
	return i.NewGeometryFigures(closed)
}

// GeometryFigure 几何里的一条子路径: 扁平 x,y 点对 + 是否闭合.
//
// **描边用途(如 lucide 图标的折线)必须 Closed=false** —— 闭合会多画一段收尾线(三角形折线变回形针);
// 填充用途要 Closed=true. 至少 2 个点才有意义(单点既围不出面积也描不出线).
type GeometryFigure struct {
	Points []float32
	Closed bool
}

// NewGeometryFigures 按给定子路径构造几何(每条可分别开/闭), 填充模式 WINDING, 末尾做合法性自检.
//
// 与 NewGeometry 的关系: 后者是"全部闭合 + 至少 3 个点"的便捷形式, 内部就是转成本函数再建.
func (i *ID2D1Factory) NewGeometryFigures(figures []GeometryFigure) (*ID2D1PathGeometry, error) {
	geometry, err := i.CreatePathGeometry()
	if err != nil {
		return nil, err
	}
	sink, err := geometry.Open()
	if err != nil {
		geometry.Release()
		return nil, err
	}
	sink.SetFillMode(D2D1_FILL_MODE_WINDING)
	n := 0
	for _, figure := range figures {
		if len(figure.Points) < 4 { // 少于 2 个点
			continue
		}
		end := D2D1_FIGURE_END_OPEN
		if figure.Closed {
			end = D2D1_FIGURE_END_CLOSED
		}
		sink.BeginFigure(D2D1_POINT_2F{X: figure.Points[0], Y: figure.Points[1]}, D2D1_FIGURE_BEGIN_FILLED)
		sink.AddLines(figure.Points[2:])
		sink.EndFigure(end)
		n++
	}
	if n == 0 {
		sink.Release()
		geometry.Release()
		return nil, errors.New("d2d: 没有一个子路径至少有 2 个点")
	}
	if err := sink.Close(); err != nil { // Close 成功后几何写入器自动释放
		geometry.Release()
		return nil, err
	}
	if !geometry.Usable() {
		geometry.Release()
		return nil, errors.New("d2d: 几何自检未通过(GetBounds 失败)")
	}
	return geometry, nil
}
