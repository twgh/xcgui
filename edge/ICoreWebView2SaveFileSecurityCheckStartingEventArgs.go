package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2SaveFileSecurityCheckStartingEventArgs 提供保存文件安全检查开始事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2savefilesecuritycheckstartingeventargs
type ICoreWebView2SaveFileSecurityCheckStartingEventArgs struct {
	Vtbl *ICoreWebView2SaveFileSecurityCheckStartingEventArgsVtbl
}

type ICoreWebView2SaveFileSecurityCheckStartingEventArgsVtbl struct {
	IUnknownVtbl
	GetCancelSave            ComProc
	PutCancelSave            ComProc
	GetDocumentOriginUri     ComProc
	GetFileExtension         ComProc
	GetFilePath              ComProc
	GetSuppressDefaultPolicy ComProc
	PutSuppressDefaultPolicy ComProc
	GetDeferral              ComProc
}

func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCancelSave 获取是否取消即将进行的保存/下载。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetCancelSave() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetCancelSave.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetCancelSave 设置是否取消即将进行的保存/下载。
//   - 默认值为 FALSE。设置为 TRUE 将在验证前取消该操作。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) SetCancelSave(value bool) error {
	r, _, _ := i.Vtbl.PutCancelSave.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDocumentOriginUri 获取发起保存操作的文档的源 URI。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetDocumentOriginUri() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetDocumentOriginUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return uri, nil
}

// GetFileExtension 获取要保存的文件的扩展名。
//   - 仅提供带句点“.”的最终扩展名。例如，“*.tar.gz”是双重扩展名，其中“.gz”是其最终扩展名。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetFileExtension() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetFileExtension.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	ext := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return ext, nil
}

// GetFilePath 获取要保存的文件的完整路径。
//   - 此方法不提供路径验证，返回的字符串可能长于 MAX_PATH。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetFilePath() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return path, nil
}

// GetSuppressDefaultPolicy 获取是否禁用默认的策略检查和安全警告。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetSuppressDefaultPolicy() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetSuppressDefaultPolicy.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetSuppressDefaultPolicy 设置是否要禁用默认的策略检查和安全警告。
//   - 默认值为 FALSE。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) SetSuppressDefaultPolicy(value bool) error {
	r, _, _ := i.Vtbl.PutSuppressDefaultPolicy.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取一个延迟对象，用于推迟处理该事件。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var value *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetCancelSave 获取是否取消即将进行的保存/下载。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetCancelSave() bool {
	result, err := i.GetCancelSave()
	ReportErrorAuto(err)
	return result
}

// MustGetDocumentOriginUri 获取发起保存操作的文档的源 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetDocumentOriginUri() string {
	result, err := i.GetDocumentOriginUri()
	ReportErrorAuto(err)
	return result
}

// MustGetFileExtension 获取要保存的文件的扩展名。出错时会触发全局错误回调。
//   - 仅提供带句点“.”的最终扩展名。例如，“*.tar.gz”是双重扩展名，其中“.gz”是其最终扩展名。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetFileExtension() string {
	result, err := i.GetFileExtension()
	ReportErrorAuto(err)
	return result
}

// MustGetFilePath 获取要保存的文件的完整路径。出错时会触发全局错误回调。
//   - 此方法不提供路径验证，返回的字符串可能长于 MAX_PATH。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetFilePath() string {
	result, err := i.GetFilePath()
	ReportErrorAuto(err)
	return result
}

// MustGetSuppressDefaultPolicy 获取是否禁用默认的策略检查和安全警告。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetSuppressDefaultPolicy() bool {
	result, err := i.GetSuppressDefaultPolicy()
	ReportErrorAuto(err)
	return result
}

// MustGetDeferral 获取一个延迟对象，用于在异步操作完成后通知 WebView2。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	result, err := i.GetDeferral()
	ReportErrorAuto(err)
	return result
}
