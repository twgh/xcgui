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
	shellExecuteW  *syscall.LazyProc
)

func init() {
	// Library.
	shell32 = syscall.NewLazyDLL("shell32.dll")

	// Functions.
	dragQueryFileW = shell32.NewProc("DragQueryFileW")
	dragFinish = shell32.NewProc("DragFinish")
	dragQueryPoint = shell32.NewProc("DragQueryPoint")
	shellExecuteW = shell32.NewProc("ShellExecuteW")
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

// 对指定文件执行操作. 如果函数成功，则返回大于 32 的值。如果函数失败，则返回指示失败原因的错误值.
//
// hwnd: 用于显示 UI 或错误消息的父窗口的句柄。如果操作与窗口无关，则此值可以为0.
//
// lpOperation: 填“open”则打开lpFlie文档. 其它操作详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew.
//
// lpFile: 想用关联的程序打印或打开的一个程序名或文件名.
//
// lpParameters: 如果lpFile是一个可执行文件，则这个字串包含了传递给执行程序的参数.
//
// lpDirectory: 想使用的默认路径完整路径.
//
// nShowCmd: 定义了如何显示启动程序的常数值, SW_.
func ShellExecuteW(hwnd int, lpOperation, lpFile, lpParameters, lpDirectory string, nShowCmd int) int {
	r, _, _ := shellExecuteW.Call(uintptr(hwnd), StrPtr(lpOperation), StrPtr(lpFile), StrPtr(lpParameters), StrPtr(lpDirectory), uintptr(nShowCmd))
	return int(r)
}
