package wapi_test

import (
	"github.com/twgh/xcgui/wapi"
	"testing"
)

func TestGetDesktopWindow(t *testing.T) {
	t.Log(wapi.GetDesktopWindow())
}

func TestMessageBoxW(t *testing.T) {
	id := wapi.MessageBoxW(0, "context", "title", wapi.MB_CanaelTryContinue|wapi.MB_IconInformation)
	switch id {
	case wapi.ID_Cancel:
		t.Log("Cancel")
	case wapi.ID_TryAgain:
		t.Log("TryAgain")
	case wapi.ID_Continue:
		t.Log("Continue")
	default:
		t.Log(id)
	}
}

func TestFindWindowExW(t *testing.T) {
	t.Log(wapi.FindWindowExW(0, 0, "", "任务管理器"))
	t.Log(wapi.FindWindowExW(0, 0, "TaskManagerWindow", ""))
	t.Log(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	t.Log(wapi.GetWindowTextLengthW(wapi.FindWindowExW(0, 0, "", "任务管理器")))
}
