package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2HttpResponseHeaders 是 HTTP 响应标头。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders
type ICoreWebView2HttpResponseHeaders struct {
	Vtbl *_ICoreWebView2HttpResponseHeadersVtbl
}

type _ICoreWebView2HttpResponseHeadersVtbl struct {
	IUnknownVtbl
	AppendHeader ComProc
	Contains     ComProc
	GetHeader    ComProc
	GetHeaders   ComProc
	GetIterator  ComProc
}

func (i *ICoreWebView2HttpResponseHeaders) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpResponseHeaders) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpResponseHeaders) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AppendHeader 用名称和值附加标题行。
//
// name: 标头名称。
//
// value: 标头值。
func (i *ICoreWebView2HttpResponseHeaders) AppendHeader(name string, value string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.AppendHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Contains 检查是否存在具有给定名称的标头。
//
// name: 要检查的标头名称。
func (i *ICoreWebView2HttpResponseHeaders) Contains(name string) (bool, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}
	var contains bool
	r, _, err := i.Vtbl.Contains.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&contains)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return contains, nil
}

// MustContains 检查是否存在具有给定名称的标头。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpResponseHeaders) MustContains(name string) bool {
	contains, _ := i.Contains(name)
	return contains
}

// GetHeader 获取集合中与该名称匹配的第一个标头值。
//
// name: 要获取的标头名称。
func (i *ICoreWebView2HttpResponseHeaders) GetHeader(name string) (string, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}
	var _value *uint16
	r, _, err := i.Vtbl.GetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// MustGetHeader 获取具有给定名称的标头值。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpResponseHeaders) MustGetHeader(name string) string {
	value, err := i.GetHeader(name)
	ReportErrorAtuo(err)
	return value
}

// GetHeaders 获取具有给定名称的所有标头值。
//
// name: 要获取的标头名称。
func (i *ICoreWebView2HttpResponseHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// MustGetHeaders 获取具有给定名称的所有标头值。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpResponseHeaders) MustGetHeaders(name string) *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, err := i.GetHeaders(name)
	ReportErrorAtuo(err)
	return iterator
}

// GetIterator 获取集合中所有标头的迭代器。
func (i *ICoreWebView2HttpResponseHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	var iterator *ICoreWebView2HttpHeadersCollectionIterator
	r, _, err := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return iterator, nil
}

// MustGetIterator 获取集合中所有标头的迭代器。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpResponseHeaders) MustGetIterator() *ICoreWebView2HttpHeadersCollectionIterator {
	iterator, err := i.GetIterator()
	ReportErrorAtuo(err)
	return iterator
}
