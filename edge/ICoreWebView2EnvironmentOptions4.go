package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2EnvironmentOptions4 提供用于创建管理自定义方案注册的 WebView2 环境的附加选项。。
//
// 110.0.1587.40
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4
type ICoreWebView2EnvironmentOptions4 struct {
	ICoreWebView2EnvironmentOptions3
}

// GetCustomSchemeRegistrations 获取自定义方案注册列表。
//   - 返回的对象指针不再使用时必须调用 ReleaseCustomSchemeRegistrations 统一释放, 以免内存泄漏。
//   - 也可以自行遍历数组, 调用对象的 Release 方法释放。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) GetCustomSchemeRegistrations() ([]*ICoreWebView2CustomSchemeRegistration, error) {
	var count uint32
	var registrations **ICoreWebView2CustomSchemeRegistration
	r, _, _ := i.Vtbl.GetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&registrations)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(registrations))

	if count == 0 || registrations == nil {
		return nil, nil
	}

	result := make([]*ICoreWebView2CustomSchemeRegistration, count)
	slice := unsafe.Slice(registrations, count)
	for i := range slice {
		if slice[i] != nil { // todo 测试是否需要 AddRef.
			slice[i].AddRef() // 增加引用计数，确保对象存活
		}
		result[i] = slice[i]
	}
	return result, nil
}

// ReleaseCustomSchemeRegistrations 释放不再使用的 *ICoreWebView2CustomSchemeRegistration.
func (i *ICoreWebView2EnvironmentOptions4) ReleaseCustomSchemeRegistrations(registrations []*ICoreWebView2CustomSchemeRegistration) {
	for _, reg := range registrations {
		if reg != nil {
			reg.Release() // 释放每个对象的引用
		}
	}
}

// SetCustomSchemeRegistrations 设置自定义方案注册列表。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) SetCustomSchemeRegistrations(registrations []*ICoreWebView2CustomSchemeRegistration) error {
	r, _, _ := i.Vtbl.SetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(len(registrations)),
		uintptr(unsafe.Pointer(&registrations[0])),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetCustomSchemeRegistrations 获取自定义方案注册列表。出错时会触发全局错误回调。
//   - 返回的对象指针不再使用时必须调用 ReleaseCustomSchemeRegistrations 统一释放, 以免内存泄漏。
//   - 也可以自行遍历数组, 调用对象的 Release 方法释放。
//
// 110.0.1587.40
func (i *ICoreWebView2EnvironmentOptions4) MustGetCustomSchemeRegistrations() []*ICoreWebView2CustomSchemeRegistration {
	registrations, err := i.GetCustomSchemeRegistrations()
	ReportErrorAtuo(err)
	return registrations
}
