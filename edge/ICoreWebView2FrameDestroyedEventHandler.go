package edge

import "unsafe"

type _ICoreWebView2FrameDestroyedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameDestroyedEventHandler 接收框架销毁事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framedestroyedeventhandler
type ICoreWebView2FrameDestroyedEventHandler struct {
	vtbl *_ICoreWebView2FrameDestroyedEventHandlerVtbl
	impl _ICoreWebView2FrameDestroyedEventHandlerImpl
}

func (i *ICoreWebView2FrameDestroyedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameDestroyedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameDestroyedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameDestroyedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameDestroyedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameDestroyedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameDestroyedEventHandlerIUnknownRelease(this *ICoreWebView2FrameDestroyedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameDestroyedEventHandlerInvoke(this *ICoreWebView2FrameDestroyedEventHandler, sender *ICoreWebView2Frame, args *IUnknown) uintptr {
	return this.impl.FrameDestroyed(sender, args)
}

type _ICoreWebView2FrameDestroyedEventHandlerImpl interface {
	IUnknownImpl
	FrameDestroyed(sender *ICoreWebView2Frame, args *IUnknown) uintptr
}

var _ICoreWebView2FrameDestroyedEventHandlerFn = _ICoreWebView2FrameDestroyedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameDestroyedEventHandlerInvoke),
}

func NewICoreWebView2FrameDestroyedEventHandler(impl _ICoreWebView2FrameDestroyedEventHandlerImpl) *ICoreWebView2FrameDestroyedEventHandler {
	return &ICoreWebView2FrameDestroyedEventHandler{
		vtbl: &_ICoreWebView2FrameDestroyedEventHandlerFn,
		impl: impl,
	}
}
