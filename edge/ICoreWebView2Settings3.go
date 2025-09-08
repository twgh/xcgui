package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2Settings3 是 ICoreWebView2Settings2 接口的延续, 支持管理是否启用浏览器快捷键。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings3
type ICoreWebView2Settings3 struct {
	ICoreWebView2Settings2
}

// GetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。
func (i *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetAreBrowserAcceleratorKeysEnabled 设置是否允许浏览器快捷键。默认值为 true。此设置对 AcceleratorKeyPressed 事件没有影响。
func (i *ICoreWebView2Settings3) SetAreBrowserAcceleratorKeysEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings3) MustGetAreBrowserAcceleratorKeysEnabled() bool {
	enabled, err := i.GetAreBrowserAcceleratorKeysEnabled()
	ReportErrorAuto(err)
	return enabled
}
