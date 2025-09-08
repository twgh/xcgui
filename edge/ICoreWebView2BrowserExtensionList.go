package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2BrowserExtensionList 浏览器扩展列表.
type ICoreWebView2BrowserExtensionList struct {
	Vtbl *ICoreWebView2BrowserExtensionListVtbl
}

type ICoreWebView2BrowserExtensionListVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2BrowserExtensionList) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtensionList) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtensionList) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取列表中包含的浏览器扩展数量。
func (i *ICoreWebView2BrowserExtensionList) GetCount() (uint32, error) {
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

// GetValueAtIndex 获取指定索引处的浏览器扩展。
func (i *ICoreWebView2BrowserExtensionList) GetValueAtIndex(index uint32) (*ICoreWebView2BrowserExtension, error) {
	var value *ICoreWebView2BrowserExtension
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

// MustGetCount 获取列表中包含的浏览器扩展数量。出错时会触发全局错误回调。
func (i *ICoreWebView2BrowserExtensionList) MustGetCount() uint32 {
	result, err := i.GetCount()
	ReportErrorAuto(err)
	return result
}

// MustGetValueAtIndex 获取指定索引处的浏览器扩展。出错时会触发全局错误回调。
func (i *ICoreWebView2BrowserExtensionList) MustGetValueAtIndex(index uint32) *ICoreWebView2BrowserExtension {
	result, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return result
}
