package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ComProc 存储 COM 过程。
type ComProc uintptr

// NewComProc 从 Go 函数创建一个新的 COM 过程。
func NewComProc(fn interface{}) ComProc {
	return ComProc(syscall.NewCallback(fn))
}

const ptrSize = unsafe.Sizeof(uintptr(0))

// ComQueryInterface 通用的 COM 对象查询接口.
func ComQueryInterface(obj unsafe.Pointer, refiid unsafe.Pointer, object unsafe.Pointer) error {
	// 获取 vtable 指针
	vtable := *(*uintptr)(obj)
	// 获取 QueryInterface 函数地址, 在 vtable 中的索引是0
	queryInterfaceFunc := *(*uintptr)(unsafe.Pointer(vtable))
	// 调用 QueryInterface 函数 (stdcall约定)
	r, _, err := syscall.SyscallN(queryInterfaceFunc, uintptr(obj), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ComAddRef 通用的 COM 对象添加引用计数.
func ComAddRef(obj unsafe.Pointer) uintptr {
	// 获取 vtable 指针
	vtable := *(*uintptr)(obj)
	// 获取 AddRef 函数地址, 在 vtable 中的索引是1
	releaseFunc := *(*uintptr)(unsafe.Pointer(vtable + ptrSize))
	// 调用 AddRef 函数 (stdcall约定)
	r, _, _ := syscall.SyscallN(releaseFunc, uintptr(obj))
	return r
}

// ComRelease 通用的 COM 对象减少引用计数.
func ComRelease(obj unsafe.Pointer) uintptr {
	// 获取 vtable 指针
	vtable := *(*uintptr)(obj)
	// 获取 Release 函数地址, 在 vtable 中的索引是2
	releaseFunc := *(*uintptr)(unsafe.Pointer(vtable + 2*ptrSize))
	// 调用 Release 函数 (stdcall约定)
	r, _, _ := syscall.SyscallN(releaseFunc, uintptr(obj))
	return r
}
