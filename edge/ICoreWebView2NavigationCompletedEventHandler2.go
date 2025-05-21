package edge

type _ICoreWebView2NavigationCompletedEventHandler2Vtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NavigationCompletedEventHandler2 接收 NavigationCompleted2 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2NavigationCompletedEventHandler
type ICoreWebView2NavigationCompletedEventHandler2 struct {
	vtbl *_ICoreWebView2NavigationCompletedEventHandler2Vtbl
	impl _ICoreWebView2NavigationCompletedEventHandler2Impl
}

func _ICoreWebView2NavigationCompletedEventHandler2IUnknownQueryInterface(this *ICoreWebView2NavigationCompletedEventHandler2, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NavigationCompletedEventHandler2IUnknownAddRef(this *ICoreWebView2NavigationCompletedEventHandler2) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NavigationCompletedEventHandler2IUnknownRelease(this *ICoreWebView2NavigationCompletedEventHandler2) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NavigationCompletedEventHandler2Invoke(this *ICoreWebView2NavigationCompletedEventHandler2, sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return this.impl.NavigationCompleted2(sender, args)
}

type _ICoreWebView2NavigationCompletedEventHandler2Impl interface {
	IUnknownImpl
	NavigationCompleted2(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
}

var _ICoreWebView2NavigationCompletedEventHandler2Fn = _ICoreWebView2NavigationCompletedEventHandler2Vtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2NavigationCompletedEventHandler2IUnknownQueryInterface),
		NewComProc(_ICoreWebView2NavigationCompletedEventHandler2IUnknownAddRef),
		NewComProc(_ICoreWebView2NavigationCompletedEventHandler2IUnknownRelease),
	},
	NewComProc(_ICoreWebView2NavigationCompletedEventHandler2Invoke),
}

func NewICoreWebView2NavigationCompletedEventHandler2(impl _ICoreWebView2NavigationCompletedEventHandler2Impl) *ICoreWebView2NavigationCompletedEventHandler2 {
	return &ICoreWebView2NavigationCompletedEventHandler2{
		vtbl: &_ICoreWebView2NavigationCompletedEventHandler2Fn,
		impl: impl,
	}
}
