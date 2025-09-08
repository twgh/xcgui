package edge

import "unsafe"

type _ICoreWebView2FindActiveMatchIndexChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FindActiveMatchIndexChangedEventHandler 接收查找活动匹配项索引更改事件。
//
// https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2findactivematchindexchangedeventhandler
type ICoreWebView2FindActiveMatchIndexChangedEventHandler struct {
	vtbl *_ICoreWebView2FindActiveMatchIndexChangedEventHandlerVtbl
	impl _ICoreWebView2FindActiveMatchIndexChangedEventHandlerImpl
}

func (i *ICoreWebView2FindActiveMatchIndexChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FindActiveMatchIndexChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FindActiveMatchIndexChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FindActiveMatchIndexChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownRelease(this *ICoreWebView2FindActiveMatchIndexChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FindActiveMatchIndexChangedEventHandlerInvoke(this *ICoreWebView2FindActiveMatchIndexChangedEventHandler, sender *ICoreWebView2Find, args *IUnknown) uintptr {
	return this.impl.FindActiveMatchIndexChanged(sender, args)
}

type _ICoreWebView2FindActiveMatchIndexChangedEventHandlerImpl interface {
	IUnknownImpl
	FindActiveMatchIndexChanged(sender *ICoreWebView2Find, args *IUnknown) uintptr
}

var _ICoreWebView2FindActiveMatchIndexChangedEventHandlerFn *_ICoreWebView2FindActiveMatchIndexChangedEventHandlerVtbl

func NewICoreWebView2FindActiveMatchIndexChangedEventHandler(impl _ICoreWebView2FindActiveMatchIndexChangedEventHandlerImpl) *ICoreWebView2FindActiveMatchIndexChangedEventHandler {
	if _ICoreWebView2FindActiveMatchIndexChangedEventHandlerFn == nil {
		_ICoreWebView2FindActiveMatchIndexChangedEventHandlerFn = &_ICoreWebView2FindActiveMatchIndexChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FindActiveMatchIndexChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FindActiveMatchIndexChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FindActiveMatchIndexChangedEventHandler{
		vtbl: _ICoreWebView2FindActiveMatchIndexChangedEventHandlerFn,
		impl: impl,
	}
}
