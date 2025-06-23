package edge

import "unsafe"

type _ICoreWebView2BrowserProcessExitedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2BrowserProcessExitedEventHandler 接收浏览器进程退出事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventhandler
type ICoreWebView2BrowserProcessExitedEventHandler struct {
	vtbl *_ICoreWebView2BrowserProcessExitedEventHandlerVtbl
	impl _ICoreWebView2BrowserProcessExitedEventHandlerImpl
}

func (i *ICoreWebView2BrowserProcessExitedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserProcessExitedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BrowserProcessExitedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownAddRef(this *ICoreWebView2BrowserProcessExitedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownRelease(this *ICoreWebView2BrowserProcessExitedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BrowserProcessExitedEventHandlerInvoke(this *ICoreWebView2BrowserProcessExitedEventHandler, sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr {
	return this.impl.BrowserProcessExited(sender, args)
}

type _ICoreWebView2BrowserProcessExitedEventHandlerImpl interface {
	IUnknownImpl
	BrowserProcessExited(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr
}

var _ICoreWebView2BrowserProcessExitedEventHandlerFn *_ICoreWebView2BrowserProcessExitedEventHandlerVtbl

func NewICoreWebView2BrowserProcessExitedEventHandler(impl _ICoreWebView2BrowserProcessExitedEventHandlerImpl) *ICoreWebView2BrowserProcessExitedEventHandler {
	if _ICoreWebView2BrowserProcessExitedEventHandlerFn == nil {
		_ICoreWebView2BrowserProcessExitedEventHandlerFn = &_ICoreWebView2BrowserProcessExitedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2BrowserProcessExitedEventHandler{
		vtbl: _ICoreWebView2BrowserProcessExitedEventHandlerFn,
		impl: impl,
	}
}
