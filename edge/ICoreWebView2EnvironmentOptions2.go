package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2EnvironmentOptions2 提供用于创建 WebView2 环境的额外配置选项。
//
// 100.0.1185.39
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2
type ICoreWebView2EnvironmentOptions2 struct {
	Vtbl *ICoreWebView2EnvironmentOptions2Vtbl
}

type ICoreWebView2EnvironmentOptions2Vtbl struct {
	IUnknownVtbl
	GetExclusiveUserDataFolderAccess ComProc
	PutExclusiveUserDataFolderAccess ComProc
}

func (i *ICoreWebView2EnvironmentOptions2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions2) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetExclusiveUserDataFolderAccess 获取其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。
func (i *ICoreWebView2EnvironmentOptions2) GetExclusiveUserDataFolderAccess() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetExclusiveUserDataFolderAccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetExclusiveUserDataFolderAccess 设置其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。默认为 false。
func (i *ICoreWebView2EnvironmentOptions2) SetExclusiveUserDataFolderAccess(value bool) error {
	r, _, _ := i.Vtbl.PutExclusiveUserDataFolderAccess.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetExclusiveUserDataFolderAccess 获取其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions2) MustGetExclusiveUserDataFolderAccess() bool {
	value, err := i.GetExclusiveUserDataFolderAccess()
	ReportErrorAuto(err)
	return value
}
