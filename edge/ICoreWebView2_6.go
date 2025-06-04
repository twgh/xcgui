package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2_6 是 ICoreWebView2_5 接口的延续，用于管理打开浏览器任务管理器窗口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_6
type ICoreWebView2_6 struct {
	Vtbl *ICoreWebView2_6Vtbl
}

type ICoreWebView2_6Vtbl struct {
	ICoreWebView2_5Vtbl
	OpenTaskManagerWindow ComProc
}

func (i *ICoreWebView2_6) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_6) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_6) QueryInterface(refiid, object uintptr) error {
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
OpenTaskManagerWindow
*/
