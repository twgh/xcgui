package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2WebResourceResponseViewGetContentCompletedHandler 接收 ICoreWebView2WebResourceResponseView.GetContent 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler struct {
	vtbl *_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl
	impl _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl
}

func (i *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownAddRef(this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownRelease(this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInvoke(this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler, errorCode syscall.Errno, result *IStream) uintptr {
	return this.impl.WebResourceResponseViewGetContentCompleted(errorCode, result)
}

type _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl interface {
	IUnknownImpl
	WebResourceResponseViewGetContentCompleted(errorCode syscall.Errno, result *IStream) uintptr
}

var _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerFn = _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInvoke),
}

func NewICoreWebView2WebResourceResponseViewGetContentCompletedHandler(impl _ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl) *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	return &ICoreWebView2WebResourceResponseViewGetContentCompletedHandler{
		vtbl: &_ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerFn,
		impl: impl,
	}
}
