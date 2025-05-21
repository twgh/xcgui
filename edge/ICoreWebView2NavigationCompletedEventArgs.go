package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2NavigationCompletedEventArgs 是 NavigationCompleted 事件的参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs
type ICoreWebView2NavigationCompletedEventArgs struct {
	Vtbl *ICoreWebView2NavigationCompletedEventArgsVtbl
}

type ICoreWebView2NavigationCompletedEventArgsVtbl struct {
	IUnknownVtbl
	GetIsSuccess      ComProc
	GetWebErrorStatus ComProc
	GetNavigationId   ComProc
}

func (i *ICoreWebView2NavigationCompletedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationCompletedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationCompletedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsSuccess 导航是否成功。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetIsSuccess() (bool, error) {
	var isSuccess bool
	r, _, err := i.Vtbl.GetIsSuccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSuccess)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return isSuccess, err
	}
	if r != 0 {
		return isSuccess, syscall.Errno(r)
	}
	return isSuccess, nil
}

// MustGetIsSuccess 导航是否成功。忽略错误。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetIsSuccess() bool {
	isSuccess, _ := i.GetIsSuccess()
	return isSuccess
}

// GetWebErrorStatus 获取导航失败时的错误代码。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetWebErrorStatus() (COREWEBVIEW2_WEB_ERROR_STATUS, error) {
	var status COREWEBVIEW2_WEB_ERROR_STATUS
	r, _, err := i.Vtbl.GetWebErrorStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&status)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return status, err
	}
	if r != 0 {
		return status, syscall.Errno(r)
	}
	return status, nil
}

// MustGetWebErrorStatus 获取导航错误状态。忽略错误。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetWebErrorStatus() COREWEBVIEW2_WEB_ERROR_STATUS {
	status, _ := i.GetWebErrorStatus()
	return status
}

// GetNavigationId 获取导航 ID。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetNavigationId() (uint64, error) {
	var navigationId uint64
	r, _, err := i.Vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return navigationId, err
	}
	if r != 0 {
		return navigationId, syscall.Errno(r)
	}
	return navigationId, nil
}

// MustGetNavigationId 获取导航 ID。忽略错误。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetNavigationId() uint64 {
	id, _ := i.GetNavigationId()
	return id
}
