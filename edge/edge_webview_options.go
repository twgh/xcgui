package edge

// WebView 选项函数类型
type WebViewOption func(*WebViewOptions)

// WebViewOptions 选项.
type WebViewOptions struct {
	// WebView 的默认背景颜色。
	//   - 此属性允许用户提前初始化 DefaultBackgroundColor，从而防止在 WebView2 加载期间，当背景色设置为白色以外的颜色时可能出现的白色闪烁。
	//   - 通过早期初始化，颜色从一开始就保持一致。
	//   - DefaultBackgroundColor 是在所有网页内容下方渲染的颜色。这意味着当没有加载网页内容时，WebView2 会渲染此颜色。如果 WebView2 中未定义背景色，它会使用 DefaultBackgroundColor 属性来渲染背景。默认情况下，此颜色设置为白色。
	//   - 此 API 仅支持不透明颜色和完全透明。对于 alpha 值不等于 0 或 255 的颜色，它会失效。当 WebView2 设置为完全透明时，它不会渲染背景，从而使其后部窗口的内容可见。
	//   - ≥ 137.0.3296.44
	DefaultBackgroundColor *COREWEBVIEW2_COLOR

	// 脚本区域设置, 如: zh-CN。
	//   - 它为所有 Intl JavaScript API 以及依赖于它的其他 JavaScript API 设置默认区域设置，即 Intl.DateTimeFormat()，这会影响字符串格式，例如时间/日期格式。示例：Intl.DateTimeFormat().format(new Date()) 预期的区域设置值采用 BCP 47 语言标记格式。
	//   - 对此属性的更改适用于更改后创建的渲染器进程。任何现有的渲染器进程将继续使用之前的 ScriptLocale 值。要确保更改应用于所有渲染器进程，请关闭并重新启动 ICoreWebView2Environment 以及所有关联的 WebView2 对象。
	//   - ScriptLocale 的默认值取决于 WebView2 语言和操作系统区域。如果 WebView2 语言和操作系统区域的语言部分匹配，则会使用操作系统区域。否则，将使用 WebView2 语言。
	//   - 你可以将 ScriptLocale 设置为空字符串以获取默认的 ScriptLocale 值。
	//   - ≥ 111.0.1661.34
	ScriptLocale string

	// ProfileName 配置文件名称。
	//   - 它的最大长度为64个字符（不包括空字符终止符）。它不区分 ASCII 大小写。
	//   - 注意：文本不能以句号“.”或空格“ ”结尾。
	//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
	ProfileName string
	// WebView 宿主窗口标题
	Title string
	// WebView 宿主窗口类名
	ClassName string

	// 资源文件中的图标资源 id, 给原生窗口设置图标, 如果为 0, 则使用默认图标.
	IconId uint

	// 原生窗口的圆角半径, 单位 px.
	//   - 需注意炫彩窗口的圆角需通过 SetShadowInfo 设置.
	RoundRadius int32

	WebViewSize

	// 填充父, 如果为 true, 则 WebView 会填充父, WebViewSize 里的固定坐标和尺寸会失效.
	FillParent bool

	// Debug 是否可打开开发者工具.
	Debug bool
	// AppDrag 是否启用非客户区域支持。
	//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
	//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 WebView 及其宿主应用程序窗口；
	//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
	//   - ≥ 123.0.2420.47
	AppDrag bool

	// AutoFocus 将在窗口获得焦点时尝试保持 WebView 的焦点。
	AutoFocus bool
	// PrivateMode 是否启用隐私模式。
	PrivateMode bool

	DefaultEnabledWebViewOption
}

// WebViewSize 是 WebView 的固定位置与大小.
type WebViewSize struct {
	Left   int32
	Top    int32
	Width  int32
	Height int32
}

// DefaultEnabledWebViewOption 里的 WebView 选项都是默认会开启的.
type DefaultEnabledWebViewOption struct {
	// 默认的上下文菜单, 默认为 true.
	DefaultContextMenus bool
	// 状态栏, 默认为 true.
	StatusBar bool
	// 缩放控件, 控制是否可以缩放, 默认为 true.
	ZoomControl bool
	// 浏览器快捷键, 是浏览器里面默认的一些快捷键, 组合键, 默认为 true.
	BrowserAcceleratorKeys bool
}

// WebView 默认选项
func defaultWebViewOptions() *WebViewOptions {
	return &WebViewOptions{
		ScriptLocale: "默认不设置",
		DefaultEnabledWebViewOption: DefaultEnabledWebViewOption{
			DefaultContextMenus:    true,
			StatusBar:              true,
			ZoomControl:            true,
			BrowserAcceleratorKeys: true,
		},
	}
}

// WithProfileName 设置配置文件名称.
//   - 它的最大长度为64个字符（不包括空字符终止符）。它不区分 ASCII 大小写。
//   - 注意：文本不能以句号"."或空格" "结尾。
//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
func WithProfileName(name string) WebViewOption {
	return func(o *WebViewOptions) {
		o.ProfileName = name
	}
}

