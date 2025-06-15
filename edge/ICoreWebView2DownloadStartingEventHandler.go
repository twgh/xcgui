package edge

import "unsafe"

type _ICoreWebView2DownloadStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DownloadStartingEventHandler 接收下载开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventhandler
type ICoreWebView2DownloadStartingEventHandler struct {
	vtbl *_ICoreWebView2DownloadStartingEventHandlerVtbl
	impl _ICoreWebView2DownloadStartingEventHandlerImpl
}

func (i *ICoreWebView2DownloadStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2DownloadStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownRelease(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DownloadStartingEventHandlerInvoke(this *ICoreWebView2DownloadStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	return this.impl.DownloadStarting(sender, args)
}

type _ICoreWebView2DownloadStartingEventHandlerImpl interface {
	IUnknownImpl
	DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr
}

var _ICoreWebView2DownloadStartingEventHandlerFn *_ICoreWebView2DownloadStartingEventHandlerVtbl

func NewICoreWebView2DownloadStartingEventHandler(impl _ICoreWebView2DownloadStartingEventHandlerImpl) *ICoreWebView2DownloadStartingEventHandler {
	if _ICoreWebView2DownloadStartingEventHandlerFn == nil {
		_ICoreWebView2DownloadStartingEventHandlerFn = &_ICoreWebView2DownloadStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2DownloadStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2DownloadStartingEventHandler{
		vtbl: _ICoreWebView2DownloadStartingEventHandlerFn,
		impl: impl,
	}
}
