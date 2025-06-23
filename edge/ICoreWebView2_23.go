package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_23 是 ICoreWebView2_22 接口的延续，支持发送带有附加对象的 JSON Web 消息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_23
type ICoreWebView2_23 struct {
	Vtbl *ICoreWebView2_23Vtbl
}

type ICoreWebView2_23Vtbl struct {
	ICoreWebView2_22Vtbl
	PostWebMessageAsJsonWithAdditionalObjects ComProc
}

func (i *ICoreWebView2_23) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_23) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_23) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
PostWebMessageAsJsonWithAdditionalObjects
*/
