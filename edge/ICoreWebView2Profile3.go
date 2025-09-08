package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile3 是 ICoreWebView2Profile2 的延续, 提供了跟踪防护级别的设置功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile3
type ICoreWebView2Profile3 struct {
	ICoreWebView2Profile2
}

// GetPreferredTrackingPreventionLevel 获取首选的跟踪防护级别。
func (i *ICoreWebView2Profile3) GetPreferredTrackingPreventionLevel() (COREWEBVIEW2_TRACKING_PREVENTION_LEVEL, error) {
	var _value COREWEBVIEW2_TRACKING_PREVENTION_LEVEL
	r, _, _ := i.Vtbl.GetPreferredTrackingPreventionLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// SetPreferredTrackingPreventionLevel 设置首选的跟踪防护级别。
//   - 所有共享同一配置文件的 WebView2 都将受到影响，并且该值会保存在用户数据文件夹中。
func (i *ICoreWebView2Profile3) SetPreferredTrackingPreventionLevel(value COREWEBVIEW2_TRACKING_PREVENTION_LEVEL) error {
	r, _, _ := i.Vtbl.PutPreferredTrackingPreventionLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetPreferredTrackingPreventionLevel 获取首选的跟踪防护级别。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile3) MustGetPreferredTrackingPreventionLevel() COREWEBVIEW2_TRACKING_PREVENTION_LEVEL {
	result, err := i.GetPreferredTrackingPreventionLevel()
	ReportErrorAuto(err)
	return result
}
