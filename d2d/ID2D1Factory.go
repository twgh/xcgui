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
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createpathgeometry(id2d1pathgeometry)
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

// NewGeometry 一步构造一条可用几何(便利方法, 不是 SDK 原生 API).
//
// figures 是若干**闭合**折线(扁平 x,y 点对, 每条至少 3 个点; 点数不足的会被跳过).
// 填充模式固定 D2D1_FILL_MODE_WINDING —— 大量互相重叠的多边形(例如把描边展开成的一条条
// 轮廓)必须取并集, 用默认的 ALTERNATE 会让重叠处互相镂空.
//
// 内部按 CreatePathGeometry → Open → BeginFigure/AddLines/EndFigure → Close 走完整流程,
// 并做合法性自检; 任一步失败返回错误与 nil(调用方应走回落路径, 而不是硬画).
func (i *ID2D1Factory) NewGeometry(figures [][]float32) (*ID2D1PathGeometry, error) {
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
		if len(figure) < 6 { // 少于 3 个点围不出面积
			continue
		}
		sink.BeginFigure(D2D1_POINT_2F{X: figure[0], Y: figure[1]}, D2D1_FIGURE_BEGIN_FILLED)
		sink.AddLines(figure[2:])
		sink.EndFigure(D2D1_FIGURE_END_CLOSED)
		n++
	}
	if n == 0 {
		sink.Release()
		geometry.Release()
		return nil, errors.New("d2d: 没有一个 figure 至少有 3 个点")
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
