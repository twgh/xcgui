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

// 空函数指针必须安全返回 0(否则会 CALL 到地址 0).
func TestCallWithNilFn(t *testing.T) {
	if CallWithFloat32(0, [4]uintptr{}, [4]float32{}) != 0 {
		t.Error("fn=0 应返回 0")
	}
	if CallWithFloat64(0, [4]uintptr{}, [4]float64{}) != 0 {
		t.Error("fn=0 应返回 0")
	}
}
