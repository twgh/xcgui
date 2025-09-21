package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_16 是 ICoreWebView2_15 接口的延续，支持打印相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_16
type ICoreWebView2_16 struct {
	ICoreWebView2_15
}

// Print 使用提供的设置，将当前网页异步打印到指定打印机。
//   - 一次只能有一个打印操作在进行中。
//
// printSettings: 打印设置，用于配置打印作业, 可为nil。
//
// handler: 打印操作完成后的回调处理程序。
func (i *ICoreWebView2_16) Print(printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintCompletedHandler) error {
	r, _, _ := i.Vtbl.Print.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PrintEx 使用提供的设置，将当前网页异步打印到指定打印机，并通过回调函数返回结果。
//
// impl: *WebViewEventImpl。
//
// printSettings: 打印设置，用于配置打印作业, 可为nil。
//
// cb: 打印操作完成后的回调函数，可以为 nil。
func (i *ICoreWebView2_16) PrintEx(impl *WebViewEventImpl, printSettings *ICoreWebView2PrintSettings, cb func(errorCode syscall.Errno, result COREWEBVIEW2_PRINT_STATUS) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "PrintCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "PrintCompleted")
	return i.Print(printSettings, (*ICoreWebView2PrintCompletedHandler)(handler))
}

// ShowPrintUI 显示打印对话框。
//   - 如果浏览器或系统的打印对话框已打开，再次调用时不会打开新的打印对话框。
//
// printDialogKind: 指定要显示的打印对话框类型，可以是浏览器的打印对话框或系统的打印对话框。
func (i *ICoreWebView2_16) ShowPrintUI(printDialogKind COREWEBVIEW2_PRINT_DIALOG_KIND) error {
	r, _, _ := i.Vtbl.ShowPrintUI.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(printDialogKind),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PrintToPdfStream 根据提供的设置，异步提供当前网页的 PDF 数据。
//   - 一次只能有一个打印操作在进行中。
//
// printSettings: 打印设置，用于配置打印作业, 可为nil。
//
// handler: 操作完成后的回调处理程序。
func (i *ICoreWebView2_16) PrintToPdfStream(printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintToPdfStreamCompletedHandler) error {
	r, _, _ := i.Vtbl.PrintToPdfStream.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PrintToPdfStreamEx 根据提供的设置，异步提供当前网页的 PDF 数据，并通过回调函数返回结果。
//
// impl: *WebViewEventImpl。
//
// printSettings: 打印设置，用于配置打印作业, 可为nil。
//
// cb: 操作完成后的回调函数，接收返回结果。
func (i *ICoreWebView2_16) PrintToPdfStreamEx(impl *WebViewEventImpl, printSettings *ICoreWebView2PrintSettings, cb func(errorCode syscall.Errno, pdfBytes []byte) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "PrintToPdfStreamCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "PrintToPdfStreamCompleted")
	return i.PrintToPdfStream(printSettings, (*ICoreWebView2PrintToPdfStreamCompletedHandler)(handler))
}
