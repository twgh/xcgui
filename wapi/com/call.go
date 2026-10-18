package com

import "math"

// 跨架构的 COM 调用: 按"被调方签名里实参的出现顺序"逐个描述参数, 由各架构的实现负责把它们摆到
// 该去的地方(amd64: 寄存器 + shadow space + 栈; 386: 全部压栈).
//
// 与 CallWithFloat32 / CallWithFloat32Stack(仅 amd64, 需要调用方自己算寄存器槽位)不同, 这里不用
// 关心槽位: this 也算第一个实参, 照着 Windows SDK 的声明顺序写即可.
//
//	com.Call(uintptr(i.vtbl().DrawLine),
//		com.U(uintptr(unsafe.Pointer(i))),           // this
//		com.U64(pointArg(p0)),                       // 按值传递的 8 字节结构体
//		com.U64(pointArg(p1)),
//		com.U(uintptr(unsafe.Pointer(brush))),
//		com.F32(strokeWidth),                        // 浮点形参
//		com.U(uintptr(unsafe.Pointer(style))))
//
// # 各架构的落位规则
//
//   - amd64(Microsoft x64 约定): 第 N 个实参占第 N 个**位置槽**(this 是位置 0). 位置 0~3 里整数/指针
//     走 RCX/RDX/R8/R9、浮点走 XMM0~XMM3; 位置 4 及以后压栈(栈上的浮点占 8 字节槽, 值放低 4 字节).
//     最多 8 个实参(4 个寄存器槽 + 4 个栈槽).
//   - 386(stdcall, 全部走栈): 实参按顺序压栈, 每个占 1 个 32 位槽; 64 位值与 8 字节结构体拆成
//     低/高两个槽. 没有 XMM 的问题, 也不需要汇编转发器.
//
// 浮点形参一律用 F32/F64 包一层: 这一步顺带把"按值传的 float"与"指向浮点结构的指针"区分开 ——
// 后者照常走 U(指针).
//
// 注意: 位置 0~3 上"整数与浮点共用同一个位置编号"这一条是实测结论(教科书式的"整数/浮点各自独立
// 计数"在 D2D 上不成立, 见 CallWithFloat32 的说明), 本函数就是按位置编号实现的.

// Arg 是 Call 的一个实参.
//
// 用 U / U64 / F32 / F64 构造, 不要直接写结构体字面量(字段含义与架构相关).
type Arg struct {
	kind argKind
	bits uint64 // 整数/指针值, 或浮点的位模式
}

type argKind uint8

const (
	argInt     argKind = iota // 整数/指针/句柄: amd64 占 1 个 8 字节槽, 386 占 1 个 4 字节槽
	argInt64                  // 64 位整数或按值传递的 8 字节结构体: amd64 占 1 个槽, 386 占 2 个槽
	argFloat32                // float32 形参
	argFloat64                // float64 形参
)

// U 描述一个整数/指针实参(句柄、COM 接口指针、长度、枚举等都算).
func U(v uintptr) Arg {
	return Arg{kind: argInt, bits: uint64(v)}
}

// U64 描述一个 64 位整数或**按值传递的 8 字节结构体**实参(例如 D2D1_POINT_2F 打包后的值).
func U64(v uint64) Arg {
	return Arg{kind: argInt64, bits: v}
}

// F32 描述一个 float32 形参.
func F32(f float32) Arg {
	return Arg{kind: argFloat32, bits: uint64(math.Float32bits(f))}
}

// F64 描述一个 float64 形参.
func F64(f float64) Arg {
	return Arg{kind: argFloat64, bits: math.Float64bits(f)}
}
