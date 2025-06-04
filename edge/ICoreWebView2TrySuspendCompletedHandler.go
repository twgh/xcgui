package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2TrySuspendCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc // COM 接口的 Invoke 方法
}

// ICoreWebView2TrySuspendCompletedHandler 接收 TrySuspend 方法的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2trysuspendcompletedhandler
type ICoreWebView2TrySuspendCompletedHandler struct {
	vtbl *_ICoreWebView2TrySuspendCompletedHandlerVtbl
	impl _ICoreWebView2TrySuspendCompletedHandlerImpl
}

func (i *ICoreWebView2TrySuspendCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2TrySuspendCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

// IUnknown 方法实现
func _ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface(
	this *ICoreWebView2TrySuspendCompletedHandler,
	refiid, object uintptr,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef(
	this *ICoreWebView2TrySuspendCompletedHandler,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease(
	this *ICoreWebView2TrySuspendCompletedHandler,
) uintptr {
	return this.impl.Release()
}

// Invoke 方法实现
func _ICoreWebView2TrySuspendCompletedHandlerInvoke(
	this *ICoreWebView2TrySuspendCompletedHandler,
	errorCode syscall.Errno, // HRESULT 错误码
	isSuccessful uintptr, // BOOL 类型（Windows 中为 4 字节）
) uintptr {
	// 将 Windows 的 BOOL（4 字节）转换为 Go 的 bool
	isSuccess := isSuccessful != 0
	return this.impl.TrySuspendCompleted(errorCode, isSuccess)
}

type _ICoreWebView2TrySuspendCompletedHandlerImpl interface {
	IUnknownImpl
	TrySuspendCompleted(errorCode syscall.Errno, isSuccessful bool) uintptr
}

var _ICoreWebView2TrySuspendCompletedHandlerFn = _ICoreWebView2TrySuspendCompletedHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2TrySuspendCompletedHandlerInvoke),
}

func NewICoreWebView2TrySuspendCompletedHandler(
	impl _ICoreWebView2TrySuspendCompletedHandlerImpl,
) *ICoreWebView2TrySuspendCompletedHandler {
	return &ICoreWebView2TrySuspendCompletedHandler{
		vtbl: &_ICoreWebView2TrySuspendCompletedHandlerFn,
		impl: impl,
	}
}
