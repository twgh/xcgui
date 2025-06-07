package edge

import "unsafe"

type _ICoreWebView2WebResourceRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2WebResourceRequestedEventHandler 接收 WebResourceRequested 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventhandler
type ICoreWebView2WebResourceRequestedEventHandler struct {
	vtbl *_ICoreWebView2WebResourceRequestedEventHandlerVtbl
	impl _ICoreWebView2WebResourceRequestedEventHandlerImpl
}

func (i *ICoreWebView2WebResourceRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebResourceRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2WebResourceRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownRelease(this *ICoreWebView2WebResourceRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebResourceRequestedEventHandlerInvoke(this *ICoreWebView2WebResourceRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	return this.impl.WebResourceRequested(sender, args)
}

type _ICoreWebView2WebResourceRequestedEventHandlerImpl interface {
	IUnknownImpl
	WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr
}

var _ICoreWebView2WebResourceRequestedEventHandlerFn = _ICoreWebView2WebResourceRequestedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerInvoke),
}

func NewICoreWebView2WebResourceRequestedEventHandler(impl _ICoreWebView2WebResourceRequestedEventHandlerImpl) *ICoreWebView2WebResourceRequestedEventHandler {
	return &ICoreWebView2WebResourceRequestedEventHandler{
		vtbl: &_ICoreWebView2WebResourceRequestedEventHandlerFn,
		impl: impl,
	}
}
