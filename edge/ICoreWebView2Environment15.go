package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment15 是 ICoreWebView2Environment14 接口的延续，用于创建查找选项对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment15
type ICoreWebView2Environment15 struct {
	ICoreWebView2Environment14
}

// CreateFindOptions 创建查找选项对象。
func (i *ICoreWebView2Environment15) CreateFindOptions() (*ICoreWebView2FindOptions, error) {
	var value *ICoreWebView2FindOptions
	r, _, _ := i.Vtbl.CreateFindOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustCreateFindOptions 创建查找选项对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment15) MustCreateFindOptions() *ICoreWebView2FindOptions {
	value, err := i.CreateFindOptions()
	ReportErrorAuto(err)
	return value
}
