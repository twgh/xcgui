package edge

// --------------------------- 事件 ---------------------------

// Event_DocumentTitleChanged 文档标题改变事件.
//   - 当 Webview 的 DocumentTitle 属性发生变化时，DocumentTitleChanged 会运行，并且可能在 NavigationCompleted 事件之前或之后运行。
func (w *WebViewEventImpl) Event_DocumentTitleChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if w.HandlerDocumentTitleChangedEvent == nil {
		w.HandlerDocumentTitleChangedEvent = NewICoreWebView2DocumentTitleChangedEventHandler(w)
		err := w.CoreWebView.AddDocumentTitleChanged(w.HandlerDocumentTitleChangedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbDocumentTitleChangedEvent = cb
	return nil
}

// Event_RasterizationScaleChanged 光栅化缩放比例改变事件.
//   - 当 Webview 检测到显示器 DPI 缩放比例已更改、ShouldDetectMonitorScaleChanges 为 true 且 Webview 已更改 RasterizationScale 属性时，将引发此事件。
func (w *WebViewEventImpl) Event_RasterizationScaleChanged(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr) error {
	c3, err := w.Controller.GetICoreWebView2Controller3()
	if err != nil {
		return err
	}
	defer c3.Release()
	if w.HandlerRasterizationScaleChangedEvent == nil {
		w.HandlerRasterizationScaleChangedEvent = NewICoreWebView2RasterizationScaleChangedEventHandler(w)
		w.TokenRasterizationScaleChangedEvent, err = c3.AddRasterizationScaleChanged(w.HandlerRasterizationScaleChangedEvent)
		if err != nil {
			return err
		}
	}
	w.cbRasterizationScaleChangedEvent = cb
	return nil
}

// Event_WindowCloseRequested 窗口关闭请求事件.
//   - WindowCloseRequested 在 Webview 内部的内容请求关闭窗口时触发，例如在运行 window.close 之后。如果这对应用程序有意义，应用程序应该关闭 Webview 和相关的应用程序窗口。
//   - 在首次调用 window.close() 之后，对于任何紧接着连续调用的 window.close()，此事件可能不会触发。
func (w *WebViewEventImpl) Event_WindowCloseRequested(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if w.HandlerWindowCloseRequestedEvent == nil {
		w.HandlerWindowCloseRequestedEvent = NewICoreWebView2WindowCloseRequestedEventHandler(w)
		err := w.CoreWebView.AddWindowCloseRequested(w.HandlerWindowCloseRequestedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbWindowCloseRequestedEvent = cb
	return nil
}

// Event_SourceChanged 源改变事件.
//   - SourceChanged 会在 Source 属性发生变化时触发。
//   - SourceChanged 会在导航到不同站点或进行片段导航时运行。
//   - 对于其他类型的导航，例如页面刷新或 history.pushState（使用与当前页面相同的 URL），它不会触发。
//   - SourceChanged 会在导航到新文档时，在 ContentLoading 之前运行。
func (w *WebViewEventImpl) Event_SourceChanged(cb func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr) error {
	if w.HandlerSourceChangedEvent == nil {
		w.HandlerSourceChangedEvent = NewICoreWebView2SourceChangedEventHandler(w)
		err := w.CoreWebView.AddSourceChanged(w.HandlerSourceChangedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbSourceChangedEvent = cb
	return nil
}

// Event_NewWindowRequested 新窗口请求事件.
//   - 当 Webview 内的内容请求打开新窗口时（例如通过 NewWindowRequested），NewWindowRequested 事件将运行。
//   - 应用可以传递一个目标 Webview 作为打开的窗口，或者将该事件标记为 Handled，在这种情况下，WebView2 不会打开窗口。
//   - 如果 Handled 或 NewWindow 属性均未设置，目标内容将在弹出窗口中打开。
func (w *WebViewEventImpl) Event_NewWindowRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr) error {
	if w.HandlerNewWindowRequestedEvent == nil {
		w.HandlerNewWindowRequestedEvent = NewICoreWebView2NewWindowRequestedEventHandler(w)
		err := w.CoreWebView.AddNewWindowRequested(w.HandlerNewWindowRequestedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbNewWindowRequestedEvent = cb
	return nil
}

// Event_PermissionRequested 权限请求事件.
//   - PermissionRequested 在 Webview 中的内容请求访问某些特权资源的权限时运行。
func (w *WebViewEventImpl) Event_PermissionRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr) {
	w.cbPermissionRequestedEvent = cb
}

// Event_MessageReceived 网页消息事件.
//   - 当 Webview 的顶级文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 事件会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何受 JSON 转换支持的对象。
func (w *WebViewEventImpl) Event_MessageReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr) {
	w.cbMessageReceivedEvent = cb
}

// Event_WebResourceRequested 网页资源请求事件.
func (w *WebViewEventImpl) Event_WebResourceRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr) error {
	if w.HandlerWebResourceRequestedEvent == nil {
		w.HandlerWebResourceRequestedEvent = NewICoreWebView2WebResourceRequestedEventHandler(w)
		err := w.CoreWebView.AddWebResourceRequested(w.HandlerWebResourceRequestedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbWebResourceRequestedEvent = cb
	return nil
}

// Event_NavigationCompleted 导航完成事件.
//   - NavigationCompleted 事件会在 Webview 完全加载完毕（与 body.onload 事件同时发生）或加载因错误而停止时触发。
func (w *WebViewEventImpl) Event_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) error {
	if w.HandlerNavigationCompletedEvent == nil {
		w.HandlerNavigationCompletedEvent = NewICoreWebView2NavigationCompletedEventHandler(w)
		err := w.CoreWebView.AddNavigationCompleted(w.HandlerNavigationCompletedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbNavigationCompletedEvent = cb
	return nil
}

// Event_FrameNavigationCompleted 框架导航完成事件.
//   - 框架导航完成(NavigationCompleted2)事件会在 Webview 中的子框架完全加载完毕（与 body.onload 触发同时）或加载因错误而停止时触发。
func (w *WebViewEventImpl) Event_FrameNavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr) error {
	if w.HandlerFrameNavigationCompletedEvent == nil {
		w.HandlerFrameNavigationCompletedEvent = NewICoreWebView2NavigationCompletedEventHandler2(w)
		err := w.CoreWebView.AddFrameNavigationCompleted(w.HandlerFrameNavigationCompletedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbFrameNavigationCompletedEvent = cb
	return nil
}

// Event_FrameNavigationStarting 框架导航开始事件。
//   - 框架导航开始(NavigationStarting2)事件会在 Webview 中的子框架请求导航到不同 URI 的权限时触发。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有框架导航开始((NavigationStarting2))事件处理程序返回之前，导航将被阻止。
func (w *WebViewEventImpl) Event_FrameNavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) error {
	if w.HandlerFrameNavigationStartingEvent == nil {
		w.HandlerFrameNavigationStartingEvent = NewICoreWebView2NavigationStartingEventHandler2(w)
		err := w.CoreWebView.AddFrameNavigationStarting(w.HandlerFrameNavigationStartingEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbFrameNavigationStartingEvent = cb
	return nil
}

// Event_NavigationStarting 导航开始事件.
//   - NavigationStarting 在 Webview 主框架请求导航到不同的统一资源标识符 (URI) 权限时运行。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有 NavigationStarting 事件处理程序返回之前，导航将被阻止。
func (w *WebViewEventImpl) Event_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr) error {
	if w.HandlerNavigationStartingEvent == nil {
		w.HandlerNavigationStartingEvent = NewICoreWebView2NavigationStartingEventHandler(w)
		err := w.CoreWebView.AddNavigationStarting(w.HandlerNavigationStartingEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbNavigationStartingEvent = cb
	return nil
}

// Event_AcceleratorKeyPressed 快捷键事件.
//   - AcceleratorKeyPressed 在 Webview 获得焦点时，当按下或释放快捷键或组合键时运行。
func (w *WebViewEventImpl) Event_AcceleratorKeyPressed(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr) error {
	if w.HandlerAcceleratorKeyPressedEvent == nil {
		w.HandlerAcceleratorKeyPressedEvent = NewICoreWebView2AcceleratorKeyPressedEventHandler(w)
		err := w.Controller.AddAcceleratorKeyPressed(w.HandlerAcceleratorKeyPressedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbAcceleratorKeyPressedEvent = cb
	return nil
}

// Event_ContentLoading 网页内容正在加载事件.
//   - ContentLoading 会在任何内容加载之前触发，包括使用 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - ContentLoading 在发生相同页面导航时（例如通过 fragment 导航或 history.pushState 导航）不会触发。
//   - 此操作在 NavigationStarting 和 SourceChanged 事件之后，以及 HistoryChanged 和 NavigationCompleted 事件之前发生。
func (w *WebViewEventImpl) Event_ContentLoading(cb func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr) error {
	if w.HandlerContentLoadingEvent == nil {
		w.HandlerContentLoadingEvent = NewICoreWebView2ContentLoadingEventHandler(w)
		err := w.CoreWebView.AddContentLoading(w.HandlerContentLoadingEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbContentLoadingEvent = cb
	return nil
}

// Event_ContainsFullScreenElementChanged 是 ContainsFullScreenElement(是否包含全屏元素) 属性改变事件.
//   - 当 ContainsFullScreenElement 属性发生变化时触发
func (w *WebViewEventImpl) Event_ContainsFullScreenElementChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if w.HandlerContainsFullScreenElementChangedEvent == nil {
		w.HandlerContainsFullScreenElementChangedEvent = NewICoreWebView2ContainsFullScreenElementChangedEventHandler(w)
		err := w.CoreWebView.AddContainsFullScreenElementChanged(w.HandlerContainsFullScreenElementChangedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbContainsFullScreenElementChangedEvent = cb
	return nil
}

// Event_ProcessFailed 进程失败事件.
//   - 当 Webview 的进程因任何原因失败时，ProcessFailed 事件将运行。
func (w *WebViewEventImpl) Event_ProcessFailed(cb func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr) error {
	if w.HandlerProcessFailedEvent == nil {
		w.HandlerProcessFailedEvent = NewICoreWebView2ProcessFailedEventHandler(w)
		err := w.CoreWebView.AddProcessFailed(w.HandlerProcessFailedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbProcessFailedEvent = cb
	return nil
}

// Event_HistoryChanged 历史记录改变事件.
//   - HistoryChanged 在联合会话历史记录发生更改时引发，该历史记录由顶级和手动框架导航组成。
func (w *WebViewEventImpl) Event_HistoryChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr) error {
	if w.HandlerHistoryChangedEvent == nil {
		w.HandlerHistoryChangedEvent = NewICoreWebView2HistoryChangedEventHandler(w)
		err := w.CoreWebView.AddHistoryChanged(w.HandlerHistoryChangedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbHistoryChangedEvent = cb
	return nil
}

// Event_ScriptDialogOpening 脚本对话框打开事件.
//   - 当 JavaScript 代码调用 alert、confirm、prompt 或 beforeunload 时触发此事件。
//   - 仅当 ICoreWebView2Settings.AreDefaultScriptDialogsEnabled 属性设置为 FALSE 时，此事件才会触发。
//   - ScriptDialogOpening 事件会抑制对话框，或使用自定义对话框替换默认对话框。
func (w *WebViewEventImpl) Event_ScriptDialogOpening(cb func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr) error {
	if w.HandlerScriptDialogOpeningEvent == nil {
		w.HandlerScriptDialogOpeningEvent = NewICoreWebView2ScriptDialogOpeningEventHandler(w)
		err := w.CoreWebView.AddScriptDialogOpening(w.HandlerScriptDialogOpeningEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbScriptDialogOpeningEvent = cb
	return nil
}

// Event_DevToolsProtocolEventReceived 是 DevTools 协议事件接收事件.
//   - DevToolsProtocolEventReceived 在收到来自 DevTools 协议的事件时运行。
func (w *WebViewEventImpl) Event_DevToolsProtocolEventReceived(eventName string, cb func(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr) error {
	var err error
	w.DevToolsProtocolEventReceiver, err = w.CoreWebView.GetDevToolsProtocolEventReceiver(eventName)
	if err != nil {
		return err
	}
	if w.HandlerDevToolsProtocolEventReceivedEvent == nil {
		w.HandlerDevToolsProtocolEventReceivedEvent = NewICoreWebView2DevToolsProtocolEventReceivedEventHandler(w)
	}
	err = w.DevToolsProtocolEventReceiver.AddDevToolsProtocolEventReceived(w.HandlerDevToolsProtocolEventReceivedEvent, w.EventRegistrationToken)
	if err != nil {
		return err
	}
	w.cbDevToolsProtocolEventReceivedEvent = cb
	return nil
}

// Event_WebResourceResponseReceived 是 Web 资源响应接收事件.
//   - 当 WebView 接收到对网络资源请求的响应时，会引发 WebResourceResponseReceived 事件（WebView执行的任何URI解析；例如HTTP/HTTPS、来自重定向、导航、HTML声明、隐式 favicon 图标的查找和数据请求，以及文档中 fetch API 的使用）。
//   - 宿主应用可以使用此事件查看网络资源的实际请求和响应。
//   - 无法保证 WebView 处理响应的顺序与宿主应用的处理程序运行的顺序。
//   - 应用的处理程序不会阻止 WebView 处理响应。
func (w *WebViewEventImpl) Event_WebResourceResponseReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr) error {
	w2_2, err := w.CoreWebView.GetICoreWebView2_2()
	if err != nil {
		return err
	}
	defer w2_2.Release()
	if w.HandlerWebResourceResponseReceivedEvent == nil {
		w.HandlerWebResourceResponseReceivedEvent = NewICoreWebView2WebResourceResponseReceivedEventHandler(w)
		w.TokenWebResourceResponseReceivedEvent, err = w2_2.AddWebResourceResponseReceived(w.HandlerWebResourceResponseReceivedEvent)
		if err != nil {
			return err
		}
	}
	w.cbWebResourceResponseReceivedEvent = cb
	return nil
}

// Event_DOMContentLoaded 是 DOM 内容加载完成事件.
//   - 当初始 HTML 文档解析完成时，会触发 DOMContentLoaded 事件。这与 HTML 中文档的 DOMContentLoaded 事件一致。
func (w *WebViewEventImpl) Event_DOMContentLoaded(cb func(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr) error {
	w2_2, err := w.CoreWebView.GetICoreWebView2_2()
	if err != nil {
		return err
	}
	defer w2_2.Release()
	if w.HandlerDOMContentLoadedEvent == nil {
		w.HandlerDOMContentLoadedEvent = NewICoreWebView2DOMContentLoadedEventHandler(w)
		w.TokenDOMContentLoadedEvent, err = w2_2.AddDomContentLoaded(w.HandlerDOMContentLoadedEvent)
		if err != nil {
			return err
		}
	}
	w.cbDOMContentLoadedEvent = cb
	return nil
}

// Event_FrameCreated 框架创建完成事件.
//   - 当创建新的 iframe 时触发。
//   - 处理此事件以访问 ICoreWebView2Frame 对象。
func (w *WebViewEventImpl) Event_FrameCreated(cb func(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr) error {
	w2_4, err := w.CoreWebView.GetICoreWebView2_4()
	if err != nil {
		return err
	}
	defer w2_4.Release()
	if w.HandlerFrameCreatedEvent == nil {
		w.HandlerFrameCreatedEvent = NewICoreWebView2FrameCreatedEventHandler(w)
		err = w2_4.AddFrameCreated(w.HandlerFrameCreatedEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbFrameCreatedEvent = cb
	return nil
}

// Event_FrameNameChanged 框架名称改变事件.
//   - 当 iframe 更改其 window.name 属性时触发。
func (w *WebViewEventImpl) Event_FrameNameChanged(frame *ICoreWebView2Frame, cb func(sender *ICoreWebView2Frame, args *IUnknown) uintptr) error {
	if w.HandlerFrameNameChangedEvent == nil {
		w.HandlerFrameNameChangedEvent = NewICoreWebView2FrameNameChangedEventHandler(w)
	}
	err := frame.AddNameChanged(w.HandlerFrameNameChangedEvent, w.EventRegistrationToken)
	if err != nil {
		return err
	}
	w.cbFrameNameChangedEvent = cb
	return nil
}

// Event_FrameDestroyed 框架销毁事件.
//   - 当与此 ICoreWebView2Frame 对象对应的 iframe 被移除或包含该 iframe 的文档被销毁时触发。
func (w *WebViewEventImpl) Event_FrameDestroyed(frame *ICoreWebView2Frame, cb func(sender *ICoreWebView2Frame, args *IUnknown) uintptr) error {
	if w.HandlerFrameDestroyedEvent == nil {
		w.HandlerFrameDestroyedEvent = NewICoreWebView2FrameDestroyedEventHandler(w)
	}
	err := frame.AddDestroyed(w.HandlerFrameDestroyedEvent, w.EventRegistrationToken)
	if err != nil {
		return err
	}
	w.cbFrameDestroyedEvent = cb
	return nil
}

// Event_DownloadStarting 下载开始事件.
//   - 当下载开始时触发，该事件会阻止默认的下载对话框弹出，但不会阻止下载进程。
//   - 主机可以选择取消下载、更改结果文件路径以及隐藏默认下载对话框。
//   - 如果主机选择取消下载，则不会保存下载内容，不会显示对话框，并且状态将更改为 COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED，中断原因是 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED.
//   - 否则，事件完成后，下载内容将保存到默认路径，如果主机未选择隐藏默认下载对话框，则会显示该对话框。
//   - 主机可以使用 Handled 属性更改下载对话框的可见性。
//   - 如果未处理该事件，下载将正常完成并显示默认对话框。
func (w *WebViewEventImpl) Event_DownloadStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr) error {
	w2_4, err := w.CoreWebView.GetICoreWebView2_4()
	if err != nil {
		return err
	}
	defer w2_4.Release()
	if w.HandlerDownloadStartingEvent == nil {
		w.HandlerDownloadStartingEvent = NewICoreWebView2DownloadStartingEventHandler(w)
		err = w2_4.AddDownloadStarting(w.HandlerDownloadStartingEvent, w.EventRegistrationToken)
		if err != nil {
			return err
		}
	}
	w.cbDownloadStartingEvent = cb
	return nil
}

// Event_BytesReceivedChanged 下载字节改变事件.
//   - 当下载的字节数发生更改时触发。
func (w *WebViewEventImpl) Event_BytesReceivedChanged(downloadOperation *ICoreWebView2DownloadOperation, cb func(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr) error {
	if w.HandlerBytesReceivedChangedEvent == nil {
		w.HandlerBytesReceivedChangedEvent = NewICoreWebView2BytesReceivedChangedEventHandler(w)
	}
	err := downloadOperation.AddBytesReceivedChanged(w.HandlerBytesReceivedChangedEvent, w.EventRegistrationToken)
	if err != nil {
		return err
	}
	w.cbBytesReceivedChangedEvent = cb
	return nil
}
