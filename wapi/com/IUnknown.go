package com

import (
	"syscall"
	"unsafe"
)

// IUnknownVtbl 是 IUnknown 的虚表, 所有 COM 接口的虚表都以它开头.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknownVtbl struct {
	QueryInterface ComProc
	AddRef         ComProc
	Release        ComProc
}

// IUnknown 是 COM 对象的映射(IUnknown 部分).
//
// 之所以能"嵌入它来复用方法": COM 对象的内存布局就是**首字段一个虚表指针**, 而本结构体只有
// 这一个字段, 因此把裸接口指针转成上层接口结构体(以值嵌入的方式接在它后面)之后:
//
//   - AddRef / Release / QueryInterface 只需要在这里实现一次, 上层接口全部继承;
//   - 上层接口自己的方法用各自的虚表类型重解释 `Vtbl` 指针即可(同一个对象、同一个虚表指针,
//     只是槽位更多) —— 注意是 `Vtbl`(虚表指针), **不是**对象地址, 后者会读到对象内存里的垃圾.
//
// 代价是不可给这些嵌入链上的结构体加字段(加了布局就不再等于一个虚表指针); 需要基类指针时直接
// 取嵌入字段地址.
type IUnknown struct {
	Vtbl *IUnknownVtbl
}

// AddRef 增加引用计数.
func (i *IUnknown) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

// Release 减少引用计数.
//
// 判空是刻意的: `defer obj.Release()` 与"零值对象"这两种用法很常见, 而映射结构体(nil)直接解引用
// 虚表就是崩进程 —— 这里让 nil 安全返回 0.
func (i *IUnknown) Release() uintptr {
	if i == nil {
		return 0
	}
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

// QueryInterface 查询该对象是否支持某个接口(refiid 传接口 IID, object 传出接口指针).
func (i *IUnknown) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// 下面三个是"独立函数版": 手上有裸对象指针(不便嵌入结构体, 例如要实现 COM 对象、或临时借用
// 一个指针)时用. nil 安全.

// QueryInterface 对裸对象指针查询接口.
func QueryInterface(obj unsafe.Pointer, refiid, object unsafe.Pointer) error {
	if obj == nil {
		return syscall.EINVAL
	}
	vtbl := *(*unsafe.Pointer)(obj)
	fn := *(*uintptr)(vtbl)
	r, _, _ := syscall.SyscallN(fn, uintptr(obj), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddRef 对裸对象指针增加引用计数.
func AddRef(obj unsafe.Pointer) uintptr {
	if obj == nil {
		return 0
	}
	vtbl := *(*unsafe.Pointer)(obj)
	fn := *(*uintptr)(unsafe.Add(vtbl, unsafe.Sizeof(uintptr(0))))
	r, _, _ := syscall.SyscallN(fn, uintptr(obj))
	return r
}

// Release 对裸对象指针减少引用计数.
func Release(obj unsafe.Pointer) uintptr {
	if obj == nil {
		return 0
	}
	vtbl := *(*unsafe.Pointer)(obj)
	fn := *(*uintptr)(unsafe.Add(vtbl, 2*unsafe.Sizeof(uintptr(0))))
	r, _, _ := syscall.SyscallN(fn, uintptr(obj))
	return r
}
