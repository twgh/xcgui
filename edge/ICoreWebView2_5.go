package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_5 是 ICoreWebView2_4 接口的延续，以支持 请求客户端证书 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_5
type ICoreWebView2_5 struct {
	Vtbl *ICoreWebView2_5Vtbl
}

type ICoreWebView2_5Vtbl struct {
	ICoreWebView2_4Vtbl
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
}

func (i *ICoreWebView2_5) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_5) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_5) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddClientCertificateRequested
RemoveClientCertificateRequested
*/
