package edge

import (
	"embed"
	"errors"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"io/fs"
	"syscall"
)

type Webview2 struct {
	// WebView2_ 系列对象
	Webview2_Objs

	// 事件接口实现
	Webview2Impl

	CoreWebView *ICoreWebView2           // CoreWebView2
	Controller  *ICoreWebView2Controller // WebView2 控制器
	Edge        *Edge                    // WebView2 环境

	// 注册事件时使用的Token
	EventRegistrationToken *EventRegistrationToken

	// 宿主窗口句柄
	hwnd uintptr
	// WebView2 控制器创建完成后是否自动获取焦点
	focusOnInit bool
}

func (e *Webview2) init() {
	e.edge = e.Edge
	e.handlerWebMessageReceivedEvent = NewICoreWebView2WebMessageReceivedEventHandler(e)
	e.handlerPermissionRequestedEvent = NewICoreWebView2PermissionRequestedEventHandler(e)

	e.permissions = make(map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE)
}

// Navigate 导航 webview 到给定的 URL。
//   - URL 可能是数据 URI，即 "data:text/text,<html>...</html>"。
//   - 通常不进行适当的 url 编码也是可以的, webview 会为你重新编码。
func (e *Webview2) Navigate(url string) error {
	return e.CoreWebView.Navigate(url)
}

// NavigateToString 直接设置 webview 的 HTML。
// 页面的来源是 `about:blank`。
//
// htmlContent: 参数的总大小不能大于2 MB（2*1024*1024字节）.
func (e *Webview2) NavigateToString(htmlContent string) error {
	return e.CoreWebView.NavigateToString(htmlContent)
}

// Eval 执行 JS 代码(异步). 必须在 UI 线程执行.
func (e *Webview2) Eval(script string) error {
	return e.CoreWebView.ExecuteScript(script, nil)
}

// Refresh 网页_刷新, 调用js代码刷新(location.reload). 必须在 UI 线程执行.
//
// forceReload: 是否强制刷新, 默认为false. 为 true 时，浏览器会强制重新加载页面，忽略缓存。这意味着无论页面是否已经在本地缓存中，都会从服务器重新获取资源。
func (e *Webview2) Refresh(forceReload ...bool) error {
	b := ""
	if len(forceReload) > 0 && forceReload[0] {
		b = "true"
	}
	return e.Eval("location.reload(" + b + ");")
}

// Show 显示 webview。
func (e *Webview2) Show() error {
	return e.Controller.PutIsVisible(true)
}

// Hide 隐藏 webview。
func (e *Webview2) Hide() error {
	return e.Controller.PutIsVisible(false)
}

// Resize 重新设置 webview 窗口大小和宿主窗口的客户区大小一致.
func (e *Webview2) Resize() error {
	if e.Controller == nil {
		return errors.New("ICoreWebView2Controller is nil")
	}
	var bounds xc.RECT
	wapi.GetClientRect(e.hwnd, &bounds)
	return e.Controller.PutBounds(bounds)
}

// SetPermission 设置权限。设置后如果网页请求该权限, 会根据设置的 state 来允许或拒绝请求。
func (e *Webview2) SetPermission(kind COREWEBVIEW2_PERMISSION_KIND, state COREWEBVIEW2_PERMISSION_STATE) {
	e.rwxPermissions.Lock()
	e.permissions[kind] = state
	e.rwxPermissions.Unlock()
}

// AddWebResourceRequestedFilter 添加web资源请求过滤器。[已过时]
//   - 请改用 ICoreWebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds，该方法将针对文档中的所有 iframe 触发此事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter
func (e *Webview2) AddWebResourceRequestedFilter(filter string, ctx COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	return e.CoreWebView.AddWebResourceRequestedFilter(filter, ctx)
}

// GetSettings 获取设置。
func (e *Webview2) GetSettings() (*ICoreWebView2Settings, error) {
	return e.CoreWebView.GetSettings()
}

// Close 关闭 WebView 并清理底层浏览器实例, 且关闭原生窗口。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (e *Webview2) Close() error {
	var err error
	if e.Controller != nil {
		err = e.Controller.Close()
		e.Controller.Release()
		e.Controller = nil
	}

	e.releaseAllWebView2_Objs()
	wapi.PostMessageW(e.hwnd, wapi.WM_CLOSE, 0, 0)
	return err
}

// NotifyParentWindowPositionChanged 通知父窗口位置已更改。
func (e *Webview2) NotifyParentWindowPositionChanged() error {
	if e.Controller == nil {
		return errors.New("controller is nil")
	}
	return e.Controller.NotifyParentWindowPositionChanged()
}

// Focus 设置焦点。
func (e *Webview2) Focus() error {
	if e.Controller == nil {
		e.focusOnInit = true
		return errors.New("controller is nil")
	}
	return e.Controller.MoveFocus(COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC)
}

// ExecuteScript 在主页面上下文中执行 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取 js 代码返回值.
func (e *Webview2) ExecuteScript(javaScript string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	if e.HandlerExecuteScriptCompleted == nil {
		e.HandlerExecuteScriptCompleted = NewICoreWebView2ExecuteScriptCompletedHandler(e)
	}
	e.cbExecuteScriptCompleted = cb
	return e.CoreWebView.ExecuteScript(javaScript, e.HandlerExecuteScriptCompleted)
}

