package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2 是 WebView2 核心.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2
type ICoreWebView2 struct {
	Vtbl *ICoreWebView2Vtbl
}

type ICoreWebView2Vtbl struct {
	IUnknownVtbl
	GetSettings                            ComProc
	GetSource                              ComProc
	Navigate                               ComProc
	NavigateToString                       ComProc
	AddNavigationStarting                  ComProc
	RemoveNavigationStarting               ComProc
	AddContentLoading                      ComProc
	RemoveContentLoading                   ComProc
	AddSourceChanged                       ComProc
	RemoveSourceChanged                    ComProc
	AddHistoryChanged                      ComProc
	RemoveHistoryChanged                   ComProc
	AddNavigationCompleted                 ComProc
	RemoveNavigationCompleted              ComProc
	AddFrameNavigationStarting             ComProc
	RemoveFrameNavigationStarting          ComProc
	AddFrameNavigationCompleted            ComProc
	RemoveFrameNavigationCompleted         ComProc
	AddScriptDialogOpening                 ComProc
	RemoveScriptDialogOpening              ComProc
	AddPermissionRequested                 ComProc
	RemovePermissionRequested              ComProc
	AddProcessFailed                       ComProc
	RemoveProcessFailed                    ComProc
	AddScriptToExecuteOnDocumentCreated    ComProc
	RemoveScriptToExecuteOnDocumentCreated ComProc
	ExecuteScript                          ComProc
	CapturePreview                         ComProc
	Reload                                 ComProc
	PostWebMessageAsJSON                   ComProc
	PostWebMessageAsString                 ComProc
	AddWebMessageReceived                  ComProc
	RemoveWebMessageReceived               ComProc
	CallDevToolsProtocolMethod             ComProc
	GetBrowserProcessID                    ComProc
	GetCanGoBack                           ComProc
	GetCanGoForward                        ComProc
	GoBack                                 ComProc
	GoForward                              ComProc
	GetDevToolsProtocolEventReceiver       ComProc
	Stop                                   ComProc
	AddNewWindowRequested                  ComProc
	RemoveNewWindowRequested               ComProc
	AddDocumentTitleChanged                ComProc
	RemoveDocumentTitleChanged             ComProc
	GetDocumentTitle                       ComProc
	AddHostObjectToScript                  ComProc
	RemoveHostObjectFromScript             ComProc
	OpenDevToolsWindow                     ComProc
	AddContainsFullScreenElementChanged    ComProc
	RemoveContainsFullScreenElementChanged ComProc
	GetContainsFullScreenElement           ComProc
	AddWebResourceRequested                ComProc
	RemoveWebResourceRequested             ComProc
	AddWebResourceRequestedFilter          ComProc
	RemoveWebResourceRequestedFilter       ComProc
	AddWindowCloseRequested                ComProc
	RemoveWindowCloseRequested             ComProc
	// 2
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDomContentLoaded               ComProc
	RemoveDomContentLoaded            ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc
	// 3
	TrySuspend                          ComProc
	Resume                              ComProc
	GetIsSuspended                      ComProc
	SetVirtualHostNameToFolderMapping   ComProc
	ClearVirtualHostNameToFolderMapping ComProc
	// 4
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
	// 5
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
	// 6
	OpenTaskManagerWindow ComProc
	// 7
	PrintToPdf ComProc
	// 8
	AddIsMutedChanged                   ComProc
	RemoveIsMutedChanged                ComProc
	GetIsMuted                          ComProc
	PutIsMuted                          ComProc
	AddIsDocumentPlayingAudioChanged    ComProc
	RemoveIsDocumentPlayingAudioChanged ComProc
	GetIsDocumentPlayingAudio           ComProc
	// 9
	AddIsDefaultDownloadDialogOpenChanged    ComProc
	RemoveIsDefaultDownloadDialogOpenChanged ComProc
	GetIsDefaultDownloadDialogOpen           ComProc
	OpenDefaultDownloadDialog                ComProc
	CloseDefaultDownloadDialog               ComProc
	GetDefaultDownloadDialogCornerAlignment  ComProc
	PutDefaultDownloadDialogCornerAlignment  ComProc
	GetDefaultDownloadDialogMargin           ComProc
	PutDefaultDownloadDialogMargin           ComProc
	// 10
	AddBasicAuthenticationRequested    ComProc
	RemoveBasicAuthenticationRequested ComProc
	// 11
	CallDevToolsProtocolMethodForSession ComProc
	AddContextMenuRequested              ComProc
	RemoveContextMenuRequested           ComProc
	// 12
	AddStatusBarTextChanged    ComProc
	RemoveStatusBarTextChanged ComProc
	GetStatusBarText           ComProc
	// 13
	GetProfile ComProc
	// 14
	AddServerCertificateErrorDetected    ComProc
	RemoveServerCertificateErrorDetected ComProc
	ClearServerCertificateErrorActions   ComProc
	// 15
	AddFaviconChanged    ComProc
	RemoveFaviconChanged ComProc
	GetFaviconUri        ComProc
	GetFavicon           ComProc
	// 16
	Print            ComProc
	ShowPrintUI      ComProc
	PrintToPdfStream ComProc
	// 17
	PostSharedBufferToScript ComProc
	// 18
	AddLaunchingExternalUriScheme    ComProc
	RemoveLaunchingExternalUriScheme ComProc
	// 19
	GetMemoryUsageTargetLevel ComProc
	PutMemoryUsageTargetLevel ComProc
	// 20
	GetFrameId ComProc
	// 21
	ExecuteScriptWithResult ComProc
	// 22
	AddWebResourceRequestedFilterWithRequestSourceKinds    ComProc
	RemoveWebResourceRequestedFilterWithRequestSourceKinds ComProc
	// 23
	PostWebMessageAsJsonWithAdditionalObjects ComProc
	// 24
	AddNotificationReceived    ComProc
	RemoveNotificationReceived ComProc
	// 25
	AddSaveAsUIShowing    ComProc
	RemoveSaveAsUIShowing ComProc
	ShowSaveAsUI          ComProc
	// 26
	AddSaveFileSecurityCheckStarting    ComProc
	RemoveSaveFileSecurityCheckStarting ComProc
	// 27
	AddScreenCaptureStarting    ComProc
	RemoveScreenCaptureStarting ComProc
	// 28
	GetFind ComProc
}

