package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_12 是 ICoreWebView2_11 接口的延续，支持 状态栏文本已更改 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_12
type ICoreWebView2_12 struct {
	Vtbl *ICoreWebView2_12Vtbl
}

type ICoreWebView2_12Vtbl struct {
	ICoreWebView2_11Vtbl
	AddStatusBarTextChanged    ComProc
	RemoveStatusBarTextChanged ComProc
	GetStatusBarText           ComProc
}

func (i *ICoreWebView2_12) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_12) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_12) QueryInterface(refiid, object unsafe.Pointer) error {
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
AddStatusBarTextChanged
RemoveStatusBarTextChanged
GetStatusBarText
*/
