package wutil

import (
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

type getWindowCountData struct {
	Pid   int       // 指定进程id
	Hwnds []uintptr // 窗口句柄数组
}

// GetWindows 枚举指定进程的所有窗口句柄.
//
// pid: 进程ID, 如果填-1, 则使用当前进程id. 如果填0, 则枚举所有进程的窗口句柄.
func GetWindows(pid ...int) []uintptr {
	var pid2 int
	if len(pid) > 0 {
		pid2 = pid[0]
	}
	if pid2 == -1 {
		pid2 = int(wapi.GetCurrentProcessId())
	}
	data := getWindowCountData{Pid: pid2}

	// 枚举窗口
	wapi.EnumWindows(syscall.NewCallback(cbGetWindowCount), uintptr(unsafe.Pointer(&data)))
	return data.Hwnds
}

func cbGetWindowCount(hwnd, lParam uintptr) uintptr {
	data := (*getWindowCountData)(unsafe.Pointer(lParam))
	if data.Pid == 0 { // 枚举所有进程
		data.Hwnds = append(data.Hwnds, hwnd)
		return 1 // 继续枚举
	}

	// 枚举指定进程
	processId := wapi.GetWindowThreadProcessId(hwnd)
	if int(processId) == data.Pid {
		data.Hwnds = append(data.Hwnds, hwnd)
	}
	return 1 // 继续枚举
}
