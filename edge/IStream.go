package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// IStream 允许读取和写入数据流对象。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/nn-objidl-istream
type IStream struct {
	Vtbl *IStreamVtbl
}

// NewStreamMem 创建内存流对象.
//
// data: 用于设置内存流的初始内容, 如果此参数为 nil，则返回的内存流没有任何初始内容。
func NewStreamMem(data []byte) (*IStream, error) {
	streamPtr, err := wapi.SHCreateMemStream(data)
	if err != nil {
		return nil, err
	}
	if streamPtr == 0 {
		return nil, errors.New("createMemStream failed")
	}
	return NewIStreamByPtr(streamPtr), nil
}

// NewIStreamByPtr 从 streamPtr 创建一个新的 IStream 实例。
//
// streamPtr: stream指针.
func NewIStreamByPtr(streamPtr uintptr) *IStream {
	return (*IStream)(unsafe.Pointer(streamPtr))
}

type IStreamVtbl struct {
	IUnknownVtbl
	Read         ComProc
	Write        ComProc
	Seek         ComProc
	SetSize      ComProc
	CopyTo       ComProc
	Commit       ComProc
	Revert       ComProc
	LockRegion   ComProc
	UnlockRegion ComProc
	Stat         ComProc
	Clone        ComProc
}

func (i *IStream) GetPtr() uintptr {
	return uintptr(unsafe.Pointer(i))
}

func (i *IStream) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IStream) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *IStream) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Read 从当前搜寻指针开始，将指定数量的字节从流对象读取到内存中。
//
// buffer: 存储读取数据的缓冲区.
func (i *IStream) Read(buffer []byte) (int, error) {
	var bytesRead uint32
	hr, _, _ := i.Vtbl.Read.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if hr != 0 {
		return 0, syscall.Errno(hr)
	}
	return int(bytesRead), nil
}

// Write 从当前搜寻指针处开始，将指定数量的字节写入流对象。
//
// buffer: 存储写入数据的缓冲区.
func (i *IStream) Write(buffer []byte) (int, error) {
	var bytesWritten uint32
	hr, _, _ := i.Vtbl.Write.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	if hr != 0 {
		return 0, syscall.Errno(hr)
	}
	return int(bytesWritten), nil
}

// STREAM_SEEK 表示流式数据中的定位方式
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/ne-objidl-stream_seek
type STREAM_SEEK uint32

const (
	STREAM_SEEK_SET STREAM_SEEK = iota // 从流的起始位置开始
	STREAM_SEEK_CUR                    // 从当前指针位置开始
	STREAM_SEEK_END                    // 从流的末尾位置开始
)

// Seek 将查找指针更改为新位置。 新位置相对于流的开头、流的结束或当前查找指针。
//
// dlibMove: 相对于 whence 的偏移量
//
// dwOrigin: STREAM_SEEK 常量.
func (i *IStream) Seek(dlibMove int64, dwOrigin STREAM_SEEK) error {
	var newPosition uint64
	hr, _, _ := i.Vtbl.Seek.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dlibMove),
		uintptr(dwOrigin),
		uintptr(unsafe.Pointer(&newPosition)),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// SetSize 更改流对象的大小。
//
// size: 新大小.
func (i *IStream) SetSize(size uint64) error {
	hr, _, _ := i.Vtbl.SetSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(size),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// CopyTo 将指定的字节数从流中的当前搜索指针复制到其他流中的当前搜索指针。
//
// pstm: 指向目标流的指针。 pstm 指向的流可以是新流或源流的克隆。
//
// count: 要从源流复制的字节数。
//
// 返回值: 实际复制的字节数。
func (i *IStream) CopyTo(pstm *IStream, count uint64) (uint64, error) {
	var bytesWritten uint64
	hr, _, _ := i.Vtbl.CopyTo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(pstm)),
		uintptr(count),
		uintptr(0),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	if hr != 0 {
		return 0, syscall.Errno(hr)
	}
	return bytesWritten, nil
}

