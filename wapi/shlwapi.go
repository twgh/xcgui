package wapi

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

var (
	// Library.
	Shlwapi = syscall.NewLazyDLL("shlwapi")

	// Functions.
	procSHCreateMemStream      = Shlwapi.NewProc("SHCreateMemStream")
	procSHCreateStreamOnFileEx = Shlwapi.NewProc("SHCreateStreamOnFileEx")
)

// SHCreateMemStream 创建内存流.
//
// initData: 此缓冲区的内容用于设置内存流的初始内容。 如果此参数为 NULL，则返回的内存流没有任何初始内容。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(initData []byte) uintptr {
	var pInit uintptr
	var cbInit uint32
	if len(initData) > 0 {
		cbInit = uint32(len(initData))
		pInit = uintptr(unsafe.Pointer(&initData[0]))
	}
	ret, _, _ := procSHCreateMemStream.Call(
		pInit,
		uintptr(cbInit),
	)
	return ret
}

// SHCreateStreamOnFileEx 打开或创建文件，并检索要读取或写入该文件的流。
//
// pszFile: 指定文件名。
//
// grfMode: 用于指定文件访问模式以及如何创建和删除公开流的对象的一个或多个 wapi.STGM 值。
//
// dwAttributes: 一个或多个标志值，用于在创建新文件时指定文件属性。wapi.FILE_ATTRIBUTE_, wapi.FILE_FLAG_ .
//
// fCreate: 可帮助与 grfMode 一起指定在创建流时应如何处理现有文件。
//
// pstmTemplate: 保留参数.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/shlwapi/nf-shlwapi-shcreatestreamonfileex
func SHCreateStreamOnFileEx(pszFile string, grfMode STGM, dwAttributes uint32, fCreate bool, pstmTemplate ...uintptr) (uintptr, error) {
	var pstmTemp uintptr
	if len(pstmTemplate) > 0 {
		pstmTemp = pstmTemplate[0]
	}
	var stream uintptr
	r, _, _ := procSHCreateStreamOnFileEx.Call(
		common.StrPtr(pszFile),
		uintptr(grfMode),
		uintptr(dwAttributes),
		common.BoolPtr(fCreate),
		pstmTemp,
		uintptr(unsafe.Pointer(&stream)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return stream, nil
}

// STGM 是指示创建和删除对象的条件以及对象的访问模式的标志。
//
// https://learn.microsoft.com/zh-cn/windows/win32/stg/stgm-constants
type STGM uint32

const (
	// STGM_READ 只读访问模式
	STGM_READ STGM = 0x00000000

	// STGM_WRITE 只写访问模式
	STGM_WRITE STGM = 0x00000001

	// STGM_READWRITE 读写访问模式
	STGM_READWRITE STGM = 0x00000002

	// STGM_SHARE_DENY_NONE 允许其他进程以任何模式打开
	STGM_SHARE_DENY_NONE STGM = 0x00000040

	// STGM_SHARE_DENY_READ 禁止其他进程读取访问
	STGM_SHARE_DENY_READ STGM = 0x00000030

	// STGM_SHARE_DENY_WRITE 禁止其他进程写入访问
	STGM_SHARE_DENY_WRITE STGM = 0x00000020

	// STGM_SHARE_EXCLUSIVE 独占访问，禁止其他进程访问
	STGM_SHARE_EXCLUSIVE STGM = 0x00000010

	// STGM_CREATE 如果不存在则创建新的存储对象
	STGM_CREATE STGM = 0x00001000

	// STGM_CONVERT 转换存储格式
	STGM_CONVERT STGM = 0x00020000

	// STGM_FAILIFTHERE 如果已存在则失败
	STGM_FAILIFTHERE STGM = 0x00000000

	// STGM_DIRECT 直接访问模式，绕过缓冲
	STGM_DIRECT STGM = 0x00000000

	// STGM_TRANSACTED 事务访问模式
	STGM_TRANSACTED STGM = 0x00010000

	// STGM_NOSCRATCH 不使用临时文件
	STGM_NOSCRATCH STGM = 0x00100000

	// STGM_NOSNAPSHOT 不创建快照
	STGM_NOSNAPSHOT STGM = 0x00200000

	// STGM_SIMPLE 简单访问模式
	STGM_SIMPLE STGM = 0x08000000

	// STGM_DIRECT_SWMR 直接单写多读模式
	STGM_DIRECT_SWMR STGM = 0x00400000

	// STGM_DELETEONRELEASE 释放时删除
	STGM_DELETEONRELEASE STGM = 0x04000000
)

// 文件属性常量，用于指定文件的属性
const (
	FILE_ATTRIBUTE_ARCHIVE   uint32 = 0x20   // 文件应存档，应用程序使用此属性标记文件以供备份或删除
	FILE_ATTRIBUTE_ENCRYPTED uint32 = 0x4000 // 文件或目录已加密. 对于文件，这意味着文件中的所有数据都已加密。 对于目录，这意味着加密是新创建的文件和子目录的默认值。
	FILE_ATTRIBUTE_HIDDEN    uint32 = 0x2    // 文件已隐藏，不要将其包含在普通目录列表中
	FILE_ATTRIBUTE_NORMAL    uint32 = 0x80   // 文件没有设置其他属性，仅当单独使用时有效
	FILE_ATTRIBUTE_OFFLINE   uint32 = 0x1000 // 文件数据不会立即可用，指示文件数据已移至脱机存储
	FILE_ATTRIBUTE_READONLY  uint32 = 0x1    // 文件是只读的，应用程序可以读取但无法写入或删除
	FILE_ATTRIBUTE_SYSTEM    uint32 = 0x4    // 文件是操作系统的一部分或专供其使用
	FILE_ATTRIBUTE_TEMPORARY uint32 = 0x100  // 文件用于临时存储
)

// 文件标志常量，用于CreateFile等函数的dwFlagsAndAttributes参数
const (
	FILE_FLAG_BACKUP_SEMANTICS   uint32 = 0x02000000 // 为备份或还原操作打开文件
	FILE_FLAG_DELETE_ON_CLOSE    uint32 = 0x04000000 // 文件将在关闭所有句柄后立即删除
	FILE_FLAG_NO_BUFFERING       uint32 = 0x20000000 // 文件或设备正在打开，没有系统缓存
	FILE_FLAG_OPEN_NO_RECALL     uint32 = 0x00100000 // 请求文件数据，但应继续位于远程存储中
	FILE_FLAG_OPEN_REPARSE_POINT uint32 = 0x00200000 // 不进行正常的重新分析点处理
	FILE_FLAG_OVERLAPPED         uint32 = 0x40000000 // 文件或设备正在为异步 I/O 打开或创建
	FILE_FLAG_POSIX_SEMANTICS    uint32 = 0x01000000 // 访问将按照 POSIX 规则进行
	FILE_FLAG_RANDOM_ACCESS      uint32 = 0x10000000 // 访问旨在随机访问
	FILE_FLAG_SESSION_AWARE      uint32 = 0x00800000 // 正在打开具有会话感知的文件或设备
	FILE_FLAG_SEQUENTIAL_SCAN    uint32 = 0x08000000 // 访问旨在从头到尾顺序进行
	FILE_FLAG_WRITE_THROUGH      uint32 = 0x80000000 // 写入操作将直接转到磁盘
)
