package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Controller2 是 ICoreWebView2Controller 接口的延续。支持设置 WebView2 的默认背景色。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2controller2
type ICoreWebView2Controller2 struct {
	ICoreWebView2Controller
}

// GetDefaultBackgroundColor 获取 WebView2 的默认背景色。
func (i *ICoreWebView2Controller2) GetDefaultBackgroundColor() (*COREWEBVIEW2_COLOR, error) {
	var backgroundColor *COREWEBVIEW2_COLOR
	r, _, _ := i.Vtbl.GetDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&backgroundColor)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return backgroundColor, nil
}

// SetDefaultBackgroundColor 设置 WebView2 的默认背景色。
func (i *ICoreWebView2Controller2) SetDefaultBackgroundColor(backgroundColor COREWEBVIEW2_COLOR) error {
	// Cast to a uint32 as that's what the call is expecting
	col := *(*uint32)(unsafe.Pointer(&backgroundColor))

	r, _, _ := i.Vtbl.PutDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(col),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// COREWEBVIEW2_COLOR 表示 WebView2 的 RGBA 颜色值。
type COREWEBVIEW2_COLOR struct {
	A uint8
	R uint8
	G uint8
	B uint8
}
