package com

import (
	"testing"
	"unsafe"
)

func TestInterfacePtr(t *testing.T) {
	if InterfacePtr(0) != nil {
		t.Error("InterfacePtr(0) 应为 nil")
	}
	p := InterfacePtr(0x1234)
	if uintptr(p) != 0x1234 {
		t.Errorf("InterfacePtr(0x1234) 搬出来的指针 = %#x", uintptr(p))
	}
}

// IUnknown 的结构体布局就是 COM 对象布局: 只有一个虚表指针字段, 虚表是三个连续函数指针.
func TestLayout(t *testing.T) {
	ptrSize := unsafe.Sizeof(uintptr(0))
	if got := unsafe.Sizeof(IUnknown{}); got != ptrSize {
		t.Errorf("IUnknown 大小 = %d 字节, 期望 %d(一个虚表指针)", got, ptrSize)
	}
	if got := unsafe.Sizeof(IUnknownVtbl{}); got != 3*ptrSize {
		t.Errorf("IUnknownVtbl 大小 = %d 字节, 期望 %d(三个函数指针)", got, 3*ptrSize)
	}
	if got := unsafe.Offsetof(IUnknownVtbl{}.Release); got != 2*ptrSize {
		t.Errorf("Release 槽位偏移 = %d, 期望 %d", got, 2*ptrSize)
	}
}

func TestNilSafety(t *testing.T) {
	var obj *IUnknown
	if obj.Release() != 0 {
		t.Error("nil 对象 Release 应返回 0")
	}
	if Release(nil) != 0 || AddRef(nil) != 0 {
		t.Error("独立函数版对 nil 应返回 0")
	}
	if err := QueryInterface(nil, nil, nil); err == nil {
		t.Error("独立函数版对 nil 应报错")
	}
}
