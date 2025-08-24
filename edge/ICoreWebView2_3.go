package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_3 是 ICoreWebView2_2 接口的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_3
type ICoreWebView2_3 struct {
	ICoreWebView2_2
}

// GetIsSuspended 获取 WebView 控件是否已挂起。
func (i *ICoreWebView2_3) GetIsSuspended() (bool, error) {
	var result bool
	r, _, _ := i.Vtbl.GetIsSuspended.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&result)),
	)
	if r != 0 {
		return result, syscall.Errno(r)
	}
	return result, nil
}

// Resume 恢复 WebView，以便它恢复网页上的活动。
func (i *ICoreWebView2_3) Resume() error {
	r, _, _ := i.Vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// TrySuspend 尝试挂起 WebView 控件, 以节省内存。
//
// handler: 接收 TrySuspend 方法的结果。
func (i *ICoreWebView2_3) TrySuspend(handler *ICoreWebView2TrySuspendCompletedHandler) error {
	r, _, _ := i.Vtbl.TrySuspend.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// TrySuspendEx 尝试挂起 WebView 控件, 以节省内存。
//
// impl: *WebViewEventImpl.
//
// cb: 接收 TrySuspend 方法的结果。
func (i *ICoreWebView2_3) TrySuspendEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, isSuccessful bool) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "TrySuspendCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "TrySuspendCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "TrySuspendCompleted")
	}
	return i.TrySuspend((*ICoreWebView2TrySuspendCompletedHandler)(handler))
}

// ClearVirtualHostNameToFolderMapping 清除 SetVirtualHostNameToFolderMapping 添加的本地文件夹的主机名映射。
//
// hostName: 要清除的主机名。
func (i *ICoreWebView2_3) ClearVirtualHostNameToFolderMapping(hostName string) error {
	_hostName, err := syscall.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}

	r, _, _ := i.Vtbl.ClearVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetVirtualHostNameToFolderMapping 设置虚拟主机名和文件夹路径之间的映射，以便通过该主机名对网站可用.
//   - 设置映射后，在 WebView 中加载的文档可以使用由 hostName 指定的主机名处的 HTTP 或 HTTPS URL，来访问由 folderPath 指定的本地文件夹中的文件。
//   - 此映射适用于顶级文档和 iframe 导航，以及文档中的子资源引用。这也适用于 Web Worker，包括专用/共享 Worker 和服务 Worker，用于加载 Worker 脚本或从 Worker 内部发出的子资源（importScripts()、fetch()、XHR等）。为了使虚拟主机映射与服务 Worker 一起正常工作，请确保在共享同一浏览器实例的所有 WebView 之间保持虚拟主机名映射一致。由于服务 Worker 独立于 WebView 工作，我们在解析虚拟主机名时会合并所有 WebView 的映射，WebView 之间不一致的映射可能会导致意外行为。
//   - 由于当前的一个实现限制，使用虚拟主机名访问的媒体文件加载可能非常缓慢。由于当前页面的资源加载器可能已经创建并正在运行，映射的更改可能不会应用到当前页面，需要重新加载页面才能应用新的映射。
//
// hostName: 要映射的主机名。映射不同文件夹应使用不同的主机名。
//
// folderPath: 要映射的文件夹路径, 长度不得超过 Windows 系统的 MAX_PATH 限制。
//
// accessKind: 指定了其他站点对虚拟主机下资源的访问级别。
func (i *ICoreWebView2_3) SetVirtualHostNameToFolderMapping(hostName, folderPath string, accessKind COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND) error {
	_hostName, err := syscall.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}
	_folderPath, err := syscall.UTF16PtrFromString(folderPath)
	if err != nil {
		return err
	}

	r, _, _ := i.Vtbl.SetVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
		uintptr(unsafe.Pointer(_folderPath)),
		uintptr(accessKind),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsSuspended 获取 WebView 控件是否已挂起。出错时会触发全局错误回调。
func (i *ICoreWebView2_3) MustGetIsSuspended() bool {
	r, err := i.GetIsSuspended()
	ReportErrorAtuo(err)
	return r
}
