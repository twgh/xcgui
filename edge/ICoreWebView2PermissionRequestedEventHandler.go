package edge

import "unsafe"

type iCoreWebView2PermissionRequestedEventHandlerImpl interface {
	IUnknownImpl
	PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr
}

type iCoreWebView2PermissionRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2PermissionRequestedEventHandler 接收请求权限事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventhandler
type ICoreWebView2PermissionRequestedEventHandler struct {
	vtbl *iCoreWebView2PermissionRequestedEventHandlerVtbl
	impl iCoreWebView2PermissionRequestedEventHandlerImpl
}

func (i *ICoreWebView2PermissionRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2PermissionRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease(this *ICoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PermissionRequestedEventHandlerInvoke(this *ICoreWebView2PermissionRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	return this.impl.PermissionRequested(sender, args)
}

var iCoreWebView2PermissionRequestedEventHandlerFn = iCoreWebView2PermissionRequestedEventHandlerVtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PermissionRequestedEventHandlerInvoke),
}

func NewICoreWebView2PermissionRequestedEventHandler(impl iCoreWebView2PermissionRequestedEventHandlerImpl) *ICoreWebView2PermissionRequestedEventHandler {
	return &ICoreWebView2PermissionRequestedEventHandler{
		vtbl: &iCoreWebView2PermissionRequestedEventHandlerFn,
		impl: impl,
	}
}
