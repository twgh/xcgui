package edge

import "unsafe"

type _ICoreWebView2HistoryChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2HistoryChangedEventHandler 处理历史记录改变事件
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2historychangedeventhandler
type ICoreWebView2HistoryChangedEventHandler struct {
	vtbl *_ICoreWebView2HistoryChangedEventHandlerVtbl
	impl _ICoreWebView2HistoryChangedEventHandlerImpl
}

func (i *ICoreWebView2HistoryChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2HistoryChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2HistoryChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2HistoryChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2HistoryChangedEventHandlerIUnknownAddRef(this *ICoreWebView2HistoryChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2HistoryChangedEventHandlerIUnknownRelease(this *ICoreWebView2HistoryChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2HistoryChangedEventHandlerInvoke(this *ICoreWebView2HistoryChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.HistoryChanged(sender, args)
}

type _ICoreWebView2HistoryChangedEventHandlerImpl interface {
	IUnknownImpl
	HistoryChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2HistoryChangedEventHandlerFn = _ICoreWebView2HistoryChangedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2HistoryChangedEventHandlerInvoke),
}

func NewICoreWebView2HistoryChangedEventHandler(impl _ICoreWebView2HistoryChangedEventHandlerImpl) *ICoreWebView2HistoryChangedEventHandler {
	return &ICoreWebView2HistoryChangedEventHandler{
		vtbl: &_ICoreWebView2HistoryChangedEventHandlerFn,
		impl: impl,
	}
}
