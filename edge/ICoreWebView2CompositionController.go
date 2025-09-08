package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
)

// ICoreWebView2CompositionController 用于支持可视化承载 WebView.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller
type ICoreWebView2CompositionController struct {
	Vtbl *ICoreWebView2CompositionControllerVtbl
}

type ICoreWebView2CompositionControllerVtbl struct {
	IUnknownVtbl
	GetRootVisualTarget ComProc
	PutRootVisualTarget ComProc
	SendMouseInput      ComProc
	SendPointerInput    ComProc
	GetCursor           ComProc
	GetSystemCursorID   ComProc
	AddCursorChanged    ComProc
	RemoveCursorChanged ComProc
	// 2
	GetAutomationProvider ComProc
	// 3
	DragEnter ComProc
	DragLeave ComProc
	DragOver  ComProc
	Drop      ComProc
	// 4
	GetNonClientRegionAtPoint    ComProc
	QueryNonClientRegion         ComProc
	AddNonClientRegionChanged    ComProc
	RemoveNonClientRegionChanged ComProc
}

// todo: ICoreWebView2CompositionController2-4

func (i *ICoreWebView2CompositionController) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CompositionController) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CompositionController) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetRootVisualTarget 获取根视觉目标. RootVisualTarget 是宿主应用程序可视化树中的一个可视化对象。
//   - 此视觉对象是 WebView 将连接其视觉树的位置。应用使用此视觉对象在应用内定位 WebView。应用仍需要使用 Bounds 属性来调整 WebView 的大小。
//   - RootVisualTarget 属性可以是 IDCompositionVisual 或 Windows::UI::Composition::ContainerVisual。 WebView 将在从属性设置器返回之前，将其视觉树连接到提供的视觉对象。
//   - 应用需要在其设备上提交设置 RootVisualTarget 属性。 RootVisualTarget 属性支持设置为 nullptr，以断开 WebView 与应用视觉树的连接。
func (i *ICoreWebView2CompositionController) GetRootVisualTarget() (*IUnknown, error) {
	var target *IUnknown
	r, _, _ := i.Vtbl.GetRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&target)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return target, nil
}

// SetRootVisualTarget 设置根视觉目标.
func (i *ICoreWebView2CompositionController) SetRootVisualTarget(target *IUnknown) error {
	r, _, _ := i.Vtbl.PutRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(target)),
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
		uintptr(unsafe.Pointer(&point)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SendPointerInput 发送指针输入事件.
//   - 系统的任何指针输入都必须首先转换为 ICoreWebView2PointerInfo.
func (i *ICoreWebView2CompositionController) SendPointerInput(eventKind COREWEBVIEW2_POINTER_EVENT_KIND, pointerInfo *ICoreWebView2PointerInfo) error {
	r, _, _ := i.Vtbl.SendPointerInput.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(eventKind),
		uintptr(unsafe.Pointer(pointerInfo)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCursor 获取当前光标句柄.
func (i *ICoreWebView2CompositionController) GetCursor() (uintptr, error) {
	var cursor uintptr
	r, _, _ := i.Vtbl.GetCursor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cursor)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return cursor, nil
}

// GetSystemCursorID 获取当前系统光标ID。
//   - 仅当渲染引擎报告默认的 Windows 光标资源值时，systemCursorId 才有效。
//   - 否则，如果正在使用自定义 CSS 光标，则此操作将返回 0。
//   - 要在 LoadCursor 或 LoadImage 中实际使用 systemCursorId，必须首先对其调用 MAKEINTRESOURCE。
func (i *ICoreWebView2CompositionController) GetSystemCursorID() (uint32, error) {
	var systemCursorID uint32
	r, _, _ := i.Vtbl.GetSystemCursorID.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&systemCursorID)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return systemCursorID, nil
}

// Event_CursorChanged 光标改变事件.
func (i *ICoreWebView2CompositionController) Event_CursorChanged(impl *WebViewEventImpl, cb func(sender *ICoreWebView2CompositionController, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "CursorChanged", cb, i, allowAddingMultiple...)
}

// AddCursorChanged 添加光标改变事件处理程序.
func (i *ICoreWebView2CompositionController) AddCursorChanged(eventHandler *ICoreWebView2CursorChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddCursorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveCursorChanged 移除光标改变事件处理程序.
func (i *ICoreWebView2CompositionController) RemoveCursorChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveCursorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetRootVisualTarget 获取根视觉目标。RootVisualTarget 是宿主应用程序可视化树中的一个可视化对象。出错时会触发全局错误回调。
//   - 此视觉对象是 WebView 将连接其视觉树的位置。应用使用此视觉对象在应用内定位 WebView。应用仍需要使用 Bounds 属性来调整 WebView 的大小。
//   - RootVisualTarget 属性可以是 IDCompositionVisual 或 Windows::UI::Composition::ContainerVisual。 WebView 将在从属性设置器返回之前，将其视觉树连接到提供的视觉对象。
//   - 应用需要在其设备上提交设置 RootVisualTarget 属性。 RootVisualTarget 属性支持设置为 nullptr，以断开 WebView 与应用视觉树的连接。
func (i *ICoreWebView2CompositionController) MustGetRootVisualTarget() *IUnknown {
	target, err := i.GetRootVisualTarget()
	ReportErrorAuto(err)
	return target
}

// MustGetCursor 获取当前光标句柄。出错时会触发全局错误回调。
func (i *ICoreWebView2CompositionController) MustGetCursor() uintptr {
	cursor, err := i.GetCursor()
	ReportErrorAuto(err)
	return cursor
}

// MustGetSystemCursorID 获取当前系统光标ID。出错时会触发全局错误回调。
//   - 仅当渲染引擎报告默认的 Windows 光标资源值时，systemCursorId 才有效。
//   - 否则，如果正在使用自定义 CSS 光标，则此操作将返回 0。
//   - 要在 LoadCursor 或 LoadImage 中实际使用 systemCursorId，必须首先对其调用 MAKEINTRESOURCE。
func (i *ICoreWebView2CompositionController) MustGetSystemCursorID() uint32 {
	cursorID, err := i.GetSystemCursorID()
	ReportErrorAuto(err)
	return cursorID
}
