package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2HttpRequestHeaders 提供了一个迭代器来列出HTTP请求头，以及用于检查和操作HTTP请求头的方法。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders
type ICoreWebView2HttpRequestHeaders struct {
	Vtbl *_ICoreWebView2HttpRequestHeadersVtbl
}

type _ICoreWebView2HttpRequestHeadersVtbl struct {
	IUnknownVtbl
	GetHeader    ComProc
	GetHeaders   ComProc
	Contains     ComProc
	SetHeader    ComProc
	RemoveHeader ComProc
	GetIterator  ComProc
}

func (i *ICoreWebView2HttpRequestHeaders) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpRequestHeaders) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpRequestHeaders) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHeader 获取指定名称的 HTTP 头的值。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) GetHeader(name string) (string, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}

	var _value *uint16
	r, _, _ := i.Vtbl.GetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// Contains 检查是否存在指定名称的HTTP头。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) Contains(name string) (bool, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}

	var contains int32
	r, _, _ := i.Vtbl.Contains.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&contains)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return contains != 0, nil
}

// RemoveHeader 删除指定名称的HTTP头。
func (i *ICoreWebView2HttpRequestHeaders) RemoveHeader(name string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.RemoveHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetHeader 设置指定名称的HTTP头的值。
//
// name: 标头名称。
//
// value: 标头值。
func (i *ICoreWebView2HttpRequestHeaders) SetHeader(name string, value string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	r, _, _ := i.Vtbl.SetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIterator 获取一个迭代器来枚举所有HTTP头。
func (i *ICoreWebView2HttpRequestHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, _ := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// GetHeaders 获取指定名称的所有HTTP头的值。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, _ := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// MustGetHeader 获取指定名称的HTTP头的值。出错时会触发全局错误回调。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) MustGetHeader(name string) string {
	value, err := i.GetHeader(name)
	ReportErrorAtuo(err)
	return value
}

// MustContains 检查是否存在指定名称的HTTP头。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpRequestHeaders) MustContains(name string) bool {
	contains, _ := i.Contains(name)
	return contains
}

// MustGetIterator 获取一个迭代器来枚举所有HTTP头。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpRequestHeaders) MustGetIterator() *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, err := i.GetIterator()
	ReportErrorAtuo(err)
	return iterator
}

// MustGetHeaders 获取指定名称的所有HTTP头的值。出错时会触发全局错误回调。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) MustGetHeaders(name string) *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, err := i.GetHeaders(name)
	ReportErrorAtuo(err)
	return iterator
}
