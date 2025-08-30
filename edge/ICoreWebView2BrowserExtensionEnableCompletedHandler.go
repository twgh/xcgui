package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2BrowserExtensionEnableCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2BrowserExtensionEnableCompletedHandler 接收浏览器扩展启用完成事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionenablecompletedhandler
type ICoreWebView2BrowserExtensionEnableCompletedHandler struct {
	vtbl *_ICoreWebView2BrowserExtensionEnableCompletedHandlerVtbl
	impl _ICoreWebView2BrowserExtensionEnableCompletedHandlerImpl
}

func (i *ICoreWebView2BrowserExtensionEnableCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtensionEnableCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2BrowserExtensionEnableCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownAddRef(this *ICoreWebView2BrowserExtensionEnableCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownRelease(this *ICoreWebView2BrowserExtensionEnableCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BrowserExtensionEnableCompletedHandlerInvoke(this *ICoreWebView2BrowserExtensionEnableCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.BrowserExtensionEnableCompleted(errorCode)
}

type _ICoreWebView2BrowserExtensionEnableCompletedHandlerImpl interface {
	IUnknownImpl
	BrowserExtensionEnableCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2BrowserExtensionEnableCompletedHandlerFn *_ICoreWebView2BrowserExtensionEnableCompletedHandlerVtbl

func NewICoreWebView2BrowserExtensionEnableCompletedHandler(impl _ICoreWebView2BrowserExtensionEnableCompletedHandlerImpl) *ICoreWebView2BrowserExtensionEnableCompletedHandler {
	if _ICoreWebView2BrowserExtensionEnableCompletedHandlerFn == nil {
		_ICoreWebView2BrowserExtensionEnableCompletedHandlerFn = &_ICoreWebView2BrowserExtensionEnableCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2BrowserExtensionEnableCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2BrowserExtensionEnableCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2BrowserExtensionEnableCompletedHandler{
		vtbl: _ICoreWebView2BrowserExtensionEnableCompletedHandlerFn,
		impl: impl,
	}
}