package edge

import "unsafe"

type _ICoreWebView2NonClientRegionChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NonClientRegionChangedEventHandler 接收非客户端区域变更事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventhandler
type ICoreWebView2NonClientRegionChangedEventHandler struct {
	vtbl *_ICoreWebView2NonClientRegionChangedEventHandlerVtbl
	impl _ICoreWebView2NonClientRegionChangedEventHandlerImpl
}

func (i *ICoreWebView2NonClientRegionChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NonClientRegionChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NonClientRegionChangedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2NonClientRegionChangedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NonClientRegionChangedEventHandlerIUnknownAddRef(
	this *ICoreWebView2NonClientRegionChangedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NonClientRegionChangedEventHandlerIUnknownRelease(
	this *ICoreWebView2NonClientRegionChangedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NonClientRegionChangedEventHandlerInvoke(
	this *ICoreWebView2NonClientRegionChangedEventHandler,
	sender *ICoreWebView2CompositionController,
	args *ICoreWebView2NonClientRegionChangedEventArgs,
) uintptr {
	return this.impl.NonClientRegionChanged(sender, args)
}

type _ICoreWebView2NonClientRegionChangedEventHandlerImpl interface {
	IUnknownImpl
	NonClientRegionChanged(sender *ICoreWebView2CompositionController, args *ICoreWebView2NonClientRegionChangedEventArgs) uintptr
}

var _ICoreWebView2NonClientRegionChangedEventHandlerFn *_ICoreWebView2NonClientRegionChangedEventHandlerVtbl

func NewICoreWebView2NonClientRegionChangedEventHandler(
	impl _ICoreWebView2NonClientRegionChangedEventHandlerImpl,
) *ICoreWebView2NonClientRegionChangedEventHandler {
	if _ICoreWebView2NonClientRegionChangedEventHandlerFn == nil {
		_ICoreWebView2NonClientRegionChangedEventHandlerFn = &_ICoreWebView2NonClientRegionChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2NonClientRegionChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2NonClientRegionChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2NonClientRegionChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2NonClientRegionChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2NonClientRegionChangedEventHandler{
		vtbl: _ICoreWebView2NonClientRegionChangedEventHandlerFn,
		impl: impl,
	}
}
