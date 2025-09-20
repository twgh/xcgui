package edge

import (
	"errors"
	"syscall"

	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

type WebView2 struct {
	// 事件接口实现
	WebViewEventImpl
	// WebView2_ 系列对象
	WebView2_Objs

	// 宿主窗口句柄
	hwnd uintptr
	// WebView2 控制器创建完成后是否自动获取焦点
	focusOnInit bool
}

func (w *WebView2) init() {
	w.permissions = make(map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE)
}

func (w *WebView2) GetWebViewEventImpl() *WebViewEventImpl {
	return &w.WebViewEventImpl
}

// Navigate 导航 webview 到给定的 URL。
//   - URL 可能是数据 URI，即 "data:text/html,<html>...</html>"。
//   - 通常不进行适当的 url 编码也是可以的, webview 会为你重新编码。
func (w *WebView2) Navigate(url string) error {
	return w.CoreWebView.Navigate(url)
}

// NavigateToString 直接设置 webview 的 HTML。
// 页面的来源是 `about:blank`。
//
// htmlContent: 参数的总大小不能大于2 MB（2*1024*1024字节）.
func (w *WebView2) NavigateToString(htmlContent string) error {
	return w.CoreWebView.NavigateToString(htmlContent)
}

// Eval 执行 JS 代码(异步). 必须在 UI 线程执行.
func (w *WebView2) Eval(script string) error {
	if w.CoreWebView == nil {
		return errors.New("CoreWebView is nil")
	}
	return w.CoreWebView.ExecuteScript(script, nil)
}

// PostWebMessageAsString 将指定的 webMessage 发布到此 WebView 中的顶级文档。
//   - 其行为方式与 PostWebMessageAsJson 完全相同.
//   - 但 window.chrome.webview 消息的事件参数的 data 属性是一个字符串，其值与 webMessageAsString 相同。
//   - 如果希望使用简单字符串而非 JSON 对象进行通信，请使用此方法代替 PostWebMessageAsJson。
//
// webMessageAsString: JavaScript 对象的简单字符串，而不是 JSON 字符串表示。
func (w *WebView2) PostWebMessageAsString(webMessageAsString string) error {
	return w.CoreWebView.PostWebMessageAsString(webMessageAsString)
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
func (w *WebView2) PostWebMessageAsJSON(webMessageAsJSON string) error {
	return w.CoreWebView.PostWebMessageAsJSON(webMessageAsJSON)
}

// GetSource 获取当前顶级文档的URI。如果导航正在进行中，则返回即将导航到的URI。
func (w *WebView2) GetSource() string {
	return w.CoreWebView.MustGetSource()
}

// Refresh 网页_刷新, 调用 js 代码刷新(location.reload). 必须在 UI 线程执行.
//
// forceReload: 是否强制刷新, 默认为false. 为 true 时，浏览器会强制重新加载页面，忽略缓存。这意味着无论页面是否已经在本地缓存中，都会从服务器重新获取资源。
func (w *WebView2) Refresh(forceReload ...bool) error {
	b := ""
	if len(forceReload) > 0 && forceReload[0] {
		b = "true"
	}
	return w.Eval("location.reload(" + b + ");")
}

// Reload 重新加载当前页面。
func (w *WebView2) Reload() error {
	return w.CoreWebView.Reload()
}

// GoBack 将 WebView 导航到导航历史记录中的上一页。
func (w *WebView2) GoBack() error {
	return w.CoreWebView.GoBack()
}

// GoForward 将 WebView 导航到导航历史记录中的下一页。
func (w *WebView2) GoForward() error {
	return w.CoreWebView.GoForward()
}

// Stop 停止所有导航和挂起的资源获取。不停止脚本。
func (w *WebView2) Stop() error {
	return w.CoreWebView.Stop()
}

// GetDocumentTitle 获取当前顶级文档的标题。
func (w *WebView2) GetDocumentTitle() string {
	return w.CoreWebView.MustGetDocumentTitle()
}

// Show 显示或隐藏 webview。
func (w *WebView2) Show(isShow bool) error {
	return w.Controller.SetIsVisible(isShow)
}

// Resize 重新设置 webview 窗口大小和宿主窗口的客户区大小一致.
func (w *WebView2) Resize() error {
	if w.Controller == nil {
		return errors.New("ICoreWebView2Controller is nil")
	}
	var bounds xc.RECT
	wapi.GetClientRect(w.hwnd, &bounds)
	return w.Controller.SetBounds(bounds)
}

// AddWebResourceRequestedFilter 添加web资源请求过滤器。[已过时]
//   - 请改用 ICoreWebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds，该方法将针对文档中的所有 iframe 触发此事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter
func (w *WebView2) AddWebResourceRequestedFilter(filter string, ctx COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	return w.CoreWebView.AddWebResourceRequestedFilter(filter, ctx)
}

// GetSettings 获取设置。
func (w *WebView2) GetSettings() (*ICoreWebView2Settings, error) {
	return w.CoreWebView.GetSettings()
}

// EnableBrowserAcceleratorKeys 设置是否启用浏览器快捷键。默认是启用的。
//   - 此设置对 AcceleratorKeyPressed 事件没有影响。
func (w *WebView2) EnableBrowserAcceleratorKeys(enable bool) error {
	settings, err := w.GetSettings()
	if err != nil {
		return err
	}
	defer settings.Release()

	s3, err := settings.GetICoreWebView2Settings3()
	if err != nil {
		return err
	}
	defer s3.Release()

	return s3.SetAreBrowserAcceleratorKeysEnabled(enable)
}

// Close 关闭 WebView 并清理底层浏览器实例, 且销毁原生窗口。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView2) Close() error {
	var err error
	if w.Controller != nil {
		err = w.Controller.Close()
		w.Controller.Release()
		w.Controller = nil
	}

	WvEventHandler.ReleaseAllEventHandler(&w.WebViewEventImpl)
	w.ReleaseWebView2_Objs()

	if w.CoreWebView != nil {
		w.CoreWebView.Release()
		w.CoreWebView = nil
	}
	wapi.DestroyWindow(w.hwnd)
	return err
}

// NotifyParentWindowPositionChanged 通知父窗口位置已更改。
func (w *WebView2) NotifyParentWindowPositionChanged() error {
	if w.Controller == nil {
		return errors.New("controller is nil")
	}
	return w.Controller.NotifyParentWindowPositionChanged()
}

// Focus 设置焦点。
func (w *WebView2) Focus() error {
	if w.Controller == nil {
		w.focusOnInit = true
		return errors.New("controller is nil")
	}
	return w.Controller.MoveFocus(COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC)
}

// ExecuteScript 在主页面上下文中执行 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取 js 代码返回值.
func (w *WebView2) ExecuteScript(javaScript string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	return w.CoreWebView.ExecuteScriptEx(&w.WebViewEventImpl, javaScript, cb)
}

// AddScriptToExecuteOnDocumentCreated 在主页面上下文中添加 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取添加脚本结果和 id.
func (w *WebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, cb func(errorCode syscall.Errno, id string) uintptr) error {
	return w.CoreWebView.AddScriptToExecuteOnDocumentCreatedEx(&w.WebViewEventImpl, javaScript, cb)
}

// RemoveScriptToExecuteOnDocumentCreated 移除在创建文档时要执行的脚本。
//
// id: 要移除的脚本的 ID。id 是 AddScriptToExecuteOnDocumentCreated 返回的.
func (w *WebView2) RemoveScriptToExecuteOnDocumentCreated(id string) error {
	return w.CoreWebView.RemoveScriptToExecuteOnDocumentCreated(id)
}

// TrySuspend 尝试挂起 WebView 控件, 以节省内存。
//
// cb: 回调函数, 在回调函数里可获取执行结果.
func (w *WebView2) TrySuspend(cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) error {
	if w.WebView2_3 == nil {
		var err error
		w.WebView2_3, err = w.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return err
		}
	}
	return w.WebView2_3.TrySuspendEx(&w.WebViewEventImpl, cb)
}

