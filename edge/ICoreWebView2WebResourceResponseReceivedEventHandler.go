package edge

import "unsafe"

type _ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2WebResourceResponseReceivedEventHandler 接收 WebResourceResponseReceived 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventhandler
type ICoreWebView2WebResourceResponseReceivedEventHandler struct {
	vtbl *_ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl
	impl _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl
}

func (i *ICoreWebView2WebResourceResponseReceivedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseReceivedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebResourceResponseReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2WebResourceResponseReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownRelease(this *ICoreWebView2WebResourceResponseReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerInvoke(this *ICoreWebView2WebResourceResponseReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
	return this.impl.WebResourceResponseReceived(sender, args)
}

type _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl interface {
	IUnknownImpl
	WebResourceResponseReceived(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr
}

var _ICoreWebView2WebResourceResponseReceivedEventHandlerFn = _ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerInvoke),
}

func NewICoreWebView2WebResourceResponseReceivedEventHandler(impl _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl) *ICoreWebView2WebResourceResponseReceivedEventHandler {
	return &ICoreWebView2WebResourceResponseReceivedEventHandler{
		vtbl: &_ICoreWebView2WebResourceResponseReceivedEventHandlerFn,
		impl: impl,
	}
}
