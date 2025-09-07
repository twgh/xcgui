package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2NonClientRegionChangedEventArgs 是非客户区区域变更事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventargs
type ICoreWebView2NonClientRegionChangedEventArgs struct {
	Vtbl *ICoreWebView2NonClientRegionChangedEventArgsVtbl
}

type ICoreWebView2NonClientRegionChangedEventArgsVtbl struct {
	IUnknownVtbl
	GetRegionKind ComProc
}

func (i *ICoreWebView2NonClientRegionChangedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NonClientRegionChangedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NonClientRegionChangedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetRegionKind 获取区域更改事件所对应的非客户区的区域类型。
func (i *ICoreWebView2NonClientRegionChangedEventArgs) GetRegionKind() (COREWEBVIEW2_NON_CLIENT_REGION_KIND, error) {
	var value COREWEBVIEW2_NON_CLIENT_REGION_KIND
	r, _, _ := i.Vtbl.GetRegionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetRegionKind 获取区域更改事件所对应的非客户区的区域类型。出错时会触发全局错误回调。
func (i *ICoreWebView2NonClientRegionChangedEventArgs) MustGetRegionKind() COREWEBVIEW2_NON_CLIENT_REGION_KIND {
	value, err := i.GetRegionKind()
	ReportErrorAtuo(err)
	return value
}
