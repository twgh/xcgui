package gdi

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

// 本文件封装 gdi32.dll 的常用函数, 以及"32bpp 位图 + 内存 DC"这一常用的离屏渲染套路.
//
// DIB 的用途: 给离屏渲染一块能直接读写的像素内存 —— CreateDIBSection 返回的就是**指向像素的
// 裸指针**, 不用再走 GetDIBits 拷一遍(例如 Direct2D 的 DC 渲染目标就要求这样一块内存).
// 纯 Win32, 与 xcgui.dll 无关.
//
// 约定: top-down 的 32bpp BGRA(DIB 的第一行就是图像的第一行, 与 Go 里 image.RGBA 的行序一致;
// 用 bottom-up 的话每次读像素都要自己翻转行序).

var (
	// Gdi32 是 Windows 操作系统中的核心图形接口，负责所有图形和字体的屏幕显示与打印输出。
	Gdi32 = syscall.NewLazyDLL("gdi32.dll")

	// Functions.
	procCreateRoundRectRgn = Gdi32.NewProc("CreateRoundRectRgn")
	procCreateCompatibleDC = Gdi32.NewProc("CreateCompatibleDC")
	procCreateDIBSection   = Gdi32.NewProc("CreateDIBSection")
	procSelectObject       = Gdi32.NewProc("SelectObject")
	procDeleteObject       = Gdi32.NewProc("DeleteObject")
	procDeleteDC           = Gdi32.NewProc("DeleteDC")
)

// CreateRoundRectRgn 创建具有圆角的矩形区域, 返回区域句柄。
//
// left, top, right, bottom: 指定矩形的边界坐标。
//
// width, height: 指定椭圆的宽度和高度，用于创建圆角。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(left, top, right, bottom, width, height int32) uintptr {
	ret, _, _ := procCreateRoundRectRgn.Call(
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		uintptr(width),
		uintptr(height),
	)
	return ret
}

// CreateCompatibleDC 创建一个与指定设备兼容的内存设备上下文(DC)。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
//
// hdc: 现有 DC 的句柄。 传 0 表示与应用程序的当前屏幕兼容。
//
// 返回值: 成功时返回内存 DC 的句柄, 失败时返回 0. 用完必须 DeleteDC。
func CreateCompatibleDC(hdc uintptr) uintptr {
	ret, _, _ := procCreateCompatibleDC.Call(hdc)
	return ret
}

// CreateDIBSection 创建 DIB(设备无关位图), 返回位图句柄与像素内存指针。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-createdibsection
//
// hdc: DC 的句柄。 传 0 表示只创建 DIB section, 不创建任何兼容 DC。
//
// pbmi: 指向 BITMAPINFOHEADER 的指针, 描述位图的尺寸 / 位深 / 压缩方式。
//
// usage: 调色板类型, 取值之一: gdi.DIB_RGB_COLORS。
//
// 返回值:
//   - hbm: 位图句柄, 失败时为 0。
//   - bits: 指向位图像素数据的指针, 失败时为 nil。 这块内存由系统分配, 位图释放时一并释放。
//
// 本封装只覆盖"像素内存由系统分配"这一种用法: 原型的 lpBits / hSection / dwOffset 三个参数固定
// 传 0(不传外部缓冲、不映射文件)。
func CreateDIBSection(hdc uintptr, pbmi *BITMAPINFOHEADER, usage uint32) (hbm uintptr, bits unsafe.Pointer) {
	ret, _, _ := procCreateDIBSection.Call(
		hdc,
		uintptr(unsafe.Pointer(pbmi)),
		uintptr(usage),
		uintptr(unsafe.Pointer(&bits)),
		0, // hSection
		0, // dwOffset
	)
	return ret, bits
}

