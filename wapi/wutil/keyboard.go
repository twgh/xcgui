package wutil

import "github.com/twgh/xcgui/wapi"

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
