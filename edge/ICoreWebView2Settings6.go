package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings6 是 ICoreWebView2Settings5 接口的延续, 支持管理是否启用滑动导航功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings6
type ICoreWebView2Settings6 struct {
	Vtbl *ICoreWebView2Settings6Vtbl
}

type ICoreWebView2Settings6Vtbl struct {
	ICoreWebView2Settings5Vtbl
	GetIsSwipeNavigationEnabled ComProc
	PutIsSwipeNavigationEnabled ComProc
}

func (i *ICoreWebView2Settings6) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings6) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings6) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsSwipeNavigationEnabled 获取是否允许滑动导航。
func (i *ICoreWebView2Settings6) GetIsSwipeNavigationEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetIsSwipeNavigationEnabled.Call(
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

// MustGetIsSwipeNavigationEnabled 获取是否允许滑动导航。忽略错误。
func (i *ICoreWebView2Settings6) MustGetIsSwipeNavigationEnabled() bool {
	enabled, _ := i.GetIsSwipeNavigationEnabled()
	return enabled
}

// PutIsSwipeNavigationEnabled 设置是否允许滑动导航。默认值为 true。
func (i *ICoreWebView2Settings6) PutIsSwipeNavigationEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutIsSwipeNavigationEnabled.Call(
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
