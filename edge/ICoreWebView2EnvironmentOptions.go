package edge

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

// MustGetAdditionalBrowserArguments 获取创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。忽略错误。
func (i *ICoreWebView2EnvironmentOptions) MustGetAdditionalBrowserArguments() string {
	value, _ := i.GetAdditionalBrowserArguments()
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

// MustGetTargetCompatibleBrowserVersion 获取目标兼容的浏览器版本。忽略错误。
func (i *ICoreWebView2EnvironmentOptions) MustGetTargetCompatibleBrowserVersion() string {
	value, _ := i.GetTargetCompatibleBrowserVersion()
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

// MustGetAllowSingleSignOnUsingOSPrimaryAccount 获取是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。忽略错误。
func (i *ICoreWebView2EnvironmentOptions) MustGetAllowSingleSignOnUsingOSPrimaryAccount() bool {
	value, _ := i.GetAllowSingleSignOnUsingOSPrimaryAccount()
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

func (i *ICoreWebView2EnvironmentOptions2) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetExclusiveUserDataFolderAccess 获取其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。
//
// 100.0.1185.39
func (i *ICoreWebView2EnvironmentOptions2) GetExclusiveUserDataFolderAccess() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetExclusiveUserDataFolderAccess.Call(
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

// MustGetExclusiveUserDataFolderAccess 获取其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。忽略错误。
//
// 100.0.1185.39
func (i *ICoreWebView2EnvironmentOptions2) MustGetExclusiveUserDataFolderAccess() bool {
	value, _ := i.GetExclusiveUserDataFolderAccess()
	return value
}

// PutExclusiveUserDataFolderAccess 设置其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。默认为 false。
//
// 100.0.1185.39
func (i *ICoreWebView2EnvironmentOptions2) PutExclusiveUserDataFolderAccess(value bool) error {
	r, _, err := i.Vtbl.PutExclusiveUserDataFolderAccess.Call(
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

// ICoreWebView2EnvironmentOptions3 提供用于创建 WebView2 环境以管理崩溃报告的附加选项。
//
// 109.0.1518.46
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3
type ICoreWebView2EnvironmentOptions3 struct {
	Vtbl *ICoreWebView2EnvironmentOptions3Vtbl
}

type ICoreWebView2EnvironmentOptions3Vtbl struct {
	IUnknownVtbl
	GetIsCustomCrashReportingEnabled ComProc
	PutIsCustomCrashReportingEnabled ComProc
}

func (i *ICoreWebView2EnvironmentOptions3) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions3) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions3) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) GetIsCustomCrashReportingEnabled() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetIsCustomCrashReportingEnabled.Call(
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

// MustGetIsCustomCrashReportingEnabled 获取 Windows 是否会将崩溃数据发送到 Microsoft 端点。忽略错误。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) MustGetIsCustomCrashReportingEnabled() bool {
	value, _ := i.GetIsCustomCrashReportingEnabled()
	return value
}

// PutIsCustomCrashReportingEnabled 设置 Windows 是否会将崩溃数据发送到 Microsoft 端点。默认为 false。
//
// 109.0.1518.46
func (i *ICoreWebView2EnvironmentOptions3) PutIsCustomCrashReportingEnabled(value bool) error {
	r, _, err := i.Vtbl.PutIsCustomCrashReportingEnabled.Call(
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

// ICoreWebView2EnvironmentOptions4 提供用于创建管理自定义方案注册的 WebView2 环境的附加选项。。
//
// 110.0.1587.40
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4
type ICoreWebView2EnvironmentOptions4 struct {
	Vtbl *ICoreWebView2EnvironmentOptions4Vtbl
}

type ICoreWebView2EnvironmentOptions4Vtbl struct {
	IUnknownVtbl
	GetCustomSchemeRegistrations ComProc
	SetCustomSchemeRegistrations ComProc
}

func (i *ICoreWebView2EnvironmentOptions4) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions4) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions4) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCustomSchemeRegistrations 获取自定义方案注册列表。
//   - 返回的对象指针不再使用时必须调用 ReleaseCustomSchemeRegistrations 统一释放, 以免内存泄漏。
//   - 也可以自行遍历数组, 调用对象的 Release 方法释放。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) GetCustomSchemeRegistrations() ([]*ICoreWebView2CustomSchemeRegistration, error) {
	var count uint32
	var registrations **ICoreWebView2CustomSchemeRegistration
	r, _, err := i.Vtbl.GetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&registrations)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	defer windows.CoTaskMemFree(unsafe.Pointer(registrations))

	result := make([]*ICoreWebView2CustomSchemeRegistration, count)
	slice := unsafe.Slice(registrations, count)
	for i := range slice {
		if slice[i] != nil { // todo 测试是否需要 AddRef.
			slice[i].AddRef() // 增加引用计数，确保对象存活
		}
		result[i] = slice[i]
	}
	return result, nil
}

// MustGetCustomSchemeRegistrations 获取自定义方案注册列表。忽略错误。
//   - 返回的对象指针不再使用时必须调用 ReleaseCustomSchemeRegistrations 统一释放, 以免内存泄漏。
//   - 也可以自行遍历数组, 调用对象的 Release 方法释放。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) MustGetCustomSchemeRegistrations() []*ICoreWebView2CustomSchemeRegistration {
	registrations, _ := i.GetCustomSchemeRegistrations()
	return registrations
}

