package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2DevToolsProtocolEventReceiver 接收器是为特定的 DevTools 协议事件创建的，可用于订阅和取消订阅该事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceiver
type ICoreWebView2DevToolsProtocolEventReceiver struct {
	Vtbl *ICoreWebView2DevToolsProtocolEventReceiverVtbl
}

type ICoreWebView2DevToolsProtocolEventReceiverVtbl struct {
	IUnknownVtbl
	AddDevToolsProtocolEventReceived    ComProc
	RemoveDevToolsProtocolEventReceived ComProc
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddDevToolsProtocolEventReceived 添加 DevTools 协议事件接收处理程序.
func (i *ICoreWebView2DevToolsProtocolEventReceiver) AddDevToolsProtocolEventReceived(eventHandler *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddDevToolsProtocolEventReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDevToolsProtocolEventReceived 移除 DevTools 协议事件接收处理程序.
func (i *ICoreWebView2DevToolsProtocolEventReceiver) RemoveDevToolsProtocolEventReceived(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveDevToolsProtocolEventReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
