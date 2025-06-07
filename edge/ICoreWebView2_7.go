package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2_7 是 ICoreWebView2_6 接口的延续，支持打印为PDF。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_7
type ICoreWebView2_7 struct {
	Vtbl *ICoreWebView2_7Vtbl
}

type ICoreWebView2_7Vtbl struct {
	ICoreWebView2_6Vtbl
	PrintToPdf ComProc
}

func (i *ICoreWebView2_7) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_7) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_7) QueryInterface(refiid, object unsafe.Pointer) error {
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
PrintToPdf
*/
