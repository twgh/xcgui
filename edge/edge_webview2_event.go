package edge

import "strings"

// --------------------------- 事件 ---------------------------

// DocumentTitleChanged 在文档标题发生变化时调用.
func (e *Webview2) DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.cbDocumentTitleChangedEvent != nil {
		e.cbDocumentTitleChangedEvent(sender, args)
	}
	return 0
}

// Event_DocumentTitleChanged 文档标题改变事件.
//   - 当 Webview 的 DocumentTitle 属性发生变化时，DocumentTitleChanged 会运行，并且可能在 NavigationCompleted 事件之前或之后运行。
func (e *Webview2) Event_DocumentTitleChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if e.HandlerDocumentTitleChangedEvent == nil {
		e.HandlerDocumentTitleChangedEvent = NewICoreWebView2DocumentTitleChangedEventHandler(e)
		err := e.CoreWebView.AddDocumentTitleChanged(e.HandlerDocumentTitleChangedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerDocumentTitleChangedEvent = nil
			return err
		}
	}
	e.cbDocumentTitleChangedEvent = cb
	return nil
}

// RasterizationScaleChanged 在光栅化缩放比例发生变化时调用.
func (e *Webview2) RasterizationScaleChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	if e.cbRasterizationScaleChangedEvent != nil {
		e.cbRasterizationScaleChangedEvent(sender, args)
	}
	return 0
}

// Event_RasterizationScaleChanged 光栅化缩放比例改变事件.
//   - 当 Webview 检测到显示器 DPI 缩放比例已更改、ShouldDetectMonitorScaleChanges 为 true 且 Webview 已更改 RasterizationScale 属性时，将引发此事件。
func (e *Webview2) Event_RasterizationScaleChanged(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr) error {
	c3, err := e.Controller.GetICoreWebView2Controller3()
	if err != nil {
		return err
	}
	defer c3.Release()
	if e.HandlerRasterizationScaleChangedEvent == nil {
		e.HandlerRasterizationScaleChangedEvent = NewICoreWebView2RasterizationScaleChangedEventHandler(e)
		e.TokenRasterizationScaleChangedEvent, err = c3.AddRasterizationScaleChanged(e.HandlerRasterizationScaleChangedEvent)
		if err != nil {
			e.HandlerRasterizationScaleChangedEvent = nil
			return err
		}
	}
	e.cbRasterizationScaleChangedEvent = cb
	return nil
}

// WindowCloseRequested 在窗口关闭请求时调用.
func (e *Webview2) WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.cbWindowCloseRequestedEvent != nil {
		e.cbWindowCloseRequestedEvent(sender, args)
	}
	return 0
}

// Event_WindowCloseRequested 窗口关闭请求事件.
//   - WindowCloseRequested 在 Webview 内部的内容请求关闭窗口时触发，例如在运行 window.close 之后。如果这对应用程序有意义，应用程序应该关闭 Webview 和相关的应用程序窗口。
//   - 在首次调用 window.close() 之后，对于任何紧接着连续调用的 window.close()，此事件可能不会触发。
func (e *Webview2) Event_WindowCloseRequested(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if e.HandlerWindowCloseRequestedEvent == nil {
		e.HandlerWindowCloseRequestedEvent = NewICoreWebView2WindowCloseRequestedEventHandler(e)
		err := e.CoreWebView.AddWindowCloseRequested(e.HandlerWindowCloseRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerWindowCloseRequestedEvent = nil
			return err
		}
	}
	e.cbWindowCloseRequestedEvent = cb
	return nil
}

// Event_SourceChanged 源改变事件.
//   - SourceChanged 会在 Source 属性发生变化时触发。
//   - SourceChanged 会在导航到不同站点或进行片段导航时运行。
//   - 对于其他类型的导航，例如页面刷新或 history.pushState（使用与当前页面相同的 URL），它不会触发。
//   - SourceChanged 会在导航到新文档时，在 ContentLoading 之前运行。
func (e *Webview2) Event_SourceChanged(cb func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr) error {
	if e.HandlerSourceChangedEvent == nil {
		e.HandlerSourceChangedEvent = NewICoreWebView2SourceChangedEventHandler(e)
		err := e.CoreWebView.AddSourceChanged(e.HandlerSourceChangedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerSourceChangedEvent = nil
			return err
		}
	}
	e.cbSourceChangedEvent = cb
	return nil
}

// SourceChanged 当源改变时调用。
func (e *Webview2) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	if e.cbSourceChangedEvent != nil {
		e.cbSourceChangedEvent(sender, args)
	}
	return 0
}

