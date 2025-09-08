package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2Notification 表示一个 HTML 通知对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2notification
type ICoreWebView2Notification struct {
	Vtbl *_ICoreWebView2NotificationVtbl
}

type _ICoreWebView2NotificationVtbl struct {
	IUnknownVtbl
	AddCloseRequested      ComProc
	RemoveCloseRequested   ComProc
	ReportShown            ComProc
	ReportClicked          ComProc
	ReportClosed           ComProc
	GetBody                ComProc
	GetDirection           ComProc
	GetLanguage            ComProc
	GetTag                 ComProc
	GetIconUri             ComProc
	GetTitle               ComProc
	GetBadgeUri            ComProc
	GetBodyImageUri        ComProc
	GetShouldRenotify      ComProc
	GetRequiresInteraction ComProc
	GetIsSilent            ComProc
	GetTimestamp           ComProc
	GetVibrationPattern    ComProc
}

func (i *ICoreWebView2Notification) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Notification) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Notification) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Event_NotificationCloseRequested 通知关闭请求事件。
//   - 当通知被网页代码关闭时（例如通过 notification.close()）触发。由于这是来自网页代码的操作，因此无需调用 ReportClosed。
func (i *ICoreWebView2Notification) Event_NotificationCloseRequested(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Notification, args *IUnknown) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "NotificationCloseRequested", cb, i, allowAddingMultiple...)
}

// AddCloseRequested 添加通知关闭请求事件处理程序。
//   - 当通知被网页代码关闭时（例如通过 notification.close()）触发。由于这是来自网页代码的操作，因此无需调用 ReportClosed。
func (i *ICoreWebView2Notification) AddCloseRequested(eventHandler *ICoreWebView2NotificationCloseRequestedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveCloseRequested 移除通知关闭请求事件处理程序。
func (i *ICoreWebView2Notification) RemoveCloseRequested(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ReportShown 宿主可以运行此操作来报告通知已显示，这将导致非持久性通知触发显示事件。
//   - 除非你正在处理 NotificationReceived 事件，否则不得运行此操作。
//   - 如果调用此操作时 Handled 为 FALSE，则返回 ERROR_INVALID_STATE。
func (i *ICoreWebView2Notification) ReportShown() error {
	r, _, _ := i.Vtbl.ReportShown.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ReportClicked 宿主可以运行此操作来报告通知已被点击，这将为非持久性通知触发点击事件，为持久性通知触发 notificationclick 事件。
//   - 使用 ReportClickedWithActionIndex 指定激活持久通知的操作。
//   - 只有在处理 NotificationReceived 事件时，才能运行此操作。
//   - 如果调用此操作时 Handled 为 FALSE 或 ReportShown 尚未运行，将返回 ERROR_INVALID_STATE。
func (i *ICoreWebView2Notification) ReportClicked() error {
	r, _, _ := i.Vtbl.ReportClicked.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ReportClosed 主机可以运行此操作来报告通知已被关闭，这将为非持久性通知触发关闭事件，为持久性通知触发通知关闭事件。
//   - 除非你正在处理 NotificationReceived 事件，否则不得运行此操作。
//   - 如果调用此操作时 Handled 为 FALSE 或 ReportShown 尚未运行，将返回 ERROR_INVALID_STATE。
func (i *ICoreWebView2Notification) ReportClosed() error {
	r, _, _ := i.Vtbl.ReportClosed.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetBody 获取通知的正文内容。
func (i *ICoreWebView2Notification) GetBody() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetBody.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetDirection 获取通知的文本方向。
func (i *ICoreWebView2Notification) GetDirection() (COREWEBVIEW2_TEXT_DIRECTION_KIND, error) {
	var value COREWEBVIEW2_TEXT_DIRECTION_KIND
	r, _, _ := i.Vtbl.GetDirection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetLanguage 获取通知的语言。
//   - 请注意，不会对此属性执行任何验证，它可以是通知发送者指定的任何字符串。这对应于 Notification.lang DOM API。默认值是空字符串。
func (i *ICoreWebView2Notification) GetLanguage() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetTag 获取通知的标签。
//   - 这对应于 Notification.tag DOM API。默认值是空字符串。
func (i *ICoreWebView2Notification) GetTag() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetTag.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetIconUri 获取通知的图标 URI。
func (i *ICoreWebView2Notification) GetIconUri() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetIconUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetTitle 获取通知的标题。
func (i *ICoreWebView2Notification) GetTitle() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetBadgeUri 获取通知的徽章 URI。
//   - 当没有足够空间显示通知本身时，一个包含用于表示该通知的图像 URI 的字符串。
func (i *ICoreWebView2Notification) GetBadgeUri() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetBadgeUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetBodyImageUri 获取通知正文图像的 URI。
func (i *ICoreWebView2Notification) GetBodyImageUri() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetBodyImageUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetShouldRenotify 获取在新通知替换旧通知后是否应通知用户。
//   - 这对应于 Notification.renotify DOM API。默认值为FALSE。
func (i *ICoreWebView2Notification) GetShouldRenotify() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetShouldRenotify.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetRequiresInteraction 获取通知是否需要用户交互。
//   - 指示通知应保持活跃状态，直到用户点击或关闭它，而不是自动关闭。
//   - 这对应于 Notification.requireInteraction DOM API。请注意，由于原生API的限制，您可能不一定能够实现这一点。默认值为 FALSE。
func (i *ICoreWebView2Notification) GetRequiresInteraction() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetRequiresInteraction.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetIsSilent 获取通知是否为静默通知。
//   - 指示通知是否应设为静音——即无论设备设置如何，都不应发出声音或振动。
//   - 这对应于 Notification.silent DOM API。默认值为 FALSE。
func (i *ICoreWebView2Notification) GetIsSilent() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetIsSilent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetTimestamp 获取通知创建或适用的时间（过去、现在或未来）的时间戳，以自 UNIX 纪元以来的毫秒数计。
func (i *ICoreWebView2Notification) GetTimestamp() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetTimestamp.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetVibrationPattern 获取通知的振动模式。
//   - 这对应于 Notification.vibrate DOM API。如果未指定任何振动模式，则返回一个空数组。
func (i *ICoreWebView2Notification) GetVibrationPattern() ([]uint64, error) {
	var count uint32
	var pattern *uint64
	r, _, _ := i.Vtbl.GetVibrationPattern.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&pattern)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(pattern))

	if count == 0 || pattern == nil {
		return nil, nil
	}

	// 将 C 数组转换为 Go 切片
	result := make([]uint64, count)
	slice := unsafe.Slice(pattern, count)
	copy(result, slice)
	return result, nil
}

// MustGetBody 获取通知的正文内容。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetBody() string {
	value, err := i.GetBody()
	ReportErrorAuto(err)
	return value
}

// MustGetDirection 获取通知的文本方向。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetDirection() COREWEBVIEW2_TEXT_DIRECTION_KIND {
	value, err := i.GetDirection()
	ReportErrorAuto(err)
	return value
}

// MustGetLanguage 获取通知的语言。出错时会触发全局错误回调。
//   - 请注意，不会对此属性执行任何验证，它可以是通知发送者指定的任何字符串。这对应于 Notification.lang DOM API。默认值是空字符串。
func (i *ICoreWebView2Notification) MustGetLanguage() string {
	value, err := i.GetLanguage()
	ReportErrorAuto(err)
	return value
}

// MustGetTag 获取通知的标签。出错时会触发全局错误回调。
//   - 这对应于 Notification.tag DOM API。默认值是空字符串。
func (i *ICoreWebView2Notification) MustGetTag() string {
	value, err := i.GetTag()
	ReportErrorAuto(err)
	return value
}

// MustGetIconUri 获取通知的图标 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetIconUri() string {
	value, err := i.GetIconUri()
	ReportErrorAuto(err)
	return value
}

// MustGetTitle 获取通知的标题。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetTitle() string {
	value, err := i.GetTitle()
	ReportErrorAuto(err)
	return value
}

