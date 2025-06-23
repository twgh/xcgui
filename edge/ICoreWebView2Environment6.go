package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment6 是 ICoreWebView2Environment5 的延续，支持创建用于打印为 PDF 的打印设置。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment6
type ICoreWebView2Environment6 struct {
	Vtbl *ICoreWebView2Environment6Vtbl
}

type ICoreWebView2Environment6Vtbl struct {
	ICoreWebView2Environment5Vtbl
	CreatePrintSettings ComProc
}

func (i *ICoreWebView2Environment6) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment6) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment6) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreatePrintSettings 创建 PrintToPdf 方法使用的 ICoreWebView2PrintSettings。
func (i *ICoreWebView2Environment6) CreatePrintSettings() (*ICoreWebView2PrintSettings, error) {
	var value *ICoreWebView2PrintSettings
	r, _, _ := i.Vtbl.CreatePrintSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustCreatePrintSettings 创建 PrintToPdf 方法使用的 ICoreWebView2PrintSettings。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2Environment6) MustCreatePrintSettings() *ICoreWebView2PrintSettings {
	value, err := i.CreatePrintSettings()
	ReportErrorAtuo(err)
	return value
}
