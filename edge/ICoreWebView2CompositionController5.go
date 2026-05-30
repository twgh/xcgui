package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2CompositionController5 是 ICoreWebView2CompositionController4 的延续，用于支持拖拽开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller5
type ICoreWebView2CompositionController5 struct {
	ICoreWebView2CompositionController4
}

// Event_DragStarting 拖拽开始事件. 返回回调函数 ID.
//   - 当用户开始拖拽操作时触发。
//
// impl: *WebViewEventImpl.
//
// cb: 回调函数.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 不填默认为 true.
func (i *ICoreWebView2CompositionController5) Event_DragStarting(impl *WebViewEventImpl, cb func(sender *ICoreWebView2CompositionController, args *ICoreWebView2DragStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventBus.AddCallback(impl, "DragStarting", cb, i, allowAddingMultiple...)
}

// AddDragStarting 添加拖拽开始事件处理程序.
func (i *ICoreWebView2CompositionController5) AddDragStarting(eventHandler *ICoreWebView2DragStartingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddDragStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDragStarting 移除拖拽开始事件处理程序.
func (i *ICoreWebView2CompositionController5) RemoveDragStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveDragStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
