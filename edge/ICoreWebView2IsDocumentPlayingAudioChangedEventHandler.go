package edge

import "unsafe"

type _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2IsDocumentPlayingAudioChangedEventHandler 接收 IsDocumentPlayingAudioChanged 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2isdocumentplayingaudiochangedeventhandler
type ICoreWebView2IsDocumentPlayingAudioChangedEventHandler struct {
	vtbl *_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl
	impl _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl
}

func (i *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownAddRef(
	this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownRelease(
	this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInvoke(
	this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler,
	sender *ICoreWebView2,
	args *IUnknown,
) uintptr {
	return this.impl.DocumentPlayingAudioChanged(sender, args)
}

type _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl interface {
	IUnknownImpl
	DocumentPlayingAudioChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn *_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl

func NewICoreWebView2IsDocumentPlayingAudioChangedEventHandler(
	impl _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl,
) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler {
	if _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn == nil {
		_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn = &_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2IsDocumentPlayingAudioChangedEventHandler{
		vtbl: _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn,
		impl: impl,
	}
}
