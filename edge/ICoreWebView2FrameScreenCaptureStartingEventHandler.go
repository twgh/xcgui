package edge

import "unsafe"

type _ICoreWebView2FrameScreenCaptureStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameScreenCaptureStartingEventHandler 接收框架屏幕截图开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framescreencapturestartingeventhandler
type ICoreWebView2FrameScreenCaptureStartingEventHandler struct {
	vtbl *_ICoreWebView2FrameScreenCaptureStartingEventHandlerVtbl
	impl _ICoreWebView2FrameScreenCaptureStartingEventHandlerImpl
}

func (i *ICoreWebView2FrameScreenCaptureStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameScreenCaptureStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameScreenCaptureStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownAddRef(this *ICoreWebView2FrameScreenCaptureStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownRelease(this *ICoreWebView2FrameScreenCaptureStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameScreenCaptureStartingEventHandlerInvoke(this *ICoreWebView2FrameScreenCaptureStartingEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr {
	return this.impl.FrameScreenCaptureStarting(sender, args)
}

type _ICoreWebView2FrameScreenCaptureStartingEventHandlerImpl interface {
	IUnknownImpl
	FrameScreenCaptureStarting(sender *ICoreWebView2Frame, args *ICoreWebView2ScreenCaptureStartingEventArgs) uintptr
}

var _ICoreWebView2FrameScreenCaptureStartingEventHandlerFn *_ICoreWebView2FrameScreenCaptureStartingEventHandlerVtbl

func NewICoreWebView2FrameScreenCaptureStartingEventHandler(impl _ICoreWebView2FrameScreenCaptureStartingEventHandlerImpl) *ICoreWebView2FrameScreenCaptureStartingEventHandler {
	if _ICoreWebView2FrameScreenCaptureStartingEventHandlerFn == nil {
		_ICoreWebView2FrameScreenCaptureStartingEventHandlerFn = &_ICoreWebView2FrameScreenCaptureStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameScreenCaptureStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameScreenCaptureStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameScreenCaptureStartingEventHandler{
		vtbl: _ICoreWebView2FrameScreenCaptureStartingEventHandlerFn,
		impl: impl,
	}
}
