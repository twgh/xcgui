package edge

import (
	"errors"
	"io/fs"
	"strings"
	"sync"
	"syscall"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// NewWebViewEventImpl 创建一个 WebView 事件接口实现对象.
//   - 此种方法是高度封装的版本, 你也可以用更原生的方法来添加事件处理函数, 自行创建新的 EventHandler.
//   - 需要注意的是你自己创建的这个对象不能让它被 GC 回收, 也就是不能创建为局部变量. 或可使用 runtime.Pinner.
func NewWebViewEventImpl(wv *WebView) *WebViewEventImpl {
	return &WebViewEventImpl{
		CoreWebView: wv.CoreWebView,
		Controller:  wv.Controller,
		Edge:        wv.Edge,
	}
}

// WebViewEventImpl 是 WebView 事件接口实现.
type WebViewEventImpl struct {
	IUnknown_Impl

	CoreWebView *ICoreWebView2           // CoreWebView2
	Controller  *ICoreWebView2Controller // WebView2 控制器
	Edge        *Edge                    // WebView2 环境

	// 仅供内部使用的网页消息事件回调
	msgcb_xcgui func(string)

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

// CreateNewWebViewEventImpl 获取一个新的 WebView 事件接口实现对象.
//   - 此种方法是高度封装的版本, 你也可以用更原生的方法来添加事件处理函数, 自行创建新的 EventHandler.
//   - 需要注意的是你自己创建的这个对象不能让它被 GC 回收, 也就是不能创建为局部变量. 或可使用 runtime.Pinner.
func (w *WebViewEventImpl) CreateNewWebViewEventImpl() *WebViewEventImpl {
	return &WebViewEventImpl{
		CoreWebView: w.CoreWebView,
		Controller:  w.Controller,
		Edge:        w.Edge,
	}
}

// --------------------------- 回调实现 ---------------------------

// TrySuspendCompleted 尝试挂起 webview 后调用, 以获取执行结果.
func (w *WebViewEventImpl) TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "TrySuspendCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, isSuccessful bool) uintptr)(errorCode, isSuccessful)
	}
	return ret
}

// ExecuteScriptCompleted 执行 js 脚本完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr {
	var str string
	cbs := WvEventHandler.GetCallBacks(w, "ExecuteScriptCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && result != nil && *result != 0 {
			str = common.UTF16PtrToString(result)
		}
	}
	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result string) uintptr)(errorCode, str)
	}
	return ret
}

// CallDevToolsProtocolMethodCompleted 调用 DevTools 协议方法完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) CallDevToolsProtocolMethodCompleted(errorCode syscall.Errno, result *uint16) uintptr {
	var str string
	cbs := WvEventHandler.GetCallBacks(w, "CallDevToolsProtocolMethodCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && result != nil && *result != 0 {
			str = common.UTF16PtrToString(result)
		}
	}
	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, id string) uintptr)(errorCode, str)
	}
	return ret
}

// AddScriptToExecuteOnDocumentCreatedCompleted 添加 js 脚本完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr {
	var str string
	cbs := WvEventHandler.GetCallBacks(w, "AddScriptToExecuteOnDocumentCreatedCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && id != nil && *id != 0 {
			str = common.UTF16PtrToString(id)
		}
	}
	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, id string) uintptr)(errorCode, str)
	}
	return ret
}

// GetCookiesCompleted 获取 cookies 完成后调用, 以获取执行结果.
func (w *WebViewEventImpl) GetCookiesCompleted(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "GetCookiesCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr)(errorCode, cookies)
	}
	return ret
}

// CapturePreviewCompleted 截图完成时调用.
func (w *WebViewEventImpl) CapturePreviewCompleted(errorCode syscall.Errno) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "CapturePreviewCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno) uintptr)(errorCode)
	}
	return ret
}