// MustGetBadgeUri 获取通知的徽章 URI。出错时会触发全局错误回调。
//   - 当没有足够空间显示通知本身时，一个包含用于表示该通知的图像 URI 的字符串。
func (i *ICoreWebView2Notification) MustGetBadgeUri() string {
	value, err := i.GetBadgeUri()
	ReportErrorAuto(err)
	return value
}

// MustGetBodyImageUri 获取通知正文图像的 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetBodyImageUri() string {
	value, err := i.GetBodyImageUri()
	ReportErrorAuto(err)
	return value
}

// MustGetShouldRenotify 获取在新通知替换旧通知后是否应通知用户。出错时会触发全局错误回调。
//   - 这对应于 Notification.renotify DOM API。默认值为FALSE。
func (i *ICoreWebView2Notification) MustGetShouldRenotify() bool {
	value, err := i.GetShouldRenotify()
	ReportErrorAuto(err)
	return value
}

// MustGetRequiresInteraction 获取通知是否需要用户交互。出错时会触发全局错误回调。
//   - 指示通知应保持活跃状态，直到用户点击或关闭它，而不是自动关闭。
//   - 这对应于 Notification.requireInteraction DOM API。请注意，由于原生API的限制，您可能不一定能够实现这一点。默认值为 FALSE。
func (i *ICoreWebView2Notification) MustGetRequiresInteraction() bool {
	value, err := i.GetRequiresInteraction()
	ReportErrorAuto(err)
	return value
}

// MustGetIsSilent 获取通知是否为静默通知。出错时会触发全局错误回调。
//   - 指示通知是否应设为静音——即无论设备设置如何，都不应发出声音或振动。
//   - 这对应于 Notification.silent DOM API。默认值为 FALSE。
func (i *ICoreWebView2Notification) MustGetIsSilent() bool {
	value, err := i.GetIsSilent()
	ReportErrorAuto(err)
	return value
}

// MustGetTimestamp 获取通知创建或适用的时间（过去、现在或未来）的时间戳，以自 UNIX 纪元以来的毫秒数计。出错时会触发全局错误回调。
func (i *ICoreWebView2Notification) MustGetTimestamp() float64 {
	value, err := i.GetTimestamp()
	ReportErrorAuto(err)
	return value
}

// MustGetVibrationPattern 获取通知的振动模式。出错时会触发全局错误回调。
//   - 这对应于 Notification.vibrate DOM API。如果未指定任何振动模式，则返回一个空数组。
func (i *ICoreWebView2Notification) MustGetVibrationPattern() []uint64 {
	value, err := i.GetVibrationPattern()
	ReportErrorAuto(err)
	return value
}
