package wapi

import (
	"syscall"
)

var (
	// Library.
	Ole32 = syscall.NewLazyDLL("ole32.dll")

	// Functions.
	procCoInitializeEx = Ole32.NewProc("CoInitializeEx")
)

// CoInitializeEx 初始化 COM 库供调用线程使用，设置线程的并发模型，并根据需要为线程创建新的单元。
//
// pvReserved: 保留参数, 必须为0.
//
// dwCoInit: 线程的并发模型和初始化选项, 可以是以下值的组合: wapi.COINIT_ .
//   - 默认值为 COINIT_MULTITHREADED.
//   - 不能同时设置 COINIT_APARTMENTTHREADED 和 COINIT_MULTITHREADED 标志。
//
// 返回: S_OK, S_FALSE, E_INVALIDARG, E_OUTOFMEMORY, E_UNEXPECTED, RPC_E_CHANGED_MODE.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex.
func CoInitializeEx(pvReserved uintptr, dwCoInit COINIT_) syscall.Errno {
	r, _, _ := procCoInitializeEx.Call(pvReserved, uintptr(dwCoInit))
	return syscall.Errno(r)
}

// COINIT_ 确定用于对此线程创建的对象的传入调用的并发模型。 此并发模型可以是单元线程模型，也可以是多线程模型。
type COINIT_ uint32

const (
	COINIT_APARTMENTTHREADED COINIT_ = 0x2 // 初始化单元线程对象并发的线程
	COINIT_MULTITHREADED     COINIT_ = 0x0 // 初始化多线程对象并发的线程
	COINIT_DISABLE_OLE1DDE   COINIT_ = 0x4 // 禁用 DDE 以支持 OLE1
	COINIT_SPEED_OVER_MEMORY COINIT_ = 0x8 // 增加内存使用量，以尝试提高性能
)

// COM 初始化状态码

const (
	S_OK               syscall.Errno = 0x00000000
	S_FALSE            syscall.Errno = 0x00000001
	RPC_E_CHANGED_MODE syscall.Errno = 0x80010106 // 并发模型冲突（如 CoInitializeEx 调用不兼容）
	E_INVALIDARG       syscall.Errno = 0x80070057 // 参数无效
	E_OUTOFMEMORY      syscall.Errno = 0x8007000E // 内存不足
	E_UNEXPECTED       syscall.Errno = 0x8000FFFF // 未知意外错误
)
