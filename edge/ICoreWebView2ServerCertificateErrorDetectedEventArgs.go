package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ServerCertificateErrorDetectedEventArgs 提供有关服务器证书错误的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs
type ICoreWebView2ServerCertificateErrorDetectedEventArgs struct {
	Vtbl *ICoreWebView2ServerCertificateErrorDetectedEventArgsVtbl
}

type ICoreWebView2ServerCertificateErrorDetectedEventArgsVtbl struct {
	IUnknownVtbl
	GetErrorStatus       ComProc
	GetRequestUri        ComProc
	GetServerCertificate ComProc
	GetAction            ComProc
	PutAction            ComProc
	GetDeferral          ComProc
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetErrorStatus 获取无效证书的 TLS 错误代码。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetErrorStatus() (COREWEBVIEW2_WEB_ERROR_STATUS, error) {
	var _value COREWEBVIEW2_WEB_ERROR_STATUS
	r, _, _ := i.Vtbl.GetErrorStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// GetRequestUri 获取导致证书错误的请求URI。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetRequestUri() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetRequestUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetServerCertificate 获取有问题的服务器证书。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetServerCertificate() (*ICoreWebView2Certificate, error) {
	var _value *ICoreWebView2Certificate
	r, _, _ := i.Vtbl.GetServerCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return _value, nil
}

// GetAction 获取当前设置的处理证书错误的操作。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetAction() (COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION, error) {
	var _value COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION
	r, _, _ := i.Vtbl.GetAction.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// SetAction 设置处理证书错误的操作。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) SetAction(value COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION) error {
	r, _, _ := i.Vtbl.PutAction.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取一个延迟对象，使用此操作可在稍后完成该事件。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var _deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return _deferral, nil
}

// MustGetErrorStatus 获取无效证书的 TLS 错误代码。出错时会触发全局错误回调。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) MustGetErrorStatus() COREWEBVIEW2_WEB_ERROR_STATUS {
	result, err := i.GetErrorStatus()
	ReportErrorAuto(err)
	return result
}

// MustGetRequestUri 获取导致证书错误的请求URI。出错时会触发全局错误回调。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) MustGetRequestUri() string {
	result, err := i.GetRequestUri()
	ReportErrorAuto(err)
	return result
}

// MustGetServerCertificate 获取有问题的服务器证书。出错时会触发全局错误回调。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) MustGetServerCertificate() *ICoreWebView2Certificate {
	result, err := i.GetServerCertificate()
	ReportErrorAuto(err)
	return result
}

// MustGetAction 获取当前设置的处理证书错误的操作。出错时会触发全局错误回调。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) MustGetAction() COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION {
	result, err := i.GetAction()
	ReportErrorAuto(err)
	return result
}

// MustGetDeferral 获取一个延迟对象，使用此操作可在稍后完成该事件。出错时会触发全局错误回调。
func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	result, err := i.GetDeferral()
	ReportErrorAuto(err)
	return result
}
