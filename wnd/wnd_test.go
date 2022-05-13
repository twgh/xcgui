package wnd

import (
	"github.com/twgh/xcgui/wapi"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	t.Log(GetWindowTitle(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	t.Log(SetTop(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	t.Log(GetWindowHWND("", "管理器"))
}
