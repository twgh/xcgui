package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2Environment7 是 ICoreWebView2Environment6 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment7
type ICoreWebView2Environment7 struct {
	ICoreWebView2Environment6
}

// GetUserDataFolder 获取从此环境创建的所有 CoreWebView2 正在使用的用户数据文件夹。
//   - 这可能是开发人员在创建环境对象时传入的值，也可能是默认处理时计算得出的值。它始终是一个绝对路径。
func (i *ICoreWebView2Environment7) GetUserDataFolder() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetUserDataFolder.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	dir := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return dir, nil
}

// MustGetUserDataFolder 获取从此环境创建的所有 CoreWebView2 正在使用的用户数据文件夹。
// 出错时会触发全局错误回调。
//   - 这可能是开发人员在创建环境对象时传入的值，也可能是默认处理时计算得出的值。它始终是一个绝对路径。
func (i *ICoreWebView2Environment7) MustGetUserDataFolder() string {
	value, err := i.GetUserDataFolder()
	ReportErrorAuto(err)
	return value
}
