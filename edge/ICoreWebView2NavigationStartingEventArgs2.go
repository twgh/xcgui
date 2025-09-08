package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2NavigationStartingEventArgs2 是 ICoreWebView2NavigationStartingEventArgs 的延续, 使开发者能够提供额外允许的框架祖先。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs2
type ICoreWebView2NavigationStartingEventArgs2 struct {
	ICoreWebView2NavigationStartingEventArgs
}

// GetAdditionalAllowedFrameAncestors 获取额外允许的框架祖先。
func (i *ICoreWebView2NavigationStartingEventArgs2) GetAdditionalAllowedFrameAncestors() (string, error) {
	var _ancestors *uint16
	r, _, _ := i.Vtbl.GetAdditionalAllowedFrameAncestors.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_ancestors)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	ancestors := common.UTF16PtrToString(_ancestors)
	wapi.CoTaskMemFree(unsafe.Pointer(_ancestors))
	return ancestors, nil
}

// SetAdditionalAllowedFrameAncestors 设置额外允许的框架祖先。
func (i *ICoreWebView2NavigationStartingEventArgs2) SetAdditionalAllowedFrameAncestors(ancestors string) error {
	_ancestors, err := syscall.UTF16PtrFromString(ancestors)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutAdditionalAllowedFrameAncestors.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_ancestors)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetAdditionalAllowedFrameAncestors 获取额外允许的框架祖先。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationStartingEventArgs2) MustGetAdditionalAllowedFrameAncestors() string {
	ancestors, err := i.GetAdditionalAllowedFrameAncestors()
	ReportErrorAuto(err)
	return ancestors
}
