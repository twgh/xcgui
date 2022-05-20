package wnd_test

import (
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wnd"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	t.Log(wnd.GetTitle(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	t.Log(wnd.SetTop(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	t.Log(wnd.GetHWND("", "管理器"))
}
