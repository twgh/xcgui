package edge

import "unsafe"

type _ICoreWebView2ProcessInfosChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ProcessInfosChangedEventHandler 接收进程信息变更事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processinfoschangedeventhandler
type ICoreWebView2ProcessInfosChangedEventHandler struct {
	vtbl *_ICoreWebView2ProcessInfosChangedEventHandlerVtbl
	impl _ICoreWebView2ProcessInfosChangedEventHandlerImpl
}

func (i *ICoreWebView2ProcessInfosChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessInfosChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ProcessInfosChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ProcessInfosChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownRelease(this *ICoreWebView2ProcessInfosChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProcessInfosChangedEventHandlerInvoke(this *ICoreWebView2ProcessInfosChangedEventHandler, sender *ICoreWebView2Environment, args *IUnknown) uintptr {
	return this.impl.ProcessInfosChanged(sender, args)
}

type _ICoreWebView2ProcessInfosChangedEventHandlerImpl interface {
	IUnknownImpl
	ProcessInfosChanged(sender *ICoreWebView2Environment, args *IUnknown) uintptr
}

var _ICoreWebView2ProcessInfosChangedEventHandlerFn *_ICoreWebView2ProcessInfosChangedEventHandlerVtbl

func NewICoreWebView2ProcessInfosChangedEventHandler(impl _ICoreWebView2ProcessInfosChangedEventHandlerImpl) *ICoreWebView2ProcessInfosChangedEventHandler {
	if _ICoreWebView2ProcessInfosChangedEventHandlerFn == nil {
		_ICoreWebView2ProcessInfosChangedEventHandlerFn = &_ICoreWebView2ProcessInfosChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ProcessInfosChangedEventHandler{
		vtbl: _ICoreWebView2ProcessInfosChangedEventHandlerFn,
		impl: impl,
	}
}
