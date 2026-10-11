#include "textflag.h"

// xmmTrampoline32 / xmmTrampoline64 / xmmTrampoline32RetFloat 按 **Windows x64 约定**被调用:
// RCX = 参数块指针, 返回值在 RAX(RetFloat 版是"目标方法按值返回 float", 它把 XMM0 搬到 RAX).
//
// 它们由 syscall.SyscallN 调用(所以跑在**系统栈**上, 见 xmm_amd64.go 的说明), 职责是:
// 把参数块里的整数/指针摆进 RCX/RDX/R8/R9、浮点摆进 XMM0~XMM3, 再间接调用目标函数.
//
// 调用目标函数前必须自己摆好栈, 这三条都是实测定出来的:
//  1. **必须经 syscall.SyscallN 进来**(它内部 cgocall 到系统栈) —— 直接在 Go 的 goroutine 栈上
//     CALL 目标, 简单 API(GetTickCount)能过, 稍复杂的(user32.GetSystemMetrics)会因 Windows 的
//     栈探测基于 TEB 栈范围而落到非法地址崩溃.
//  2. **调用点 RSP 要 16 字节对齐** —— 否则被调方入口就不是约定的 ≡8, 用到 SSE 对齐存取的 API 崩;
//     实测这条路径进来时 RSP ≡ 0 (mod 16), 所以减 32(16 的倍数)刚好, 减 40 就会崩.
//  3. **要留 32 字节 shadow space**(被调方的"槽位", 它会往里写寄存器参数) —— SUBQ $32 正好留出.
//
// 参数块布局与 Go 侧一致:
//   xmmArgs   : fn=0, this=8, a1=16, a2=24, a3=32, x0=40, x1=44, x2=48, x3=52  (float32)
//   xmmArgs64 : 同上, 但 x0=40, x1=48, x2=56, x3=64                              (float64)
//
// 指针先搬到 R10(调用者保存寄存器, 且不跨越任何会破坏它的调用)再逐个读取.

TEXT ·xmmTrampoline32(SB), NOSPLIT, $0-0
	MOVQ  CX, R10
	MOVQ  0(R10), AX
	MOVQ  8(R10), CX
	MOVQ  16(R10), DX
	MOVQ  24(R10), R8
	MOVQ  32(R10), R9
	MOVSS 40(R10), X0
	MOVSS 44(R10), X1
	MOVSS 48(R10), X2
	MOVSS 52(R10), X3
	SUBQ  $32, SP
	CALL  AX
	ADDQ  $32, SP
	RET

TEXT ·xmmTrampoline64(SB), NOSPLIT, $0-0
	MOVQ  CX, R10
	MOVQ  0(R10), AX
	MOVQ  8(R10), CX
	MOVQ  16(R10), DX
	MOVQ  24(R10), R8
	MOVQ  32(R10), R9
	MOVSD 40(R10), X0
	MOVSD 48(R10), X1
	MOVSD 56(R10), X2
	MOVSD 64(R10), X3
	SUBQ  $32, SP
	CALL  AX
	ADDQ  $32, SP
	RET

TEXT ·xmmTrampoline32RetFloat(SB), NOSPLIT, $0-0
	MOVQ  CX, R10
	MOVQ  0(R10), AX
	MOVQ  8(R10), CX
	MOVQ  16(R10), DX
	MOVQ  24(R10), R8
	MOVQ  32(R10), R9
	MOVSS 40(R10), X0
	MOVSS 44(R10), X1
	MOVSS 48(R10), X2
	MOVSS 52(R10), X3
	SUBQ  $32, SP
	CALL  AX
	ADDQ  $32, SP
	MOVQ  X0, AX
	RET

// func xmmTrampoline32S1()  —— 同 xmmTrampoline32, 另把参数块里的 stack0 放到**第 5 个参数位**.
//
// 第 5 个参数(x64 约定)在 [RSP+32]: 前 32 字节是 shadow space, 之后才是栈参数. 这里 SUBQ $48
// (16 的倍数, 保持调用点对齐; 同时给出 32 shadow + 8 栈参数位 + 8 余量).
TEXT ·xmmTrampoline32S1(SB), NOSPLIT, $0-0
	MOVQ  CX, R10
	MOVQ  0(R10), AX
	MOVQ  8(R10), CX
	MOVQ  16(R10), DX
	MOVQ  24(R10), R8
	MOVQ  32(R10), R9
	MOVSS 40(R10), X0
	MOVSS 44(R10), X1
	MOVSS 48(R10), X2
	MOVSS 52(R10), X3
	MOVQ  56(R10), R11
	SUBQ  $48, SP
	MOVQ  R11, 32(SP)
	CALL  AX
	ADDQ  $48, SP
	RET
