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

// MustGetIsNewDocument 获取导航到的页面是否为一个新文档. 忽略错误.
func (i *ICoreWebView2SourceChangedEventArgs) MustGetIsNewDocument() bool {
	isNewDocument, _ := i.GetIsNewDocument()
	return isNewDocument
}
