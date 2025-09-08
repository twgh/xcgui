package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ContextMenuItem 一个 WebView2 上下文菜单项.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem
type ICoreWebView2ContextMenuItem struct {
	Vtbl *ICoreWebView2ContextMenuItemVtbl
}

type ICoreWebView2ContextMenuItemVtbl struct {
	IUnknownVtbl
	GetName                   ComProc
	GetLabel                  ComProc
	GetCommandId              ComProc
	GetShortcutKeyDescription ComProc
	GetIcon                   ComProc
	GetKind                   ComProc
	PutIsEnabled              ComProc
	GetIsEnabled              ComProc
	PutIsChecked              ComProc
	GetIsChecked              ComProc
	GetChildren               ComProc
	AddCustomItemSelected     ComProc
	RemoveCustomItemSelected  ComProc
}

func (i *ICoreWebView2ContextMenuItem) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuItem) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuItem) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetName 获取菜单项的未本地化名称.
//   - 使用此来区分上下文菜单项类型。这将是菜单项的英文标签，采用小驼峰式命名法。
//   - 例如，“另存为”菜单项将是“saveAs”。扩展菜单项将是“extension”，自定义菜单项将是“custom”，拼写检查项将是“spellCheck”。
//   - 还有: copyImage, openLinkInNewWindow, cut, copy, paste等.
func (i *ICoreWebView2ContextMenuItem) GetName() (string, error) {
	var name *uint16
	r, _, _ := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&name)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(name)
	wapi.CoTaskMemFree(unsafe.Pointer(name))
	return result, nil
}

// GetLabel 获取菜单项显示的标签. 将包含用作键盘快捷键的字符的 & 号。
func (i *ICoreWebView2ContextMenuItem) GetLabel() (string, error) {
	var label *uint16
	r, _, _ := i.Vtbl.GetLabel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&label)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(label)
	wapi.CoTaskMemFree(unsafe.Pointer(label))
	return result, nil
}

// GetCommandId 获取菜单项命令ID.
//   - 使用此方法在 ContextMenuRequested 事件中报告 SelectedCommandId。
func (i *ICoreWebView2ContextMenuItem) GetCommandId() (int32, error) {
	var commandId int32
	r, _, _ := i.Vtbl.GetCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&commandId)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return commandId, nil
}

// GetShortcutKeyDescription 获取此上下文菜单项的快捷键描述.
//   - 如果没有键盘快捷键，该值将为空字符串。
//   - 此文本旨在向最终用户显示键盘快捷键。例如，“检查”菜单项的此属性为 Ctrl+Shift+I。
func (i *ICoreWebView2ContextMenuItem) GetShortcutKeyDescription() (string, error) {
	var desc *uint16
	r, _, _ := i.Vtbl.GetShortcutKeyDescription.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&desc)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(desc)
	wapi.CoTaskMemFree(unsafe.Pointer(desc))
	return result, nil
}

// GetIcon 获取菜单项图标.
//   - 以 IStream 的形式获取菜单项的 PNG、位图或 SVG 格式的图标。
//   - 获取后要释放 IStream.
func (i *ICoreWebView2ContextMenuItem) GetIcon() ([]byte, error) {
	var streamPtr uintptr
	r, _, _ := i.Vtbl.GetIcon.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&streamPtr)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	if streamPtr == 0 {
		return nil, nil
	}
	stream := NewIStreamByPtr(streamPtr)
	icon, err := stream.GetBytesAndRelease()
	if err != nil {
		return nil, err
	}
	return icon, nil
}

// GetKind 获取菜单项类型.
func (i *ICoreWebView2ContextMenuItem) GetKind() (COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND, error) {
	var kind COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// SetIsEnabled 设置菜单项是否启用.
//   - 此选项仅用于自定义上下文菜单项的情况。其默认值为 TRUE。
func (i *ICoreWebView2ContextMenuItem) SetIsEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsEnabled 获取菜单项是否启用.
func (i *ICoreWebView2ContextMenuItem) GetIsEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetIsChecked 设置菜单项是否选中. 仅可用于类型为复选框或单选按钮的自定义上下文菜单项。
func (i *ICoreWebView2ContextMenuItem) SetIsChecked(checked bool) error {
	r, _, _ := i.Vtbl.PutIsChecked.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(checked),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsChecked 获取菜单项是否选中，在类型为“复选框”或“单选框”时使用。
func (i *ICoreWebView2ContextMenuItem) GetIsChecked() (bool, error) {
	var checked bool
	r, _, _ := i.Vtbl.GetIsChecked.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&checked)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return checked, nil
}

// GetChildren 获取子菜单项集合, 如果类型为子菜单.
func (i *ICoreWebView2ContextMenuItem) GetChildren() (*ICoreWebView2ContextMenuItemCollection, error) {
	var children *ICoreWebView2ContextMenuItemCollection
	r, _, _ := i.Vtbl.GetChildren.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&children)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return children, nil
}

