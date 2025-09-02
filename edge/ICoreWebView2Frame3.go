package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame3 是 ICoreWebView2Frame2 的延续, 支持框架权限请求事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame3
type ICoreWebView2Frame3 struct {
	ICoreWebView2Frame2
}

// Event_FramePermissionRequested 框架权限请求事件.
//   - 当 iframe 及其任何子 iframe 中的内容请求访问特权资源时触发。
//   - 这与 PermissionRequested 事件以及 ICoreWebView2 有关。在 iframe 请求权限的情况下，这两个事件都会被触发。ICoreWebView2Frame的事件处理程序将在 ICoreWebView2 的事件处理程序之前被调用。
//   - 如果在 ICoreWebView2Frame 事件处理程序中，PermissionRequestedEventArgs 的 Handled 属性被设置为 TRUE，那么该事件将不会在 ICoreWebView2 上触发，其事件处理程序也不会被调用。
//   - 在嵌套 iframe 的情况下，如果在当前嵌套 iframe 中处理了 PermissionRequested 事件（即 PermissionRequestedEventArgs 的 Handled 属性设置为 TRUE），则不会在父级 ICoreWebView2Frame 上引发该事件。但是，如果该嵌套 iframe 中未处理 PermissionRequested 事件，该事件将从其最近的已跟踪父级 ICoreWebView2Frame 引发。它会遍历父框架链直至主框架，直到某个父框架处理该请求为止。
func (i *ICoreWebView2Frame3) Event_FramePermissionRequested(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FramePermissionRequested", cb, i, allowAddingMultiple...)
}

// AddPermissionRequested 为权限请求事件添加处理程序。
//   - 当 iframe 及其任何子 iframe 中的内容请求访问特权资源时触发。
//   - 这与 PermissionRequested 事件以及 ICoreWebView2 有关。在 iframe 请求权限的情况下，这两个事件都会被触发。ICoreWebView2Frame的事件处理程序将在 ICoreWebView2 的事件处理程序之前被调用。
//   - 如果在 ICoreWebView2Frame 事件处理程序中，PermissionRequestedEventArgs 的 Handled 属性被设置为 TRUE，那么该事件将不会在 ICoreWebView2 上触发，其事件处理程序也不会被调用。
//   - 在嵌套 iframe 的情况下，如果在当前嵌套 iframe 中处理了 PermissionRequested 事件（即 PermissionRequestedEventArgs 的 Handled 属性设置为 TRUE），则不会在父级 ICoreWebView2Frame 上引发该事件。但是，如果该嵌套 iframe 中未处理 PermissionRequested 事件，该事件将从其最近的已跟踪父级 ICoreWebView2Frame 引发。它会遍历父框架链直至主框架，直到某个父框架处理该请求为止。
func (i *ICoreWebView2Frame3) AddPermissionRequested(eventHandler *ICoreWebView2FramePermissionRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemovePermissionRequested 移除 PermissionRequested 事件处理程序。
func (i *ICoreWebView2Frame3) RemovePermissionRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemovePermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
