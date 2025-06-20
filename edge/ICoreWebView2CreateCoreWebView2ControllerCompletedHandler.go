package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CreateCoreWebView2ControllerCompletedHandler 接收 CreateCoreWebView2Controller 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2createcorewebview2controllercompletedhandler
type ICoreWebView2CreateCoreWebView2ControllerCompletedHandler struct {
	vtbl *_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl
	impl _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownAddRef(this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownRelease(this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInvoke(this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler, errorCode syscall.Errno, createdController *ICoreWebView2Controller) uintptr {
	return this.impl.CreateCoreWebView2ControllerCompleted(errorCode, createdController)
}

type _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl interface {
	IUnknownImpl
	CreateCoreWebView2ControllerCompleted(errorCode syscall.Errno, createdController *ICoreWebView2Controller) uintptr
}

var _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn *_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl

func NewICoreWebView2CreateCoreWebView2ControllerCompletedHandler(impl _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl) *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler {
	if _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn == nil {
		_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn = &_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2CreateCoreWebView2ControllerCompletedHandler{
		vtbl: _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn,
		impl: impl,
	}
}
