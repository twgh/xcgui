package wapi

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

var (
	// Kernel32 为应用程序提供最基础的操作系统功能，如内存管理、进程/线程控制、文件输入输出等底层系统服务。
	Kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// Functions.
	sleep                   = Kernel32.NewProc("Sleep")
	sleepEx                 = Kernel32.NewProc("SleepEx")
	closeHandle             = Kernel32.NewProc("CloseHandle")
	globalLock              = Kernel32.NewProc("GlobalLock")
	globalSize              = Kernel32.NewProc("GlobalSize")
	globalUnlock            = Kernel32.NewProc("GlobalUnlock")
	globalAlloc             = Kernel32.NewProc("GlobalAlloc")
	lstrcpyW                = Kernel32.NewProc("lstrcpyW")
	globalFree              = Kernel32.NewProc("GlobalFree")
	getLastError            = Kernel32.NewProc("GetLastError")
	getModuleHandleW        = Kernel32.NewProc("GetModuleHandleW")
	procGetCurrentProcessId = Kernel32.NewProc("GetCurrentProcessId")
	procGetModuleHandleEx   = Kernel32.NewProc("GetModuleHandleExW")
	procGetCurrentThreadId  = Kernel32.NewProc("GetCurrentThreadId")
	mulDiv                  = Kernel32.NewProc("MulDiv")
	procVirtualAlloc        = Kernel32.NewProc("VirtualAlloc")
	procVirtualProtect      = Kernel32.NewProc("VirtualProtect")
	procMultiByteToWideChar = Kernel32.NewProc("MultiByteToWideChar")
)

// GetCurrentThreadId 检索调用线程的线程标识符。在线程终止之前，线程标识符将在整个系统中唯一标识线程。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := procGetCurrentThreadId.Call()
	return uint32(ret)
}

// GET_MODULE_HANDLE_EX_FLAG_ 在 GetModuleHandleEx 中使用.
type GET_MODULE_HANDLE_EX_FLAG_ uint32

const (
	GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS GET_MODULE_HANDLE_EX_FLAG_ = 0x00000004 // lpModuleName 参数是模块中的地址。

	// 无论调用 FreeLibrary 多少次，模块都会一直加载到进程终止。
	//
	// 此选项不能用于 GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT。
	GET_MODULE_HANDLE_EX_FLAG_PIN GET_MODULE_HANDLE_EX_FLAG_ = 0x00000001

	// 模块的引用计数不会递增。 此选项等效于 GetModuleHandleW 的行为。不要将检索到的模块句柄传递给 FreeLibrary 函数;这样做可能会导致 DLL 过早地取消映射。 有关详细信息，请参阅“备注”。
	//
	// 此选项不能用于 GET_MODULE_HANDLE_EX_FLAG_PIN。
	GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT GET_MODULE_HANDLE_EX_FLAG_ = 0x00000002
)

// GetModuleHandleEx 检索调用进程的进程标识符。在进程终止之前，进程标识符将在整个系统中唯一标识进程。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/libloaderapi/nf-libloaderapi-GetModuleHandleExW.
//
// dwFlags: 此参数可以是零个或以下一个或多个值: wapi.GET_MODULE_HANDLE_EX_FLAG_。如果模块的引用计数递增，调用方必须使用 FreeLibrary 函数在不再需要模块句柄时递减引用计数。
//
// lpModuleName: 使用 common.StrPtr 生成. 加载的模块的名称（.dll 或 .exe 文件），或模块中的地址（如果 dwFlags = GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS）。
//   - 对于模块名称，如果省略文件扩展名，则会追加默认库扩展名 .dll。 文件名字符串可以包含尾随点字符 （.），以指示模块名称没有扩展名。 该字符串不必指定路径。 指定路径时，请务必使用反斜杠（\），而不是正斜杠（/）。 名称与当前映射到调用进程的地址空间的模块的名称进行比较（大小写）。
//   - 如果此参数为空，则该函数将返回用于创建调用进程的文件的句柄（.exe 文件）。
//   - GetModuleHandleEx 函数不会检索使用 LOAD_LIBRARY_AS_DATAFILE 标志加载的模块的句柄。
func GetModuleHandleEx(dwFlags GET_MODULE_HANDLE_EX_FLAG_, lpModuleName string) uintptr {
	var phModule uintptr
	_, _, _ = procGetModuleHandleEx.Call(
		uintptr(dwFlags),
		common.StrPtr(lpModuleName),
		uintptr(unsafe.Pointer(&phModule)),
	)
	return phModule
}

