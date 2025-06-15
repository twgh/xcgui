package edge

import "unsafe"

type _ICoreWebView2_Frame_NavigationCompletedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2_Frame_NavigationCompletedEventHandler 接收 ICoreWebView2 的框架导航完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2NavigationCompletedEventHandler
type ICoreWebView2_Frame_NavigationCompletedEventHandler struct {
	vtbl *_ICoreWebView2_Frame_NavigationCompletedEventHandlerVtbl
	impl _ICoreWebView2_Frame_NavigationCompletedEventHandlerImpl
}

func (i *ICoreWebView2_Frame_NavigationCompletedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_Frame_NavigationCompletedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownQueryInterface(this *ICoreWebView2_Frame_NavigationCompletedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownAddRef(this *ICoreWebView2_Frame_NavigationCompletedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownRelease(this *ICoreWebView2_Frame_NavigationCompletedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2_Frame_NavigationCompletedEventHandlerInvoke(this *ICoreWebView2_Frame_NavigationCompletedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return this.impl.Frame_NavigationCompleted(sender, args)
}

type _ICoreWebView2_Frame_NavigationCompletedEventHandlerImpl interface {
	IUnknownImpl
	Frame_NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
}

var _ICoreWebView2_Frame_NavigationCompletedEventHandlerFn *_ICoreWebView2_Frame_NavigationCompletedEventHandlerVtbl

func NewICoreWebView2_Frame_NavigationCompletedEventHandler(impl _ICoreWebView2_Frame_NavigationCompletedEventHandlerImpl) *ICoreWebView2_Frame_NavigationCompletedEventHandler {
	if _ICoreWebView2_Frame_NavigationCompletedEventHandlerFn == nil {
		_ICoreWebView2_Frame_NavigationCompletedEventHandlerFn = &_ICoreWebView2_Frame_NavigationCompletedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2_Frame_NavigationCompletedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2_Frame_NavigationCompletedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2_Frame_NavigationCompletedEventHandler{
		vtbl: _ICoreWebView2_Frame_NavigationCompletedEventHandlerFn,
		impl: impl,
	}
}
