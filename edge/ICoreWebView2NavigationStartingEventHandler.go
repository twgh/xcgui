package edge

import "unsafe"

type _ICoreWebView2NavigationStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NavigationStartingEventHandler 接收 NavigationStarting 事件。
//
//	https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventhandler
type ICoreWebView2NavigationStartingEventHandler struct {
	vtbl *_ICoreWebView2NavigationStartingEventHandlerVtbl
	impl _ICoreWebView2NavigationStartingEventHandlerImpl
}

func (i *ICoreWebView2NavigationStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NavigationStartingEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2NavigationStartingEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NavigationStartingEventHandlerIUnknownAddRef(
	this *ICoreWebView2NavigationStartingEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NavigationStartingEventHandlerIUnknownRelease(
	this *ICoreWebView2NavigationStartingEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NavigationStartingEventHandlerInvoke(
	this *ICoreWebView2NavigationStartingEventHandler,
	sender *ICoreWebView2,
	args *ICoreWebView2NavigationStartingEventArgs,
) uintptr {
	return this.impl.NavigationStarting(sender, args)
}

type _ICoreWebView2NavigationStartingEventHandlerImpl interface {
	IUnknownImpl
	NavigationStarting(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
}

var _ICoreWebView2NavigationStartingEventHandlerFn *_ICoreWebView2NavigationStartingEventHandlerVtbl

func NewICoreWebView2NavigationStartingEventHandler(
	impl _ICoreWebView2NavigationStartingEventHandlerImpl,
) *ICoreWebView2NavigationStartingEventHandler {
	if _ICoreWebView2NavigationStartingEventHandlerFn == nil {
		_ICoreWebView2NavigationStartingEventHandlerFn = &_ICoreWebView2NavigationStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2NavigationStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2NavigationStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2NavigationStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2NavigationStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2NavigationStartingEventHandler{
		vtbl: _ICoreWebView2NavigationStartingEventHandlerFn,
		impl: impl,
	}
}
