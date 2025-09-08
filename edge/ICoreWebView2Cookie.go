package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

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
	PutValue      ComProc
	GetDomain     ComProc
	GetPath       ComProc
	GetExpires    ComProc
	PutExpires    ComProc
	GetIsHttpOnly ComProc
	PutIsHttpOnly ComProc
	GetSameSite   ComProc
	PutSameSite   ComProc
	GetIsSecure   ComProc
	PutIsSecure   ComProc
	GetIsSession  ComProc
}

func (i *ICoreWebView2Cookie) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Cookie) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Cookie) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
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
	r, _, _ := i.Vtbl.GetDomain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_domain)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	domain := common.UTF16PtrToString(_domain)
	wapi.CoTaskMemFree(unsafe.Pointer(_domain))
	return domain, nil
}

// GetName 获取 Cookie 的名称。
func (i *ICoreWebView2Cookie) GetName() (string, error) {
	var _name *uint16
	r, _, _ := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_name)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	name := common.UTF16PtrToString(_name)
	wapi.CoTaskMemFree(unsafe.Pointer(_name))
	return name, nil
}

// GetValue 获取 Cookie 的值。
func (i *ICoreWebView2Cookie) GetValue() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetValue.Call(
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

// GetPath 获取 Cookie 的路径。
//   - 默认值为“/”，这意味着此 cookie 将发送到域上的所有页面。
func (i *ICoreWebView2Cookie) GetPath() (string, error) {
	var _path *uint16
	r, _, _ := i.Vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_path)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(_path)
	wapi.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}

// GetExpires 获取 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。
//   - 默认值为 -1.0，这意味着默认情况下 Cookie 是会话 Cookie。
func (i *ICoreWebView2Cookie) GetExpires() (float64, error) {
	var expires float64
	r, _, _ := i.Vtbl.GetExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return expires, nil
}

// GetIsHttpOnly 获取 Cookie 是否为 HttpOnly。
//   - 如果页面脚本或其他活动内容无法访问此 cookie，则为 true。默认值为 false。
func (i *ICoreWebView2Cookie) GetIsHttpOnly() (bool, error) {
	var isHttpOnly bool
	r, _, _ := i.Vtbl.GetIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isHttpOnly)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isHttpOnly, nil
}

// GetSameSite 获取 Cookie 的 SameSite 状态，表示cookie的强制模式。
//   - 默认为 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX。
func (i *ICoreWebView2Cookie) GetSameSite() (COREWEBVIEW2_COOKIE_SAME_SITE_KIND, error) {
	var sameSite COREWEBVIEW2_COOKIE_SAME_SITE_KIND
	r, _, _ := i.Vtbl.GetSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&sameSite)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return sameSite, nil
}

// GetIsSecure 获取 Cookie 的安全级别。
//   - 如果客户端仅在后续请求中返回 cookie（如果这些请求使用 HTTPS），则为 true。
//   - 默认值为 false。
//   - 请注意，请求 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE 但未标记为安全的 cookie 将被拒绝。
func (i *ICoreWebView2Cookie) GetIsSecure() (bool, error) {
	var isSecure bool
	r, _, _ := i.Vtbl.GetIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSecure)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isSecure, nil
}

// GetIsSession 获取 Cookie 是否为会话 Cookie, 默认值为false。
func (i *ICoreWebView2Cookie) GetIsSession() (bool, error) {
	var isSession bool
	r, _, _ := i.Vtbl.GetIsSession.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSession)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isSession, nil
}

