package edge

import (
	"unsafe"
)

type _ICoreWebView2NotificationCloseRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NotificationCloseRequestedEventHandler 接收通知关闭请求事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2notificationcloserequestedeventhandler
type ICoreWebView2NotificationCloseRequestedEventHandler struct {
	vtbl *_ICoreWebView2NotificationCloseRequestedEventHandlerVtbl
	impl _ICoreWebView2NotificationCloseRequestedEventHandlerImpl
}

func (i *ICoreWebView2NotificationCloseRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NotificationCloseRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2NotificationCloseRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2NotificationCloseRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownRelease(this *ICoreWebView2NotificationCloseRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NotificationCloseRequestedEventHandlerInvoke(this *ICoreWebView2NotificationCloseRequestedEventHandler, sender *ICoreWebView2Notification, args *IUnknown) uintptr {
	return this.impl.NotificationCloseRequested(sender, args)
}

type _ICoreWebView2NotificationCloseRequestedEventHandlerImpl interface {
	IUnknownImpl
	NotificationCloseRequested(sender *ICoreWebView2Notification, args *IUnknown) uintptr
}

var _ICoreWebView2NotificationCloseRequestedEventHandlerFn *_ICoreWebView2NotificationCloseRequestedEventHandlerVtbl

func NewICoreWebView2NotificationCloseRequestedEventHandler(impl _ICoreWebView2NotificationCloseRequestedEventHandlerImpl) *ICoreWebView2NotificationCloseRequestedEventHandler {
	if _ICoreWebView2NotificationCloseRequestedEventHandlerFn == nil {
		_ICoreWebView2NotificationCloseRequestedEventHandlerFn = &_ICoreWebView2NotificationCloseRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2NotificationCloseRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2NotificationCloseRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2NotificationCloseRequestedEventHandler{
		vtbl: _ICoreWebView2NotificationCloseRequestedEventHandlerFn,
		impl: impl,
	}
}
