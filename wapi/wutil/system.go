package wutil

import "github.com/twgh/xcgui/wapi"

// GetScreenWidth 取屏幕宽度.
func GetScreenWidth() int32 {
	return wapi.GetSystemMetrics(wapi.SM_CXSCREEN)
}

// GetScreenHeight 取屏幕高度.
func GetScreenHeight() int32 {
	return wapi.GetSystemMetrics(wapi.SM_CYSCREEN)
}
