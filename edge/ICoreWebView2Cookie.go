package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Cookie 提供一组用于管理 Cookie 的方法。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2cookie
type ICoreWebView2Cookie struct {
	Vtbl *ICoreWebView2CookieVtbl
}

type ICoreWebView2CookieVtbl struct {
	IUnknownVtbl
	GetName       ComProc
	GetValue      ComProc
	GetDomain     ComProc
	GetPath       ComProc
	GetExpires    ComProc
	GetIsHttpOnly ComProc
	GetSameSite   ComProc
	GetIsSecure   ComProc
	GetIsSession  ComProc
	PutValue      ComProc
	PutExpires    ComProc
	PutIsHttpOnly ComProc
	PutSameSite   ComProc
	PutIsSecure   ComProc
}

func (i *ICoreWebView2Cookie) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Cookie) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Cookie) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDomain 获取 Cookie 的域。
//   - 默认值是从中接收此 cookie 的主机。
//   - 请注意，例如: “.bing.com”、“bing.com”和“www.bing.com”被认为是不同的域。
func (i *ICoreWebView2Cookie) GetDomain() (string, error) {
	var _domain *uint16
	r, _, err := i.Vtbl.GetDomain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_domain)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	domain := windows.UTF16PtrToString(_domain)
	windows.CoTaskMemFree(unsafe.Pointer(_domain))
	return domain, nil
}

// MustGetDomain 获取 Cookie 的域。忽略错误.
//   - 默认值是从中接收此 cookie 的主机。
//   - 请注意，例如: “.bing.com”、“bing.com”和“www.bing.com”被认为是不同的域。
func (i *ICoreWebView2Cookie) MustGetDomain() string {
	domain, _ := i.GetDomain()
	return domain
}

// GetName 获取 Cookie 的名称。
func (i *ICoreWebView2Cookie) GetName() (string, error) {
	var _name *uint16
	r, _, err := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_name)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	name := windows.UTF16PtrToString(_name)
	windows.CoTaskMemFree(unsafe.Pointer(_name))
	return name, nil
}

// MustGetName 获取 Cookie 的名称。忽略错误。
func (i *ICoreWebView2Cookie) MustGetName() string {
	name, _ := i.GetName()
	return name
}

// GetValue 获取 Cookie 的值。
func (i *ICoreWebView2Cookie) GetValue() (string, error) {
	var _value *uint16
	r, _, err := i.Vtbl.GetValue.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// MustGetValue 获取 Cookie 的值。忽略错误。
func (i *ICoreWebView2Cookie) MustGetValue() string {
	value, _ := i.GetValue()
	return value
}

// GetPath 获取 Cookie 的路径。
//   - 默认值为“/”，这意味着此 cookie 将发送到域上的所有页面。
func (i *ICoreWebView2Cookie) GetPath() (string, error) {
	var _path *uint16
	r, _, err := i.Vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_path)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := windows.UTF16PtrToString(_path)
	windows.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}

// MustGetPath 获取 Cookie 的路径。忽略错误。
//   - 默认值为“/”，这意味着此 cookie 将发送到域上的所有页面。
func (i *ICoreWebView2Cookie) MustGetPath() string {
	path, _ := i.GetPath()
	return path
}

// GetExpires 获取 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。
//   - 默认值为 -1.0，这意味着默认情况下 Cookie 是会话 Cookie。
func (i *ICoreWebView2Cookie) GetExpires() (float64, error) {
	var expires float64
	r, _, err := i.Vtbl.GetExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return expires, nil
}

// MustGetExpires 获取 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。忽略错误。
//   - 默认值为 -1.0，这意味着默认情况下 Cookie 是会话 Cookie。
func (i *ICoreWebView2Cookie) MustGetExpires() float64 {
	expires, _ := i.GetExpires()
	return expires
}

// GetIsHttpOnly 获取 Cookie 是否为 HttpOnly。
//   - 如果页面脚本或其他活动内容无法访问此 cookie，则为 true。默认值为 false。
func (i *ICoreWebView2Cookie) GetIsHttpOnly() (bool, error) {
	var isHttpOnly bool
	r, _, err := i.Vtbl.GetIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isHttpOnly)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isHttpOnly, nil
}

