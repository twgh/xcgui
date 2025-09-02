package edge

import "unsafe"

type _ICoreWebView2FrameNavigationStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameNavigationStartingEventHandler 接收框架导航开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framenavigationstartingeventhandler
type ICoreWebView2FrameNavigationStartingEventHandler struct {
	vtbl *_ICoreWebView2FrameNavigationStartingEventHandlerVtbl
	impl _ICoreWebView2FrameNavigationStartingEventHandlerImpl
}

func (i *ICoreWebView2FrameNavigationStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameNavigationStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNavigationStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNavigationStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownRelease(this *ICoreWebView2FrameNavigationStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNavigationStartingEventHandlerInvoke(this *ICoreWebView2FrameNavigationStartingEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	return this.impl.FrameNavigationStarting(sender, args)
}

type _ICoreWebView2FrameNavigationStartingEventHandlerImpl interface {
	IUnknownImpl
	FrameNavigationStarting(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr
}

var _ICoreWebView2FrameNavigationStartingEventHandlerFn *_ICoreWebView2FrameNavigationStartingEventHandlerVtbl

func NewICoreWebView2FrameNavigationStartingEventHandler(impl _ICoreWebView2FrameNavigationStartingEventHandlerImpl) *ICoreWebView2FrameNavigationStartingEventHandler {
	if _ICoreWebView2FrameNavigationStartingEventHandlerFn == nil {
		_ICoreWebView2FrameNavigationStartingEventHandlerFn = &_ICoreWebView2FrameNavigationStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameNavigationStartingEventHandler{
		vtbl: _ICoreWebView2FrameNavigationStartingEventHandlerFn,
		impl: impl,
	}
}
