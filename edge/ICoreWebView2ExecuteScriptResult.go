package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ExecuteScriptResult 表示脚本执行的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2executescriptresult
type ICoreWebView2ExecuteScriptResult struct {
	Vtbl *ICoreWebView2ExecuteScriptResultVtbl
}

type ICoreWebView2ExecuteScriptResultVtbl struct {
	IUnknownVtbl
	GetSucceeded         ComProc
	GetResultAsJson      ComProc
	TryGetResultAsString ComProc
	GetException         ComProc
}

func (i *ICoreWebView2ExecuteScriptResult) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ExecuteScriptResult) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ExecuteScriptResult) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSucceeded 如果 ExecuteScriptWithResult 成功执行脚本且没有未处理的异常，并且结果可通过 GetResultAsJson 或 TryGetResultAsString 方法获取，那么返回 true。
//   - 如果为 false，则脚本执行发生了未处理的异常，可通过 GetException 方法获取该异常。
func (i *ICoreWebView2ExecuteScriptResult) GetSucceeded() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetSucceeded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetResultAsJson 获取脚本执行结果的 JSON 表示。
//   - 没有显式返回值的函数会返回 undefined。
//   - 如果运行的脚本抛出未处理的异常，那么结果为“null”。
//   - 如果在 ContentLoading 之前运行该方法，脚本将不会执行，并且会返回字符串“null”。
func (i *ICoreWebView2ExecuteScriptResult) GetResultAsJson() (string, error) {
	var _jsonResult *uint16
	r, _, _ := i.Vtbl.GetResultAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_jsonResult)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	jsonResult := common.UTF16PtrToString(_jsonResult)
	wapi.CoTaskMemFree(unsafe.Pointer(_jsonResult))
	return jsonResult, nil
}

// TryGetResultAsString 尝试将脚本执行结果作为字符串获取。
//   - 如果 GetSucceeded 为 true 且脚本执行结果是字符串，此方法会提供该字符串结果的值.
//   - 当 js 结果不是字符串类型时，我们将得到 false 变量值。
func (i *ICoreWebView2ExecuteScriptResult) TryGetResultAsString() (string, bool, error) {
	var _stringResult *uint16
	var value bool
	r, _, _ := i.Vtbl.TryGetResultAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_stringResult)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", false, syscall.Errno(r)
	}
	stringResult := common.UTF16PtrToString(_stringResult)
	wapi.CoTaskMemFree(unsafe.Pointer(_stringResult))
	return stringResult, value, nil
}

// GetException 获取脚本执行过程中发生的异常信息。
func (i *ICoreWebView2ExecuteScriptResult) GetException() (*ICoreWebView2ScriptException, error) {
	var exception *ICoreWebView2ScriptException
	r, _, _ := i.Vtbl.GetException.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&exception)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return exception, nil
}

// MustGetSucceeded 如果 ExecuteScriptWithResult 成功执行脚本且没有未处理的异常，并且结果可通过 GetResultAsJson 或 TryGetResultAsString 方法获取，那么返回 true。出错时会触发全局错误回调。
//   - 如果为 false，则脚本执行发生了未处理的异常，可通过 GetException 方法获取该异常。
func (i *ICoreWebView2ExecuteScriptResult) MustGetSucceeded() bool {
	value, err := i.GetSucceeded()
	ReportErrorAuto(err)
	return value
}

// MustGetResultAsJson 获取脚本执行结果的 JSON 表示。出错时会触发全局错误回调。
//   - 没有显式返回值的函数会返回 undefined。
//   - 如果运行的脚本抛出未处理的异常，那么结果为“null”。
//   - 如果在 ContentLoading 之前运行该方法，脚本将不会执行，并且会返回字符串“null”。
func (i *ICoreWebView2ExecuteScriptResult) MustGetResultAsJson() string {
	jsonResult, err := i.GetResultAsJson()
	ReportErrorAuto(err)
	return jsonResult
}

// MustGetException 获取脚本执行过程中发生的异常信息。出错时会触发全局错误回调。
func (i *ICoreWebView2ExecuteScriptResult) MustGetException() *ICoreWebView2ScriptException {
	exception, err := i.GetException()
	ReportErrorAuto(err)
	return exception
}

// MustTryGetResultAsString 尝试将脚本执行结果作为字符串获取。出错时会触发全局错误回调。
//   - 如果 GetSucceeded 为 true 且脚本执行结果是字符串，此方法会提供该字符串结果的值.
//   - 当 js 结果不是字符串类型时，我们将得到 false 变量值。
func (i *ICoreWebView2ExecuteScriptResult) MustTryGetResultAsString() (string, bool) {
	stringResult, ok, err := i.TryGetResultAsString()
	ReportErrorAuto(err)
	return stringResult, ok
}
