package edge

import "unsafe"

type _ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2EstimatedEndTimeChangedEventHandler 接收预计结束时间改变事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2estimatedendtimechangedeventhandler
type ICoreWebView2EstimatedEndTimeChangedEventHandler struct {
	vtbl *_ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl
	impl _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl
}

func (i *ICoreWebView2EstimatedEndTimeChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EstimatedEndTimeChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2EstimatedEndTimeChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownAddRef(this *ICoreWebView2EstimatedEndTimeChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownRelease(this *ICoreWebView2EstimatedEndTimeChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerInvoke(this *ICoreWebView2EstimatedEndTimeChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	return this.impl.EstimatedEndTimeChanged(sender, args)
}

type _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl interface {
	IUnknownImpl
	EstimatedEndTimeChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr
}

var _ICoreWebView2EstimatedEndTimeChangedEventHandlerFn *_ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl

func NewICoreWebView2EstimatedEndTimeChangedEventHandler(impl _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl) *ICoreWebView2EstimatedEndTimeChangedEventHandler {
	if _ICoreWebView2EstimatedEndTimeChangedEventHandlerFn == nil {
		_ICoreWebView2EstimatedEndTimeChangedEventHandlerFn = &_ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2EstimatedEndTimeChangedEventHandler{
		vtbl: _ICoreWebView2EstimatedEndTimeChangedEventHandlerFn,
		impl: impl,
	}
}
