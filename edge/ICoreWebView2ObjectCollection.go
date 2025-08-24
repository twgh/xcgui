package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ObjectCollection 提供对对象集合的访问和修改功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2objectcollection
type ICoreWebView2ObjectCollection struct {
	Vtbl *ICoreWebView2ObjectCollectionVtbl
}

type ICoreWebView2ObjectCollectionVtbl struct {
	IUnknownVtbl
	GetCount           ComProc
	GetValueAtIndex    ComProc
	RemoveValueAtIndex ComProc
	InsertValueAtIndex ComProc
}

func (i *ICoreWebView2ObjectCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ObjectCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ObjectCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的对象数量。
func (i *ICoreWebView2ObjectCollection) GetCount() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetValueAtIndex 获取指定索引位置的对象。
func (i *ICoreWebView2ObjectCollection) GetValueAtIndex(index uint32) (*IUnknown, error) {
	var value *IUnknown
	r, _, _ := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// RemoveValueAtIndex 移除指定索引位置的对象。
func (i *ICoreWebView2ObjectCollection) RemoveValueAtIndex(index uint32) error {
	r, _, _ := i.Vtbl.RemoveValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// InsertValueAtIndex 在指定索引位置插入对象。
func (i *ICoreWebView2ObjectCollection) InsertValueAtIndex(index uint32, value *IUnknown) error {
	r, _, _ := i.Vtbl.InsertValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetCount 获取集合中的对象数量。出错时会触发全局错误回调。
func (i *ICoreWebView2ObjectCollection) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAtuo(err)
	return value
}

// MustGetValueAtIndex 获取指定索引位置的对象。出错时会触发全局错误回调。
func (i *ICoreWebView2ObjectCollection) MustGetValueAtIndex(index uint32) *IUnknown {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAtuo(err)
	return value
}
