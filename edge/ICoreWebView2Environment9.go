package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment9 继承自ICoreWebView2Environment8，提供创建上下文菜单项的功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment9
type ICoreWebView2Environment9 struct {
	ICoreWebView2Environment8
}

// CreateContextMenuItem 创建一个上下文菜单项。
//   - CoreWebView2 将在解码前倒回图标流。
//   - 在给定时间内，活动的自定义上下文菜单项限制为 1000 个。
//   - 在删除现有项之前尝试创建更多项将因 wapi.ERROR_NOT_ENOUGH_QUOTA 而失败。
//   - 为提高性能，建议在 ContextMenuRequested 事件中重复使用 ContextMenuItem。
//   - 返回的 ContextMenuItem 对象的 IsEnabled 属性默认为 TRUE，IsChecked 属性默认为 FALSE。
//   - 将为 ContextMenuItem 对象分配一个 CommandId，该 ID 在活动的自定义上下文菜单项中是唯一的，但已删除的 ContextMenuItem 的命令 ID 值可以重新分配。
//
// label: 菜单项的标签文本。
//
// iconStream: 菜单项的图标流，可以为nil。
//
// kind: 菜单项的类型。
func (i *ICoreWebView2Environment9) CreateContextMenuItem(label string, iconStream *IStream, kind COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND) (*ICoreWebView2ContextMenuItem, error) {
	_label, err := syscall.UTF16PtrFromString(label)
	if err != nil {
		return nil, err
	}
	var value *ICoreWebView2ContextMenuItem
	r, _, _ := i.Vtbl.CreateContextMenuItem.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_label)),
		uintptr(unsafe.Pointer(iconStream)),
		uintptr(kind),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustCreateContextMenuItem 创建一个上下文菜单项。
// 出错时会触发全局错误回调。
//
// label: 菜单项的标签文本。
//
// iconStream: 菜单项的图标流，可以为nil。
//
// kind: 菜单项的类型。
func (i *ICoreWebView2Environment9) MustCreateContextMenuItem(label string, iconStream *IStream, kind COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND) *ICoreWebView2ContextMenuItem {
	value, err := i.CreateContextMenuItem(label, iconStream, kind)
	ReportErrorAuto(err)
	return value
}
