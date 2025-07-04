package edge

import "unsafe"

type _ICoreWebView2WindowCloseRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2WindowCloseRequestedEventHandler 接收 WindowCloseRequested 事件.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2windowcloserequestedeventhandler
type ICoreWebView2WindowCloseRequestedEventHandler struct {
	vtbl *_ICoreWebView2WindowCloseRequestedEventHandlerVtbl
	impl _ICoreWebView2WindowCloseRequestedEventHandlerImpl
}

func (i *ICoreWebView2WindowCloseRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WindowCloseRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WindowCloseRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WindowCloseRequestedEventHandlerInvoke(this *ICoreWebView2WindowCloseRequestedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.WindowCloseRequested(sender, args)
}

type _ICoreWebView2WindowCloseRequestedEventHandlerImpl interface {
	IUnknownImpl
	WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2WindowCloseRequestedEventHandlerFn *_ICoreWebView2WindowCloseRequestedEventHandlerVtbl

func NewICoreWebView2WindowCloseRequestedEventHandler(impl _ICoreWebView2WindowCloseRequestedEventHandlerImpl) *ICoreWebView2WindowCloseRequestedEventHandler {
	if _ICoreWebView2WindowCloseRequestedEventHandlerFn == nil {
		_ICoreWebView2WindowCloseRequestedEventHandlerFn = &_ICoreWebView2WindowCloseRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2WindowCloseRequestedEventHandler{
		vtbl: _ICoreWebView2WindowCloseRequestedEventHandlerFn,
		impl: impl,
	}
}
