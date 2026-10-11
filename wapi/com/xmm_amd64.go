//go:build amd64

package com

import (
	"math"
	"reflect"
	"syscall"
	"unsafe"
)

// xmmArgs 是交给汇编转发器的参数块(float32 版).
//
// **字段顺序/类型不要改**: xmm_amd64.s 按固定偏移读它(fn=0, this=8, a1=16, a2=24, a3=32,
// x0=40, x1=44, x2=48, x3=52). 浮点用位模式存放, 由汇编搬进 XMM.
type xmmArgs struct {
	fn   uintptr
	this uintptr
	a1   uintptr
	a2   uintptr
	a3   uintptr
	x0   uint32
	x1   uint32
	x2   uint32
	x3   uint32
}

// xmmArgs64 同 xmmArgs, 但浮点是 float64(x0=40, x1=48, x2=56, x3=64).
type xmmArgs64 struct {
	fn   uintptr
	this uintptr
	a1   uintptr
	a2   uintptr
	a3   uintptr
	x0   uint64
	x1   uint64
	x2   uint64
	x3   uint64
}

// xmmArgsStack 在 xmmArgs 之上多 4 个**栈参数**(第 5~8 个参数; x64 约定里超出 4 个寄存器的部分压栈).
//
// 偏移: fn=0, this=8, a1=16, a2=24, a3=32, x0=40, x1=44, x2=48, x3=52, **stack=56(4×8)**.
//
// 用不到的栈槽填 0 即可 —— 多写的字节落在转发器自己的栈帧里, 目标只读它声明要的参数.
type xmmArgsStack struct {
	fn, this, a1, a2, a3 uintptr
	x0, x1, x2, x3       uint32
	stack                [4]uintptr
}

// xmmTrampoline32 / xmmTrampoline64 由 xmm_amd64.s 实现, 按 **Windows x64 约定**接收参数
// (RCX = 参数块指针, 返回 RAX).
//
// 这里声明成无参函数只是为了让 Go 能取到它们的入口地址 —— **不要直接调用**(直接调用走 Go 约定,
// 而实现读的是 RCX). 实际调用路径是 syscall.SyscallN(pc, &args):
//
//	要的不是 SyscallN 设寄存器(它设不了 XMM), 而是它内部会把调用**切到系统栈**上执行 ——
//	Windows API 的栈探测看的是 TEB 里的栈范围, 而 Go 的 goroutine 栈不在其中: 直接在 goroutine
//	栈上 CALL, 简单 API(GetTickCount)能过, 稍复杂的(GetSystemMetrics)会因栈探测落到非法地址而崩.
//	Go 自己的 syscall 就是靠 runtime.cgocall 切系统栈来避免这个坑的.
//
//go:noescape
func xmmTrampoline32()

//go:noescape
func xmmTrampoline64()

// xmmTrampoline32RetFloat 同 xmmTrampoline32, 但目标方法**按值返回 float**(返回值在 XMM0),
// 它把 XMM0 搬到 RAX 交回来.
//
//go:noescape
func xmmTrampoline32RetFloat()

// xmmTrampoline32Stack 同 xmmTrampoline32, 另把参数块里的 4 个栈参数放到第 5~8 个参数位(见 .s).
//
//go:noescape
func xmmTrampoline32Stack()

var (
	xmmTrampoline32PC         = reflect.ValueOf(xmmTrampoline32).Pointer()
	xmmTrampoline64PC         = reflect.ValueOf(xmmTrampoline64).Pointer()
	xmmTrampoline32RetFloatPC = reflect.ValueOf(xmmTrampoline32RetFloat).Pointer()
	xmmTrampoline32StackPC    = reflect.ValueOf(xmmTrampoline32Stack).Pointer()
)

