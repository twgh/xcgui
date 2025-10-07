package edge

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

// 返回内置的 WebView2Loader.dll 的版本号。
func getWebView2LoaderVersion() string {
	return "1.0.3537.50"
}

// 返回内置的 WebView2Helper.dll 的版本号。
func getWebView2HelperVersion() string {
	return "1.0.0.0"
}

func init() {
	// 程序运行目录没有时才写出
	if !xc.PathExists2("WebView2Loader.dll") {
		err := writeDll(WebView2Loader_Dll, "WebView2Loader")
		if err != nil {
			log.Println("写出 WebView2Loader.dll 时报错:", err.Error())
		}
	}
	if !xc.PathExists2("WebView2Helper.dll") {
		err := writeDll(WebView2Helper_Dll, "WebView2Helper")
		if err != nil {
			log.Println("写出 WebView2Helper.dll 时报错:", err.Error())
		}
	}
	loadDll()

	hr := wapi.CoInitializeEx(0, wapi.COINIT_APARTMENTTHREADED)
	if !errors.Is(hr, wapi.S_OK) {
		log.Println("Warning: CoInitializeEx call failed:", hr.Error())
	}
}

var (
	dllPath_WebView2Loader = "WebView2Loader.dll"
	dllPath_WebView2Helper = "WebView2Helper.dll"
)

