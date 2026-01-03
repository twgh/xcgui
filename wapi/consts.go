package wapi

import "syscall"

const (
	MAX_PATH = 260

	ERROR_SUCCESS          syscall.Errno = 0
	ERROR_FILE_NOT_FOUND   syscall.Errno = 2
	ERROR_NOT_ENOUGH_QUOTA syscall.Errno = 0x718 // 处理此命令的配额不够。
	ERROR_INVALID_INDEX    syscall.Errno = 0x585 // 索引无效。
)

const (
	S_OK               syscall.Errno = 0x00000000
	S_FALSE            syscall.Errno = 0x00000001
	RPC_E_CHANGED_MODE syscall.Errno = 0x80010106 // 并发模型冲突（如 CoInitializeEx 调用不兼容）
	E_INVALIDARG       syscall.Errno = 0x80070057 // 参数无效
	E_OUTOFMEMORY      syscall.Errno = 0x8007000E // 内存不足
	E_UNEXPECTED       syscall.Errno = 0x8000FFFF // 未知意外错误
	E_FAIL             syscall.Errno = 0x80004005 // 未指定的失败
	RO_E_CLOSED        syscall.Errno = 0x80000013 // 对象已关闭, 无法再使用
)
