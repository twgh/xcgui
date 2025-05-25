package edge

type _ICoreWebView2DocumentTitleChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2DocumentTitleChangedEventHandler 接收 DocumentTitleChanged 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2documenttitlechangedeventhandler
type ICoreWebView2DocumentTitleChangedEventHandler struct {
	vtbl *_ICoreWebView2DocumentTitleChangedEventHandlerVtbl
	impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DocumentTitleChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerInvoke(this *ICoreWebView2DocumentTitleChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.DocumentTitleChanged(sender, args)
}

type _ICoreWebView2DocumentTitleChangedEventHandlerImpl interface {
	IUnknownImpl
	DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2DocumentTitleChangedEventHandlerFn = _ICoreWebView2DocumentTitleChangedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerInvoke),
}

func NewICoreWebView2DocumentTitleChangedEventHandler(impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl) *ICoreWebView2DocumentTitleChangedEventHandler {
	return &ICoreWebView2DocumentTitleChangedEventHandler{
		vtbl: &_ICoreWebView2DocumentTitleChangedEventHandlerFn,
		impl: impl,
	}
}
