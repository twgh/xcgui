package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment2 是 ICoreWebView2Environment 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment2
type ICoreWebView2Environment2 struct {
	Vtbl *ICoreWebView2Environment2Vtbl
}

type ICoreWebView2Environment2Vtbl struct {
	ICoreWebView2EnvironmentVtbl
	CreateWebResourceRequest ComProc
}

func (i *ICoreWebView2Environment2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment2) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateWebResourceRequest 创建一个新的网络资源请求对象。
//   - URI 参数必须是绝对 URI。
//   - 标头字符串是由 CRLF 分隔的原始请求标头字符串（最后一个标头中可选）。也可以使用空标头字符串创建此对象，然后使用 ICoreWebView2HttpRequestHeaders 逐行构造标头。
func (i *ICoreWebView2Environment2) CreateWebResourceRequest(uri string, method string, postData *IStream, headers string) (*ICoreWebView2WebResourceRequest, error) {
	var request *ICoreWebView2WebResourceRequest
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return nil, err
	}
	_method, err := syscall.UTF16PtrFromString(method)
	if err != nil {
		return nil, err
	}
	_headers, err := syscall.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}
	r, _, err := i.Vtbl.CreateWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(unsafe.Pointer(_method)),
		uintptr(unsafe.Pointer(postData)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&request)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return request, nil
}
