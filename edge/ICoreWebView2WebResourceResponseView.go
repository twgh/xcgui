package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceResponseView 网页资源响应的 HTTP 表示形式视图。
//   - 此对象的属性不可变。
//   - 此响应视图与 WebResourceResponseReceived 事件一起使用。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview
type ICoreWebView2WebResourceResponseView struct {
	Vtbl *ICoreWebView2WebResourceResponseViewVtbl
}

type ICoreWebView2WebResourceResponseViewVtbl struct {
	IUnknownVtbl
	GetHeaders      ComProc
	GetStatusCode   ComProc
	GetReasonPhrase ComProc
	GetContent      ComProc
}

func (i *ICoreWebView2WebResourceResponseView) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseView) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseView) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHeaders 获取HTTP响应头。
func (i *ICoreWebView2WebResourceResponseView) GetHeaders() (*ICoreWebView2HttpResponseHeaders, error) {
	var headers *ICoreWebView2HttpResponseHeaders
	r, _, _ := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return headers, nil
}

// GetStatusCode 获取 HTTP 响应状态码。
func (i *ICoreWebView2WebResourceResponseView) GetStatusCode() (int, error) {
	var statusCode int
	r, _, _ := i.Vtbl.GetStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&statusCode)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return statusCode, nil
}

// GetReasonPhrase 获取 HTTP 响应原因短语。
func (i *ICoreWebView2WebResourceResponseView) GetReasonPhrase() (string, error) {
	var reasonPhrase *uint16
	r, _, _ := i.Vtbl.GetReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reasonPhrase)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	phrase := common.UTF16PtrToString(reasonPhrase)
	wapi.CoTaskMemFree(unsafe.Pointer(reasonPhrase))
	return phrase, nil
}

// GetContent 获取响应内容。处理程序将接收响应内容流。
//   - 如果内容大小超过 123MB，或者对于会变成下载的导航，又或者如果响应是可下载内容类型（例如，application/octet-stream），则此方法返回 null。请参阅 DownloadStarting 事件来处理该响应。
//   - 如果在首次调用完成之前再次调用此方法，将在调用先前调用的处理程序的同时调用该处理程序。
//   - 如果在首次调用完成之后调用此方法，该处理程序将立即被调用。
func (i *ICoreWebView2WebResourceResponseView) GetContent(handler *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) error {
	r, _, _ := i.Vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetContentEx 获取响应内容。处理程序将接收响应内容流。
//   - 如果内容大小超过 123MB，或者对于会变成下载的导航，又或者如果响应是可下载内容类型（例如，application/octet-stream），则此方法返回 null。请参阅 DownloadStarting 事件来处理该响应。
//   - 如果在首次调用完成之前再次调用此方法，将在调用先前调用的处理程序的同时调用该处理程序。
//   - 如果在首次调用完成之后调用此方法，该处理程序将立即被调用。
//
// impl: *WebViewEventImpl。
//
// cb: 执行完成后的回调处理程序，接收返回结果。
func (i *ICoreWebView2WebResourceResponseView) GetContentEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, content []byte) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "WebResourceResponseViewGetContentCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "WebResourceResponseViewGetContentCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "WebResourceResponseViewGetContentCompleted")
	}
	return i.GetContent((*ICoreWebView2WebResourceResponseViewGetContentCompletedHandler)(handler))
}

// MustGetHeaders 获取 HTTP 响应头，出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponseView) MustGetHeaders() *ICoreWebView2HttpResponseHeaders {
	headers, err := i.GetHeaders()
	ReportErrorAtuo(err)
	return headers
}

// MustGetStatusCode 获取 HTTP 响应状态码，出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponseView) MustGetStatusCode() int {
	statusCode, err := i.GetStatusCode()
	ReportErrorAtuo(err)
	return statusCode
}

// MustGetReasonPhrase 获取 HTTP 响应原因短语，出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponseView) MustGetReasonPhrase() string {
	phrase, err := i.GetReasonPhrase()
	ReportErrorAtuo(err)
	return phrase
}