// WithRoundRadius 设置原生窗口的圆角半径, 单位 px.
//   - 需注意炫彩窗口的圆角需通过 SetShadowInfo 设置.
func WithRoundRadius(radius int32) WebViewOption {
	return func(o *WebViewOptions) {
		o.RoundRadius = radius
	}
}

// WithScriptLocale 设置脚本区域设置, 如: zh-CN。
//   - 它为所有 Intl JavaScript API 以及依赖于它的其他 JavaScript API 设置默认区域设置，即 Intl.DateTimeFormat()，这会影响字符串格式，例如时间/日期格式。示例：Intl.DateTimeFormat().format(new Date()) 预期的区域设置值采用 BCP 47 语言标记格式。
//   - 对此属性的更改适用于更改后创建的渲染器进程。任何现有的渲染器进程将继续使用之前的 ScriptLocale 值。要确保更改应用于所有渲染器进程，请关闭并重新启动 ICoreWebView2Environment 以及所有关联的 WebView2 对象。
//   - ScriptLocale 的默认值取决于 WebView2 语言和操作系统区域。如果 WebView2 语言和操作系统区域的语言部分匹配，则会使用操作系统区域。否则，将使用 WebView2 语言。
//   - 你可以将 ScriptLocale 设置为空字符串以获取默认的 ScriptLocale 值。
//   - ≥ 111.0.1661.34
func WithScriptLocale(locale string) WebViewOption {
	return func(o *WebViewOptions) {
		o.ScriptLocale = locale
	}
}

// WithDefaultBackgroundColor 设置 WebView 的默认背景颜色.
//   - 此属性允许用户提前初始化 DefaultBackgroundColor，从而防止在 WebView2 加载期间，当背景色设置为白色以外的颜色时可能出现的白色闪烁。
//   - 通过早期初始化，颜色从一开始就保持一致。
//   - DefaultBackgroundColor 是在所有网页内容下方渲染的颜色。这意味着当没有加载网页内容时，WebView2 会渲染此颜色。如果 WebView2 中未定义背景色，它会使用 DefaultBackgroundColor 属性来渲染背景。默认情况下，此颜色设置为白色。
//   - 此 API 仅支持不透明颜色和完全透明。对于 alpha 值不等于 0 或 255 的颜色，它会失效。当 WebView2 设置为完全透明时，它不会渲染背景，从而使其后部窗口的内容可见。
//   - ≥ 137.0.3296.44
func WithDefaultBackgroundColor(color *COREWEBVIEW2_COLOR) WebViewOption {
	return func(o *WebViewOptions) {
		o.DefaultBackgroundColor = color
	}
}

// WithTitle 设置 WebView 宿主窗口标题
func WithTitle(title string) WebViewOption {
	return func(o *WebViewOptions) {
		o.Title = title
	}
}

// WithClassName 设置 WebView 宿主窗口类名
func WithClassName(className string) WebViewOption {
	return func(o *WebViewOptions) {
		o.ClassName = className
	}
}

// WithIconId 设置资源文件中的图标资源id, 给原生窗口设置图标, 如果为0, 则使用默认图标.
func WithIconId(id uint) WebViewOption {
	return func(o *WebViewOptions) {
		o.IconId = id
	}
}

// WithWebViewSize 设置 WebView 固定位置与大小.
func WithWebViewSize(left, top, width, height int32) WebViewOption {
	return func(o *WebViewOptions) {
		o.WebViewSize = WebViewSize{
			Left:   left,
			Top:    top,
			Width:  width,
			Height: height,
		}
	}
}

// WithFillParent 设置是否填充父, 如果为 true, 则 WebView 会填充父,  WebViewSize 里的固定坐标和尺寸会失效.
func WithFillParent(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.FillParent = enable
	}
}

// WithDebug 设置是否可打开开发者工具.
func WithDebug(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.Debug = enable
	}
}

// WithAppDrag 设置是否启用非客户区域支持。
//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 WebView 及其宿主应用程序窗口；
//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
//   - ≥ 123.0.2420.47
func WithAppDrag(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.AppDrag = enable
	}
}

// WithAutoFocus 设置是否在窗口获得焦点时尝试保持 WebView 的焦点.
func WithAutoFocus(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.AutoFocus = enable
	}
}

// WithPrivateMode 设置是否启用隐私模式.
func WithPrivateMode(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.PrivateMode = enable
	}
}

// WithDefaultContextMenus 设置是否启用默认的上下文菜单, 默认为 true.
func WithDefaultContextMenus(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.DefaultContextMenus = enable
	}
}

// WithStatusBar 设置是否启用状态栏, 默认为 true.
func WithStatusBar(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.StatusBar = enable
	}
}

// WithZoomControl 设置是否启用缩放控件, 控制是否可以缩放, 默认为 true.
func WithZoomControl(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.ZoomControl = enable
	}
}

// WithBrowserAcceleratorKeys 设置是否启用浏览器快捷键, 是浏览器里面默认的一些快捷键, 组合键, 默认为 true.
func WithBrowserAcceleratorKeys(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.BrowserAcceleratorKeys = enable
	}
}
