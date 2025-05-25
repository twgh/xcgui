package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
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

func (i *ICoreWebView2EnvironmentOptions) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAdditionalBrowserArguments 获取创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
func (i *ICoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments() (string, error) {
	var value *uint16
	r, _, err := i.Vtbl.GetAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := windows.UTF16PtrToString(value)
	windows.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// MustGetAdditionalBrowserArguments 获取创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetAdditionalBrowserArguments() string {
	value, err := i.GetAdditionalBrowserArguments()
	ReportError2(err)
	return value
}

// PutAdditionalBrowserArguments 设置创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
//   - 请注意，调用此 API 多次时将替换之前的值，而不是追加到其中。
//   - 如果有多个命令行参数，它们之间应该用空格分隔。
//   - 唯一的例外是，如果为单个命令行参数启用/禁用了多个功能，则这些功能应该用逗号分隔。例如：“--disable-features=feature1,feature2 --some-other-switch --do-something”
func (i *ICoreWebView2EnvironmentOptions) PutAdditionalBrowserArguments(value string) error {
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetLanguage 获取 WebView2 环境的语言。
func (i *ICoreWebView2EnvironmentOptions) GetLanguage() (string, error) {
	var value *uint16
	r, _, err := i.Vtbl.GetLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := windows.UTF16PtrToString(value)
	windows.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// PutLanguage 设置 WebView2 环境的语言。
//   - 它适用于浏览器用户界面，如上下文菜单和对话框。
//   - 它也适用于 WebView 向网站发送的 accept - languages HTTP 头部。
//   - 预期的语言环境值采用 BCP 47 语言标签的格式。更多信息可从 IETF BCP47 中获取。
func (i *ICoreWebView2EnvironmentOptions) PutLanguage(value string) error {
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTargetCompatibleBrowserVersion 获取目标兼容的浏览器版本。
func (i *ICoreWebView2EnvironmentOptions) GetTargetCompatibleBrowserVersion() (string, error) {
	var value *uint16
	r, _, err := i.Vtbl.GetTargetCompatibleBrowserVersion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := windows.UTF16PtrToString(value)
	windows.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// MustGetTargetCompatibleBrowserVersion 获取目标兼容的浏览器版本。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetTargetCompatibleBrowserVersion() string {
	value, err := i.GetTargetCompatibleBrowserVersion()
	ReportError2(err)
	return value
}

// PutTargetCompatibleBrowserVersion 设置目标兼容的浏览器版本。
func (i *ICoreWebView2EnvironmentOptions) PutTargetCompatibleBrowserVersion(value string) error {
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutTargetCompatibleBrowserVersion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowSingleSignOnUsingOSPrimaryAccount 获取是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。
func (i *ICoreWebView2EnvironmentOptions) GetAllowSingleSignOnUsingOSPrimaryAccount() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// MustGetAllowSingleSignOnUsingOSPrimaryAccount 获取是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。出错时会触发全局错误回调。
func (i *ICoreWebView2EnvironmentOptions) MustGetAllowSingleSignOnUsingOSPrimaryAccount() bool {
	value, err := i.GetAllowSingleSignOnUsingOSPrimaryAccount()
	ReportError2(err)
	return value
}

// PutAllowSingleSignOnUsingOSPrimaryAccount 设置是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。
func (i *ICoreWebView2EnvironmentOptions) PutAllowSingleSignOnUsingOSPrimaryAccount(value bool) error {
	r, _, err := i.Vtbl.PutAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
