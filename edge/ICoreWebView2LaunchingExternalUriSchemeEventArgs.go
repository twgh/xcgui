package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2LaunchingExternalUriSchemeEventArgs 是启动外部URI方案事件的参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs
type ICoreWebView2LaunchingExternalUriSchemeEventArgs struct {
	Vtbl *ICoreWebView2LaunchingExternalUriSchemeEventArgsVtbl
}

type ICoreWebView2LaunchingExternalUriSchemeEventArgsVtbl struct {
	IUnknownVtbl
	GetUri              ComProc
	GetInitiatingOrigin ComProc
	GetIsUserInitiated  ComProc
	GetCancel           ComProc
	PutCancel           ComProc
	GetDeferral         ComProc
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取要启动的 URI。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetUri() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return uri, nil
}

// GetInitiatingOrigin 获取发起导航到 URI 的源。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetInitiatingOrigin() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetInitiatingOrigin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	origin := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return origin, nil
}

// GetIsUserInitiated 获取外部 URI 方案请求是否通过用户操作发起。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetIsUserInitiated() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetCancel 获取是否取消启动外部 URI 方案。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetCancel() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetCancel 设置是否取消启动外部 URI 方案。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) SetCancel(value bool) error {
	r, _, _ := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取延迟对象，使用此操作可在稍后完成该事件。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var value *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetUri 获取要启动的 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetInitiatingOrigin 获取发起导航到 URI 的源。出错时会触发全局错误回调。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) MustGetInitiatingOrigin() string {
	origin, err := i.GetInitiatingOrigin()
	ReportErrorAtuo(err)
	return origin
}

// MustGetIsUserInitiated 获取外部 URI 方案请求是否通过用户操作发起。出错时会触发全局错误回调。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) MustGetIsUserInitiated() bool {
	result, err := i.GetIsUserInitiated()
	ReportErrorAtuo(err)
	return result
}

// MustGetCancel 获取是否取消启动外部 URI 方案。出错时会触发全局错误回调。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) MustGetCancel() bool {
	result, err := i.GetCancel()
	ReportErrorAtuo(err)
	return result
}

// MustGetDeferral 获取延迟对象，使用此操作可在稍后完成该事件。出错时会触发全局错误回调。
func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}
