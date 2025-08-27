package edge

import "unsafe"

type _ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2LaunchingExternalUriSchemeEventHandler 接收启动外部 URI 方案事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventhandler
type ICoreWebView2LaunchingExternalUriSchemeEventHandler struct {
	vtbl *_ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl
	impl _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownQueryInterface(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownAddRef(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownRelease(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerInvoke(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler, sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr {
	return this.impl.LaunchingExternalUriScheme(sender, args)
}

type _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl interface {
	IUnknownImpl
	LaunchingExternalUriScheme(sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr
}

var _ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn *_ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl

func NewICoreWebView2LaunchingExternalUriSchemeEventHandler(impl _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl) *ICoreWebView2LaunchingExternalUriSchemeEventHandler {
	if _ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn == nil {
		_ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn = &_ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerInvoke),
		}
	}
	return &ICoreWebView2LaunchingExternalUriSchemeEventHandler{
		vtbl: _ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn,
		impl: impl,
	}
}
