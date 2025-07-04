package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_20 是 ICoreWebView2_19 接口的延续，用于获取框架ID。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_20
type ICoreWebView2_20 struct {
	Vtbl *ICoreWebView2_20Vtbl
}

type ICoreWebView2_20Vtbl struct {
	ICoreWebView2_19Vtbl
	GetFrameId ComProc
}

func (i *ICoreWebView2_20) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_20) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_20) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
GetFrameId
*/
