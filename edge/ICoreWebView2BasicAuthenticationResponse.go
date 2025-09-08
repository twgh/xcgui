package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2BasicAuthenticationResponse 表示对基本 HTTP 身份验证请求的响应，该响应包含根据 RFC7617（https://tools.ietf.org/html/rfc7617）规定的用户名和密码。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse
type ICoreWebView2BasicAuthenticationResponse struct {
	Vtbl *_ICoreWebView2BasicAuthenticationResponseVtbl
}

type _ICoreWebView2BasicAuthenticationResponseVtbl struct {
	IUnknownVtbl
	GetUserName ComProc
	PutUserName ComProc
	GetPassword ComProc
	PutPassword ComProc
}

func (i *ICoreWebView2BasicAuthenticationResponse) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BasicAuthenticationResponse) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BasicAuthenticationResponse) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUserName 获取用于基本身份验证的用户名。
func (i *ICoreWebView2BasicAuthenticationResponse) GetUserName() (string, error) {
	var _userName *uint16
	r, _, _ := i.Vtbl.GetUserName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_userName)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	userName := common.UTF16PtrToString(_userName)
	wapi.CoTaskMemFree(unsafe.Pointer(_userName))
	return userName, nil
}

// SetUserName 设置用于基本身份验证的用户名。
func (i *ICoreWebView2BasicAuthenticationResponse) SetUserName(userName string) error {
	_userName, err := syscall.UTF16PtrFromString(userName)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutUserName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userName)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPassword 获取用于基本身份验证的密码。
func (i *ICoreWebView2BasicAuthenticationResponse) GetPassword() (string, error) {
	var _password *uint16
	r, _, _ := i.Vtbl.GetPassword.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_password)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	password := common.UTF16PtrToString(_password)
	wapi.CoTaskMemFree(unsafe.Pointer(_password))
	return password, nil
}

// SetPassword 设置用于基本身份验证的密码。
func (i *ICoreWebView2BasicAuthenticationResponse) SetPassword(password string) error {
	_password, err := syscall.UTF16PtrFromString(password)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutPassword.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_password)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetUserName 获取用于基本身份验证的用户名。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationResponse) MustGetUserName() string {
	result, err := i.GetUserName()
	ReportErrorAuto(err)
	return result
}

// MustGetPassword 获取用于基本身份验证的密码。出错时会触发全局错误回调。
func (i *ICoreWebView2BasicAuthenticationResponse) MustGetPassword() string {
	result, err := i.GetPassword()
	ReportErrorAuto(err)
	return result
}
