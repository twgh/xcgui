package edge

import "unsafe"

type _ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ServerCertificateErrorDetectedEventHandler 接收检测到服务器证书错误事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventhandler
type ICoreWebView2ServerCertificateErrorDetectedEventHandler struct {
	vtbl *_ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl
	impl _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownAddRef(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownRelease(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerInvoke(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler, sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr {
	return this.impl.ServerCertificateErrorDetected(sender, args)
}

type _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl interface {
	IUnknownImpl
	ServerCertificateErrorDetected(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr
}

var _ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn *_ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl

func NewICoreWebView2ServerCertificateErrorDetectedEventHandler(impl _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl) *ICoreWebView2ServerCertificateErrorDetectedEventHandler {
	if _ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn == nil {
		_ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn = &_ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ServerCertificateErrorDetectedEventHandler{
		vtbl: _ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn,
		impl: impl,
	}
}
