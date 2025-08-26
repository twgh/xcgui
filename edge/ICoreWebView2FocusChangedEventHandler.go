package edge

import "unsafe"

type _ICoreWebView2FocusChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FocusChangedEventHandler 接收焦点变化事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2focuschangedeventhandler
type ICoreWebView2FocusChangedEventHandler struct {
	vtbl *_ICoreWebView2FocusChangedEventHandlerVtbl
	impl _ICoreWebView2FocusChangedEventHandlerImpl
}

func (i *ICoreWebView2FocusChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FocusChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FocusChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FocusChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FocusChangedEventHandlerIUnknownRelease(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FocusChangedEventHandlerInvoke(this *ICoreWebView2FocusChangedEventHandler, sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	return this.impl.FocusChanged(sender, args)
}

type _ICoreWebView2FocusChangedEventHandlerImpl interface {
	IUnknownImpl
	FocusChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr
}

var _ICoreWebView2FocusChangedEventHandlerFn *_ICoreWebView2FocusChangedEventHandlerVtbl

func NewICoreWebView2FocusChangedEventHandler(impl _ICoreWebView2FocusChangedEventHandlerImpl) *ICoreWebView2FocusChangedEventHandler {
	if _ICoreWebView2FocusChangedEventHandlerFn == nil {
		_ICoreWebView2FocusChangedEventHandlerFn = &_ICoreWebView2FocusChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FocusChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FocusChangedEventHandler{
		vtbl: _ICoreWebView2FocusChangedEventHandlerFn,
		impl: impl,
	}
}