package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2NewWindowRequestedEventArgs2 是 ICoreWebView2NewWindowRequestedEventArgs 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs2
type ICoreWebView2NewWindowRequestedEventArgs2 struct {
	ICoreWebView2NewWindowRequestedEventArgs
}

// GetName 获取新窗口的名称。
//   - 此窗口可通过 window.open(url, windowName) 创建，其中 windowName 参数对应 Name 属性。
//   - 如果没有向 window.open 传递 windowName，那么 Name 属性将被设置为空字符串。
//   - 此外，如果通过其他方式打开窗口，例如 <a target="windowName">...</a> 或 <iframe name="windowName">...</iframe>，那么 Name 属性将被相应地设置。
//   - 对于 target=_blank 的情况，Name 属性将是空字符串。
//   - 通过按住 ctrl 并点击链接的方式打开窗口，会导致 Name 属性被设置为空字符串。
func (i *ICoreWebView2NewWindowRequestedEventArgs2) GetName() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	name := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return name, nil
}

// MustGetName 获取新窗口的名称。出错时会触发全局错误回调。
//   - 此窗口可通过 window.open(url, windowName) 创建，其中 windowName 参数对应 Name 属性。
//   - 如果没有向 window.open 传递 windowName，那么 Name 属性将被设置为空字符串。
//   - 此外，如果通过其他方式打开窗口，例如 <a target="windowName">...</a> 或 <iframe name="windowName">...</iframe>，那么 Name 属性将被相应地设置。
//   - 对于 target=_blank 的情况，Name 属性将是空字符串。
//   - 通过按住 ctrl 并点击链接的方式打开窗口，会导致 Name 属性被设置为空字符串。
func (i *ICoreWebView2NewWindowRequestedEventArgs2) MustGetName() string {
	name, err := i.GetName()
	ReportErrorAuto(err)
	return name
}
