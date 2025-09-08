package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ProcessInfoCollection 是 ICoreWebView2ProcessInfo 的集合。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection
type ICoreWebView2ProcessInfoCollection struct {
	Vtbl *ICoreWebView2ProcessInfoCollectionVtbl
}

type ICoreWebView2ProcessInfoCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2ProcessInfoCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessInfoCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessInfoCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的进程信息数量。
func (i *ICoreWebView2ProcessInfoCollection) GetCount() (uint32, error) {
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

// GetValueAtIndex 获取指定索引处的进程信息。
func (i *ICoreWebView2ProcessInfoCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ProcessInfo, error) {
	var value *ICoreWebView2ProcessInfo
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

// MustGetCount 获取集合中的进程信息数量。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessInfoCollection) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAuto(err)
	return value
}

// MustGetValueAtIndex 获取指定索引处的进程信息。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessInfoCollection) MustGetValueAtIndex(index uint32) *ICoreWebView2ProcessInfo {
	value, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return value
}
