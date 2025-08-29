package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ShowSaveAsUICompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ShowSaveAsUICompletedHandler 接收"另存为"UI完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2showsaveasuicompletedhandler
type ICoreWebView2ShowSaveAsUICompletedHandler struct {
	vtbl *_ICoreWebView2ShowSaveAsUICompletedHandlerVtbl
	impl _ICoreWebView2ShowSaveAsUICompletedHandlerImpl
}

func (i *ICoreWebView2ShowSaveAsUICompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ShowSaveAsUICompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ShowSaveAsUICompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownAddRef(this *ICoreWebView2ShowSaveAsUICompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownRelease(this *ICoreWebView2ShowSaveAsUICompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ShowSaveAsUICompletedHandlerInvoke(this *ICoreWebView2ShowSaveAsUICompletedHandler, errorCode syscall.Errno, result COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr {
	return this.impl.ShowSaveAsUICompleted(errorCode, result)
}

type _ICoreWebView2ShowSaveAsUICompletedHandlerImpl interface {
	IUnknownImpl
	ShowSaveAsUICompleted(errorCode syscall.Errno, result COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr
}

var _ICoreWebView2ShowSaveAsUICompletedHandlerFn *_ICoreWebView2ShowSaveAsUICompletedHandlerVtbl

func NewICoreWebView2ShowSaveAsUICompletedHandler(impl _ICoreWebView2ShowSaveAsUICompletedHandlerImpl) *ICoreWebView2ShowSaveAsUICompletedHandler {
	if _ICoreWebView2ShowSaveAsUICompletedHandlerFn == nil {
		_ICoreWebView2ShowSaveAsUICompletedHandlerFn = &_ICoreWebView2ShowSaveAsUICompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ShowSaveAsUICompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ShowSaveAsUICompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ShowSaveAsUICompletedHandler{
		vtbl: _ICoreWebView2ShowSaveAsUICompletedHandlerFn,
		impl: impl,
	}
}