// WebResourceResponseViewGetContentCompleted 当 ICoreWebView2WebResourceResponseView.GetContent 完成时调用。
func (w *WebViewEventImpl) WebResourceResponseViewGetContentCompleted(errorCode syscall.Errno, result *IStream) uintptr {
	if result != nil {
		defer result.Release()
	}

	var bs []byte
	cbs := WvEventHandler.GetCallBacks(w, "WebResourceResponseViewGetContentCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && result != nil {
			var err error
			bs, err = result.GetBytes()
			if err != nil {
				ReportErrorAtuo(errors.New("WebResourceResponseViewGetContentCompleted, GetBytes failed: " + err.Error()))
			}
		}
	}

	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result []byte) uintptr)(errorCode, bs)
	}
	return ret
}

// --------------------------- 事件接口实现 ---------------------------

// DocumentTitleChanged 在文档标题发生变化时调用.
func (w *WebViewEventImpl) DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "DocumentTitleChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// RasterizationScaleChanged 在光栅化缩放比例发生变化时调用.
func (w *WebViewEventImpl) RasterizationScaleChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "RasterizationScaleChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// WindowCloseRequested 在窗口关闭请求时调用.
func (w *WebViewEventImpl) WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "WindowCloseRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// SourceChanged 当源改变时调用。
func (w *WebViewEventImpl) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "SourceChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// NewWindowRequested 当收到新窗口请求时调用。
func (w *WebViewEventImpl) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NewWindowRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
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
		err := args.SetState(result)
		ReportErrorAtuo(err)
	}

	cbs := WvEventHandler.GetCallBacks(w, "PermissionRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// WebMessageReceived 当从 webview 收到消息时调用。
func (w *WebViewEventImpl) WebMessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	if w.msgcb_xcgui != nil {
		message, err := args.TryGetWebMessageAsString()
		if err == nil {
			if strings.HasPrefix(message, "{\"id\":") {
				w.msgcb_xcgui(message)
			}
		}
	}

	cbs := WvEventHandler.GetCallBacks(w, "WebMessageReceived")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr)(sender, args)
	}
	return ret
}

var once bool

// WebResourceRequested 当收到资源请求时调用。
func (w *WebViewEventImpl) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	var ret uintptr
	defer func() {
		cbs := WvEventHandler.GetCallBacks(w, "WebResourceRequested")
		for i := len(cbs) - 1; i >= 0; i-- {
			ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr)(sender, args)
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
			err = args.SetResponse(res)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, SetResponse failed1: " + err.Error()))
				return 0
			}
			return 0
		}

		var res *ICoreWebView2WebResourceResponse
		if !once && w.firstResponse != nil { // 仅第一次时这样
			once = true
			stream, err := NewStreamMem(data)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, create stream failed: " + err.Error()))
			} else {
				defer stream.Release()
			}
			err = w.firstResponse.SetContent(stream)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, SetContent failed: " + err.Error()))
			}
			res = w.firstResponse
			defer func() {
				w.firstResponse.Release()
				w.firstResponse = nil
			}()
		} else { // 返回动态生成的响应
			// 创建响应流
			stream, err := NewStreamMem(data)
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, create stream failed2: " + err.Error()))
			} else {
				defer stream.Release()
			}
			res, err = w.Edge.Environment.CreateWebResourceResponse(stream, 200, "OK", "")
			if err != nil {
				ReportErrorAtuo(errors.New("webResourceRequested, CreateWebResourceResponse failed2: " + err.Error()))
				return 0
			}
		}

		err = args.SetResponse(res)
		if err != nil {
			ReportErrorAtuo(errors.New("webResourceRequested, SetResponse failed2: " + err.Error()))
			return 0
		}
	}
	return ret
}

