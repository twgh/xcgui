package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_14 是 ICoreWebView2_13 接口的延续，支持检测到服务器证书错误事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_14
type ICoreWebView2_14 struct {
	ICoreWebView2_13
}

// AddServerCertificateErrorDetected 添加检测到服务器证书错误事件处理程序。
//   - 当 WebView2 在加载网页时无法验证服务器的数字证书时触发。
//   - 此事件将针对所有网络资源触发，并紧随 WebResourceRequested 事件之后。
//   - 如果不处理该事件，WebView2 会在导航时向用户显示默认的 TLS 插页式错误页面，而对于非导航操作，Web 请求会被取消。
func (i *ICoreWebView2_14) AddServerCertificateErrorDetected(eventHandler *ICoreWebView2ServerCertificateErrorDetectedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveServerCertificateErrorDetected 移除检测到服务器证书错误事件处理程序。
func (i *ICoreWebView2_14) RemoveServerCertificateErrorDetected(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearServerCertificateErrorActions 清除所有共享同一会话的 WebView2 中来自 ServerCertificateErrorDetected 事件的所有缓存的继续处理 TLS 证书错误的决定。
func (i *ICoreWebView2_14) ClearServerCertificateErrorActions(handler *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) error {
	r, _, _ := i.Vtbl.ClearServerCertificateErrorActions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearServerCertificateErrorActionsEx 清除所有共享同一会话的 WebView2 中来自 ServerCertificateErrorDetected 事件的所有缓存的继续处理 TLS 证书错误的决定。
//
// impl: *WebViewEventImpl。
//
// cb: 清除操作完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2_14) ClearServerCertificateErrorActionsEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "ClearServerCertificateErrorActionsCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "ClearServerCertificateErrorActionsCompleted")
	return i.ClearServerCertificateErrorActions((*ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler)(handler))
}
