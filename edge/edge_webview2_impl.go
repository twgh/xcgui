package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"io/fs"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
)

// Webview2Impl 事件接口实现.
type Webview2Impl struct {
	// -------------------- Handlers --------------------

	// HandlerTrySuspendCompleted 在调用 ICoreWebView2_3.TrySuspend 时使用. 如果为 nil, 需自行调用 NewICoreWebView2TrySuspendCompletedHandler 来赋值.
	HandlerTrySuspendCompleted *ICoreWebView2TrySuspendCompletedHandler
	// HandlerExecuteScriptCompleted 在调用 ICoreWebView2.ExecuteScript 时使用. 如果为 nil, 需自行调用 NewICoreWebView2ExecuteScriptCompletedHandler 来赋值.
	HandlerExecuteScriptCompleted *ICoreWebView2ExecuteScriptCompletedHandler
	// HandlerAddScriptToExecuteOnDocumentCreatedCompleted 在调用 ICoreWebView2.AddScriptToExecuteOnDocumentCreated 时使用. 如果为 nil, 需自行调用 NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler 来赋值.
	HandlerAddScriptToExecuteOnDocumentCreatedCompleted *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler
	// HandlerGetCookiesCompleted 在调用 ICoreWebView2CookieManager.GetCookies 时使用. 如果为 nil, 需自行调用 NewICoreWebView2GetCookiesCompletedHandler 来赋值.
	HandlerGetCookiesCompleted *ICoreWebView2GetCookiesCompletedHandler

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

	ref int32 // 引用计数
	// 权限map读写锁
	rwxPermissions sync.RWMutex
	// 权限map
	permissions map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE

	edge *Edge // WebView2 环境
	// 主机名
	hostName string
	// 嵌入的文件系统
	embedFS fs.FS
	// 第一次加载时使用的
	firstResponse *ICoreWebView2WebResourceResponse
}

func (e *Webview2Impl) QueryInterface(_, _ uintptr) uintptr {
	return 0
}

func (e *Webview2Impl) AddRef() uintptr {
	atomic.AddInt32(&e.ref, 1)
	return uintptr(e.ref)
}

func (e *Webview2Impl) Release() uintptr {
	atomic.AddInt32(&e.ref, -1)
	return uintptr(e.ref)
}

// --------------------------- 回调实现 ---------------------------

// TrySuspendCompleted 尝试挂起 webview 后调用, 以获取执行结果.
func (e *Webview2Impl) TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	if e.cbTrySuspendCompleted != nil {
		e.cbTrySuspendCompleted(errorCode, isSuccessful)
	}
	return 0
}

// ExecuteScriptCompleted 执行 js 脚本完成后调用, 以获取执行结果.
func (e *Webview2Impl) ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr {
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
func (e *Webview2Impl) AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr {
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
func (e *Webview2Impl) GetCookiesCompleted(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr {
	if e.cbGetCookiesCompleted != nil {
		e.cbGetCookiesCompleted(errorCode, cookies)
	}
	return 0
}

// --------------------------- 事件接口实现 ---------------------------

// DocumentTitleChanged 在文档标题发生变化时调用.
func (e *Webview2Impl) DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.cbDocumentTitleChangedEvent != nil {
		e.cbDocumentTitleChangedEvent(sender, args)
	}
	return 0
}

// RasterizationScaleChanged 在光栅化缩放比例发生变化时调用.
func (e *Webview2Impl) RasterizationScaleChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	if e.cbRasterizationScaleChangedEvent != nil {
		e.cbRasterizationScaleChangedEvent(sender, args)
	}
	return 0
}

// WindowCloseRequested 在窗口关闭请求时调用.
func (e *Webview2Impl) WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.cbWindowCloseRequestedEvent != nil {
		e.cbWindowCloseRequestedEvent(sender, args)
	}
	return 0
}

// SourceChanged 当源改变时调用。
func (e *Webview2Impl) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	if e.cbSourceChangedEvent != nil {
		e.cbSourceChangedEvent(sender, args)
	}
	return 0
}

// NewWindowRequested 当收到新窗口请求时调用。
func (e *Webview2Impl) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	if e.cbNewWindowRequestedEvent != nil {
		e.cbNewWindowRequestedEvent(sender, args)
	}
	return 0
}

