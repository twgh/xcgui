package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"

	"io/fs"
	"strings"
	"sync"
	"syscall"
)

// NewWebViewEventImpl 创建一个 WebView 事件接口实现对象.
//   - 使用此函数创建 WebViewEventImpl 对象, 用于给指定 WebView 添加更多的事件处理函数, 也就是一个类型的事件可以添加多个事件处理函数.
//   - 此种方法是高度封装的版本, 你也可以用更原生的方法来添加事件处理函数, 自行创建新的 EventHandler.
func NewWebViewEventImpl(wv *WebView) *WebViewEventImpl {
	return &WebViewEventImpl{
		CoreWebView: wv.CoreWebView,
		Controller:  wv.Controller,
		Edge:        wv.Edge,
	}
}

// WebViewEventImpl WebView 事件接口实现.
type WebViewEventImpl struct {
	IUnknown_Impl

	CoreWebView *ICoreWebView2           // CoreWebView2
	Controller  *ICoreWebView2Controller // WebView2 控制器
	Edge        *Edge                    // WebView2 环境

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
	// HandlerCapturePreviewCompleted 在调用 ICoreWebView2.CapturePreview 时使用. 如果为 nil, 需自行调用 NewICoreWebView2CapturePreviewCompletedHandler 来赋值.
	HandlerCapturePreviewCompleted *ICoreWebView2CapturePreviewCompletedHandler

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
	// HandlerContainsFullScreenElementChangedEvent 是 ContainsFullScreenElement(是否包含全屏元素) 属性改变事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerContainsFullScreenElementChangedEvent *ICoreWebView2ContainsFullScreenElementChangedEventHandler
	cbContainsFullScreenElementChangedEvent      func(sender *ICoreWebView2, args *IUnknown) uintptr
	// HandlerProcessFailedEvent 进程失败事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerProcessFailedEvent *ICoreWebView2ProcessFailedEventHandler
	cbProcessFailedEvent      func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr
	// HandlerHistoryChangedEvent 历史记录改变事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerHistoryChangedEvent *ICoreWebView2HistoryChangedEventHandler
	cbHistoryChangedEvent      func(sender *ICoreWebView2, args *IUnknown) uintptr
	// HandlerScriptDialogOpeningEvent 脚本对话框打开事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerScriptDialogOpeningEvent *ICoreWebView2ScriptDialogOpeningEventHandler
	cbScriptDialogOpeningEvent      func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr
	// HandlerDevToolsProtocolEventReceived DevTools 协议事件处理程序. 在调用 Event_ 时会自动赋值.
	HandlerDevToolsProtocolEventReceivedEvent *ICoreWebView2DevToolsProtocolEventReceivedEventHandler
	cbDevToolsProtocolEventReceivedEvent      func(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr
	// DevToolsProtocolEventReceiver 是 DevTools 协议事件接收器. 在调用 Event_ 时会自动赋值.
	DevToolsProtocolEventReceiver *ICoreWebView2DevToolsProtocolEventReceiver

	// -------------------- Callbacks --------------------

	// 仅供内部使用的网页消息事件回调
	msgcb_xcgui func(string)
	// 挂起事件结果回调
	cbTrySuspendCompleted func(errorCode syscall.Errno, isSuccessful bool) uintptr
	// 执行脚本完成事件回调
	cbExecuteScriptCompleted func(errorCode syscall.Errno, result string) uintptr
	// 在文档创建前添加脚本完成回调
	cbAddScriptToExecuteOnDocumentCreatedCompleted func(errorCode syscall.Errno, id string) uintptr
	// 获取 Cookies 完成回调
	cbGetCookiesCompleted func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr
	// 截图完成回调
	cbCapturePreviewCompleted func(errorCode syscall.Errno) uintptr

	// 权限map读写锁
	rwxPermissions sync.RWMutex
	// 权限map
	permissions map[COREWEBVIEW2_PERMISSION_KIND]COREWEBVIEW2_PERMISSION_STATE

	// 主机名
	hostName string
	// 嵌入的文件系统
	embedFS fs.FS
	// 第一次加载时使用的
	firstResponse *ICoreWebView2WebResourceResponse
}

// ReleaseEventHandler 释放事件处理程序, 防止内存泄漏.
//   - 此方法会在关闭 WebView 前自动调用, 无需手动调用.
//   - 如果此 WebViewEventImpl 对象实例是你另外手动创建的, 需要在关闭 WebView 前自行调用此方法释放.
func (w *WebViewEventImpl) ReleaseEventHandler() {
	if w.DevToolsProtocolEventReceiver != nil {
		w.DevToolsProtocolEventReceiver.Release()
		w.DevToolsProtocolEventReceiver = nil
	}
	if w.HandlerTrySuspendCompleted != nil {
		w.HandlerTrySuspendCompleted.Release()
		w.HandlerTrySuspendCompleted = nil
	}
	if w.HandlerExecuteScriptCompleted != nil {
		w.HandlerExecuteScriptCompleted.Release()
		w.HandlerExecuteScriptCompleted = nil
	}
	if w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted != nil {
		w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted.Release()
		w.HandlerAddScriptToExecuteOnDocumentCreatedCompleted = nil
	}
	if w.HandlerGetCookiesCompleted != nil {
		w.HandlerGetCookiesCompleted.Release()
		w.HandlerGetCookiesCompleted = nil
	}
	if w.HandlerCapturePreviewCompleted != nil {
		w.HandlerCapturePreviewCompleted.Release()
		w.HandlerCapturePreviewCompleted = nil
	}
	if w.handlerWebMessageReceivedEvent != nil {
		w.handlerWebMessageReceivedEvent.Release()
		w.handlerWebMessageReceivedEvent = nil
	}
	if w.handlerPermissionRequestedEvent != nil {
		w.handlerPermissionRequestedEvent.Release()
		w.handlerPermissionRequestedEvent = nil
	}
	if w.HandlerWebResourceRequestedEvent != nil {
		w.HandlerWebResourceRequestedEvent.Release()
		w.HandlerWebResourceRequestedEvent = nil
	}
	if w.HandlerAcceleratorKeyPressedEvent != nil {
		w.HandlerAcceleratorKeyPressedEvent.Release()
		w.HandlerAcceleratorKeyPressedEvent = nil
	}
	if w.HandlerNavigationCompletedEvent != nil {
		w.HandlerNavigationCompletedEvent.Release()
		w.HandlerNavigationCompletedEvent = nil
	}
	if w.HandlerNavigationStartingEvent != nil {
		w.HandlerNavigationStartingEvent.Release()
		w.HandlerNavigationStartingEvent = nil
	}
	if w.HandlerContentLoadingEvent != nil {
		w.HandlerContentLoadingEvent.Release()
		w.HandlerContentLoadingEvent = nil
	}
	if w.HandlerNewWindowRequestedEvent != nil {
		w.HandlerNewWindowRequestedEvent.Release()
		w.HandlerNewWindowRequestedEvent = nil
	}
	if w.HandlerSourceChangedEvent != nil {
		w.HandlerSourceChangedEvent.Release()
		w.HandlerSourceChangedEvent = nil
	}
	if w.HandlerFrameNavigationStartingEvent != nil {
		w.HandlerFrameNavigationStartingEvent.Release()
		w.HandlerFrameNavigationStartingEvent = nil
	}
	if w.HandlerFrameNavigationCompletedEvent != nil {
		w.HandlerFrameNavigationCompletedEvent.Release()
		w.HandlerFrameNavigationCompletedEvent = nil
	}
	if w.HandlerWindowCloseRequestedEvent != nil {
		w.HandlerWindowCloseRequestedEvent.Release()
		w.HandlerWindowCloseRequestedEvent = nil
	}
	if w.HandlerDocumentTitleChangedEvent != nil {
		w.HandlerDocumentTitleChangedEvent.Release()
		w.HandlerDocumentTitleChangedEvent = nil
	}
	if w.HandlerRasterizationScaleChangedEvent != nil {
		w.HandlerRasterizationScaleChangedEvent.Release()
		w.HandlerRasterizationScaleChangedEvent = nil
	}
	if w.HandlerContainsFullScreenElementChangedEvent != nil {
		w.HandlerContainsFullScreenElementChangedEvent.Release()
		w.HandlerContainsFullScreenElementChangedEvent = nil
	}
	if w.HandlerProcessFailedEvent != nil {
		w.HandlerProcessFailedEvent.Release()
		w.HandlerProcessFailedEvent = nil
	}
	if w.HandlerHistoryChangedEvent != nil {
		w.HandlerHistoryChangedEvent.Release()
		w.HandlerHistoryChangedEvent = nil
	}
	if w.HandlerScriptDialogOpeningEvent != nil {
		w.HandlerScriptDialogOpeningEvent.Release()
		w.HandlerScriptDialogOpeningEvent = nil
	}
	if w.HandlerDevToolsProtocolEventReceivedEvent != nil {
		w.HandlerDevToolsProtocolEventReceivedEvent.Release()
		w.HandlerDevToolsProtocolEventReceivedEvent = nil
	}
}

// --------------------------- 回调实现 ---------------------------

// TrySuspendCompleted 尝试挂起 webview 后调用, 以获取执行结果.
func (w *WebViewEventImpl) TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	if w.cbTrySuspendCompleted != nil {
		w.cbTrySuspendCompleted(errorCode, isSuccessful)
	}
	return 0
}

