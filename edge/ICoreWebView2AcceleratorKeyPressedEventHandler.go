package edge

import "unsafe"

type _ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2AcceleratorKeyPressedEventHandler 接收按键事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventhandler
type ICoreWebView2AcceleratorKeyPressedEventHandler struct {
	vtbl *_ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl
	impl _ICoreWebView2AcceleratorKeyPressedEventHandlerImpl
}

func (i *ICoreWebView2AcceleratorKeyPressedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AcceleratorKeyPressedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownQueryInterface(this *ICoreWebView2AcceleratorKeyPressedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownAddRef(this *ICoreWebView2AcceleratorKeyPressedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownRelease(this *ICoreWebView2AcceleratorKeyPressedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2AcceleratorKeyPressedEventHandlerInvoke(this *ICoreWebView2AcceleratorKeyPressedEventHandler, sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	return this.impl.AcceleratorKeyPressed(sender, args)
}

type _ICoreWebView2AcceleratorKeyPressedEventHandlerImpl interface {
	IUnknownImpl
	AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr
}

var _ICoreWebView2AcceleratorKeyPressedEventHandlerFn *_ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl

func NewICoreWebView2AcceleratorKeyPressedEventHandler(impl _ICoreWebView2AcceleratorKeyPressedEventHandlerImpl) *ICoreWebView2AcceleratorKeyPressedEventHandler {
	if _ICoreWebView2AcceleratorKeyPressedEventHandlerFn == nil {
		_ICoreWebView2AcceleratorKeyPressedEventHandlerFn = &_ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2AcceleratorKeyPressedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2AcceleratorKeyPressedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2AcceleratorKeyPressedEventHandler{
		vtbl: _ICoreWebView2AcceleratorKeyPressedEventHandlerFn,
		impl: impl,
	}
}
