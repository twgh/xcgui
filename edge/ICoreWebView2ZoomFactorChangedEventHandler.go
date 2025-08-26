package edge

import "unsafe"

type _ICoreWebView2ZoomFactorChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ZoomFactorChangedEventHandler 接收缩放因子变化事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2zoomfactorchangedeventhandler
type ICoreWebView2ZoomFactorChangedEventHandler struct {
	vtbl *_ICoreWebView2ZoomFactorChangedEventHandlerVtbl
	impl _ICoreWebView2ZoomFactorChangedEventHandlerImpl
}

func (i *ICoreWebView2ZoomFactorChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ZoomFactorChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ZoomFactorChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease(this *ICoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ZoomFactorChangedEventHandlerInvoke(this *ICoreWebView2ZoomFactorChangedEventHandler, sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	return this.impl.ZoomFactorChanged(sender, args)
}

type _ICoreWebView2ZoomFactorChangedEventHandlerImpl interface {
	IUnknownImpl
	ZoomFactorChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr
}

var _ICoreWebView2ZoomFactorChangedEventHandlerFn *_ICoreWebView2ZoomFactorChangedEventHandlerVtbl

func NewICoreWebView2ZoomFactorChangedEventHandler(impl _ICoreWebView2ZoomFactorChangedEventHandlerImpl) *ICoreWebView2ZoomFactorChangedEventHandler {
	if _ICoreWebView2ZoomFactorChangedEventHandlerFn == nil {
		_ICoreWebView2ZoomFactorChangedEventHandlerFn = &_ICoreWebView2ZoomFactorChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ZoomFactorChangedEventHandler{
		vtbl: _ICoreWebView2ZoomFactorChangedEventHandlerFn,
		impl: impl,
	}
}