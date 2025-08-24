package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment5 是 ICoreWebView2Environment4 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment5
type ICoreWebView2Environment5 struct {
	ICoreWebView2Environment4
}

// AddBrowserProcessExited 添加浏览器进程退出事件处理程序。
//   - 当与WebView2关联的浏览器进程意外终止时触发此事件。
func (i *ICoreWebView2Environment5) AddBrowserProcessExited(eventHandler *ICoreWebView2BrowserProcessExitedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddBrowserProcessExited.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveBrowserProcessExited 移除浏览器进程退出事件处理程序。
func (i *ICoreWebView2Environment5) RemoveBrowserProcessExited(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveBrowserProcessExited.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