// NavigationCompleted 当导航完成时调用。
func (w *WebViewEventImpl) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NavigationCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// Frame_NavigationCompleted 当 ICoreWebView2 的框架导航完成时调用。
func (w *WebViewEventImpl) Frame_NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "Frame_NavigationCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// Frame_NavigationStarting 当 ICoreWebView2 的框架导航开始时调用。
func (w *WebViewEventImpl) Frame_NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "Frame_NavigationStarting")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// NavigationStarting 当导航开始时调用。
func (w *WebViewEventImpl) NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NavigationStarting")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// AcceleratorKeyPressed 当使用快捷键时调用。
func (w *WebViewEventImpl) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "AcceleratorKeyPressed")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ContentLoading 当 WebView 控件开始加载内容时调用。
func (w *WebViewEventImpl) ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ContentLoading")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ContainsFullScreenElementChanged 当 ContainsFullScreenElement(是否包含全屏元素) 属性改变时调用。
func (w *WebViewEventImpl) ContainsFullScreenElementChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ContainsFullScreenElementChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// ProcessFailed 当 Webview 的进程因任何原因失败时调用。
func (w *WebViewEventImpl) ProcessFailed(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ProcessFailed")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// HistoryChanged 在联合会话历史记录发生更改时引发，该历史记录由顶级和手动框架导航组成。
func (w *WebViewEventImpl) HistoryChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "HistoryChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// ScriptDialogOpening 当脚本对话框打开时调用。alert、confirm、prompt 或 beforeunload。
func (w *WebViewEventImpl) ScriptDialogOpening(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ScriptDialogOpening")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr)(sender, args)
	}
	return ret
}

// DevToolsProtocolEventReceived 当收到来自 DevTools 协议的事件时调用。
func (w *WebViewEventImpl) DevToolsProtocolEventReceived(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "DevToolsProtocolEventReceived")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// WebResourceResponseReceived 当 WebView 接收到对网络资源请求的响应时调用。
//   - WebView 执行的任何 URI 解析；例如 HTTP/HTTPS、来自重定向、导航、HTML 声明、隐式 favicon 图标的查找和数据请求，以及文档中 fetch API 的使用都会触发此事件。
//   - 宿主应用可以使用此事件查看网络资源的实际请求和响应。
//   - 无法保证 WebView 处理响应的顺序与宿主应用的处理程序运行的顺序。
//   - 应用的处理程序不会阻止 WebView 处理响应。
func (w *WebViewEventImpl) WebResourceResponseReceived(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "WebResourceResponseReceived")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// DOMContentLoaded 当初始 HTML 文档解析完成时触发。
func (w *WebViewEventImpl) DOMContentLoaded(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "DOMContentLoaded")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// FrameCreated 当创建新的 iframe 时触发。
func (w *WebViewEventImpl) FrameCreated(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "FrameCreated")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// FrameNameChanged 当框架名称更改时调用。
func (w *WebViewEventImpl) FrameNameChanged(sender *ICoreWebView2Frame, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "FrameNameChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Frame, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// FrameDestroyed 当框架被销毁时调用。
func (w *WebViewEventImpl) FrameDestroyed(sender *ICoreWebView2Frame, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "FrameDestroyed")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Frame, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// DownloadStarting 当下载开始时调用。
func (w *WebViewEventImpl) DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "DownloadStarting")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// BytesReceivedChanged 当下载的字节数更改时调用。
func (w *WebViewEventImpl) BytesReceivedChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "BytesReceivedChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// ClientCertificateRequested 当 WebView2 向需要客户端证书进行 HTTP 身份验证的 HTTP 服务器发出请求时触发.
func (w *WebViewEventImpl) ClientCertificateRequested(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ClientCertificateRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// IsMutedChanged 当 WebView 的静音状态更改时调用。
func (w *WebViewEventImpl) IsMutedChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "IsMutedChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// DocumentPlayingAudioChanged 当文档播放音频状态改变时调用。
func (w *WebViewEventImpl) DocumentPlayingAudioChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "DocumentPlayingAudioChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// ContextMenuRequested 当用户请求上下文菜单，且 WebView 内部的内容未禁用上下文菜单时调用。
func (w *WebViewEventImpl) ContextMenuRequested(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ContextMenuRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// CustomItemSelected 当用户选择自定义上下文菜单项时调用。
func (w *WebViewEventImpl) CustomItemSelected(sender *ICoreWebView2ContextMenuItem, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "CustomItemSelected")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2ContextMenuItem, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// CreateCoreWebView2CompositionControllerCompleted 当创建 WebView2CompositionController 完成时调用。
func (w *WebViewEventImpl) CreateCoreWebView2CompositionControllerCompleted(errorCode syscall.Errno, result *ICoreWebView2CompositionController) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "CreateCoreWebView2CompositionControllerCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result *ICoreWebView2CompositionController) uintptr)(errorCode, result)
	}
	return ret
}

