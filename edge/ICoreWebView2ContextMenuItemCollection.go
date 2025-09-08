package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ContextMenuItemCollection 上下文菜单项的集合.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection
type ICoreWebView2ContextMenuItemCollection struct {
	Vtbl *ICoreWebView2ContextMenuItemCollectionVtbl
}

type ICoreWebView2ContextMenuItemCollectionVtbl struct {
	IUnknownVtbl
	GetCount           ComProc
	GetValueAtIndex    ComProc
	RemoveValueAtIndex ComProc
	InsertValueAtIndex ComProc
}

func (i *ICoreWebView2ContextMenuItemCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuItemCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuItemCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取菜单项数量.
func (i *ICoreWebView2ContextMenuItemCollection) GetCount() (uint32, error) {
	var count uint32
	r, _, _ := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return count, nil
}

// GetValueAtIndex 获取指定索引处的菜单项.
func (i *ICoreWebView2ContextMenuItemCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ContextMenuItem, error) {
	var item *ICoreWebView2ContextMenuItem
	r, _, _ := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&item)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return item, nil
}

// RemoveValueAtIndex 移除指定索引处的菜单项.
func (i *ICoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) error {
	r, _, _ := i.Vtbl.RemoveValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// InsertValueAtIndex 在指定索引处插入菜单项.
func (i *ICoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, item *ICoreWebView2ContextMenuItem) error {
	r, _, _ := i.Vtbl.InsertValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(item)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetCount 获取菜单项数量。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItemCollection) MustGetCount() uint32 {
	count, err := i.GetCount()
	ReportErrorAuto(err)
	return count
}

// MustGetValueAtIndex 获取指定索引处的菜单项。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItemCollection) MustGetValueAtIndex(index uint32) *ICoreWebView2ContextMenuItem {
	item, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return item
}
