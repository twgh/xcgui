package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_11 是 ICoreWebView2_10 接口的延续，支持 CDP 方法调用的 sessionId 和 已请求上下文菜单 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_11
type ICoreWebView2_11 struct {
	Vtbl *ICoreWebView2_11Vtbl
}

type ICoreWebView2_11Vtbl struct {
	ICoreWebView2_10Vtbl
	CallDevToolsProtocolMethodForSession ComProc
	AddContextMenuRequested              ComProc
	RemoveContextMenuRequested           ComProc
}

func (i *ICoreWebView2_11) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_11) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_11) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CallDevToolsProtocolMethodForSession 针对已连接目标的特定会话，运行异步的 DevToolsProtocol 方法。
//
// sessionId: DevTools 协议会话 ID。
//
// methodName: DevTools 协议方法的完整名称，格式为 {domain}.{method}。
//
// parametersAsJson: JSON 字符串，其中包含相应方法的参数。
//
// handler: 执行完成后的回调处理程序，接收返回的 JSON 结果。
func (i *ICoreWebView2_11) CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) error {
	_sessionId, err := syscall.UTF16PtrFromString(sessionId)
	if err != nil {
		return err
	}
	_methodName, err := syscall.UTF16PtrFromString(methodName)
	if err != nil {
		return err
	}
	_parametersAsJson, err := syscall.UTF16PtrFromString(parametersAsJson)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.CallDevToolsProtocolMethodForSession.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_sessionId)),
		uintptr(unsafe.Pointer(_methodName)),
		uintptr(unsafe.Pointer(_parametersAsJson)),
		uintptr(unsafe.Pointer(handler)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddContextMenuRequested
RemoveContextMenuRequested
*/
