package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment13 提供获取进程扩展信息的功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment13
type ICoreWebView2Environment13 struct {
	ICoreWebView2Environment12
}

// GetProcessExtendedInfos 获取与该 CoreWebView2Environment 关联的所有当前正在运行的进程（不包括崩溃报告进程）对应的 ProcessExtendedInfo 的快照集合。
//   - 这提供了与 ProcessInfo 中相同的 GetProcessInfos 列表，但另外还提供了渲染器进程中正在主动运行（显示或隐藏用户界面元素）的相关 FrameInfo 列表。有关更多信息，请参阅 AssociatedFrameInfos。
//
// handler: 获取进程扩展信息完成后的回调处理程序。
func (i *ICoreWebView2Environment13) GetProcessExtendedInfos(handler *ICoreWebView2GetProcessExtendedInfosCompletedHandler) error {
	r, _, _ := i.Vtbl.GetProcessExtendedInfos.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProcessExtendedInfosEx 获取与该 CoreWebView2Environment 关联的所有当前正在运行的进程（不包括崩溃报告进程）对应的 ProcessExtendedInfo 的快照集合。
//   - 这提供了与 ProcessInfo 中相同的 GetProcessInfos 列表，但另外还提供了渲染器进程中正在主动运行（显示或隐藏用户界面元素）的相关 FrameInfo 列表。有关更多信息，请参阅 AssociatedFrameInfos。
func (i *ICoreWebView2Environment13) GetProcessExtendedInfosEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, result *ICoreWebView2ProcessExtendedInfoCollection) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "GetProcessExtendedInfosCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "GetProcessExtendedInfosCompleted")
	return i.GetProcessExtendedInfos((*ICoreWebView2GetProcessExtendedInfosCompletedHandler)(handler))
}
