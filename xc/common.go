package xc

import (
	"math"
	"os"
	"strconv"
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

// RGBA 根据r, g, b, a组合成炫彩使用的颜色.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
//
// a: 透明度.
func RGBA(r, g, b, a byte) int {
	return int(uint32(a)<<24 | uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// RGBA2 根据 RGB, a组合成炫彩使用的颜色.
//
// rgb: RGB 颜色.
//
// a: 透明度.
func RGBA2(rgb int, a byte) int {
	return int((uint32(rgb) & 16777215) | uint32(a)<<24)
}

// RGB 根据r, g, b组合成炫彩以前使用的颜色. 现在已经不用这个了.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
func RGB(r, g, b byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// RGB2RGBA 将 RGB 颜色转换到炫彩使用的颜色.
//
// rgb: RGB 颜色.
//
// a: 透明度.
func RGB2RGBA(rgb int, a byte) int {
	r := byte(rgb & 255)
	g := byte((rgb >> 8) & 255)
	b := byte((rgb >> 16) & 255)
	return RGBA(r, g, b, a)
}

// HexRGB2RGB 将十六进制 RGB 颜色转换到十进制 RGB 颜色.
//
// str: 十六进制 RGB 颜色, 开头有没有#都可以.
func HexRGB2RGB(str string) int {
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
func HexRGB2RGBA(str string, a byte) int {
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
	return strconv.FormatInt(int64(i32), 10)
}

// Atoi 将 string 转换到 int32.
func Atoi(s string) int32 {
	i, _ := strconv.Atoi(s)
	return int32(i)
}

// Ftoa 将 float32 转换到 string.
func Ftoa(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// Atof 将 string 转换到 float32.
func Atof(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
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
