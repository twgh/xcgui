package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"golang.org/x/sys/windows"
	"sync"
	"syscall"
)

type Webview2 struct {
	CoreWebView *ICoreWebView2           // CoreWebView2
	Controller  *ICoreWebView2Controller // WebView2 控制器
	Edge        *Edge

	// ------------------- WebView2_ 系列对象  -------------------

	WebView2_2  *ICoreWebView2_2  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_3  *ICoreWebView2_3  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_4  *ICoreWebView2_4  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_5  *ICoreWebView2_5  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_6  *ICoreWebView2_6  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_7  *ICoreWebView2_7  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_8  *ICoreWebView2_8  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_9  *ICoreWebView2_9  // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_10 *ICoreWebView2_10 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_11 *ICoreWebView2_11 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_12 *ICoreWebView2_12 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_13 *ICoreWebView2_13 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_14 *ICoreWebView2_14 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_15 *ICoreWebView2_15 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_16 *ICoreWebView2_16 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_17 *ICoreWebView2_17 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_18 *ICoreWebView2_18 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_19 *ICoreWebView2_19 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_20 *ICoreWebView2_20 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_21 *ICoreWebView2_21 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_22 *ICoreWebView2_22 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_23 *ICoreWebView2_23 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_24 *ICoreWebView2_24 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_25 *ICoreWebView2_25 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_26 *ICoreWebView2_26 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_
	WebView2_27 *ICoreWebView2_27 // 默认值为 nil, 使用前需自行赋值, 或调用一次 InitAllWebView2_

	// 注册事件时使用的Token
	EventRegistrationToken *EventRegistrationToken

	// -------------------- Handlers --------------------

	// HandlerTrySuspendCompleted 在调用 ICoreWebView2_3.TrySuspend 时使用. 如果为 nil, 需自行调用 NewICoreWebView2TrySuspendCompletedHandler 来赋值.
	HandlerTrySuspendCompleted *ICoreWebView2TrySuspendCompletedHandler
	// HandlerExecuteScriptCompleted 在调用 ICoreWebView2.ExecuteScript 时使用. 如果为 nil, 需自行调用 NewICoreWebView2ExecuteScriptCompletedHandler 来赋值.
	HandlerExecuteScriptCompleted *ICoreWebView2ExecuteScriptCompletedHandler
	// HandlerAddScriptToExecuteOnDocumentCreatedCompleted 在调用 ICoreWebView2.AddScriptToExecuteOnDocumentCreated 时使用. 如果为 nil, 需自行调用 NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler 来赋值.
	HandlerAddScriptToExecuteOnDocumentCreatedCompleted *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler
	// HandlerGetCookiesCompleted 在调用 ICoreWebView2CookieManager.GetCookies 时使用. 如果为 nil, 需自行调用 NewICoreWebView2GetCookiesCompletedHandler 来赋值.
	HandlerGetCookiesCompleted *ICoreWebView2GetCookiesCompletedHandler

	// -------------------- Callbacks --------------------

	// 挂起事件结果回调
	cbTrySuspendCompleted func(errorCode syscall.Errno, isSuccessful bool) uintptr
	// 执行脚本完成事件回调
	cbExecuteScriptCompleted func(errorCode syscall.Errno, result string) uintptr
	// 在文档创建前添加脚本完成回调
	cbAddScriptToExecuteOnDocumentCreatedCompleted func(errorCode syscall.Errno, id string) uintptr
	// 获取 Cookies 完成回调
	cbGetCookiesCompleted func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr
	// 仅供内部使用的网页消息事件回调
	msgcb_xcgui func(string)

	// -------------------- 事件 Handlers 和 Callbacks --------------------

	// 网页消息接收事件处理程序. [内部需要使用, 不能移除]
	handlerWebMessageReceivedEvent *ICoreWebView2WebMessageReceivedEventHandler
	cbMessageReceivedEvent         func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
	// 权限请求事件处理程序. [内部需要使用, 不能移除]
	handlerPermissionRequestedEvent *ICoreWebView2PermissionRequestedEventHandler
	cbPermissionRequestedEvent      func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr
	// HandlerWebResourceRequestedEvent 网络资源请求事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerWebResourceRequestedEvent *ICoreWebView2WebResourceRequestedEventHandler
	cbWebResourceRequestedEvent      func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr
	// HandlerAcceleratorKeyPressedEvent 快捷键事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerAcceleratorKeyPressedEvent *ICoreWebView2AcceleratorKeyPressedEventHandler
	cbAcceleratorKeyPressedEvent      func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr
	// HandlerNavigationCompletedEvent 导航完成事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerNavigationCompletedEvent *ICoreWebView2NavigationCompletedEventHandler
	cbNavigationCompletedEvent      func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
	// HandlerNavigationStartingEvent 导航开始事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerNavigationStartingEvent *ICoreWebView2NavigationStartingEventHandler
	cbNavigationStartingEvent      func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
	// HandlerContentLoadingEvent 内容加载事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerContentLoadingEvent *ICoreWebView2ContentLoadingEventHandler
	cbContentLoadingEvent      func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr
	// HandlerNewWindowRequestedEvent 新窗口请求事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerNewWindowRequestedEvent *ICoreWebView2NewWindowRequestedEventHandler
	cbNewWindowRequestedEvent      func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr
	// HandlerSourceChangedEvent 源改变事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerSourceChangedEvent *ICoreWebView2SourceChangedEventHandler
	cbSourceChangedEvent      func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr
	// HandlerFrameNavigationStartingEvent 框架导航开始事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerFrameNavigationStartingEvent *ICoreWebView2NavigationStartingEventHandler2
	cbFrameNavigationStartingEvent      func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
	// HandlerFrameNavigationCompletedEvent 框架导航完成事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerFrameNavigationCompletedEvent *ICoreWebView2NavigationCompletedEventHandler2
	cbFrameNavigationCompletedEvent      func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
	// HandlerWindowCloseRequestedEvent 窗口关闭请求事件处理程序. 在调用 Event_ 时自动赋值.
	HandlerWindowCloseRequestedEvent *ICoreWebView2WindowCloseRequestedEventHandler
	cbWindowCloseRequestedEvent      func(sender *ICoreWebView2, args *IUnknown) uintptr
	// HandlerDocumentTitleChangedEvent 文档标题改变事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerDocumentTitleChangedEvent *ICoreWebView2DocumentTitleChangedEventHandler
	cbDocumentTitleChangedEvent      func(sender *ICoreWebView2, args *IUnknown) uintptr

	// 添加窗口光栅化缩放比例改变事件处理程序时返回的 Token
	TokenRasterizationScaleChangedEvent EventRegistrationToken
	// HandlerRasterizationScaleChanged 窗口光栅化缩放比例改变事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerRasterizationScaleChangedEvent *ICoreWebView2RasterizationScaleChangedEventHandler
	cbRasterizationScaleChangedEvent      func(sender *ICoreWebView2Controller, args *IUnknown) uintptr

	// 宿主窗口句柄
	hwnd uintptr
	// 权限map读写锁
	rwxPermissions sync.RWMutex
	// 权限map
	permissions map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE
	// WebView2 控制器创建完成后是否自动获取焦点
	focusOnInit bool
}