// CallWithFloat32 调用"整数/指针 + float32 混合形参"的 COM 方法(仅 amd64; 386 见包注释).
//
// Windows x64 约定里整数/指针参数走 RCX/RDX/R8/R9、浮点参数走 XMM0~XMM3; 而 syscall.SyscallN 只能
// 设置通用寄存器 —— 形参里的 float32 位打包根本传不进去(实测会崩或无效, 见包注释"参数传递"一节).
// 本函数借一小段汇编转发器把两类寄存器都设置好再调用, 于是这类方法也能封装了.
//
//	ints: 依次填入 RCX/RDX/R8/R9 —— 第 0 个通常是 this 指针, 其余按签名里整数/指针实参的出现顺序
//	flts: 依次填入 XMM0~XMM3 —— **槽号要按接口实测**, 见下面的"槽位不一定等于参数序号"
//
// 用不到的位置随便填(被调方只读它签名里声明的那些). 返回 HRESULT 或方法的整数返回值.
//
// # 槽位不一定等于参数序号(实测)
//
// 按"整数与浮点各自独立计数"的教科书写法, ID2D1SolidColorBrush::SetOpacity(FLOAT opacity) 的浮点
// 参数应在 XMM0(this 走 RCX) —— **实测却是 XMM1**: 逐个槽单独填 0.5, 只有 XMM1 生效(读回 0.500),
// XMM0/XMM2/XMM3 都是 0. 也就是说这类接口把参数**按位置**编号(this 占位置 0 → 第一个浮点参数落在
// XMM1). 不同接口来源(D3D/DWrite/win32)可能不同, 封装前先用"写值 + 读回"验证一次:
//
//	// 逐槽试: 哪个槽能读回目标值, 就是它
//	for slot := 0; slot < 4; slot++ {
//		var flts [4]float32
//		flts[slot] = 0.5
//		CallWithFloat32(setterFn, [4]uintptr{this}, flts)
//		// 读回...
//	}
//
// 另一个坑: 目标方法内部可能因对象状态而崩(与传参无关) —— 例如 ID2D1RenderTarget::SetDpi 在一个
// DPI=0 的渲染目标上直接崩, 所以它不是好的验证对象; 纯属性 setter(如画刷 opacity)才是.
func CallWithFloat32(fn uintptr, ints [4]uintptr, flts [4]float32) uintptr {
	if fn == 0 {
		return 0
	}
	args := xmmArgs{
		fn:   fn,
		this: ints[0],
		a1:   ints[1],
		a2:   ints[2],
		a3:   ints[3],
		x0:   math.Float32bits(flts[0]),
		x1:   math.Float32bits(flts[1]),
		x2:   math.Float32bits(flts[2]),
		x3:   math.Float32bits(flts[3]),
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline32PC, uintptr(unsafe.Pointer(&args)))
	return r
}

// CallWithFloat32RetFloat 同 CallWithFloat32, 但目标方法**按值返回 float**(在 XMM0 里返回),
// 本函数把那个返回值交回来.
//
// 例: ID2D1SolidColorBrush::GetOpacity() 的返回值就是 FLOAT —— 用 syscall 调只会拿到 RAX(垃圾).
func CallWithFloat32RetFloat(fn uintptr, ints [4]uintptr, flts [4]float32) float32 {
	if fn == 0 {
		return 0
	}
	args := xmmArgs{
		fn:   fn,
		this: ints[0],
		a1:   ints[1],
		a2:   ints[2],
		a3:   ints[3],
		x0:   math.Float32bits(flts[0]),
		x1:   math.Float32bits(flts[1]),
		x2:   math.Float32bits(flts[2]),
		x3:   math.Float32bits(flts[3]),
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline32RetFloatPC, uintptr(unsafe.Pointer(&args)))
	return math.Float32frombits(uint32(r))
}

// CallWithFloat32Stack 同 CallWithFloat32, 并支持**最多 4 个栈参数**(第 5~8 个参数).
//
// 用不到的栈槽填 0 —— 多写的字节落在转发器自己的栈帧里, 目标方法只读它声明要的个数, 无害.
//
// 典型用例(都是 D2D 的"按位置计数"布局, 见包注释第 3 条):
//
//	DrawGeometry(geometry, brush, FLOAT strokeWidth, style)     线宽→XMM3, style→栈[0]
//	DrawLine(p0, p1, brush, FLOAT strokeWidth, style)           线宽→**栈[0]**(位置 4, 前 4 个 XMM 槽已被
//	                                                            位置 0~3 占掉), style→栈[1]
//
// 栈上的**浮点**参数同样是 8 字节槽, 值放低 4 字节(调用方用 math.Float32bits 填进 uintptr).
//
// 这套布局是真机实测出来的: 按"独立计数"把线宽塞 XMM0、样式塞 R9 时, 调用不崩但整帧渲染被搞坏
// (窗口全黑); 按位置计数(线宽 XMM3 + 样式压栈)则正确画出描边 —— 与 SetOpacity 读 XMM1 的规律一致.
func CallWithFloat32Stack(fn uintptr, ints [4]uintptr, flts [4]float32, stack [4]uintptr) uintptr {
	if fn == 0 {
		return 0
	}
	args := xmmArgsStack{
		fn:    fn,
		this:  ints[0],
		a1:    ints[1],
		a2:    ints[2],
		a3:    ints[3],
		x0:    math.Float32bits(flts[0]),
		x1:    math.Float32bits(flts[1]),
		x2:    math.Float32bits(flts[2]),
		x3:    math.Float32bits(flts[3]),
		stack: stack,
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline32StackPC, uintptr(unsafe.Pointer(&args)))
	return r
}

// CallWithFloat64 同 CallWithFloat32, 但浮点形参是 float64(填 XMM0~XMM3 用 MOVSD).
func CallWithFloat64(fn uintptr, ints [4]uintptr, flts [4]float64) uintptr {
	if fn == 0 {
		return 0
	}
	args := xmmArgs64{
		fn:   fn,
		this: ints[0],
		a1:   ints[1],
		a2:   ints[2],
		a3:   ints[3],
		x0:   math.Float64bits(flts[0]),
		x1:   math.Float64bits(flts[1]),
		x2:   math.Float64bits(flts[2]),
		x3:   math.Float64bits(flts[3]),
	}
	r, _, _ := syscall.SyscallN(xmmTrampoline64PC, uintptr(unsafe.Pointer(&args)))
	return r
}
