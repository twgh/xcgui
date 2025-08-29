package edge

import "unsafe"

type _ICoreWebView2ScreenCaptureStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ScreenCaptureStartingEventHandler 接收屏幕截图开始事件。
//
// https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2screencapturestartingeventhandler
type ICoreWebView2ScreenCaptureStartingEventHandler struct {
	vtbl *_ICoreWebView2ScreenCaptureStartingEventHandlerVtbl
	impl _ICoreWebView2ScreenCaptureStartingEventHandlerImpl
}

func (i *ICoreWebView2ScreenCaptureStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScreenCaptureStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2ScreenCaptureStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownAddRef(this *ICoreWebView2ScreenCaptureStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownRelease(this *ICoreWebView2ScreenCaptureStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ScreenCaptureStartingEventHandlerInvoke(this *ICoreWebView2ScreenCaptureStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr {
	return this.impl.ScreenCaptureStarting(sender, args)
}

type _ICoreWebView2ScreenCaptureStartingEventHandlerImpl interface {
	IUnknownImpl
	ScreenCaptureStarting(sender *ICoreWebView2, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr
}

var _ICoreWebView2ScreenCaptureStartingEventHandlerFn *_ICoreWebView2ScreenCaptureStartingEventHandlerVtbl

func NewICoreWebView2ScreenCaptureStartingEventHandler(impl _ICoreWebView2ScreenCaptureStartingEventHandlerImpl) *ICoreWebView2ScreenCaptureStartingEventHandler {
	if _ICoreWebView2ScreenCaptureStartingEventHandlerFn == nil {
		_ICoreWebView2ScreenCaptureStartingEventHandlerFn = &_ICoreWebView2ScreenCaptureStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ScreenCaptureStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ScreenCaptureStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ScreenCaptureStartingEventHandler{
		vtbl: _ICoreWebView2ScreenCaptureStartingEventHandlerFn,
		impl: impl,
	}
}