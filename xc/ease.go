package xc

// 缓动_Linear.
//
// p: 位置, 0.0f - 1.0f.
func XEase_Linear(p float32) int {
	r, _, _ := xEase_Linear.Call(float32Ptr(p))
	return int(r)
}

// 缓动_Quad.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quad(p float32, flag int) int {
	r, _, _ := xEase_Quad.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Cubic.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Cubic(p float32, flag int) int {
	r, _, _ := xEase_Cubic.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Quart.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quart(p float32, flag int) int {
	r, _, _ := xEase_Quart.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Quint.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quint(p float32, flag int) int {
	r, _, _ := xEase_Quint.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Sine.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Sine(p float32, flag int) int {
	r, _, _ := xEase_Sine.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Expo.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Expo(p float32, flag int) int {
	r, _, _ := xEase_Expo.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Circ.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Circ(p float32, flag int) int {
	r, _, _ := xEase_Circ.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Elastic.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Elastic(p float32, flag int) int {
	r, _, _ := xEase_Elastic.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Back.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Back(p float32, flag int) int {
	r, _, _ := xEase_Back.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}

// 缓动_Bounce.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Bounce(p float32, flag int) int {
	r, _, _ := xEase_Bounce.Call(float32Ptr(p), uintptr(flag))
	return int(r)
}
