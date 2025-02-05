package wutil

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// HIcon 从文件加载图标. 返回 HICON 句柄.
//   - 可用于需要 HICON 句柄的函数, 如设置托盘图标.
//   - 如果失败, 可使用 syscall.GetLastError() 获取错误信息.
//   - 当图标句柄不再使用时, 可使用 wapi.DestroyIcon 函数释放.
//
// iconPath: 图标路径.
func HIcon(iconPath string) uintptr {
	return wapi.LoadImageW(0, common.StrPtr(iconPath), wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE)
}
