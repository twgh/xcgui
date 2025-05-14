package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2HttpHeadersCollectionIterator 提供了一个迭代器来枚举HTTP头集合。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator
type ICoreWebView2HttpHeadersCollectionIterator struct {
	Vtbl *ICoreWebView2HttpHeadersCollectionIteratorVtbl
}

type ICoreWebView2HttpHeadersCollectionIteratorVtbl struct {
	IUnknownVtbl
	GetCurrentHeader    ComProc
	MoveNext            ComProc
	GetHasCurrentHeader ComProc
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) Release() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCurrentHeader 获取当前HTTP头的名称和值。
func (i *ICoreWebView2HttpHeadersCollectionIterator) GetCurrentHeader() (name string, value string, err error) {
	var _name, _value *uint16
	r, _, err := i.Vtbl.GetCurrentHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_name)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", "", err
	}
	if r != 0 {
		return "", "", syscall.Errno(r)
	}
	name = windows.UTF16PtrToString(_name)
	value = windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_name))
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return name, value, nil
}

// MustGetCurrentHeader 获取当前HTTP头的名称和值。忽略错误。
func (i *ICoreWebView2HttpHeadersCollectionIterator) MustGetCurrentHeader() (name string, value string) {
	name, value, _ = i.GetCurrentHeader()
	return
}

// MoveNext 将迭代器移动到下一个HTTP头。
func (i *ICoreWebView2HttpHeadersCollectionIterator) MoveNext() (bool, error) {
	var hasNext int32
	r, _, err := i.Vtbl.MoveNext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasNext)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasNext != 0, nil
}

// MustMoveNext 将迭代器移动到下一个HTTP头。忽略错误。
func (i *ICoreWebView2HttpHeadersCollectionIterator) MustMoveNext() bool {
	hasNext, _ := i.MoveNext()
	return hasNext
}

// GetHasCurrentHeader 检查迭代器是否有当前HTTP头。
func (i *ICoreWebView2HttpHeadersCollectionIterator) GetHasCurrentHeader() (bool, error) {
	var hasCurrent int32
	r, _, err := i.Vtbl.GetHasCurrentHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasCurrent)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasCurrent != 0, nil
}

// MustGetHasCurrentHeader 检查迭代器是否有当前HTTP头。忽略错误。
func (i *ICoreWebView2HttpHeadersCollectionIterator) MustGetHasCurrentHeader() bool {
	hasCurrent, _ := i.GetHasCurrentHeader()
	return hasCurrent
}
