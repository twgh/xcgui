package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2Settings6 是 ICoreWebView2Settings5 接口的延续, 支持管理是否启用滑动导航功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings6
type ICoreWebView2Settings6 struct {
	ICoreWebView2Settings5
}

// GetIsSwipeNavigationEnabled 获取是否允许滑动导航。
func (i *ICoreWebView2Settings6) GetIsSwipeNavigationEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsSwipeNavigationEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsSwipeNavigationEnabled 设置是否允许滑动导航。默认值为 true。
func (i *ICoreWebView2Settings6) SetIsSwipeNavigationEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsSwipeNavigationEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsSwipeNavigationEnabled 获取是否允许滑动导航。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings6) MustGetIsSwipeNavigationEnabled() bool {
	enabled, err := i.GetIsSwipeNavigationEnabled()
	ReportErrorAuto(err)
	return enabled
}