// ExecuteScriptCompleted 执行 js 脚本完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr {
	if w.cbExecuteScriptCompleted != nil {
		var str string
		if result != nil && *result != 0 {
			str = common.UTF16PtrToString(result)
		}
		w.cbExecuteScriptCompleted(errorCode, str)
	}
	return 0
}

// AddScriptToExecuteOnDocumentCreatedCompleted 添加 js 脚本完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr {
	if w.cbAddScriptToExecuteOnDocumentCreatedCompleted != nil {
		var idStr string
		if id != nil && *id != 0 {
			idStr = common.UTF16PtrToString(id)
		}
		w.cbAddScriptToExecuteOnDocumentCreatedCompleted(errorCode, idStr)
	}
	return 0
}

// GetCookiesCompleted 获取 cookies 完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) GetCookiesCompleted(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr {
	if w.cbGetCookiesCompleted != nil {
		w.cbGetCookiesCompleted(errorCode, cookies)
	}
	return 0
}

// CapturePreviewCompleted 截图完成时调用.
func (w *WebViewEventImpl) CapturePreviewCompleted(errorCode syscall.Errno) uintptr {
	if w.cbCapturePreviewCompleted != nil {
		w.cbCapturePreviewCompleted(errorCode)
	}
	return 0
}

// --------------------------- 事件接口实现 ---------------------------

