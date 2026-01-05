package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ControllerOptions3 是 ICoreWebView2ControllerOptions2 的延续，提供了设置默认背景颜色的功能。
type ICoreWebView2ControllerOptions3 struct {
	ICoreWebView2ControllerOptions2
}

// GetDefaultBackgroundColor 获取 WebView 的默认背景颜色。
func (i *ICoreWebView2ControllerOptions3) GetDefaultBackgroundColor() (*COREWEBVIEW2_COLOR, error) {
	var value COREWEBVIEW2_COLOR
	r, _, _ := i.Vtbl.GetDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return &value, nil
}

// SetDefaultBackgroundColor 设置 WebView 的默认背景颜色。
//   - 使用 edge.NewColor 创建颜色.
//   - 此属性允许用户提前初始化 DefaultBackgroundColor，从而防止在 WebView2 加载期间，当背景色设置为白色以外的颜色时可能出现的白色闪烁。
//   - 通过早期初始化，颜色从一开始就保持一致。初始化后，DefaultBackgroundColor 将返回使用此 API 设置的值。
//   - DefaultBackgroundColor 可通过 WEBVIEW2_DEFAULT_BACKGROUND_COLOR 环境变量进行设置，对于正在使用此解决方案的情况，该环境变量将继续得到支持。建议从环境变量过渡到使用此 API 解决方案来应用该属性。需要重点强调的是，设置环境变量后，它会覆盖 DefaultBackgroundColor，并成为 DefaultBackgroundColor 的初始值。
//   - DefaultBackgroundColor 是在所有网页内容下方渲染的颜色。这意味着当没有加载网页内容时，WebView2 会渲染此颜色。如果 WebView2 中未定义背景色，它会使用 DefaultBackgroundColor 属性来渲染背景。默认情况下，此颜色设置为白色。
//   - 此 API 仅支持不透明颜色和完全透明。对于 alpha 值不等于 0 或 255 的颜色，它会失效。当 WebView2 设置为完全透明时，它不会渲染背景，从而使其后部窗口的内容可见。
func (i *ICoreWebView2ControllerOptions3) SetDefaultBackgroundColor(backgroundColor *COREWEBVIEW2_COLOR) error {
	r, _, _ := i.Vtbl.PutDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(backgroundColor.ToUint32()),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetDefaultBackgroundColor 获取 WebView 的默认背景颜色。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions3) MustGetDefaultBackgroundColor() *COREWEBVIEW2_COLOR {
	result, err := i.GetDefaultBackgroundColor()
	ReportErrorAuto(err)
	return result
}
