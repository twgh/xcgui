package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2DevToolsProtocolEventReceivedEventArgs 是 DevTools 协议事件的参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs
type ICoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	Vtbl *ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl
}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetParameterObjectAsJson ComProc
	// 2
	GetSessionId ComProc
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetParameterObjectAsJson 获取作为 JSON 字符串的事件参数对象.
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) GetParameterObjectAsJson() (string, error) {
	var _json *uint16
	r, _, _ := i.Vtbl.GetParameterObjectAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_json)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	jsonStr := common.UTF16PtrToString(_json)
	wapi.CoTaskMemFree(unsafe.Pointer(_json))
	return jsonStr, nil
}

// GetICoreWebView2DevToolsProtocolEventReceivedEventArgs2 获取 ICoreWebView2DevToolsProtocolEventReceivedEventArgs2。
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) GetICoreWebView2DevToolsProtocolEventReceivedEventArgs2() (*ICoreWebView2DevToolsProtocolEventReceivedEventArgs2, error) {
	var result *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2DevToolsProtocolEventReceivedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetParameterObjectAsJson 获取作为 JSON 字符串的事件参数对象。出错时会触发全局错误回调。
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) MustGetParameterObjectAsJson() string {
	jsonStr, err := i.GetParameterObjectAsJson()
	ReportErrorAtuo(err)
	return jsonStr
}

// MustGetICoreWebView2DevToolsProtocolEventReceivedEventArgs2 获取 ICoreWebView2DevToolsProtocolEventReceivedEventArgs2。
func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) MustGetICoreWebView2DevToolsProtocolEventReceivedEventArgs2() *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 {
	result, err := i.GetICoreWebView2DevToolsProtocolEventReceivedEventArgs2()
	ReportErrorAtuo(err)
	return result
}
