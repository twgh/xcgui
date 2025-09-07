package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2SetPermissionStateCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2SetPermissionStateCompletedHandler 接收设置权限状态操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler
type ICoreWebView2SetPermissionStateCompletedHandler struct {
	vtbl *_ICoreWebView2SetPermissionStateCompletedHandlerVtbl
	impl _ICoreWebView2SetPermissionStateCompletedHandlerImpl
}

func (i *ICoreWebView2SetPermissionStateCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SetPermissionStateCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2SetPermissionStateCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownAddRef(this *ICoreWebView2SetPermissionStateCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownRelease(this *ICoreWebView2SetPermissionStateCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2SetPermissionStateCompletedHandlerInvoke(this *ICoreWebView2SetPermissionStateCompletedHandler, errorCode syscall.Errno) uintptr {
	return this.impl.SetPermissionStateCompleted(errorCode)
}

type _ICoreWebView2SetPermissionStateCompletedHandlerImpl interface {
	IUnknownImpl
	SetPermissionStateCompleted(errorCode syscall.Errno) uintptr
}

var _ICoreWebView2SetPermissionStateCompletedHandlerFn *_ICoreWebView2SetPermissionStateCompletedHandlerVtbl

func NewICoreWebView2SetPermissionStateCompletedHandler(impl _ICoreWebView2SetPermissionStateCompletedHandlerImpl) *ICoreWebView2SetPermissionStateCompletedHandler {
	if _ICoreWebView2SetPermissionStateCompletedHandlerFn == nil {
		_ICoreWebView2SetPermissionStateCompletedHandlerFn = &_ICoreWebView2SetPermissionStateCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2SetPermissionStateCompletedHandler{
		vtbl: _ICoreWebView2SetPermissionStateCompletedHandlerFn,
		impl: impl,
	}
}
