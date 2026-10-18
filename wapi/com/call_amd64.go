//go:build amd64

package com

import (
	"syscall"
	"unsafe"
)

// maxCallArgs amd64 下 Call 能表达的实参个数上限: 4 个寄存器位置 + 4 个栈位置.
const maxCallArgs = 8

// Call 按实参顺序调用 COM 方法(见 call.go 的说明).
func Call(fn uintptr, args ...Arg) uintptr {
	if fn == 0 {
		return 0
	}
	if len(args) > maxCallArgs {
		panic("com.Call: amd64 下最多支持 8 个实参(4 个寄存器位置 + 4 个栈位置)")
	}

	// float64 要 MOVSD 搬进 XMM, 走的是另一套转发器(且只有 4 个寄存器位置, 不支持压栈参数).
	for _, arg := range args {
		if arg.kind == argFloat64 {
			return callFloat64(fn, args)
		}
	}

	// 位置槽语义: 第 N 个实参落第 N 个位置; 位置 <4 时整数进 RCX/RDX/R8/R9、浮点进 XMM0~XMM3,
	// 位置 >=4 依次落到栈槽 [0..3].
	var a xmmArgsStack
	a.fn = fn
	for i, arg := range args {
		switch {
		case i >= 4:
			// 栈槽统一 8 字节; 浮点(仅 float32 会走到这里)放低 4 字节.
			a.stack[i-4] = uintptr(arg.bits)
		case arg.kind == argFloat32:
			switch i {
			case 0:
				a.x0 = uint32(arg.bits)
			case 1:
				a.x1 = uint32(arg.bits)
			case 2:
				a.x2 = uint32(arg.bits)
			case 3:
				a.x3 = uint32(arg.bits)
			}
		default: // argInt / argInt64
			switch i {
			case 0:
				a.this = uintptr(arg.bits)
			case 1:
				a.a1 = uintptr(arg.bits)
			case 2:
				a.a2 = uintptr(arg.bits)
			case 3:
				a.a3 = uintptr(arg.bits)
			}
		}
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline32StackPC, uintptr(unsafe.Pointer(&a)))
	return r
}

// callFloat64 是 Call 的 float64 分支: 用 MOVSD 版转发器, 只支持前 4 个位置(无压栈参数),
// 且同一次调用里不混用 float32(两套转发器一条指令只能搬一种宽度).
func callFloat64(fn uintptr, args []Arg) uintptr {
	if len(args) > 4 {
		panic("com.Call: float64 形参目前只支持前 4 个位置, 更靠后的参数需要压栈版转发器")
	}
	var a xmmArgs64
	a.fn = fn
	for i, arg := range args {
		switch arg.kind {
		case argFloat32:
			panic("com.Call: 同一次调用里不要混用 float32 与 float64 形参")
		case argFloat64:
			switch i {
			case 0:
				a.x0 = arg.bits
			case 1:
				a.x1 = arg.bits
			case 2:
				a.x2 = arg.bits
			case 3:
				a.x3 = arg.bits
			}
		default: // argInt / argInt64
			switch i {
			case 0:
				a.this = uintptr(arg.bits)
			case 1:
				a.a1 = uintptr(arg.bits)
			case 2:
				a.a2 = uintptr(arg.bits)
			case 3:
				a.a3 = uintptr(arg.bits)
			}
		}
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline64PC, uintptr(unsafe.Pointer(&a)))
	return r
}
