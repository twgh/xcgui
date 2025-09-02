package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame2 是 ICoreWebView2Frame 的延续, 提供对 iframe 的更多功能访问，包括导航、内容加载、执行脚本和 Web 消息传递。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame2
type ICoreWebView2Frame2 struct {
	ICoreWebView2Frame
}

// Event_FrameNavigationStarting 框架导航开始事件.
//   - 框架导航会引发 NavigationStarting 事件和 ICoreWebView2.FrameNavigationStarting 事件。
//   - 当前框架的所有 FrameNavigationStarting 事件处理程序将在 NavigationStarting 事件处理程序之前运行。
//   - 所有事件处理程序共享一个通用的 NavigationStartingEventArgs 对象。
//   - 最后更改 NavigationStartingEventArgs.Cancel 属性的事件处理程序将决定是否取消框架导航。
//   - 重定向也会引发此事件，且导航 ID 与原始导航 ID 相同。
//   - 在所有 NavigationStarting 和 ICoreWebView2.FrameNavigationStarting 事件处理程序返回之前，导航将被阻止。
func (i *ICoreWebView2Frame2) Event_FrameNavigationStarting(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameNavigationStarting", cb, i, allowAddingMultiple...)
}

// Event_FrameContentLoading 框架内容加载事件.
//   - ContentLoading 在任何内容加载前触发，包括通过 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - 如果发生同页面导航（例如通过 fragment 导航或 history.pushState 导航），ContentLoading 不会触发。
//   - 此操作在 NavigationStarting 之后、NavigationCompleted 事件之前发生。
func (i *ICoreWebView2Frame2) Event_FrameContentLoading(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameContentLoading", cb, i, allowAddingMultiple...)
}

// Event_FrameNavigationCompleted 框架导航完成事件.
//   - NavigationCompleted 在 ICoreWebView2Frame 完全加载时（与 body.onload 同时运行）或加载因错误而停止时运行。
func (i *ICoreWebView2Frame2) Event_FrameNavigationCompleted(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameNavigationCompleted", cb, i, allowAddingMultiple...)
}

// Event_FrameDOMContentLoaded 框架 DOM 内容加载完成事件.
//   - 当 iframe 的 html 文档被解析后，会触发 DOMContentLoaded 事件。
//   - 这与 html 中文档的 DOMContentLoaded 事件是一致的。
func (i *ICoreWebView2Frame2) Event_FrameDOMContentLoaded(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameDOMContentLoaded", cb, i, allowAddingMultiple...)
}

// Event_FrameWebMessageReceived 框架 Web 消息接收事件.
//   - 当 ICoreWebView2Settings.IsWebMessageEnabled 设置已启用且框架文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何支持 JSON 转换的对象。
//   - 当框架调用 postMessage 时，对象参数会转换为 JSON 字符串，并异步发送到宿主进程。这将导致处理程序的 Invoke 方法被调用，该 JSON 字符串作为其参数。
func (i *ICoreWebView2Frame2) Event_FrameWebMessageReceived(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FrameWebMessageReceived", cb, i, allowAddingMultiple...)
}

