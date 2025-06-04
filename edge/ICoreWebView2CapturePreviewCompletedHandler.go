package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2CapturePreviewCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CapturePreviewCompletedHandler 接收 CapturePreview 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2capturepreviewcompletedhandler
type ICoreWebView2CapturePreviewCompletedHandler struct {
	vtbl *_ICoreWebView2CapturePreviewCompletedHandlerVtbl
	impl _ICoreWebView2CapturePreviewCompletedHandlerImpl
}

func (i *ICoreWebView2CapturePreviewCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CapturePreviewCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownQueryInterface(
	this *ICoreWebView2CapturePreviewCompletedHandler,
	refiid, object uintptr,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownAddRef(
	this *ICoreWebView2CapturePreviewCompletedHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownRelease(
	this *ICoreWebView2CapturePreviewCompletedHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CapturePreviewCompletedHandlerInvoke(
	this *ICoreWebView2CapturePreviewCompletedHandler,
	errorCode syscall.Errno,
) uintptr {
	return this.impl.CapturePreviewCompleted(errorCode)
}

type _ICoreWebView2CapturePreviewCompletedHandlerImpl interface {
	IUnknownImpl
	CapturePreviewCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2CapturePreviewCompletedHandlerFn = _ICoreWebView2CapturePreviewCompletedHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerInvoke),
}

func NewICoreWebView2CapturePreviewCompletedHandler(
	impl _ICoreWebView2CapturePreviewCompletedHandlerImpl,
) *ICoreWebView2CapturePreviewCompletedHandler {
	return &ICoreWebView2CapturePreviewCompletedHandler{
		vtbl: &_ICoreWebView2CapturePreviewCompletedHandlerFn,
		impl: impl,
	}
}
