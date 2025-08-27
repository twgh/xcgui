package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2PrintCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2PrintCompletedHandler 接收 Print 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler
type ICoreWebView2PrintCompletedHandler struct {
	vtbl *_ICoreWebView2PrintCompletedHandlerVtbl
	impl _ICoreWebView2PrintCompletedHandlerImpl
}

func (i *ICoreWebView2PrintCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PrintCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2PrintCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintCompletedHandlerInvoke(this *ICoreWebView2PrintCompletedHandler, errorCode syscall.Errno, result COREWEBVIEW2_PRINT_STATUS) uintptr {
	return this.impl.PrintCompleted(errorCode, result)
}

type _ICoreWebView2PrintCompletedHandlerImpl interface {
	IUnknownImpl
	PrintCompleted(errorCode syscall.Errno, result COREWEBVIEW2_PRINT_STATUS) uintptr
}

var _ICoreWebView2PrintCompletedHandlerFn *_ICoreWebView2PrintCompletedHandlerVtbl

func NewICoreWebView2PrintCompletedHandler(impl _ICoreWebView2PrintCompletedHandlerImpl) *ICoreWebView2PrintCompletedHandler {
	if _ICoreWebView2PrintCompletedHandlerFn == nil {
		_ICoreWebView2PrintCompletedHandlerFn = &_ICoreWebView2PrintCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2PrintCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2PrintCompletedHandler{
		vtbl: _ICoreWebView2PrintCompletedHandlerFn,
		impl: impl,
	}
}
