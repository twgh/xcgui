package d2d

import (
	"errors"
	"math"
	"syscall"
	"unsafe"
)

// errGeometryNil 几何为空(句柄为 0 或已释放).
var errGeometryNil = errors.New("d2d: 几何为空")

// errno 把 HRESULT 包成 error.
func errno(r uintptr) error { return syscall.Errno(r) }

// pointArg 把 D2D1_POINT_2F(按值传递的两个 float32)打包成一个 64 位整数.
//
// 之所以要打包: syscall.SyscallN 只能传通用寄存器, 浮点形参走 XMM 寄存器传不进去, 而"按值传递
// 的 8 字节结构体"在 x64 调用约定里正好走一个通用寄存器 —— 打包成一个 uintptr 就能原样传进去.
func pointArg(p D2D1_POINT_2F) uintptr {
	return uintptr(math.Float32bits(p.X)) | uintptr(math.Float32bits(p.Y))<<32
}

// ---------------- nil 安全 ----------------
//
// 值嵌入的代价: 对 nil 的外层指针调用**被提升的基类方法**时, 编译器要先解引用到嵌入字段
// (相当于 `(&nil.ID2D1Brush.ID2D1Resource.IUnknown)`), 这会 panic —— 哪怕基类方法自己判了空.
// 于是这里给每个接口补一个"只做判空"的 Release 覆盖: Release 是唯一可能被 defer/对空对象调用的
// 方法(构造失败拿到 nil 时很常见). AddRef/QueryInterface 不覆盖 —— 对 nil 调它们本来就是 bug,
// 直接崩反而好定位.

// Release 减少引用计数(nil 安全).
func (i *ID2D1Resource) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.IUnknown.Release()
}

// Release 减少引用计数(nil 安全).
func (i *ID2D1Factory) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.IUnknown.Release()
}

// Release 减少引用计数(nil 安全).
func (i *ID2D1RenderTarget) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Resource.Release()
}

// Release 释放 DC 渲染目标(nil 安全).
func (i *ID2D1DCRenderTarget) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1RenderTarget.Release()
}

// Release 减少引用计数(nil 安全).
func (i *ID2D1Geometry) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Resource.Release()
}

// Release 释放几何(nil 安全).
func (i *ID2D1PathGeometry) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Geometry.Release()
}

// Release 释放矩形几何(nil 安全).
func (i *ID2D1RectangleGeometry) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Geometry.Release()
}

// Release 释放圆角矩形几何(nil 安全).
func (i *ID2D1RoundedRectangleGeometry) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Geometry.Release()
}

// Release 释放椭圆几何(nil 安全).
func (i *ID2D1EllipseGeometry) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Geometry.Release()
}

// Release 减少引用计数(nil 安全).
func (i *ID2D1Brush) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Resource.Release()
}

// Release 释放画刷(nil 安全).
func (i *ID2D1SolidColorBrush) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Brush.Release()
}

// Release 释放描边样式(nil 安全).
func (i *ID2D1StrokeStyle) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1Resource.Release()
}

// Release 释放几何写入器(nil 安全; Close 成功时已自动释放).
func (i *ID2D1SimplifiedGeometrySink) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.IUnknown.Release()
}

// Release 释放几何写入器(nil 安全; Close 成功时已自动释放).
func (i *ID2D1GeometrySink) Release() uintptr {
	if i == nil {
		return 0
	}
	return i.ID2D1SimplifiedGeometrySink.Release()
}

const ptrSize = int(unsafe.Sizeof(uintptr(0)))

// 布局不变量(编译期断言): 这些接口结构体都是"值嵌入单字段基类"串起来的, 必须只有一个指针大小 ——
// 那样它的地址才是 COM 对象的地址、它的首字段才是真实虚表指针, 各接口的 vtbl() 重解释才成立.
//
// 谁要是往嵌入链上加了字段, 这里会**编译失败**(数组索引非 0 越界), 而不是等运行期画出满屏乱码.
// 断言写法: [1]struct{}{}[size-ptrSize] —— size 必须是 ptrSize, 索引 0 才合法.
var (
	_ = [1]struct{}{}[int(unsafe.Sizeof(IUnknown{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1Resource{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1Factory{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1RenderTarget{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1DCRenderTarget{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1Geometry{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1PathGeometry{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1RectangleGeometry{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1RoundedRectangleGeometry{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1EllipseGeometry{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1Brush{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1SolidColorBrush{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1StrokeStyle{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1SimplifiedGeometrySink{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(ID2D1GeometrySink{}))-ptrSize]
	_ = [1]struct{}{}[int(unsafe.Sizeof(IUnknownVtbl{}))-int(3*ptrSize)]
)
