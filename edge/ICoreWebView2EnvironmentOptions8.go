package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions8 提供用于创建 WebView2 环境以管理滚动条样式的附加选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8
type ICoreWebView2EnvironmentOptions8 struct {
	Vtbl *ICoreWebView2EnvironmentOptions8Vtbl
}

type ICoreWebView2EnvironmentOptions8Vtbl struct {
	IUnknownVtbl
	GetScrollBarStyle ComProc
	PutScrollBarStyle ComProc
}

func (i *ICoreWebView2EnvironmentOptions8) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions8) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions8) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetScrollBarStyle 获取滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) GetScrollBarStyle() (COREWEBVIEW2_SCROLLBAR_STYLE, error) {
	var value COREWEBVIEW2_SCROLLBAR_STYLE
	r, _, err := i.Vtbl.GetScrollBarStyle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetScrollBarStyle 获取滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) MustGetScrollBarStyle() COREWEBVIEW2_SCROLLBAR_STYLE {
	value, err := i.GetScrollBarStyle()
	ReportErrorAtuo(err)
	return value
}

// PutScrollBarStyle 设置滚动条样式.
func (i *ICoreWebView2EnvironmentOptions8) PutScrollBarStyle(value COREWEBVIEW2_SCROLLBAR_STYLE) error {
	r, _, err := i.Vtbl.PutScrollBarStyle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
