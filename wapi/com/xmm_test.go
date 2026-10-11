//go:build amd64

package com

import (
	"syscall"
	"testing"
	"unsafe"
)

// 整数/指针参数的转发必须与 syscall.SyscallN 完全一致(汇编里只多设置了浮点寄存器).
func TestCallWithIntsMatchesSyscall(t *testing.T) {
	proc := syscall.NewLazyDLL("user32.dll").NewProc("GetSystemMetrics")
	if err := proc.Find(); err != nil {
		t.Skipf("找不到 user32.GetSystemMetrics: %v", err)
	}
	for _, index := range []uintptr{0 /*SM_CXSCREEN*/, 1 /*SM_CYSCREEN*/, 78 /*SM_CXVIRTUALSCREEN*/} {
		got := CallWithFloat32(proc.Addr(), [4]uintptr{index}, [4]float32{})
		want, _, _ := proc.Call(index)
		if got != want {
			t.Errorf("GetSystemMetrics(%d): 转发得到 %d, syscall 得到 %d", index, got, want)
		}
	}
}

// 指针参数也一样: 取当前进程句柄, 两条路必须给出同一个值.
func TestCallWithPointerMatchesSyscall(t *testing.T) {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetCurrentProcess")
	if err := proc.Find(); err != nil {
		t.Skipf("找不到 kernel32.GetCurrentProcess: %v", err)
	}
	got := CallWithFloat32(proc.Addr(), [4]uintptr{}, [4]float32{})
	want, _, _ := proc.Call()
	if got != want {
		t.Errorf("GetCurrentProcess: 转发得到 %d, syscall 得到 %d", got, want)
	}
}

// 参数块布局是汇编硬编码的偏移, 不能改 —— 这里把它固化成断言.
func TestXmmArgsLayout(t *testing.T) {
	if got := unsafe.Offsetof(xmmArgs{}.x0); got != 40 {
		t.Errorf("xmmArgs.x0 偏移 = %d, 期望 40", got)
	}
	if got := unsafe.Offsetof(xmmArgs{}.x3); got != 52 {
		t.Errorf("xmmArgs.x3 偏移 = %d, 期望 52", got)
	}
	if got := unsafe.Offsetof(xmmArgs64{}.x0); got != 40 {
		t.Errorf("xmmArgs64.x0 偏移 = %d, 期望 40", got)
	}
	if got := unsafe.Offsetof(xmmArgs64{}.x3); got != 64 {
		t.Errorf("xmmArgs64.x3 偏移 = %d, 期望 64", got)
	}
}

// 第 5 个参数必须真的落到栈上 —— 用 ReadProcessMemory 读自己进程里的一个变量验证:
// 它的第 5 个参数是"实际读到的字节数", 传对的话写回 8.
//
// (第 6~8 个栈槽的顺序不在这里验 —— 没有现成的"第 6 个参数是出参"的无副作用 API;
// 真实用例的端到端验证交给 d2d 的探针: DrawLine 的 lineWidth 在栈[0]、strokeStyle 在栈[1],
// 抓图能同时看出"线宽对不对"和"线帽圆不圆".)
func TestCallWithFloat32Stack(t *testing.T) {
	read := syscall.NewLazyDLL("kernel32.dll").NewProc("ReadProcessMemory")
	if err := read.Find(); err != nil {
		t.Skipf("找不到 kernel32.ReadProcessMemory: %v", err)
	}
	secret := uint64(0x1122334455667788)
	var buf uint64
	var n uintptr
	const curProcess = ^uintptr(0) // GetCurrentProcess() 伪句柄 = -1
	r := CallWithFloat32Stack(read.Addr(),
		[4]uintptr{curProcess, uintptr(unsafe.Pointer(&secret)), uintptr(unsafe.Pointer(&buf)), 8},
		[4]float32{},
		[4]uintptr{uintptr(unsafe.Pointer(&n))}) // 第 5 个参数; 其余槽填 0
	if r == 0 {
		t.Fatal("ReadProcessMemory 返回 0(失败)")
	}
	if buf != secret {
		t.Errorf("读回 %#x, 期望 %#x", buf, secret)
	}
	if n != 8 {
		t.Errorf("第 5 个参数(读到的字节数) = %d, 期望 8 —— 说明它没被正确压栈", n)
	}
}

// 空函数指针必须安全返回 0(否则会 CALL 到地址 0).
func TestCallWithNilFn(t *testing.T) {
	if CallWithFloat32(0, [4]uintptr{}, [4]float32{}) != 0 {
		t.Error("fn=0 应返回 0")
	}
	if CallWithFloat64(0, [4]uintptr{}, [4]float64{}) != 0 {
		t.Error("fn=0 应返回 0")
	}
}
