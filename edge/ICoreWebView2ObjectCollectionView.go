package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ObjectCollectionView 提供对对象集合的只读访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview
type ICoreWebView2ObjectCollectionView struct {
	Vtbl *ICoreWebView2ObjectCollectionViewVtbl
}

type ICoreWebView2ObjectCollectionViewVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2ObjectCollectionView) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ObjectCollectionView) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ObjectCollectionView) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的对象数量。
func (i *ICoreWebView2ObjectCollectionView) GetCount() (uint32, error) {
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
func (i *ICoreWebView2ObjectCollectionView) GetValueAtIndex(index uint32) (*IUnknown, error) {
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

// MustGetCount 获取集合中的对象数量。出错时会触发全局错误回调。
func (i *ICoreWebView2ObjectCollectionView) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAtuo(err)
	return value
}

// MustGetValueAtIndex 获取指定索引位置的对象。出错时会触发全局错误回调。
func (i *ICoreWebView2ObjectCollectionView) MustGetValueAtIndex(index uint32) *IUnknown {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAtuo(err)
	return value
}
