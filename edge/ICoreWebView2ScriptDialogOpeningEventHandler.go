package edge

import "unsafe"

type _ICoreWebView2ScriptDialogOpeningEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ScriptDialogOpeningEventHandler 接收 ScriptDialogOpening 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventhandler
type ICoreWebView2ScriptDialogOpeningEventHandler struct {
	vtbl *_ICoreWebView2ScriptDialogOpeningEventHandlerVtbl
	impl _ICoreWebView2ScriptDialogOpeningEventHandlerImpl
}

func (i *ICoreWebView2ScriptDialogOpeningEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ScriptDialogOpeningEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownQueryInterface(this *ICoreWebView2ScriptDialogOpeningEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownAddRef(this *ICoreWebView2ScriptDialogOpeningEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownRelease(this *ICoreWebView2ScriptDialogOpeningEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerInvoke(this *ICoreWebView2ScriptDialogOpeningEventHandler, sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
	return this.impl.ScriptDialogOpening(sender, args)
}

type _ICoreWebView2ScriptDialogOpeningEventHandlerImpl interface {
	IUnknownImpl
	ScriptDialogOpening(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr
}

var _ICoreWebView2ScriptDialogOpeningEventHandlerFn = _ICoreWebView2ScriptDialogOpeningEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerInvoke),
}

func NewICoreWebView2ScriptDialogOpeningEventHandler(impl _ICoreWebView2ScriptDialogOpeningEventHandlerImpl) *ICoreWebView2ScriptDialogOpeningEventHandler {
	return &ICoreWebView2ScriptDialogOpeningEventHandler{
		vtbl: &_ICoreWebView2ScriptDialogOpeningEventHandlerFn,
		impl: impl,
	}
}
