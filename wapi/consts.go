package wapi

import "syscall"

const (
	MAX_PATH = 260

	ERROR_SUCCESS        syscall.Errno = 0
	ERROR_FILE_NOT_FOUND syscall.Errno = 2
)
