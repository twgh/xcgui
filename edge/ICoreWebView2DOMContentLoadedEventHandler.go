package edge

import "unsafe"

type _ICoreWebView2DOMContentLoadedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DOMContentLoadedEventHandler 接收 DOMContentLoaded 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventhandler
type ICoreWebView2DOMContentLoadedEventHandler struct {
	vtbl *_ICoreWebView2DOMContentLoadedEventHandlerVtbl
	impl _ICoreWebView2DOMContentLoadedEventHandlerImpl
}

func (i *ICoreWebView2DOMContentLoadedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DOMContentLoadedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DOMContentLoadedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownAddRef(this *ICoreWebView2DOMContentLoadedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownRelease(this *ICoreWebView2DOMContentLoadedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DOMContentLoadedEventHandlerInvoke(this *ICoreWebView2DOMContentLoadedEventHandler, sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return this.impl.DOMContentLoaded(sender, args)
}

type _ICoreWebView2DOMContentLoadedEventHandlerImpl interface {
	IUnknownImpl
	DOMContentLoaded(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr
}

var _ICoreWebView2DOMContentLoadedEventHandlerFn = _ICoreWebView2DOMContentLoadedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerInvoke),
}

func NewICoreWebView2DOMContentLoadedEventHandler(impl _ICoreWebView2DOMContentLoadedEventHandlerImpl) *ICoreWebView2DOMContentLoadedEventHandler {
	return &ICoreWebView2DOMContentLoadedEventHandler{
		vtbl: &_ICoreWebView2DOMContentLoadedEventHandlerFn,
		impl: impl,
	}
}
