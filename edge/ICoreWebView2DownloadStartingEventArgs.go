package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2DownloadStartingEventArgs 是下载开始事件的参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs
type ICoreWebView2DownloadStartingEventArgs struct {
	Vtbl *ICoreWebView2DownloadStartingEventArgsVtbl
}

type ICoreWebView2DownloadStartingEventArgsVtbl struct {
	IUnknownVtbl
	GetDownloadOperation ComProc
	GetCancel            ComProc
	PutCancel            ComProc
	GetResultFilePath    ComProc
	PutResultFilePath    ComProc
	GetHandled           ComProc
	PutHandled           ComProc
	GetDeferral          ComProc
}

func (i *ICoreWebView2DownloadStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadStartingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadStartingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDownloadOperation 获取已开始下载的下载操作对象.
func (i *ICoreWebView2DownloadStartingEventArgs) GetDownloadOperation() (*ICoreWebView2DownloadOperation, error) {
	var downloadOperation *ICoreWebView2DownloadOperation
	r, _, err := i.Vtbl.GetDownloadOperation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&downloadOperation)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return downloadOperation, nil
}

// MustGetDownloadOperation 获取已开始下载的下载操作对象。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustGetDownloadOperation() *ICoreWebView2DownloadOperation {
	result, err := i.GetDownloadOperation()
	ReportErrorAtuo(err)
	return result
}

// GetCancel 获取是否取消下载.
func (i *ICoreWebView2DownloadStartingEventArgs) GetCancel() (bool, error) {
	var cancel bool
	r, _, err := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return cancel, nil
}

// MustGetCancel 获取是否取消下载。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustGetCancel() bool {
	result, err := i.GetCancel()
	ReportErrorAtuo(err)
	return result
}

// PutCancel 设置是否取消下载.
//   - 如果取消，则无论 Handled 属性如何，都不会显示下载保存对话框。
func (i *ICoreWebView2DownloadStartingEventArgs) PutCancel(cancel bool) error {
	r, _, err := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(cancel),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustPutCancel 设置是否取消下载.
//   - 如果取消，则无论 Handled 属性如何，都不会显示下载保存对话框。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustPutCancel(cancel bool) {
	err := i.PutCancel(cancel)
	ReportErrorAtuo(err)
}

// GetResultFilePath 获取结果文件路径.
func (i *ICoreWebView2DownloadStartingEventArgs) GetResultFilePath() (string, error) {
	var resultFilePath *uint16
	r, _, err := i.Vtbl.GetResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&resultFilePath)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(resultFilePath)
	wapi.CoTaskMemFree(unsafe.Pointer(resultFilePath))
	return path, nil
}

// MustGetResultFilePath 获取结果文件路径。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustGetResultFilePath() string {
	result, err := i.GetResultFilePath()
	ReportErrorAtuo(err)
	return result
}

// PutResultFilePath 设置结果文件路径.
//   - 如果设置路径，主机应确保该路径为绝对路径，包括文件名，且该路径不指向现有文件。
//   - 如果路径指向现有文件，该文件将被覆盖。
//   - 如果目录不存在，将创建该目录。
func (i *ICoreWebView2DownloadStartingEventArgs) PutResultFilePath(resultFilePath string) error {
	_path, err := syscall.UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHandled 获取是否已处理.
func (i *ICoreWebView2DownloadStartingEventArgs) GetHandled() (bool, error) {
	var handled bool
	r, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled, nil
}

// MustGetHandled 获取是否已处理。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustGetHandled() bool {
	result, err := i.GetHandled()
	ReportErrorAtuo(err)
	return result
}

// PutHandled 设置是否已处理.
//   - 可以将此标志设置为 TRUE，以隐藏此下载的默认下载对话框。
//   - 如果下载未取消，它将正常进行，只是不会显示默认用户界面。
//   - 默认情况下，该值为 FALSE，并显示默认下载对话框。
func (i *ICoreWebView2DownloadStartingEventArgs) PutHandled(handled bool) error {
	r, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取延迟对象.
//   - 使用此操作可在稍后时间完成该事件。
func (i *ICoreWebView2DownloadStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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

// MustGetDeferral 获取延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadStartingEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	result, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return result
}
