package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_5 是 ICoreWebView2_4 接口的延续，以支持 请求客户端证书 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_5
type ICoreWebView2_5 struct {
	ICoreWebView2_4
}

// AddClientCertificateRequested 添加客户端证书请求事件处理程序.
//   - 当 WebView2 向需要客户端证书进行 HTTP 身份验证的 HTTP 服务器发出请求时触发.
//   - 如果不处理该事件， WebView2 将向用户显示默认的客户端证书选择对话框提示。
func (i *ICoreWebView2_5) AddClientCertificateRequested(eventHandler *ICoreWebView2ClientCertificateRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveClientCertificateRequested 移除客户端证书请求事件处理程序
func (i *ICoreWebView2_5) RemoveClientCertificateRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
