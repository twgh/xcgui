package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2AcceleratorKeyPressedEventArgs2 是 ICoreWebView2AcceleratorKeyPressedEventArgs 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs2
type ICoreWebView2AcceleratorKeyPressedEventArgs2 struct {
	ICoreWebView2AcceleratorKeyPressedEventArgs
}

// GetIsBrowserAcceleratorKeyEnabled 获取特定浏览器快捷键是否启用.
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) GetIsBrowserAcceleratorKeyEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsBrowserAcceleratorKeyEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsBrowserAcceleratorKeyEnabled 设置特定浏览器快捷键是否启用.
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) SetIsBrowserAcceleratorKeyEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsBrowserAcceleratorKeyEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsBrowserAcceleratorKeyEnabled 获取特定浏览器快捷键是否启用. 出错时会触发全局错误回调。
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) MustGetIsBrowserAcceleratorKeyEnabled() bool {
	enabled, err := i.GetIsBrowserAcceleratorKeyEnabled()
	ReportErrorAuto(err)
	return enabled
}
