package edge

import "unsafe"

type _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler 接收默认下载对话框打开状态变化事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2isdefaultdownloaddialogopenchangedeventhandler
type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler struct {
	vtbl *_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl
	impl _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl
}

func (i *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownAddRef(this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownRelease(this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInvoke(this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.IsDefaultDownloadDialogOpenChanged(sender, args)
}

type _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl interface {
	IUnknownImpl
	IsDefaultDownloadDialogOpenChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerFn *_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl

func NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler(impl _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl) *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler {
	if _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerFn == nil {
		_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerFn = &_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler{
		vtbl: _ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerFn,
		impl: impl,
	}
}