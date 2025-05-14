package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"golang.org/x/sys/windows"
	"sync"
	"syscall"
)

type webview2 struct {
	Controller  *ICoreWebView2Controller // WebView2 控制器
	CoreWebView *ICoreWebView2           // CoreWebView2
	Edge        *Edge
	hwnd        uintptr

	// 注册事件时使用的Token
	eventRegistrationToken *EventRegistrationToken

	// 权限map读写锁
	rwxPermissions sync.RWMutex
	// 权限map
	permissions map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE

	// -------------------- Handlers --------------------
	// todo Handlers 全部移动到 Edge里? 需要测试

	// 在调用 ICoreWebView2_3.TrySuspend 时使用
	trySuspendCompletedHandler *ICoreWebView2TrySuspendCompletedHandler
	// 执行脚本完成
	executeScriptCompletedHandler *ICoreWebView2ExecuteScriptCompletedHandler
	// 在文档创建前添加脚本完成
	addScriptToExecuteOnDocumentCreatedCompletedHandler *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler

	// 网页消息接收事件
	webMessageReceivedEventHandler *ICoreWebView2WebMessageReceivedEventHandler
	// 权限请求事件
	permissionRequestedEventHandler *ICoreWebView2PermissionRequestedEventHandler
	// 网络资源请求事件
	webResourceRequestedEventHandler *ICoreWebView2WebResourceRequestedEventHandler
	// 快捷键事件
	acceleratorKeyPressedEventHandler *ICoreWebView2AcceleratorKeyPressedEventHandler
	// 导航完成事件
	navigationCompletedEventHandler *ICoreWebView2NavigationCompletedEventHandler
	// 导航开始事件
	navigationStartingEventHandler *ICoreWebView2NavigationStartingEventHandler
	// 内容加载事件
	contentLoadingEventHandler *ICoreWebView2ContentLoadingEventHandler

	// -------------------- Callbacks --------------------

	// 仅供内部使用的网页消息事件回调
	msgcb_xcgui func(string)
	// 挂起事件结果回调
	trySuspendCompletedCallBack func(errorCode syscall.Errno, isSuccessful bool) uintptr
	// 执行脚本完成事件回调
	executeScriptCompletedCallback func(errorCode syscall.Errno, result string) uintptr
	// 在文档创建前添加脚本完成回调
	addScriptToExecuteOnDocumentCreatedCompletedCallback func(errorCode syscall.Errno, id string) uintptr

	// 网页消息事件回调
	messageReceivedCallback func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
	// 网络资源请求事件回调
	webResourceRequestedCallback func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr
	// 导航完成事件回调
	navigationCompletedCallback func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
	// 导航开始事件回调
	navigationStartingCallback func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
	// 快捷键事件回调
	acceleratorKeyPressedCallback func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr
	// 内容加载事件回调
	contentLoadingCallback func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr
	// 权限请求事件回调
	permissionRequestedCallback func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr

	focusOnInit bool // WebView2 控制器创建完成后是否自动获取焦点
}

func (e *webview2) init() {
	e.webMessageReceivedEventHandler = NewICoreWebView2WebMessageReceivedEventHandler(e)
	e.permissionRequestedEventHandler = NewICoreWebView2PermissionRequestedEventHandler(e)

	e.permissions = make(map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE)
}

// Navigate 导航 webview 到给定的 URL。
//   - URL 可能是数据 URI，即 "data:text/text,<html>...</html>"。
//   - 通常不进行适当的 url 编码也是可以的, webview 会为你重新编码。
func (e *webview2) Navigate(url string) error {
	return e.CoreWebView.Navigate(url)
}

// NavigateToString 直接设置 webview 的 HTML。
// 页面的来源是 `about:blank`。
//
// htmlContent: 参数的总大小不能大于2 MB（2*1024*1024字节）.
func (e *webview2) NavigateToString(htmlContent string) error {
	return e.CoreWebView.NavigateToString(htmlContent)
}

// Eval 执行 JS 代码(异步). 必须在 UI 线程执行.
func (e *webview2) Eval(script string) error {
	return e.CoreWebView.ExecuteScript(script, nil)
}

// Refresh 网页_刷新, 调用js代码刷新. 必须在 UI 线程执行.
//
// forceReload: 是否强制刷新, 默认为false. 为 true 时，浏览器会强制重新加载页面，忽略缓存。这意味着无论页面是否已经在本地缓存中，都会从服务器重新获取资源。
func (e *webview2) Refresh(forceReload ...bool) error {
	b := ""
	if len(forceReload) > 0 && forceReload[0] {
		b = "true"
	}
	return e.Eval("location.reload(" + b + ");")
}

