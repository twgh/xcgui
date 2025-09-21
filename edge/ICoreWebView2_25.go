package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_25 是 ICoreWebView2_24 接口的延续，支持"另存为"UI相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_25
type ICoreWebView2_25 struct {
	ICoreWebView2_24
}

// AddSaveAsUIShowing 添加"另存为"UI显示事件处理程序。
//   - 当通过编程方式或手动方式触发“另存为”时，会引发此事件。
func (i *ICoreWebView2_25) AddSaveAsUIShowing(eventHandler *ICoreWebView2SaveAsUIShowingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddSaveAsUIShowing.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveSaveAsUIShowing 移除"另存为"UI显示事件处理程序。
func (i *ICoreWebView2_25) RemoveSaveAsUIShowing(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveSaveAsUIShowing.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ShowSaveAsUI 显示"另存为"UI，允许用户保存当前页面。
//   - 默认情况下会打开系统模态对话框。如果 SuppressDefaultDialog 属性为 TRUE，则不会打开系统对话框。
func (i *ICoreWebView2_25) ShowSaveAsUI(handler *ICoreWebView2ShowSaveAsUICompletedHandler) error {
	r, _, _ := i.Vtbl.ShowSaveAsUI.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ShowSaveAsUIEx 显示"另存为"UI，允许用户保存当前页面。
//
// impl: *WebViewEventImpl。
//
// cb: 执行完成后的回调处理程序，接收操作结果。
func (i *ICoreWebView2_25) ShowSaveAsUIEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, result COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "ShowSaveAsUICompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "ShowSaveAsUICompleted")
	return i.ShowSaveAsUI((*ICoreWebView2ShowSaveAsUICompletedHandler)(handler))
}
