package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ExecuteScriptCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ExecuteScriptCompletedHandler 接收 ExecuteScript 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2executescriptcompletedhandler
type ICoreWebView2ExecuteScriptCompletedHandler struct {
	vtbl *_ICoreWebView2ExecuteScriptCompletedHandlerVtbl
	impl _ICoreWebView2ExecuteScriptCompletedHandlerImpl
}

func (i *ICoreWebView2ExecuteScriptCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ExecuteScriptCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ExecuteScriptCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownAddRef(this *ICoreWebView2ExecuteScriptCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownRelease(this *ICoreWebView2ExecuteScriptCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ExecuteScriptCompletedHandlerInvoke(this *ICoreWebView2ExecuteScriptCompletedHandler, errorCode syscall.Errno, result *uint16) uintptr {
	return this.impl.ExecuteScriptCompleted(errorCode, result)
}

type _ICoreWebView2ExecuteScriptCompletedHandlerImpl interface {
	IUnknownImpl
	ExecuteScriptCompleted(errorCode syscall.Errno, result *uint16) uintptr
}

var _ICoreWebView2ExecuteScriptCompletedHandlerFn *_ICoreWebView2ExecuteScriptCompletedHandlerVtbl

func NewICoreWebView2ExecuteScriptCompletedHandler(impl _ICoreWebView2ExecuteScriptCompletedHandlerImpl) *ICoreWebView2ExecuteScriptCompletedHandler {
	if _ICoreWebView2ExecuteScriptCompletedHandlerFn == nil {
		_ICoreWebView2ExecuteScriptCompletedHandlerFn = &_ICoreWebView2ExecuteScriptCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ExecuteScriptCompletedHandler{
		vtbl: _ICoreWebView2ExecuteScriptCompletedHandlerFn,
		impl: impl,
	}
}
