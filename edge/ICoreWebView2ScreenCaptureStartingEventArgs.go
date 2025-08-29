package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2ScreenCaptureStartingEventArgs 是屏幕截图开始事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2screencapturestartingeventargs
type ICoreWebView2ScreenCaptureStartingEventArgs struct {
	Vtbl *ICoreWebView2ScreenCaptureStartingEventArgsVtbl
}

type ICoreWebView2ScreenCaptureStartingEventArgsVtbl struct {
	IUnknownVtbl
	GetCancel                  ComProc
	PutCancel                  ComProc
	GetHandled                 ComProc
	PutHandled                 ComProc
	GetOriginalSourceFrameInfo ComProc
	GetDeferral                ComProc
}

func (i *ICoreWebView2ScreenCaptureStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScreenCaptureStartingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScreenCaptureStartingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCancel 获取是否取消屏幕截图。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) GetCancel() (bool, error) {
	var cancel bool
	r, _, _ := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return cancel, nil
}

// SetCancel 设置是否取消屏幕截图。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) SetCancel(cancel bool) error {
	r, _, _ := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(cancel),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHandled 获取是否已处理屏幕截图事件。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) GetHandled() (bool, error) {
	var handled bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled, nil
}

// SetHandled 设置是否已处理屏幕截图事件。
//   - 默认情况下，ScreenCaptureStarting 事件处理程序会在 ICoreWebView2Frame 和 ICoreWebView2 上均被调用，其中 ICoreWebView2Frame 的事件处理程序会先被调用。
//   - 可以在 ICoreWebView2Frame 事件处理程序中将此标志设置为 TRUE，以阻止调用其余的 ICoreWebView2 事件处理程序。
//   - 如果在 ICoreWebView2Frame 事件处理程序中将该标志设置为 FALSE，则下游处理程序可以更新 Cancel 属性。
//   - 如果对事件参数进行了延迟处理，则必须在进行延迟处理之前同步将 Handled 设置为 TRUE，以防止调用 ICoreWebView2 的事件处理程序。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) SetHandled(handled bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetOriginalSourceFrameInfo 获取请求屏幕捕获权限的相关框架信息。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) GetOriginalSourceFrameInfo() (*ICoreWebView2FrameInfo, error) {
	var frameInfo *ICoreWebView2FrameInfo
	r, _, _ := i.Vtbl.GetOriginalSourceFrameInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameInfo)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return frameInfo, nil
}

// GetDeferral 获取延迟对象。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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

// MustGetCancel 获取是否取消屏幕截图。出错时会触发全局错误回调。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) MustGetCancel() bool {
	result, err := i.GetCancel()
	ReportErrorAtuo(err)
	return result
}

// MustGetHandled 获取是否已处理屏幕截图事件。出错时会触发全局错误回调。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) MustGetHandled() bool {
	result, err := i.GetHandled()
	ReportErrorAtuo(err)
	return result
}

// MustGetOriginalSourceFrameInfo 获取请求屏幕捕获权限的相关框架信息。出错时会触发全局错误回调。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) MustGetOriginalSourceFrameInfo() *ICoreWebView2FrameInfo {
	result, err := i.GetOriginalSourceFrameInfo()
	ReportErrorAtuo(err)
	return result
}

// MustGetDeferral 获取延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2ScreenCaptureStartingEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	result, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return result
}