// GetCurrentProcessId 检索调用进程的进程标识符。在进程终止之前，进程标识符将在整个系统中唯一标识进程。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/libloaderapi/nf-libloaderapi-GetCurrentProcessId.
func GetCurrentProcessId() uint32 {
	ret, _, _ := procGetCurrentProcessId.Call()
	return uint32(ret)
}

// GetModuleHandleW 检索指定模块的模块句柄。 模块必须已由调用进程加载.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew.
//
// lpModuleName: 加载的模块的名称 (.dll 或 .exe 文件). 如果省略文件扩展名，则会追加默认库扩展名 .dll。 文件名字符串可以包含尾随点字符 (.) ，以指示模块名称没有扩展名。 字符串不必指定路径。 指定路径时，请务必使用反斜杠 (\) ，而不是使用 /) (正斜杠。 名称 (大小写独立比较，) 当前映射到调用进程的地址空间的模块的名称。如果此参数为空， 则 GetModuleHandleW 返回用于创建调用进程 (.exe 文件) 的文件的句柄.
func GetModuleHandleW(lpModuleName string) uintptr {
	r, _, _ := getModuleHandleW.Call(common.StrPtr(lpModuleName))
	return r
}

// Sleep 暂停当前线程的执行，直到超时间隔结束。若要进入可警报等待状态，请使用 SleepEx 函数.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/synchapi/nf-synchapi-sleep.
//
// ms: 毫秒.
func Sleep(ms uint32) {
	sleep.Call(uintptr(ms))
}

// SleepEx 挂起当前线程，直到满足指定的条件。 发生以下情况之一时，将继续执行：
//   - 调用 I/O 完成回调函数.
//   - 异步过程调用 (APC) 排队到线程.
//   - 超时间隔已过.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/synchapi/nf-synchapi-sleepex.
//
// dwMilliseconds: 暂停执行的时间间隔（以毫秒为单位）.
//
// bAlertable: 如果此参数为 FALSE，则函数在超时期限过后才会返回。 如果发生 I/O 完成回调，该函数不会立即返回，并且不会执行 I/O 完成函数。 如果 APC 已排队到线程，该函数不会立即返回，并且不会执行 APC 函数。如果 参数为 TRUE，并且调用此函数的线程与调用扩展 I/O 函数 (ReadFileEx 或 WriteFileEx) 的线程相同，则当超时期限已过或发生 I/O 完成回调函数时，函数将返回 。 如果发生 I/O 完成回调，则调用 I/O 完成函数。 如果将 APC 排队到 queueUserAPC) (线程，则当超时期限已过或调用 APC 函数时，函数将返回 .
//
// 返回值: 如果指定的时间间隔过期，则返回值为零。如果函数由于一个或多个 I/O 完成回调函数而返回，则返回值 WAIT_IO_COMPLETION. 仅当 bAlertable 为 TRUE，并且调用 SleepEx 函数的线程与调用扩展 I/O 函数的线程相同时，才会发生这种情况.
func SleepEx(dwMilliseconds uint32, bAlertable bool) uint32 {
	r, _, _ := sleepEx.Call(uintptr(dwMilliseconds), common.BoolPtr(bAlertable))
	return uint32(r)
}

// CloseHandle 关闭一个内核对象.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/handleapi/nf-handleapi-closehandle.
//
// handle: 对象句柄.
func CloseHandle(handle uintptr) bool {
	r, _, _ := closeHandle.Call(handle)
	return r != 0
}

