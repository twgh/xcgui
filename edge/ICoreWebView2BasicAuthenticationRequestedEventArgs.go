package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2BasicAuthenticationRequestedEventArgs 提供有关基本身份验证请求的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs
type ICoreWebView2BasicAuthenticationRequestedEventArgs struct {
	Vtbl *_ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl
}

type _ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetUri       ComProc
	GetChallenge ComProc
	GetResponse  ComProc
	GetCancel    ComProc
	PutCancel    ComProc
	GetDeferral  ComProc
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取需要基本身份验证的URI。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetUri() (string, error) {
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

// GetChallenge 获取身份验证质询字符串。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetChallenge() (string, error) {
	var _challenge *uint16
	r, _, _ := i.Vtbl.GetChallenge.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_challenge)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	challenge := common.UTF16PtrToString(_challenge)
	wapi.CoTaskMemFree(unsafe.Pointer(_challenge))
	return challenge, nil
}

// GetResponse 获取用于响应基本身份验证请求的响应对象。
//   - 如果主机希望提供身份验证凭据，此对象将由应用程序填充。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetResponse() (*ICoreWebView2BasicAuthenticationResponse, error) {
	var response *ICoreWebView2BasicAuthenticationResponse
	r, _, _ := i.Vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// GetCancel 获取是否取消基本身份验证请求。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetCancel() (bool, error) {
	var cancel bool
	r, _, _ := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return cancel, nil
}

// SetCancel 设置是否取消基本身份验证请求。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) SetCancel(cancel bool) error {
	r, _, _ := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(cancel),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取延迟对象，使用此延迟来推迟显示基本身份验证对话框的决定。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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

// MustGetUri 获取需要基本身份验证的URI。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) MustGetUri() string {
	result, err := i.GetUri()
	ReportErrorAtuo(err)
	return result
}

// MustGetChallenge 获取身份验证质询字符串。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) MustGetChallenge() string {
	result, err := i.GetChallenge()
	ReportErrorAtuo(err)
	return result
}

// MustGetResponse 获取用于响应基本身份验证请求的响应对象。出错时会触发全局错误回调。
//   - 如果主机希望提供身份验证凭据，此对象将由应用程序填充。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) MustGetResponse() *ICoreWebView2BasicAuthenticationResponse {
	result, err := i.GetResponse()
	ReportErrorAtuo(err)
	return result
}

// MustGetCancel 获取是否取消基本身份验证请求。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) MustGetCancel() bool {
	result, err := i.GetCancel()
	ReportErrorAtuo(err)
	return result
}

// MustGetDeferral 获取延迟对象，用于异步处理基本身份验证请求。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	result, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return result
}
