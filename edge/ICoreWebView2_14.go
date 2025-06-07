package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_14 是 ICoreWebView2_13 接口的延续，支持 检测到服务器证书错误 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_14
type ICoreWebView2_14 struct {
	Vtbl *ICoreWebView2_14Vtbl
}

type ICoreWebView2_14Vtbl struct {
	ICoreWebView2_13Vtbl
	AddServerCertificateErrorDetected    ComProc
	RemoveServerCertificateErrorDetected ComProc
	ClearServerCertificateErrorActions   ComProc
}

func (i *ICoreWebView2_14) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_14) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_14) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddServerCertificateErrorDetected
RemoveServerCertificateErrorDetected
ClearServerCertificateErrorActions
*/
