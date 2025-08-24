package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2FileSystemHandle 提供对文件系统句柄的访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle
type ICoreWebView2FileSystemHandle struct {
	Vtbl *ICoreWebView2FileSystemHandleVtbl
}

type ICoreWebView2FileSystemHandleVtbl struct {
	IUnknownVtbl
	GetKind       ComProc
	GetPath       ComProc
	GetPermission ComProc
}

func (i *ICoreWebView2FileSystemHandle) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FileSystemHandle) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FileSystemHandle) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKind 获取文件系统句柄的类型。
func (i *ICoreWebView2FileSystemHandle) GetKind() (COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND, error) {
	var value COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetPath 获取文件系统句柄的路径。
func (i *ICoreWebView2FileSystemHandle) GetPath() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return path, nil
}

// GetPermission 获取文件系统句柄的权限。
func (i *ICoreWebView2FileSystemHandle) GetPermission() (COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION, error) {
	var value COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION
	r, _, _ := i.Vtbl.GetPermission.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetKind 获取文件系统句柄的类型。出错时会触发全局错误回调。
func (i *ICoreWebView2FileSystemHandle) MustGetKind() COREWEBVIEW2_FILE_SYSTEM_HANDLE_KIND {
	value, err := i.GetKind()
	ReportErrorAtuo(err)
	return value
}

// MustGetPath 获取文件系统句柄的路径。出错时会触发全局错误回调。
func (i *ICoreWebView2FileSystemHandle) MustGetPath() string {
	value, err := i.GetPath()
	ReportErrorAtuo(err)
	return value
}

// MustGetPermission 获取文件系统句柄的权限。出错时会触发全局错误回调。
func (i *ICoreWebView2FileSystemHandle) MustGetPermission() COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION {
	value, err := i.GetPermission()
	ReportErrorAtuo(err)
	return value
}
