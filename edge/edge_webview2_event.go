package edge

// --------------------------- 事件 ---------------------------

// Event_DocumentTitleChanged 文档标题改变事件.
//   - 当 Webview 的 DocumentTitle 属性发生变化时, DocumentTitleChanged 会运行，并且可能在 NavigationCompleted 事件之前或之后运行。
func (w *WebViewEventImpl) Event_DocumentTitleChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "DocumentTitleChanged", c, nil, allowAddingMultiple...)
}

// Event_RasterizationScaleChanged 光栅化缩放比例改变事件.
//   - 当 Webview 检测到显示器 DPI 缩放比例已更改、ShouldDetectMonitorScaleChanges 为 true 且 Webview 已更改 RasterizationScale 属性时，将引发此事件。
func (w *WebViewEventImpl) Event_RasterizationScaleChanged(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "RasterizationScaleChanged", c, nil, allowAddingMultiple...)
}

// Event_WindowCloseRequested 窗口关闭请求事件.
//   - WindowCloseRequested 在 Webview 内部的内容请求关闭窗口时触发，例如在运行 window.close 之后。如果这对应用程序有意义，应用程序应该关闭 Webview 和相关的应用程序窗口。
//   - 在首次调用 window.close() 之后，对于任何紧接着连续调用的 window.close()，此事件可能不会触发。
func (w *WebViewEventImpl) Event_WindowCloseRequested(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "WindowCloseRequested", c, nil, allowAddingMultiple...)
}

// Event_SourceChanged 源改变事件.
//   - SourceChanged 会在 Source 属性发生变化时触发。
//   - SourceChanged 会在导航到不同站点或进行片段导航时运行。
//   - 对于其他类型的导航，例如页面刷新或 history.pushState（使用与当前页面相同的 URL），它不会触发。
//   - SourceChanged 会在导航到新文档时，在 ContentLoading 之前运行。
func (w *WebViewEventImpl) Event_SourceChanged(cb func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "SourceChanged", c, nil, allowAddingMultiple...)
}

// Event_NewWindowRequested 新窗口请求事件.
//   - 当 Webview 内的内容请求打开新窗口时（例如通过 NewWindowRequested）, NewWindowRequested 事件将运行。
//   - 应用可以传递一个目标 Webview 作为打开的窗口，或者将该事件标记为 Handled，在这种情况下，WebView2 不会打开窗口。
//   - 如果 Handled 或 NewWindow 属性均未设置，目标内容将在弹出窗口中打开。
func (w *WebViewEventImpl) Event_NewWindowRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "NewWindowRequested", c, nil, allowAddingMultiple...)
}

// Event_PermissionRequested 权限请求事件.
//   - PermissionRequested 在 Webview 中的内容请求访问某些特权资源的权限时运行。
func (w *WebViewEventImpl) Event_PermissionRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "PermissionRequested", c, nil, allowAddingMultiple...)
}

// Event_WebMessageReceived 网页消息事件.
//   - 当 WebView 的顶级文档运行 window.chrome.webview.postMessage 时, WebMessageReceived 事件会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何受 JSON 转换支持的对象。
func (w *WebViewEventImpl) Event_WebMessageReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "WebMessageReceived", c, nil, allowAddingMultiple...)
}

// Event_WebResourceRequested 网页资源请求事件.
func (w *WebViewEventImpl) Event_WebResourceRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "WebResourceRequested", c, nil, allowAddingMultiple...)
}

// Event_NavigationCompleted 导航完成事件.
//   - NavigationCompleted 事件会在 Webview 完全加载完毕（与 body.onload 事件同时发生）或加载因错误而停止时触发。
func (w *WebViewEventImpl) Event_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "NavigationCompleted", c, nil, allowAddingMultiple...)
}

