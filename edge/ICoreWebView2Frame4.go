package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Frame4 是 ICoreWebView2Frame3 的延续，支持基于文件映射的共享缓冲区功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame4
type ICoreWebView2Frame4 struct {
	ICoreWebView2Frame3
}

// PostSharedBufferToScript 与 WebView 中 iframe 的脚本共享一个共享缓冲区对象。
//   - 该脚本将从 chrome.webview 接收一个 sharedbufferreceived 事件。
//   - 该事件的事件参数将具有以下方法和属性：
//   - getBuffer()：返回一个 ArrayBuffer 对象，该对象包含共享缓冲区的支持内容。
//   - additionalData：一个对象，是将 additionalDataAsJson 作为 JSON 字符串解析的结果。如果 additionalDataAsJson 为 nullptr 或空字符串，此属性将为 undefined。
//   - source：其值设置为 chrome.webview 对象。如果提供的字符串作为 additionalDataAsJson 但不是有效的 JSON 字符串，API 将失败并返回 wapi.E_INVALIDARG。如果 access 为 COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY，脚本将仅对缓冲区具有读取权限。如果脚本尝试修改只读缓冲区中的内容，将导致 WebView 渲染器进程中出现访问冲突，并使渲染器进程崩溃。如果共享缓冲区已关闭，API 将失败并返回 wapi.RO_E_CLOSED。
//   - 脚本代码在不再需要访问共享缓冲区时，应调用 chrome.webview.releaseBuffer 并将共享缓冲区作为参数，以释放底层资源。
//   - 该应用程序可以将同一个共享缓冲区对象发送到多个网页或 iframe，也可以多次发送到同一个网页或 iframe。每次调用 PostSharedBufferToScript 都会创建一个独立的 ArrayBuffer 对象，该对象拥有自己的内存视图，并且会被单独释放。当所有视图都被释放后，底层的共享内存才会被释放。
func (i *ICoreWebView2Frame4) PostSharedBufferToScript(sharedBuffer *ICoreWebView2SharedBuffer, access COREWEBVIEW2_SHARED_BUFFER_ACCESS, additionalDataAsJson string) error {
	_additionalDataAsJson, err := syscall.UTF16PtrFromString(additionalDataAsJson)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PostSharedBufferToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(sharedBuffer)),
		uintptr(access),
		uintptr(unsafe.Pointer(_additionalDataAsJson)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
