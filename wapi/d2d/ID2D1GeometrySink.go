package d2d

import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/twgh/xcgui/wapi/com"
)

// ID2D1GeometrySink 是几何写入器(在 SimplifiedGeometrySink 之上再追加曲线与弧).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1geometrysink
type ID2D1GeometrySink struct {
	ID2D1SimplifiedGeometrySink
}

type ID2D1GeometrySinkVtbl struct {
	ID2D1SimplifiedGeometrySinkVtbl
	AddLine             ComProc // 10
	AddBezier           ComProc // 11
	AddQuadraticBezier  ComProc // 12
	AddQuadraticBeziers ComProc // 13
	AddArc              ComProc // 14
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1GeometrySink) vtbl() *ID2D1GeometrySinkVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1GeometrySinkVtbl)(unsafe.Pointer(i.Vtbl))
}

// SetFillMode 设置填充模式(void 方法).
func (i *ID2D1GeometrySink) SetFillMode(fillMode D2D1_FILL_MODE) {
	if i == nil {
		return
	}
	i.vtbl().SetFillMode.Call(uintptr(unsafe.Pointer(i)), uintptr(fillMode))
}

// BeginFigure 开始一条子路径(void 方法).
//
// 点按值传递(8 字节结构体): 打包成 64 位整数后用 com.U64 传, 386 下它会自动拆成两个槽.
func (i *ID2D1GeometrySink) BeginFigure(point D2D1_POINT_2F, figureBegin D2D1_FIGURE_BEGIN) {
	if i == nil {
		return
	}
	com.Call(uintptr(i.vtbl().BeginFigure),
		com.U(uintptr(unsafe.Pointer(i))),
		com.U64(pointArg(point)),
		com.U(uintptr(figureBegin)))
}

// AddLines 追加折线顶点(void 方法).
//
// points 是扁平 x,y 点对, **不含**起点(BeginFigure 已给出).
func (i *ID2D1GeometrySink) AddLines(points []float32) {
	if i == nil || len(points) < 2 {
		return
	}
	i.vtbl().AddLines.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&points[0])),
		uintptr(len(points)/2))
	runtime.KeepAlive(points) // points 是调用期间才读的缓冲区
}

// EndFigure 结束当前子路径(void 方法).
func (i *ID2D1GeometrySink) EndFigure(figureEnd D2D1_FIGURE_END) {
	if i == nil {
		return
	}
	i.vtbl().EndFigure.Call(uintptr(unsafe.Pointer(i)), uintptr(figureEnd))
}

// Close 提交几何 —— 唯一返回 HRESULT 的写入方法, 必须调用, 否则几何不可用(见 Usable).
//
// 成功时内部会释放几何写入器(它是一次性的), 调用方无需再调 Release; 失败时返回错误, 调用方
// 应放弃这条几何.
func (i *ID2D1GeometrySink) Close() error {
	if i == nil {
		return errors.New("d2d: 几何写入器为空")
	}
	defer i.Release()
	r, _, _ := i.vtbl().Close.Call(uintptr(unsafe.Pointer(i)))
	if r != 0 {
		return errno(r)
	}
	return nil
}
