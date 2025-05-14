package webviewloader

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func init() {
	if !xc.PathExists2("WebView2Loader.dll") {
		err := writeDll(Dll)
		if err != nil {
			panic("写出 WebView2Loader.dll 时报错: " + err.Error())
		}
	}
	loadDll()

	hr := wapi.CoInitializeEx(0, wapi.COINIT_APARTMENTTHREADED)
	if hr < 0 {
		log.Printf("Warning: CoInitializeEx call failed: 0x%08x", hr)
	}
}

var dllPath = "WebView2Loader.dll"

// GetVersion 返回内置的 WebView2Loader.dll 的版本号。
func GetVersion() string {
	return "1.0.3240.44"
}

// writeDll 把 WebView2Loader.dll 写出到 windows 临时目录中 'WebView2Loader+版本号+_编译时的目标架构' 文件夹里.
//
// 使用完本函数后无需再调用 SetDllPath(), 内部已自动操作.
func writeDll(dll []byte) error {
	tmpDir := os.TempDir()
	tmpPath := filepath.Join(tmpDir, "WebView2Loader"+GetVersion()+"_"+runtime.GOARCH)

	err := os.Mkdir(tmpPath, 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}

	DllPath := filepath.Join(tmpPath, "WebView2Loader.dll")
	err = os.WriteFile(DllPath, dll, 0777)
	if err != nil {
		return err
	}

	dllPath = DllPath
	return nil
}

var (
	// 保证 loadDll 只运行一次.
	once = sync.Once{}

	// dll
	dllModule *syscall.LazyDLL

	// func
	procCreateCoreWebView2EnvironmentWithOptions                *syscall.LazyProc
	procCompareBrowserVersions                                  *syscall.LazyProc
	procGetAvailableCoreWebView2BrowserVersionString            *syscall.LazyProc
	procGetAvailableCoreWebView2BrowserVersionStringWithOptions *syscall.LazyProc
)

func loadDll() {
	once.Do(func() {
		dllModule = syscall.NewLazyDLL(dllPath)
		if dllModule.Handle() == 0 {
			panic("载入失败: " + dllPath)
		}

		procCreateCoreWebView2EnvironmentWithOptions = dllModule.NewProc("CreateCoreWebView2EnvironmentWithOptions")
		procCompareBrowserVersions = dllModule.NewProc("CompareBrowserVersions")
		procGetAvailableCoreWebView2BrowserVersionString = dllModule.NewProc("GetAvailableCoreWebView2BrowserVersionString")
		procGetAvailableCoreWebView2BrowserVersionStringWithOptions = dllModule.NewProc("GetAvailableCoreWebView2BrowserVersionStringWithOptions")
	})
}

// CompareBrowserVersions 将比较两个给定的版本号并返回如下结果:
//
//	-1 = v1 < v2
//	 0 = v1 == v2
//	 1 = v1 > v2
func CompareBrowserVersions(v1 string, v2 string) (int, error) {
	var result int
	r, _, err := procCompareBrowserVersions.Call(
		common.StrPtr(v1),
		common.StrPtr(v2),
		uintptr(unsafe.Pointer(&result)))
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return result, err
	}
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
	r, _, err := procGetAvailableCoreWebView2BrowserVersionString.Call(
		common.StrPtr(_browserExecutableFolder),
		uintptr(unsafe.Pointer(&result)))
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		// HRESULT 的低16位（错误代码本身）是 ERROR_FILE_NOT_FOUND，这意味着无可匹配版本
		if r&0xFFFF == uintptr(windows.ERROR_FILE_NOT_FOUND) {
			return "", nil // 返回一个空字符串，但不返回错误，因为我们成功检测到没有可匹配版本
		}
		return "", syscall.Errno(r)
	}
	version := windows.UTF16PtrToString(result)
	windows.CoTaskMemFree(unsafe.Pointer(result))
	return version, nil
}

// GetAvailableBrowserVersionWithOptions 获取本机安装或指定目录的可用的 webview2 运行时的版本号。没有可用的则返回空字符串。
//
// browserExecutableFolder: webview2 可执行文件的文件夹路径, 为空则获取本机安装的。
//
// environmentOptions: 包含 WebView2 环境选项的 ICoreWebView2EnvironmentOptions 对象指针。
func GetAvailableBrowserVersionWithOptions(browserExecutableFolder string, environmentOptions uintptr) (string, error) {
	var result *uint16
	r, _, err := procGetAvailableCoreWebView2BrowserVersionStringWithOptions.Call(
		common.StrPtr(browserExecutableFolder),
		environmentOptions,
		uintptr(unsafe.Pointer(&result)))
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		// HRESULT 的低16位（错误代码本身）是 ERROR_FILE_NOT_FOUND，这意味着无可匹配版本
		if r&0xFFFF == uintptr(windows.ERROR_FILE_NOT_FOUND) {
			return "", nil // 返回一个空字符串，但不返回错误，因为我们成功检测到没有可匹配版本
		}
		return "", syscall.Errno(r)
	}
	version := windows.UTF16PtrToString(result)
	windows.CoTaskMemFree(unsafe.Pointer(result))
	return version, nil
}

// CreateCoreWebView2EnvironmentWithOptions 使用指定选项创建一个 WebView2 环境。
//
// browserExecutableFolder: 浏览器可执行文件的文件夹路径, 可为空字符串。
//
// userDataFolder: 用户数据文件夹路径。
//
// environmentOptions: WebView2 环境选项的 ICoreWebView2EnvironmentOptions 对象指针。
//
// environmentCompletedHandler: 一个回调指针，当 WebView2 环境创建完成时调用。
func CreateCoreWebView2EnvironmentWithOptions(browserExecutableFolder, userDataFolder string, environmentOptions uintptr, environmentCompletedHandler uintptr) (uintptr, error) {
	result, _, _ := procCreateCoreWebView2EnvironmentWithOptions.Call(
		common.StrPtr(browserExecutableFolder),
		common.StrPtr(userDataFolder),
		environmentOptions,
		environmentCompletedHandler,
	)
	return result, nil
}
