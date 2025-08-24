package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions3 提供用于创建 WebView2 环境以管理崩溃报告的附加选项。
//
// 109.0.1518.46
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3
type ICoreWebView2EnvironmentOptions3 struct {
	ICoreWebView2EnvironmentOptions2
}

// GetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) GetIsCustomCrashReportingEnabled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetIsCustomCrashReportingEnabled 设置 Windows 是否会将崩溃数据发送到 Microsoft 端点。默认为 false。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) SetIsCustomCrashReportingEnabled(value bool) error {
	r, _, _ := i.Vtbl.PutIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。出错时会触发全局错误回调。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) MustGetIsCustomCrashReportingEnabled() bool {
	value, err := i.GetIsCustomCrashReportingEnabled()
	ReportErrorAtuo(err)
	return value
}
