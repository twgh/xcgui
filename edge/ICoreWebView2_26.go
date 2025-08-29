package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_26 是 ICoreWebView2_25 接口的延续，支持 开始保存文件安全检查 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_26
type ICoreWebView2_26 struct {
	ICoreWebView2_25
}

// AddSaveFileSecurityCheckStarting 添加保存文件安全检查开始事件的处理程序。
//   - 在系统 FileTypePolicy 检查危险文件扩展名列表期间，将触发此事件。
func (i *ICoreWebView2_26) AddSaveFileSecurityCheckStarting(eventHandler *ICoreWebView2SaveFileSecurityCheckStartingEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddSaveFileSecurityCheckStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveSaveFileSecurityCheckStarting 移除保存文件安全检查开始事件的处理程序。
func (i *ICoreWebView2_26) RemoveSaveFileSecurityCheckStarting(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveSaveFileSecurityCheckStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
