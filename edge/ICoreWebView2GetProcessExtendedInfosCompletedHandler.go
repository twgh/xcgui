package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2GetProcessExtendedInfosCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2GetProcessExtendedInfosCompletedHandler 接收获取进程扩展信息操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2getprocessextendedinfoscompletedhandler
type ICoreWebView2GetProcessExtendedInfosCompletedHandler struct {
	vtbl *_ICoreWebView2GetProcessExtendedInfosCompletedHandlerVtbl
	impl _ICoreWebView2GetProcessExtendedInfosCompletedHandlerImpl
}

func (i *ICoreWebView2GetProcessExtendedInfosCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2GetProcessExtendedInfosCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetProcessExtendedInfosCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetProcessExtendedInfosCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownRelease(this *ICoreWebView2GetProcessExtendedInfosCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2GetProcessExtendedInfosCompletedHandlerInvoke(this *ICoreWebView2GetProcessExtendedInfosCompletedHandler, errorCode syscall.Errno, result *ICoreWebView2ProcessExtendedInfoCollection) uintptr {
	return this.impl.GetProcessExtendedInfosCompleted(errorCode, result)
}

type _ICoreWebView2GetProcessExtendedInfosCompletedHandlerImpl interface {
	IUnknownImpl
	GetProcessExtendedInfosCompleted(errorCode syscall.Errno, result *ICoreWebView2ProcessExtendedInfoCollection) uintptr
}

var _ICoreWebView2GetProcessExtendedInfosCompletedHandlerFn *_ICoreWebView2GetProcessExtendedInfosCompletedHandlerVtbl

func NewICoreWebView2GetProcessExtendedInfosCompletedHandler(impl _ICoreWebView2GetProcessExtendedInfosCompletedHandlerImpl) *ICoreWebView2GetProcessExtendedInfosCompletedHandler {
	if _ICoreWebView2GetProcessExtendedInfosCompletedHandlerFn == nil {
		_ICoreWebView2GetProcessExtendedInfosCompletedHandlerFn = &_ICoreWebView2GetProcessExtendedInfosCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2GetProcessExtendedInfosCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2GetProcessExtendedInfosCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2GetProcessExtendedInfosCompletedHandler{
		vtbl: _ICoreWebView2GetProcessExtendedInfosCompletedHandlerFn,
		impl: impl,
	}
}
