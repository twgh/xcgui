package edge

import "unsafe"

type _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DevToolsProtocolEventReceivedEventHandler 接收开发者工具协议事件.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventhandler
type ICoreWebView2DevToolsProtocolEventReceivedEventHandler struct {
	vtbl *_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl
	impl _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownRelease(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInvoke(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr {
	return this.impl.DevToolsProtocolEventReceived(sender, args)
}

type _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl interface {
	IUnknownImpl
	DevToolsProtocolEventReceived(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr
}

var _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerFn = _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInvoke),
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventHandler(impl _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl) *ICoreWebView2DevToolsProtocolEventReceivedEventHandler {
	return &ICoreWebView2DevToolsProtocolEventReceivedEventHandler{
		vtbl: &_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerFn,
		impl: impl,
	}
}