// Show 显示 webview。
func (e *webview2) Show() error {
	return e.Controller.PutIsVisible(true)
}

// Hide 隐藏 webview。
func (e *webview2) Hide() error {
	return e.Controller.PutIsVisible(false)
}

// Resize 重新设置 webview 窗口大小和宿主窗口的客户区大小一致.
func (e *webview2) Resize() error {
	if e.Controller == nil {
		return errors.New("ICoreWebView2Controller is nil")
	}
	var bounds xc.RECT
	wapi.GetClientRect(e.hwnd, &bounds)
	return e.Controller.PutBounds(bounds)
}

func (e *webview2) QueryInterface(refiid, object uintptr) uintptr {
	err := e.CoreWebView.QueryInterface(refiid, object)
	errno, _ := ErrorToErrno(err)
	return uintptr(errno)
}

func (e *webview2) AddRef() uintptr {
	return 1
}

func (e *webview2) Release() uintptr {
	return 1
}

// SetPermission 设置权限。设置后如果网页请求该权限, 会根据设置的 state 来允许或拒绝请求。
func (e *webview2) SetPermission(kind COREWEBVIEW2_PERMISSION_KIND, state COREWEBVIEW2_PERMISSION_STATE) {
	e.rwxPermissions.Lock()
	e.permissions[kind] = state
	e.rwxPermissions.Unlock()
}

// AddWebResourceRequestedFilter 添加web资源请求过滤器。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter
func (e *webview2) AddWebResourceRequestedFilter(filter string, ctx COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	return e.CoreWebView.AddWebResourceRequestedFilter(filter, ctx)
}

// GetSettings 获取设置。
func (e *webview2) GetSettings() (*ICoreWebView2Settings, error) {
	return e.CoreWebView.GetSettings()
}

// Close 关闭 WebView 并清理底层浏览器实例, 且关闭原生窗口。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口里的.
func (e *webview2) Close() error {
	var err error
	if e.Controller != nil {
		err = e.Controller.Close()
		e.Controller.Release()
		e.Controller = nil
	}

	if e.CoreWebView != nil {
		e.CoreWebView.Release()
		e.CoreWebView = nil
	}
	wapi.PostMessageW(e.hwnd, wapi.WM_CLOSE, 0, 0)
	return err
}

// NotifyParentWindowPositionChanged 通知父窗口位置已更改。
func (e *webview2) NotifyParentWindowPositionChanged() error {
	if e.Controller == nil {
		return errors.New("controller is nil")
	}
	return e.Controller.NotifyParentWindowPositionChanged()
}

// Focus 设置焦点。
func (e *webview2) Focus() error {
	if e.Controller == nil {
		e.focusOnInit = true
		return errors.New("controller is nil")
	}
	return e.Controller.MoveFocus(COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC)
}

// GetEventRegistrationToken 获取事件注册令牌.
func (e *webview2) GetEventRegistrationToken() *EventRegistrationToken {
	return e.eventRegistrationToken
}

// GetTrySuspendCompletedHandler 在调用 ICoreWebView2_3.TrySuspend 时作为参数使用.
func (e *webview2) GetTrySuspendCompletedHandler() *ICoreWebView2TrySuspendCompletedHandler {
	return e.trySuspendCompletedHandler
}

// GetExecuteScriptCompletedHandler 在调用 ICoreWebView2.ExecuteScript 时作为参数使用.
func (e *webview2) GetExecuteScriptCompletedHandler() *ICoreWebView2ExecuteScriptCompletedHandler {
	return e.executeScriptCompletedHandler
}

// GetAddScriptToExecuteOnDocumentCreatedCompletedHandler 在调用 ICoreWebView2.AddScriptToExecuteOnDocumentCreated 时作为参数使用.
func (e *webview2) GetAddScriptToExecuteOnDocumentCreatedCompletedHandler() *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler {
	return e.addScriptToExecuteOnDocumentCreatedCompletedHandler
}

// setTrySuspendCompletedCallBack 设置挂起事件结果回调.
func (e *webview2) setTrySuspendCompletedCallBack(cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) {
	if e.trySuspendCompletedHandler == nil {
		e.trySuspendCompletedHandler = NewICoreWebView2TrySuspendCompletedHandler(e)
	}
	e.trySuspendCompletedCallBack = cb
}

