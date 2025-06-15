package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_5 是 ICoreWebView2_4 接口的延续，以支持 请求客户端证书 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_5
type ICoreWebView2_5 struct {
	Vtbl *ICoreWebView2_5Vtbl
}

type ICoreWebView2_5Vtbl struct {
	ICoreWebView2_4Vtbl
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
}

func (i *ICoreWebView2_5) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_5) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_5) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddClientCertificateRequested 添加客户端证书请求事件处理程序.
//   - 当 WebView2 向需要客户端证书进行 HTTP 身份验证的 HTTP 服务器发出请求时触发.
//   - 如果不处理该事件， WebView2 将向用户显示默认的客户端证书选择对话框提示。
func (i *ICoreWebView2_5) AddClientCertificateRequested(eventHandler *ICoreWebView2ClientCertificateRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveClientCertificateRequested 移除客户端证书请求事件处理程序
func (i *ICoreWebView2_5) RemoveClientCertificateRequested(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
