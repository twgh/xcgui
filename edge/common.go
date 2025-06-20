package edge

import (
	"github.com/twgh/xcgui/edge/webviewloader"
	"unsafe"
)

// CompareBrowserVersions 将比较两个给定的版本号并返回如下结果:
//
//	-1 = v1 < v2
//	 0 = v1 == v2
//	 1 = v1 > v2
func CompareBrowserVersions(v1 string, v2 string) (int, error) {
	return webviewloader.CompareBrowserVersions(v1, v2)
}

// GetAvailableBrowserVersion 获取本机安装或指定目录的可用的 webview2 运行时的版本号。没有可用的则返回空字符串。
//
// browserExecutableFolder: webview2 可执行文件的文件夹路径, 为空则获取本机安装的。
func GetAvailableBrowserVersion(browserExecutableFolder ...string) (string, error) {
	return webviewloader.GetAvailableBrowserVersion(browserExecutableFolder...)
}

// CreateCoreWebView2EnvironmentWithOptions 使用指定选项创建一个 WebView2 环境。
//
// browserExecutableFolder: 浏览器可执行文件的文件夹路径, 可为空字符串。如包含\Edge\Application\则失败.
//
// userDataFolder: 用户数据文件夹路径。
//
// environmentOptions: WebView2 环境选项。
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
