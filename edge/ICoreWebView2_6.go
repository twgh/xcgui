package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_6 是 ICoreWebView2_5 接口的延续，用于管理打开浏览器任务管理器窗口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_6
type ICoreWebView2_6 struct {
	ICoreWebView2_5
}

// OpenTaskManagerWindow 打开浏览器任务管理器窗口。
func (i *ICoreWebView2_6) OpenTaskManagerWindow() error {
	r, _, _ := i.Vtbl.OpenTaskManagerWindow.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
