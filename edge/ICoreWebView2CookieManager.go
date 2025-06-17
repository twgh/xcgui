package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2CookieManager 创建、添加或更新、获取或查看Cookie。
//   - 这些更改将应用于用户配置文件的上下文。也就是说，相同用户配置文件下的其他 WebView 可能会受到影响。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager
type ICoreWebView2CookieManager struct {
	Vtbl *ICoreWebView2CookieManagerVtbl
}

type ICoreWebView2CookieManagerVtbl struct {
	IUnknownVtbl
	CreateCookie                   ComProc
	CopyCookie                     ComProc
	GetCookies                     ComProc
	AddOrUpdateCookie              ComProc
	DeleteCookie                   ComProc
	DeleteCookies                  ComProc
	DeleteCookiesWithDomainAndPath ComProc
	DeleteAllCookies               ComProc
}

func (i *ICoreWebView2CookieManager) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CookieManager) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CookieManager) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCookie 创建一个 ICoreWebView2Cookie 对象，使用指定的参数。
//
// name: cookie 的名称.
//
// value: cookie 的值.
//
// domain: cookie 的域.
//
// path: cookie 的路径.
func (i *ICoreWebView2CookieManager) CreateCookie(name, value, domain, path string) (*ICoreWebView2Cookie, error) {
	var cookie *ICoreWebView2Cookie
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return nil, err
	}
	_domain, err := syscall.UTF16PtrFromString(domain)
	if err != nil {
		return nil, err
	}
	_path, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}

	r, _, err := i.Vtbl.CreateCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
		uintptr(unsafe.Pointer(_domain)),
		uintptr(unsafe.Pointer(_path)),
		uintptr(unsafe.Pointer(&cookie)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cookie, nil
}

// CopyCookie 创建一个 ICoreWebView2Cookie 对象，该对象是指定 cookie 的副本。
func (i *ICoreWebView2CookieManager) CopyCookie(cookieParam *ICoreWebView2Cookie) (*ICoreWebView2Cookie, error) {
	var cookie *ICoreWebView2Cookie
	r, _, err := i.Vtbl.CopyCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookieParam)),
		uintptr(unsafe.Pointer(&cookie)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cookie, nil
}

// AddOrUpdateCookie 使用给定的 cookie 数据添加或更新 cookie；可以用匹配的名称、域和路径覆盖 Cookie（如果存在）。
func (i *ICoreWebView2CookieManager) AddOrUpdateCookie(cookie *ICoreWebView2Cookie) error {
	r, _, err := i.Vtbl.AddOrUpdateCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookie)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// DeleteCookie 删除名称和域/路径对与指定 cookie 的名称和路径对匹配的 cookie。
func (i *ICoreWebView2CookieManager) DeleteCookie(cookie *ICoreWebView2Cookie) error {
	r, _, err := i.Vtbl.DeleteCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookie)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// DeleteAllCookies 删除同一配置文件下的所有 Cookie。
func (i *ICoreWebView2CookieManager) DeleteAllCookies() error {
	r, _, err := i.Vtbl.DeleteAllCookies.Call(
		uintptr(unsafe.Pointer(i)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// DeleteCookiesWithDomainAndPath 删除具有匹配名称和域/路径对的 Cookie。
//   - Cookie名称是必需的。
//   - 如果指定了域，则仅删除具有确切域的 Cookie。
//   - 如果指定了路径，则仅删除具有准确路径的 Cookie。
//
// name: cookie 的名称.
//
// domain: cookie 的域.
//
// path: cookie 的路径.
func (i *ICoreWebView2CookieManager) DeleteCookiesWithDomainAndPath(name, domain, path string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_domain, err := syscall.UTF16PtrFromString(domain)
	if err != nil {
		return err
	}
	_path, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.DeleteCookiesWithDomainAndPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_domain)),
		uintptr(unsafe.Pointer(_path)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// DeleteCookies 删除具有匹配名称和 uri 的 Cookie。
//   - Cookie 名称是必需的。删除域和路径与提供的 URI 匹配的具有给定名称的所有 Cookie。
//
// name: cookie 的名称.
//
// uri: .
func (i *ICoreWebView2CookieManager) DeleteCookies(name, uri string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.DeleteCookies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_uri)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCookies 获取与指定 URI 匹配的所有 Cookie。
//   - 如果 uri 为空字符串，则返回同一配置文件下的所有 Cookie。
//   - 你可以通过调用 ICoreWebView2CookieManager.AddOrUpdateCookie 来修改 Cookie 对象，所做的更改将应用到WebView中。
//
// uri: 要匹配的 URI.
func (i *ICoreWebView2CookieManager) GetCookies(uri string, handler *ICoreWebView2GetCookiesCompletedHandler) error {
	_uri, err := syscall.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.GetCookies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(unsafe.Pointer(handler)),
	)

	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCookiesEx 获取与指定 URI 匹配的所有 Cookie。
//   - 如果 uri 为空字符串，则返回同一配置文件下的所有 Cookie。
//   - 你可以通过调用 ICoreWebView2CookieManager.AddOrUpdateCookie 来修改 Cookie 对象，所做的更改将应用到WebView中。
//
// impl: *WebViewEventImpl.
//
// uri: 要匹配的 URI.
//
// cb: 接收结果的回调函数.
func (i *ICoreWebView2CookieManager) GetCookiesEx(impl *WebViewEventImpl, uri string, cb func(errorCode syscall.Errno, cookies *ICoreWebView2CookieList) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "GetCookiesCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "GetCookiesCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "GetCookiesCompleted")
	}
	return i.GetCookies(uri, (*ICoreWebView2GetCookiesCompletedHandler)(handler))
}

// MustCopyCookie 创建一个 ICoreWebView2Cookie 对象，该对象是指定 cookie 的副本。出错时会触发全局错误回调。
func (i *ICoreWebView2CookieManager) MustCopyCookie(cookieParam *ICoreWebView2Cookie) *ICoreWebView2Cookie {
	cookie, _ := i.CopyCookie(cookieParam)
	return cookie
}
