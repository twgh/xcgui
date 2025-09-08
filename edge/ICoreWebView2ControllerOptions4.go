package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2ControllerOptions4 是 ICoreWebView2ControllerOptions3 的延续，用于允许用户输入通过浏览器传递，并使其在主机应用程序进程中被接收。
type ICoreWebView2ControllerOptions4 struct {
	ICoreWebView2ControllerOptions3
}

// GetAllowHostInputProcessing 获取是否启用输入在传递到 WebView2 之前通过应用程序。
func (i *ICoreWebView2ControllerOptions4) GetAllowHostInputProcessing() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetAllowHostInputProcessing.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetAllowHostInputProcessing 设置是否启用输入在传递到 WebView2 之前通过应用程序。
//   - 此属性仅适用于通过 ICoreWebView2Environment.CreateCoreWebView2ControllerAsync 创建的控制器，而不适用于通过 ICoreWebView2Environment.CreateCoreWebView2CompositionControllerAsync 创建的合成控制器。
//   - 默认值为 FALSE。
//   - 使用视觉托管时，设置此属性无效。
func (i *ICoreWebView2ControllerOptions4) SetAllowHostInputProcessing(value bool) error {
	r, _, _ := i.Vtbl.PutAllowHostInputProcessing.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetAllowHostInputProcessing 获取是否启用输入在传递到 WebView2 之前通过应用程序。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions4) MustGetAllowHostInputProcessing() bool {
	result, err := i.GetAllowHostInputProcessing()
	ReportErrorAuto(err)
	return result
}
