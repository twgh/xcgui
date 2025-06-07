package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_13 是 ICoreWebView2_12 接口的延续，支持获取关联的 ICoreWebView2Profile 对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_13
type ICoreWebView2_13 struct {
	Vtbl *ICoreWebView2_13Vtbl
}

type ICoreWebView2_13Vtbl struct {
	ICoreWebView2_12Vtbl
	GetProfile ComProc
}

func (i *ICoreWebView2_13) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_13) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_13) QueryInterface(refiid, object unsafe.Pointer) error {
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
GetProfile
*/
