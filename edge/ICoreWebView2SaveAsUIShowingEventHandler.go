package edge

import "unsafe"

type _ICoreWebView2SaveAsUIShowingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2SaveAsUIShowingEventHandler 接收"另存为"UI显示事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2saveasuishowingeventhandler
type ICoreWebView2SaveAsUIShowingEventHandler struct {
	vtbl *_ICoreWebView2SaveAsUIShowingEventHandlerVtbl
	impl _ICoreWebView2SaveAsUIShowingEventHandlerImpl
}

func (i *ICoreWebView2SaveAsUIShowingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveAsUIShowingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2SaveAsUIShowingEventHandlerIUnknownQueryInterface(this *ICoreWebView2SaveAsUIShowingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2SaveAsUIShowingEventHandlerIUnknownAddRef(this *ICoreWebView2SaveAsUIShowingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2SaveAsUIShowingEventHandlerIUnknownRelease(this *ICoreWebView2SaveAsUIShowingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2SaveAsUIShowingEventHandlerInvoke(this *ICoreWebView2SaveAsUIShowingEventHandler, sender *ICoreWebView2, args *ICoreWebView2SaveAsUIShowingEventArgs) uintptr {
	return this.impl.SaveAsUIShowing(sender, args)
}

type _ICoreWebView2SaveAsUIShowingEventHandlerImpl interface {
	IUnknownImpl
	SaveAsUIShowing(sender *ICoreWebView2, args *ICoreWebView2SaveAsUIShowingEventArgs) uintptr
}

var _ICoreWebView2SaveAsUIShowingEventHandlerFn *_ICoreWebView2SaveAsUIShowingEventHandlerVtbl

func NewICoreWebView2SaveAsUIShowingEventHandler(impl _ICoreWebView2SaveAsUIShowingEventHandlerImpl) *ICoreWebView2SaveAsUIShowingEventHandler {
	if _ICoreWebView2SaveAsUIShowingEventHandlerFn == nil {
		_ICoreWebView2SaveAsUIShowingEventHandlerFn = &_ICoreWebView2SaveAsUIShowingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2SaveAsUIShowingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2SaveAsUIShowingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2SaveAsUIShowingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2SaveAsUIShowingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2SaveAsUIShowingEventHandler{
		vtbl: _ICoreWebView2SaveAsUIShowingEventHandlerFn,
		impl: impl,
	}
}
