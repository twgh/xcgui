package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Controller3 是 ICoreWebView2Controller2 的延续，添加了光栅化缩放和边界模式控制功能.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2controller3
type ICoreWebView2Controller3 struct {
	ICoreWebView2Controller2
}

// GetRasterizationScale 获取当前光栅化缩放比例.
func (i *ICoreWebView2Controller3) GetRasterizationScale() (float64, error) {
	var scale float64
	r, _, _ := i.Vtbl.GetRasterizationScale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scale)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return scale, nil
}

// SetRasterizationScale 设置光栅化缩放比例.
func (i *ICoreWebView2Controller3) SetRasterizationScale(scale float64) error {
	r, _, _ := i.Vtbl.PutRasterizationScale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(scale),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldDetectMonitorScaleChanges 获取是否尝试跟踪显示器DPI缩放变化.
func (i *ICoreWebView2Controller3) GetShouldDetectMonitorScaleChanges() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetShouldDetectMonitorScaleChanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetShouldDetectMonitorScaleChanges 设置是否尝试跟踪显示器DPI缩放变化.
func (i *ICoreWebView2Controller3) SetShouldDetectMonitorScaleChanges(value bool) error {
	r, _, _ := i.Vtbl.PutShouldDetectMonitorScaleChanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetBoundsMode 获取当前边界模式.
//   - BoundsMode 会影响 Bounds 和 RasterizationScale 属性的设置方式。
func (i *ICoreWebView2Controller3) GetBoundsMode() (COREWEBVIEW2_BOUNDS_MODE, error) {
	var boundsMode COREWEBVIEW2_BOUNDS_MODE
	r, _, _ := i.Vtbl.GetBoundsMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&boundsMode)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return boundsMode, nil
}

// SetBoundsMode 设置边界模式.
//   - BoundsMode 会影响 Bounds 和 RasterizationScale 属性的设置方式。
func (i *ICoreWebView2Controller3) SetBoundsMode(boundsMode COREWEBVIEW2_BOUNDS_MODE) error {
	r, _, _ := i.Vtbl.PutBoundsMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boundsMode),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddRasterizationScaleChanged 添加光栅化缩放比例改变事件处理程序.
//   - 当 Webview 检测到显示器 DPI 缩放比例已更改、ShouldDetectMonitorScaleChanges 为 true 且 Webview 已更改 RasterizationScale 属性时，将引发此事件。
func (i *ICoreWebView2Controller3) AddRasterizationScaleChanged(eventHandler *ICoreWebView2RasterizationScaleChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddRasterizationScaleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveRasterizationScaleChanged 移除光栅化缩放比例改变事件处理程序.
func (i *ICoreWebView2Controller3) RemoveRasterizationScaleChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveRasterizationScaleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetRasterizationScale 获取当前光栅化缩放比例。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller3) MustGetRasterizationScale() float64 {
	scale, err := i.GetRasterizationScale()
	ReportErrorAuto(err)
	return scale
}

// MustGetShouldDetectMonitorScaleChanges 获取是否尝试跟踪显示器DPI缩放变化。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller3) MustGetShouldDetectMonitorScaleChanges() bool {
	value, err := i.GetShouldDetectMonitorScaleChanges()
	ReportErrorAuto(err)
	return value
}

// MustGetBoundsMode 获取当前边界模式。出错时会触发全局错误回调。
//   - BoundsMode 会影响 Bounds 和 RasterizationScale 属性的设置方式。
func (i *ICoreWebView2Controller3) MustGetBoundsMode() COREWEBVIEW2_BOUNDS_MODE {
	boundsMode, err := i.GetBoundsMode()
	ReportErrorAuto(err)
	return boundsMode
}
