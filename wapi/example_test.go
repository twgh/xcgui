package wapi_test

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func ExampleLoadImageW() {
	hIcon := wapi.LoadImageW(0, "C:\\Windows\\System32\\OneDrive.ico", wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
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

func ExampleSHBrowseForFolderW() {
	buf := make([]uint16, 260)
	bi := wapi.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: common.Uint16SliceDataPtr(&buf),
		LpszTitle:      common.StrPtr("显示在对话框中树视图控件上方的文本"),
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	ret := wapi.SHGetPathFromIDListW(wapi.SHBrowseForFolderW(&bi), &pszPath)
	fmt.Println(ret)
	fmt.Println("pszPath:", pszPath)                           // 用户选择的文件夹完整路径
	fmt.Println("PszDisplayName:", syscall.UTF16ToString(buf)) // 用户选择的文件夹的名称
}

func ExampleSHGetPathFromIDListW() {
	buf := make([]uint16, 260)
	bi := wapi.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: common.Uint16SliceDataPtr(&buf),
		LpszTitle:      common.StrPtr("显示在对话框中树视图控件上方的文本"),
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	wapi.SHGetPathFromIDListW(wapi.SHBrowseForFolderW(&bi), &pszPath)
	fmt.Println("pszPath:", pszPath)                           // 用户选择的文件夹完整路径
	fmt.Println("PszDisplayName:", syscall.UTF16ToString(buf)) // 用户选择的文件夹的名称
}

func ExampleGetOpenFileNameW() {
	// 多个过滤器, 打开单个文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("打开文件"),
		Flags:             wapi.OFN_PATHMUTEXIST, // 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetOpenFileNameW(&ofn)
	fmt.Println(ret)
	fmt.Println("lpstrFile:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("lpstrFileTitle:", syscall.UTF16ToString(lpstrFileTitle))
}

func ExampleGetOpenFileNameW_2() {
	// 多个过滤器, 打开多个文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 512)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      2,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          512,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("打开文件(可选多个)"),
		Flags:             wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST, // 允许文件多选 | 使用新界面 | 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetOpenFileNameW(&ofn)
	fmt.Println(ret)

	s := common.Uint16SliceToStringSlice(lpstrFile)
	fmt.Println("选择的文件个数:", len(s)-1) // -1是因为切片中第一个元素是文件目录, 不是文件
	fmt.Println("lpstrFile:", s)
}

func ExampleGetSaveFileNameW() {
	// 多个过滤器, 保存文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("保存文件"),
		Flags:             wapi.OFN_OVERWRITEPROMPT, // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       common.StrPtr("txt"), // 如果用户没有输入文件扩展名, 则默认使用这个
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetSaveFileNameW(&ofn)
	fmt.Println(ret)
	fmt.Println("lpstrFile:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("lpstrFileTitle:", syscall.UTF16ToString(lpstrFileTitle))
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