// SelectObject 将对象选入指定的 DC, 返回 DC 中原本选着的同类型对象。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-selectobject
//
// hdc: DC 的句柄。
//
// h: 要选入的对象的句柄。 该对象必须与 DC 类型匹配(位图只能选入内存 DC)。
//
// 返回值:
//   - 成功时返回被替换下来的对象句柄, 失败时返回 0。 调用方必须保存它, 用完再选回去,
//     否则被替换的对象释放不掉。
//   - 传入区域(HRGN)时返回值另有含义: 0 表示出错, 非 0 是区域类型(见 MSDN)。
func SelectObject(hdc, h uintptr) uintptr {
	ret, _, _ := procSelectObject.Call(hdc, h)
	return ret
}

// DeleteObject 删除逻辑笔 / 画笔 / 字体 / 位图 / 区域或调色板, 释放其占用的系统资源。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-deleteobject
//
// h: 对象的句柄。
//
// 返回值: 成功返回 true; 句柄无效, 或对象仍被选在某个 DC 里时返回 false(所以删位图前先把它选出去)。
func DeleteObject(h uintptr) bool {
	ret, _, _ := procDeleteObject.Call(h)
	return ret != 0
}

// DeleteDC 删除指定的设备上下文(DC)。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-deletedc
//
// hdc: DC 的句柄。
//
// 返回值: 成功返回 true, 失败返回 false。
//
// 注意: DC 里选着的对象不会随 DC 一起释放, 且删除 DC 前必须先把它们选回原位(见 SelectObject)。
func DeleteDC(hdc uintptr) bool {
	ret, _, _ := procDeleteDC.Call(hdc)
	return ret != 0
}

const (
	// BI_RGB 是 BITMAPINFOHEADER.Compression 的取值: 不压缩(32bpp 下就是"每像素 4 字节, 蓝绿红 alpha")。
	BI_RGB = 0
	// DIB_RGB_COLORS 是 CreateDIBSection 的 usage 参数取值: 调色板是 RGB 值(32bpp 无调色板, 与 0 等价)。
	DIB_RGB_COLORS = 0
)

// maxDIBDim 是 DIB 边长的上限(常见的渲染 API, 如 D2D, 在桌面上的上限是 16384)。
const maxDIBDim = 16384

// BITMAPINFOHEADER 描述 DIB 的尺寸 / 位深 / 压缩方式, 共 40 字节, 字段顺序即 C 布局。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	Size          uint32 // 结构大小(字节), 固定为 sizeof(BITMAPINFOHEADER)。
	Width         int32  // 位图宽度(像素)。
	Height        int32  // 位图高度(像素)。 负数 = top-down(第一行在内存最前), 正数 = bottom-up。
	Planes        uint16 // 颜色平面数, 固定为 1。
	BitCount      uint16 // 每像素位数, 32 即每像素 4 字节(BGRA)。
	Compression   uint32 // 压缩方式, 不压缩用 gdi.BI_RGB。
	SizeImage     uint32 // 图像字节数, BI_RGB 下可传 0。
	XPelsPerMeter int32  // 水平分辨率(像素/米), 仅供显示。
	YPelsPerMeter int32  // 垂直分辨率(像素/米), 仅供显示。
	ClrUsed       uint32 // 调色板实际用到的颜色数, 0 表示用满 2^BitCount 种。
	ClrImportant  uint32 // 显示所必需的颜色数, 0 表示全都必需。
}

// DIB 是一块 32bpp 的 DIB section + 一个已选好它的内存 DC.
//
// 生命周期: NewDIB 建, **必须 Close**(内部按序恢复 DC 原对象 → 删位图 → 删 DC)。
//
// 建好之后像素内存就可以直接读写(见 Pixels), 也可以把 HDC 交给离屏渲染目标(例如
// d2d.ID2D1DCRenderTarget)去画.
type DIB struct {
	dc   uintptr        // 内存 DC(CreateCompatibleDC)
	hbm  uintptr        // DIB section 位图句柄(CreateDIBSection)
	old  uintptr        // DC 里原本选着的对象(Close 时还回去, 否则位图删不掉)
	bits unsafe.Pointer // 像素内存(宽×高×4 字节, BGRA, top-down)
	w, h int32
}

