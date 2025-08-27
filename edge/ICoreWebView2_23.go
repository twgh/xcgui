package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_23 是 ICoreWebView2_22 接口的延续，允许使用附加对象发送 Web 消息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_23
type ICoreWebView2_23 struct {
	ICoreWebView2_22
}

// PostWebMessageAsJsonWithAdditionalObjects 将指定的 webMessage 作为 JSON 字符串发布到 WebView，还支持向页面内容发送 DOM 对象。
//
// webMessageAsJson: 要发布到 WebView 的 JSON 格式的 Web 消息。
//
// additionalObjects: 包含要传递给 WebView 的 DOM 对象的集合。
//   - 目前，这些对象可以是以下类型： ICoreWebView2FileSystemHandle, FileSystemHandle, nullptr, null.
func (i *ICoreWebView2_23) PostWebMessageAsJsonWithAdditionalObjects(webMessageAsJson string, additionalObjects *ICoreWebView2ObjectCollectionView) error {
	_webMessageAsJson, err := syscall.UTF16PtrFromString(webMessageAsJson)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostWebMessageAsJsonWithAdditionalObjects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
		uintptr(unsafe.Pointer(additionalObjects)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
