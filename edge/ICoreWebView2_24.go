package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_24 是 ICoreWebView2_23 接口的延续，支持 通知接收 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_24
type ICoreWebView2_24 struct {
	Vtbl *ICoreWebView2_24Vtbl
}

type ICoreWebView2_24Vtbl struct {
	ICoreWebView2_23Vtbl
	AddNotificationReceived    ComProc
	RemoveNotificationReceived ComProc
}

func (i *ICoreWebView2_24) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_24) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_24) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddNotificationReceived
RemoveNotificationReceived
*/
