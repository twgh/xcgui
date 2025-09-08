package edge

import "unsafe"

type _ICoreWebView2FindMatchCountChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FindMatchCountChangedEventHandler 接收查找匹配项数量更改事件。
//
// https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2findmatchcountchangedeventhandler
type ICoreWebView2FindMatchCountChangedEventHandler struct {
	vtbl *_ICoreWebView2FindMatchCountChangedEventHandlerVtbl
	impl _ICoreWebView2FindMatchCountChangedEventHandlerImpl
}

func (i *ICoreWebView2FindMatchCountChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FindMatchCountChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FindMatchCountChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FindMatchCountChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FindMatchCountChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FindMatchCountChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FindMatchCountChangedEventHandlerIUnknownRelease(this *ICoreWebView2FindMatchCountChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FindMatchCountChangedEventHandlerInvoke(this *ICoreWebView2FindMatchCountChangedEventHandler, sender *ICoreWebView2Find, args *IUnknown) uintptr {
	return this.impl.FindMatchCountChanged(sender, args)
}

type _ICoreWebView2FindMatchCountChangedEventHandlerImpl interface {
	IUnknownImpl
	FindMatchCountChanged(sender *ICoreWebView2Find, args *IUnknown) uintptr
}

var _ICoreWebView2FindMatchCountChangedEventHandlerFn *_ICoreWebView2FindMatchCountChangedEventHandlerVtbl

func NewICoreWebView2FindMatchCountChangedEventHandler(impl _ICoreWebView2FindMatchCountChangedEventHandlerImpl) *ICoreWebView2FindMatchCountChangedEventHandler {
	if _ICoreWebView2FindMatchCountChangedEventHandlerFn == nil {
		_ICoreWebView2FindMatchCountChangedEventHandlerFn = &_ICoreWebView2FindMatchCountChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FindMatchCountChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FindMatchCountChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FindMatchCountChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FindMatchCountChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FindMatchCountChangedEventHandler{
		vtbl: _ICoreWebView2FindMatchCountChangedEventHandlerFn,
		impl: impl,
	}
}
