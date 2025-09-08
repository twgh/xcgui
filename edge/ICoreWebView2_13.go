package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_13 是 ICoreWebView2_12 接口的延续，支持获取关联的 ICoreWebView2Profile 对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_13
type ICoreWebView2_13 struct {
	ICoreWebView2_12
}

// GetProfile 获取与此 WebView2 关联的配置文件对象。
//   - 如果此 ICoreWebView2 是使用 ICoreWebView2ControllerOptions 创建的，那么 ICoreWebView2Profile 将与指定的选项匹配。
//   - 否则，如果此 ICoreWebView2 是在没有 ICoreWebView2ControllerOptions 的情况下创建的，那么它将是相应 ICoreWebView2Environment 的默认 ICoreWebView2Profile。
func (i *ICoreWebView2_13) GetProfile() (*ICoreWebView2Profile, error) {
	var _profile *ICoreWebView2Profile
	r, _, _ := i.Vtbl.GetProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_profile)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return _profile, nil
}

// MustGetProfile 获取与此 WebView2 关联的配置文件对象。出错时会触发全局错误回调。
//   - 如果此 ICoreWebView2 是使用 ICoreWebView2ControllerOptions 创建的，那么 ICoreWebView2Profile 将与指定的选项匹配。
//   - 否则，如果此 ICoreWebView2 是在没有 ICoreWebView2ControllerOptions 的情况下创建的，那么它将是相应 ICoreWebView2Environment 的默认 ICoreWebView2Profile。
func (i *ICoreWebView2_13) MustGetProfile() *ICoreWebView2Profile {
	result, err := i.GetProfile()
	ReportErrorAuto(err)
	return result
}
