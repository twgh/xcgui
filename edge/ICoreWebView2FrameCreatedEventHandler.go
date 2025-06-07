package edge

import "unsafe"

type _ICoreWebView2FrameCreatedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameCreatedEventHandler 接收框架创建事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventhandler
type ICoreWebView2FrameCreatedEventHandler struct {
	vtbl *_ICoreWebView2FrameCreatedEventHandlerVtbl
	impl _ICoreWebView2FrameCreatedEventHandlerImpl
}

func (i *ICoreWebView2FrameCreatedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameCreatedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameCreatedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameCreatedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameCreatedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameCreatedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameCreatedEventHandlerIUnknownRelease(this *ICoreWebView2FrameCreatedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameCreatedEventHandlerInvoke(this *ICoreWebView2FrameCreatedEventHandler, sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr {
	return this.impl.FrameCreated(sender, args)
}

type _ICoreWebView2FrameCreatedEventHandlerImpl interface {
	IUnknownImpl
	FrameCreated(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr
}

var _ICoreWebView2FrameCreatedEventHandlerFn = _ICoreWebView2FrameCreatedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameCreatedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameCreatedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameCreatedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameCreatedEventHandlerInvoke),
}

func NewICoreWebView2FrameCreatedEventHandler(impl _ICoreWebView2FrameCreatedEventHandlerImpl) *ICoreWebView2FrameCreatedEventHandler {
	return &ICoreWebView2FrameCreatedEventHandler{
		vtbl: &_ICoreWebView2FrameCreatedEventHandlerFn,
		impl: impl,
	}
}
