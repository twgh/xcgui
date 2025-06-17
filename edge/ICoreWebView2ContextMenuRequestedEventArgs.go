package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"syscall"
	"unsafe"
)

// ICoreWebView2ContextMenuRequestedEventArgs 上下文菜单请求事件的参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs
type ICoreWebView2ContextMenuRequestedEventArgs struct {
	Vtbl *ICoreWebView2ContextMenuRequestedEventArgsVtbl
}

type ICoreWebView2ContextMenuRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetMenuItems         ComProc
	GetContextMenuTarget ComProc
	GetLocation          ComProc
	PutSelectedCommandId ComProc
	GetSelectedCommandId ComProc
	PutHandled           ComProc
	GetHandled           ComProc
	GetDeferral          ComProc
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMenuItems 获取上下文菜单项集合.
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetMenuItems() (*ICoreWebView2ContextMenuItemCollection, error) {
	var items *ICoreWebView2ContextMenuItemCollection
	r, _, err := i.Vtbl.GetMenuItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&items)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return items, nil
}

// GetContextMenuTarget 获取与请求的上下文菜单关联的目标信息。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetContextMenuTarget() (*ICoreWebView2ContextMenuTarget, error) {
	var target *ICoreWebView2ContextMenuTarget
	r, _, err := i.Vtbl.GetContextMenuTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&target)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return target, nil
}

// GetLocation 获取上下文菜单请求发生的坐标，该坐标相对于 WebView 边界的左上角。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetLocation() (xc.POINT, error) {
	var point xc.POINT
	r, _, err := i.Vtbl.GetLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.POINT{}, err
	}
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// SetSelectedCommandId 设置选中的上下文菜单项的命令 ID。
//   - 设置此值后， WebView 将执行选定的命令。此值应始终通过选定的 ContextMenuItem 的 CommandId 属性获取。默认值为 -1，表示未进行选择。
//   - 应用还可以报告自定义上下文菜单项的选定命令 ID，这将导致为自定义项触发 CustomItemSelected 事件，但是，虽然在 ContextMenuRequested 事件期间每个自定义上下文菜单项的命令 ID 是唯一的，但 CoreWebView2 可能会将已删除的自定义 ContextMenuItem 的命令 ID 值重新分配给新对象，并且分配给同一自定义项的命令 ID 在每次应用运行时可能会有所不同。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) SetSelectedCommandId(id int32) error {
	r, _, err := i.Vtbl.PutSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(id),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSelectedCommandId 获取选中的命令 ID.
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetSelectedCommandId() (int32, error) {
	var id int32
	r, _, err := i.Vtbl.GetSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return id, nil
}

// SetHandled 设置是否已处理.
//   - 如果 Handled 设置为 TRUE，则 WebView 将不显示上下文菜单，而是使用 SelectedCommandId 属性来指示（如果有）要调用哪个上下文菜单项命令。
//   - 如果在事件处理程序或延期完成后， Handled 设置为 FALSE，则 WebView 将根据 MenuItems 属性的内容显示上下文菜单。
//   - 默认值为 FALSE。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) SetHandled(handled bool) error {
	r, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHandled 获取是否已处理.
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetHandled() (bool, error) {
	var handled bool
	r, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled, nil
}

// GetDeferral 获取延迟对象.
//   - 使用此操作可在自定义上下文菜单关闭时完成该事件。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// MustGetMenuItems 获取上下文菜单项集合。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetMenuItems() *ICoreWebView2ContextMenuItemCollection {
	items, err := i.GetMenuItems()
	ReportErrorAtuo(err)
	return items
}

// MustGetContextMenuTarget 获取与请求的上下文菜单关联的目标信息。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetContextMenuTarget() *ICoreWebView2ContextMenuTarget {
	target, err := i.GetContextMenuTarget()
	ReportErrorAtuo(err)
	return target
}

// MustGetLocation 获取上下文菜单请求发生的坐标，该坐标相对于 WebView 边界的左上角。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetLocation() xc.POINT {
	point, err := i.GetLocation()
	ReportErrorAtuo(err)
	return point
}

// MustGetSelectedCommandId 获取选中的命令 ID。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetSelectedCommandId() int32 {
	id, err := i.GetSelectedCommandId()
	ReportErrorAtuo(err)
	return id
}

// MustGetHandled 获取是否已处理。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetHandled() bool {
	handled, err := i.GetHandled()
	ReportErrorAtuo(err)
	return handled
}

// MustGetDeferral 获取延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}
