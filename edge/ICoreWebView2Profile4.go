package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile4 是 ICoreWebView2Profile3 的延续, 提供了权限管理功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile4
type ICoreWebView2Profile4 struct {
	ICoreWebView2Profile3
}

// SetPermissionState 异步为给定的权限类型和来源设置权限状态。
//
// permissionKind: 权限类型。
//
// origin: 权限的来源。
//
// state: 要设置的权限状态。
func (i *ICoreWebView2Profile4) SetPermissionState(permissionKind COREWEBVIEW2_PERMISSION_KIND, origin string, state COREWEBVIEW2_PERMISSION_STATE, handler *ICoreWebView2SetPermissionStateCompletedHandler) error {
	_origin, err := syscall.UTF16PtrFromString(origin)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.SetPermissionState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(permissionKind),
		uintptr(unsafe.Pointer(_origin)),
		uintptr(state),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetNonDefaultPermissionSettings 获取所有非默认的权限设置。
//   - 使用此方法可获取当前会话和先前会话中设置的权限状态。
func (i *ICoreWebView2Profile4) GetNonDefaultPermissionSettings(handler *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) error {
	r, _, _ := i.Vtbl.GetNonDefaultPermissionSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetPermissionStateEx 在指定的 WebView 事件实现上下文中异步为给定的权限类型和来源设置权限状态。
//
// impl: *WebViewEventImpl。
//
// permissionKind: 权限类型。
//
// origin: 权限的来源。
//
// state: 要设置的权限状态。
//
// cb: 设置完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2Profile4) SetPermissionStateEx(impl *WebViewEventImpl, permissionKind COREWEBVIEW2_PERMISSION_KIND, origin string, state COREWEBVIEW2_PERMISSION_STATE, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "SetPermissionStateCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "SetPermissionStateCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "SetPermissionStateCompleted")
	}
	return i.SetPermissionState(permissionKind, origin, state, (*ICoreWebView2SetPermissionStateCompletedHandler)(handler))
}

// GetNonDefaultPermissionSettingsEx 在指定的 WebView 事件实现上下文中获取所有非默认的权限设置。
//
// impl: *WebViewEventImpl。
//
// cb: 获取完成后的回调处理程序，可以为 nil。如果获取成功，将通过 cb 返回权限设置集合。
func (i *ICoreWebView2Profile4) GetNonDefaultPermissionSettingsEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, result *ICoreWebView2PermissionSettingCollectionView) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "GetNonDefaultPermissionSettingsCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "GetNonDefaultPermissionSettingsCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "GetNonDefaultPermissionSettingsCompleted")
	}
	return i.GetNonDefaultPermissionSettings((*ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler)(handler))
}
