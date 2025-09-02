package edge

import "unsafe"

type _ICoreWebView2FrameNavigationCompletedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameNavigationCompletedEventHandler 接收框架导航完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framenavigationcompletedeventhandler
type ICoreWebView2FrameNavigationCompletedEventHandler struct {
	vtbl *_ICoreWebView2FrameNavigationCompletedEventHandlerVtbl
	impl _ICoreWebView2FrameNavigationCompletedEventHandlerImpl
}

func (i *ICoreWebView2FrameNavigationCompletedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameNavigationCompletedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNavigationCompletedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNavigationCompletedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownRelease(this *ICoreWebView2FrameNavigationCompletedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerInvoke(this *ICoreWebView2FrameNavigationCompletedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return this.impl.FrameNavigationCompleted(sender, args)
}

type _ICoreWebView2FrameNavigationCompletedEventHandlerImpl interface {
	IUnknownImpl
	FrameNavigationCompleted(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
}

var _ICoreWebView2FrameNavigationCompletedEventHandlerFn *_ICoreWebView2FrameNavigationCompletedEventHandlerVtbl

func NewICoreWebView2FrameNavigationCompletedEventHandler(impl _ICoreWebView2FrameNavigationCompletedEventHandlerImpl) *ICoreWebView2FrameNavigationCompletedEventHandler {
	if _ICoreWebView2FrameNavigationCompletedEventHandlerFn == nil {
		_ICoreWebView2FrameNavigationCompletedEventHandlerFn = &_ICoreWebView2FrameNavigationCompletedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameNavigationCompletedEventHandler{
		vtbl: _ICoreWebView2FrameNavigationCompletedEventHandlerFn,
		impl: impl,
	}
}