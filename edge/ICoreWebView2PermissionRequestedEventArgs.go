package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

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
	// 2
	GetHandled ComProc
	PutHandled ComProc
	// 3
	GetSavesInProfile ComProc
	PutSavesInProfile ComProc
}

func (i *ICoreWebView2PermissionRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionRequestedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取请求权限的 web 内容的来源。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetUri() (string, error) {
	var _uri *uint16
	r, _, _ := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(_uri)
	wapi.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// GetPermissionKind 获取请求的权限类型。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetPermissionKind() (COREWEBVIEW2_PERMISSION_KIND, error) {
	var kind COREWEBVIEW2_PERMISSION_KIND
	r, _, _ := i.Vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// GetIsUserInitiated 获取权限请求是否由用户发起。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetIsUserInitiated() (bool, error) {
	var isUserInitiated bool
	r, _, _ := i.Vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isUserInitiated)),
	)
	if r != 0 {
		return isUserInitiated, syscall.Errno(r)
	}
	return isUserInitiated, nil
}

// GetState 获取权限请求的状态。默认值为 COREWEBVIEW2_PERMISSION_STATE_DEFAULT。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetState() (COREWEBVIEW2_PERMISSION_STATE, error) {
	var state COREWEBVIEW2_PERMISSION_STATE
	r, _, _ := i.Vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&state)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return state, nil
}

// SetState 设置权限请求的状态。
func (i *ICoreWebView2PermissionRequestedEventArgs) SetState(state COREWEBVIEW2_PERMISSION_STATE) error {
	r, _, _ := i.Vtbl.PutState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(state),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取权限请求的延迟对象。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// GetICoreWebView2PermissionRequestedEventArgs2 获取 ICoreWebView2PermissionRequestedEventArgs2。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetICoreWebView2PermissionRequestedEventArgs2() (*ICoreWebView2PermissionRequestedEventArgs2, error) {
	var result *ICoreWebView2PermissionRequestedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2PermissionRequestedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2PermissionRequestedEventArgs3 获取 ICoreWebView2PermissionRequestedEventArgs3。
func (i *ICoreWebView2PermissionRequestedEventArgs) GetICoreWebView2PermissionRequestedEventArgs3() (*ICoreWebView2PermissionRequestedEventArgs3, error) {
	var result *ICoreWebView2PermissionRequestedEventArgs3
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2PermissionRequestedEventArgs3)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetUri 获取请求权限的 web 内容的来源。出错时会触发全局错误回调.
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetPermissionKind 获取请求的权限类型。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetPermissionKind() COREWEBVIEW2_PERMISSION_KIND {
	kind, err := i.GetPermissionKind()
	ReportErrorAtuo(err)
	return kind
}

// MustGetIsUserInitiated 获取权限请求是否由用户发起。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetIsUserInitiated() bool {
	isUserInitiated, err := i.GetIsUserInitiated()
	ReportErrorAtuo(err)
	return isUserInitiated
}

// MustGetState 获取权限请求的状态。默认值为 COREWEBVIEW2_PERMISSION_STATE_DEFAULT。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetState() COREWEBVIEW2_PERMISSION_STATE {
	state, err := i.GetState()
	ReportErrorAtuo(err)
	return state
}

// MustGetDeferral 获取权限请求的延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}

// MustGetICoreWebView2PermissionRequestedEventArgs2 获取 ICoreWebView2PermissionRequestedEventArgs2。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetICoreWebView2PermissionRequestedEventArgs2() *ICoreWebView2PermissionRequestedEventArgs2 {
	result, err := i.GetICoreWebView2PermissionRequestedEventArgs2()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2PermissionRequestedEventArgs3 获取 ICoreWebView2PermissionRequestedEventArgs3。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs) MustGetICoreWebView2PermissionRequestedEventArgs3() *ICoreWebView2PermissionRequestedEventArgs3 {
	result, err := i.GetICoreWebView2PermissionRequestedEventArgs3()
	ReportErrorAtuo(err)
	return result
}
