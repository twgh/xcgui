package wnd_test

import (
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(wnd.GetTitle(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(wnd.SetTop(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	fmt.Println(wnd.GetHWND("", "管理器"))
}