// Event_Frame_NavigationCompleted 框架导航完成事件.
//   - 框架导航完成事件会在 Webview 中的子框架完全加载完毕（与 body.onload 触发同时）或加载因错误而停止时触发。
func (w *WebViewEventImpl) Event_Frame_NavigationCompleted(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "Frame_NavigationCompleted", c, nil, allowAddingMultiple...)
}

// Event_Frame_NavigationStarting 是 ICoreWebView2 的框架导航开始事件。
//   - 框架导航开始事件会在 Webview 中的子框架请求导航到不同 URI 的权限时触发。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有框架导航开始事件处理程序返回之前，导航将被阻止。
func (w *WebViewEventImpl) Event_Frame_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "Frame_NavigationStarting", c, nil, allowAddingMultiple...)
}

// Event_NavigationStarting 导航开始事件.
//   - NavigationStarting 在 Webview 主框架请求导航到不同的统一资源标识符 (URI) 权限时运行。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有 NavigationStarting 事件处理程序返回之前，导航将被阻止。
func (w *WebViewEventImpl) Event_NavigationStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "NavigationStarting", c, nil, allowAddingMultiple...)
}

// Event_AcceleratorKeyPressed 快捷键事件.
//   - AcceleratorKeyPressed 在 WebView 获得焦点时，当按下或释放快捷键或组合键时运行。
func (w *WebViewEventImpl) Event_AcceleratorKeyPressed(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "AcceleratorKeyPressed", c, nil, allowAddingMultiple...)
}

// Event_ContentLoading 网页内容正在加载事件.
//   - ContentLoading 会在任何内容加载之前触发，包括使用 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - ContentLoading 在发生相同页面导航时（例如通过 fragment 导航或 history.pushState 导航）不会触发。
//   - 此操作在 NavigationStarting 和 SourceChanged 事件之后，以及 HistoryChanged 和 NavigationCompleted 事件之前发生。
func (w *WebViewEventImpl) Event_ContentLoading(cb func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ContentLoading", c, nil, allowAddingMultiple...)
}

// Event_ContainsFullScreenElementChanged 全屏元素状态改变事件.
//   - 当 ContainsFullScreenElement 属性发生变化时触发。
func (w *WebViewEventImpl) Event_ContainsFullScreenElementChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ContainsFullScreenElementChanged", c, nil, allowAddingMultiple...)
}

// Event_ProcessFailed 进程失败事件.
//   - 当 WebView 的进程因任何原因失败时, ProcessFailed 事件将运行。
func (w *WebViewEventImpl) Event_ProcessFailed(cb func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ProcessFailed", c, nil, allowAddingMultiple...)
}

// Event_HistoryChanged 历史记录改变事件.
//   - HistoryChanged 在联合会话历史记录发生更改时引发，该历史记录由顶级和手动框架导航组成。
func (w *WebViewEventImpl) Event_HistoryChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "HistoryChanged", c, nil, allowAddingMultiple...)
}

// Event_ScriptDialogOpening 脚本对话框打开事件.
//   - 当 JavaScript 代码调用 alert、confirm、prompt 或 beforeunload 时触发此事件。
//   - 需 ICoreWebView2Settings.SetAreDefaultScriptDialogsEnabled(false) 后，此事件才会触发。
//   - ScriptDialogOpening 事件会抑制对话框，或使用自定义对话框替换默认对话框。
func (w *WebViewEventImpl) Event_ScriptDialogOpening(cb func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ScriptDialogOpening", c, nil, allowAddingMultiple...)
}

// Event_WebResourceResponseReceived 是 Web 资源响应接收事件.
//   - 当 WebView 接收到对网络资源请求的响应时，会引发 WebResourceResponseReceived 事件（WebView执行的任何URI解析；例如HTTP/HTTPS、来自重定向、导航、HTML声明、隐式 favicon 图标的查找和数据请求，以及文档中 fetch API 的使用）。
//   - 宿主应用可以使用此事件查看网络资源的实际请求和响应。
//   - 无法保证 WebView 处理响应的顺序与宿主应用的处理程序运行的顺序。
//   - 应用的处理程序不会阻止 WebView 处理响应。
func (w *WebViewEventImpl) Event_WebResourceResponseReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "WebResourceResponseReceived", c, nil, allowAddingMultiple...)
}

