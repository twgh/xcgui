package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment4 是 ICoreWebView2Environment3 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment4
type ICoreWebView2Environment4 struct {
	Vtbl *ICoreWebView2Environment4Vtbl
}

type ICoreWebView2Environment4Vtbl struct {
	ICoreWebView2Environment3Vtbl
	GetAutomationProviderForWindow ComProc
}

func (i *ICoreWebView2Environment4) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment4) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment4) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAutomationProviderForWindow 获取指定窗口的 WebView 的自动化提供程序。
func (i *ICoreWebView2Environment4) GetAutomationProviderForWindow(hwnd uintptr) (*IUnknown, error) {
	var value *IUnknown
	r, _, _ := i.Vtbl.GetAutomationProviderForWindow.Call(
		uintptr(unsafe.Pointer(i)),
		hwnd,
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetAutomationProviderForWindow 获取指定窗口的 WebView 的自动化提供程序。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment4) MustGetAutomationProviderForWindow(hwnd uintptr) *IUnknown {
	provider, err := i.GetAutomationProviderForWindow(hwnd)
	ReportErrorAtuo(err)
	return provider
}
