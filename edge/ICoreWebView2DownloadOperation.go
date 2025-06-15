package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2DownloadOperation 表示下载操作.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation
type ICoreWebView2DownloadOperation struct {
	Vtbl *ICoreWebView2DownloadOperationVtbl
}

type ICoreWebView2DownloadOperationVtbl struct {
	IUnknownVtbl
	AddBytesReceivedChanged       ComProc
	RemoveBytesReceivedChanged    ComProc
	AddEstimatedEndTimeChanged    ComProc
	RemoveEstimatedEndTimeChanged ComProc
	AddStateChanged               ComProc
	RemoveStateChanged            ComProc
	GetUri                        ComProc
	GetContentDisposition         ComProc
	GetMimeType                   ComProc
	GetTotalBytesToReceive        ComProc
	GetBytesReceived              ComProc
	GetEstimatedEndTime           ComProc
	GetResultFilePath             ComProc
	GetState                      ComProc
	GetInterruptReason            ComProc
	Cancel                        ComProc
	Pause                         ComProc
	Resume                        ComProc
	GetCanResume                  ComProc
}

func (i *ICoreWebView2DownloadOperation) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadOperation) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadOperation) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddBytesReceivedChanged 添加接收到的字节数改变事件处理程序.
func (i *ICoreWebView2DownloadOperation) AddBytesReceivedChanged(eventHandler *ICoreWebView2BytesReceivedChangedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveBytesReceivedChanged 移除下载字节数改变事件处理程序
func (i *ICoreWebView2DownloadOperation) RemoveBytesReceivedChanged(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取下载的 URI.
func (i *ICoreWebView2DownloadOperation) GetUri() (string, error) {
	var uri *uint16
	r, _, err := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(uri)
	wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return result, nil
}

// MustGetUri 获取下载的 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadOperation) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// Cancel 取消下载.
//   - 如果取消下载，默认下载对话框会显示下载已取消。
//   - 如果要取消下载且不显示默认下载对话框，宿主应设置 Cancel 中的 ICoreWebView2DownloadStartingEventArgs 属性。
func (i *ICoreWebView2DownloadOperation) Cancel() error {
	r, _, err := i.Vtbl.Cancel.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Pause 暂停下载.
//   - 如果处于暂停状态，默认下载对话框会显示下载已暂停。
//   - 如果下载已暂停，则无效果。
//   - 暂停下载会将状态更改为 COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED，并将 InterruptReason 设置为 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED.
func (i *ICoreWebView2DownloadOperation) Pause() error {
	r, _, err := i.Vtbl.Pause.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Resume 恢复下载.
//   - 如果 CanResume 返回 true，还可以恢复因其他原因中断的下载。
//   - 恢复下载会将状态从 COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED 更改为 COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS.
func (i *ICoreWebView2DownloadOperation) Resume() error {
	r, _, err := i.Vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetContentDisposition 获取下载的HTTP响应中的 Content-Disposition 标头值。
func (i *ICoreWebView2DownloadOperation) GetContentDisposition() (string, error) {
	var contentDisposition *uint16
	r, _, err := i.Vtbl.GetContentDisposition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&contentDisposition)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(contentDisposition)
	wapi.CoTaskMemFree(unsafe.Pointer(contentDisposition))
	return result, nil
}

// MustGetContentDisposition 获取下载的HTTP响应中的 Content-Disposition 标头值。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadOperation) MustGetContentDisposition() string {
	contentDisposition, err := i.GetContentDisposition()
	ReportErrorAtuo(err)
	return contentDisposition
}

// GetMimeType 获取下载内容的 MIME 类型.
func (i *ICoreWebView2DownloadOperation) GetMimeType() (string, error) {
	var mimeType *uint16
	r, _, err := i.Vtbl.GetMimeType.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&mimeType)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(mimeType)
	wapi.CoTaskMemFree(unsafe.Pointer(mimeType))
	return result, nil
}

// MustGetMimeType 获取下载内容的 MIME 类型。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadOperation) MustGetMimeType() string {
	mimeType, err := i.GetMimeType()
	ReportErrorAtuo(err)
	return mimeType
}

// GetTotalBytesToReceive 获取根据 HTTP Content-Length 标头计算出的下载内容总字节数的预期大小。
//   - 如果大小未知，则返回 -1。
func (i *ICoreWebView2DownloadOperation) GetTotalBytesToReceive() (int64, error) {
	var totalBytes int64
	r, _, err := i.Vtbl.GetTotalBytesToReceive.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&totalBytes)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return totalBytes, nil
}

// MustGetTotalBytesToReceive 获取根据 HTTP Content-Length 标头计算出的下载内容总字节数的预期大小。出错时会触发全局错误回调。
//   - 如果大小未知，则返回 -1。
func (i *ICoreWebView2DownloadOperation) MustGetTotalBytesToReceive() int64 {
	totalBytes, err := i.GetTotalBytesToReceive()
	ReportErrorAtuo(err)
	return totalBytes
}

