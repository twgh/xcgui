package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceResponseReceivedEventArgs 是 Web 资源响应接收事件的参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs
type ICoreWebView2WebResourceResponseReceivedEventArgs struct {
	Vtbl *ICoreWebView2WebResourceResponseReceivedEventArgsVtbl
}

type ICoreWebView2WebResourceResponseReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetRequest  ComProc
	GetResponse ComProc
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetRequest 获取已提交的 Web 资源请求对象。
//   - 这包括网络堆栈添加的、在关联的 WebResourceRequested 事件期间未包含的标头，例如身份验证标头。
//   - 对此对象的修改对请求的处理方式没有影响，因为请求已被发送。
func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetRequest() (*ICoreWebView2WebResourceRequest, error) {
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

// GetResponse 获取收到的网络资源响应对象的视图。
func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetResponse() (*ICoreWebView2WebResourceResponseView, error) {
	var response *ICoreWebView2WebResourceResponseView
	r, _, _ := i.Vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// MustGetRequest 获取已提交的 Web 资源请求对象。出错时会触发全局错误回调。
//   - 这包括网络堆栈添加的、在关联的 WebResourceRequested 事件期间未包含的标头，例如身份验证标头。
//   - 对此对象的修改对请求的处理方式没有影响，因为请求已被发送。
func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) MustGetRequest() *ICoreWebView2WebResourceRequest {
	request, err := i.GetRequest()
	ReportErrorAuto(err)
	return request
}

// MustGetResponse 获取收到的网络资源响应对象的视图。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) MustGetResponse() *ICoreWebView2WebResourceResponseView {
	response, err := i.GetResponse()
	ReportErrorAuto(err)
	return response
}
