package edge

import "unsafe"

// ICoreWebView2FramePermissionRequestedEventHandler 接收框架权限请求事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framepermissionrequestedeventhandler
type ICoreWebView2FramePermissionRequestedEventHandler struct {
	vtbl *iCoreWebView2FramePermissionRequestedEventHandlerVtbl
	impl iCoreWebView2FramePermissionRequestedEventHandlerImpl
}

type iCoreWebView2FramePermissionRequestedEventHandlerImpl interface {
	IUnknownImpl
	FramePermissionRequested(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr
}

type iCoreWebView2FramePermissionRequestedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

func (i *ICoreWebView2FramePermissionRequestedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FramePermissionRequestedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FramePermissionRequestedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2FramePermissionRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownRelease(this *ICoreWebView2FramePermissionRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FramePermissionRequestedEventHandlerInvoke(this *ICoreWebView2FramePermissionRequestedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr {
	return this.impl.FramePermissionRequested(sender, args)
}

var iCoreWebView2FramePermissionRequestedEventHandlerFn *iCoreWebView2FramePermissionRequestedEventHandlerVtbl

func NewICoreWebView2FramePermissionRequestedEventHandler(impl iCoreWebView2FramePermissionRequestedEventHandlerImpl) *ICoreWebView2FramePermissionRequestedEventHandler {
	if iCoreWebView2FramePermissionRequestedEventHandlerFn == nil {
		iCoreWebView2FramePermissionRequestedEventHandlerFn = &iCoreWebView2FramePermissionRequestedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2FramePermissionRequestedEventHandler{
		vtbl: iCoreWebView2FramePermissionRequestedEventHandlerFn,
		impl: impl,
	}
}
