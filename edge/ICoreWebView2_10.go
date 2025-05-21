package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_10 是 ICoreWebView2_9 接口的延续，支持 请求基本身份验证 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_10
type ICoreWebView2_10 struct {
	Vtbl *ICoreWebView2_10Vtbl
}

type ICoreWebView2_10Vtbl struct {
	ICoreWebView2_9Vtbl
	AddBasicAuthenticationRequested    ComProc
	RemoveBasicAuthenticationRequested ComProc
}

func (i *ICoreWebView2_10) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_10) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_10) QueryInterface(refiid, object uintptr) error {
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
AddBasicAuthenticationRequested
RemoveBasicAuthenticationRequested
*/
