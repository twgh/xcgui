package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_21 是 ICoreWebView2_20 接口的延续，提供了带结果的脚本执行功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_21
type ICoreWebView2_21 struct {
	ICoreWebView2_20
}

// ExecuteScriptWithResult 异步执行 JavaScript 代码并返回执行结果。
//
// javaScript: 要执行的 JavaScript 代码。
//
// handler: 执行完成后的回调处理程序，接收执行结果或错误信息。
func (i *ICoreWebView2_21) ExecuteScriptWithResult(javaScript string, handler *ICoreWebView2ExecuteScriptWithResultCompletedHandler) error {
	_javaScript, err := syscall.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.ExecuteScriptWithResult.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ExecuteScriptWithResultEx 异步执行 JavaScript 代码并返回执行结果。
//
// impl: *WebViewEventImpl。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 执行完成后的回调处理程序，接收执行结果或错误信息。
func (i *ICoreWebView2_21) ExecuteScriptWithResultEx(impl *WebViewEventImpl, javaScript string, cb func(errorCode syscall.Errno, result *ICoreWebView2ExecuteScriptResult) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "ExecuteScriptWithResultCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "ExecuteScriptWithResultCompleted")
	return i.ExecuteScriptWithResult(javaScript, (*ICoreWebView2ExecuteScriptWithResultCompletedHandler)(handler))
}