// Event_NewWindowRequested 新窗口请求事件.
//   - 当 Webview 内的内容请求打开新窗口时（例如通过 NewWindowRequested），NewWindowRequested 事件将运行。
//   - 应用可以传递一个目标 Webview 作为打开的窗口，或者将该事件标记为 Handled，在这种情况下，WebView2 不会打开窗口。
//   - 如果 Handled 或 NewWindow 属性均未设置，目标内容将在弹出窗口中打开。
func (e *Webview2) Event_NewWindowRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr) error {
	if e.HandlerNewWindowRequestedEvent == nil {
		e.HandlerNewWindowRequestedEvent = NewICoreWebView2NewWindowRequestedEventHandler(e)
		err := e.CoreWebView.AddNewWindowRequested(e.HandlerNewWindowRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerNewWindowRequestedEvent = nil
			return err
		}
	}
	e.cbNewWindowRequestedEvent = cb
	return nil
}

// NewWindowRequested 当收到新窗口请求时调用。
func (e *Webview2) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	if e.cbNewWindowRequestedEvent != nil {
		e.cbNewWindowRequestedEvent(sender, args)
	}
	return 0
}

// PermissionRequested 当收到权限请求时调用。
func (e *Webview2) PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	kind, _ := args.GetPermissionKind()
	e.rwxPermissions.RLock()
	result, ok := e.permissions[kind]
	e.rwxPermissions.RUnlock()
	if !ok {
		result = COREWEBVIEW2_PERMISSION_STATE_DEFAULT
	}
	_ = args.PutState(result)

	if e.cbPermissionRequestedEvent != nil {
		e.cbPermissionRequestedEvent(sender, args)
	}
	return 0
}

// Event_PermissionRequested 权限请求事件.
//   - PermissionRequested 在 Webview 中的内容请求访问某些特权资源的权限时运行。
func (e *Webview2) Event_PermissionRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr) {
	e.cbPermissionRequestedEvent = cb
}

// MessageReceived 当从 webview 收到消息时调用。
func (e *Webview2) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	if e.msgcb_xcgui != nil {
		message, err := args.TryGetWebMessageAsString()
		if err == nil && strings.HasPrefix(message, "{\"id\":") {
			e.msgcb_xcgui(message)
		}
	}
	if e.cbMessageReceivedEvent != nil {
		e.cbMessageReceivedEvent(sender, args)
	}
	return 0
}

// Event_MessageReceived 网页消息事件.
//   - 当 Webview 的顶级文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 事件会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何受 JSON 转换支持的对象。
func (e *Webview2) Event_MessageReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr) {
	e.cbMessageReceivedEvent = cb
}

// WebResourceRequested 当收到资源请求时调用。
func (e *Webview2) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	if e.cbWebResourceRequestedEvent != nil {
		e.cbWebResourceRequestedEvent(sender, args)
	}
	return 0
}

// Event_WebResourceRequested 网页资源请求事件.
func (e *Webview2) Event_WebResourceRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr) error {
	if e.HandlerWebResourceRequestedEvent == nil {
		e.HandlerWebResourceRequestedEvent = NewICoreWebView2WebResourceRequestedEventHandler(e)
		err := e.CoreWebView.AddWebResourceRequested(e.HandlerWebResourceRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerWebResourceRequestedEvent = nil
			return err
		}
	}
	e.cbWebResourceRequestedEvent = cb
	return nil
}

// NavigationCompleted 当导航完成时调用。
func (e *Webview2) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.cbNavigationCompletedEvent != nil {
		e.cbNavigationCompletedEvent(sender, args)
	}
	return 0
}

// Event_NavigationCompleted 导航完成事件.
//   - NavigationCompleted 事件会在 Webview 完全加载完毕（与 body.onload 事件同时发生）或加载因错误而停止时触发。
func (e *Webview2) Event_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) error {
	if e.HandlerNavigationCompletedEvent == nil {
		e.HandlerNavigationCompletedEvent = NewICoreWebView2NavigationCompletedEventHandler(e)
		err := e.CoreWebView.AddNavigationCompleted(e.HandlerNavigationCompletedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerNavigationCompletedEvent = nil
			return err
		}
	}
	e.cbNavigationCompletedEvent = cb
	return nil
}

// NavigationCompleted2 当框架导航完成时调用。
func (e *Webview2) NavigationCompleted2(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.cbFrameNavigationCompletedEvent != nil {
		e.cbFrameNavigationCompletedEvent(sender, args)
	}
	return 0
}

