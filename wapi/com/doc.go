// Package com 提供 COM 互操作的基础件: 过程调用(ComProc)、IUnknown(虚表 + 映射结构体)、
// 接口指针搬运, 以及跨架构的按序调用入口 Call(形参含浮点时也走它).
//
// 使用方:
//   - d2d(直通 Direct2D): 用它映射 ID2D1* 接口, 再按各接口的虚表结构体(定义在 d2d 包里)重解释调用;
//   - edge(WebView2): 它自有一套 ComProc, 但形参里带浮点/按值结构体的方法都走本包的 Call 了
//     (见 edge/edge_call.go).
//
// # 参数传递(386 与 amd64 不同)
//
// 首选 **Call**: 按被调方签名里的实参顺序(连 this 一起)用 U / Struct / F32 / F64 列出来, 各架构的
// 落位规则由它处理 —— 细节见 call.go. 下面两段是它背后的机制, 手写封装(以及读懂历史代码)时才会用到.
//
// 调用一律经过 syscall.SyscallN, 而它**只能设置通用寄存器**; 两种架构下的写法因此不同:
//
//   - **386**(stdcall, 参数全部走栈):
//     结构体(如 RECT/POINT)**按值拆成 4/2 个 32 位槽**逐个压栈即可;
//     float32/float64 也**按位拆成 32 位槽**压栈(float64 占两个槽) —— 这是完全符合 ABI 的做法,
//     Call 的 386 实现就是这么铺的(不需要汇编转发器).
//   - **amd64**(Microsoft x64 约定):
//     8 字节以上的结构体**按引用传递**, 所以"按值传结构体"的方法要**传指针**;
//     浮点参数走 XMM0~XMM3, 与通用寄存器是**两套独立**的 —— 用 SyscallN 直接传(哪怕是位打包)是
//     **无效的**: 实测 ID2D1RenderTarget::SetDpi(float, float) 位打包调用直接崩溃
//     (现场 rax/r9 = 0x42C00000, 正是 96.0f 的位模式 —— 值确实进了通用寄存器, 被调方从 XMM
//     读到的却是垃圾).
//     要用这类方法, 走本包的 **Call**(它内部就是靠下面这套转发器); 手写封装时也可以直接用
//     **CallWithFloat32 / CallWithFloat64 / CallWithFloat32RetFloat**(最后那个用于"按值返回 float"
//     的方法) —— 靠一小段汇编转发器把两类寄存器都摆好, 见下节.
//
// # 汇编转发器(仅 amd64): 把浮点送进去的三个坑
//
//  1. **必须由 syscall.SyscallN 调用转发器** —— 借它内部的 runtime.cgocall 切到**系统栈**: Windows 的
//     栈探测看的是 TEB 里的栈范围, 而 Go 的 goroutine 栈不在其中. 直接在 goroutine 栈上 CALL 目标:
//     简单 API(GetTickCount)能过, 稍复杂的(user32.GetSystemMetrics)会因栈探测落到非法地址而崩.
//  2. **转发器自己要把栈摆正**: 调用点 RSP 必须 16 字节对齐(否则被调方入口不是约定的 ≡8, 用到 SSE
//     对齐存取的 API 会崩), 并留 32 字节 shadow space. 实测这条路径进来时 RSP ≡ 0 (mod 16) ——
//     SUBQ $32 刚好, SUBQ $40 就崩.
//  3. **浮点槽号要按接口实测**: 教科书说法是"整数/浮点各自独立计数"(SetOpacity(FLOAT) 的浮点应在
//     XMM0, this 走 RCX), 但实测 ID2D1SolidColorBrush::SetOpacity 读的是 **XMM1** —— 逐个槽单独
//     填 0.5, 只有 XMM1 生效; 这类接口按"参数位置"编号(this 占位置 0). 换新接口前先用
//     "写一个值 + 用 getter 读回"验证一次槽位; 也要注意目标方法可能因**对象状态**而崩(例如 DPI=0
//     的渲染目标上 SetDpi 必崩), 那不是传参问题, 换纯属性 setter 验证.
//  4. **第 5 个及以后的参数要压栈**(x64 约定: 超出 RCX/RDX/R8/R9 的放 [RSP+32] 起, 前 32 字节是
//     shadow space): 用 CallWithFloat32Stack1("4 个整数/指针 + 浮点 + 1 个栈参数"). 实测用例
//     ID2D1RenderTarget::DrawGeometry(geometry, brush, FLOAT strokeWidth, ID2D1StrokeStyle *style) ——
//     按位置计数落位是 this→RCX / geometry→RDX / brush→R8 / strokeWidth→XMM3 / style→栈.
package com
