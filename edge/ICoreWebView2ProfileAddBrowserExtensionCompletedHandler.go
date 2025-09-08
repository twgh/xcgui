package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ProfileAddBrowserExtensionCompletedHandler 接收添加浏览器扩展操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profileaddbrowserextensioncompletedhandler
type ICoreWebView2ProfileAddBrowserExtensionCompletedHandler struct {
	vtbl *_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerVtbl
	impl _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerImpl
}

func (i *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownAddRef(this *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownRelease(this *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerInvoke(this *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler, errorCode syscall.Errno, result *ICoreWebView2BrowserExtension) uintptr {
	return this.impl.ProfileAddBrowserExtensionCompleted(errorCode, result)
}

type _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerImpl interface {
	IUnknownImpl
	ProfileAddBrowserExtensionCompleted(errorCode syscall.Errno, result *ICoreWebView2BrowserExtension) uintptr
}

var _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerFn *_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerVtbl

func NewICoreWebView2ProfileAddBrowserExtensionCompletedHandler(impl _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerImpl) *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	if _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerFn == nil {
		_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerFn = &_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ProfileAddBrowserExtensionCompletedHandler{
		vtbl: _ICoreWebView2ProfileAddBrowserExtensionCompletedHandlerFn,
		impl: impl,
	}
}
