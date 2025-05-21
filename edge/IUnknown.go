package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"sync/atomic"
	"syscall"
	"unsafe"
)

type IUnknownVtbl struct {
	QueryInterface ComProc
	AddRef         ComProc
	Release        ComProc
}

type IUnknownImpl interface {
	QueryInterface(refiid, object uintptr) uintptr
	AddRef() uintptr
	Release() uintptr
}

// IUnknown 封装了COM基础的IUnknown接口
type IUnknown struct {
	Vtbl *IUnknownVtbl
	impl IUnknownImpl
	ref  int32
}

// NewIUnknown 创建一个新的 IUnknown 实例
func NewIUnknown(impl IUnknownImpl) *IUnknown {
	return &IUnknown{
		Vtbl: &IUnknownVtbl{
			QueryInterface: NewComProc(queryInterface),
			AddRef:         NewComProc(addRef),
			Release:        NewComProc(release),
		},
		impl: impl,
		ref:  1,
	}
}

func queryInterface(this *IUnknown, refiid uintptr, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func addRef(this *IUnknown) uintptr {
	atomic.AddInt32(&this.ref, 1)
	return this.impl.AddRef()
}

func release(this *IUnknown) uintptr {
	ref := atomic.AddInt32(&this.ref, -1)
	if ref == 0 {
		this.impl.Release()
	}
	return uintptr(ref)
}

func (i *IUnknown) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IUnknown) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IUnknown) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
