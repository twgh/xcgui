package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment12 用于创建共享缓冲区对象.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment12
type ICoreWebView2Environment12 struct {
	ICoreWebView2Environment11
}

// CreateSharedBuffer 创建一个基于共享内存的缓冲区，其大小以字节为单位指定。
//   - 通过在 PostSharedBufferToScript 或 CoreWebView2Frame 对象上调用 PostSharedBufferToScript，可以在 WebView 中与网页内容共享缓冲区。
//   - 一旦共享，应用进程和 WebView 中的脚本都可以访问缓冲区的相同内容。对内容的修改对所有有权访问该缓冲区的各方都可见。
//   - 共享缓冲区以 ArrayBuffer 的形式呈现给脚本。所有适用于 ArrayBuffer 的 JavaScript API（包括 Atomics API）都可以在其上使用。
//   - 目前存在一个限制，即仅支持小于 2GB 的大小。
func (i *ICoreWebView2Environment12) CreateSharedBuffer(size uint64) (*ICoreWebView2SharedBuffer, error) {
	var buffer *ICoreWebView2SharedBuffer
	r, _, _ := i.Vtbl.CreateSharedBuffer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(size),
		uintptr(unsafe.Pointer(&buffer)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return buffer, nil
}

// MustCreateSharedBuffer 创建一个基于共享内存的缓冲区，其大小以字节为单位指定。出错时会触发全局错误回调。
//   - 通过在 PostSharedBufferToScript 或 CoreWebView2Frame 对象上调用 PostSharedBufferToScript，可以在 WebView 中与网页内容共享缓冲区。
//   - 一旦共享，应用进程和 WebView 中的脚本都可以访问缓冲区的相同内容。对内容的修改对所有有权访问该缓冲区的各方都可见。
//   - 共享缓冲区以 ArrayBuffer 的形式呈现给脚本。所有适用于 ArrayBuffer 的 JavaScript API（包括 Atomics API）都可以在其上使用。
//   - 目前存在一个限制，即仅支持小于 2GB 的大小。
func (i *ICoreWebView2Environment12) MustCreateSharedBuffer(size uint64) *ICoreWebView2SharedBuffer {
	buffer, err := i.CreateSharedBuffer(size)
	ReportErrorAuto(err)
	return buffer
}
