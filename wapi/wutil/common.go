package wutil

import (
	"syscall"

	"github.com/twgh/xcgui/wapi"
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