// GlobalLock 锁定一个全局内存对象并返回一个指向对象内存块第一个字节的指针.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalLock.
//
// hMem: 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
//
// 返回值: 如果函数成功，则返回值是指向内存块第一个字节的指针. 如果函数失败，则返回值为0.
func GlobalLock(hMem uintptr) uintptr {
	r, _, _ := globalLock.Call(hMem)
	return r
}

// GMEM_ 内存分配属性.
type GMEM_ uint32

const (
	GHND          GMEM_ = 0x0042 // 结合 GMEM_Moveable 和 GMEM_ZeroInit
	GMEM_Fixed    GMEM_ = 0x0000 // 分配固定内存。返回值是一个指针.
	GMEM_Moveable GMEM_ = 0x0002 // 分配可移动内存。内存块永远不会在物理内存中移动，但它们可以在默认堆内移动。返回值是内存对象的句柄。要将句柄转换为指针，请使用 GlobalLock 函数。此值不能与 GMEM_Fixed 结合使用.

	GMEM_ZeroInit GMEM_ = 0x0040 // 将内存内容初始化为零.
	GPTR          GMEM_ = 0x0040 // 结合 GMEM_Fixed 和 GMEM_ZeroInit
)

// GlobalAlloc 从堆中分配指定数量的字节.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalAlloc.
//
// uFlags: 内存分配属性。如果指定为零，则默认为 GMEM_Fixed. 该参数可以是以下值中的一个或多个: wapi.GMEM_ .
//
// dwBytes: 要分配的字节数。如果此参数为0并且uFlags参数指定 GMEM_Moveable ，则该函数返回标记为已丢弃的内存对象的句柄.
//
// 返回值: 如果函数成功，则返回值是新分配的内存对象的句柄. 如果函数失败，则返回值为0.
func GlobalAlloc(uFlags GMEM_, dwBytes uint) uintptr {
	r, _, _ := globalAlloc.Call(uintptr(uFlags), uintptr(dwBytes))
	return r
}

// GlobalUnlock 减少与使用 GMEM_Moveable 分配的内存对象关联的锁计数。此函数对使用 GMEM_Fixed 分配的内存对象没有影响.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalUnlock.
//
// hMem: 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
func GlobalUnlock(hMem uintptr) bool {
	r, _, _ := globalUnlock.Call(hMem)
	return r != 0
}

// GlobalSize 检索指定全局内存对象的当前大小，以字节为单位.
//   - 内存块的大小可能大于分配内存时请求的大小. 要验证指定对象的内存块是否未被丢弃, 请在调用 GlobalSize 之前使用 GlobalFlags 函数.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalSize.
//
// hMem: 全局内存对象的句柄。此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回.
//
// 返回值: 如果函数成功，则返回值是指定全局内存对象的大小，以字节为单位. 如果指定的句柄无效或对象已被丢弃，则返回值为0.
func GlobalSize(hMem uintptr) uint {
	r, _, _ := globalSize.Call(hMem)
	return uint(r)
}

// LstrcpyW 将字符串复制到缓冲区.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-LstrcpyW.
//
// lpString1: 用于接收 lpString2 参数指向的字符串内容的缓冲区. 缓冲区必须足够大以包含字符串，包括终止空字符.
//
// lpString2: 要复制的以 NullStr 结尾的字符串.
//
// 返回值: 如果函数成功，则返回值是指向缓冲区的指针. 如果函数失败，则返回值为0, 并且lpString1可能不是以 NullStr 结尾的.
func LstrcpyW(lpString1, lpString2 uintptr) uintptr {
	r, _, _ := lstrcpyW.Call(lpString1, lpString2)
	return r
}

