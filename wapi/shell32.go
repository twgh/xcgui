package wapi

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	// Library.
	shell32 = syscall.NewLazyDLL("shell32.dll")

	// Functions.
	dragQueryFileW = shell32.NewProc("DragQueryFileW")
	dragFinish     = shell32.NewProc("DragFinish")
	dragQueryPoint = shell32.NewProc("DragQueryPoint")
	shellExecuteW  = shell32.NewProc("ShellExecuteW")
)

// DragQueryFileW 检索由成功的拖放操作产生的文件路径.
//	@Description: 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragQueryFileW.
//	@param hDrop 句柄.
//	@param iFile 文件索引.
//	@param lpszFile 返回的文件路径.
//	@param cch 接收的文件路径的字符数, 通常为260.
//	@return int 返回文件路径的字符数.
//
func DragQueryFileW(hDrop int, iFile int, lpszFile *string, cch uint32) int {
	buf := make([]uint16, cch)
	r, _, _ := dragQueryFileW.Call(uintptr(hDrop), uintptr(iFile), common.Uint16SliceDataPtr(&buf), uintptr(cch))
	*lpszFile = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// DragFinish 释放系统分配用于将文件名传输到应用程序的内存.
//	@Description: 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragFinish.
//	@param hDrop 句柄.
//
func DragFinish(hDrop int) {
	dragFinish.Call(uintptr(hDrop))
}

// DragQueryPoint 检索在拖放文件时鼠标指针的位置.
//	@Description: 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragQueryPoint.
//	@param hDrop 句柄.
//	@param ppt 接收鼠标指针的坐标.
//	@return bool 如果拖放发生在窗口的客户区, 返回true；否则返回false.
//
func DragQueryPoint(hDrop int, ppt *xc.POINT) bool {
	r, _, _ := dragQueryPoint.Call(uintptr(hDrop), uintptr(unsafe.Pointer(ppt)))
	return int(r) != 0
}

// ShellExecuteW 对指定文件执行操作.
//	@Description: 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew.
//	@param hwnd 用于显示 UI 或错误消息的父窗口的句柄。如果操作与窗口无关，则此值可以为0.
//	@param lpOperation 填“open”则打开lpFlie文档.
//	@param lpFile 想用关联的程序打印或打开的一个程序名或文件名.
//	@param lpParameters 如果lpFile是一个可执行文件，则这个字串包含了传递给执行程序的参数.
//	@param lpDirectory 想使用的默认路径完整路径.
//	@param nShowCmd 定义了如何显示启动程序的常数值, xcc.SW_.
//	@return int 如果函数成功，则返回大于32的值。如果函数失败，则返回指示失败原因的错误值.
//
func ShellExecuteW(hwnd int, lpOperation, lpFile, lpParameters, lpDirectory string, nShowCmd xcc.SW_) int {
	r, _, _ := shellExecuteW.Call(uintptr(hwnd), common.StrPtr(lpOperation), common.StrPtr(lpFile), common.StrPtr(lpParameters), common.StrPtr(lpDirectory), uintptr(nShowCmd))
	return int(r)
}
