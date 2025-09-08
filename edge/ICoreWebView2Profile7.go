package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile7 是 ICoreWebView2Profile6 的延续，提供了浏览器扩展管理功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile7
type ICoreWebView2Profile7 struct {
	ICoreWebView2Profile6
}

// AddBrowserExtension 使用本地设备中非打包扩展程序的扩展路径添加浏览器扩展。
//   - 扩展程序在安装后立即运行。
//   - 扩展程序文件夹路径是未打包浏览器扩展的最顶层文件夹，且包含浏览器扩展清单文件。
//   - 如果 extensionFolderPath 是无效路径或不包含扩展程序 manifest.json 文件，此函数将向调用者返回 ERROR_FILE_NOT_FOUND。
//   - 已安装的扩展程序的 IsEnabled 默认值为 true。
//   - 当 AreBrowserExtensionsEnabled 为 FALSE 时，AddBrowserExtension 将失败并返回 HRESULT ERROR_NOT_SUPPORTED。
//   - 安装过程中，扩展程序的内容不会复制到用户数据文件夹。一旦扩展程序安装完成，修改其内容会导致该扩展程序从已安装的配置文件中被移除。
//   - 添加扩展程序后，该扩展程序会持久化到相应的配置文件中。下次使用此配置文件时，该扩展程序仍处于安装状态。
//   - 从文件夹路径安装扩展程序时，从相同文件夹路径添加相同的扩展程序意味着重新安装该扩展程序。当安装了两个具有相同 ID 的扩展程序时，只会保留较晚安装的那个扩展程序。
//
// extensionFolderPath: 扩展文件夹路径.
func (i *ICoreWebView2Profile7) AddBrowserExtension(extensionFolderPath string, handler *ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) error {
	_extensionFolderPath, err := syscall.UTF16PtrFromString(extensionFolderPath)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.AddBrowserExtension.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_extensionFolderPath)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetBrowserExtensions 获取已安装的扩展程序集的快照。
//   - 如果在 GetBrowserExtensions 完成后安装或卸载扩展，GetBrowserExtensions 返回的列表将保持不变。
//   - 当 AreBrowserExtensionsEnabled 为 FALSE 时，不会返回当前用户配置文件上的任何扩展。
func (i *ICoreWebView2Profile7) GetBrowserExtensions(handler *ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) error {
	r, _, _ := i.Vtbl.GetBrowserExtensions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddBrowserExtensionEx 使用本地设备中非打包扩展程序的扩展路径添加浏览器扩展。
//   - 扩展程序在安装后立即运行。
//   - 扩展程序文件夹路径是未打包浏览器扩展的最顶层文件夹，且包含浏览器扩展清单文件。
//   - 如果 extensionFolderPath 是无效路径或不包含扩展程序 manifest.json 文件，此函数将向调用者返回 ERROR_FILE_NOT_FOUND。
//   - 已安装的扩展程序的 IsEnabled 默认值为 true。
//   - 当 AreBrowserExtensionsEnabled 为 FALSE 时，AddBrowserExtension 将失败并返回 HRESULT ERROR_NOT_SUPPORTED。
//   - 安装过程中，扩展程序的内容不会复制到用户数据文件夹。一旦扩展程序安装完成，修改其内容会导致该扩展程序从已安装的配置文件中被移除。
//   - 添加扩展程序后，该扩展程序会持久化到相应的配置文件中。下次使用此配置文件时，该扩展程序仍处于安装状态。
//   - 从文件夹路径安装扩展程序时，从相同文件夹路径添加相同的扩展程序意味着重新安装该扩展程序。当安装了两个具有相同 ID 的扩展程序时，只会保留较晚安装的那个扩展程序。
//
// impl: *WebViewEventImpl。
//
// extensionFolderPath: 扩展文件夹路径.
//
// cb: 添加完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2Profile7) AddBrowserExtensionEx(impl *WebViewEventImpl, extensionFolderPath string, cb func(errorCode syscall.Errno, result *ICoreWebView2BrowserExtension) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ProfileAddBrowserExtensionCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ProfileAddBrowserExtensionCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ProfileAddBrowserExtensionCompleted")
	}
	return i.AddBrowserExtension(extensionFolderPath, (*ICoreWebView2ProfileAddBrowserExtensionCompletedHandler)(handler))
}

// GetBrowserExtensionsEx 获取已安装的扩展程序集的快照。
//   - 如果在 GetBrowserExtensions 完成后安装或卸载扩展，GetBrowserExtensions 返回的列表将保持不变。
//   - 当 AreBrowserExtensionsEnabled 为 FALSE 时，不会返回当前用户配置文件上的任何扩展。
//
// impl: *WebViewEventImpl。
//
// cb: 获取完成后的回调处理程序，可以为 nil。如果获取成功，将通过 cb 返回扩展列表。
func (i *ICoreWebView2Profile7) GetBrowserExtensionsEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno, result *ICoreWebView2BrowserExtensionList) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ProfileGetBrowserExtensionsCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ProfileGetBrowserExtensionsCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ProfileGetBrowserExtensionsCompleted")
	}
	return i.GetBrowserExtensions((*ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler)(handler))
}
