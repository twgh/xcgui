package edge

type _ICoreWebView2ContentLoadingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ContentLoadingEventHandler 接收 ContentLoading 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventhandler
type ICoreWebView2ContentLoadingEventHandler struct {
	vtbl *_ICoreWebView2ContentLoadingEventHandlerVtbl
	impl _ICoreWebView2ContentLoadingEventHandlerImpl
}

func _ICoreWebView2ContentLoadingEventHandlerIUnknownQueryInterface(
	this *ICoreWebView2ContentLoadingEventHandler,
	refiid, object uintptr,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContentLoadingEventHandlerIUnknownAddRef(
	this *ICoreWebView2ContentLoadingEventHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContentLoadingEventHandlerIUnknownRelease(
	this *ICoreWebView2ContentLoadingEventHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContentLoadingEventHandlerInvoke(
	this *ICoreWebView2ContentLoadingEventHandler,
	sender *ICoreWebView2,
	args *ICoreWebView2ContentLoadingEventArgs,
) uintptr {
	return this.impl.ContentLoading(sender, args)
}

type _ICoreWebView2ContentLoadingEventHandlerImpl interface {
	IUnknownImpl
	ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr
}

var _ICoreWebView2ContentLoadingEventHandlerFn = _ICoreWebView2ContentLoadingEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ContentLoadingEventHandlerInvoke),
}

func NewICoreWebView2ContentLoadingEventHandler(
	impl _ICoreWebView2ContentLoadingEventHandlerImpl,
) *ICoreWebView2ContentLoadingEventHandler {
	return &ICoreWebView2ContentLoadingEventHandler{
		vtbl: &_ICoreWebView2ContentLoadingEventHandlerFn,
		impl: impl,
	}
}
