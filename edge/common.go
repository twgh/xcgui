package edge

import (
	"unsafe"
)

// COREWEBVIEW2_COLOR 表示 WebView2 的 RGBA 颜色值。
type COREWEBVIEW2_COLOR struct {
	A uint8
	R uint8
	G uint8
	B uint8
}

func NewColor(r, g, b, a uint8) *COREWEBVIEW2_COLOR {
	return &COREWEBVIEW2_COLOR{R: r, G: g, B: b, A: a}
}

// ToUint32 将其转换为 uint32, 这是 COM 方法底层所期望的.
func (c *COREWEBVIEW2_COLOR) ToUint32() uint32 {
	return *(*uint32)(unsafe.Pointer(c))
}
