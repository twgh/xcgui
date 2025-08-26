package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2PrintToPdfCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2PrintToPdfCompletedHandler 接收打印到PDF操作完成的通知。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfcompletedhandler
type ICoreWebView2PrintToPdfCompletedHandler struct {
	vtbl *_ICoreWebView2PrintToPdfCompletedHandlerVtbl
	impl _ICoreWebView2PrintToPdfCompletedHandlerImpl
}

func (i *ICoreWebView2PrintToPdfCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PrintToPdfCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintToPdfCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintToPdfCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintToPdfCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintToPdfCompletedHandlerInvoke(this *ICoreWebView2PrintToPdfCompletedHandler, errorCode syscall.Errno, isSuccessful uintptr) uintptr {
	return this.impl.PrintToPdfCompleted(errorCode, isSuccessful != 0)
}

type _ICoreWebView2PrintToPdfCompletedHandlerImpl interface {
	IUnknownImpl
	PrintToPdfCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr
}

var _ICoreWebView2PrintToPdfCompletedHandlerFn *_ICoreWebView2PrintToPdfCompletedHandlerVtbl

func NewICoreWebView2PrintToPdfCompletedHandler(impl _ICoreWebView2PrintToPdfCompletedHandlerImpl) *ICoreWebView2PrintToPdfCompletedHandler {
	if _ICoreWebView2PrintToPdfCompletedHandlerFn == nil {
		_ICoreWebView2PrintToPdfCompletedHandlerFn = &_ICoreWebView2PrintToPdfCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2PrintToPdfCompletedHandler{
		vtbl: _ICoreWebView2PrintToPdfCompletedHandlerFn,
		impl: impl,
	}
}
