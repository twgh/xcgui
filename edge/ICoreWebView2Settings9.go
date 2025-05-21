package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings9 是 ICoreWebView2Settings8 接口的延续, 支持管理是否启用非客户端区域支持。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings9
type ICoreWebView2Settings9 struct {
	Vtbl *ICoreWebView2Settings9Vtbl
}

type ICoreWebView2Settings9Vtbl struct {
	ICoreWebView2Settings8Vtbl
	GetIsNonClientRegionSupportEnabled ComProc
	PutIsNonClientRegionSupportEnabled ComProc
}

func (i *ICoreWebView2Settings9) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings9) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings9) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsNonClientRegionSupportEnabled 获取是否启用非客户区域支持。
func (i *ICoreWebView2Settings9) GetIsNonClientRegionSupportEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetIsNonClientRegionSupportEnabled.Call(
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

// MustGetIsNonClientRegionSupportEnabled 获取是否启用非客户区域支持。忽略错误。
func (i *ICoreWebView2Settings9) MustGetIsNonClientRegionSupportEnabled() bool {
	enabled, _ := i.GetIsNonClientRegionSupportEnabled()
	return enabled
}

// PutIsNonClientRegionSupportEnabled 设置是否启用非客户区域支持。默认值为 false。
//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 WebView 及其宿主应用程序窗口；
//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
func (i *ICoreWebView2Settings9) PutIsNonClientRegionSupportEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutIsNonClientRegionSupportEnabled.Call(
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