func (i *ICoreWebView2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSettings 获取 ICoreWebView2Settings 对象, 它包含正在运行的 WebView 的各种可修改设置。
func (i *ICoreWebView2) GetSettings() (*ICoreWebView2Settings, error) {
	var settings *ICoreWebView2Settings
	r, _, _ := i.Vtbl.GetSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&settings)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return settings, nil
}

// AddWebResourceRequestedFilter 添加 web 资源请求过滤器. [已过时]
//   - 请改用 ICoreWebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds，该方法将针对文档中的所有 iframe 触发此事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter
func (i *ICoreWebView2) AddWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.AddWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebResourceRequestedFilter 移除 web 资源请求过滤器.
func (i *ICoreWebView2) RemoveWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.RemoveWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// EventRegistrationToken 事件注册令牌.
type EventRegistrationToken struct {
	Value int64
}

// AddNavigationCompleted 添加导航完成事件处理程序.
//   - NavigationCompleted 事件会在 Webview 完全加载完毕（与 body.onload 事件同时发生）或加载因错误而停止时触发。
func (i *ICoreWebView2) AddNavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNavigationCompleted 移除导航完成事件处理程序。
func (i *ICoreWebView2) RemoveNavigationCompleted(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Navigate 导航 webview 到给定的 URL。
//   - URL 可能是数据 URI，即 "data:text/html,<html>...</html>"。
//   - 通常不进行适当的 url 编码也是可以的, webview 会为你重新编码。
func (i *ICoreWebView2) Navigate(uri string) error {
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.Navigate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// NavigateToString 直接设置 webview 的 HTML。
// 页面的来源是 `about:blank`。
//
// htmlContent: 参数的总大小不能大于2 MB（2*1024*1024字节）.
func (i *ICoreWebView2) NavigateToString(htmlContent string) error {
	_htmlContent, err := syscall.UTF16PtrFromString(htmlContent)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.NavigateToString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_htmlContent)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSource 获取当前顶级文档的URI。如果导航正在进行中，则返回即将导航到的URI。
func (i *ICoreWebView2) GetSource() (string, error) {
	var _uri *uint16
	r, _, _ := i.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(_uri)
	wapi.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// AddNavigationStarting 添加导航开始事件处理程序。
//   - NavigationStarting 在 Webview 主框架请求导航到不同的统一资源标识符 (URI) 权限时运行。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有 NavigationStarting 事件处理程序返回之前，导航将被阻止。
func (i *ICoreWebView2) AddNavigationStarting(eventHandler *ICoreWebView2NavigationStartingEventHandler, token *EventRegistrationToken) error {
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
func (i *ICoreWebView2) RemoveNavigationStarting(token EventRegistrationToken) error {
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
//   - ContentLoading 会在任何内容加载之前触发，包括使用 AddScriptToExecuteOnDocumentCreated 添加的脚本。
//   - ContentLoading 在发生相同页面导航时（例如通过 fragment 导航或 history.pushState 导航）不会触发。
//   - 此操作在 NavigationStarting 和 SourceChanged 事件之后，以及 HistoryChanged 和 NavigationCompleted 事件之前发生。
func (i *ICoreWebView2) AddContentLoading(eventHandler *ICoreWebView2ContentLoadingEventHandler, token *EventRegistrationToken) error {
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
func (i *ICoreWebView2) RemoveContentLoading(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddWebResourceRequested 添加 web 资源请求事件处理程序。
func (i *ICoreWebView2) AddWebResourceRequested(eventHandler *ICoreWebView2WebResourceRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebResourceRequested 移除 web 资源请求事件处理程序。
func (i *ICoreWebView2) RemoveWebResourceRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddWebMessageReceived 添加 web 消息接收事件处理程序。
//   - 当 Webview 的顶级文档运行 window.chrome.webview.postMessage 时，WebMessageReceived 事件会运行。
//   - postMessage 函数为 void postMessage(object)，其中 object 是任何受 JSON 转换支持的对象。
func (i *ICoreWebView2) AddWebMessageReceived(eventHandler *ICoreWebView2WebMessageReceivedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebMessageReceived 移除 web 消息接收事件处理程序。
func (i *ICoreWebView2) RemoveWebMessageReceived(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddPermissionRequested 添加权限请求事件处理程序。
//   - PermissionRequested 在 Webview 中的内容请求访问某些特权资源的权限时运行。
func (i *ICoreWebView2) AddPermissionRequested(eventHandler *ICoreWebView2PermissionRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemovePermissionRequested 移除权限请求事件处理程序。
func (i *ICoreWebView2) RemovePermissionRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemovePermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PostWebMessageAsString 将指定的 webMessage 发布到此 WebView 中的顶级文档。
//   - 其行为方式与 PostWebMessageAsJson 完全相同.
//   - 但 window.chrome.webview 消息的事件参数的 data 属性是一个字符串，其值与 webMessageAsString 相同。
//   - 如果希望使用简单字符串而非 JSON 对象进行通信，请使用此方法代替 PostWebMessageAsJson。
//
// webMessageAsString: JavaScript 对象的简单字符串，而不是 JSON 字符串表示。
func (i *ICoreWebView2) PostWebMessageAsString(webMessageAsString string) error {
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

// PostWebMessageAsJSON 将指定的 webMessage 发布到此 WebView 中的顶级文档。
//   - 通过订阅页面文档的 message 事件（属于window.chrome.webview）来接收消息。
//   - window.chrome.webview.addEventListener('message', handler)
//   - 事件参数是 MessageEvent 的实例。
//   - 事件参数的 data 属性是将 webMessage 字符串参数解析为 JSON 字符串后得到的 JavaScript 对象。
//   - 事件参数的 source 属性是对 window.chrome.webview 对象的引用。
//   - 消息以异步方式传递。如果在消息发布到页面之前发生导航，该消息将被丢弃。
//
// webMessageAsJSON: JavaScript 对象的 JSON 字符串.
func (i *ICoreWebView2) PostWebMessageAsJSON(webMessageAsJSON string) error {
	_webMessageAsJSON, err := syscall.UTF16PtrFromString(webMessageAsJSON)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostWebMessageAsJSON.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJSON)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Reload 重新加载当前页面。
func (i *ICoreWebView2) Reload() error {
	r, _, _ := i.Vtbl.Reload.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GoBack 将 WebView 导航到导航历史记录中的上一页。
func (i *ICoreWebView2) GoBack() error {
	r, _, _ := i.Vtbl.GoBack.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GoForward 将 WebView 导航到导航历史记录中的下一页。
func (i *ICoreWebView2) GoForward() error {
	r, _, _ := i.Vtbl.GoForward.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Stop 停止所有导航和挂起的资源获取。不停止脚本。
func (i *ICoreWebView2) Stop() error {
	r, _, _ := i.Vtbl.Stop.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ExecuteScript 在主页面上下文中执行 JavaScript 代码。
//
// javaScript: 要执行的 JavaScript 代码。
//
// handler: 执行完成后的回调处理程序，可以为 nil。如果 JavaScript 返回值，将通过 handler 返回。
func (i *ICoreWebView2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) error {
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

// ExecuteScriptEx 在主页面上下文中执行 JavaScript 代码。
//
// impl: *WebViewEventImpl。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 执行完成后的回调处理程序，可以为 nil。如果 JavaScript 返回值，将通过 cb 返回。
func (i *ICoreWebView2) ExecuteScriptEx(impl *WebViewEventImpl, javaScript string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "ExecuteScriptCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "ExecuteScriptCompleted")
	return i.ExecuteScript(javaScript, (*ICoreWebView2ExecuteScriptCompletedHandler)(handler))
}

// OpenDevToolsWindow 打开开发者工具窗口。
func (i *ICoreWebView2) OpenDevToolsWindow() error {
	r, _, _ := i.Vtbl.OpenDevToolsWindow.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetContainsFullScreenElement 获取 WebView 是否包含全屏元素。
func (i *ICoreWebView2) GetContainsFullScreenElement() (bool, error) {
	var containsFullScreenElement bool
	r, _, _ := i.Vtbl.GetContainsFullScreenElement.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&containsFullScreenElement)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return containsFullScreenElement, nil
}

// GetDocumentTitle 获取当前顶级文档的标题。
func (i *ICoreWebView2) GetDocumentTitle() (string, error) {
	var title *uint16
	r, _, _ := i.Vtbl.GetDocumentTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&title)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(title)
	wapi.CoTaskMemFree(unsafe.Pointer(title))
	return result, nil
}

// GetBrowserProcessID 获取承载 WebView 的浏览器进程的 ID。
func (i *ICoreWebView2) GetBrowserProcessID() (uint32, error) {
	var pid uint32
	r, _, _ := i.Vtbl.GetBrowserProcessID.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pid)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return pid, nil
}

// GetCanGoBack 获取 WebView 是否可以导航到上一页。
func (i *ICoreWebView2) GetCanGoBack() (bool, error) {
	var canGoBack bool
	r, _, _ := i.Vtbl.GetCanGoBack.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canGoBack)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return canGoBack, nil
}

// GetCanGoForward 获取 WebView 是否可以导航到下一页。
func (i *ICoreWebView2) GetCanGoForward() (bool, error) {
	var canGoForward bool
	r, _, _ := i.Vtbl.GetCanGoForward.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canGoForward)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return canGoForward, nil
}

// AddScriptToExecuteOnDocumentCreated 添加在创建文档时要执行的脚本。
//
// javaScript: 要执行的 JavaScript 代码。
//
// handler: 添加脚本完成后的回调处理程序, 可为 nil。
func (i *ICoreWebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, handler *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) error {
	_javaScript, err := syscall.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.AddScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddScriptToExecuteOnDocumentCreated 添加在创建文档时要执行的脚本。
//
// impl: *WebViewEventImpl。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 添加脚本完成后的回调处理程序。
func (i *ICoreWebView2) AddScriptToExecuteOnDocumentCreatedEx(impl *WebViewEventImpl, javaScript string, cb func(errorCode syscall.Errno, id string) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "AddScriptToExecuteOnDocumentCreatedCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "AddScriptToExecuteOnDocumentCreatedCompleted")
	return i.AddScriptToExecuteOnDocumentCreated(javaScript, (*ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler)(handler))
}

// RemoveScriptToExecuteOnDocumentCreated 移除在创建文档时要执行的脚本。
//
// id: 要移除的脚本的 ID。
func (i *ICoreWebView2) RemoveScriptToExecuteOnDocumentCreated(id string) error {
	_id, err := syscall.UTF16PtrFromString(id)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.RemoveScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_id)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddNewWindowRequested 添加新窗口请求事件处理程序.
//   - 当 Webview 内的内容请求打开新窗口时（例如通过 NewWindowRequested），NewWindowRequested 事件将运行。
//   - 应用可以传递一个被视为已打开窗口的目标 Webview，或者将该事件标记为 Handled，在此情况下，WebView2 不会打开窗口。
//   - 如果 Handled 或 NewWindow 属性均未设置，目标内容将在弹出窗口中打开。
func (i *ICoreWebView2) AddNewWindowRequested(eventHandler *ICoreWebView2NewWindowRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNewWindowRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNewWindowRequested 移除新窗口请求事件处理程序.
func (i *ICoreWebView2) RemoveNewWindowRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNewWindowRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddSourceChanged 添加源改变事件处理程序.
//   - SourceChanged 会在 Source 属性发生变化时触发。
//   - SourceChanged 会在导航到不同站点或进行片段导航时运行。
//   - 对于其他类型的导航，例如页面刷新或 history.pushState（使用与当前页面相同的 URL），它不会触发。
//   - SourceChanged 会在导航到新文档时，在 ContentLoading 之前运行。
func (i *ICoreWebView2) AddSourceChanged(eventHandler *ICoreWebView2SourceChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddSourceChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveSourceChanged 移除源改变事件处理程序.
func (i *ICoreWebView2) RemoveSourceChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveSourceChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddFrameNavigationStarting 添加框架导航开始事件处理程序。
//   - 框架导航开始事件会在 Webview 中的子框架请求导航到不同 URI 的权限时触发。重定向也会触发此操作，并且导航 ID 与原始 ID 相同。
//   - 在所有框架导航开始事件处理程序返回之前，导航将被阻止。
func (i *ICoreWebView2) AddFrameNavigationStarting(eventHandler *ICoreWebView2_Frame_NavigationStartingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddFrameNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveFrameNavigationStarting 移除框架导航开始事件处理程序。
func (i *ICoreWebView2) RemoveFrameNavigationStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveFrameNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddFrameNavigationCompleted 添加框架导航完成事件处理程序。
//   - 框架导航完成事件会在 Webview 中的子框架完全加载完毕（与 body.onload 触发同时）或加载因错误而停止时触发。
func (i *ICoreWebView2) AddFrameNavigationCompleted(eventHandler *ICoreWebView2_Frame_NavigationCompletedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddFrameNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveFrameNavigationCompleted 移除框架导航完成事件处理程序。
func (i *ICoreWebView2) RemoveFrameNavigationCompleted(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveFrameNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddWindowCloseRequested 添加窗口关闭请求事件处理程序.
//   - WindowCloseRequested 在 Webview 内部的内容请求关闭窗口时触发，例如在运行 window.close 之后。如果这对应用程序有意义，应用程序应该关闭 Webview 和相关的应用程序窗口。
//   - 在首次调用 window.close() 之后，对于任何紧接着连续调用的 window.close()，此事件可能不会触发。
func (i *ICoreWebView2) AddWindowCloseRequested(eventHandler *ICoreWebView2WindowCloseRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddWindowCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWindowCloseRequested 移除窗口关闭请求事件处理程序.
func (i *ICoreWebView2) RemoveWindowCloseRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveWindowCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddHostObjectToScript 向脚本添加主机对象。
//
// name: 对象名称，将在 JavaScript 中作为 window.chrome.webview.hostObjects.sync.name 或 window.chrome.webview.hostObjects.name 访问。
//
// object: 要添加的对象指针, VARIANT *object.
//
// todo: 这个对象的格式在 go 中还没测试明白. 需结合 go-ole 库, 在对象中嵌入 ole.IDispatch, 并实现其方法.
func (i *ICoreWebView2) AddHostObjectToScript(name string, object unsafe.Pointer) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.AddHostObjectToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(object),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveHostObjectFromScript 从脚本中移除主机对象。
//
// name: 要移除的对象名称.
func (i *ICoreWebView2) RemoveHostObjectFromScript(name string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.RemoveHostObjectFromScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddDocumentTitleChanged 添加文档标题改变事件处理程序.
//   - 当 Webview 的 DocumentTitle 属性发生变化时，DocumentTitleChanged 会运行，并且可能在 NavigationCompleted 事件之前或之后运行。
func (i *ICoreWebView2) AddDocumentTitleChanged(eventHandler *ICoreWebView2DocumentTitleChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddDocumentTitleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDocumentTitleChanged 移除文档标题改变事件处理程序.
func (i *ICoreWebView2) RemoveDocumentTitleChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveDocumentTitleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CapturePreview 捕获 Webview 的预览图像。
//
// imageFormat: 图像格式枚举值.
//
// imageStream: 接收图像数据的流对象.
//
// handler: 捕获完成后的回调处理程序. 写入流完毕后触发.
func (i *ICoreWebView2) CapturePreview(imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, imageStream *IStream, handler *ICoreWebView2CapturePreviewCompletedHandler) error {
	r, _, _ := i.Vtbl.CapturePreview.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(imageFormat),
		uintptr(unsafe.Pointer(imageStream)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CapturePreviewEx 捕获 Webview 的预览图像。
//
// impl: *WebViewEventImpl.
//
// imageFormat: 图像格式枚举值.
//
// imageStream: 接收图像数据的流对象.
//
// cb: 捕获完成后的回调处理程序. 写入流完毕后触发.
func (i *ICoreWebView2) CapturePreviewEx(impl *WebViewEventImpl, imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, imageStream *IStream, cb func(errorCode syscall.Errno) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "CapturePreviewCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "CapturePreviewCompleted")
	return i.CapturePreview(imageFormat, imageStream, (*ICoreWebView2CapturePreviewCompletedHandler)(handler))
}

// AddContainsFullScreenElementChanged 添加全屏元素状态改变事件处理程序.
func (i *ICoreWebView2) AddContainsFullScreenElementChanged(eventHandler *ICoreWebView2ContainsFullScreenElementChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveContainsFullScreenElementChanged 移除是否包含全屏元素属性改变事件处理程序.
func (i *ICoreWebView2) RemoveContainsFullScreenElementChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddProcessFailed 添加进程失败事件处理程序.
//   - 当 WebView 的浏览器进程意外终止时触发此事件.
func (i *ICoreWebView2) AddProcessFailed(eventHandler *ICoreWebView2ProcessFailedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddProcessFailed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveProcessFailed 移除进程失败事件处理程序.
func (i *ICoreWebView2) RemoveProcessFailed(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveProcessFailed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddHistoryChanged 添加历史记录改变事件处理程序.
//   - HistoryChanged 在联合会话历史记录发生更改时引发，该历史记录由顶级和手动框架导航组成.
func (i *ICoreWebView2) AddHistoryChanged(eventHandler *ICoreWebView2HistoryChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddHistoryChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveHistoryChanged 移除历史记录改变事件处理程序.
func (i *ICoreWebView2) RemoveHistoryChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveHistoryChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddScriptDialogOpening 添加脚本对话框打开事件处理程序。
//   - 当 JavaScript 代码调用 alert、confirm、prompt 或 beforeunload 时触发此事件。
//   - 仅当 ICoreWebView2Settings.AreDefaultScriptDialogsEnabled 属性设置为 FALSE 时，此事件才会触发。
//   - ScriptDialogOpening 事件会抑制对话框，或使用自定义对话框替换默认对话框。
func (i *ICoreWebView2) AddScriptDialogOpening(eventHandler *ICoreWebView2ScriptDialogOpeningEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddScriptDialogOpening.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveScriptDialogOpening 移除脚本对话框打开事件处理程序。
func (i *ICoreWebView2) RemoveScriptDialogOpening(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveScriptDialogOpening.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDevToolsProtocolEventReceiver 获取指定 DevTools 协议事件的接收器。
//
// eventName: DevTools 协议事件的完整名称，格式为 {domain}.{event}。
func (i *ICoreWebView2) GetDevToolsProtocolEventReceiver(eventName string) (*ICoreWebView2DevToolsProtocolEventReceiver, error) {
	_eventName, err := syscall.UTF16PtrFromString(eventName)
	if err != nil {
		return nil, err
	}
	var receiver *ICoreWebView2DevToolsProtocolEventReceiver
	r, _, _ := i.Vtbl.GetDevToolsProtocolEventReceiver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_eventName)),
		uintptr(unsafe.Pointer(&receiver)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return receiver, nil
}

// CallDevToolsProtocolMethod 调用 Chrome DevTools 协议方法。
//
// methodName: DevTools 协议方法的完整名称，格式为 {domain}.{method}。
//
// parametersAsJson: JSON 字符串，其中包含相应方法的参数。
//
// handler: 执行完成后的回调处理程序，接收返回的 JSON 结果。
func (i *ICoreWebView2) CallDevToolsProtocolMethod(methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) error {
	_methodName, err := syscall.UTF16PtrFromString(methodName)
	if err != nil {
		return err
	}
	_parametersAsJson, err := syscall.UTF16PtrFromString(parametersAsJson)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.CallDevToolsProtocolMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_methodName)),
		uintptr(unsafe.Pointer(_parametersAsJson)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CallDevToolsProtocolMethodEx 调用 Chrome DevTools 协议方法。
//
// impl: *WebViewEventImpl。
//
// methodName: DevTools 协议方法的完整名称，格式为 {domain}.{method}。
//
// parametersAsJson: JSON 字符串，其中包含相应方法的参数。
//
// cb: 执行完成后的回调处理程序，接收返回的 JSON 结果。
func (i *ICoreWebView2) CallDevToolsProtocolMethodEx(impl *WebViewEventImpl, methodName string, parametersAsJson string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "CallDevToolsProtocolMethodCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "CallDevToolsProtocolMethodCompleted")
	return i.CallDevToolsProtocolMethod(methodName, parametersAsJson, (*ICoreWebView2CallDevToolsProtocolMethodCompletedHandler)(handler))
}

// MustGetSettings 获取 ICoreWebView2Settings 对象, 它包含正在运行的 WebView 的各种可修改设置。出错时会触发全局错误回调.
func (i *ICoreWebView2) MustGetSettings() *ICoreWebView2Settings {
	s, err := i.GetSettings()
	ReportErrorAuto(err)
	return s
}

// MustGetSource 获取当前顶级文档的URI。如果导航正在进行中，则返回即将导航到的URI。出错时会触发全局错误回调.
func (i *ICoreWebView2) MustGetSource() string {
	uri, err := i.GetSource()
	ReportErrorAuto(err)
	return uri
}

// MustGetContainsFullScreenElement 获取 WebView 是否包含全屏元素。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetContainsFullScreenElement() bool {
	result, err := i.GetContainsFullScreenElement()
	ReportErrorAuto(err)
	return result
}

// MustGetDocumentTitle 获取当前顶级文档的标题。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetDocumentTitle() string {
	title, err := i.GetDocumentTitle()
	ReportErrorAuto(err)
	return title
}

// MustGetBrowserProcessID 获取承载 WebView 的浏览器进程的 ID。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetBrowserProcessID() uint32 {
	pid, err := i.GetBrowserProcessID()
	ReportErrorAuto(err)
	return pid
}

// MustGetCanGoBack 获取 WebView 是否可以导航到上一页。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetCanGoBack() bool {
	result, err := i.GetCanGoBack()
	ReportErrorAuto(err)
	return result
}

// MustGetCanGoForward 获取 WebView 是否可以导航到下一页。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetCanGoForward() bool {
	result, err := i.GetCanGoForward()
	ReportErrorAuto(err)
	return result
}

// MustGetDevToolsProtocolEventReceiver 获取指定 DevTools 协议事件的接收器。出错时会触发全局错误回调。
//
// eventName: DevTools 协议事件的完整名称，格式为 {domain}.{event}。
func (i *ICoreWebView2) MustGetDevToolsProtocolEventReceiver(eventName string) *ICoreWebView2DevToolsProtocolEventReceiver {
	receiver, err := i.GetDevToolsProtocolEventReceiver(eventName)
	ReportErrorAuto(err)
	return receiver
}
