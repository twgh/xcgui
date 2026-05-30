package edge

import "unsafe"

type _ICoreWebView2DragStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DragStartingEventHandler 接收拖拽开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2dragstartingeventhandler
type ICoreWebView2DragStartingEventHandler struct {
	vtbl *_ICoreWebView2DragStartingEventHandlerVtbl
	impl _ICoreWebView2DragStartingEventHandlerImpl
}

func (i *ICoreWebView2DragStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DragStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2DragStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2DragStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DragStartingEventHandlerIUnknownAddRef(this *ICoreWebView2DragStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DragStartingEventHandlerIUnknownRelease(this *ICoreWebView2DragStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DragStartingEventHandlerInvoke(this *ICoreWebView2DragStartingEventHandler, sender *ICoreWebView2CompositionController, args *ICoreWebView2DragStartingEventArgs) uintptr {
	return this.impl.DragStarting(sender, args)
}

type _ICoreWebView2DragStartingEventHandlerImpl interface {
	IUnknownImpl
	DragStarting(sender *ICoreWebView2CompositionController, args *ICoreWebView2DragStartingEventArgs) uintptr
}

var _ICoreWebView2DragStartingEventHandlerFn *_ICoreWebView2DragStartingEventHandlerVtbl

func NewICoreWebView2DragStartingEventHandler(impl _ICoreWebView2DragStartingEventHandlerImpl) *ICoreWebView2DragStartingEventHandler {
	if _ICoreWebView2DragStartingEventHandlerFn == nil {
		_ICoreWebView2DragStartingEventHandlerFn = &_ICoreWebView2DragStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2DragStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2DragStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2DragStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2DragStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2DragStartingEventHandler{
		vtbl: _ICoreWebView2DragStartingEventHandlerFn,
		impl: impl,
	}
}
