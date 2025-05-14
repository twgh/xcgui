package edge

// ok

import (
	"errors"
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequest) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
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
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// MustGetUri 获取请求URI。忽略错误.
func (i *ICoreWebView2WebResourceRequest) MustGetUri() string {
	uri, _ := i.GetUri()
	return uri
}

// GetContent 获取请求的内容。
func (i *ICoreWebView2WebResourceRequest) GetContent() ([]byte, error) {
	var streamPtr uintptr
	hr, _, err := i.Vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&streamPtr)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	if streamPtr == 0 {
		return nil, nil
	}
	stream := (*IStream)(unsafe.Pointer(streamPtr))
	defer stream.Release()

	const bufferSize = 4096
	var content []byte

	for {
		buffer := make([]byte, bufferSize)
		n, hr := stream.Read(buffer)
		if hr != nil && !errors.Is(hr, wapi.S_FALSE) {
			return nil, fmt.Errorf("stream read failed: 0x%08X", hr)
		}
		if n == 0 {
			break
		}
		content = append(content, buffer[:n]...)
	}

	return content, nil
}

// MustGetContent 获取请求的内容。忽略错误。
func (i *ICoreWebView2WebResourceRequest) MustGetContent() []byte {
	content, _ := i.GetContent()
	return content
}

// GetMethod 获取请求的HTTP方法。
func (i *ICoreWebView2WebResourceRequest) GetMethod() (string, error) {
	var _method *uint16
	r, _, err := i.Vtbl.GetMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_method)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	method := windows.UTF16PtrToString(_method)
	windows.CoTaskMemFree(unsafe.Pointer(_method))
	return method, nil
}

// MustGetMethod 获取请求的HTTP方法。忽略错误。
func (i *ICoreWebView2WebResourceRequest) MustGetMethod() string {
	method, _ := i.GetMethod()
	return method
}

// GetHeaders 获取请求的HTTP标头。
func (i *ICoreWebView2WebResourceRequest) GetHeaders() (*ICoreWebView2HttpRequestHeaders, error) {
	var headers *ICoreWebView2HttpRequestHeaders
	r, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return headers, nil
}

// MustGetHeaders 获取请求的HTTP标头。忽略错误。
func (i *ICoreWebView2WebResourceRequest) MustGetHeaders() *ICoreWebView2HttpRequestHeaders {
	headers, _ := i.GetHeaders()
	return headers
}

// PutUri 设置请求URI。
func (i *ICoreWebView2WebResourceRequest) PutUri(uri string) error {
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutContent 设置请求的内容。
func (i *ICoreWebView2WebResourceRequest) PutContent(content []byte) error {
	var err error
	var stream uintptr
	// 创建内存流
	if len(content) > 0 {
		stream, err = wapi.SHCreateMemStream(content)
		if err != nil {
			return err
		}
	}
	r, _, err := i.Vtbl.PutContent.Call(
		uintptr(unsafe.Pointer(i)),
		stream,
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutMethod 设置请求的HTTP方法。
func (i *ICoreWebView2WebResourceRequest) PutMethod(method string) error {
	_method, err := windows.UTF16PtrFromString(method)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_method)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
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

// MustGetHeadersMap 获取请求的HTTP标头并转换为map。忽略错误。
func (i *ICoreWebView2WebResourceRequest) MustGetHeadersMap() map[string]string {
	headers, _ := i.GetHeadersMap()
	return headers
}