// AddScriptToExecuteOnDocumentCreated 在主页面上下文中添加 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取添加脚本结果和 id.
func (e *Webview2) AddScriptToExecuteOnDocumentCreated(javaScript string, cb func(errorCode syscall.Errno, id string) uintptr) error {
	if e.HandlerAddScriptToExecuteOnDocumentCreatedCompleted == nil {
		e.HandlerAddScriptToExecuteOnDocumentCreatedCompleted = NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(e)
	}
	e.cbAddScriptToExecuteOnDocumentCreatedCompleted = cb
	return e.CoreWebView.AddScriptToExecuteOnDocumentCreated(javaScript, e.HandlerAddScriptToExecuteOnDocumentCreatedCompleted)
}

// RemoveScriptToExecuteOnDocumentCreated 移除在创建文档时要执行的脚本。
//
// id: 要移除的脚本的 ID。id 是 AddScriptToExecuteOnDocumentCreated 返回的.
func (e *Webview2) RemoveScriptToExecuteOnDocumentCreated(id string) error {
	return e.CoreWebView.RemoveScriptToExecuteOnDocumentCreated(id)
}

// TrySuspend 尝试挂起 WebView 控件, 以节省内存。
//
// cb: 回调函数, 在回调函数里可获取执行结果.
func (e *Webview2) TrySuspend(cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) error {
	if e.WebView2_3 == nil {
		var err error
		e.WebView2_3, err = e.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return err
		}
	}

	if e.HandlerTrySuspendCompleted == nil {
		e.HandlerTrySuspendCompleted = NewICoreWebView2TrySuspendCompletedHandler(e)
	}
	e.cbTrySuspendCompleted = cb
	return e.WebView2_3.TrySuspend(e.HandlerTrySuspendCompleted)
}

// Resume 尝试挂起 WebView 控件, 以节省内存。
func (e *Webview2) Resume() error {
	if e.WebView2_3 == nil {
		var err error
		e.WebView2_3, err = e.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return err
		}
	}
	return e.WebView2_3.Resume()
}

// IsSuspended 获取 WebView 控件是否已挂起。
func (e *Webview2) IsSuspended() bool {
	if e.WebView2_3 == nil {
		var err error
		e.WebView2_3, err = e.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return false
		}
	}
	return e.WebView2_3.MustGetIsSuspended()
}

// GetCookies 获取与指定 URI 匹配的所有 Cookie，失败返回 nil。
//   - 如果 uri 为空字符串，则返回同一配置文件下的所有 Cookie。
//   - 你可以通过调用 ICoreWebView2CookieManager.AddOrUpdateCookie 来修改 Cookie 对象，所做的更改将应用到WebView中。
//
// uri: 要匹配的 URI.
func (e *Webview2) GetCookies(uri string, cb func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr) error {
	if e.WebView2_2 == nil {
		var err error
		e.WebView2_2, err = e.CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return err
		}
	}
	CookieManager, err := e.WebView2_2.GetCookieManager()
	if err != nil {
		return err
	}
	if e.HandlerGetCookiesCompleted == nil {
		e.HandlerGetCookiesCompleted = NewICoreWebView2GetCookiesCompletedHandler(e)
	}
	e.cbGetCookiesCompleted = cb
	return CookieManager.GetCookies(uri, e.HandlerGetCookiesCompleted)
}

// 设置虚拟主机名和嵌入文件系统之间的映射，以便通过该主机名对网站可用.
//
// hostName: 要映射的主机名。
//
// embedFS: 要映射的文件系统。
//
// dir: 上面嵌入的目录名。例如: "assets". 不填则自动获取.
func (e *Webview2) SetVirtualHostNameToEmbedFSMapping(hostName string, embedFS embed.FS, dir ...string) error {
	var dirName string

	if len(dir) > 0 { // 使用指定的目录名
		dirName = dir[0]
	} else {
		// 自动检测目录名
		entries, err := embedFS.ReadDir(".")
		if err != nil {
			return errors.New("自动获取目录名时报错: " + err.Error())
		}
		if len(entries) != 1 || !entries[0].IsDir() {
			return errors.New("自动获取目录名失败")
		}
		dirName = entries[0].Name()
	}

	subFS, err := fs.Sub(embedFS, dirName)
	if err != nil {
		return errors.New("创建子文件系统失败: " + err.Error())
	}

	filter := "*://" + hostName + "/*"
	if e.WebView2_22 == nil {
		e.WebView2_22, err = e.CoreWebView.GetICoreWebView2_22()
		if err == nil {
			err = e.WebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL, COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL)
		} else {
			err = e.AddWebResourceRequestedFilter(filter, COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
		}
	}
	if err != nil {
		return err
	}

	if e.HandlerWebResourceRequestedEvent == nil {
		e.HandlerWebResourceRequestedEvent = NewICoreWebView2WebResourceRequestedEventHandler(e)
		err := e.CoreWebView.AddWebResourceRequested(e.HandlerWebResourceRequestedEvent, e.EventRegistrationToken)
		if err != nil {
			e.HandlerWebResourceRequestedEvent = nil
			return err
		}
	}

	e.firstResponse, err = e.Edge.Environment.CreateWebResourceResponse(nil, 200, "OK", "")
	ReportError2(err)
	e.hostName = "http://" + hostName + "/"
	e.embedFS = subFS
	return nil
}
