package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceRequest 是与 WebResourceRequested 事件一起使用的HTTP请求。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest
type ICoreWebView2WebResourceRequest struct {
	Vtbl *ICoreWebView2WebResourceRequestVtbl
}

type ICoreWebView2WebResourceRequestVtbl struct {
	IUnknownVtbl
	GetUri     ComProc
	PutUri     ComProc
	GetMethod  ComProc
	PutMethod  ComProc
	GetContent ComProc
	PutContent ComProc
	GetHeaders ComProc
}

func (i *ICoreWebView2WebResourceRequest) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequest) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequest) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取请求URI。
func (i *ICoreWebView2WebResourceRequest) GetUri() (string, error) {
	var _uri *uint16
	r, _, err := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(_uri)
	wapi.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// GetContent 获取请求的内容。
func (i *ICoreWebView2WebResourceRequest) GetContent() ([]byte, error) {
	var streamPtr uintptr
	hr, _, err := i.Vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&streamPtr)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	if streamPtr == 0 {
		return nil, nil
	}
	stream := NewIStreamByPtr(streamPtr)
	content, err := stream.GetBytesAndRelease()
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetMethod 获取请求的HTTP方法。
func (i *ICoreWebView2WebResourceRequest) GetMethod() (string, error) {
	var _method *uint16
	r, _, err := i.Vtbl.GetMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_method)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	method := common.UTF16PtrToString(_method)
	wapi.CoTaskMemFree(unsafe.Pointer(_method))
	return method, nil
}

// GetHeaders 获取请求的HTTP标头。
func (i *ICoreWebView2WebResourceRequest) GetHeaders() (*ICoreWebView2HttpRequestHeaders, error) {
	var headers *ICoreWebView2HttpRequestHeaders
	r, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return headers, nil
}

// SetUri 设置请求URI。
func (i *ICoreWebView2WebResourceRequest) SetUri(uri string) error {
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetContent 设置请求的内容。
func (i *ICoreWebView2WebResourceRequest) SetContent(content *IStream) error {
	r, _, err := i.Vtbl.PutContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(content)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetMethod 设置请求的HTTP方法。
func (i *ICoreWebView2WebResourceRequest) SetMethod(method string) error {
	_method, err := syscall.UTF16PtrFromString(method)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_method)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHeadersMap 获取请求的HTTP标头并转换为map。
func (i *ICoreWebView2WebResourceRequest) GetHeadersMap() (map[string]string, error) {
	headers, err := i.GetHeaders()
	if err != nil {
		return nil, err
	}

	iterator, err := headers.GetIterator()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for {
		hasCurrent, err := iterator.GetHasCurrentHeader()
		if err != nil || !hasCurrent {
			break
		}

		name, value, err := iterator.GetCurrentHeader()
		if err != nil {
			return nil, err
		}
		result[name] = value

		hasNext, err := iterator.MoveNext()
		if err != nil {
			return nil, err
		}
		if !hasNext {
			break
		}
	}

	return result, nil
}

// MustGetUri 获取请求URI。出错时会触发全局错误回调.
func (i *ICoreWebView2WebResourceRequest) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetContent 获取请求的内容。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequest) MustGetContent() []byte {
	content, err := i.GetContent()
	ReportErrorAtuo(err)
	return content
}

// MustGetMethod 获取请求的HTTP方法。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequest) MustGetMethod() string {
	method, err := i.GetMethod()
	ReportErrorAtuo(err)
	return method
}

// MustGetHeaders 获取请求的HTTP标头。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequest) MustGetHeaders() *ICoreWebView2HttpRequestHeaders {
	headers, err := i.GetHeaders()
	ReportErrorAtuo(err)
	return headers
}

// MustGetHeadersMap 获取请求的HTTP标头并转换为map。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequest) MustGetHeadersMap() map[string]string {
	headers, err := i.GetHeadersMap()
	ReportErrorAtuo(err)
	return headers
}
