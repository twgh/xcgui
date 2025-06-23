package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings3 是 ICoreWebView2Settings2 接口的延续, 支持管理是否启用浏览器快捷键。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings3
type ICoreWebView2Settings3 struct {
	Vtbl *ICoreWebView2Settings3Vtbl
}

type ICoreWebView2Settings3Vtbl struct {
	ICoreWebView2Settings2Vtbl
	GetAreBrowserAcceleratorKeysEnabled ComProc
	PutAreBrowserAcceleratorKeysEnabled ComProc
}

func (i *ICoreWebView2Settings3) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings3) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings3) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。
func (i *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	var enabled bool
	r, _, _ := i.Vtbl.GetAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return enabled, nil
}

// SetAreBrowserAcceleratorKeysEnabled 设置是否允许浏览器快捷键。默认值为 true。此设置对 AcceleratorKeyPressed 事件没有影响。
func (i *ICoreWebView2Settings3) SetAreBrowserAcceleratorKeysEnabled(enabled bool) error {
	r, _, _ := i.Vtbl.PutAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(enabled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetAreBrowserAcceleratorKeysEnabled 获取是否允许浏览器快捷键。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings3) MustGetAreBrowserAcceleratorKeysEnabled() bool {
	enabled, err := i.GetAreBrowserAcceleratorKeysEnabled()
	ReportErrorAtuo(err)
	return enabled
}
