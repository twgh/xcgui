package xc

import (
	"github.com/twgh/xcgui/xcc"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// PathExists 判断文件或文件夹是否存在.
//
//	@param path 文件或文件夹.
//	@return error 如果出错, 则不确定是否存在.
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

// Font_Info_Name 将[32]uint16转换到string.
//
//	@param arr [32]uint16.
//	@return string
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

// ABGR 根据r, g, b, a组合成ABGR颜色.
//
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@param a 透明度.
//	@return int ABGR颜色.
func ABGR(r, g, b, a byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16 | uint32(a)<<24)
}

// ABGR2 根据bgr, a组合成ABGR颜色.
//
//	@param bgr BGR颜色.
//	@param a 透明度.
//	@return int ABGR颜色.
func ABGR2(bgr int, a byte) int {
	return int((uint32(bgr) & 16777215) | (uint32(a)&255)<<24)
}

// RGBA 根据r, g, b, a组合成ABGR颜色. 和 ABGR 函数一模一样, 只是为了符合部分人使用习惯.
//
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@param a 透明度.
//	@return int ABGR颜色.
func RGBA(r, g, b, a byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16 | uint32(a)<<24)
}

// RGBA2 根据bgr, a组合成十进制ABGR颜色. 和 ABGR2 函数一模一样, 只是为了符合部分人使用习惯.
//
//	@param bgr BGR颜色.
//	@param a 透明度.
//	@return int ABGR颜色.
func RGBA2(bgr int, a byte) int {
	return int((uint32(bgr) & 16777215) | uint32(a)<<24)
}

// BGR 根据r, g, b组合成BGR颜色.
//
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@return int BGR颜色.
func BGR(r, g, b byte) int {
	return int(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

// RGB 根据r, g, b组合成RGB颜色.
//
//	@param r 红色分量.
//	@param g 绿色分量.
//	@param b 蓝色分量.
//	@return int RGB颜色.
func RGB(r, g, b byte) int {
	return int(uint32(b) | uint32(g)<<8 | uint32(r)<<16)
}

// RGB2ABGR 将RGB颜色转换到ABGR颜色.
//
//	@param rgb RGB颜色.
//	@param a 透明度.
//	@return int ABGR颜色.
func RGB2ABGR(rgb int, a byte) int {
	r := rgb % 256
	g := (rgb / 256) % 256
	b := rgb / 65536
	return ABGR(byte(r), byte(g), byte(b), a)
}

// RGB2BGR 将RGB颜色转换到BGR颜色.
//
//	@param rgb RGB颜色.
//	@return int BGR颜色.
func RGB2BGR(rgb int) int {
	r := byte(rgb % 256)
	g := byte((rgb / 256) % 256)
	b := byte(rgb / 65536)
	return BGR(r, g, b)
}

// HexRGB2BGR 将十六进制RGB颜色转换到十进制BGR颜色.
//
//	@param str 十六进制RGB颜色, 开头有没有#都可以.
//	@return int BGR颜色.
func HexRGB2BGR(str string) int {
	s := strings.TrimLeft(str, "#")
	r, _ := strconv.ParseInt(s[:2], 16, 10)
	g, _ := strconv.ParseInt(s[2:4], 16, 18)
	b, _ := strconv.ParseInt(s[4:], 16, 10)
	return BGR(byte(r), byte(g), byte(b))
}

// HexRGB2ABGR 将十六进制RGB颜色转换到十进制ABGR颜色.
//
//	@param str 十六进制RGB颜色, 开头有没有#都可以.
//	@param a 透明度.
//	@return int ABGR颜色.
func HexRGB2ABGR(str string, a byte) int {
	s := strings.TrimLeft(str, "#")
	r, _ := strconv.ParseInt(s[:2], 16, 10)
	g, _ := strconv.ParseInt(s[2:4], 16, 18)
	b, _ := strconv.ParseInt(s[4:], 16, 10)
	return ABGR(byte(r), byte(g), byte(b), a)
}

// HexRGB2RGB 将十六进制RGB颜色转换到十进制RGB颜色.
//
//	@param str 十六进制RGB颜色, 开头有没有#都可以.
//	@return int RGB颜色.
func HexRGB2RGB(str string) int {
	s := strings.TrimLeft(str, "#")
	r, _ := strconv.ParseInt(s[:2], 16, 10)
	g, _ := strconv.ParseInt(s[2:4], 16, 18)
	b, _ := strconv.ParseInt(s[4:], 16, 10)
	return RGB(byte(r), byte(g), byte(b))
}

// SetBnClicks 给一个窗口或布局元素中所有的按钮注册按钮单击事件.
//
// 说明: 本函数是通过遍历子元素实现的, 只会给窗口和布局元素中的按钮注册事件, 像List, Tree, 滑块条等元素中的按钮不会注册, 想要更多可以自己实现一个.
//
//	@param HXCGUI 炫彩窗口句柄或布局元素句柄.
//	@param onBnClick 按钮单击事件回调函数.
func SetBnClicks(HXCGUI int, onBnClick func(hEle int, pbHandled *bool) int) {
	if XC_IsHWINDOW(HXCGUI) {
		for i := 0; i < XWnd_GetChildCount(HXCGUI); i++ {
			hEle := XWnd_GetChildByIndex(HXCGUI, i)
			switch XC_GetObjectType(hEle) {
			case xcc.XC_ELE_LAYOUT:
				SetBnClicks(hEle, onBnClick)
			case xcc.XC_BUTTON:
				XEle_RegEventC1(hEle, xcc.XE_BNCLICK, onBnClick)
			}
		}
	} else if XC_GetObjectType(HXCGUI) == xcc.XC_ELE_LAYOUT {
		for i := 0; i < XEle_GetChildCount(HXCGUI); i++ {
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

// Deprecated
//
// !已废弃, 请使用 xc.XWnd_ClientToScreen 或 wapi.ClientToScreen
func ClientToScreen(hWindow int, pPoint *POINT) {
	XWnd_ClientToScreen(hWindow, pPoint)
}
