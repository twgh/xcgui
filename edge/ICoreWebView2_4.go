package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_4 是 ICoreWebView2_3 接口的延续，以支持 FrameCreated 和 下载开始 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_4
type ICoreWebView2_4 struct {
	Vtbl *ICoreWebView2_4Vtbl
}

type ICoreWebView2_4Vtbl struct {
	ICoreWebView2_3Vtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

func (i *ICoreWebView2_4) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_4) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_4) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddFrameCreated 添加框架创建完成事件处理程序。
//   - 当创建新的 iframe 时触发。
//   - 处理此事件以访问 ICoreWebView2Frame 对象。
func (i *ICoreWebView2_4) AddFrameCreated(eventHandler *ICoreWebView2FrameCreatedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddFrameCreated.Call(
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

// RemoveFrameCreated 移除框架创建事件处理程序。
func (i *ICoreWebView2_4) RemoveFrameCreated(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveFrameCreated.Call(
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

// AddDownloadStarting 添加下载开始事件处理程序。
//   - 当下载开始时触发，该事件会阻止默认的下载对话框弹出，但不会阻止下载进程。
//   - 主机可以选择取消下载、更改结果文件路径以及隐藏默认下载对话框。
//   - 如果主机选择取消下载，则不会保存下载内容，不会显示对话框，并且状态将更改为 COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED，中断原因是 COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED.
//   - 否则，事件完成后，下载内容将保存到默认路径，如果主机未选择隐藏默认下载对话框，则会显示该对话框。
//   - 主机可以使用 Handled 属性更改下载对话框的可见性。
//   - 如果未处理该事件，下载将正常完成并显示默认对话框。
func (i *ICoreWebView2_4) AddDownloadStarting(eventHandler *ICoreWebView2DownloadStartingEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddDownloadStarting.Call(
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

// RemoveDownloadStarting 移除下载开始事件处理程序。
func (i *ICoreWebView2_4) RemoveDownloadStarting(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveDownloadStarting.Call(
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