// GlobalFree 释放指定的全局内存对象并使其句柄无效.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GlobalFree.
//
// hMem: 全局内存对象的句柄. 此句柄由 GlobalAlloc 或 GlobalReAlloc 函数返回. 释放使用 LocalAlloc 分配的内存是不安全的.
//
// 返回值: 如果函数成功, 则返回值为0. 如果函数失败, 则返回值等于全局内存对象的句柄.
func GlobalFree(hMem uintptr) uintptr {
	r, _, _ := globalFree.Call(hMem)
	return r
}

// GetLastError 检索调用线程的最后一个错误代码值。最后一个错误代码是在每个线程的基础上维护的。多个线程不会覆盖彼此的最后一个错误代码.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-GetLastError.
//
// 返回值: 返回值是调用线程的最后一个错误代码.
func GetLastError() int32 {
	r, _, _ := getLastError.Call()
	return int32(r)
}

// MulDiv 将两个 32 位值相乘，对得到的 64 位中间结果除以第三个 32 位值，最终结果四舍五入为最接近的整数。
//   - 物理像素是屏幕硬件上真实存在的发光点，逻辑像素是为了适应不同屏幕，抽象出来的一种“虚拟尺寸单位”。
//   - Windows 规定：96 DPI 是标准基准
//   - 将物理像素转为逻辑像素: MulDiv(物理像素, 96, dpi), dpi = 144 相当于 150% 缩放
//   - 将逻辑像素转为物理像素: MulDiv(逻辑像素, dpi, 96)
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winbase/nf-winbase-muldiv.
//
// nNumber: 被乘数.
//
// nNumerator: 乘数.
//
// nDenominator: 除数.
//
// 返回值:
//   - 如果函数成功，则返回值是乘除法的最终结果，四舍五入为最接近的整数。
//   - 如果函数失败，则返回值为 -1。失败的原因包括溢出或 nDenominator 参数为 0。
//   - 如果失败，请调用 GetLastError 获取更多错误信息.
func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
	r, _, _ := mulDiv.Call(
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator),
	)
	return int32(r)
}

// PAGE_ 内存页保护属性, 在 VirtualAlloc 和 VirtualProtect 中使用.
type PAGE_ uint32

const (
	PAGE_NOACCESS          PAGE_ = 0x01 // 禁止访问该区域. 如果尝试读取、写入或执行, 将引发访问冲突.
	PAGE_READONLY          PAGE_ = 0x02 // 允许读取.
	PAGE_READWRITE         PAGE_ = 0x04 // 允许读写.
	PAGE_WRITECOPY         PAGE_ = 0x08 // 允许写入, 但写时复制.
	PAGE_EXECUTE           PAGE_ = 0x10 // 允许执行.
	PAGE_EXECUTE_READ      PAGE_ = 0x20 // 允许执行和读取.
	PAGE_EXECUTE_READWRITE PAGE_ = 0x40 // 允许执行、读取和写入.
	PAGE_EXECUTE_WRITECOPY PAGE_ = 0x80 // 允许执行和写入, 但写时复制.
)

// MEM_ 内存分配类型, 在 VirtualAlloc 中使用.
type MEM_ uint32

const (
	MEM_COMMIT  MEM_ = 0x00001000 // 为指定的保留内存页分配内存.
	MEM_RESERVE MEM_ = 0x00002000 // 保留进程的虚拟地址空间范围.
	MEM_RELEASE MEM_ = 0x00008000 // 释放指定的保留内存页.
	MEM_FREE    MEM_ = 0x00010000 // 释放已占用的内存页.
	MEM_RESET   MEM_ = 0x00080000 // 将指定的内存页重置为未提交状态.
)

