package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_17 是 ICoreWebView2_16 接口的延续，支持基于文件映射的共享缓冲区。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_17
type ICoreWebView2_17 struct {
	Vtbl *ICoreWebView2_17Vtbl
}

type ICoreWebView2_17Vtbl struct {
	ICoreWebView2_16Vtbl
	PostSharedBufferToScript ComProc
}

func (i *ICoreWebView2_17) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_17) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_17) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PostSharedBufferToScript 在 WebView 中与主框架的脚本共享一个共享缓冲区对象。
func (i *ICoreWebView2_17) PostSharedBufferToScript(sharedBuffer *ICoreWebView2SharedBuffer, access COREWEBVIEW2_SHARED_BUFFER_ACCESS, additionalDataAsJson string) error {
	_additionalDataAsJson, err := syscall.UTF16PtrFromString(additionalDataAsJson)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostSharedBufferToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(sharedBuffer)),
		uintptr(access),
		uintptr(unsafe.Pointer(_additionalDataAsJson)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustPostSharedBufferToScript 在 WebView 中与主框架的脚本共享一个共享缓冲区对象。出错时会触发全局错误回调。
func (i *ICoreWebView2_17) MustPostSharedBufferToScript(sharedBuffer *ICoreWebView2SharedBuffer, access COREWEBVIEW2_SHARED_BUFFER_ACCESS, additionalDataAsJson string) {
	err := i.PostSharedBufferToScript(sharedBuffer, access, additionalDataAsJson)
	ReportErrorAtuo(err)
}
