package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings3 是 ICoreWebView2Settings2 接口的延续, 支持管理是否启用浏览器快捷键。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings3
type ICoreWebView2Settings3 struct {
	Vtbl *ICoreWebView2Settings3Vtbl
}

type ICoreWebView2Settings3Vtbl struct {
	ICoreWebView2Settings2Vtbl
	GetAreBrowserAcceleratorKeysEnabled ComProc
	PutAreBrowserAcceleratorKeysEnabled ComProc
}

// GetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。
func (i *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// MustGetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。忽略错误。
func (i *ICoreWebView2Settings3) MustGetAreBrowserAcceleratorKeysEnabled() bool {
	enabled, _ := i.GetAreBrowserAcceleratorKeysEnabled()
	return enabled
}

// PutAreBrowserAcceleratorKeysEnabled 设置是否允许浏览器快捷键。默认值为 true。此设置对 AcceleratorKeyPressed 事件没有影响。
func (i *ICoreWebView2Settings3) PutAreBrowserAcceleratorKeysEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutAreBrowserAcceleratorKeysEnabled.Call(
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
