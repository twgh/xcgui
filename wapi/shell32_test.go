package wapi_test

import (
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xcc"
	"syscall"
	"testing"
)

func TestShellExecuteW(t *testing.T) {
	// 打开指定网址
	wapi.ShellExecuteW(0, "open", "https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew", "", "", xcc.SW_SHOWNORMAL)
	// 打开指定文件
	wapi.ShellExecuteW(0, "open", "C:\\Windows\\System32\\calc.exe", "", "", xcc.SW_SHOWNORMAL)
}

func TestSHBrowseForFolderW(t *testing.T) {
	displayNameBuffer := make([]uint16, 260)
	lpszTitle, _ := syscall.UTF16PtrFromString("显示在对话框中树视图控件上方的文本")

	bi := wapi.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: &displayNameBuffer[0],
		LpszTitle:      lpszTitle,
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}

	pidl := wapi.SHBrowseForFolderW(&bi)
	if pidl == 0 {
		fmt.Println("用户取消了选择")
		return
	}

	var pszPath string
	ret := wapi.SHGetPathFromIDListW(pidl, &pszPath)
	fmt.Println(ret)
	fmt.Println("用户选择的文件夹完整路径:", pszPath)
	fmt.Println("用户选择的文件夹的名称:", syscall.UTF16ToString(displayNameBuffer))
}
