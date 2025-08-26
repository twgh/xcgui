package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2MoveFocusRequestedEventArgs 提供有关移动焦点请求事件的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs
type ICoreWebView2MoveFocusRequestedEventArgs struct {
	Vtbl *ICoreWebView2MoveFocusRequestedEventArgsVtbl
}

type ICoreWebView2MoveFocusRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetReason ComProc
	GetHandled ComProc
	PutHandled ComProc
}

// GetReason 获取焦点移动的原因。
func (i *ICoreWebView2MoveFocusRequestedEventArgs) GetReason() (COREWEBVIEW2_MOVE_FOCUS_REASON, error) {
	var reason COREWEBVIEW2_MOVE_FOCUS_REASON
	r, _, _ := i.Vtbl.GetReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if r != 0 {
		return reason, syscall.Errno(r)
	}
	return reason, nil
}

// GetHandled 获取事件是否已被处理。
func (i *ICoreWebView2MoveFocusRequestedEventArgs) GetHandled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return value, syscall.Errno(r)
	}
	return value, nil
}

// SetHandled 设置事件是否已被处理。
func (i *ICoreWebView2MoveFocusRequestedEventArgs) SetHandled(value bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetReason 获取焦点移动的原因。出错时会触发全局错误回调。
func (i *ICoreWebView2MoveFocusRequestedEventArgs) MustGetReason() COREWEBVIEW2_MOVE_FOCUS_REASON {
	reason, err := i.GetReason()
	ReportErrorAtuo(err)
	return reason
}

// MustGetHandled 获取事件是否已被处理。出错时会触发全局错误回调。
func (i *ICoreWebView2MoveFocusRequestedEventArgs) MustGetHandled() bool {
	value, err := i.GetHandled()
	ReportErrorAtuo(err)
	return value
}