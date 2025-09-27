package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions 提供用于创建 WebView2 环境的配置选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions
type ICoreWebView2EnvironmentOptions struct {
	Vtbl *ICoreWebView2EnvironmentOptionsVtbl
}

type ICoreWebView2EnvironmentOptionsVtbl struct {
	IUnknownVtbl
	GetAdditionalBrowserArguments             ComProc
	PutAdditionalBrowserArguments             ComProc
	GetLanguage                               ComProc
	PutLanguage                               ComProc
	GetTargetCompatibleBrowserVersion         ComProc
	PutTargetCompatibleBrowserVersion         ComProc
	GetAllowSingleSignOnUsingOSPrimaryAccount ComProc
	PutAllowSingleSignOnUsingOSPrimaryAccount ComProc
}

func (i *ICoreWebView2EnvironmentOptions) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAdditionalBrowserArguments 获取创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
func (i *ICoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// SetAdditionalBrowserArguments 设置创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
//   - 请注意，调用此 API 多次时将替换之前的值，而不是追加到其中。
//   - 如果有多个命令行参数，它们之间应该用空格分隔。
//   - 唯一的例外是，如果为单个命令行参数启用/禁用了多个功能，则这些功能应该用逗号分隔。例如：“--disable-features=feature1,feature2 --some-other-switch --do-something”
func (i *ICoreWebView2EnvironmentOptions) SetAdditionalBrowserArguments(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetLanguage 获取 WebView2 环境的语言。
func (i *ICoreWebView2EnvironmentOptions) GetLanguage() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// SetLanguage 设置 WebView2 环境的语言。
//   - 它适用于浏览器用户界面，如上下文菜单和对话框。
//   - 它也适用于 WebView 向网站发送的 accept - languages HTTP 头部。
//   - 预期的语言环境值采用 BCP 47 语言标签的格式。更多信息可从 IETF BCP47 中获取。
func (i *ICoreWebView2EnvironmentOptions) SetLanguage(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTargetCompatibleBrowserVersion 获取目标兼容的浏览器版本。
func (i *ICoreWebView2EnvironmentOptions) GetTargetCompatibleBrowserVersion() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetTargetCompatibleBrowserVersion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// SetTargetCompatibleBrowserVersion 设置目标兼容的浏览器版本。
func (i *ICoreWebView2EnvironmentOptions) SetTargetCompatibleBrowserVersion(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutTargetCompatibleBrowserVersion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowSingleSignOnUsingOSPrimaryAccount 获取是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。
func (i *ICoreWebView2EnvironmentOptions) GetAllowSingleSignOnUsingOSPrimaryAccount() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetAllowSingleSignOnUsingOSPrimaryAccount 设置是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。
func (i *ICoreWebView2EnvironmentOptions) SetAllowSingleSignOnUsingOSPrimaryAccount(value bool) error {
	r, _, _ := i.Vtbl.PutAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2EnvironmentOptions2 获取 ICoreWebView2EnvironmentOptions2。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions2() (*ICoreWebView2EnvironmentOptions2, error) {
	var result *ICoreWebView2EnvironmentOptions2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions2)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions3 获取 ICoreWebView2EnvironmentOptions3。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions3() (*ICoreWebView2EnvironmentOptions3, error) {
	var result *ICoreWebView2EnvironmentOptions3
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions3)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions4 获取 ICoreWebView2EnvironmentOptions4。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions4() (*ICoreWebView2EnvironmentOptions4, error) {
	var result *ICoreWebView2EnvironmentOptions4
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions4)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions5 获取 ICoreWebView2EnvironmentOptions5。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions5() (*ICoreWebView2EnvironmentOptions5, error) {
	var result *ICoreWebView2EnvironmentOptions5
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions5)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions6 获取 ICoreWebView2EnvironmentOptions6。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions6() (*ICoreWebView2EnvironmentOptions6, error) {
	var result *ICoreWebView2EnvironmentOptions6
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions6)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions7 获取 ICoreWebView2EnvironmentOptions7。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions7() (*ICoreWebView2EnvironmentOptions7, error) {
	var result *ICoreWebView2EnvironmentOptions7
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions7)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2EnvironmentOptions8 获取 ICoreWebView2EnvironmentOptions8。
func (i *ICoreWebView2EnvironmentOptions) GetICoreWebView2EnvironmentOptions8() (*ICoreWebView2EnvironmentOptions8, error) {
	var result *ICoreWebView2EnvironmentOptions8
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2EnvironmentOptions8)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetLanguage 获取 WebView2 环境的语言。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetLanguage() string {
	value, err := i.GetLanguage()
	ReportErrorAuto(err)
	return value
}

// MustGetAdditionalBrowserArguments 获取创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetAdditionalBrowserArguments() string {
	value, err := i.GetAdditionalBrowserArguments()
	ReportErrorAuto(err)
	return value
}

// MustGetTargetCompatibleBrowserVersion 获取目标兼容的浏览器版本。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetTargetCompatibleBrowserVersion() string {
	value, err := i.GetTargetCompatibleBrowserVersion()
	ReportErrorAuto(err)
	return value
}

// MustGetAllowSingleSignOnUsingOSPrimaryAccount 获取是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetAllowSingleSignOnUsingOSPrimaryAccount() bool {
	value, err := i.GetAllowSingleSignOnUsingOSPrimaryAccount()
	ReportErrorAuto(err)
	return value
}

// MustGetICoreWebView2EnvironmentOptions2 获取 ICoreWebView2EnvironmentOptions2。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions2() *ICoreWebView2EnvironmentOptions2 {
	result, err := i.GetICoreWebView2EnvironmentOptions2()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions3 获取 ICoreWebView2EnvironmentOptions3。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions3() *ICoreWebView2EnvironmentOptions3 {
	result, err := i.GetICoreWebView2EnvironmentOptions3()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions4 获取 ICoreWebView2EnvironmentOptions4。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions4() *ICoreWebView2EnvironmentOptions4 {
	result, err := i.GetICoreWebView2EnvironmentOptions4()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions5 获取 ICoreWebView2EnvironmentOptions5。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions5() *ICoreWebView2EnvironmentOptions5 {
	result, err := i.GetICoreWebView2EnvironmentOptions5()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions6 获取 ICoreWebView2EnvironmentOptions6。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions6() *ICoreWebView2EnvironmentOptions6 {
	result, err := i.GetICoreWebView2EnvironmentOptions6()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions7 获取 ICoreWebView2EnvironmentOptions7。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions7() *ICoreWebView2EnvironmentOptions7 {
	result, err := i.GetICoreWebView2EnvironmentOptions7()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2EnvironmentOptions8 获取 ICoreWebView2EnvironmentOptions8。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetICoreWebView2EnvironmentOptions8() *ICoreWebView2EnvironmentOptions8 {
	result, err := i.GetICoreWebView2EnvironmentOptions8()
	ReportErrorAuto(err)
	return result
}
