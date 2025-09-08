package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2SourceChangedEventArgs 表示源改变事件的事件参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs
type ICoreWebView2SourceChangedEventArgs struct {
	Vtbl *ICoreWebView2SourceChangedEventArgsVtbl
}

type ICoreWebView2SourceChangedEventArgsVtbl struct {
	IUnknownVtbl
	GetIsNewDocument ComProc
}

func (i *ICoreWebView2SourceChangedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SourceChangedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SourceChangedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsNewDocument 获取导航到的页面是否为一个新文档.
func (i *ICoreWebView2SourceChangedEventArgs) GetIsNewDocument() (bool, error) {
	var isNewDocument bool
	r, _, _ := i.Vtbl.GetIsNewDocument.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isNewDocument)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isNewDocument, nil
}

// MustGetIsNewDocument 获取导航到的页面是否为一个新文档. 出错时会触发全局错误回调.
func (i *ICoreWebView2SourceChangedEventArgs) MustGetIsNewDocument() bool {
	isNewDocument, err := i.GetIsNewDocument()
	ReportErrorAuto(err)
	return isNewDocument
}
