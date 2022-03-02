// 调用系统API.
// 封装对窗口的操作.
package wnd

import "github.com/twgh/xcgui/user32"

// 窗口_置顶.
//
// hWnd: 窗口真实句柄.
//
// b: 是否置顶.
func SetTop(hWnd int, b bool) bool {
	hWndInsertAfter := user32.HWND_TOPMOST
	if !b {
		hWndInsertAfter = user32.HWND_NOTOPMOST
	}
	return user32.SetWindowPos(hWnd, hWndInsertAfter, 0, 0, 0, 0, user32.SWP_NOMOVE|user32.SWP_NOSIZE)
}
