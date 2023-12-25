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
// hDropInfo 拖放信息句柄.
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

// OpenDir 打开文件夹.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@return string 返回选择的文件夹完整路径.
func OpenDir(hParent int) string {
	buf := make([]uint16, 260)
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	bi := wapi.BrowseInfoW{
		HwndOwner:      hwnd,
		PidlRoot:       0,
		PszDisplayName: common.Uint16SliceDataPtr(&buf),
		LpszTitle:      common.StrPtr("请选择目录"),
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	wapi.SHGetPathFromIDListW(wapi.SHBrowseForFolderW(&bi), &pszPath)
	return pszPath
}

// OpenFile 打开单个文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@return string 返回文件完整路径.
func OpenFile(hParent int, filters []string, defaultDir string) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var LpstrFilter *uint16 = nil
	if len(filters) > 0 {
		LpstrFilter = common.StringToUint16Ptr(strings.Join(filters, wapi.NULL) + wapi.NULL2)
	}

	lpstrFile := make([]uint16, 260)
	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       LpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   common.StrPtr(defaultDir),
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
	if !wapi.GetOpenFileNameW(&ofn) {
		return ""
	}
	return syscall.UTF16ToString(lpstrFile)
}

// OpenFiles 打开多个文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@return string 返回文件完整路径数组.
func OpenFiles(hParent int, filters []string, defaultDir string) []string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var LpstrFilter *uint16 = nil
	if len(filters) > 0 {
		LpstrFilter = common.StringToUint16Ptr(strings.Join(filters, wapi.NULL) + wapi.NULL2)
	}

	lpstrFile := make([]uint16, 512)
	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       LpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          512,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   common.StrPtr(defaultDir),
		LpstrTitle:        common.StrPtr("打开文件"),
		Flags:             wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST, // 允许文件多选 | 使用新界面 | 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !wapi.GetOpenFileNameW(&ofn) {
		return nil
	}

	slice := common.Uint16SliceToStringSlice(lpstrFile)
	if len(slice) < 2 {
		return nil
	}

	dir := slice[0]
	var s []string
	for i := 1; i < len(slice); i++ {
		s = append(s, filepath.Join(dir, slice[i]))
	}
	return s
}

// SaveFile 保存文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@param defaultFileName 默认文件名.
//	@return string 返回文件完整路径.
func SaveFile(hParent int, filters []string, defaultDir, defaultFileName string) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var lpstrFilter *uint16 = nil
	if len(filters) > 0 {
		lpstrFilter = common.StringToUint16Ptr(strings.Join(filters, wapi.NULL) + wapi.NULL2)
	}

	var lpstrFile *uint16 = nil
	if defaultFileName != "" {
		lpstrFile = common.StringToUint16Ptr(strings.ReplaceAll(defaultFileName, " ", ""))
	} else {
		lpstrFile = &make([]uint16, 260)[0]
	}
	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       lpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         lpstrFile,
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   common.StrPtr(defaultDir),
		LpstrTitle:        common.StrPtr("保存文件"),
		Flags:             wapi.OFN_OVERWRITEPROMPT | wapi.OFN_PATHMUTEXIST | wapi.OFN_PATHMUTEXIST, // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。| 检测文件路径是否合法
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0, // 如果用户没有输入文件扩展名, 则默认使用这个
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

// ChooseColor 选择颜色.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@return int 返回rgb颜色.
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
