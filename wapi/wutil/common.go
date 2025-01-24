package wutil

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

// GetDropFiles 获取拖放进来的文件.
//
// hDropInfo: 拖放信息句柄.
func GetDropFiles(hDropInfo uintptr) []string {
	var filePath string
	files := make([]string, 0)
	var i uint32
	for {
		length := wapi.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}
		files = append(files, filePath)
		i++ // 索引+1检索下一个文件
	}
	wapi.DragFinish(hDropInfo)
	return files
}

// OpenDir 打开文件夹. 返回选择的文件夹完整路径.
//
// hParent: 父炫彩窗口句柄, 可为0.
func OpenDir(hParent int) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}

	displayNameBuffer := make([]uint16, 260)
	lpszTitle, _ := syscall.UTF16PtrFromString("请选择目录")

	bi := wapi.BrowseInfoW{
		HwndOwner:      hwnd,
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
		return ""
	}

	var pszPath string
	wapi.SHGetPathFromIDListW(pidl, &pszPath)
	return pszPath
}

// OpenFile 打开单个文件. 返回文件完整路径.
//
// hParent: 父炫彩窗口句柄, 可为0.
//
// filters: 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}.
//
// defaultDir: 初始目录, 即默认打开的目录.
func OpenFile(hParent int, filters []string, defaultDir string) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}

	var lpstrFilter, lpstrInitialDir *uint16 = nil, nil
	// 拼接过滤器
	if len(filters) > 1 {
		lpstrFilter = common.StringToUint16Ptr(strings.Join(filters, wapi.NullStr) + wapi.NullStr2)
	}

	lpstrFile := make([]uint16, 260)
	lpstrTitle, _ := syscall.UTF16PtrFromString("打开文件")
	if defaultDir != "" {
		lpstrInitialDir, _ = syscall.UTF16PtrFromString(defaultDir)
	}

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       lpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   lpstrInitialDir,
		LpstrTitle:        lpstrTitle,
		Flags:             wapi.OFN_PATHMUTEXIST, // 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       nil,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !wapi.GetOpenFileNameW(&ofn) {
		return ""
	}
	return syscall.UTF16ToString(lpstrFile)
}

// OpenFiles 打开多个文件. 返回文件完整路径数组. 注意:
//   - 如果用户选择了超级多个文件, 可能会超过缓冲区大小, 一旦超过了缓冲区的大小, 会返回nil, 且 wapi.CommDlgExtendedError() == wapi.FNERR_BUFFERTOOSMALL
//
// hParent: 父炫彩窗口句柄, 可为0.
//
// filters: 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}.
//
// defaultDir: 初始目录, 即默认打开的目录.
func OpenFiles(hParent int, filters []string, defaultDir string) []string {
	opt := OpenFileOption{
		Filters:      filters,
		DefDir:       defaultDir,
		MaxOpenFiles: 10,
		Flags:        wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST,
	}
	if hParent > 0 {
		opt.HwndOwner = xc.XWnd_GetHWND(hParent)
	}
	return OpenFileEx(opt)
}

// OpenFileEx 和 SaveFileEx 通用的参数, 都是选填, 可填可不填.
type OpenFileOption struct {
	// 父窗口句柄, 可为0.
	HwndOwner uintptr
	// 打开的对话框标题.
	Title string
	// 初始目录, 即默认打开的目录.
	DefDir string
	// 默认文件名. 在 SaveFileEx 中有效.
	DefFileName string
	// 默认扩展名, 如果用户没有输入文件扩展名, 则默认使用这个.
	DefExt string
	// 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 不填则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}.
	Filters []string
	// 最多打开的文件数量. 打开多个文件时可用.
	//   - 只是为了分配合适的缓冲区大小, 并不能真的限制用户选择多少个文件, 如果用户的选择超过数量, 只返回 maxOpenFiles 个.
	//   - 如果用户选择了超级多个文件, 可能会超过缓冲区大小, 一旦超过了缓冲区的大小, 会返回nil, 且 wapi.CommDlgExtendedError() == wapi.FNERR_BUFFERTOOSMALL
	MaxOpenFiles uint32

	// 标志: wapi.OFN_ , 可组合.
	//  - 打开单个文件时可不填, 默认为 wapi.OFN_PATHMUTEXIST.
	//  - 打开多个文件时, 可填 wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST
	//  - 保存文件时可不填, 默认为 wapi.OFN_OVERWRITEPROMPT | wapi.OFN_PATHMUTEXIST | wapi.OFN_PATHMUTEXIST
	Flags wapi.OFN_
}

