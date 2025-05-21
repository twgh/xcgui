package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
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

func (i *ICoreWebView2HttpRequestHeaders) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHeader 获取指定名称的 HTTP 头的值。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) GetHeader(name string) (string, error) {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}

	var _value *uint16
	r, _, err := i.Vtbl.GetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// MustGetHeader 获取指定名称的HTTP头的值。忽略错误。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) MustGetHeader(name string) string {
	value, _ := i.GetHeader(name)
	return value
}

// Contains 检查是否存在指定名称的HTTP头。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) Contains(name string) (bool, error) {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}

	var contains int32
	r, _, err := i.Vtbl.Contains.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&contains)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return contains != 0, nil
}

// MustContains 检查是否存在指定名称的HTTP头。忽略错误。
func (i *ICoreWebView2HttpRequestHeaders) MustContains(name string) bool {
	contains, _ := i.Contains(name)
	return contains
}

// RemoveHeader 删除指定名称的HTTP头。
func (i *ICoreWebView2HttpRequestHeaders) RemoveHeader(name string) error {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.RemoveHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
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
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.SetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIterator 获取一个迭代器来枚举所有HTTP头。
func (i *ICoreWebView2HttpRequestHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, err := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// MustGetIterator 获取一个迭代器来枚举所有HTTP头。忽略错误。
func (i *ICoreWebView2HttpRequestHeaders) MustGetIterator() *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, _ := i.GetIterator()
	return iterator
}

// GetHeaders 获取指定名称的所有HTTP头的值。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// MustGetHeaders 获取指定名称的所有HTTP头的值。忽略错误。
//
// name: 标头名称。
func (i *ICoreWebView2HttpRequestHeaders) MustGetHeaders(name string) *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, _ := i.GetHeaders(name)
	return iterator
}
