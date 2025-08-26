package edge

import "unsafe"

type _ICoreWebView2StateChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2StateChangedEventHandler 接收下载状态改变事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2statechangedeventhandler
type ICoreWebView2StateChangedEventHandler struct {
	vtbl *_ICoreWebView2StateChangedEventHandlerVtbl
	impl _ICoreWebView2StateChangedEventHandlerImpl
}

func (i *ICoreWebView2StateChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2StateChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2StateChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2StateChangedEventHandlerIUnknownAddRef(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2StateChangedEventHandlerIUnknownRelease(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2StateChangedEventHandlerInvoke(this *ICoreWebView2StateChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	return this.impl.StateChanged(sender, args)
}

type _ICoreWebView2StateChangedEventHandlerImpl interface {
	IUnknownImpl
	StateChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr
}

var _ICoreWebView2StateChangedEventHandlerFn *_ICoreWebView2StateChangedEventHandlerVtbl

func NewICoreWebView2StateChangedEventHandler(impl _ICoreWebView2StateChangedEventHandlerImpl) *ICoreWebView2StateChangedEventHandler {
	if _ICoreWebView2StateChangedEventHandlerFn == nil {
		_ICoreWebView2StateChangedEventHandlerFn = &_ICoreWebView2StateChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2StateChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2StateChangedEventHandler{
		vtbl: _ICoreWebView2StateChangedEventHandlerFn,
		impl: impl,
	}
}