// CursorChanged 当光标更改时调用。
func (w *WebViewEventImpl) CursorChanged(sender *ICoreWebView2CompositionController, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "CursorChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2CompositionController, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// BrowserProcessExited 当浏览器进程退出时调用。
func (w *WebViewEventImpl) BrowserProcessExited(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "BrowserProcessExited")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ProcessInfosChanged 当 WebView2 环境中的进程信息发生更改时调用。
func (w *WebViewEventImpl) ProcessInfosChanged(sender *ICoreWebView2Environment, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ProcessInfosChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Environment, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// GetProcessExtendedInfosCompleted 获取进程扩展信息完成时调用。
func (w *WebViewEventImpl) GetProcessExtendedInfosCompleted(errorCode syscall.Errno, result *ICoreWebView2ProcessExtendedInfoCollection) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "GetProcessExtendedInfosCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result *ICoreWebView2ProcessExtendedInfoCollection) uintptr)(errorCode, result)
	}
	return ret
}

// FaviconChanged 当网站图标（favicon）的 URL 与之前的 URL 不同时触发。
//   - 首次导航到新文档时，无论该文档是否在 HTML 中声明了网站图标，只要其图标与之前的图标不同，就会触发 FaviconChanged 事件。
//   - 如果 HTML 中声明了网站图标或通过脚本设置了网站图标，该事件将再次触发。随后可以通过 GetFavicon 和 GetFaviconUri 方法获取网站图标信息。
func (w *WebViewEventImpl) FaviconChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "FaviconChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// GetFaviconCompleted 获取网站图标完成时调用。
func (w *WebViewEventImpl) GetFaviconCompleted(errorCode syscall.Errno, faviconStream *IStream) uintptr {
	if faviconStream != nil {
		defer faviconStream.Release()
	}

	var bs []byte
	cbs := WvEventHandler.GetCallBacks(w, "GetFaviconCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && faviconStream != nil {
			var err error
			bs, err = faviconStream.GetBytes()
			if err != nil {
				ReportErrorAtuo(errors.New("GetFaviconCompleted, GetBytes failed: " + err.Error()))
			}
		}
	}

	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result []byte) uintptr)(errorCode, bs)
	}
	return ret
}

// ZoomFactorChanged 当 WebView 的缩放因子改变时调用。
//   - WebView 会关联每个网站最后使用的缩放比例。导航到不同页面时，缩放比例可能会发生变化。当因导航变化导致缩放比例改变时，ZoomFactorChanged 事件会在 ContentLoading 事件之后立即运行。
func (w *WebViewEventImpl) ZoomFactorChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ZoomFactorChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// MoveFocusRequested 移动焦点请求时调用。
//   - MoveFocusRequested 在用户尝试按 Tab 键离开 WebView 时运行。此事件运行时，WebView 的焦点尚未改变。
func (w *WebViewEventImpl) MoveFocusRequested(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "MoveFocusRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// FocusChanged 焦点变化时调用(当前仅用于 GotFocus 事件)。
func (w *WebViewEventImpl) FocusChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "FocusChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// LostFocus 失去焦点时调用。
func (w *WebViewEventImpl) LostFocus(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "LostFocus")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Controller, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// NewBrowserVersionAvailable 当有新版本的 WebView2 运行时可用时调用。
func (w *WebViewEventImpl) NewBrowserVersionAvailable(sender *ICoreWebView2Environment, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NewBrowserVersionAvailable")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Environment, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// EstimatedEndTimeChanged 当下载的预计结束时间更改时调用。
func (w *WebViewEventImpl) EstimatedEndTimeChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "EstimatedEndTimeChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// PrintToPdfCompleted 当打印到 PDF 操作完成时调用。
func (w *WebViewEventImpl) PrintToPdfCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "PrintToPdfCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, isSuccessful bool) uintptr)(errorCode, isSuccessful)
	}
	return ret
}

