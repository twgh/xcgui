package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
)

// ICoreWebView2RegionRectCollectionView 是区域矩形的集合。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2regionrectcollectionview
type ICoreWebView2RegionRectCollectionView struct {
	Vtbl *ICoreWebView2RegionRectCollectionViewVtbl
}

type ICoreWebView2RegionRectCollectionViewVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2RegionRectCollectionView) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2RegionRectCollectionView) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2RegionRectCollectionView) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的区域矩形数量。
func (i *ICoreWebView2RegionRectCollectionView) GetCount() (uint32, error) {
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

// GetValueAtIndex 获取指定索引位置的区域矩形。
func (i *ICoreWebView2RegionRectCollectionView) GetValueAtIndex(index uint32) (*xc.RECT, error) {
	var value *xc.RECT
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

// MustGetCount 获取集合中的区域矩形数量。出错时会触发全局错误回调。
func (i *ICoreWebView2RegionRectCollectionView) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAuto(err)
	return value
}

// MustGetValueAtIndex 获取指定索引位置的区域矩形。出错时会触发全局错误回调。
func (i *ICoreWebView2RegionRectCollectionView) MustGetValueAtIndex(index uint32) *xc.RECT {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return value
}
