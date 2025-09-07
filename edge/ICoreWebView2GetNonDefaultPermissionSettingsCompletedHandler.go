package edge

import (
	"syscall"
	"unsafe"
)

type _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler 接收获取非默认权限设置操作的结果。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler
type ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler struct {
	vtbl *_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerVtbl
	impl _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImpl
}

func (i *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownRelease(this *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerInvoke(this *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler, errorCode syscall.Errno, result *ICoreWebView2PermissionSettingCollectionView) uintptr {
	return this.impl.GetNonDefaultPermissionSettingsCompleted(errorCode, result)
}

type _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImpl interface {
	IUnknownImpl
	GetNonDefaultPermissionSettingsCompleted(errorCode syscall.Errno, result *ICoreWebView2PermissionSettingCollectionView) uintptr
}

var _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerFn *_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerVtbl

func NewICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(impl _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImpl) *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	if _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerFn == nil {
		_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerFn = &_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerInvoke),
		}
	}
	return &ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler{
		vtbl: _ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandlerFn,
		impl: impl,
	}
}
