package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Deferral 表示一个延迟对象。此接口用于完成对支持使用 GetDeferral 方法获取延迟的事件参数的延迟。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2deferral
type ICoreWebView2Deferral struct {
	Vtbl *ICoreWebView2DeferralVtbl
}

type ICoreWebView2DeferralVtbl struct {
	IUnknownVtbl
	Complete ComProc
}

func (i *ICoreWebView2Deferral) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Deferral) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Deferral) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Complete 完成相关的延迟事件。每次延期只应运行一次完成。
func (i *ICoreWebView2Deferral) Complete() error {
	r, _, _ := i.Vtbl.Complete.Call(uintptr(unsafe.Pointer(i)))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
