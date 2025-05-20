package edge

import "strings"

// --------------------------- 事件 ---------------------------

// Event_SourceChanged 源改变事件.
//   - SourceChanged 会在 Source 属性发生变化时触发。
//   - SourceChanged 会在导航到不同站点或进行片段导航时运行。
//   - 对于其他类型的导航，例如页面刷新或 history.pushState（使用与当前页面相同的 URL），它不会触发。
//   - SourceChanged 会在导航到新文档时，在 ContentLoading 之前运行。
func (e *webview2) Event_SourceChanged(cb func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr) error {
	if e.Handler_SourceChangedEvent == nil {
		e.Handler_SourceChangedEvent = NewICoreWebView2SourceChangedEventHandler(e)
		err := e.CoreWebView.AddSourceChanged(e.Handler_SourceChangedEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_SourceChangedEvent = nil
			return err
		}
	}
	e.cb_SourceChangedEvent = cb
	return nil
}

// SourceChanged 当源改变时调用。
func (e *webview2) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	if e.cb_SourceChangedEvent != nil {
		e.cb_SourceChangedEvent(sender, args)
	}
	return 0
}

// Event_NewWindowRequested 新窗口请求事件.
//   - 当 Webview 内的内容请求打开新窗口时（例如通过 NewWindowRequested），NewWindowRequested 事件将运行。
//   - 应用可以传递一个目标 Webview 作为打开的窗口，或者将该事件标记为 Handled，在这种情况下，WebView2 不会打开窗口。
//   - 如果 Handled 或 NewWindow 属性均未设置，目标内容将在弹出窗口中打开。
func (e *webview2) Event_NewWindowRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr) error {
	if e.Handler_NewWindowRequestedEvent == nil {
		e.Handler_NewWindowRequestedEvent = NewICoreWebView2NewWindowRequestedEventHandler(e)
		err := e.CoreWebView.AddNewWindowRequested(e.Handler_NewWindowRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_NewWindowRequestedEvent = nil
			return err
		}
	}
	e.cb_NewWindowRequestedEvent = cb
	return nil
}

// NewWindowRequested 当收到新窗口请求时调用。
func (e *webview2) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	if e.cb_NewWindowRequestedEvent != nil {
		e.cb_NewWindowRequestedEvent(sender, args)
	}
	return 0
}

// PermissionRequested 当收到权限请求时调用。
func (e *webview2) PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	kind, _ := args.GetPermissionKind()
	e.rwxPermissions.RLock()
	result, ok := e.permissions[kind]
	e.rwxPermissions.RUnlock()
	if !ok {
		result = COREWEBVIEW2_PERMISSION_STATE_DEFAULT
	}
	_ = args.PutState(result)

	if e.cb_PermissionRequestedEvent != nil {
		e.cb_PermissionRequestedEvent(sender, args)
	}
	return 0
}

// Event_PermissionRequested 权限请求事件.
//   - PermissionRequested 在 Webview 中的内容请求访问某些特权资源的权限时运行。
func (e *webview2) Event_PermissionRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr) {
	e.cb_PermissionRequestedEvent = cb
}

// MessageReceived 当从 webview 收到消息时调用。
func (e *webview2) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	if e.msgcb_xcgui != nil {
		message, err := args.TryGetWebMessageAsString()
		if err == nil && strings.HasPrefix(message, "{\"id\":") {
			e.msgcb_xcgui(message)
		}
	}
	if e.cb_MessageReceivedEvent != nil {
		e.cb_MessageReceivedEvent(sender, args)
	}
	return 0
}

// Event_MessageReceived 网页消息事件.
//   - 当 Webview 的顶级文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 事件会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何受 JSON 转换支持的对象。
func (e *webview2) Event_MessageReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr) {
	e.cb_MessageReceivedEvent = cb
}

// WebResourceRequested 当收到资源请求时调用。
func (e *webview2) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	if e.cb_WebResourceRequestedEvent != nil {
		e.cb_WebResourceRequestedEvent(sender, args)
	}
	return 0
}