// setExecuteScriptCompletedCallBack 设置执行脚本完成事件回调.
func (e *webview2) setExecuteScriptCompletedCallBack(cb func(errorCode syscall.Errno, result string) uintptr) {
	if e.executeScriptCompletedHandler == nil {
		e.executeScriptCompletedHandler = NewICoreWebView2ExecuteScriptCompletedHandler(e)
	}
	e.executeScriptCompletedCallback = cb
}

// setAddScriptToExecuteOnDocumentCreatedCompletedCallBack 设置添加脚本完成事件回调.
func (e *webview2) setAddScriptToExecuteOnDocumentCreatedCompletedCallBack(cb func(errorCode syscall.Errno, id string) uintptr) {
	if e.addScriptToExecuteOnDocumentCreatedCompletedHandler == nil {
		e.addScriptToExecuteOnDocumentCreatedCompletedHandler = NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(e)
	}
	e.addScriptToExecuteOnDocumentCreatedCompletedCallback = cb
}

// ExecuteScript 在主页面上下文中执行 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取 js 代码返回值.
func (e *webview2) ExecuteScript(javaScript string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	e.setExecuteScriptCompletedCallBack(cb)
	return e.CoreWebView.ExecuteScript(javaScript, e.executeScriptCompletedHandler)
}

// AddScriptToExecuteOnDocumentCreated 在主页面上下文中添加 JavaScript 代码(异步)。
//
// javaScript: 要执行的 JavaScript 代码。
//
// cb: 回调函数, 在回调函数里可获取添加脚本结果和 id.
func (e *webview2) AddScriptToExecuteOnDocumentCreated(javaScript string, cb func(errorCode syscall.Errno, id string) uintptr) error {
	e.setAddScriptToExecuteOnDocumentCreatedCompletedCallBack(cb)
	return e.CoreWebView.AddScriptToExecuteOnDocumentCreated(javaScript, e.addScriptToExecuteOnDocumentCreatedCompletedHandler)
}

// RemoveScriptToExecuteOnDocumentCreated 移除在创建文档时要执行的脚本。
//
// id: 要移除的脚本的 ID。id 是 AddScriptToExecuteOnDocumentCreated 返回的.
func (e *webview2) RemoveScriptToExecuteOnDocumentCreated(id string) error {
	return e.CoreWebView.RemoveScriptToExecuteOnDocumentCreated(id)
}

// TrySuspend 尝试挂起 WebView 控件, 以节省内存。
//
// cb: 回调函数, 在回调函数里可获取执行结果.
func (e *webview2) TrySuspend(cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) error {
	WebView2_3, err := e.CoreWebView.GetICoreWebView2_3()
	if err != nil {
		return err
	}
	if WebView2_3 == nil {
		return errors.New("GetICoreWebView2_3 = nil")
	}
	e.setTrySuspendCompletedCallBack(cb)
	return WebView2_3.TrySuspend(e.trySuspendCompletedHandler)
}

// Resume 尝试挂起 WebView 控件, 以节省内存。
func (e *webview2) Resume() error {
	WebView2_3, err := e.CoreWebView.GetICoreWebView2_3()
	if err != nil {
		return err
	}
	if WebView2_3 == nil {
		return errors.New("GetICoreWebView2_3 = nil")
	}
	return WebView2_3.Resume()
}

// IsSuspended 获取 WebView 控件是否已挂起。
func (e *webview2) IsSuspended() bool {
	WebView2_3, _ := e.CoreWebView.GetICoreWebView2_3()
	if WebView2_3 == nil {
		return false
	}
	return WebView2_3.MustGetIsSuspended()
}

// --------------------------- 回调实现 ---------------------------

// TrySuspendCompleted 尝试挂起 webview 后调用, 以获取执行结果.
func (e *webview2) TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	if e.trySuspendCompletedCallBack != nil {
		e.trySuspendCompletedCallBack(errorCode, isSuccessful)
	}
	return 0
}

// ExecuteScriptCompleted 执行 js 脚本完成后调用, 以获取执行结果.
func (e *webview2) ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr {
	if e.executeScriptCompletedCallback != nil {
		var str string
		if result != nil && *result != 0 {
			str = windows.UTF16PtrToString(result)
		}
		e.executeScriptCompletedCallback(errorCode, str)
	}
	return 0
}

// AddScriptToExecuteOnDocumentCreatedCompleted 添加 js 脚本完成后调用, 以获取执行结果.
func (e *webview2) AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr {
	if e.addScriptToExecuteOnDocumentCreatedCompletedCallback != nil {
		var idStr string
		if id != nil && *id != 0 {
			idStr = windows.UTF16PtrToString(id)
		}
		e.addScriptToExecuteOnDocumentCreatedCompletedCallback(errorCode, idStr)
	}
	return 0
}
