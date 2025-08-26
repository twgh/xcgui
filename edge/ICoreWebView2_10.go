package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_10 是 ICoreWebView2_9 接口的延续，支持请求基本身份验证事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_10
type ICoreWebView2_10 struct {
	ICoreWebView2_9
}

// AddBasicAuthenticationRequested 添加基本身份验证请求事件处理程序。
//   - 主机可以提供包含身份验证凭据的响应，也可以取消请求。
//   - 如果主机将 Cancel 属性设置为 false，但未在 Response 属性上提供 UserName 或 Password 属性，那么 WebView2 将向用户显示默认的身份验证质询对话框提示。
func (i *ICoreWebView2_10) AddBasicAuthenticationRequested(handler *ICoreWebView2BasicAuthenticationRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddBasicAuthenticationRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveBasicAuthenticationRequested 移除基本身份验证请求事件处理程序。
func (i *ICoreWebView2_10) RemoveBasicAuthenticationRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveBasicAuthenticationRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
