package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_26 是 ICoreWebView2_25 接口的延续，支持 开始保存文件安全检查 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_26
type ICoreWebView2_26 struct {
	Vtbl *ICoreWebView2_26Vtbl
}

type ICoreWebView2_26Vtbl struct {
	ICoreWebView2_25Vtbl
	AddSaveFileSecurityCheckStarting    ComProc
	RemoveSaveFileSecurityCheckStarting ComProc
}

func (i *ICoreWebView2_26) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_26) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_26) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddSaveFileSecurityCheckStarting
RemoveSaveFileSecurityCheckStarting
*/
