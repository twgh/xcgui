package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_18 是 ICoreWebView2_17 接口的延续，它管理对在操作系统中注册的 URI 方案的导航请求。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_18
type ICoreWebView2_18 struct {
	ICoreWebView2_17
}

// AddLaunchingExternalUriScheme 添加启动外部 URI 方案事件的处理程序。
//   - 当导航请求指向已在操作系统中注册的 URI 方案时触发。
//   - 事件处理程序可以抑制默认对话框，或者用自定义对话框替换默认对话框。
//   - 如果未对事件参数执行延迟操作，外部 URI 方案的启动将被阻止，直到事件处理程序返回。
//   - 如果执行了延迟操作，外部 URI 方案的启动将被阻止，直到延迟完成。主机还可以选择取消 URI 方案的启动。
//   - 无论 Cancel 属性设置为 TRUE 还是 FALSE，都会触发 NavigationStarting 和 NavigationCompleted 事件。
//   - 无论主机是否在 ICoreWebView2LaunchingExternalUriSchemeEventArgs 上设置 Cancel 属性，NavigationCompleted 事件触发时，IsSuccess 属性都将设置为 FALSE，WebErrorStatus 属性都将设置为 ConnectionAborted。
//   - 对于此次导航到外部 URI 方案，无论 Cancel 属性如何设置，都不会触发 SourceChanged、ContentLoading 和 HistoryChanged 事件。
//   - LaunchingExternalUriScheme 事件将在 NavigationStarting 事件之后、NavigationCompleted 事件之前触发。
//   - 导航到外部 URI 方案时，默认的 ICoreWebView2Settings 也会更新。如果 ICoreWebView2Settings 接口上的某个设置已更改，导航到外部 URI 方案将触发 ICoreWebView2Settings 更新。
func (i *ICoreWebView2_18) AddLaunchingExternalUriScheme(eventHandler *ICoreWebView2LaunchingExternalUriSchemeEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveLaunchingExternalUriScheme 移除启动外部 URI 方案事件的处理程序。
func (i *ICoreWebView2_18) RemoveLaunchingExternalUriScheme(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
