package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment14 提供创建文件系统句柄和对象集合的功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment14
type ICoreWebView2Environment14 struct {
	ICoreWebView2Environment13
}

// CreateWebFileSystemFileHandle 根据表示 Web FileSystemFileHandle 的路径创建一个 ICoreWebView2FileSystemHandle 对象。
//
// path: 文件路径.
//
// permission: 文件访问权限.
func (i *ICoreWebView2Environment14) CreateWebFileSystemFileHandle(path string, permission COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION) (*ICoreWebView2FileSystemHandle, error) {
	_path, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	var value *ICoreWebView2FileSystemHandle
	r, _, _ := i.Vtbl.CreateWebFileSystemFileHandle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
		uintptr(permission),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// CreateWebFileSystemDirectoryHandle 根据表示 Web FileSystemDirectoryHandle 的路径创建 ICoreWebView2FileSystemHandle 对象。
//
// path: 目录路径.
//
// permission: 目录访问权限.
func (i *ICoreWebView2Environment14) CreateWebFileSystemDirectoryHandle(path string, permission COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION) (*ICoreWebView2FileSystemHandle, error) {
	_path, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	var value *ICoreWebView2FileSystemHandle
	r, _, _ := i.Vtbl.CreateWebFileSystemDirectoryHandle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
		uintptr(permission),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// CreateObjectCollection 创建一个通用对象集合。
//
// items: 要添加到集合中的对象数组.
func (i *ICoreWebView2Environment14) CreateObjectCollection(items []*IUnknown) (*ICoreWebView2ObjectCollection, error) {
	var value *ICoreWebView2ObjectCollection
	r, _, _ := i.Vtbl.CreateObjectCollection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(len(items)),
		uintptr(unsafe.Pointer(&items[0])),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustCreateWebFileSystemFileHandle 根据表示 Web FileSystemFileHandle 的路径创建一个 ICoreWebView2FileSystemHandle 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment14) MustCreateWebFileSystemFileHandle(path string, permission COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION) *ICoreWebView2FileSystemHandle {
	value, err := i.CreateWebFileSystemFileHandle(path, permission)
	ReportErrorAuto(err)
	return value
}

// MustCreateWebFileSystemDirectoryHandle 根据表示 Web FileSystemDirectoryHandle 的路径创建 ICoreWebView2FileSystemHandle 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment14) MustCreateWebFileSystemDirectoryHandle(path string, permission COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION) *ICoreWebView2FileSystemHandle {
	value, err := i.CreateWebFileSystemDirectoryHandle(path, permission)
	ReportErrorAuto(err)
	return value
}

// MustCreateObjectCollection 创建一个通用对象集合。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment14) MustCreateObjectCollection(items []*IUnknown) *ICoreWebView2ObjectCollection {
	value, err := i.CreateObjectCollection(items)
	ReportErrorAuto(err)
	return value
}
