package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2WebResourceRequestedEventArgs2 是 ICoreWebView2WebResourceRequestedEventArgs 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs2
type ICoreWebView2WebResourceRequestedEventArgs2 struct {
	ICoreWebView2WebResourceRequestedEventArgs
}

// GetRequestedSourceKinds 获取请求的网络资源的来源类型。
func (i *ICoreWebView2WebResourceRequestedEventArgs2) GetRequestedSourceKinds() (COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS, error) {
	var value COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS
	r, _, _ := i.Vtbl.GetRequestedSourceKinds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetRequestedSourceKinds 获取请求的网络资源的来源类型。出错时会触发全局错误回调。
func (i *ICoreWebView2WebResourceRequestedEventArgs2) MustGetRequestedSourceKinds() COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS {
	result, err := i.GetRequestedSourceKinds()
	ReportErrorAuto(err)
	return result
}
