package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings 是 WebView2 的设置。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings
type ICoreWebView2Settings struct {
	Vtbl *ICoreWebView2SettingsVtbl
}

type ICoreWebView2SettingsVtbl struct {
	IUnknownVtbl
	GetIsScriptEnabled                ComProc
	PutIsScriptEnabled                ComProc
	GetIsWebMessageEnabled            ComProc
	PutIsWebMessageEnabled            ComProc
	GetAreDefaultScriptDialogsEnabled ComProc
	PutAreDefaultScriptDialogsEnabled ComProc
	GetIsStatusBarEnabled             ComProc
	PutIsStatusBarEnabled             ComProc
	GetAreDevToolsEnabled             ComProc
	PutAreDevToolsEnabled             ComProc
	GetAreDefaultContextMenusEnabled  ComProc
	PutAreDefaultContextMenusEnabled  ComProc
	GetAreHostObjectsAllowed          ComProc
	PutAreHostObjectsAllowed          ComProc
	GetIsZoomControlEnabled           ComProc
	PutIsZoomControlEnabled           ComProc
	GetIsBuiltInErrorPageEnabled      ComProc
	PutIsBuiltInErrorPageEnabled      ComProc
}

func (i *ICoreWebView2Settings) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsScriptEnabled 获取是否在 WebView 中的所有未来导航中启用运行 JavaScript。
//   - 这仅影响文档中的脚本。
//   - 即使脚本被禁用，用 ExecuteScript 注入的脚本也会运行。
func (i *ICoreWebView2Settings) GetIsScriptEnabled() (bool, error) {
	var isScriptEnabled bool
	r, _, _ := i.Vtbl.GetIsScriptEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isScriptEnabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isScriptEnabled, nil
}