// Resume 恢复WebView，以便它恢复网页上的活动。
func (w *WebView2) Resume() error {
	if w.WebView2_3 == nil {
		var err error
		w.WebView2_3, err = w.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return err
		}
	}
	return w.WebView2_3.Resume()
}

// IsSuspended 获取 WebView 控件是否已挂起。
func (w *WebView2) IsSuspended() bool {
	if w.WebView2_3 == nil {
		var err error
		w.WebView2_3, err = w.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return false
		}
	}
	return w.WebView2_3.MustGetIsSuspended()
}

// GetGetCookieManager 获取 ICoreWebView2CookieManager 对象，用于管理 Cookie。
func (w *WebView2) GetGetCookieManager() (*ICoreWebView2CookieManager, error) {
	if w.WebView2_2 == nil {
		var err error
		w.WebView2_2, err = w.CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return nil, err
		}
	}
	return w.WebView2_2.GetCookieManager()
}

// GetCookies 获取与指定 URI 匹配的所有 Cookie，失败返回 nil。
//   - 如果 uri 为空字符串，则返回同一配置文件下的所有 Cookie。
//   - 你可以通过调用 ICoreWebView2CookieManager.AddOrUpdateCookie 来修改 Cookie 对象，所做的更改将应用到 WebView 中。
//
// uri: 要匹配的 URI.
//
// cb: 接收结果的回调函数.
func (w *WebView2) GetCookies(uri string, cb func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr) error {
	CookieManager, err := w.GetGetCookieManager()
	if err != nil {
		return err
	}
	return CookieManager.GetCookiesEx(&w.WebViewEventImpl, uri, cb)
}

