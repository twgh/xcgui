package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2_12 是 ICoreWebView2_11 接口的延续，支持状态栏文本改变事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_12
type ICoreWebView2_12 struct {
	ICoreWebView2_11
}

// AddStatusBarTextChanged 添加状态栏文本改变事件处理程序。
//   - 当 WebView 显示状态消息、URL 或空字符串（表示隐藏状态栏）时触发。
func (i *ICoreWebView2_12) AddStatusBarTextChanged(handler *ICoreWebView2StatusBarTextChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddStatusBarTextChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveStatusBarTextChanged 移除状态栏文本改变事件处理程序。
func (i *ICoreWebView2_12) RemoveStatusBarTextChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveStatusBarTextChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetStatusBarText 获取当前状态栏文本。
func (i *ICoreWebView2_12) GetStatusBarText() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetStatusBarText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// MustGetStatusBarText 获取当前状态栏文本。出错时会触发全局错误回调。
func (i *ICoreWebView2_12) MustGetStatusBarText() string {
	result, err := i.GetStatusBarText()
	ReportErrorAuto(err)
	return result
}