// Event_FrameNavigationCompleted 框架导航完成事件.
//   - 框架导航完成(NavigationCompleted2)事件会在 Webview 中的子框架完全加载完毕（与 body.onload 触发同时）或加载因错误而停止时触发。
func (e *Webview2) Event_FrameNavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) error {
	if e.HandlerFrameNavigationCompletedEvent == nil {
		e.HandlerFrameNavigationCompletedEvent = NewICoreWebView2NavigationCompletedEventHandler2(e)
		err := e.CoreWebView.AddFrameNavigationCompleted(e.HandlerFrameNavigationCompletedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerFrameNavigationCompletedEvent = nil
			return err
		}
	}
	e.cbFrameNavigationCompletedEvent = cb
	return nil
}

// NavigationStarting2 当框架导航开始时调用。
func (e *Webview2) NavigationStarting2(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.cbFrameNavigationStartingEvent != nil {
		e.cbFrameNavigationStartingEvent(sender, args)
	}
	return 0
}

// Event_FrameNavigationStarting 框架导航开始事件。
//   - 框架导航开始(NavigationStarting2)事件会在 Webview 中的子框架请求导航到不同 URI 的权限时触发。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有框架导航开始((NavigationStarting2))事件处理程序返回之前，导航将被阻止。
func (e *Webview2) Event_FrameNavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) error {
	if e.HandlerFrameNavigationStartingEvent == nil {
		e.HandlerFrameNavigationStartingEvent = NewICoreWebView2NavigationStartingEventHandler2(e)
		err := e.CoreWebView.AddFrameNavigationStarting(e.HandlerFrameNavigationStartingEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerFrameNavigationStartingEvent = nil
			return err
		}
	}
	e.cbFrameNavigationStartingEvent = cb
	return nil
}

// NavigationStarting 当导航开始时调用。
func (e *Webview2) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.cbNavigationStartingEvent != nil {
		e.cbNavigationStartingEvent(sender, args)
	}
	return 0
}

// Event_NavigationStarting 导航开始事件.
//   - NavigationStarting 在 Webview 主框架请求导航到不同的统一资源标识符 (URI) 权限时运行。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有 NavigationStarting 事件处理程序返回之前，导航将被阻止。
func (e *Webview2) Event_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) error {
	if e.HandlerNavigationStartingEvent == nil {
		e.HandlerNavigationStartingEvent = NewICoreWebView2NavigationStartingEventHandler(e)
		err := e.CoreWebView.AddNavigationStarting(e.HandlerNavigationStartingEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerNavigationStartingEvent = nil
			return err
		}
	}
	e.cbNavigationStartingEvent = cb
	return nil
}

// AcceleratorKeyPressed 当使用快捷键时调用。
func (e *Webview2) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if e.cbAcceleratorKeyPressedEvent != nil {
		e.cbAcceleratorKeyPressedEvent(sender, args)
	}
	return 0
}

// Event_AcceleratorKeyPressed 快捷键事件.
//   - AcceleratorKeyPressed 在 Webview 获得焦点时，当按下或释放快捷键或组合键时运行。
func (e *Webview2) Event_AcceleratorKeyPressed(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr) error {
	if e.HandlerAcceleratorKeyPressedEvent == nil {
		e.HandlerAcceleratorKeyPressedEvent = NewICoreWebView2AcceleratorKeyPressedEventHandler(e)
		err := e.Controller.AddAcceleratorKeyPressed(e.HandlerAcceleratorKeyPressedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerAcceleratorKeyPressedEvent = nil
			return err
		}
	}
	e.cbAcceleratorKeyPressedEvent = cb
	return nil
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (e *Webview2) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	if e.cbContentLoadingEvent != nil {
		e.cbContentLoadingEvent(sender, args)
	}
	return 0
}

// Event_ContentLoading 网页内容正在加载事件.
//   - ContentLoading 会在任何内容加载之前触发，包括使用 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - ContentLoading 在发生相同页面导航时（例如通过 fragment 导航或 history.pushState 导航）不会触发。
//   - 此操作在 NavigationStarting 和 SourceChanged 事件之后，以及 HistoryChanged 和 NavigationCompleted 事件之前发生。
func (e *Webview2) Event_ContentLoading(cb func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr) error {
	if e.HandlerContentLoadingEvent == nil {
		e.HandlerContentLoadingEvent = NewICoreWebView2ContentLoadingEventHandler(e)
		err := e.CoreWebView.AddContentLoading(e.HandlerContentLoadingEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerContentLoadingEvent = nil
			return err
		}
	}
	e.cbContentLoadingEvent = cb
	return nil
}