// SetValue 设置 Cookie 的值。
func (i *ICoreWebView2Cookie) SetValue(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutValue.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetIsHttpOnly 设置 Cookie 是否为 HttpOnly。
func (i *ICoreWebView2Cookie) SetIsHttpOnly(isHttpOnly bool) error {
	r, _, _ := i.Vtbl.PutIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isHttpOnly),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetSameSite 设置 Cookie 的 SameSite。
func (i *ICoreWebView2Cookie) SetSameSite(sameSite COREWEBVIEW2_COOKIE_SAME_SITE_KIND) error {
	r, _, _ := i.Vtbl.PutSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(sameSite),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetExpires 设置 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。
//   - Cookie 是会话 Cookie，如果 Expires 设置为 -1.0，则 Cookie 不会持久。
//   - 不允许 NaN、无穷大和除 -1.0 以外的任何负值集。
func (i *ICoreWebView2Cookie) SetExpires(expires float64) error {
	r, _, _ := i.Vtbl.PutExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetIsSecure 设置 Cookie 的是否安全属性。
func (i *ICoreWebView2Cookie) SetIsSecure(isSecure bool) error {
	r, _, _ := i.Vtbl.PutIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isSecure),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetDomain 获取 Cookie 的域。出错时会触发全局错误回调.
//   - 默认值是从中接收此 cookie 的主机。
//   - 请注意，例如: “.bing.com”、“bing.com”和“www.bing.com”被认为是不同的域。
func (i *ICoreWebView2Cookie) MustGetDomain() string {
	domain, err := i.GetDomain()
	ReportErrorAuto(err)
	return domain
}

// MustGetName 获取 Cookie 的名称。出错时会触发全局错误回调。
func (i *ICoreWebView2Cookie) MustGetName() string {
	name, err := i.GetName()
	ReportErrorAuto(err)
	return name
}

// MustGetValue 获取 Cookie 的值。出错时会触发全局错误回调。
func (i *ICoreWebView2Cookie) MustGetValue() string {
	value, err := i.GetValue()
	ReportErrorAuto(err)
	return value
}

// MustGetPath 获取 Cookie 的路径。出错时会触发全局错误回调。
//   - 默认值为“/”，这意味着此 cookie 将发送到域上的所有页面。
func (i *ICoreWebView2Cookie) MustGetPath() string {
	path, err := i.GetPath()
	ReportErrorAuto(err)
	return path
}

// MustGetExpires 获取 Cookie 的到期日期和时间，自 UNIX 纪元以来的秒数。出错时会触发全局错误回调。
//   - 默认值为 -1.0，这意味着默认情况下 Cookie 是会话 Cookie。
func (i *ICoreWebView2Cookie) MustGetExpires() float64 {
	expires, err := i.GetExpires()
	ReportErrorAuto(err)
	return expires
}

// MustGetIsHttpOnly 获取 Cookie 是否为 HttpOnly。出错时会触发全局错误回调。
//   - 如果页面脚本或其他活动内容无法访问此 cookie，则为 true。默认值为 false。
func (i *ICoreWebView2Cookie) MustGetIsHttpOnly() bool {
	isHttpOnly, err := i.GetIsHttpOnly()
	ReportErrorAuto(err)
	return isHttpOnly
}

// MustGetSameSite 获取 Cookie 的 SameSite 状态，表示cookie的强制模式。出错时会触发全局错误回调。
//   - 默认为 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX。
func (i *ICoreWebView2Cookie) MustGetSameSite() COREWEBVIEW2_COOKIE_SAME_SITE_KIND {
	sameSite, err := i.GetSameSite()
	ReportErrorAuto(err)
	return sameSite
}

// MustGetIsSecure 获取 Cookie 的安全级别。出错时会触发全局错误回调。
//   - 如果客户端仅在后续请求中返回 cookie（如果这些请求使用 HTTPS），则为 true。
//   - 默认值为 false。
//   - 请注意，请求 COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE 但未标记为安全的 cookie 将被拒绝。
func (i *ICoreWebView2Cookie) MustGetIsSecure() bool {
	isSecure, err := i.GetIsSecure()
	ReportErrorAuto(err)
	return isSecure
}

// MustGetIsSession 获取 Cookie 是否为会话 Cookie, 默认值为false。出错时会触发全局错误回调。
func (i *ICoreWebView2Cookie) MustGetIsSession() bool {
	isSession, err := i.GetIsSession()
	ReportErrorAuto(err)
	return isSession
}