// NewDIB 创建宽 w 高 h 的 32bpp 位图及配套内存 DC(像素初始内容**不保证是零**, 要干净画面
// 请先自己清一遍, 比如给 D2D 使用时用 ID2D1RenderTarget.Clear).
func NewDIB(w, h int32) (*DIB, error) {
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("DIB 尺寸必须为正, 得到 %dx%d", w, h)
	}
	// 上限同时兜住两件事: 渲染目标的尺寸上限(见 maxDIBDim), 以及 w*h*4 在 int32 里不溢出.
	if w > maxDIBDim || h > maxDIBDim {
		return nil, fmt.Errorf("DIB 尺寸超过 %d×%d 上限, 得到 %dx%d", maxDIBDim, maxDIBDim, w, h)
	}
	for _, p := range []*syscall.LazyProc{procCreateCompatibleDC, procCreateDIBSection} {
		if err := p.Find(); err != nil {
			return nil, fmt.Errorf("加载 gdi32.dll 失败: %w", err)
		}
	}

	dc := CreateCompatibleDC(0) // 0 = 与屏幕兼容的内存 DC
	if dc == 0 {
		return nil, errors.New("CreateCompatibleDC 失败")
	}
	d := &DIB{dc: dc, w: w, h: h}

	bi := BITMAPINFOHEADER{
		Size:        uint32(unsafe.Sizeof(BITMAPINFOHEADER{})),
		Width:       w,
		Height:      -h, // 负数 = top-down(渲染目标与 Go 的 image 都按这个方向读)
		Planes:      1,
		BitCount:    32,
		Compression: BI_RGB,
		SizeImage:   uint32(w * h * 4),
	}
	hbm, bits := CreateDIBSection(dc, &bi, DIB_RGB_COLORS)
	if hbm == 0 || bits == nil {
		d.Close()
		return nil, errors.New("CreateDIBSection 失败")
	}
	d.hbm, d.bits = hbm, bits

	old := SelectObject(dc, hbm)
	if old == 0 {
		d.Close()
		return nil, errors.New("SelectObject(DIB) 失败")
	}
	d.old = old
	return d, nil
}

// HDC 返回内存 DC 的句柄.
func (d *DIB) HDC() uintptr {
	if d == nil {
		return 0
	}
	return d.dc
}

// Width 返回位图宽度(像素).
func (d *DIB) Width() int32 {
	if d == nil {
		return 0
	}
	return d.w
}

// Height 返回位图高度(像素).
func (d *DIB) Height() int32 {
	if d == nil {
		return 0
	}
	return d.h
}

// Pixels 返回像素内存的视图: **宽×高×4 字节, BGRA, top-down**(第一行在上).
//
// 它与 HDC 是同一块内存 —— 把 HDC 交给渲染目标画完之后, 直接读它就能拿到结果(不需要
// GetDIBits 再拷一道).
func (d *DIB) Pixels() []byte {
	if d == nil || d.bits == nil {
		return nil
	}
	return unsafe.Slice((*byte)(d.bits), int(d.w)*int(d.h)*4)
}

// Close 释放位图与 DC(nil 安全, 可重复调用).
//
// 顺序不能错: 先把原对象选回 DC(位图还选着的话 DeleteObject 会失败), 再删位图, 最后删 DC.
func (d *DIB) Close() {
	if d == nil {
		return
	}
	if d.dc != 0 {
		if d.old != 0 {
			SelectObject(d.dc, d.old)
			d.old = 0
		}
		if d.hbm != 0 {
			DeleteObject(d.hbm)
			d.hbm = 0
		}
		DeleteDC(d.dc)
		d.dc = 0
	}
	d.bits = nil
}