// ReleaseCustomSchemeRegistrations 释放不再使用的 *ICoreWebView2CustomSchemeRegistration.
func (i *ICoreWebView2EnvironmentOptions4) ReleaseCustomSchemeRegistrations(registrations []*ICoreWebView2CustomSchemeRegistration) {
	for _, reg := range registrations {
		if reg != nil {
			reg.Release() // 释放每个对象的引用
		}
	}
}

// SetCustomSchemeRegistrations 设置自定义方案注册列表。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) SetCustomSchemeRegistrations(registrations []*ICoreWebView2CustomSchemeRegistration) error {
	r, _, err := i.Vtbl.SetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(len(registrations)),
		uintptr(unsafe.Pointer(&registrations[0])),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ICoreWebView2CustomSchemeRegistration 表示向 ICoreWebView2Environment 注册自定义方案。
//
// 110.0.1587.40
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration
type ICoreWebView2CustomSchemeRegistration struct {
	Vtbl *ICoreWebView2CustomSchemeRegistrationVtbl
}

type ICoreWebView2CustomSchemeRegistrationVtbl struct {
	IUnknownVtbl
	GetSchemeName            ComProc
	GetTreatAsSecure         ComProc
	PutTreatAsSecure         ComProc
	GetAllowedOrigins        ComProc
	SetAllowedOrigins        ComProc
	GetHasAuthorityComponent ComProc
	PutHasAuthorityComponent ComProc
}

