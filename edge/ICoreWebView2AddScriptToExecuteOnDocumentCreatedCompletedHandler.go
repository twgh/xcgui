package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler 接收 ICoreWebView2.AddScriptToExecuteOnDocumentCreated 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2addscripttoexecuteondocumentcreatedcompletedhandler
type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler struct {
	vtbl *_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl
	impl _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl
}

func (i *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownAddRef(this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownRelease(this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInvoke(this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler, errorCode syscall.Errno, id *uint16) uintptr {
	return this.impl.AddScriptToExecuteOnDocumentCreatedCompleted(errorCode, id)
}

type _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl interface {
	IUnknownImpl
	AddScriptToExecuteOnDocumentCreatedCompleted(errorCode syscall.Errno, id *uint16) uintptr
}

var _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerFn *_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl

func NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(impl _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl) *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler {
	if _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerFn == nil {
		_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerFn = &_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler{
		vtbl: _ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerFn,
		impl: impl,
	}
}
