package edge

import (
	"errors"
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xcc"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"path/filepath"
	"sync/atomic"
)

// GetVersion 返回本库使用的 WebView2 运行时版本号。更高版本也可用, 是兼容低版本的。
func GetVersion() string {
	return "136.0.3240.44"
}

// WebView2DownloadLink 是 WebView2 运行时小型引导程序下载地址.
//   - 引导程序是一个小型安装程序，用于下载常青运行时匹配设备体系结构并将其安装在本地。
//   - 建议引导用户使用管理员权限打开下载后的安装包。
//   - 在 Win7 系统使用该引导程序会安装 109.0.1518.140 版本的 WebView2 运行时, 这是最后支持 Win7 系统的版本.
const WebView2DownloadLink = "https://go.microsoft.com/fwlink/p/?LinkId=2124703"

// DownloadWebView2 使用默认浏览器打开 WebView2 运行时小型引导程序下载地址.
//   - 引导程序是一个小型安装程序，用于下载常青运行时匹配设备体系结构并将其安装在本地。
//   - 建议引导用户使用管理员权限打开下载后的安装包。
//   - 在 Win7 系统使用该引导程序会安装 109.0.1518.140 版本的 WebView2 运行时, 这是最后支持 Win7 系统的版本.
func DownloadWebView2() {
	wapi.ShellExecuteW(0, "open", WebView2DownloadLink, "", "", xcc.SW_SHOWNORMAL)
}

func init() {
	hr := wapi.CoInitializeEx(0, wapi.COINIT_APARTMENTTHREADED)
	if hr < 0 {
		log.Printf("Warning: CoInitializeEx call failed: 0x%08x", hr)
	}
}

var (
	// WebView2 环境是否已初始化
	envInited uintptr
)

// Option Edge选项.
type Option struct {
	// 浏览器可执行文件的文件夹路径, 为空则使用本机安装的 webview2 运行时。
	BrowserExecutableFolder string

	// 用户数据文件夹路径。为空则使用当前可执行文件所在目录下的 AppData 文件夹。
	UserDataFolder string

	// 创建 WebView2 环境的配置选项。为 nil 则使用默认配置。
	EnvironmentOptions *ICoreWebView2EnvironmentOptions

	// 环境创建完成回调函数
	EnvCompletedCallback func(result uintptr, env *ICoreWebView2Environment)
}

type Edge struct {
	// WebView2 环境
	Environment *ICoreWebView2Environment

	// 环境初始化完成事件处理程序
	handlerCreateCoreWebView2EnvironmentCompleted *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler
	cbEnvCompleted                                func(result uintptr, env *ICoreWebView2Environment)
	// 控制器创建完成事件处理程序
	handler_CreateCoreWebView2ControllerCompleted *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler
	cbCreateCoreWebView2ControllerCompleted       func(result uintptr, controller *ICoreWebView2Controller)

	// WebView2 环境创建完成回调 [内部使用]
	_cbEnvCompleted func(result uintptr, env *ICoreWebView2Environment)
	// WebView2 控制器创建完成回调 [内部使用]
	_cbCreateCoreWebView2ControllerCompleted func(result uintptr, controller *ICoreWebView2Controller)
}

// New Edge 在整个应用程序的生命周期里应该只创建一次.
//
// opt: Edge选项.
func New(opt Option) (*Edge, error) {
	e := &Edge{}
	e.cbEnvCompleted = opt.EnvCompletedCallback

	var err error
	if atomic.LoadUintptr(&envInited) == 0 {
		// 环境初始化完成事件
		e.handlerCreateCoreWebView2EnvironmentCompleted = NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(e)
		// 控制器创建完成事件
		e.handler_CreateCoreWebView2ControllerCompleted = NewICoreWebView2CreateCoreWebView2ControllerCompletedHandler(e)

		// 处理用户数据文件夹路径, 如果为空, 则使用当前可执行文件所在目录下的 AppData 文件夹
		dataPath := opt.UserDataFolder
		if dataPath == "" {
			currentExePath := make([]uint16, windows.MAX_PATH)
			_, err := windows.GetModuleFileName(windows.Handle(0), &currentExePath[0], windows.MAX_PATH)
			if err != nil {
				return nil, errors.New("error calling GetModuleFileName: " + err.Error())
			}
			currentExeName := filepath.Base(windows.UTF16ToString(currentExePath))
			dataPath = filepath.Join(os.Getenv("AppData"), currentExeName)
		}

		e._cbEnvCompleted = func(result uintptr, env *ICoreWebView2Environment) {
			if result == 0 { // 成功
				env.AddRef()
				e.Environment = env
				atomic.StoreUintptr(&envInited, 1)
			} else {
				err = fmt.Errorf("CreateCoreWebView2EnvironmentWithOptions failed: 0x%08x", result)
			}
		}

		// 创建 WebView2 环境
		result, err := CreateCoreWebView2EnvironmentWithOptions(opt.BrowserExecutableFolder, dataPath, opt.EnvironmentOptions, e.handlerCreateCoreWebView2EnvironmentCompleted)
		if err != nil {
			return nil, errors.New("error calling CreateCoreWebView2EnvironmentWithOptions: " + err.Error())
		} else if result != 0 {
			return nil, fmt.Errorf("calling CreateCoreWebView2EnvironmentWithOptions return: 0x%08x", result)
		}

		// 等待 webview2 环境创建完成
		var msg wapi.MSG
		for {
			if atomic.LoadUintptr(&envInited) != 0 {
				break
			}
			if wapi.GetMessage(&msg, 0, 0, 0) == 0 {
				break
			}
			wapi.TranslateMessage(&msg)
			wapi.DispatchMessage(&msg)
		}
	}
	return e, err
}

func (e *Edge) QueryInterface(refiid, object uintptr) uintptr {
	err := e.Environment.QueryInterface(refiid, object)
	errno, _ := ErrorToErrno(err)
	return uintptr(errno)
}

func (e *Edge) AddRef() uintptr {
	return 1
}

func (e *Edge) Release() uintptr {
	return 1
}

// --------------------------- 事件 ---------------------------

// EnvironmentCompleted 环境创建完成.
func (e *Edge) EnvironmentCompleted(result uintptr, env *ICoreWebView2Environment) uintptr {
	if e._cbEnvCompleted != nil {
		e._cbEnvCompleted(result, env)
	}
	if e.cbEnvCompleted != nil {
		e.cbEnvCompleted(result, env)
	}
	return 0
}

// CreateCoreWebView2ControllerCompleted WebView2 控制器创建完成.
func (e *Edge) CreateCoreWebView2ControllerCompleted(result uintptr, controller *ICoreWebView2Controller) uintptr {
	if e._cbCreateCoreWebView2ControllerCompleted != nil {
		e._cbCreateCoreWebView2ControllerCompleted(result, controller)
	}
	if e.cbCreateCoreWebView2ControllerCompleted != nil {
		e.cbCreateCoreWebView2ControllerCompleted(result, controller)
	}
	return 0
}

// Event_CreateCoreWebView2ControllerCompleted 是 WebView2 控制器创建完成事件. 需用在 NewWebView 之前.
func (e *Edge) Event_CreateCoreWebView2ControllerCompleted(cb func(result uintptr, controller *ICoreWebView2Controller)) {
	e.cbCreateCoreWebView2ControllerCompleted = cb
}
