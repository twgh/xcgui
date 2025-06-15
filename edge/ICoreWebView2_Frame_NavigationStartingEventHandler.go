package edge

import "unsafe"

type _ICoreWebView2_Frame_NavigationStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2_Frame_NavigationStartingEventHandler 接收 ICoreWebView2 的框架导航开始事件。
//
//	https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventhandler
type ICoreWebView2_Frame_NavigationStartingEventHandler struct {
	vtbl *_ICoreWebView2_Frame_NavigationStartingEventHandlerVtbl
	impl _ICoreWebView2_Frame_NavigationStartingEventHandlerImpl
}

func (i *ICoreWebView2_Frame_NavigationStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_Frame_NavigationStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2_Frame_NavigationStartingEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownAddRef(
	this *ICoreWebView2_Frame_NavigationStartingEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownRelease(
	this *ICoreWebView2_Frame_NavigationStartingEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2_Frame_NavigationStartingEventHandlerInvoke(
	this *ICoreWebView2_Frame_NavigationStartingEventHandler,
	sender *ICoreWebView2,
	args *ICoreWebView2NavigationStartingEventArgs,
) uintptr {
	return this.impl.Frame_NavigationStarting(sender, args)
}

type _ICoreWebView2_Frame_NavigationStartingEventHandlerImpl interface {
	IUnknownImpl
	Frame_NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
}

var _ICoreWebView2_Frame_NavigationStartingEventHandlerFn *_ICoreWebView2_Frame_NavigationStartingEventHandlerVtbl

func NewICoreWebView2_Frame_NavigationStartingEventHandler(
	impl _ICoreWebView2_Frame_NavigationStartingEventHandlerImpl,
) *ICoreWebView2_Frame_NavigationStartingEventHandler {
	if _ICoreWebView2_Frame_NavigationStartingEventHandlerFn == nil {
		_ICoreWebView2_Frame_NavigationStartingEventHandlerFn = &_ICoreWebView2_Frame_NavigationStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2_Frame_NavigationStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2_Frame_NavigationStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2_Frame_NavigationStartingEventHandler{
		vtbl: _ICoreWebView2_Frame_NavigationStartingEventHandlerFn,
		impl: impl,
	}
}