// Event_CustomItemSelected 自定义菜单项选中事件.
//   - 此事件仅针对最终开发人员创建的上下文菜单项引发。
func (i *ICoreWebView2ContextMenuItem) Event_CustomItemSelected(w *WebViewEventImpl, cb func(sender *ICoreWebView2ContextMenuItem, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(w, "CustomItemSelected", cb, i, allowAddingMultiple...)
}

// AddCustomItemSelected 添加自定义菜单项选中事件处理程序.
//   - 此事件仅针对最终开发人员创建的上下文菜单项引发。
func (i *ICoreWebView2ContextMenuItem) AddCustomItemSelected(eventHandler *ICoreWebView2CustomItemSelectedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddCustomItemSelected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveCustomItemSelected 移除自定义菜单项选中事件处理程序
func (i *ICoreWebView2ContextMenuItem) RemoveCustomItemSelected(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveCustomItemSelected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetName 获取菜单项的未本地化名称。出错时会触发全局错误回调。
//   - 使用此来区分上下文菜单项类型。这将是菜单项的英文标签，采用小驼峰式命名法。
//   - 例如，“另存为”菜单项将是“saveAs”。扩展菜单项将是“extension”，自定义菜单项将是“custom”，拼写检查项将是“spellCheck”。
//   - 还有: copyImage, openLinkInNewWindow, cut, copy, paste等.
func (i *ICoreWebView2ContextMenuItem) MustGetName() string {
	name, err := i.GetName()
	ReportErrorAuto(err)
	return name
}

// MustGetLabel 获取菜单项显示的标签. 将包含用作键盘快捷键的字符的 & 号。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItem) MustGetLabel() string {
	label, err := i.GetLabel()
	ReportErrorAuto(err)
	return label
}

// MustGetCommandId 获取菜单项命令ID。出错时会触发全局错误回调。
//   - 使用此方法在 ContextMenuRequested 事件中报告 SelectedCommandId。
func (i *ICoreWebView2ContextMenuItem) MustGetCommandId() int32 {
	commandId, err := i.GetCommandId()
	ReportErrorAuto(err)
	return commandId
}

// MustGetShortcutKeyDescription 获取此上下文菜单项的快捷键描述。出错时会触发全局错误回调。
//   - 如果没有键盘快捷键，该值将为空字符串。
//   - 此文本旨在向最终用户显示键盘快捷键。例如，“检查”菜单项的此属性为 Ctrl+Shift+I。
func (i *ICoreWebView2ContextMenuItem) MustGetShortcutKeyDescription() string {
	desc, err := i.GetShortcutKeyDescription()
	ReportErrorAuto(err)
	return desc
}

// MustGetIcon 获取菜单项图标。出错时会触发全局错误回调。
//   - 以 IStream 的形式获取菜单项的 PNG、位图或 SVG 格式的图标。
//   - 获取后要释放 IStream.
func (i *ICoreWebView2ContextMenuItem) MustGetIcon() []byte {
	icon, err := i.GetIcon()
	ReportErrorAuto(err)
	return icon
}

// MustGetKind 获取菜单项类型。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItem) MustGetKind() COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND {
	kind, err := i.GetKind()
	ReportErrorAuto(err)
	return kind
}

// MustGetIsEnabled 获取菜单项是否启用。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItem) MustGetIsEnabled() bool {
	enabled, err := i.GetIsEnabled()
	ReportErrorAuto(err)
	return enabled
}

// MustGetIsChecked 获取菜单项是否选中，在类型为“复选框”或“单选框”时使用。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItem) MustGetIsChecked() bool {
	checked, err := i.GetIsChecked()
	ReportErrorAuto(err)
	return checked
}

// MustGetChildren 获取子菜单项集合, 如果类型为子菜单。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuItem) MustGetChildren() *ICoreWebView2ContextMenuItemCollection {
	children, err := i.GetChildren()
	ReportErrorAuto(err)
	return children
}