// Event_DOMContentLoaded 是 DOM 内容加载完成事件.
//   - 当初始 HTML 文档解析完成时，会触发 DOMContentLoaded 事件。这与 HTML 中文档的 DOMContentLoaded 事件一致。
func (w *WebViewEventImpl) Event_DOMContentLoaded(cb func(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "DOMContentLoaded", c, nil, allowAddingMultiple...)
}

// Event_FrameCreated 框架创建完成事件.
//   - 当创建新的 iframe 时触发。
//   - 处理此事件以访问 ICoreWebView2Frame 对象。
func (w *WebViewEventImpl) Event_FrameCreated(cb func(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "FrameCreated", c, nil, allowAddingMultiple...)
}

// Event_DownloadStarting 下载开始事件.
//   - 当下载开始时触发，该事件会阻止默认的下载对话框弹出，但不会阻止下载进程。
//   - 主机可以选择取消下载、更改结果文件路径以及隐藏默认下载对话框。
//   - 如果主机选择取消下载，则不会保存下载内容，不会显示对话框，并且状态将更改为 COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED，中断原因是 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED.
//   - 否则，事件完成后，下载内容将保存到默认路径，如果主机未选择隐藏默认下载对话框，则会显示该对话框。
//   - 主机可以使用 Handled 属性更改下载对话框的可见性。
//   - 如果未处理该事件，下载将正常完成并显示默认对话框。
func (w *WebViewEventImpl) Event_DownloadStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "DownloadStarting", c, nil, allowAddingMultiple...)
}

// Event_ClientCertificateRequested 客户端证书请求事件.
//   - 当 WebView2 向需要客户端证书进行 HTTP 身份验证的 HTTP 服务器发出请求时触发.
//   - 如果不处理该事件， WebView2 将向用户显示默认的客户端证书选择对话框提示。
func (w *WebViewEventImpl) Event_ClientCertificateRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ClientCertificateRequested", c, nil, allowAddingMultiple...)
}

// Event_IsMutedChanged 静音状态改变事件.
//   - 当 WebView 的静音状态更改时调用。
func (w *WebViewEventImpl) Event_IsMutedChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "IsMutedChanged", c, nil, allowAddingMultiple...)
}

// Event_DocumentPlayingAudioChanged 文档播放音频状态改变事件.
func (w *WebViewEventImpl) Event_DocumentPlayingAudioChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "DocumentPlayingAudioChanged", c, nil, allowAddingMultiple...)
}

// Event_ContextMenuRequested 上下文菜单请求事件.
//   - 当用户请求上下文菜单，且 WebView 内部的内容未禁用上下文菜单时，将引发 ContextMenuRequested 事件。
//   - 宿主可以选择使用事件中提供的信息创建自己的上下文菜单，也可以向 WebView 上下文菜单添加或删除菜单项。
//   - 如果宿主不处理该事件， WebView 将显示默认上下文菜单。
func (w *WebViewEventImpl) Event_ContextMenuRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ContextMenuRequested", c, nil, allowAddingMultiple...)
}

// Event_BrowserProcessExited 浏览器进程退出事件.
//   - 当与 WebView2 关联的浏览器进程退出时触发此事件。
func (w *WebViewEventImpl) Event_BrowserProcessExited(cb func(sender *ICoreWebView2, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "BrowserProcessExited", c, nil, allowAddingMultiple...)
}

