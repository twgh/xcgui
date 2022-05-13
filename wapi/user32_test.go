package wapi

import (
	"testing"
)

func TestGetDesktopWindow(t *testing.T) {
	t.Log(GetDesktopWindow())
}

func TestMessageBoxW(t *testing.T) {
	id := MessageBoxW(0, "context", "title", MB_CanaelTryContinue|MB_IconInformation)
	switch id {
	case ID_Cancel:
		t.Log("Cancel")
	case ID_TryAgain:
		t.Log("TryAgain")
	case ID_Continue:
		t.Log("Continue")
	default:
		t.Log(id)
	}
}

func TestFindWindowExW(t *testing.T) {
	t.Log(FindWindowExW(0, 0, "", "任务管理器"))
	t.Log(FindWindowExW(0, 0, "TaskManagerWindow", ""))
	t.Log(FindWindowExW(0, 0, "TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	t.Log(GetWindowTextLengthW(FindWindowExW(0, 0, "", "任务管理器")))
}
