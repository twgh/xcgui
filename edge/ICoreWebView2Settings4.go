package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2Settings4 是 ICoreWebView2Settings3 接口的延续, 支持管理密码自动保存和通用自动填充功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings4
type ICoreWebView2Settings4 struct {
	ICoreWebView2Settings3
}

// GetIsPasswordAutosaveEnabled 获取是否允许自动保存密码。
func (i *ICoreWebView2Settings4) GetIsPasswordAutosaveEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsPasswordAutosaveEnabled 设置是否允许自动保存密码。默认值为 false。
func (i *ICoreWebView2Settings4) SetIsPasswordAutosaveEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsGeneralAutofillEnabled 获取是否允许自动填充表单。
func (i *ICoreWebView2Settings4) GetIsGeneralAutofillEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsGeneralAutofillEnabled 设置是否允许自动填充表单。默认值为 true。
func (i *ICoreWebView2Settings4) SetIsGeneralAutofillEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsPasswordAutosaveEnabled 获取是否允许自动保存密码。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings4) MustGetIsPasswordAutosaveEnabled() bool {
	enabled, err := i.GetIsPasswordAutosaveEnabled()
	ReportErrorAuto(err)
	return enabled
}

// MustGetIsGeneralAutofillEnabled 获取是否允许自动填充表单。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings4) MustGetIsGeneralAutofillEnabled() bool {
	enabled, err := i.GetIsGeneralAutofillEnabled()
	ReportErrorAuto(err)
	return enabled
}
