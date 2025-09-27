package edge

import (
	"unsafe"

	"github.com/twgh/xcgui/edge/webviewloader"
)

// CreateEnvironmentOptions 创建 WebView2 环境选项.
func CreateEnvironmentOptions() (*ICoreWebView2EnvironmentOptions, error) {
	var opts *ICoreWebView2EnvironmentOptions
	err := webviewloader.CreateEnvironmentOptions(unsafe.Pointer(&opts))
	return opts, err
}

// CreateCustomSchemeRegistration 创建自定义方案对象。
func CreateCustomSchemeRegistration(schemeName string) (*ICoreWebView2CustomSchemeRegistration, error) {
	var reg *ICoreWebView2CustomSchemeRegistration
	err := webviewloader.CreateCustomSchemeRegistration(schemeName, unsafe.Pointer(&reg))
	return reg, err
}
