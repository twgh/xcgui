package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2CookieList 是 ICoreWebView2Cookie 的集合.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2cookielist
type ICoreWebView2CookieList struct {
	Vtbl *ICoreWebView2CookieListVtbl
}

type ICoreWebView2CookieListVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2CookieList) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CookieList) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CookieList) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取 Cookie 列表中的 Cookie 数量.
func (i *ICoreWebView2CookieList) GetCount() (uint32, error) {
	var count uint32
	r, _, _ := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)

	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return count, nil
}

// GetValueAtIndex 根据索引获取 Cookie 列表中的指定 Cookie.
func (i *ICoreWebView2CookieList) GetValueAtIndex(index uint32) (*ICoreWebView2Cookie, error) {
	var cookie *ICoreWebView2Cookie
	r, _, _ := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&cookie)),
	)

	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cookie, nil
}

// MustGetCount 获取 Cookie 列表中的 Cookie 数量. 出错时会触发全局错误回调.
func (i *ICoreWebView2CookieList) MustGetCount() uint32 {
	count, err := i.GetCount()
	ReportErrorAuto(err)
	return count
}

// MustGetValueAtIndex 根据索引获取 Cookie 列表中的指定 Cookie. 出错时会触发全局错误回调.
func (i *ICoreWebView2CookieList) MustGetValueAtIndex(index uint32) *ICoreWebView2Cookie {
	cookie, err := i.GetValueAtIndex(index)
	ReportErrorAuto(err)
	return cookie
}