// AddNavigationStarting 添加导航开始事件处理程序。
//   - 框架导航会引发 NavigationStarting 事件和 ICoreWebView2.FrameNavigationStarting 事件。
//   - 当前框架的所有 FrameNavigationStarting 事件处理程序将在 NavigationStarting 事件处理程序之前运行。
//   - 所有事件处理程序共享一个通用的 NavigationStartingEventArgs 对象。
//   - 最后更改 NavigationStartingEventArgs.Cancel 属性的事件处理程序将决定是否取消框架导航。
//   - 重定向也会引发此事件，且导航 ID 与原始导航 ID 相同。
//   - 在所有 NavigationStarting 和 ICoreWebView2.FrameNavigationStarting 事件处理程序返回之前，导航将被阻止。
func (i *ICoreWebView2Frame2) AddNavigationStarting(eventHandler *ICoreWebView2FrameNavigationStartingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNavigationStarting 移除导航开始事件处理程序。
func (i *ICoreWebView2Frame2) RemoveNavigationStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddContentLoading 添加内容加载事件处理程序。
//   - ContentLoading 在任何内容加载前触发，包括通过 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - 如果发生同页面导航（例如通过 fragment 导航或 history.pushState 导航），ContentLoading 不会触发。
//   - 此操作在 NavigationStarting 之后、NavigationCompleted 事件之前发生。
func (i *ICoreWebView2Frame2) AddContentLoading(eventHandler *ICoreWebView2FrameContentLoadingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveContentLoading 移除内容加载事件处理程序。
func (i *ICoreWebView2Frame2) RemoveContentLoading(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddNavigationCompleted 添加导航完成事件处理程序。
//   - NavigationCompleted 在 ICoreWebView2Frame 完全加载时（与 body.onload 同时运行）或加载因错误而停止时运行。
func (i *ICoreWebView2Frame2) AddNavigationCompleted(eventHandler *ICoreWebView2FrameNavigationCompletedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNavigationCompleted 移除导航完成事件处理程序。
func (i *ICoreWebView2Frame2) RemoveNavigationCompleted(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddDOMContentLoaded 添加 DOM 内容加载完成事件处理程序。
//   - 当 iframe 的 html 文档被解析后，会触发 DOMContentLoaded 事件。
//   - 这与 html 中文档的 DOMContentLoaded 事件是一致的。
func (i *ICoreWebView2Frame2) AddDOMContentLoaded(eventHandler *ICoreWebView2FrameDOMContentLoadedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDOMContentLoaded 移除 DOM 内容加载完成事件处理程序。
func (i *ICoreWebView2Frame2) RemoveDOMContentLoaded(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ExecuteScript 在当前框架的上下文中执行 JavaScript 代码。
//   - 对所提供的 JavaScript 进行求值的结果会传递给完成处理程序。
//   - 该结果值是一个 JSON 编码的字符串。
//   - 如果结果为未定义、包含引用循环，或者因其他原因无法编码为 JSON，那么该结果将被视为 null，在 JSON 中编码为字符串"null"。
//   - 没有显式返回值的函数会返回 undefined。
//   - 如果运行的脚本抛出未处理的异常，那么结果也会是“null”。
//   - 此方法以异步方式应用。如果该方法在 ContentLoading 之前运行，脚本将不会执行，并且会返回字符串“null”。
//   - 即使 ICoreWebView2Settings.IsScriptEnabled 被设置为 FALSE，此操作也会执行该脚本。
func (i *ICoreWebView2Frame2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) error {
	_javaScript, err := syscall.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.ExecuteScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ExecuteScriptEx 在当前框架的上下文中执行 JavaScript 代码。
//   - 对所提供的 JavaScript 进行求值的结果会传递给完成处理程序。
//   - 该结果值是一个 JSON 编码的字符串。
//   - 如果结果为未定义、包含引用循环，或者因其他原因无法编码为 JSON，那么该结果将被视为 null，在 JSON 中编码为字符串"null"。
//   - 没有显式返回值的函数会返回 undefined。
//   - 如果运行的脚本抛出未处理的异常，那么结果也会是“null”。
//   - 此方法以异步方式应用。如果该方法在 ContentLoading 之前运行，脚本将不会执行，并且会返回字符串“null”。
//   - 即使 ICoreWebView2Settings.IsScriptEnabled 被设置为 FALSE，此操作也会执行该脚本。
//
// impl: *WebViewEventImpl。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 执行完成后的回调处理程序，可以为 nil。如果 JavaScript 返回值，将通过 cb 返回。
func (i *ICoreWebView2Frame2) ExecuteScriptEx(impl *WebViewEventImpl, javaScript string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ExecuteScriptCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ExecuteScriptCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ExecuteScriptCompleted")
	}
	return i.ExecuteScript(javaScript, (*ICoreWebView2ExecuteScriptCompletedHandler)(handler))
}

// PostWebMessageAsJson 将指定的 webMessage 发布到框架。
//   - 在 iframe 中运行的脚本可以通过 window.chrome.webview.addEventListener('message', handler) 接收消息。
//   - 事件参数是 MessageEvent 的实例。
//   - ICoreWebView2Settings.IsWebMessageEnabled 设置必须为 TRUE，否则消息将不会被发送。
//   - 事件参数的 data 属性是将 webMessage 字符串参数解析为 JSON 字符串后得到的 JavaScript 对象。
//   - 事件参数的 source 属性是对 window.chrome.webview 对象的引用。
//   - 有关从 WebView 中的 HTML 文档向主机发送消息的信息，请访问 [add_WebMessageReceived](https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived)。
//   - 消息以异步方式传递。如果在消息发布到页面之前发生导航，该消息将被丢弃。
func (i *ICoreWebView2Frame2) PostWebMessageAsJson(webMessageAsJson string) error {
	_webMessageAsJson, err := syscall.UTF16PtrFromString(webMessageAsJson)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostWebMessageAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PostWebMessageAsString 发布一条消息，该消息是一个简单字符串，而非 JavaScript 对象的 JSON 字符串表示形式。
//   - 其行为方式与 PostWebMessageAsJson 完全相同，但 window.chrome.webview 消息的事件参数的 data 属性是一个字符串，其值与 webMessageAsString 相同。
//   - 如果您希望使用简单字符串而非 JSON 对象进行通信，请使用此方法替代 PostWebMessageAsJson。
func (i *ICoreWebView2Frame2) PostWebMessageAsString(webMessageAsString string) error {
	_webMessageAsString, err := syscall.UTF16PtrFromString(webMessageAsString)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsString)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddWebMessageReceived 添加 Web 消息接收事件处理程序。
//   - 当 ICoreWebView2Settings.IsWebMessageEnabled 设置已启用且框架文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何支持 JSON 转换的对象。
//   - 当框架调用 postMessage 时，对象参数会转换为 JSON 字符串，并异步发送到宿主进程。这将导致处理程序的 Invoke 方法被调用，该 JSON 字符串作为其参数。
func (i *ICoreWebView2Frame2) AddWebMessageReceived(handler *ICoreWebView2FrameWebMessageReceivedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebMessageReceived 移除 Web 消息接收事件处理程序。
func (i *ICoreWebView2Frame2) RemoveWebMessageReceived(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
