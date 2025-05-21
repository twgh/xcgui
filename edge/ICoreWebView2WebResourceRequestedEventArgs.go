package edge

// ok

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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

func (i *ICoreWebView2WebResourceRequestedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutResponse 设置响应。
//
// response: 使用 ICoreWebView2Environment.CreateWebResourceResponse 创建一个空的 web 资源响应对象，然后修改它以构造响应。
func (i *ICoreWebView2WebResourceRequestedEventArgs) PutResponse(response *ICoreWebView2WebResourceResponse) error {
	r, _, err := i.Vtbl.PutResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(response)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetRequest 获取请求。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetRequest() (*ICoreWebView2WebResourceRequest, error) {
	var request *ICoreWebView2WebResourceRequest
	r, _, err := i.Vtbl.GetRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&request)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return request, nil
}

// MustGetRequest 获取请求。忽略错误。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetRequest() *ICoreWebView2WebResourceRequest {
	request, _ := i.GetRequest()
	return request
}

// GetResponse 获取响应。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetResponse() (*ICoreWebView2WebResourceResponse, error) {
	var response *ICoreWebView2WebResourceResponse
	r, _, err := i.Vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// MustGetResponse 获取响应。忽略错误。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetResponse() *ICoreWebView2WebResourceResponse {
	response, _ := i.GetResponse()
	return response
}

// GetDeferral 获取 ICoreWebView2Deferral 对象，并将事件置于延迟状态。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// MustGetDeferral 获取 ICoreWebView2Deferral 对象，并将事件置于延迟状态。忽略错误。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, _ := i.GetDeferral()
	return deferral
}

// GetResourceContext 获取资源上下文。
func (i *ICoreWebView2WebResourceRequestedEventArgs) GetResourceContext() (COREWEBVIEW2_WEB_RESOURCE_CONTEXT, error) {
	var context COREWEBVIEW2_WEB_RESOURCE_CONTEXT
	r, _, err := i.Vtbl.GetResourceContext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&context)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return context, nil
}

// MustGetResourceContext 获取资源上下文。忽略错误。
func (i *ICoreWebView2WebResourceRequestedEventArgs) MustGetResourceContext() COREWEBVIEW2_WEB_RESOURCE_CONTEXT {
	context, _ := i.GetResourceContext()
	return context
}