// CapturePreview 捕获 Webview 的预览图像。
//
// imageFormat: 图像格式枚举值.
//
// imageStream: 接收图像数据的流对象.
//
// cb: 捕获完成后的回调处理程序. 写入流完毕后触发.
func (w *WebView2) CapturePreview(imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, imageStream *IStream, cb func(errorCode syscall.Errno) uintptr) error {
	return w.CoreWebView.CapturePreviewEx(&w.WebViewEventImpl, imageFormat, imageStream, cb)
}

// SetVirtualHostNameToFolderMapping 设置虚拟主机名和文件夹路径之间的映射，以便通过该主机名对网站可用.
//
// hostName: 要映射的主机名。
//
// folderPath: 要映射的文件夹路径。
//
// accessKind: 资源访问权限。
func (w *WebView2) SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND) error {
	if w.WebView2_3 == nil {
		var err error
		w.WebView2_3, err = w.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return err
		}
	}
	return w.WebView2_3.SetVirtualHostNameToFolderMapping(hostName, folderPath, accessKind)
}

// EnableVirtualHostNameToEmbedFSMapping 启用虚拟主机名和嵌入文件系统之间的映射。
//   - 启用后, 该 WebView 可使用全局的虚拟主机名和嵌入文件系统之间的映射来处理请求。
func (w *WebView2) EnableVirtualHostNameToEmbedFSMapping(enable bool) error {
	var err error
	filter := "http*://*"

	if enable {
		if w.WebView2_22 == nil {
			w.WebView2_22, err = w.CoreWebView.GetICoreWebView2_22()
			if err == nil {
				err = w.WebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL, COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL)
			} else {
				err = w.AddWebResourceRequestedFilter(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
			}
		}
		if err != nil {
			return err
		}

		// 添加网页资源请求事件, 添加空回调, 主要目的是注册一下事件
		_, err = w.Event_WebResourceRequested(nil, true)
		if err != nil {
			return err
		}
		if !once { // 只创建一次
			w.firstResponse, err = w.Edge.Environment.CreateWebResourceResponse(nil, 200, "OK", "")
			ReportErrorAuto(err)
		}
	} else {
		if w.WebView2_22 == nil {
			w.WebView2_22, err = w.CoreWebView.GetICoreWebView2_22()
			if err == nil {
				err = w.WebView2_22.RemoveWebResourceRequestedFilterWithRequestSourceKinds(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL, COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL)
			} else {
				err = w.CoreWebView.RemoveWebResourceRequestedFilter(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
			}
		}
		if err != nil {
			return err
		}
	}

	w.enableVirtualHostNameToEmbedFSMapping = enable
	return nil
}
