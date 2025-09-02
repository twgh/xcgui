package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 是 ICoreWebView2DevToolsProtocolEventReceivedEventArgs 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs2
type ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 struct {
	ICoreWebView2DevToolsProtocolEventReceivedEventArgs
}

// GetSessionId 获取事件起源的目标的会话ID。
//   - 如果事件来自顶级页面的默认会话，则会返回空字符串作为 sessionId。
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) GetSessionId() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetSessionId.Call(
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

// MustGetSessionId 获取事件起源的目标的会话ID。出错时会触发全局错误回调。
//   - 如果事件来自顶级页面的默认会话，则会返回空字符串作为 sessionId。
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) MustGetSessionId() string {
	value, err := i.GetSessionId()
	ReportErrorAtuo(err)
	return value
}
