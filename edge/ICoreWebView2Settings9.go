package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings9 是 ICoreWebView2Settings8 接口的延续, 支持管理是否启用非客户端区域支持。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings9
type ICoreWebView2Settings9 struct {
	ICoreWebView2Settings8
}

// GetIsNonClientRegionSupportEnabled 获取是否启用非客户区域支持。
func (i *ICoreWebView2Settings9) GetIsNonClientRegionSupportEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsNonClientRegionSupportEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsNonClientRegionSupportEnabled 设置是否启用非客户区域支持。默认值为 false。
//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 WebView 及其宿主应用程序窗口；
//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
func (i *ICoreWebView2Settings9) SetIsNonClientRegionSupportEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsNonClientRegionSupportEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsNonClientRegionSupportEnabled 获取是否启用非客户区域支持。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings9) MustGetIsNonClientRegionSupportEnabled() bool {
	enabled, err := i.GetIsNonClientRegionSupportEnabled()
	ReportErrorAtuo(err)
	return enabled
}
