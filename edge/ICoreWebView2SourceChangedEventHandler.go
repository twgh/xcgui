package edge

import "unsafe"

type _ICoreWebView2SourceChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2SourceChangedEventHandler 源改变事件.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventhandler
type ICoreWebView2SourceChangedEventHandler struct {
	vtbl *_ICoreWebView2SourceChangedEventHandlerVtbl
	impl _ICoreWebView2SourceChangedEventHandlerImpl
}

func (i *ICoreWebView2SourceChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SourceChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2SourceChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2SourceChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2SourceChangedEventHandlerIUnknownAddRef(this *ICoreWebView2SourceChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2SourceChangedEventHandlerIUnknownRelease(this *ICoreWebView2SourceChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2SourceChangedEventHandlerInvoke(this *ICoreWebView2SourceChangedEventHandler, sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	return this.impl.SourceChanged(sender, args)
}

type _ICoreWebView2SourceChangedEventHandlerImpl interface {
	IUnknownImpl
	SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr
}

var _ICoreWebView2SourceChangedEventHandlerFn = _ICoreWebView2SourceChangedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2SourceChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2SourceChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2SourceChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2SourceChangedEventHandlerInvoke),
}

func NewICoreWebView2SourceChangedEventHandler(impl _ICoreWebView2SourceChangedEventHandlerImpl) *ICoreWebView2SourceChangedEventHandler {
	return &ICoreWebView2SourceChangedEventHandler{
		vtbl: &_ICoreWebView2SourceChangedEventHandlerFn,
		impl: impl,
	}
}
