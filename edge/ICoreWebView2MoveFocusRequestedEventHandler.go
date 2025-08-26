package edge

import "unsafe"

type _ICoreWebView2MoveFocusRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2MoveFocusRequestedEventHandler 接收移动焦点请求事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventhandler
type ICoreWebView2MoveFocusRequestedEventHandler struct {
	vtbl *_ICoreWebView2MoveFocusRequestedEventHandlerVtbl
	impl _ICoreWebView2MoveFocusRequestedEventHandlerImpl
}

func (i *ICoreWebView2MoveFocusRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2MoveFocusRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2MoveFocusRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2MoveFocusRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2MoveFocusRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2MoveFocusRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2MoveFocusRequestedEventHandlerIUnknownRelease(this *ICoreWebView2MoveFocusRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2MoveFocusRequestedEventHandlerInvoke(this *ICoreWebView2MoveFocusRequestedEventHandler, sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr {
	return this.impl.MoveFocusRequested(sender, args)
}

type _ICoreWebView2MoveFocusRequestedEventHandlerImpl interface {
	IUnknownImpl
	MoveFocusRequested(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr
}

var _ICoreWebView2MoveFocusRequestedEventHandlerFn *_ICoreWebView2MoveFocusRequestedEventHandlerVtbl

func NewICoreWebView2MoveFocusRequestedEventHandler(impl _ICoreWebView2MoveFocusRequestedEventHandlerImpl) *ICoreWebView2MoveFocusRequestedEventHandler {
	if _ICoreWebView2MoveFocusRequestedEventHandlerFn == nil {
		_ICoreWebView2MoveFocusRequestedEventHandlerFn = &_ICoreWebView2MoveFocusRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2MoveFocusRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2MoveFocusRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2MoveFocusRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2MoveFocusRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2MoveFocusRequestedEventHandler{
		vtbl: _ICoreWebView2MoveFocusRequestedEventHandlerFn,
		impl: impl,
	}
}