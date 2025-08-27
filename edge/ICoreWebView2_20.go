package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_20 是 ICoreWebView2_19 接口的延续，用于获取框架ID。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_20
type ICoreWebView2_20 struct {
	ICoreWebView2_19
}

// GetFrameId 获取当前 WebView2 主框架的唯一标识符。
//   - 这与 FrameId 在 ICoreWebView2Frame 中以及通过 ICoreWebView2FrameInfo 获取的 ID 类型相同。
//   - 请注意，如果 ICoreWebView2 尚未进行任何导航，FrameId 可能无效。
//   - 在第一次 ContentLoading 事件期间或之后获取此值是安全的。
//   - 否则，它可能会返回无效的框架 ID: 0。
func (i *ICoreWebView2_20) GetFrameId() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetFrameId 获取当前 WebView2 主框架的唯一标识符。出错时会触发全局错误回调。
//   - 这与 FrameId 在 ICoreWebView2Frame 中以及通过 ICoreWebView2FrameInfo 获取的 ID 类型相同。
//   - 请注意，如果 ICoreWebView2 尚未进行任何导航，FrameId 可能无效。
//   - 在第一次 ContentLoading 事件期间或之后获取此值是安全的。
//   - 否则，它可能会返回无效的框架 ID: 0。
func (i *ICoreWebView2_20) MustGetFrameId() uint32 {
	value, err := i.GetFrameId()
	ReportErrorAtuo(err)
	return value
}