func (e *Webview2) init() {
	e.handlerWebMessageReceivedEvent = NewICoreWebView2WebMessageReceivedEventHandler(e)
	e.handlerPermissionRequestedEvent = NewICoreWebView2PermissionRequestedEventHandler(e)

	e.permissions = make(map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE)
}

// InitAllWebView2_ 初始化所有的 WebView2_ 系列对象, 给本类成员赋值.
//   - 从 2 到 27 按顺序获取, 出错时返回出错序号.
//   - 比如获取 WebView2_2 时失败返回 2, 获取 WebView2_3 时失败返回 3, 以此类推.
//   - 如果返回0, 代表成功获取到了 WebView2_2 到 WebView2_27.
//   - 这些对象, 只要运行时版本支持, 基本是不会获取出错的.
func (e *Webview2) InitAllWebView2_() int {
	var err error
	if e.WebView2_2 == nil {
		e.WebView2_2, err = e.CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return 2
		}
	}
	if e.WebView2_3 == nil {
		e.WebView2_3, err = e.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return 3
		}
	}
	if e.WebView2_4 == nil {
		e.WebView2_4, err = e.CoreWebView.GetICoreWebView2_4()
		if err != nil {
			return 4
		}
	}
	if e.WebView2_5 == nil {
		e.WebView2_5, err = e.CoreWebView.GetICoreWebView2_5()
		if err != nil {
			return 5
		}
	}
	if e.WebView2_6 == nil {
		e.WebView2_6, err = e.CoreWebView.GetICoreWebView2_6()
		if err != nil {
			return 6
		}
	}
	if e.WebView2_7 == nil {
		e.WebView2_7, err = e.CoreWebView.GetICoreWebView2_7()
		if err != nil {
			return 7
		}
	}
	if e.WebView2_8 == nil {
		e.WebView2_8, err = e.CoreWebView.GetICoreWebView2_8()
		if err != nil {
			return 8
		}
	}
	if e.WebView2_9 == nil {
		e.WebView2_9, err = e.CoreWebView.GetICoreWebView2_9()
		if err != nil {
			return 9
		}
	}
	if e.WebView2_10 == nil {
		e.WebView2_10, err = e.CoreWebView.GetICoreWebView2_10()
		if err != nil {
			return 10
		}
	}
	if e.WebView2_11 == nil {
		e.WebView2_11, err = e.CoreWebView.GetICoreWebView2_11()
		if err != nil {
			return 11
		}
	}
	if e.WebView2_12 == nil {
		e.WebView2_12, err = e.CoreWebView.GetICoreWebView2_12()
		if err != nil {
			return 12
		}
	}
	if e.WebView2_13 == nil {
		e.WebView2_13, err = e.CoreWebView.GetICoreWebView2_13()
		if err != nil {
			return 13
		}
	}
	if e.WebView2_14 == nil {
		e.WebView2_14, err = e.CoreWebView.GetICoreWebView2_14()
		if err != nil {
			return 14
		}
	}
	if e.WebView2_15 == nil {
		e.WebView2_15, err = e.CoreWebView.GetICoreWebView2_15()
		if err != nil {
			return 15
		}
	}
	if e.WebView2_16 == nil {
		e.WebView2_16, err = e.CoreWebView.GetICoreWebView2_16()
		if err != nil {
			return 16
		}
	}
	if e.WebView2_17 == nil {
		e.WebView2_17, err = e.CoreWebView.GetICoreWebView2_17()
		if err != nil {
			return 17
		}
	}
	if e.WebView2_18 == nil {
		e.WebView2_18, err = e.CoreWebView.GetICoreWebView2_18()
		if err != nil {
			return 18
		}
	}
	if e.WebView2_19 == nil {
		e.WebView2_19, err = e.CoreWebView.GetICoreWebView2_19()
		if err != nil {
			return 19
		}
	}
	if e.WebView2_20 == nil {
		e.WebView2_20, err = e.CoreWebView.GetICoreWebView2_20()
		if err != nil {
			return 20
		}
	}
	if e.WebView2_21 == nil {
		e.WebView2_21, err = e.CoreWebView.GetICoreWebView2_21()
		if err != nil {
			return 21
		}
	}
	if e.WebView2_22 == nil {
		e.WebView2_22, err = e.CoreWebView.GetICoreWebView2_22()
		if err != nil {
			return 22
		}
	}
	if e.WebView2_23 == nil {
		e.WebView2_23, err = e.CoreWebView.GetICoreWebView2_23()
		if err != nil {
			return 23
		}
	}
	if e.WebView2_24 == nil {
		e.WebView2_24, err = e.CoreWebView.GetICoreWebView2_24()
		if err != nil {
			return 24
		}
	}
	if e.WebView2_25 == nil {
		e.WebView2_25, err = e.CoreWebView.GetICoreWebView2_25()
		if err != nil {
			return 25
		}
	}
	if e.WebView2_26 == nil {
		e.WebView2_26, err = e.CoreWebView.GetICoreWebView2_26()
		if err != nil {
			return 26
		}
	}
	if e.WebView2_27 == nil {
		e.WebView2_27, err = e.CoreWebView.GetICoreWebView2_27()
		if err != nil {
			return 27
		}
	}
	return 0
}

