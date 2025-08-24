package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment10 是 ICoreWebView2Environment9 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment10
type ICoreWebView2Environment10 struct {
	ICoreWebView2Environment9
}

// CreateCoreWebView2ControllerOptions 创建一个 ICoreWebView2ControllerOptions 对象，用于配置 WebView2 控制器的创建。
//   - “options” 是可设置的，其中配置文件名称的默认值为空字符串，“IsInPrivateModeEnabled” 的默认值为 false。
//   - 此外，配置文件名称可以重复使用。
func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerOptions() (*ICoreWebView2ControllerOptions, error) {
	var value *ICoreWebView2ControllerOptions
	r, _, _ := i.Vtbl.CreateCoreWebView2ControllerOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// CreateCoreWebView2ControllerWithOptions 使用指定的选项创建一个新的 WebView 控制器。
//
// parentWindow: 父窗口的句柄。
//
// options: 控制器选项。
//
// handler: 创建完成后的回调处理程序。
func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerWithOptions(parentWindow uintptr, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {
	r, _, _ := i.Vtbl.CreateCoreWebView2ControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		parentWindow,
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCoreWebView2CompositionControllerWithOptions 使用指定的选项创建一个新的 WebView 合成控制器。
//
// parentWindow: 父窗口的句柄。
//
// options: 控制器选项。
//
// handler: 创建完成后的回调处理程序。
func (i *ICoreWebView2Environment10) CreateCoreWebView2CompositionControllerWithOptions(parentWindow uintptr, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) error {
	r, _, _ := i.Vtbl.CreateCoreWebView2CompositionControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		parentWindow,
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustCreateCoreWebView2ControllerOptions 创建一个 ICoreWebView2ControllerOptions 对象，用于配置 WebView2 控制器的创建。出错时会触发全局错误回调。
//   - “options” 是可设置的，其中配置文件名称的默认值为空字符串，“IsInPrivateModeEnabled” 的默认值为 false。
//   - 此外，配置文件名称可以重复使用。
func (i *ICoreWebView2Environment10) MustCreateCoreWebView2ControllerOptions() *ICoreWebView2ControllerOptions {
	value, err := i.CreateCoreWebView2ControllerOptions()
	ReportErrorAtuo(err)
	return value
}
