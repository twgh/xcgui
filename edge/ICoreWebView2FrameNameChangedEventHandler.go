package edge

import "unsafe"

type _ICoreWebView2FrameNameChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameNameChangedEventHandler 接收框架名称更改事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framenamechangedeventhandler
type ICoreWebView2FrameNameChangedEventHandler struct {
	vtbl *_ICoreWebView2FrameNameChangedEventHandlerVtbl
	impl _ICoreWebView2FrameNameChangedEventHandlerImpl
}

func (i *ICoreWebView2FrameNameChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameNameChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameNameChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNameChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNameChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNameChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNameChangedEventHandlerIUnknownRelease(this *ICoreWebView2FrameNameChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNameChangedEventHandlerInvoke(this *ICoreWebView2FrameNameChangedEventHandler, sender *ICoreWebView2Frame, args *IUnknown) uintptr {
	return this.impl.FrameNameChanged(sender, args)
}

type _ICoreWebView2FrameNameChangedEventHandlerImpl interface {
	IUnknownImpl
	FrameNameChanged(sender *ICoreWebView2Frame, args *IUnknown) uintptr
}

var _ICoreWebView2FrameNameChangedEventHandlerFn *_ICoreWebView2FrameNameChangedEventHandlerVtbl

func NewICoreWebView2FrameNameChangedEventHandler(impl _ICoreWebView2FrameNameChangedEventHandlerImpl) *ICoreWebView2FrameNameChangedEventHandler {
	if _ICoreWebView2FrameNameChangedEventHandlerFn == nil {
		_ICoreWebView2FrameNameChangedEventHandlerFn = &_ICoreWebView2FrameNameChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameNameChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameNameChangedEventHandler{
		vtbl: _ICoreWebView2FrameNameChangedEventHandlerFn,
		impl: impl,
	}
}
