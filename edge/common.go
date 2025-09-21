package edge

import (
	"math"
	"unsafe"

	"github.com/twgh/xcgui/edge/webviewloader"
	"github.com/twgh/xcgui/xc"
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

// COREWEBVIEW2_COLOR 表示 WebView2 的 RGBA 颜色值。
type COREWEBVIEW2_COLOR struct {
	A uint8
	R uint8
	G uint8
	B uint8
}

func NewColor(r, g, b, a uint8) *COREWEBVIEW2_COLOR {
	return &COREWEBVIEW2_COLOR{R: r, G: g, B: b, A: a}
}

// ToUint32 将其转换为 uint32, 这是 COM 方法底层所期望的.
func (c *COREWEBVIEW2_COLOR) ToUint32() uint32 {
	return *(*uint32)(unsafe.Pointer(c))
}

// Float64ToUint32Pair 将一个 64 位浮点数转换为两个 32 位整数。用于传递 float64 给 COM 方法。
func Float64ToUint32Pair(f float64) (low, high uint32) {
	// 将浮点数转换为 uint64（保持位模式）
	bits := math.Float64bits(f)
	// 在 32 位下，浮点数被拆分为两个 32 位整数
	low = uint32(bits)        // 浮点数的低32位
	high = uint32(bits >> 32) // 浮点数的高32位
	return
}

// PointToUintptr 将一个 POINT 结构强制转换为 uintptr。用于传递 POINT 结构给 COM 方法。
func PointToUintptr(pt xc.POINT) uintptr {
	return *(*uintptr)(unsafe.Pointer(&pt))
}