func (i *ICoreWebView2CustomSchemeRegistration) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CustomSchemeRegistration) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CustomSchemeRegistration) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSchemeName 获取自定义方案名称。
func (i *ICoreWebView2CustomSchemeRegistration) GetSchemeName() (string, error) {
	var value *uint16
	r, _, err := i.Vtbl.GetSchemeName.Call(
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

// MustGetSchemeName 获取自定义方案名称。忽略错误。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetSchemeName() string {
	value, _ := i.GetSchemeName()
	return value
}

// GetTreatAsSecure 获取采用此方案的网站是否会像 HTTPS 网站一样被视为安全上下文。
func (i *ICoreWebView2CustomSchemeRegistration) GetTreatAsSecure() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetTreatAsSecure.Call(
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

// MustGetTreatAsSecure 获取采用此方案的网站是否会像 HTTPS 网站一样被视为安全上下文。忽略错误。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetTreatAsSecure() bool {
	value, _ := i.GetTreatAsSecure()
	return value
}

// PutTreatAsSecure 设置该方案是否将被视为安全上下文。
func (i *ICoreWebView2CustomSchemeRegistration) PutTreatAsSecure(value bool) error {
	r, _, err := i.Vtbl.PutTreatAsSecure.Call(
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

// GetAllowedOrigins 获取允许使用自定义方案（如带有 Origin 标头的 XHR 请求和子资源请求）发出请求的来源列表。
func (i *ICoreWebView2CustomSchemeRegistration) GetAllowedOrigins() ([]string, error) {
	var count uint32
	var origins **uint16
	r, _, err := i.Vtbl.GetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&origins)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	defer windows.CoTaskMemFree(unsafe.Pointer(origins))

	result := make([]string, count)
	slice := unsafe.Slice(origins, count)
	for i := range slice {
		result[i] = windows.UTF16PtrToString(slice[i])
		windows.CoTaskMemFree(unsafe.Pointer(slice[i]))
	}
	return result, nil
}

// MustGetAllowedOrigins 获取允许使用自定义方案（如带有 Origin 标头的 XHR 请求和子资源请求）发出请求的来源列表。忽略错误。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetAllowedOrigins() []string {
	origins, _ := i.GetAllowedOrigins()
	return origins
}

// SetAllowedOrigins 设置被允许使用该方案（协议）的源数组。
func (i *ICoreWebView2CustomSchemeRegistration) SetAllowedOrigins(origins []string) error {
	if len(origins) == 0 {
		return errors.New("origins is empty")
	}
	_origins := make([]*uint16, len(origins))
	for i := range origins {
		_origin, err := windows.UTF16PtrFromString(origins[i])
		if err != nil {
			return err
		}
		_origins[i] = _origin
	}
	r, _, err := i.Vtbl.SetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(len(origins)),
		uintptr(unsafe.Pointer(&_origins[0])),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHasAuthorityComponent 获取此方案是否具有权限组件。
func (i *ICoreWebView2CustomSchemeRegistration) GetHasAuthorityComponent() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetHasAuthorityComponent.Call(
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

// PutHasAuthorityComponent 设置具有此自定义方案的 URI 是否将包含一个授权组件（自定义方案的主机）。
func (i *ICoreWebView2CustomSchemeRegistration) PutHasAuthorityComponent(value bool) error {
	r, _, err := i.Vtbl.PutHasAuthorityComponent.Call(
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

// ICoreWebView2EnvironmentOptions5 提供用于创建 WebView2 环境以管理跟踪防护的附加选项。
//
// 111.0.1661.34
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5
type ICoreWebView2EnvironmentOptions5 struct {
	Vtbl *ICoreWebView2EnvironmentOptions5Vtbl
}

type ICoreWebView2EnvironmentOptions5Vtbl struct {
	IUnknownVtbl
	GetEnableTrackingPrevention ComProc
	PutEnableTrackingPrevention ComProc
}

func (i *ICoreWebView2EnvironmentOptions5) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions5) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions5) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetEnableTrackingPrevention 获取是否启用 WebView2 中的跟踪防护功能。
//
// 111.0.1661.34
func (i *ICoreWebView2EnvironmentOptions5) GetEnableTrackingPrevention() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetEnableTrackingPrevention.Call(
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

// PutEnableTrackingPrevention 设置是否启用 WebView2 中的跟踪防护功能。默认为 true。
//   - 此属性可启用/禁用在同一环境中创建的所有 WebView2 的跟踪防护功能。默认情况下，该功能处于启用状态，用于阻止潜在有害的跟踪器以及来自从未访问过的网站的跟踪器，并设置为 COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED 或配置文件上次更改/保存的任何值。
//   - 如果应用程序仅在 WebView2 中渲染已知安全的内容，则可以将此属性设置为 false 以禁用跟踪防护功能。在创建环境时禁用此功能还可以通过跳过相关代码来提高运行时性能。
//   - 如果将 WebView2 用作具有任意导航功能的“完整浏览器”，则不应禁用此属性，并且应保护最终用户的隐私。
//
// 111.0.1661.34
func (i *ICoreWebView2EnvironmentOptions5) PutEnableTrackingPrevention(value bool) error {
	r, _, err := i.Vtbl.PutEnableTrackingPrevention.Call(
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
