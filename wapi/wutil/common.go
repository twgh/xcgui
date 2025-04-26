package wutil

import (
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
		length := wapi.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}
		files = append(files, filePath)
		i++ // 索引+1检索下一个文件
	}
	wapi.DragFinish(hDropInfo)
	return files
}

// IsKeyPressed 使用 wapi.GetAsyncKeyState 判断按键是否按下.
//
// vkCode: 按键的虚拟键码. xcc.VK_ .
func IsKeyPressed(vkCode int32) bool {
	return wapi.GetAsyncKeyState(vkCode) < 0
}

// IsComboKey 判断组合键是否按下.
func IsComboKey(mainKey int32, modifierKeys ...int32) bool {
	for _, modKey := range modifierKeys {
		if !IsKeyPressed(modKey) {
			return false
		}
	}
	return IsKeyPressed(mainKey)
}
