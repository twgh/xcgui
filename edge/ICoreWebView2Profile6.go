package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2Profile6 是 ICoreWebView2Profile5 的延续，提供了密码自动保存和通用自动填充功能的管理。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile6
type ICoreWebView2Profile6 struct {
	ICoreWebView2Profile5
}

// GetIsPasswordAutosaveEnabled 获取是否启用密码自动保存。
//   - IsPasswordAutosaveEnabled 属性的行为独立于 IsGeneralAutofillEnabled 属性。
//   - 当 IsPasswordAutosaveEnabled 为 false 时，不会保存新的密码数据，也不会显示“保存/更新密码”提示。但是，如果在禁用此设置之前已经保存了密码数据，则会自动填充该密码信息，显示建议，点击其中一个建议就会填充字段。
//   - 当 IsPasswordAutosaveEnabled 为 true 时，会自动填充密码信息，显示建议，点击其中一个建议会填充字段，保存新数据，并显示“保存/更新密码”提示。设置后立即生效。默认值为 FALSE。
//   - 此属性与 ICoreWebView2Settings.IsPasswordAutosaveEnabled 的值相同，更改其中一个会改变另一个。
//   - 所有具有相同 ICoreWebView2Profile 的 ICoreWebView2 将共享此属性的相同值，因此对于具有相同配置文件的 ICoreWebView2，它们的 ICoreWebView2Settings.IsPasswordAutosaveEnabled 和 ICoreWebView2Profile.IsPasswordAutosaveEnabled 将始终具有相同的值。
func (i *ICoreWebView2Profile6) GetIsPasswordAutosaveEnabled() (bool, error) {
	var value int32
	r, _, _ := i.Vtbl.GetIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value != 0, nil
}

// SetIsPasswordAutosaveEnabled 设置是否启用密码自动保存。
func (i *ICoreWebView2Profile6) SetIsPasswordAutosaveEnabled(value bool) error {
	r, _, _ := i.Vtbl.PutIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsGeneralAutofillEnabled 获取是否启用通用自动填充。
func (i *ICoreWebView2Profile6) GetIsGeneralAutofillEnabled() (bool, error) {
	var value int32
	r, _, _ := i.Vtbl.GetIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value != 0, nil
}

// SetIsGeneralAutofillEnabled 设置是否启用通用自动填充。
//   - 这不包括密码和信用卡信息。
//   - 当 IsGeneralAutofillEnabled 为 false 时，不会显示任何建议，也不会保存新信息。
//   - 当 IsGeneralAutofillEnabled 为 true 时，会保存信息，显示建议，点击其中一个建议会填充表单字段。设置后立即生效。默认值为 TRUE。
//   - 此属性与 ICoreWebView2Settings.IsGeneralAutofillEnabled 的值相同，更改其中一个会使另一个也发生更改。
//   - 所有具有相同 ICoreWebView2Profile 的 ICoreWebView2 将共享此属性的相同值，因此对于具有相同配置文件的 ICoreWebView2，它们的 ICoreWebView2Settings.IsGeneralAutofillEnabled 和 ICoreWebView2Profile.IsGeneralAutofillEnabled 将始终具有相同的值。
func (i *ICoreWebView2Profile6) SetIsGeneralAutofillEnabled(value bool) error {
	r, _, _ := i.Vtbl.PutIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsPasswordAutosaveEnabled 获取是否启用密码自动保存。出错时会触发全局错误回调。
//   - IsPasswordAutosaveEnabled 属性的行为独立于 IsGeneralAutofillEnabled 属性。
//   - 当 IsPasswordAutosaveEnabled 为 false 时，不会保存新的密码数据，也不会显示“保存/更新密码”提示。但是，如果在禁用此设置之前已经保存了密码数据，则会自动填充该密码信息，显示建议，点击其中一个建议就会填充字段。
//   - 当 IsPasswordAutosaveEnabled 为 true 时，会自动填充密码信息，显示建议，点击其中一个建议会填充字段，保存新数据，并显示“保存/更新密码”提示。设置后立即生效。默认值为 FALSE。
//   - 此属性与 ICoreWebView2Settings.IsPasswordAutosaveEnabled 的值相同，更改其中一个会改变另一个。
//   - 所有具有相同 ICoreWebView2Profile 的 ICoreWebView2 将共享此属性的相同值，因此对于具有相同配置文件的 ICoreWebView2，它们的 ICoreWebView2Settings.IsPasswordAutosaveEnabled 和 ICoreWebView2Profile.IsPasswordAutosaveEnabled 将始终具有相同的值。
func (i *ICoreWebView2Profile6) MustGetIsPasswordAutosaveEnabled() bool {
	result, err := i.GetIsPasswordAutosaveEnabled()
	ReportErrorAuto(err)
	return result
}

// MustGetIsGeneralAutofillEnabled 获取是否启用通用自动填充。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile6) MustGetIsGeneralAutofillEnabled() bool {
	result, err := i.GetIsGeneralAutofillEnabled()
	ReportErrorAuto(err)
	return result
}
