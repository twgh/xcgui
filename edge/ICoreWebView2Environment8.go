package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment8 是 ICoreWebView2Environment7 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment8
type ICoreWebView2Environment8 struct {
	ICoreWebView2Environment7
}

// AddProcessInfosChanged 添加进程信息变更事件处理程序。
//   - 当 WebView2 运行时的进程列表发生变化时触发。
func (i *ICoreWebView2Environment8) AddProcessInfosChanged(eventHandler *ICoreWebView2ProcessInfosChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveProcessInfosChanged 移除进程信息变更事件处理程序。
func (i *ICoreWebView2Environment8) RemoveProcessInfosChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProcessInfos 获取当前 WebView2 运行时的进程信息集合，提供除崩溃报告进程外使用相同用户数据文件夹的所有进程的列表。
func (i *ICoreWebView2Environment8) GetProcessInfos() (*ICoreWebView2ProcessInfoCollection, error) {
	var value *ICoreWebView2ProcessInfoCollection
	r, _, _ := i.Vtbl.GetProcessInfos.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetProcessInfos 获取当前 WebView2 运行时的进程信息集合，提供除崩溃报告进程外使用相同用户数据文件夹的所有进程的列表。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2Environment8) MustGetProcessInfos() *ICoreWebView2ProcessInfoCollection {
	value, err := i.GetProcessInfos()
	ReportErrorAtuo(err)
	return value
}
