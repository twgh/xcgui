package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2CompositionController3 是 ICoreWebView2CompositionController2 的延续，用于管理拖放操作。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3
type ICoreWebView2CompositionController3 struct {
	ICoreWebView2CompositionController2
}

// DragLeave 处理拖拽离开事件。
//   - 此函数对应于 IDropTarget::DragLeave.
//   - 此函数依赖于 ICoreWebView2Controller 的 AllowExternalDrop 属性，如果 AllowExternalDrop 属性设置为 false，会向调用方返回 E_FAIL 以表明不允许此操作。
//   - 宿主应用程序必须注册为 IDropTarget，并实现 DragLeave 调用并将其转发给此函数。
func (i *ICoreWebView2CompositionController3) DragLeave() error {
	r, _, _ := i.Vtbl.DragLeave.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
