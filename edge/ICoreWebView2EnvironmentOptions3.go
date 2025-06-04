package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions3 提供用于创建 WebView2 环境以管理崩溃报告的附加选项。
//
// 109.0.1518.46
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3
type ICoreWebView2EnvironmentOptions3 struct {
	Vtbl *ICoreWebView2EnvironmentOptions3Vtbl
}

type ICoreWebView2EnvironmentOptions3Vtbl struct {
	IUnknownVtbl
	GetIsCustomCrashReportingEnabled ComProc
	PutIsCustomCrashReportingEnabled ComProc
}

func (i *ICoreWebView2EnvironmentOptions3) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions3) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions3) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) GetIsCustomCrashReportingEnabled() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// MustGetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。出错时会触发全局错误回调。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) MustGetIsCustomCrashReportingEnabled() bool {
	value, err := i.GetIsCustomCrashReportingEnabled()
	ReportErrorAtuo(err)
	return value
}

// PutIsCustomCrashReportingEnabled 设置 Windows 是否会将崩溃数据发送到 Microsoft 端点。默认为 false。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) PutIsCustomCrashReportingEnabled(value bool) error {
	r, _, err := i.Vtbl.PutIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
