package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2PermissionSetting 提供对权限设置的访问，包括权限类型、来源和状态。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting
type ICoreWebView2PermissionSetting struct {
	Vtbl *ICoreWebView2PermissionSettingVtbl
}

type ICoreWebView2PermissionSettingVtbl struct {
	IUnknownVtbl
	GetPermissionKind   ComProc
	GetPermissionOrigin ComProc
	GetPermissionState  ComProc
}

func (i *ICoreWebView2PermissionSetting) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionSetting) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionSetting) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPermissionKind 获取权限的类型。
func (i *ICoreWebView2PermissionSetting) GetPermissionKind() (COREWEBVIEW2_PERMISSION_KIND, error) {
	var value COREWEBVIEW2_PERMISSION_KIND
	r, _, _ := i.Vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetPermissionOrigin 获取权限的源URL。
func (i *ICoreWebView2PermissionSetting) GetPermissionOrigin() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetPermissionOrigin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	origin := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return origin, nil
}

// GetPermissionState 获取权限的状态。
func (i *ICoreWebView2PermissionSetting) GetPermissionState() (COREWEBVIEW2_PERMISSION_STATE, error) {
	var value COREWEBVIEW2_PERMISSION_STATE
	r, _, _ := i.Vtbl.GetPermissionState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetPermissionKind 获取权限的类型。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionSetting) MustGetPermissionKind() COREWEBVIEW2_PERMISSION_KIND {
	value, err := i.GetPermissionKind()
	ReportErrorAtuo(err)
	return value
}

// MustGetPermissionOrigin 获取权限的源URL。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionSetting) MustGetPermissionOrigin() string {
	value, err := i.GetPermissionOrigin()
	ReportErrorAtuo(err)
	return value
}

// MustGetPermissionState 获取权限的状态。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionSetting) MustGetPermissionState() COREWEBVIEW2_PERMISSION_STATE {
	value, err := i.GetPermissionState()
	ReportErrorAtuo(err)
	return value
}