// Event_WebResourceRequested 网页资源请求事件.
func (e *webview2) Event_WebResourceRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr) error {
	if e.Handler_WebResourceRequestedEvent == nil {
		e.Handler_WebResourceRequestedEvent = NewICoreWebView2WebResourceRequestedEventHandler(e)
		err := e.CoreWebView.AddWebResourceRequested(e.Handler_WebResourceRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_WebResourceRequestedEvent = nil
			return err
		}
	}
	e.cb_WebResourceRequestedEvent = cb
	return nil
}

// NavigationCompleted 当导航完成时调用。
func (e *webview2) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.cb_NavigationCompletedEvent != nil {
		e.cb_NavigationCompletedEvent(sender, args)
	}
	return 0
}

// Event_NavigationCompleted 导航完成事件.
//   - NavigationCompleted 事件会在 Webview 完全加载完毕（与 body.onload 事件同时发生）或加载因错误而停止时触发。
func (e *webview2) Event_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) error {
	if e.Handler_NavigationCompletedEvent == nil {
		e.Handler_NavigationCompletedEvent = NewICoreWebView2NavigationCompletedEventHandler(e)
		err := e.CoreWebView.AddNavigationCompleted(e.Handler_NavigationCompletedEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_NavigationCompletedEvent = nil
			return err
		}
	}
	e.cb_NavigationCompletedEvent = cb
	return nil
}

// NavigationStarting 当导航开始时调用。
func (e *webview2) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.cb_NavigationStartingEvent != nil {
		e.cb_NavigationStartingEvent(sender, args)
	}
	return 0
}

// Event_NavigationStarting 导航开始事件.
//   - NavigationStarting 在 Webview 主框架请求导航到不同的统一资源标识符 (URI) 权限时运行。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有 NavigationStarting 事件处理程序返回之前，导航将被阻止。
func (e *webview2) Event_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) error {
	if e.Handler_NavigationStartingEvent == nil {
		e.Handler_NavigationStartingEvent = NewICoreWebView2NavigationStartingEventHandler(e)
		err := e.CoreWebView.AddNavigationStarting(e.Handler_NavigationStartingEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_NavigationStartingEvent = nil
			return err
		}
	}
	e.cb_NavigationStartingEvent = cb
	return nil
}

// AcceleratorKeyPressed 当使用快捷键时调用。
func (e *webview2) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if e.cb_AcceleratorKeyPressedEvent != nil {
		e.cb_AcceleratorKeyPressedEvent(sender, args)
	}
	return 0
}

// Event_AcceleratorKeyPressed 快捷键事件.
//   - AcceleratorKeyPressed 在 Webview 获得焦点时，当按下或释放快捷键或组合键时运行。
func (e *webview2) Event_AcceleratorKeyPressed(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr) error {
	if e.Handler_AcceleratorKeyPressedEvent == nil {
		e.Handler_AcceleratorKeyPressedEvent = NewICoreWebView2AcceleratorKeyPressedEventHandler(e)
		err := e.Controller.AddAcceleratorKeyPressed(e.Handler_AcceleratorKeyPressedEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_AcceleratorKeyPressedEvent = nil
			return err
		}
	}
	e.cb_AcceleratorKeyPressedEvent = cb
	return nil
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (e *webview2) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	if e.cb_ContentLoadingEvent != nil {
		e.cb_ContentLoadingEvent(sender, args)
	}
	return 0
}

// Event_ContentLoading 网页内容正在加载事件.
//   - ContentLoading 会在任何内容加载之前触发，包括使用 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - ContentLoading 在发生相同页面导航时（例如通过 fragment 导航或 history.pushState 导航）不会触发。
//   - 此操作在 NavigationStarting 和 SourceChanged 事件之后，以及 HistoryChanged 和 NavigationCompleted 事件之前发生。
func (e *webview2) Event_ContentLoading(cb func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr) error {
	if e.Handler_ContentLoadingEvent == nil {
		e.Handler_ContentLoadingEvent = NewICoreWebView2ContentLoadingEventHandler(e)
		err := e.CoreWebView.AddContentLoading(e.Handler_ContentLoadingEvent, e.EventRegistrationToken)
		if err != nil {
			e.Handler_ContentLoadingEvent = nil
			return err
		}
	}
	e.cb_ContentLoadingEvent = cb
	return nil
}
