package xc

import (
	"github.com/twgh/xcgui/xcc"
	"os"
	"strconv"
	"syscall"
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

// PathExists 判断文件或文件夹是否存在. 不考虑极端情况.
//
// path: 文件或文件夹路径.
func PathExists2(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Font_Info_Name 将[32]uint16转换到string.
//
// arr: [32]uint16.
func Font_Info_Name(arr [32]uint16) string {
	return syscall.UTF16ToString(arr[0:])
}

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

// ARGB 根据r, g, b, a组合成ARGB颜色.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
//
// a: 透明度.
func ARGB(r, g, b, a byte) int {
	return int(uint32(a)<<24 | uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// ARGB2 根据rgb, a组合成十进制ARGB颜色.
//
// rgb: RGB颜色.
//
// a: 透明度.
func ARGB2(rgb int, a byte) int {
	return int((uint32(rgb) & 16777215) | uint32(a)<<24)
}

// RGBA 根据r, g, b, a组合成ARGB颜色. 和 ARGB 函数一模一样, 只是为了符合部分人使用习惯.
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

// RGBA2 根据rgb, a组合成十进制ARGB颜色. 和 ARGB2 函数一模一样, 只是为了符合部分人使用习惯.
//
// rgb: RGB颜色.
//
// a: 透明度.
func RGBA2(rgb int, a byte) int {
	return int((uint32(rgb) & 16777215) | uint32(a)<<24)
}

// RGB 根据r, g, b组合成RGB颜色.
//
// r: 红色分量.
//
// g: 绿色分量.
//
// b: 蓝色分量.
func RGB(r, g, b byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// RGB2ARGB 将RGB颜色转换到ARGB颜色.
//
// rgb: RGB颜色.
//
// a: 透明度.
func RGB2ARGB(rgb int, a byte) int {
	r := byte(rgb & 255)
	g := byte((rgb >> 8) & 255)
	b := byte((rgb >> 16) & 255)
	return ARGB(r, g, b, a)
}

// HexRGB2RGB 将十六进制RGB颜色转换到十进制RGB颜色.
//
// str: 十六进制RGB颜色, 开头有没有#都可以.
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

// HexRGB2ARGB 将十六进制RGB颜色转换到十进制ARGB颜色.
//
// str: 十六进制RGB颜色, 开头有没有#都可以.
//
// a: 透明度.
func HexRGB2ARGB(str string, a byte) int {
	return ARGB2(HexRGB2RGB(str), a)
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

// SetBnClicks 给一个窗口或布局元素中所有的按钮注册按钮单击事件.
//
// 说明: 本函数是通过遍历子元素实现的, 只会给窗口和布局元素中的按钮注册事件, 像List, Tree, 滑块条等元素中的按钮不会注册, 想要更多可以自己实现一个.
//
//	HXCGUI 炫彩窗口句柄或布局元素句柄.
//
//	onBnClick 按钮单击事件回调函数.
func SetBnClicks(HXCGUI int, onBnClick func(hEle int, pbHandled *bool) int) {
	if XC_IsHWINDOW(HXCGUI) {
		for i := int32(0); i < XWnd_GetChildCount(HXCGUI); i++ {
			hEle := XWnd_GetChildByIndex(HXCGUI, i)
			switch XC_GetObjectType(hEle) {
			case xcc.XC_ELE_LAYOUT:
				SetBnClicks(hEle, onBnClick)
			case xcc.XC_BUTTON:
				XEle_RegEventC1(hEle, xcc.XE_BNCLICK, onBnClick)
			}
		}
	} else if XC_GetObjectType(HXCGUI) == xcc.XC_ELE_LAYOUT {
		for i := int32(0); i < XEle_GetChildCount(HXCGUI); i++ {
			hEle := XEle_GetChildByIndex(HXCGUI, i)
			switch XC_GetObjectType(hEle) {
			case xcc.XC_ELE_LAYOUT:
				SetBnClicks(hEle, onBnClick)
			case xcc.XC_BUTTON:
				XEle_RegEventC1(hEle, xcc.XE_BNCLICK, onBnClick)
			}
		}
	}
}

// Itoa 将int32转换到string.
func Itoa(i32 int32) string {
	return strconv.FormatInt(int64(i32), 10)
}

// Atoi 将string转换到int32.
func Atoi(s string) int32 {
	i, _ := strconv.Atoi(s)
	return int32(i)
}
