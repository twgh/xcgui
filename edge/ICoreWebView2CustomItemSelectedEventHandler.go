package edge

import "unsafe"

type _ICoreWebView2CustomItemSelectedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CustomItemSelectedEventHandler 接收 CustomItemSelected 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2customitemselectedeventhandler
type ICoreWebView2CustomItemSelectedEventHandler struct {
	vtbl *_ICoreWebView2CustomItemSelectedEventHandlerVtbl
	impl _ICoreWebView2CustomItemSelectedEventHandlerImpl
}

func (i *ICoreWebView2CustomItemSelectedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CustomItemSelectedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2CustomItemSelectedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownAddRef(
	this *ICoreWebView2CustomItemSelectedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownRelease(
	this *ICoreWebView2CustomItemSelectedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CustomItemSelectedEventHandlerInvoke(
	this *ICoreWebView2CustomItemSelectedEventHandler,
	sender *ICoreWebView2ContextMenuItem,
	args *IUnknown,
) uintptr {
	return this.impl.CustomItemSelected(sender, args)
}

type _ICoreWebView2CustomItemSelectedEventHandlerImpl interface {
	IUnknownImpl
	CustomItemSelected(sender *ICoreWebView2ContextMenuItem, args *IUnknown) uintptr
}

var _ICoreWebView2CustomItemSelectedEventHandlerFn *_ICoreWebView2CustomItemSelectedEventHandlerVtbl

func NewICoreWebView2CustomItemSelectedEventHandler(
	impl _ICoreWebView2CustomItemSelectedEventHandlerImpl,
) *ICoreWebView2CustomItemSelectedEventHandler {
	if _ICoreWebView2CustomItemSelectedEventHandlerFn == nil {
		_ICoreWebView2CustomItemSelectedEventHandlerFn = &_ICoreWebView2CustomItemSelectedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2CustomItemSelectedEventHandler{
		vtbl: _ICoreWebView2CustomItemSelectedEventHandlerFn,
		impl: impl,
	}
}
