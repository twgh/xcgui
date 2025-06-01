package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2PermissionRequestedEventArgs 是请求权限事件的参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs
type ICoreWebView2PermissionRequestedEventArgs struct {
	Vtbl *ICoreWebView2PermissionRequestedEventArgsVtbl
}

type ICoreWebView2PermissionRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetUri             ComProc
	GetPermissionKind  ComProc
	GetIsUserInitiated ComProc
	GetState           ComProc
	PutState           ComProc
	GetDeferral        ComProc
}

func (i *ICoreWebView2PermissionRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionRequestedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取请求权限的 web 内容的来源。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetUri() (string, error) {
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

// MustGetUri 获取请求权限的 web 内容的来源。出错时会触发全局错误回调.
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// GetPermissionKind 获取请求的权限类型。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetPermissionKind() (COREWEBVIEW2_PERMISSION_KIND, error) {
	var kind COREWEBVIEW2_PERMISSION_KIND
	r, _, err := i.Vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetPermissionKind 获取请求的权限类型。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetPermissionKind() COREWEBVIEW2_PERMISSION_KIND {
	kind, err := i.GetPermissionKind()
	ReportErrorAtuo(err)
	return kind
}

// GetIsUserInitiated 获取权限请求是否由用户发起。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetIsUserInitiated() (bool, error) {
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

// MustGetIsUserInitiated 获取权限请求是否由用户发起。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetIsUserInitiated() bool {
	isUserInitiated, err := i.GetIsUserInitiated()
	ReportErrorAtuo(err)
	return isUserInitiated
}

// GetState 获取权限请求的状态。默认值为 COREWEBVIEW2_PERMISSION_STATE_DEFAULT。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetState() (COREWEBVIEW2_PERMISSION_STATE, error) {
	var state COREWEBVIEW2_PERMISSION_STATE
	r, _, err := i.Vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&state)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return state, nil
}

// MustGetState 获取权限请求的状态。默认值为 COREWEBVIEW2_PERMISSION_STATE_DEFAULT。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetState() COREWEBVIEW2_PERMISSION_STATE {
	state, err := i.GetState()
	ReportErrorAtuo(err)
	return state
}

// PutState 设置权限请求的状态。
func (i *ICoreWebView2PermissionRequestedEventArgs) PutState(state COREWEBVIEW2_PERMISSION_STATE) error {
	r, _, err := i.Vtbl.PutState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(state),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取权限请求的延迟对象。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// MustGetDeferral 获取权限请求的延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}