// Event_ProcessInfosChanged 进程信息变更事件.
//   - 当 WebView2 运行时的进程列表发生变化时触发
func (w *WebViewEventImpl) Event_ProcessInfosChanged(cb func(sender *ICoreWebView2Environment, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ProcessInfosChanged", c, nil, allowAddingMultiple...)
}

// Event_FaviconChanged 网站图标改变事件.
//   - 当网站图标（favicon）的 URL 与之前的 URL 不同时，会触发 FaviconChanged 事件。
//   - 首次导航到新文档时，无论该文档是否在 HTML 中声明了网站图标，只要其图标与之前的图标不同，就会触发 FaviconChanged 事件。
//   - 如果 HTML 中声明了网站图标或通过脚本设置了网站图标，该事件将再次触发。随后可以通过 GetFavicon 和 GetFaviconUri 方法获取网站图标信息。
func (w *WebViewEventImpl) Event_FaviconChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "FaviconChanged", c, nil, allowAddingMultiple...)
}

// Event_ZoomFactorChanged 缩放因子改变事件.
//   - 当 WebView 的 ZoomFactor 属性发生改变时触发。
//   - WebView 会关联每个网站最后使用的缩放比例。导航到不同页面时，缩放比例可能会发生变化。当因导航变化导致缩放比例改变时，ZoomFactorChanged 事件会在 ContentLoading 事件之后立即运行。
func (w *WebViewEventImpl) Event_ZoomFactorChanged(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ZoomFactorChanged", c, nil, allowAddingMultiple...)
}

// Event_MoveFocusRequested 移动焦点请求事件.
//   - MoveFocusRequested 在用户尝试按 Tab 键离开 WebView 时运行。此事件运行时，WebView 的焦点尚未改变。
func (w *WebViewEventImpl) Event_MoveFocusRequested(cb func(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "MoveFocusRequested", c, nil, allowAddingMultiple...)
}

// Event_GotFocus 获得焦点事件.
//   - 当 WebView 获得焦点时触发。
func (w *WebViewEventImpl) Event_GotFocus(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "FocusChanged", c, nil, allowAddingMultiple...)
}

// Event_LostFocus 失去焦点事件.
//   - 当 WebView 失去焦点时触发。
func (w *WebViewEventImpl) Event_LostFocus(cb func(sender *ICoreWebView2Controller, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "LostFocus", c, nil, allowAddingMultiple...)
}

// Event_NewBrowserVersionAvailable 新浏览器版本可用事件.
//   - 当有新版本的 WebView2 运行时可用时触发。
func (w *WebViewEventImpl) Event_NewBrowserVersionAvailable(cb func(sender *ICoreWebView2Environment, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "NewBrowserVersionAvailable", c, nil, allowAddingMultiple...)
}

// Event_IsDefaultDownloadDialogOpenChanged 默认下载对话框打开状态变化事件.
//   - 此事件在 DownloadStarting 事件之后发生。
//   - 在 DownloadStartingEventArgs 上设置 Handled 属性会禁用默认下载对话框，并确保此事件永远不会被触发。
func (w *WebViewEventImpl) Event_IsDefaultDownloadDialogOpenChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "IsDefaultDownloadDialogOpenChanged", c, nil, allowAddingMultiple...)
}

// Event_BasicAuthenticationRequested 基本身份验证请求事件.
//   - 主机可以提供包含身份验证凭据的响应，也可以取消请求。
//   - 如果主机将 Cancel 属性设置为 false，但未在 Response 属性上提供 UserName 或 Password 属性，那么 WebView2 将向用户显示默认的身份验证质询对话框提示。
func (w *WebViewEventImpl) Event_BasicAuthenticationRequested(cb func(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "BasicAuthenticationRequested", c, nil, allowAddingMultiple...)
}

// Event_StatusBarTextChanged 状态栏文本改变事件.
//   - 当 WebView 显示状态消息、URL 或空字符串（表示隐藏状态栏）时触发。
func (w *WebViewEventImpl) Event_StatusBarTextChanged(cb func(sender *ICoreWebView2, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "StatusBarTextChanged", c, nil, allowAddingMultiple...)
}

