package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame7 是 ICoreWebView2Frame6 的延续，添加了对框架创建事件的支持。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame7
type ICoreWebView2Frame7 struct {
	ICoreWebView2Frame6
}

// Event_FrameChildFrameCreated 框架创建事件.
//   - 当新的直接子级 iframe 创建时触发。
//   - 处理此事件可获取 ICoreWebView2Frame 对象。使用 ICoreWebView2Frame 的销毁事件来监听该 iframe 何时消失。
func (i *ICoreWebView2Frame7) Event_FrameChildFrameCreated(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2FrameCreatedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameChildFrameCreated", cb, i, allowAddingMultiple...)
}

// AddFrameCreated 添加框架创建事件处理程序。
//   - 当新的直接子级 iframe 创建时触发。
//   - 处理此事件可获取 ICoreWebView2Frame 对象。使用 ICoreWebView2Frame 的销毁事件来监听该 iframe 何时消失。
func (i *ICoreWebView2Frame7) AddFrameCreated(eventHandler *ICoreWebView2FrameChildFrameCreatedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveFrameCreated 移除框架创建事件处理程序。
func (i *ICoreWebView2Frame7) RemoveFrameCreated(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
