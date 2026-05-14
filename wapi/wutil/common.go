package wutil

import (
	"errors"
	"syscall"

	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

// GetDropFiles 获取拖放进来的文件.
//
// hDropInfo: 拖放信息句柄.
func GetDropFiles(hDropInfo uintptr) []string {
	var filePath string
	files := make([]string, 0)
	var i uint32
	for {
		length := wapi.DragQueryFileW(hDropInfo, i, &filePath, syscall.MAX_PATH)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}
		files = append(files, filePath)
		i++ // 索引+1检索下一个文件
	}
	wapi.DragFinish(hDropInfo)
	return files
}

// GetTaskbarRect 获取任务栏矩形.
func GetTaskbarRect() (xc.RECT, error) {
	var rc xc.RECT
	// 1. 查找任务栏窗口句柄
	// 任务栏的类名固定为 "Shell_TrayWnd"
	hwnd := wapi.FindWindowW("Shell_TrayWnd", "")
	if hwnd == 0 {
		return rc, errors.New("未找到任务栏窗口")
	}

	// 2. 获取窗口矩形区域
	// GetWindowRect 返回的是屏幕坐标
	if ok := wapi.GetWindowRect(hwnd, &rc); !ok {
		return rc, errors.New("获取窗口矩形区域失败")
	}

	return rc, nil
}
