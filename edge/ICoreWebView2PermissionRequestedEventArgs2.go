package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2PermissionRequestedEventArgs2 是 ICoreWebView2PermissionRequestedEventArgs 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs2
type ICoreWebView2PermissionRequestedEventArgs2 struct {
	ICoreWebView2PermissionRequestedEventArgs
}

// GetHandled 获取 Handled 属性。
func (i *ICoreWebView2PermissionRequestedEventArgs2) GetHandled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetHandled 设置 Handled 属性。
//   - 默认情况下，PermissionRequested 事件处理程序会在 ICoreWebView2Frame 和 ICoreWebView2 上都被调用，其中 ICoreWebView2Frame 的事件处理程序会先被调用。
//   - 主机可以在  ICoreWebView2Frame 事件处理程序中将此标志设置为 TRUE，以阻止其余 ICoreWebView2 事件处理程序被调用。
//   - 如果对事件参数使用延迟，则必须在使用延迟之前同步将 Handled 设置为 TRUE，以防止调用 ICoreWebView2 的事件处理程序。
func (i *ICoreWebView2PermissionRequestedEventArgs2) SetHandled(value bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetHandled 获取 Handled 属性。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs2) MustGetHandled() bool {
	value, err := i.GetHandled()
	ReportErrorAtuo(err)
	return value
}
