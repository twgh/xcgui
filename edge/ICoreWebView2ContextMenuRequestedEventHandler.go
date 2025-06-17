package edge

import "unsafe"

type _ICoreWebView2ContextMenuRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ContextMenuRequestedEventHandler 接收 ContextMenuRequested 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventhandler
type ICoreWebView2ContextMenuRequestedEventHandler struct {
	vtbl *_ICoreWebView2ContextMenuRequestedEventHandlerVtbl
	impl _ICoreWebView2ContextMenuRequestedEventHandlerImpl
}

func (i *ICoreWebView2ContextMenuRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2ContextMenuRequestedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownAddRef(
	this *ICoreWebView2ContextMenuRequestedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownRelease(
	this *ICoreWebView2ContextMenuRequestedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContextMenuRequestedEventHandlerInvoke(
	this *ICoreWebView2ContextMenuRequestedEventHandler,
	sender *ICoreWebView2,
	args *ICoreWebView2ContextMenuRequestedEventArgs,
) uintptr {
	return this.impl.ContextMenuRequested(sender, args)
}

type _ICoreWebView2ContextMenuRequestedEventHandlerImpl interface {
	IUnknownImpl
	ContextMenuRequested(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr
}

var _ICoreWebView2ContextMenuRequestedEventHandlerFn *_ICoreWebView2ContextMenuRequestedEventHandlerVtbl

func NewICoreWebView2ContextMenuRequestedEventHandler(
	impl _ICoreWebView2ContextMenuRequestedEventHandlerImpl,
) *ICoreWebView2ContextMenuRequestedEventHandler {
	if _ICoreWebView2ContextMenuRequestedEventHandlerFn == nil {
		_ICoreWebView2ContextMenuRequestedEventHandlerFn = &_ICoreWebView2ContextMenuRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ContextMenuRequestedEventHandler{
		vtbl: _ICoreWebView2ContextMenuRequestedEventHandlerFn,
		impl: impl,
	}
}
