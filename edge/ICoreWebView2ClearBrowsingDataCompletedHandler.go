package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ClearBrowsingDataCompletedHandler 接收清除浏览数据操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clearbrowsingdatacompletedhandler
type ICoreWebView2ClearBrowsingDataCompletedHandler struct {
	vtbl *_ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl
	impl _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl
}

func (i *ICoreWebView2ClearBrowsingDataCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClearBrowsingDataCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ClearBrowsingDataCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownAddRef(this *ICoreWebView2ClearBrowsingDataCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownRelease(this *ICoreWebView2ClearBrowsingDataCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerInvoke(this *ICoreWebView2ClearBrowsingDataCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.ClearBrowsingDataCompleted(errorCode)
}

type _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl interface {
	IUnknownImpl
	ClearBrowsingDataCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2ClearBrowsingDataCompletedHandlerFn *_ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl

func NewICoreWebView2ClearBrowsingDataCompletedHandler(impl _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl) *ICoreWebView2ClearBrowsingDataCompletedHandler {
	if _ICoreWebView2ClearBrowsingDataCompletedHandlerFn == nil {
		_ICoreWebView2ClearBrowsingDataCompletedHandlerFn = &_ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2ClearBrowsingDataCompletedHandler{
		vtbl: _ICoreWebView2ClearBrowsingDataCompletedHandlerFn,
		impl: impl,
	}
}
