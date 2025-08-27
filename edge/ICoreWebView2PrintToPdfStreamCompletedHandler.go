package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2PrintToPdfStreamCompletedHandler 接收 PrintToPdfStream 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler
type ICoreWebView2PrintToPdfStreamCompletedHandler struct {
	vtbl *_ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl
	impl _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl
}

func (i *ICoreWebView2PrintToPdfStreamCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PrintToPdfStreamCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintToPdfStreamCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintToPdfStreamCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintToPdfStreamCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerInvoke(this *ICoreWebView2PrintToPdfStreamCompletedHandler, errorCode syscall.Errno, pdfStream *IStream) uintptr {
	return this.impl.PrintToPdfStreamCompleted(errorCode, pdfStream)
}

type _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl interface {
	IUnknownImpl
	PrintToPdfStreamCompleted(errorCode syscall.Errno, pdfStream *IStream) uintptr
}

var _ICoreWebView2PrintToPdfStreamCompletedHandlerFn *_ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl

func NewICoreWebView2PrintToPdfStreamCompletedHandler(impl _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl) *ICoreWebView2PrintToPdfStreamCompletedHandler {
	if _ICoreWebView2PrintToPdfStreamCompletedHandlerFn == nil {
		_ICoreWebView2PrintToPdfStreamCompletedHandlerFn = &_ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2PrintToPdfStreamCompletedHandler{
		vtbl: _ICoreWebView2PrintToPdfStreamCompletedHandlerFn,
		impl: impl,
	}
}
