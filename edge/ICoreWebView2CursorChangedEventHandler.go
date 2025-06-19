package edge

import "unsafe"

type _ICoreWebView2CursorChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CursorChangedEventHandler 接收光标改变事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2cursorchangedeventhandler
type ICoreWebView2CursorChangedEventHandler struct {
	vtbl *_ICoreWebView2CursorChangedEventHandlerVtbl
	impl _ICoreWebView2CursorChangedEventHandlerImpl
}

func (i *ICoreWebView2CursorChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CursorChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2CursorChangedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2CursorChangedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CursorChangedEventHandlerIUnknownAddRef(
	this *ICoreWebView2CursorChangedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CursorChangedEventHandlerIUnknownRelease(
	this *ICoreWebView2CursorChangedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CursorChangedEventHandlerInvoke(
	this *ICoreWebView2CursorChangedEventHandler,
	sender *ICoreWebView2CompositionController,
	args *IUnknown,
) uintptr {
	return this.impl.CursorChanged(sender, args)
}

type _ICoreWebView2CursorChangedEventHandlerImpl interface {
	IUnknownImpl
	CursorChanged(sender *ICoreWebView2CompositionController, args *IUnknown) uintptr
}

var _ICoreWebView2CursorChangedEventHandlerFn *_ICoreWebView2CursorChangedEventHandlerVtbl

func NewICoreWebView2CursorChangedEventHandler(
	impl _ICoreWebView2CursorChangedEventHandlerImpl,
) *ICoreWebView2CursorChangedEventHandler {
	if _ICoreWebView2CursorChangedEventHandlerFn == nil {
		_ICoreWebView2CursorChangedEventHandlerFn = &_ICoreWebView2CursorChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2CursorChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2CursorChangedEventHandler{
		vtbl: _ICoreWebView2CursorChangedEventHandlerFn,
		impl: impl,
	}
}
