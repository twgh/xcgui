package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_19 是 ICoreWebView2_18 接口的延续，用于管理内存使用目标级别。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_19
type ICoreWebView2_19 struct {
	ICoreWebView2_18
}

// GetMemoryUsageTargetLevel 获取 WebView2 的内存使用目标级别。
func (i *ICoreWebView2_19) GetMemoryUsageTargetLevel() (COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL, error) {
	var value COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL
	r, _, _ := i.Vtbl.GetMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMemoryUsageTargetLevel 设置 WebView2 的内存使用目标级别。
//   - 脚本不会受到影响，会继续运行。这对于那些处于非活动状态但仍希望运行脚本和 / 或保持网络连接活跃的应用程序很有用，因此它们无法调用 TrySuspend 和 Resume 来减少内存消耗。
//   - 这些应用程序可以在变为非活动状态时，将内存使用目标级别设置为 COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW，在变为活动状态时，再将其设置回 COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_NORMAL。
//   - 更改内存使用级别是一种尽力而为的操作。
func (i *ICoreWebView2_19) SetMemoryUsageTargetLevel(value COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL) error {
	r, _, _ := i.Vtbl.PutMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetMemoryUsageTargetLevel 获取 WebView2 的内存使用目标级别。出错时会触发全局错误回调。
func (i *ICoreWebView2_19) MustGetMemoryUsageTargetLevel() COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL {
	value, err := i.GetMemoryUsageTargetLevel()
	ReportErrorAuto(err)
	return value
}
