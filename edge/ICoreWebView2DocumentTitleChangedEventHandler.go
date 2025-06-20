package edge

import "unsafe"

type _ICoreWebView2DocumentTitleChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DocumentTitleChangedEventHandler 接收 DocumentTitleChanged 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2documenttitlechangedeventhandler
type ICoreWebView2DocumentTitleChangedEventHandler struct {
	vtbl *_ICoreWebView2DocumentTitleChangedEventHandlerVtbl
	impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl
}

func (i *ICoreWebView2DocumentTitleChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DocumentTitleChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DocumentTitleChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerInvoke(this *ICoreWebView2DocumentTitleChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.DocumentTitleChanged(sender, args)
}

type _ICoreWebView2DocumentTitleChangedEventHandlerImpl interface {
	IUnknownImpl
	DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2DocumentTitleChangedEventHandlerFn *_ICoreWebView2DocumentTitleChangedEventHandlerVtbl

func NewICoreWebView2DocumentTitleChangedEventHandler(impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl) *ICoreWebView2DocumentTitleChangedEventHandler {
	if _ICoreWebView2DocumentTitleChangedEventHandlerFn == nil {
		_ICoreWebView2DocumentTitleChangedEventHandlerFn = &_ICoreWebView2DocumentTitleChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2DocumentTitleChangedEventHandler{
		vtbl: _ICoreWebView2DocumentTitleChangedEventHandlerFn,
		impl: impl,
	}
}