// MustGetIsHttpOnly 获取 Cookie 是否为 HttpOnly。忽略错误。
//   - 如果页面脚本或其他活动内容无法访问此 cookie，则为 true。默认值为 false。
func (i *ICoreWebView2Cookie) MustGetIsHttpOnly() bool {
	isHttpOnly, _ := i.GetIsHttpOnly()
	return isHttpOnly
}

// GetSameSite 获取 Cookie 的 SameSite 状态，表示cookie的强制模式。
//   - 默认为 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX。
func (i *ICoreWebView2Cookie) GetSameSite() (COREWEBVIEW2_COOKIE_SAME_SITE_KIND, error) {
	var sameSite COREWEBVIEW2_COOKIE_SAME_SITE_KIND
	r, _, err := i.Vtbl.GetSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&sameSite)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return sameSite, nil
}

// MustGetSameSite 获取 Cookie 的 SameSite 状态，表示cookie的强制模式。忽略错误。
//   - 默认为 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX。
func (i *ICoreWebView2Cookie) MustGetSameSite() COREWEBVIEW2_COOKIE_SAME_SITE_KIND {
	sameSite, _ := i.GetSameSite()
	return sameSite
}

// GetIsSecure 获取 Cookie 的安全级别。
//   - 如果客户端仅在后续请求中返回 cookie（如果这些请求使用 HTTPS），则为 true。
//   - 默认值为 false。
//   - 请注意，请求 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE 但未标记为安全的 cookie 将被拒绝。
func (i *ICoreWebView2Cookie) GetIsSecure() (bool, error) {
	var isSecure bool
	r, _, err := i.Vtbl.GetIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSecure)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isSecure, nil
}

// MustGetIsSecure 获取 Cookie 的安全级别。忽略错误。
//   - 如果客户端仅在后续请求中返回 cookie（如果这些请求使用 HTTPS），则为 true。
//   - 默认值为 false。
//   - 请注意，请求 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE 但未标记为安全的 cookie 将被拒绝。
func (i *ICoreWebView2Cookie) MustGetIsSecure() bool {
	isSecure, _ := i.GetIsSecure()
	return isSecure
}

// GetIsSession 获取 Cookie 是否为会话 Cookie, 默认值为false。
func (i *ICoreWebView2Cookie) GetIsSession() (bool, error) {
	var isSession bool
	r, _, err := i.Vtbl.GetIsSession.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSession)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isSession, nil
}

// MustGetIsSession 获取 Cookie 是否为会话 Cookie, 默认值为false。忽略错误。
func (i *ICoreWebView2Cookie) MustGetIsSession() bool {
	isSession, _ := i.GetIsSession()
	return isSession
}

// PutValue 设置 Cookie 的值。
func (i *ICoreWebView2Cookie) PutValue(value string) error {
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutValue.Call(
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

// PutIsHttpOnly 设置 Cookie 是否为 HttpOnly。
func (i *ICoreWebView2Cookie) PutIsHttpOnly(isHttpOnly bool) error {
	r, _, err := i.Vtbl.PutIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(isHttpOnly)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutSameSite 设置 Cookie 的 SameSite。
func (i *ICoreWebView2Cookie) PutSameSite(sameSite COREWEBVIEW2_COOKIE_SAME_SITE_KIND) error {
	r, _, err := i.Vtbl.PutSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(sameSite),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutExpires 设置 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。
//   - Cookie 是会话 Cookie，如果 Expires 设置为 -1.0，则 Cookie 不会持久。
//   - 不允许 NaN、无穷大和除 -1.0 以外的任何负值集。
func (i *ICoreWebView2Cookie) PutExpires(expires float64) error {
	r, _, err := i.Vtbl.PutExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// PutIsSecure 设置 Cookie 的是否安全属性。
func (i *ICoreWebView2Cookie) PutIsSecure(isSecure bool) error {
	r, _, err := i.Vtbl.PutIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(isSecure)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