// writeDll 把 WebView2Loader.dll 或 WebView2Helper.dll 写出到 windows 临时目录中 'dll名+版本号+_编译时的目标架构' 文件夹里.
//
// dll: 数据.
//
// name: dll文件名, 只能填 WebView2Loader 或 WebView2Helper。
func writeDll(dll []byte, name string) error {
	version := ""
	switch name {
	case "WebView2Loader":
		version = getWebView2LoaderVersion()
	case "WebView2Helper":
		version = getWebView2HelperVersion()
	}

	tmpDir := os.TempDir()
	tmpPath := filepath.Join(tmpDir, name+version+"_"+runtime.GOARCH)
	dllPath := filepath.Join(tmpPath, name+".dll")
	if xc.PathExists2(dllPath) { // 已存在就不写出了
		switch name {
		case "WebView2Loader":
			dllPath_WebView2Loader = dllPath
		case "WebView2Helper":
			dllPath_WebView2Helper = dllPath
		}
		return nil
	}

	err := os.Mkdir(tmpPath, 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = os.WriteFile(dllPath, dll, 0777)
	if err != nil {
		return err
	}

	switch name {
	case "WebView2Loader":
		dllPath_WebView2Loader = dllPath
	case "WebView2Helper":
		dllPath_WebView2Helper = dllPath
	}
	return nil
}

var (
	// 保证 loadDll 只运行一次.
	once = sync.Once{}

	// dll
	moduleWebView2Loader *syscall.LazyDLL
	moduleWebView2Helper *syscall.LazyDLL

	// func
	procCreateCoreWebView2EnvironmentWithOptions                *syscall.LazyProc
	procCompareBrowserVersions                                  *syscall.LazyProc
	procGetAvailableCoreWebView2BrowserVersionString            *syscall.LazyProc
	procGetAvailableCoreWebView2BrowserVersionStringWithOptions *syscall.LazyProc

	procCreateEnvironmentOptions       *syscall.LazyProc
	procCreateCustomSchemeRegistration *syscall.LazyProc
)

func loadDll() {
	once.Do(func() {
		moduleWebView2Loader = syscall.NewLazyDLL(dllPath_WebView2Loader)
		if moduleWebView2Loader.Handle() == 0 {
			log.Println("载入 WebView2Loader.dll 失败:", dllPath_WebView2Loader)
		}
		moduleWebView2Helper = syscall.NewLazyDLL(dllPath_WebView2Helper)
		if moduleWebView2Helper.Handle() == 0 {
			log.Println("载入 WebView2Helper.dll 失败:", dllPath_WebView2Helper)
		}

		procCreateCoreWebView2EnvironmentWithOptions = moduleWebView2Loader.NewProc("CreateCoreWebView2EnvironmentWithOptions")
		procCompareBrowserVersions = moduleWebView2Loader.NewProc("CompareBrowserVersions")
		procGetAvailableCoreWebView2BrowserVersionString = moduleWebView2Loader.NewProc("GetAvailableCoreWebView2BrowserVersionString")
		procGetAvailableCoreWebView2BrowserVersionStringWithOptions = moduleWebView2Loader.NewProc("GetAvailableCoreWebView2BrowserVersionStringWithOptions")

		procCreateEnvironmentOptions = moduleWebView2Helper.NewProc("CreateEnvironmentOptions")
		procCreateCustomSchemeRegistration = moduleWebView2Helper.NewProc("CreateCustomSchemeRegistration")
	})
}

// CompareBrowserVersions 将比较两个给定的版本号并返回如下结果:
//
//	-1 = v1 < v2
//	 0 = v1 == v2
//	 1 = v1 > v2
func CompareBrowserVersions(v1 string, v2 string) (int, error) {
	var result int
	r, _, _ := procCompareBrowserVersions.Call(
		common.StrPtr(v1),
		common.StrPtr(v2),
		uintptr(unsafe.Pointer(&result)))
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return result, nil
}

// GetAvailableBrowserVersion 获取本机安装或指定目录的可用的 webview2 运行时的版本号。没有可用的则返回空字符串。
//
// browserExecutableFolder: webview2 可执行文件的文件夹路径, 为空则获取本机安装的。
func GetAvailableBrowserVersion(browserExecutableFolder ...string) (string, error) {
	var _browserExecutableFolder string
	if len(browserExecutableFolder) > 0 {
		_browserExecutableFolder = browserExecutableFolder[0]
	}

	var result *uint16
	r, _, _ := procGetAvailableCoreWebView2BrowserVersionString.Call(
		common.StrPtr(_browserExecutableFolder),
		uintptr(unsafe.Pointer(&result)))
	if r != 0 {
		// HRESULT 的低16位（错误代码本身）是 ERROR_FILE_NOT_FOUND，这意味着无可匹配版本
		if r&0xFFFF == uintptr(wapi.ERROR_FILE_NOT_FOUND) {
			return "", nil // 返回一个空字符串，但不返回错误，因为我们成功检测到没有可匹配版本
		}
		return "", syscall.Errno(r)
	}
	version := common.UTF16PtrToString(result)
	wapi.CoTaskMemFree(unsafe.Pointer(result))
	return version, nil
}

// GetAvailableBrowserVersionWithOptions 获取本机安装或指定目录的可用的 webview2 运行时的版本号。没有可用的则返回空字符串。
//
// browserExecutableFolder: webview2 可执行文件的文件夹路径, 为空则获取本机安装的。
//
// environmentOptions:  WebView2 环境的配置选项。
func GetAvailableBrowserVersionWithOptions(browserExecutableFolder string, environmentOptions *ICoreWebView2EnvironmentOptions) (string, error) {
	var result *uint16
	r, _, _ := procGetAvailableCoreWebView2BrowserVersionStringWithOptions.Call(
		common.StrPtr(browserExecutableFolder),
		uintptr(unsafe.Pointer(environmentOptions)),
		uintptr(unsafe.Pointer(&result)))
	if r != 0 {
		// HRESULT 的低16位（错误代码本身）是 ERROR_FILE_NOT_FOUND，这意味着无可匹配版本
		if r&0xFFFF == uintptr(wapi.ERROR_FILE_NOT_FOUND) {
			return "", nil // 返回一个空字符串，但不返回错误，因为我们成功检测到没有可匹配版本
		}
		return "", syscall.Errno(r)
	}
	version := common.UTF16PtrToString(result)
	wapi.CoTaskMemFree(unsafe.Pointer(result))
	return version, nil
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
	r, _, _ := procCreateCoreWebView2EnvironmentWithOptions.Call(
		common.StrPtr(browserExecutableFolder),
		common.StrPtr(userDataFolder),
		uintptr(unsafe.Pointer(environmentOptions)),
		uintptr(unsafe.Pointer(environmentCompletedHandler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateEnvironmentOptions 创建 WebView2 环境选项.
func CreateEnvironmentOptions() (*ICoreWebView2EnvironmentOptions, error) {
	var opts *ICoreWebView2EnvironmentOptions
	hr, _, _ := procCreateEnvironmentOptions.Call(uintptr(unsafe.Pointer(&opts)))
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	return opts, nil
}

// CreateCustomSchemeRegistration 创建自定义方案对象。
func CreateCustomSchemeRegistration(schemeName string) (*ICoreWebView2CustomSchemeRegistration, error) {
	var reg *ICoreWebView2CustomSchemeRegistration
	hr, _, _ := procCreateCustomSchemeRegistration.Call(common.StrPtr(schemeName), uintptr(unsafe.Pointer(&reg)))
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	return reg, nil
}
