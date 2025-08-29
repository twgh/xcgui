package edge

import "unsafe"

type _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2SaveFileSecurityCheckStartingEventHandler 接收保存文件安全检查开始事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2savefilesecuritycheckstartingeventhandler
type ICoreWebView2SaveFileSecurityCheckStartingEventHandler struct {
	vtbl *_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerVtbl
	impl _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerImpl
}

func (i *ICoreWebView2SaveFileSecurityCheckStartingEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SaveFileSecurityCheckStartingEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2SaveFileSecurityCheckStartingEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownAddRef(this *ICoreWebView2SaveFileSecurityCheckStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownRelease(this *ICoreWebView2SaveFileSecurityCheckStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerInvoke(this *ICoreWebView2SaveFileSecurityCheckStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr {
	return this.impl.SaveFileSecurityCheckStarting(sender, args)
}

type _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerImpl interface {
	IUnknownImpl
	SaveFileSecurityCheckStarting(sender *ICoreWebView2, args *ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr
}

var _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerFn *_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerVtbl

func NewICoreWebView2SaveFileSecurityCheckStartingEventHandler(impl _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerImpl) *ICoreWebView2SaveFileSecurityCheckStartingEventHandler {
	if _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerFn == nil {
		_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerFn = &_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2SaveFileSecurityCheckStartingEventHandlerInvoke),
		}
	}
	return &ICoreWebView2SaveFileSecurityCheckStartingEventHandler{
		vtbl: _ICoreWebView2SaveFileSecurityCheckStartingEventHandlerFn,
		impl: impl,
	}
}