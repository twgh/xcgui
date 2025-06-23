package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_21 是 ICoreWebView2_20 接口的延续，支持执行脚本并获取字符串和异常。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_21
type ICoreWebView2_21 struct {
	Vtbl *ICoreWebView2_21Vtbl
}

type ICoreWebView2_21Vtbl struct {
	ICoreWebView2_20Vtbl
	ExecuteScriptWithResult ComProc
}

func (i *ICoreWebView2_21) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_21) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_21) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
ExecuteScriptWithResult
*/
