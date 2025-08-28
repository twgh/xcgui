package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_24 是 ICoreWebView2_23 接口的延续，支持管理 WebView2 网页通知功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_24
type ICoreWebView2_24 struct {
	ICoreWebView2_23
}

// AddNotificationReceived 添加通知接收事件处理程序。
//   - 如果未对事件参数执行延迟操作，那么在 DOM 通知创建调用（即 Notification()）之后的后续脚本会被阻塞，直到事件处理程序返回。
//   - 如果执行了延迟操作，脚本会被阻塞，直到延迟完成。
func (i *ICoreWebView2_24) AddNotificationReceived(eventHandler *ICoreWebView2NotificationReceivedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNotificationReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNotificationReceived 移除通知接收事件处理程序。
func (i *ICoreWebView2_24) RemoveNotificationReceived(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNotificationReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
