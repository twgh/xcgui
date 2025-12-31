package edge

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"syscall"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xcc"
)

// GetVersion 返回本库使用的 WebView2 运行时版本号。更高版本也可用, 是兼容低版本的。
func GetVersion() string {
	return "143.0.3650.58"
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
	// 浏览器可执行文件的文件夹路径, 为空则使用本机安装的 WebView2 运行时。
	BrowserExecutableFolder string

	// 用户数据文件夹路径。为空则使用 AppData 环境变量下的当前进程名文件夹。
	UserDataFolder string

	// WebView2 环境的配置选项。可为 nil.
	//   - 内部会调用 edge.CreateEnvironmentOptions().
	//   - 与 EnvironmentOptions 二选一使用, 两种方法区别是一个手动一个自动.
	//   - 如果 EnvironmentOptions 不为 nil，将忽略本字段。
	EnvOptions *EnvOptions

	// 创建 WebView2 环境的配置选项。可为 nil.
	//   - 使用 edge.CreateEnvironmentOptions() 创建.
	//   - 与 EnvOptions 二选一使用, 两种方法区别是一个手动一个自动.
	EnvironmentOptions *ICoreWebView2EnvironmentOptions

	// 环境创建完成回调函数
	EnvCompletedCallback func(result syscall.Errno, env *ICoreWebView2Environment)
}

// EnvOptions WebView2 环境选项简化配置.
type EnvOptions struct {
	// 创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
	//   - 如: --autoplay-policy=no-user-gesture-required
	AdditionalBrowserArguments []string

	// WebView2 环境的语言。如: en-us, zh-cn
	//   - 它适用于浏览器用户界面，如上下文菜单和对话框。
	//   - 它也适用于 WebView 向网站发送的 accept - languages HTTP 头部。
	//   - 预期的语言环境值采用 BCP 47 语言标签的格式。更多信息可从 IETF BCP47 中获取。
	Language string

	// 目标兼容的浏览器版本。
	TargetCompatibleBrowserVersion string

	// 是否允许 Azure Active Directory (AAD) 和个人 Microsoft Account (MSA) 资源的单点登录。
	//   - 所有与 Windows 关联并供所有应用程序共享的 AAD 账户均受支持。
	//   - 对于 MSA，单点登录 (SSO) 仅对用于 Windows 账户登录的账户启用（如果存在）。
	//   - 默认情况下为禁用状态。
	//   - 通用 Windows 平台应用还必须声明 enterpriseCloudSSORestricted 功能，单点登录 (SSO) 才能正常工作。
	AllowSingleSignOnUsingOSPrimaryAccount bool

	// 其他进程是否可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例。默认为 false。
	ExclusiveUserDataFolderAccess bool

	// Windows 是否会将崩溃数据发送到 Microsoft 端点。默认为 false。
	IsCustomCrashReportingEnabled bool

	// 自定义方案注册列表。
	CustomSchemeRegistrations []*ICoreWebView2CustomSchemeRegistration

	// 是否禁用 WebView2 中的跟踪防护功能。默认为 false。
	//   - 此属性可禁用在同一环境中创建的所有 WebView2 的跟踪防护功能。默认情况下，该功能处于启用状态，用于阻止潜在有害的跟踪器以及来自从未访问过的网站的跟踪器，并设置为 COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED 或配置文件上次更改/保存的任何值。
	//   - 如果应用程序仅在 WebView2 中渲染已知安全的内容，则可以将此属性设置为 true 以禁用跟踪防护功能。在创建环境时禁用此功能还可以通过跳过相关代码来提高运行时性能。
	//   - 如果将 WebView2 用作具有任意导航功能的“完整浏览器”，则不应禁用，并且应保护最终用户的隐私。
	DisableTrackingPrevention bool

	// 是否启用浏览器扩展功能。默认为 false。
	AreBrowserExtensionsEnabled bool

	// 频道搜索类型。
	//   - 默认值为 COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE.
	//   - 环境创建会按稳定性从高到低的顺序在计算机上搜索发布通道，并使用找到的第一个通道。
	//   - 默认搜索顺序为：WebView2 运行时 -> Beta -> Dev -> Canary。
	ChannelSearchKind COREWEBVIEW2_CHANNEL_SEARCH_KIND

	// 发布频道。
	//   - 指示环境创建应搜索哪些频道。
	//   - 默认值是所有通道的组合.
	ReleaseChannels *ReleaseChannels

	// 滚动条样式。
	ScrollBarStyle COREWEBVIEW2_SCROLLBAR_STYLE
}