// PermissionRequested 当收到权限请求时调用。
func (e *Webview2Impl) PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	kind, _ := args.GetPermissionKind()
	e.rwxPermissions.RLock()
	result, ok := e.permissions[kind]
	e.rwxPermissions.RUnlock()
	if !ok {
		result = COREWEBVIEW2_PERMISSION_STATE_DEFAULT
	}
	err := args.PutState(result)
	ReportErrorAtuo(err)

	if e.cbPermissionRequestedEvent != nil {
		e.cbPermissionRequestedEvent(sender, args)
	}
	return 0
}

// MessageReceived 当从 webview 收到消息时调用。
func (e *Webview2Impl) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
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

var once bool

// WebResourceRequested 当收到资源请求时调用。
func (e *Webview2Impl) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	defer func() {
		if e.cbWebResourceRequestedEvent != nil {
			e.cbWebResourceRequestedEvent(sender, args)
		}
	}()

	if e.hostName != "" { // 使用嵌入的文件系统映射
		request, err := args.GetRequest()
		if err != nil {
			ReportErrorAtuo(errors.New("webResourceRequested, GetRequest failed: " + err.Error()))
			return 0
		}

		// 获取请求uri
		uri := request.MustGetUri()
		// 判断请求路径是否以 hostName 开头
		if !strings.HasPrefix(uri, e.hostName) {
			return 0
		}

		// 去除 hostName
		embedPath := strings.TrimPrefix(uri, e.hostName)
		if embedPath == "" {
			return 0
		}

		// 从嵌入文件系统读取内容
		data, err := fs.ReadFile(e.embedFS, embedPath)
		if err != nil {
			// 固定会有一个对 favicon.ico 的请求, 没有这个文件的话, 这里肯定会触发一次, 没啥影响
			ReportErrorAtuo(errors.New("webResourceRequested, ReadFile failed1: " + err.Error()))
			// 返回 404
			res, err := e.edge.Environment.CreateWebResourceResponse(nil, 404, "Not Found", "")
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, CreateWebResourceResponse failed1: " + err.Error()))
				return 0
			}
			err = args.PutResponse(res)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, PutResponse failed1: " + err.Error()))
				return 0
			}
			return 0
		}

		var res *ICoreWebView2WebResourceResponse
		if !once && e.firstResponse != nil { // 仅第一次时这样
			once = true
			err = e.firstResponse.PutContent(data)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, PutContent failed: " + err.Error()))
			}
			res = e.firstResponse
			defer func() {
				e.firstResponse.Release()
				e.firstResponse = nil
			}()
		} else {
			// 返回动态生成的响应
			res, err = e.edge.Environment.CreateWebResourceResponse(data, 200, "OK", "")
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, CreateWebResourceResponse failed2: " + err.Error()))
				return 0
			}
		}

		err = args.PutResponse(res)
		if err != nil {
			ReportErrorAtuo(errors.New("webResourceRequested, PutResponse failed2: " + err.Error()))
			return 0
		}
	}
	return 0
}

// NavigationCompleted 当导航完成时调用。
func (e *Webview2Impl) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.cbNavigationCompletedEvent != nil {
		e.cbNavigationCompletedEvent(sender, args)
	}
	return 0
}

// NavigationCompleted2 当框架导航完成时调用。
func (e *Webview2Impl) NavigationCompleted2(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.cbFrameNavigationCompletedEvent != nil {
		e.cbFrameNavigationCompletedEvent(sender, args)
	}
	return 0
}

// NavigationStarting2 当框架导航开始时调用。
func (e *Webview2Impl) NavigationStarting2(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.cbFrameNavigationStartingEvent != nil {
		e.cbFrameNavigationStartingEvent(sender, args)
	}
	return 0
}

// NavigationStarting 当导航开始时调用。
func (e *Webview2Impl) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if e.cbNavigationStartingEvent != nil {
		e.cbNavigationStartingEvent(sender, args)
	}
	return 0
}

// AcceleratorKeyPressed 当使用快捷键时调用。
func (e *Webview2Impl) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if e.cbAcceleratorKeyPressedEvent != nil {
		e.cbAcceleratorKeyPressedEvent(sender, args)
	}
	return 0
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (e *Webview2Impl) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	if e.cbContentLoadingEvent != nil {
		e.cbContentLoadingEvent(sender, args)
	}
	return 0
}
