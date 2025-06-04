package edge

import "unsafe"

type _ICoreWebView2NavigationCompletedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NavigationCompletedEventHandler 接收 NavigationCompleted 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventhandler
type ICoreWebView2NavigationCompletedEventHandler struct {
	vtbl *_ICoreWebView2NavigationCompletedEventHandlerVtbl
	impl _ICoreWebView2NavigationCompletedEventHandlerImpl
}

func (i *ICoreWebView2NavigationCompletedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationCompletedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NavigationCompletedEventHandlerIUnknownQueryInterface(this *ICoreWebView2NavigationCompletedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NavigationCompletedEventHandlerIUnknownAddRef(this *ICoreWebView2NavigationCompletedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NavigationCompletedEventHandlerIUnknownRelease(this *ICoreWebView2NavigationCompletedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NavigationCompletedEventHandlerInvoke(this *ICoreWebView2NavigationCompletedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return this.impl.NavigationCompleted(sender, args)
}

type _ICoreWebView2NavigationCompletedEventHandlerImpl interface {
	IUnknownImpl
	NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
}

var _ICoreWebView2NavigationCompletedEventHandlerFn = _ICoreWebView2NavigationCompletedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2NavigationCompletedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2NavigationCompletedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2NavigationCompletedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2NavigationCompletedEventHandlerInvoke),
}

func NewICoreWebView2NavigationCompletedEventHandler(impl _ICoreWebView2NavigationCompletedEventHandlerImpl) *ICoreWebView2NavigationCompletedEventHandler {
	return &ICoreWebView2NavigationCompletedEventHandler{
		vtbl: &_ICoreWebView2NavigationCompletedEventHandlerFn,
		impl: impl,
	}
}
