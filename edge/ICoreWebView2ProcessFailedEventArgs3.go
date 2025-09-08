package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ProcessFailedEventArgs3 是 ICoreWebView2ProcessFailedEventArgs2 的延续，用于获取代码完整性进程失败时被阻止的文件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs3
type ICoreWebView2ProcessFailedEventArgs3 struct {
	ICoreWebView2ProcessFailedEventArgs2
}

// GetFailureSourceModulePath 获取导致失败的模块路径。
func (i *ICoreWebView2ProcessFailedEventArgs3) GetFailureSourceModulePath() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetFailureSourceModulePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// MustGetFailureSourceModulePath 获取导致失败的模块路径。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs3) MustGetFailureSourceModulePath() string {
	value, err := i.GetFailureSourceModulePath()
	ReportErrorAuto(err)
	return value
}
