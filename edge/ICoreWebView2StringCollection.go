package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2StringCollection 字符串集合.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection
type ICoreWebView2StringCollection struct {
	Vtbl *ICoreWebView2StringCollectionVtbl
}

type ICoreWebView2StringCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2StringCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2StringCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2StringCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的字符串数量.
func (i *ICoreWebView2StringCollection) GetCount() (uint32, error) {
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

// GetValueAtIndex 获取指定索引处的字符串值.
func (i *ICoreWebView2StringCollection) GetValueAtIndex(index uint32) (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(value))
	return common.UTF16PtrToString(value), nil
}

// MustGetCount 获取集合中的字符串数量. 出错时会触发全局错误回调.
func (i *ICoreWebView2StringCollection) MustGetCount() uint32 {
	count, err := i.GetCount()
	ReportErrorAuto(err)
	return count
}

// MustGetValueAtIndex 获取指定索引处的字符串值. 出错时会触发全局错误回调.
func (i *ICoreWebView2StringCollection) MustGetValueAtIndex(index uint32) string {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return value
}
