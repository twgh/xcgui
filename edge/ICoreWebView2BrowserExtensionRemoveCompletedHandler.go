package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2BrowserExtensionRemoveCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2BrowserExtensionRemoveCompletedHandler 接收浏览器扩展移除完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionremovecompletedhandler
type ICoreWebView2BrowserExtensionRemoveCompletedHandler struct {
	vtbl *_ICoreWebView2BrowserExtensionRemoveCompletedHandlerVtbl
	impl _ICoreWebView2BrowserExtensionRemoveCompletedHandlerImpl
}

func (i *ICoreWebView2BrowserExtensionRemoveCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtensionRemoveCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2BrowserExtensionRemoveCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownAddRef(this *ICoreWebView2BrowserExtensionRemoveCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownRelease(this *ICoreWebView2BrowserExtensionRemoveCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BrowserExtensionRemoveCompletedHandlerInvoke(this *ICoreWebView2BrowserExtensionRemoveCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.BrowserExtensionRemoveCompleted(errorCode)
}

type _ICoreWebView2BrowserExtensionRemoveCompletedHandlerImpl interface {
	IUnknownImpl
	BrowserExtensionRemoveCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2BrowserExtensionRemoveCompletedHandlerFn *_ICoreWebView2BrowserExtensionRemoveCompletedHandlerVtbl

func NewICoreWebView2BrowserExtensionRemoveCompletedHandler(impl _ICoreWebView2BrowserExtensionRemoveCompletedHandlerImpl) *ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	if _ICoreWebView2BrowserExtensionRemoveCompletedHandlerFn == nil {
		_ICoreWebView2BrowserExtensionRemoveCompletedHandlerFn = &_ICoreWebView2BrowserExtensionRemoveCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2BrowserExtensionRemoveCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2BrowserExtensionRemoveCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2BrowserExtensionRemoveCompletedHandler{
		vtbl: _ICoreWebView2BrowserExtensionRemoveCompletedHandlerFn,
		impl: impl,
	}
}