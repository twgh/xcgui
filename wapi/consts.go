package wapi

import "syscall"

const (
	MAX_PATH = 260

	ERROR_SUCCESS          syscall.Errno = 0
	ERROR_FILE_NOT_FOUND   syscall.Errno = 2
	ERROR_NOT_ENOUGH_QUOTA syscall.Errno = 0x718 // 处理此命令的配额不够。
)
