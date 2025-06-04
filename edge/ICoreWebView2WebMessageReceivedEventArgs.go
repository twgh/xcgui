package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2WebMessageReceivedEventArgs WebMessageReceived 事件的事件参数.
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
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// TryGetWebMessageAsString 尝试获取 web 消息作为字符串。
//   - 如果发布的消息是其他类型的 JavaScript 类型，则该方法失败，并返回错误: wapi.E_INVALIDARG。
func (e *ICoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (string, error) {
	var _message *uint16
	r, _, err := e.Vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_message)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	message := common.UTF16PtrToString(_message)
	wapi.CoTaskMemFree(unsafe.Pointer(_message))
	return message, nil
}

// MustTryGetWebMessageAsString 尝试获取 web 消息作为字符串。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustTryGetWebMessageAsString() string {
	message, _ := e.TryGetWebMessageAsString()
	return message
}

// GetWebMessageAsJSON 获取 web 消息作为 JSON 字符串。
func (e *ICoreWebView2WebMessageReceivedEventArgs) GetWebMessageAsJSON() (string, error) {
	var _json *uint16
	r, _, err := e.Vtbl.GetWebMessageAsJSON.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_json)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	json := common.UTF16PtrToString(_json)
	wapi.CoTaskMemFree(unsafe.Pointer(_json))
	return json, nil
}

// MustGetWebMessageAsJSON 获取 web 消息作为 JSON 字符串。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustGetWebMessageAsJSON() string {
	json, _ := e.GetWebMessageAsJSON()
	return json
}

// GetSource 获取发送此 web 消息的文档的 URI。
func (e *ICoreWebView2WebMessageReceivedEventArgs) GetSource() (string, error) {
	var _source *uint16
	r, _, err := e.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_source)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	source := common.UTF16PtrToString(_source)
	wapi.CoTaskMemFree(unsafe.Pointer(_source))
	return source, nil
}

// MustGetSource 获取发送此 web 消息的文档的 URI。出错时会触发全局错误回调。
func (e *ICoreWebView2WebMessageReceivedEventArgs) MustGetSource() string {
	source, _ := e.GetSource()
	return source
}