// DocumentTitleChanged 在文档标题发生变化时调用.
func (w *WebViewEventImpl) DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if w.cbDocumentTitleChangedEvent != nil {
		w.cbDocumentTitleChangedEvent(sender, args)
	}
	return 0
}

// RasterizationScaleChanged 在光栅化缩放比例发生变化时调用.
func (w *WebViewEventImpl) RasterizationScaleChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	if w.cbRasterizationScaleChangedEvent != nil {
		w.cbRasterizationScaleChangedEvent(sender, args)
	}
	return 0
}

// WindowCloseRequested 在窗口关闭请求时调用.
func (w *WebViewEventImpl) WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr {
	if w.cbWindowCloseRequestedEvent != nil {
		w.cbWindowCloseRequestedEvent(sender, args)
	}
	return 0
}

// SourceChanged 当源改变时调用。
func (w *WebViewEventImpl) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	if w.cbSourceChangedEvent != nil {
		w.cbSourceChangedEvent(sender, args)
	}
	return 0
}

// NewWindowRequested 当收到新窗口请求时调用。
func (w *WebViewEventImpl) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	if w.cbNewWindowRequestedEvent != nil {
		w.cbNewWindowRequestedEvent(sender, args)
	}
	return 0
}

// PermissionRequested 当收到权限请求时调用。
func (w *WebViewEventImpl) PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	if w.permissions != nil {
		kind, _ := args.GetPermissionKind()
		w.rwxPermissions.RLock()
		result, ok := w.permissions[kind]
		w.rwxPermissions.RUnlock()
		if !ok {
			result = COREWEBVIEW2_PERMISSION_STATE_DEFAULT
		}
		err := args.PutState(result)
		ReportErrorAtuo(err)
	}

	if w.cbPermissionRequestedEvent != nil {
		w.cbPermissionRequestedEvent(sender, args)
	}
	return 0
}

// MessageReceived 当从 webview 收到消息时调用。
func (w *WebViewEventImpl) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	if w.msgcb_xcgui != nil {
		message, err := args.TryGetWebMessageAsString()
		if err == nil && strings.HasPrefix(message, "{\"id\":") {
			w.msgcb_xcgui(message)
		}
	}
	if w.cbMessageReceivedEvent != nil {
		w.cbMessageReceivedEvent(sender, args)
	}
	return 0
}

var once bool

