package wapi

import (
	"syscall"
	"unsafe"
)

var (
	// Library.
	Ole32 = syscall.NewLazyDLL("ole32.dll")

	// Functions.
	procCoInitializeEx   = Ole32.NewProc("CoInitializeEx")
	procCoCreateInstance = Ole32.NewProc("CoCreateInstance")
)

// CoCreateInstance 创建并默认初始化与指定 CLSID 关联的类的单个对象。
//
// rclsid: 要创建的对象的 CLSID。
//
// pUnkOuter: 如果为 NULL，则指示对象不是作为聚合的一部分创建的。 如果不是 NULL，则指向聚合对象的 IUnknown 接口的指针 (控制 IUnknown) 。
//
// dwClsContext: 管理新创建对象的代码将在其中运行的上下文。这些值取自枚举 CLSCTX。
//
// riid: 要用于与对象通信的接口的 IID。
//
// ppv: 指向接口指针变量的指针。如果创建成功，*ppv 包含请求的接口指针。
//
// 返回: S_OK 表示成功，其他值表示失败。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(rclsid, pUnkOuter uintptr, dwClsContext uint32, riid uintptr) (unsafe.Pointer, syscall.Errno) {
	var ppv unsafe.Pointer
	r, _, _ := procCoCreateInstance.Call(rclsid, pUnkOuter, uintptr(dwClsContext), riid, uintptr(unsafe.Pointer(&ppv)))
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return ppv, S_OK
}

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
