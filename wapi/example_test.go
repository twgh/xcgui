package wapi_test

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func ExampleLoadImageW() {
	hIcon := wapi.LoadImageW(0, common.StrPtr("C:\\Windows\\System32\\OneDrive.ico"), wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	fmt.Println("hIcon:", hIcon)
	fmt.Println("LastErr:", syscall.GetLastError())
}

func ExampleMessageBoxW() {
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

func ExampleFindWindowExW() {
	fmt.Println(wapi.FindWindowExW(0, 0, "", "任务管理器"))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", ""))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "任务管理器"))
}

func ExampleClientToScreen() {
	a := app.New(true)
	w := window.New(0, 0, 300, 300, "", 0, xcc.Window_Style_Default)

	pt := wapi.POINT{X: 0, Y: 0}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	fmt.Println(pt)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func ExampleShellExecuteW() {
	// 打开指定网址
	wapi.ShellExecuteW(0, "open", "https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew", "", "", xcc.SW_SHOWNORMAL)
	// 打开指定文件
	wapi.ShellExecuteW(0, "open", "C:\\Windows\\System32\\calc.exe", "", "", xcc.SW_SHOWNORMAL)
}

func ExampleChooseColorW() {
	var lpCustColors [16]uint32
	cc := wapi.ChooseColor{
		LStructSize:    36,
		HwndOwner:      0,
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &lpCustColors[0],
		Flags:          wapi.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	ret := wapi.ChooseColorW(&cc)
	fmt.Println(ret)
	fmt.Println(cc.RgbResult) // rgb颜色
	fmt.Println(lpCustColors) // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在
}

func ExampleGetCursorPos() {
	var pt wapi.POINT
	wapi.GetCursorPos(&pt)
	fmt.Println(pt)
}
