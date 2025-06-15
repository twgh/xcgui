package edge

import "unsafe"

type _ICoreWebView2BytesReceivedChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2BytesReceivedChangedEventHandler 接收 BytesReceivedChanged 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2bytesreceivedchangedeventhandler
type ICoreWebView2BytesReceivedChangedEventHandler struct {
	vtbl *_ICoreWebView2BytesReceivedChangedEventHandlerVtbl
	impl _ICoreWebView2BytesReceivedChangedEventHandlerImpl
}

func (i *ICoreWebView2BytesReceivedChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BytesReceivedChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BytesReceivedChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownAddRef(this *ICoreWebView2BytesReceivedChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownRelease(this *ICoreWebView2BytesReceivedChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BytesReceivedChangedEventHandlerInvoke(this *ICoreWebView2BytesReceivedChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	return this.impl.BytesReceivedChanged(sender, args)
}

type _ICoreWebView2BytesReceivedChangedEventHandlerImpl interface {
	IUnknownImpl
	BytesReceivedChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr
}

var _ICoreWebView2BytesReceivedChangedEventHandlerFn *_ICoreWebView2BytesReceivedChangedEventHandlerVtbl

func NewICoreWebView2BytesReceivedChangedEventHandler(impl _ICoreWebView2BytesReceivedChangedEventHandlerImpl) *ICoreWebView2BytesReceivedChangedEventHandler {
	if _ICoreWebView2BytesReceivedChangedEventHandlerFn == nil {
		_ICoreWebView2BytesReceivedChangedEventHandlerFn = &_ICoreWebView2BytesReceivedChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2BytesReceivedChangedEventHandler{
		vtbl: _ICoreWebView2BytesReceivedChangedEventHandlerFn,
		impl: impl,
	}
}
