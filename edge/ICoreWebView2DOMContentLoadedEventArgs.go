package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2DOMContentLoadedEventArgs 是 DOMContentLoaded 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs
type ICoreWebView2DOMContentLoadedEventArgs struct {
	Vtbl *ICoreWebView2DOMContentLoadedEventArgsVtbl
}

type ICoreWebView2DOMContentLoadedEventArgsVtbl struct {
	IUnknownVtbl
	GetNavigationId ComProc
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetNavigationId 获取导航的 ID，与其他导航事件中的其他导航 ID 属性相对应。
func (i *ICoreWebView2DOMContentLoadedEventArgs) GetNavigationId() (uint64, error) {
	var id uint64
	r, _, err := i.Vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return id, err
	}
	if r != 0 {
		return id, syscall.Errno(r)
	}
	return id, nil
}

// MustGetNavigationId 获取导航的 ID，与其他导航事件中的其他导航 ID 属性相对应。出错时会触发全局错误回调。
func (i *ICoreWebView2DOMContentLoadedEventArgs) MustGetNavigationId() uint64 {
	id, err := i.GetNavigationId()
	ReportErrorAtuo(err)
	return id
}
