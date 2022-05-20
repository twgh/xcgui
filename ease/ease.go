package ease

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 缓动_Linear, 线性.
//
// p: 位置, 0.0f - 1.0f.
func Linear(p float32) float32 {
	return xc.XEase_Linear(p)
}

// 缓动_Quad, 二次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Quad(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quad(p, flag)
}

// 缓动_Cubic, 三次方曲线, 圆弧.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Cubic(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Cubic(p, flag)
}

// 缓动_Quart, 四次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Quart(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quart(p, flag)
}

// 缓动_Quint, 五次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Quint(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quint(p, flag)
}

// 缓动_Sine, 正弦曲线, 在末端变化.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Sine(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Sine(p, flag)
}

// 缓动_Expo, 突击曲线, 突然一下.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Expo(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Expo(p, flag)
}

// 缓动_Circ, 圆环, 好比绕过一个圆环.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Circ(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Circ(p, flag)
}

// 缓动_Elastic, 强力回弹.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Elastic(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Elastic(p, flag)
}

// 缓动_Back, 回弹, 比较缓慢.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Back(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Back(p, flag)
}

// 缓动_Bounce.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func Bounce(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Bounce(p, flag)
}

// 缓动_扩展, 全部缓动类型.
//
// pos: 位置, 0.0f - 1.0f.
//
// flag: 缓动标识, Ease_Flag_.
func Ex(pos float32, flag xcc.Ease_Flag_) float32 {
	return xc.XEase_Ex(pos, flag)
}
