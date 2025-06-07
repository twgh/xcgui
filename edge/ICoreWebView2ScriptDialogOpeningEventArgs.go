package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ScriptDialogOpeningEventArgs 表示脚本对话框打开事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs
type ICoreWebView2ScriptDialogOpeningEventArgs struct {
	Vtbl *ICoreWebView2ScriptDialogOpeningEventArgsVtbl
}

type ICoreWebView2ScriptDialogOpeningEventArgsVtbl struct {
	IUnknownVtbl
	GetUri         ComProc
	GetKind        ComProc
	GetMessage     ComProc
	Accept         ComProc
	GetDefaultText ComProc
	GetResultText  ComProc
	PutResultText  ComProc
	GetDeferral    ComProc
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取对话框来源的URI。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetUri() (string, error) {
	var _uri *uint16
	r, _, err := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(_uri)
	wapi.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

// MustGetUri 获取对话框来源的URI。出错时会触发全局错误回调。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// GetKind 获取 JavaScript 对话框的类型。alert、confirm、prompt 或 beforeunload。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetKind() (COREWEBVIEW2_SCRIPT_DIALOG_KIND, error) {
	var kind COREWEBVIEW2_SCRIPT_DIALOG_KIND
	r, _, err := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return kind, err
	}
	if r != 0 {
		return kind, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetKind 获取 JavaScript 对话框的类型。alert、confirm、prompt 或 beforeunload。出错时会触发全局错误回调。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) MustGetKind() COREWEBVIEW2_SCRIPT_DIALOG_KIND {
	kind, err := i.GetKind()
	ReportErrorAtuo(err)
	return kind
}

// GetMessage 获取对话框消息内容。
//   - 从 JavaScript 角度来看，这是传递给 alert、confirm 和 prompt 的第一个参数，对于 beforeunload 来说，该参数为空。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetMessage() (string, error) {
	var _message *uint16
	r, _, err := i.Vtbl.GetMessage.Call(
		uintptr(unsafe.Pointer(i)),
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

// MustGetMessage 获取对话框消息内容。出错时会触发全局错误回调。
//   - 从 JavaScript 角度来看，这是传递给 alert、confirm 和 prompt 的第一个参数，对于 beforeunload 来说，该参数为空。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) MustGetMessage() string {
	message, err := i.GetMessage()
	ReportErrorAtuo(err)
	return message
}

// Accept 接受对话框。
//   - 主机可以运行此操作，以使用 OK 来响应 confirm、prompt 和 beforeunload 对话框。
//   - 请勿运行此方法以表示取消。在JavaScript中，这意味着如果运行 confirm，则 beforeunload 和 TRUE 函数返回 Accept。
//   - 对于 ResultText 函数，如果运行 Accept，它将返回 FALSE 的值，否则返回 FALSE。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) Accept() error {
	r, _, err := i.Vtbl.Accept.Call(uintptr(unsafe.Pointer(i)))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDefaultText 获取对话框默认文本。即传递给 JavaScript 提示对话框的第二个参数。
//   - prompt JavaScript 函数的结果将此值用作默认值。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetDefaultText() (string, error) {
	var _defaultText *uint16
	r, _, err := i.Vtbl.GetDefaultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_defaultText)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defaultText := common.UTF16PtrToString(_defaultText)
	wapi.CoTaskMemFree(unsafe.Pointer(_defaultText))
	return defaultText, nil
}

// MustGetDefaultText 获取对话框默认文本。即传递给 JavaScript 提示对话框的第二个参数。
//   - prompt JavaScript 函数的结果将此值用作默认值。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) MustGetDefaultText() string {
	defaultText, err := i.GetDefaultText()
	ReportErrorAtuo(err)
	return defaultText
}

// GetResultText 获取对话框结果文本。如果运行 Accept，JavaScript 提示函数的返回值。
//   - 对于除提示之外的其他对话框类型，此值将被忽略。
//   - 如果未运行 Accept，则此值将被忽略，并且提示将返回 FALSE。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetResultText() (string, error) {
	var _resultText *uint16
	r, _, err := i.Vtbl.GetResultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_resultText)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	resultText := common.UTF16PtrToString(_resultText)
	wapi.CoTaskMemFree(unsafe.Pointer(_resultText))
	return resultText, nil
}

// MustGetResultText 获取对话框结果文本。如果运行 Accept，JavaScript 提示函数的返回值。
//   - 对于除提示之外的其他对话框类型，此值将被忽略。
//   - 如果未运行 Accept，则此值将被忽略，并且提示将返回 FALSE。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) MustGetResultText() string {
	resultText, err := i.GetResultText()
	ReportErrorAtuo(err)
	return resultText
}

// PutResultText 设置对话框结果文本。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) PutResultText(resultText string) error {
	_resultText, err := syscall.UTF16PtrFromString(resultText)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutResultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultText)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取延迟对象 ICoreWebView2Deferral。使用此操作可在稍后完成该事件。
func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}
