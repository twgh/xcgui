package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ProcessInfo 为 ICoreWebView2Environment 中的进程提供一组属性。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processinfo
type ICoreWebView2ProcessInfo struct {
	Vtbl *ICoreWebView2ProcessInfoVtbl
}

type ICoreWebView2ProcessInfoVtbl struct {
	IUnknownVtbl
	GetProcessId ComProc
	GetKind      ComProc
}

func (i *ICoreWebView2ProcessInfo) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessInfo) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessInfo) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProcessId 获取进程ID。
func (i *ICoreWebView2ProcessInfo) GetProcessId() (int32, error) {
	var value int32
	r, _, _ := i.Vtbl.GetProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetKind 获取进程类型。
func (i *ICoreWebView2ProcessInfo) GetKind() (COREWEBVIEW2_PROCESS_KIND, error) {
	var value COREWEBVIEW2_PROCESS_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetProcessId 获取进程ID。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessInfo) MustGetProcessId() int32 {
	value, err := i.GetProcessId()
	ReportErrorAuto(err)
	return value
}

// MustGetKind 获取进程类型。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessInfo) MustGetKind() COREWEBVIEW2_PROCESS_KIND {
	value, err := i.GetKind()
	ReportErrorAuto(err)
	return value
}
