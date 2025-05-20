package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_22 是 ICoreWebView2_21 接口的延续，允许设置筛选器，以便接收服务工作线程、共享工作线程和不同源 iframe 的 已请求web资源 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_22
type ICoreWebView2_22 struct {
	Vtbl *ICoreWebView2_22Vtbl
}

type ICoreWebView2_22Vtbl struct {
	ICoreWebView2_21Vtbl
	AddWebResourceRequestedFilterWithRequestSourceKinds    ComProc
	RemoveWebResourceRequestedFilterWithRequestSourceKinds ComProc
}

// AddWebResourceRequestedFilterWithRequestSourceKinds 添加带有请求源类型的 Web 资源请求过滤器.
//
// uri: 要过滤的 URI.
//
// ResourceContext: 资源上下文类型.
//
// requestSourceKinds: 请求源类型.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_22
func (i *ICoreWebView2_22) AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, ResourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT, requestSourceKinds COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS) error {
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.AddWebResourceRequestedFilterWithRequestSourceKinds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(ResourceContext),
		uintptr(requestSourceKinds),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveWebResourceRequestedFilterWithRequestSourceKinds 移除带有请求源类型的 Web 资源请求过滤器.
//
// uri: 要移除的 URI.
//
// ResourceContext: 资源上下文类型.
//
// requestSourceKinds: 请求源类型.
func (i *ICoreWebView2_22) RemoveWebResourceRequestedFilterWithRequestSourceKinds(uri string, ResourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT, requestSourceKinds COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS) error {
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.RemoveWebResourceRequestedFilterWithRequestSourceKinds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(ResourceContext),
		uintptr(requestSourceKinds),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
