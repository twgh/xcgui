package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// 新窗口请求事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs
type ICoreWebView2NewWindowRequestedEventArgs struct {
	Vtbl *ICoreWebView2NewWindowRequestedEventArgsVtbl
}

type ICoreWebView2NewWindowRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetUri             ComProc
	PutNewWindow       ComProc
	GetNewWindow       ComProc
	PutHandled         ComProc
	GetHandled         ComProc
	GetIsUserInitiated ComProc
	GetDeferral        ComProc
	GetWindowFeatures  ComProc
}

func (i *ICoreWebView2NewWindowRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NewWindowRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NewWindowRequestedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUri 获取新窗口请求的 URI.
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetUri() (string, error) {
	var uri *uint16
	r, _, err := i.Vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return common.UTF16PtrToString(uri), nil
}

// MustGetUri 获取新窗口请求的 URI. 出错时会触发全局错误回调.
func (i *ICoreWebView2NewWindowRequestedEventArgs) MustGetUri() string {
	uri, err := i.GetUri()
	ReportErrorAtuo(err)
	return uri
}

// PutNewWindow 设置一个 ICoreWebView2 作为 NewWindowRequested 事件的结果。
//   - 为来自请求 Webview 内部的 window.open() 提供一个 Webview 作为目标。
//   - 如果设置了此项，此 Webview 的顶级窗口会作为已打开的 WindowProxy 返回给打开者脚本。
//   - 如果未设置此项，则会检查 Handled 以确定 NewWindowRequested 事件的行为。
//   - NewWindow 属性中提供的 CoreWebView2 必须与打开者 Webview 处于同一环境中，并且此前不应已进行过导航。
//   - 在目标内容加载完成之前，请勿对该 CoreWebView2 使用会导致导航或与 DOM 交互的方法。
//   - 设置事件处理程序、更改“Settings”属性或调用其他方法都是可行的。
//   - 应在调用 PutNewWindow 之前更改设置，以确保这些设置对新设置的 WebView 生效。
//   - 一旦设置了“NewWindow”，此 CoreWebView2 的基础网页内容将被替换，并会针对新窗口进行适当导航。
//   - 设置新窗口后无法更改，否则将返回错误。
func (i *ICoreWebView2NewWindowRequestedEventArgs) PutNewWindow(newWindow *ICoreWebView2) error {
	r, _, err := i.Vtbl.PutNewWindow.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(newWindow)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetNewWindow 获取新窗口的 ICoreWebView2 实例.
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetNewWindow() (*ICoreWebView2, error) {
	var newWindow *ICoreWebView2
	r, _, err := i.Vtbl.GetNewWindow.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&newWindow)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return newWindow, nil
}

// MustGetNewWindow 获取新窗口的 ICoreWebView2 实例. 出错时会触发全局错误回调.
func (i *ICoreWebView2NewWindowRequestedEventArgs) MustGetNewWindow() *ICoreWebView2 {
	newWindow, err := i.GetNewWindow()
	ReportErrorAtuo(err)
	return newWindow
}

// PutHandled 设置是否处理新窗口请求.
//   - 如果此值为 false 且未设置 NewWindow，WebView 将打开一个弹出窗口，并返回已打开的 WindowProxy。
//   - 如果设置为 true 且未为 NewWindow 设置 window.open，则打开的 WindowProxy 用于测试窗口对象，且不会加载任何窗口。
//   - 默认值为 false。
func (i *ICoreWebView2NewWindowRequestedEventArgs) PutHandled(handled bool) error {
	r, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHandled 获取是否处理了新窗口请求.
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetHandled() (bool, error) {
	var handled bool
	r, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled, nil
}

// MustGetHandled 获取是否处理了新窗口请求. 出错时会触发全局错误回调.
func (i *ICoreWebView2NewWindowRequestedEventArgs) MustGetHandled() bool {
	handled, err := i.GetHandled()
	ReportErrorAtuo(err)
	return handled
}

// GetIsUserInitiated 获取新窗口请求是否由用户发起.
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetIsUserInitiated() (bool, error) {
	var isUserInitiated bool
	r, _, err := i.Vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isUserInitiated)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isUserInitiated, nil
}

// MustGetIsUserInitiated 获取新窗口请求是否由用户发起. 出错时会触发全局错误回调.
func (i *ICoreWebView2NewWindowRequestedEventArgs) MustGetIsUserInitiated() bool {
	isUserInitiated, err := i.GetIsUserInitiated()
	ReportErrorAtuo(err)
	return isUserInitiated
}

// GetDeferral 获取延迟对象，并将事件置于延迟状态。
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// MustGetDeferral 获取延迟对象，并将事件置于延迟状态. 出错时会触发全局错误回调.
func (i *ICoreWebView2NewWindowRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}

// GetWindowFeatures 获取由 window.open 指定的窗口特性。
//   - 新 Webview 窗口的定位和大小调整应考虑这些特性。
func (i *ICoreWebView2NewWindowRequestedEventArgs) GetWindowFeatures() (*ICoreWebView2WindowFeatures, error) {
	var features *ICoreWebView2WindowFeatures
	r, _, err := i.Vtbl.GetWindowFeatures.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&features)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return features, nil
}
