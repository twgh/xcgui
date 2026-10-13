package edge

import (
	"syscall"
	"unsafe"
)

type IUnknownImpl interface {
	QueryInterface(refiid, object unsafe.Pointer) uintptr
	AddRef() uintptr
	Release() uintptr
}

type IUnknownVtbl struct {
	QueryInterface ComProc
	AddRef         ComProc
	Release        ComProc
}

// IUnknown 封装了 COM 基础的 IUnknown 接口.
type IUnknown struct {
	Vtbl *IUnknownVtbl
}

func (i *IUnknown) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IUnknown) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IUnknown) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

type IUnknown_Impl struct {
}

func (w *IUnknown_Impl) QueryInterface(_, _ unsafe.Pointer) uintptr {
	return 0
}

func (w *IUnknown_Impl) AddRef() uintptr {
	return 1
}

func (w *IUnknown_Impl) Release() uintptr {
	return 1
}
