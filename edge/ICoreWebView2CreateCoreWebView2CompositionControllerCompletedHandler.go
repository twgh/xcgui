package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler 接收 CreateCoreWebView2CompositionControllerCompleted 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2createcorewebview2compositioncontrollercompletedhandler
type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler struct {
	vtbl *_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl
	impl _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl
}

func (i *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownQueryInterface(
	this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownAddRef(
	this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownRelease(
	this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInvoke(
	this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler,
	errorCode syscall.Errno,
	result *ICoreWebView2CompositionController,
) uintptr {
	return this.impl.CreateCoreWebView2CompositionControllerCompleted(errorCode, result)
}

type _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl interface {
	IUnknownImpl
	CreateCoreWebView2CompositionControllerCompleted(errorCode syscall.Errno, result *ICoreWebView2CompositionController) uintptr
}

var _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn *_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl

func NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler(
	impl _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl,
) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
	if _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn == nil {
		_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn = &_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler{
		vtbl: _ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn,
		impl: impl,
	}
}
