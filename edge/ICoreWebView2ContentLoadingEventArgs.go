package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ContentLoadingEventArgs 是 ContentLoading 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs
type ICoreWebView2ContentLoadingEventArgs struct {
	Vtbl *ICoreWebView2ContentLoadingEventArgsVtbl
}

type ICoreWebView2ContentLoadingEventArgsVtbl struct {
	IUnknownVtbl
	GetIsErrorPage  ComProc
	GetNavigationId ComProc
}

func (i *ICoreWebView2ContentLoadingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContentLoadingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContentLoadingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsErrorPage 获取当前导航是否为错误页面。
func (i *ICoreWebView2ContentLoadingEventArgs) GetIsErrorPage() (bool, error) {
	var isErrorPage bool
	r, _, err := i.Vtbl.GetIsErrorPage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isErrorPage)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return isErrorPage, err
	}
	if r != 0 {
		return isErrorPage, syscall.Errno(r)
	}
	return isErrorPage, nil
}

// GetNavigationId 获取导航的 ID。
func (i *ICoreWebView2ContentLoadingEventArgs) GetNavigationId() (uint64, error) {
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

// MustGetIsErrorPage 获取当前导航是否为错误页面。出错时会触发全局错误回调。
func (i *ICoreWebView2ContentLoadingEventArgs) MustGetIsErrorPage() bool {
	isErrorPage, err := i.GetIsErrorPage()
	ReportErrorAtuo(err)
	return isErrorPage
}

// MustGetNavigationId 获取导航的 ID。出错时会触发全局错误回调。
func (i *ICoreWebView2ContentLoadingEventArgs) MustGetNavigationId() uint64 {
	id, err := i.GetNavigationId()
	ReportErrorAtuo(err)
	return id
}
