package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2NavigationCompletedEventArgs2 是 ICoreWebView2NavigationCompletedEventArgs 的延续, 提供 StatusCode 属性。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs2
type ICoreWebView2NavigationCompletedEventArgs2 struct {
	ICoreWebView2NavigationCompletedEventArgs
}

// GetHttpStatusCode 获取导航的 HTTP 状态代码 (如果导航涉及 HTTP 请求)。
func (i *ICoreWebView2NavigationCompletedEventArgs2) GetHttpStatusCode() (int32, error) {
	var statusCode int32
	r, _, _ := i.Vtbl.GetHttpStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&statusCode)),
	)
	if r != 0 {
		return statusCode, syscall.Errno(r)
	}
	return statusCode, nil
}

// MustGetHttpStatusCode 获取导航的 HTTP 状态代码 (如果导航涉及 HTTP 请求)。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationCompletedEventArgs2) MustGetHttpStatusCode() int32 {
	statusCode, err := i.GetHttpStatusCode()
	ReportErrorAtuo(err)
	return statusCode
}
