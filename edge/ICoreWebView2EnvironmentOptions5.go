package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions5 提供用于创建 WebView2 环境以管理跟踪防护的附加选项。
//
// 111.0.1661.34
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5
type ICoreWebView2EnvironmentOptions5 struct {
	Vtbl *ICoreWebView2EnvironmentOptions5Vtbl
}

type ICoreWebView2EnvironmentOptions5Vtbl struct {
	IUnknownVtbl
	GetEnableTrackingPrevention ComProc
	PutEnableTrackingPrevention ComProc
}

func (i *ICoreWebView2EnvironmentOptions5) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions5) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions5) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetEnableTrackingPrevention 获取是否启用 WebView2 中的跟踪防护功能。
//
// 111.0.1661.34
func (i *ICoreWebView2EnvironmentOptions5) GetEnableTrackingPrevention() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetEnableTrackingPrevention.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// PutEnableTrackingPrevention 设置是否启用 WebView2 中的跟踪防护功能。默认为 true。
//   - 此属性可启用/禁用在同一环境中创建的所有 WebView2 的跟踪防护功能。默认情况下，该功能处于启用状态，用于阻止潜在有害的跟踪器以及来自从未访问过的网站的跟踪器，并设置为 COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED 或配置文件上次更改/保存的任何值。
//   - 如果应用程序仅在 WebView2 中渲染已知安全的内容，则可以将此属性设置为 false 以禁用跟踪防护功能。在创建环境时禁用此功能还可以通过跳过相关代码来提高运行时性能。
//   - 如果将 WebView2 用作具有任意导航功能的“完整浏览器”，则不应禁用此属性，并且应保护最终用户的隐私。
//
// 111.0.1661.34
func (i *ICoreWebView2EnvironmentOptions5) PutEnableTrackingPrevention(value bool) error {
	r, _, err := i.Vtbl.PutEnableTrackingPrevention.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