// STGC 表示存储提交标志，用于控制存储行为的位掩码类型
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/wtypes/ne-wtypes-stgc
type STGC uint32

const (
	// STGC_DEFAULT 默认存储行为，遵循标准提交规则
	STGC_DEFAULT STGC = 0

	// STGC_OVERWRITE 覆盖现有数据，即使存在版本冲突
	STGC_OVERWRITE STGC = 1

	// STGC_ONLYIFCURRENT 仅当对象当前未被修改时提交
	STGC_ONLYIFCURRENT STGC = 2

	// STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE
	// 仅提交到磁盘缓存（可能丢失数据），用于临时存储场景。
	// 注意：系统崩溃可能导致数据丢失，使用时必须明确风险
	STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE STGC = 4

	// STGC_CONSOLIDATE 合并存储空间，优化存储结构
	STGC_CONSOLIDATE STGC = 8
)

// Commit 提交流中的更改。
//
// flags: 控制提交对流对象的更改的方式: STGC 常量.
func (i *IStream) Commit(flags STGC) error {
	hr, _, _ := i.Vtbl.Commit.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(flags),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// Revert 撤销自上次 Commit 以来的更改。
func (i *IStream) Revert() error {
	hr, _, _ := i.Vtbl.Revert.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// LOCK 表示资源锁定类型，用于控制并发访问的位掩码类型.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/ne-objidl-locktype
type LOCK uint32

const (
	// LOCK_WRITE 写锁定，允许并发读取但阻止其他写入
	// 注意：与 LOCK_EXCLUSIVE 同时使用时将升级为独占锁
	LOCK_WRITE LOCK = 1

	// LOCK_EXCLUSIVE 独占锁定，禁止其他所有访问
	// 使用该模式会完全阻塞资源，建议设置超时机制
	LOCK_EXCLUSIVE LOCK = 2

	// LOCK_ONLYONCE 一次性锁定，确保资源只被锁定一次
	// 适用于初始化保护场景，重复锁定将返回错误
	LOCK_ONLYONCE LOCK = 4
)

// LockRegion 限制对流中指定字节范围的访问。
//
// offset: 要锁定的流的字节偏移量。
//
// count: 要锁定的字节数。
//
// lockType: 锁定类型: LOCK 常量。
func (i *IStream) LockRegion(offset uint64, count uint64, lockType LOCK) error {
	hr, _, _ := i.Vtbl.LockRegion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(offset),
		uintptr(count),
		uintptr(lockType),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// UnlockRegion 删除对以前使用 IStream.LockRegion 限制的字节范围的访问限制。
//
// offset: 要解锁的流的字节偏移量。
//
// count: 要解锁的字节数。
//
// lockType: 以前对范围施加的访问限制类型: LOCK 常量。
func (i *IStream) UnlockRegion(offset uint64, count uint64, lockType LOCK) error {
	hr, _, _ := i.Vtbl.UnlockRegion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(offset),
		uintptr(count),
		uintptr(lockType),
	)
	if hr != 0 {
		return syscall.Errno(hr)
	}
	return nil
}

// STATSTG 存储统计信息结构，用于描述存储对象的属性和状态.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/ns-objidl-statstg
type STATSTG struct {
	// 对象名称指针 (OLE字符串)，使用后需调用 wapi.CoTaskMemFree 释放.
	//
	// 转换字符串方法：syscall.UTF16PtrToString(PwcsName)
	//
	// 若要不返回此成员，请在调用返回 STATSTG 结构的方法时指定 STATFLAG_NONAME 值.
	PwcsName *uint16

	// 对象类型标识，对应 STGTY 常量
	Type STGTY

	// 对象大小（以字节为单位）
	CbSize uint64

	// 最后修改时间，使用 syscall.Filetime 结构保存原始数据.
	//
	// 转换方法：time.Unix(0, Mtime.Nanoseconds()).UTC()
	Mtime syscall.Filetime

	// 创建时间，FILETIME 格式存储
	Ctime syscall.Filetime

	// 最后访问时间，FILETIME 格式存储
	Atime syscall.Filetime

	// 访问模式标志，对应 STGM 常量组合
	GrfMode STGM

	// 支持的锁类型，对应 LOCK 位掩码
	GrfLocksSupported LOCK

	// 类标识符（CLSID），128位全局唯一标识
	Clsid syscall.GUID

	// 状态位标志，用于对象持久化状态
	GrfStateBits uint32

	// 保留字段，必须初始化为0
	Reserved uint32
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

// STATFLAG 控制统计信息获取方式的枚举，决定返回数据的详细程度和操作行为.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/ns-objidl-statstg
type STATFLAG uint32

const (
	// STATFLAG_DEFAULT 默认获取模式，返回所有可用信息.
	//
	// 包含对象名称、完整属性和打开状态信息.
	STATFLAG_DEFAULT STATFLAG = 0

	// STATFLAG_NONAME 省略对象名称信息，提升性能.
	//
	// 使用此标志时，返回结构体的 pwcsName 字段将为nil.
	//
	// 注意：调用方不得尝试释放或访问名称字段.
	STATFLAG_NONAME STATFLAG = 1

	// STATFLAG_NOOPEN 以最小开销方式获取统计信息.
	//
	// 避免完全打开对象，可能返回部分元数据.
	//
	// 警告：可能影响返回数据的实时性.
	STATFLAG_NOOPEN STATFLAG = 2
)

// STGTY 表示存储对象类型枚举，用于标识不同的存储实体类型
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/ne-objidl-stgty
type STGTY uint32

const (
	// STGTY_STORAGE 表示复合存储对象，可包含其他存储对象和流.
	//
	// 用于目录式的层次化存储结构，对应IStorage接口.
	STGTY_STORAGE STGTY = 1

	// STGTY_STREAM 表示字节流对象，用于存储序列化数据.
	//
	// 基本数据存储单元，对应IStream接口.
	STGTY_STREAM STGTY = 2

	// STGTY_LOCKBYTES 表示字节数组对象，支持原子访问.
	//
	// 用于低级别字节操作，对应ILockBytes接口.
	STGTY_LOCKBYTES STGTY = 3

	// STGTY_PROPERTY 表示属性存储对象，用于元数据管理.
	//
	// 存储键值对属性信息，需配合属性集使用.
	STGTY_PROPERTY STGTY = 4
)

// Stat 获取流的统计信息。
//
// statFlag: 控制统计信息获取方式，决定返回数据的详细程度和操作行为. STATFLAG 常量.
func (i *IStream) Stat(statFlag STATFLAG) (*STATSTG, error) {
	var stat STATSTG
	hr, _, _ := i.Vtbl.Stat.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&stat)),
		uintptr(statFlag),
	)
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	return &stat, nil
}

// Clone 创建流的副本。
func (i *IStream) Clone() (*IStream, error) {
	var clone *IStream
	hr, _, _ := i.Vtbl.Clone.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&clone)),
	)
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	return clone, nil
}

// GetBytes 获取流的内容。
func (i *IStream) GetBytes() ([]byte, error) {
	const bufferSize = 4096
	var content []byte
	for {
		buffer := make([]byte, bufferSize)
		n, hr := i.Read(buffer)
		if hr != nil && !errors.Is(hr, wapi.S_FALSE) {
			return nil, errors.New("stream read failed: " + hr.Error())
		}
		if n == 0 {
			break
		}
		content = append(content, buffer[:n]...)
	}
	return content, nil
}

// GetBytesAndRelease 获取流的内容并释放流。
func (i *IStream) GetBytesAndRelease() ([]byte, error) {
	defer i.Release()
	return i.GetBytes()
}