// OpenFileEx 打开单个或多个文件. 返回文件完整路径数组. 注意:
//   - 打开多个文件时, 需要填 OpenFileOption 中的 Flags 字段.
//   - 如果用户选择了超级多个文件, 可能会超过缓冲区大小, 一旦超过了缓冲区的大小, 会返回nil, 且 wapi.CommDlgExtendedError() == wapi.FNERR_BUFFERTOOSMALL
func OpenFileEx(opt OpenFileOption) []string {
	var lpstrFilter, lpstrInitialDir, lpstrDefExt *uint16 = nil, nil, nil
	// 拼接过滤器
	if len(opt.Filters) > 1 {
		lpstrFilter = common.StringToUint16Ptr(strings.Join(opt.Filters, wapi.NullStr) + wapi.NullStr2)
	}

	maxOpenFiles := opt.MaxOpenFiles
	if maxOpenFiles == 0 {
		maxOpenFiles = 1
	}
	nBufSize := 261*maxOpenFiles + 1
	lpstrFile := make([]uint16, nBufSize)

	title := opt.Title
	if title == "" {
		title = "打开文件"
	}
	lpstrTitle, _ := syscall.UTF16PtrFromString(title)

	if opt.DefDir != "" {
		lpstrInitialDir, _ = syscall.UTF16PtrFromString(opt.DefDir)
	}

	if opt.DefExt != "" {
		lpstrDefExt, _ = syscall.UTF16PtrFromString(opt.DefExt)
	}

	flags := opt.Flags
	if flags == 0 {
		flags = wapi.OFN_PATHMUTEXIST
	}

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         opt.HwndOwner,
		HInstance:         0,
		LpstrFilter:       lpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          nBufSize,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   lpstrInitialDir,
		LpstrTitle:        lpstrTitle,
		Flags:             flags,
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       lpstrDefExt,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !wapi.GetOpenFileNameW(&ofn) {
		return nil
	}

	slice := common.Uint16SliceToStringSlice(lpstrFile)
	length := len(slice)
	if length == 0 {
		return nil
	} else if length == 1 { // 只选择了一个文件
		return slice
	}

	// 选择了多个文件, 第一个成员是目录, 后面是文件名
	dir := slice[0]
	var s []string
	for i := 1; i < len(slice) && uint32(i) <= maxOpenFiles; i++ {
		s = append(s, filepath.Join(dir, slice[i])) // 拼接文件路径
	}
	return s
}

// SaveFile 保存文件. 返回文件完整路径.
//
// hParent: 父炫彩窗口句柄, 可为0.
//
// filters: 过滤器数组, 每两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}.
//
// defaultDir: 初始目录, 即默认打开的目录.
//
// defaultFileName: 默认文件名.
func SaveFile(hParent int, filters []string, defaultDir, defaultFileName string) string {
	opt := OpenFileOption{
		Filters:     filters,
		DefDir:      defaultDir,
		DefFileName: defaultFileName,
	}
	if hParent > 0 {
		opt.HwndOwner = xc.XWnd_GetHWND(hParent)
	}
	return SaveFileEx(opt)
}

// SaveFileEx 保存文件. 返回文件完整路径.
func SaveFileEx(opt OpenFileOption) string {
	var lpstrFilter, lpstrInitialDir, lpstrDefExt *uint16 = nil, nil, nil
	// 拼接过滤器
	if len(opt.Filters) > 1 {
		lpstrFilter = common.StringToUint16Ptr(strings.Join(opt.Filters, wapi.NullStr) + wapi.NullStr2)
	}

	fileNameBuffer := make([]uint16, 260)
	if opt.DefFileName != "" {
		fromString, err := syscall.UTF16FromString(opt.DefFileName)
		if err == nil {
			copy(fileNameBuffer, fromString)
		}
	}
	lpstrFile := &fileNameBuffer[0]

	title := opt.Title
	if title == "" {
		title = "保存文件"
	}
	lpstrTitle, _ := syscall.UTF16PtrFromString(title)

	if opt.DefDir != "" {
		lpstrInitialDir, _ = syscall.UTF16PtrFromString(opt.DefDir)
	}
	if opt.DefExt != "" {
		lpstrDefExt, _ = syscall.UTF16PtrFromString(opt.DefExt)
	}

	flags := opt.Flags
	if flags == 0 {
		flags = wapi.OFN_OVERWRITEPROMPT | wapi.OFN_PATHMUTEXIST | wapi.OFN_PATHMUTEXIST // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。| 检测文件路径是否合法
	}

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         opt.HwndOwner,
		HInstance:         0,
		LpstrFilter:       lpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         lpstrFile,
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   lpstrInitialDir,
		LpstrTitle:        lpstrTitle,
		Flags:             flags,
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       lpstrDefExt, // 如果用户没有输入文件扩展名, 则默认使用这个
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !wapi.GetSaveFileNameW(&ofn) {
		return ""
	}
	return common.UintPtrToString(uintptr(unsafe.Pointer(lpstrFile)))
}

// ChooseColor 选择颜色. 返回rgb颜色.
//
// hParent: 父炫彩窗口句柄, 可为0.
func ChooseColor(hParent int) int {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	var lpCustColors [16]uint32
	cc := wapi.ChooseColor{
		LStructSize:    36,
		HwndOwner:      hwnd,
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &lpCustColors[0],
		Flags:          wapi.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	if !wapi.ChooseColorW(&cc) {
		return 0
	}
	return int(cc.RgbResult)
}

// HIcon 从文件加载图标. 返回 HICON 句柄.
//   - 可用于需要 HICON 句柄的函数, 如设置托盘图标.
//   - 如果失败, 可使用 syscall.GetLastError() 获取错误信息.
//   - 当图标句柄不再使用时, 可使用 wapi.DestroyIcon 函数释放.
//
// iconPath: 图标路径.
func HIcon(iconPath string) uintptr {
	return wapi.LoadImageW(0, common.StrPtr(iconPath), wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE)
}
