package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_28 是 ICoreWebView2_27 接口的延续，支持查找功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_28
type ICoreWebView2_28 struct {
	ICoreWebView2_27
}

// GetFind 获取当前网页视图的查找会话接口。
func (i *ICoreWebView2_28) GetFind() (*ICoreWebView2Find, error) {
	var value *ICoreWebView2Find
	r, _, _ := i.Vtbl.GetFind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetFind 获取当前网页视图的查找会话接口。出错时会触发全局错误回调。
func (i *ICoreWebView2_28) MustGetFind() *ICoreWebView2Find {
	value, err := i.GetFind()
	ReportErrorAuto(err)
	return value
}
