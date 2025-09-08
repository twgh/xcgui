package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2FindStartCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2FindStartCompletedHandler 接收查找开始完成事件。
//
// https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2findstartcompletedhandler
type ICoreWebView2FindStartCompletedHandler struct {
	vtbl *_ICoreWebView2FindStartCompletedHandlerVtbl
	impl _ICoreWebView2FindStartCompletedHandlerImpl
}

func (i *ICoreWebView2FindStartCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FindStartCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FindStartCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2FindStartCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FindStartCompletedHandlerIUnknownAddRef(this *ICoreWebView2FindStartCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FindStartCompletedHandlerIUnknownRelease(this *ICoreWebView2FindStartCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FindStartCompletedHandlerInvoke(this *ICoreWebView2FindStartCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.FindStartCompleted(errorCode)
}

type _ICoreWebView2FindStartCompletedHandlerImpl interface {
	IUnknownImpl
	FindStartCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2FindStartCompletedHandlerFn *_ICoreWebView2FindStartCompletedHandlerVtbl

func NewICoreWebView2FindStartCompletedHandler(impl _ICoreWebView2FindStartCompletedHandlerImpl) *ICoreWebView2FindStartCompletedHandler {
	if _ICoreWebView2FindStartCompletedHandlerFn == nil {
		_ICoreWebView2FindStartCompletedHandlerFn = &_ICoreWebView2FindStartCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FindStartCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FindStartCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FindStartCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FindStartCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2FindStartCompletedHandler{
		vtbl: _ICoreWebView2FindStartCompletedHandlerFn,
		impl: impl,
	}
}
