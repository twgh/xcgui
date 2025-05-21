package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings2 是 ICoreWebView2Settings 接口的延续, 支持管理 UserAgent.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings2
type ICoreWebView2Settings2 struct {
	Vtbl *ICoreWebView2Settings2Vtbl
}

type ICoreWebView2Settings2Vtbl struct {
	ICoreWebView2SettingsVtbl
	GetUserAgent ComProc
	PutUserAgent ComProc
}

func (i *ICoreWebView2Settings2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings2) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUserAgent 获取 UserAgent。
func (i *ICoreWebView2Settings2) GetUserAgent() (string, error) {
	var _userAgent *uint16
	r, _, err := i.Vtbl.GetUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_userAgent)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	userAgent := windows.UTF16PtrToString(_userAgent)
	windows.CoTaskMemFree(unsafe.Pointer(_userAgent))
	return userAgent, nil
}

// MustGetUserAgent 获取 UserAgent。忽略错误。
func (i *ICoreWebView2Settings2) MustGetUserAgent() string {
	userAgent, _ := i.GetUserAgent()
	return userAgent
}

// PutUserAgent 设置 UserAgent。
func (i *ICoreWebView2Settings2) PutUserAgent(userAgent string) error {
	_userAgent, err := windows.UTF16PtrFromString(userAgent)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.PutUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userAgent)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
