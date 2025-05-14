package edge

import "strings"

// --------------------------- 事件 ---------------------------

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

	if e.permissionRequestedCallback != nil {
		e.permissionRequestedCallback(sender, args)
	}
	return 0
}

// Event_PermissionRequested 权限请求事件.
func (e *webview2) Event_PermissionRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr) {
	e.permissionRequestedCallback = cb
}

// MessageReceived 当从 webview 收到消息时调用。
func (e *webview2) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	if e.msgcb_xcgui != nil {
		message, err := args.TryGetWebMessageAsString()
		if err == nil && strings.HasPrefix(message, "{\"id\":") {
			e.msgcb_xcgui(message)
		}
	}
	if e.messageReceivedCallback != nil {
		e.messageReceivedCallback(sender, args)
	}
	return 0
}

// Event_MessageReceived 网页消息事件.
func (e *webview2) Event_MessageReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr) {
	e.messageReceivedCallback = cb
}

// WebResourceRequested 当收到资源请求时调用。
func (e *webview2) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	if e.webResourceRequestedCallback != nil {
		e.webResourceRequestedCallback(sender, args)
	}
	return 0
}

// Event_WebResourceRequested 网页资源请求事件.
func (e *webview2) Event_WebResourceRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr) {
	if e.webResourceRequestedEventHandler == nil {
		e.webResourceRequestedEventHandler = NewICoreWebView2WebResourceRequestedEventHandler(e)
		_ = e.CoreWebView.AddWebResourceRequested(e.webResourceRequestedEventHandler, e.eventRegistrationToken)
	}
	e.webResourceRequestedCallback = cb
}

// NavigationCompleted 当导航完成时调用。
func (e *webview2) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.navigationCompletedCallback != nil {
		e.navigationCompletedCallback(sender, args)
	}
	return 0
}

// Event_NavigationCompleted 导航完成事件.
func (e *webview2) Event_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) {
	if e.navigationCompletedEventHandler == nil {
		e.navigationCompletedEventHandler = NewICoreWebView2NavigationCompletedEventHandler(e)
		_ = e.CoreWebView.AddNavigationCompleted(e.navigationCompletedEventHandler, e.eventRegistrationToken)
	}
	e.navigationCompletedCallback = cb
}

// NavigationStarting 当导航开始时调用。
func (e *webview2) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.navigationStartingCallback != nil {
		e.navigationStartingCallback(sender, args)
	}
	return 0
}

// Event_NavigationStarting 导航开始事件.
func (e *webview2) Event_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) {
	if e.navigationStartingEventHandler == nil {
		e.navigationStartingEventHandler = NewICoreWebView2NavigationStartingEventHandler(e)
		_ = e.CoreWebView.AddNavigationStarting(e.navigationStartingEventHandler, e.eventRegistrationToken)
	}
	e.navigationStartingCallback = cb
}

// AcceleratorKeyPressed 当使用快捷键时，会调用。
func (e *webview2) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if e.acceleratorKeyPressedCallback != nil {
		e.acceleratorKeyPressedCallback(sender, args)
	}
	return 0
}

// Event_AcceleratorKeyPressed 快捷键事件.
func (e *webview2) Event_AcceleratorKeyPressed(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr) {
	if e.acceleratorKeyPressedEventHandler == nil {
		e.acceleratorKeyPressedEventHandler = NewICoreWebView2AcceleratorKeyPressedEventHandler(e)
		_ = e.Controller.AddAcceleratorKeyPressed(e.acceleratorKeyPressedEventHandler, e.eventRegistrationToken)
	}
	e.acceleratorKeyPressedCallback = cb
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (e *webview2) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	if e.contentLoadingCallback != nil {
		e.contentLoadingCallback(sender, args)
	}
	return 0
}

// Event_ContentLoading 网页内容正在加载事件.
//
// cb: 当 WebView 控件开始加载内容时调用。
func (e *webview2) Event_ContentLoading(cb func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr) {
	if e.contentLoadingEventHandler == nil {
		e.contentLoadingEventHandler = NewICoreWebView2ContentLoadingEventHandler(e)
		_ = e.CoreWebView.AddContentLoading(e.contentLoadingEventHandler, e.eventRegistrationToken)
	}
	e.contentLoadingCallback = cb
}
