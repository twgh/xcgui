package edge

import "unsafe"

type _ICoreWebView2LostFocusEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2LostFocusEventHandler 接收失去焦点事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2focuschangedeventhandler
type ICoreWebView2LostFocusEventHandler struct {
	vtbl *_ICoreWebView2LostFocusEventHandlerVtbl
	impl _ICoreWebView2LostFocusEventHandlerImpl
}

func (i *ICoreWebView2LostFocusEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2LostFocusEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2LostFocusEventHandlerIUnknownQueryInterface(this *ICoreWebView2LostFocusEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2LostFocusEventHandlerIUnknownAddRef(this *ICoreWebView2LostFocusEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2LostFocusEventHandlerIUnknownRelease(this *ICoreWebView2LostFocusEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2LostFocusEventHandlerInvoke(this *ICoreWebView2LostFocusEventHandler, sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	return this.impl.LostFocus(sender, args)
}

type _ICoreWebView2LostFocusEventHandlerImpl interface {
	IUnknownImpl
	LostFocus(sender *ICoreWebView2Controller, args *IUnknown) uintptr
}

var _ICoreWebView2LostFocusEventHandlerFn *_ICoreWebView2LostFocusEventHandlerVtbl

func NewICoreWebView2LostFocusEventHandler(impl _ICoreWebView2LostFocusEventHandlerImpl) *ICoreWebView2LostFocusEventHandler {
	if _ICoreWebView2LostFocusEventHandlerFn == nil {
		_ICoreWebView2LostFocusEventHandlerFn = &_ICoreWebView2LostFocusEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2LostFocusEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2LostFocusEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2LostFocusEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2LostFocusEventHandlerInvoke),
		}
	}
	return &ICoreWebView2LostFocusEventHandler{
		vtbl: _ICoreWebView2LostFocusEventHandlerFn,
		impl: impl,
	}
}
