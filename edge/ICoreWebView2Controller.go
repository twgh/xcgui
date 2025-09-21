package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

// ICoreWebView2Controller 是 CoreWebView2 对象的所有者，该对象支持调整大小、显示和隐藏、聚焦以及与窗口和合成相关的其他功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2controller
type ICoreWebView2Controller struct {
	Vtbl *ICoreWebView2ControllerVtbl
}

type ICoreWebView2ControllerVtbl struct {
	IUnknownVtbl
	GetIsVisible                      ComProc
	PutIsVisible                      ComProc
	GetBounds                         ComProc
	PutBounds                         ComProc
	GetZoomFactor                     ComProc
	PutZoomFactor                     ComProc
	AddZoomFactorChanged              ComProc
	RemoveZoomFactorChanged           ComProc
	SetBoundsAndZoomFactor            ComProc
	MoveFocus                         ComProc
	AddMoveFocusRequested             ComProc
	RemoveMoveFocusRequested          ComProc
	AddGotFocus                       ComProc
	RemoveGotFocus                    ComProc
	AddLostFocus                      ComProc
	RemoveLostFocus                   ComProc
	AddAcceleratorKeyPressed          ComProc
	RemoveAcceleratorKeyPressed       ComProc
	GetParentWindow                   ComProc
	PutParentWindow                   ComProc
	NotifyParentWindowPositionChanged ComProc
	Close                             ComProc
	GetCoreWebView2                   ComProc
	// 2
	GetDefaultBackgroundColor ComProc
	PutDefaultBackgroundColor ComProc
	// 3
	GetRasterizationScale              ComProc
	PutRasterizationScale              ComProc
	GetShouldDetectMonitorScaleChanges ComProc
	PutShouldDetectMonitorScaleChanges ComProc
	AddRasterizationScaleChanged       ComProc
	RemoveRasterizationScaleChanged    ComProc
	GetBoundsMode                      ComProc
	PutBoundsMode                      ComProc
	// 4
	GetAllowExternalDrop ComProc
	PutAllowExternalDrop ComProc
}

func (i *ICoreWebView2Controller) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Controller) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Controller) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetBounds 获取 WebView 的边界。
func (i *ICoreWebView2Controller) GetBounds() (xc.RECT, error) {
	var bounds xc.RECT
	r, _, _ := i.Vtbl.GetBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if r != 0 {
		return bounds, syscall.Errno(r)
	}
	return bounds, nil
}

// AddAcceleratorKeyPressed 添加键盘快捷键事件处理程序.
//   - AcceleratorKeyPressed 在 Webview 获得焦点时，当按下或释放快捷键或组合键时运行。
func (i *ICoreWebView2Controller) AddAcceleratorKeyPressed(eventHandler *ICoreWebView2AcceleratorKeyPressedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveAcceleratorKeyPressed 移除键盘快捷键事件处理程序.
func (i *ICoreWebView2Controller) RemoveAcceleratorKeyPressed(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsVisible 获取 WebView 是否可见。
func (i *ICoreWebView2Controller) GetIsVisible() (bool, error) {
	var isVisible bool
	r, _, _ := i.Vtbl.GetIsVisible.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isVisible)),
	)
	if r != 0 {
		return isVisible, syscall.Errno(r)
	}
	return isVisible, nil
}

// SetIsVisible 设置 WebView 是否可见。
func (i *ICoreWebView2Controller) SetIsVisible(isVisible bool) error {
	r, _, _ := i.Vtbl.PutIsVisible.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isVisible),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// NotifyParentWindowPositionChanged 通知父窗口位置已更改。
func (i *ICoreWebView2Controller) NotifyParentWindowPositionChanged() error {
	r, _, _ := i.Vtbl.NotifyParentWindowPositionChanged.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MoveFocus 将焦点移动到 WebView。
//
// reason: 用于确定焦点移动原因的 COREWEBVIEW2_MOVE_FOCUS_REASON 常量值。
func (i *ICoreWebView2Controller) MoveFocus(reason COREWEBVIEW2_MOVE_FOCUS_REASON) error {
	r, _, _ := i.Vtbl.MoveFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(reason),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCoreWebView2 获取 WebView2 对象。
func (i *ICoreWebView2Controller) GetCoreWebView2() (*ICoreWebView2, error) {
	var webView *ICoreWebView2
	r, _, _ := i.Vtbl.GetCoreWebView2.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&webView)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return webView, nil
}

// GetZoomFactor 获取 WebView 的缩放系数。
func (i *ICoreWebView2Controller) GetZoomFactor() (float64, error) {
	var zoomFactor float64
	r, _, _ := i.Vtbl.GetZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&zoomFactor)),
	)
	if r != 0 {
		return zoomFactor, syscall.Errno(r)
	}
	return zoomFactor, nil
}

// Close 关闭 WebView 并清理底层浏览器实例。
func (i *ICoreWebView2Controller) Close() error {
	r, _, _ := i.Vtbl.Close.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetParentWindow 获取 WebView 的父窗口句柄。
func (i *ICoreWebView2Controller) GetParentWindow() (uintptr, error) {
	var parentWindow uintptr
	r, _, _ := i.Vtbl.GetParentWindow.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
	)
	if r != 0 {
		return parentWindow, syscall.Errno(r)
	}
	return parentWindow, nil
}

