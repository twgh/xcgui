package xc

import (
	"math"
	"os"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// PathExists 判断文件或文件夹是否存在. 如果出错, 则不确定是否存在.
//
// path: 文件或文件夹路径.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 如果返回的错误为nil,说明文件或文件夹存在
		return true, nil
	}
	if os.IsNotExist(err) { // 如果返回的错误类型使用 os.IsNotExist() 判断为true, 说明文件或文件夹不存在
		return false, nil
	}
	return false, err // 如果返回的错误为其它类型, 则不确定是否在存在
}

// PathExists2 判断文件或文件夹是否存在. 不考虑极端情况.
//
// path: 文件或文件夹路径.
func PathExists2(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ParseRGBA 将 RGBA 生成的颜色还原回 r, g, b, a.
func ParseRGBA(color uint32) (r, g, b, a byte) {
	r = byte(color & 0xFF)
	g = byte((color >> 8) & 0xFF)
	b = byte((color >> 16) & 0xFF)
	a = byte((color >> 24) & 0xFF)
	return
}

// ParseRGB 将 RGB 生成的颜色还原回 r, g, b.
func ParseRGB(color uint32) (r, g, b byte) {
	r = byte(color & 0xFF)
	g = byte((color >> 8) & 0xFF)
	b = byte((color >> 16) & 0xFF)
	return
}

// RGBA 根据 r, g, b, a 组合成炫彩使用的颜色.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
//
// a: 透明度.
func RGBA(r, g, b, a byte) uint32 {
	return uint32(a)<<24 | uint32(r) | uint32(g)<<8 | uint32(b)<<16
}

// RGBA2 根据 RGB, a组合成炫彩使用的颜色.
//
// rgb: RGB 颜色.
//
// a: 透明度.
func RGBA2(rgb uint32, a byte) uint32 {
	return (uint32(rgb) & 16777215) | uint32(a)<<24
}

// RGB 根据 r, g, b 组合成 RGB 颜色.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
func RGB(r, g, b byte) uint32 {
	return uint32(r) | uint32(g)<<8 | uint32(b)<<16
}

// RGB2RGBA 将 RGB 颜色转换到炫彩使用的颜色.
//
// rgb: RGB 颜色.
//
// a: 透明度.
func RGB2RGBA(rgb uint32, a byte) uint32 {
	r := byte(rgb & 255)
	g := byte((rgb >> 8) & 255)
	b := byte((rgb >> 16) & 255)
	return RGBA(r, g, b, a)
}

// HexRGB2RGB 将十六进制 RGB 颜色转换到十进制 RGB 颜色.
//
// str: 十六进制 RGB 颜色, 开头有没有#都可以.
func HexRGB2RGB(str string) uint32 {
	if len(str) > 0 && str[0] == '#' {
		str = str[1:]
	}
	if len(str) != 6 {
		return 0 // 返回一个默认值或错误
	}
	r := hexToByte(str[0:2])
	g := hexToByte(str[2:4])
	b := hexToByte(str[4:6])
	return RGB(r, g, b)
}

// HexRGB2RGBA 将十六进制 RGB 颜色转换到十进制炫彩使用的颜色.
//
// str: 十六进制 RGB 颜色, 开头有没有#都可以.
//
// a: 透明度.
func HexRGB2RGBA(str string, a byte) uint32 {
	return RGBA2(HexRGB2RGB(str), a)
}

func hexToByte(s string) byte {
	if len(s) != 2 {
		return 0 // 返回一个默认值或错误
	}
	nib1 := hexToNibble(s[0])
	nib2 := hexToNibble(s[1])
	return (nib1 << 4) | nib2
}

func hexToNibble(c byte) byte {
	if c >= '0' && c <= '9' {
		return c - '0'
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else {
		return 0 // 返回一个默认值或错误
	}
}

// Itoa 将 int32 转换到 string.
func Itoa(i32 int32) string {
	return common.Itoa(i32)
}

// Atoi 将 string 转换到 int32.
func Atoi(s string) int32 {
	return common.Atoi(s)
}

// Ftoa 将 float32 转换到 string.
func Ftoa(f float32) string {
	return common.Ftoa(f)
}

// Atof 将 string 转换到 float32.
func Atof(s string) float32 {
	return common.Atof(s)
}

// DpiConv 将 int32 类型的整数根据窗口 dpi 进行换算.
//   - 开启自动 DPI 后, 你可能感觉到一些函数获取的坐标不对了, 因为你在用高分辨率屏幕, 然后系统里会有个缩放的设置, 可能是 150%, 这时获取到的坐标应该乘以 1.5 才是屏幕上显示的真实的坐标.
//
// dpi: 使用窗口.GetDPI() 函数获取.
//
// i: int32 类型的整数.
func DpiConv(dpi int32, i int32) int32 {
	ret := float32(dpi*i) / 96.0
	return int32(ret)
}

// DpiConvRound 将 int32 类型的整数根据窗口 dpi 进行换算, 计算结果会四舍五入.
//   - 开启自动 DPI 后, 你可能感觉到一些函数获取的坐标不对了, 因为你在用高分辨率屏幕, 然后系统里会有个缩放的设置, 可能是 150%, 这时获取到的坐标应该乘以 1.5 才是屏幕上显示的真实的坐标.
//
// dpi: 使用窗口.GetDPI() 函数获取.
//
// i: int32 类型的整数.
func DpiConvRound(dpi int32, i int32) int32 {
	ret := float64(dpi) * float64(i) / 96.0
	return int32(math.Round(ret))
}

// OffsetRect 传入 RECT, 返回偏移后的新的 RECT.
func OffsetRect(rc RECT, left, top, right, bottom int32) RECT {
	var rc2 RECT
	rc2.Left = rc.Left + left
	rc2.Top = rc.Top + top
	rc2.Right = rc.Right + right
	rc2.Bottom = rc.Bottom + bottom
	return rc2
}

// Font_Info_Name 将 [32]uint16 转换到 string.
//
// arr: [32]uint16.
func Font_Info_Name(arr [32]uint16) string {
	return syscall.UTF16ToString(arr[0:])
}

// Rect2RectF 将 RECT 转换到 RECTF.
func Rect2RectF(rc RECT) RECTF {
	return RECTF{
		Left:   float32(rc.Left),
		Right:  float32(rc.Right),
		Top:    float32(rc.Top),
		Bottom: float32(rc.Bottom),
	}
}

// PointInRect 判断一个点是否在矩形范围内.
func PointInRect(pt POINT, rc RECT) bool {
	return pt.X <= rc.Right && pt.X >= rc.Left && pt.Y <= rc.Bottom && pt.Y >= rc.Top
}

// PointToUintptr 将一个 POINT 结构强制转换为 uintptr。用于传递 POINT 结构给 COM 方法。
func PointToUintptr(pt POINT) uintptr {
	return *(*uintptr)(unsafe.Pointer(&pt))
}
