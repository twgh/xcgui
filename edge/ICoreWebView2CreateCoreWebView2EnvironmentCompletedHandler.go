package edge

import (
	"syscall"
	"unsafe"
)

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl interface {
	IUnknownImpl
	EnvironmentCompleted(result syscall.Errno, env *ICoreWebView2Environment) uintptr
}

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler 接收 CreateCoreWebView2Environment 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2createcorewebview2environmentcompletedhandler
type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler struct {
	vtbl *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl
	impl iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownAddRef(this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownRelease(this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke(this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler, result syscall.Errno, env *ICoreWebView2Environment) uintptr {
	return this.impl.EnvironmentCompleted(result, env)
}

var iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl

func NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(impl iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl) *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler {
	if iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn == nil {
		iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn = &iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler{
		vtbl: iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn,
		impl: impl,
	}
}
