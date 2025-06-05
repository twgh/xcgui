package edge

import (
	"embed"
	"errors"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"io/fs"
	"syscall"
)

type WebView2 struct {
	// 事件接口实现
	WebViewEventImpl
	// WebView2_ 系列对象
	Webview2_Objs

	// 宿主窗口句柄
	hwnd uintptr
	// WebView2 控制器创建完成后是否自动获取焦点
	focusOnInit bool
}

func (w *WebView2) init() {
	w.handlerWebMessageReceivedEvent = NewICoreWebView2WebMessageReceivedEventHandler(w)
	w.handlerPermissionRequestedEvent = NewICoreWebView2PermissionRequestedEventHandler(w)

	w.permissions = make(map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE)
}

// Navigate 导航 webview 到给定的 URL。
//   - URL 可能是数据 URI，即 "data:text/text,<html>...</html>"。
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

// Refresh 网页_刷新, 调用js代码刷新(location.reload). 必须在 UI 线程执行.
//
// forceReload: 是否强制刷新, 默认为false. 为 true 时，浏览器会强制重新加载页面，忽略缓存。这意味着无论页面是否已经在本地缓存中，都会从服务器重新获取资源。
func (w *WebView2) Refresh(forceReload ...bool) error {
	b := ""
	if len(forceReload) > 0 && forceReload[0] {
		b = "true"
	}
	return w.Eval("location.reload(" + b + ");")
}

// Show 显示 webview。
func (w *WebView2) Show() error {
	return w.Controller.PutIsVisible(true)
}

// Hide 隐藏 webview。
func (w *WebView2) Hide() error {
	return w.Controller.PutIsVisible(false)
}

// Resize 重新设置 webview 窗口大小和宿主窗口的客户区大小一致.
func (w *WebView2) Resize() error {
	if w.Controller == nil {
		return errors.New("ICoreWebView2Controller is nil")
	}
	var bounds xc.RECT
	wapi.GetClientRect(w.hwnd, &bounds)
	return w.Controller.PutBounds(bounds)
}

// SetPermission 设置权限。设置后如果网页请求该权限, 会根据设置的 state 来允许或拒绝请求。
func (w *WebView2) SetPermission(kind COREWEBVIEW2_PERMISSION_KIND, state COREWEBVIEW2_PERMISSION_STATE) {
	w.rwxPermissions.Lock()
	w.permissions[kind] = state
	w.rwxPermissions.Unlock()
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

// Close 关闭 WebView 并清理底层浏览器实例, 且销毁原生窗口。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView2) Close() error {
	var err error
	if w.Controller != nil {
		err = w.Controller.Close()
		w.Controller.Release()
		w.Controller = nil
	}

	w.ReleaseEventHandler()
	w.releaseAllWebView2_Objs()
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
	if w.HandlerExecuteScriptCompleted == nil {
		w.HandlerExecuteScriptCompleted = NewICoreWebView2ExecuteScriptCompletedHandler(w)
	}
	w.cbExecuteScriptCompleted = cb
	return w.CoreWebView.ExecuteScript(javaScript, w.HandlerExecuteScriptCompleted)
}

// AddScriptToExecuteOnDocumentCreated 在主页面上下文中添加 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取添加脚本结果和 id.
func (w *WebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, cb func(errorCode syscall.Errno, id string) uintptr) error {
	if w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted == nil {
		w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted = NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(w)
	}
	w.cbAddScriptToExecuteOnDocumentCreatedCompleted = cb
	return w.CoreWebView.AddScriptToExecuteOnDocumentCreated(javaScript, w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted)
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

	if w.HandlerTrySuspendCompleted == nil {
		w.HandlerTrySuspendCompleted = NewICoreWebView2TrySuspendCompletedHandler(w)
	}
	w.cbTrySuspendCompleted = cb
	return w.WebView2_3.TrySuspend(w.HandlerTrySuspendCompleted)
}

// Resume 尝试挂起 WebView 控件, 以节省内存。
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

// GetCookies 获取与指定 URI 匹配的所有 Cookie，失败返回 nil。
//   - 如果 uri 为空字符串，则返回同一配置文件下的所有 Cookie。
//   - 你可以通过调用 ICoreWebView2CookieManager.AddOrUpdateCookie 来修改 Cookie 对象，所做的更改将应用到WebView中。
//
// uri: 要匹配的 URI.
func (w *WebView2) GetCookies(uri string, cb func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr) error {
	if w.WebView2_2 == nil {
		var err error
		w.WebView2_2, err = w.CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return err
		}
	}
	CookieManager, err := w.WebView2_2.GetCookieManager()
	if err != nil {
		return err
	}
	if w.HandlerGetCookiesCompleted == nil {
		w.HandlerGetCookiesCompleted = NewICoreWebView2GetCookiesCompletedHandler(w)
	}
	w.cbGetCookiesCompleted = cb
	return CookieManager.GetCookies(uri, w.HandlerGetCookiesCompleted)
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

// 设置虚拟主机名和嵌入文件系统之间的映射，以便通过该主机名对网站可用.
//
// hostName: 要映射的主机名。
//
// embedFS: 要映射的文件系统。
//
// dir: 上面嵌入的目录名。例如: "assets". 不填则自动获取.
func (w *WebView2) SetVirtualHostNameToEmbedFSMapping(hostName string, embedFS embed.FS, dir ...string) error {
	var dirName string

	if len(dir) > 0 { // 使用指定的目录名
		dirName = dir[0]
	} else {
		// 自动检测目录名
		entries, err := embedFS.ReadDir(".")
		if err != nil {
			return errors.New("error reported when automatically obtaining directory name: " + err.Error())
		}
		if len(entries) != 1 || !entries[0].IsDir() {
			return errors.New("automatic retrieval of directory name failed")
		}
		dirName = entries[0].Name()
	}

	subFS, err := fs.Sub(embedFS, dirName)
	if err != nil {
		return errors.New("failed to create sub file system: " + err.Error())
	}

	filter := "*://" + hostName + "/*"
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

	if w.HandlerWebResourceRequestedEvent == nil {
		w.HandlerWebResourceRequestedEvent = NewICoreWebView2WebResourceRequestedEventHandler(w)
		err := w.CoreWebView.AddWebResourceRequested(w.HandlerWebResourceRequestedEvent, w.EventRegistrationToken)
		if err != nil {
			w.HandlerWebResourceRequestedEvent = nil
			return err
		}
	}

	if !once { // 只创建一次
		w.firstResponse, err = w.Edge.Environment.CreateWebResourceResponse(nil, 200, "OK", "")
		ReportErrorAtuo(err)
	}
	w.hostName = "http://" + hostName + "/"
	w.embedFS = subFS
	return nil
}

// CapturePreview 捕获 Webview 的预览图像。
//
// imageFormat: 图像格式枚举值.
func (w *WebView2) CapturePreview(imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, stream *IStream, cb func(errorCode syscall.Errno) uintptr) error {
	if w.HandlerCapturePreviewCompleted == nil {
		w.HandlerCapturePreviewCompleted = NewICoreWebView2CapturePreviewCompletedHandler(w)
	}

	w.cbCapturePreviewCompleted = cb
	err := w.CoreWebView.CapturePreview(imageFormat, stream, w.HandlerCapturePreviewCompleted)
	if err != nil {
		return err
	}
	return nil
}
