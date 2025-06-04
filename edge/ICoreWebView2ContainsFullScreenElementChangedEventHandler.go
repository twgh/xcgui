package edge

import "unsafe"

type _ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ContainsFullScreenElementChangedEventHandler 接收 ContainsFullScreenElementChanged 事件
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2containsfullscreenelementchangedeventhandler
type ICoreWebView2ContainsFullScreenElementChangedEventHandler struct {
	vtbl *_ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl
	impl _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownRelease(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerInvoke(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.ContainsFullScreenElementChanged(sender, args)
}

type _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl interface {
	IUnknownImpl
	ContainsFullScreenElementChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2ContainsFullScreenElementChangedEventHandlerFn = _ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerInvoke),
}

func NewICoreWebView2ContainsFullScreenElementChangedEventHandler(impl _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl) *ICoreWebView2ContainsFullScreenElementChangedEventHandler {
	return &ICoreWebView2ContainsFullScreenElementChangedEventHandler{
		vtbl: &_ICoreWebView2ContainsFullScreenElementChangedEventHandlerFn,
		impl: impl,
	}
}
