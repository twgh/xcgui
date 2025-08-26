package edge

import "unsafe"

type _ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2BasicAuthenticationRequestedEventHandler 接收基本身份验证请求事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventhandler
type ICoreWebView2BasicAuthenticationRequestedEventHandler struct {
	vtbl *_ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl
	impl _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BasicAuthenticationRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2BasicAuthenticationRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownRelease(this *ICoreWebView2BasicAuthenticationRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerInvoke(this *ICoreWebView2BasicAuthenticationRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr {
	return this.impl.BasicAuthenticationRequested(sender, args)
}

type _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl interface {
	IUnknownImpl
	BasicAuthenticationRequested(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr
}

var _ICoreWebView2BasicAuthenticationRequestedEventHandlerFn *_ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl

func NewICoreWebView2BasicAuthenticationRequestedEventHandler(impl _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl) *ICoreWebView2BasicAuthenticationRequestedEventHandler {
	if _ICoreWebView2BasicAuthenticationRequestedEventHandlerFn == nil {
		_ICoreWebView2BasicAuthenticationRequestedEventHandlerFn = &_ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2BasicAuthenticationRequestedEventHandler{
		vtbl: _ICoreWebView2BasicAuthenticationRequestedEventHandlerFn,
		impl: impl,
	}
}