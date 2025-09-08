package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2Environment11 提供了获取 WebView2 环境失败报告文件夹路径的方法。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment11
type ICoreWebView2Environment11 struct {
	ICoreWebView2Environment10
}

// GetFailureReportFolderPath 获取 WebView2 环境的失败报告文件夹路径。
//   - 每当 WebView2 进程崩溃时，都会在崩溃转储文件夹中创建一个崩溃转储文件。
//   - 崩溃转储格式为小型转储文件。有关详细信息，请参阅 小型转储文件文档。
//   - 通常，当单个子进程出现故障时，会生成一个小型转储并写入磁盘，然后引发 ProcessFailed 事件。
//   - 但对于意外崩溃，可能根本不会生成小型转储文件，无论是否引发 ProcessFailed 事件。
//   - 如果同时发生多个进程故障，则可能会生成多个小型转储文件。因此，FailureReportFolderPath 可能包含与特定 ProcessFailed 事件无关的旧小型转储文件。
func (i *ICoreWebView2Environment11) GetFailureReportFolderPath() (string, error) {
	var _path *uint16
	r, _, _ := i.Vtbl.GetFailureReportFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_path)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(_path)
	wapi.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}

// MustGetFailureReportFolderPath 获取 WebView2 环境的失败报告文件夹路径。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2Environment11) MustGetFailureReportFolderPath() string {
	path, err := i.GetFailureReportFolderPath()
	ReportErrorAuto(err)
	return path
}
