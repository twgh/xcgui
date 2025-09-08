package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler 接收获取浏览器扩展列表操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profilegetbrowserextensionscompletedhandler
type ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler struct {
	vtbl *_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerVtbl
	impl _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerImpl
}

func (i *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownAddRef(this *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownRelease(this *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerInvoke(this *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler, errorCode syscall.Errno, result *ICoreWebView2BrowserExtensionList) uintptr {
	return this.impl.ProfileGetBrowserExtensionsCompleted(errorCode, result)
}

type _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerImpl interface {
	IUnknownImpl
	ProfileGetBrowserExtensionsCompleted(errorCode syscall.Errno, result *ICoreWebView2BrowserExtensionList) uintptr
}

var _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerFn *_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerVtbl

func NewICoreWebView2ProfileGetBrowserExtensionsCompletedHandler(impl _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerImpl) *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	if _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerFn == nil {
		_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerFn = &_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler{
		vtbl: _ICoreWebView2ProfileGetBrowserExtensionsCompletedHandlerFn,
		impl: impl,
	}
}
