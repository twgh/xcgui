//go:build 386

package com

// Call 按实参顺序调用 COM 方法(见 call.go 的说明).
//
// 386 是 stdcall, 实参全部走栈, 没有 XMM 的问题 —— 只要把每个实参摊成 32 位槽交给
// syscall.SyscallN 即可, 不需要 amd64 那套汇编转发器.
func Call(fn uintptr, args ...Arg) uintptr {
	if fn == 0 {
		return 0
	}
	slots := make([]uintptr, 0, len(args)+2)
	for _, arg := range args {
		switch arg.kind {
		case argInt, argFloat32:
			// uintptr / 指针 / float32 位模式: 一个 4 字节槽.
			slots = append(slots, uintptr(uint32(arg.bits)))
		case argInt64, argFloat64:
			// 64 位值(含按值传递的 8 字节结构体): 低 4 字节在前, 高 4 字节在后.
			slots = append(slots, uintptr(uint32(arg.bits)), uintptr(uint32(arg.bits>>32)))
		}
	}
	r, _, _ := ComProc(fn).Call(slots...)
	return r
}
