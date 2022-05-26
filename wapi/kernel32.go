package wapi

import (
	"syscall"
)

var (
	// Library.
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// Functions.
	sleep        = kernel32.NewProc("Sleep")
	closeHandle  = kernel32.NewProc("CloseHandle")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalSize   = kernel32.NewProc("GlobalSize")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	globalAlloc  = kernel32.NewProc("GlobalAlloc")
	lstrcpyW     = kernel32.NewProc("lstrcpyW")
	globalFree   = kernel32.NewProc("GlobalFree")
)

// Sleep 延时.
//	@param ms 毫秒.
//	@return int
//
func Sleep(ms int) int {
	r, _, _ := sleep.Call(uintptr(ms))
	return int(r)
}

// CloseHandle 关闭一个内核对象.
//	@param handle 对象句柄.
//	@return int
//
func CloseHandle(handle int) int {
	r, _, _ := closeHandle.Call(uintptr(handle))
	return int(r)
}

// GlobalLock 锁定一个全局内存对象并返回一个指向对象内存块第一个字节的指针.
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalLock.
//	@param hMem 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
//	@return uintptr 如果函数成功，则返回值是指向内存块第一个字节的指针. 如果函数失败，则返回值为NULL.
//
func GlobalLock(hMem uintptr) uintptr {
	r, _, _ := globalLock.Call(hMem)
	return r
}

// GMEM_ 内存分配属性.
type GMEM_ uint32

const (
	GHND          GMEM_ = 0x0042 // 结合 GMEM_Moveable 和 GMEM_ZeroInit
	GMEM_Fixed    GMEM_ = 0x0000 // 分配固定内存。返回值是一个指针。
	GMEM_Moveable GMEM_ = 0x0002 // 分配可移动内存。内存块永远不会在物理内存中移动，但它们可以在默认堆内移动。返回值是内存对象的句柄。要将句柄转换为指针，请使用 GlobalLock 函数。此值不能与 GMEM_Fixed 结合使用。

	GMEM_ZeroInit GMEM_ = 0x0040 // 将内存内容初始化为零。
	GPTR          GMEM_ = 0x0040 // 结合 GMEM_Fixed 和 GMEM_ZeroInit
)

// GlobalAlloc 从堆中分配指定数量的字节.
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalAlloc.
//	@param uFlags 内存分配属性。如果指定为零，则默认为 GMEM_Fixed. 该参数可以是以下值中的一个或多个: wapi.GMEM_ .
//	@param dwBytes 要分配的字节数。如果此参数为0并且uFlags参数指定 GMEM_Moveable ，则该函数返回标记为已丢弃的内存对象的句柄.
//	@return uintptr 如果函数成功，则返回值是新分配的内存对象的句柄. 如果函数失败，则返回值为NULL.
//
func GlobalAlloc(uFlags GMEM_, dwBytes uint64) uintptr {
	r, _, _ := globalAlloc.Call(uintptr(uFlags), uintptr(dwBytes))
	return r
}

// GlobalUnlock 减少与使用 GMEM_Moveable 分配的内存对象关联的锁计数。此函数对使用 GMEM_Fixed 分配的内存对象没有影响.
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalUnlock.
//	@param hMem 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
//	@return bool
//
func GlobalUnlock(hMem uintptr) bool {
	r, _, _ := globalUnlock.Call(hMem)
	return r != 0
}

// GlobalSize 检索指定全局内存对象的当前大小，以字节为单位.
//	@Description 内存块的大小可能大于分配内存时请求的大小. 要验证指定对象的内存块是否未被丢弃, 请在调用 GlobalSize 之前使用 GlobalFlags 函数.
//	详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalSize.
//	@param hMem 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
//	@return uint64 如果函数成功，则返回值是指定全局内存对象的大小，以字节为单位. 如果指定的句柄无效或对象已被丢弃，则返回值为0.
//
func GlobalSize(hMem uintptr) uint64 {
	r, _, _ := globalSize.Call(hMem)
	return uint64(r)
}

// LstrcpyW 将字符串复制到缓冲区.
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-LstrcpyW.
//	@param lpString1 用于接收 lpString2 参数指向的字符串内容的缓冲区. 缓冲区必须足够大以包含字符串，包括终止空字符.
//	@param lpString2 要复制的以 null 结尾的字符串.
//	@return uintptr 如果函数成功，则返回值是指向缓冲区的指针. 如果函数失败，则返回值为NULL, 并且lpString1可能不是以 null 结尾的.
//
func LstrcpyW(lpString1, lpString2 uintptr) uintptr {
	r, _, _ := lstrcpyW.Call(lpString1, lpString2)
	return r
}

// GlobalFree 释放指定的全局内存对象并使其句柄无效.
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalFree.
//	@param hMem 全局内存对象的句柄. 此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回. 释放使用 LocalAlloc 分配的内存是不安全的.
//	@return uintptr 如果函数成功, 则返回值为NULL. 如果函数失败, 则返回值等于全局内存对象的句柄.
//
func GlobalFree(hMem uintptr) uintptr {
	r, _, _ := globalFree.Call(hMem)
	return r
}
