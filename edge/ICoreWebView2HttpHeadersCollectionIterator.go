package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

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
	GetHasCurrentHeader ComProc
	MoveNext            ComProc
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
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
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", "", err
	}
	if r != 0 {
		return "", "", syscall.Errno(r)
	}
	name = common.UTF16PtrToString(_name)
	value = common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_name))
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return name, value, nil
}

// MustGetCurrentHeader 获取当前HTTP头的名称和值。出错时会触发全局错误回调。
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
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasNext != 0, nil
}

// MustMoveNext 将迭代器移动到下一个HTTP头。出错时会触发全局错误回调。
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
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasCurrent != 0, nil
}

// MustGetHasCurrentHeader 检查迭代器是否有当前HTTP头。出错时会触发全局错误回调。
func (i *ICoreWebView2HttpHeadersCollectionIterator) MustGetHasCurrentHeader() bool {
	hasCurrent, err := i.GetHasCurrentHeader()
	ReportErrorAtuo(err)
	return hasCurrent
}
