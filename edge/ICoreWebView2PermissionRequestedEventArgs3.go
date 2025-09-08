package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2PermissionRequestedEventArgs3 是 ICoreWebView2PermissionRequestedEventArgs2 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3
type ICoreWebView2PermissionRequestedEventArgs3 struct {
	ICoreWebView2PermissionRequestedEventArgs2
}

// GetSavesInProfile 获取 SavesInProfile 属性。
func (i *ICoreWebView2PermissionRequestedEventArgs3) GetSavesInProfile() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetSavesInProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetSavesInProfile 设置 SavesInProfile 属性。
//   - 通过 PermissionRequested 事件设置的权限状态默认保存在配置文件中；它会跨会话保持，并成为未来 PermissionRequested 事件的新默认行为。
//   - 浏览器的启发式算法会影响在配置文件中保存状态时事件是否继续触发。将 SavesInProfile 属性设置为 FALSE，可使状态不会在当前请求之外保留，并且会继续接收针对此来源和权限类型的 PermissionRequested 事件。
func (i *ICoreWebView2PermissionRequestedEventArgs3) SetSavesInProfile(value bool) error {
	r, _, _ := i.Vtbl.PutSavesInProfile.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetSavesInProfile 获取 SavesInProfile 属性。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionRequestedEventArgs3) MustGetSavesInProfile() bool {
	value, err := i.GetSavesInProfile()
	ReportErrorAuto(err)
	return value
}
