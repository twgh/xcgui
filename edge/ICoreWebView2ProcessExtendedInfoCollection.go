package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ProcessExtendedInfoCollection 是 ICoreWebView2ProcessExtendedInfo 的集合。。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection
type ICoreWebView2ProcessExtendedInfoCollection struct {
	Vtbl *ICoreWebView2ProcessExtendedInfoCollectionVtbl
}

type ICoreWebView2ProcessExtendedInfoCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的项目数。
func (i *ICoreWebView2ProcessExtendedInfoCollection) GetCount() (uint32, error) {
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

// GetValueAtIndex 获取指定索引处的进程扩展信息。
//
// index: 要获取的项目的从零开始的索引。
func (i *ICoreWebView2ProcessExtendedInfoCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ProcessExtendedInfo, error) {
	var value *ICoreWebView2ProcessExtendedInfo
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

// MustGetCount 获取集合中的项目数。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessExtendedInfoCollection) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAuto(err)
	return value
}

// MustGetValueAtIndex 获取指定索引处的进程扩展信息。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessExtendedInfoCollection) MustGetValueAtIndex(index uint32) *ICoreWebView2ProcessExtendedInfo {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return value
}
