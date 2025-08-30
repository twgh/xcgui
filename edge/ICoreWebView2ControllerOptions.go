package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ControllerOptions 提供用于创建 WebView2 控制器的选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions
type ICoreWebView2ControllerOptions struct {
	Vtbl *ICoreWebView2ControllerOptionsVtbl
}

type ICoreWebView2ControllerOptionsVtbl struct {
	IUnknownVtbl
	GetProfileName            ComProc
	PutProfileName            ComProc
	GetIsInPrivateModeEnabled ComProc
	PutIsInPrivateModeEnabled ComProc
	// 2
	GetScriptLocale ComProc
	PutScriptLocale ComProc
	// 3
	GetDefaultBackgroundColor ComProc
	PutDefaultBackgroundColor ComProc
	// 4
	GetAllowHostInputProcessing ComProc
	PutAllowHostInputProcessing ComProc
}

func (i *ICoreWebView2ControllerOptions) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ControllerOptions) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ControllerOptions) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProfileName 获取用于 WebView2 控制器的配置文件名称。
func (i *ICoreWebView2ControllerOptions) GetProfileName() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetProfileName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	profileName := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return profileName, nil
}

// SetProfileName 设置用于 WebView2 控制器的配置文件名称。
//   - 它的最大长度为64个字符（不包括空字符终止符）。它不区分ASCII大小写。
//   - 注意：文本不能以句号“.”或空格“ ”结尾。
//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
func (i *ICoreWebView2ControllerOptions) SetProfileName(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutProfileName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsInPrivateModeEnabled 获取是否启用隐私模式。
func (i *ICoreWebView2ControllerOptions) GetIsInPrivateModeEnabled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetIsInPrivateModeEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetIsInPrivateModeEnabled 设置是否启用隐私模式。
func (i *ICoreWebView2ControllerOptions) SetIsInPrivateModeEnabled(value bool) error {
	r, _, _ := i.Vtbl.PutIsInPrivateModeEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2ControllerOptions2 获取 ICoreWebView2ControllerOptions2 接口。
func (i *ICoreWebView2ControllerOptions) GetICoreWebView2ControllerOptions2() (*ICoreWebView2ControllerOptions2, error) {
	var result *ICoreWebView2ControllerOptions2
	iidICoreWebView2ControllerOptions2 := wapi.NewGUID(IID_ICoreWebView2ControllerOptions2)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2ControllerOptions2),
		unsafe.Pointer(&result))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetICoreWebView2ControllerOptions3 获取 ICoreWebView2ControllerOptions3 接口。
func (i *ICoreWebView2ControllerOptions) GetICoreWebView2ControllerOptions3() (*ICoreWebView2ControllerOptions3, error) {
	var result *ICoreWebView2ControllerOptions3
	iidICoreWebView2ControllerOptions3 := wapi.NewGUID(IID_ICoreWebView2ControllerOptions3)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2ControllerOptions3),
		unsafe.Pointer(&result))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetICoreWebView2ControllerOptions4 获取 ICoreWebView2ControllerOptions4 接口。
func (i *ICoreWebView2ControllerOptions) GetICoreWebView2ControllerOptions4() (*ICoreWebView2ControllerOptions4, error) {
	var result *ICoreWebView2ControllerOptions4
	iidICoreWebView2ControllerOptions4 := wapi.NewGUID(IID_ICoreWebView2ControllerOptions4)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2ControllerOptions4),
		unsafe.Pointer(&result))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MustGetProfileName 获取用于 WebView2 控制器的配置文件名称。出错时会触发全局错误回调。
//   - 注意：文本不能以句号“.”或空格“ ”结尾。
//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
func (i *ICoreWebView2ControllerOptions) MustGetProfileName() string {
	value, err := i.GetProfileName()
	ReportErrorAtuo(err)
	return value
}

// MustGetIsInPrivateModeEnabled 获取是否启用隐私模式。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions) MustGetIsInPrivateModeEnabled() bool {
	value, err := i.GetIsInPrivateModeEnabled()
	ReportErrorAtuo(err)
	return value
}

// MustGetICoreWebView2ControllerOptions2 获取 ICoreWebView2ControllerOptions2 接口。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions) MustGetICoreWebView2ControllerOptions2() *ICoreWebView2ControllerOptions2 {
	result, err := i.GetICoreWebView2ControllerOptions2()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2ControllerOptions3 获取 ICoreWebView2ControllerOptions3 接口。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions) MustGetICoreWebView2ControllerOptions3() *ICoreWebView2ControllerOptions3 {
	result, err := i.GetICoreWebView2ControllerOptions3()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2ControllerOptions4 获取 ICoreWebView2ControllerOptions4 接口。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions) MustGetICoreWebView2ControllerOptions4() *ICoreWebView2ControllerOptions4 {
	result, err := i.GetICoreWebView2ControllerOptions4()
	ReportErrorAtuo(err)
	return result
}
