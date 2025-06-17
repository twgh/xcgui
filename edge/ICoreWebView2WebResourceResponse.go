package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceResponse 是与 WebResourceRequested 事件一起使用的 HTTP 响应。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/ICoreWebView2WebResourceResponse
type ICoreWebView2WebResourceResponse struct {
	Vtbl *ICoreWebView2WebResourceResponseVtbl
}

type ICoreWebView2WebResourceResponseVtbl struct {
	IUnknownVtbl
	GetContent      ComProc
	PutContent      ComProc
	GetHeaders      ComProc
	GetStatusCode   ComProc
	PutStatusCode   ComProc
	GetReasonPhrase ComProc
	PutReasonPhrase ComProc
}

func (i *ICoreWebView2WebResourceResponse) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponse) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponse) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetContent 获取请求的内容。
func (i *ICoreWebView2WebResourceResponse) GetContent() ([]byte, error) {
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

// SetContent 设置请求的内容。
func (i *ICoreWebView2WebResourceResponse) SetContent(content *IStream) error {
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

// SetStatusCode 设置响应的 HTTP 状态码。
//
// statusCode: 要设置的 HTTP 状态码。
func (i *ICoreWebView2WebResourceResponse) SetStatusCode(statusCode int32) error {
	r, _, err := i.Vtbl.PutStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(statusCode),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetStatusCode 获取响应的 HTTP 状态码。
func (i *ICoreWebView2WebResourceResponse) GetStatusCode() (int32, error) {
	var statusCode int32
	r, _, err := i.Vtbl.GetStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&statusCode)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return statusCode, nil
}

// SetReasonPhrase 设置响应的 HTTP 状态码描述。
//
// reasonPhrase: 要设置的 HTTP 状态码描述。
func (i *ICoreWebView2WebResourceResponse) SetReasonPhrase(reasonPhrase string) error {
	ptr, err := syscall.UTF16PtrFromString(reasonPhrase)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(ptr)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetReasonPhrase 获取响应的 HTTP 状态码描述。
func (i *ICoreWebView2WebResourceResponse) GetReasonPhrase() (string, error) {
	var _reasonPhrase *uint16
	r, _, err := i.Vtbl.GetReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_reasonPhrase)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	reasonPhrase := common.UTF16PtrToString(_reasonPhrase)
	wapi.CoTaskMemFree(unsafe.Pointer(_reasonPhrase))
	return reasonPhrase, nil
}

// GetHeaders 获取响应的 HTTP 标头。
func (i *ICoreWebView2WebResourceResponse) GetHeaders() (*ICoreWebView2HttpResponseHeaders, error) {
	var headers *ICoreWebView2HttpResponseHeaders
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

// GetHeadersMap 获取请求的HTTP标头并转换为map。
func (i *ICoreWebView2WebResourceResponse) GetHeadersMap() (map[string]string, error) {
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

// MustGetStatusCode 获取响应的 HTTP 状态码。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponse) MustGetStatusCode() int32 {
	statusCode, err := i.GetStatusCode()
	ReportErrorAtuo(err)
	return statusCode
}

// MustGetReasonPhrase 获取响应的 HTTP 状态码描述。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponse) MustGetReasonPhrase() string {
	reasonPhrase, err := i.GetReasonPhrase()
	ReportErrorAtuo(err)
	return reasonPhrase
}

// MustGetHeaders 获取响应的 HTTP 标头。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponse) MustGetHeaders() *ICoreWebView2HttpResponseHeaders {
	headers, err := i.GetHeaders()
	ReportErrorAtuo(err)
	return headers
}

// MustGetHeadersMap 获取请求的HTTP标头并转换为map。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceResponse) MustGetHeadersMap() map[string]string {
	headers, err := i.GetHeadersMap()
	ReportErrorAtuo(err)
	return headers
}
