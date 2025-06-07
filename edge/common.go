package edge

import (
	"github.com/twgh/xcgui/edge/webviewloader"
	"unsafe"
)

// CreateCoreWebView2EnvironmentWithOptions 使用指定选项创建一个 WebView2 环境。
//
// browserExecutableFolder: 浏览器可执行文件的文件夹路径, 可为空字符串。
//
// userDataFolder: 用户数据文件夹路径。
//
// environmentOptions: 包含 WebView2 环境选项的 ICoreWebView2EnvironmentOptions 对象指针。
//
// environmentCompletedHandler: 当 WebView2 环境创建完成时调用。
func CreateCoreWebView2EnvironmentWithOptions(browserExecutableFolder, userDataFolder string, environmentOptions *ICoreWebView2EnvironmentOptions, environmentCompletedHandler *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) error {
	return webviewloader.CreateCoreWebView2EnvironmentWithOptions(browserExecutableFolder, userDataFolder, unsafe.Pointer(environmentOptions), unsafe.Pointer(environmentCompletedHandler))
}

// GetAvailableBrowserVersionWithOptions 获取本机安装或指定目录的可用的 webview2 运行时的版本号。没有可用的则返回空字符串。
//
// browserExecutableFolder: webview2 可执行文件的文件夹路径, 为空则获取本机安装的。
//
// environmentOptions: WebView2 环境选项的 ICoreWebView2EnvironmentOptions 对象指针。
func GetAvailableBrowserVersionWithOptions(browserExecutableFolder string, environmentOptions *ICoreWebView2EnvironmentOptions) (string, error) {
	return webviewloader.GetAvailableBrowserVersionWithOptions(browserExecutableFolder, unsafe.Pointer(environmentOptions))
}