// GetBytesReceived 获取已写入下载文件的字节数。
func (i *ICoreWebView2DownloadOperation) GetBytesReceived() (int64, error) {
	var bytesReceived int64
	r, _, err := i.Vtbl.GetBytesReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bytesReceived)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return bytesReceived, nil
}

// MustGetBytesReceived 获取已写入下载文件的字节数。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadOperation) MustGetBytesReceived() int64 {
	bytesReceived, err := i.GetBytesReceived()
	ReportErrorAtuo(err)
	return bytesReceived
}

// GetEstimatedEndTime 获取预计结束时间，采用 ISO 8601 日期和时间格式。
func (i *ICoreWebView2DownloadOperation) GetEstimatedEndTime() (string, error) {
	var estimatedEndTime *uint16
	r, _, err := i.Vtbl.GetEstimatedEndTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&estimatedEndTime)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(estimatedEndTime)
	wapi.CoTaskMemFree(unsafe.Pointer(estimatedEndTime))
	return result, nil
}

// MustGetEstimatedEndTime 获取预计结束时间，采用 ISO 8601 日期和时间格式。出错时会触发全局错误回调。
func (i *ICoreWebView2DownloadOperation) MustGetEstimatedEndTime() string {
	estimatedEndTime, err := i.GetEstimatedEndTime()
	ReportErrorAtuo(err)
	return estimatedEndTime
}

// GetResultFilePath 获取下载文件的绝对路径，包括文件名。
//   - 宿主可以从 ICoreWebView2DownloadStartingEventArgs 更改此设置。
func (i *ICoreWebView2DownloadOperation) GetResultFilePath() (string, error) {
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
	result := common.UTF16PtrToString(resultFilePath)
	wapi.CoTaskMemFree(unsafe.Pointer(resultFilePath))
	return result, nil
}

// MustGetResultFilePath 获取下载文件的绝对路径，包括文件名。出错时会触发全局错误回调。
//   - 宿主可以从 ICoreWebView2DownloadStartingEventArgs 更改此设置。
func (i *ICoreWebView2DownloadOperation) MustGetResultFilePath() string {
	resultFilePath, err := i.GetResultFilePath()
	ReportErrorAtuo(err)
	return resultFilePath
}

// GetState 获取下载状态。
func (i *ICoreWebView2DownloadOperation) GetState() (COREWEBVIEW2_DOWNLOAD_STATE, error) {
	var state COREWEBVIEW2_DOWNLOAD_STATE
	r, _, err := i.Vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&state)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return state, nil
}

// MustGetState 获取下载状态。出错时会触发全局错误回调.
func (i *ICoreWebView2DownloadOperation) MustGetState() COREWEBVIEW2_DOWNLOAD_STATE {
	state, err := i.GetState()
	ReportErrorAtuo(err)
	return state
}

// GetInterruptReason 获取下载中断的原因.
func (i *ICoreWebView2DownloadOperation) GetInterruptReason() (COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON, error) {
	var reason COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON
	r, _, err := i.Vtbl.GetInterruptReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return reason, nil
}

// MustGetInterruptReason 获取下载中断的原因。出错时会触发全局错误回调.
func (i *ICoreWebView2DownloadOperation) MustGetInterruptReason() COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON {
	reason, err := i.GetInterruptReason()
	ReportErrorAtuo(err)
	return reason
}

// GetCanResume 获取下载是否可以恢复.
//   - 以下中断原因导致的下载可能无需你调用任何方法即可自动恢复: COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE、 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH、 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT。在这些情况下，下载进度可能会重新开始，且 BytesReceived 将重置为0。
func (i *ICoreWebView2DownloadOperation) GetCanResume() (bool, error) {
	var canResume bool
	r, _, err := i.Vtbl.GetCanResume.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canResume)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return canResume, nil
}

// MustGetCanResume 获取下载是否可以恢复。出错时会触发全局错误回调.
//   - 以下中断原因导致的下载可能无需你调用任何方法即可自动恢复: COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE、 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH、 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT。在这些情况下，下载进度可能会重新开始，且 BytesReceived 将重置为0。
func (i *ICoreWebView2DownloadOperation) MustGetCanResume() bool {
	canResume, err := i.GetCanResume()
	ReportErrorAtuo(err)
	return canResume
}

// Event_BytesReceivedChanged 下载字节改变事件.
//   - 当下载的字节数发生更改时触发。
func (i *ICoreWebView2DownloadOperation) Event_BytesReceivedChanged(w *WebViewEventImpl, cb func(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(w, "BytesReceivedChanged", cb, i, allowAddingMultiple...)
}

/*TODO:
AddEstimatedEndTimeChanged
RemoveEstimatedEndTimeChanged
AddStateChanged
RemoveStateChanged
*/
