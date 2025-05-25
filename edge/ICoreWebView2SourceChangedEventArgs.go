package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
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

func (i *ICoreWebView2SourceChangedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsNewDocument 获取导航到的页面是否为一个新文档.
func (i *ICoreWebView2SourceChangedEventArgs) GetIsNewDocument() (bool, error) {
	var isNewDocument bool
	r, _, err := i.Vtbl.GetIsNewDocument.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isNewDocument)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isNewDocument, nil
}

// MustGetIsNewDocument 获取导航到的页面是否为一个新文档. 出错时会触发全局错误回调.
func (i *ICoreWebView2SourceChangedEventArgs) MustGetIsNewDocument() bool {
	isNewDocument, err := i.GetIsNewDocument()
	ReportError2(err)
	return isNewDocument
}
