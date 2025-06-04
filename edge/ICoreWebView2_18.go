package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_18 是 ICoreWebView2_17 接口的延续，它管理对在操作系统中注册的 URI 方案的导航请求。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_18
type ICoreWebView2_18 struct {
	Vtbl *ICoreWebView2_18Vtbl
}

type ICoreWebView2_18Vtbl struct {
	ICoreWebView2_17Vtbl
	AddLaunchingExternalUriScheme    ComProc
	RemoveLaunchingExternalUriScheme ComProc
}

func (i *ICoreWebView2_18) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_18) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_18) QueryInterface(refiid, object uintptr) error {
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
AddLaunchingExternalUriScheme
RemoveLaunchingExternalUriScheme
*/
