package edge

import "unsafe"

type iCoreWebView2WebMessageReceivedEventHandlerImpl interface {
	IUnknownImpl
	WebMessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
}

type iCoreWebView2WebMessageReceivedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2WebMessageReceivedEventHandler 接收 WebMessageReceived 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventhandler
type ICoreWebView2WebMessageReceivedEventHandler struct {
	vtbl *iCoreWebView2WebMessageReceivedEventHandlerVtbl
	impl iCoreWebView2WebMessageReceivedEventHandlerImpl
}

func (i *ICoreWebView2WebMessageReceivedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebMessageReceivedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebMessageReceivedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease(this *ICoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebMessageReceivedEventHandlerInvoke(this *ICoreWebView2WebMessageReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return this.impl.WebMessageReceived(sender, args)
}

var iCoreWebView2WebMessageReceivedEventHandlerFn *iCoreWebView2WebMessageReceivedEventHandlerVtbl

func NewICoreWebView2WebMessageReceivedEventHandler(impl iCoreWebView2WebMessageReceivedEventHandlerImpl) *ICoreWebView2WebMessageReceivedEventHandler {
	if iCoreWebView2WebMessageReceivedEventHandlerFn == nil {
		iCoreWebView2WebMessageReceivedEventHandlerFn = &iCoreWebView2WebMessageReceivedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2WebMessageReceivedEventHandler{
		vtbl: iCoreWebView2WebMessageReceivedEventHandlerFn,
		impl: impl,
	}
}
