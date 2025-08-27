package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler 提供清除服务器证书错误操作完成后的回调方法。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clearservercertificateerroractionscompletedhandler
type ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler struct {
	vtbl *_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl
	impl _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl
}

func (i *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownAddRef(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownRelease(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerInvoke(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.ClearServerCertificateErrorActionsCompleted(errorCode)
}

type _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl interface {
	IUnknownImpl
	ClearServerCertificateErrorActionsCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn *_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl

func NewICoreWebView2ClearServerCertificateErrorActionsCompletedHandler(impl _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl) *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler {
	if _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn == nil {
		_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn = &_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler{
		vtbl: _ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn,
		impl: impl,
	}
}
