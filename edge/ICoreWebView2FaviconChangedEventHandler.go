package edge

import "unsafe"

type _ICoreWebView2FaviconChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FaviconChangedEventHandler 接收网站图标(Favicon)变更事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2faviconchangedeventhandler
type ICoreWebView2FaviconChangedEventHandler struct {
	vtbl *_ICoreWebView2FaviconChangedEventHandlerVtbl
	impl _ICoreWebView2FaviconChangedEventHandlerImpl
}

func (i *ICoreWebView2FaviconChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FaviconChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FaviconChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownRelease(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FaviconChangedEventHandlerInvoke(this *ICoreWebView2FaviconChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.FaviconChanged(sender, args)
}

type _ICoreWebView2FaviconChangedEventHandlerImpl interface {
	IUnknownImpl
	FaviconChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2FaviconChangedEventHandlerFn *_ICoreWebView2FaviconChangedEventHandlerVtbl

func NewICoreWebView2FaviconChangedEventHandler(impl _ICoreWebView2FaviconChangedEventHandlerImpl) *ICoreWebView2FaviconChangedEventHandler {
	if _ICoreWebView2FaviconChangedEventHandlerFn == nil {
		_ICoreWebView2FaviconChangedEventHandlerFn = &_ICoreWebView2FaviconChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FaviconChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FaviconChangedEventHandler{
		vtbl: _ICoreWebView2FaviconChangedEventHandlerFn,
		impl: impl,
	}
}
