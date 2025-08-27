package edge

import "unsafe"

type _ICoreWebView2StatusBarTextChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2StatusBarTextChangedEventHandler 接收状态栏文本改变事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2statusbartextchangedeventhandler
type ICoreWebView2StatusBarTextChangedEventHandler struct {
	vtbl *_ICoreWebView2StatusBarTextChangedEventHandlerVtbl
	impl _ICoreWebView2StatusBarTextChangedEventHandlerImpl
}

func (i *ICoreWebView2StatusBarTextChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2StatusBarTextChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2StatusBarTextChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownAddRef(this *ICoreWebView2StatusBarTextChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownRelease(this *ICoreWebView2StatusBarTextChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2StatusBarTextChangedEventHandlerInvoke(this *ICoreWebView2StatusBarTextChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.StatusBarTextChanged(sender, args)
}

type _ICoreWebView2StatusBarTextChangedEventHandlerImpl interface {
	IUnknownImpl
	StatusBarTextChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2StatusBarTextChangedEventHandlerFn *_ICoreWebView2StatusBarTextChangedEventHandlerVtbl

func NewICoreWebView2StatusBarTextChangedEventHandler(impl _ICoreWebView2StatusBarTextChangedEventHandlerImpl) *ICoreWebView2StatusBarTextChangedEventHandler {
	if _ICoreWebView2StatusBarTextChangedEventHandlerFn == nil {
		_ICoreWebView2StatusBarTextChangedEventHandlerFn = &_ICoreWebView2StatusBarTextChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2StatusBarTextChangedEventHandler{
		vtbl: _ICoreWebView2StatusBarTextChangedEventHandlerFn,
		impl: impl,
	}
}
