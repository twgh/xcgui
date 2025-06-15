package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2GetCookiesCompletedHandler 接收 ICoreWebView2CookieManager.GetCookies 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2getcookiescompletedhandler
type ICoreWebView2GetCookiesCompletedHandler struct {
	vtbl *_ICoreWebView2GetCookiesCompletedHandlerVtbl
	impl _ICoreWebView2GetCookiesCompletedHandlerImpl
}

func (i *ICoreWebView2GetCookiesCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2GetCookiesCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

type _ICoreWebView2GetCookiesCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

func _ICoreWebView2GetCookiesCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetCookiesCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2GetCookiesCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetCookiesCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2GetCookiesCompletedHandlerIUnknownRelease(this *ICoreWebView2GetCookiesCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2GetCookiesCompletedHandlerInvoke(this *ICoreWebView2GetCookiesCompletedHandler, errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr {
	return this.impl.GetCookiesCompleted(errorCode, cookies)
}

type _ICoreWebView2GetCookiesCompletedHandlerImpl interface {
	IUnknownImpl
	GetCookiesCompleted(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr
}

var _ICoreWebView2GetCookiesCompletedHandlerFn *_ICoreWebView2GetCookiesCompletedHandlerVtbl

func NewICoreWebView2GetCookiesCompletedHandler(impl _ICoreWebView2GetCookiesCompletedHandlerImpl) *ICoreWebView2GetCookiesCompletedHandler {
	if _ICoreWebView2GetCookiesCompletedHandlerFn == nil {
		_ICoreWebView2GetCookiesCompletedHandlerFn = &_ICoreWebView2GetCookiesCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2GetCookiesCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2GetCookiesCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2GetCookiesCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2GetCookiesCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2GetCookiesCompletedHandler{
		vtbl: _ICoreWebView2GetCookiesCompletedHandlerFn,
		impl: impl,
	}
}
