package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_27 是 ICoreWebView2_26 接口的延续，支持屏幕捕获开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_27
type ICoreWebView2_27 struct {
	ICoreWebView2_26
}

// AddScreenCaptureStarting 添加屏幕截图开始事件处理程序。
func (i *ICoreWebView2_27) AddScreenCaptureStarting(eventHandler *ICoreWebView2ScreenCaptureStartingEventHandler, token *EventRegistrationToken) error {
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
func (i *ICoreWebView2_27) RemoveScreenCaptureStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveScreenCaptureStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
