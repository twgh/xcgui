package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2NotificationReceivedEventArgs 提供有关接收到的通知的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2notificationreceivedeventargs
type ICoreWebView2NotificationReceivedEventArgs struct {
	Vtbl *_ICoreWebView2NotificationReceivedEventArgsVtbl
}

type _ICoreWebView2NotificationReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetSenderOrigin ComProc
	GetNotification ComProc
	PutHandled      ComProc
	GetHandled      ComProc
	GetDeferral     ComProc
}

func (i *ICoreWebView2NotificationReceivedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NotificationReceivedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NotificationReceivedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSenderOrigin 获取发送通知的网页内容的来源，例如 https://example.com/ 或 https://www.example.com/。
func (i *ICoreWebView2NotificationReceivedEventArgs) GetSenderOrigin() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetSenderOrigin.Call(
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

// GetNotification 获取通知对象。
//   - 你可以访问 Notification 对象的属性来显示自己的通知。
func (i *ICoreWebView2NotificationReceivedEventArgs) GetNotification() (*ICoreWebView2Notification, error) {
	var notification *ICoreWebView2Notification
	r, _, _ := i.Vtbl.GetNotification.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&notification)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return notification, nil
}

// SetHandled 设置通知是否已被处理。
func (i *ICoreWebView2NotificationReceivedEventArgs) SetHandled(value bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHandled 获取通知是否已被处理。
func (i *ICoreWebView2NotificationReceivedEventArgs) GetHandled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetDeferral 获取延迟对象，用于延迟完成事件处理。
func (i *ICoreWebView2NotificationReceivedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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

// MustGetSenderOrigin 获取发送通知的网页内容的来源，例如 https://example.com/ 或 https://www.example.com/。出错时会触发全局错误回调。
func (i *ICoreWebView2NotificationReceivedEventArgs) MustGetSenderOrigin() string {
	value, err := i.GetSenderOrigin()
	ReportErrorAtuo(err)
	return value
}

// MustGetNotification 获取通知对象。出错时会触发全局错误回调。
//   - 你可以访问 Notification 对象的属性来显示自己的通知。
func (i *ICoreWebView2NotificationReceivedEventArgs) MustGetNotification() *ICoreWebView2Notification {
	notification, err := i.GetNotification()
	ReportErrorAtuo(err)
	return notification
}

// MustGetHandled 获取通知是否已被处理。出错时会触发全局错误回调。
func (i *ICoreWebView2NotificationReceivedEventArgs) MustGetHandled() bool {
	value, err := i.GetHandled()
	ReportErrorAtuo(err)
	return value
}

// MustGetDeferral 获取延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2NotificationReceivedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}
