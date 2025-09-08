package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2SharedBuffer 表示由 ICoreWebView2Environment12.CreateSharedBuffer 创建的共享缓冲区对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer
type ICoreWebView2SharedBuffer struct {
	Vtbl *ICoreWebView2SharedBufferVtbl
}

type ICoreWebView2SharedBufferVtbl struct {
	IUnknownVtbl
	GetSize              ComProc
	GetBuffer            ComProc
	OpenStream           ComProc
	GetFileMappingHandle ComProc
	Close                ComProc
}

func (i *ICoreWebView2SharedBuffer) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SharedBuffer) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2SharedBuffer) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSize 获取共享缓冲区的大小（以字节为单位）.
func (i *ICoreWebView2SharedBuffer) GetSize() (uint64, error) {
	var size uint64
	r, _, _ := i.Vtbl.GetSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&size)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return size, nil
}

// GetBuffer 获取共享缓冲区的内存地址.
func (i *ICoreWebView2SharedBuffer) GetBuffer() (unsafe.Pointer, error) {
	var buffer unsafe.Pointer
	r, _, _ := i.Vtbl.GetBuffer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&buffer)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return buffer, nil
}

// OpenStream 获取一个可用于访问共享缓冲区的 IStream 对象。
func (i *ICoreWebView2SharedBuffer) OpenStream() (*IStream, error) {
	var stream *IStream
	r, _, _ := i.Vtbl.OpenStream.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&stream)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return stream, nil
}

// GetFileMappingHandle 返回指向支持此共享缓冲区的文件映射对象的句柄。
func (i *ICoreWebView2SharedBuffer) GetFileMappingHandle() (uintptr, error) {
	var handle uintptr
	r, _, _ := i.Vtbl.GetFileMappingHandle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handle)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return handle, nil
}

// Close 关闭共享缓冲区.
//   - 当不再需要访问缓冲区时，应用程序应调用此API，以确保即使由于某些引用泄漏而未释放共享缓冲区对象本身，底层资源也能及时释放。共享缓冲区关闭后，先前获取的缓冲区地址和文件映射句柄将失效，无法再使用。访问该对象的属性将失败，并返回 RO_E_CLOSED。对从 OpenStream 返回的 IStream 对象执行诸如读取或写入等操作将失败，并返回 RO_E_CLOSED。PostSharedBufferToScript 也将失败，并返回 RO_E_CLOSED。
//   - 脚本代码一旦不再需要访问共享缓冲区，就应该以共享缓冲区作为参数调用 chrome.webview.releaseBuffer 来释放底层资源。在调用 chrome.webview.releaseBuffer 后，如果脚本尝试访问该缓冲区，将会抛出 JavaScript TypeError 异常，提示正在访问一个已分离的 ArrayBuffer，这与尝试访问一个已转移的 ArrayBuffer 时抛出的异常相同。
//   - 在原生端关闭缓冲区对象不会影响脚本对其的访问，而从脚本释放缓冲区也不会影响原生端对该缓冲区的访问。当原生端和脚本端都释放缓冲区时，底层的共享内存将由操作系统释放。
func (i *ICoreWebView2SharedBuffer) Close() error {
	r, _, _ := i.Vtbl.Close.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetSize 获取共享缓冲区的大小（以字节为单位），出错时会触发全局错误回调。
func (i *ICoreWebView2SharedBuffer) MustGetSize() uint64 {
	size, err := i.GetSize()
	ReportErrorAuto(err)
	return size
}

// MustGetBuffer 获取共享缓冲区的内存地址，出错时会触发全局错误回调。
func (i *ICoreWebView2SharedBuffer) MustGetBuffer() unsafe.Pointer {
	buffer, err := i.GetBuffer()
	ReportErrorAuto(err)
	return buffer
}

// MustOpenStream 获取一个可用于访问共享缓冲区的 IStream 对象，出错时会触发全局错误回调。
func (i *ICoreWebView2SharedBuffer) MustOpenStream() *IStream {
	stream, err := i.OpenStream()
	ReportErrorAuto(err)
	return stream
}

// MustGetFileMappingHandle 返回指向支持此共享缓冲区的文件映射对象的句柄，出错时会触发全局错误回调。
func (i *ICoreWebView2SharedBuffer) MustGetFileMappingHandle() uintptr {
	handle, err := i.GetFileMappingHandle()
	ReportErrorAuto(err)
	return handle
}
