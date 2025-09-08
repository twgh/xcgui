package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2WebMessageReceivedEventArgs 是 WebMessageReceived 事件的事件参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs
type ICoreWebView2WebMessageReceivedEventArgs struct {
	Vtbl *ICoreWebView2WebMessageReceivedEventArgsVtbl
}

type ICoreWebView2WebMessageReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJSON      ComProc
	TryGetWebMessageAsString ComProc
	// 2
	GetAdditionalObjects ComProc
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// TryGetWebMessageAsString 尝试获取 web 消息作为字符串。
//   - 如果发布的消息是其他类型的 JavaScript 类型，则该方法失败，并返回错误: wapi.E_INVALIDARG。
func (e *ICoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (string, error) {
	var _message *uint16
	r, _, _ := e.Vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_message)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	message := common.UTF16PtrToString(_message)
	wapi.CoTaskMemFree(unsafe.Pointer(_message))
	return message, nil
}

// GetWebMessageAsJSON 获取 web 消息作为 JSON 字符串。
func (e *ICoreWebView2WebMessageReceivedEventArgs) GetWebMessageAsJSON() (string, error) {
	var _json *uint16
	r, _, _ := e.Vtbl.GetWebMessageAsJSON.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_json)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	json := common.UTF16PtrToString(_json)
	wapi.CoTaskMemFree(unsafe.Pointer(_json))
	return json, nil
}

// GetSource 获取发送此 web 消息的文档的 URI。
func (e *ICoreWebView2WebMessageReceivedEventArgs) GetSource() (string, error) {
	var _source *uint16
	r, _, _ := e.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_source)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	source := common.UTF16PtrToString(_source)
	wapi.CoTaskMemFree(unsafe.Pointer(_source))
	return source, nil
}

// GetICoreWebView2WebMessageReceivedEventArgs2 获取 ICoreWebView2WebMessageReceivedEventArgs2。
func (i *ICoreWebView2WebMessageReceivedEventArgs) GetICoreWebView2WebMessageReceivedEventArgs2() (*ICoreWebView2WebMessageReceivedEventArgs2, error) {
	var result *ICoreWebView2WebMessageReceivedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2WebMessageReceivedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustTryGetWebMessageAsString 尝试获取 web 消息作为字符串。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustTryGetWebMessageAsString() string {
	message, _ := e.TryGetWebMessageAsString()
	return message
}

// MustGetWebMessageAsJSON 获取 web 消息作为 JSON 字符串。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustGetWebMessageAsJSON() string {
	json, _ := e.GetWebMessageAsJSON()
	return json
}

// MustGetSource 获取发送此 web 消息的文档的 URI。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustGetSource() string {
	source, _ := e.GetSource()
	return source
}

// MustGetICoreWebView2WebMessageReceivedEventArgs2 获取 ICoreWebView2WebMessageReceivedEventArgs2。出错时会触发全局错误回调。
func (i *ICoreWebView2WebMessageReceivedEventArgs) MustGetICoreWebView2WebMessageReceivedEventArgs2() *ICoreWebView2WebMessageReceivedEventArgs2 {
	result, err := i.GetICoreWebView2WebMessageReceivedEventArgs2()
	ReportErrorAuto(err)
	return result
}
