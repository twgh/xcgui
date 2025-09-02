package edge

import "unsafe"

type _ICoreWebView2FrameContentLoadingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameContentLoadingEventHandler 接收框架内容加载事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framecontentloadingeventhandler
type ICoreWebView2FrameContentLoadingEventHandler struct {
	vtbl *_ICoreWebView2FrameContentLoadingEventHandlerVtbl
	impl _ICoreWebView2FrameContentLoadingEventHandlerImpl
}

func (i *ICoreWebView2FrameContentLoadingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameContentLoadingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameContentLoadingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownAddRef(this *ICoreWebView2FrameContentLoadingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownRelease(this *ICoreWebView2FrameContentLoadingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameContentLoadingEventHandlerInvoke(this *ICoreWebView2FrameContentLoadingEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	return this.impl.FrameContentLoading(sender, args)
}

type _ICoreWebView2FrameContentLoadingEventHandlerImpl interface {
	IUnknownImpl
	FrameContentLoading(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr
}

var _ICoreWebView2FrameContentLoadingEventHandlerFn *_ICoreWebView2FrameContentLoadingEventHandlerVtbl

func NewICoreWebView2FrameContentLoadingEventHandler(impl _ICoreWebView2FrameContentLoadingEventHandlerImpl) *ICoreWebView2FrameContentLoadingEventHandler {
	if _ICoreWebView2FrameContentLoadingEventHandlerFn == nil {
		_ICoreWebView2FrameContentLoadingEventHandlerFn = &_ICoreWebView2FrameContentLoadingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameContentLoadingEventHandler{
		vtbl: _ICoreWebView2FrameContentLoadingEventHandlerFn,
		impl: impl,
	}
}