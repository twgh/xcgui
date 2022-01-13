package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Library.
	shell32 *syscall.LazyDLL

	// Functions.
	dragQueryFileW *syscall.LazyProc
	dragFinish     *syscall.LazyProc
	dragQueryPoint *syscall.LazyProc
)

func init() {
	// Library.
	shell32 = syscall.NewLazyDLL("shell32.dll")

	// Functions.
	dragQueryFileW = shell32.NewProc("DragQueryFileW")
	dragFinish = shell32.NewProc("DragFinish")
	dragQueryPoint = shell32.NewProc("DragQueryPoint")
}

// 检索由成功的拖放操作产生的文件路径, 返回文件路径的字符数.
//
// hDrop: 句柄.
//
// iFile: 文件索引.
//
// lpszFile: 返回的文件路径.
//
// cch: 接收的文件路径的字符数, 通常为260.
func DragQueryFileW(hDrop, iFile int, lpszFile *string, cch int) int {
	buf := make([]uint16, cch)
	r, _, _ := dragQueryFileW.Call(uintptr(hDrop), uintptr(iFile), Uint16SliceDataPtr(&buf), uintptr(cch))
	*lpszFile = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 释放系统分配用于将文件名传输到应用程序的内存.
//
// hDrop: 句柄.
func DragFinish(hDrop int) {
	dragFinish.Call(uintptr(hDrop))
}

// 检索在拖放文件时鼠标指针的位置. 如果拖放发生在窗口的客户区, 返回true；否则返回false.
//
// hDrop: 句柄.
//
// ppt: 接收鼠标指针的坐标.
func DragQueryPoint(hDrop int, ppt *POINT) bool {
	r, _, _ := dragQueryPoint.Call(uintptr(hDrop), uintptr(unsafe.Pointer(ppt)))
	return int(r) != 0
}
