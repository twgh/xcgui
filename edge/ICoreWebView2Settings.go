package edge

// ok

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
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
	r, _, err := i.Vtbl.GetIsScriptEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isScriptEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isScriptEnabled, nil
}

// MustGetIsScriptEnabled 获取是否在 WebView 中的所有未来导航中启用运行 JavaScript。忽略错误。
//   - 这仅影响文档中的脚本。
//   - 即使脚本被禁用，用 ExecuteScript 注入的脚本也会运行。
func (i *ICoreWebView2Settings) MustGetIsScriptEnabled() bool {
	enabled, _ := i.GetIsScriptEnabled()
	return enabled
}

// PutIsScriptEnabled 设置是否在 WebView 中的所有未来导航中启用运行 JavaScript。
//   - 这仅影响文档中的脚本。
//   - 即使脚本被禁用，用 ExecuteScript 注入的脚本也会运行。
//   - 默认值为 true。
func (i *ICoreWebView2Settings) PutIsScriptEnabled(isScriptEnabled bool) error {
	r, _, err := i.Vtbl.PutIsScriptEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(isScriptEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsWebMessageEnabled 获取是否允许 WebView 发送和接收 WebMessage。
func (i *ICoreWebView2Settings) GetIsWebMessageEnabled() (bool, error) {
	var isWebMessageEnabled bool
	r, _, err := i.Vtbl.GetIsWebMessageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isWebMessageEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isWebMessageEnabled, nil
}

// MustGetIsWebMessageEnabled 获取是否允许 WebView 发送和接收 WebMessage。忽略错误。
func (i *ICoreWebView2Settings) MustGetIsWebMessageEnabled() bool {
	enabled, _ := i.GetIsWebMessageEnabled()
	return enabled
}

// PutIsWebMessageEnabled 设置是否允许 WebView 发送和接收 WebMessage。默认值为 true。
func (i *ICoreWebView2Settings) PutIsWebMessageEnabled(isWebMessageEnabled bool) error {
	r, _, err := i.Vtbl.PutIsWebMessageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(isWebMessageEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDefaultScriptDialogsEnabled 获取是否允许在 WebView 中显示默认的 JavaScript 对话框。
func (i *ICoreWebView2Settings) GetAreDefaultScriptDialogsEnabled() (bool, error) {
	var areDefaultScriptDialogsEnabled bool
	r, _, err := i.Vtbl.GetAreDefaultScriptDialogsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areDefaultScriptDialogsEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return areDefaultScriptDialogsEnabled, nil
}

// MustGetAreDefaultScriptDialogsEnabled 获取是否允许在 WebView 中显示默认的 JavaScript 对话框。忽略错误。
func (i *ICoreWebView2Settings) MustGetAreDefaultScriptDialogsEnabled() bool {
	enabled, _ := i.GetAreDefaultScriptDialogsEnabled()
	return enabled
}

// PutAreDefaultScriptDialogsEnabled 设置是否允许在 WebView 中显示默认的 JavaScript 对话框。默认值为 true。
func (i *ICoreWebView2Settings) PutAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled bool) error {
	r, _, err := i.Vtbl.PutAreDefaultScriptDialogsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(areDefaultScriptDialogsEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
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
	r, _, err := i.Vtbl.GetIsStatusBarEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isStatusBarEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isStatusBarEnabled, nil
}

// MustGetIsStatusBarEnabled 获取是否在 WebView 中显示状态栏。忽略错误。
//   - 状态栏通常显示在WebView的左下角，显示用户悬停在链接上时链接的 URI 等信息。
//   - 状态栏 UI 可以通过 web 内容进行更改，不应被视为安全。
func (i *ICoreWebView2Settings) MustGetIsStatusBarEnabled() bool {
	enabled, _ := i.GetIsStatusBarEnabled()
	return enabled
}

// PutIsStatusBarEnabled 设置是否在 WebView 中显示状态栏。
//   - 状态栏通常显示在WebView的左下角，显示用户悬停在链接上时链接的 URI 等信息。
//   - 默认值为 true。
//   - 状态栏 UI 可以通过 web 内容进行更改，不应被视为安全。
func (i *ICoreWebView2Settings) PutIsStatusBarEnabled(isStatusBarEnabled bool) error {
	r, _, err := i.Vtbl.PutIsStatusBarEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(isStatusBarEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDevToolsEnabled 获取是否允许在 WebView 中显示开发者工具。
func (i *ICoreWebView2Settings) GetAreDevToolsEnabled() (bool, error) {
	var areDevToolsEnabled bool
	r, _, err := i.Vtbl.GetAreDevToolsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areDevToolsEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return areDevToolsEnabled, nil
}

// MustGetAreDevToolsEnabled 获取是否允许在 WebView 中显示开发者工具。忽略错误。
func (i *ICoreWebView2Settings) MustGetAreDevToolsEnabled() bool {
	enabled, _ := i.GetAreDevToolsEnabled()
	return enabled
}

// PutAreDevToolsEnabled 设置是否允许在 WebView 中显示开发者工具。默认值为 true。
func (i *ICoreWebView2Settings) PutAreDevToolsEnabled(areDevToolsEnabled bool) error {
	r, _, err := i.Vtbl.PutAreDevToolsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(areDevToolsEnabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreDefaultContextMenusEnabled 获取是否允许在 WebView 中显示默认的上下文菜单。
func (i *ICoreWebView2Settings) GetAreDefaultContextMenusEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetAreDefaultContextMenusEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// MustGetAreDefaultContextMenusEnabled 获取是否允许在 WebView 中显示默认的上下文菜单。忽略错误。
func (i *ICoreWebView2Settings) MustGetAreDefaultContextMenusEnabled() bool {
	enabled, _ := i.GetAreDefaultContextMenusEnabled()
	return enabled
}

// PutAreDefaultContextMenusEnabled 设置是否允许在 WebView 中显示默认的上下文菜单。默认值为 true。
func (i *ICoreWebView2Settings) PutAreDefaultContextMenusEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutAreDefaultContextMenusEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreHostObjectsAllowed 获取是否允许在 WebView 中显示主机对象。
func (i *ICoreWebView2Settings) GetAreHostObjectsAllowed() (bool, error) {
	var allowed bool
	r, _, err := i.Vtbl.GetAreHostObjectsAllowed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowed)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return allowed, nil
}

// MustGetAreHostObjectsAllowed 获取是否允许在 WebView 中显示主机对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetAreHostObjectsAllowed() bool {
	allowed, _ := i.GetAreHostObjectsAllowed()
	return allowed
}

// PutAreHostObjectsAllowed 设置是否允许在 WebView 中显示主机对象。默认值为 true。
func (i *ICoreWebView2Settings) PutAreHostObjectsAllowed(allowed bool) error {
	r, _, err := i.Vtbl.PutAreHostObjectsAllowed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(allowed)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsZoomControlEnabled 获取是否允许在 WebView 中显示缩放控件。
func (i *ICoreWebView2Settings) GetIsZoomControlEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetIsZoomControlEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// MustGetIsZoomControlEnabled 获取是否允许在 WebView 中显示缩放控件。忽略错误。
func (i *ICoreWebView2Settings) MustGetIsZoomControlEnabled() bool {
	enabled, _ := i.GetIsZoomControlEnabled()
	return enabled
}

// PutIsZoomControlEnabled 设置是否允许在 WebView 中显示缩放控件。默认值为 true。
func (i *ICoreWebView2Settings) PutIsZoomControlEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutIsZoomControlEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsBuiltInErrorPageEnabled 获取是否显示内置的错误页面。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) GetIsBuiltInErrorPageEnabled() (bool, error) {
	var enabled bool
	r, _, err := i.Vtbl.GetIsBuiltInErrorPageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// MustGetIsBuiltInErrorPageEnabled 获取是否显示内置的错误页面。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) MustGetIsBuiltInErrorPageEnabled() bool {
	enabled, _ := i.GetIsBuiltInErrorPageEnabled()
	return enabled
}

// PutIsBuiltInErrorPageEnabled 设置是否显示内置的错误页面。默认值为 true。禁用时，当发生相关错误时，会显示一个空白页。
func (i *ICoreWebView2Settings) PutIsBuiltInErrorPageEnabled(enabled bool) error {
	r, _, err := i.Vtbl.PutIsBuiltInErrorPageEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(BoolToInt(enabled)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2Settings2 获取 ICoreWebView2Settings2 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings2() (*ICoreWebView2Settings2, error) {
	var result *ICoreWebView2Settings2
	iidICoreWebView2Settings2 := NewGUID(IID_ICoreWebView2Settings2)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings2)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings2 获取 ICoreWebView2Settings2 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings2() *ICoreWebView2Settings2 {
	c, _ := i.GetICoreWebView2Settings2()
	return c
}

// GetICoreWebView2Settings3 获取 ICoreWebView2Settings3 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings3() (*ICoreWebView2Settings3, error) {
	var result *ICoreWebView2Settings3
	iidICoreWebView2Settings3 := NewGUID(IID_ICoreWebView2Settings3)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings3)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings3 获取 ICoreWebView2Settings3 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings3() *ICoreWebView2Settings3 {
	c, _ := i.GetICoreWebView2Settings3()
	return c
}

// GetICoreWebView2Settings4 获取 ICoreWebView2Settings4 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings4() (*ICoreWebView2Settings4, error) {
	var result *ICoreWebView2Settings4
	iidICoreWebView2Settings4 := NewGUID(IID_ICoreWebView2Settings4)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings4)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings4 获取 ICoreWebView2Settings4 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings4() *ICoreWebView2Settings4 {
	c, _ := i.GetICoreWebView2Settings4()
	return c
}

// GetICoreWebView2Settings5 获取 ICoreWebView2Settings5 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings5() (*ICoreWebView2Settings5, error) {
	var result *ICoreWebView2Settings5
	iidICoreWebView2Settings5 := NewGUID(IID_ICoreWebView2Settings5)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings5)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings5 获取 ICoreWebView2Settings5 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings5() *ICoreWebView2Settings5 {
	c, _ := i.GetICoreWebView2Settings5()
	return c
}

// GetICoreWebView2Settings6 获取 ICoreWebView2Settings6 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings6() (*ICoreWebView2Settings6, error) {
	var result *ICoreWebView2Settings6
	iidICoreWebView2Settings6 := NewGUID(IID_ICoreWebView2Settings6)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings6)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings6 获取 ICoreWebView2Settings6 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings6() *ICoreWebView2Settings6 {
	c, _ := i.GetICoreWebView2Settings6()
	return c
}

// GetICoreWebView2Settings7 获取 ICoreWebView2Settings7 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings7() (*ICoreWebView2Settings7, error) {
	var result *ICoreWebView2Settings7
	iidICoreWebView2Settings7 := NewGUID(IID_ICoreWebView2Settings7)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings7)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings7 获取 ICoreWebView2Settings7 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings7() *ICoreWebView2Settings7 {
	c, _ := i.GetICoreWebView2Settings7()
	return c
}

// GetICoreWebView2Settings8 获取 ICoreWebView2Settings8 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings8() (*ICoreWebView2Settings8, error) {
	var result *ICoreWebView2Settings8
	iidICoreWebView2Settings8 := NewGUID(IID_ICoreWebView2Settings8)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings8)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings8 获取 ICoreWebView2Settings8 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings8() *ICoreWebView2Settings8 {
	c, _ := i.GetICoreWebView2Settings8()
	return c
}

// GetICoreWebView2Settings9 获取 ICoreWebView2Settings9 对象。
func (i *ICoreWebView2Settings) GetICoreWebView2Settings9() (*ICoreWebView2Settings9, error) {
	var result *ICoreWebView2Settings9
	iidICoreWebView2Settings9 := NewGUID(IID_ICoreWebView2Settings9)
	err := i.QueryInterface(
		uintptr(unsafe.Pointer(iidICoreWebView2Settings9)),
		uintptr(unsafe.Pointer(&result)))
	return result, err
}

// MustGetICoreWebView2Settings9 获取 ICoreWebView2Settings9 对象。忽略错误。
func (i *ICoreWebView2Settings) MustGetICoreWebView2Settings9() *ICoreWebView2Settings9 {
	c, _ := i.GetICoreWebView2Settings9()
	return c
}
