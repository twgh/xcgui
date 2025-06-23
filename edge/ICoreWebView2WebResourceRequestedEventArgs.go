package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceRequestedEventArgs 是 WebResourceRequested 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs
type ICoreWebView2WebResourceRequestedEventArgs struct {
	Vtbl *ICoreWebView2WebResourceRequestedEventArgsVtbl
}

type ICoreWebView2WebResourceRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetRequest         ComProc
	GetResponse        ComProc
	PutResponse        ComProc
	GetDeferral        ComProc
	GetResourceContext ComProc
}

func (i *ICoreWebView2WebResourceRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequestedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetResponse 设置响应。
//
// response: 使用 ICoreWebView2Environment.CreateWebResourceResponse 创建一个空的 web 资源响应对象，然后修改它以构造响应。
func (i *ICoreWebView2WebResourceRequestedEventArgs) SetResponse(response *ICoreWebView2WebResourceResponse) error {
	r, _, _ := i.Vtbl.PutResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(response)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetRequest 获取请求。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetRequest() (*ICoreWebView2WebResourceRequest, error) {
	var request *ICoreWebView2WebResourceRequest
	r, _, _ := i.Vtbl.GetRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&request)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return request, nil
}

// GetResponse 获取响应。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetResponse() (*ICoreWebView2WebResourceResponse, error) {
	var response *ICoreWebView2WebResourceResponse
	r, _, _ := i.Vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// GetDeferral 获取 ICoreWebView2Deferral 对象，并将事件置于延迟状态。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// GetResourceContext 获取资源上下文。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetResourceContext() (COREWEBVIEW2_WEB_RESOURCE_CONTEXT, error) {
	var context COREWEBVIEW2_WEB_RESOURCE_CONTEXT
	r, _, _ := i.Vtbl.GetResourceContext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&context)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return context, nil
}

// MustGetRequest 获取请求。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetRequest() *ICoreWebView2WebResourceRequest {
	request, err := i.GetRequest()
	ReportErrorAtuo(err)
	return request
}

// MustGetResponse 获取响应。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetResponse() *ICoreWebView2WebResourceResponse {
	response, err := i.GetResponse()
	ReportErrorAtuo(err)
	return response
}

// MustGetDeferral 获取 ICoreWebView2Deferral 对象，并将事件置于延迟状态。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}

// MustGetResourceContext 获取资源上下文。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetResourceContext() COREWEBVIEW2_WEB_RESOURCE_CONTEXT {
	context, err := i.GetResourceContext()
	ReportErrorAtuo(err)
	return context
}
