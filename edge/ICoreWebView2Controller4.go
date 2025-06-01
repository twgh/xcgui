package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Controller4 是 ICoreWebView2Controller3 的延续, 提供用于启用/禁用外部拖放的方法。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2controller4
type ICoreWebView2Controller4 struct {
	Vtbl *ICoreWebView2Controller4Vtbl
}

type ICoreWebView2Controller4Vtbl struct {
	ICoreWebView2Controller3Vtbl
	GetAllowExternalDrop ComProc
	PutAllowExternalDrop ComProc
}

// GetAllowExternalDrop 获取是否允许外部拖放.
func (i *ICoreWebView2Controller4) GetAllowExternalDrop() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// MustGetAllowExternalDrop 获取是否允许外部拖放. 出错时会触发全局错误回调。
func (i *ICoreWebView2Controller4) MustGetAllowExternalDrop() bool {
	value, err := i.GetAllowExternalDrop()
	ReportErrorAtuo(err)
	return value
}

// PutAllowExternalDrop 设置是否允许外部拖放. 默认值为 true.
func (i *ICoreWebView2Controller4) PutAllowExternalDrop(value bool) error {
	r, _, err := i.Vtbl.PutAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