// WebResourceRequested 当收到资源请求时调用。
func (w *WebViewEventImpl) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	defer func() {
		if w.cbWebResourceRequestedEvent != nil {
			w.cbWebResourceRequestedEvent(sender, args)
		}
	}()

	if w.hostName != "" { // 使用嵌入的文件系统映射
		request, err := args.GetRequest()
		if err != nil {
			ReportErrorAtuo(errors.New("webResourceRequested, GetRequest failed: " + err.Error()))
			return 0
		}

		// 获取请求uri
		uri := request.MustGetUri()
		// 判断请求路径是否以 hostName 开头
		if !strings.HasPrefix(uri, w.hostName) {
			return 0
		}

		// 去除 hostName
		embedPath := strings.TrimPrefix(uri, w.hostName)
		if embedPath == "" {
			return 0
		}

		// 从嵌入文件系统读取内容
		data, err := fs.ReadFile(w.embedFS, embedPath)
		if err != nil {
			// 固定会有一个对 favicon.ico 的请求, 没有这个文件的话, 这里肯定会触发一次, 没啥影响
			ReportErrorAtuo(errors.New("webResourceRequested, ReadFile failed1: " + err.Error()))
			// 返回 404
			res, err := w.Edge.Environment.CreateWebResourceResponse(nil, 404, "Not Found", "")
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
		if !once && w.firstResponse != nil { // 仅第一次时这样
			once = true
			err = w.firstResponse.PutContent(data)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, PutContent failed: " + err.Error()))
			}
			res = w.firstResponse
			defer func() {
				w.firstResponse.Release()
				w.firstResponse = nil
			}()
		} else {
			// 返回动态生成的响应
			res, err = w.Edge.Environment.CreateWebResourceResponse(data, 200, "OK", "")
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
func (w *WebViewEventImpl) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if w.cbNavigationCompletedEvent != nil {
		w.cbNavigationCompletedEvent(sender, args)
	}
	return 0
}

// NavigationCompleted2 当框架导航完成时调用。
func (w *WebViewEventImpl) NavigationCompleted2(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if w.cbFrameNavigationCompletedEvent != nil {
		w.cbFrameNavigationCompletedEvent(sender, args)
	}
	return 0
}

// NavigationStarting2 当框架导航开始时调用。
func (w *WebViewEventImpl) NavigationStarting2(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if w.cbFrameNavigationStartingEvent != nil {
		w.cbFrameNavigationStartingEvent(sender, args)
	}
	return 0
}

// NavigationStarting 当导航开始时调用。
func (w *WebViewEventImpl) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	if w.cbNavigationStartingEvent != nil {
		w.cbNavigationStartingEvent(sender, args)
	}
	return 0
}

// AcceleratorKeyPressed 当使用快捷键时调用。
func (w *WebViewEventImpl) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if w.cbAcceleratorKeyPressedEvent != nil {
		w.cbAcceleratorKeyPressedEvent(sender, args)
	}
	return 0
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (w *WebViewEventImpl) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	if w.cbContentLoadingEvent != nil {
		w.cbContentLoadingEvent(sender, args)
	}
	return 0
}

// ContainsFullScreenElementChanged 当 ContainsFullScreenElement(是否包含全屏元素) 属性改变时调用。
func (w *WebViewEventImpl) ContainsFullScreenElementChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if w.cbContainsFullScreenElementChangedEvent != nil {
		w.cbContainsFullScreenElementChangedEvent(sender, args)
	}
	return 0
}

// ProcessFailed 当 Webview 的进程因任何原因失败时调用。
func (w *WebViewEventImpl) ProcessFailed(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr {
	if w.cbProcessFailedEvent != nil {
		w.cbProcessFailedEvent(sender, args)
	}
	return 0
}

// HistoryChanged 在联合会话历史记录发生更改时引发，该历史记录由顶级和手动框架导航组成。
func (w *WebViewEventImpl) HistoryChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if w.cbHistoryChangedEvent != nil {
		w.cbHistoryChangedEvent(sender, args)
	}
	return 0
}

// ScriptDialogOpening 当脚本对话框打开时调用。alert、confirm、prompt 或 beforeunload。
func (w *WebViewEventImpl) ScriptDialogOpening(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
	if w.cbScriptDialogOpeningEvent != nil {
		w.cbScriptDialogOpeningEvent(sender, args)
	}
	return 0
}

// DevToolsProtocolEventReceived 当收到来自 DevTools 协议的事件时调用。
func (w *WebViewEventImpl) DevToolsProtocolEventReceived(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr {
	if w.cbDevToolsProtocolEventReceivedEvent != nil {
		w.cbDevToolsProtocolEventReceivedEvent(sender, args)
	}
	return 0
}
