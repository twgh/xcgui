package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_19 是 ICoreWebView2_18 接口的延续，用于管理内存使用目标级别。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_19
type ICoreWebView2_19 struct {
	Vtbl *ICoreWebView2_19Vtbl
}

type ICoreWebView2_19Vtbl struct {
	ICoreWebView2_18Vtbl
	GetMemoryUsageTargetLevel ComProc
	PutMemoryUsageTargetLevel ComProc
}

func (i *ICoreWebView2_19) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_19) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_19) QueryInterface(refiid, object unsafe.Pointer) error {
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
GetMemoryUsageTargetLevel
PutMemoryUsageTargetLevel
*/
