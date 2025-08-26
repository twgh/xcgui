package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_7 是 ICoreWebView2_6 接口的延续，支持打印为PDF。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_7
type ICoreWebView2_7 struct {
	ICoreWebView2_6
}

// PrintToPdf 异步将当前页面打印为 PDF 文件。
//
// resultFilePath: 保存 PDF 文件的路径。
//
// printSettings: 用于打印的设置。
//
// handler: 打印完成后的回调处理程序。
func (i *ICoreWebView2_7) PrintToPdf(resultFilePath string, printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintToPdfCompletedHandler) error {
	_resultFilePath, err := syscall.UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PrintToPdf.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultFilePath)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PrintToPdfEx 异步将当前页面打印为 PDF 文件。
//
// impl: *WebViewEventImpl.
//
// resultFilePath: 保存 PDF 文件的路径。
//
// printSettings: 用于打印的设置。
//
// cb: 打印完成后的回调函数。
func (i *ICoreWebView2_7) PrintToPdfEx(impl *WebViewEventImpl, resultFilePath string, printSettings *ICoreWebView2PrintSettings, cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "PrintToPdfCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "PrintToPdfCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "PrintToPdfCompleted")
	}
	return i.PrintToPdf(resultFilePath, printSettings, (*ICoreWebView2PrintToPdfCompletedHandler)(handler))
}
