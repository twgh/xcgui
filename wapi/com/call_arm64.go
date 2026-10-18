//go:build arm64

package com

import "runtime"

// Call 按实参顺序调用 COM 方法(见 call.go 的说明).
//
// arm64 目前只支持**不含浮点形参**的调用: 整数/指针走 X0~X7, syscall.SyscallN 能设置;
// 浮点走 V0~V7, 与 amd64 的 XMM 同理需要汇编转发器(尚未实现).
func Call(fn uintptr, args ...Arg) uintptr {
	if fn == 0 {
		return 0
	}
	slots := make([]uintptr, len(args))
	for i, arg := range args {
		kind, bits := normArg(arg)
		switch kind {
		case argInt, argInt64:
			slots[i] = uintptr(bits)
		default:
			panic("com.Call: arm64 尚未实现浮点形参(需要汇编转发器把值放进 V 寄存器)")
		}
	}
	r, _, _ := ComProc(fn).Call(slots...)
	runtime.KeepAlive(args) // 结构体实参靠 args 里的 unsafe.Pointer 续命
	return r
}
