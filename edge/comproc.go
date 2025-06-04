package edge

import (
	"syscall"
)

// ComProc 存储COM过程。
type ComProc uintptr

// NewComProc 从Go函数创建一个新的COM过程。
func NewComProc(fn interface{}) ComProc {
	return ComProc(syscall.NewCallback(fn))
}
