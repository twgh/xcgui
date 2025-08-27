package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ExecuteScriptWithResultCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ExecuteScriptWithResultCompletedHandler 接收脚本执行结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2executescriptwithresultcompletedhandler
type ICoreWebView2ExecuteScriptWithResultCompletedHandler struct {
	vtbl *_ICoreWebView2ExecuteScriptWithResultCompletedHandlerVtbl
	impl _ICoreWebView2ExecuteScriptWithResultCompletedHandlerImpl
}

func (i *ICoreWebView2ExecuteScriptWithResultCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ExecuteScriptWithResultCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ExecuteScriptWithResultCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownAddRef(this *ICoreWebView2ExecuteScriptWithResultCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownRelease(this *ICoreWebView2ExecuteScriptWithResultCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ExecuteScriptWithResultCompletedHandlerInvoke(this *ICoreWebView2ExecuteScriptWithResultCompletedHandler, errorCode syscall.Errno, result *ICoreWebView2ExecuteScriptResult) uintptr {
	return this.impl.ExecuteScriptWithResultCompleted(errorCode, result)
}

type _ICoreWebView2ExecuteScriptWithResultCompletedHandlerImpl interface {
	IUnknownImpl
	ExecuteScriptWithResultCompleted(errorCode syscall.Errno, result *ICoreWebView2ExecuteScriptResult) uintptr
}

var _ICoreWebView2ExecuteScriptWithResultCompletedHandlerFn *_ICoreWebView2ExecuteScriptWithResultCompletedHandlerVtbl

func NewICoreWebView2ExecuteScriptWithResultCompletedHandler(impl _ICoreWebView2ExecuteScriptWithResultCompletedHandlerImpl) *ICoreWebView2ExecuteScriptWithResultCompletedHandler {
	if _ICoreWebView2ExecuteScriptWithResultCompletedHandlerFn == nil {
		_ICoreWebView2ExecuteScriptWithResultCompletedHandlerFn = &_ICoreWebView2ExecuteScriptWithResultCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ExecuteScriptWithResultCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ExecuteScriptWithResultCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ExecuteScriptWithResultCompletedHandler{
		vtbl: _ICoreWebView2ExecuteScriptWithResultCompletedHandlerFn,
		impl: impl,
	}
}