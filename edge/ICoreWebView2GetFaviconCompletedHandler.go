package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2GetFaviconCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2GetFaviconCompletedHandler 接收获取网站图标(Favicon)操作完成的回调。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2getfaviconcompletedhandler
type ICoreWebView2GetFaviconCompletedHandler struct {
	vtbl *_ICoreWebView2GetFaviconCompletedHandlerVtbl
	impl _ICoreWebView2GetFaviconCompletedHandlerImpl
}

func (i *ICoreWebView2GetFaviconCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2GetFaviconCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetFaviconCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2GetFaviconCompletedHandlerInvoke(this *ICoreWebView2GetFaviconCompletedHandler, errorCode syscall.Errno, faviconStream *IStream) uintptr {
	return this.impl.GetFaviconCompleted(errorCode, faviconStream)
}

type _ICoreWebView2GetFaviconCompletedHandlerImpl interface {
	IUnknownImpl
	GetFaviconCompleted(errorCode syscall.Errno, faviconStream *IStream) uintptr
}

var _ICoreWebView2GetFaviconCompletedHandlerFn *_ICoreWebView2GetFaviconCompletedHandlerVtbl

func NewICoreWebView2GetFaviconCompletedHandler(impl _ICoreWebView2GetFaviconCompletedHandlerImpl) *ICoreWebView2GetFaviconCompletedHandler {
	if _ICoreWebView2GetFaviconCompletedHandlerFn == nil {
		_ICoreWebView2GetFaviconCompletedHandlerFn = &_ICoreWebView2GetFaviconCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2GetFaviconCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2GetFaviconCompletedHandler{
		vtbl: _ICoreWebView2GetFaviconCompletedHandlerFn,
		impl: impl,
	}
}