// StateChanged 当下载的状态更改时调用。
func (w *WebViewEventImpl) StateChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "StateChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// IsDefaultDownloadDialogOpenChanged 当默认下载对话框的打开状态发生变化时调用。
//   - 此事件在 DownloadStarting 事件之后发生。
//   - 在 DownloadStartingEventArgs 上设置 Handled 属性会禁用默认下载对话框，并确保此事件永远不会被触发。
func (w *WebViewEventImpl) IsDefaultDownloadDialogOpenChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "IsDefaultDownloadDialogOpenChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// BasicAuthenticationRequested 当网页请求基本身份验证时调用。
//   - 主机可以提供包含身份验证凭据的响应，也可以取消请求。
//   - 如果主机将 Cancel 属性设置为 false，但未在 Response 属性上提供 UserName 或 Password 属性，那么 WebView2 将向用户显示默认的身份验证质询对话框提示。
func (w *WebViewEventImpl) BasicAuthenticationRequested(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "BasicAuthenticationRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// StatusBarTextChanged 当状态栏文本改变时调用。
//   - 当 WebView 显示状态消息、URL 或空字符串（表示隐藏状态栏）时触发。
func (w *WebViewEventImpl) StatusBarTextChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "StatusBarTextChanged")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// ServerCertificateErrorDetected 当 WebView2 在加载网页时无法验证服务器的数字证书时调用。
//   - 此事件将针对所有网络资源触发，并紧随 WebResourceRequested 事件之后。
//   - 如果不处理该事件，WebView2 会在导航时向用户显示默认的 TLS 插页式错误页面，而对于非导航操作，Web 请求会被取消。
func (w *WebViewEventImpl) ServerCertificateErrorDetected(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ServerCertificateErrorDetected")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ClearServerCertificateErrorActionsCompleted 当清除服务器证书错误操作完成时调用。
func (w *WebViewEventImpl) ClearServerCertificateErrorActionsCompleted(errorCode syscall.Errno) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ClearServerCertificateErrorActionsCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno) uintptr)(errorCode)
	}
	return ret
}

// PrintCompleted 当 ICoreWebView2_16.Print 方法执行完成时调用。
func (w *WebViewEventImpl) PrintCompleted(errorCode syscall.Errno, result COREWEBVIEW2_PRINT_STATUS) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "PrintCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result COREWEBVIEW2_PRINT_STATUS) uintptr)(errorCode, result)
	}
	return ret
}

// PrintToPdfStreamCompleted 当 ICoreWebView2_16.PrintToPdfStream 执行完成时调用。
func (w *WebViewEventImpl) PrintToPdfStreamCompleted(errorCode syscall.Errno, pdfStream *IStream) uintptr {
	if pdfStream != nil {
		defer pdfStream.Release()
	}

	var bs []byte
	cbs := WvEventHandler.GetCallBacks(w, "PrintToPdfStreamCompleted")
	n := len(cbs)
	if n > 0 {
		if errors.Is(errorCode, wapi.S_OK) && pdfStream != nil {
			var err error
			bs, err = pdfStream.GetBytes()
			if err != nil {
				ReportErrorAtuo(errors.New("PrintToPdfStreamCompleted, GetBytes failed: " + err.Error()))
			}
		}
	}

	var ret uintptr
	for i := n - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, pdfBytes []byte) uintptr)(errorCode, bs)
	}
	return ret
}

