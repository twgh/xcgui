//go:build 386

package com

import (
	"runtime"
	"unsafe"
)

// Call 按实参顺序调用 COM 方法(见 call.go 的说明).
//
// 386 是 stdcall, 实参全部走栈, 没有 XMM 的问题 —— 只要把每个实参摊成 32 位槽交给
// syscall.SyscallN 即可, 不需要 amd64 那套汇编转发器.
func Call(fn uintptr, args ...Arg) uintptr {
	if fn == 0 {
		return 0
	}
	slots := make([]uintptr, 0, len(args)+4)
	for _, arg := range args {
		switch arg.kind {
		case argInt, argFloat32:
			// uintptr / 指针 / float32 位模式: 一个 4 字节槽.
			slots = append(slots, uintptr(uint32(arg.bits)))
		case argInt64, argFloat64:
			// 64 位值: 低 4 字节在前, 高 4 字节在后.
			slots = append(slots, uintptr(uint32(arg.bits)), uintptr(uint32(arg.bits>>32)))
		case argStruct:
			// 按值结构体: 按 4 字节槽逐个压栈, 末尾不足 4 字节也占一个槽(高位补 0).
			for off := uintptr(0); off < arg.size; off += 4 {
				var w uint32
				for b := uintptr(0); b < 4 && off+b < arg.size; b++ {
					w |= uint32(*(*byte)(unsafe.Add(arg.ptr, off+b))) << (8 * b)
				}
				slots = append(slots, uintptr(w))
			}
		}
	}
	r, _, _ := ComProc(fn).Call(slots...)
	runtime.KeepAlive(args) // 结构体实参靠 args 里的 unsafe.Pointer 续命
	return r
}
