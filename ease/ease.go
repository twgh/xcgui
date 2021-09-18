// 缓动
package ease

import "github.com/twgh/xcgui/xc"

// 缓动_Linear
// p: 位置, 0.0f - 1.0f
func Linear(p float32) int {
	return xc.XEase_Linear(p)
}

// 缓动_Quad
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Quad(p float32, flag int) int {
	return xc.XEase_Quad(p, flag)
}

// 缓动_Cubic
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Cubic(p float32, flag int) int {
	return xc.XEase_Cubic(p, flag)
}

// 缓动_Quart
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Quart(p float32, flag int) int {
	return xc.XEase_Quart(p, flag)
}

// 缓动_Quint
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Quint(p float32, flag int) int {
	return xc.XEase_Quint(p, flag)
}

// 缓动_Sine
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Sine(p float32, flag int) int {
	return xc.XEase_Sine(p, flag)
}

// 缓动_Expo
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Expo(p float32, flag int) int {
	return xc.XEase_Expo(p, flag)
}

// 缓动_Circ
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Circ(p float32, flag int) int {
	return xc.XEase_Circ(p, flag)
}

// 缓动_Elastic
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Elastic(p float32, flag int) int {
	return xc.XEase_Elastic(p, flag)
}

// 缓动_Back
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Back(p float32, flag int) int {
	return xc.XEase_Back(p, flag)
}

// 缓动_Bounce
// p: 位置, 0.0f - 1.0f
// flag: 缓动方式, Ease_
func Bounce(p float32, flag int) int {
	return xc.XEase_Bounce(p, flag)
}