// LaunchingExternalUriScheme 当导航请求指向已在操作系统中注册的 URI 方案时触发。
//   - 事件处理程序可以抑制默认对话框，或者用自定义对话框替换默认对话框。
//   - 如果未对事件参数执行延迟操作，外部 URI 方案的启动将被阻止，直到事件处理程序返回。
//   - 如果执行了延迟操作，外部 URI 方案的启动将被阻止，直到延迟完成。主机还可以选择取消 URI 方案的启动。
//   - 无论 Cancel 属性设置为 TRUE 还是 FALSE，都会触发 NavigationStarting 和 NavigationCompleted 事件。
//   - 无论主机是否在 ICoreWebView2LaunchingExternalUriSchemeEventArgs 上设置 Cancel 属性，NavigationCompleted 事件触发时，IsSuccess 属性都将设置为 FALSE，WebErrorStatus 属性都将设置为 ConnectionAborted。
//   - 对于此次导航到外部 URI 方案，无论 Cancel 属性如何设置，都不会触发 SourceChanged、ContentLoading 和 HistoryChanged 事件。
//   - LaunchingExternalUriScheme 事件将在 NavigationStarting 事件之后、NavigationCompleted 事件之前触发。
//   - 导航到外部 URI 方案时，默认的 ICoreWebView2Settings 也会更新。如果 ICoreWebView2Settings 接口上的某个设置已更改，导航到外部 URI 方案将触发 ICoreWebView2Settings 更新。
func (w *WebViewEventImpl) LaunchingExternalUriScheme(sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "LaunchingExternalUriScheme")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ExecuteScriptWithResultCompleted 当 ICoreWebView2_21.ExecuteScriptWithResult 方法执行完成时调用。
func (w *WebViewEventImpl) ExecuteScriptWithResultCompleted(errorCode syscall.Errno, args *ICoreWebView2ExecuteScriptResult) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ExecuteScriptWithResultCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, args *ICoreWebView2ExecuteScriptResult) uintptr)(errorCode, args)
	}
	return ret
}

// NotificationReceived 当 WebView2 接收到网页通知时调用。
//   - 如果未对事件参数执行延迟操作，那么在 DOM 通知创建调用（即 Notification()）之后的后续脚本会被阻塞，直到事件处理程序返回。
//   - 如果执行了延迟操作，脚本会被阻塞，直到延迟完成。
func (w *WebViewEventImpl) NotificationReceived(sender *ICoreWebView2, args *ICoreWebView2NotificationReceivedEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NotificationReceived")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2NotificationReceivedEventArgs) uintptr)(sender, args)
	}
	return ret
}

// NotificationCloseRequested 当通知被网页代码关闭时（例如通过 notification.close()）触发。
//   - 由于这是来自网页代码的操作，因此无需调用 ReportClosed。
func (w *WebViewEventImpl) NotificationCloseRequested(sender *ICoreWebView2Notification, args *IUnknown) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "NotificationCloseRequested")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2Notification, args *IUnknown) uintptr)(sender, args)
	}
	return ret
}

// SaveAsUIShowing 当通过编程方式或手动方式触发“另存为”时调用。
func (w *WebViewEventImpl) SaveAsUIShowing(sender *ICoreWebView2, args *ICoreWebView2SaveAsUIShowingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "SaveAsUIShowing")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2SaveAsUIShowingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ShowSaveAsUICompleted 当"另存为"UI操作完成时调用。
func (w *WebViewEventImpl) ShowSaveAsUICompleted(errorCode syscall.Errno, result COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ShowSaveAsUICompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno, result COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr)(errorCode, result)
	}
	return ret
}

// SaveFileSecurityCheckStarting 在系统 FileTypePolicy 检查危险文件扩展名列表时触发。
func (w *WebViewEventImpl) SaveFileSecurityCheckStarting(sender *ICoreWebView2, args *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "SaveFileSecurityCheckStarting")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// ScreenCaptureStarting 当屏幕截图开始时调用。
func (w *WebViewEventImpl) ScreenCaptureStarting(sender *ICoreWebView2, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "ScreenCaptureStarting")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(sender *ICoreWebView2, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr)(sender, args)
	}
	return ret
}

// BrowserExtensionRemoveCompleted 浏览器扩展移除完成时调用。
func (w *WebViewEventImpl) BrowserExtensionRemoveCompleted(errorCode syscall.Errno) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "BrowserExtensionRemoveCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno) uintptr)(errorCode)
	}
	return ret
}

// BrowserExtensionEnableCompleted 浏览器扩展启用完成时调用。
func (w *WebViewEventImpl) BrowserExtensionEnableCompleted(errorCode syscall.Errno) uintptr {
	cbs := WvEventHandler.GetCallBacks(w, "BrowserExtensionEnableCompleted")
	var ret uintptr
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(func(errorCode syscall.Errno) uintptr)(errorCode)
	}
	return ret
}