// 释放所有的 WebView2_ 系列对象. 会在使用 Close 时自动调用.
func (e *Webview2) releaseAllWebView2_() {
	if e.WebView2_2 != nil {
		e.WebView2_2.Release()
		e.WebView2_2 = nil
	}
	if e.WebView2_3 != nil {
		e.WebView2_3.Release()
		e.WebView2_3 = nil
	}
	if e.WebView2_4 != nil {
		e.WebView2_4.Release()
		e.WebView2_4 = nil
	}
	if e.WebView2_5 != nil {
		e.WebView2_5.Release()
		e.WebView2_5 = nil
	}
	if e.WebView2_6 != nil {
		e.WebView2_6.Release()
		e.WebView2_6 = nil
	}
	if e.WebView2_7 != nil {
		e.WebView2_7.Release()
		e.WebView2_7 = nil
	}
	if e.WebView2_8 != nil {
		e.WebView2_8.Release()
		e.WebView2_8 = nil
	}
	if e.WebView2_9 != nil {
		e.WebView2_9.Release()
		e.WebView2_9 = nil
	}
	if e.WebView2_10 != nil {
		e.WebView2_10.Release()
		e.WebView2_10 = nil
	}
	if e.WebView2_11 != nil {
		e.WebView2_11.Release()
		e.WebView2_11 = nil
	}
	if e.WebView2_12 != nil {
		e.WebView2_12.Release()
		e.WebView2_12 = nil
	}
	if e.WebView2_13 != nil {
		e.WebView2_13.Release()
		e.WebView2_13 = nil
	}
	if e.WebView2_14 != nil {
		e.WebView2_14.Release()
		e.WebView2_14 = nil
	}
	if e.WebView2_15 != nil {
		e.WebView2_15.Release()
		e.WebView2_15 = nil
	}
	if e.WebView2_16 != nil {
		e.WebView2_16.Release()
		e.WebView2_16 = nil
	}
	if e.WebView2_17 != nil {
		e.WebView2_17.Release()
		e.WebView2_17 = nil
	}
	if e.WebView2_18 != nil {
		e.WebView2_18.Release()
		e.WebView2_18 = nil
	}
	if e.WebView2_19 != nil {
		e.WebView2_19.Release()
		e.WebView2_19 = nil
	}
	if e.WebView2_20 != nil {
		e.WebView2_20.Release()
		e.WebView2_20 = nil
	}
	if e.WebView2_21 != nil {
		e.WebView2_21.Release()
		e.WebView2_21 = nil
	}
	if e.WebView2_22 != nil {
		e.WebView2_22.Release()
		e.WebView2_22 = nil
	}
	if e.WebView2_23 != nil {
		e.WebView2_23.Release()
		e.WebView2_23 = nil
	}
	if e.WebView2_24 != nil {
		e.WebView2_24.Release()
		e.WebView2_24 = nil
	}
	if e.WebView2_25 != nil {
		e.WebView2_25.Release()
		e.WebView2_25 = nil
	}
	if e.WebView2_26 != nil {
		e.WebView2_26.Release()
		e.WebView2_26 = nil
	}
	if e.WebView2_27 != nil {
		e.WebView2_27.Release()
		e.WebView2_27 = nil
	}

	if e.CoreWebView != nil {
		e.CoreWebView.Release()
		e.CoreWebView = nil
	}
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

// Refresh 网页_刷新, 调用js代码刷新. 必须在 UI 线程执行.
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

func (e *Webview2) QueryInterface(refiid, object uintptr) uintptr {
	err := e.CoreWebView.QueryInterface(refiid, object)
	errno, _ := ErrorToErrno(err)
	return uintptr(errno)
}

func (e *Webview2) AddRef() uintptr {
	return 1
}

func (e *Webview2) Release() uintptr {
	return 1
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

	e.releaseAllWebView2_()
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

// --------------------------- 回调实现 ---------------------------

// TrySuspendCompleted 尝试挂起 webview 后调用, 以获取执行结果.
func (e *Webview2) TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	if e.cbTrySuspendCompleted != nil {
		e.cbTrySuspendCompleted(errorCode, isSuccessful)
	}
	return 0
}

// ExecuteScriptCompleted 执行 js 脚本完成后调用, 以获取执行结果.
func (e *Webview2) ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr {
	if e.cbExecuteScriptCompleted != nil {
		var str string
		if result != nil && *result != 0 {
			str = windows.UTF16PtrToString(result)
		}
		e.cbExecuteScriptCompleted(errorCode, str)
	}
	return 0
}

// AddScriptToExecuteOnDocumentCreatedCompleted 添加 js 脚本完成后调用, 以获取执行结果.
func (e *Webview2) AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr {
	if e.cbAddScriptToExecuteOnDocumentCreatedCompleted != nil {
		var idStr string
		if id != nil && *id != 0 {
			idStr = windows.UTF16PtrToString(id)
		}
		e.cbAddScriptToExecuteOnDocumentCreatedCompleted(errorCode, idStr)
	}
	return 0
}

// GetCookiesCompleted 获取 cookies 完成后调用, 以获取执行结果.
func (e *Webview2) GetCookiesCompleted(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr {
	if e.cbGetCookiesCompleted != nil {
		e.cbGetCookiesCompleted(errorCode, cookies)
	}
	return 0
}
