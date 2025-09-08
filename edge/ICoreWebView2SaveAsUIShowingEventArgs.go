package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2SaveAsUIShowingEventArgs 提供"另存为"UI显示事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2saveasuishowingeventargs
type ICoreWebView2SaveAsUIShowingEventArgs struct {
	Vtbl *ICoreWebView2SaveAsUIShowingEventArgsVtbl
}

type ICoreWebView2SaveAsUIShowingEventArgsVtbl struct {
	IUnknownVtbl
	GetContentMimeType       ComProc
	PutCancel                ComProc
	GetCancel                ComProc
	PutSuppressDefaultDialog ComProc
	GetSuppressDefaultDialog ComProc
	GetDeferral              ComProc
	PutSaveAsFilePath        ComProc
	GetSaveAsFilePath        ComProc
	PutAllowReplace          ComProc
	GetAllowReplace          ComProc
	PutKind                  ComProc
	GetKind                  ComProc
}

func (i *ICoreWebView2SaveAsUIShowingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveAsUIShowingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveAsUIShowingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetContentMimeType 获取内容的 MIME 类型。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetContentMimeType() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetContentMimeType.Call(
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

// SetCancel 设置是否取消"另存为"操作。
//   - 将此属性设置为 TRUE 可取消“另存为”操作并阻止下载开始。在这种情况下，ShowSaveAsUI 会返回 COREWEBVIEW2_SAVE_AS_UI_RESULT_CANCELLED。
//   - 默认值为 FALSE。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) SetCancel(value bool) error {
	r, _, _ := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCancel 获取是否取消"另存为"操作。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetCancel() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetSuppressDefaultDialog 设置是否禁止系统默认对话框。
//   - 设置为 FALSE 时，会显示默认的“另存为”对话框，并且在事件参数调用完成时，通过 SaveAsFilePath、AllowReplace 和 Kind 分配的值会被忽略。
//   - 设置为 TRUE 可执行静默“另存为”操作。系统对话框将被跳过，并使用 SaveAsFilePath、AllowReplace 和 Kind 的值。
//   - 默认值为 FALSE。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) SetSuppressDefaultDialog(value bool) error {
	r, _, _ := i.Vtbl.PutSuppressDefaultDialog.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSuppressDefaultDialog 获取是否禁止系统默认对话框。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetSuppressDefaultDialog() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetSuppressDefaultDialog.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetDeferral 获取延迟对象，用于推迟处理事件。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// SetSaveAsFilePath 设置保存文件的路径。
//   - 如果 SaveAsFilePath 无效（例如，根驱动器不存在），则“另存为”操作被拒绝，并返回  COREWEBVIEW2_SAVE_AS_INVALID_PATH。
//   - 默认值是系统根据用户的本地环境建议的路径。
//   - 如果 Kind 属性为 COREWEBVIEW2_SAVE_AS_KIND_COMPLETE，则会有一个包含资源文件的附加目录。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) SetSaveAsFilePath(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutSaveAsFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSaveAsFilePath 获取保存文件的路径。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetSaveAsFilePath() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetSaveAsFilePath.Call(
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

// SetAllowReplace 设置是否允许替换现有文件。
//   - 将此属性设置为 false 则不会替换现有文件，并将返回 COREWEBVIEW2_SAVE_AS_UI_RESULT_FILE_ALREADY_EXISTS。
//   - 默认值为 false。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) SetAllowReplace(value bool) error {
	r, _, _ := i.Vtbl.PutAllowReplace.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowReplace 获取是否允许替换现有文件。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetAllowReplace() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetAllowReplace.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetKind 设置"另存为"操作的类型。
//   - 如果当前文档不允许该类型，ShowSaveAsUI 将返回 COREWEBVIEW2_SAVE_AS_UI_RESULT_KIND_NOT_SUPPORTED。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) SetKind(value COREWEBVIEW2_SAVE_AS_KIND) error {
	r, _, _ := i.Vtbl.PutKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKind 获取"另存为"操作的类型。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) GetKind() (COREWEBVIEW2_SAVE_AS_KIND, error) {
	var value COREWEBVIEW2_SAVE_AS_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetContentMimeType 获取内容的 MIME 类型。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetContentMimeType() string {
	value, err := i.GetContentMimeType()
	ReportErrorAuto(err)
	return value
}

// MustGetCancel 获取是否取消"另存为"操作。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetCancel() bool {
	value, err := i.GetCancel()
	ReportErrorAuto(err)
	return value
}

// MustGetSuppressDefaultDialog 获取是否禁止系统默认对话框。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetSuppressDefaultDialog() bool {
	value, err := i.GetSuppressDefaultDialog()
	ReportErrorAuto(err)
	return value
}

// MustGetDeferral 获取延迟对象，用于异步处理事件。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAuto(err)
	return deferral
}

// MustGetSaveAsFilePath 获取保存文件的路径。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetSaveAsFilePath() string {
	value, err := i.GetSaveAsFilePath()
	ReportErrorAuto(err)
	return value
}

// MustGetAllowReplace 获取是否允许替换现有文件。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetAllowReplace() bool {
	value, err := i.GetAllowReplace()
	ReportErrorAuto(err)
	return value
}

// MustGetKind 获取"另存为"操作的类型。出错时会触发全局错误回调。
func (i *ICoreWebView2SaveAsUIShowingEventArgs) MustGetKind() COREWEBVIEW2_SAVE_AS_KIND {
	value, err := i.GetKind()
	ReportErrorAuto(err)
	return value
}
