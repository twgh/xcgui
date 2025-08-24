package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions8 提供用于创建 WebView2 环境以管理滚动条样式的附加选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8
type ICoreWebView2EnvironmentOptions8 struct {
	ICoreWebView2EnvironmentOptions7
}

// GetScrollBarStyle 获取滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) GetScrollBarStyle() (COREWEBVIEW2_SCROLLBAR_STYLE, error) {
	var value COREWEBVIEW2_SCROLLBAR_STYLE
	r, _, _ := i.Vtbl.GetScrollBarStyle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetScrollBarStyle 设置滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) SetScrollBarStyle(value COREWEBVIEW2_SCROLLBAR_STYLE) error {
	r, _, _ := i.Vtbl.PutScrollBarStyle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetScrollBarStyle 获取滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) MustGetScrollBarStyle() COREWEBVIEW2_SCROLLBAR_STYLE {
	value, err := i.GetScrollBarStyle()
	ReportErrorAtuo(err)
	return value
}
