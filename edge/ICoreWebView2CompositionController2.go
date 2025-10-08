package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2CompositionController2 是 ICoreWebView2CompositionController 的延续，用于获取 WebView 的自动化提供程序。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller2
type ICoreWebView2CompositionController2 struct {
	ICoreWebView2CompositionController
}

// GetAutomationProvider 获取 WebView 的自动化提供程序。
//   - 该对象实现了 IRawElementProviderSimple.
func (i *ICoreWebView2CompositionController2) GetAutomationProvider() (*IUnknown, error) {
	var value *IUnknown
	r, _, _ := i.Vtbl.GetAutomationProvider.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetAutomationProvider 获取 WebView 的自动化提供程序。出错时会触发全局错误回调。
//   - 该对象实现了 IRawElementProviderSimple.
func (i *ICoreWebView2CompositionController2) MustGetAutomationProvider() *IUnknown {
	value, err := i.GetAutomationProvider()
	ReportErrorAuto(err)
	return value
}
