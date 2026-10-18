package com

import (
	"math"
	"unsafe"
)

// 跨架构的 COM 调用: 按"被调方签名里实参的出现顺序"逐个描述参数, 由各架构的实现负责把它们摆到
// 该去的地方(amd64: 寄存器 + shadow space + 栈; 386: 全部压栈).
//
// 与 CallWithFloat32 / CallWithFloat32Stack(仅 amd64, 需要调用方自己算寄存器槽位)不同, 这里不用
// 关心槽位: this 也算第一个实参, 照着 Windows SDK 的声明顺序写即可.
//
//	com.Call(uintptr(i.vtbl().DrawLine),
//		com.U(uintptr(unsafe.Pointer(i))),           // this
//		com.StructOf(&p0),                           // 按值传递的结构体
//		com.StructOf(&p1),
//		com.U(uintptr(unsafe.Pointer(brush))),
//		com.F32(strokeWidth),                        // 浮点形参
//		com.U(uintptr(unsafe.Pointer(style))))
//
// # 各架构的落位规则
//
//   - amd64(Microsoft x64 约定): 第 N 个实参占第 N 个**位置槽**(this 是位置 0). 位置 0~3 里整数/指针
//     走 RCX/RDX/R8/R9、浮点走 XMM0~XMM3; 位置 4 及以后压栈(栈上的浮点占 8 字节槽, 值放低 4 字节).
//     最多 8 个实参(4 个寄存器槽 + 4 个栈槽). 结构体超过 8 字节按引用传(占一个槽), 8 字节及以下
//     打包进一个槽.
//   - 386(stdcall, 全部走栈): 实参按顺序压栈, 每个占 1 个 32 位槽; 64 位值占 2 个槽, 结构体按
//     4 字节槽展开. 没有 XMM 的问题, 也不需要汇编转发器.
//
// 浮点形参一律用 F32/F64 包一层: 这一步顺带把"按值传的 float"与"指向浮点结构的指针"区分开 ——
// 后者照常走 U(指针). 同理, 按值传的结构体要用 Struct/StructOf, 传它的指针才用 U.
//
// 注意: 位置 0~3 上"整数与浮点共用同一个位置编号"这一条是实测结论(教科书式的"整数/浮点各自独立
// 计数"在 D2D 上不成立, 见 CallWithFloat32 的说明), 本函数就是按位置编号实现的.

// Arg 是 Call 的一个实参.
//
// 用 U / U64 / F32 / F64 / Struct 构造, 不要直接写结构体字面量(字段含义与架构相关).
type Arg struct {
	kind argKind
	bits uint64 // 整数/指针值, 或浮点的位模式
	// 仅 argStruct 使用: 结构体地址与大小. 故意用 unsafe.Pointer 而不是 uintptr,
	// 这样 GC 知道它引用着那块内存, 调用期间不会把它回收.
	ptr  unsafe.Pointer
	size uintptr
}

type argKind uint8

const (
	argInt     argKind = iota // 整数/指针/句柄: amd64 占 1 个 8 字节槽, 386 占 1 个 4 字节槽
	argInt64                  // 64 位整数: amd64 占 1 个槽, 386 占 2 个槽
	argFloat32                // float32 形参
	argFloat64                // float64 形参
	argStruct                 // 按值传递的结构体, 由各架构按自身规则铺开
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

// Struct 描述一个**按值传递的结构体**实参(例如 WebView2 的 RECT / POINT).
//   - amd64/arm64: 超过 8 字节的结构体按引用传递(传指针), 8 字节及以下打包进一个位置槽;
//   - 386: 按 4 字节槽展开压栈(长度向上取整到 4 的倍数).
//
// p 必须是 **unsafe.Pointer**(而不是 uintptr): 调用期间 GC 要能看见这个引用, 否则对象可能被回收.
func Struct(p unsafe.Pointer, size uintptr) Arg {
	return Arg{kind: argStruct, ptr: p, size: size}
}

// StructOf 是 Struct 的便捷写法, 取值的地址与大小: com.StructOf(&rect).
func StructOf[T any](v *T) Arg {
	return Struct(unsafe.Pointer(v), unsafe.Sizeof(*v))
}

// structBits 把 size(<=8) 个字节按小端拼成一个 64 位整数(amd64/arm64 上"按值小结构体"走一个槽).
func structBits(p unsafe.Pointer, size uintptr) uint64 {
	var bits uint64
	for i := uintptr(0); i < size && i < 8; i++ {
		bits |= uint64(*(*byte)(unsafe.Add(p, i))) << (8 * i)
	}
	return bits
}

// normArg 把 argStruct 折成等价的整数实参: 超过 8 字节的结构体按引用传递(指针), 否则打包成一个槽.
//
// 386 不走这里 —— 那边的结构体要展开成一串 32 位槽, 见 call_386.go.
func normArg(arg Arg) (argKind, uint64) {
	if arg.kind != argStruct {
		return arg.kind, arg.bits
	}
	if arg.size > 8 {
		return argInt, uint64(uintptr(arg.ptr))
	}
	return argInt, structBits(arg.ptr, arg.size)
}
