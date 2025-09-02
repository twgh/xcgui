package edge

import "unsafe"

type _ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameDOMContentLoadedEventHandler 接收框架 DOM 内容加载完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framedomcontentloadedeventhandler
type ICoreWebView2FrameDOMContentLoadedEventHandler struct {
	vtbl *_ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl
	impl _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl
}

func (i *ICoreWebView2FrameDOMContentLoadedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameDOMContentLoadedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameDOMContentLoadedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameDOMContentLoadedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownRelease(this *ICoreWebView2FrameDOMContentLoadedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerInvoke(this *ICoreWebView2FrameDOMContentLoadedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return this.impl.FrameDOMContentLoaded(sender, args)
}

type _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl interface {
	IUnknownImpl
	FrameDOMContentLoaded(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr
}

var _ICoreWebView2FrameDOMContentLoadedEventHandlerFn *_ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl

func NewICoreWebView2FrameDOMContentLoadedEventHandler(impl _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl) *ICoreWebView2FrameDOMContentLoadedEventHandler {
	if _ICoreWebView2FrameDOMContentLoadedEventHandlerFn == nil {
		_ICoreWebView2FrameDOMContentLoadedEventHandlerFn = &_ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameDOMContentLoadedEventHandler{
		vtbl: _ICoreWebView2FrameDOMContentLoadedEventHandlerFn,
		impl: impl,
	}
}