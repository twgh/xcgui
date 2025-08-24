package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings5 是 ICoreWebView2Settings4 接口的延续, 支持管理是否启用捏合缩放功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings5
type ICoreWebView2Settings5 struct {
	ICoreWebView2Settings4
}

// GetIsPinchZoomEnabled 获取是否允许缩放。
func (i *ICoreWebView2Settings5) GetIsPinchZoomEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsPinchZoomEnabled 设置是否允许缩放。默认值为 true。
func (i *ICoreWebView2Settings5) SetIsPinchZoomEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsPinchZoomEnabled 获取是否允许缩放。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings5) MustGetIsPinchZoomEnabled() bool {
	enabled, err := i.GetIsPinchZoomEnabled()
	ReportErrorAtuo(err)
	return enabled
}
