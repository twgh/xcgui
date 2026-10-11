package com

import "syscall"

// ComProc 存储 COM 过程(虚表里的一个函数指针).
type ComProc uintptr

// NewComProc 从 Go 函数创建一个新的 COM 过程(自己实现 COM 对象、需要把 Go 回调塞进虚表时用).
func NewComProc(fn interface{}) ComProc {
	return ComProc(syscall.NewCallback(fn))
}

// Call 调用 COM 过程.
//
// 参数怎么摆见包注释"参数传递"一节: 386 下结构体/浮点都要按 32 位槽拆开传, amd64 下结构体传指针;
// 而**浮点形参(以及按值返回的 float)在本方法里传不进去** —— SyscallN 只能设置通用寄存器.
//
//go:uintptrescapes
func (p ComProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	// uintptrescapes 注解必不可少: 它让编译器把 uintptr 实参当作可能指向内存来处理,
	// 从而满足 unsafe.Pointer 规则"(4) 调用 syscall.Syscall 时把指针转成 uintptr" ——
	// 否则调用期间指针可能被搬走(尤其是位于会动态增长的 Go 栈上的指针).
	// 见 https://pkg.go.dev/unsafe#Pointer 与 https://github.com/golang/go/issues/34474
	return syscall.SyscallN(uintptr(p), a...)
}
