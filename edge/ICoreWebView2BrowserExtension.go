package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2BrowserExtension 提供了一组用于管理扩展的属性，包括ID、名称、是否启用，以及移除扩展、启用或禁用扩展的功能。
type ICoreWebView2BrowserExtension struct {
	Vtbl *ICoreWebView2BrowserExtensionVtbl
}

type ICoreWebView2BrowserExtensionVtbl struct {
	IUnknownVtbl
	GetId        ComProc
	GetName      ComProc
	Remove       ComProc
	GetIsEnabled ComProc
	Enable       ComProc
}

func (i *ICoreWebView2BrowserExtension) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtension) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2BrowserExtension) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetId 获取浏览器扩展的ID。
//   - 这与浏览器扩展 API chrome.runtime.id 返回的浏览器扩展 ID 相同。有关该 ID 的生成方式的更多详细信息，请参阅该文档。
//   - 扩展被移除后，调用 Id 将返回被移除的扩展的 ID。
func (i *ICoreWebView2BrowserExtension) GetId() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	id := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return id, nil
}

// GetName 获取浏览器扩展的名称。
//   - 该值在本浏览器扩展的 manifest.json 文件中定义。
//   - 如果 manifest.json 定义了扩展的本地化名称，此值将是该名称的本地化版本。
//   - 扩展被移除后，调用 Name 将返回已移除扩展的名称。
func (i *ICoreWebView2BrowserExtension) GetName() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	name := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return name, nil
}

// Remove 移除浏览器扩展。
//   - 浏览器扩展会立即被移除，包括从与该 WebView2 配置文件相关联的所有当前运行的 HTML 文档中移除。
//   - 此移除操作会被持久化，该配置文件未来的使用将不会安装此扩展。
//   - 扩展被移除后，再次调用 Remove 将会导致异常。
func (i *ICoreWebView2BrowserExtension) Remove(handler *ICoreWebView2BrowserExtensionRemoveCompletedHandler) error {
	r, _, _ := i.Vtbl.Remove.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveEx 移除浏览器扩展。
//   - 浏览器扩展会立即被移除，包括从与该 WebView2 配置文件相关联的所有当前运行的 HTML 文档中移除。
//   - 此移除操作会被持久化，该配置文件未来的使用将不会安装此扩展。
//   - 扩展被移除后，再次调用 Remove 将会导致异常。
//
// impl: *WebViewEventImpl。
//
// cb: 完成回调函数，接收错误码。
func (i *ICoreWebView2BrowserExtension) RemoveEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "BrowserExtensionRemoveCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "BrowserExtensionRemoveCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "BrowserExtensionRemoveCompleted")
	}
	return i.Remove((*ICoreWebView2BrowserExtensionRemoveCompletedHandler)(handler))
}

// GetIsEnabled 获取浏览器扩展是否启用。
//   - 扩展程序首次安装时，IsEnable 的默认值为 TRUE。
//   - isEnabled 按配置文件保存。
//   - 扩展程序被移除后，调用 isEnabled 将返回其被移除时的值。
func (i *ICoreWebView2BrowserExtension) GetIsEnabled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// Enable 启用或禁用浏览器扩展。
//   - 此更改会立即应用于与该配置文件相关联的所有 WebView2 中所有 HTML 文档内的扩展。
//   - 扩展被移除后，调用 Enable 将不会改变 IsEnabled 的值。
func (i *ICoreWebView2BrowserExtension) Enable(isEnabled bool, handler *ICoreWebView2BrowserExtensionEnableCompletedHandler) error {
	r, _, _ := i.Vtbl.Enable.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isEnabled),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// EnableEx 启用或禁用浏览器扩展。
//   - 此更改会立即应用于与该配置文件相关联的所有 WebView2 中所有 HTML 文档内的扩展。
//   - 扩展被移除后，调用 Enable 将不会改变 IsEnabled 的值。
//
// impl: *WebViewEventImpl。
//
// isEnabled: 是否启用扩展。
//
// cb: 完成回调函数，接收错误码。
func (i *ICoreWebView2BrowserExtension) EnableEx(impl *WebViewEventImpl, isEnabled bool, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "BrowserExtensionEnableCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "BrowserExtensionEnableCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "BrowserExtensionEnableCompleted")
	}
	return i.Enable(isEnabled, (*ICoreWebView2BrowserExtensionEnableCompletedHandler)(handler))
}

// MustGetId 获取浏览器扩展的ID。出错时会触发全局错误回调。
//   - 这与浏览器扩展 API chrome.runtime.id 返回的浏览器扩展 ID 相同。有关该 ID 的生成方式的更多详细信息，请参阅该文档。
//   - 扩展被移除后，调用 Id 将返回被移除的扩展的 ID。
func (i *ICoreWebView2BrowserExtension) MustGetId() string {
	result, err := i.GetId()
	ReportErrorAuto(err)
	return result
}

// MustGetName 获取浏览器扩展的名称。出错时会触发全局错误回调。
//   - 该值在本浏览器扩展的 manifest.json 文件中定义。
//   - 如果 manifest.json 定义了扩展的本地化名称，此值将是该名称的本地化版本。
//   - 扩展被移除后，调用 Name 将返回已移除扩展的名称。
func (i *ICoreWebView2BrowserExtension) MustGetName() string {
	result, err := i.GetName()
	ReportErrorAuto(err)
	return result
}

// MustGetIsEnabled 获取浏览器扩展是否启用。出错时会触发全局错误回调。
//   - 扩展程序首次安装时，IsEnable 的默认值为 TRUE。
//   - isEnabled 按配置文件保存。
//   - 扩展程序被移除后，调用 isEnabled 将返回其被移除时的值。
func (i *ICoreWebView2BrowserExtension) MustGetIsEnabled() bool {
	result, err := i.GetIsEnabled()
	ReportErrorAuto(err)
	return result
}
