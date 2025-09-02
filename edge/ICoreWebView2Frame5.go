package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame5 是 ICoreWebView2Frame4 的延续，添加了对 FrameId 的支持。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame5
type ICoreWebView2Frame5 struct {
	ICoreWebView2Frame4
}

// GetFrameId 获取此框架的唯一 ID。
//   - 它与 FrameId 在 ICoreWebView2 中以及通过 ICoreWebView2FrameInfo 所使用的是同一种类型的 ID。
func (i *ICoreWebView2Frame5) GetFrameId() (uint32, error) {
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

// MustGetFrameId 获取此框架的唯一 ID。出错时会触发全局错误回调。
//   - 它与 FrameId 在 ICoreWebView2 中以及通过 ICoreWebView2FrameInfo 所使用的是同一种类型的 ID。
func (i *ICoreWebView2Frame5) MustGetFrameId() uint32 {
	r, err := i.GetFrameId()
	ReportErrorAtuo(err)
	return r
}
