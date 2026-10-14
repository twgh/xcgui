package d2d

import (
	"errors"
	"unsafe"

	"github.com/twgh/xcgui/wapi/types"
)

// ID2D1DCRenderTarget 是画到 GDI 设备上下文(DC)上的渲染目标 —— 离屏渲染就靠它:
// 把 32bpp DIB 的 HDC 绑上来, 画完之后直接读 DIB 内存里的像素.
//
// 它比 ID2D1HwndRenderTarget 少掉"窗口/交换链"这一层: 不需要窗口、不需要 D2D 设备,
// 因此在 **GDI 渲染模式**(app.New(false))甚至"还没建窗口"时都能出图.
//
// 由本包创建(见 ID2D1Factory.CreateDCRenderTarget), 用完必须 Release.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1dcrendertarget
type ID2D1DCRenderTarget struct {
	ID2D1RenderTarget
}

type ID2D1DCRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	BindDC ComProc // 57
}

// vtbl 取本接口的虚表(见 ID2D1Factory.vtbl 的说明: 必须取 i.Vtbl, 不能把对象地址当虚表).
func (i *ID2D1DCRenderTarget) vtbl() *ID2D1DCRenderTargetVtbl {
	if i == nil {
		return nil
	}
	return (*ID2D1DCRenderTargetVtbl)(unsafe.Pointer(i.Vtbl))
}

// BindDC 把渲染目标绑到一个 GDI 设备上下文(dc), 绘制的落点由 rect 指定(dc 内的子区域,
// 通常就是整个 DIB: {0, 0, 宽, 高}).
//
// 注意顺序: **先 BindDC 再 BeginDraw**; 同一渲染目标可以换 DC 反复绑定(重新绑定前要保证
// 上一次的绘制已经 EndDraw 提交).
func (i *ID2D1DCRenderTarget) BindDC(dc uintptr, rect *types.RECT) error {
	if i == nil {
		return errors.New("d2d: DC 渲染目标为空")
	}
	if dc == 0 {
		return errors.New("d2d: 设备上下文为空")
	}
	if rect == nil {
		return errors.New("d2d: 绑定区域为空")
	}
	r, _, _ := i.vtbl().BindDC.Call(uintptr(unsafe.Pointer(i)), dc, uintptr(unsafe.Pointer(rect)))
	if r != 0 {
		return errno(r)
	}
	return nil
}
