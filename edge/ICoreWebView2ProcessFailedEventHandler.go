package edge

import "unsafe"

type _ICoreWebView2ProcessFailedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ProcessFailedEventHandler 接收 ProcessFailed 事件.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventhandler
type ICoreWebView2ProcessFailedEventHandler struct {
	vtbl *_ICoreWebView2ProcessFailedEventHandlerVtbl
	impl _ICoreWebView2ProcessFailedEventHandlerImpl
}

func (i *ICoreWebView2ProcessFailedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessFailedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ProcessFailedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ProcessFailedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProcessFailedEventHandlerIUnknownAddRef(this *ICoreWebView2ProcessFailedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProcessFailedEventHandlerIUnknownRelease(this *ICoreWebView2ProcessFailedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProcessFailedEventHandlerInvoke(this *ICoreWebView2ProcessFailedEventHandler, sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr {
	return this.impl.ProcessFailed(sender, args)
}

type _ICoreWebView2ProcessFailedEventHandlerImpl interface {
	IUnknownImpl
	ProcessFailed(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr
}

var _ICoreWebView2ProcessFailedEventHandlerFn = _ICoreWebView2ProcessFailedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2ProcessFailedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ProcessFailedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ProcessFailedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ProcessFailedEventHandlerInvoke),
}

func NewICoreWebView2ProcessFailedEventHandler(impl _ICoreWebView2ProcessFailedEventHandlerImpl) *ICoreWebView2ProcessFailedEventHandler {
	return &ICoreWebView2ProcessFailedEventHandler{
		vtbl: &_ICoreWebView2ProcessFailedEventHandlerFn,
		impl: impl,
	}
}
