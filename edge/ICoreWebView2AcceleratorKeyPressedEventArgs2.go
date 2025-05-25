package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2AcceleratorKeyPressedEventArgs2 是 ICoreWebView2AcceleratorKeyPressedEventArgs 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs2
type ICoreWebView2AcceleratorKeyPressedEventArgs2 struct {
	Vtbl *ICoreWebView2AcceleratorKeyPressedEventArgs2Vtbl
}

type ICoreWebView2AcceleratorKeyPressedEventArgs2Vtbl struct {
	ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
	GetIsBrowserAcceleratorKeyEnabled ComProc
	PutIsBrowserAcceleratorKeyEnabled ComProc
}

// GetIsBrowserAcceleratorKeyEnabled 获取特定浏览器快捷键是否启用.
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) GetIsBrowserAcceleratorKeyEnabled() (bool, error) {
	var enabled int32
	r, _, err := i.Vtbl.GetIsBrowserAcceleratorKeyEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled != 0, nil
}

// MustGetIsBrowserAcceleratorKeyEnabled 获取特定浏览器快捷键是否启用. 出错时会触发全局错误回调。
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) MustGetIsBrowserAcceleratorKeyEnabled() bool {
	enabled, err := i.GetIsBrowserAcceleratorKeyEnabled()
	ReportError2(err)
	return enabled
}

// PutIsBrowserAcceleratorKeyEnabled 设置特定浏览器快捷键是否启用.
//   - 浏览器快捷键是用于访问网络浏览器特定功能的按键或按键组合.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs2) PutIsBrowserAcceleratorKeyEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutIsBrowserAcceleratorKeyEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