type ReleaseChannels struct {
	// 发布频道。
	//   - 一个或多个 COREWEBVIEW2_RELEASE_CHANNELS 的组合，指示环境创建应搜索哪些频道。
	//   - 默认值是所有通道的组合.
	ReleaseChannels COREWEBVIEW2_RELEASE_CHANNELS
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

// createEnvOptions 根据 EnvOptions 创建 ICoreWebView2EnvironmentOptions.
func createEnvOptions(opts *EnvOptions) (*ICoreWebView2EnvironmentOptions, error) {
	if opts == nil {
		return nil, errors.New("EnvOptions is nil")
	}

	// 创建环境选项
	envOpts, err := CreateEnvironmentOptions()
	if err != nil {
		return nil, err
	}

	// 设置 Options
	if len(opts.AdditionalBrowserArguments) > 0 {
		if err = envOpts.SetAdditionalBrowserArguments(strings.Join(opts.AdditionalBrowserArguments, " ")); err != nil {
			return nil, err
		}
	}
	if opts.Language != "" {
		if err = envOpts.SetLanguage(opts.Language); err != nil {
			return nil, err
		}
	}
	if opts.TargetCompatibleBrowserVersion != "" {
		if err = envOpts.SetTargetCompatibleBrowserVersion(opts.TargetCompatibleBrowserVersion); err != nil {
			return nil, err
		}
	}
	if opts.AllowSingleSignOnUsingOSPrimaryAccount {
		if err = envOpts.SetAllowSingleSignOnUsingOSPrimaryAccount(opts.AllowSingleSignOnUsingOSPrimaryAccount); err != nil {
			return nil, err
		}
	}

	// 设置 Options2
	if opts.ExclusiveUserDataFolderAccess {
		envOpts2, err := envOpts.GetICoreWebView2EnvironmentOptions2()
		if err != nil {
			return nil, err
		}
		defer envOpts2.Release()

		if err = envOpts2.SetExclusiveUserDataFolderAccess(opts.ExclusiveUserDataFolderAccess); err != nil {
			return nil, err
		}
	}

	// 设置 Options3
	if opts.IsCustomCrashReportingEnabled {
		envOpts3, err := envOpts.GetICoreWebView2EnvironmentOptions3()
		if err != nil {
			return nil, err
		}
		defer envOpts3.Release()

		if err = envOpts3.SetIsCustomCrashReportingEnabled(opts.IsCustomCrashReportingEnabled); err != nil {
			return nil, err
		}
	}

	// 设置 Options4
	if len(opts.CustomSchemeRegistrations) > 0 {
		envOpts4, err := envOpts.GetICoreWebView2EnvironmentOptions4()
		if err != nil {
			return nil, err
		}
		defer envOpts4.Release()

		if err = envOpts4.SetCustomSchemeRegistrations(opts.CustomSchemeRegistrations); err != nil {
			return nil, err
		}
	}

	// 设置 Options5
	if opts.DisableTrackingPrevention {
		envOpts5, err := envOpts.GetICoreWebView2EnvironmentOptions5()
		if err != nil {
			return nil, err
		}
		defer envOpts5.Release()

		if err = envOpts5.SetEnableTrackingPrevention(false); err != nil {
			return nil, err
		}
	}

	// 设置 Options6
	if opts.AreBrowserExtensionsEnabled {
		envOpts6, err := envOpts.GetICoreWebView2EnvironmentOptions6()
		if err != nil {
			return nil, err
		}
		defer envOpts6.Release()

		if err = envOpts6.SetAreBrowserExtensionsEnabled(opts.AreBrowserExtensionsEnabled); err != nil {
			return nil, err
		}
	}

	// 设置 Options7
	if opts.ChannelSearchKind > 0 || opts.ReleaseChannels != nil {
		envOpts7, err := envOpts.GetICoreWebView2EnvironmentOptions7()
		if err != nil {
			return nil, err
		}
		defer envOpts7.Release()

		if opts.ChannelSearchKind > 0 {
			if err = envOpts7.SetChannelSearchKind(opts.ChannelSearchKind); err != nil {
				return nil, err
			}
		}

		if opts.ReleaseChannels != nil {
			if err = envOpts7.SetReleaseChannels(opts.ReleaseChannels.ReleaseChannels); err != nil {
				return nil, err
			}
		}
	}

	// 设置 Options8
	if opts.ScrollBarStyle > 0 {
		envOpts8, err := envOpts.GetICoreWebView2EnvironmentOptions8()
		if err != nil {
			return nil, err
		}
		defer envOpts8.Release()

		if err = envOpts8.SetScrollBarStyle(opts.ScrollBarStyle); err != nil {
			return nil, err
		}
	}

	return envOpts, nil
}

// New Edge 在整个应用程序的生命周期里应该只创建一次.
//
// opt: Edge选项.
func New(opt Option) (*Edge, error) {
	if atomic.LoadUintptr(&envInited) != 0 {
		return nil, errors.New("edge environment already inited")
	}

	var err error
	// 处理环境选项
	envOpts := opt.EnvironmentOptions
	if envOpts == nil && opt.EnvOptions != nil {
		// 根据 EnvOptions 创建环境选项
		envOpts, err = createEnvOptions(opt.EnvOptions)
		if err != nil {
			return nil, errors.New("create environment options failed: " + err.Error())
		}
		defer envOpts.Release()
	}

	e := &Edge{}
	e.cbEnvCompleted = opt.EnvCompletedCallback
	// 环境初始化完成事件
	e.handlerCreateCoreWebView2EnvironmentCompleted = NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(e)
	// 控制器创建完成事件
	e.handlerCreateCoreWebView2ControllerCompleted = NewICoreWebView2CreateCoreWebView2ControllerCompletedHandler(e)

	// 处理用户数据文件夹路径, 如果为空, 则使用 AppData 环境变量下的当前进程名文件夹
	dataPath := opt.UserDataFolder
	if dataPath == "" {
		currentExeName := common.GetProcessNameWithoutExt()
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
	err2 := CreateCoreWebView2EnvironmentWithOptions(opt.BrowserExecutableFolder, dataPath, envOpts, e.handlerCreateCoreWebView2EnvironmentCompleted)
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
