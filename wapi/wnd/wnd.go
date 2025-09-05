package wnd

import (
	"errors"
	"strings"

	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

// SetTop 窗口_置顶.
//
// hWnd: 窗口真实句柄.
//
// b: 是否置顶.
func SetTop(hWnd uintptr, b bool) bool {
	hWndInsertAfter := wapi.HWND_TOPMOST
	if !b {
		hWndInsertAfter = wapi.HWND_NOTOPMOST
	}
	return wapi.SetWindowPos(hWnd, hWndInsertAfter, 0, 0, 0, 0, wapi.SWP_NOMOVE|wapi.SWP_NOSIZE)
}

// GetTitle 取窗口标题.
//
// hWnd: 窗口真实句柄.
func GetTitle(hWnd uintptr) string {
	dwSize := wapi.GetWindowTextLengthW(hWnd)
	if dwSize == 0 {
		return ""
	}
	dwSize++ // 必须算上空字符

	var title string
	wapi.GetWindowTextW(hWnd, &title, dwSize)
	return title
}

// GetHWND 取窗口句柄, 标题支持模糊. 返回窗口真实句柄.
//
// className: 窗口类名, 不支持模糊, 可空.
//
// title: 窗口标题, 可输入关键字, 支持模糊, 可空.
func GetHWND(className, title string) uintptr {
	if className == "" && title == "" {
		return 0
	}
	var hWnd uintptr
	for {
		hWnd = wapi.FindWindowExW(0, hWnd, className, "")
		if hWnd != 0 {
			titleName := strings.ToLower(GetTitle(hWnd))
			if strings.Contains(titleName, strings.ToLower(title)) {
				return hWnd
			}
		} else {
			break
		}
	}
	return 0
}

// SetWindowRound 设置窗口圆角.
//
// hwnd: 窗口句柄.
//
// radius: 圆角半径, 单位 px. 没有处理dpi, 需要的话自行传入处理后的值.
func SetWindowRound(hwnd uintptr, radius int32) error {
	// 获取窗口尺寸
	var rect xc.RECT
	if !wapi.GetWindowRect(hwnd, &rect) {
		return errors.New("failed to get window rect")
	}
	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top

	// 创建圆角区域
	rgn := wapi.CreateRoundRectRgn(0, 0, width, height, radius*2, radius*2)
	if rgn == 0 {
		return errors.New("failed to create round rect region")
	}
	// 应用区域到窗口
	if !wapi.SetWindowRgn(hwnd, rgn, true) {
		return errors.New("failed to set window region")
	}
	return nil
}
