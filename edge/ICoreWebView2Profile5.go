package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile5 是 ICoreWebView2Profile4 的延续，提供了 Cookie 管理器访问功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile5
type ICoreWebView2Profile5 struct {
	ICoreWebView2Profile4
}

// GetCookieManager 获取与此配置文件关联的 Cookie 管理器。
//   - 与该配置文件关联的所有 ICoreWebView2 共享相同的 Cookie 值。
//   - 在此 Cookie 管理器中对 Cookie 所做的更改将应用于与该配置文件关联的所有 ICoreWebView2。
func (i *ICoreWebView2Profile5) GetCookieManager() (*ICoreWebView2CookieManager, error) {
	var value *ICoreWebView2CookieManager
	r, _, _ := i.Vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetCookieManager 获取与此配置文件关联的 Cookie 管理器。出错时会触发全局错误回调。
//   - 与该配置文件关联的所有 ICoreWebView2 共享相同的 Cookie 值。
//   - 在此 Cookie 管理器中对 Cookie 所做的更改将应用于与该配置文件关联的所有 ICoreWebView2。
func (i *ICoreWebView2Profile5) MustGetCookieManager() *ICoreWebView2CookieManager {
	result, err := i.GetCookieManager()
	ReportErrorAuto(err)
	return result
}