// VirtualAlloc 在调用进程的虚拟地址空间中保留、提交或更改内存区域的状态. 函数分配的内存初始化为零.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/memoryapi/nf-memoryapi-virtualalloc.
//
// lpAddress: 要分配的区域的起始地址. 如果为 0, 则由系统决定分配位置.
//
// dwSize: 要分配的区域的大小, 以字节为单位.
//
// flAllocationType: 内存分配类型, 是以下值的组合: wapi.MEM_.
//
// flProtect: 内存页保护属性, 是以下值之一: wapi.PAGE_.
//
// 返回值: 如果函数成功, 则返回值是分配的内存区域的基址; 如果函数失败, 则返回值为 0.
func VirtualAlloc(lpAddress, dwSize uintptr, flAllocationType MEM_, flProtect PAGE_) uintptr {
	r, _, _ := procVirtualAlloc.Call(lpAddress, dwSize, uintptr(flAllocationType), uintptr(flProtect))
	return r
}

// VirtualProtect 更改调用进程的虚拟地址空间中已提交页面的保护.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/memoryapi/nf-memoryapi-virtualprotect.
//
// lpAddress: 要更改访问保护的页面区域的起始地址.
//
// dwSize: 要更改访问保护的区域的大小, 以字节为单位.
//
// flNewProtect: 新的内存保护属性, 是以下值之一: wapi.PAGE_.
//
// lpflOldProtect: 指向一个变量的指针, 该变量接收指定页区域的第一个页的当前访问保护属性. 如果此参数为 nil, 则该函数会失败.
//
// 返回值: 如果函数成功, 则返回值为 true; 如果函数失败, 则返回值为 false.
func VirtualProtect(lpAddress, dwSize uintptr, flNewProtect PAGE_, lpflOldProtect *PAGE_) bool {
	r, _, _ := procVirtualProtect.Call(lpAddress, dwSize, uintptr(flNewProtect), uintptr(unsafe.Pointer(lpflOldProtect)))
	return r != 0
}

// CODE_PAGE_ 代码页, 在 MultiByteToWideChar 和 WideCharToMultiByte 中使用.
type CODE_PAGE_ uint32

const (
	CP_ACP        CODE_PAGE_ = 0     // 系统的默认 ANSI 代码页.
	CP_OEMCP      CODE_PAGE_ = 1     // 系统的默认 OEM 代码页.
	CP_MACCP      CODE_PAGE_ = 2     // 系统的默认 Macintosh 代码页.
	CP_THREAD_ACP CODE_PAGE_ = 3     // 当前线程的 ANSI 代码页.
	CP_SYMBOL     CODE_PAGE_ = 42    // 符号代码页.
	CP_UTF7       CODE_PAGE_ = 65000 // UTF-7 代码页.
	CP_UTF8       CODE_PAGE_ = 65001 // UTF-8 代码页.
)

// MultiByteToWideChar 将字符串映射到 UTF-16 (宽字符) 字符串.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/stringapiset/nf-stringapiset-multibytetowidechar.
//
// codePage: 代码页, 是以下值之一: wapi.CODE_PAGE_.
//
// dwFlags: 指示转换类型的标志. 如果该值为 0, 则函数在转换失败时返回 0, 并且不会更改输出缓冲区. 详见官方文档.
//
// lpMultiByteStr: 使用 common.StrPtr 生成. 要转换的字符串.
//
// cbMultiByte: lpMultiByteStr 字符串的大小, 以字节为单位. 如果为 -1, 则该字符串以 NullStr 结尾, 并且函数会计算长度.
//
// lpWideCharStr: 接收转换后字符串的缓冲区. 如果为 0, 则该函数返回所需的缓冲区大小.
//
// cchWideChar: lpWideCharStr 缓冲区的大小, 以字符为单位(UTF-16 码元).
//
// 返回值: 如果函数成功, 则返回值是写入缓冲区的字符数; 如果函数失败, 则返回值为 0.
func MultiByteToWideChar(codePage CODE_PAGE_, dwFlags uint32, lpMultiByteStr uintptr, cbMultiByte int32, lpWideCharStr uintptr, cchWideChar int32) int32 {
	r, _, _ := procMultiByteToWideChar.Call(uintptr(codePage), uintptr(dwFlags), lpMultiByteStr, uintptr(cbMultiByte), lpWideCharStr, uintptr(cchWideChar))
	return int32(r)
}
