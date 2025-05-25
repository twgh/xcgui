package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2NavigationStartingEventArgs 是 NavigationStarting 和 FrameNavigationStarting 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs
type ICoreWebView2NavigationStartingEventArgs struct {
	Vtbl *ICoreWebView2NavigationStartingEventArgsVtbl
}

type ICoreWebView2NavigationStartingEventArgsVtbl struct {
	IUnknownVtbl
	GetUri             ComProc
	GetIsUserInitiated ComProc
	GetIsRedirected    ComProc
	GetRequestHeaders  ComProc
	GetCancel          ComProc
	PutCancel          ComProc
	GetNavigationId    ComProc
}

func (i *ICoreWebView2NavigationStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationStartingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationStartingEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取请求的导航的 uri。
func (i *ICoreWebView2NavigationStartingEventArgs) GetUri() (string, error) {
	var _uri *uint16
	r, _, err := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// MustGetUri 获取请求的导航的 uri。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportError2(err)
	return uri
}

// GetNavigationId 获取导航的 ID。
func (i *ICoreWebView2NavigationStartingEventArgs) GetNavigationId() (uint64, error) {
	var id uint64
	r, _, err := i.Vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return id, err
	}
	if r != 0 {
		return id, syscall.Errno(r)
	}
	return id, nil
}

// MustGetNavigationId 获取导航的 ID。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetNavigationId() uint64 {
	id, err := i.GetNavigationId()
	ReportError2(err)
	return id
}

// GetIsUserInitiated 获取导航是否由用户发起。
func (i *ICoreWebView2NavigationStartingEventArgs) GetIsUserInitiated() (bool, error) {
	var isUserInitiated bool
	r, _, err := i.Vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isUserInitiated)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return isUserInitiated, err
	}
	if r != 0 {
		return isUserInitiated, syscall.Errno(r)
	}
	return isUserInitiated, nil
}

// MustGetIsUserInitiated 获取导航是否由用户发起。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetIsUserInitiated() bool {
	isUserInitiated, err := i.GetIsUserInitiated()
	ReportError2(err)
	return isUserInitiated
}

// GetIsRedirected 获取导航是否是重定向的结果。
func (i *ICoreWebView2NavigationStartingEventArgs) GetIsRedirected() (bool, error) {
	var isRedirected bool
	r, _, err := i.Vtbl.GetIsRedirected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isRedirected)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return isRedirected, err
	}
	if r != 0 {
		return isRedirected, syscall.Errno(r)
	}
	return isRedirected, nil
}

// MustGetIsRedirected 获取导航是否是重定向的结果。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetIsRedirected() bool {
	isRedirected, err := i.GetIsRedirected()
	ReportError2(err)
	return isRedirected
}

// GetRequestHeaders 获取导航请求的 HTTP 请求头。你无法在 NavigationStarting 事件中修改HTTP请求标头。
func (i *ICoreWebView2NavigationStartingEventArgs) GetRequestHeaders() (*ICoreWebView2HttpRequestHeaders, error) {
	var headers *ICoreWebView2HttpRequestHeaders
	r, _, err := i.Vtbl.GetRequestHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return headers, nil
}

// MustGetRequestHeaders 获取导航请求的 HTTP 请求头。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetRequestHeaders() *ICoreWebView2HttpRequestHeaders {
	headers, err := i.GetRequestHeaders()
	ReportError2(err)
	return headers
}

// GetCancel 获取是否取消导航。
func (i *ICoreWebView2NavigationStartingEventArgs) GetCancel() (bool, error) {
	var cancel bool
	r, _, err := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return cancel, err
	}
	if r != 0 {
		return cancel, syscall.Errno(r)
	}
	return cancel, nil
}

// MustGetCancel 获取是否取消导航。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs) MustGetCancel() bool {
	cancel, err := i.GetCancel()
	ReportError2(err)
	return cancel
}

// PutCancel 设置是否取消导航。
func (i *ICoreWebView2NavigationStartingEventArgs) PutCancel(cancel bool) error {
	r, _, err := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(cancel),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
