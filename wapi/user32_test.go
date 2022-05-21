package wapi_test

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestGetDesktopWindow(t *testing.T) {
	fmt.Println(wapi.GetDesktopWindow())
}

func TestMessageBoxW(t *testing.T) {
	id := wapi.MessageBoxW(0, "context", "title", wapi.MB_CanaelTryContinue|wapi.MB_IconInformation)
	switch id {
	case wapi.ID_Cancel:
		fmt.Println("Cancel")
	case wapi.ID_TryAgain:
		fmt.Println("TryAgain")
	case wapi.ID_Continue:
		fmt.Println("Continue")
	default:
		fmt.Println(id)
	}
}

func TestFindWindowExW(t *testing.T) {
	fmt.Println(wapi.FindWindowExW(0, 0, "", "任务管理器"))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", ""))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	// 取任务管理器的窗口标题长度
	fmt.Println(wapi.GetWindowTextLengthW(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "")))
}

func TestClientToScreen(t *testing.T) {
	a := app.New(true)
	w := window.NewWindow(0, 0, 300, 300, "", 0, xcc.Window_Style_Default)

	pt := xc.POINT{X: 0, Y: 0}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	fmt.Println(pt)

	a.ShowAndRun(w.Handle)
	a.Exit()
}
