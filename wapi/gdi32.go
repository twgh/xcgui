package wapi

import "syscall"

var (
	// Library.
	Gdi32 = syscall.NewLazyDLL("gdi32.dll")

	// Functions.
	procCreateRoundRectRgn = Gdi32.NewProc("CreateRoundRectRgn")
)

// CreateRoundRectRgn 创建具有圆角的矩形区域, 返回区域句柄。
//
// left, top, right, bottom: 指定矩形的边界坐标。
//
// width, height: 指定椭圆的宽度和高度，用于创建圆角。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
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
