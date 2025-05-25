package edge

// ok

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2WindowFeatures 是 window.open 指定的窗口特性。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures
type ICoreWebView2WindowFeatures struct {
	Vtbl *ICoreWebView2WindowFeaturesVtbl
}

type ICoreWebView2WindowFeaturesVtbl struct {
	IUnknownVtbl
	GetHasPosition             ComProc
	GetHasSize                 ComProc
	GetLeft                    ComProc
	GetTop                     ComProc
	GetHeight                  ComProc
	GetWidth                   ComProc
	GetShouldDisplayMenuBar    ComProc
	GetShouldDisplayStatus     ComProc
	GetShouldDisplayToolbar    ComProc
	GetShouldDisplayScrollBars ComProc
}

func (i *ICoreWebView2WindowFeatures) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WindowFeatures) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WindowFeatures) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHasPosition 获取窗口是否指定了位置.
func (i *ICoreWebView2WindowFeatures) GetHasPosition() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetHasPosition.Call(
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

// MustGetHasPosition 获取窗口是否指定了位置. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetHasPosition() bool {
	value, err := i.GetHasPosition()
	ReportError2(err)
	return value
}

// GetHasSize 获取窗口是否指定了大小
func (i *ICoreWebView2WindowFeatures) GetHasSize() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetHasSize.Call(
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

// MustGetHasSize 获取窗口是否指定了大小. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetHasSize() bool {
	value, err := i.GetHasSize()
	ReportError2(err)
	return value
}

// GetLeft 获取窗口左侧位置.
func (i *ICoreWebView2WindowFeatures) GetLeft() (uint32, error) {
	var value uint32
	r, _, err := i.Vtbl.GetLeft.Call(
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

// MustGetLeft 获取窗口顶部位置. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetLeft() uint32 {
	value, err := i.GetLeft()
	ReportError2(err)
	return value
}

// GetTop 获取窗口顶部位置.
func (i *ICoreWebView2WindowFeatures) GetTop() (uint32, error) {
	var value uint32
	r, _, err := i.Vtbl.GetTop.Call(
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

// MustGetTop 获取窗口顶部位置. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetTop() uint32 {
	value, err := i.GetTop()
	ReportError2(err)
	return value
}

// GetHeight 获取窗口高度.
func (i *ICoreWebView2WindowFeatures) GetHeight() (uint32, error) {
	var value uint32
	r, _, err := i.Vtbl.GetHeight.Call(
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

// MustGetHeight 获取窗口高度. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetHeight() uint32 {
	value, err := i.GetHeight()
	ReportError2(err)
	return value
}

// GetWidth 获取窗口宽度.
func (i *ICoreWebView2WindowFeatures) GetWidth() (uint32, error) {
	var value uint32
	r, _, err := i.Vtbl.GetWidth.Call(
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

// MustGetWidth 获取窗口宽度. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetWidth() uint32 {
	value, err := i.GetWidth()
	ReportError2(err)
	return value
}

// GetShouldDisplayMenuBar 获取是否显示菜单栏.
func (i *ICoreWebView2WindowFeatures) GetShouldDisplayMenuBar() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetShouldDisplayMenuBar.Call(
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

// MustGetShouldDisplayMenuBar 获取是否显示菜单栏. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetShouldDisplayMenuBar() bool {
	value, err := i.GetShouldDisplayMenuBar()
	ReportError2(err)
	return value
}

// GetShouldDisplayStatus 获取是否显示状态栏.
func (i *ICoreWebView2WindowFeatures) GetShouldDisplayStatus() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetShouldDisplayStatus.Call(
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

// MustGetShouldDisplayStatus 获取是否显示状态栏. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetShouldDisplayStatus() bool {
	value, err := i.GetShouldDisplayStatus()
	ReportError2(err)
	return value
}

// GetShouldDisplayToolbar 获取是否显示工具栏.
func (i *ICoreWebView2WindowFeatures) GetShouldDisplayToolbar() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetShouldDisplayToolbar.Call(
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

// MustGetShouldDisplayToolbar 获取是否显示工具栏. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetShouldDisplayToolbar() bool {
	value, err := i.GetShouldDisplayToolbar()
	ReportError2(err)
	return value
}

// GetShouldDisplayScrollBars 获取是否显示滚动条.
func (i *ICoreWebView2WindowFeatures) GetShouldDisplayScrollBars() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetShouldDisplayScrollBars.Call(
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

// MustGetShouldDisplayScrollBars 获取是否显示滚动条. 出错时会触发全局错误回调.
func (i *ICoreWebView2WindowFeatures) MustGetShouldDisplayScrollBars() bool {
	value, err := i.GetShouldDisplayScrollBars()
	ReportError2(err)
	return value
}
