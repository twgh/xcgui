//go:build 386

package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xc"
)

// SetBounds 设置 WebView 的边界。
//
// bounds: WebView 的边界矩形。
func (i *ICoreWebView2Controller) SetBounds(bounds xc.RECT) error {
	r, _, _ := i.Vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(bounds.Left),
		uintptr(bounds.Top),
		uintptr(bounds.Right),
		uintptr(bounds.Bottom),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetZoomFactor 设置 WebView 的缩放系数。
func (i *ICoreWebView2Controller) SetZoomFactor(zoomFactor float64) error {
	low, high := common.Float64ToUint32Pair(zoomFactor)
	r, _, _ := i.Vtbl.PutZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(low),
		uintptr(high),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetBoundsAndZoomFactor 设置 WebView 的边界和缩放系数。
//
// bounds: WebView 的边界矩形.
//
// zoomFactor: 缩放系数.
func (i *ICoreWebView2Controller) SetBoundsAndZoomFactor(bounds xc.RECT, zoomFactor float64) error {
	low, high := common.Float64ToUint32Pair(zoomFactor)
	r, _, _ := i.Vtbl.SetBoundsAndZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(bounds.Left),
		uintptr(bounds.Top),
		uintptr(bounds.Right),
		uintptr(bounds.Bottom),
		uintptr(low),
		uintptr(high),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetPointerDeviceRect 设置指针设备矩形区域.
func (i *ICoreWebView2PointerInfo) SetPointerDeviceRect(rect xc.RECT) error {
	r, _, _ := i.Vtbl.PutPointerDeviceRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(rect.Left),
		uintptr(rect.Top),
		uintptr(rect.Right),
		uintptr(rect.Bottom),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetDisplayRect 设置显示矩形区域.
func (i *ICoreWebView2PointerInfo) SetDisplayRect(rect xc.RECT) error {
	r, _, _ := i.Vtbl.PutDisplayRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(rect.Left),
		uintptr(rect.Top),
		uintptr(rect.Right),
		uintptr(rect.Bottom),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetTouchContact 设置触摸接触区域.
func (i *ICoreWebView2PointerInfo) SetTouchContact(rect xc.RECT) error {
	r, _, _ := i.Vtbl.PutTouchContact.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(rect.Left),
		uintptr(rect.Top),
		uintptr(rect.Right),
		uintptr(rect.Bottom),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetTouchContactRaw 设置原始触摸接触区域.
func (i *ICoreWebView2PointerInfo) SetTouchContactRaw(rect xc.RECT) error {
	r, _, _ := i.Vtbl.PutTouchContactRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(rect.Left),
		uintptr(rect.Top),
		uintptr(rect.Right),
		uintptr(rect.Bottom),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetDefaultDownloadDialogMargin 设置默认下载对话框的边距.
//   - 边距是一个点，用于描述所选 WebView 角与距离它最近的默认下载对话框角之间的垂直和水平距离。
//   - 正值会使对话框从所选 WebView 角向 WebView 中心移动，负值则会使对话框远离该角。
//   - 使用(0, 0)可使对话框与 WebView 角对齐且无边距。
//   - 应在初始化期间设置角对齐方式和边距，以确保在首次计算布局时它们能正确应用，否则它们将不会生效，直到下次 WebView 位置或大小更新。
func (i *ICoreWebView2_9) SetDefaultDownloadDialogMargin(margin xc.POINT) error {
	r, _, _ := i.Vtbl.PutDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(margin.X),
		uintptr(margin.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetPixelLocation 设置像素位置.
func (i *ICoreWebView2PointerInfo) SetPixelLocation(point xc.POINT) error {
	r, _, _ := i.Vtbl.PutPixelLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(point.X),
		uintptr(point.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetHimetricLocation 设置 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) SetHimetricLocation(point xc.POINT) error {
	r, _, _ := i.Vtbl.PutHimetricLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(point.X),
		uintptr(point.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetPixelLocationRaw 设置原始像素坐标位置.
func (i *ICoreWebView2PointerInfo) SetPixelLocationRaw(point xc.POINT) error {
	r, _, _ := i.Vtbl.PutPixelLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(point.X),
		uintptr(point.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetHimetricLocationRaw 设置原始 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) SetHimetricLocationRaw(point xc.POINT) error {
	r, _, _ := i.Vtbl.PutHimetricLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(point.X),
		uintptr(point.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SendMouseInput 发送鼠标输入事件.
//
// mouseData: 指定鼠标事件的数据。
//   - 如果 eventKind 为 COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL 或 COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL，则 mouseData 指定滚轮移动量。正值表示滚轮向前旋转，远离用户；负值表示滚轮向后旋转，朝向用户。一次滚轮点击定义为 WHEEL_DELTA，即 120。
//   - 如果 eventKind 为 COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK、COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN 或 COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP，则 mouseData 指定哪些 X 按钮被按下或释放。如果第一个 X 按钮被按下/释放，此值应为 1；如果第二个 X 按钮被按下/释放，此值应为 2。
//   - 如果 eventKind 为 COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE，则 virtualKeys、mouseData 和 point 都应为零。
//   - 如果 eventKind 为任何其他值，则 mouseData 应为零。
//   - point 应位于 WebView 的客户端坐标空间中。
//   - 若要跟踪在 WebView 中开始且可能移动到 WebView 和宿主应用程序外部的鼠标事件，建议调用 SetCapture 和 ReleaseCapture。
//   - 若要关闭悬停弹出窗口，也建议发送 COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE 消息。
func (i *ICoreWebView2CompositionController) SendMouseInput(eventKind COREWEBVIEW2_MOUSE_EVENT_KIND, virtualKeys COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS, mouseData uint32, point xc.POINT) error {
	r, _, _ := i.Vtbl.SendMouseInput.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(eventKind),
		uintptr(virtualKeys),
		uintptr(mouseData),
		uintptr(point.X),
		uintptr(point.Y),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// DragEnter 处理拖拽进入事件。
//   - 此函数对应于 IDropTarget::DragEnter.
//   - 此函数依赖于 ICoreWebView2Controller 的 AllowExternalDrop 属性，如果 AllowExternalDrop 属性设置为 false，会向调用方返回 wapi.E_FAIL 以表明不允许此操作。
//   - 宿主应用程序必须注册为 IDropTarget，并实现 DragEnter 调用并将其转发给此函数。
//
// dataObject: 包含拖拽数据的 IDataObject 接口。
//
// keyState: 指示在拖拽操作期间按键的状态。
//
// point: 指定当前光标坐标。
//
// 返回值: 指示目标将对拖拽操作执行的操作。
func (i *ICoreWebView2CompositionController3) DragEnter(dataObject unsafe.Pointer, keyState uint32, point xc.POINT) (uint32, error) {
	var effect uint32
	r, _, _ := i.Vtbl.DragEnter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataObject),
		uintptr(keyState),
		uintptr(point.X),
		uintptr(point.Y),
		uintptr(unsafe.Pointer(&effect)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return effect, nil
}

// DragOver 处理拖拽悬停事件。
//   - 此函数对应于 IDropTarget::DragOver.
//   - 此函数依赖于 ICoreWebView2Controller 的 AllowExternalDrop 属性，如果 AllowExternalDrop 属性设置为 false，会向调用方返回 wapi.E_FAIL 以表明不允许此操作。
//   - 宿主应用程序必须注册为 IDropTarget，并实现 DragOver 调用并将其转发给此函数。
//
// keyState: 指示在拖拽操作期间按键的状态。
//
// point: 指定当前光标坐标。
//
// 返回值: 指示目标将对拖拽操作执行的操作。
func (i *ICoreWebView2CompositionController3) DragOver(keyState uint32, point xc.POINT) (uint32, error) {
	var effect uint32
	r, _, _ := i.Vtbl.DragOver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(keyState),
		uintptr(point.X),
		uintptr(point.Y),
		uintptr(unsafe.Pointer(&effect)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return effect, nil
}

// Drop 处理拖拽放置事件。
//   - 此函数对应于 IDropTarget::Drop.
//   - 此函数依赖于 ICoreWebView2Controller 的 AllowExternalDrop 属性，如果 AllowExternalDrop 属性设置为 false，会向调用方返回 wapi.E_FAIL 以表明不允许此操作。
//   - 宿主应用程序必须注册为 IDropTarget，并实现拖放调用并将其转发给此函数。
//
// dataObject: 包含拖拽数据的 IDataObject 接口。
//
// keyState: 指示在拖拽操作期间按键的状态。
//
// point: 指定当前光标坐标。
//
// 返回值: 指示目标将对拖拽操作执行的操作。
func (i *ICoreWebView2CompositionController3) Drop(dataObject unsafe.Pointer, keyState uint32, point xc.POINT) (uint32, error) {
	var effect uint32
	r, _, _ := i.Vtbl.Drop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataObject),
		uintptr(keyState),
		uintptr(point.X),
		uintptr(point.Y),
		uintptr(unsafe.Pointer(&effect)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return effect, nil
}

// GetNonClientRegionAtPoint 获取指定点的非客户端区域类型。
//   - 如果您使用 ICoreWebView2CompositionController 托管 WebView2，可以在 Win32 窗口过程（WndProc）中调用此方法，以确定鼠标是否正在移动到 WebView2 网页内容上或在其上点击，而这些内容应被视为非客户区的一部分。
//
// point: 指定要检查的点坐标。
func (i *ICoreWebView2CompositionController4) GetNonClientRegionAtPoint(point xc.POINT) (COREWEBVIEW2_NON_CLIENT_REGION_KIND, error) {
	var value COREWEBVIEW2_NON_CLIENT_REGION_KIND
	r, _, _ := i.Vtbl.GetNonClientRegionAtPoint.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(point.X),
		uintptr(point.Y),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}
