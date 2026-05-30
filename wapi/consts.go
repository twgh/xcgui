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

// DROPEFFECT_* 表示有关拖放操作效果的信息。DoDragDrop 函数以及 IDropSource 和 IDropTarget 中的许多方法都会使用此枚举的值。
// https://learn.microsoft.com/en-us/windows/win32/com/dropeffect-constants

const (
	DROPEFFECT_NONE   uint32 = 0          // 放置目标无法接收该数据。
	DROPEFFECT_COPY   uint32 = 1          // 放置操作将生成副本，拖拽源处的原始数据保持不变。
	DROPEFFECT_MOVE   uint32 = 2          // 拖拽源应移除该数据。
	DROPEFFECT_LINK   uint32 = 4          // 拖拽源应创建指向原始数据的链接。
	DROPEFFECT_SCROLL uint32 = 0x80000000 // 目标区域即将开始滚动或正处于滚动状态。此值需与其他值结合使用。
)
