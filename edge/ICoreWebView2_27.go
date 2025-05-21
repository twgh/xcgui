package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_27 是 ICoreWebView2_26 接口的延续，支持 屏幕捕获开始 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_27
type ICoreWebView2_27 struct {
	Vtbl *ICoreWebView2_27Vtbl
}

type ICoreWebView2_27Vtbl struct {
	ICoreWebView2_26Vtbl
	AddScreenCaptureStarting    ComProc
	RemoveScreenCaptureStarting ComProc
}

func (i *ICoreWebView2_27) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_27) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_27) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddScreenCaptureStarting
RemoveScreenCaptureStarting
*/
