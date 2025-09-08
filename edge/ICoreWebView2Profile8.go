package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile8 是 ICoreWebView2Profile7 的延续，提供了配置文件删除功能和删除事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile8
type ICoreWebView2Profile8 struct {
	ICoreWebView2Profile7
}

// Delete 将此配置文件标记为待删除。
//   - 本地配置文件的目录将在浏览器进程退出时被删除。
//   - 如果删除失败（因为其他进程打开了相关文件），WebView2 将在未来所有浏览器进程启动时尝试删除该配置文件，直到成功为止。
//   - 相应的 ICoreWebView2 将被关闭，并且会触发 ICoreWebView2Profile.Deleted 事件。
//   - 如果您尝试创建一个与已标记为删除但尚未删除的现有配置文件同名的新配置文件，创建操作将失败，并返回 ERROR_DELETE_PENDING。
func (i *ICoreWebView2Profile8) Delete() error {
	r, _, _ := i.Vtbl.Delete.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Event_ProfileDeleted 配置文件删除事件。
//   - 当配置文件被标记为删除时触发。
//   - 触发此事件时，ICoreWebView2Profile 及其对应的 ICoreWebView2 已关闭，且无法再使用。
func (i *ICoreWebView2Profile8) Event_ProfileDeleted(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Profile, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "ProfileDeleted", cb, i, allowAddingMultiple...)
}

// AddDeleted 添加配置文件删除事件的处理程序。
//   - 当配置文件被标记为删除时触发。
//   - 触发此事件时，ICoreWebView2Profile 及其对应的 ICoreWebView2 已关闭，且无法再使用。
func (i *ICoreWebView2Profile8) AddDeleted(eventHandler *ICoreWebView2ProfileDeletedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddDeleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDeleted 移除配置文件删除事件的处理程序。
//
// token: 之前通过 AddDeleted 获得的事件注册令牌
func (i *ICoreWebView2Profile8) RemoveDeleted(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveDeleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
