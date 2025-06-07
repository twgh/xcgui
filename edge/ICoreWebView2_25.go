package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_25 是 ICoreWebView2_24 接口的延续，支持"另存为"UI相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_25
type ICoreWebView2_25 struct {
	Vtbl *ICoreWebView2_25Vtbl
}

type ICoreWebView2_25Vtbl struct {
	ICoreWebView2_24Vtbl
	AddSaveAsUIShowing    ComProc
	RemoveSaveAsUIShowing ComProc
	ShowSaveAsUI          ComProc
}

func (i *ICoreWebView2_25) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_25) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_25) QueryInterface(refiid, object unsafe.Pointer) error {
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
AddSaveAsUIShowing
RemoveSaveAsUIShowing
ShowSaveAsUI
*/
