package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_2 是 ICoreWebView2 接口的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_2
type ICoreWebView2_2 struct {
	ICoreWebView2
}

// GetEnvironment 获取用于创建此 ICoreWebView2 的 ICoreWebView2Environment.
func (i *ICoreWebView2_2) GetEnvironment() (*ICoreWebView2Environment, error) {
	var environment *ICoreWebView2Environment
	r, _, _ := i.Vtbl.GetEnvironment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&environment)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return environment, nil
}

// NavigateWithWebResourceRequest 使用构造的 ICoreWebView2WebResourceRequest 对象进行导航。
func (i *ICoreWebView2_2) NavigateWithWebResourceRequest(request *ICoreWebView2WebResourceRequest) error {
	r, _, _ := i.Vtbl.NavigateWithWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(request)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCookieManager 获取与此 ICoreWebView2 关联的 cookie 管理器对象。
func (i *ICoreWebView2_2) GetCookieManager() (*ICoreWebView2CookieManager, error) {
	var cookieManager *ICoreWebView2CookieManager
	r, _, _ := i.Vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cookieManager)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cookieManager, nil
}

// AddWebResourceResponseReceived 添加 WebResourceResponseReceived 事件处理程序。
//   - 当 WebView 接收到对网络资源请求的响应时，会引发 WebResourceResponseReceived 事件（WebView执行的任何URI解析；例如HTTP/HTTPS、来自重定向、导航、HTML声明、隐式 favicon 图标的查找和数据请求，以及文档中 fetch API 的使用）。
//   - 宿主应用可以使用此事件查看网络资源的实际请求和响应。
//   - 无法保证 WebView 处理响应的顺序与宿主应用的处理程序运行的顺序。
//   - 应用的处理程序不会阻止 WebView 处理响应。
func (i *ICoreWebView2_2) AddWebResourceResponseReceived(handler *ICoreWebView2WebResourceResponseReceivedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebResourceResponseReceived 移除 WebResourceResponseReceived 事件处理程序。
func (i *ICoreWebView2_2) RemoveWebResourceResponseReceived(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddDomContentLoaded 添加 DomContentLoaded 事件处理程序。
//   - 当初始 HTML 文档解析完成时，会触发 DOMContentLoaded 事件。这与 HTML 中文档的 DOMContentLoaded 事件一致。
func (i *ICoreWebView2_2) AddDomContentLoaded(handler *ICoreWebView2DOMContentLoadedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddDomContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDomContentLoaded 移除 DomContentLoaded 事件处理程序。
func (i *ICoreWebView2_2) RemoveDomContentLoaded(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveDomContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetEnvironment 获取用于创建此 ICoreWebView2 的 ICoreWebView2Environment. 出错时会触发全局错误回调.
func (i *ICoreWebView2_2) MustGetEnvironment() *ICoreWebView2Environment {
	environment, err := i.GetEnvironment()
	ReportErrorAuto(err)
	return environment
}

// MustGetCookieManager 获取与此 ICoreWebView2 关联的 cookie 管理器对象。出错时会触发全局错误回调。
func (i *ICoreWebView2_2) MustGetCookieManager() *ICoreWebView2CookieManager {
	cookieManager, err := i.GetCookieManager()
	ReportErrorAuto(err)
	return cookieManager
}
