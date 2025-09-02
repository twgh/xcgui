package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2WebMessageReceivedEventArgs2 是 ICoreWebView2WebMessageReceivedEventArgs 的延续, 提供对额外 WebMessage 对象的访问.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2
type ICoreWebView2WebMessageReceivedEventArgs2 struct {
	ICoreWebView2WebMessageReceivedEventArgs
}

// GetAdditionalObjects 获取包含随 Web 消息一起发送的附加对象的对象集合视图。
//   - 要通过 WebMessage 向应用程序传递 additionalObjects，请使用 chrome.webview.postMessageWithAdditionalObjects 内容API。
//   - 传入 additionalObjects 参数的、任何可原生表示的 DOM 对象类型在此处都可访问。
//   - 目前，WebMessage 对象可以是 ICoreWebView2File 类型。
//   - 如果传入了 null 或 undefined，集合中的条目可能为 nullptr。
func (i *ICoreWebView2WebMessageReceivedEventArgs2) GetAdditionalObjects() (*ICoreWebView2ObjectCollectionView, error) {
	var value *ICoreWebView2ObjectCollectionView
	r, _, _ := i.Vtbl.GetAdditionalObjects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetAdditionalObjects 获取包含随 Web 消息一起发送的附加对象的对象集合视图。出错时会触发全局错误回调。
//   - 要通过 WebMessage 向应用程序传递 additionalObjects，请使用 chrome.webview.postMessageWithAdditionalObjects 内容API。
//   - 传入 additionalObjects 参数的、任何可原生表示的 DOM 对象类型在此处都可访问。
//   - 目前，WebMessage 对象可以是 ICoreWebView2File 类型。
//   - 如果传入了 null 或 undefined，集合中的条目可能为 nullptr。
func (i *ICoreWebView2WebMessageReceivedEventArgs2) MustGetAdditionalObjects() *ICoreWebView2ObjectCollectionView {
	value, err := i.GetAdditionalObjects()
	ReportErrorAtuo(err)
	return value
}
