package edge

import (
	"unsafe"
)

type _ICoreWebView2NotificationReceivedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NotificationReceivedEventHandler 接收通知接收事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2notificationreceivedeventhandler
type ICoreWebView2NotificationReceivedEventHandler struct {
	vtbl *_ICoreWebView2NotificationReceivedEventHandlerVtbl
	impl _ICoreWebView2NotificationReceivedEventHandlerImpl
}

func (i *ICoreWebView2NotificationReceivedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NotificationReceivedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NotificationReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2NotificationReceivedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NotificationReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2NotificationReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NotificationReceivedEventHandlerIUnknownRelease(this *ICoreWebView2NotificationReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NotificationReceivedEventHandlerInvoke(this *ICoreWebView2NotificationReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NotificationReceivedEventArgs) uintptr {
	return this.impl.NotificationReceived(sender, args)
}

type _ICoreWebView2NotificationReceivedEventHandlerImpl interface {
	IUnknownImpl
	NotificationReceived(sender *ICoreWebView2, args *ICoreWebView2NotificationReceivedEventArgs) uintptr
}

var _ICoreWebView2NotificationReceivedEventHandlerFn *_ICoreWebView2NotificationReceivedEventHandlerVtbl

func NewICoreWebView2NotificationReceivedEventHandler(impl _ICoreWebView2NotificationReceivedEventHandlerImpl) *ICoreWebView2NotificationReceivedEventHandler {
	if _ICoreWebView2NotificationReceivedEventHandlerFn == nil {
		_ICoreWebView2NotificationReceivedEventHandlerFn = &_ICoreWebView2NotificationReceivedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2NotificationReceivedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2NotificationReceivedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2NotificationReceivedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2NotificationReceivedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2NotificationReceivedEventHandler{
		vtbl: _ICoreWebView2NotificationReceivedEventHandlerFn,
		impl: impl,
	}
}
