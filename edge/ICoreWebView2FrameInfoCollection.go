package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2FrameInfoCollection 是 FrameInfo （名称和来源）的集合。用于在 ICoreWebView2 中仅渲染框架的进程失败时列出受影响的框架信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection
type ICoreWebView2FrameInfoCollection struct {
	Vtbl *ICoreWebView2FrameInfoCollectionVtbl
}

type ICoreWebView2FrameInfoCollectionVtbl struct {
	IUnknownVtbl
	GetIterator ComProc
}

func (i *ICoreWebView2FrameInfoCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfoCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfoCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIterator 获取框架信息集合的迭代器。
func (i *ICoreWebView2FrameInfoCollection) GetIterator() (*ICoreWebView2FrameInfoCollectionIterator, error) {
	var value *ICoreWebView2FrameInfoCollectionIterator
	r, _, _ := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetIterator 获取框架信息集合的迭代器。出错时会触发全局错误回调。
func (i *ICoreWebView2FrameInfoCollection) MustGetIterator() *ICoreWebView2FrameInfoCollectionIterator {
	value, err := i.GetIterator()
	ReportErrorAuto(err)
	return value
}
