package xc

// 缓动_Linear.
//
// p: 位置, 0.0f - 1.0f.
func XEase_Linear(p float32) float32 {
	_, r, _ := xEase_Linear.Call(float32Ptr(p))
	return uintPtrToFloat32(r)
}

// 缓动_Quad.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quad(p float32, flag int) float32 {
	_, r, _ := xEase_Quad.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Cubic.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Cubic(p float32, flag int) float32 {
	_, r, _ := xEase_Cubic.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Quart.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quart(p float32, flag int) float32 {
	_, r, _ := xEase_Quart.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Quint.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Quint(p float32, flag int) float32 {
	_, r, _ := xEase_Quint.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Sine.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Sine(p float32, flag int) float32 {
	_, r, _ := xEase_Sine.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Expo.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Expo(p float32, flag int) float32 {
	_, r, _ := xEase_Expo.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Circ.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Circ(p float32, flag int) float32 {
	_, r, _ := xEase_Circ.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Elastic.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Elastic(p float32, flag int) float32 {
	_, r, _ := xEase_Elastic.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Back.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Back(p float32, flag int) float32 {
	_, r, _ := xEase_Back.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}

// 缓动_Bounce.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动方式, Ease_.
func XEase_Bounce(p float32, flag int) float32 {
	_, r, _ := xEase_Bounce.Call(float32Ptr(p), uintptr(flag))
	return uintPtrToFloat32(r)
}
