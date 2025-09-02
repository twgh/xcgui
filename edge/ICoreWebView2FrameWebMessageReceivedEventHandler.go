package edge

import "unsafe"

type _ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FrameWebMessageReceivedEventHandler 接收框架 Web 消息事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framewebmessagereceivedeventhandler
type ICoreWebView2FrameWebMessageReceivedEventHandler struct {
	vtbl *_ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl
	impl _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl
}

func (i *ICoreWebView2FrameWebMessageReceivedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameWebMessageReceivedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameWebMessageReceivedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameWebMessageReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownRelease(this *ICoreWebView2FrameWebMessageReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerInvoke(this *ICoreWebView2FrameWebMessageReceivedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return this.impl.FrameWebMessageReceived(sender, args)
}

type _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl interface {
	IUnknownImpl
	FrameWebMessageReceived(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
}

var _ICoreWebView2FrameWebMessageReceivedEventHandlerFn *_ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl

func NewICoreWebView2FrameWebMessageReceivedEventHandler(impl _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl) *ICoreWebView2FrameWebMessageReceivedEventHandler {
	if _ICoreWebView2FrameWebMessageReceivedEventHandlerFn == nil {
		_ICoreWebView2FrameWebMessageReceivedEventHandlerFn = &_ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FrameWebMessageReceivedEventHandler{
		vtbl: _ICoreWebView2FrameWebMessageReceivedEventHandlerFn,
		impl: impl,
	}
}