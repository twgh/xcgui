package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xcc"
	"os"
	"path/filepath"
	"sync/atomic"
	"syscall"
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

var (
	// WebView2 环境是否已初始化
	envInited uintptr
)

// Option Edge选项.
type Option struct {
	// 浏览器可执行文件的文件夹路径, 为空则使用本机安装的 webview2 运行时。
	BrowserExecutableFolder string

	// 用户数据文件夹路径。为空则使用 AppData 环境变量下的当前进程名文件夹。
	UserDataFolder string

	// 创建 WebView2 环境的配置选项。为 nil 则使用默认配置。
	EnvironmentOptions *ICoreWebView2EnvironmentOptions

	// 环境创建完成回调函数
	EnvCompletedCallback func(result syscall.Errno, env *ICoreWebView2Environment)
}

// Edge WebView2 环境.
type Edge struct {
	IUnknown_Impl
	// WebView2 环境
	Environment *ICoreWebView2Environment

	// 环境初始化完成事件处理程序
	handlerCreateCoreWebView2EnvironmentCompleted *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler
	cbEnvCompleted                                func(result syscall.Errno, env *ICoreWebView2Environment)
	// 控制器创建完成事件处理程序
	handlerCreateCoreWebView2ControllerCompleted *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler
	cbCreateCoreWebView2ControllerCompleted      func(result syscall.Errno, controller *ICoreWebView2Controller)

	// WebView2 环境创建完成回调 [内部使用]
	_cbEnvCompleted func(result syscall.Errno, env *ICoreWebView2Environment)
	// WebView2 控制器创建完成回调 [内部使用]
	_cbCreateCoreWebView2ControllerCompleted func(result syscall.Errno, controller *ICoreWebView2Controller)
}

// New Edge 在整个应用程序的生命周期里应该只创建一次.
//
// opt: Edge选项.
func New(opt Option) (*Edge, error) {
	if atomic.LoadUintptr(&envInited) != 0 {
		return nil, errors.New("edge environment already inited")
	}

	var err error
	e := &Edge{}
	e.cbEnvCompleted = opt.EnvCompletedCallback
	// 环境初始化完成事件
	e.handlerCreateCoreWebView2EnvironmentCompleted = NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(e)
	// 控制器创建完成事件
	e.handlerCreateCoreWebView2ControllerCompleted = NewICoreWebView2CreateCoreWebView2ControllerCompletedHandler(e)

	// 处理用户数据文件夹路径, 如果为空, 则使用 AppData 环境变量下的当前进程名文件夹
	dataPath := opt.UserDataFolder
	if dataPath == "" {
		currentExeName := common.GetProcessNameNoExt()
		if currentExeName == "" {
			currentExeName = "xcgui_webview"
		}
		dataPath = filepath.Join(os.Getenv("AppData"), currentExeName)
	}

	e._cbEnvCompleted = func(result syscall.Errno, env *ICoreWebView2Environment) {
		if errors.Is(result, wapi.S_OK) { // 成功
			env.AddRef()
			e.Environment = env
			atomic.StoreUintptr(&envInited, 1)
		} else {
			err = errors.New("CreateCoreWebView2EnvironmentWithOptions failed: " + result.Error())
		}
	}

	// 创建 WebView2 环境
	err2 := CreateCoreWebView2EnvironmentWithOptions(opt.BrowserExecutableFolder, dataPath, opt.EnvironmentOptions, e.handlerCreateCoreWebView2EnvironmentCompleted)
	if err2 != nil {
		return nil, errors.New("calling CreateCoreWebView2EnvironmentWithOptions failed: " + err2.Error())
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
	return e, err
}

// --------------------------- 事件接口实现 ---------------------------

// EnvironmentCompleted 环境创建完成.
func (e *Edge) EnvironmentCompleted(result syscall.Errno, env *ICoreWebView2Environment) uintptr {
	if e._cbEnvCompleted != nil {
		e._cbEnvCompleted(result, env)
	}
	if e.cbEnvCompleted != nil {
		e.cbEnvCompleted(result, env)
	}
	return 0
}

// CreateCoreWebView2ControllerCompleted WebView2 控制器创建完成.
func (e *Edge) CreateCoreWebView2ControllerCompleted(result syscall.Errno, controller *ICoreWebView2Controller) uintptr {
	if e._cbCreateCoreWebView2ControllerCompleted != nil {
		e._cbCreateCoreWebView2ControllerCompleted(result, controller)
	}
	if e.cbCreateCoreWebView2ControllerCompleted != nil {
		e.cbCreateCoreWebView2ControllerCompleted(result, controller)
	}
	return 0
}

// --------------------------- 事件 ---------------------------

// Event_CreateCoreWebView2ControllerCompleted 是 WebView2 控制器创建完成事件. 需用在 NewWebView 之前.
func (e *Edge) Event_CreateCoreWebView2ControllerCompleted(cb func(result syscall.Errno, controller *ICoreWebView2Controller)) {
	e.cbCreateCoreWebView2ControllerCompleted = cb
}
