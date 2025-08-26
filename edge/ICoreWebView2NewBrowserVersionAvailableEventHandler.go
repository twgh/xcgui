package edge

import "unsafe"

type _ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NewBrowserVersionAvailableEventHandler 接收新浏览器版本可用事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2newbrowserversionavailableeventhandler
type ICoreWebView2NewBrowserVersionAvailableEventHandler struct {
	vtbl *_ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl
	impl _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl
}

func (i *ICoreWebView2NewBrowserVersionAvailableEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NewBrowserVersionAvailableEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownQueryInterface(this *ICoreWebView2NewBrowserVersionAvailableEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownAddRef(this *ICoreWebView2NewBrowserVersionAvailableEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownRelease(this *ICoreWebView2NewBrowserVersionAvailableEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerInvoke(this *ICoreWebView2NewBrowserVersionAvailableEventHandler, sender *ICoreWebView2Environment, args *IUnknown) uintptr {
	return this.impl.NewBrowserVersionAvailable(sender, args)
}

type _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl interface {
	IUnknownImpl
	NewBrowserVersionAvailable(sender *ICoreWebView2Environment, args *IUnknown) uintptr
}

var _ICoreWebView2NewBrowserVersionAvailableEventHandlerFn *_ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl

func NewICoreWebView2NewBrowserVersionAvailableEventHandler(impl _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl) *ICoreWebView2NewBrowserVersionAvailableEventHandler {
	if _ICoreWebView2NewBrowserVersionAvailableEventHandlerFn == nil {
		_ICoreWebView2NewBrowserVersionAvailableEventHandlerFn = &_ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerInvoke),
		}
	}
	return &ICoreWebView2NewBrowserVersionAvailableEventHandler{
		vtbl: _ICoreWebView2NewBrowserVersionAvailableEventHandlerFn,
		impl: impl,
	}
}