// SetIsScriptEnabled 设置是否在 WebView 中的所有未来导航中启用运行 JavaScript。
//   - 这仅影响文档中的脚本。
//   - 即使脚本被禁用，用 ExecuteScript 注入的脚本也会运行。
//   - 默认值为 true。
func (i *ICoreWebView2Settings) SetIsScriptEnabled(isScriptEnabled bool) error {
	r, _, _ := i.Vtbl.PutIsScriptEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isScriptEnabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsWebMessageEnabled 获取是否允许 WebView 发送和接收 WebMessage。
func (i *ICoreWebView2Settings) GetIsWebMessageEnabled() (bool, error) {
	var isWebMessageEnabled bool
	r, _, _ := i.Vtbl.GetIsWebMessageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isWebMessageEnabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isWebMessageEnabled, nil
}

// SetIsWebMessageEnabled 设置是否允许 WebView 发送和接收 WebMessage。默认值为 true。
func (i *ICoreWebView2Settings) SetIsWebMessageEnabled(isWebMessageEnabled bool) error {
	r, _, _ := i.Vtbl.PutIsWebMessageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isWebMessageEnabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDefaultScriptDialogsEnabled 获取是否允许在 WebView 中显示默认的 JavaScript 对话框。
func (i *ICoreWebView2Settings) GetAreDefaultScriptDialogsEnabled() (bool, error) {
	var areDefaultScriptDialogsEnabled bool
	r, _, _ := i.Vtbl.GetAreDefaultScriptDialogsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areDefaultScriptDialogsEnabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return areDefaultScriptDialogsEnabled, nil
}

// SetAreDefaultScriptDialogsEnabled 设置是否允许在 WebView 中显示默认的 JavaScript 对话框。默认值为 true。
func (i *ICoreWebView2Settings) SetAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled bool) error {
	r, _, _ := i.Vtbl.PutAreDefaultScriptDialogsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(areDefaultScriptDialogsEnabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsStatusBarEnabled 获取是否在 WebView 中显示状态栏。
//   - 状态栏通常显示在WebView的左下角，显示用户悬停在链接上时链接的 URI 等信息。
//   - 状态栏 UI 可以通过 web 内容进行更改，不应被视为安全。
func (i *ICoreWebView2Settings) GetIsStatusBarEnabled() (bool, error) {
	var isStatusBarEnabled bool
	r, _, _ := i.Vtbl.GetIsStatusBarEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isStatusBarEnabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isStatusBarEnabled, nil
}

// SetIsStatusBarEnabled 设置是否在 WebView 中显示状态栏。
//   - 状态栏通常显示在WebView的左下角，显示用户悬停在链接上时链接的 URI 等信息。
//   - 默认值为 true。
//   - 状态栏 UI 可以通过 web 内容进行更改，不应被视为安全。
func (i *ICoreWebView2Settings) SetIsStatusBarEnabled(isStatusBarEnabled bool) error {
	r, _, _ := i.Vtbl.PutIsStatusBarEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isStatusBarEnabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDevToolsEnabled 获取是否允许在 WebView 中显示开发者工具。
func (i *ICoreWebView2Settings) GetAreDevToolsEnabled() (bool, error) {
	var areDevToolsEnabled bool
	r, _, _ := i.Vtbl.GetAreDevToolsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areDevToolsEnabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return areDevToolsEnabled, nil
}

// SetAreDevToolsEnabled 设置是否允许在 WebView 中显示开发者工具。默认值为 true。
func (i *ICoreWebView2Settings) SetAreDevToolsEnabled(areDevToolsEnabled bool) error {
	r, _, _ := i.Vtbl.PutAreDevToolsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(areDevToolsEnabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDefaultContextMenusEnabled 获取是否允许在 WebView 中显示默认的上下文菜单。
func (i *ICoreWebView2Settings) GetAreDefaultContextMenusEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetAreDefaultContextMenusEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetAreDefaultContextMenusEnabled 设置是否允许在 WebView 中显示默认的上下文菜单。默认值为 true。
func (i *ICoreWebView2Settings) SetAreDefaultContextMenusEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutAreDefaultContextMenusEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreHostObjectsAllowed 获取是否允许在 WebView 中显示主机对象。
func (i *ICoreWebView2Settings) GetAreHostObjectsAllowed() (bool, error) {
	var allowed bool
	r, _, _ := i.Vtbl.GetAreHostObjectsAllowed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowed)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return allowed, nil
}

// SetAreHostObjectsAllowed 设置是否允许在 WebView 中显示主机对象。默认值为 true。
func (i *ICoreWebView2Settings) SetAreHostObjectsAllowed(allowed bool) error {
	r, _, _ := i.Vtbl.PutAreHostObjectsAllowed.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(allowed),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsZoomControlEnabled 获取是否允许在 WebView 中显示缩放控件。
func (i *ICoreWebView2Settings) GetIsZoomControlEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsZoomControlEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsZoomControlEnabled 设置是否允许在 WebView 中显示缩放控件。默认值为 true。
func (i *ICoreWebView2Settings) SetIsZoomControlEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsZoomControlEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsBuiltInErrorPageEnabled 获取是否显示内置的错误页面。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) GetIsBuiltInErrorPageEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsBuiltInErrorPageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsBuiltInErrorPageEnabled 设置是否显示内置的错误页面。默认值为 true。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) SetIsBuiltInErrorPageEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsBuiltInErrorPageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2Settings2 获取 ICoreWebView2Settings2 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings2() (*ICoreWebView2Settings2, error) {
	var result *ICoreWebView2Settings2
	iidICoreWebView2Settings2 := wapi.NewGUID(IID_ICoreWebView2Settings2)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings2),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings3 获取 ICoreWebView2Settings3 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings3() (*ICoreWebView2Settings3, error) {
	var result *ICoreWebView2Settings3
	iidICoreWebView2Settings3 := wapi.NewGUID(IID_ICoreWebView2Settings3)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings3),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings4 获取 ICoreWebView2Settings4 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings4() (*ICoreWebView2Settings4, error) {
	var result *ICoreWebView2Settings4
	iidICoreWebView2Settings4 := wapi.NewGUID(IID_ICoreWebView2Settings4)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings4),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings5 获取 ICoreWebView2Settings5 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings5() (*ICoreWebView2Settings5, error) {
	var result *ICoreWebView2Settings5
	iidICoreWebView2Settings5 := wapi.NewGUID(IID_ICoreWebView2Settings5)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings5),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings6 获取 ICoreWebView2Settings6 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings6() (*ICoreWebView2Settings6, error) {
	var result *ICoreWebView2Settings6
	iidICoreWebView2Settings6 := wapi.NewGUID(IID_ICoreWebView2Settings6)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings6),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings7 获取 ICoreWebView2Settings7 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings7() (*ICoreWebView2Settings7, error) {
	var result *ICoreWebView2Settings7
	iidICoreWebView2Settings7 := wapi.NewGUID(IID_ICoreWebView2Settings7)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings7),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings8 获取 ICoreWebView2Settings8 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings8() (*ICoreWebView2Settings8, error) {
	var result *ICoreWebView2Settings8
	iidICoreWebView2Settings8 := wapi.NewGUID(IID_ICoreWebView2Settings8)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings8),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Settings9 获取 ICoreWebView2Settings9 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings9() (*ICoreWebView2Settings9, error) {
	var result *ICoreWebView2Settings9
	iidICoreWebView2Settings9 := wapi.NewGUID(IID_ICoreWebView2Settings9)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Settings9),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetIsScriptEnabled 获取是否在 WebView 中的所有未来导航中启用运行 JavaScript。出错时会触发全局错误回调。
//   - 这仅影响文档中的脚本。
//   - 即使脚本被禁用，用 ExecuteScript 注入的脚本也会运行。
func (i *ICoreWebView2Settings) MustGetIsScriptEnabled() bool {
	enabled, err := i.GetIsScriptEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetIsWebMessageEnabled 获取是否允许 WebView 发送和接收 WebMessage。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetIsWebMessageEnabled() bool {
	enabled, err := i.GetIsWebMessageEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetAreDefaultScriptDialogsEnabled 获取是否允许在 WebView 中显示默认的 JavaScript 对话框。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetAreDefaultScriptDialogsEnabled() bool {
	enabled, err := i.GetAreDefaultScriptDialogsEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetIsStatusBarEnabled 获取是否在 WebView 中显示状态栏。出错时会触发全局错误回调。
//   - 状态栏通常显示在WebView的左下角，显示用户悬停在链接上时链接的 URI 等信息。
//   - 状态栏 UI 可以通过 web 内容进行更改，不应被视为安全。
func (i *ICoreWebView2Settings) MustGetIsStatusBarEnabled() bool {
	enabled, err := i.GetIsStatusBarEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetAreDevToolsEnabled 获取是否允许在 WebView 中显示开发者工具。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetAreDevToolsEnabled() bool {
	enabled, err := i.GetAreDevToolsEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetAreDefaultContextMenusEnabled 获取是否允许在 WebView 中显示默认的上下文菜单。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetAreDefaultContextMenusEnabled() bool {
	enabled, err := i.GetAreDefaultContextMenusEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetAreHostObjectsAllowed 获取是否允许在 WebView 中显示主机对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetAreHostObjectsAllowed() bool {
	allowed, err := i.GetAreHostObjectsAllowed()
	ReportErrorAtuo(err)
	return allowed
}

// MustGetIsZoomControlEnabled 获取是否允许在 WebView 中显示缩放控件。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetIsZoomControlEnabled() bool {
	enabled, err := i.GetIsZoomControlEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetIsBuiltInErrorPageEnabled 获取是否显示内置的错误页面。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) MustGetIsBuiltInErrorPageEnabled() bool {
	enabled, err := i.GetIsBuiltInErrorPageEnabled()
	ReportErrorAtuo(err)
	return enabled
}

// MustGetICoreWebView2Settings2 获取 ICoreWebView2Settings2 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings2() *ICoreWebView2Settings2 {
	c, err := i.GetICoreWebView2Settings2()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings3 获取 ICoreWebView2Settings3 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings3() *ICoreWebView2Settings3 {
	c, err := i.GetICoreWebView2Settings3()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings4 获取 ICoreWebView2Settings4 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings4() *ICoreWebView2Settings4 {
	c, err := i.GetICoreWebView2Settings4()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings5 获取 ICoreWebView2Settings5 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings5() *ICoreWebView2Settings5 {
	c, err := i.GetICoreWebView2Settings5()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings6 获取 ICoreWebView2Settings6 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings6() *ICoreWebView2Settings6 {
	c, err := i.GetICoreWebView2Settings6()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings7 获取 ICoreWebView2Settings7 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings7() *ICoreWebView2Settings7 {
	c, err := i.GetICoreWebView2Settings7()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings8 获取 ICoreWebView2Settings8 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings8() *ICoreWebView2Settings8 {
	c, err := i.GetICoreWebView2Settings8()
	ReportErrorAtuo(err)
	return c
}

// MustGetICoreWebView2Settings9 获取 ICoreWebView2Settings9 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings9() *ICoreWebView2Settings9 {
	c, err := i.GetICoreWebView2Settings9()
	ReportErrorAtuo(err)
	return c
}
