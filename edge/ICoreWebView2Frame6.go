package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame6 是 ICoreWebView2Frame5 的延续，添加了对框架屏幕截图开始事件的支持。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame6
type ICoreWebView2Frame6 struct {
	ICoreWebView2Frame5
}

// Event_FrameScreenCaptureStarting 框架屏幕截图开始事件.
//   - 当 iframe 或其任何子 iframe 中的内容通过 getDisplayMedia() 请求使用屏幕捕获 API 的权限时触发.
//   - 这与 ICoreWebView2 上的 ScreenCaptureStarting 事件相关。在 iframe 请求屏幕捕获的情况下，这两个事件都会被触发。ICoreWebView2 上的事件处理程序之前，会先调用 ICoreWebView2Frame 的事件处理程序。
//   - 如果在 ICoreWebView2Frame 事件处理程序中将 ScreenCaptureStartingEventArgs 的 Handled 属性设置为 TRUE，那么 ICoreWebView2 上将不会触发该事件，其事件处理程序也不会被调用。
func (i *ICoreWebView2Frame6) Event_FrameScreenCaptureStarting(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameScreenCaptureStarting", cb, i, allowAddingMultiple...)
}

// AddScreenCaptureStarting 添加屏幕截图开始事件处理程序。
//   - 当 iframe 或其任何子 iframe 中的内容通过 getDisplayMedia() 请求使用屏幕捕获 API 的权限时触发.
//   - 这与 ICoreWebView2 上的 ScreenCaptureStarting 事件相关。在 iframe 请求屏幕捕获的情况下，这两个事件都会被触发。ICoreWebView2 上的事件处理程序之前，会先调用 ICoreWebView2Frame 的事件处理程序。
//   - 如果在 ICoreWebView2Frame 事件处理程序中将 ScreenCaptureStartingEventArgs 的 Handled 属性设置为 TRUE，那么 ICoreWebView2 上将不会触发该事件，其事件处理程序也不会被调用。
func (i *ICoreWebView2Frame6) AddScreenCaptureStarting(eventHandler *ICoreWebView2FrameScreenCaptureStartingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddScreenCaptureStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveScreenCaptureStarting 移除屏幕截图开始事件处理程序。
func (i *ICoreWebView2Frame6) RemoveScreenCaptureStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveScreenCaptureStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
