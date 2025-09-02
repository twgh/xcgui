package edge

import "unsafe"

type _ICoreWebView2FrameChildFrameCreatedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameChildFrameCreatedEventHandler 接收子框架创建事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framechildframecreatedeventhandler
type ICoreWebView2FrameChildFrameCreatedEventHandler struct {
	vtbl *_ICoreWebView2FrameChildFrameCreatedEventHandlerVtbl
	impl _ICoreWebView2FrameChildFrameCreatedEventHandlerImpl
}

func (i *ICoreWebView2FrameChildFrameCreatedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameChildFrameCreatedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameChildFrameCreatedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameChildFrameCreatedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownRelease(this *ICoreWebView2FrameChildFrameCreatedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameChildFrameCreatedEventHandlerInvoke(this *ICoreWebView2FrameChildFrameCreatedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2FrameCreatedEventArgs) uintptr {
	return this.impl.FrameChildFrameCreated(sender, args)
}

type _ICoreWebView2FrameChildFrameCreatedEventHandlerImpl interface {
	IUnknownImpl
	FrameChildFrameCreated(sender *ICoreWebView2Frame, args *ICoreWebView2FrameCreatedEventArgs) uintptr
}

var _ICoreWebView2FrameChildFrameCreatedEventHandlerFn *_ICoreWebView2FrameChildFrameCreatedEventHandlerVtbl

func NewICoreWebView2FrameChildFrameCreatedEventHandler(impl _ICoreWebView2FrameChildFrameCreatedEventHandlerImpl) *ICoreWebView2FrameChildFrameCreatedEventHandler {
	if _ICoreWebView2FrameChildFrameCreatedEventHandlerFn == nil {
		_ICoreWebView2FrameChildFrameCreatedEventHandlerFn = &_ICoreWebView2FrameChildFrameCreatedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameChildFrameCreatedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameChildFrameCreatedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameChildFrameCreatedEventHandler{
		vtbl: _ICoreWebView2FrameChildFrameCreatedEventHandlerFn,
		impl: impl,
	}
}
