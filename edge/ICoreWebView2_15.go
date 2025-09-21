package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2_15 是 ICoreWebView2_14 接口的延续，支持网站图标(Favicon)相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_15
type ICoreWebView2_15 struct {
	ICoreWebView2_14
}

// AddFaviconChanged 添加网站图标(Favicon)变更事件的处理程序。
//   - 当网站图标（favicon）的 URL 与之前的 URL 不同时，会触发 FaviconChanged 事件。
//   - 首次导航到新文档时，无论该文档是否在 HTML 中声明了网站图标，只要其图标与之前的图标不同，就会触发 FaviconChanged 事件。
//   - 如果 HTML 中声明了网站图标或通过脚本设置了网站图标，该事件将再次触发。随后可以通过 GetFavicon 和 GetFaviconUri 方法获取网站图标信息。
//
// eventHandler: 事件处理程序。
//
// token: 事件注册令牌。
func (i *ICoreWebView2_15) AddFaviconChanged(eventHandler *ICoreWebView2FaviconChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveFaviconChanged 移除网站图标(Favicon)变更事件的处理程序。
//
// token: 事件注册令牌。
func (i *ICoreWebView2_15) RemoveFaviconChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFaviconUri 获取当前网站图标(Favicon)的URI。
func (i *ICoreWebView2_15) GetFaviconUri() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetFaviconUri.Call(
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

// GetFavicon 异步获取当前网站图标(Favicon)的图像数据。
//
// format: 图像格式，如 PNG 或 JPEG。
//
// completedHandler: 获取完成后的回调处理程序。
func (i *ICoreWebView2_15) GetFavicon(format COREWEBVIEW2_FAVICON_IMAGE_FORMAT, completedHandler *ICoreWebView2GetFaviconCompletedHandler) error {
	r, _, _ := i.Vtbl.GetFavicon.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(format),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFaviconEx 异步获取当前网站图标(Favicon)的图像数据。
//
// impl: *WebViewEventImpl。
//
// format: 图像格式，如 PNG 或 JPEG。
//
// cb: 获取完成后的回调处理程序。
func (i *ICoreWebView2_15) GetFaviconEx(impl *WebViewEventImpl, format COREWEBVIEW2_FAVICON_IMAGE_FORMAT, cb func(errorCode syscall.Errno, favicon []byte) uintptr) error {
	var c interface{}
	if cb == nil {
		c = nil
	} else {
		c = cb
	}
	_, _ = WvEventHandler.AddCallBack(impl, "GetFaviconCompleted", c, nil)
	handler := WvEventHandler.GetHandler(impl, "GetFaviconCompleted")
	return i.GetFavicon(format, (*ICoreWebView2GetFaviconCompletedHandler)(handler))
}

// MustGetFaviconUri 获取当前网站图标(Favicon)的URI。出错时会触发全局错误回调。
func (i *ICoreWebView2_15) MustGetFaviconUri() string {
	value, err := i.GetFaviconUri()
	ReportErrorAuto(err)
	return value
}