// Event_ServerCertificateErrorDetected 检测到服务器证书错误事件.
//   - 当 WebView2 在加载网页时无法验证服务器的数字证书时触发。
//   - 此事件将针对所有网络资源触发，并紧随 WebResourceRequested 事件之后。
//   - 如果不处理该事件，WebView2 会在导航时向用户显示默认的 TLS 插页式错误页面，而对于非导航操作，Web 请求会被取消。
func (w *WebViewEventImpl) Event_ServerCertificateErrorDetected(cb func(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ServerCertificateErrorDetected", c, nil, allowAddingMultiple...)
}

// Event_LaunchingExternalUriScheme 启动外部 URI 方案事件。
//   - 当导航请求指向已在操作系统中注册的 URI 方案时触发。
//   - 事件处理程序可以抑制默认对话框，或者用自定义对话框替换默认对话框。
//   - 如果未对事件参数执行延迟操作，外部 URI 方案的启动将被阻止，直到事件处理程序返回。
//   - 如果执行了延迟操作，外部 URI 方案的启动将被阻止，直到延迟完成。主机还可以选择取消 URI 方案的启动。
//   - 无论 Cancel 属性设置为 TRUE 还是 FALSE，都会触发 NavigationStarting 和 NavigationCompleted 事件。
//   - 无论主机是否在 ICoreWebView2LaunchingExternalUriSchemeEventArgs 上设置 Cancel 属性，NavigationCompleted 事件触发时，IsSuccess 属性都将设置为 FALSE，WebErrorStatus 属性都将设置为 ConnectionAborted。
//   - 对于此次导航到外部 URI 方案，无论 Cancel 属性如何设置，都不会触发 SourceChanged、ContentLoading 和 HistoryChanged 事件。
//   - LaunchingExternalUriScheme 事件将在 NavigationStarting 事件之后、NavigationCompleted 事件之前触发。
//   - 导航到外部 URI 方案时，默认的 ICoreWebView2Settings 也会更新。如果 ICoreWebView2Settings 接口上的某个设置已更改，导航到外部 URI 方案将触发 ICoreWebView2Settings 更新。
func (w *WebViewEventImpl) Event_LaunchingExternalUriScheme(cb func(sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "LaunchingExternalUriScheme", c, nil, allowAddingMultiple...)
}

// Event_NotificationReceived 通知接收事件。
//   - 如果未对事件参数执行延迟操作，那么在 DOM 通知创建调用（即 Notification()）之后的后续脚本会被阻塞，直到事件处理程序返回。
//   - 如果执行了延迟操作，脚本会被阻塞，直到延迟完成。
func (w *WebViewEventImpl) Event_NotificationReceived(cb func(sender *ICoreWebView2, args *ICoreWebView2NotificationReceivedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "NotificationReceived", c, nil, allowAddingMultiple...)
}

// Event_SaveAsUIShowing 另存为界面显示事件。
//   - 当通过编程方式或手动方式触发“另存为”时触发。
func (w *WebViewEventImpl) Event_SaveAsUIShowing(cb func(sender *ICoreWebView2, args *ICoreWebView2SaveAsUIShowingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "SaveAsUIShowing", c, nil, allowAddingMultiple...)
}

// Event_SaveFileSecurityCheckStarting 保存文件安全检查开始事件
//   - 在系统 FileTypePolicy 检查危险文件扩展名列表期间，将触发此事件。
func (w *WebViewEventImpl) Event_SaveFileSecurityCheckStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "SaveFileSecurityCheckStarting", c, nil, allowAddingMultiple...)
}

// Event_ScreenCaptureStarting 屏幕截图开始事件。
func (w *WebViewEventImpl) Event_ScreenCaptureStarting(cb func(sender *ICoreWebView2, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	return WvEventHandler.AddCallBack(w, "ScreenCaptureStarting", c, nil, allowAddingMultiple...)
}
