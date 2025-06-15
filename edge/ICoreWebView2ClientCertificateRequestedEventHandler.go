package edge

import "unsafe"

type _ICoreWebView2ClientCertificateRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ClientCertificateRequestedEventHandler 接收 ClientCertificateRequested 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventhandler
type ICoreWebView2ClientCertificateRequestedEventHandler struct {
	vtbl *_ICoreWebView2ClientCertificateRequestedEventHandlerVtbl
	impl _ICoreWebView2ClientCertificateRequestedEventHandlerImpl
}

func (i *ICoreWebView2ClientCertificateRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificateRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2ClientCertificateRequestedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownAddRef(
	this *ICoreWebView2ClientCertificateRequestedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownRelease(
	this *ICoreWebView2ClientCertificateRequestedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerInvoke(
	this *ICoreWebView2ClientCertificateRequestedEventHandler,
	sender *ICoreWebView2,
	args *ICoreWebView2ClientCertificateRequestedEventArgs,
) uintptr {
	return this.impl.ClientCertificateRequested(sender, args)
}

type _ICoreWebView2ClientCertificateRequestedEventHandlerImpl interface {
	IUnknownImpl
	ClientCertificateRequested(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr
}

var _ICoreWebView2ClientCertificateRequestedEventHandlerFn *_ICoreWebView2ClientCertificateRequestedEventHandlerVtbl

func NewICoreWebView2ClientCertificateRequestedEventHandler(
	impl _ICoreWebView2ClientCertificateRequestedEventHandlerImpl,
) *ICoreWebView2ClientCertificateRequestedEventHandler {
	if _ICoreWebView2ClientCertificateRequestedEventHandlerFn == nil {
		_ICoreWebView2ClientCertificateRequestedEventHandlerFn = &_ICoreWebView2ClientCertificateRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ClientCertificateRequestedEventHandler{
		vtbl: _ICoreWebView2ClientCertificateRequestedEventHandlerFn,
		impl: impl,
	}
}
