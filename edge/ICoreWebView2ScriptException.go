package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ScriptException 表示一个 JavaScript 异常。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2scriptexception
type ICoreWebView2ScriptException struct {
	Vtbl *ICoreWebView2ScriptExceptionVtbl
}

type ICoreWebView2ScriptExceptionVtbl struct {
	IUnknownVtbl
	GetLineNumber   ComProc
	GetColumnNumber ComProc
	GetName         ComProc
	GetMessage      ComProc
	GetToJson       ComProc
}

func (i *ICoreWebView2ScriptException) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScriptException) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScriptException) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetLineNumber 获取发生异常的源代码的行号。
//   - 在 JSON 中，它是 exceptionDetail.lineNumber。请注意，该位置从 0 开始。
func (i *ICoreWebView2ScriptException) GetLineNumber() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetLineNumber.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetColumnNumber 获取异常发生的源代码列号。
//   - 在 JSON 中，它是 exceptionDetail.columnNumber。请注意，此位置从 0 开始。
func (i *ICoreWebView2ScriptException) GetColumnNumber() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetColumnNumber.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetName 获取异常的类名。
//   - 在 JSON 中，它是 exceptionDetail.exception.className。
//   - 如果异常没有类名，它就是空字符串。如果脚本抛出非 Error 对象（例如 throw "abc";），就会出现这种情况。
func (i *ICoreWebView2ScriptException) GetName() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetName.Call(
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

// GetMessage 获取异常的信息，可能还包含堆栈。
//   - 在 JSON 中，它是 exceptionDetail.exception.description。
//   - 如果异常没有描述，这就是空字符串。如果脚本抛出非 Error 对象（例如 throw "abc";），就会出现这种情况。
func (i *ICoreWebView2ScriptException) GetMessage() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetMessage.Call(
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

// GetToJson 获取异常的 JSON 字符串形式的所有详细信息。
//   - 如果脚本抛出了非 Error 对象，例如 throw "abc"; 或任何其他非 Error 对象，你可以获取该对象的特定属性。
func (i *ICoreWebView2ScriptException) GetToJson() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetToJson.Call(
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

// MustGetLineNumber 获取发生异常的源代码的行号。出错时会触发全局错误回调。
//   - 在 JSON 中，它是 exceptionDetail.lineNumber。请注意，该位置从 0 开始。
func (i *ICoreWebView2ScriptException) MustGetLineNumber() uint32 {
	value, err := i.GetLineNumber()
	ReportErrorAtuo(err)
	return value
}

// MustGetColumnNumber 获取异常发生的源代码列号。出错时会触发全局错误回调。
//   - 在 JSON 中，它是 exceptionDetail.columnNumber。请注意，此位置从 0 开始。
func (i *ICoreWebView2ScriptException) MustGetColumnNumber() uint32 {
	value, err := i.GetColumnNumber()
	ReportErrorAtuo(err)
	return value
}

// MustGetName 获取异常的类名。出错时会触发全局错误回调。
//   - 在 JSON 中，它是 exceptionDetail.exception.className。
//   - 如果异常没有类名，它就是空字符串。如果脚本抛出非 Error 对象（例如 throw "abc";），就会出现这种情况。
func (i *ICoreWebView2ScriptException) MustGetName() string {
	value, err := i.GetName()
	ReportErrorAtuo(err)
	return value
}

// MustGetMessage 获取异常的信息，可能还包含堆栈。出错时会触发全局错误回调。
//   - 在 JSON 中，它是 exceptionDetail.exception.description。
//   - 如果异常没有描述，这就是空字符串。如果脚本抛出非 Error 对象（例如 throw "abc";），就会出现这种情况。
func (i *ICoreWebView2ScriptException) MustGetMessage() string {
	value, err := i.GetMessage()
	ReportErrorAtuo(err)
	return value
}

// MustGetToJson 获取异常的 JSON 字符串形式的所有详细信息。出错时会触发全局错误回调。
//   - 如果脚本抛出了非 Error 对象，例如 throw "abc"; 或任何其他非 Error 对象，你可以获取该对象的特定属性。
func (i *ICoreWebView2ScriptException) MustGetToJson() string {
	value, err := i.GetToJson()
	ReportErrorAtuo(err)
	return value
}
