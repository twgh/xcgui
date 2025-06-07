package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_11 是 ICoreWebView2_10 接口的延续，支持 CDP 方法调用的 sessionId 和 已请求上下文菜单 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_11
type ICoreWebView2_11 struct {
	Vtbl *ICoreWebView2_11Vtbl
}

type ICoreWebView2_11Vtbl struct {
	ICoreWebView2_10Vtbl
	CallDevToolsProtocolMethodForSession ComProc
	AddContextMenuRequested              ComProc
	RemoveContextMenuRequested           ComProc
}

func (i *ICoreWebView2_11) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_11) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_11) QueryInterface(refiid, object unsafe.Pointer) error {
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
CallDevToolsProtocolMethodForSession
AddContextMenuRequested
RemoveContextMenuRequested
*/