// SetParentWindow 设置 WebView 的父窗口句柄。
func (i *ICoreWebView2Controller) SetParentWindow(parentWindow uintptr) error {
	r, _, _ := i.Vtbl.PutParentWindow.Call(
		uintptr(unsafe.Pointer(i)),
		parentWindow,
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddZoomFactorChanged 添加缩放因子变化事件处理程序。
//   - 当 WebView 的 ZoomFactor 属性发生变化时触发。
//   - WebView 会关联每个网站最后使用的缩放比例。导航到不同页面时，缩放比例可能会发生变化。当因导航变化导致缩放比例改变时，ZoomFactorChanged 事件会在 ContentLoading 事件之后立即运行。
func (i *ICoreWebView2Controller) AddZoomFactorChanged(eventHandler *ICoreWebView2ZoomFactorChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddZoomFactorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveZoomFactorChanged 移除缩放因子变化事件处理程序。
func (i *ICoreWebView2Controller) RemoveZoomFactorChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveZoomFactorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddMoveFocusRequested 添加移动焦点请求事件处理程序。
//   - MoveFocusRequested 在用户尝试按 Tab 键离开 WebView 时运行。此事件运行时，WebView 的焦点尚未改变。
func (i *ICoreWebView2Controller) AddMoveFocusRequested(eventHandler *ICoreWebView2MoveFocusRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddMoveFocusRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveMoveFocusRequested 移除移动焦点请求事件处理程序。
func (i *ICoreWebView2Controller) RemoveMoveFocusRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveMoveFocusRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddGotFocus 添加获得焦点事件处理程序。
//   - 当 WebView 获得焦点时触发。
func (i *ICoreWebView2Controller) AddGotFocus(eventHandler *ICoreWebView2FocusChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddGotFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveGotFocus 移除获得焦点事件处理程序。
func (i *ICoreWebView2Controller) RemoveGotFocus(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveGotFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddLostFocus 添加失去焦点事件处理程序。
//   - 当 WebView 失去焦点时触发。
func (i *ICoreWebView2Controller) AddLostFocus(eventHandler *ICoreWebView2LostFocusEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddLostFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveLostFocus 移除失去焦点事件处理程序。
func (i *ICoreWebView2Controller) RemoveLostFocus(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveLostFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2Controller2 获取 ICoreWebView2Controller2 对象。
func (i *ICoreWebView2Controller) GetICoreWebView2Controller2() (*ICoreWebView2Controller2, error) {
	var result *ICoreWebView2Controller2
	iidICoreWebView2Controller2 := wapi.NewGUID(IID_ICoreWebView2Controller2)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Controller2),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Controller3 获取 ICoreWebView2Controller3 对象。
func (i *ICoreWebView2Controller) GetICoreWebView2Controller3() (*ICoreWebView2Controller3, error) {
	var result *ICoreWebView2Controller3
	iidICoreWebView2Controller3 := wapi.NewGUID(IID_ICoreWebView2Controller3)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Controller3),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Controller4 获取 ICoreWebView2Controller4 对象。
func (i *ICoreWebView2Controller) GetICoreWebView2Controller4() (*ICoreWebView2Controller4, error) {
	var result *ICoreWebView2Controller4
	iidICoreWebView2Controller4 := wapi.NewGUID(IID_ICoreWebView2Controller4)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2Controller4),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetBounds 获取 WebView 的边界。出错时会触发全局错误回调.
func (i *ICoreWebView2Controller) MustGetBounds() xc.RECT {
	b, err := i.GetBounds()
	ReportErrorAuto(err)
	return b
}

// MustGetIsVisible 获取 WebView 是否可见。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetIsVisible() bool {
	v, err := i.GetIsVisible()
	ReportErrorAuto(err)
	return v
}

// MustGetCoreWebView2 获取 WebView2 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetCoreWebView2() *ICoreWebView2 {
	w, err := i.GetCoreWebView2()
	ReportErrorAuto(err)
	return w
}

// MustGetZoomFactor 获取 WebView 的缩放系数。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetZoomFactor() float64 {
	z, err := i.GetZoomFactor()
	ReportErrorAuto(err)
	return z
}

// MustGetParentWindow 获取 WebView 的父窗口句柄。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetParentWindow() uintptr {
	h, err := i.GetParentWindow()
	ReportErrorAuto(err)
	return h
}

// MustGetICoreWebView2Controller2 获取 ICoreWebView2Controller2 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetICoreWebView2Controller2() *ICoreWebView2Controller2 {
	c, err := i.GetICoreWebView2Controller2()
	ReportErrorAuto(err)
	return c
}

// MustGetICoreWebView2Controller3 获取 ICoreWebView2Controller3 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetICoreWebView2Controller3() *ICoreWebView2Controller3 {
	c, err := i.GetICoreWebView2Controller3()
	ReportErrorAuto(err)
	return c
}

// MustGetICoreWebView2Controller4 获取 ICoreWebView2Controller4 对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Controller) MustGetICoreWebView2Controller4() *ICoreWebView2Controller4 {
	c, err := i.GetICoreWebView2Controller4()
	ReportErrorAuto(err)
	return c
}
