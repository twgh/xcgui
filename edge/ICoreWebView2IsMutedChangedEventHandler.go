package edge

import "unsafe"

type _ICoreWebView2IsMutedChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2IsMutedChangedEventHandler 接收 IsMutedChanged 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2ismutedchangedeventhandler
type ICoreWebView2IsMutedChangedEventHandler struct {
	vtbl *_ICoreWebView2IsMutedChangedEventHandlerVtbl
	impl _ICoreWebView2IsMutedChangedEventHandlerImpl
}

func (i *ICoreWebView2IsMutedChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2IsMutedChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2IsMutedChangedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2IsMutedChangedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2IsMutedChangedEventHandlerIUnknownAddRef(
	this *ICoreWebView2IsMutedChangedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2IsMutedChangedEventHandlerIUnknownRelease(
	this *ICoreWebView2IsMutedChangedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2IsMutedChangedEventHandlerInvoke(
	this *ICoreWebView2IsMutedChangedEventHandler,
	sender *ICoreWebView2,
	args *IUnknown,
) uintptr {
	return this.impl.IsMutedChanged(sender, args)
}

type _ICoreWebView2IsMutedChangedEventHandlerImpl interface {
	IUnknownImpl
	IsMutedChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2IsMutedChangedEventHandlerFn *_ICoreWebView2IsMutedChangedEventHandlerVtbl

func NewICoreWebView2IsMutedChangedEventHandler(
	impl _ICoreWebView2IsMutedChangedEventHandlerImpl,
) *ICoreWebView2IsMutedChangedEventHandler {
	if _ICoreWebView2IsMutedChangedEventHandlerFn == nil {
		_ICoreWebView2IsMutedChangedEventHandlerFn = &_ICoreWebView2IsMutedChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2IsMutedChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2IsMutedChangedEventHandler{
		vtbl: _ICoreWebView2IsMutedChangedEventHandlerFn,
		impl: impl,
	}
}
