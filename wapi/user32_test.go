package wapi_test

import (
	"fmt"
	"github.com/twgh/xcgui/tf"
	"testing"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
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

func TestFindWindowW(t *testing.T) {
	fmt.Println(wapi.FindWindowW("", "任务管理器"))
	fmt.Println(wapi.FindWindowW("TaskManagerWindow", ""))
	fmt.Println(wapi.FindWindowW("TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	// 取任务管理器的窗口标题长度
	fmt.Println(wapi.GetWindowTextLengthW(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "")))
}

func TestClientToScreen(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		pt := wapi.POINT{X: 0, Y: 0}
		wapi.ClientToScreen(w.GetHWND(), &pt)
		fmt.Println(pt)
	})
}

func TestGetCursorPos(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(20, 50, 450, 300, "GetCursorPos", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			var pt wapi.POINT
			wapi.GetCursorPos(&pt)
			a.MessageBox("GetCursorPos", fmt.Sprintf("x: %d  y: %d", pt.X, pt.Y), xcc.MessageBox_Flag_Ok, w.GetHWND(), xcc.Window_Style_Default)
			return 0
		})
	})
}

func TestIsWindow(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(20, 50, 100, 30, "IsWindow", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			w.MessageBox("IsWindow", fmt.Sprintf("IsWindow: %v", wapi.IsWindow(w.GetHWND())), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		})
	})
}

func TestPostQuitMessage(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(20, 50, 200, 30, "PostQuitMessage", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			a.EnableAutoExitApp(false)
			w.CloseWindow()
			wapi.PostQuitMessage(0)
			return 0
		})
	})
}

func TestSetForegroundWindow(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(20, 50, 200, 30, "SetForegroundWindow", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			hwnd := wapi.FindWindowW("TaskManagerWindow", "")
			if !wapi.IsWindow(hwnd) {
				a.Alert("提示", "请先打开任务管理器, 本按钮功能是把任务管理器窗口激活")
				return 0
			}
			wapi.SetForegroundWindow(hwnd)
			return 0
		})
	})
}
