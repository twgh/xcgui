package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2NavigationStartingEventArgs3 是 ICoreWebView2NavigationStartingEventArgs2 的延续, 使开发者能够获取有关导航类型的更多信息。。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3
type ICoreWebView2NavigationStartingEventArgs3 struct {
	ICoreWebView2NavigationStartingEventArgs2
}

// GetNavigationKind 获取导航的导航类型。
func (i *ICoreWebView2NavigationStartingEventArgs3) GetNavigationKind() (COREWEBVIEW2_NAVIGATION_KIND, error) {
	var kind COREWEBVIEW2_NAVIGATION_KIND
	r, _, _ := i.Vtbl.GetNavigationKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return kind, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetNavigationKind 获取导航的导航类型。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs3) MustGetNavigationKind() COREWEBVIEW2_NAVIGATION_KIND {
	kind, err := i.GetNavigationKind()
	ReportErrorAuto(err)
	return kind
}
