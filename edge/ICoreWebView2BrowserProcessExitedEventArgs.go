package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2BrowserProcessExitedEventArgs 提供有关浏览器进程退出的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs
type ICoreWebView2BrowserProcessExitedEventArgs struct {
	Vtbl *ICoreWebView2BrowserProcessExitedEventArgsVtbl
}

type ICoreWebView2BrowserProcessExitedEventArgsVtbl struct {
	IUnknownVtbl
	GetBrowserProcessExitKind ComProc
	GetBrowserProcessId       ComProc
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetBrowserProcessExitKind 获取浏览器进程退出的类型。
func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessExitKind() (COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND, error) {
	var value COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND
	r, _, _ := i.Vtbl.GetBrowserProcessExitKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetBrowserProcessId 获取退出的浏览器进程的进程ID。
func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessId() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetBrowserProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetBrowserProcessExitKind 获取浏览器进程退出的类型。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2BrowserProcessExitedEventArgs) MustGetBrowserProcessExitKind() COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND {
	value, err := i.GetBrowserProcessExitKind()
	ReportErrorAuto(err)
	return value
}

// MustGetBrowserProcessId 获取退出的浏览器进程的进程ID。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2BrowserProcessExitedEventArgs) MustGetBrowserProcessId() uint32 {
	value, err := i.GetBrowserProcessId()
	ReportErrorAuto(err)
	return value
}
